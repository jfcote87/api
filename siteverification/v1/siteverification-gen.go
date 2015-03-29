// Package siteverification provides access to the Google Site Verification API.
//
// See https://developers.google.com/site-verification/
//
// Usage example:
//
//   import "github.com/jfcote87/api/siteverification/v1"
//   ...
//   siteverificationService, err := siteverification.New(oauthHttpClient)
package siteverification

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

const apiId = "siteVerification:v1"
const apiName = "siteVerification"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/siteVerification/v1/"}

// OAuth2 scopes used by this API.
const (
	// Manage the list of sites and domains you control
	SiteverificationScope = "https://www.googleapis.com/auth/siteverification"

	// Manage your new site verifications with Google
	SiteverificationVerify_onlyScope = "https://www.googleapis.com/auth/siteverification.verify_only"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.WebResource = NewWebResourceService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	WebResource *WebResourceService
}

func NewWebResourceService(s *Service) *WebResourceService {
	rs := &WebResourceService{s: s}
	return rs
}

type WebResourceService struct {
	s *Service
}

type SiteVerificationWebResourceGettokenRequest struct {
	// Site: The site for which a verification token will be generated.
	Site *SiteVerificationWebResourceGettokenRequestSite `json:"site,omitempty"`

	// VerificationMethod: The verification method that will be used to
	// verify this site. For sites, 'FILE' or 'META' methods may be used.
	// For domains, only 'DNS' may be used.
	VerificationMethod string `json:"verificationMethod,omitempty"`
}

type SiteVerificationWebResourceGettokenRequestSite struct {
	// Identifier: The site identifier. If the type is set to SITE, the
	// identifier is a URL. If the type is set to INET_DOMAIN, the site
	// identifier is a domain name.
	Identifier string `json:"identifier,omitempty"`

	// Type: The type of resource to be verified. Can be SITE or INET_DOMAIN
	// (domain name).
	Type string `json:"type,omitempty"`
}

type SiteVerificationWebResourceGettokenResponse struct {
	// Method: The verification method to use in conjunction with this
	// token. For FILE, the token should be placed in the top-level
	// directory of the site, stored inside a file of the same name. For
	// META, the token should be placed in the HEAD tag of the default page
	// that is loaded for the site. For DNS, the token should be placed in a
	// TXT record of the domain.
	Method string `json:"method,omitempty"`

	// Token: The verification token. The token must be placed appropriately
	// in order for verification to succeed.
	Token string `json:"token,omitempty"`
}

type SiteVerificationWebResourceListResponse struct {
	// Items: The list of sites that are owned by the authenticated user.
	Items []*SiteVerificationWebResourceResource `json:"items,omitempty"`
}

type SiteVerificationWebResourceResource struct {
	// Id: The string used to identify this site. This value should be used
	// in the "id" portion of the REST URL for the Get, Update, and Delete
	// operations.
	Id string `json:"id,omitempty"`

	// Owners: The email addresses of all verified owners.
	Owners []string `json:"owners,omitempty"`

	// Site: The address and type of a site that is verified or will be
	// verified.
	Site *SiteVerificationWebResourceResourceSite `json:"site,omitempty"`
}

type SiteVerificationWebResourceResourceSite struct {
	// Identifier: The site identifier. If the type is set to SITE, the
	// identifier is a URL. If the type is set to INET_DOMAIN, the site
	// identifier is a domain name.
	Identifier string `json:"identifier,omitempty"`

	// Type: The site type. Can be SITE or INET_DOMAIN (domain name).
	Type string `json:"type,omitempty"`
}

// method id "siteVerification.webResource.delete":

type WebResourceDeleteCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Relinquish ownership of a website or domain.

func (r *WebResourceService) Delete(id string) *WebResourceDeleteCall {
	return &WebResourceDeleteCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "webResource/{id}",
	}
}

