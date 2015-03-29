// Package logging provides access to the Google Cloud Logging API.
//
// Usage example:
//
//   import "github.com/jfcote87/api/logging/v1beta3"
//   ...
//   loggingService, err := logging.New(oauthHttpClient)
package logging

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

const apiId = "logging:v1beta3"
const apiName = "logging"
const apiVersion = "v1beta3"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "logging.googleapis.com", Path: "/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.LogEntries = NewProjectsLogEntriesService(s)
	rs.LogServices = NewProjectsLogServicesService(s)
	rs.Logs = NewProjectsLogsService(s)
	rs.Metrics = NewProjectsMetricsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	LogEntries *ProjectsLogEntriesService

	LogServices *ProjectsLogServicesService

	Logs *ProjectsLogsService

	Metrics *ProjectsMetricsService
}

func NewProjectsLogEntriesService(s *Service) *ProjectsLogEntriesService {
	rs := &ProjectsLogEntriesService{s: s}
	return rs
}

type ProjectsLogEntriesService struct {
	s *Service
}

func NewProjectsLogServicesService(s *Service) *ProjectsLogServicesService {
	rs := &ProjectsLogServicesService{s: s}
	rs.Indexes = NewProjectsLogServicesIndexesService(s)
	rs.Sinks = NewProjectsLogServicesSinksService(s)
	return rs
}

type ProjectsLogServicesService struct {
	s *Service

	Indexes *ProjectsLogServicesIndexesService

	Sinks *ProjectsLogServicesSinksService
}

func NewProjectsLogServicesIndexesService(s *Service) *ProjectsLogServicesIndexesService {
	rs := &ProjectsLogServicesIndexesService{s: s}
	return rs
}

type ProjectsLogServicesIndexesService struct {
	s *Service
}

func NewProjectsLogServicesSinksService(s *Service) *ProjectsLogServicesSinksService {
	rs := &ProjectsLogServicesSinksService{s: s}
	return rs
}

type ProjectsLogServicesSinksService struct {
	s *Service
}

func NewProjectsLogsService(s *Service) *ProjectsLogsService {
	rs := &ProjectsLogsService{s: s}
	rs.Entries = NewProjectsLogsEntriesService(s)
	rs.Sinks = NewProjectsLogsSinksService(s)
	return rs
}

type ProjectsLogsService struct {
	s *Service

	Entries *ProjectsLogsEntriesService

	Sinks *ProjectsLogsSinksService
}

func NewProjectsLogsEntriesService(s *Service) *ProjectsLogsEntriesService {
	rs := &ProjectsLogsEntriesService{s: s}
	return rs
}

type ProjectsLogsEntriesService struct {
	s *Service
}

func NewProjectsLogsSinksService(s *Service) *ProjectsLogsSinksService {
	rs := &ProjectsLogsSinksService{s: s}
	return rs
}

type ProjectsLogsSinksService struct {
	s *Service
}

func NewProjectsMetricsService(s *Service) *ProjectsMetricsService {
	rs := &ProjectsMetricsService{s: s}
	return rs
}

type ProjectsMetricsService struct {
	s *Service
}

type Empty struct {
}

type ListLogEntriesResponse struct {
	// Entries: A list of log entry resources. Fewer than page_size entries
	// may be returned, but next_page_token is the only indication of no
	// more results.
	Entries []*LogEntry `json:"entries,omitempty"`

	// NextPageToken: If there are more results, they may be retrieved by
	// invoking ListLogEntries again, supplying this next_page_token as
	// page_token in the request. If next_page_token is empty, there were no
	// further results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogMetricsResponse struct {
	// Metrics: The list of metrics returned.
	Metrics []*LogMetric `json:"metrics,omitempty"`

	// NextPageToken: If non-empty, then there are more results. This token
	// should be used in the next call to ListLogMetrics.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogServiceIndexesResponse struct {
	// NextPageToken: A token to retrieve more log service indexes. If
	// next_page_token is not empty, addition results may be retrieved by
	// calling ListLogServiceIndexes again with page_token set to this
	// value.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServiceIndexPrefixes: A list of log service index prefixes.
	ServiceIndexPrefixes []string `json:"serviceIndexPrefixes,omitempty"`
}

