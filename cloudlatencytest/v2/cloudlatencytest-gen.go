// Package cloudlatencytest provides access to the Google Cloud Network Performance Monitoring API.
//
// Usage example:
//
//   import "github.com/jfcote87/api/cloudlatencytest/v2"
//   ...
//   cloudlatencytestService, err := cloudlatencytest.New(oauthHttpClient)
package cloudlatencytest

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

const apiId = "cloudlatencytest:v2"
const apiName = "cloudlatencytest"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "cloudlatencytest-pa.googleapis.com", Path: "/v2/statscollection/"}

// OAuth2 scopes used by this API.
const (
	// View monitoring data for all of your Google Cloud and API projects
	MonitoringReadonlyScope = "https://www.googleapis.com/auth/monitoring.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Statscollection = NewStatscollectionService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Statscollection *StatscollectionService
}

func NewStatscollectionService(s *Service) *StatscollectionService {
	rs := &StatscollectionService{s: s}
	return rs
}

type StatscollectionService struct {
	s *Service
}

type AggregatedStats struct {
	Stats []*Stats `json:"stats,omitempty"`
}

type AggregatedStatsReply struct {
	TestValue string `json:"testValue,omitempty"`
}

type DoubleValue struct {
	Label string `json:"label,omitempty"`

	Value float64 `json:"value,omitempty"`
}

type IntValue struct {
	Label string `json:"label,omitempty"`

	Value int64 `json:"value,omitempty,string"`
}

type Stats struct {
	DoubleValues []*DoubleValue `json:"doubleValues,omitempty"`

	IntValues []*IntValue `json:"intValues,omitempty"`

	StringValues []*StringValue `json:"stringValues,omitempty"`

	Time float64 `json:"time,omitempty"`
}

type StatsReply struct {
	TestValue string `json:"testValue,omitempty"`
}

type StringValue struct {
	Label string `json:"label,omitempty"`

	Value string `json:"value,omitempty"`
}

// method id "cloudlatencytest.statscollection.updateaggregatedstats":

type StatscollectionUpdateaggregatedstatsCall struct {
	s               *Service
	aggregatedstats *AggregatedStats
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Updateaggregatedstats: RPC to update the new TCP stats.

func (r *StatscollectionService) Updateaggregatedstats(aggregatedstats *AggregatedStats) *StatscollectionUpdateaggregatedstatsCall {
	return &StatscollectionUpdateaggregatedstatsCall{
		s:               r.s,
		aggregatedstats: aggregatedstats,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "updateaggregatedstats",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatscollectionUpdateaggregatedstatsCall) Fields(s ...googleapi.Field) *StatscollectionUpdateaggregatedstatsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StatscollectionUpdateaggregatedstatsCall) Do() (*AggregatedStatsReply, error) {
	var returnValue *AggregatedStatsReply
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.aggregatedstats,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "RPC to update the new TCP stats.",
	//   "httpMethod": "POST",
	//   "id": "cloudlatencytest.statscollection.updateaggregatedstats",
	//   "path": "updateaggregatedstats",
	//   "request": {
	//     "$ref": "AggregatedStats"
	//   },
	//   "response": {
	//     "$ref": "AggregatedStatsReply"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring.readonly"
	//   ]
	// }

}

// method id "cloudlatencytest.statscollection.updatestats":

type StatscollectionUpdatestatsCall struct {
	s             *Service
	stats         *Stats
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Updatestats: RPC to update the new TCP stats.

func (r *StatscollectionService) Updatestats(stats *Stats) *StatscollectionUpdatestatsCall {
	return &StatscollectionUpdatestatsCall{
		s:             r.s,
		stats:         stats,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "updatestats",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatscollectionUpdatestatsCall) Fields(s ...googleapi.Field) *StatscollectionUpdatestatsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StatscollectionUpdatestatsCall) Do() (*StatsReply, error) {
	var returnValue *StatsReply
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.stats,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "RPC to update the new TCP stats.",
	//   "httpMethod": "POST",
	//   "id": "cloudlatencytest.statscollection.updatestats",
	//   "path": "updatestats",
	//   "request": {
	//     "$ref": "Stats"
	//   },
	//   "response": {
	//     "$ref": "StatsReply"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/monitoring.readonly"
	//   ]
	// }

}