func (c *WebResourceDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Relinquish ownership of a website or domain.",
	//   "httpMethod": "DELETE",
	//   "id": "siteVerification.webResource.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.get":

type WebResourceGetCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get the most current data for a website or domain.

func (r *WebResourceService) Get(id string) *WebResourceGetCall {
	return &WebResourceGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "webResource/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceGetCall) Fields(s ...googleapi.Field) *WebResourceGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WebResourceGetCall) Do() (*SiteVerificationWebResourceResource, error) {
	var returnValue *SiteVerificationWebResourceResource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get the most current data for a website or domain.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.getToken":

type WebResourceGetTokenCall struct {
	s                                          *Service
	siteverificationwebresourcegettokenrequest *SiteVerificationWebResourceGettokenRequest
	caller_                                    googleapi.Caller
	params_                                    url.Values
	pathTemplate_                              string
}

// GetToken: Get a verification token for placing on a website or
// domain.

func (r *WebResourceService) GetToken(siteverificationwebresourcegettokenrequest *SiteVerificationWebResourceGettokenRequest) *WebResourceGetTokenCall {
	return &WebResourceGetTokenCall{
		s: r.s,
		siteverificationwebresourcegettokenrequest: siteverificationwebresourcegettokenrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "token",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceGetTokenCall) Fields(s ...googleapi.Field) *WebResourceGetTokenCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WebResourceGetTokenCall) Do() (*SiteVerificationWebResourceGettokenResponse, error) {
	var returnValue *SiteVerificationWebResourceGettokenResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.siteverificationwebresourcegettokenrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get a verification token for placing on a website or domain.",
	//   "httpMethod": "POST",
	//   "id": "siteVerification.webResource.getToken",
	//   "path": "token",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceGettokenRequest"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceGettokenResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification",
	//     "https://www.googleapis.com/auth/siteverification.verify_only"
	//   ]
	// }

}

// method id "siteVerification.webResource.insert":

type WebResourceInsertCall struct {
	s                                   *Service
	verificationMethod                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	caller_                             googleapi.Caller
	params_                             url.Values
	pathTemplate_                       string
}

// Insert: Attempt verification of a website or domain.

func (r *WebResourceService) Insert(verificationMethod string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourceInsertCall {
	return &WebResourceInsertCall{
		s:                                   r.s,
		verificationMethod:                  verificationMethod,
		siteverificationwebresourceresource: siteverificationwebresourceresource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "webResource",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceInsertCall) Fields(s ...googleapi.Field) *WebResourceInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WebResourceInsertCall) Do() (*SiteVerificationWebResourceResource, error) {
	var returnValue *SiteVerificationWebResourceResource
	c.params_.Set("verificationMethod", fmt.Sprintf("%v", c.verificationMethod))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.siteverificationwebresourceresource,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Attempt verification of a website or domain.",
	//   "httpMethod": "POST",
	//   "id": "siteVerification.webResource.insert",
	//   "parameterOrder": [
	//     "verificationMethod"
	//   ],
	//   "parameters": {
	//     "verificationMethod": {
	//       "description": "The method to use for verifying a site or domain.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification",
	//     "https://www.googleapis.com/auth/siteverification.verify_only"
	//   ]
	// }

}

// method id "siteVerification.webResource.list":

type WebResourceListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Get the list of your verified websites and domains.

func (r *WebResourceService) List() *WebResourceListCall {
	return &WebResourceListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "webResource",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceListCall) Fields(s ...googleapi.Field) *WebResourceListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WebResourceListCall) Do() (*SiteVerificationWebResourceListResponse, error) {
	var returnValue *SiteVerificationWebResourceListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get the list of your verified websites and domains.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.list",
	//   "path": "webResource",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.patch":

type WebResourcePatchCall struct {
	s                                   *Service
	id                                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	caller_                             googleapi.Caller
	params_                             url.Values
	pathTemplate_                       string
}

// Patch: Modify the list of owners for your website or domain. This
// method supports patch semantics.

func (r *WebResourceService) Patch(id string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourcePatchCall {
	return &WebResourcePatchCall{
		s:  r.s,
		id: id,
		siteverificationwebresourceresource: siteverificationwebresourceresource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "webResource/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourcePatchCall) Fields(s ...googleapi.Field) *WebResourcePatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WebResourcePatchCall) Do() (*SiteVerificationWebResourceResource, error) {
	var returnValue *SiteVerificationWebResourceResource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.siteverificationwebresourceresource,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Modify the list of owners for your website or domain. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "siteVerification.webResource.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.update":

type WebResourceUpdateCall struct {
	s                                   *Service
	id                                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	caller_                             googleapi.Caller
	params_                             url.Values
	pathTemplate_                       string
}

// Update: Modify the list of owners for your website or domain.

func (r *WebResourceService) Update(id string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourceUpdateCall {
	return &WebResourceUpdateCall{
		s:  r.s,
		id: id,
		siteverificationwebresourceresource: siteverificationwebresourceresource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "webResource/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceUpdateCall) Fields(s ...googleapi.Field) *WebResourceUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *WebResourceUpdateCall) Do() (*SiteVerificationWebResourceResource, error) {
	var returnValue *SiteVerificationWebResourceResource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.siteverificationwebresourceresource,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Modify the list of owners for your website or domain.",
	//   "httpMethod": "PUT",
	//   "id": "siteVerification.webResource.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}
