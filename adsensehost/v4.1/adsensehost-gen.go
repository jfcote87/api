// Package adsensehost provides access to the AdSense Host API.
//
// See https://developers.google.com/adsense/host/
//
// Usage example:
//
//   import "github.com/jfcote87/api/adsensehost/v4.1"
//   ...
//   adsensehostService, err := adsensehost.New(oauthHttpClient)
package adsensehost

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

const apiId = "adsensehost:v4.1"
const apiName = "adsensehost"
const apiVersion = "v4.1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/adsensehost/v4.1/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your AdSense host data and associated accounts
	AdsensehostScope = "https://www.googleapis.com/auth/adsensehost"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Accounts = NewAccountsService(s)
	s.Adclients = NewAdclientsService(s)
	s.Associationsessions = NewAssociationsessionsService(s)
	s.Customchannels = NewCustomchannelsService(s)
	s.Reports = NewReportsService(s)
	s.Urlchannels = NewUrlchannelsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Accounts *AccountsService

	Adclients *AdclientsService

	Associationsessions *AssociationsessionsService

	Customchannels *CustomchannelsService

	Reports *ReportsService

	Urlchannels *UrlchannelsService
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	rs.Adclients = NewAccountsAdclientsService(s)
	rs.Adunits = NewAccountsAdunitsService(s)
	rs.Reports = NewAccountsReportsService(s)
	return rs
}

type AccountsService struct {
	s *Service

	Adclients *AccountsAdclientsService

	Adunits *AccountsAdunitsService

	Reports *AccountsReportsService
}

func NewAccountsAdclientsService(s *Service) *AccountsAdclientsService {
	rs := &AccountsAdclientsService{s: s}
	return rs
}

type AccountsAdclientsService struct {
	s *Service
}

func NewAccountsAdunitsService(s *Service) *AccountsAdunitsService {
	rs := &AccountsAdunitsService{s: s}
	return rs
}

type AccountsAdunitsService struct {
	s *Service
}

func NewAccountsReportsService(s *Service) *AccountsReportsService {
	rs := &AccountsReportsService{s: s}
	return rs
}

type AccountsReportsService struct {
	s *Service
}

func NewAdclientsService(s *Service) *AdclientsService {
	rs := &AdclientsService{s: s}
	return rs
}

type AdclientsService struct {
	s *Service
}

func NewAssociationsessionsService(s *Service) *AssociationsessionsService {
	rs := &AssociationsessionsService{s: s}
	return rs
}

type AssociationsessionsService struct {
	s *Service
}

func NewCustomchannelsService(s *Service) *CustomchannelsService {
	rs := &CustomchannelsService{s: s}
	return rs
}

type CustomchannelsService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	return rs
}

type ReportsService struct {
	s *Service
}

func NewUrlchannelsService(s *Service) *UrlchannelsService {
	rs := &UrlchannelsService{s: s}
	return rs
}

type UrlchannelsService struct {
	s *Service
}

type Account struct {
	// Id: Unique identifier of this account.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsensehost#account.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this account.
	Name string `json:"name,omitempty"`

	// Status: Approval status of this account. One of: PENDING, APPROVED,
	// DISABLED.
	Status string `json:"status,omitempty"`
}

type Accounts struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The accounts returned in this list response.
	Items []*Account `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsensehost#accounts.
	Kind string `json:"kind,omitempty"`
}

type AdClient struct {
	// ArcOptIn: Whether this ad client is opted in to ARC.
	ArcOptIn bool `json:"arcOptIn,omitempty"`

	// Id: Unique identifier of this ad client.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsensehost#adClient.
	Kind string `json:"kind,omitempty"`

	// ProductCode: This ad client's product code, which corresponds to the
	// PRODUCT_CODE report dimension.
	ProductCode string `json:"productCode,omitempty"`

	// SupportsReporting: Whether this ad client supports being reported on.
	SupportsReporting bool `json:"supportsReporting,omitempty"`
}

type AdClients struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The ad clients returned in this list response.
	Items []*AdClient `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsensehost#adClients.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through ad clients. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AdCode struct {
	// AdCode: The ad code snippet.
	AdCode string `json:"adCode,omitempty"`

	// Kind: Kind this is, in this case adsensehost#adCode.
	Kind string `json:"kind,omitempty"`
}

type AdStyle struct {
	// Colors: The colors included in the style. These are represented as
	// six hexadecimal characters, similar to HTML color codes, but without
	// the leading hash.
	Colors *AdStyleColors `json:"colors,omitempty"`

	// Corners: The style of the corners in the ad. Possible values are
	// SQUARE, SLIGHTLY_ROUNDED and VERY_ROUNDED.
	Corners string `json:"corners,omitempty"`

	// Font: The font which is included in the style.
	Font *AdStyleFont `json:"font,omitempty"`

	// Kind: Kind this is, in this case adsensehost#adStyle.
	Kind string `json:"kind,omitempty"`
}

type AdStyleColors struct {
	// Background: The color of the ad background.
	Background string `json:"background,omitempty"`

	// Border: The color of the ad border.
	Border string `json:"border,omitempty"`

	// Text: The color of the ad text.
	Text string `json:"text,omitempty"`

	// Title: The color of the ad title.
	Title string `json:"title,omitempty"`

	// Url: The color of the ad url.
	Url string `json:"url,omitempty"`
}

