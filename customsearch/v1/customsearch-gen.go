// Package customsearch provides access to the CustomSearch API.
//
// See https://developers.google.com/custom-search/v1/using_rest
//
// Usage example:
//
//   import "github.com/jfcote87/api/customsearch/v1"
//   ...
//   customsearchService, err := customsearch.New(oauthHttpClient)
package customsearch

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

const apiId = "customsearch:v1"
const apiName = "customsearch"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/customsearch/"}

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Cse = NewCseService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Cse *CseService
}

func NewCseService(s *Service) *CseService {
	rs := &CseService{s: s}
	return rs
}

type CseService struct {
	s *Service
}

type Context struct {
	Facets [][]*ContextFacetsItem `json:"facets,omitempty"`

	Title string `json:"title,omitempty"`
}

type ContextFacetsItem struct {
	Anchor string `json:"anchor,omitempty"`

	Label string `json:"label,omitempty"`

	Label_with_op string `json:"label_with_op,omitempty"`
}

type Promotion struct {
	BodyLines []*PromotionBodyLines `json:"bodyLines,omitempty"`

	DisplayLink string `json:"displayLink,omitempty"`

	HtmlTitle string `json:"htmlTitle,omitempty"`

	Image *PromotionImage `json:"image,omitempty"`

	Link string `json:"link,omitempty"`

	Title string `json:"title,omitempty"`
}

type PromotionBodyLines struct {
	HtmlTitle string `json:"htmlTitle,omitempty"`

	Link string `json:"link,omitempty"`

	Title string `json:"title,omitempty"`

	Url string `json:"url,omitempty"`
}

type PromotionImage struct {
	Height int64 `json:"height,omitempty"`

	Source string `json:"source,omitempty"`

	Width int64 `json:"width,omitempty"`
}

type Query struct {
	Count int64 `json:"count,omitempty"`

	Cr string `json:"cr,omitempty"`

	Cref string `json:"cref,omitempty"`

	Cx string `json:"cx,omitempty"`

	DateRestrict string `json:"dateRestrict,omitempty"`

	DisableCnTwTranslation string `json:"disableCnTwTranslation,omitempty"`

	ExactTerms string `json:"exactTerms,omitempty"`

	ExcludeTerms string `json:"excludeTerms,omitempty"`

	FileType string `json:"fileType,omitempty"`

	Filter string `json:"filter,omitempty"`

	Gl string `json:"gl,omitempty"`

	GoogleHost string `json:"googleHost,omitempty"`

	HighRange string `json:"highRange,omitempty"`

	Hl string `json:"hl,omitempty"`

	Hq string `json:"hq,omitempty"`

	ImgColorType string `json:"imgColorType,omitempty"`

	ImgDominantColor string `json:"imgDominantColor,omitempty"`

	ImgSize string `json:"imgSize,omitempty"`

	ImgType string `json:"imgType,omitempty"`

	InputEncoding string `json:"inputEncoding,omitempty"`

	Language string `json:"language,omitempty"`

	LinkSite string `json:"linkSite,omitempty"`

	LowRange string `json:"lowRange,omitempty"`

	OrTerms string `json:"orTerms,omitempty"`

	OutputEncoding string `json:"outputEncoding,omitempty"`

	RelatedSite string `json:"relatedSite,omitempty"`

	Rights string `json:"rights,omitempty"`

	Safe string `json:"safe,omitempty"`

	SearchTerms string `json:"searchTerms,omitempty"`

	SearchType string `json:"searchType,omitempty"`

	SiteSearch string `json:"siteSearch,omitempty"`

	SiteSearchFilter string `json:"siteSearchFilter,omitempty"`

	Sort string `json:"sort,omitempty"`

	StartIndex int64 `json:"startIndex,omitempty"`

	StartPage int64 `json:"startPage,omitempty"`

	Title string `json:"title,omitempty"`

	TotalResults int64 `json:"totalResults,omitempty,string"`
}

type Result struct {
	CacheId string `json:"cacheId,omitempty"`

	DisplayLink string `json:"displayLink,omitempty"`

	FileFormat string `json:"fileFormat,omitempty"`

	FormattedUrl string `json:"formattedUrl,omitempty"`

	HtmlFormattedUrl string `json:"htmlFormattedUrl,omitempty"`

	HtmlSnippet string `json:"htmlSnippet,omitempty"`

	HtmlTitle string `json:"htmlTitle,omitempty"`

	Image *ResultImage `json:"image,omitempty"`

	Kind string `json:"kind,omitempty"`

	Labels []*ResultLabels `json:"labels,omitempty"`

	Link string `json:"link,omitempty"`

	Mime string `json:"mime,omitempty"`

	Pagemap *ResultPagemap `json:"pagemap,omitempty"`

	Snippet string `json:"snippet,omitempty"`

	Title string `json:"title,omitempty"`
}

