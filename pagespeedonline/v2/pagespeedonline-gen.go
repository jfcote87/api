// Package pagespeedonline provides access to the PageSpeed Insights API.
//
// See https://developers.google.com/speed/docs/insights/v2/getting-started
//
// Usage example:
//
//   import "github.com/jfcote87/api/pagespeedonline/v2"
//   ...
//   pagespeedonlineService, err := pagespeedonline.New(oauthHttpClient)
package pagespeedonline

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

const apiId = "pagespeedonline:v2"
const apiName = "pagespeedonline"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/pagespeedonline/v2/"}

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Pagespeedapi = NewPagespeedapiService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Pagespeedapi *PagespeedapiService
}

func NewPagespeedapiService(s *Service) *PagespeedapiService {
	rs := &PagespeedapiService{s: s}
	return rs
}

type PagespeedapiService struct {
	s *Service
}

type PagespeedApiFormatStringV2 struct {
	// Args: List of arguments for the format string.
	Args []*PagespeedApiFormatStringV2Args `json:"args,omitempty"`

	// Format: A localized format string with {{FOO}} placeholders, where
	// 'FOO' is the key of the argument whose value should be substituted.
	// For HYPERLINK arguments, the format string will instead contain
	// {{BEGIN_FOO}} and {{END_FOO}} for the argument with key 'FOO'.
	Format string `json:"format,omitempty"`
}

type PagespeedApiFormatStringV2Args struct {
	// Key: The placeholder key for this arg, as a string.
	Key string `json:"key,omitempty"`

	// Rects: The screen rectangles being referred to, with dimensions
	// measured in CSS pixels. This is only ever used for SNAPSHOT_RECT
	// arguments. If this is absent for a SNAPSHOT_RECT argument, it means
	// that that argument refers to the entire snapshot.
	Rects []*PagespeedApiFormatStringV2ArgsRects `json:"rects,omitempty"`

	// Secondary_rects: Secondary screen rectangles being referred to, with
	// dimensions measured in CSS pixels. This is only ever used for
	// SNAPSHOT_RECT arguments.
	Secondary_rects []*PagespeedApiFormatStringV2ArgsSecondary_rects `json:"secondary_rects,omitempty"`

	// Type: Type of argument. One of URL, STRING_LITERAL, INT_LITERAL,
	// BYTES, DURATION, VERBATIM_STRING, PERCENTAGE, HYPERLINK, or
	// SNAPSHOT_RECT.
	Type string `json:"type,omitempty"`

	// Value: Argument value, as a localized string.
	Value string `json:"value,omitempty"`
}

type PagespeedApiFormatStringV2ArgsRects struct {
	// Height: The height of the rect.
	Height int64 `json:"height,omitempty"`

	// Left: The left coordinate of the rect, in page coordinates.
	Left int64 `json:"left,omitempty"`

	// Top: The top coordinate of the rect, in page coordinates.
	Top int64 `json:"top,omitempty"`

	// Width: The width of the rect.
	Width int64 `json:"width,omitempty"`
}

type PagespeedApiFormatStringV2ArgsSecondary_rects struct {
	// Height: The height of the rect.
	Height int64 `json:"height,omitempty"`

	// Left: The left coordinate of the rect, in page coordinates.
	Left int64 `json:"left,omitempty"`

	// Top: The top coordinate of the rect, in page coordinates.
	Top int64 `json:"top,omitempty"`

	// Width: The width of the rect.
	Width int64 `json:"width,omitempty"`
}

type PagespeedApiImageV2 struct {
	// Data: Image data base64 encoded.
	Data string `json:"data,omitempty"`

	// Height: Height of screenshot in pixels.
	Height int64 `json:"height,omitempty"`

	// Key: Unique string key, if any, identifying this image.
	Key string `json:"key,omitempty"`

	// Mime_type: Mime type of image data (e.g. "image/jpeg").
	Mime_type string `json:"mime_type,omitempty"`

	// Page_rect: The region of the page that is captured by this image,
	// with dimensions measured in CSS pixels.
	Page_rect *PagespeedApiImageV2Page_rect `json:"page_rect,omitempty"`

	// Width: Width of screenshot in pixels.
	Width int64 `json:"width,omitempty"`
}

