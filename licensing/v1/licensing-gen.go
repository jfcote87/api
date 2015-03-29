// Package licensing provides access to the Enterprise License Manager API.
//
// See https://developers.google.com/google-apps/licensing/
//
// Usage example:
//
//   import "github.com/jfcote87/api/licensing/v1"
//   ...
//   licensingService, err := licensing.New(oauthHttpClient)
package licensing

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

const apiId = "licensing:v1"
const apiName = "licensing"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/apps/licensing/v1/product/"}

// OAuth2 scopes used by this API.
const (
	// View and manage Google Apps licenses for your domain
	AppsLicensingScope = "https://www.googleapis.com/auth/apps.licensing"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.LicenseAssignments = NewLicenseAssignmentsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	LicenseAssignments *LicenseAssignmentsService
}

func NewLicenseAssignmentsService(s *Service) *LicenseAssignmentsService {
	rs := &LicenseAssignmentsService{s: s}
	return rs
}

type LicenseAssignmentsService struct {
	s *Service
}

type LicenseAssignment struct {
	// Etags: ETag of the resource.
	Etags string `json:"etags,omitempty"`

	// Kind: Identifies the resource as a LicenseAssignment.
	Kind string `json:"kind,omitempty"`

	// ProductId: Name of the product.
	ProductId string `json:"productId,omitempty"`

	// SelfLink: Link to this page.
	SelfLink string `json:"selfLink,omitempty"`

	// SkuId: Name of the sku of the product.
	SkuId string `json:"skuId,omitempty"`

	// UserId: Email id of the user.
	UserId string `json:"userId,omitempty"`
}

type LicenseAssignmentInsert struct {
	// UserId: Email id of the user
	UserId string `json:"userId,omitempty"`
}

type LicenseAssignmentList struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: The LicenseAssignments in this page of results.
	Items []*LicenseAssignment `json:"items,omitempty"`

	// Kind: Identifies the resource as a collection of LicenseAssignments.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "licensing.licenseAssignments.delete":

