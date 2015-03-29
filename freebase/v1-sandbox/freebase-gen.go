// Package freebase provides access to the Freebase Search.
//
// See https://developers.google.com/freebase/
//
// Usage example:
//
//   import "github.com/jfcote87/api/freebase/v1-sandbox"
//   ...
//   freebaseService, err := freebase.New(oauthHttpClient)
package freebase

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

const apiId = "freebase:v1-sandbox"
const apiName = "freebase"
const apiVersion = "v1-sandbox"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/freebase/v1-sandbox/"}

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	return s, nil
}

type Service struct {
	client *http.Client
}

type ReconcileCandidate struct {
	// Confidence: Percentage likelihood that this candidate is the unique
	// matching entity. Value will be between 0.0 and 1.0
	Confidence float64 `json:"confidence,omitempty"`

	// Lang: Language code that candidate and notable names are displayed
	// in.
	Lang string `json:"lang,omitempty"`

	// Mid: Freebase MID of candidate entity.
	Mid string `json:"mid,omitempty"`

	// Name: Freebase name of matching entity in specified language.
	Name string `json:"name,omitempty"`

	// Notable: Type or profession the candidate is notable for.
	Notable *ReconcileCandidateNotable `json:"notable,omitempty"`
}

type ReconcileCandidateNotable struct {
	// Id: MID of notable category.
	Id string `json:"id,omitempty"`

	// Name: Name of notable category in specified language.
	Name string `json:"name,omitempty"`
}

type ReconcileGet struct {
	// Candidate: If filled, then the listed candidates are potential
	// matches, and such should be evaluated by a more discerning algorithm
	// or human. The matches are ordered by confidence.
	Candidate []*ReconcileCandidate `json:"candidate,omitempty"`

	// Costs: Server costs for reconciling.
	Costs *ReconcileGetCosts `json:"costs,omitempty"`

	// Match: If filled, this entity is guaranteed to match at requested
	// confidence probability (default 99%).
	Match *ReconcileCandidate `json:"match,omitempty"`

	// Warning: If filled, then there were recoverable problems that
	// affected the request. For example, some of the properties were
	// ignored because they either are not valid Freebase predicates or are
	// not indexed for reconciliation. The candidates returned should be
	// considered valid results, with the caveat that sections of the
	// request were ignored as specified by the warning text.
	Warning []*ReconcileGetWarning `json:"warning,omitempty"`
}

type ReconcileGetCosts struct {
	// Hits: Total number of hits found.
	Hits int64 `json:"hits,omitempty"`

	// Ms: Total milliseconds spent.
	Ms int64 `json:"ms,omitempty"`
}

type ReconcileGetWarning struct {
	// Location: Location of warning in the request e.g. invalid predicate.
	Location string `json:"location,omitempty"`

	// Message: Warning message to display to the user.
	Message string `json:"message,omitempty"`

	// Reason: Code for identifying classes of warnings.
	Reason string `json:"reason,omitempty"`
}

// method id "freebase.reconcile":

type ReconcileCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reconcile: Reconcile entities to Freebase open data.

func (s *Service) Reconcile() *ReconcileCall {
	return &ReconcileCall{
		s:             s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reconcile",
	}
}

// Confidence sets the optional parameter "confidence": Required
// confidence for a candidate to match. Must be between .5 and 1.0
func (c *ReconcileCall) Confidence(confidence float64) *ReconcileCall {
	c.params_.Set("confidence", fmt.Sprintf("%v", confidence))
	return c
}

// Kind sets the optional parameter "kind": Classifications of entity
// e.g. type, category, title.
func (c *ReconcileCall) Kind(kind ...string) *ReconcileCall {
	c.params_["kind"] = kind
	return c
}

// Lang sets the optional parameter "lang": Languages for names and
// values. First language is used for display. Default is 'en'.
func (c *ReconcileCall) Lang(lang ...string) *ReconcileCall {
	c.params_["lang"] = lang
	return c
}

// Limit sets the optional parameter "limit": Maximum number of
// candidates to return.
func (c *ReconcileCall) Limit(limit int64) *ReconcileCall {
	c.params_.Set("limit", fmt.Sprintf("%v", limit))
	return c
}

// Name sets the optional parameter "name": Name of entity.
func (c *ReconcileCall) Name(name string) *ReconcileCall {
	c.params_.Set("name", fmt.Sprintf("%v", name))
	return c
}

// Prop sets the optional parameter "prop": Property values for entity
// formatted as
// :
func (c *ReconcileCall) Prop(prop ...string) *ReconcileCall {
	c.params_["prop"] = prop
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReconcileCall) Fields(s ...googleapi.Field) *ReconcileCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReconcileCall) Do() (*ReconcileGet, error) {
	var returnValue *ReconcileGet
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Reconcile entities to Freebase open data.",
	//   "httpMethod": "GET",
	//   "id": "freebase.reconcile",
	//   "parameters": {
	//     "confidence": {
	//       "default": "0.99",
	//       "description": "Required confidence for a candidate to match. Must be between .5 and 1.0",
	//       "format": "float",
	//       "location": "query",
	//       "maximum": "1.0",
	//       "minimum": "0.0",
	//       "type": "number"
	//     },
	//     "kind": {
	//       "description": "Classifications of entity e.g. type, category, title.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "lang": {
	//       "description": "Languages for names and values. First language is used for display. Default is 'en'.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "limit": {
	//       "default": "3",
	//       "description": "Maximum number of candidates to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "25",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "name": {
	//       "description": "Name of entity.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "prop": {
	//       "description": "Property values for entity formatted as\n:",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "reconcile",
	//   "response": {
	//     "$ref": "ReconcileGet"
	//   }
	// }

}

