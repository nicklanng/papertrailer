package events

import (
	"log"

	"github.com/go-resty/resty"
	"github.com/nicklanng/papertrailer/pkg/config"
)

var filter = "program:application severity:(error warning) -\"Investigating zombie call channels\" -\"initial applet not set\" -\"Possible Data Provider config issue\" -\"This error will be suppressed for 60 minutes\" -\"Unknown call state when attempting to write stats event\" -REQUEST_LIMIT_EXCEEDED -\"inactive user\" -ProtocolFaultResponseException"

func Fetch(groupID, minId string) *EventResponse {
	queryParams := map[string]string{
		"group_id": groupID,
		"q":        filter,
	}

	if minId != "" {
		queryParams["min_id"] = minId
	}

	resp, err := resty.R().
		SetQueryParams(queryParams).
		SetHeader("X-Papertrail-Token", config.Config.Token).
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		SetResult(&EventResponse{}).
		Get("https://papertrailapp.com/api/v1/events/search.json")

	if err != nil {
		log.Fatal(err)
	}

	return resp.Result().(*EventResponse)
}