type AdStyleFont struct {
	// Family: The family of the font. Possible values are:
	// ACCOUNT_DEFAULT_FAMILY, ADSENSE_DEFAULT_FAMILY, ARIAL, TIMES and
	// VERDANA.
	Family string `json:"family,omitempty"`

	// Size: The size of the font. Possible values are:
	// ACCOUNT_DEFAULT_SIZE, ADSENSE_DEFAULT_SIZE, SMALL, MEDIUM and LARGE.
	Size string `json:"size,omitempty"`
}

type AdUnit struct {
	// Code: Identity code of this ad unit, not necessarily unique across ad
	// clients.
	Code string `json:"code,omitempty"`

	// ContentAdsSettings: Settings specific to content ads (AFC) and
	// highend mobile content ads (AFMC).
	ContentAdsSettings *AdUnitContentAdsSettings `json:"contentAdsSettings,omitempty"`

	// CustomStyle: Custom style information specific to this ad unit.
	CustomStyle *AdStyle `json:"customStyle,omitempty"`

	// Id: Unique identifier of this ad unit. This should be considered an
	// opaque identifier; it is not safe to rely on it being in any
	// particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsensehost#adUnit.
	Kind string `json:"kind,omitempty"`

	// MobileContentAdsSettings: Settings specific to WAP mobile content ads
	// (AFMC).
	MobileContentAdsSettings *AdUnitMobileContentAdsSettings `json:"mobileContentAdsSettings,omitempty"`

	// Name: Name of this ad unit.
	Name string `json:"name,omitempty"`

	// Status: Status of this ad unit. Possible values are:
	// NEW: Indicates
	// that the ad unit was created within the last seven days and does not
	// yet have any activity associated with it.
	//
	// ACTIVE: Indicates that
	// there has been activity on this ad unit in the last seven
	// days.
	//
	// INACTIVE: Indicates that there has been no activity on this ad
	// unit in the last seven days.
	Status string `json:"status,omitempty"`
}

type AdUnitContentAdsSettings struct {
	// BackupOption: The backup option to be used in instances where no ad
	// is available.
	BackupOption *AdUnitContentAdsSettingsBackupOption `json:"backupOption,omitempty"`

	// Size: Size of this ad unit. Size values are in the form
	// SIZE_{width}_{height}.
	Size string `json:"size,omitempty"`

	// Type: Type of this ad unit. Possible values are TEXT, TEXT_IMAGE,
	// IMAGE and LINK.
	Type string `json:"type,omitempty"`
}

type AdUnitContentAdsSettingsBackupOption struct {
	// Color: Color to use when type is set to COLOR. These are represented
	// as six hexadecimal characters, similar to HTML color codes, but
	// without the leading hash.
	Color string `json:"color,omitempty"`

	// Type: Type of the backup option. Possible values are BLANK, COLOR and
	// URL.
	Type string `json:"type,omitempty"`

	// Url: URL to use when type is set to URL.
	Url string `json:"url,omitempty"`
}

type AdUnitMobileContentAdsSettings struct {
	// MarkupLanguage: The markup language to use for this ad unit.
	MarkupLanguage string `json:"markupLanguage,omitempty"`

	// ScriptingLanguage: The scripting language to use for this ad unit.
	ScriptingLanguage string `json:"scriptingLanguage,omitempty"`

	// Size: Size of this ad unit.
	Size string `json:"size,omitempty"`

	// Type: Type of this ad unit.
	Type string `json:"type,omitempty"`
}

type AdUnits struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The ad units returned in this list response.
	Items []*AdUnit `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsensehost#adUnits.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through ad units. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AssociationSession struct {
	// AccountId: Hosted account id of the associated publisher after
	// association. Present if status is ACCEPTED.
	AccountId string `json:"accountId,omitempty"`

	// Id: Unique identifier of this association session.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adsensehost#associationSession.
	Kind string `json:"kind,omitempty"`

	// ProductCodes: The products to associate with the user. Options: AFC,
	// AFF, AFS, AFMC
	ProductCodes []string `json:"productCodes,omitempty"`

	// RedirectUrl: Redirect URL of this association session. Used to
	// redirect users into the AdSense association flow.
	RedirectUrl string `json:"redirectUrl,omitempty"`

	// Status: Status of the completed association, available once the
	// association callback token has been verified. One of ACCEPTED,
	// REJECTED, or ERROR.
	Status string `json:"status,omitempty"`

	// UserLocale: The preferred locale of the user themselves when going
	// through the AdSense association flow.
	UserLocale string `json:"userLocale,omitempty"`

	// WebsiteLocale: The locale of the user's hosted website.
	WebsiteLocale string `json:"websiteLocale,omitempty"`

	// WebsiteUrl: The URL of the user's hosted website.
	WebsiteUrl string `json:"websiteUrl,omitempty"`
}

type CustomChannel struct {
	// Code: Code of this custom channel, not necessarily unique across ad
	// clients.
	Code string `json:"code,omitempty"`

	// Id: Unique identifier of this custom channel. This should be
	// considered an opaque identifier; it is not safe to rely on it being
	// in any particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adsensehost#customChannel.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this custom channel.
	Name string `json:"name,omitempty"`
}

type CustomChannels struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The custom channels returned in this list response.
	Items []*CustomChannel `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsensehost#customChannels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through custom
	// channels. To retrieve the next page of results, set the next
	// request's "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Report struct {
	// Averages: The averages of the report. This is the same length as any
	// other row in the report; cells corresponding to dimension columns are
	// empty.
	Averages []string `json:"averages,omitempty"`

	// Headers: The header information of the columns requested in the
	// report. This is a list of headers; one for each dimension in the
	// request, followed by one for each metric in the request.
	Headers []*ReportHeaders `json:"headers,omitempty"`

	// Kind: Kind this is, in this case adsensehost#report.
	Kind string `json:"kind,omitempty"`

	// Rows: The output rows of the report. Each row is a list of cells; one
	// for each dimension in the request, followed by one for each metric in
	// the request. The dimension cells contain strings, and the metric
	// cells contain numbers.
	Rows [][]string `json:"rows,omitempty"`

	// TotalMatchedRows: The total number of rows matched by the report
	// request. Fewer rows may be returned in the response due to being
	// limited by the row count requested or the report row limit.
	TotalMatchedRows int64 `json:"totalMatchedRows,omitempty,string"`

	// Totals: The totals of the report. This is the same length as any
	// other row in the report; cells corresponding to dimension columns are
	// empty.
	Totals []string `json:"totals,omitempty"`

	// Warnings: Any warnings associated with generation of the report.
	Warnings []string `json:"warnings,omitempty"`
}

