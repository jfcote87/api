// Package webmasters provides access to the Webmaster Tools API.
//
// See https://developers.google.com/webmaster-tools/v3/welcome
//
// Usage example:
//
//   import "github.com/jfcote87/api/webmasters/v3"
//   ...
//   webmastersService, err := webmasters.New(oauthHttpClient)
package webmasters

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

const apiId = "webmasters:v3"
const apiName = "webmasters"
const apiVersion = "v3"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/webmasters/v3/"}

// OAuth2 scopes used by this API.
const (
	// View and modify Webmaster Tools data for your verified sites
	WebmastersScope = "https://www.googleapis.com/auth/webmasters"

	// View Webmaster Tools data for your verified sites
	WebmastersReadonlyScope = "https://www.googleapis.com/auth/webmasters.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Sitemaps = NewSitemapsService(s)
	s.Sites = NewSitesService(s)
	s.Urlcrawlerrorscounts = NewUrlcrawlerrorscountsService(s)
	s.Urlcrawlerrorssamples = NewUrlcrawlerrorssamplesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Sitemaps *SitemapsService

	Sites *SitesService

	Urlcrawlerrorscounts *UrlcrawlerrorscountsService

	Urlcrawlerrorssamples *UrlcrawlerrorssamplesService
}

func NewSitemapsService(s *Service) *SitemapsService {
	rs := &SitemapsService{s: s}
	return rs
}

type SitemapsService struct {
	s *Service
}

func NewSitesService(s *Service) *SitesService {
	rs := &SitesService{s: s}
	return rs
}

type SitesService struct {
	s *Service
}

func NewUrlcrawlerrorscountsService(s *Service) *UrlcrawlerrorscountsService {
	rs := &UrlcrawlerrorscountsService{s: s}
	return rs
}

type UrlcrawlerrorscountsService struct {
	s *Service
}

func NewUrlcrawlerrorssamplesService(s *Service) *UrlcrawlerrorssamplesService {
	rs := &UrlcrawlerrorssamplesService{s: s}
	return rs
}

type UrlcrawlerrorssamplesService struct {
	s *Service
}

type SitemapsListResponse struct {
	// Sitemap: Information about a sitemap entry.
	Sitemap []*WmxSitemap `json:"sitemap,omitempty"`
}

type SitesListResponse struct {
	// SiteEntry: Access level information for a Webmaster Tools site.
	SiteEntry []*WmxSite `json:"siteEntry,omitempty"`
}

type UrlCrawlErrorCount struct {
	// Count: The error count at the given timestamp.
	Count int64 `json:"count,omitempty,string"`

	// Timestamp: The time (well, date) when errors were detected, in RFC
	// 3339 format.
	Timestamp string `json:"timestamp,omitempty"`
}

type UrlCrawlErrorCountsPerType struct {
	// Category: The crawl error type.
	Category string `json:"category,omitempty"`

	// Entries: The error count entries time series.
	Entries []*UrlCrawlErrorCount `json:"entries,omitempty"`

	// Platform: Corresponding to the user agent that made the request.
	Platform string `json:"platform,omitempty"`
}

type UrlCrawlErrorsCountsQueryResponse struct {
	// CountPerTypes: The time series of the number of URL crawl errors for
	// per error category and platform.
	CountPerTypes []*UrlCrawlErrorCountsPerType `json:"countPerTypes,omitempty"`
}

type UrlCrawlErrorsSample struct {
	// First_detected: The time the error was first detected, in RFC 3339
	// format.
	First_detected string `json:"first_detected,omitempty"`

	// Last_crawled: The time when the URL was last crawled, in RFC 3339
	// format.
	Last_crawled string `json:"last_crawled,omitempty"`

	// PageUrl: The URL of an error, relative to the site.
	PageUrl string `json:"pageUrl,omitempty"`

	// ResponseCode: The HTTP response code, if any.
	ResponseCode int64 `json:"responseCode,omitempty"`

	// UrlDetails: Additional details about the URL, set only when calling
	// get().
	UrlDetails *UrlSampleDetails `json:"urlDetails,omitempty"`
}

