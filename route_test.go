package amixr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var testSlackChannelId = "TEST_SLACK_CHANNEL_ID"

var testRoute = &Route{
	ID:             "RH2V5FYIPYJ1M",
	IntegrationId:  "CGEXJ922S7TXQ",
	Position:       0,
	RoutingRegex:   "us-west",
	IsTheLastRoute: false,
	SlackRoute: &SlackRoute{
		&testSlackChannelId,
	},
}

var testRouteBody = `{
	"id": "RH2V5FYIPYJ1M",
	"integration_id": "CGEXJ922S7TXQ",
	"routing_regex": "us-west",
	"position": 0,
	"is_the_last_route": false,
	"slack": {
		"channel_id": "TEST_SLACK_CHANNEL_ID"
	}
}`

func TestCreateRoute(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/routes/", func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "POST")
		fmt.Fprint(w, testRouteBody)
	})

	createOptions := &CreateRouteOptions{
		IntegrationId: "CGEXJ922S7TXQ",
		RoutingRegex:  "us-west",
		Slack: &SlackRoute{
			&testSlackChannelId,
		},
	}
	route, _, err := client.Routes.CreateRoute(createOptions)

	if err != nil {
		t.Fatal(err)
	}

	want := testRoute

	if !reflect.DeepEqual(want, route) {
		t.Errorf("returned\n %+v\n want\n %+v\n", route, want)
	}
}

func TestDeleteRoute(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/routes/RH2V5FYIPYJ1M/", func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "DELETE")
	})

	options := &DeleteRouteOptions{}

	_, err := client.Routes.DeleteRoute("RH2V5FYIPYJ1M", options)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListRoutes(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/routes/", func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "GET")
		fmt.Fprint(w, fmt.Sprintf(`{"count": 1, "next": null, "previous": null, "results": [%s]}`, testRouteBody))
	})

	options := &ListRouteOptions{
		IntegrationId: "CGEXJ922S7TXQ",
	}

	routes, _, err := client.Routes.ListRoutes(options)
	if err != nil {
		t.Fatal(err)
	}

	want := &PaginatedRoutesResponse{
		PaginatedResponse: PaginatedResponse{
			Count:    1,
			Next:     nil,
			Previous: nil,
		},
		Routes: []*Route{
			testRoute,
		},
	}
	if !reflect.DeepEqual(want, routes) {

		t.Errorf(" returned\n %+v, \nwant\n %+v", routes, want)
	}
}

func TestGetRoute(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/routes/RH2V5FYIPYJ1M/", func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "GET")
		fmt.Fprint(w, testRouteBody)
	})

	options := &GetRouteOptions{}

	route, _, err := client.Routes.GetRoute("RH2V5FYIPYJ1M", options)

	if err != nil {
		t.Fatal(err)
	}

	want := testRoute

	if !reflect.DeepEqual(want, route) {
		t.Errorf("returned\n %+v\n want\n %+v\n", route, want)
	}
}
