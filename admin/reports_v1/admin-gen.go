// Package admin provides access to the Admin Reports API.
//
// See https://developers.google.com/admin-sdk/reports/
//
// Usage example:
//
//   import "github.com/jfcote87/api/admin/reports_v1"
//   ...
//   adminService, err := admin.New(oauthHttpClient)
package admin

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

const apiId = "admin:reports_v1"
const apiName = "admin"
const apiVersion = "reports_v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/admin/reports/v1/"}

// OAuth2 scopes used by this API.
const (
	// View audit reports of Google Apps for your domain
	AdminReportsAuditReadonlyScope = "https://www.googleapis.com/auth/admin.reports.audit.readonly"

	// View usage reports of Google Apps for your domain
	AdminReportsUsageReadonlyScope = "https://www.googleapis.com/auth/admin.reports.usage.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Activities = NewActivitiesService(s)
	s.Channels = NewChannelsService(s)
	s.CustomerUsageReports = NewCustomerUsageReportsService(s)
	s.UserUsageReport = NewUserUsageReportService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Activities *ActivitiesService

	Channels *ChannelsService

	CustomerUsageReports *CustomerUsageReportsService

	UserUsageReport *UserUsageReportService
}

func NewActivitiesService(s *Service) *ActivitiesService {
	rs := &ActivitiesService{s: s}
	return rs
}

type ActivitiesService struct {
	s *Service
}

func NewChannelsService(s *Service) *ChannelsService {
	rs := &ChannelsService{s: s}
	return rs
}

type ChannelsService struct {
	s *Service
}

func NewCustomerUsageReportsService(s *Service) *CustomerUsageReportsService {
	rs := &CustomerUsageReportsService{s: s}
	return rs
}

type CustomerUsageReportsService struct {
	s *Service
}

func NewUserUsageReportService(s *Service) *UserUsageReportService {
	rs := &UserUsageReportService{s: s}
	return rs
}

type UserUsageReportService struct {
	s *Service
}

type Activities struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: Each record in read response.
	Items []*Activity `json:"items,omitempty"`

	// Kind: Kind of list response this is.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token for retrieving the next page
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Activity struct {
	// Actor: User doing the action.
	Actor *ActivityActor `json:"actor,omitempty"`

	// Etag: ETag of the entry.
	Etag string `json:"etag,omitempty"`

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
	// CallerType: User or OAuth 2LO request.
	CallerType string `json:"callerType,omitempty"`

	// Email: Email address of the user.
	Email string `json:"email,omitempty"`

	// Key: For OAuth 2LO API requests, consumer_key of the requestor.
	Key string `json:"key,omitempty"`

	// ProfileId: Obfuscated user id of the user.
	ProfileId string `json:"profileId,omitempty"`
}

type ActivityEvents struct {
	// Name: Name of event.
	Name string `json:"name,omitempty"`

	// Parameters: Parameter value pairs for various applications.
	Parameters []*ActivityEventsParameters `json:"parameters,omitempty"`

	// Type: Type of event.
	Type string `json:"type,omitempty"`
}

type ActivityEventsParameters struct {
	// BoolValue: Boolean value of the parameter.
	BoolValue bool `json:"boolValue,omitempty"`

	// IntValue: Integral value of the parameter.
	IntValue int64 `json:"intValue,omitempty,string"`

	// MultiIntValue: Multi-int value of the parameter.
	MultiIntValue googleapi.Int64s `json:"multiIntValue,omitempty"`

	// MultiValue: Multi-string value of the parameter.
	MultiValue []string `json:"multiValue,omitempty"`

	// Name: The name of the parameter.
	Name string `json:"name,omitempty"`

	// Value: String value of the parameter.
	Value string `json:"value,omitempty"`
}