type UrlCrawlErrorsSamplesListResponse struct {
	// UrlCrawlErrorSample: Information about the sample URL and its crawl
	// error.
	UrlCrawlErrorSample []*UrlCrawlErrorsSample `json:"urlCrawlErrorSample,omitempty"`
}

type UrlSampleDetails struct {
	// ContainingSitemaps: List of sitemaps pointing at this URL.
	ContainingSitemaps []string `json:"containingSitemaps,omitempty"`

	// LinkedFromUrls: A sample set of URLs linking to this URL.
	LinkedFromUrls []string `json:"linkedFromUrls,omitempty"`
}

type WmxSite struct {
	// PermissionLevel: The user's permission level for the site.
	PermissionLevel string `json:"permissionLevel,omitempty"`

	// SiteUrl: The URL of the site.
	SiteUrl string `json:"siteUrl,omitempty"`
}

type WmxSitemap struct {
	// Contents: The various content types in the sitemap.
	Contents []*WmxSitemapContent `json:"contents,omitempty"`

	// Errors: Number of errors in the sitemap - issues with the sitemap
	// itself, that needs to be fixed before it can be processed correctly.
	Errors int64 `json:"errors,omitempty,string"`

	// IsPending: If true, the sitemap has not been processed.
	IsPending bool `json:"isPending,omitempty"`

	// IsSitemapsIndex: If true, the sitemap is a collection of sitemaps.
	IsSitemapsIndex bool `json:"isSitemapsIndex,omitempty"`

	// LastDownloaded: Date & time in which this sitemap was last
	// downloaded. Date format is in RFC 3339 format (yyyy-mm-dd).
	LastDownloaded string `json:"lastDownloaded,omitempty"`

	// LastSubmitted: Date & time in which this sitemap was submitted. Date
	// format is in RFC 3339 format (yyyy-mm-dd).
	LastSubmitted string `json:"lastSubmitted,omitempty"`

	// Path: The url of the sitemap.
	Path string `json:"path,omitempty"`

	// Type: The type of the sitemap (for example "sitemap").
	Type string `json:"type,omitempty"`

	// Warnings: Number of warnings for the sitemap - issues with URLs in
	// the sitemaps.
	Warnings int64 `json:"warnings,omitempty,string"`
}

type WmxSitemapContent struct {
	// Indexed: The number of URLs from the sitemap that were indexed (of
	// the content type).
	Indexed int64 `json:"indexed,omitempty,string"`

	// Submitted: The number of URLs in the sitemap (of the content type).
	Submitted int64 `json:"submitted,omitempty,string"`

	// Type: The specific type of content in this sitemap (for example
	// "web", "images").
	Type string `json:"type,omitempty"`
}

// method id "webmasters.sitemaps.delete":

