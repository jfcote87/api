// Package cloudmonitoring provides access to the Cloud Monitoring API.
//
// See https://cloud.google.com/monitoring/v2beta2/
//
// Usage example:
//
//   import "github.com/jfcote87/api/cloudmonitoring/v2beta2"
//   ...
//   cloudmonitoringService, err := cloudmonitoring.New(oauthHttpClient)
package cloudmonitoring

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

const apiId = "cloudmonitoring:v2beta2"
const apiName = "cloudmonitoring"
const apiVersion = "v2beta2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/cloudmonitoring/v2beta2/projects/"}

// OAuth2 scopes used by this API.
const (
	// View and write monitoring data for all of your Google and third-party
	// Cloud and API projects
	MonitoringScope = "https://www.googleapis.com/auth/monitoring"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.MetricDescriptors = NewMetricDescriptorsService(s)
	s.Timeseries = NewTimeseriesService(s)
	s.TimeseriesDescriptors = NewTimeseriesDescriptorsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	MetricDescriptors *MetricDescriptorsService

	Timeseries *TimeseriesService

	TimeseriesDescriptors *TimeseriesDescriptorsService
}

func NewMetricDescriptorsService(s *Service) *MetricDescriptorsService {
	rs := &MetricDescriptorsService{s: s}
	return rs
}

type MetricDescriptorsService struct {
	s *Service
}

func NewTimeseriesService(s *Service) *TimeseriesService {
	rs := &TimeseriesService{s: s}
	return rs
}

type TimeseriesService struct {
	s *Service
}

func NewTimeseriesDescriptorsService(s *Service) *TimeseriesDescriptorsService {
	rs := &TimeseriesDescriptorsService{s: s}
	return rs
}

type TimeseriesDescriptorsService struct {
	s *Service
}

type DeleteMetricDescriptorResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#deleteMetricDescriptorResponse".
	Kind string `json:"kind,omitempty"`
}

type ListMetricDescriptorsRequest struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listMetricDescriptorsRequest".
	Kind string `json:"kind,omitempty"`
}

type ListMetricDescriptorsResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listMetricDescriptorsResponse".
	Kind string `json:"kind,omitempty"`

	// Metrics: The returned metric descriptors.
	Metrics []*MetricDescriptor `json:"metrics,omitempty"`

	// NextPageToken: Pagination token. If present, indicates that
	// additional results are available for retrieval. To access the results
	// past the pagination limit, pass this value to the pageToken query
	// parameter.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListTimeseriesDescriptorsRequest struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesDescriptorsRequest".
	Kind string `json:"kind,omitempty"`
}

type ListTimeseriesDescriptorsResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesDescriptorsResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token. If present, indicates that
	// additional results are available for retrieval. To access the results
	// past the pagination limit, set this value to the pageToken query
	// parameter.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Oldest: The oldest timestamp of the interval of this query, as an RFC
	// 3339 string.
	Oldest string `json:"oldest,omitempty"`

	// Timeseries: The returned time series descriptors.
	Timeseries []*TimeseriesDescriptor `json:"timeseries,omitempty"`

	// Youngest: The youngest timestamp of the interval of this query, as an
	// RFC 3339 string.
	Youngest string `json:"youngest,omitempty"`
}

type ListTimeseriesRequest struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesRequest".
	Kind string `json:"kind,omitempty"`
}

type ListTimeseriesResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#listTimeseriesResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token. If present, indicates that
	// additional results are available for retrieval. To access the results
	// past the pagination limit, set the pageToken query parameter to this
	// value. All of the points of a time series will be returned before
	// returning any point of the subsequent time series.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Oldest: The oldest timestamp of the interval of this query as an RFC
	// 3339 string.
	Oldest string `json:"oldest,omitempty"`

	// Timeseries: The returned time series.
	Timeseries []*Timeseries `json:"timeseries,omitempty"`

	// Youngest: The youngest timestamp of the interval of this query as an
	// RFC 3339 string.
	Youngest string `json:"youngest,omitempty"`
}

