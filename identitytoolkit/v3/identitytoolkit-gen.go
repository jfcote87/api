// Package identitytoolkit provides access to the Google Identity Toolkit API.
//
// See https://developers.google.com/identity-toolkit/v3/
//
// Usage example:
//
//   import "github.com/jfcote87/api/identitytoolkit/v3"
//   ...
//   identitytoolkitService, err := identitytoolkit.New(oauthHttpClient)
package identitytoolkit

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

const apiId = "identitytoolkit:v3"
const apiName = "identitytoolkit"
const apiVersion = "v3"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/identitytoolkit/v3/relyingparty/"}

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Relyingparty = NewRelyingpartyService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Relyingparty *RelyingpartyService
}

func NewRelyingpartyService(s *Service) *RelyingpartyService {
	rs := &RelyingpartyService{s: s}
	return rs
}

type RelyingpartyService struct {
	s *Service
}

type CreateAuthUriResponse struct {
	// AuthUri: The URI used by the IDP to authenticate the user.
	AuthUri string `json:"authUri,omitempty"`

	// CaptchaRequired: True if captcha is required.
	CaptchaRequired bool `json:"captchaRequired,omitempty"`

	// ForExistingProvider: True if the authUri is for user's existing
	// provider.
	ForExistingProvider bool `json:"forExistingProvider,omitempty"`

	// Kind: The fixed string identitytoolkit#CreateAuthUriResponse".
	Kind string `json:"kind,omitempty"`

	// ProviderId: The provider ID of the auth URI.
	ProviderId string `json:"providerId,omitempty"`

	// Registered: Whether the user is registered if the identifier is an
	// email.
	Registered bool `json:"registered,omitempty"`
}

type DeleteAccountResponse struct {
	// Kind: The fixed string "identitytoolkit#DeleteAccountResponse".
	Kind string `json:"kind,omitempty"`
}

type DownloadAccountResponse struct {
	// Kind: The fixed string "identitytoolkit#DownloadAccountResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The next page token. To be used in a subsequent
	// request to return the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Users: The user accounts data.
	Users []*UserInfo `json:"users,omitempty"`
}

type GetAccountInfoResponse struct {
	// Kind: The fixed string "identitytoolkit#GetAccountInfoResponse".
	Kind string `json:"kind,omitempty"`

	// Users: The info of the users.
	Users []*UserInfo `json:"users,omitempty"`
}

type GetOobConfirmationCodeResponse struct {
	// Kind: The fixed string
	// "identitytoolkit#GetOobConfirmationCodeResponse".
	Kind string `json:"kind,omitempty"`

	// OobCode: The code to be send to the user.
	OobCode string `json:"oobCode,omitempty"`
}

type IdentitytoolkitRelyingpartyCreateAuthUriRequest struct {
	// AppId: The app ID of the mobile app, base64(CERT_SHA1):PACKAGE_NAME
	// for Android, BUNDLE_ID for iOS.
	AppId string `json:"appId,omitempty"`

	// ClientId: The relying party OAuth client ID.
	ClientId string `json:"clientId,omitempty"`

	// Context: The opaque value used by the client to maintain context info
	// between the authentication request and the IDP callback.
	Context string `json:"context,omitempty"`

	// ContinueUri: The URI to which the IDP redirects the user after the
	// federated login flow.
	ContinueUri string `json:"continueUri,omitempty"`

	// Identifier: The email or federated ID of the user.
	Identifier string `json:"identifier,omitempty"`

	// OpenidRealm: Optional realm for OpenID protocol. The sub string
	// "scheme://domain:port" of the param "continueUri" is used if this is
	// not set.
	OpenidRealm string `json:"openidRealm,omitempty"`

	// OtaApp: The native app package for OTA installation.
	OtaApp string `json:"otaApp,omitempty"`

	// ProviderId: The IdP ID. For white listed IdPs it's a short domain
	// name e.g. google.com, aol.com, live.net and yahoo.com. For other
	// OpenID IdPs it's the OP identifier.
	ProviderId string `json:"providerId,omitempty"`
}

type IdentitytoolkitRelyingpartyDeleteAccountRequest struct {
	// LocalId: The local ID of the user.
	LocalId string `json:"localId,omitempty"`
}

