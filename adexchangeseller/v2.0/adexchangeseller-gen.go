// Package adexchangeseller provides access to the Ad Exchange Seller API.
//
// See https://developers.google.com/ad-exchange/seller-rest/
//
// Usage example:
//
//   import "github.com/jfcote87/api/adexchangeseller/v2.0"
//   ...
//   adexchangesellerService, err := adexchangeseller.New(oauthHttpClient)
package adexchangeseller

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

const apiId = "adexchangeseller:v2.0"
const apiName = "adexchangeseller"
const apiVersion = "v2.0"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/adexchangeseller/v2.0/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your Ad Exchange data
	AdexchangeSellerScope = "https://www.googleapis.com/auth/adexchange.seller"

	// View your Ad Exchange data
	AdexchangeSellerReadonlyScope = "https://www.googleapis.com/auth/adexchange.seller.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Accounts = NewAccountsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Accounts *AccountsService
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	rs.Adclients = NewAccountsAdclientsService(s)
	rs.Alerts = NewAccountsAlertsService(s)
	rs.Customchannels = NewAccountsCustomchannelsService(s)
	rs.Metadata = NewAccountsMetadataService(s)
	rs.Preferreddeals = NewAccountsPreferreddealsService(s)
	rs.Reports = NewAccountsReportsService(s)
	rs.Urlchannels = NewAccountsUrlchannelsService(s)
	return rs
}

type AccountsService struct {
	s *Service

	Adclients *AccountsAdclientsService

	Alerts *AccountsAlertsService

	Customchannels *AccountsCustomchannelsService

	Metadata *AccountsMetadataService

	Preferreddeals *AccountsPreferreddealsService

	Reports *AccountsReportsService

	Urlchannels *AccountsUrlchannelsService
}

func NewAccountsAdclientsService(s *Service) *AccountsAdclientsService {
	rs := &AccountsAdclientsService{s: s}
	return rs
}

type AccountsAdclientsService struct {
	s *Service
}

func NewAccountsAlertsService(s *Service) *AccountsAlertsService {
	rs := &AccountsAlertsService{s: s}
	return rs
}

type AccountsAlertsService struct {
	s *Service
}

func NewAccountsCustomchannelsService(s *Service) *AccountsCustomchannelsService {
	rs := &AccountsCustomchannelsService{s: s}
	return rs
}

type AccountsCustomchannelsService struct {
	s *Service
}

func NewAccountsMetadataService(s *Service) *AccountsMetadataService {
	rs := &AccountsMetadataService{s: s}
	rs.Dimensions = NewAccountsMetadataDimensionsService(s)
	rs.Metrics = NewAccountsMetadataMetricsService(s)
	return rs
}

type AccountsMetadataService struct {
	s *Service

	Dimensions *AccountsMetadataDimensionsService

	Metrics *AccountsMetadataMetricsService
}

func NewAccountsMetadataDimensionsService(s *Service) *AccountsMetadataDimensionsService {
	rs := &AccountsMetadataDimensionsService{s: s}
	return rs
}

type AccountsMetadataDimensionsService struct {
	s *Service
}

func NewAccountsMetadataMetricsService(s *Service) *AccountsMetadataMetricsService {
	rs := &AccountsMetadataMetricsService{s: s}
	return rs
}

type AccountsMetadataMetricsService struct {
	s *Service
}

func NewAccountsPreferreddealsService(s *Service) *AccountsPreferreddealsService {
	rs := &AccountsPreferreddealsService{s: s}
	return rs
}

type AccountsPreferreddealsService struct {
	s *Service
}

func NewAccountsReportsService(s *Service) *AccountsReportsService {
	rs := &AccountsReportsService{s: s}
	rs.Saved = NewAccountsReportsSavedService(s)
	return rs
}

type AccountsReportsService struct {
	s *Service

	Saved *AccountsReportsSavedService
}

func NewAccountsReportsSavedService(s *Service) *AccountsReportsSavedService {
	rs := &AccountsReportsSavedService{s: s}
	return rs
}

type AccountsReportsSavedService struct {
	s *Service
}

