package amixr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var testSchedule = &Schedule{
	ID:        "SBM7DV7BKFUYU",
	Type:      "ical",
	OnCallNow: []string{"U4DNY931HHJS5", "U6RV9WPSL6DFW"},
}

var testScheduleBody = `{
	"id": "SBM7DV7BKFUYU",
	"type": "ical",
	"on_call_now": [
	"U4DNY931HHJS5",
	"U6RV9WPSL6DFW"
	]
}`

func TestListSchedules(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/schedules/", func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "GET")
		fmt.Fprint(w, fmt.Sprintf(`{"count": 1, "next": null, "previous": null, "results": [%s]}`, testScheduleBody))
	})

	options := &ListScheduleOptions{}

	schedules, _, err := client.Schedules.ListSchedules(options)
	if err != nil {
		t.Fatal(err)
	}

	want := &PaginatedSchedulesResponse{
		PaginatedResponse: PaginatedResponse{
			Count:    1,
			Next:     nil,
			Previous: nil,
		},
		Schedules: []*Schedule{
			testSchedule,
		},
	}
	if !reflect.DeepEqual(want, schedules) {
		t.Errorf("returned\n %+v, \nwant\n %+v", schedules, want)
	}
}