type ListLogServiceSinksResponse struct {
	// Sinks: These may be partial sinks where only name is populated.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogServicesResponse struct {
	// LogServices: A list of log services.
	LogServices []*LogService `json:"logServices,omitempty"`

	// NextPageToken: A token to retrieve more ServiceIndex objects. If
	// next_page_token is not empty, addition results may be retrieved by
	// calling ListLogServices again with page_token set to this value.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogSinksResponse struct {
	// Sinks: These may be partial sinks where only name is populated.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogsResponse struct {
	// Logs: A list of log resources.
	Logs []*Log `json:"logs,omitempty"`

	// NextPageToken: Pagination: If there are more results, retrieve them
	// by invoking ListLogs again with the same arguments and this
	// next_page_token.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Log struct {
	// DisplayName: Name to be used when displaying the log to the user
	// (e.g., in a UI)
	DisplayName string `json:"displayName,omitempty"`

	// Name: REQUIRED: log name
	Name string `json:"name,omitempty"`

	// PayloadType: Type URL describing the expected payload type for the
	// log.
	PayloadType string `json:"payloadType,omitempty"`
}

type LogEntry struct {
	// InsertId: Unique ID used to deduplicate log entries. If provided, the
	// logging service will attempt to reject multiple log entries on the
	// same log with the same insert_id that are sent within the previous N
	// minutes. Deduplication may occur asynchronously, so the client may
	// not receive an error if the entry is recognized as a duplicate.
	InsertId string `json:"insertId,omitempty"`

	// Log: The log resource that this entry belongs to. This is ignored by
	// WriteLogEntries (the log name is instead specified in the URL).
	Log string `json:"log,omitempty"`

	// Metadata: Metadata for the log entry.
	Metadata *LogEntryMetadata `json:"metadata,omitempty"`

	// ProtoPayload: Contains a structured (protocol buffer) log entry. If
	// the data type in proto_payload is known by the logging system, JSON
	// clients will receive it in JSON; otherwise, the client will receive
	// it as a serialized proto and must decode it. See google.protobuf.Any
	// for more details.
	ProtoPayload *LogEntryProtoPayload `json:"protoPayload,omitempty"`

	// StructPayload: Contains "JSON-like" structured data.
	StructPayload *LogEntryStructPayload `json:"structPayload,omitempty"`

	// TextPayload: Contains a text representation of the log entry.
	TextPayload string `json:"textPayload,omitempty"`
}

type LogEntryProtoPayload struct {
}

type LogEntryStructPayload struct {
}

type LogEntryMetadata struct {
	// Labels: Callers are expected to populate one of the following sets of
	// labels describing the source of the log entry. App Engine: service:
	// "appengine.googleapis.com" labels: appengine.googleapis.com/module_id
	// appengine.googleapis.com/version_id and one of:
	// appengine.googleapis.com/replica_index
	// appengine.googleapis.com/clone_id or the Compute Engine labels
	// (resource_type/resource_id) below Compute Engine: service:
	// "compute.googleapis.com" labels: compute.googleapis.com/resource_type
	// = "instance" compute.googleapis.com/resource_id
	Labels map[string]string `json:"labels,omitempty"`

	// ProjectId: If the log entry is from a Google Cloud Platform source,
	// this must be the project ID of the service that generated the entry.
	// Otherwise, the caller may populate this with whatever project name or
	// ID is meaningful, or leave it unset.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: This field is populated by the API at ingestion time
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// Region: If the log entry is from a Google Cloud Platform source, this
	// must be the region of the source (e.g., us-central1) if it has one,
	// or unset if "region" isn't meaningful for the service. For non-Google
	// Cloud Platform sources, the caller may populate this with whatever
	// location name is meaningful, or leave it unset.
	Region string `json:"region,omitempty"`

	// ServiceName: If the log entry is from a Google Cloud Platform source,
	// this must be the official API name of the service (e.g.,
	// compute.googleapis.com). Otherwise, the caller may populate this with
	// whatever service name is meaningful, or leave it unset.
	ServiceName string `json:"serviceName,omitempty"`

	// Severity: The severity of the log entry.
	Severity string `json:"severity,omitempty"`

	TimeNanos int64 `json:"timeNanos,omitempty,string"`

	// Timestamp: REQUIRED: The time the event described by the log entry
	// occurred. Timestamps must be later than January 1 1970.
	Timestamp string `json:"timestamp,omitempty"`

	// UserId: If the log entry applies to an action taken by an
	// authenticated user, this field must contain the user identifier (a
	// fully qualified email address) of the user that requested or
	// performed the action. May be empty if the event described by the log
	// entry doesn't have an associated user.
	UserId string `json:"userId,omitempty"`

	// Zone: If the log entry is from a Google Cloud Platform source, this
	// must be the zone of the source (e.g., us-central1-a) if it has one,
	// or unset if "zone" isn't meaningful for the service. For non-Google
	// Cloud Platform sources, the caller may populate this with whatever
	// location name is meaningful, or leave it unset.
	Zone string `json:"zone,omitempty"`
}

