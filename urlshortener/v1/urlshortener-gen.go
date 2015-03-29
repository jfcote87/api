// Package urlshortener provides access to the URL Shortener API.
//
// See https://developers.google.com/url-shortener/v1/getting_started
//
// Usage example:
//
//   import "github.com/jfcote87/api/urlshortener/v1"
//   ...
//   urlshortenerService, err := urlshortener.New(oauthHttpClient)
package urlshortener

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

const apiId = "urlshortener:v1"
const apiName = "urlshortener"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/urlshortener/v1/"}

// OAuth2 scopes used by this API.
const (
	// Manage your goo.gl short URLs
	UrlshortenerScope = "https://www.googleapis.com/auth/urlshortener"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Url = NewUrlService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Url *UrlService
}

func NewUrlService(s *Service) *UrlService {
	rs := &UrlService{s: s}
	return rs
}

type UrlService struct {
	s *Service
}

type AnalyticsSnapshot struct {
	// Browsers: Top browsers, e.g. "Chrome"; sorted by (descending) click
	// counts. Only present if this data is available.
	Browsers []*StringCount `json:"browsers,omitempty"`

	// Countries: Top countries (expressed as country codes), e.g. "US" or
	// "DE"; sorted by (descending) click counts. Only present if this data
	// is available.
	Countries []*StringCount `json:"countries,omitempty"`

	// LongUrlClicks: Number of clicks on all goo.gl short URLs pointing to
	// this long URL.
	LongUrlClicks int64 `json:"longUrlClicks,omitempty,string"`

	// Platforms: Top platforms or OSes, e.g. "Windows"; sorted by
	// (descending) click counts. Only present if this data is available.
	Platforms []*StringCount `json:"platforms,omitempty"`

	// Referrers: Top referring hosts, e.g. "www.google.com"; sorted by
	// (descending) click counts. Only present if this data is available.
	Referrers []*StringCount `json:"referrers,omitempty"`

	// ShortUrlClicks: Number of clicks on this short URL.
	ShortUrlClicks int64 `json:"shortUrlClicks,omitempty,string"`
}

type AnalyticsSummary struct {
	// AllTime: Click analytics over all time.
	AllTime *AnalyticsSnapshot `json:"allTime,omitempty"`

	// Day: Click analytics over the last day.
	Day *AnalyticsSnapshot `json:"day,omitempty"`

	// Month: Click analytics over the last month.
	Month *AnalyticsSnapshot `json:"month,omitempty"`

	// TwoHours: Click analytics over the last two hours.
	TwoHours *AnalyticsSnapshot `json:"twoHours,omitempty"`

	// Week: Click analytics over the last week.
	Week *AnalyticsSnapshot `json:"week,omitempty"`
}

type StringCount struct {
	// Count: Number of clicks for this top entry, e.g. for this particular
	// country or browser.
	Count int64 `json:"count,omitempty,string"`

	// Id: Label assigned to this top entry, e.g. "US" or "Chrome".
	Id string `json:"id,omitempty"`
}

type Url struct {
	// Analytics: A summary of the click analytics for the short and long
	// URL. Might not be present if not requested or currently unavailable.
	Analytics *AnalyticsSummary `json:"analytics,omitempty"`

	// Created: Time the short URL was created; ISO 8601 representation
	// using the yyyy-MM-dd'T'HH:mm:ss.SSSZZ format, e.g.
	// "2010-10-14T19:01:24.944+00:00".
	Created string `json:"created,omitempty"`

	// Id: Short URL, e.g. "http://goo.gl/l6MS".
	Id string `json:"id,omitempty"`

	// Kind: The fixed string "urlshortener#url".
	Kind string `json:"kind,omitempty"`

	// LongUrl: Long URL, e.g. "http://www.google.com/". Might not be
	// present if the status is "REMOVED".
	LongUrl string `json:"longUrl,omitempty"`

	// Status: Status of the target URL. Possible values: "OK", "MALWARE",
	// "PHISHING", or "REMOVED". A URL might be marked "REMOVED" if it was
	// flagged as spam, for example.
	Status string `json:"status,omitempty"`
}