func NewAccountsUrlchannelsService(s *Service) *AccountsUrlchannelsService {
	rs := &AccountsUrlchannelsService{s: s}
	return rs
}

type AccountsUrlchannelsService struct {
	s *Service
}

type Account struct {
	// Id: Unique identifier of this account.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adexchangeseller#account.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this account.
	Name string `json:"name,omitempty"`
}

type Accounts struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The accounts returned in this list response.
	Items []*Account `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adexchangeseller#accounts.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through accounts. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AdClient struct {
	// ArcOptIn: Whether this ad client is opted in to ARC.
	ArcOptIn bool `json:"arcOptIn,omitempty"`

	// Id: Unique identifier of this ad client.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adexchangeseller#adClient.
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

	// Kind: Kind of list this is, in this case adexchangeseller#adClients.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through ad clients. To
	// retrieve the next page of results, set the next request's "pageToken"
	// value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Alert struct {
	// Id: Unique identifier of this alert. This should be considered an
	// opaque identifier; it is not safe to rely on it being in any
	// particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case adexchangeseller#alert.
	Kind string `json:"kind,omitempty"`

	// Message: The localized alert message.
	Message string `json:"message,omitempty"`

	// Severity: Severity of this alert. Possible values: INFO, WARNING,
	// SEVERE.
	Severity string `json:"severity,omitempty"`

	// Type: Type of this alert. Possible values: SELF_HOLD,
	// MIGRATED_TO_BILLING3, ADDRESS_PIN_VERIFICATION,
	// PHONE_PIN_VERIFICATION, CORPORATE_ENTITY, GRAYLISTED_PUBLISHER,
	// API_HOLD.
	Type string `json:"type,omitempty"`
}

type Alerts struct {
	// Items: The alerts returned in this list response.
	Items []*Alert `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adexchangeseller#alerts.
	Kind string `json:"kind,omitempty"`
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
	// adexchangeseller#customChannel.
	Kind string `json:"kind,omitempty"`

	// Name: Name of this custom channel.
	Name string `json:"name,omitempty"`

	// TargetingInfo: The targeting information of this custom channel, if
	// activated.
	TargetingInfo *CustomChannelTargetingInfo `json:"targetingInfo,omitempty"`
}

type CustomChannelTargetingInfo struct {
	// AdsAppearOn: The name used to describe this channel externally.
	AdsAppearOn string `json:"adsAppearOn,omitempty"`

	// Description: The external description of the channel.
	Description string `json:"description,omitempty"`

	// Location: The locations in which ads appear. (Only valid for content
	// and mobile content ads). Acceptable values for content ads are:
	// TOP_LEFT, TOP_CENTER, TOP_RIGHT, MIDDLE_LEFT, MIDDLE_CENTER,
	// MIDDLE_RIGHT, BOTTOM_LEFT, BOTTOM_CENTER, BOTTOM_RIGHT,
	// MULTIPLE_LOCATIONS. Acceptable values for mobile content ads are:
	// TOP, MIDDLE, BOTTOM, MULTIPLE_LOCATIONS.
	Location string `json:"location,omitempty"`

	// SiteLanguage: The language of the sites ads will be displayed on.
	SiteLanguage string `json:"siteLanguage,omitempty"`
}

type CustomChannels struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The custom channels returned in this list response.
	Items []*CustomChannel `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case
	// adexchangeseller#customChannels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through custom
	// channels. To retrieve the next page of results, set the next
	// request's "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Metadata struct {
	Items []*ReportingMetadataEntry `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case adexchangeseller#metadata.
	Kind string `json:"kind,omitempty"`
}