type SitemapsDeleteCall struct {
	s             *Service
	siteUrl       string
	feedpath      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a sitemap from this site.

func (r *SitemapsService) Delete(siteUrl string, feedpath string) *SitemapsDeleteCall {
	return &SitemapsDeleteCall{
		s:             r.s,
		siteUrl:       siteUrl,
		feedpath:      feedpath,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/sitemaps/{feedpath}",
	}
}

func (c *SitemapsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl":  c.siteUrl,
		"feedpath": c.feedpath,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a sitemap from this site.",
	//   "httpMethod": "DELETE",
	//   "id": "webmasters.sitemaps.delete",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "feedpath"
	//   ],
	//   "parameters": {
	//     "feedpath": {
	//       "description": "The URL of the actual sitemap (for example http://www.example.com/sitemap.xml).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps/{feedpath}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sitemaps.get":

type SitemapsGetCall struct {
	s             *Service
	siteUrl       string
	feedpath      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves information about a specific sitemap.

func (r *SitemapsService) Get(siteUrl string, feedpath string) *SitemapsGetCall {
	return &SitemapsGetCall{
		s:             r.s,
		siteUrl:       siteUrl,
		feedpath:      feedpath,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/sitemaps/{feedpath}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitemapsGetCall) Fields(s ...googleapi.Field) *SitemapsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitemapsGetCall) Do() (*WmxSitemap, error) {
	var returnValue *WmxSitemap
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl":  c.siteUrl,
		"feedpath": c.feedpath,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves information about a specific sitemap.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sitemaps.get",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "feedpath"
	//   ],
	//   "parameters": {
	//     "feedpath": {
	//       "description": "The URL of the actual sitemap (for example http://www.example.com/sitemap.xml).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps/{feedpath}",
	//   "response": {
	//     "$ref": "WmxSitemap"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sitemaps.list":

type SitemapsListCall struct {
	s             *Service
	siteUrl       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists sitemaps uploaded to the site.

func (r *SitemapsService) List(siteUrl string) *SitemapsListCall {
	return &SitemapsListCall{
		s:             r.s,
		siteUrl:       siteUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/sitemaps",
	}
}

// SitemapIndex sets the optional parameter "sitemapIndex": A URL of a
// site's sitemap index.
func (c *SitemapsListCall) SitemapIndex(sitemapIndex string) *SitemapsListCall {
	c.params_.Set("sitemapIndex", fmt.Sprintf("%v", sitemapIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitemapsListCall) Fields(s ...googleapi.Field) *SitemapsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitemapsListCall) Do() (*SitemapsListResponse, error) {
	var returnValue *SitemapsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists sitemaps uploaded to the site.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sitemaps.list",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sitemapIndex": {
	//       "description": "A URL of a site's sitemap index.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps",
	//   "response": {
	//     "$ref": "SitemapsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sitemaps.submit":

type SitemapsSubmitCall struct {
	s             *Service
	siteUrl       string
	feedpath      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Submit: Submits a sitemap for a site.

func (r *SitemapsService) Submit(siteUrl string, feedpath string) *SitemapsSubmitCall {
	return &SitemapsSubmitCall{
		s:             r.s,
		siteUrl:       siteUrl,
		feedpath:      feedpath,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/sitemaps/{feedpath}",
	}
}

func (c *SitemapsSubmitCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl":  c.siteUrl,
		"feedpath": c.feedpath,
	})
	call := &googleapi.Call{
		Method: "PUT",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Submits a sitemap for a site.",
	//   "httpMethod": "PUT",
	//   "id": "webmasters.sitemaps.submit",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "feedpath"
	//   ],
	//   "parameters": {
	//     "feedpath": {
	//       "description": "The URL of the sitemap to add.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/sitemaps/{feedpath}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sites.add":

type SitesAddCall struct {
	s             *Service
	siteUrl       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Add: Adds a site to the set of the user's sites in Webmaster Tools.

func (r *SitesService) Add(siteUrl string) *SitesAddCall {
	return &SitesAddCall{
		s:             r.s,
		siteUrl:       siteUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}",
	}
}

func (c *SitesAddCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
	})
	call := &googleapi.Call{
		Method: "PUT",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds a site to the set of the user's sites in Webmaster Tools.",
	//   "httpMethod": "PUT",
	//   "id": "webmasters.sites.add",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The URL of the site to add.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sites.delete":

type SitesDeleteCall struct {
	s             *Service
	siteUrl       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Removes a site from the set of the user's Webmaster Tools
// sites.

func (r *SitesService) Delete(siteUrl string) *SitesDeleteCall {
	return &SitesDeleteCall{
		s:             r.s,
		siteUrl:       siteUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}",
	}
}

func (c *SitesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Removes a site from the set of the user's Webmaster Tools sites.",
	//   "httpMethod": "DELETE",
	//   "id": "webmasters.sites.delete",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}

// method id "webmasters.sites.get":

type SitesGetCall struct {
	s             *Service
	siteUrl       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves information about specific site.

func (r *SitesService) Get(siteUrl string) *SitesGetCall {
	return &SitesGetCall{
		s:             r.s,
		siteUrl:       siteUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesGetCall) Fields(s ...googleapi.Field) *SitesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesGetCall) Do() (*WmxSite, error) {
	var returnValue *WmxSite
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves information about specific site.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sites.get",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}",
	//   "response": {
	//     "$ref": "WmxSite"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.sites.list":

type SitesListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists your Webmaster Tools sites.

func (r *SitesService) List() *SitesListCall {
	return &SitesListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesListCall) Fields(s ...googleapi.Field) *SitesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesListCall) Do() (*SitesListResponse, error) {
	var returnValue *SitesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists your Webmaster Tools sites.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.sites.list",
	//   "path": "sites",
	//   "response": {
	//     "$ref": "SitesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorscounts.query":

type UrlcrawlerrorscountsQueryCall struct {
	s             *Service
	siteUrl       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Query: Retrieves a time series of the number of URL crawl errors per
// error category and platform.

func (r *UrlcrawlerrorscountsService) Query(siteUrl string) *UrlcrawlerrorscountsQueryCall {
	return &UrlcrawlerrorscountsQueryCall{
		s:             r.s,
		siteUrl:       siteUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/urlCrawlErrorsCounts/query",
	}
}

// Category sets the optional parameter "category": The crawl error
// category, for example 'serverError'. If not specified, we return
// results for all categories.
func (c *UrlcrawlerrorscountsQueryCall) Category(category string) *UrlcrawlerrorscountsQueryCall {
	c.params_.Set("category", fmt.Sprintf("%v", category))
	return c
}

// LatestCountsOnly sets the optional parameter "latestCountsOnly": If
// true, returns only the latest crawl error counts.
func (c *UrlcrawlerrorscountsQueryCall) LatestCountsOnly(latestCountsOnly bool) *UrlcrawlerrorscountsQueryCall {
	c.params_.Set("latestCountsOnly", fmt.Sprintf("%v", latestCountsOnly))
	return c
}

// Platform sets the optional parameter "platform": The user agent type
// (platform) that made the request, for example 'web'. If not
// specified, we return results for all platforms.
func (c *UrlcrawlerrorscountsQueryCall) Platform(platform string) *UrlcrawlerrorscountsQueryCall {
	c.params_.Set("platform", fmt.Sprintf("%v", platform))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorscountsQueryCall) Fields(s ...googleapi.Field) *UrlcrawlerrorscountsQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlcrawlerrorscountsQueryCall) Do() (*UrlCrawlErrorsCountsQueryResponse, error) {
	var returnValue *UrlCrawlErrorsCountsQueryResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a time series of the number of URL crawl errors per error category and platform.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.urlcrawlerrorscounts.query",
	//   "parameterOrder": [
	//     "siteUrl"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category, for example 'serverError'. If not specified, we return results for all categories.",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "latestCountsOnly": {
	//       "default": "true",
	//       "description": "If true, returns only the latest crawl error counts.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request, for example 'web'. If not specified, we return results for all platforms.",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsCounts/query",
	//   "response": {
	//     "$ref": "UrlCrawlErrorsCountsQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorssamples.get":

type UrlcrawlerrorssamplesGetCall struct {
	s             *Service
	siteUrl       string
	url           string
	category      string
	platform      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves details about crawl errors for a site's sample URL.

func (r *UrlcrawlerrorssamplesService) Get(siteUrl string, url string, category string, platform string) *UrlcrawlerrorssamplesGetCall {
	return &UrlcrawlerrorssamplesGetCall{
		s:             r.s,
		siteUrl:       siteUrl,
		url:           url,
		category:      category,
		platform:      platform,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/urlCrawlErrorsSamples/{url}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorssamplesGetCall) Fields(s ...googleapi.Field) *UrlcrawlerrorssamplesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlcrawlerrorssamplesGetCall) Do() (*UrlCrawlErrorsSample, error) {
	var returnValue *UrlCrawlErrorsSample
	c.params_.Set("category", fmt.Sprintf("%v", c.category))
	c.params_.Set("platform", fmt.Sprintf("%v", c.platform))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
		"url":     c.url,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves details about crawl errors for a site's sample URL.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.urlcrawlerrorssamples.get",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "url",
	//     "category",
	//     "platform"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category, for example 'authPermissions'",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request, for example 'web'",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The relative path (without the site) of the sample URL; must be one of the URLs returned by list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsSamples/{url}",
	//   "response": {
	//     "$ref": "UrlCrawlErrorsSample"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorssamples.list":

type UrlcrawlerrorssamplesListCall struct {
	s             *Service
	siteUrl       string
	category      string
	platform      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists a site's sample URLs for the specified crawl error
// category and platform.

func (r *UrlcrawlerrorssamplesService) List(siteUrl string, category string, platform string) *UrlcrawlerrorssamplesListCall {
	return &UrlcrawlerrorssamplesListCall{
		s:             r.s,
		siteUrl:       siteUrl,
		category:      category,
		platform:      platform,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/urlCrawlErrorsSamples",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlcrawlerrorssamplesListCall) Fields(s ...googleapi.Field) *UrlcrawlerrorssamplesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlcrawlerrorssamplesListCall) Do() (*UrlCrawlErrorsSamplesListResponse, error) {
	var returnValue *UrlCrawlErrorsSamplesListResponse
	c.params_.Set("category", fmt.Sprintf("%v", c.category))
	c.params_.Set("platform", fmt.Sprintf("%v", c.platform))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists a site's sample URLs for the specified crawl error category and platform.",
	//   "httpMethod": "GET",
	//   "id": "webmasters.urlcrawlerrorssamples.list",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "category",
	//     "platform"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category, for example 'authPermissions'",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request, for example 'web'",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsSamples",
	//   "response": {
	//     "$ref": "UrlCrawlErrorsSamplesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters",
	//     "https://www.googleapis.com/auth/webmasters.readonly"
	//   ]
	// }

}

// method id "webmasters.urlcrawlerrorssamples.markAsFixed":

type UrlcrawlerrorssamplesMarkAsFixedCall struct {
	s             *Service
	siteUrl       string
	url           string
	category      string
	platform      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// MarkAsFixed: Marks the provided site's sample URL as fixed, and
// removes it from the samples list.

func (r *UrlcrawlerrorssamplesService) MarkAsFixed(siteUrl string, url string, category string, platform string) *UrlcrawlerrorssamplesMarkAsFixedCall {
	return &UrlcrawlerrorssamplesMarkAsFixedCall{
		s:             r.s,
		siteUrl:       siteUrl,
		url:           url,
		category:      category,
		platform:      platform,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "sites/{siteUrl}/urlCrawlErrorsSamples/{url}",
	}
}

func (c *UrlcrawlerrorssamplesMarkAsFixedCall) Do() error {
	c.params_.Set("category", fmt.Sprintf("%v", c.category))
	c.params_.Set("platform", fmt.Sprintf("%v", c.platform))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"siteUrl": c.siteUrl,
		"url":     c.url,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Marks the provided site's sample URL as fixed, and removes it from the samples list.",
	//   "httpMethod": "DELETE",
	//   "id": "webmasters.urlcrawlerrorssamples.markAsFixed",
	//   "parameterOrder": [
	//     "siteUrl",
	//     "url",
	//     "category",
	//     "platform"
	//   ],
	//   "parameters": {
	//     "category": {
	//       "description": "The crawl error category, for example 'authPermissions'",
	//       "enum": [
	//         "authPermissions",
	//         "manyToOneRedirect",
	//         "notFollowed",
	//         "notFound",
	//         "other",
	//         "roboted",
	//         "serverError",
	//         "soft404"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "platform": {
	//       "description": "The user agent type (platform) that made the request, for example 'web'",
	//       "enum": [
	//         "mobile",
	//         "smartphoneOnly",
	//         "web"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "siteUrl": {
	//       "description": "The site's URL, including protocol, for example 'http://www.example.com/'",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The relative path (without the site) of the sample URL; must be one of the URLs returned by list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "sites/{siteUrl}/urlCrawlErrorsSamples/{url}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/webmasters"
	//   ]
	// }

}