type MetricDescriptor struct {
	// Description: Description of this metric.
	Description string `json:"description,omitempty"`

	// Labels: Labels defined for this metric.
	Labels []*MetricDescriptorLabelDescriptor `json:"labels,omitempty"`

	// Name: The name of this metric.
	Name string `json:"name,omitempty"`

	// Project: The project ID to which the metric belongs.
	Project string `json:"project,omitempty"`

	// TypeDescriptor: Type description for this metric.
	TypeDescriptor *MetricDescriptorTypeDescriptor `json:"typeDescriptor,omitempty"`
}

type MetricDescriptorLabelDescriptor struct {
	// Description: Label description.
	Description string `json:"description,omitempty"`

	// Key: Label key.
	Key string `json:"key,omitempty"`
}

type MetricDescriptorTypeDescriptor struct {
	// MetricType: The method of collecting data for the metric.
	MetricType string `json:"metricType,omitempty"`

	// ValueType: The type of data that is written to a timeseries point for
	// this metric.
	ValueType string `json:"valueType,omitempty"`
}

type Point struct {
	// BoolValue: The value of this data point. Either "true" or "false".
	BoolValue bool `json:"boolValue,omitempty"`

	// DistributionValue: The value of this data point as a distribution. A
	// distribution value can contain a list of buckets and/or an
	// underflowBucket and an overflowBucket. The values of these points can
	// be used to create a histogram.
	DistributionValue *PointDistribution `json:"distributionValue,omitempty"`

	// DoubleValue: The value of this data point as a double-precision
	// floating-point number.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// End: The interval [start, end] is the time period to which the
	// point's value applies. For gauge metrics, whose values are
	// instantaneous measurements, this interval should be empty (start
	// should equal end). For cumulative metrics (of which deltas and rates
	// are special cases), the interval should be non-empty. Both start and
	// end are RFC 3339 strings.
	End string `json:"end,omitempty"`

	// Int64Value: The value of this data point as a 64-bit integer.
	Int64Value int64 `json:"int64Value,omitempty,string"`

	// Start: The interval [start, end] is the time period to which the
	// point's value applies. For gauge metrics, whose values are
	// instantaneous measurements, this interval should be empty (start
	// should equal end). For cumulative metrics (of which deltas and rates
	// are special cases), the interval should be non-empty. Both start and
	// end are RFC 3339 strings.
	Start string `json:"start,omitempty"`

	// StringValue: The value of this data point in string format.
	StringValue string `json:"stringValue,omitempty"`
}

type PointDistribution struct {
	// Buckets: The finite buckets.
	Buckets []*PointDistributionBucket `json:"buckets,omitempty"`

	// OverflowBucket: The overflow bucket.
	OverflowBucket *PointDistributionOverflowBucket `json:"overflowBucket,omitempty"`

	// UnderflowBucket: The underflow bucket.
	UnderflowBucket *PointDistributionUnderflowBucket `json:"underflowBucket,omitempty"`
}

type PointDistributionBucket struct {
	// Count: The number of events whose values are in the interval defined
	// by this bucket.
	Count int64 `json:"count,omitempty,string"`

	// LowerBound: The lower bound of the value interval of this bucket
	// (inclusive).
	LowerBound float64 `json:"lowerBound,omitempty"`

	// UpperBound: The upper bound of the value interval of this bucket
	// (exclusive).
	UpperBound float64 `json:"upperBound,omitempty"`
}

type PointDistributionOverflowBucket struct {
	// Count: The number of events whose values are in the interval defined
	// by this bucket.
	Count int64 `json:"count,omitempty,string"`

	// LowerBound: The lower bound of the value interval of this bucket
	// (inclusive).
	LowerBound float64 `json:"lowerBound,omitempty"`
}