type LicenseAssignmentsDeleteCall struct {
	s             *Service
	productId     string
	skuId         string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Revoke License.

func (r *LicenseAssignmentsService) Delete(productId string, skuId string, userId string) *LicenseAssignmentsDeleteCall {
	return &LicenseAssignmentsDeleteCall{
		s:             r.s,
		productId:     productId,
		skuId:         skuId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{productId}/sku/{skuId}/user/{userId}",
	}
}

func (c *LicenseAssignmentsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Revoke License.",
	//   "httpMethod": "DELETE",
	//   "id": "licensing.licenseAssignments.delete",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.get":

type LicenseAssignmentsGetCall struct {
	s             *Service
	productId     string
	skuId         string
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get license assignment of a particular product and sku for a
// user

func (r *LicenseAssignmentsService) Get(productId string, skuId string, userId string) *LicenseAssignmentsGetCall {
	return &LicenseAssignmentsGetCall{
		s:             r.s,
		productId:     productId,
		skuId:         skuId,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{productId}/sku/{skuId}/user/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsGetCall) Fields(s ...googleapi.Field) *LicenseAssignmentsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicenseAssignmentsGetCall) Do() (*LicenseAssignment, error) {
	var returnValue *LicenseAssignment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get license assignment of a particular product and sku for a user",
	//   "httpMethod": "GET",
	//   "id": "licensing.licenseAssignments.get",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.insert":

type LicenseAssignmentsInsertCall struct {
	s                       *Service
	productId               string
	skuId                   string
	licenseassignmentinsert *LicenseAssignmentInsert
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Insert: Assign License.

func (r *LicenseAssignmentsService) Insert(productId string, skuId string, licenseassignmentinsert *LicenseAssignmentInsert) *LicenseAssignmentsInsertCall {
	return &LicenseAssignmentsInsertCall{
		s:         r.s,
		productId: productId,
		skuId:     skuId,
		licenseassignmentinsert: licenseassignmentinsert,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "{productId}/sku/{skuId}/user",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsInsertCall) Fields(s ...googleapi.Field) *LicenseAssignmentsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicenseAssignmentsInsertCall) Do() (*LicenseAssignment, error) {
	var returnValue *LicenseAssignment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.licenseassignmentinsert,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Assign License.",
	//   "httpMethod": "POST",
	//   "id": "licensing.licenseAssignments.insert",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user",
	//   "request": {
	//     "$ref": "LicenseAssignmentInsert"
	//   },
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.listForProduct":

type LicenseAssignmentsListForProductCall struct {
	s             *Service
	productId     string
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ListForProduct: List license assignments for given product of the
// customer.

func (r *LicenseAssignmentsService) ListForProduct(productId string, customerId string) *LicenseAssignmentsListForProductCall {
	return &LicenseAssignmentsListForProductCall{
		s:             r.s,
		productId:     productId,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{productId}/users",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of campaigns to return at one time. Must be positive.  Default value
// is 100.
func (c *LicenseAssignmentsListForProductCall) MaxResults(maxResults int64) *LicenseAssignmentsListForProductCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to fetch the
// next page. By default server will return first page
func (c *LicenseAssignmentsListForProductCall) PageToken(pageToken string) *LicenseAssignmentsListForProductCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsListForProductCall) Fields(s ...googleapi.Field) *LicenseAssignmentsListForProductCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicenseAssignmentsListForProductCall) Do() (*LicenseAssignmentList, error) {
	var returnValue *LicenseAssignmentList
	c.params_.Set("customerId", fmt.Sprintf("%v", c.customerId))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List license assignments for given product of the customer.",
	//   "httpMethod": "GET",
	//   "id": "licensing.licenseAssignments.listForProduct",
	//   "parameterOrder": [
	//     "productId",
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "CustomerId represents the customer for whom licenseassignments are queried",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "100",
	//       "description": "Maximum number of campaigns to return at one time. Must be positive. Optional. Default value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "default": "",
	//       "description": "Token to fetch the next page.Optional. By default server will return first page",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/users",
	//   "response": {
	//     "$ref": "LicenseAssignmentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.listForProductAndSku":

type LicenseAssignmentsListForProductAndSkuCall struct {
	s             *Service
	productId     string
	skuId         string
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ListForProductAndSku: List license assignments for given product and
// sku of the customer.

func (r *LicenseAssignmentsService) ListForProductAndSku(productId string, skuId string, customerId string) *LicenseAssignmentsListForProductAndSkuCall {
	return &LicenseAssignmentsListForProductAndSkuCall{
		s:             r.s,
		productId:     productId,
		skuId:         skuId,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{productId}/sku/{skuId}/users",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of campaigns to return at one time. Must be positive.  Default value
// is 100.
func (c *LicenseAssignmentsListForProductAndSkuCall) MaxResults(maxResults int64) *LicenseAssignmentsListForProductAndSkuCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to fetch the
// next page. By default server will return first page
func (c *LicenseAssignmentsListForProductAndSkuCall) PageToken(pageToken string) *LicenseAssignmentsListForProductAndSkuCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsListForProductAndSkuCall) Fields(s ...googleapi.Field) *LicenseAssignmentsListForProductAndSkuCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicenseAssignmentsListForProductAndSkuCall) Do() (*LicenseAssignmentList, error) {
	var returnValue *LicenseAssignmentList
	c.params_.Set("customerId", fmt.Sprintf("%v", c.customerId))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List license assignments for given product and sku of the customer.",
	//   "httpMethod": "GET",
	//   "id": "licensing.licenseAssignments.listForProductAndSku",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "CustomerId represents the customer for whom licenseassignments are queried",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "100",
	//       "description": "Maximum number of campaigns to return at one time. Must be positive. Optional. Default value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "default": "",
	//       "description": "Token to fetch the next page.Optional. By default server will return first page",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/users",
	//   "response": {
	//     "$ref": "LicenseAssignmentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.patch":

type LicenseAssignmentsPatchCall struct {
	s                 *Service
	productId         string
	skuId             string
	userId            string
	licenseassignment *LicenseAssignment
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// Patch: Assign License. This method supports patch semantics.

func (r *LicenseAssignmentsService) Patch(productId string, skuId string, userId string, licenseassignment *LicenseAssignment) *LicenseAssignmentsPatchCall {
	return &LicenseAssignmentsPatchCall{
		s:                 r.s,
		productId:         productId,
		skuId:             skuId,
		userId:            userId,
		licenseassignment: licenseassignment,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "{productId}/sku/{skuId}/user/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsPatchCall) Fields(s ...googleapi.Field) *LicenseAssignmentsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicenseAssignmentsPatchCall) Do() (*LicenseAssignment, error) {
	var returnValue *LicenseAssignment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.licenseassignment,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Assign License. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "licensing.licenseAssignments.patch",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku for which license would be revoked",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "request": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.update":

type LicenseAssignmentsUpdateCall struct {
	s                 *Service
	productId         string
	skuId             string
	userId            string
	licenseassignment *LicenseAssignment
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// Update: Assign License.

func (r *LicenseAssignmentsService) Update(productId string, skuId string, userId string, licenseassignment *LicenseAssignment) *LicenseAssignmentsUpdateCall {
	return &LicenseAssignmentsUpdateCall{
		s:                 r.s,
		productId:         productId,
		skuId:             skuId,
		userId:            userId,
		licenseassignment: licenseassignment,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "{productId}/sku/{skuId}/user/{userId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsUpdateCall) Fields(s ...googleapi.Field) *LicenseAssignmentsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LicenseAssignmentsUpdateCall) Do() (*LicenseAssignment, error) {
	var returnValue *LicenseAssignment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.licenseassignment,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Assign License.",
	//   "httpMethod": "PUT",
	//   "id": "licensing.licenseAssignments.update",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku for which license would be revoked",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "request": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}
