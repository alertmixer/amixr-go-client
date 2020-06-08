package amixr

import (
	"fmt"
	"log"
	"net/http"
)

// Handles requests to escalation endpoint
// Use NewEscalationService instead of direct creation EscalationService
//
// http://api-docs.amixr.io/#escalations
type EscalationService struct {
	client *Client
	url    string
}

// NewEscalationsService creates EscalationService with defined url
func NewEscalationService(client *Client) *EscalationService {
	escalationService := EscalationService{}
	escalationService.client = client
	escalationService.url = "escalation_policies"
	return &escalationService
}

type PaginatedEscalationsResponse struct {
	PaginatedResponse
	Escalations []*Escalation `json:"results"`
}

type Escalation struct {
	ID                       string    `json:"id"`
	RouteId                  string    `json:"route_id"`
	Position                 int       `json:"position"`
	Type                     string    `json:"type"`
	Duration                 *int      `json:"duration"`
	PersonsToNotify          *[]string `json:"persons_to_notify"`
	PersonsToNotifyEachTime  *[]string `json:"persons_to_notify_next_each_time"`
	NotifyOnCallFromSchedule *string   `json:"notify_on_call_from_schedule"`
}

// Empty struct is here in case we want to add request params to ListEscalations.
type ListEscalationOptions struct {
	ListOptions
}

// ListEscalations gets all escalations for authorized team
//
// http://api-docs.amixr.io/#list-escalations
func (service *EscalationService) ListEscalations(opt *ListEscalationOptions) (*PaginatedEscalationsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var escalations *PaginatedEscalationsResponse
	resp, err := service.client.Do(req, &escalations)
	if err != nil {
		return nil, resp, err
	}

	return escalations, resp, err
}

type GetEscalationOptions struct {
}

// Get escalation by given id
//
// http://api-docs.amixr.io/#get-escalation
func (service *EscalationService) GetEscalation(id string, opt *GetEscalationOptions) (*Escalation, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	escalation := new(Escalation)
	resp, err := service.client.Do(req, escalation)
	if err != nil {
		return nil, resp, err
	}

	return escalation, resp, err
}

type CreateEscalationOptions struct {
	RouteId                     string    `json:"route_id,omitempty"`
	Position                    *int      `json:"position,omitempty"`
	Type                        string    `json:"type,omitempty"`
	Duration                    int       `json:"duration,omitempty"`
	PersonsToNotify             *[]string `json:"persons_to_notify,omitempty"`
	PersonsToNotifyNextEachTime *[]string `json:"persons_to_notify_next_each_time,omitempty"`
	NotifyOnCallFromSchedule    string    `json:"notify_on_call_from_schedule,omitempty"`
	ManualOrder                 bool      `url:"manual_order,omitempty" json:"manual_order,omitempty"`
}

// Create escalation with given name and type
//
// http://api-docs.amixr.io/#create-escalation
func (service *EscalationService) CreateEscalation(opt *CreateEscalationOptions) (*Escalation, *http.Response, error) {
	log.Printf("[DEBUG] create amixr escalation")
	u := fmt.Sprintf("%s/", service.url)
	req, err := service.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	escalation := new(Escalation)

	resp, err := service.client.Do(req, escalation)
	log.Printf("[DEBUG] request success")

	if err != nil {
		return nil, resp, err
	}

	return escalation, resp, err
}

type UpdateEscalationOptions struct {
	Position                 *int      `json:"position,omitempty"`
	Type                     string    `json:"type,omitempty"`
	Duration                 int       `json:"duration,omitempty"`
	PersonsToNotify          *[]string `json:"persons_to_notify,omitempty"`
	PersonsToNotifyEachTime  *[]string `json:"persons_to_notify_next_each_time,omitempty"`
	NotifyOnCallFromSchedule string    `json:"notify_on_call_from_schedule,omitempty"`
	ManualOrder              bool      `url:"manual_order,omitempty" json:"manual_order,omitempty"`
}

// Updates escalation with new templates and/or name. At least one field in template is required
//
// http://api-docs.amixr.io/#update-escalation
func (service *EscalationService) UpdateEscalation(id string, opt *UpdateEscalationOptions) (*Escalation, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	escalation := new(Escalation)
	resp, err := service.client.Do(req, escalation)
	if err != nil {
		return nil, resp, err
	}

	return escalation, resp, err
}

type DeleteEscalationOptions struct {
}

// Deletes escalation
//
// http://api-docs.amixr.io/#delete-escalation
func (service *EscalationService) DeleteEscalation(id string, opt *DeleteEscalationOptions) (*http.Response, error) {

	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("DELETE", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := service.client.Do(req, nil)
	return resp, err
}