type LogError struct {
	// Resource: The resource associated with the error. It may be different
	// from the sink destination. E.g. the sink may point to a BigQuery
	// dataset, but the error may refer to a table resource inside the
	// dataset.
	Resource string `json:"resource,omitempty"`

	// Status: The description of the last error observed.
	Status *Status `json:"status,omitempty"`

	// TimeNanos: The last time the error was observed, in nanoseconds since
	// the Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`
}

type LogMetric struct {
	// Description: Description of this metric.
	Description string `json:"description,omitempty"`

	// Filter: A filter that is applied to a LogEntry that determines
	// whether the given LogEntry matches this metric.
	Filter string `json:"filter,omitempty"`

	// Name: The name of this metric. This is a user defined identifier for
	// the resource. Allowed characters include letters, numbers, and: / \ $
	// - _ . + ! * ' () %
	Name string `json:"name,omitempty"`

	// ProjectId: Project ID of the project that this metric belongs to. The
	// user does not need to set this value. It is automatically set when
	// the metric is created.
	ProjectId string `json:"projectId,omitempty"`
}

type LogService struct {
	// IndexKeys: Label keys used when labeling log entries for this
	// service. The order of the keys is significant, with higher priority
	// keys coming earlier in the list.
	IndexKeys []string `json:"indexKeys,omitempty"`

	// Name: The service's name.
	Name string `json:"name,omitempty"`
}

type LogSink struct {
	// Destination: The resource to send log entries to. The supported sink
	// resource types are: Google Cloud Storage:
	// storage.googleapis.com/{bucket} bucket.storage.googleapis.com/ Google
	// BigQuery:
	// bigquery.googleapis.com/projects/{projectId}/datasets/{datasetId}
	// Currently the logging service supports at most one sink of each type
	// per log or log service resource.
	Destination string `json:"destination,omitempty"`

	// Errors: All active errors found for this sink. [output only]
	Errors []*LogError `json:"errors,omitempty"`

	// Name: The name of this sink. This is a client-assigned identifier for
	// the resource. This is ignored by UpdateLogSink and
	// UpdateLogServicesSink.
	Name string `json:"name,omitempty"`
}

type Status struct {
	// Code: The status code, which should be an enum value of
	// [google.rpc.Code][].
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There will
	// be a common set of message types for APIs to use.
	Details []*StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. The user-facing error message should be localized and stored
	// in the [google.rpc.Status.details][google.rpc.Status.details] field.
	Message string `json:"message,omitempty"`
}

type StatusDetails struct {
}

type WriteLogEntriesRequest struct {
	// CommonLabels: Labels that apply to all entries in this request. If a
	// conflicting label key is present in the per-entry
	// LogEntryMetadata.label list, it overrides the value specified here.
	// See the documentation for LogEntryMetadata.labels for additional
	// notes.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// Entries: Log entries to insert.
	Entries []*LogEntry `json:"entries,omitempty"`
}

type WriteLogEntriesResponse struct {
}

// method id "logging.projects.logEntries.list":

