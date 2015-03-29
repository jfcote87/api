// Package appstate provides access to the Google App State API.
//
// See https://developers.google.com/games/services/web/api/states
//
// Usage example:
//
//   import "github.com/jfcote87/api/appstate/v1"
//   ...
//   appstateService, err := appstate.New(oauthHttpClient)
package appstate

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

const apiId = "appstate:v1"
const apiName = "appstate"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/appstate/v1/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data for this application
	AppstateScope = "https://www.googleapis.com/auth/appstate"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.States = NewStatesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	States *StatesService
}

func NewStatesService(s *Service) *StatesService {
	rs := &StatesService{s: s}
	return rs
}

type StatesService struct {
	s *Service
}

type GetResponse struct {
	// CurrentStateVersion: The current app state version.
	CurrentStateVersion string `json:"currentStateVersion,omitempty"`

	// Data: The requested data.
	Data string `json:"data,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#getResponse.
	Kind string `json:"kind,omitempty"`

	// StateKey: The key for the data.
	StateKey int64 `json:"stateKey,omitempty"`
}

type ListResponse struct {
	// Items: The app state data.
	Items []*GetResponse `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#listResponse.
	Kind string `json:"kind,omitempty"`

	// MaximumKeyCount: The maximum number of keys allowed for this user.
	MaximumKeyCount int64 `json:"maximumKeyCount,omitempty"`
}

type UpdateRequest struct {
	// Data: The new app state data that your application is trying to
	// update with.
	Data string `json:"data,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#updateRequest.
	Kind string `json:"kind,omitempty"`
}

type WriteResult struct {
	// CurrentStateVersion: The version of the data for this key on the
	// server.
	CurrentStateVersion string `json:"currentStateVersion,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string appstate#writeResult.
	Kind string `json:"kind,omitempty"`

	// StateKey: The written key.
	StateKey int64 `json:"stateKey,omitempty"`
}

// method id "appstate.states.clear":

