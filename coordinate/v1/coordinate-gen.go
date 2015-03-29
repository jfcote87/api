// Package coordinate provides access to the Google Maps Coordinate API.
//
// See https://developers.google.com/coordinate/
//
// Usage example:
//
//   import "github.com/jfcote87/api/coordinate/v1"
//   ...
//   coordinateService, err := coordinate.New(oauthHttpClient)
package coordinate

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

const apiId = "coordinate:v1"
const apiName = "coordinate"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/coordinate/v1/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Maps Coordinate jobs
	CoordinateScope = "https://www.googleapis.com/auth/coordinate"

	// View your Google Coordinate jobs
	CoordinateReadonlyScope = "https://www.googleapis.com/auth/coordinate.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.CustomFieldDef = NewCustomFieldDefService(s)
	s.Jobs = NewJobsService(s)
	s.Location = NewLocationService(s)
	s.Schedule = NewScheduleService(s)
	s.Team = NewTeamService(s)
	s.Worker = NewWorkerService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	CustomFieldDef *CustomFieldDefService

	Jobs *JobsService

	Location *LocationService

	Schedule *ScheduleService

	Team *TeamService

	Worker *WorkerService
}

func NewCustomFieldDefService(s *Service) *CustomFieldDefService {
	rs := &CustomFieldDefService{s: s}
	return rs
}

type CustomFieldDefService struct {
	s *Service
}

func NewJobsService(s *Service) *JobsService {
	rs := &JobsService{s: s}
	return rs
}

type JobsService struct {
	s *Service
}

func NewLocationService(s *Service) *LocationService {
	rs := &LocationService{s: s}
	return rs
}

type LocationService struct {
	s *Service
}

func NewScheduleService(s *Service) *ScheduleService {
	rs := &ScheduleService{s: s}
	return rs
}

type ScheduleService struct {
	s *Service
}

func NewTeamService(s *Service) *TeamService {
	rs := &TeamService{s: s}
	return rs
}

type TeamService struct {
	s *Service
}

func NewWorkerService(s *Service) *WorkerService {
	rs := &WorkerService{s: s}
	return rs
}

type WorkerService struct {
	s *Service
}

type CustomField struct {
	// CustomFieldId: Custom field id.
	CustomFieldId int64 `json:"customFieldId,omitempty,string"`

	// Kind: Identifies this object as a custom field.
	Kind string `json:"kind,omitempty"`

	// Value: Custom field value.
	Value string `json:"value,omitempty"`
}

type CustomFieldDef struct {
	// Enabled: Whether the field is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// Enumitems: List of enum items for this custom field. Populated only
	// if the field type is enum. Enum fields appear as 'lists' in the
	// Coordinate web and mobile UI.
	Enumitems []*EnumItemDef `json:"enumitems,omitempty"`

	// Id: Custom field id.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies this object as a custom field definition.
	Kind string `json:"kind,omitempty"`

	// Name: Custom field name.
	Name string `json:"name,omitempty"`

	// RequiredForCheckout: Whether the field is required for checkout.
	RequiredForCheckout bool `json:"requiredForCheckout,omitempty"`

	// Type: Custom field type.
	Type string `json:"type,omitempty"`
}

type CustomFieldDefListResponse struct {
	// Items: Collection of custom field definitions in a team.
	Items []*CustomFieldDef `json:"items,omitempty"`

	// Kind: Identifies this object as a collection of custom field
	// definitions in a team.
	Kind string `json:"kind,omitempty"`
}

type CustomFields struct {
	// CustomField: Collection of custom fields.
	CustomField []*CustomField `json:"customField,omitempty"`

	// Kind: Identifies this object as a collection of custom fields.
	Kind string `json:"kind,omitempty"`
}

type EnumItemDef struct {
	// Active: Whether the enum item is active. Jobs may contain inactive
	// enum values; however, setting an enum to an inactive value when
	// creating or updating a job will result in a 500 error.
	Active bool `json:"active,omitempty"`

	// Kind: Identifies this object as an enum item definition.
	Kind string `json:"kind,omitempty"`

	// Value: Custom field value.
	Value string `json:"value,omitempty"`
}

type Job struct {
	// Id: Job id.
	Id uint64 `json:"id,omitempty,string"`

	// JobChange: List of job changes since it was created. The first change
	// corresponds to the state of the job when it was created.
	JobChange []*JobChange `json:"jobChange,omitempty"`

	// Kind: Identifies this object as a job.
	Kind string `json:"kind,omitempty"`

	// State: Current job state.
	State *JobState `json:"state,omitempty"`
}