type PreferredDeal struct {
	// AdvertiserName: The name of the advertiser this deal is for.
	AdvertiserName string `json:"advertiserName,omitempty"`

	// BuyerNetworkName: The name of the buyer network this deal is for.
	BuyerNetworkName string `json:"buyerNetworkName,omitempty"`

	// CurrencyCode: The currency code that applies to the fixed_cpm value.
	// If not set then assumed to be USD.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// EndTime: Time when this deal stops being active in seconds since the
	// epoch (GMT). If not set then this deal is valid until manually
	// disabled by the publisher.
	EndTime uint64 `json:"endTime,omitempty,string"`

	// FixedCpm: The fixed price for this preferred deal. In cpm micros of
	// currency according to currencyCode. If set, then this preferred deal
	// is eligible for the fixed price tier of buying (highest priority, pay
	// exactly the configured fixed price).
	FixedCpm int64 `json:"fixedCpm,omitempty,string"`

	// Id: Unique identifier of this preferred deal.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Kind of resource this is, in this case
	// adexchangeseller#preferredDeal.
	Kind string `json:"kind,omitempty"`

	// StartTime: Time when this deal becomes active in seconds since the
	// epoch (GMT). If not set then this deal is active immediately upon
	// creation.
	StartTime uint64 `json:"startTime,omitempty,string"`
}