type PointDistributionUnderflowBucket struct {
	// Count: The number of events whose values are in the interval defined
	// by this bucket.
	Count int64 `json:"count,omitempty,string"`

	// UpperBound: The upper bound of the value interval of this bucket
	// (exclusive).
	UpperBound float64 `json:"upperBound,omitempty"`
}

type Timeseries struct {
	// Points: The data points of this time series. The points are listed in
	// order of their end timestamp, from younger to older.
	Points []*Point `json:"points,omitempty"`

	// TimeseriesDesc: The descriptor of this time series.
	TimeseriesDesc *TimeseriesDescriptor `json:"timeseriesDesc,omitempty"`
}

type TimeseriesDescriptor struct {
	// Labels: The label's name.
	Labels map[string]string `json:"labels,omitempty"`

	// Metric: The name of the metric.
	Metric string `json:"metric,omitempty"`

	// Project: The Developers Console project number to which this time
	// series belongs.
	Project string `json:"project,omitempty"`
}

type TimeseriesDescriptorLabel struct {
	// Key: The label's name.
	Key string `json:"key,omitempty"`

	// Value: The label's value.
	Value string `json:"value,omitempty"`
}

type TimeseriesPoint struct {
	// Point: The data point in this time series snapshot.
	Point *Point `json:"point,omitempty"`

	// TimeseriesDesc: The descriptor of this time series.
	TimeseriesDesc *TimeseriesDescriptor `json:"timeseriesDesc,omitempty"`
}

type WriteTimeseriesRequest struct {
	// CommonLabels: The label's name.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// Timeseries: Provide time series specific labels and the data points
	// for each time series. The labels in timeseries and the common_labels
	// should form a complete list of labels that required by the metric.
	Timeseries []*TimeseriesPoint `json:"timeseries,omitempty"`
}

type WriteTimeseriesResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "cloudmonitoring#writeTimeseriesResponse".
	Kind string `json:"kind,omitempty"`
}

// method id "cloudmonitoring.metricDescriptors.create":

type MetricDescriptorsCreateCall struct {
	s                *Service
	project          string
	metricdescriptor *MetricDescriptor
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
}

// Create: Create a new metric.