type PagespeedApiImageV2Page_rect struct {
	// Height: The height of the rect.
	Height int64 `json:"height,omitempty"`

	// Left: The left coordinate of the rect, in page coordinates.
	Left int64 `json:"left,omitempty"`

	// Top: The top coordinate of the rect, in page coordinates.
	Top int64 `json:"top,omitempty"`

	// Width: The width of the rect.
	Width int64 `json:"width,omitempty"`
}

type Result struct {
	// FormattedResults: Localized PageSpeed results. Contains a ruleResults
	// entry for each PageSpeed rule instantiated and run by the server.
	FormattedResults *ResultFormattedResults `json:"formattedResults,omitempty"`

	// Id: Canonicalized and final URL for the document, after following
	// page redirects (if any).
	Id string `json:"id,omitempty"`

	// InvalidRules: List of rules that were specified in the request, but
	// which the server did not know how to instantiate.
	InvalidRules []string `json:"invalidRules,omitempty"`

	// Kind: Kind of result.
	Kind string `json:"kind,omitempty"`

	// PageStats: Summary statistics for the page, such as number of
	// JavaScript bytes, number of HTML bytes, etc.
	PageStats *ResultPageStats `json:"pageStats,omitempty"`

	// ResponseCode: Response code for the document. 200 indicates a normal
	// page load. 4xx/5xx indicates an error.
	ResponseCode int64 `json:"responseCode,omitempty"`

	// RuleGroups: A map with one entry for each rule group in these
	// results.
	RuleGroups *ResultRuleGroups `json:"ruleGroups,omitempty"`

	// Screenshot: Base64-encoded screenshot of the page that was analyzed.
	Screenshot *PagespeedApiImageV2 `json:"screenshot,omitempty"`

	// Title: Title of the page, as displayed in the browser's title bar.
	Title string `json:"title,omitempty"`

	// Version: The version of PageSpeed used to generate these results.
	Version *ResultVersion `json:"version,omitempty"`
}

type ResultFormattedResults struct {
	// Locale: The locale of the formattedResults, e.g. "en_US".
	Locale string `json:"locale,omitempty"`

	// RuleResults: Dictionary of formatted rule results, with one entry for
	// each PageSpeed rule instantiated and run by the server.
	RuleResults *ResultFormattedResultsRuleResults `json:"ruleResults,omitempty"`
}

type ResultFormattedResultsRuleResults struct {
}

type ResultPageStats struct {
	// CssResponseBytes: Number of uncompressed response bytes for CSS
	// resources on the page.
	CssResponseBytes int64 `json:"cssResponseBytes,omitempty,string"`

	// FlashResponseBytes: Number of response bytes for flash resources on
	// the page.
	FlashResponseBytes int64 `json:"flashResponseBytes,omitempty,string"`

	// HtmlResponseBytes: Number of uncompressed response bytes for the main
	// HTML document and all iframes on the page.
	HtmlResponseBytes int64 `json:"htmlResponseBytes,omitempty,string"`

	// ImageResponseBytes: Number of response bytes for image resources on
	// the page.
	ImageResponseBytes int64 `json:"imageResponseBytes,omitempty,string"`

	// JavascriptResponseBytes: Number of uncompressed response bytes for JS
	// resources on the page.
	JavascriptResponseBytes int64 `json:"javascriptResponseBytes,omitempty,string"`

	// NumberCssResources: Number of CSS resources referenced by the page.
	NumberCssResources int64 `json:"numberCssResources,omitempty"`

	// NumberHosts: Number of unique hosts referenced by the page.
	NumberHosts int64 `json:"numberHosts,omitempty"`

	// NumberJsResources: Number of JavaScript resources referenced by the
	// page.
	NumberJsResources int64 `json:"numberJsResources,omitempty"`

	// NumberResources: Number of HTTP resources loaded by the page.
	NumberResources int64 `json:"numberResources,omitempty"`

	// NumberStaticResources: Number of static (i.e. cacheable) resources on
	// the page.
	NumberStaticResources int64 `json:"numberStaticResources,omitempty"`

	// OtherResponseBytes: Number of response bytes for other resources on
	// the page.
	OtherResponseBytes int64 `json:"otherResponseBytes,omitempty,string"`

	// TextResponseBytes: Number of uncompressed response bytes for text
	// resources not covered by other statistics (i.e non-HTML, non-script,
	// non-CSS resources) on the page.
	TextResponseBytes int64 `json:"textResponseBytes,omitempty,string"`

	// TotalRequestBytes: Total size of all request bytes sent by the page.
	TotalRequestBytes int64 `json:"totalRequestBytes,omitempty,string"`
}