type PreferredDeals struct {
	// Items: The preferred deals returned in this list response.
	Items []*PreferredDeal `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case
	// adexchangeseller#preferredDeals.
	Kind string `json:"kind,omitempty"`
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

	// Kind: Kind this is, in this case adexchangeseller#report.
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

type ReportingMetadataEntry struct {
	// CompatibleDimensions: For metrics this is a list of dimension IDs
	// which the metric is compatible with, for dimensions it is a list of
	// compatibility groups the dimension belongs to.
	CompatibleDimensions []string `json:"compatibleDimensions,omitempty"`

	// CompatibleMetrics: The names of the metrics the dimension or metric
	// this reporting metadata entry describes is compatible with.
	CompatibleMetrics []string `json:"compatibleMetrics,omitempty"`

	// Id: Unique identifier of this reporting metadata entry, corresponding
	// to the name of the appropriate dimension or metric.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adexchangeseller#reportingMetadataEntry.
	Kind string `json:"kind,omitempty"`

	// RequiredDimensions: The names of the dimensions which the dimension
	// or metric this reporting metadata entry describes requires to also be
	// present in order for the report to be valid. Omitting these will not
	// cause an error or warning, but may result in data which cannot be
	// correctly interpreted.
	RequiredDimensions []string `json:"requiredDimensions,omitempty"`

	// RequiredMetrics: The names of the metrics which the dimension or
	// metric this reporting metadata entry describes requires to also be
	// present in order for the report to be valid. Omitting these will not
	// cause an error or warning, but may result in data which cannot be
	// correctly interpreted.
	RequiredMetrics []string `json:"requiredMetrics,omitempty"`

	// SupportedProducts: The codes of the projects supported by the
	// dimension or metric this reporting metadata entry describes.
	SupportedProducts []string `json:"supportedProducts,omitempty"`
}

type SavedReport struct {
	// Id: Unique identifier of this saved report.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adexchangeseller#savedReport.
	Kind string `json:"kind,omitempty"`

	// Name: This saved report's name.
	Name string `json:"name,omitempty"`
}

type SavedReports struct {
	// Etag: ETag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The saved reports returned in this list response.
	Items []*SavedReport `json:"items,omitempty"`

	// Kind: Kind of list this is, in this case
	// adexchangeseller#savedReports.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through saved reports.
	// To retrieve the next page of results, set the next request's
	// "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type UrlChannel struct {
	// Id: Unique identifier of this URL channel. This should be considered
	// an opaque identifier; it is not safe to rely on it being in any
	// particular format.
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is, in this case
	// adexchangeseller#urlChannel.
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

	// Kind: Kind of list this is, in this case
	// adexchangeseller#urlChannels.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through URL channels.
	// To retrieve the next page of results, set the next request's
	// "pageToken" value to this.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "adexchangeseller.accounts.get":

type AccountsGetCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about the selected Ad Exchange account.

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
	//   "description": "Get information about the selected Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.get",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to get information about. Tip: 'myaccount' is a valid ID.",
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
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.list":

type AccountsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all accounts available to this Ad Exchange account.

func (r *AccountsService) List() *AccountsListCall {
	return &AccountsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of accounts to include in the response, used for paging.
func (c *AccountsListCall) MaxResults(maxResults int64) *AccountsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through accounts. To retrieve the next page, set
// this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsListCall) PageToken(pageToken string) *AccountsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
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
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all accounts available to this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of accounts to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through accounts. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "Accounts"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.adclients.list":

type AccountsAdclientsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all ad clients in this Ad Exchange account.

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
	//   "description": "List all ad clients in this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.adclients.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
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
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.alerts.list":

type AccountsAlertsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List the alerts for this Ad Exchange account.

func (r *AccountsAlertsService) List(accountId string) *AccountsAlertsListCall {
	return &AccountsAlertsListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/alerts",
	}
}

// Locale sets the optional parameter "locale": The locale to use for
// translating alert messages. The account locale will be used if this
// is not supplied. The AdSense default (English) will be used if the
// supplied locale is invalid or unsupported.
func (c *AccountsAlertsListCall) Locale(locale string) *AccountsAlertsListCall {
	c.params_.Set("locale", fmt.Sprintf("%v", locale))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsAlertsListCall) Fields(s ...googleapi.Field) *AccountsAlertsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsAlertsListCall) Do() (*Alerts, error) {
	var returnValue *Alerts
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
	//   "description": "List the alerts for this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.alerts.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account owning the alerts.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "locale": {
	//       "description": "The locale to use for translating alert messages. The account locale will be used if this is not supplied. The AdSense default (English) will be used if the supplied locale is invalid or unsupported.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/alerts",
	//   "response": {
	//     "$ref": "Alerts"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.customchannels.get":

type AccountsCustomchannelsGetCall struct {
	s               *Service
	accountId       string
	adClientId      string
	customChannelId string
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Get: Get the specified custom channel from the specified ad client.

func (r *AccountsCustomchannelsService) Get(accountId string, adClientId string, customChannelId string) *AccountsCustomchannelsGetCall {
	return &AccountsCustomchannelsGetCall{
		s:               r.s,
		accountId:       accountId,
		adClientId:      adClientId,
		customChannelId: customChannelId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "accounts/{accountId}/adclients/{adClientId}/customchannels/{customChannelId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustomchannelsGetCall) Fields(s ...googleapi.Field) *AccountsCustomchannelsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsCustomchannelsGetCall) Do() (*CustomChannel, error) {
	var returnValue *CustomChannel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":       c.accountId,
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
	//   "description": "Get the specified custom channel from the specified ad client.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.customchannels.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId",
	//     "customChannelId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "adClientId": {
	//       "description": "Ad client which contains the custom channel.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "customChannelId": {
	//       "description": "Custom channel to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/adclients/{adClientId}/customchannels/{customChannelId}",
	//   "response": {
	//     "$ref": "CustomChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.customchannels.list":

type AccountsCustomchannelsListCall struct {
	s             *Service
	accountId     string
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all custom channels in the specified ad client for this Ad
// Exchange account.

func (r *AccountsCustomchannelsService) List(accountId string, adClientId string) *AccountsCustomchannelsListCall {
	return &AccountsCustomchannelsListCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/customchannels",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of custom channels to include in the response, used for
// paging.
func (c *AccountsCustomchannelsListCall) MaxResults(maxResults int64) *AccountsCustomchannelsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through custom channels. To retrieve the next
// page, set this parameter to the value of "nextPageToken" from the
// previous response.
func (c *AccountsCustomchannelsListCall) PageToken(pageToken string) *AccountsCustomchannelsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustomchannelsListCall) Fields(s ...googleapi.Field) *AccountsCustomchannelsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsCustomchannelsListCall) Do() (*CustomChannels, error) {
	var returnValue *CustomChannels
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
	//   "description": "List all custom channels in the specified ad client for this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.customchannels.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
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
	//   "path": "accounts/{accountId}/adclients/{adClientId}/customchannels",
	//   "response": {
	//     "$ref": "CustomChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.metadata.dimensions.list":

type AccountsMetadataDimensionsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List the metadata for the dimensions available to this
// AdExchange account.

func (r *AccountsMetadataDimensionsService) List(accountId string) *AccountsMetadataDimensionsListCall {
	return &AccountsMetadataDimensionsListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/metadata/dimensions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsMetadataDimensionsListCall) Fields(s ...googleapi.Field) *AccountsMetadataDimensionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsMetadataDimensionsListCall) Do() (*Metadata, error) {
	var returnValue *Metadata
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
	//   "description": "List the metadata for the dimensions available to this AdExchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.metadata.dimensions.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account with visibility to the dimensions.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/metadata/dimensions",
	//   "response": {
	//     "$ref": "Metadata"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.metadata.metrics.list":

type AccountsMetadataMetricsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List the metadata for the metrics available to this AdExchange
// account.

func (r *AccountsMetadataMetricsService) List(accountId string) *AccountsMetadataMetricsListCall {
	return &AccountsMetadataMetricsListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/metadata/metrics",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsMetadataMetricsListCall) Fields(s ...googleapi.Field) *AccountsMetadataMetricsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsMetadataMetricsListCall) Do() (*Metadata, error) {
	var returnValue *Metadata
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
	//   "description": "List the metadata for the metrics available to this AdExchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.metadata.metrics.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account with visibility to the metrics.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/metadata/metrics",
	//   "response": {
	//     "$ref": "Metadata"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.preferreddeals.get":

type AccountsPreferreddealsGetCall struct {
	s             *Service
	accountId     string
	dealId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about the selected Ad Exchange Preferred Deal.

func (r *AccountsPreferreddealsService) Get(accountId string, dealId string) *AccountsPreferreddealsGetCall {
	return &AccountsPreferreddealsGetCall{
		s:             r.s,
		accountId:     accountId,
		dealId:        dealId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/preferreddeals/{dealId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPreferreddealsGetCall) Fields(s ...googleapi.Field) *AccountsPreferreddealsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPreferreddealsGetCall) Do() (*PreferredDeal, error) {
	var returnValue *PreferredDeal
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
		"dealId":    c.dealId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get information about the selected Ad Exchange Preferred Deal.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.preferreddeals.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "dealId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account owning the deal.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "dealId": {
	//       "description": "Preferred deal to get information about.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/preferreddeals/{dealId}",
	//   "response": {
	//     "$ref": "PreferredDeal"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.preferreddeals.list":

type AccountsPreferreddealsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List the preferred deals for this Ad Exchange account.

func (r *AccountsPreferreddealsService) List(accountId string) *AccountsPreferreddealsListCall {
	return &AccountsPreferreddealsListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/preferreddeals",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPreferreddealsListCall) Fields(s ...googleapi.Field) *AccountsPreferreddealsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPreferreddealsListCall) Do() (*PreferredDeals, error) {
	var returnValue *PreferredDeals
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
	//   "description": "List the preferred deals for this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.preferreddeals.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account owning the deals.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/preferreddeals",
	//   "response": {
	//     "$ref": "PreferredDeals"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.reports.generate":

type AccountsReportsGenerateCall struct {
	s             *Service
	accountId     string
	startDate     string
	endDate       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generate: Generate an Ad Exchange report based on the report request
// sent in the query parameters. Returns the result as JSON; to retrieve
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

func (c *AccountsReportsGenerateCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	c.params_.Set("endDate", fmt.Sprintf("%v", c.endDate))
	c.params_.Set("startDate", fmt.Sprintf("%v", c.startDate))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
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
	//   "description": "Generate an Ad Exchange report based on the report request sent in the query parameters. Returns the result as JSON; to retrieve output in CSV format specify \"alt=csv\" as a query parameter.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.reports.generate",
	//   "parameterOrder": [
	//     "accountId",
	//     "startDate",
	//     "endDate"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account which owns the generated report.",
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
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "adexchangeseller.accounts.reports.saved.generate":

type AccountsReportsSavedGenerateCall struct {
	s             *Service
	accountId     string
	savedReportId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generate: Generate an Ad Exchange report based on the saved report ID
// sent in the query parameters.

func (r *AccountsReportsSavedService) Generate(accountId string, savedReportId string) *AccountsReportsSavedGenerateCall {
	return &AccountsReportsSavedGenerateCall{
		s:             r.s,
		accountId:     accountId,
		savedReportId: savedReportId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/reports/{savedReportId}",
	}
}

// Locale sets the optional parameter "locale": Optional locale to use
// for translating report output to a local language. Defaults to
// "en_US" if not specified.
func (c *AccountsReportsSavedGenerateCall) Locale(locale string) *AccountsReportsSavedGenerateCall {
	c.params_.Set("locale", fmt.Sprintf("%v", locale))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows of report data to return.
func (c *AccountsReportsSavedGenerateCall) MaxResults(maxResults int64) *AccountsReportsSavedGenerateCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first row of report data to return.
func (c *AccountsReportsSavedGenerateCall) StartIndex(startIndex int64) *AccountsReportsSavedGenerateCall {
	c.params_.Set("startIndex", fmt.Sprintf("%v", startIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsReportsSavedGenerateCall) Fields(s ...googleapi.Field) *AccountsReportsSavedGenerateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsReportsSavedGenerateCall) Do() (*Report, error) {
	var returnValue *Report
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":     c.accountId,
		"savedReportId": c.savedReportId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generate an Ad Exchange report based on the saved report ID sent in the query parameters.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.reports.saved.generate",
	//   "parameterOrder": [
	//     "accountId",
	//     "savedReportId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account owning the saved report.",
	//       "location": "path",
	//       "required": true,
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
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "50000",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "savedReportId": {
	//       "description": "The saved report to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first row of report data to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "5000",
	//       "minimum": "0",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "accounts/{accountId}/reports/{savedReportId}",
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.reports.saved.list":

type AccountsReportsSavedListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all saved reports in this Ad Exchange account.

func (r *AccountsReportsSavedService) List(accountId string) *AccountsReportsSavedListCall {
	return &AccountsReportsSavedListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/reports/saved",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of saved reports to include in the response, used for paging.
func (c *AccountsReportsSavedListCall) MaxResults(maxResults int64) *AccountsReportsSavedListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through saved reports. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsReportsSavedListCall) PageToken(pageToken string) *AccountsReportsSavedListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsReportsSavedListCall) Fields(s ...googleapi.Field) *AccountsReportsSavedListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsReportsSavedListCall) Do() (*SavedReports, error) {
	var returnValue *SavedReports
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
	//   "description": "List all saved reports in this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.reports.saved.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account owning the saved reports.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of saved reports to include in the response, used for paging.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token, used to page through saved reports. To retrieve the next page, set this parameter to the value of \"nextPageToken\" from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/reports/saved",
	//   "response": {
	//     "$ref": "SavedReports"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}

// method id "adexchangeseller.accounts.urlchannels.list":

type AccountsUrlchannelsListCall struct {
	s             *Service
	accountId     string
	adClientId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all URL channels in the specified ad client for this Ad
// Exchange account.

func (r *AccountsUrlchannelsService) List(accountId string, adClientId string) *AccountsUrlchannelsListCall {
	return &AccountsUrlchannelsListCall{
		s:             r.s,
		accountId:     accountId,
		adClientId:    adClientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/adclients/{adClientId}/urlchannels",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of URL channels to include in the response, used for paging.
func (c *AccountsUrlchannelsListCall) MaxResults(maxResults int64) *AccountsUrlchannelsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token, used to page through URL channels. To retrieve the next page,
// set this parameter to the value of "nextPageToken" from the previous
// response.
func (c *AccountsUrlchannelsListCall) PageToken(pageToken string) *AccountsUrlchannelsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsUrlchannelsListCall) Fields(s ...googleapi.Field) *AccountsUrlchannelsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsUrlchannelsListCall) Do() (*UrlChannels, error) {
	var returnValue *UrlChannels
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
	//   "description": "List all URL channels in the specified ad client for this Ad Exchange account.",
	//   "httpMethod": "GET",
	//   "id": "adexchangeseller.accounts.urlchannels.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "adClientId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "Account to which the ad client belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
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
	//   "path": "accounts/{accountId}/adclients/{adClientId}/urlchannels",
	//   "response": {
	//     "$ref": "UrlChannels"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/adexchange.seller",
	//     "https://www.googleapis.com/auth/adexchange.seller.readonly"
	//   ]
	// }

}