type UrlHistory struct {
	// Items: A list of URL resources.
	Items []*Url `json:"items,omitempty"`

	// ItemsPerPage: Number of items returned with each full "page" of
	// results. Note that the last page could have fewer items than the
	// "itemsPerPage" value.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`

	// Kind: The fixed string "urlshortener#urlHistory".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to provide to get the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of short URLs associated with this user (may
	// be approximate).
	TotalItems int64 `json:"totalItems,omitempty"`
}

// method id "urlshortener.url.get":

type UrlGetCall struct {
	s             *Service
	shortUrl      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Expands a short URL or gets creation time and analytics.

func (r *UrlService) Get(shortUrl string) *UrlGetCall {
	return &UrlGetCall{
		s:             r.s,
		shortUrl:      shortUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "url",
	}
}

// Projection sets the optional parameter "projection": Additional
// information to return.
func (c *UrlGetCall) Projection(projection string) *UrlGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlGetCall) Fields(s ...googleapi.Field) *UrlGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlGetCall) Do() (*Url, error) {
	var returnValue *Url
	c.params_.Set("shortUrl", fmt.Sprintf("%v", c.shortUrl))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Expands a short URL or gets creation time and analytics.",
	//   "httpMethod": "GET",
	//   "id": "urlshortener.url.get",
	//   "parameterOrder": [
	//     "shortUrl"
	//   ],
	//   "parameters": {
	//     "projection": {
	//       "description": "Additional information to return.",
	//       "enum": [
	//         "ANALYTICS_CLICKS",
	//         "ANALYTICS_TOP_STRINGS",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Returns only click counts.",
	//         "Returns only top string counts.",
	//         "Returns the creation timestamp and all available analytics."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "shortUrl": {
	//       "description": "The short URL, including the protocol.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "url",
	//   "response": {
	//     "$ref": "Url"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/urlshortener"
	//   ]
	// }

}

// method id "urlshortener.url.insert":

type UrlInsertCall struct {
	s             *Service
	url           *Url
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new short URL.

func (r *UrlService) Insert(url *Url) *UrlInsertCall {
	return &UrlInsertCall{
		s:             r.s,
		url:           url,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "url",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlInsertCall) Fields(s ...googleapi.Field) *UrlInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlInsertCall) Do() (*Url, error) {
	var returnValue *Url
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.url,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new short URL.",
	//   "httpMethod": "POST",
	//   "id": "urlshortener.url.insert",
	//   "path": "url",
	//   "request": {
	//     "$ref": "Url"
	//   },
	//   "response": {
	//     "$ref": "Url"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/urlshortener"
	//   ]
	// }

}

// method id "urlshortener.url.list":

type UrlListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of URLs shortened by a user.

func (r *UrlService) List() *UrlListCall {
	return &UrlListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "url/history",
	}
}

// Projection sets the optional parameter "projection": Additional
// information to return.
func (c *UrlListCall) Projection(projection string) *UrlListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// StartToken sets the optional parameter "start-token": Token for
// requesting successive pages of results.
func (c *UrlListCall) StartToken(startToken string) *UrlListCall {
	c.params_.Set("start-token", fmt.Sprintf("%v", startToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlListCall) Fields(s ...googleapi.Field) *UrlListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlListCall) Do() (*UrlHistory, error) {
	var returnValue *UrlHistory
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of URLs shortened by a user.",
	//   "httpMethod": "GET",
	//   "id": "urlshortener.url.list",
	//   "parameters": {
	//     "projection": {
	//       "description": "Additional information to return.",
	//       "enum": [
	//         "ANALYTICS_CLICKS",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Returns short URL click counts.",
	//         "Returns short URL click counts."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "start-token": {
	//       "description": "Token for requesting successive pages of results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "url/history",
	//   "response": {
	//     "$ref": "UrlHistory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/urlshortener"
	//   ]
	// }

}
