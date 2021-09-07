package amixr

import (
	"fmt"
	"log"
	"net/http"
)

// ScheduleService handles requests to schedule endpoint
// Use NewScheduleService instead of direct creation ScheduleService
//
// http://api-docs.amixr.io/#schedules
type ScheduleService struct {
	client *Client
	url    string
}

// NewScheduleService creates ScheduleService with defined url
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
	ID               string         `json:"id"`
	Type             string         `json:"type"`
	OnCallNow        []string       `json:"on_call_now"`
	Name             string         `json:"name"`
	ICalUrlPrimary   *string        `json:"ical_url_primary"`
	ICalUrlOverrides *string        `json:"ical_url_overrides"`
	TimeZone         string         `json:"time_zone"`
	Slack            *SlackSchedule `json:"slack"`
}

type SlackSchedule struct {
	ChannelId *string `json:"channel_id"`
}

type ListScheduleOptions struct {
	ListOptions
	Name string `url:"name,omitempty" json:"name,omitempty"`
}

// ListSchedules gets all schedules for authorized organization
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

type GetScheduleOptions struct {
}

// Get schedule shift by given id
func (service *ScheduleService) GetSchedule(id string, opt *GetScheduleOptions) (*Schedule, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	schedule := new(Schedule)
	resp, err := service.client.Do(req, schedule)
	if err != nil {
		return nil, resp, err
	}

	return schedule, resp, err
}

type CreateScheduleOptions struct {
	Name             string         `json:"name"`
	Type             string         `json:"type"`
	ICalUrlPrimary   *string        `json:"ical_url_primary"`
	ICalUrlOverrides *string        `json:"ical_url_overrides"`
	TimeZone         string         `json:"time_zone"`
	Slack            *SlackSchedule `json:"slack,omitempty"`
}

// Create schedule with given name
func (service *ScheduleService) CreateSchedule(opt *CreateScheduleOptions) (*Schedule, *http.Response, error) {
	log.Printf("[DEBUG] create amixr schedule")
	u := fmt.Sprintf("%s/", service.url)
	req, err := service.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	schedule := new(Schedule)

	resp, err := service.client.Do(req, schedule)

	if err != nil {
		return nil, resp, err
	}

	return schedule, resp, err
}

type UpdateScheduleOptions struct {
	Name             string         `json:"name,omitempty"`
	ICalUrlPrimary   *string        `json:"ical_url_primary"`
	ICalUrlOverrides *string        `json:"ical_url_overrides"`
	TimeZone         string         `json:"time_zone"`
	Slack            *SlackSchedule `json:"slack,omitempty"`
}

// Updates schedule
func (service *ScheduleService) UpdateSchedule(id string, opt *UpdateScheduleOptions) (*Schedule, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	schedule := new(Schedule)
	resp, err := service.client.Do(req, schedule)
	if err != nil {
		return nil, resp, err
	}

	return schedule, resp, err
}

type DeleteScheduleOptions struct {
}

// Deletes schedule
func (service *ScheduleService) DeleteSchedule(id string, opt *DeleteScheduleOptions) (*http.Response, error) {

	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("DELETE", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := service.client.Do(req, nil)
	return resp, err
}