type IdentitytoolkitRelyingpartyDownloadAccountRequest struct {
	// MaxResults: The max number of results to return in the response.
	MaxResults int64 `json:"maxResults,omitempty"`

	// NextPageToken: The token for the next page. This should be taken from
	// the previous response.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type IdentitytoolkitRelyingpartyGetAccountInfoRequest struct {
	// Email: The list of emails of the users to inquiry.
	Email []string `json:"email,omitempty"`

	// IdToken: The GITKit token of the authenticated user.
	IdToken string `json:"idToken,omitempty"`

	// LocalId: The list of local ID's of the users to inquiry.
	LocalId []string `json:"localId,omitempty"`
}

type IdentitytoolkitRelyingpartyResetPasswordRequest struct {
	// Email: The email address of the user.
	Email string `json:"email,omitempty"`

	// NewPassword: The new password inputted by the user.
	NewPassword string `json:"newPassword,omitempty"`

	// OldPassword: The old password inputted by the user.
	OldPassword string `json:"oldPassword,omitempty"`

	// OobCode: The confirmation code.
	OobCode string `json:"oobCode,omitempty"`
}

type IdentitytoolkitRelyingpartySetAccountInfoRequest struct {
	// CaptchaChallenge: The captcha challenge.
	CaptchaChallenge string `json:"captchaChallenge,omitempty"`

	// CaptchaResponse: Response to the captcha.
	CaptchaResponse string `json:"captchaResponse,omitempty"`

	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// EmailVerified: Mark the email as verified or not.
	EmailVerified bool `json:"emailVerified,omitempty"`

	// IdToken: The GITKit token of the authenticated user.
	IdToken string `json:"idToken,omitempty"`

	// LocalId: The local ID of the user.
	LocalId string `json:"localId,omitempty"`

	// OobCode: The out-of-band code of the change email request.
	OobCode string `json:"oobCode,omitempty"`

	// Password: The new password of the user.
	Password string `json:"password,omitempty"`

	// Provider: The associated IDPs of the user.
	Provider []string `json:"provider,omitempty"`

	// UpgradeToFederatedLogin: Mark the user to upgrade to federated login.
	UpgradeToFederatedLogin bool `json:"upgradeToFederatedLogin,omitempty"`
}

type IdentitytoolkitRelyingpartyUploadAccountRequest struct {
	// HashAlgorithm: The password hash algorithm.
	HashAlgorithm string `json:"hashAlgorithm,omitempty"`

	// MemoryCost: Memory cost for hash calculation. Used by scrypt similar
	// algorithms.
	MemoryCost int64 `json:"memoryCost,omitempty"`

	// Rounds: Rounds for hash calculation. Used by scrypt and similar
	// algorithms.
	Rounds int64 `json:"rounds,omitempty"`

	// SaltSeparator: The salt separator.
	SaltSeparator string `json:"saltSeparator,omitempty"`

	// SignerKey: The key for to hash the password.
	SignerKey string `json:"signerKey,omitempty"`

	// Users: The account info to be stored.
	Users []*UserInfo `json:"users,omitempty"`
}

type IdentitytoolkitRelyingpartyVerifyAssertionRequest struct {
	// PendingIdToken: The GITKit token for the non-trusted IDP pending to
	// be confirmed by the user.
	PendingIdToken string `json:"pendingIdToken,omitempty"`

	// PostBody: The post body if the request is a HTTP POST.
	PostBody string `json:"postBody,omitempty"`

	// RequestUri: The URI to which the IDP redirects the user back. It may
	// contain federated login result params added by the IDP.
	RequestUri string `json:"requestUri,omitempty"`
}

type IdentitytoolkitRelyingpartyVerifyPasswordRequest struct {
	// CaptchaChallenge: The captcha challenge.
	CaptchaChallenge string `json:"captchaChallenge,omitempty"`

	// CaptchaResponse: Response to the captcha.
	CaptchaResponse string `json:"captchaResponse,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// Password: The password inputed by the user.
	Password string `json:"password,omitempty"`

	// PendingIdToken: The GITKit token for the non-trusted IDP, which is to
	// be confirmed by the user.
	PendingIdToken string `json:"pendingIdToken,omitempty"`
}

type Relyingparty struct {
	// CaptchaResp: The recaptcha response from the user.
	CaptchaResp string `json:"captchaResp,omitempty"`

	// Challenge: The recaptcha challenge presented to the user.
	Challenge string `json:"challenge,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// IdToken: The user's Gitkit login token for email change.
	IdToken string `json:"idToken,omitempty"`

	// Kind: The fixed string "identitytoolkit#relyingparty".
	Kind string `json:"kind,omitempty"`

	// NewEmail: The new email if the code is for email change.
	NewEmail string `json:"newEmail,omitempty"`

	// RequestType: The request type.
	RequestType string `json:"requestType,omitempty"`

	// UserIp: The IP address of the user.
	UserIp string `json:"userIp,omitempty"`
}