// method id "freebase.search":

type SearchCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Search: Search Freebase open data.

func (s *Service) Search() *SearchCall {
	return &SearchCall{
		s:             s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "search",
	}
}

// As_of_time sets the optional parameter "as_of_time": A mql as_of_time
// value to use with mql_output queries.
func (c *SearchCall) As_of_time(as_of_time string) *SearchCall {
	c.params_.Set("as_of_time", fmt.Sprintf("%v", as_of_time))
	return c
}

// Callback sets the optional parameter "callback": JS method name for
// JSONP callbacks.
func (c *SearchCall) Callback(callback string) *SearchCall {
	c.params_.Set("callback", fmt.Sprintf("%v", callback))
	return c
}

// Cursor sets the optional parameter "cursor": The cursor value to use
// for the next page of results.
func (c *SearchCall) Cursor(cursor int64) *SearchCall {
	c.params_.Set("cursor", fmt.Sprintf("%v", cursor))
	return c
}

// Domain sets the optional parameter "domain": Restrict to topics with
// this Freebase domain id.
func (c *SearchCall) Domain(domain ...string) *SearchCall {
	c.params_["domain"] = domain
	return c
}

// Encode sets the optional parameter "encode": The encoding of the
// response. You can use this parameter to enable html encoding.
func (c *SearchCall) Encode(encode string) *SearchCall {
	c.params_.Set("encode", fmt.Sprintf("%v", encode))
	return c
}

// Exact sets the optional parameter "exact": Query on exact name and
// keys only.
func (c *SearchCall) Exact(exact bool) *SearchCall {
	c.params_.Set("exact", fmt.Sprintf("%v", exact))
	return c
}

// Filter sets the optional parameter "filter": A filter to apply to the
// query.
func (c *SearchCall) Filter(filter ...string) *SearchCall {
	c.params_["filter"] = filter
	return c
}

// Format sets the optional parameter "format": Structural format of the
// json response.
func (c *SearchCall) Format(format string) *SearchCall {
	c.params_.Set("format", fmt.Sprintf("%v", format))
	return c
}

// Help sets the optional parameter "help": The keyword to request help
// on.
func (c *SearchCall) Help(help string) *SearchCall {
	c.params_.Set("help", fmt.Sprintf("%v", help))
	return c
}

// Indent sets the optional parameter "indent": Whether to indent the
// json results or not.
func (c *SearchCall) Indent(indent bool) *SearchCall {
	c.params_.Set("indent", fmt.Sprintf("%v", indent))
	return c
}

// Lang sets the optional parameter "lang": The code of the language to
// run the query with. Default is 'en'.
func (c *SearchCall) Lang(lang ...string) *SearchCall {
	c.params_["lang"] = lang
	return c
}

// Limit sets the optional parameter "limit": Maximum number of results
// to return.
func (c *SearchCall) Limit(limit int64) *SearchCall {
	c.params_.Set("limit", fmt.Sprintf("%v", limit))
	return c
}

// Mid sets the optional parameter "mid": A mid to use instead of a
// query.
func (c *SearchCall) Mid(mid ...string) *SearchCall {
	c.params_["mid"] = mid
	return c
}

// Mql_output sets the optional parameter "mql_output": The MQL query to
// run againist the results to extract more data.
func (c *SearchCall) Mql_output(mql_output string) *SearchCall {
	c.params_.Set("mql_output", fmt.Sprintf("%v", mql_output))
	return c
}

// Output sets the optional parameter "output": An output expression to
// request data from matches.
func (c *SearchCall) Output(output string) *SearchCall {
	c.params_.Set("output", fmt.Sprintf("%v", output))
	return c
}

// Prefixed sets the optional parameter "prefixed": Prefix match against
// names and aliases.
func (c *SearchCall) Prefixed(prefixed bool) *SearchCall {
	c.params_.Set("prefixed", fmt.Sprintf("%v", prefixed))
	return c
}

// Query sets the optional parameter "query": Query term to search for.
func (c *SearchCall) Query(query string) *SearchCall {
	c.params_.Set("query", fmt.Sprintf("%v", query))
	return c
}

// Scoring sets the optional parameter "scoring": Relevance scoring
// algorithm to use.
func (c *SearchCall) Scoring(scoring string) *SearchCall {
	c.params_.Set("scoring", fmt.Sprintf("%v", scoring))
	return c
}

// Spell sets the optional parameter "spell": Request 'did you mean'
// suggestions
func (c *SearchCall) Spell(spell string) *SearchCall {
	c.params_.Set("spell", fmt.Sprintf("%v", spell))
	return c
}