type ReportHeaders struct {
	// Currency: The currency of this column. Only present if the header
	// type is METRIC_CURRENCY.
	Currency string `json:"currency,omitempty"`

	// Name: The name of the header.
	Name string `json:"name,omitempty"`

	// Type: The type of the header; one of DIMENSION, METRIC_TALLY,
	// METRIC_RATIO, or METRIC_CURRENCY.
	Type string `json:"type,omitempty"`
}

type UrlChannel struct {
	// Id: Unique identifier of this URL channel. This should be considered
	// an opaque identifier; it is not safe to rely on it being in any
	// particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adsensehost#urlChannel.
	Kind string `json:"kind,omitempty"`

	// UrlPattern: URL Pattern of this URL channel. Does not include
	// "http://" or "https://". Example: www.example.com/home
	UrlPattern string `json:"urlPattern,omitempty"`
}

type UrlChannels struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The URL channels returned in this list response.
	Items []*UrlChannel `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adsensehost#urlChannels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through URL channels.
	// To retrieve the next page of results, set the next request's
	// "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "adsensehost.accounts.get":

type AccountsGetCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about the selected associated AdSense account.

func (r *AccountsService) Get(accountId string) *AccountsGetCall {
	return &AccountsGetCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsGetCall) Fields(s ...googleapi.Field) *AccountsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsGetCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get information about the selected associated AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.get",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to get information about.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}",
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.list":

type AccountsListCall struct {
	s                *Service
	filterAdClientId []string
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
}

// List: List hosted accounts associated with this AdSense account by ad
// client id.

func (r *AccountsService) List(filterAdClientId []string) *AccountsListCall {
	return &AccountsListCall{
		s:                r.s,
		filterAdClientId: filterAdClientId,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "accounts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsListCall) Fields(s ...googleapi.Field) *AccountsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsListCall) Do() (*Accounts, error) {
	var returnValue *Accounts
	c.params_.Del("filterAdClientId")
	for _, v := range c.filterAdClientId {
		c.params_.Add("filterAdClientId", fmt.Sprintf("%v", v))
	}
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List hosted accounts associated with this AdSense account by ad client id.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.list",
	//   "parameterOrder": [
	//     "filterAdClientId"
	//   ],
	//   "parameters": {
	//     "filterAdClientId": {
	//       "description": "Ad clients to list accounts for.",
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "Accounts"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adclients.get":

type AccountsAdclientsGetCall struct {
	s             *Service
	accountId     string
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about one of the ad clients in the specified
// publisher's AdSense account.

func (r *AccountsAdclientsService) Get(accountId string, adClientId string) *AccountsAdclientsGetCall {
	return &AccountsAdclientsGetCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdclientsGetCall) Fields(s ...googleapi.Field) *AccountsAdclientsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdclientsGetCall) Do() (*AdClient, error) {
	var returnValue *AdClient
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get information about one of the ad clients in the specified publisher's AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.adclients.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad client.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}",
	//   "response": {
	//     "$ref": "AdClient"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adclients.list":

type AccountsAdclientsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all hosted ad clients in the specified hosted account.

func (r *AccountsAdclientsService) List(accountId string) *AccountsAdclientsListCall {
	return &AccountsAdclientsListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad clients to include in the response, used for paging.
func (c *AccountsAdclientsListCall) MaxResults(maxResults int64) *AccountsAdclientsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad clients. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsAdclientsListCall) PageToken(pageToken string) *AccountsAdclientsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdclientsListCall) Fields(s ...googleapi.Field) *AccountsAdclientsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdclientsListCall) Do() (*AdClients, error) {
	var returnValue *AdClients
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all hosted ad clients in the specified hosted account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.adclients.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account for which to list ad clients.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad clients to include in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad clients. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients",
	//   "response": {
	//     "$ref": "AdClients"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.delete":

type AccountsAdunitsDeleteCall struct {
	s             *Service
	accountId     string
	adClientId    string
	adUnitId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete the specified ad unit from the specified publisher
// AdSense account.

func (r *AccountsAdunitsService) Delete(accountId string, adClientId string, adUnitId string) *AccountsAdunitsDeleteCall {
	return &AccountsAdunitsDeleteCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		adUnitId:      adUnitId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsDeleteCall) Fields(s ...googleapi.Field) *AccountsAdunitsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsDeleteCall) Do() (*AdUnit, error) {
	var returnValue *AdUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete the specified ad unit from the specified publisher AdSense account.",
	//   "httpMethod": "DELETE",
	//   "id": "adsensehost.accounts.adunits.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to get ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}",
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.get":

type AccountsAdunitsGetCall struct {
	s             *Service
	accountId     string
	adClientId    string
	adUnitId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get the specified host ad unit in this AdSense account.

func (r *AccountsAdunitsService) Get(accountId string, adClientId string, adUnitId string) *AccountsAdunitsGetCall {
	return &AccountsAdunitsGetCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		adUnitId:      adUnitId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsGetCall) Fields(s ...googleapi.Field) *AccountsAdunitsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsGetCall) Do() (*AdUnit, error) {
	var returnValue *AdUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get the specified host ad unit in this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.adunits.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to get ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}",
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.getAdCode":

type AccountsAdunitsGetAdCodeCall struct {
	s             *Service
	accountId     string
	adClientId    string
	adUnitId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetAdCode: Get ad code for the specified ad unit, attaching the
// specified host custom channels.

func (r *AccountsAdunitsService) GetAdCode(accountId string, adClientId string, adUnitId string) *AccountsAdunitsGetAdCodeCall {
	return &AccountsAdunitsGetAdCodeCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		adUnitId:      adUnitId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}/adcode",
	}
}

// HostCustomChannelId sets the optional parameter
// "hostCustomChannelId": Host custom channel to attach to the ad code.
func (c *AccountsAdunitsGetAdCodeCall) HostCustomChannelId(hostCustomChannelId ...string) *AccountsAdunitsGetAdCodeCall {
	c.params_["hostCustomChannelId"] = hostCustomChannelId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsGetAdCodeCall) Fields(s ...googleapi.Field) *AccountsAdunitsGetAdCodeCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsGetAdCodeCall) Do() (*AdCode, error) {
	var returnValue *AdCode
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
		"adUnitId":   c.adUnitId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get ad code for the specified ad unit, attaching the specified host custom channels.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.adunits.getAdCode",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad client.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client with contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit to get the code for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "hostCustomChannelId": {
	//       "description": "Host custom channel to attach to the ad code.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits/{adUnitId}/adcode",
	//   "response": {
	//     "$ref": "AdCode"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.insert":

type AccountsAdunitsInsertCall struct {
	s             *Service
	accountId     string
	adClientId    string
	adunit        *AdUnit
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Insert the supplied ad unit into the specified publisher
// AdSense account.

func (r *AccountsAdunitsService) Insert(accountId string, adClientId string, adunit *AdUnit) *AccountsAdunitsInsertCall {
	return &AccountsAdunitsInsertCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		adunit:        adunit,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsInsertCall) Fields(s ...googleapi.Field) *AccountsAdunitsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsInsertCall) Do() (*AdUnit, error) {
	var returnValue *AdUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.adunit,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Insert the supplied ad unit into the specified publisher AdSense account.",
	//   "httpMethod": "POST",
	//   "id": "adsensehost.accounts.adunits.insert",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which will contain the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client into which to insert the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits",
	//   "request": {
	//     "$ref": "AdUnit"
	//   },
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.list":

type AccountsAdunitsListCall struct {
	s             *Service
	accountId     string
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all ad units in the specified publisher's AdSense account.

func (r *AccountsAdunitsService) List(accountId string, adClientId string) *AccountsAdunitsListCall {
	return &AccountsAdunitsListCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits",
	}
}

// IncludeInactive sets the optional parameter "includeInactive":
// Whether to include inactive ad units. Default: true.
func (c *AccountsAdunitsListCall) IncludeInactive(includeInactive bool) *AccountsAdunitsListCall {
	c.params_.Set("includeInactive", fmt.Sprintf("%v", includeInactive))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad units to include in the response, used for paging.
func (c *AccountsAdunitsListCall) MaxResults(maxResults int64) *AccountsAdunitsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad units. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsAdunitsListCall) PageToken(pageToken string) *AccountsAdunitsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsListCall) Fields(s ...googleapi.Field) *AccountsAdunitsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsListCall) Do() (*AdUnits, error) {
	var returnValue *AdUnits
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all ad units in the specified publisher's AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.adunits.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad client.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client for which to list ad units.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "includeInactive": {
	//       "description": "Whether to include inactive ad units. Default: true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of ad units to include in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad units. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits",
	//   "response": {
	//     "$ref": "AdUnits"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.patch":

type AccountsAdunitsPatchCall struct {
	s             *Service
	accountId     string
	adClientId    string
	adUnitId      string
	adunit        *AdUnit
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update the supplied ad unit in the specified publisher AdSense
// account. This method supports patch semantics.

func (r *AccountsAdunitsService) Patch(accountId string, adClientId string, adUnitId string, adunit *AdUnit) *AccountsAdunitsPatchCall {
	return &AccountsAdunitsPatchCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		adUnitId:      adUnitId,
		adunit:        adunit,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsPatchCall) Fields(s ...googleapi.Field) *AccountsAdunitsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsPatchCall) Do() (*AdUnit, error) {
	var returnValue *AdUnit
	c.params_.Set("adUnitId", fmt.Sprintf("%v", c.adUnitId))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.adunit,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the supplied ad unit in the specified publisher AdSense account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "adsensehost.accounts.adunits.patch",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "adUnitId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad client.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client which contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adUnitId": {
	//       "description": "Ad unit to get.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits",
	//   "request": {
	//     "$ref": "AdUnit"
	//   },
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.adunits.update":

type AccountsAdunitsUpdateCall struct {
	s             *Service
	accountId     string
	adClientId    string
	adunit        *AdUnit
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update the supplied ad unit in the specified publisher
// AdSense account.

func (r *AccountsAdunitsService) Update(accountId string, adClientId string, adunit *AdUnit) *AccountsAdunitsUpdateCall {
	return &AccountsAdunitsUpdateCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		adunit:        adunit,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/adunits",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAdunitsUpdateCall) Fields(s ...googleapi.Field) *AccountsAdunitsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAdunitsUpdateCall) Do() (*AdUnit, error) {
	var returnValue *AdUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":  c.accountId,
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.adunit,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the supplied ad unit in the specified publisher AdSense account.",
	//   "httpMethod": "PUT",
	//   "id": "adsensehost.accounts.adunits.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which contains the ad client.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client which contains the ad unit.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/adunits",
	//   "request": {
	//     "$ref": "AdUnit"
	//   },
	//   "response": {
	//     "$ref": "AdUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.accounts.reports.generate":

type AccountsReportsGenerateCall struct {
	s             *Service
	accountId     string
	startDate     string
	endDate       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generate: Generate an AdSense report based on the report request sent
// in the query parameters. Returns the result as JSON; to retrieve
// output in CSV format specify "alt=csv" as a query parameter.

func (r *AccountsReportsService) Generate(accountId string, startDate string, endDate string) *AccountsReportsGenerateCall {
	return &AccountsReportsGenerateCall{
		s:             r.s,
		accountId:     accountId,
		startDate:     startDate,
		endDate:       endDate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/reports",
	}
}

// Dimension sets the optional parameter "dimension": Dimensions to base
// the report on.
func (c *AccountsReportsGenerateCall) Dimension(dimension ...string) *AccountsReportsGenerateCall {
	c.params_["dimension"] = dimension
	return c
}

// Filter sets the optional parameter "filter": Filters to be run on the
// report.
func (c *AccountsReportsGenerateCall) Filter(filter ...string) *AccountsReportsGenerateCall {
	c.params_["filter"] = filter
	return c
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *AccountsReportsGenerateCall) Locale(locale string) *AccountsReportsGenerateCall {
	c.params_.Set("locale", fmt.Sprintf("%v", locale))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *AccountsReportsGenerateCall) MaxResults(maxResults int64) *AccountsReportsGenerateCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// Metric sets the optional parameter "metric": Numeric columns to
// include in the report.
func (c *AccountsReportsGenerateCall) Metric(metric ...string) *AccountsReportsGenerateCall {
	c.params_["metric"] = metric
	return c
}

// Sort sets the optional parameter "sort": The name of a dimension or
// metric to sort the resulting report on, optionally prefixed with "+"
// to sort ascending or "-" to sort descending. If no prefix is
// specified, the column is sorted ascending.
func (c *AccountsReportsGenerateCall) Sort(sort ...string) *AccountsReportsGenerateCall {
	c.params_["sort"] = sort
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *AccountsReportsGenerateCall) StartIndex(startIndex int64) *AccountsReportsGenerateCall {
	c.params_.Set("startIndex", fmt.Sprintf("%v", startIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsReportsGenerateCall) Fields(s ...googleapi.Field) *AccountsReportsGenerateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsReportsGenerateCall) Do() (*Report, error) {
	var returnValue *Report
	c.params_.Set("endDate", fmt.Sprintf("%v", c.endDate))
	c.params_.Set("startDate", fmt.Sprintf("%v", c.startDate))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generate an AdSense report based on the report request sent in the query parameters. Returns the result as JSON; to retrieve output in CSV format specify \"alt=csv\" as a query parameter.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.accounts.reports.generate",
	//   "parameterOrder": [
	//     "accountId",
	//     "startDate",
	//     "endDate"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Hosted account upon which to report.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "dimension": {
	//       "description": "Dimensions to base the report on.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "End of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Filters to be run on the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+(==|=@).+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "Optional locale to use for translating report output to a local language. Defaults to \"en_US\" if not specified.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows of report data to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "metric": {
	//       "description": "Numeric columns to include in the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The name of a dimension or metric to sort the resulting report on, optionally prefixed with \"+\" to sort ascending or \"-\" to sort descending. If no prefix is specified, the column is sorted ascending.",
	//       "location": "query",
	//       "pattern": "(\\+|-)?[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Start of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{accountId}/reports",
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.adclients.get":

type AdclientsGetCall struct {
	s             *Service
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about one of the ad clients in the Host AdSense
// account.

func (r *AdclientsService) Get(adClientId string) *AdclientsGetCall {
	return &AdclientsGetCall{
		s:             r.s,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdclientsGetCall) Fields(s ...googleapi.Field) *AdclientsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdclientsGetCall) Do() (*AdClient, error) {
	var returnValue *AdClient
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get information about one of the ad clients in the Host AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.adclients.get",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}",
	//   "response": {
	//     "$ref": "AdClient"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.adclients.list":

type AdclientsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all host ad clients in this AdSense account.

func (r *AdclientsService) List() *AdclientsListCall {
	return &AdclientsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of ad clients to include in the response, used for paging.
func (c *AdclientsListCall) MaxResults(maxResults int64) *AdclientsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through ad clients. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AdclientsListCall) PageToken(pageToken string) *AdclientsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdclientsListCall) Fields(s ...googleapi.Field) *AdclientsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdclientsListCall) Do() (*AdClients, error) {
	var returnValue *AdClients
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all host ad clients in this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.adclients.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of ad clients to include in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through ad clients. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients",
	//   "response": {
	//     "$ref": "AdClients"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.associationsessions.start":

type AssociationsessionsStartCall struct {
	s             *Service
	productCode   []string
	websiteUrl    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Start: Create an association session for initiating an association
// with an AdSense user.

func (r *AssociationsessionsService) Start(productCode []string, websiteUrl string) *AssociationsessionsStartCall {
	return &AssociationsessionsStartCall{
		s:             r.s,
		productCode:   productCode,
		websiteUrl:    websiteUrl,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "associationsessions/start",
	}
}

// UserLocale sets the optional parameter "userLocale": The preferred
// locale of the user.
func (c *AssociationsessionsStartCall) UserLocale(userLocale string) *AssociationsessionsStartCall {
	c.params_.Set("userLocale", fmt.Sprintf("%v", userLocale))
	return c
}

// WebsiteLocale sets the optional parameter "websiteLocale": The locale
// of the user's hosted website.
func (c *AssociationsessionsStartCall) WebsiteLocale(websiteLocale string) *AssociationsessionsStartCall {
	c.params_.Set("websiteLocale", fmt.Sprintf("%v", websiteLocale))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssociationsessionsStartCall) Fields(s ...googleapi.Field) *AssociationsessionsStartCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AssociationsessionsStartCall) Do() (*AssociationSession, error) {
	var returnValue *AssociationSession
	c.params_.Set("websiteUrl", fmt.Sprintf("%v", c.websiteUrl))
	c.params_.Del("productCode")
	for _, v := range c.productCode {
		c.params_.Add("productCode", fmt.Sprintf("%v", v))
	}
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create an association session for initiating an association with an AdSense user.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.associationsessions.start",
	//   "parameterOrder": [
	//     "productCode",
	//     "websiteUrl"
	//   ],
	//   "parameters": {
	//     "productCode": {
	//       "description": "Products to associate with the user.",
	//       "enum": [
	//         "AFC",
	//         "AFG",
	//         "AFMC",
	//         "AFS",
	//         "AFV"
	//       ],
	//       "enumDescriptions": [
	//         "AdSense For Content",
	//         "AdSense For Games",
	//         "AdSense For Mobile Content",
	//         "AdSense For Search",
	//         "AdSense For Video"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userLocale": {
	//       "description": "The preferred locale of the user.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "websiteLocale": {
	//       "description": "The locale of the user's hosted website.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "websiteUrl": {
	//       "description": "The URL of the user's hosted website.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "associationsessions/start",
	//   "response": {
	//     "$ref": "AssociationSession"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.associationsessions.verify":

type AssociationsessionsVerifyCall struct {
	s             *Service
	token         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Verify: Verify an association session after the association callback
// returns from AdSense signup.

func (r *AssociationsessionsService) Verify(token string) *AssociationsessionsVerifyCall {
	return &AssociationsessionsVerifyCall{
		s:             r.s,
		token:         token,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "associationsessions/verify",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssociationsessionsVerifyCall) Fields(s ...googleapi.Field) *AssociationsessionsVerifyCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AssociationsessionsVerifyCall) Do() (*AssociationSession, error) {
	var returnValue *AssociationSession
	c.params_.Set("token", fmt.Sprintf("%v", c.token))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Verify an association session after the association callback returns from AdSense signup.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.associationsessions.verify",
	//   "parameterOrder": [
	//     "token"
	//   ],
	//   "parameters": {
	//     "token": {
	//       "description": "The token returned to the association callback URL.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "associationsessions/verify",
	//   "response": {
	//     "$ref": "AssociationSession"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.customchannels.delete":

type CustomchannelsDeleteCall struct {
	s               *Service
	adClientId      string
	customChannelId string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Delete: Delete a specific custom channel from the host AdSense
// account.

func (r *CustomchannelsService) Delete(adClientId string, customChannelId string) *CustomchannelsDeleteCall {
	return &CustomchannelsDeleteCall{
		s:               r.s,
		adClientId:      adClientId,
		customChannelId: customChannelId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "adclients/{adClientId}/customchannels/{customChannelId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsDeleteCall) Fields(s ...googleapi.Field) *CustomchannelsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomchannelsDeleteCall) Do() (*CustomChannel, error) {
	var returnValue *CustomChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId":      c.adClientId,
		"customChannelId": c.customChannelId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete a specific custom channel from the host AdSense account.",
	//   "httpMethod": "DELETE",
	//   "id": "adsensehost.customchannels.delete",
	//   "parameterOrder": [
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client from which to delete the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels/{customChannelId}",
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.customchannels.get":

type CustomchannelsGetCall struct {
	s               *Service
	adClientId      string
	customChannelId string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Get: Get a specific custom channel from the host AdSense account.

func (r *CustomchannelsService) Get(adClientId string, customChannelId string) *CustomchannelsGetCall {
	return &CustomchannelsGetCall{
		s:               r.s,
		adClientId:      adClientId,
		customChannelId: customChannelId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "adclients/{adClientId}/customchannels/{customChannelId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsGetCall) Fields(s ...googleapi.Field) *CustomchannelsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomchannelsGetCall) Do() (*CustomChannel, error) {
	var returnValue *CustomChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId":      c.adClientId,
		"customChannelId": c.customChannelId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get a specific custom channel from the host AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.customchannels.get",
	//   "parameterOrder": [
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client from which to get the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels/{customChannelId}",
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.customchannels.insert":

type CustomchannelsInsertCall struct {
	s             *Service
	adClientId    string
	customchannel *CustomChannel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Add a new custom channel to the host AdSense account.

func (r *CustomchannelsService) Insert(adClientId string, customchannel *CustomChannel) *CustomchannelsInsertCall {
	return &CustomchannelsInsertCall{
		s:             r.s,
		adClientId:    adClientId,
		customchannel: customchannel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}/customchannels",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsInsertCall) Fields(s ...googleapi.Field) *CustomchannelsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomchannelsInsertCall) Do() (*CustomChannel, error) {
	var returnValue *CustomChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.customchannel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add a new custom channel to the host AdSense account.",
	//   "httpMethod": "POST",
	//   "id": "adsensehost.customchannels.insert",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client to which the new custom channel will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels",
	//   "request": {
	//     "$ref": "CustomChannel"
	//   },
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.customchannels.list":

type CustomchannelsListCall struct {
	s             *Service
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all host custom channels in this AdSense account.

func (r *CustomchannelsService) List(adClientId string) *CustomchannelsListCall {
	return &CustomchannelsListCall{
		s:             r.s,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}/customchannels",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of custom channels to include in the response, used for
// paging.
func (c *CustomchannelsListCall) MaxResults(maxResults int64) *CustomchannelsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through custom channels. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *CustomchannelsListCall) PageToken(pageToken string) *CustomchannelsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsListCall) Fields(s ...googleapi.Field) *CustomchannelsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomchannelsListCall) Do() (*CustomChannels, error) {
	var returnValue *CustomChannels
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all host custom channels in this AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.customchannels.list",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client for which to list custom channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of custom channels to include in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through custom channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels",
	//   "response": {
	//     "$ref": "CustomChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.customchannels.patch":

type CustomchannelsPatchCall struct {
	s               *Service
	adClientId      string
	customChannelId string
	customchannel   *CustomChannel
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Patch: Update a custom channel in the host AdSense account. This
// method supports patch semantics.

func (r *CustomchannelsService) Patch(adClientId string, customChannelId string, customchannel *CustomChannel) *CustomchannelsPatchCall {
	return &CustomchannelsPatchCall{
		s:               r.s,
		adClientId:      adClientId,
		customChannelId: customChannelId,
		customchannel:   customchannel,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "adclients/{adClientId}/customchannels",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsPatchCall) Fields(s ...googleapi.Field) *CustomchannelsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomchannelsPatchCall) Do() (*CustomChannel, error) {
	var returnValue *CustomChannel
	c.params_.Set("customChannelId", fmt.Sprintf("%v", c.customChannelId))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.customchannel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update a custom channel in the host AdSense account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "adsensehost.customchannels.patch",
	//   "parameterOrder": [
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client in which the custom channel will be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel to get.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels",
	//   "request": {
	//     "$ref": "CustomChannel"
	//   },
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.customchannels.update":

type CustomchannelsUpdateCall struct {
	s             *Service
	adClientId    string
	customchannel *CustomChannel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update a custom channel in the host AdSense account.

func (r *CustomchannelsService) Update(adClientId string, customchannel *CustomChannel) *CustomchannelsUpdateCall {
	return &CustomchannelsUpdateCall{
		s:             r.s,
		adClientId:    adClientId,
		customchannel: customchannel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}/customchannels",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomchannelsUpdateCall) Fields(s ...googleapi.Field) *CustomchannelsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CustomchannelsUpdateCall) Do() (*CustomChannel, error) {
	var returnValue *CustomChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.customchannel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update a custom channel in the host AdSense account.",
	//   "httpMethod": "PUT",
	//   "id": "adsensehost.customchannels.update",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client in which the custom channel will be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/customchannels",
	//   "request": {
	//     "$ref": "CustomChannel"
	//   },
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.reports.generate":

type ReportsGenerateCall struct {
	s             *Service
	startDate     string
	endDate       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generate: Generate an AdSense report based on the report request sent
// in the query parameters. Returns the result as JSON; to retrieve
// output in CSV format specify "alt=csv" as a query parameter.

func (r *ReportsService) Generate(startDate string, endDate string) *ReportsGenerateCall {
	return &ReportsGenerateCall{
		s:             r.s,
		startDate:     startDate,
		endDate:       endDate,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reports",
	}
}

// Dimension sets the optional parameter "dimension": Dimensions to base
// the report on.
func (c *ReportsGenerateCall) Dimension(dimension ...string) *ReportsGenerateCall {
	c.params_["dimension"] = dimension
	return c
}

// Filter sets the optional parameter "filter": Filters to be run on the
// report.
func (c *ReportsGenerateCall) Filter(filter ...string) *ReportsGenerateCall {
	c.params_["filter"] = filter
	return c
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *ReportsGenerateCall) Locale(locale string) *ReportsGenerateCall {
	c.params_.Set("locale", fmt.Sprintf("%v", locale))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *ReportsGenerateCall) MaxResults(maxResults int64) *ReportsGenerateCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// Metric sets the optional parameter "metric": Numeric columns to
// include in the report.
func (c *ReportsGenerateCall) Metric(metric ...string) *ReportsGenerateCall {
	c.params_["metric"] = metric
	return c
}

// Sort sets the optional parameter "sort": The name of a dimension or
// metric to sort the resulting report on, optionally prefixed with "+"
// to sort ascending or "-" to sort descending. If no prefix is
// specified, the column is sorted ascending.
func (c *ReportsGenerateCall) Sort(sort ...string) *ReportsGenerateCall {
	c.params_["sort"] = sort
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *ReportsGenerateCall) StartIndex(startIndex int64) *ReportsGenerateCall {
	c.params_.Set("startIndex", fmt.Sprintf("%v", startIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsGenerateCall) Fields(s ...googleapi.Field) *ReportsGenerateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsGenerateCall) Do() (*Report, error) {
	var returnValue *Report
	c.params_.Set("endDate", fmt.Sprintf("%v", c.endDate))
	c.params_.Set("startDate", fmt.Sprintf("%v", c.startDate))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generate an AdSense report based on the report request sent in the query parameters. Returns the result as JSON; to retrieve output in CSV format specify \"alt=csv\" as a query parameter.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.reports.generate",
	//   "parameterOrder": [
	//     "startDate",
	//     "endDate"
	//   ],
	//   "parameters": {
	//     "dimension": {
	//       "description": "Dimensions to base the report on.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "End of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "Filters to be run on the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+(==|=@).+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "Optional locale to use for translating report output to a local language. Defaults to \"en_US\" if not specified.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows of report data to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "metric": {
	//       "description": "Numeric columns to include in the report.",
	//       "location": "query",
	//       "pattern": "[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "The name of a dimension or metric to sort the resulting report on, optionally prefixed with \"+\" to sort ascending or \"-\" to sort descending. If no prefix is specified, the column is sorted ascending.",
	//       "location": "query",
	//       "pattern": "(\\+|-)?[a-zA-Z_]+",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "Start of the date range to report on in \"YYYY-MM-DD\" format, inclusive.",
	//       "location": "query",
	//       "pattern": "\\d{4}-\\d{2}-\\d{2}|(today|startOfMonth|startOfYear)(([\\-\\+]\\d+[dwmy]){0,3}?)",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "reports",
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.urlchannels.delete":

type UrlchannelsDeleteCall struct {
	s             *Service
	adClientId    string
	urlChannelId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete a URL channel from the host AdSense account.

func (r *UrlchannelsService) Delete(adClientId string, urlChannelId string) *UrlchannelsDeleteCall {
	return &UrlchannelsDeleteCall{
		s:             r.s,
		adClientId:    adClientId,
		urlChannelId:  urlChannelId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}/urlchannels/{urlChannelId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlchannelsDeleteCall) Fields(s ...googleapi.Field) *UrlchannelsDeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlchannelsDeleteCall) Do() (*UrlChannel, error) {
	var returnValue *UrlChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId":   c.adClientId,
		"urlChannelId": c.urlChannelId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete a URL channel from the host AdSense account.",
	//   "httpMethod": "DELETE",
	//   "id": "adsensehost.urlchannels.delete",
	//   "parameterOrder": [
	//     "adClientId",
	//     "urlChannelId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client from which to delete the URL channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "urlChannelId": {
	//       "description": "URL channel to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/urlchannels/{urlChannelId}",
	//   "response": {
	//     "$ref": "UrlChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.urlchannels.insert":

type UrlchannelsInsertCall struct {
	s             *Service
	adClientId    string
	urlchannel    *UrlChannel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Add a new URL channel to the host AdSense account.

func (r *UrlchannelsService) Insert(adClientId string, urlchannel *UrlChannel) *UrlchannelsInsertCall {
	return &UrlchannelsInsertCall{
		s:             r.s,
		adClientId:    adClientId,
		urlchannel:    urlchannel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}/urlchannels",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlchannelsInsertCall) Fields(s ...googleapi.Field) *UrlchannelsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlchannelsInsertCall) Do() (*UrlChannel, error) {
	var returnValue *UrlChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.urlchannel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add a new URL channel to the host AdSense account.",
	//   "httpMethod": "POST",
	//   "id": "adsensehost.urlchannels.insert",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client to which the new URL channel will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/urlchannels",
	//   "request": {
	//     "$ref": "UrlChannel"
	//   },
	//   "response": {
	//     "$ref": "UrlChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}

// method id "adsensehost.urlchannels.list":

type UrlchannelsListCall struct {
	s             *Service
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all host URL channels in the host AdSense account.

func (r *UrlchannelsService) List(adClientId string) *UrlchannelsListCall {
	return &UrlchannelsListCall{
		s:             r.s,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "adclients/{adClientId}/urlchannels",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of URL channels to include in the response, used for paging.
func (c *UrlchannelsListCall) MaxResults(maxResults int64) *UrlchannelsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through URL channels. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *UrlchannelsListCall) PageToken(pageToken string) *UrlchannelsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UrlchannelsListCall) Fields(s ...googleapi.Field) *UrlchannelsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UrlchannelsListCall) Do() (*UrlChannels, error) {
	var returnValue *UrlChannels
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"adClientId": c.adClientId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all host URL channels in the host AdSense account.",
	//   "httpMethod": "GET",
	//   "id": "adsensehost.urlchannels.list",
	//   "parameterOrder": [
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "adClientId": {
	//       "description": "Ad client for which to list URL channels.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of URL channels to include in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through URL channels. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "adclients/{adClientId}/urlchannels",
	//   "response": {
	//     "$ref": "UrlChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adsensehost"
	//   ]
	// }

}