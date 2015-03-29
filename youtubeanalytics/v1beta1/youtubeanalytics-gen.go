// Package youtubeanalytics provides access to the YouTube Analytics API.
//
// See http://developers.google.com/youtube/analytics/
//
// Usage example:
//
//   import "github.com/jfcote87/api/youtubeanalytics/v1beta1"
//   ...
//   youtubeanalyticsService, err := youtubeanalytics.New(oauthHttpClient)
package youtubeanalytics

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

const apiId = "youtubeAnalytics:v1beta1"
const apiName = "youtubeAnalytics"
const apiVersion = "v1beta1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/youtube/analytics/v1beta1/"}

// OAuth2 scopes used by this API.
const (
	// Manage your YouTube account
	YoutubeScope = "https://www.googleapis.com/auth/youtube"

	// View your YouTube account
	YoutubeReadonlyScope = "https://www.googleapis.com/auth/youtube.readonly"

	// View and manage your assets and associated content on YouTube
	YoutubepartnerScope = "https://www.googleapis.com/auth/youtubepartner"

	// View YouTube Analytics monetary reports for your YouTube content
	YtAnalyticsMonetaryReadonlyScope = "https://www.googleapis.com/auth/yt-analytics-monetary.readonly"

	// View YouTube Analytics reports for your YouTube content
	YtAnalyticsReadonlyScope = "https://www.googleapis.com/auth/yt-analytics.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.BatchReportDefinitions = NewBatchReportDefinitionsService(s)
	s.BatchReports = NewBatchReportsService(s)
	s.GroupItems = NewGroupItemsService(s)
	s.Groups = NewGroupsService(s)
	s.Reports = NewReportsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	BatchReportDefinitions *BatchReportDefinitionsService

	BatchReports *BatchReportsService

	GroupItems *GroupItemsService

	Groups *GroupsService

	Reports *ReportsService
}

func NewBatchReportDefinitionsService(s *Service) *BatchReportDefinitionsService {
	rs := &BatchReportDefinitionsService{s: s}
	return rs
}

type BatchReportDefinitionsService struct {
	s *Service
}

func NewBatchReportsService(s *Service) *BatchReportsService {
	rs := &BatchReportsService{s: s}
	return rs
}

type BatchReportsService struct {
	s *Service
}

func NewGroupItemsService(s *Service) *GroupItemsService {
	rs := &GroupItemsService{s: s}
	return rs
}

type GroupItemsService struct {
	s *Service
}

func NewGroupsService(s *Service) *GroupsService {
	rs := &GroupsService{s: s}
	return rs
}

type GroupsService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	return rs
}

type ReportsService struct {
	s *Service
}

type BatchReport struct {
	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// report.
	Id string `json:"id,omitempty"`

	// Kind: This value specifies the type of data of this item. For batch
	// report the kind property value is youtubeAnalytics#batchReport.
	Kind string `json:"kind,omitempty"`

	// Outputs: Report outputs.
	Outputs []*BatchReportOutputs `json:"outputs,omitempty"`

	// ReportId: The ID of the the report definition.
	ReportId string `json:"reportId,omitempty"`

	// TimeSpan: Period included in the report. For reports containing all
	// entities endTime is not set. Both startTime and endTime are
	// inclusive.
	TimeSpan *BatchReportTimeSpan `json:"timeSpan,omitempty"`

	// TimeUpdated: The time when the report was updated.
	TimeUpdated string `json:"timeUpdated,omitempty"`
}

type BatchReportOutputs struct {
	// DownloadUrl: Cloud storage URL to download this report. This URL is
	// valid for 30 minutes.
	DownloadUrl string `json:"downloadUrl,omitempty"`

	// Format: Format of the output.
	Format string `json:"format,omitempty"`

	// Type: Type of the output.
	Type string `json:"type,omitempty"`
}