type StatesClearCall struct {
	s             *Service
	stateKey      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Clear: Clears (sets to empty) the data for the passed key if and only
// if the passed version matches the currently stored version. This
// method results in a conflict error on version mismatch.

func (r *StatesService) Clear(stateKey int64) *StatesClearCall {
	return &StatesClearCall{
		s:             r.s,
		stateKey:      stateKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "states/{stateKey}/clear",
	}
}

// CurrentDataVersion sets the optional parameter "currentDataVersion":
// The version of the data to be cleared. Version strings are returned
// by the server.
func (c *StatesClearCall) CurrentDataVersion(currentDataVersion string) *StatesClearCall {
	c.params_.Set("currentDataVersion", fmt.Sprintf("%v", currentDataVersion))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesClearCall) Fields(s ...googleapi.Field) *StatesClearCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StatesClearCall) Do() (*WriteResult, error) {
	var returnValue *WriteResult
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Clears (sets to empty) the data for the passed key if and only if the passed version matches the currently stored version. This method results in a conflict error on version mismatch.",
	//   "httpMethod": "POST",
	//   "id": "appstate.states.clear",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "currentDataVersion": {
	//       "description": "The version of the data to be cleared. Version strings are returned by the server.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}/clear",
	//   "response": {
	//     "$ref": "WriteResult"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.delete":

type StatesDeleteCall struct {
	s             *Service
	stateKey      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a key and the data associated with it. The key is
// removed and no longer counts against the key quota. Note that since
// this method is not safe in the face of concurrent modifications, it
// should only be used for development and testing purposes. Invoking
// this method in shipping code can result in data loss and data
// corruption.

func (r *StatesService) Delete(stateKey int64) *StatesDeleteCall {
	return &StatesDeleteCall{
		s:             r.s,
		stateKey:      stateKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "states/{stateKey}",
	}
}

func (c *StatesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a key and the data associated with it. The key is removed and no longer counts against the key quota. Note that since this method is not safe in the face of concurrent modifications, it should only be used for development and testing purposes. Invoking this method in shipping code can result in data loss and data corruption.",
	//   "httpMethod": "DELETE",
	//   "id": "appstate.states.delete",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.get":

type StatesGetCall struct {
	s             *Service
	stateKey      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the data corresponding to the passed key. If the key
// does not exist on the server, an HTTP 404 will be returned.

func (r *StatesService) Get(stateKey int64) *StatesGetCall {
	return &StatesGetCall{
		s:             r.s,
		stateKey:      stateKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "states/{stateKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesGetCall) Fields(s ...googleapi.Field) *StatesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StatesGetCall) Do() (*GetResponse, error) {
	var returnValue *GetResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the data corresponding to the passed key. If the key does not exist on the server, an HTTP 404 will be returned.",
	//   "httpMethod": "GET",
	//   "id": "appstate.states.get",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}",
	//   "response": {
	//     "$ref": "GetResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.list":

type StatesListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all the states keys, and optionally the state data.

func (r *StatesService) List() *StatesListCall {
	return &StatesListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "states",
	}
}

// IncludeData sets the optional parameter "includeData": Whether to
// include the full data in addition to the version number
func (c *StatesListCall) IncludeData(includeData bool) *StatesListCall {
	c.params_.Set("includeData", fmt.Sprintf("%v", includeData))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesListCall) Fields(s ...googleapi.Field) *StatesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StatesListCall) Do() (*ListResponse, error) {
	var returnValue *ListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all the states keys, and optionally the state data.",
	//   "httpMethod": "GET",
	//   "id": "appstate.states.list",
	//   "parameters": {
	//     "includeData": {
	//       "default": "false",
	//       "description": "Whether to include the full data in addition to the version number",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "states",
	//   "response": {
	//     "$ref": "ListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}

// method id "appstate.states.update":

type StatesUpdateCall struct {
	s             *Service
	stateKey      int64
	updaterequest *UpdateRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update the data associated with the input key if and only if
// the passed version matches the currently stored version. This method
// is safe in the face of concurrent writes. Maximum per-key size is
// 128KB.

func (r *StatesService) Update(stateKey int64, updaterequest *UpdateRequest) *StatesUpdateCall {
	return &StatesUpdateCall{
		s:             r.s,
		stateKey:      stateKey,
		updaterequest: updaterequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "states/{stateKey}",
	}
}

// CurrentStateVersion sets the optional parameter
// "currentStateVersion": The version of the app state your application
// is attempting to update. If this does not match the current version,
// this method will return a conflict error. If there is no data stored
// on the server for this key, the update will succeed irrespective of
// the value of this parameter.
func (c *StatesUpdateCall) CurrentStateVersion(currentStateVersion string) *StatesUpdateCall {
	c.params_.Set("currentStateVersion", fmt.Sprintf("%v", currentStateVersion))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StatesUpdateCall) Fields(s ...googleapi.Field) *StatesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StatesUpdateCall) Do() (*WriteResult, error) {
	var returnValue *WriteResult
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"stateKey": strconv.FormatInt(c.stateKey, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.updaterequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the data associated with the input key if and only if the passed version matches the currently stored version. This method is safe in the face of concurrent writes. Maximum per-key size is 128KB.",
	//   "httpMethod": "PUT",
	//   "id": "appstate.states.update",
	//   "parameterOrder": [
	//     "stateKey"
	//   ],
	//   "parameters": {
	//     "currentStateVersion": {
	//       "description": "The version of the app state your application is attempting to update. If this does not match the current version, this method will return a conflict error. If there is no data stored on the server for this key, the update will succeed irrespective of the value of this parameter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stateKey": {
	//       "description": "The key for the data to be retrieved.",
	//       "format": "int32",
	//       "location": "path",
	//       "maximum": "3",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "states/{stateKey}",
	//   "request": {
	//     "$ref": "UpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteResult"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/appstate"
	//   ]
	// }

}