type ActivityId struct {
	// ApplicationName: Application name to which the event belongs.
	ApplicationName string `json:"applicationName,omitempty"`

	// CustomerId: Obfuscated customer ID of the source customer.
	CustomerId string `json:"customerId,omitempty"`

	// Time: Time of occurrence of the activity.
	Time string `json:"time,omitempty"`

	// UniqueQualifier: Unique qualifier if multiple events have the same
	// time.
	UniqueQualifier int64 `json:"uniqueQualifier,omitempty,string"`
}

type Channel struct {
	// Address: The address where notifications are delivered for this
	// channel.
	Address string `json:"address,omitempty"`

	// Expiration: Date and time of notification channel expiration,
	// expressed as a Unix timestamp, in milliseconds. Optional.
	Expiration int64 `json:"expiration,omitempty,string"`

	// Id: A UUID or similar unique string that identifies this channel.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this as a notification channel used to watch for
	// changes to a resource. Value: the fixed string "api#channel".
	Kind string `json:"kind,omitempty"`

	// Params: Additional parameters controlling delivery channel behavior.
	// Optional.
	Params map[string]string `json:"params,omitempty"`

	// Payload: A Boolean value to indicate whether payload is wanted.
	// Optional.
	Payload bool `json:"payload,omitempty"`

	// ResourceId: An opaque ID that identifies the resource being watched
	// on this channel. Stable across different API versions.
	ResourceId string `json:"resourceId,omitempty"`

	// ResourceUri: A version-specific identifier for the watched resource.
	ResourceUri string `json:"resourceUri,omitempty"`

	// Token: An arbitrary string delivered to the target address with each
	// notification delivered over this channel. Optional.
	Token string `json:"token,omitempty"`

	// Type: The type of delivery mechanism used for this channel.
	Type string `json:"type,omitempty"`
}

