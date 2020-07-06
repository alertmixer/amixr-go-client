package amixr

import (
	"fmt"
	"log"
	"net/http"
)

// Handles requests to integration endpoint
// Use NewIntegrationService instead of direct creation IntegrationService
//
// http://api-docs.amixr.io/#integrations
type IntegrationService struct {
	client *Client
	url    string
}

// NewIntegrationsService creates IntegrationService with defined url
func NewIntegrationService(client *Client) *IntegrationService {
	integrationService := IntegrationService{}
	integrationService.client = client
	integrationService.url = "integrations"
	return &integrationService
}

type PaginatedIntegrationsResponse struct {
	PaginatedResponse
	Integrations []*Integration `json:"results"`
}

type Integration struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Link           string     `json:"link"`
	IncidentsCount int        `json:"incidents_count"`
	Type           string     `json:"type"`
	DefaultRouteId string     `json:"default_route_id"`
	Templates      *Templates `json:"templates"`
}

type Templates struct {
	GroupingKey   *string        `json:"grouping_key"`
	ResolveSignal *string        `json:"resolve_signal"`
	Slack         *SlackTemplate `json:"slack"`
}

type SlackTemplate struct {
	Title    *string `json:"title"`
	Message  *string `json:"message"`
	ImageURL *string `json:"image_url"`
}

type ListIntegrationOptions struct {
	ListOptions
}

// ListIntegrations gets all integrations for authorized team
//
// http://api-docs.amixr.io/#list-integrations
func (service *IntegrationService) ListIntegrations(opt *ListIntegrationOptions) (*PaginatedIntegrationsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var integrations *PaginatedIntegrationsResponse
	resp, err := service.client.Do(req, &integrations)
	if err != nil {
		return nil, resp, err
	}

	return integrations, resp, err
}

type GetIntegrationOptions struct {
}

// Get integration by given id
//
// http://api-docs.amixr.io/#get-integration
func (service *IntegrationService) GetIntegration(id string, opt *GetIntegrationOptions) (*Integration, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	integration := new(Integration)
	resp, err := service.client.Do(req, integration)
	if err != nil {
		return nil, resp, err
	}

	return integration, resp, err
}

type CreateIntegrationOptions struct {
	Name      string     `url:"name,omitempty" json:"name,omitempty"`
	Type      string     `url:"type,omitempty" json:"type,omitempty"`
	Templates *Templates `url:"type,omitempty" json:"templates,omitempty"`
}

// Create integration with given name and type
//
// http://api-docs.amixr.io/#create-integration
func (service *IntegrationService) CreateIntegration(opt *CreateIntegrationOptions) (*Integration, *http.Response, error) {
	log.Printf("[DEBUG] create amixr integration")
	u := fmt.Sprintf("%s/", service.url)

	req, err := service.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	integration := new(Integration)
	resp, err := service.client.Do(req, integration)
	if err != nil {
		return nil, resp, err
	}

	return integration, resp, err
}

type UpdateIntegrationOptions struct {
	Name      string     `json:"name, omitempty"`
	Templates *Templates `json:"templates,omitempty"`
}

// Updates integration with new templates and/or name. At least one field in template is required
//
// http://api-docs.amixr.io/#update-integration
func (service *IntegrationService) UpdateIntegration(id string, opt *UpdateIntegrationOptions) (*Integration, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	integration := new(Integration)
	resp, err := service.client.Do(req, integration)
	if err != nil {
		return nil, resp, err
	}

	return integration, resp, err
}

type DeleteIntegrationOptions struct {
}

// Deletes integration
//
// http://api-docs.amixr.io/#delete-integration
func (service *IntegrationService) DeleteIntegration(id string, opt *DeleteIntegrationOptions) (*http.Response, error) {

	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("DELETE", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := service.client.Do(req, nil)
	return resp, err
}
