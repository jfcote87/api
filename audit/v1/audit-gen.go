// Package audit provides access to the Enterprise Audit API.
//
// See https://developers.google.com/google-apps/admin-audit/get_started
//
// Usage example:
//
//   import "github.com/jfcote87/api/audit/v1"
//   ...
//   auditService, err := audit.New(oauthHttpClient)
package audit

import (
	"errors"
	"fmt"
	"github.com/jfcote87/api/googleapi"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "audit:v1"
const apiName = "audit"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/apps/reporting/audit/v1/"}

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Activities = NewActivitiesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Activities *ActivitiesService
}

func NewActivitiesService(s *Service) *ActivitiesService {
	rs := &ActivitiesService{s: s}
	return rs
}

type ActivitiesService struct {
	s *Service
}

type Activities struct {
	// Items: Each record in read response.
	Items []*Activity `json:"items,omitempty"`

	// Kind: Kind of list response this is.
	Kind string `json:"kind,omitempty"`

	// Next: Next page URL.
	Next string `json:"next,omitempty"`
}

type Activity struct {
	// Actor: User doing the action.
	Actor *ActivityActor `json:"actor,omitempty"`

	// Events: Activity events.
	Events []*ActivityEvents `json:"events,omitempty"`

	// Id: Unique identifier for each activity record.
	Id *ActivityId `json:"id,omitempty"`

	// IpAddress: IP Address of the user doing the action.
	IpAddress string `json:"ipAddress,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// OwnerDomain: Domain of source customer.
	OwnerDomain string `json:"ownerDomain,omitempty"`
}

type ActivityActor struct {
	// ApplicationId: ID of application which interacted on behalf of the
	// user.
	ApplicationId int64 `json:"applicationId,omitempty,string"`

	// CallerType: User or OAuth 2LO request.
	CallerType string `json:"callerType,omitempty"`

	// Email: Email address of the user.
	Email string `json:"email,omitempty"`

	// Key: For OAuth 2LO API requests, consumer_key of the requestor.
	Key string `json:"key,omitempty"`
}

type ActivityEvents struct {
	// EventType: Type of event.
	EventType string `json:"eventType,omitempty"`

	// Name: Name of event.
	Name string `json:"name,omitempty"`

	// Parameters: Event parameters.
	Parameters []*ActivityEventsParameters `json:"parameters,omitempty"`
}

type ActivityEventsParameters struct {
	// Name: Name of the parameter.
	Name string `json:"name,omitempty"`

	// Value: Value of the parameter.
	Value string `json:"value,omitempty"`
}

type ActivityId struct {
	// ApplicationId: Application ID of the source application.
	ApplicationId int64 `json:"applicationId,omitempty,string"`

	// CustomerId: Obfuscated customer ID of the source customer.
	CustomerId string `json:"customerId,omitempty"`

	// Time: Time of occurrence of the activity.
	Time string `json:"time,omitempty"`

	// UniqQualifier: Unique qualifier if multiple events have the same
	// time.
	UniqQualifier int64 `json:"uniqQualifier,omitempty,string"`
}

// method id "audit.activities.list":

type ActivitiesListCall struct {
	s             *Service
	customerId    string
	applicationId int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of activities for a specific customer and
// application.

func (r *ActivitiesService) List(customerId string, applicationId int64) *ActivitiesListCall {
	return &ActivitiesListCall{
		s:             r.s,
		customerId:    customerId,
		applicationId: applicationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{customerId}/{applicationId}",
	}
}

// ActorApplicationId sets the optional parameter "actorApplicationId":
// Application ID of the application which interacted on behalf of the
// user while performing the event.
func (c *ActivitiesListCall) ActorApplicationId(actorApplicationId int64) *ActivitiesListCall {
	c.params_.Set("actorApplicationId", fmt.Sprintf("%v", actorApplicationId))
	return c
}

// ActorEmail sets the optional parameter "actorEmail": Email address of
// the user who performed the action.
func (c *ActivitiesListCall) ActorEmail(actorEmail string) *ActivitiesListCall {
	c.params_.Set("actorEmail", fmt.Sprintf("%v", actorEmail))
	return c
}

// ActorIpAddress sets the optional parameter "actorIpAddress": IP
// Address of host where the event was performed. Supports both IPv4 and
// IPv6 addresses.
func (c *ActivitiesListCall) ActorIpAddress(actorIpAddress string) *ActivitiesListCall {
	c.params_.Set("actorIpAddress", fmt.Sprintf("%v", actorIpAddress))
	return c
}

// Caller sets the optional parameter "caller": Type of the caller.
func (c *ActivitiesListCall) Caller(caller string) *ActivitiesListCall {
	c.params_.Set("caller", fmt.Sprintf("%v", caller))
	return c
}

// ContinuationToken sets the optional parameter "continuationToken":
// Next page URL.
func (c *ActivitiesListCall) ContinuationToken(continuationToken string) *ActivitiesListCall {
	c.params_.Set("continuationToken", fmt.Sprintf("%v", continuationToken))
	return c
}

// EndTime sets the optional parameter "endTime": Return events which
// occured at or before this time.
func (c *ActivitiesListCall) EndTime(endTime string) *ActivitiesListCall {
	c.params_.Set("endTime", fmt.Sprintf("%v", endTime))
	return c
}

// EventName sets the optional parameter "eventName": Name of the event
// being queried.
func (c *ActivitiesListCall) EventName(eventName string) *ActivitiesListCall {
	c.params_.Set("eventName", fmt.Sprintf("%v", eventName))
	return c
}

// MaxResults sets the optional parameter "maxResults": Number of
// activity records to be shown in each page.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// StartTime sets the optional parameter "startTime": Return events
// which occured at or after this time.
func (c *ActivitiesListCall) StartTime(startTime string) *ActivitiesListCall {
	c.params_.Set("startTime", fmt.Sprintf("%v", startTime))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ActivitiesListCall) Fields(s ...googleapi.Field) *ActivitiesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ActivitiesListCall) Do() (*Activities, error) {
	var returnValue *Activities
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":    c.customerId,
		"applicationId": strconv.FormatInt(c.applicationId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of activities for a specific customer and application.",
	//   "httpMethod": "GET",
	//   "id": "audit.activities.list",
	//   "parameterOrder": [
	//     "customerId",
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "actorApplicationId": {
	//       "description": "Application ID of the application which interacted on behalf of the user while performing the event.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "actorEmail": {
	//       "description": "Email address of the user who performed the action.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "actorIpAddress": {
	//       "description": "IP Address of host where the event was performed. Supports both IPv4 and IPv6 addresses.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "applicationId": {
	//       "description": "Application ID of the application on which the event was performed.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "caller": {
	//       "description": "Type of the caller.",
	//       "enum": [
	//         "application_owner",
	//         "customer"
	//       ],
	//       "enumDescriptions": [
	//         "Caller is an application owner.",
	//         "Caller is a customer."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "continuationToken": {
	//       "description": "Next page URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Represents the customer who is the owner of target object on which action was performed.",
	//       "location": "path",
	//       "pattern": "C.+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Return events which occured at or before this time.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "eventName": {
	//       "description": "Name of the event being queried.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Number of activity records to be shown in each page.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "startTime": {
	//       "description": "Return events which occured at or after this time.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{customerId}/{applicationId}",
	//   "response": {
	//     "$ref": "Activities"
	//   }
	// }

}