func (r *MetricDescriptorsService) Create(project string, metricdescriptor *MetricDescriptor) *MetricDescriptorsCreateCall {
	return &MetricDescriptorsCreateCall{
		s:                r.s,
		project:          project,
		metricdescriptor: metricdescriptor,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "{project}/metricDescriptors",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsCreateCall) Fields(s ...googleapi.Field) *MetricDescriptorsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MetricDescriptorsCreateCall) Do() (*MetricDescriptor, error) {
	var returnValue *MetricDescriptor
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.metricdescriptor,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create a new metric.",
	//   "httpMethod": "POST",
	//   "id": "cloudmonitoring.metricDescriptors.create",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project id. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors",
	//   "request": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "response": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.metricDescriptors.delete":

type MetricDescriptorsDeleteCall struct {
	s             *Service
	project       string
	metric        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete an existing metric.

func (r *MetricDescriptorsService) Delete(project string, metric string) *MetricDescriptorsDeleteCall {
	return &MetricDescriptorsDeleteCall{
		s:             r.s,
		project:       project,
		metric:        metric,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/metricDescriptors/{metric}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsDeleteCall) Fields(s ...googleapi.Field) *MetricDescriptorsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MetricDescriptorsDeleteCall) Do() (*DeleteMetricDescriptorResponse, error) {
	var returnValue *DeleteMetricDescriptorResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete an existing metric.",
	//   "httpMethod": "DELETE",
	//   "id": "cloudmonitoring.metricDescriptors.delete",
	//   "parameterOrder": [
	//     "project",
	//     "metric"
	//   ],
	//   "parameters": {
	//     "metric": {
	//       "description": "Name of the metric.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID to which the metric belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors/{metric}",
	//   "response": {
	//     "$ref": "DeleteMetricDescriptorResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.metricDescriptors.list":

type MetricDescriptorsListCall struct {
	s                            *Service
	project                      string
	listmetricdescriptorsrequest *ListMetricDescriptorsRequest
	caller_                      googleapi.Caller
	params_                      url.Values
	pathTemplate_                string
}

// List: List metric descriptors that match the query. If the query is
// not set, then all of the metric descriptors will be returned. Large
// responses will be paginated, use the nextPageToken returned in the
// response to request subsequent pages of results by setting the
// pageToken query parameter to the value of the nextPageToken.

func (r *MetricDescriptorsService) List(project string, listmetricdescriptorsrequest *ListMetricDescriptorsRequest) *MetricDescriptorsListCall {
	return &MetricDescriptorsListCall{
		s:       r.s,
		project: project,
		listmetricdescriptorsrequest: listmetricdescriptorsrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/metricDescriptors",
	}
}

// Count sets the optional parameter "count": Maximum number of metric
// descriptors per page. Used for pagination. If not specified, count =
// 100.
func (c *MetricDescriptorsListCall) Count(count int64) *MetricDescriptorsListCall {
	c.params_.Set("count", fmt.Sprintf("%v", count))
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token, which is used to page through large result sets. Set this
// value to the value of the nextPageToken to retrieve the next page of
// results.
func (c *MetricDescriptorsListCall) PageToken(pageToken string) *MetricDescriptorsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Query sets the optional parameter "query": The query used to search
// against existing metrics. Separate keywords with a space; the service
// joins all keywords with AND, meaning that all keywords must match for
// a metric to be returned. If this field is omitted, all metrics are
// returned. If an empty string is passed with this field, no metrics
// are returned.
func (c *MetricDescriptorsListCall) Query(query string) *MetricDescriptorsListCall {
	c.params_.Set("query", fmt.Sprintf("%v", query))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetricDescriptorsListCall) Fields(s ...googleapi.Field) *MetricDescriptorsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MetricDescriptorsListCall) Do() (*ListMetricDescriptorsResponse, error) {
	var returnValue *ListMetricDescriptorsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "GET",
		URL:     u,
		Params:  c.params_,
		Payload: c.listmetricdescriptorsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List metric descriptors that match the query. If the query is not set, then all of the metric descriptors will be returned. Large responses will be paginated, use the nextPageToken returned in the response to request subsequent pages of results by setting the pageToken query parameter to the value of the nextPageToken.",
	//   "httpMethod": "GET",
	//   "id": "cloudmonitoring.metricDescriptors.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "count": {
	//       "default": "100",
	//       "description": "Maximum number of metric descriptors per page. Used for pagination. If not specified, count = 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token, which is used to page through large result sets. Set this value to the value of the nextPageToken to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project id. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "The query used to search against existing metrics. Separate keywords with a space; the service joins all keywords with AND, meaning that all keywords must match for a metric to be returned. If this field is omitted, all metrics are returned. If an empty string is passed with this field, no metrics are returned.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/metricDescriptors",
	//   "request": {
	//     "$ref": "ListMetricDescriptorsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListMetricDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.timeseries.list":

type TimeseriesListCall struct {
	s                     *Service
	project               string
	metric                string
	youngest              string
	listtimeseriesrequest *ListTimeseriesRequest
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
}

// List: List the data points of the time series that match the metric
// and labels values and that have data points in the interval. Large
// responses are paginated; use the nextPageToken returned in the
// response to request subsequent pages of results by setting the
// pageToken query parameter to the value of the nextPageToken.

func (r *TimeseriesService) List(project string, metric string, youngest string, listtimeseriesrequest *ListTimeseriesRequest) *TimeseriesListCall {
	return &TimeseriesListCall{
		s:                     r.s,
		project:               project,
		metric:                metric,
		youngest:              youngest,
		listtimeseriesrequest: listtimeseriesrequest,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "{project}/timeseries/{metric}",
	}
}

// Aggregator sets the optional parameter "aggregator": The aggregation
// function that will reduce the data points in each window to a single
// point. This parameter is only valid for non-cumulative metric types.
func (c *TimeseriesListCall) Aggregator(aggregator string) *TimeseriesListCall {
	c.params_.Set("aggregator", fmt.Sprintf("%v", aggregator))
	return c
}

// Count sets the optional parameter "count": Maximum number of data
// points per page, which is used for pagination of results.
func (c *TimeseriesListCall) Count(count int64) *TimeseriesListCall {
	c.params_.Set("count", fmt.Sprintf("%v", count))
	return c
}

// Labels sets the optional parameter "labels": A collection of labels
// for the matching time series, which are represented as:
// -
// key==value: key equals the value
// - key=~value: key regex matches the
// value
// - key!=value: key does not equal the value
// - key!~value: key
// regex does not match the value  For example, to list all of the time
// series descriptors for the region us-central1, you could
// specify:
// label=cloud.googleapis.com%2Flocation=~us-central1.*
func (c *TimeseriesListCall) Labels(labels ...string) *TimeseriesListCall {
	c.params_["labels"] = labels
	return c
}

// Oldest sets the optional parameter "oldest": Start of the time
// interval (exclusive), which is expressed as an RFC 3339 timestamp. If
// neither oldest nor timespan is specified, the default time interval
// will be (youngest - 4 hours, youngest]
func (c *TimeseriesListCall) Oldest(oldest string) *TimeseriesListCall {
	c.params_.Set("oldest", fmt.Sprintf("%v", oldest))
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token, which is used to page through large result sets. Set this
// value to the value of the nextPageToken to retrieve the next page of
// results.
func (c *TimeseriesListCall) PageToken(pageToken string) *TimeseriesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Timespan sets the optional parameter "timespan": Length of the time
// interval to query, which is an alternative way to declare the
// interval: (youngest - timespan, youngest]. The timespan and oldest
// parameters should not be used together. Units:
// - s: second
// - m:
// minute
// - h: hour
// - d: day
// - w: week  Examples: 2s, 3m, 4w. Only
// one unit is allowed, for example: 2w3d is not allowed; you should use
// 17d instead.
//
// If neither oldest nor timespan is specified, the
// default time interval will be (youngest - 4 hours, youngest].
func (c *TimeseriesListCall) Timespan(timespan string) *TimeseriesListCall {
	c.params_.Set("timespan", fmt.Sprintf("%v", timespan))
	return c
}

// Window sets the optional parameter "window": The sampling window. At
// most one data point will be returned for each window in the requested
// time interval. This parameter is only valid for non-cumulative metric
// types. Units:
// - m: minute
// - h: hour
// - d: day
// - w: week
// Examples: 3m, 4w. Only one unit is allowed, for example: 2w3d is not
// allowed; you should use 17d instead.
func (c *TimeseriesListCall) Window(window string) *TimeseriesListCall {
	c.params_.Set("window", fmt.Sprintf("%v", window))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesListCall) Fields(s ...googleapi.Field) *TimeseriesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimeseriesListCall) Do() (*ListTimeseriesResponse, error) {
	var returnValue *ListTimeseriesResponse
	c.params_.Set("youngest", fmt.Sprintf("%v", c.youngest))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	call := &googleapi.Call{
		Method:  "GET",
		URL:     u,
		Params:  c.params_,
		Payload: c.listtimeseriesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List the data points of the time series that match the metric and labels values and that have data points in the interval. Large responses are paginated; use the nextPageToken returned in the response to request subsequent pages of results by setting the pageToken query parameter to the value of the nextPageToken.",
	//   "httpMethod": "GET",
	//   "id": "cloudmonitoring.timeseries.list",
	//   "parameterOrder": [
	//     "project",
	//     "metric",
	//     "youngest"
	//   ],
	//   "parameters": {
	//     "aggregator": {
	//       "description": "The aggregation function that will reduce the data points in each window to a single point. This parameter is only valid for non-cumulative metric types.",
	//       "enum": [
	//         "max",
	//         "mean",
	//         "min",
	//         "sum"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "count": {
	//       "default": "6000",
	//       "description": "Maximum number of data points per page, which is used for pagination of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "12000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "labels": {
	//       "description": "A collection of labels for the matching time series, which are represented as:  \n- key==value: key equals the value \n- key=~value: key regex matches the value \n- key!=value: key does not equal the value \n- key!~value: key regex does not match the value  For example, to list all of the time series descriptors for the region us-central1, you could specify:\nlabel=cloud.googleapis.com%2Flocation=~us-central1.*",
	//       "location": "query",
	//       "pattern": "(.+?)(==|=~|!=|!~)(.+)",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "metric": {
	//       "description": "Metric names are protocol-free URLs as listed in the Supported Metrics page. For example, compute.googleapis.com/instance/disk/read_ops_count.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldest": {
	//       "description": "Start of the time interval (exclusive), which is expressed as an RFC 3339 timestamp. If neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest]",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token, which is used to page through large result sets. Set this value to the value of the nextPageToken to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID to which this time series belongs. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timespan": {
	//       "description": "Length of the time interval to query, which is an alternative way to declare the interval: (youngest - timespan, youngest]. The timespan and oldest parameters should not be used together. Units:  \n- s: second \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 2s, 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.\n\nIf neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest].",
	//       "location": "query",
	//       "pattern": "[0-9]+[smhdw]?",
	//       "type": "string"
	//     },
	//     "window": {
	//       "description": "The sampling window. At most one data point will be returned for each window in the requested time interval. This parameter is only valid for non-cumulative metric types. Units:  \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.",
	//       "location": "query",
	//       "pattern": "[0-9]+[mhdw]?",
	//       "type": "string"
	//     },
	//     "youngest": {
	//       "description": "End of the time interval (inclusive), which is expressed as an RFC 3339 timestamp.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseries/{metric}",
	//   "request": {
	//     "$ref": "ListTimeseriesRequest"
	//   },
	//   "response": {
	//     "$ref": "ListTimeseriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.timeseries.write":

type TimeseriesWriteCall struct {
	s                      *Service
	project                string
	writetimeseriesrequest *WriteTimeseriesRequest
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
}

// Write: Put data points to one or more time series for one or more
// metrics. If a time series does not exist, a new time series will be
// created. It is not allowed to write a time series point that is older
// than the existing youngest point of that time series. Points that are
// older than the existing youngest point of that time series will be
// discarded silently. Therefore, users should make sure that points of
// a time series are written sequentially in the order of their end
// time.

func (r *TimeseriesService) Write(project string, writetimeseriesrequest *WriteTimeseriesRequest) *TimeseriesWriteCall {
	return &TimeseriesWriteCall{
		s:                      r.s,
		project:                project,
		writetimeseriesrequest: writetimeseriesrequest,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "{project}/timeseries:write",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesWriteCall) Fields(s ...googleapi.Field) *TimeseriesWriteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimeseriesWriteCall) Do() (*WriteTimeseriesResponse, error) {
	var returnValue *WriteTimeseriesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.writetimeseriesrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Put data points to one or more time series for one or more metrics. If a time series does not exist, a new time series will be created. It is not allowed to write a time series point that is older than the existing youngest point of that time series. Points that are older than the existing youngest point of that time series will be discarded silently. Therefore, users should make sure that points of a time series are written sequentially in the order of their end time.",
	//   "httpMethod": "POST",
	//   "id": "cloudmonitoring.timeseries.write",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The project ID. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseries:write",
	//   "request": {
	//     "$ref": "WriteTimeseriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteTimeseriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "cloudmonitoring.timeseriesDescriptors.list":

type TimeseriesDescriptorsListCall struct {
	s                                *Service
	project                          string
	metric                           string
	youngest                         string
	listtimeseriesdescriptorsrequest *ListTimeseriesDescriptorsRequest
	caller_                          googleapi.Caller
	params_                          url.Values
	pathTemplate_                    string
}

// List: List the descriptors of the time series that match the metric
// and labels values and that have data points in the interval. Large
// responses are paginated; use the nextPageToken returned in the
// response to request subsequent pages of results by setting the
// pageToken query parameter to the value of the nextPageToken.

func (r *TimeseriesDescriptorsService) List(project string, metric string, youngest string, listtimeseriesdescriptorsrequest *ListTimeseriesDescriptorsRequest) *TimeseriesDescriptorsListCall {
	return &TimeseriesDescriptorsListCall{
		s:        r.s,
		project:  project,
		metric:   metric,
		youngest: youngest,
		listtimeseriesdescriptorsrequest: listtimeseriesdescriptorsrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{project}/timeseriesDescriptors/{metric}",
	}
}

// Aggregator sets the optional parameter "aggregator": The aggregation
// function that will reduce the data points in each window to a single
// point. This parameter is only valid for non-cumulative metric types.
func (c *TimeseriesDescriptorsListCall) Aggregator(aggregator string) *TimeseriesDescriptorsListCall {
	c.params_.Set("aggregator", fmt.Sprintf("%v", aggregator))
	return c
}

// Count sets the optional parameter "count": Maximum number of time
// series descriptors per page. Used for pagination. If not specified,
// count = 100.
func (c *TimeseriesDescriptorsListCall) Count(count int64) *TimeseriesDescriptorsListCall {
	c.params_.Set("count", fmt.Sprintf("%v", count))
	return c
}

// Labels sets the optional parameter "labels": A collection of labels
// for the matching time series, which are represented as:
// -
// key==value: key equals the value
// - key=~value: key regex matches the
// value
// - key!=value: key does not equal the value
// - key!~value: key
// regex does not match the value  For example, to list all of the time
// series descriptors for the region us-central1, you could
// specify:
// label=cloud.googleapis.com%2Flocation=~us-central1.*
func (c *TimeseriesDescriptorsListCall) Labels(labels ...string) *TimeseriesDescriptorsListCall {
	c.params_["labels"] = labels
	return c
}

// Oldest sets the optional parameter "oldest": Start of the time
// interval (exclusive), which is expressed as an RFC 3339 timestamp. If
// neither oldest nor timespan is specified, the default time interval
// will be (youngest - 4 hours, youngest]
func (c *TimeseriesDescriptorsListCall) Oldest(oldest string) *TimeseriesDescriptorsListCall {
	c.params_.Set("oldest", fmt.Sprintf("%v", oldest))
	return c
}

// PageToken sets the optional parameter "pageToken": The pagination
// token, which is used to page through large result sets. Set this
// value to the value of the nextPageToken to retrieve the next page of
// results.
func (c *TimeseriesDescriptorsListCall) PageToken(pageToken string) *TimeseriesDescriptorsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Timespan sets the optional parameter "timespan": Length of the time
// interval to query, which is an alternative way to declare the
// interval: (youngest - timespan, youngest]. The timespan and oldest
// parameters should not be used together. Units:
// - s: second
// - m:
// minute
// - h: hour
// - d: day
// - w: week  Examples: 2s, 3m, 4w. Only
// one unit is allowed, for example: 2w3d is not allowed; you should use
// 17d instead.
//
// If neither oldest nor timespan is specified, the
// default time interval will be (youngest - 4 hours, youngest].
func (c *TimeseriesDescriptorsListCall) Timespan(timespan string) *TimeseriesDescriptorsListCall {
	c.params_.Set("timespan", fmt.Sprintf("%v", timespan))
	return c
}

// Window sets the optional parameter "window": The sampling window. At
// most one data point will be returned for each window in the requested
// time interval. This parameter is only valid for non-cumulative metric
// types. Units:
// - m: minute
// - h: hour
// - d: day
// - w: week
// Examples: 3m, 4w. Only one unit is allowed, for example: 2w3d is not
// allowed; you should use 17d instead.
func (c *TimeseriesDescriptorsListCall) Window(window string) *TimeseriesDescriptorsListCall {
	c.params_.Set("window", fmt.Sprintf("%v", window))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimeseriesDescriptorsListCall) Fields(s ...googleapi.Field) *TimeseriesDescriptorsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimeseriesDescriptorsListCall) Do() (*ListTimeseriesDescriptorsResponse, error) {
	var returnValue *ListTimeseriesDescriptorsResponse
	c.params_.Set("youngest", fmt.Sprintf("%v", c.youngest))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"project": c.project,
		"metric":  c.metric,
	})
	call := &googleapi.Call{
		Method:  "GET",
		URL:     u,
		Params:  c.params_,
		Payload: c.listtimeseriesdescriptorsrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List the descriptors of the time series that match the metric and labels values and that have data points in the interval. Large responses are paginated; use the nextPageToken returned in the response to request subsequent pages of results by setting the pageToken query parameter to the value of the nextPageToken.",
	//   "httpMethod": "GET",
	//   "id": "cloudmonitoring.timeseriesDescriptors.list",
	//   "parameterOrder": [
	//     "project",
	//     "metric",
	//     "youngest"
	//   ],
	//   "parameters": {
	//     "aggregator": {
	//       "description": "The aggregation function that will reduce the data points in each window to a single point. This parameter is only valid for non-cumulative metric types.",
	//       "enum": [
	//         "max",
	//         "mean",
	//         "min",
	//         "sum"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "count": {
	//       "default": "100",
	//       "description": "Maximum number of time series descriptors per page. Used for pagination. If not specified, count = 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "labels": {
	//       "description": "A collection of labels for the matching time series, which are represented as:  \n- key==value: key equals the value \n- key=~value: key regex matches the value \n- key!=value: key does not equal the value \n- key!~value: key regex does not match the value  For example, to list all of the time series descriptors for the region us-central1, you could specify:\nlabel=cloud.googleapis.com%2Flocation=~us-central1.*",
	//       "location": "query",
	//       "pattern": "(.+?)(==|=~|!=|!~)(.+)",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "metric": {
	//       "description": "Metric names are protocol-free URLs as listed in the Supported Metrics page. For example, compute.googleapis.com/instance/disk/read_ops_count.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldest": {
	//       "description": "Start of the time interval (exclusive), which is expressed as an RFC 3339 timestamp. If neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest]",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The pagination token, which is used to page through large result sets. Set this value to the value of the nextPageToken to retrieve the next page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID to which this time series belongs. The value can be the numeric project ID or string-based project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "timespan": {
	//       "description": "Length of the time interval to query, which is an alternative way to declare the interval: (youngest - timespan, youngest]. The timespan and oldest parameters should not be used together. Units:  \n- s: second \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 2s, 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.\n\nIf neither oldest nor timespan is specified, the default time interval will be (youngest - 4 hours, youngest].",
	//       "location": "query",
	//       "pattern": "[0-9]+[smhdw]?",
	//       "type": "string"
	//     },
	//     "window": {
	//       "description": "The sampling window. At most one data point will be returned for each window in the requested time interval. This parameter is only valid for non-cumulative metric types. Units:  \n- m: minute \n- h: hour \n- d: day \n- w: week  Examples: 3m, 4w. Only one unit is allowed, for example: 2w3d is not allowed; you should use 17d instead.",
	//       "location": "query",
	//       "pattern": "[0-9]+[mhdw]?",
	//       "type": "string"
	//     },
	//     "youngest": {
	//       "description": "End of the time interval (inclusive), which is expressed as an RFC 3339 timestamp.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/timeseriesDescriptors/{metric}",
	//   "request": {
	//     "$ref": "ListTimeseriesDescriptorsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListTimeseriesDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}