type ResetPasswordResponse struct {
	// Email: The user's email.
	Email string `json:"email,omitempty"`

	// Kind: The fixed string "identitytoolkit#ResetPasswordResponse".
	Kind string `json:"kind,omitempty"`
}

type SetAccountInfoResponse struct {
	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// IdToken: The Gitkit id token to login the newly sign up user.
	IdToken string `json:"idToken,omitempty"`

	// Kind: The fixed string "identitytoolkit#SetAccountInfoResponse".
	Kind string `json:"kind,omitempty"`

	// ProviderUserInfo: The user's profiles at the associated IdPs.
	ProviderUserInfo []*SetAccountInfoResponseProviderUserInfo `json:"providerUserInfo,omitempty"`
}

type SetAccountInfoResponseProviderUserInfo struct {
	// DisplayName: The user's display name at the IDP.
	DisplayName string `json:"displayName,omitempty"`

	// PhotoUrl: The user's photo url at the IDP.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderId: The IdP ID. For whitelisted IdPs it's a short domain
	// name, e.g., google.com, aol.com, live.net and yahoo.com. For other
	// OpenID IdPs it's the OP identifier.
	ProviderId string `json:"providerId,omitempty"`
}

type UploadAccountResponse struct {
	// Error: The error encountered while processing the account info.
	Error []*UploadAccountResponseError `json:"error,omitempty"`

	// Kind: The fixed string "identitytoolkit#UploadAccountResponse".
	Kind string `json:"kind,omitempty"`
}

type UploadAccountResponseError struct {
	// Index: The index of the malformed account, starting from 0.
	Index int64 `json:"index,omitempty"`

	// Message: Detailed error message for the account info.
	Message string `json:"message,omitempty"`
}

type UserInfo struct {
	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email of the user.
	Email string `json:"email,omitempty"`

	// EmailVerified: Whether the email has been verified.
	EmailVerified bool `json:"emailVerified,omitempty"`

	// LocalId: The local ID of the user.
	LocalId string `json:"localId,omitempty"`

	// PasswordHash: The user's hashed password.
	PasswordHash string `json:"passwordHash,omitempty"`

	// PasswordUpdatedAt: The timestamp when the password was last updated.
	PasswordUpdatedAt float64 `json:"passwordUpdatedAt,omitempty"`

	// PhotoUrl: The URL of the user profile photo.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderUserInfo: The IDP of the user.
	ProviderUserInfo []*UserInfoProviderUserInfo `json:"providerUserInfo,omitempty"`

	// Salt: The user's password salt.
	Salt string `json:"salt,omitempty"`

	// Version: Version of the user's password.
	Version int64 `json:"version,omitempty"`
}

type UserInfoProviderUserInfo struct {
	// DisplayName: The user's display name at the IDP.
	DisplayName string `json:"displayName,omitempty"`

	// FederatedId: User's identifier at IDP.
	FederatedId string `json:"federatedId,omitempty"`

	// PhotoUrl: The user's photo url at the IDP.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderId: The IdP ID. For white listed IdPs it's a short domain
	// name, e.g., google.com, aol.com, live.net and yahoo.com. For other
	// OpenID IdPs it's the OP identifier.
	ProviderId string `json:"providerId,omitempty"`
}