// Stemmed sets the optional parameter "stemmed": Query on stemmed names
// and aliases. May not be used with prefixed.
func (c *SearchCall) Stemmed(stemmed bool) *SearchCall {
	c.params_.Set("stemmed", fmt.Sprintf("%v", stemmed))
	return c
}

// Type sets the optional parameter "type": Restrict to topics with this
// Freebase type id.
func (c *SearchCall) Type(type_ ...string) *SearchCall {
	c.params_["type"] = type_
	return c
}

// With sets the optional parameter "with": A rule to match against.
func (c *SearchCall) With(with ...string) *SearchCall {
	c.params_["with"] = with
	return c
}

// Without sets the optional parameter "without": A rule to not match
// against.
func (c *SearchCall) Without(without ...string) *SearchCall {
	c.params_["without"] = without
	return c
}

func (c *SearchCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *SearchCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Search Freebase open data.",
	//   "httpMethod": "GET",
	//   "id": "freebase.search",
	//   "parameters": {
	//     "as_of_time": {
	//       "description": "A mql as_of_time value to use with mql_output queries.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "callback": {
	//       "description": "JS method name for JSONP callbacks.",
	//       "location": "query",
	//       "pattern": "([A-Za-z0-9_$.]|\\[|\\])+",
	//       "type": "string"
	//     },
	//     "cursor": {
	//       "description": "The cursor value to use for the next page of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "domain": {
	//       "description": "Restrict to topics with this Freebase domain id.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "encode": {
	//       "default": "off",
	//       "description": "The encoding of the response. You can use this parameter to enable html encoding.",
	//       "enum": [
	//         "html",
	//         "off"
	//       ],
	//       "enumDescriptions": [
	//         "Encode certain characters in the response (such as tags and ambersands) using html encoding.",
	//         "No encoding of the response. You should not print the results directly on an web page without html-escaping the content first."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "exact": {
	//       "description": "Query on exact name and keys only.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "filter": {
	//       "description": "A filter to apply to the query.",
	//       "location": "query",
	//       "pattern": "^\\(.*\\)$",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "format": {
	//       "default": "entity",
	//       "description": "Structural format of the json response.",
	//       "enum": [
	//         "ac",
	//         "classic",
	//         "entity",
	//         "guids",
	//         "ids",
	//         "mids"
	//       ],
	//       "enumDescriptions": [
	//         "Compact format useful for autocomplete/suggest UIs.",
	//         "[DEPRECATED] Same format as was returned by api.freebase.com.",
	//         "Basic information about the entities.",
	//         "[DEPRECATED] Ordered list of a freebase guids.",
	//         "Ordered list of freebase ids.",
	//         "Ordered list of freebase mids."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "help": {
	//       "description": "The keyword to request help on.",
	//       "enum": [
	//         "langs",
	//         "mappings",
	//         "predicates"
	//       ],
	//       "enumDescriptions": [
	//         "The language codes served by the service.",
	//         "The property/path mappings supported by the filter and output request parameters.",
	//         "The predicates and path-terminating properties supported by the filter and output request parameters."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "indent": {
	//       "description": "Whether to indent the json results or not.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "lang": {
	//       "description": "The code of the language to run the query with. Default is 'en'.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "limit": {
	//       "default": "20",
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "mid": {
	//       "description": "A mid to use instead of a query.",
	//       "location": "query",
	//       "pattern": "^/[mgtx]/[0-2][0-9bcdfghjklmnpqrstvwxyz_]{1,24}$",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "mql_output": {
	//       "description": "The MQL query to run againist the results to extract more data.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "output": {
	//       "description": "An output expression to request data from matches.",
	//       "location": "query",
	//       "pattern": "^\\(.*\\)$",
	//       "type": "string"
	//     },
	//     "prefixed": {
	//       "description": "Prefix match against names and aliases.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "query": {
	//       "description": "Query term to search for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "scoring": {
	//       "default": "entity",
	//       "description": "Relevance scoring algorithm to use.",
	//       "enum": [
	//         "entity",
	//         "freebase",
	//         "schema"
	//       ],
	//       "enumDescriptions": [
	//         "Use freebase and popularity entity ranking.",
	//         "Use freebase entity ranking.",
	//         "Use schema ranking for properties and types."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "spell": {
	//       "default": "no_spelling",
	//       "description": "Request 'did you mean' suggestions",
	//       "enum": [
	//         "always",
	//         "no_results",
	//         "no_spelling"
	//       ],
	//       "enumDescriptions": [
	//         "Request spelling suggestions for any query at least three characters long.",
	//         "Request spelling suggestions if no results were found.",
	//         "Don't request spelling suggestions."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stemmed": {
	//       "description": "Query on stemmed names and aliases. May not be used with prefixed.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "type": {
	//       "description": "Restrict to topics with this Freebase type id.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "with": {
	//       "description": "A rule to match against.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "without": {
	//       "description": "A rule to not match against.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "search",
	//   "supportsMediaDownload": true
	// }

}