type ResultRuleGroups struct {
}

type ResultVersion struct {
	// Major: The major version number of PageSpeed used to generate these
	// results.
	Major int64 `json:"major,omitempty"`

	// Minor: The minor version number of PageSpeed used to generate these
	// results.
	Minor int64 `json:"minor,omitempty"`
}

// method id "pagespeedonline.pagespeedapi.runpagespeed":

type PagespeedapiRunpagespeedCall struct {
	s             *Service
	url           string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Runpagespeed: Runs PageSpeed analysis on the page at the specified
// URL, and returns PageSpeed scores, a list of suggestions to make that
// page faster, and other information.

func (r *PagespeedapiService) Runpagespeed(url string) *PagespeedapiRunpagespeedCall {
	return &PagespeedapiRunpagespeedCall{
		s:             r.s,
		url:           url,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "runPagespeed",
	}
}

// Filter_third_party_resources sets the optional parameter
// "filter_third_party_resources": Indicates if third party resources
// should be filtered out before PageSpeed analysis.
func (c *PagespeedapiRunpagespeedCall) Filter_third_party_resources(filter_third_party_resources bool) *PagespeedapiRunpagespeedCall {
	c.params_.Set("filter_third_party_resources", fmt.Sprintf("%v", filter_third_party_resources))
	return c
}

// Locale sets the optional parameter "locale": The locale used to
// localize formatted results
func (c *PagespeedapiRunpagespeedCall) Locale(locale string) *PagespeedapiRunpagespeedCall {
	c.params_.Set("locale", fmt.Sprintf("%v", locale))
	return c
}

// Rule sets the optional parameter "rule": A PageSpeed rule to run; if
// none are given, all rules are run
func (c *PagespeedapiRunpagespeedCall) Rule(rule ...string) *PagespeedapiRunpagespeedCall {
	c.params_["rule"] = rule
	return c
}

// Screenshot sets the optional parameter "screenshot": Indicates if
// binary data containing a screenshot should be included
func (c *PagespeedapiRunpagespeedCall) Screenshot(screenshot bool) *PagespeedapiRunpagespeedCall {
	c.params_.Set("screenshot", fmt.Sprintf("%v", screenshot))
	return c
}

// Strategy sets the optional parameter "strategy": The analysis
// strategy to use
func (c *PagespeedapiRunpagespeedCall) Strategy(strategy string) *PagespeedapiRunpagespeedCall {
	c.params_.Set("strategy", fmt.Sprintf("%v", strategy))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PagespeedapiRunpagespeedCall) Fields(s ...googleapi.Field) *PagespeedapiRunpagespeedCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PagespeedapiRunpagespeedCall) Do() (*Result, error) {
	var returnValue *Result
	c.params_.Set("url", fmt.Sprintf("%v", c.url))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Runs PageSpeed analysis on the page at the specified URL, and returns PageSpeed scores, a list of suggestions to make that page faster, and other information.",
	//   "httpMethod": "GET",
	//   "id": "pagespeedonline.pagespeedapi.runpagespeed",
	//   "parameterOrder": [
	//     "url"
	//   ],
	//   "parameters": {
	//     "filter_third_party_resources": {
	//       "default": "false",
	//       "description": "Indicates if third party resources should be filtered out before PageSpeed analysis.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "locale": {
	//       "description": "The locale used to localize formatted results",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+(_[a-zA-Z]+)?",
	//       "type": "string"
	//     },
	//     "rule": {
	//       "description": "A PageSpeed rule to run; if none are given, all rules are run",
	//       "location": "query",
	//       "pattern": "[a-zA-Z]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "screenshot": {
	//       "default": "false",
	//       "description": "Indicates if binary data containing a screenshot should be included",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "strategy": {
	//       "description": "The analysis strategy to use",
	//       "enum": [
	//         "desktop",
	//         "mobile"
	//       ],
	//       "enumDescriptions": [
	//         "Fetch and analyze the URL for desktop browsers",
	//         "Fetch and analyze the URL for mobile devices"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "url": {
	//       "description": "The URL to fetch and analyze",
	//       "location": "query",
	//       "pattern": "http(s)?://.*",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "runPagespeed",
	//   "response": {
	//     "$ref": "Result"
	//   }
	// }

}