type UsageReport struct {
	// Date: The date to which the record belongs.
	Date string `json:"date,omitempty"`

	// Entity: Information about the type of the item.
	Entity *UsageReportEntity `json:"entity,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: The kind of object.
	Kind string `json:"kind,omitempty"`

	// Parameters: Parameter value pairs for various applications.
	Parameters []*UsageReportParameters `json:"parameters,omitempty"`
}

type UsageReportEntity struct {
	// CustomerId: Obfuscated customer id for the record.
	CustomerId string `json:"customerId,omitempty"`

	// ProfileId: Obfuscated user id for the record.
	ProfileId string `json:"profileId,omitempty"`

	// Type: The type of item, can be a customer or user.
	Type string `json:"type,omitempty"`

	// UserEmail: user's email.
	UserEmail string `json:"userEmail,omitempty"`
}

type UsageReportParameters struct {
	// BoolValue: Boolean value of the parameter.
	BoolValue bool `json:"boolValue,omitempty"`

	// DatetimeValue: RFC 3339 formatted value of the parameter.
	DatetimeValue string `json:"datetimeValue,omitempty"`

	// IntValue: Integral value of the parameter.
	IntValue int64 `json:"intValue,omitempty,string"`

	// MsgValue: Nested message value of the parameter.
	MsgValue []*UsageReportParametersMsgValue `json:"msgValue,omitempty"`

	// Name: The name of the parameter.
	Name string `json:"name,omitempty"`

	// StringValue: String value of the parameter.
	StringValue string `json:"stringValue,omitempty"`
}

type UsageReportParametersMsgValue struct {
}

type UsageReports struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: The kind of object.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token for retrieving the next page
	NextPageToken string `json:"nextPageToken,omitempty"`

	// UsageReports: Various application parameter records.
	UsageReports []*UsageReport `json:"usageReports,omitempty"`

	// Warnings: Warnings if any.
	Warnings []*UsageReportsWarnings `json:"warnings,omitempty"`
}

type UsageReportsWarnings struct {
	// Code: Machine readable code / warning type.
	Code string `json:"code,omitempty"`

	// Data: Key-Value pairs to give detailed information on the warning.
	Data []*UsageReportsWarningsData `json:"data,omitempty"`

	// Message: Human readable message for the warning.
	Message string `json:"message,omitempty"`
}

type UsageReportsWarningsData struct {
	// Key: Key associated with a key-value pair to give detailed
	// information on the warning.
	Key string `json:"key,omitempty"`

	// Value: Value associated with a key-value pair to give detailed
	// information on the warning.
	Value string `json:"value,omitempty"`
}

// method id "reports.activities.list":

type ActivitiesListCall struct {
	s               *Service
	userKey         string
	applicationName string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// List: Retrieves a list of activities for a specific customer and
// application.

func (r *ActivitiesService) List(userKey string, applicationName string) *ActivitiesListCall {
	return &ActivitiesListCall{
		s:               r.s,
		userKey:         userKey,
		applicationName: applicationName,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "activity/users/{userKey}/applications/{applicationName}",
	}
}

// ActorIpAddress sets the optional parameter "actorIpAddress": IP
// Address of host where the event was performed. Supports both IPv4 and
// IPv6 addresses.
func (c *ActivitiesListCall) ActorIpAddress(actorIpAddress string) *ActivitiesListCall {
	c.params_.Set("actorIpAddress", fmt.Sprintf("%v", actorIpAddress))
	return c
}

// CustomerId sets the optional parameter "customerId": Represents the
// customer for which the data is to be fetched.
func (c *ActivitiesListCall) CustomerId(customerId string) *ActivitiesListCall {
	c.params_.Set("customerId", fmt.Sprintf("%v", customerId))
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

// Filters sets the optional parameter "filters": Event parameters in
// the form [parameter1 name][operator][parameter1 value],[parameter2
// name][operator][parameter2 value],...
func (c *ActivitiesListCall) Filters(filters string) *ActivitiesListCall {
	c.params_.Set("filters", fmt.Sprintf("%v", filters))
	return c
}

// MaxResults sets the optional parameter "maxResults": Number of
// activity records to be shown in each page.
func (c *ActivitiesListCall) MaxResults(maxResults int64) *ActivitiesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *ActivitiesListCall) PageToken(pageToken string) *ActivitiesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
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
		"userKey":         c.userKey,
		"applicationName": c.applicationName,
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
	//   "id": "reports.activities.list",
	//   "parameterOrder": [
	//     "userKey",
	//     "applicationName"
	//   ],
	//   "parameters": {
	//     "actorIpAddress": {
	//       "description": "IP Address of host where the event was performed. Supports both IPv4 and IPv6 addresses.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "applicationName": {
	//       "description": "Application name for which the events are to be retrieved.",
	//       "location": "path",
	//       "pattern": "(admin)|(calendar)|(docs)|(drive)|(login)|(token)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Represents the customer for which the data is to be fetched.",
	//       "location": "query",
	//       "pattern": "C.+",
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Return events which occured at or before this time.",
	//       "location": "query",
	//       "pattern": "(\\d\\d\\d\\d)-(\\d\\d)-(\\d\\d)T(\\d\\d):(\\d\\d):(\\d\\d)(?:\\.(\\d+))?(?:(Z)|([-+])(\\d\\d):(\\d\\d))",
	//       "type": "string"
	//     },
	//     "eventName": {
	//       "description": "Name of the event being queried.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "Event parameters in the form [parameter1 name][operator][parameter1 value],[parameter2 name][operator][parameter2 value],...",
	//       "location": "query",
	//       "pattern": "(.+[\u003c,\u003c=,==,\u003e=,\u003e,\u003c\u003e].+,)*(.+[\u003c,\u003c=,==,\u003e=,\u003e,\u003c\u003e].+)",
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
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Return events which occured at or after this time.",
	//       "location": "query",
	//       "pattern": "(\\d\\d\\d\\d)-(\\d\\d)-(\\d\\d)T(\\d\\d):(\\d\\d):(\\d\\d)(?:\\.(\\d+))?(?:(Z)|([-+])(\\d\\d):(\\d\\d))",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Represents the profile id or the user email for which the data should be filtered. When 'all' is specified as the userKey, it returns usageReports for all users.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activity/users/{userKey}/applications/{applicationName}",
	//   "response": {
	//     "$ref": "Activities"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.audit.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "reports.activities.watch":

type ActivitiesWatchCall struct {
	s               *Service
	userKey         string
	applicationName string
	channel         *Channel
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Watch: Push changes to activities

func (r *ActivitiesService) Watch(userKey string, applicationName string, channel *Channel) *ActivitiesWatchCall {
	return &ActivitiesWatchCall{
		s:               r.s,
		userKey:         userKey,
		applicationName: applicationName,
		channel:         channel,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "activity/users/{userKey}/applications/{applicationName}/watch",
	}
}

// ActorIpAddress sets the optional parameter "actorIpAddress": IP
// Address of host where the event was performed. Supports both IPv4 and
// IPv6 addresses.
func (c *ActivitiesWatchCall) ActorIpAddress(actorIpAddress string) *ActivitiesWatchCall {
	c.params_.Set("actorIpAddress", fmt.Sprintf("%v", actorIpAddress))
	return c
}

// CustomerId sets the optional parameter "customerId": Represents the
// customer for which the data is to be fetched.
func (c *ActivitiesWatchCall) CustomerId(customerId string) *ActivitiesWatchCall {
	c.params_.Set("customerId", fmt.Sprintf("%v", customerId))
	return c
}

// EndTime sets the optional parameter "endTime": Return events which
// occured at or before this time.
func (c *ActivitiesWatchCall) EndTime(endTime string) *ActivitiesWatchCall {
	c.params_.Set("endTime", fmt.Sprintf("%v", endTime))
	return c
}

// EventName sets the optional parameter "eventName": Name of the event
// being queried.
func (c *ActivitiesWatchCall) EventName(eventName string) *ActivitiesWatchCall {
	c.params_.Set("eventName", fmt.Sprintf("%v", eventName))
	return c
}

// Filters sets the optional parameter "filters": Event parameters in
// the form [parameter1 name][operator][parameter1 value],[parameter2
// name][operator][parameter2 value],...
func (c *ActivitiesWatchCall) Filters(filters string) *ActivitiesWatchCall {
	c.params_.Set("filters", fmt.Sprintf("%v", filters))
	return c
}

// MaxResults sets the optional parameter "maxResults": Number of
// activity records to be shown in each page.
func (c *ActivitiesWatchCall) MaxResults(maxResults int64) *ActivitiesWatchCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *ActivitiesWatchCall) PageToken(pageToken string) *ActivitiesWatchCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// StartTime sets the optional parameter "startTime": Return events
// which occured at or after this time.
func (c *ActivitiesWatchCall) StartTime(startTime string) *ActivitiesWatchCall {
	c.params_.Set("startTime", fmt.Sprintf("%v", startTime))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ActivitiesWatchCall) Fields(s ...googleapi.Field) *ActivitiesWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ActivitiesWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey":         c.userKey,
		"applicationName": c.applicationName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Push changes to activities",
	//   "httpMethod": "POST",
	//   "id": "reports.activities.watch",
	//   "parameterOrder": [
	//     "userKey",
	//     "applicationName"
	//   ],
	//   "parameters": {
	//     "actorIpAddress": {
	//       "description": "IP Address of host where the event was performed. Supports both IPv4 and IPv6 addresses.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "applicationName": {
	//       "description": "Application name for which the events are to be retrieved.",
	//       "location": "path",
	//       "pattern": "(admin)|(calendar)|(docs)|(drive)|(login)|(token)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Represents the customer for which the data is to be fetched.",
	//       "location": "query",
	//       "pattern": "C.+",
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Return events which occured at or before this time.",
	//       "location": "query",
	//       "pattern": "(\\d\\d\\d\\d)-(\\d\\d)-(\\d\\d)T(\\d\\d):(\\d\\d):(\\d\\d)(?:\\.(\\d+))?(?:(Z)|([-+])(\\d\\d):(\\d\\d))",
	//       "type": "string"
	//     },
	//     "eventName": {
	//       "description": "Name of the event being queried.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "Event parameters in the form [parameter1 name][operator][parameter1 value],[parameter2 name][operator][parameter2 value],...",
	//       "location": "query",
	//       "pattern": "(.+[\u003c,\u003c=,==,\u003e=,\u003e,\u003c\u003e].+,)*(.+[\u003c,\u003c=,==,\u003e=,\u003e,\u003c\u003e].+)",
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
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Return events which occured at or after this time.",
	//       "location": "query",
	//       "pattern": "(\\d\\d\\d\\d)-(\\d\\d)-(\\d\\d)T(\\d\\d):(\\d\\d):(\\d\\d)(?:\\.(\\d+))?(?:(Z)|([-+])(\\d\\d):(\\d\\d))",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Represents the profile id or the user email for which the data should be filtered. When 'all' is specified as the userKey, it returns usageReports for all users.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "activity/users/{userKey}/applications/{applicationName}/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.audit.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "admin.channels.stop":

type ChannelsStopCall struct {
	s             *Service
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Stop: Stop watching resources through this channel

func (r *ChannelsService) Stop(channel *Channel) *ChannelsStopCall {
	return &ChannelsStopCall{
		s:             r.s,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "/admin/reports_v1/channels/stop",
	}
}

func (c *ChannelsStopCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Stop watching resources through this channel",
	//   "httpMethod": "POST",
	//   "id": "admin.channels.stop",
	//   "path": "/admin/reports_v1/channels/stop",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.audit.readonly"
	//   ]
	// }

}

// method id "reports.customerUsageReports.get":

type CustomerUsageReportsGetCall struct {
	s             *Service
	date          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a report which is a collection of properties /
// statistics for a specific customer.

func (r *CustomerUsageReportsService) Get(date string) *CustomerUsageReportsGetCall {
	return &CustomerUsageReportsGetCall{
		s:             r.s,
		date:          date,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "usage/dates/{date}",
	}
}

// CustomerId sets the optional parameter "customerId": Represents the
// customer for which the data is to be fetched.
func (c *CustomerUsageReportsGetCall) CustomerId(customerId string) *CustomerUsageReportsGetCall {
	c.params_.Set("customerId", fmt.Sprintf("%v", customerId))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *CustomerUsageReportsGetCall) PageToken(pageToken string) *CustomerUsageReportsGetCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Parameters sets the optional parameter "parameters": Represents the
// application name, parameter name pairs to fetch in csv as
// app_name1:param_name1, app_name2:param_name2.
func (c *CustomerUsageReportsGetCall) Parameters(parameters string) *CustomerUsageReportsGetCall {
	c.params_.Set("parameters", fmt.Sprintf("%v", parameters))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomerUsageReportsGetCall) Fields(s ...googleapi.Field) *CustomerUsageReportsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomerUsageReportsGetCall) Do() (*UsageReports, error) {
	var returnValue *UsageReports
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"date": c.date,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a report which is a collection of properties / statistics for a specific customer.",
	//   "httpMethod": "GET",
	//   "id": "reports.customerUsageReports.get",
	//   "parameterOrder": [
	//     "date"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Represents the customer for which the data is to be fetched.",
	//       "location": "query",
	//       "pattern": "C.+",
	//       "type": "string"
	//     },
	//     "date": {
	//       "description": "Represents the date in yyyy-mm-dd format for which the data is to be fetched.",
	//       "location": "path",
	//       "pattern": "(\\d){4}-(\\d){2}-(\\d){2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parameters": {
	//       "description": "Represents the application name, parameter name pairs to fetch in csv as app_name1:param_name1, app_name2:param_name2.",
	//       "location": "query",
	//       "pattern": "(((accounts)|(cros)|(gmail)|(calendar)|(docs)|(gplus)|(sites)|(device_management)):.+,)*(((accounts)|(cros)|(gmail)|(calendar)|(docs)|(gplus)|(sites)|(device_management)):.+)",
	//       "type": "string"
	//     }
	//   },
	//   "path": "usage/dates/{date}",
	//   "response": {
	//     "$ref": "UsageReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.usage.readonly"
	//   ]
	// }

}

// method id "reports.userUsageReport.get":

type UserUsageReportGetCall struct {
	s             *Service
	userKey       string
	date          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a report which is a collection of properties /
// statistics for a set of users.

func (r *UserUsageReportService) Get(userKey string, date string) *UserUsageReportGetCall {
	return &UserUsageReportGetCall{
		s:             r.s,
		userKey:       userKey,
		date:          date,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "usage/users/{userKey}/dates/{date}",
	}
}

// CustomerId sets the optional parameter "customerId": Represents the
// customer for which the data is to be fetched.
func (c *UserUsageReportGetCall) CustomerId(customerId string) *UserUsageReportGetCall {
	c.params_.Set("customerId", fmt.Sprintf("%v", customerId))
	return c
}

// Filters sets the optional parameter "filters": Represents the set of
// filters including parameter operator value.
func (c *UserUsageReportGetCall) Filters(filters string) *UserUsageReportGetCall {
	c.params_.Set("filters", fmt.Sprintf("%v", filters))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Maximum allowed is 1000
func (c *UserUsageReportGetCall) MaxResults(maxResults int64) *UserUsageReportGetCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page.
func (c *UserUsageReportGetCall) PageToken(pageToken string) *UserUsageReportGetCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Parameters sets the optional parameter "parameters": Represents the
// application name, parameter name pairs to fetch in csv as
// app_name1:param_name1, app_name2:param_name2.
func (c *UserUsageReportGetCall) Parameters(parameters string) *UserUsageReportGetCall {
	c.params_.Set("parameters", fmt.Sprintf("%v", parameters))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserUsageReportGetCall) Fields(s ...googleapi.Field) *UserUsageReportGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserUsageReportGetCall) Do() (*UsageReports, error) {
	var returnValue *UsageReports
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
		"date":    c.date,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a report which is a collection of properties / statistics for a set of users.",
	//   "httpMethod": "GET",
	//   "id": "reports.userUsageReport.get",
	//   "parameterOrder": [
	//     "userKey",
	//     "date"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Represents the customer for which the data is to be fetched.",
	//       "location": "query",
	//       "pattern": "C.+",
	//       "type": "string"
	//     },
	//     "date": {
	//       "description": "Represents the date in yyyy-mm-dd format for which the data is to be fetched.",
	//       "location": "path",
	//       "pattern": "(\\d){4}-(\\d){2}-(\\d){2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "Represents the set of filters including parameter operator value.",
	//       "location": "query",
	//       "pattern": "(((accounts)|(cros)|(gmail)|(calendar)|(docs)|(gplus)|(sites)|(device_management)):.+[\u003c,\u003c=,==,\u003e=,\u003e,!=].+,)*(((accounts)|(cros)|(gmail)|(calendar)|(docs)|(gplus)|(sites)|(device_management)):.+[\u003c,\u003c=,==,\u003e=,\u003e,!=].+)",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Maximum allowed is 1000",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parameters": {
	//       "description": "Represents the application name, parameter name pairs to fetch in csv as app_name1:param_name1, app_name2:param_name2.",
	//       "location": "query",
	//       "pattern": "(((accounts)|(cros)|(gmail)|(calendar)|(docs)|(gplus)|(sites)|(device_management)):.+,)*(((accounts)|(cros)|(gmail)|(calendar)|(docs)|(gplus)|(sites)|(device_management)):.+)",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Represents the profile id or the user email for which the data should be filtered.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "usage/users/{userKey}/dates/{date}",
	//   "response": {
	//     "$ref": "UsageReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.reports.usage.readonly"
	//   ]
	// }

}