type BatchReportTimeSpan struct {
	// EndTime: End of the period included in the report. Inclusive. For
	// reports containing all entities endTime is not set.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Start of the period included in the report. Inclusive.
	StartTime string `json:"startTime,omitempty"`
}

type BatchReportDefinition struct {
	// Id: The ID that YouTube assigns and uses to uniquely identify the
	// report definition.
	Id string `json:"id,omitempty"`

	// Kind: This value specifies the type of data of this item. For batch
	// report definition the kind property value is
	// youtubeAnalytics#batchReportDefinition.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the report definition.
	Name string `json:"name,omitempty"`

	// Status: Status of the report definition.
	Status string `json:"status,omitempty"`

	// Type: Type of the report definition.
	Type string `json:"type,omitempty"`
}

type BatchReportDefinitionList struct {
	// Items: A list of batchReportDefinition resources that match the
	// request criteria.
	Items []*BatchReportDefinition `json:"items,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response. For the list method, the kind property value is
	// youtubeAnalytics#batchReportDefinitionList.
	Kind string `json:"kind,omitempty"`
}

type BatchReportList struct {
	// Items: A list of batchReport resources that match the request
	// criteria.
	Items []*BatchReport `json:"items,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response. For the list method, the kind property value is
	// youtubeAnalytics#batchReportList.
	Kind string `json:"kind,omitempty"`
}

type Group struct {
	ContentDetails *GroupContentDetails `json:"contentDetails,omitempty"`

	Etag string `json:"etag,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Snippet *GroupSnippet `json:"snippet,omitempty"`
}

type GroupContentDetails struct {
	ItemCount uint64 `json:"itemCount,omitempty,string"`

	ItemType string `json:"itemType,omitempty"`
}

type GroupSnippet struct {
	PublishedAt string `json:"publishedAt,omitempty"`

	Title string `json:"title,omitempty"`
}

type GroupItem struct {
	Etag string `json:"etag,omitempty"`

	GroupId string `json:"groupId,omitempty"`

	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`

	Resource *GroupItemResource `json:"resource,omitempty"`
}

