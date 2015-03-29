// Package oauth2 provides access to the Google OAuth2 API.
//
// See https://developers.google.com/accounts/docs/OAuth2
//
// Usage example:
//
//   import "github.com/jfcote87/api/oauth2/v2"
//   ...
//   oauth2Service, err := oauth2.New(oauthHttpClient)
package oauth2

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

const apiId = "oauth2:v2"
const apiName = "oauth2"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/"}

// OAuth2 scopes used by this API.
const (
	// Know your basic profile info and list of people in your circles.
	PlusLoginScope = "https://www.googleapis.com/auth/plus.login"

	// Know who you are on Google
	PlusMeScope = "https://www.googleapis.com/auth/plus.me"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"

	// View your basic profile info
	UserinfoProfileScope = "https://www.googleapis.com/auth/userinfo.profile"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Userinfo = NewUserinfoService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Userinfo *UserinfoService
}

func NewUserinfoService(s *Service) *UserinfoService {
	rs := &UserinfoService{s: s}
	rs.V2 = NewUserinfoV2Service(s)
	return rs
}

type UserinfoService struct {
	s *Service

	V2 *UserinfoV2Service
}

func NewUserinfoV2Service(s *Service) *UserinfoV2Service {
	rs := &UserinfoV2Service{s: s}
	rs.Me = NewUserinfoV2MeService(s)
	return rs
}

type UserinfoV2Service struct {
	s *Service

	Me *UserinfoV2MeService
}

func NewUserinfoV2MeService(s *Service) *UserinfoV2MeService {
	rs := &UserinfoV2MeService{s: s}
	return rs
}

type UserinfoV2MeService struct {
	s *Service
}

type Jwk struct {
	Keys []*JwkKeys `json:"keys,omitempty"`
}

type JwkKeys struct {
	Alg string `json:"alg,omitempty"`

	E string `json:"e,omitempty"`

	Kid string `json:"kid,omitempty"`

	Kty string `json:"kty,omitempty"`

	N string `json:"n,omitempty"`

	Use string `json:"use,omitempty"`
}

type Tokeninfo struct {
	// Access_type: The access type granted with this token. It can be
	// offline or online.
	Access_type string `json:"access_type,omitempty"`

	// Audience: Who is the intended audience for this token. In general the
	// same as issued_to.
	Audience string `json:"audience,omitempty"`

	// Email: The email address of the user. Present only if the email scope
	// is present in the request.
	Email string `json:"email,omitempty"`

	// Expires_in: The expiry time of the token, as number of seconds left
	// until expiry.
	Expires_in int64 `json:"expires_in,omitempty"`

	// Issued_to: To whom was the token issued to. In general the same as
	// audience.
	Issued_to string `json:"issued_to,omitempty"`

	// Scope: The space separated list of scopes granted to this token.
	Scope string `json:"scope,omitempty"`

	// Token_handle: The token handle associated with this token.
	Token_handle string `json:"token_handle,omitempty"`

	// User_id: The obfuscated user id.
	User_id string `json:"user_id,omitempty"`

	// Verified_email: Boolean flag which is true if the email address is
	// verified. Present only if the email scope is present in the request.
	Verified_email bool `json:"verified_email,omitempty"`
}

type Userinfoplus struct {
	// Email: The user's email address.
	Email string `json:"email,omitempty"`

	// Family_name: The user's last name.
	Family_name string `json:"family_name,omitempty"`

	// Gender: The user's gender.
	Gender string `json:"gender,omitempty"`

	// Given_name: The user's first name.
	Given_name string `json:"given_name,omitempty"`

	// Hd: The hosted domain e.g. example.com if the user is Google apps
	// user.
	Hd string `json:"hd,omitempty"`

	// Id: The obfuscated ID of the user.
	Id string `json:"id,omitempty"`

	// Link: URL of the profile page.
	Link string `json:"link,omitempty"`

	// Locale: The user's preferred locale.
	Locale string `json:"locale,omitempty"`

	// Name: The user's full name.
	Name string `json:"name,omitempty"`

	// Picture: URL of the user's picture image.
	Picture string `json:"picture,omitempty"`

	// Verified_email: Boolean flag which is true if the email address is
	// verified. Always verified because we only return the user's primary
	// email address.
	Verified_email bool `json:"verified_email,omitempty"`
}