type ResultImage struct {
	ByteSize int64 `json:"byteSize,omitempty"`

	ContextLink string `json:"contextLink,omitempty"`

	Height int64 `json:"height,omitempty"`

	ThumbnailHeight int64 `json:"thumbnailHeight,omitempty"`

	ThumbnailLink string `json:"thumbnailLink,omitempty"`

	ThumbnailWidth int64 `json:"thumbnailWidth,omitempty"`

	Width int64 `json:"width,omitempty"`
}

type ResultLabels struct {
	DisplayName string `json:"displayName,omitempty"`

	Label_with_op string `json:"label_with_op,omitempty"`

	Name string `json:"name,omitempty"`
}

type ResultPagemap struct {
}

type Search struct {
	Context *Context `json:"context,omitempty"`

	Items []*Result `json:"items,omitempty"`

	Kind string `json:"kind,omitempty"`

	Promotions []*Promotion `json:"promotions,omitempty"`

	Queries map[string][]Query `json:"queries,omitempty"`

	SearchInformation *SearchSearchInformation `json:"searchInformation,omitempty"`

	Spelling *SearchSpelling `json:"spelling,omitempty"`

	Url *SearchUrl `json:"url,omitempty"`
}

type SearchSearchInformation struct {
	FormattedSearchTime string `json:"formattedSearchTime,omitempty"`

	FormattedTotalResults string `json:"formattedTotalResults,omitempty"`

	SearchTime float64 `json:"searchTime,omitempty"`

	TotalResults int64 `json:"totalResults,omitempty,string"`
}

type SearchSpelling struct {
	CorrectedQuery string `json:"correctedQuery,omitempty"`

	HtmlCorrectedQuery string `json:"htmlCorrectedQuery,omitempty"`
}

type SearchUrl struct {
	Template string `json:"template,omitempty"`

	Type string `json:"type,omitempty"`
}

// method id "search.cse.list":