type GroupItemResource struct {
	Id string `json:"id,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type GroupItemListResponse struct {
	Etag string `json:"etag,omitempty"`

	Items []*GroupItem `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type GroupListResponse struct {
	Etag string `json:"etag,omitempty"`

	Items []*Group `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type ResultTable struct {
	// ColumnHeaders: This value specifies information about the data
	// returned in the rows fields. Each item in the columnHeaders list
	// identifies a field returned in the rows value, which contains a list
	// of comma-delimited data. The columnHeaders list will begin with the
	// dimensions specified in the API request, which will be followed by
	// the metrics specified in the API request. The order of both
	// dimensions and metrics will match the ordering in the API request.
	// For example, if the API request contains the parameters
	// dimensions=ageGroup,gender&metrics=viewerPercentage, the API response
	// will return columns in this order: ageGroup,gender,viewerPercentage.
	ColumnHeaders []*ResultTableColumnHeaders `json:"columnHeaders,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response. For the query method, the kind property value will be
	// youtubeAnalytics#resultTable.
	Kind string `json:"kind,omitempty"`

	// Rows: The list contains all rows of the result table. Each item in
	// the list is an array that contains comma-delimited data corresponding
	// to a single row of data. The order of the comma-delimited data fields
	// will match the order of the columns listed in the columnHeaders
	// field. If no data is available for the given query, the rows element
	// will be omitted from the response. The response for a query with the
	// day dimension will not contain rows for the most recent days.
	Rows [][]interface{} `json:"rows,omitempty"`
}

type ResultTableColumnHeaders struct {
	// ColumnType: The type of the column (DIMENSION or METRIC).
	ColumnType string `json:"columnType,omitempty"`

	// DataType: The type of the data in the column (STRING, INTEGER, FLOAT,
	// etc.).
	DataType string `json:"dataType,omitempty"`

	// Name: The name of the dimension or metric.
	Name string `json:"name,omitempty"`
}

// method id "youtubeAnalytics.batchReportDefinitions.list":

type BatchReportDefinitionsListCall struct {
	s                      *Service
	onBehalfOfContentOwner string
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
}

// List: Retrieves a list of available batch report definitions.

func (r *BatchReportDefinitionsService) List(onBehalfOfContentOwner string) *BatchReportDefinitionsListCall {
	return &BatchReportDefinitionsListCall{
		s: r.s,
		onBehalfOfContentOwner: onBehalfOfContentOwner,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "batchReportDefinitions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BatchReportDefinitionsListCall) Fields(s ...googleapi.Field) *BatchReportDefinitionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BatchReportDefinitionsListCall) Do() (*BatchReportDefinitionList, error) {
	var returnValue *BatchReportDefinitionList
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", c.onBehalfOfContentOwner))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of available batch report definitions.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.batchReportDefinitions.list",
	//   "parameterOrder": [
	//     "onBehalfOfContentOwner"
	//   ],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "batchReportDefinitions",
	//   "response": {
	//     "$ref": "BatchReportDefinitionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.batchReports.list":

type BatchReportsListCall struct {
	s                       *Service
	batchReportDefinitionId string
	onBehalfOfContentOwner  string
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// List: Retrieves a list of processed batch reports.

func (r *BatchReportsService) List(batchReportDefinitionId string, onBehalfOfContentOwner string) *BatchReportsListCall {
	return &BatchReportsListCall{
		s: r.s,
		batchReportDefinitionId: batchReportDefinitionId,
		onBehalfOfContentOwner:  onBehalfOfContentOwner,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "batchReports",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BatchReportsListCall) Fields(s ...googleapi.Field) *BatchReportsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BatchReportsListCall) Do() (*BatchReportList, error) {
	var returnValue *BatchReportList
	c.params_.Set("batchReportDefinitionId", fmt.Sprintf("%v", c.batchReportDefinitionId))
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", c.onBehalfOfContentOwner))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of processed batch reports.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.batchReports.list",
	//   "parameterOrder": [
	//     "batchReportDefinitionId",
	//     "onBehalfOfContentOwner"
	//   ],
	//   "parameters": {
	//     "batchReportDefinitionId": {
	//       "description": "The batchReportDefinitionId parameter specifies the ID of the batch reportort definition for which you are retrieving reports.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "The onBehalfOfContentOwner parameter identifies the content owner that the user is acting on behalf of.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "batchReports",
	//   "response": {
	//     "$ref": "BatchReportList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.delete":

type GroupItemsDeleteCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Removes an item from a group.

func (r *GroupItemsService) Delete(id string) *GroupItemsDeleteCall {
	return &GroupItemsDeleteCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groupItems",
	}
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupItemsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsDeleteCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

func (c *GroupItemsDeleteCall) Do() error {
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Removes an item from a group.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubeAnalytics.groupItems.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube group item ID for the group that is being deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groupItems",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.insert":

type GroupItemsInsertCall struct {
	s             *Service
	groupitem     *GroupItem
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a group item.

func (r *GroupItemsService) Insert(groupitem *GroupItem) *GroupItemsInsertCall {
	return &GroupItemsInsertCall{
		s:             r.s,
		groupitem:     groupitem,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groupItems",
	}
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupItemsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsInsertCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsInsertCall) Fields(s ...googleapi.Field) *GroupItemsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupItemsInsertCall) Do() (*GroupItem, error) {
	var returnValue *GroupItem
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.groupitem,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a group item.",
	//   "httpMethod": "POST",
	//   "id": "youtubeAnalytics.groupItems.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groupItems",
	//   "request": {
	//     "$ref": "GroupItem"
	//   },
	//   "response": {
	//     "$ref": "GroupItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.list":

type GroupItemsListCall struct {
	s             *Service
	groupId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns a collection of group items that match the API request
// parameters.

func (r *GroupItemsService) List(groupId string) *GroupItemsListCall {
	return &GroupItemsListCall{
		s:             r.s,
		groupId:       groupId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groupItems",
	}
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupItemsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsListCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsListCall) Fields(s ...googleapi.Field) *GroupItemsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupItemsListCall) Do() (*GroupItemListResponse, error) {
	var returnValue *GroupItemListResponse
	c.params_.Set("groupId", fmt.Sprintf("%v", c.groupId))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a collection of group items that match the API request parameters.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.groupItems.list",
	//   "parameterOrder": [
	//     "groupId"
	//   ],
	//   "parameters": {
	//     "groupId": {
	//       "description": "The id parameter specifies the unique ID of the group for which you want to retrieve group items.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groupItems",
	//   "response": {
	//     "$ref": "GroupItemListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.delete":

type GroupsDeleteCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a group.

func (r *GroupsService) Delete(id string) *GroupsDeleteCall {
	return &GroupsDeleteCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups",
	}
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsDeleteCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

func (c *GroupsDeleteCall) Do() error {
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a group.",
	//   "httpMethod": "DELETE",
	//   "id": "youtubeAnalytics.groups.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies the YouTube group ID for the group that is being deleted.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.insert":

type GroupsInsertCall struct {
	s             *Service
	group         *Group
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a group.

func (r *GroupsService) Insert(group *Group) *GroupsInsertCall {
	return &GroupsInsertCall{
		s:             r.s,
		group:         group,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups",
	}
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsInsertCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsInsertCall) Fields(s ...googleapi.Field) *GroupsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsInsertCall) Do() (*Group, error) {
	var returnValue *Group
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.group,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a group.",
	//   "httpMethod": "POST",
	//   "id": "youtubeAnalytics.groups.insert",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.list":

type GroupsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns a collection of groups that match the API request
// parameters. For example, you can retrieve all groups that the
// authenticated user owns, or you can retrieve one or more groups by
// their unique IDs.

func (r *GroupsService) List() *GroupsListCall {
	return &GroupsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups",
	}
}

// Id sets the optional parameter "id": The id parameter specifies a
// comma-separated list of the YouTube group ID(s) for the resource(s)
// that are being retrieved. In a group resource, the id property
// specifies the group's YouTube group ID.
func (c *GroupsListCall) Id(id string) *GroupsListCall {
	c.params_.Set("id", fmt.Sprintf("%v", id))
	return c
}

// Mine sets the optional parameter "mine": Set this parameter's value
// to true to instruct the API to only return groups owned by the
// authenticated user.
func (c *GroupsListCall) Mine(mine bool) *GroupsListCall {
	c.params_.Set("mine", fmt.Sprintf("%v", mine))
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsListCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsListCall) Fields(s ...googleapi.Field) *GroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsListCall) Do() (*GroupListResponse, error) {
	var returnValue *GroupListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a collection of groups that match the API request parameters. For example, you can retrieve all groups that the authenticated user owns, or you can retrieve one or more groups by their unique IDs.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.groups.list",
	//   "parameters": {
	//     "id": {
	//       "description": "The id parameter specifies a comma-separated list of the YouTube group ID(s) for the resource(s) that are being retrieved. In a group resource, the id property specifies the group's YouTube group ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mine": {
	//       "description": "Set this parameter's value to true to instruct the API to only return groups owned by the authenticated user.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "response": {
	//     "$ref": "GroupListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.update":

type GroupsUpdateCall struct {
	s             *Service
	group         *Group
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Modifies a group. For example, you could change a group's
// title.

func (r *GroupsService) Update(group *Group) *GroupsUpdateCall {
	return &GroupsUpdateCall{
		s:             r.s,
		group:         group,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups",
	}
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": Note: This parameter is intended
// exclusively for YouTube content partners.
//
// The onBehalfOfContentOwner
// parameter indicates that the request's authorization credentials
// identify a YouTube CMS user who is acting on behalf of the content
// owner specified in the parameter value. This parameter is intended
// for YouTube content partners that own and manage many different
// YouTube channels. It allows content owners to authenticate once and
// get access to all their video and channel data, without having to
// provide authentication credentials for each individual channel. The
// CMS account that the user authenticates with must be linked to the
// specified YouTube content owner.
func (c *GroupsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsUpdateCall {
	c.params_.Set("onBehalfOfContentOwner", fmt.Sprintf("%v", onBehalfOfContentOwner))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsUpdateCall) Fields(s ...googleapi.Field) *GroupsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsUpdateCall) Do() (*Group, error) {
	var returnValue *Group
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.group,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Modifies a group. For example, you could change a group's title.",
	//   "httpMethod": "PUT",
	//   "id": "youtubeAnalytics.groups.update",
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "Note: This parameter is intended exclusively for YouTube content partners.\n\nThe onBehalfOfContentOwner parameter indicates that the request's authorization credentials identify a YouTube CMS user who is acting on behalf of the content owner specified in the parameter value. This parameter is intended for YouTube content partners that own and manage many different YouTube channels. It allows content owners to authenticate once and get access to all their video and channel data, without having to provide authentication credentials for each individual channel. The CMS account that the user authenticates with must be linked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtubepartner"
	//   ]
	// }

}

// method id "youtubeAnalytics.reports.query":

type ReportsQueryCall struct {
	s             *Service
	ids           string
	startDate     string
	endDate       string
	metrics       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Query: Retrieve your YouTube Analytics reports.

func (r *ReportsService) Query(ids string, startDate string, endDate string, metrics string) *ReportsQueryCall {
	return &ReportsQueryCall{
		s:             r.s,
		ids:           ids,
		startDate:     startDate,
		endDate:       endDate,
		metrics:       metrics,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reports",
	}
}

// Currency sets the optional parameter "currency": The currency to
// which financial metrics should be converted. The default is US Dollar
// (USD). If the result contains no financial metrics, this flag will be
// ignored. Responds with an error if the specified currency is not
// recognized.
func (c *ReportsQueryCall) Currency(currency string) *ReportsQueryCall {
	c.params_.Set("currency", fmt.Sprintf("%v", currency))
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of YouTube Analytics dimensions, such as views
// or ageGroup,gender. See the Available Reports document for a list of
// the reports that you can retrieve and the dimensions used for those
// reports. Also see the Dimensions document for definitions of those
// dimensions.
func (c *ReportsQueryCall) Dimensions(dimensions string) *ReportsQueryCall {
	c.params_.Set("dimensions", fmt.Sprintf("%v", dimensions))
	return c
}

// Filters sets the optional parameter "filters": A list of filters that
// should be applied when retrieving YouTube Analytics data. The
// Available Reports document identifies the dimensions that can be used
// to filter each report, and the Dimensions document defines those
// dimensions. If a request uses multiple filters, join them together
// with a semicolon (;), and the returned result table will satisfy both
// filters. For example, a filters parameter value of
// video==dMH0bHeiRNg;country==IT restricts the result set to include
// data for the given video in Italy.
func (c *ReportsQueryCall) Filters(filters string) *ReportsQueryCall {
	c.params_.Set("filters", fmt.Sprintf("%v", filters))
	return c
}

// MaxResults sets the optional parameter "max-results": The maximum
// number of rows to include in the response.
func (c *ReportsQueryCall) MaxResults(maxResults int64) *ReportsQueryCall {
	c.params_.Set("max-results", fmt.Sprintf("%v", maxResults))
	return c
}

// Sort sets the optional parameter "sort": A comma-separated list of
// dimensions or metrics that determine the sort order for YouTube
// Analytics data. By default the sort order is ascending. The '-'
// prefix causes descending sort order.
func (c *ReportsQueryCall) Sort(sort string) *ReportsQueryCall {
	c.params_.Set("sort", fmt.Sprintf("%v", sort))
	return c
}

// StartIndex sets the optional parameter "start-index": An index of the
// first entity to retrieve. Use this parameter as a pagination
// mechanism along with the max-results parameter (one-based,
// inclusive).
func (c *ReportsQueryCall) StartIndex(startIndex int64) *ReportsQueryCall {
	c.params_.Set("start-index", fmt.Sprintf("%v", startIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsQueryCall) Fields(s ...googleapi.Field) *ReportsQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsQueryCall) Do() (*ResultTable, error) {
	var returnValue *ResultTable
	c.params_.Set("end-date", fmt.Sprintf("%v", c.endDate))
	c.params_.Set("ids", fmt.Sprintf("%v", c.ids))
	c.params_.Set("metrics", fmt.Sprintf("%v", c.metrics))
	c.params_.Set("start-date", fmt.Sprintf("%v", c.startDate))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve your YouTube Analytics reports.",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.reports.query",
	//   "parameterOrder": [
	//     "ids",
	//     "start-date",
	//     "end-date",
	//     "metrics"
	//   ],
	//   "parameters": {
	//     "currency": {
	//       "description": "The currency to which financial metrics should be converted. The default is US Dollar (USD). If the result contains no financial metrics, this flag will be ignored. Responds with an error if the specified currency is not recognized.",
	//       "location": "query",
	//       "pattern": "[A-Z]{3}",
	//       "type": "string"
	//     },
	//     "dimensions": {
	//       "description": "A comma-separated list of YouTube Analytics dimensions, such as views or ageGroup,gender. See the Available Reports document for a list of the reports that you can retrieve and the dimensions used for those reports. Also see the Dimensions document for definitions of those dimensions.",
	//       "location": "query",
	//       "pattern": "[0-9a-zA-Z,]+",
	//       "type": "string"
	//     },
	//     "end-date": {
	//       "description": "The end date for fetching YouTube Analytics data. The value should be in YYYY-MM-DD format.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A list of filters that should be applied when retrieving YouTube Analytics data. The Available Reports document identifies the dimensions that can be used to filter each report, and the Dimensions document defines those dimensions. If a request uses multiple filters, join them together with a semicolon (;), and the returned result table will satisfy both filters. For example, a filters parameter value of video==dMH0bHeiRNg;country==IT restricts the result set to include data for the given video in Italy.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Identifies the YouTube channel or content owner for which you are retrieving YouTube Analytics data.\n- To request data for a YouTube user, set the ids parameter value to channel==CHANNEL_ID, where CHANNEL_ID specifies the unique YouTube channel ID.\n- To request data for a YouTube CMS content owner, set the ids parameter value to contentOwner==OWNER_NAME, where OWNER_NAME is the CMS name of the content owner.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+==[a-zA-Z0-9_+-]+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "max-results": {
	//       "description": "The maximum number of rows to include in the response.",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "metrics": {
	//       "description": "A comma-separated list of YouTube Analytics metrics, such as views or likes,dislikes. See the Available Reports document for a list of the reports that you can retrieve and the metrics available in each report, and see the Metrics document for definitions of those metrics.",
	//       "location": "query",
	//       "pattern": "[0-9a-zA-Z,]+",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort order for YouTube Analytics data. By default the sort order is ascending. The '-' prefix causes descending sort order.",
	//       "location": "query",
	//       "pattern": "[-0-9a-zA-Z,]+",
	//       "type": "string"
	//     },
	//     "start-date": {
	//       "description": "The start date for fetching YouTube Analytics data. The value should be in YYYY-MM-DD format.",
	//       "location": "query",
	//       "pattern": "[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start-index": {
	//       "description": "An index of the first entity to retrieve. Use this parameter as a pagination mechanism along with the max-results parameter (one-based, inclusive).",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "reports",
	//   "response": {
	//     "$ref": "ResultTable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}
