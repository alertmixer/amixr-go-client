package amixr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var testCustomAction = &CustomAction{
	ID:            "KGEFG74LU1D8L",
	Name:          "Test action",
	IntegrationId: "CGEXJ922S7TXQ",
}

var testCustomActionBody = `{
	"id": "KGEFG74LU1D8L",
	"name": "Test action",
	"integration_id": "CGEXJ922S7TXQ"
}`

func TestListCustomActions(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v1/actions/", func(w http.ResponseWriter, r *http.Request) {
		testRequestMethod(t, r, "GET")
		fmt.Fprint(w, fmt.Sprintf(`{"count": 1, "next": null, "previous": null, "results": [%s]}`, testCustomActionBody))
	})

	options := &ListCustomActionOptions{
		IntegrationId: "CGEXJ922S7TXQ",
	}

	customActions, _, err := client.CustomActions.ListCustomActions(options)
	if err != nil {
		t.Fatal(err)
	}

	want := &PaginatedCustomActionsResponse{
		PaginatedResponse: PaginatedResponse{
			Count:    1,
			Next:     nil,
			Previous: nil,
		},
		CustomActions: []*CustomAction{
			testCustomAction,
		},
	}
	if !reflect.DeepEqual(want, customActions) {
		t.Errorf("returned\n %+v, \nwant\n %+v", customActions, want)
	}
}