type CseListCall struct {
	s             *Service
	q             string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns metadata about the search performed, metadata about the
// custom search engine used for the search, and the search results.

func (r *CseService) List(q string) *CseListCall {
	return &CseListCall{
		s:             r.s,
		q:             q,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "v1",
	}
}

// C2coff sets the optional parameter "c2coff": Turns off the
// translation between zh-CN and zh-TW.
func (c *CseListCall) C2coff(c2coff string) *CseListCall {
	c.params_.Set("c2coff", fmt.Sprintf("%v", c2coff))
	return c
}

// Cr sets the optional parameter "cr": Country restrict(s).
func (c *CseListCall) Cr(cr string) *CseListCall {
	c.params_.Set("cr", fmt.Sprintf("%v", cr))
	return c
}

// Cref sets the optional parameter "cref": The URL of a linked custom
// search engine
func (c *CseListCall) Cref(cref string) *CseListCall {
	c.params_.Set("cref", fmt.Sprintf("%v", cref))
	return c
}

// Cx sets the optional parameter "cx": The custom search engine ID to
// scope this search query
func (c *CseListCall) Cx(cx string) *CseListCall {
	c.params_.Set("cx", fmt.Sprintf("%v", cx))
	return c
}

// DateRestrict sets the optional parameter "dateRestrict": Specifies
// all search results are from a time period
func (c *CseListCall) DateRestrict(dateRestrict string) *CseListCall {
	c.params_.Set("dateRestrict", fmt.Sprintf("%v", dateRestrict))
	return c
}

// ExactTerms sets the optional parameter "exactTerms": Identifies a
// phrase that all documents in the search results must contain
func (c *CseListCall) ExactTerms(exactTerms string) *CseListCall {
	c.params_.Set("exactTerms", fmt.Sprintf("%v", exactTerms))
	return c
}

// ExcludeTerms sets the optional parameter "excludeTerms": Identifies a
// word or phrase that should not appear in any documents in the search
// results
func (c *CseListCall) ExcludeTerms(excludeTerms string) *CseListCall {
	c.params_.Set("excludeTerms", fmt.Sprintf("%v", excludeTerms))
	return c
}

// FileType sets the optional parameter "fileType": Returns images of a
// specified type. Some of the allowed values are: bmp, gif, png, jpg,
// svg, pdf, ...
func (c *CseListCall) FileType(fileType string) *CseListCall {
	c.params_.Set("fileType", fmt.Sprintf("%v", fileType))
	return c
}

// Filter sets the optional parameter "filter": Controls turning on or
// off the duplicate content filter.
func (c *CseListCall) Filter(filter string) *CseListCall {
	c.params_.Set("filter", fmt.Sprintf("%v", filter))
	return c
}

// Gl sets the optional parameter "gl": Geolocation of end user.
func (c *CseListCall) Gl(gl string) *CseListCall {
	c.params_.Set("gl", fmt.Sprintf("%v", gl))
	return c
}

// Googlehost sets the optional parameter "googlehost": The local Google
// domain to use to perform the search.
func (c *CseListCall) Googlehost(googlehost string) *CseListCall {
	c.params_.Set("googlehost", fmt.Sprintf("%v", googlehost))
	return c
}

// HighRange sets the optional parameter "highRange": Creates a range in
// form as_nlo value..as_nhi value and attempts to append it to query
func (c *CseListCall) HighRange(highRange string) *CseListCall {
	c.params_.Set("highRange", fmt.Sprintf("%v", highRange))
	return c
}

// Hl sets the optional parameter "hl": Sets the user interface
// language.
func (c *CseListCall) Hl(hl string) *CseListCall {
	c.params_.Set("hl", fmt.Sprintf("%v", hl))
	return c
}

// Hq sets the optional parameter "hq": Appends the extra query terms to
// the query.
func (c *CseListCall) Hq(hq string) *CseListCall {
	c.params_.Set("hq", fmt.Sprintf("%v", hq))
	return c
}

// ImgColorType sets the optional parameter "imgColorType": Returns
// black and white, grayscale, or color images: mono, gray, and color.
func (c *CseListCall) ImgColorType(imgColorType string) *CseListCall {
	c.params_.Set("imgColorType", fmt.Sprintf("%v", imgColorType))
	return c
}

// ImgDominantColor sets the optional parameter "imgDominantColor":
// Returns images of a specific dominant color: yellow, green, teal,
// blue, purple, pink, white, gray, black and brown.
func (c *CseListCall) ImgDominantColor(imgDominantColor string) *CseListCall {
	c.params_.Set("imgDominantColor", fmt.Sprintf("%v", imgDominantColor))
	return c
}

// ImgSize sets the optional parameter "imgSize": Returns images of a
// specified size, where size can be one of: icon, small, medium, large,
// xlarge, xxlarge, and huge.
func (c *CseListCall) ImgSize(imgSize string) *CseListCall {
	c.params_.Set("imgSize", fmt.Sprintf("%v", imgSize))
	return c
}

// ImgType sets the optional parameter "imgType": Returns images of a
// type, which can be one of: clipart, face, lineart, news, and photo.
func (c *CseListCall) ImgType(imgType string) *CseListCall {
	c.params_.Set("imgType", fmt.Sprintf("%v", imgType))
	return c
}

// LinkSite sets the optional parameter "linkSite": Specifies that all
// search results should contain a link to a particular URL
func (c *CseListCall) LinkSite(linkSite string) *CseListCall {
	c.params_.Set("linkSite", fmt.Sprintf("%v", linkSite))
	return c
}

// LowRange sets the optional parameter "lowRange": Creates a range in
// form as_nlo value..as_nhi value and attempts to append it to query
func (c *CseListCall) LowRange(lowRange string) *CseListCall {
	c.params_.Set("lowRange", fmt.Sprintf("%v", lowRange))
	return c
}

// Lr sets the optional parameter "lr": The language restriction for the
// search results
func (c *CseListCall) Lr(lr string) *CseListCall {
	c.params_.Set("lr", fmt.Sprintf("%v", lr))
	return c
}

// Num sets the optional parameter "num": Number of search results to
// return
func (c *CseListCall) Num(num int64) *CseListCall {
	c.params_.Set("num", fmt.Sprintf("%v", num))
	return c
}

// OrTerms sets the optional parameter "orTerms": Provides additional
// search terms to check for in a document, where each document in the
// search results must contain at least one of the additional search
// terms
func (c *CseListCall) OrTerms(orTerms string) *CseListCall {
	c.params_.Set("orTerms", fmt.Sprintf("%v", orTerms))
	return c
}

// RelatedSite sets the optional parameter "relatedSite": Specifies that
// all search results should be pages that are related to the specified
// URL
func (c *CseListCall) RelatedSite(relatedSite string) *CseListCall {
	c.params_.Set("relatedSite", fmt.Sprintf("%v", relatedSite))
	return c
}

// Rights sets the optional parameter "rights": Filters based on
// licensing. Supported values include: cc_publicdomain, cc_attribute,
// cc_sharealike, cc_noncommercial, cc_nonderived and combinations of
// these.
func (c *CseListCall) Rights(rights string) *CseListCall {
	c.params_.Set("rights", fmt.Sprintf("%v", rights))
	return c
}

// Safe sets the optional parameter "safe": Search safety level
func (c *CseListCall) Safe(safe string) *CseListCall {
	c.params_.Set("safe", fmt.Sprintf("%v", safe))
	return c
}

// SearchType sets the optional parameter "searchType": Specifies the
// search type: image.
func (c *CseListCall) SearchType(searchType string) *CseListCall {
	c.params_.Set("searchType", fmt.Sprintf("%v", searchType))
	return c
}

// SiteSearch sets the optional parameter "siteSearch": Specifies all
// search results should be pages from a given site
func (c *CseListCall) SiteSearch(siteSearch string) *CseListCall {
	c.params_.Set("siteSearch", fmt.Sprintf("%v", siteSearch))
	return c
}

// SiteSearchFilter sets the optional parameter "siteSearchFilter":
// Controls whether to include or exclude results from the site named in
// the as_sitesearch parameter
func (c *CseListCall) SiteSearchFilter(siteSearchFilter string) *CseListCall {
	c.params_.Set("siteSearchFilter", fmt.Sprintf("%v", siteSearchFilter))
	return c
}

// Sort sets the optional parameter "sort": The sort expression to apply
// to the results
func (c *CseListCall) Sort(sort string) *CseListCall {
	c.params_.Set("sort", fmt.Sprintf("%v", sort))
	return c
}

// Start sets the optional parameter "start": The index of the first
// result to return
func (c *CseListCall) Start(start int64) *CseListCall {
	c.params_.Set("start", fmt.Sprintf("%v", start))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CseListCall) Fields(s ...googleapi.Field) *CseListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CseListCall) Do() (*Search, error) {
	var returnValue *Search
	c.params_.Set("q", fmt.Sprintf("%v", c.q))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns metadata about the search performed, metadata about the custom search engine used for the search, and the search results.",
	//   "httpMethod": "GET",
	//   "id": "search.cse.list",
	//   "parameterOrder": [
	//     "q"
	//   ],
	//   "parameters": {
	//     "c2coff": {
	//       "description": "Turns off the translation between zh-CN and zh-TW.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cr": {
	//       "description": "Country restrict(s).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cref": {
	//       "description": "The URL of a linked custom search engine",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cx": {
	//       "description": "The custom search engine ID to scope this search query",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dateRestrict": {
	//       "description": "Specifies all search results are from a time period",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "exactTerms": {
	//       "description": "Identifies a phrase that all documents in the search results must contain",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "excludeTerms": {
	//       "description": "Identifies a word or phrase that should not appear in any documents in the search results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "fileType": {
	//       "description": "Returns images of a specified type. Some of the allowed values are: bmp, gif, png, jpg, svg, pdf, ...",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Controls turning on or off the duplicate content filter.",
	//       "enum": [
	//         "0",
	//         "1"
	//       ],
	//       "enumDescriptions": [
	//         "Turns off duplicate content filter.",
	//         "Turns on duplicate content filter."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gl": {
	//       "description": "Geolocation of end user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "googlehost": {
	//       "description": "The local Google domain to use to perform the search.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "highRange": {
	//       "description": "Creates a range in form as_nlo value..as_nhi value and attempts to append it to query",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hl": {
	//       "description": "Sets the user interface language.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "hq": {
	//       "description": "Appends the extra query terms to the query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "imgColorType": {
	//       "description": "Returns black and white, grayscale, or color images: mono, gray, and color.",
	//       "enum": [
	//         "color",
	//         "gray",
	//         "mono"
	//       ],
	//       "enumDescriptions": [
	//         "color",
	//         "gray",
	//         "mono"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "imgDominantColor": {
	//       "description": "Returns images of a specific dominant color: yellow, green, teal, blue, purple, pink, white, gray, black and brown.",
	//       "enum": [
	//         "black",
	//         "blue",
	//         "brown",
	//         "gray",
	//         "green",
	//         "pink",
	//         "purple",
	//         "teal",
	//         "white",
	//         "yellow"
	//       ],
	//       "enumDescriptions": [
	//         "black",
	//         "blue",
	//         "brown",
	//         "gray",
	//         "green",
	//         "pink",
	//         "purple",
	//         "teal",
	//         "white",
	//         "yellow"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "imgSize": {
	//       "description": "Returns images of a specified size, where size can be one of: icon, small, medium, large, xlarge, xxlarge, and huge.",
	//       "enum": [
	//         "huge",
	//         "icon",
	//         "large",
	//         "medium",
	//         "small",
	//         "xlarge",
	//         "xxlarge"
	//       ],
	//       "enumDescriptions": [
	//         "huge",
	//         "icon",
	//         "large",
	//         "medium",
	//         "small",
	//         "xlarge",
	//         "xxlarge"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "imgType": {
	//       "description": "Returns images of a type, which can be one of: clipart, face, lineart, news, and photo.",
	//       "enum": [
	//         "clipart",
	//         "face",
	//         "lineart",
	//         "news",
	//         "photo"
	//       ],
	//       "enumDescriptions": [
	//         "clipart",
	//         "face",
	//         "lineart",
	//         "news",
	//         "photo"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "linkSite": {
	//       "description": "Specifies that all search results should contain a link to a particular URL",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lowRange": {
	//       "description": "Creates a range in form as_nlo value..as_nhi value and attempts to append it to query",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "lr": {
	//       "description": "The language restriction for the search results",
	//       "enum": [
	//         "lang_ar",
	//         "lang_bg",
	//         "lang_ca",
	//         "lang_cs",
	//         "lang_da",
	//         "lang_de",
	//         "lang_el",
	//         "lang_en",
	//         "lang_es",
	//         "lang_et",
	//         "lang_fi",
	//         "lang_fr",
	//         "lang_hr",
	//         "lang_hu",
	//         "lang_id",
	//         "lang_is",
	//         "lang_it",
	//         "lang_iw",
	//         "lang_ja",
	//         "lang_ko",
	//         "lang_lt",
	//         "lang_lv",
	//         "lang_nl",
	//         "lang_no",
	//         "lang_pl",
	//         "lang_pt",
	//         "lang_ro",
	//         "lang_ru",
	//         "lang_sk",
	//         "lang_sl",
	//         "lang_sr",
	//         "lang_sv",
	//         "lang_tr",
	//         "lang_zh-CN",
	//         "lang_zh-TW"
	//       ],
	//       "enumDescriptions": [
	//         "Arabic",
	//         "Bulgarian",
	//         "Catalan",
	//         "Czech",
	//         "Danish",
	//         "German",
	//         "Greek",
	//         "English",
	//         "Spanish",
	//         "Estonian",
	//         "Finnish",
	//         "French",
	//         "Croatian",
	//         "Hungarian",
	//         "Indonesian",
	//         "Icelandic",
	//         "Italian",
	//         "Hebrew",
	//         "Japanese",
	//         "Korean",
	//         "Lithuanian",
	//         "Latvian",
	//         "Dutch",
	//         "Norwegian",
	//         "Polish",
	//         "Portuguese",
	//         "Romanian",
	//         "Russian",
	//         "Slovak",
	//         "Slovenian",
	//         "Serbian",
	//         "Swedish",
	//         "Turkish",
	//         "Chinese (Simplified)",
	//         "Chinese (Traditional)"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "num": {
	//       "default": "10",
	//       "description": "Number of search results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "orTerms": {
	//       "description": "Provides additional search terms to check for in a document, where each document in the search results must contain at least one of the additional search terms",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "q": {
	//       "description": "Query",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "relatedSite": {
	//       "description": "Specifies that all search results should be pages that are related to the specified URL",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "rights": {
	//       "description": "Filters based on licensing. Supported values include: cc_publicdomain, cc_attribute, cc_sharealike, cc_noncommercial, cc_nonderived and combinations of these.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "safe": {
	//       "default": "off",
	//       "description": "Search safety level",
	//       "enum": [
	//         "high",
	//         "medium",
	//         "off"
	//       ],
	//       "enumDescriptions": [
	//         "Enables highest level of safe search filtering.",
	//         "Enables moderate safe search filtering.",
	//         "Disables safe search filtering."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "searchType": {
	//       "description": "Specifies the search type: image.",
	//       "enum": [
	//         "image"
	//       ],
	//       "enumDescriptions": [
	//         "custom image search"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "siteSearch": {
	//       "description": "Specifies all search results should be pages from a given site",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "siteSearchFilter": {
	//       "description": "Controls whether to include or exclude results from the site named in the as_sitesearch parameter",
	//       "enum": [
	//         "e",
	//         "i"
	//       ],
	//       "enumDescriptions": [
	//         "exclude",
	//         "include"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The sort expression to apply to the results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "start": {
	//       "description": "The index of the first result to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v1",
	//   "response": {
	//     "$ref": "Search"
	//   }
	// }

}