type JobChange struct {
	// Kind: Identifies this object as a job change.
	Kind string `json:"kind,omitempty"`

	// State: Change applied to the job. Only the fields that were changed
	// are set.
	State *JobState `json:"state,omitempty"`

	// Timestamp: Time at which this change was applied.
	Timestamp uint64 `json:"timestamp,omitempty,string"`
}

type JobListResponse struct {
	// Items: Jobs in the collection.
	Items []*Job `json:"items,omitempty"`

	// Kind: Identifies this object as a list of jobs.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to provide to get the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type JobState struct {
	// Assignee: Email address of the assignee, or the string "DELETED_USER"
	// if the account is no longer available.
	Assignee string `json:"assignee,omitempty"`

	// CustomFields: Custom fields.
	CustomFields *CustomFields `json:"customFields,omitempty"`

	// CustomerName: Customer name.
	CustomerName string `json:"customerName,omitempty"`

	// CustomerPhoneNumber: Customer phone number.
	CustomerPhoneNumber string `json:"customerPhoneNumber,omitempty"`

	// Kind: Identifies this object as a job state.
	Kind string `json:"kind,omitempty"`

	// Location: Job location.
	Location *Location `json:"location,omitempty"`

	// Note: Note added to the job.
	Note []string `json:"note,omitempty"`

	// Progress: Job progress.
	Progress string `json:"progress,omitempty"`

	// Title: Job title.
	Title string `json:"title,omitempty"`
}

type Location struct {
	// AddressLine: Address.
	AddressLine []string `json:"addressLine,omitempty"`

	// Kind: Identifies this object as a location.
	Kind string `json:"kind,omitempty"`

	// Lat: Latitude.
	Lat float64 `json:"lat,omitempty"`

	// Lng: Longitude.
	Lng float64 `json:"lng,omitempty"`
}

type LocationListResponse struct {
	// Items: Locations in the collection.
	Items []*LocationRecord `json:"items,omitempty"`

	// Kind: Identifies this object as a list of locations.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to provide to get the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TokenPagination: Pagination information for token pagination.
	TokenPagination *TokenPagination `json:"tokenPagination,omitempty"`
}

type LocationRecord struct {
	// CollectionTime: The collection time in milliseconds since the epoch.
	CollectionTime int64 `json:"collectionTime,omitempty,string"`

	// ConfidenceRadius: The location accuracy in meters. This is the radius
	// of a 95% confidence interval around the location measurement.
	ConfidenceRadius float64 `json:"confidenceRadius,omitempty"`

	// Kind: Identifies this object as a location.
	Kind string `json:"kind,omitempty"`

	// Latitude: Latitude.
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: Longitude.
	Longitude float64 `json:"longitude,omitempty"`
}

type Schedule struct {
	// AllDay: Whether the job is scheduled for the whole day. Time of day
	// in start/end times is ignored if this is true.
	AllDay bool `json:"allDay,omitempty"`

	// Duration: Job duration in milliseconds.
	Duration uint64 `json:"duration,omitempty,string"`

	// EndTime: Scheduled end time in milliseconds since epoch.
	EndTime uint64 `json:"endTime,omitempty,string"`

	// Kind: Identifies this object as a job schedule.
	Kind string `json:"kind,omitempty"`

	// StartTime: Scheduled start time in milliseconds since epoch.
	StartTime uint64 `json:"startTime,omitempty,string"`
}

type Team struct {
	// Id: Team id, as found in a coordinate team url e.g.
	// https://coordinate.google.com/f/xyz where "xyz" is the team id.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this object as a team.
	Kind string `json:"kind,omitempty"`

	// Name: Team name
	Name string `json:"name,omitempty"`
}

type TeamListResponse struct {
	// Items: Teams in the collection.
	Items []*Team `json:"items,omitempty"`

	// Kind: Identifies this object as a list of teams.
	Kind string `json:"kind,omitempty"`
}

type TokenPagination struct {
	// Kind: Identifies this object as pagination information.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to provide to get the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PreviousPageToken: A token to provide to get the previous page of
	// results.
	PreviousPageToken string `json:"previousPageToken,omitempty"`
}

type Worker struct {
	// Id: Worker email address. If a worker has been deleted from your
	// team, the email address will appear as DELETED_USER.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this object as a worker.
	Kind string `json:"kind,omitempty"`
}

type WorkerListResponse struct {
	// Items: Workers in the collection.
	Items []*Worker `json:"items,omitempty"`

	// Kind: Identifies this object as a list of workers.
	Kind string `json:"kind,omitempty"`
}