type VerifyAssertionResponse struct {
	// Action: The action code.
	Action string `json:"action,omitempty"`

	// AppInstallationUrl: URL for OTA app installation.
	AppInstallationUrl string `json:"appInstallationUrl,omitempty"`

	// AppScheme: The custom scheme used by mobile app.
	AppScheme string `json:"appScheme,omitempty"`

	// Context: The opaque value used by the client to maintain context info
	// between the authentication request and the IDP callback.
	Context string `json:"context,omitempty"`

	// DateOfBirth: The birth date of the IdP account.
	DateOfBirth string `json:"dateOfBirth,omitempty"`

	// DisplayName: The display name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email returned by the IdP. NOTE: The federated login user
	// may not own the email.
	Email string `json:"email,omitempty"`

	// EmailRecycled: It's true if the email is recycled.
	EmailRecycled bool `json:"emailRecycled,omitempty"`

	// EmailVerified: The value is true if the IDP is also the email
	// provider. It means the user owns the email.
	EmailVerified bool `json:"emailVerified,omitempty"`

	// FederatedId: The unique ID identifies the IdP account.
	FederatedId string `json:"federatedId,omitempty"`

	// FirstName: The first name of the user.
	FirstName string `json:"firstName,omitempty"`

	// FullName: The full name of the user.
	FullName string `json:"fullName,omitempty"`

	// IdToken: The ID token.
	IdToken string `json:"idToken,omitempty"`

	// InputEmail: It's the identifier param in the createAuthUri request if
	// the identifier is an email. It can be used to check whether the user
	// input email is different from the asserted email.
	InputEmail string `json:"inputEmail,omitempty"`

	// Kind: The fixed string "identitytoolkit#VerifyAssertionResponse".
	Kind string `json:"kind,omitempty"`

	// Language: The language preference of the user.
	Language string `json:"language,omitempty"`

	// LastName: The last name of the user.
	LastName string `json:"lastName,omitempty"`

	// LocalId: The RP local ID if it's already been mapped to the IdP
	// account identified by the federated ID.
	LocalId string `json:"localId,omitempty"`

	// NeedConfirmation: Whether the assertion is from a non-trusted IDP and
	// need account linking confirmation.
	NeedConfirmation bool `json:"needConfirmation,omitempty"`

	// NickName: The nick name of the user.
	NickName string `json:"nickName,omitempty"`

	// OauthRequestToken: The user approved request token for the OpenID
	// OAuth extension.
	OauthRequestToken string `json:"oauthRequestToken,omitempty"`

	// OauthScope: The scope for the OpenID OAuth extension.
	OauthScope string `json:"oauthScope,omitempty"`

	// OriginalEmail: The original email stored in the mapping storage. It's
	// returned when the federated ID is associated to a different email.
	OriginalEmail string `json:"originalEmail,omitempty"`

	// PhotoUrl: The URI of the public accessible profiel picture.
	PhotoUrl string `json:"photoUrl,omitempty"`

	// ProviderId: The IdP ID. For white listed IdPs it's a short domain
	// name e.g. google.com, aol.com, live.net and yahoo.com. If the
	// "providerId" param is set to OpenID OP identifer other than the
	// whilte listed IdPs the OP identifier is returned. If the "identifier"
	// param is federated ID in the createAuthUri request. The domain part
	// of the federated ID is returned.
	ProviderId string `json:"providerId,omitempty"`

	// TimeZone: The timezone of the user.
	TimeZone string `json:"timeZone,omitempty"`

	// VerifiedProvider: When action is 'map', contains the idps which can
	// be used for confirmation.
	VerifiedProvider []string `json:"verifiedProvider,omitempty"`
}

type VerifyPasswordResponse struct {
	// DisplayName: The name of the user.
	DisplayName string `json:"displayName,omitempty"`

	// Email: The email returned by the IdP. NOTE: The federated login user
	// may not own the email.
	Email string `json:"email,omitempty"`

	// IdToken: The GITKit token for authenticated user.
	IdToken string `json:"idToken,omitempty"`

	// Kind: The fixed string "identitytoolkit#VerifyPasswordResponse".
	Kind string `json:"kind,omitempty"`

	// LocalId: The RP local ID if it's already been mapped to the IdP
	// account identified by the federated ID.
	LocalId string `json:"localId,omitempty"`

	// PhotoUrl: The URI of the user's photo at IdP
	PhotoUrl string `json:"photoUrl,omitempty"`

	// Registered: Whether the email is registered.
	Registered bool `json:"registered,omitempty"`
}

// method id "identitytoolkit.relyingparty.createAuthUri":

type RelyingpartyCreateAuthUriCall struct {
	s                                               *Service
	identitytoolkitrelyingpartycreateauthurirequest *IdentitytoolkitRelyingpartyCreateAuthUriRequest
	caller_                                         googleapi.Caller
	params_                                         url.Values
	pathTemplate_                                   string
}

// CreateAuthUri: Creates the URI used by the IdP to authenticate the
// user.