type ProjectsLogEntriesListCall struct {
	s             *Service
	projectsId    string
	filter        string
	orderBy       string
	pageSize      int64
	pageToken     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists log entries in the specified project.

func (r *ProjectsLogEntriesService) List(projectsId string, filter string, orderBy string, pageSize int64, pageToken string) *ProjectsLogEntriesListCall {
	return &ProjectsLogEntriesListCall{
		s:             r.s,
		projectsId:    projectsId,
		filter:        filter,
		orderBy:       orderBy,
		pageSize:      pageSize,
		pageToken:     pageToken,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logEntries",
	}
}

// Filter sets the optional parameter "filter": Filter specification.
// Response includes only entries which are selected by this filter. If
// empty, response is unfiltered.
func (c *ProjectsLogEntriesListCall) Filter(filter string) *ProjectsLogEntriesListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort order. The only
// LogEntry field supported for sorting is metadata.timestamp and the
// default order is ascending (from older to newer). To have entries
// returned in the opposite (descending) order, specify
// order_by="metadata.timestamp desc".
func (c *ProjectsLogEntriesListCall) OrderBy(orderBy string) *ProjectsLogEntriesListCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// entries to return per request. Fewer entries may be returned.
func (c *ProjectsLogEntriesListCall) PageSize(pageSize int64) *ProjectsLogEntriesListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token
// obtained from a prior ListLogEntries response (in next_page_token).
// If this is supplied, other parameters in this request are ignored, in
// favor of the corresponding parameters in the original request.
func (c *ProjectsLogEntriesListCall) PageToken(pageToken string) *ProjectsLogEntriesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogEntriesListCall) Fields(s ...googleapi.Field) *ProjectsLogEntriesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogEntriesListCall) Do() (*ListLogEntriesResponse, error) {
	var returnValue *ListLogEntriesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists log entries in the specified project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logEntries.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "filter",
	//     "orderBy",
	//     "pageSize",
	//     "pageToken"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Filter specification. Response includes only entries which are selected by this filter. If empty, response is unfiltered.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Sort order. The only LogEntry field supported for sorting is metadata.timestamp and the default order is ascending (from older to newer). To have entries returned in the opposite (descending) order, specify order_by=\"metadata.timestamp desc\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of entries to return per request. Fewer entries may be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token obtained from a prior ListLogEntries response (in next_page_token). If this is supplied, other parameters in this request are ignored, in favor of the corresponding parameters in the original request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project resource from which to retrieve log entries, for example: \"projects/my_project_id\".",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logEntries",
	//   "response": {
	//     "$ref": "ListLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.list":

type ProjectsLogServicesListCall struct {
	s             *Service
	projectsId    string
	log           string
	pageSize      int64
	pageToken     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists log services associated with log entries ingested for a
// project.

func (r *ProjectsLogServicesService) List(projectsId string, log string, pageSize int64, pageToken string) *ProjectsLogServicesListCall {
	return &ProjectsLogServicesListCall{
		s:             r.s,
		projectsId:    projectsId,
		log:           log,
		pageSize:      pageSize,
		pageToken:     pageToken,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices",
	}
}

// Log sets the optional parameter "log": A log resource like
// /projects/project_id/logs/log_name identifying the log for which to
// list services. When empty, all services will be listed.
func (c *ProjectsLogServicesListCall) Log(log string) *ProjectsLogServicesListCall {
	c.params_.Set("log", fmt.Sprintf("%v", log))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of LogService objects to return.
func (c *ProjectsLogServicesListCall) PageSize(pageSize int64) *ProjectsLogServicesListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An optional
// next_page_token from a previous ListLogServicesResult. Other fields
// are ignored when the page_token is not empty.
func (c *ProjectsLogServicesListCall) PageToken(pageToken string) *ProjectsLogServicesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesListCall) Do() (*ListLogServicesResponse, error) {
	var returnValue *ListLogServicesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists log services associated with log entries ingested for a project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "log",
	//     "pageSize",
	//     "pageToken"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "description": "A log resource like /projects/project_id/logs/log_name identifying the log for which to list services. When empty, all services will be listed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of LogService objects to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An optional next_page_token from a previous ListLogServicesResult. Other fields are ignored when the page_token is not empty.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project resource for which to list the services.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices",
	//   "response": {
	//     "$ref": "ListLogServicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.indexes.list":

type ProjectsLogServicesIndexesListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	indexPrefix   string
	depth         int64
	log           string
	pageSize      int64
	pageToken     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists log service indexes associated with a log service.

func (r *ProjectsLogServicesIndexesService) List(projectsId string, logServicesId string, indexPrefix string, depth int64, log string, pageSize int64, pageToken string) *ProjectsLogServicesIndexesListCall {
	return &ProjectsLogServicesIndexesListCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		indexPrefix:   indexPrefix,
		depth:         depth,
		log:           log,
		pageSize:      pageSize,
		pageToken:     pageToken,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
	}
}

// Depth sets the optional parameter "depth": A limit to the number of
// levels of the index hierarchy that will be expanded. If the depth is
// 0, it will default to the level specified by the prefix field (the
// number of slash separators). The default empty prefix implies a depth
// of 1. It is an error for depth to be any non-zero value less than the
// number of components in index_prefix.
func (c *ProjectsLogServicesIndexesListCall) Depth(depth int64) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("depth", fmt.Sprintf("%v", depth))
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": A prefix of
// the log service indexes to be returned. The prefix is a slash
// separated list of the label values in order corresponding to the
// [LogService index_keys][google.logging.v1.LogService.index_keys]. For
// example use "/myModule/" to retrieve App Engine versions associated
// with myModule. The trailing slash terminates the value, while
// "/myModule" would allow retrieval of App Engine modules with names
// beginning with myModule. An prefix component matches all log service
// indexes. A non-empty prefix must begin with "/".
func (c *ProjectsLogServicesIndexesListCall) IndexPrefix(indexPrefix string) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("indexPrefix", fmt.Sprintf("%v", indexPrefix))
	return c
}

// Log sets the optional parameter "log": A log resource like
// /projects/project_id/logs/log_name identifying the log for which to
// list services indexes.
func (c *ProjectsLogServicesIndexesListCall) Log(log string) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("log", fmt.Sprintf("%v", log))
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of log service index resources to return.
func (c *ProjectsLogServicesIndexesListCall) PageSize(pageSize int64) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An optional
// next_page_token from a previous ListLogServicesIndexesResult. Other
// fields are ignored when the page_token is not empty.
func (c *ProjectsLogServicesIndexesListCall) PageToken(pageToken string) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesIndexesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesIndexesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesIndexesListCall) Do() (*ListLogServiceIndexesResponse, error) {
	var returnValue *ListLogServiceIndexesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists log service indexes associated with a log service.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.indexes.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "indexPrefix",
	//     "depth",
	//     "log",
	//     "pageSize",
	//     "pageToken"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "description": "A limit to the number of levels of the index hierarchy that will be expanded. If the depth is 0, it will default to the level specified by the prefix field (the number of slash separators). The default empty prefix implies a depth of 1. It is an error for depth to be any non-zero value less than the number of components in index_prefix.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "A prefix of the log service indexes to be returned. The prefix is a slash separated list of the label values in order corresponding to the [LogService index_keys][google.logging.v1.LogService.index_keys]. For example use \"/myModule/\" to retrieve App Engine versions associated with myModule. The trailing slash terminates the value, while \"/myModule\" would allow retrieval of App Engine modules with names beginning with myModule. An prefix component matches all log service indexes. A non-empty prefix must begin with \"/\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "log": {
	//       "description": "A log resource like /projects/project_id/logs/log_name identifying the log for which to list services indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of log service index resources to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An optional next_page_token from a previous ListLogServicesIndexesResult. Other fields are ignored when the page_token is not empty.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. A log service resource whose service indexes to return (e.g. /projects/myProj/logServices/appengine.googleapis.com).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
	//   "response": {
	//     "$ref": "ListLogServiceIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.create":

type ProjectsLogServicesSinksCreateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Create the specified log service sink resource.

func (r *ProjectsLogServicesSinksService) Create(projectsId string, logServicesId string, logsink *LogSink) *ProjectsLogServicesSinksCreateCall {
	return &ProjectsLogServicesSinksCreateCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksCreateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create the specified log service sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logServices.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The service in which to create a sink.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.delete":

type ProjectsLogServicesSinksDeleteCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified log service sink.

func (r *ProjectsLogServicesSinksService) Delete(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksDeleteCall {
	return &ProjectsLogServicesSinksDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified log service sink.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logServices.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.get":

type ProjectsLogServicesSinksGetCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get the specified log service sink resource.

func (r *ProjectsLogServicesSinksService) Get(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksGetCall {
	return &ProjectsLogServicesSinksGetCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksGetCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get the specified log service sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The sink to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.list":

type ProjectsLogServicesSinksListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List log service sinks associated with the specified service.

func (r *ProjectsLogServicesSinksService) List(projectsId string, logServicesId string) *ProjectsLogServicesSinksListCall {
	return &ProjectsLogServicesSinksListCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksListCall) Do() (*ListLogServiceSinksResponse, error) {
	var returnValue *ListLogServiceSinksResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List log service sinks associated with the specified service.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The service for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "response": {
	//     "$ref": "ListLogServiceSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.update":

type ProjectsLogServicesSinksUpdateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Create or update the specified log service sink resource.

func (r *ProjectsLogServicesSinksService) Update(projectsId string, logServicesId string, sinksId string, logsink *LogSink) *ProjectsLogServicesSinksUpdateCall {
	return &ProjectsLogServicesSinksUpdateCall{
		s:             r.s,
		projectsId:    projectsId,
		logServicesId: logServicesId,
		sinksId:       sinksId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogServicesSinksUpdateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create or update the specified log service sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logServices.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.delete":

type ProjectsLogsDeleteCall struct {
	s             *Service
	projectsId    string
	logsId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified log resource and all log entries
// contained in it.

func (r *ProjectsLogsService) Delete(projectsId string, logsId string) *ProjectsLogsDeleteCall {
	return &ProjectsLogsDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified log resource and all log entries contained in it.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log resource to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.list":

type ProjectsLogsListCall struct {
	s                  *Service
	projectsId         string
	serviceName        string
	serviceIndexPrefix string
	pageSize           int64
	pageToken          string
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// List: Lists log resources belonging to the specified project.

func (r *ProjectsLogsService) List(projectsId string, serviceName string, serviceIndexPrefix string, pageSize int64, pageToken string) *ProjectsLogsListCall {
	return &ProjectsLogsListCall{
		s:                  r.s,
		projectsId:         projectsId,
		serviceName:        serviceName,
		serviceIndexPrefix: serviceIndexPrefix,
		pageSize:           pageSize,
		pageToken:          pageToken,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "v1beta3/projects/{projectsId}/logs",
	}
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return.
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An optional
// next_page_token from a previous ListLogsResponse. Other query
// parameters are ignored when page_token is not empty.
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ServiceIndexPrefix sets the optional parameter "serviceIndexPrefix":
// A log service index prefix for which to list logs. Only logs
// containing entries whose metadata included these label values
// (associated with index keys) will be returned. The prefix is a slash
// separated list of values, and need not specify all index labels. An
// empty index (or a single slash) matches all log service indexes.
func (c *ProjectsLogsListCall) ServiceIndexPrefix(serviceIndexPrefix string) *ProjectsLogsListCall {
	c.params_.Set("serviceIndexPrefix", fmt.Sprintf("%v", serviceIndexPrefix))
	return c
}

// ServiceName sets the optional parameter "serviceName": A service name
// for which to list logs. Only logs containing entries whose metadata
// included this service name will be returned. If empty and
// service_index_prefix is also empty then all log names are returned.
// To list all log names (regardless of service) leave both the
// service_name and service_index_prefix empty. To list log names
// containing entries with a particular service name (or explicitly
// empty service name) set service_name to the desired value and
// service_index_prefix to "/".
func (c *ProjectsLogsListCall) ServiceName(serviceName string) *ProjectsLogsListCall {
	c.params_.Set("serviceName", fmt.Sprintf("%v", serviceName))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsListCall) Fields(s ...googleapi.Field) *ProjectsLogsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsListCall) Do() (*ListLogsResponse, error) {
	var returnValue *ListLogsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists log resources belonging to the specified project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "serviceName",
	//     "serviceIndexPrefix",
	//     "pageSize",
	//     "pageToken"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An optional next_page_token from a previous ListLogsResponse. Other query parameters are ignored when page_token is not empty.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project name for which to list the log resources.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "serviceIndexPrefix": {
	//       "description": "A log service index prefix for which to list logs. Only logs containing entries whose metadata included these label values (associated with index keys) will be returned. The prefix is a slash separated list of values, and need not specify all index labels. An empty index (or a single slash) matches all log service indexes.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceName": {
	//       "description": "A service name for which to list logs. Only logs containing entries whose metadata included this service name will be returned. If empty and service_index_prefix is also empty then all log names are returned. To list all log names (regardless of service) leave both the service_name and service_index_prefix empty. To list log names containing entries with a particular service name (or explicitly empty service name) set service_name to the desired value and service_index_prefix to \"/\".",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.entries.write":

type ProjectsLogsEntriesWriteCall struct {
	s                      *Service
	projectsId             string
	logsId                 string
	writelogentriesrequest *WriteLogEntriesRequest
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
}

// Write: Creates several log entries in a log.

func (r *ProjectsLogsEntriesService) Write(projectsId string, logsId string, writelogentriesrequest *WriteLogEntriesRequest) *ProjectsLogsEntriesWriteCall {
	return &ProjectsLogsEntriesWriteCall{
		s:                      r.s,
		projectsId:             projectsId,
		logsId:                 logsId,
		writelogentriesrequest: writelogentriesrequest,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsEntriesWriteCall) Fields(s ...googleapi.Field) *ProjectsLogsEntriesWriteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsEntriesWriteCall) Do() (*WriteLogEntriesResponse, error) {
	var returnValue *WriteLogEntriesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.writelogentriesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates several log entries in a log.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.entries.write",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log resource into which to insert the log entries. If the log resource name contains characters that are illegal in a URL (for example, \"/\") they must be URL-encoded (for example, as \"%2F\").",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
	//   "request": {
	//     "$ref": "WriteLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.create":

type ProjectsLogsSinksCreateCall struct {
	s             *Service
	projectsId    string
	logsId        string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Create the specified log sink resource.

func (r *ProjectsLogsSinksService) Create(projectsId string, logsId string, logsink *LogSink) *ProjectsLogsSinksCreateCall {
	return &ProjectsLogsSinksCreateCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksCreateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create the specified log sink resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log in which to create a sink",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.delete":

type ProjectsLogsSinksDeleteCall struct {
	s             *Service
	projectsId    string
	logsId        string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified log sink.

func (r *ProjectsLogsSinksService) Delete(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksDeleteCall {
	return &ProjectsLogsSinksDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified log sink.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The sink to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.get":

type ProjectsLogsSinksGetCall struct {
	s             *Service
	projectsId    string
	logsId        string
	sinksId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get the specified log sink resource.

func (r *ProjectsLogsSinksService) Get(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksGetCall {
	return &ProjectsLogsSinksGetCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		sinksId:       sinksId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksGetCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get the specified log sink resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The sink to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.list":

type ProjectsLogsSinksListCall struct {
	s             *Service
	projectsId    string
	logsId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List log sinks associated with the specified log.

func (r *ProjectsLogsSinksService) List(projectsId string, logsId string) *ProjectsLogsSinksListCall {
	return &ProjectsLogsSinksListCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksListCall) Do() (*ListLogSinksResponse, error) {
	var returnValue *ListLogSinksResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List log sinks associated with the specified log.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. The log for which to list sinks.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "response": {
	//     "$ref": "ListLogSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.update":

type ProjectsLogsSinksUpdateCall struct {
	s             *Service
	projectsId    string
	logsId        string
	sinksId       string
	logsink       *LogSink
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Create or update the specified log sink resource.

func (r *ProjectsLogsSinksService) Update(projectsId string, logsId string, sinksId string, logsink *LogSink) *ProjectsLogsSinksUpdateCall {
	return &ProjectsLogsSinksUpdateCall{
		s:             r.s,
		projectsId:    projectsId,
		logsId:        logsId,
		sinksId:       sinksId,
		logsink:       logsink,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsLogsSinksUpdateCall) Do() (*LogSink, error) {
	var returnValue *LogSink
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.logsink,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create or update the specified log sink resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logs.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The sink to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.create":

type ProjectsMetricsCreateCall struct {
	s             *Service
	projectsId    string
	logmetric     *LogMetric
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Create the specified log metric resource.

func (r *ProjectsMetricsService) Create(projectsId string, logmetric *LogMetric) *ProjectsMetricsCreateCall {
	return &ProjectsMetricsCreateCall{
		s:             r.s,
		projectsId:    projectsId,
		logmetric:     logmetric,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/metrics",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsCreateCall) Fields(s ...googleapi.Field) *ProjectsMetricsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMetricsCreateCall) Do() (*LogMetric, error) {
	var returnValue *LogMetric
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.logmetric,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create the specified log metric resource.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.metrics.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project in which to create a metric.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics",
	//   "request": {
	//     "$ref": "LogMetric"
	//   },
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.delete":

type ProjectsMetricsDeleteCall struct {
	s             *Service
	projectsId    string
	metricsId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified log metric.

func (r *ProjectsMetricsService) Delete(projectsId string, metricsId string) *ProjectsMetricsDeleteCall {
	return &ProjectsMetricsDeleteCall{
		s:             r.s,
		projectsId:    projectsId,
		metricsId:     metricsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/metrics/{metricsId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsDeleteCall) Fields(s ...googleapi.Field) *ProjectsMetricsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMetricsDeleteCall) Do() (*Empty, error) {
	var returnValue *Empty
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified log metric.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.metrics.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "metricsId"
	//   ],
	//   "parameters": {
	//     "metricsId": {
	//       "description": "Part of `metricName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `metricName`. The metric to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.get":

type ProjectsMetricsGetCall struct {
	s             *Service
	projectsId    string
	metricsId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get the specified log metric resource.

func (r *ProjectsMetricsService) Get(projectsId string, metricsId string) *ProjectsMetricsGetCall {
	return &ProjectsMetricsGetCall{
		s:             r.s,
		projectsId:    projectsId,
		metricsId:     metricsId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/metrics/{metricsId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsGetCall) Fields(s ...googleapi.Field) *ProjectsMetricsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMetricsGetCall) Do() (*LogMetric, error) {
	var returnValue *LogMetric
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get the specified log metric resource.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "metricsId"
	//   ],
	//   "parameters": {
	//     "metricsId": {
	//       "description": "Part of `metricName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `metricName`. The metric to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.list":

type ProjectsMetricsListCall struct {
	s             *Service
	projectsId    string
	pageToken     string
	pageSize      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List log metrics associated with the specified project.

func (r *ProjectsMetricsService) List(projectsId string, pageToken string, pageSize int64) *ProjectsMetricsListCall {
	return &ProjectsMetricsListCall{
		s:             r.s,
		projectsId:    projectsId,
		pageToken:     pageToken,
		pageSize:      pageSize,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/metrics",
	}
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// results to return.
func (c *ProjectsMetricsListCall) PageSize(pageSize int64) *ProjectsMetricsListCall {
	c.params_.Set("pageSize", fmt.Sprintf("%v", pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If non-empty,
// specifies from where to start returning results. This page_token is
// from a previous call to ListLogsMetrics.
func (c *ProjectsMetricsListCall) PageToken(pageToken string) *ProjectsMetricsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsListCall) Fields(s ...googleapi.Field) *ProjectsMetricsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMetricsListCall) Do() (*ListLogMetricsResponse, error) {
	var returnValue *ListLogMetricsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List log metrics associated with the specified project.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "pageToken",
	//     "pageSize"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If non-empty, specifies from where to start returning results. This page_token is from a previous call to ListLogsMetrics.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project name for which to list metrics.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics",
	//   "response": {
	//     "$ref": "ListLogMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "logging.projects.metrics.update":

type ProjectsMetricsUpdateCall struct {
	s             *Service
	projectsId    string
	metricsId     string
	logmetric     *LogMetric
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Create or update the specified log metric resource.

func (r *ProjectsMetricsService) Update(projectsId string, metricsId string, logmetric *LogMetric) *ProjectsMetricsUpdateCall {
	return &ProjectsMetricsUpdateCall{
		s:             r.s,
		projectsId:    projectsId,
		metricsId:     metricsId,
		logmetric:     logmetric,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1beta3/projects/{projectsId}/metrics/{metricsId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsUpdateCall) Fields(s ...googleapi.Field) *ProjectsMetricsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsMetricsUpdateCall) Do() (*LogMetric, error) {
	var returnValue *LogMetric
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.logmetric,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create or update the specified log metric resource.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.metrics.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "metricsId"
	//   ],
	//   "parameters": {
	//     "metricsId": {
	//       "description": "Part of `metricName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `metricName`. The metric to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
	//   "request": {
	//     "$ref": "LogMetric"
	//   },
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