// method id "coordinate.customFieldDef.list":

type CustomFieldDefListCall struct {
	s             *Service
	teamId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of custom field definitions for a team.

func (r *CustomFieldDefService) List(teamId string) *CustomFieldDefListCall {
	return &CustomFieldDefListCall{
		s:             r.s,
		teamId:        teamId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/custom_fields",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomFieldDefListCall) Fields(s ...googleapi.Field) *CustomFieldDefListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomFieldDefListCall) Do() (*CustomFieldDefListResponse, error) {
	var returnValue *CustomFieldDefListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of custom field definitions for a team.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.customFieldDef.list",
	//   "parameterOrder": [
	//     "teamId"
	//   ],
	//   "parameters": {
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/custom_fields",
	//   "response": {
	//     "$ref": "CustomFieldDefListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}

// method id "coordinate.jobs.get":

type JobsGetCall struct {
	s             *Service
	teamId        string
	jobId         uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a job, including all the changes made to the job.

func (r *JobsService) Get(teamId string, jobId uint64) *JobsGetCall {
	return &JobsGetCall{
		s:             r.s,
		teamId:        teamId,
		jobId:         jobId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs/{jobId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsGetCall) Fields(s ...googleapi.Field) *JobsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsGetCall) Do() (*Job, error) {
	var returnValue *Job
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
		"jobId":  strconv.FormatUint(c.jobId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a job, including all the changes made to the job.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.jobs.get",
	//   "parameterOrder": [
	//     "teamId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Job number",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}

// method id "coordinate.jobs.insert":

type JobsInsertCall struct {
	s             *Service
	teamId        string
	address       string
	lat           float64
	lng           float64
	title         string
	job           *Job
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new job. Only the state field of the job should be
// set.

func (r *JobsService) Insert(teamId string, address string, lat float64, lng float64, title string, job *Job) *JobsInsertCall {
	return &JobsInsertCall{
		s:             r.s,
		teamId:        teamId,
		address:       address,
		lat:           lat,
		lng:           lng,
		title:         title,
		job:           job,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs",
	}
}

// Assignee sets the optional parameter "assignee": Assignee email
// address, or empty string to unassign.
func (c *JobsInsertCall) Assignee(assignee string) *JobsInsertCall {
	c.params_.Set("assignee", fmt.Sprintf("%v", assignee))
	return c
}

// CustomField sets the optional parameter "customField": Sets the value
// of custom fields. To set a custom field, pass the field id (from
// /team/teamId/custom_fields), a URL escaped '=' character, and the
// desired value as a parameter. For example, customField=12%3DAlice.
// Repeat the parameter for each custom field. Note that '=' cannot
// appear in the parameter value. Specifying an invalid, or inactive
// enum field will result in an error 500.
func (c *JobsInsertCall) CustomField(customField ...string) *JobsInsertCall {
	c.params_["customField"] = customField
	return c
}

// CustomerName sets the optional parameter "customerName": Customer
// name
func (c *JobsInsertCall) CustomerName(customerName string) *JobsInsertCall {
	c.params_.Set("customerName", fmt.Sprintf("%v", customerName))
	return c
}

// CustomerPhoneNumber sets the optional parameter
// "customerPhoneNumber": Customer phone number
func (c *JobsInsertCall) CustomerPhoneNumber(customerPhoneNumber string) *JobsInsertCall {
	c.params_.Set("customerPhoneNumber", fmt.Sprintf("%v", customerPhoneNumber))
	return c
}

// Note sets the optional parameter "note": Job note as newline (Unix)
// separated string
func (c *JobsInsertCall) Note(note string) *JobsInsertCall {
	c.params_.Set("note", fmt.Sprintf("%v", note))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsInsertCall) Fields(s ...googleapi.Field) *JobsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsInsertCall) Do() (*Job, error) {
	var returnValue *Job
	c.params_.Set("address", fmt.Sprintf("%v", c.address))
	c.params_.Set("lat", fmt.Sprintf("%v", c.lat))
	c.params_.Set("lng", fmt.Sprintf("%v", c.lng))
	c.params_.Set("title", fmt.Sprintf("%v", c.title))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.job,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new job. Only the state field of the job should be set.",
	//   "httpMethod": "POST",
	//   "id": "coordinate.jobs.insert",
	//   "parameterOrder": [
	//     "teamId",
	//     "address",
	//     "lat",
	//     "lng",
	//     "title"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Job address as newline (Unix) separated string",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "assignee": {
	//       "description": "Assignee email address, or empty string to unassign.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customField": {
	//       "description": "Sets the value of custom fields. To set a custom field, pass the field id (from /team/teamId/custom_fields), a URL escaped '=' character, and the desired value as a parameter. For example, customField=12%3DAlice. Repeat the parameter for each custom field. Note that '=' cannot appear in the parameter value. Specifying an invalid, or inactive enum field will result in an error 500.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "customerName": {
	//       "description": "Customer name",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerPhoneNumber": {
	//       "description": "Customer phone number",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lat": {
	//       "description": "The latitude coordinate of this job's location.",
	//       "format": "double",
	//       "location": "query",
	//       "required": true,
	//       "type": "number"
	//     },
	//     "lng": {
	//       "description": "The longitude coordinate of this job's location.",
	//       "format": "double",
	//       "location": "query",
	//       "required": true,
	//       "type": "number"
	//     },
	//     "note": {
	//       "description": "Job note as newline (Unix) separated string",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "title": {
	//       "description": "Job title",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate"
	//   ]
	// }

}

// method id "coordinate.jobs.list":

type JobsListCall struct {
	s             *Service
	teamId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves jobs created or modified since the given timestamp.

func (r *JobsService) List(teamId string) *JobsListCall {
	return &JobsListCall{
		s:             r.s,
		teamId:        teamId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return in one page.
func (c *JobsListCall) MaxResults(maxResults int64) *JobsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// MinModifiedTimestampMs sets the optional parameter
// "minModifiedTimestampMs": Minimum time a job was modified in
// milliseconds since epoch.
func (c *JobsListCall) MinModifiedTimestampMs(minModifiedTimestampMs uint64) *JobsListCall {
	c.params_.Set("minModifiedTimestampMs", fmt.Sprintf("%v", minModifiedTimestampMs))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
func (c *JobsListCall) PageToken(pageToken string) *JobsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsListCall) Fields(s ...googleapi.Field) *JobsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsListCall) Do() (*JobListResponse, error) {
	var returnValue *JobListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves jobs created or modified since the given timestamp.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.jobs.list",
	//   "parameterOrder": [
	//     "teamId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return in one page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "minModifiedTimestampMs": {
	//       "description": "Minimum time a job was modified in milliseconds since epoch.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs",
	//   "response": {
	//     "$ref": "JobListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}

// method id "coordinate.jobs.patch":

type JobsPatchCall struct {
	s             *Service
	teamId        string
	jobId         uint64
	job           *Job
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a job. Fields that are set in the job state will be
// updated. This method supports patch semantics.

func (r *JobsService) Patch(teamId string, jobId uint64, job *Job) *JobsPatchCall {
	return &JobsPatchCall{
		s:             r.s,
		teamId:        teamId,
		jobId:         jobId,
		job:           job,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs/{jobId}",
	}
}

// Address sets the optional parameter "address": Job address as newline
// (Unix) separated string
func (c *JobsPatchCall) Address(address string) *JobsPatchCall {
	c.params_.Set("address", fmt.Sprintf("%v", address))
	return c
}

// Assignee sets the optional parameter "assignee": Assignee email
// address, or empty string to unassign.
func (c *JobsPatchCall) Assignee(assignee string) *JobsPatchCall {
	c.params_.Set("assignee", fmt.Sprintf("%v", assignee))
	return c
}

// CustomField sets the optional parameter "customField": Sets the value
// of custom fields. To set a custom field, pass the field id (from
// /team/teamId/custom_fields), a URL escaped '=' character, and the
// desired value as a parameter. For example, customField=12%3DAlice.
// Repeat the parameter for each custom field. Note that '=' cannot
// appear in the parameter value. Specifying an invalid, or inactive
// enum field will result in an error 500.
func (c *JobsPatchCall) CustomField(customField ...string) *JobsPatchCall {
	c.params_["customField"] = customField
	return c
}

// CustomerName sets the optional parameter "customerName": Customer
// name
func (c *JobsPatchCall) CustomerName(customerName string) *JobsPatchCall {
	c.params_.Set("customerName", fmt.Sprintf("%v", customerName))
	return c
}

// CustomerPhoneNumber sets the optional parameter
// "customerPhoneNumber": Customer phone number
func (c *JobsPatchCall) CustomerPhoneNumber(customerPhoneNumber string) *JobsPatchCall {
	c.params_.Set("customerPhoneNumber", fmt.Sprintf("%v", customerPhoneNumber))
	return c
}

// Lat sets the optional parameter "lat": The latitude coordinate of
// this job's location.
func (c *JobsPatchCall) Lat(lat float64) *JobsPatchCall {
	c.params_.Set("lat", fmt.Sprintf("%v", lat))
	return c
}

// Lng sets the optional parameter "lng": The longitude coordinate of
// this job's location.
func (c *JobsPatchCall) Lng(lng float64) *JobsPatchCall {
	c.params_.Set("lng", fmt.Sprintf("%v", lng))
	return c
}

// Note sets the optional parameter "note": Job note as newline (Unix)
// separated string
func (c *JobsPatchCall) Note(note string) *JobsPatchCall {
	c.params_.Set("note", fmt.Sprintf("%v", note))
	return c
}

// Progress sets the optional parameter "progress": Job progress
func (c *JobsPatchCall) Progress(progress string) *JobsPatchCall {
	c.params_.Set("progress", fmt.Sprintf("%v", progress))
	return c
}

// Title sets the optional parameter "title": Job title
func (c *JobsPatchCall) Title(title string) *JobsPatchCall {
	c.params_.Set("title", fmt.Sprintf("%v", title))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsPatchCall) Fields(s ...googleapi.Field) *JobsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsPatchCall) Do() (*Job, error) {
	var returnValue *Job
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
		"jobId":  strconv.FormatUint(c.jobId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.job,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a job. Fields that are set in the job state will be updated. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "coordinate.jobs.patch",
	//   "parameterOrder": [
	//     "teamId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Job address as newline (Unix) separated string",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "assignee": {
	//       "description": "Assignee email address, or empty string to unassign.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customField": {
	//       "description": "Sets the value of custom fields. To set a custom field, pass the field id (from /team/teamId/custom_fields), a URL escaped '=' character, and the desired value as a parameter. For example, customField=12%3DAlice. Repeat the parameter for each custom field. Note that '=' cannot appear in the parameter value. Specifying an invalid, or inactive enum field will result in an error 500.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "customerName": {
	//       "description": "Customer name",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerPhoneNumber": {
	//       "description": "Customer phone number",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobId": {
	//       "description": "Job number",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "lat": {
	//       "description": "The latitude coordinate of this job's location.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "lng": {
	//       "description": "The longitude coordinate of this job's location.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "note": {
	//       "description": "Job note as newline (Unix) separated string",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "progress": {
	//       "description": "Job progress",
	//       "enum": [
	//         "COMPLETED",
	//         "IN_PROGRESS",
	//         "NOT_ACCEPTED",
	//         "NOT_STARTED",
	//         "OBSOLETE"
	//       ],
	//       "enumDescriptions": [
	//         "Completed",
	//         "In progress",
	//         "Not accepted",
	//         "Not started",
	//         "Obsolete"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "title": {
	//       "description": "Job title",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs/{jobId}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate"
	//   ]
	// }

}

// method id "coordinate.jobs.update":

type JobsUpdateCall struct {
	s             *Service
	teamId        string
	jobId         uint64
	job           *Job
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a job. Fields that are set in the job state will be
// updated.

func (r *JobsService) Update(teamId string, jobId uint64, job *Job) *JobsUpdateCall {
	return &JobsUpdateCall{
		s:             r.s,
		teamId:        teamId,
		jobId:         jobId,
		job:           job,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs/{jobId}",
	}
}

// Address sets the optional parameter "address": Job address as newline
// (Unix) separated string
func (c *JobsUpdateCall) Address(address string) *JobsUpdateCall {
	c.params_.Set("address", fmt.Sprintf("%v", address))
	return c
}

// Assignee sets the optional parameter "assignee": Assignee email
// address, or empty string to unassign.
func (c *JobsUpdateCall) Assignee(assignee string) *JobsUpdateCall {
	c.params_.Set("assignee", fmt.Sprintf("%v", assignee))
	return c
}

// CustomField sets the optional parameter "customField": Sets the value
// of custom fields. To set a custom field, pass the field id (from
// /team/teamId/custom_fields), a URL escaped '=' character, and the
// desired value as a parameter. For example, customField=12%3DAlice.
// Repeat the parameter for each custom field. Note that '=' cannot
// appear in the parameter value. Specifying an invalid, or inactive
// enum field will result in an error 500.
func (c *JobsUpdateCall) CustomField(customField ...string) *JobsUpdateCall {
	c.params_["customField"] = customField
	return c
}

// CustomerName sets the optional parameter "customerName": Customer
// name
func (c *JobsUpdateCall) CustomerName(customerName string) *JobsUpdateCall {
	c.params_.Set("customerName", fmt.Sprintf("%v", customerName))
	return c
}

// CustomerPhoneNumber sets the optional parameter
// "customerPhoneNumber": Customer phone number
func (c *JobsUpdateCall) CustomerPhoneNumber(customerPhoneNumber string) *JobsUpdateCall {
	c.params_.Set("customerPhoneNumber", fmt.Sprintf("%v", customerPhoneNumber))
	return c
}

// Lat sets the optional parameter "lat": The latitude coordinate of
// this job's location.
func (c *JobsUpdateCall) Lat(lat float64) *JobsUpdateCall {
	c.params_.Set("lat", fmt.Sprintf("%v", lat))
	return c
}

// Lng sets the optional parameter "lng": The longitude coordinate of
// this job's location.
func (c *JobsUpdateCall) Lng(lng float64) *JobsUpdateCall {
	c.params_.Set("lng", fmt.Sprintf("%v", lng))
	return c
}

// Note sets the optional parameter "note": Job note as newline (Unix)
// separated string
func (c *JobsUpdateCall) Note(note string) *JobsUpdateCall {
	c.params_.Set("note", fmt.Sprintf("%v", note))
	return c
}

// Progress sets the optional parameter "progress": Job progress
func (c *JobsUpdateCall) Progress(progress string) *JobsUpdateCall {
	c.params_.Set("progress", fmt.Sprintf("%v", progress))
	return c
}

// Title sets the optional parameter "title": Job title
func (c *JobsUpdateCall) Title(title string) *JobsUpdateCall {
	c.params_.Set("title", fmt.Sprintf("%v", title))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsUpdateCall) Fields(s ...googleapi.Field) *JobsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsUpdateCall) Do() (*Job, error) {
	var returnValue *Job
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
		"jobId":  strconv.FormatUint(c.jobId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.job,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a job. Fields that are set in the job state will be updated.",
	//   "httpMethod": "PUT",
	//   "id": "coordinate.jobs.update",
	//   "parameterOrder": [
	//     "teamId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "Job address as newline (Unix) separated string",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "assignee": {
	//       "description": "Assignee email address, or empty string to unassign.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customField": {
	//       "description": "Sets the value of custom fields. To set a custom field, pass the field id (from /team/teamId/custom_fields), a URL escaped '=' character, and the desired value as a parameter. For example, customField=12%3DAlice. Repeat the parameter for each custom field. Note that '=' cannot appear in the parameter value. Specifying an invalid, or inactive enum field will result in an error 500.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "customerName": {
	//       "description": "Customer name",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerPhoneNumber": {
	//       "description": "Customer phone number",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobId": {
	//       "description": "Job number",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "lat": {
	//       "description": "The latitude coordinate of this job's location.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "lng": {
	//       "description": "The longitude coordinate of this job's location.",
	//       "format": "double",
	//       "location": "query",
	//       "type": "number"
	//     },
	//     "note": {
	//       "description": "Job note as newline (Unix) separated string",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "progress": {
	//       "description": "Job progress",
	//       "enum": [
	//         "COMPLETED",
	//         "IN_PROGRESS",
	//         "NOT_ACCEPTED",
	//         "NOT_STARTED",
	//         "OBSOLETE"
	//       ],
	//       "enumDescriptions": [
	//         "Completed",
	//         "In progress",
	//         "Not accepted",
	//         "Not started",
	//         "Obsolete"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "title": {
	//       "description": "Job title",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs/{jobId}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate"
	//   ]
	// }

}

// method id "coordinate.location.list":

type LocationListCall struct {
	s                *Service
	teamId           string
	workerEmail      string
	startTimestampMs uint64
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
}

// List: Retrieves a list of locations for a worker.

func (r *LocationService) List(teamId string, workerEmail string, startTimestampMs uint64) *LocationListCall {
	return &LocationListCall{
		s:                r.s,
		teamId:           teamId,
		workerEmail:      workerEmail,
		startTimestampMs: startTimestampMs,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "teams/{teamId}/workers/{workerEmail}/locations",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return in one page.
func (c *LocationListCall) MaxResults(maxResults int64) *LocationListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
func (c *LocationListCall) PageToken(pageToken string) *LocationListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LocationListCall) Fields(s ...googleapi.Field) *LocationListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LocationListCall) Do() (*LocationListResponse, error) {
	var returnValue *LocationListResponse
	c.params_.Set("startTimestampMs", fmt.Sprintf("%v", c.startTimestampMs))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId":      c.teamId,
		"workerEmail": c.workerEmail,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of locations for a worker.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.location.list",
	//   "parameterOrder": [
	//     "teamId",
	//     "workerEmail",
	//     "startTimestampMs"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return in one page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTimestampMs": {
	//       "description": "Start timestamp in milliseconds since the epoch.",
	//       "format": "uint64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "workerEmail": {
	//       "description": "Worker email address.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/workers/{workerEmail}/locations",
	//   "response": {
	//     "$ref": "LocationListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}

// method id "coordinate.schedule.get":

type ScheduleGetCall struct {
	s             *Service
	teamId        string
	jobId         uint64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the schedule for a job.

func (r *ScheduleService) Get(teamId string, jobId uint64) *ScheduleGetCall {
	return &ScheduleGetCall{
		s:             r.s,
		teamId:        teamId,
		jobId:         jobId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs/{jobId}/schedule",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ScheduleGetCall) Fields(s ...googleapi.Field) *ScheduleGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ScheduleGetCall) Do() (*Schedule, error) {
	var returnValue *Schedule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
		"jobId":  strconv.FormatUint(c.jobId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the schedule for a job.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.schedule.get",
	//   "parameterOrder": [
	//     "teamId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Job number",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs/{jobId}/schedule",
	//   "response": {
	//     "$ref": "Schedule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}

// method id "coordinate.schedule.patch":

type SchedulePatchCall struct {
	s             *Service
	teamId        string
	jobId         uint64
	schedule      *Schedule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Replaces the schedule of a job with the provided schedule.
// This method supports patch semantics.

func (r *ScheduleService) Patch(teamId string, jobId uint64, schedule *Schedule) *SchedulePatchCall {
	return &SchedulePatchCall{
		s:             r.s,
		teamId:        teamId,
		jobId:         jobId,
		schedule:      schedule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs/{jobId}/schedule",
	}
}

// AllDay sets the optional parameter "allDay": Whether the job is
// scheduled for the whole day. Time of day in start/end times is
// ignored if this is true.
func (c *SchedulePatchCall) AllDay(allDay bool) *SchedulePatchCall {
	c.params_.Set("allDay", fmt.Sprintf("%v", allDay))
	return c
}

// Duration sets the optional parameter "duration": Job duration in
// milliseconds.
func (c *SchedulePatchCall) Duration(duration uint64) *SchedulePatchCall {
	c.params_.Set("duration", fmt.Sprintf("%v", duration))
	return c
}

// EndTime sets the optional parameter "endTime": Scheduled end time in
// milliseconds since epoch.
func (c *SchedulePatchCall) EndTime(endTime uint64) *SchedulePatchCall {
	c.params_.Set("endTime", fmt.Sprintf("%v", endTime))
	return c
}

// StartTime sets the optional parameter "startTime": Scheduled start
// time in milliseconds since epoch.
func (c *SchedulePatchCall) StartTime(startTime uint64) *SchedulePatchCall {
	c.params_.Set("startTime", fmt.Sprintf("%v", startTime))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SchedulePatchCall) Fields(s ...googleapi.Field) *SchedulePatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SchedulePatchCall) Do() (*Schedule, error) {
	var returnValue *Schedule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
		"jobId":  strconv.FormatUint(c.jobId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.schedule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Replaces the schedule of a job with the provided schedule. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "coordinate.schedule.patch",
	//   "parameterOrder": [
	//     "teamId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "allDay": {
	//       "description": "Whether the job is scheduled for the whole day. Time of day in start/end times is ignored if this is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "duration": {
	//       "description": "Job duration in milliseconds.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Scheduled end time in milliseconds since epoch.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobId": {
	//       "description": "Job number",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Scheduled start time in milliseconds since epoch.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs/{jobId}/schedule",
	//   "request": {
	//     "$ref": "Schedule"
	//   },
	//   "response": {
	//     "$ref": "Schedule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate"
	//   ]
	// }

}

// method id "coordinate.schedule.update":

type ScheduleUpdateCall struct {
	s             *Service
	teamId        string
	jobId         uint64
	schedule      *Schedule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Replaces the schedule of a job with the provided schedule.

func (r *ScheduleService) Update(teamId string, jobId uint64, schedule *Schedule) *ScheduleUpdateCall {
	return &ScheduleUpdateCall{
		s:             r.s,
		teamId:        teamId,
		jobId:         jobId,
		schedule:      schedule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/jobs/{jobId}/schedule",
	}
}

// AllDay sets the optional parameter "allDay": Whether the job is
// scheduled for the whole day. Time of day in start/end times is
// ignored if this is true.
func (c *ScheduleUpdateCall) AllDay(allDay bool) *ScheduleUpdateCall {
	c.params_.Set("allDay", fmt.Sprintf("%v", allDay))
	return c
}

// Duration sets the optional parameter "duration": Job duration in
// milliseconds.
func (c *ScheduleUpdateCall) Duration(duration uint64) *ScheduleUpdateCall {
	c.params_.Set("duration", fmt.Sprintf("%v", duration))
	return c
}

// EndTime sets the optional parameter "endTime": Scheduled end time in
// milliseconds since epoch.
func (c *ScheduleUpdateCall) EndTime(endTime uint64) *ScheduleUpdateCall {
	c.params_.Set("endTime", fmt.Sprintf("%v", endTime))
	return c
}

// StartTime sets the optional parameter "startTime": Scheduled start
// time in milliseconds since epoch.
func (c *ScheduleUpdateCall) StartTime(startTime uint64) *ScheduleUpdateCall {
	c.params_.Set("startTime", fmt.Sprintf("%v", startTime))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ScheduleUpdateCall) Fields(s ...googleapi.Field) *ScheduleUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ScheduleUpdateCall) Do() (*Schedule, error) {
	var returnValue *Schedule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
		"jobId":  strconv.FormatUint(c.jobId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.schedule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Replaces the schedule of a job with the provided schedule.",
	//   "httpMethod": "PUT",
	//   "id": "coordinate.schedule.update",
	//   "parameterOrder": [
	//     "teamId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "allDay": {
	//       "description": "Whether the job is scheduled for the whole day. Time of day in start/end times is ignored if this is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "duration": {
	//       "description": "Job duration in milliseconds.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endTime": {
	//       "description": "Scheduled end time in milliseconds since epoch.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobId": {
	//       "description": "Job number",
	//       "format": "uint64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Scheduled start time in milliseconds since epoch.",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/jobs/{jobId}/schedule",
	//   "request": {
	//     "$ref": "Schedule"
	//   },
	//   "response": {
	//     "$ref": "Schedule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate"
	//   ]
	// }

}

// method id "coordinate.team.list":

type TeamListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of teams for a user.

func (r *TeamService) List() *TeamListCall {
	return &TeamListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams",
	}
}

// Admin sets the optional parameter "admin": Whether to include teams
// for which the user has the Admin role.
func (c *TeamListCall) Admin(admin bool) *TeamListCall {
	c.params_.Set("admin", fmt.Sprintf("%v", admin))
	return c
}

// Dispatcher sets the optional parameter "dispatcher": Whether to
// include teams for which the user has the Dispatcher role.
func (c *TeamListCall) Dispatcher(dispatcher bool) *TeamListCall {
	c.params_.Set("dispatcher", fmt.Sprintf("%v", dispatcher))
	return c
}

// Worker sets the optional parameter "worker": Whether to include teams
// for which the user has the Worker role.
func (c *TeamListCall) Worker(worker bool) *TeamListCall {
	c.params_.Set("worker", fmt.Sprintf("%v", worker))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TeamListCall) Fields(s ...googleapi.Field) *TeamListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TeamListCall) Do() (*TeamListResponse, error) {
	var returnValue *TeamListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of teams for a user.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.team.list",
	//   "parameters": {
	//     "admin": {
	//       "description": "Whether to include teams for which the user has the Admin role.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "dispatcher": {
	//       "description": "Whether to include teams for which the user has the Dispatcher role.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "worker": {
	//       "description": "Whether to include teams for which the user has the Worker role.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "teams",
	//   "response": {
	//     "$ref": "TeamListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}

// method id "coordinate.worker.list":

type WorkerListCall struct {
	s             *Service
	teamId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of workers in a team.

func (r *WorkerService) List(teamId string) *WorkerListCall {
	return &WorkerListCall{
		s:             r.s,
		teamId:        teamId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "teams/{teamId}/workers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WorkerListCall) Fields(s ...googleapi.Field) *WorkerListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WorkerListCall) Do() (*WorkerListResponse, error) {
	var returnValue *WorkerListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"teamId": c.teamId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of workers in a team.",
	//   "httpMethod": "GET",
	//   "id": "coordinate.worker.list",
	//   "parameterOrder": [
	//     "teamId"
	//   ],
	//   "parameters": {
	//     "teamId": {
	//       "description": "Team ID",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "teams/{teamId}/workers",
	//   "response": {
	//     "$ref": "WorkerListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/coordinate",
	//     "https://www.googleapis.com/auth/coordinate.readonly"
	//   ]
	// }

}
