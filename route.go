package amixr

import (
	"fmt"
	"log"
	"net/http"
)

// Handles requests to route endpoint
// Use NewRouteService instead of direct creation RouteService
//
// http://api-docs.amixr.io/#routes
type RouteService struct {
	client *Client
	url    string
}

// NewRoutesService creates RouteService with defined url
func NewRouteService(client *Client) *RouteService {
	routeService := RouteService{}
	routeService.client = client
	routeService.url = "routes"
	return &routeService
}

type PaginatedRoutesResponse struct {
	PaginatedResponse
	Routes []*Route `json:"results"`
}

type Route struct {
	ID             string      `json:"id"`
	IntegrationId  string      `json:"integration_id"`
	Position       int         `json:"position"`
	RoutingRegex   string      `json:"routing_regex"`
	IsTheLastRoute bool        `json:"is_the_last_route"`
	SlackRoute     *SlackRoute `json:"slack"`
}

type SlackRoute struct {
	ChannelId *string `json:"channel_id"`
}

type ListRouteOptions struct {
	ListOptions
	IntegrationId string `url:"integration_id,omitempty" json:"integration_id,omitempty"`
	RoutingRegex  string `url:"routing_regex,omitempty" json:"routing_regex,omitempty"`
}

// ListRoutes gets all routes for authorized team
//
// http://api-docs.amixr.io/#list-routes
func (service *RouteService) ListRoutes(opt *ListRouteOptions) (*PaginatedRoutesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var routes *PaginatedRoutesResponse
	resp, err := service.client.Do(req, &routes)
	if err != nil {
		return nil, resp, err
	}

	return routes, resp, err
}

type GetRouteOptions struct {
}

// Get route by given id
//
// http://api-docs.amixr.io/#get-route
func (service *RouteService) GetRoute(id string, opt *GetRouteOptions) (*Route, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	route := new(Route)
	resp, err := service.client.Do(req, route)
	if err != nil {
		return nil, resp, err
	}

	return route, resp, err
}

type CreateRouteOptions struct {
	IntegrationId string      `json:"integration_id,omitempty"`
	Position      *int        `json:"position,omitempty"`
	RoutingRegex  string      `json:"routing_regex, omitempty"`
	Slack         *SlackRoute `json:"slack,omitempty"`
	ManualOrder   bool        `url:"manual_order,omitempty" json:"manual_order,omitempty"`
}

// Create route with given name and type
//
// http://api-docs.amixr.io/#create-route
func (service *RouteService) CreateRoute(opt *CreateRouteOptions) (*Route, *http.Response, error) {
	log.Printf("[DEBUG] create amixr route")
	u := fmt.Sprintf("%s/", service.url)
	req, err := service.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	route := new(Route)

	resp, err := service.client.Do(req, route)
	log.Printf("[DEBUG] request success")

	if err != nil {
		return nil, resp, err
	}

	return route, resp, err
}

type UpdateRouteOptions struct {
	Position     *int        `json:"position,omitempty"`
	Slack        *SlackRoute `json:"slack,omitempty"`
	RoutingRegex string      `json:"routing_regex, omitempty"`
	ManualOrder  bool        `url:"manual_order,omitempty" json:"manual_order,omitempty"`
}

// Updates route with new templates and/or name. At least one field in template is required
//
// http://api-docs.amixr.io/#update-route
func (service *RouteService) UpdateRoute(id string, opt *UpdateRouteOptions) (*Route, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	route := new(Route)
	resp, err := service.client.Do(req, route)
	if err != nil {
		return nil, resp, err
	}

	return route, resp, err
}

type DeleteRouteOptions struct {
}

// Deletes route
//
// http://api-docs.amixr.io/#delete-route
func (service *RouteService) DeleteRoute(id string, opt *DeleteRouteOptions) (*http.Response, error) {

	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("DELETE", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := service.client.Do(req, nil)
	return resp, err
}
