package amixr

import (
	"fmt"
	"log"
	"net/http"
)

type OnCallShiftService struct {
	client *Client
	url    string
}

// NewOnCallShiftServiceService creates OnCallShiftService with defined url
func NewOnCallShiftService(client *Client) *OnCallShiftService {
	onCallShiftService := OnCallShiftService{}
	onCallShiftService.client = client
	onCallShiftService.url = "on_call_shifts"
	return &onCallShiftService
}

type PaginatedOnCallShiftsResponse struct {
	PaginatedResponse
	OnCallShifts []*OnCallShift `json:"results"`
}

type OnCallShift struct {
	ID                         string      `json:"id"`
	Type                       string      `json:"type"`
	Name                       string      `json:"name"`
	Level                      int         `json:"level"`
	Start                      string      `json:"start"`
	Duration                   int         `json:"duration"`
	Frequency                  *string     `json:"frequency"`
	Users                      *[]string   `json:"users"`
	Interval                   *int        `json:"interval"`
	WeekStart                  *string     `json:"week_start"`
	ByDay                      *[]string   `json:"by_day"`
	ByMonth                    *[]int      `json:"by_month"`
	ByMonthday                 *[]int      `json:"by_monthday"`
	RollingUsers               *[][]string `json:"rolling_users"`
	TimeZone                   *string     `json:"time_zone"`
	StartRotationFromUserIndex *int        `json:"start_rotation_from_user_index"`
}

type ListOnCallShiftOptions struct {
	ListOptions
	ScheduleId string `url:"schedule_id,omitempty" json:"schedule_id,omitempty"`
	Name       string `url:"name,omitempty" json:"name,omitempty"`
}

// ListOnCallShifts gets all on call shifts for authorized organization
func (service *OnCallShiftService) ListOnCallShifts(opt *ListOnCallShiftOptions) (*PaginatedOnCallShiftsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var onCallShifts *PaginatedOnCallShiftsResponse
	resp, err := service.client.Do(req, &onCallShifts)
	if err != nil {
		return nil, resp, err
	}

	return onCallShifts, resp, err
}

type GetOnCallShiftOptions struct {
}

// Get on-call shift by given id
func (service *OnCallShiftService) GetOnCallShift(id string, opt *GetOnCallShiftOptions) (*OnCallShift, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	onCallShift := new(OnCallShift)
	resp, err := service.client.Do(req, onCallShift)
	if err != nil {
		return nil, resp, err
	}

	return onCallShift, resp, err
}

type CreateOnCallShiftOptions struct {
	Type                       string      `json:"type"`
	Name                       string      `json:"name"`
	Level                      *int        `json:"level,omitempty"`
	Start                      string      `json:"start"`
	Duration                   int         `json:"duration"`
	Frequency                  *string     `json:"frequency"`
	Users                      *[]string   `json:"users"`
	Interval                   *int        `json:"interval"`
	WeekStart                  *string     `json:"week_start,omitempty"`
	ByDay                      *[]string   `json:"by_day"`
	ByMonth                    *[]int      `json:"by_month"`
	ByMonthday                 *[]int      `json:"by_monthday"`
	Source                     int         `json:"source"`
	RollingUsers               *[][]string `json:"rolling_users"`
	TimeZone                   *string     `json:"time_zone"`
	StartRotationFromUserIndex *int        `json:"start_rotation_from_user_index"`
}

// Create on-call shift
func (service *OnCallShiftService) CreateOnCallShift(opt *CreateOnCallShiftOptions) (*OnCallShift, *http.Response, error) {
	log.Printf("[DEBUG] create amixr on_call_shift")
	u := fmt.Sprintf("%s/", service.url)
	req, err := service.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	onCallShift := new(OnCallShift)

	resp, err := service.client.Do(req, onCallShift)

	if err != nil {
		return nil, resp, err
	}

	return onCallShift, resp, err
}

type UpdateOnCallShiftOptions struct {
	Type                       string      `json:"type"`
	Name                       string      `json:"name"`
	Level                      *int        `json:"level,omitempty"`
	Start                      string      `json:"start"`
	Duration                   int         `json:"duration"`
	Frequency                  *string     `json:"frequency"`
	Users                      *[]string   `json:"users"`
	Interval                   *int        `json:"interval"`
	WeekStart                  *string     `json:"week_start,omitempty"`
	ByDay                      *[]string   `json:"by_day"`
	ByMonth                    *[]int      `json:"by_month"`
	ByMonthday                 *[]int      `json:"by_monthday"`
	Source                     int         `json:"source"`
	RollingUsers               *[][]string `json:"rolling_users"`
	TimeZone                   *string     `json:"time_zone"`
	StartRotationFromUserIndex *int        `json:"start_rotation_from_user_index"`
}

// Updates on-call shift
func (service *OnCallShiftService) UpdateOnCallShift(id string, opt *UpdateOnCallShiftOptions) (*OnCallShift, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	onCallShift := new(OnCallShift)
	resp, err := service.client.Do(req, onCallShift)
	if err != nil {
		return nil, resp, err
	}

	return onCallShift, resp, err
}

type DeleteOnCallShiftOptions struct {
}

// Deletes on-call shift
func (service *OnCallShiftService) DeleteOnCallShift(id string, opt *DeleteOnCallShiftOptions) (*http.Response, error) {

	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("DELETE", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := service.client.Do(req, nil)
	return resp, err
}