func (r *RelyingpartyService) CreateAuthUri(identitytoolkitrelyingpartycreateauthurirequest *IdentitytoolkitRelyingpartyCreateAuthUriRequest) *RelyingpartyCreateAuthUriCall {
	return &RelyingpartyCreateAuthUriCall{
		s: r.s,
		identitytoolkitrelyingpartycreateauthurirequest: identitytoolkitrelyingpartycreateauthurirequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "createAuthUri",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyCreateAuthUriCall) Fields(s ...googleapi.Field) *RelyingpartyCreateAuthUriCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyCreateAuthUriCall) Do() (*CreateAuthUriResponse, error) {
	var returnValue *CreateAuthUriResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartycreateauthurirequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates the URI used by the IdP to authenticate the user.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.createAuthUri",
	//   "path": "createAuthUri",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyCreateAuthUriRequest"
	//   },
	//   "response": {
	//     "$ref": "CreateAuthUriResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.deleteAccount":

type RelyingpartyDeleteAccountCall struct {
	s                                               *Service
	identitytoolkitrelyingpartydeleteaccountrequest *IdentitytoolkitRelyingpartyDeleteAccountRequest
	caller_                                         googleapi.Caller
	params_                                         url.Values
	pathTemplate_                                   string
}

// DeleteAccount: Delete user account.

func (r *RelyingpartyService) DeleteAccount(identitytoolkitrelyingpartydeleteaccountrequest *IdentitytoolkitRelyingpartyDeleteAccountRequest) *RelyingpartyDeleteAccountCall {
	return &RelyingpartyDeleteAccountCall{
		s: r.s,
		identitytoolkitrelyingpartydeleteaccountrequest: identitytoolkitrelyingpartydeleteaccountrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "deleteAccount",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyDeleteAccountCall) Fields(s ...googleapi.Field) *RelyingpartyDeleteAccountCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyDeleteAccountCall) Do() (*DeleteAccountResponse, error) {
	var returnValue *DeleteAccountResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartydeleteaccountrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete user account.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.deleteAccount",
	//   "path": "deleteAccount",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyDeleteAccountRequest"
	//   },
	//   "response": {
	//     "$ref": "DeleteAccountResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.downloadAccount":

type RelyingpartyDownloadAccountCall struct {
	s                                                 *Service
	identitytoolkitrelyingpartydownloadaccountrequest *IdentitytoolkitRelyingpartyDownloadAccountRequest
	caller_                                           googleapi.Caller
	params_                                           url.Values
	pathTemplate_                                     string
}

// DownloadAccount: Batch download user accounts.

func (r *RelyingpartyService) DownloadAccount(identitytoolkitrelyingpartydownloadaccountrequest *IdentitytoolkitRelyingpartyDownloadAccountRequest) *RelyingpartyDownloadAccountCall {
	return &RelyingpartyDownloadAccountCall{
		s: r.s,
		identitytoolkitrelyingpartydownloadaccountrequest: identitytoolkitrelyingpartydownloadaccountrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "downloadAccount",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyDownloadAccountCall) Fields(s ...googleapi.Field) *RelyingpartyDownloadAccountCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyDownloadAccountCall) Do() (*DownloadAccountResponse, error) {
	var returnValue *DownloadAccountResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartydownloadaccountrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Batch download user accounts.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.downloadAccount",
	//   "path": "downloadAccount",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyDownloadAccountRequest"
	//   },
	//   "response": {
	//     "$ref": "DownloadAccountResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.getAccountInfo":

type RelyingpartyGetAccountInfoCall struct {
	s                                                *Service
	identitytoolkitrelyingpartygetaccountinforequest *IdentitytoolkitRelyingpartyGetAccountInfoRequest
	caller_                                          googleapi.Caller
	params_                                          url.Values
	pathTemplate_                                    string
}

// GetAccountInfo: Returns the account info.

func (r *RelyingpartyService) GetAccountInfo(identitytoolkitrelyingpartygetaccountinforequest *IdentitytoolkitRelyingpartyGetAccountInfoRequest) *RelyingpartyGetAccountInfoCall {
	return &RelyingpartyGetAccountInfoCall{
		s: r.s,
		identitytoolkitrelyingpartygetaccountinforequest: identitytoolkitrelyingpartygetaccountinforequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "getAccountInfo",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyGetAccountInfoCall) Fields(s ...googleapi.Field) *RelyingpartyGetAccountInfoCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyGetAccountInfoCall) Do() (*GetAccountInfoResponse, error) {
	var returnValue *GetAccountInfoResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartygetaccountinforequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the account info.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.getAccountInfo",
	//   "path": "getAccountInfo",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyGetAccountInfoRequest"
	//   },
	//   "response": {
	//     "$ref": "GetAccountInfoResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.getOobConfirmationCode":

type RelyingpartyGetOobConfirmationCodeCall struct {
	s             *Service
	relyingparty  *Relyingparty
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetOobConfirmationCode: Get a code for user action confirmation.

func (r *RelyingpartyService) GetOobConfirmationCode(relyingparty *Relyingparty) *RelyingpartyGetOobConfirmationCodeCall {
	return &RelyingpartyGetOobConfirmationCodeCall{
		s:             r.s,
		relyingparty:  relyingparty,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "getOobConfirmationCode",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyGetOobConfirmationCodeCall) Fields(s ...googleapi.Field) *RelyingpartyGetOobConfirmationCodeCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyGetOobConfirmationCodeCall) Do() (*GetOobConfirmationCodeResponse, error) {
	var returnValue *GetOobConfirmationCodeResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.relyingparty,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get a code for user action confirmation.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.getOobConfirmationCode",
	//   "path": "getOobConfirmationCode",
	//   "request": {
	//     "$ref": "Relyingparty"
	//   },
	//   "response": {
	//     "$ref": "GetOobConfirmationCodeResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.getPublicKeys":

type RelyingpartyGetPublicKeysCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetPublicKeys: Get token signing public key.

func (r *RelyingpartyService) GetPublicKeys() *RelyingpartyGetPublicKeysCall {
	return &RelyingpartyGetPublicKeysCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "publicKeys",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyGetPublicKeysCall) Fields(s ...googleapi.Field) *RelyingpartyGetPublicKeysCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyGetPublicKeysCall) Do() (map[string]string, error) {
	var returnValue map[string]string
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get token signing public key.",
	//   "httpMethod": "GET",
	//   "id": "identitytoolkit.relyingparty.getPublicKeys",
	//   "path": "publicKeys",
	//   "response": {
	//     "$ref": "IdentitytoolkitRelyingpartyGetPublicKeysResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.resetPassword":

type RelyingpartyResetPasswordCall struct {
	s                                               *Service
	identitytoolkitrelyingpartyresetpasswordrequest *IdentitytoolkitRelyingpartyResetPasswordRequest
	caller_                                         googleapi.Caller
	params_                                         url.Values
	pathTemplate_                                   string
}

// ResetPassword: Reset password for a user.

func (r *RelyingpartyService) ResetPassword(identitytoolkitrelyingpartyresetpasswordrequest *IdentitytoolkitRelyingpartyResetPasswordRequest) *RelyingpartyResetPasswordCall {
	return &RelyingpartyResetPasswordCall{
		s: r.s,
		identitytoolkitrelyingpartyresetpasswordrequest: identitytoolkitrelyingpartyresetpasswordrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "resetPassword",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyResetPasswordCall) Fields(s ...googleapi.Field) *RelyingpartyResetPasswordCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyResetPasswordCall) Do() (*ResetPasswordResponse, error) {
	var returnValue *ResetPasswordResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartyresetpasswordrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Reset password for a user.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.resetPassword",
	//   "path": "resetPassword",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyResetPasswordRequest"
	//   },
	//   "response": {
	//     "$ref": "ResetPasswordResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.setAccountInfo":

type RelyingpartySetAccountInfoCall struct {
	s                                                *Service
	identitytoolkitrelyingpartysetaccountinforequest *IdentitytoolkitRelyingpartySetAccountInfoRequest
	caller_                                          googleapi.Caller
	params_                                          url.Values
	pathTemplate_                                    string
}

// SetAccountInfo: Set account info for a user.

func (r *RelyingpartyService) SetAccountInfo(identitytoolkitrelyingpartysetaccountinforequest *IdentitytoolkitRelyingpartySetAccountInfoRequest) *RelyingpartySetAccountInfoCall {
	return &RelyingpartySetAccountInfoCall{
		s: r.s,
		identitytoolkitrelyingpartysetaccountinforequest: identitytoolkitrelyingpartysetaccountinforequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "setAccountInfo",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartySetAccountInfoCall) Fields(s ...googleapi.Field) *RelyingpartySetAccountInfoCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartySetAccountInfoCall) Do() (*SetAccountInfoResponse, error) {
	var returnValue *SetAccountInfoResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartysetaccountinforequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Set account info for a user.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.setAccountInfo",
	//   "path": "setAccountInfo",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartySetAccountInfoRequest"
	//   },
	//   "response": {
	//     "$ref": "SetAccountInfoResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.uploadAccount":

type RelyingpartyUploadAccountCall struct {
	s                                               *Service
	identitytoolkitrelyingpartyuploadaccountrequest *IdentitytoolkitRelyingpartyUploadAccountRequest
	caller_                                         googleapi.Caller
	params_                                         url.Values
	pathTemplate_                                   string
}

// UploadAccount: Batch upload existing user accounts.

func (r *RelyingpartyService) UploadAccount(identitytoolkitrelyingpartyuploadaccountrequest *IdentitytoolkitRelyingpartyUploadAccountRequest) *RelyingpartyUploadAccountCall {
	return &RelyingpartyUploadAccountCall{
		s: r.s,
		identitytoolkitrelyingpartyuploadaccountrequest: identitytoolkitrelyingpartyuploadaccountrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "uploadAccount",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyUploadAccountCall) Fields(s ...googleapi.Field) *RelyingpartyUploadAccountCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyUploadAccountCall) Do() (*UploadAccountResponse, error) {
	var returnValue *UploadAccountResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartyuploadaccountrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Batch upload existing user accounts.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.uploadAccount",
	//   "path": "uploadAccount",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyUploadAccountRequest"
	//   },
	//   "response": {
	//     "$ref": "UploadAccountResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.verifyAssertion":

type RelyingpartyVerifyAssertionCall struct {
	s                                                 *Service
	identitytoolkitrelyingpartyverifyassertionrequest *IdentitytoolkitRelyingpartyVerifyAssertionRequest
	caller_                                           googleapi.Caller
	params_                                           url.Values
	pathTemplate_                                     string
}

// VerifyAssertion: Verifies the assertion returned by the IdP.

func (r *RelyingpartyService) VerifyAssertion(identitytoolkitrelyingpartyverifyassertionrequest *IdentitytoolkitRelyingpartyVerifyAssertionRequest) *RelyingpartyVerifyAssertionCall {
	return &RelyingpartyVerifyAssertionCall{
		s: r.s,
		identitytoolkitrelyingpartyverifyassertionrequest: identitytoolkitrelyingpartyverifyassertionrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "verifyAssertion",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyVerifyAssertionCall) Fields(s ...googleapi.Field) *RelyingpartyVerifyAssertionCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyVerifyAssertionCall) Do() (*VerifyAssertionResponse, error) {
	var returnValue *VerifyAssertionResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartyverifyassertionrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Verifies the assertion returned by the IdP.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.verifyAssertion",
	//   "path": "verifyAssertion",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyVerifyAssertionRequest"
	//   },
	//   "response": {
	//     "$ref": "VerifyAssertionResponse"
	//   }
	// }

}

// method id "identitytoolkit.relyingparty.verifyPassword":

type RelyingpartyVerifyPasswordCall struct {
	s                                                *Service
	identitytoolkitrelyingpartyverifypasswordrequest *IdentitytoolkitRelyingpartyVerifyPasswordRequest
	caller_                                          googleapi.Caller
	params_                                          url.Values
	pathTemplate_                                    string
}

// VerifyPassword: Verifies the user entered password.

func (r *RelyingpartyService) VerifyPassword(identitytoolkitrelyingpartyverifypasswordrequest *IdentitytoolkitRelyingpartyVerifyPasswordRequest) *RelyingpartyVerifyPasswordCall {
	return &RelyingpartyVerifyPasswordCall{
		s: r.s,
		identitytoolkitrelyingpartyverifypasswordrequest: identitytoolkitrelyingpartyverifypasswordrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "verifyPassword",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RelyingpartyVerifyPasswordCall) Fields(s ...googleapi.Field) *RelyingpartyVerifyPasswordCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RelyingpartyVerifyPasswordCall) Do() (*VerifyPasswordResponse, error) {
	var returnValue *VerifyPasswordResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.identitytoolkitrelyingpartyverifypasswordrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Verifies the user entered password.",
	//   "httpMethod": "POST",
	//   "id": "identitytoolkit.relyingparty.verifyPassword",
	//   "path": "verifyPassword",
	//   "request": {
	//     "$ref": "IdentitytoolkitRelyingpartyVerifyPasswordRequest"
	//   },
	//   "response": {
	//     "$ref": "VerifyPasswordResponse"
	//   }
	// }

}