// method id "oauth2.getCertForOpenIdConnect":

type GetCertForOpenIdConnectCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetCertForOpenIdConnect:

func (s *Service) GetCertForOpenIdConnect() *GetCertForOpenIdConnectCall {
	return &GetCertForOpenIdConnectCall{
		s:             s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "oauth2/v2/certs",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GetCertForOpenIdConnectCall) Fields(s ...googleapi.Field) *GetCertForOpenIdConnectCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GetCertForOpenIdConnectCall) Do() (*Jwk, error) {
	var returnValue *Jwk
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.getCertForOpenIdConnect",
	//   "path": "oauth2/v2/certs",
	//   "response": {
	//     "$ref": "Jwk"
	//   }
	// }

}

// method id "oauth2.tokeninfo":

type TokeninfoCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Tokeninfo:

func (s *Service) Tokeninfo() *TokeninfoCall {
	return &TokeninfoCall{
		s:             s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "oauth2/v2/tokeninfo",
	}
}

// Access_token sets the optional parameter "access_token":
func (c *TokeninfoCall) Access_token(access_token string) *TokeninfoCall {
	c.params_.Set("access_token", fmt.Sprintf("%v", access_token))
	return c
}

// Id_token sets the optional parameter "id_token":
func (c *TokeninfoCall) Id_token(id_token string) *TokeninfoCall {
	c.params_.Set("id_token", fmt.Sprintf("%v", id_token))
	return c
}

// Token_handle sets the optional parameter "token_handle":
func (c *TokeninfoCall) Token_handle(token_handle string) *TokeninfoCall {
	c.params_.Set("token_handle", fmt.Sprintf("%v", token_handle))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TokeninfoCall) Fields(s ...googleapi.Field) *TokeninfoCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TokeninfoCall) Do() (*Tokeninfo, error) {
	var returnValue *Tokeninfo
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "httpMethod": "POST",
	//   "id": "oauth2.tokeninfo",
	//   "parameters": {
	//     "access_token": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id_token": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "token_handle": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "oauth2/v2/tokeninfo",
	//   "response": {
	//     "$ref": "Tokeninfo"
	//   }
	// }

}

// method id "oauth2.userinfo.get":

type UserinfoGetCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get:

func (r *UserinfoService) Get() *UserinfoGetCall {
	return &UserinfoGetCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "oauth2/v2/userinfo",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserinfoGetCall) Fields(s ...googleapi.Field) *UserinfoGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserinfoGetCall) Do() (*Userinfoplus, error) {
	var returnValue *Userinfoplus
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.userinfo.get",
	//   "path": "oauth2/v2/userinfo",
	//   "response": {
	//     "$ref": "Userinfoplus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.login",
	//     "https://www.googleapis.com/auth/plus.me",
	//     "https://www.googleapis.com/auth/userinfo.email",
	//     "https://www.googleapis.com/auth/userinfo.profile"
	//   ]
	// }

}

// method id "oauth2.userinfo.v2.me.get":

type UserinfoV2MeGetCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get:

func (r *UserinfoV2MeService) Get() *UserinfoV2MeGetCall {
	return &UserinfoV2MeGetCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userinfo/v2/me",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserinfoV2MeGetCall) Fields(s ...googleapi.Field) *UserinfoV2MeGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserinfoV2MeGetCall) Do() (*Userinfoplus, error) {
	var returnValue *Userinfoplus
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "httpMethod": "GET",
	//   "id": "oauth2.userinfo.v2.me.get",
	//   "path": "userinfo/v2/me",
	//   "response": {
	//     "$ref": "Userinfoplus"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/plus.login",
	//     "https://www.googleapis.com/auth/plus.me",
	//     "https://www.googleapis.com/auth/userinfo.email",
	//     "https://www.googleapis.com/auth/userinfo.profile"
	//   ]
	// }

}
