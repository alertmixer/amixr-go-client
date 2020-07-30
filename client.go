package amixr

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	defaultBaseURL = "https://amixr.io/"
	apiVersionPath = "api/v1/"
)

type ListOptions struct {
	Page int `url:"page,omitempty" json:"page,omitempty"`
}

type PaginatedResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type Client struct {
	// HTTP client used to communicate with the API.
	client  *retryablehttp.Client
	token   string
	baseURL *url.URL
	// List of Services. Keep in sync with func newClient
	Integrations   *IntegrationService
	Escalations    *EscalationService
	Users          *UserService
	Schedules      *ScheduleService
	Routes         *RouteService
	SlackChannels  *SlackChannelService
	UserGroups     *UserGroupService
	CustomActions  *CustomActionService
	disableRetries bool
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("Token required")
	}
	client, err := newClient()
	if err != nil {
		return nil, err
	}
	// retries disabled for test/develop purposes. Normally should be enabled.
	client.disableRetries = true
	client.token = token
	return client, nil
}

func newClient() (*Client, error) {
	c := &Client{}

	// Configure the HTTP client.
	c.client = &retryablehttp.Client{
		CheckRetry: c.retryHTTPCheck,
		RetryMax:   1,
	}

	// Set the default base URL. _ suppress error handling
	_ = c.setBaseURL(defaultBaseURL + apiVersionPath)

	// Create services. Keep in sync with Client struct
	c.Integrations = NewIntegrationService(c)
	c.Escalations = NewEscalationService(c)
	c.Users = NewUserService(c)
	c.Schedules = NewScheduleService(c)
	c.Routes = NewRouteService(c)
	c.SlackChannels = NewSlackChannelService(c)
	c.UserGroups = NewUserGroupService(c)
	c.CustomActions = NewCustomActionService(c)

	return c, nil
}

func (c *Client) setBaseURL(urlStr string) error {

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	c.baseURL = baseURL

	return nil
}

func (c *Client) NewRequest(method, path string, opt interface{}) (*retryablehttp.Request, error) {
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Authorization", c.token)

	var body interface{}
	switch {
	case method == "POST" || method == "PUT":
		reqHeaders.Set("Content-Type", "application/json")

		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			}
		}
	case opt != nil:
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := retryablehttp.NewRequest(method, u.String(), body)

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *retryablehttp.Request, v interface{}) (*http.Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		errorResponse.Body = data

		// Very naive realization if handling errors messages
		var rawError interface{}
		if err := json.Unmarshal(data, &rawError); err != nil {
			errorResponse.Message = "failed to parse unknown error format"
		} else {
			errorResponse.Message = parseError(rawError)
		}
	}
	if err != nil {
		return err
	}

	return errorResponse
}

func parseError(raw interface{}) string {
	switch raw := raw.(type) {
	case string:
		return raw

	case []interface{}:
		var errs []string
		for _, v := range raw {
			errs = append(errs, parseError(v))
		}
		return fmt.Sprintf("[%s]", strings.Join(errs, ", "))

	case map[string]interface{}:
		var errs []string
		for k, v := range raw {
			errs = append(errs, fmt.Sprintf("{%s: %s}", k, parseError(v)))
		}
		sort.Strings(errs)
		return strings.Join(errs, ", ")

	default:
		return fmt.Sprintf("failed to parse unexpected error type: %T", raw)
	}
}

type ErrorResponse struct {
	Body     []byte
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.Message)
}

func (c *Client) retryHTTPCheck(ctx context.Context, resp *http.Response, err error) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	if err != nil {
		return false, err
	}
	if !c.disableRetries && (resp.StatusCode == 429 || resp.StatusCode >= 500) {
		return true, nil
	}
	return false, nil
}

func (c *Client) BaseURL() *url.URL {
	u := *c.baseURL
	return &u
}
