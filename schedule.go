package amixr

import (
	"fmt"
	"net/http"
)

// Handles requests to schedule endpoint
// Use NewScheduleService instead of direct creation ScheduleService
//
// http://api-docs.amixr.io/#schedules
type ScheduleService struct {
	client *Client
	url    string
}

// NewSchedulesService creates ScheduleService with defined url
func NewScheduleService(client *Client) *ScheduleService {
	scheduleService := ScheduleService{}
	scheduleService.client = client
	scheduleService.url = "schedules"
	return &scheduleService
}

type PaginatedSchedulesResponse struct {
	PaginatedResponse
	Schedules []*Schedule `json:"results"`
}

type Schedule struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	OnCallNow []string `json:"on_call_now"`
	Name      string   `json:"name"`
}

type ListScheduleOptions struct {
	ListOptions
	Name string `url:"name,omitempty" json:"name,omitempty"`
}

// ListSchedules gets all schedules for authorized team
//
// http://api-docs.amixr.io/#list-schedules
func (service *ScheduleService) ListSchedules(opt *ListScheduleOptions) (*PaginatedSchedulesResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var schedules *PaginatedSchedulesResponse
	resp, err := service.client.Do(req, &schedules)
	if err != nil {
		return nil, resp, err
	}

	return schedules, resp, err
}
