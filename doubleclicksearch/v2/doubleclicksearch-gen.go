// Package doubleclicksearch provides access to the DoubleClick Search API.
//
// See https://developers.google.com/doubleclick-search/
//
// Usage example:
//
//   import "github.com/jfcote87/api/doubleclicksearch/v2"
//   ...
//   doubleclicksearchService, err := doubleclicksearch.New(oauthHttpClient)
package doubleclicksearch

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

const apiId = "doubleclicksearch:v2"
const apiName = "doubleclicksearch"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/doubleclicksearch/v2/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your advertising data in DoubleClick Search
	DoubleclicksearchScope = "https://www.googleapis.com/auth/doubleclicksearch"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Conversion = NewConversionService(s)
	s.Reports = NewReportsService(s)
	s.SavedColumns = NewSavedColumnsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Conversion *ConversionService

	Reports *ReportsService

	SavedColumns *SavedColumnsService
}

func NewConversionService(s *Service) *ConversionService {
	rs := &ConversionService{s: s}
	return rs
}

type ConversionService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	return rs
}

type ReportsService struct {
	s *Service
}

func NewSavedColumnsService(s *Service) *SavedColumnsService {
	rs := &SavedColumnsService{s: s}
	return rs
}

type SavedColumnsService struct {
	s *Service
}

type Availability struct {
	// AdvertiserId: DS advertiser ID.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AgencyId: DS agency ID.
	AgencyId int64 `json:"agencyId,omitempty,string"`

	// AvailabilityTimestamp: The time by which all conversions have been
	// uploaded, in epoch millis UTC.
	AvailabilityTimestamp uint64 `json:"availabilityTimestamp,omitempty,string"`

	// SegmentationId: The numeric segmentation identifier (for example,
	// DoubleClick Search Floodlight activity ID).
	SegmentationId int64 `json:"segmentationId,omitempty,string"`

	// SegmentationName: The friendly segmentation identifier (for example,
	// DoubleClick Search Floodlight activity name).
	SegmentationName string `json:"segmentationName,omitempty"`

	// SegmentationType: The segmentation type that this availability is for
	// (its default value is FLOODLIGHT).
	SegmentationType string `json:"segmentationType,omitempty"`
}

type Conversion struct {
	// AdGroupId: DS ad group ID.
	AdGroupId int64 `json:"adGroupId,omitempty,string"`

	// AdId: DS ad ID.
	AdId int64 `json:"adId,omitempty,string"`

	// AdvertiserId: DS advertiser ID.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AgencyId: DS agency ID.
	AgencyId int64 `json:"agencyId,omitempty,string"`

	// AttributionModel: Attribution model name. This field is ignored.
	AttributionModel string `json:"attributionModel,omitempty"`

	// CampaignId: DS campaign ID.
	CampaignId int64 `json:"campaignId,omitempty,string"`

	// ClickId: DS click ID for the conversion.
	ClickId string `json:"clickId,omitempty"`

	// ConversionId: Advertiser-provided ID for the conversion, also known
	// as the order ID.
	ConversionId string `json:"conversionId,omitempty"`

	// ConversionModifiedTimestamp: The time at which the conversion was
	// last modified, in epoch millis UTC.
	ConversionModifiedTimestamp uint64 `json:"conversionModifiedTimestamp,omitempty,string"`

	// ConversionTimestamp: The time at which the conversion took place, in
	// epoch millis UTC.
	ConversionTimestamp uint64 `json:"conversionTimestamp,omitempty,string"`

	// CountMillis: The number of conversions, formatted in millis
	// (conversions multiplied by 1000). This field is ignored.
	CountMillis int64 `json:"countMillis,omitempty,string"`

	// CriterionId: DS criterion (keyword) ID.
	CriterionId int64 `json:"criterionId,omitempty,string"`

	// CurrencyCode: The currency code for the conversion's revenue. Should
	// be in ISO 4217 alphabetic (3-char) format.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// CustomDimension: Custom dimensions for the conversion, which can be
	// used to filter data in a report.
	CustomDimension []*CustomDimension `json:"customDimension,omitempty"`

	// CustomMetric: Custom metrics for the conversion.
	CustomMetric []*CustomMetric `json:"customMetric,omitempty"`

	// DsConversionId: DS conversion ID.
	DsConversionId int64 `json:"dsConversionId,omitempty,string"`

	// EngineAccountId: DS engine account ID.
	EngineAccountId int64 `json:"engineAccountId,omitempty,string"`

	// FloodlightOrderId: The advertiser-provided order id for the
	// conversion.
	FloodlightOrderId string `json:"floodlightOrderId,omitempty"`

	// QuantityMillis: The quantity of this conversion, in millis.
	QuantityMillis int64 `json:"quantityMillis,omitempty,string"`

	// RevenueMicros: The revenue amount of this TRANSACTION conversion, in
	// micros.
	RevenueMicros int64 `json:"revenueMicros,omitempty,string"`

	// SegmentationId: The numeric segmentation identifier (for example,
	// DoubleClick Search Floodlight activity ID).
	SegmentationId int64 `json:"segmentationId,omitempty,string"`

	// SegmentationName: The friendly segmentation identifier (for example,
	// DoubleClick Search Floodlight activity name).
	SegmentationName string `json:"segmentationName,omitempty"`

	// SegmentationType: The segmentation type of this conversion (for
	// example, FLOODLIGHT).
	SegmentationType string `json:"segmentationType,omitempty"`

	// State: The state of the conversion, that is, either ACTIVE or
	// REMOVED. Note: state DELETED is deprecated.
	State string `json:"state,omitempty"`

	// Type: The type of the conversion, that is, either ACTION or
	// TRANSACTION. An ACTION conversion is an action by the user that has
	// no monetarily quantifiable value, while a TRANSACTION conversion is
	// an action that does have a monetarily quantifiable value. Examples
	// are email list signups (ACTION) versus ecommerce purchases
	// (TRANSACTION).
	Type string `json:"type,omitempty"`
}

type ConversionList struct {
	// Conversion: The conversions being requested.
	Conversion []*Conversion `json:"conversion,omitempty"`

	// Kind: Identifies this as a ConversionList resource. Value: the fixed
	// string doubleclicksearch#conversionList.
	Kind string `json:"kind,omitempty"`
}

type CustomDimension struct {
	// Name: Custom dimension name.
	Name string `json:"name,omitempty"`

	// Value: Custom dimension value.
	Value string `json:"value,omitempty"`
}

type CustomMetric struct {
	// Name: Custom metric name.
	Name string `json:"name,omitempty"`

	// Value: Custom metric numeric value.
	Value float64 `json:"value,omitempty"`
}

type Report struct {
	// Files: Asynchronous report only. Contains a list of generated report
	// files once the report has succesfully completed.
	Files []*ReportFiles `json:"files,omitempty"`

	// Id: Asynchronous report only. Id of the report.
	Id string `json:"id,omitempty"`

	// IsReportReady: Asynchronous report only. True if and only if the
	// report has completed successfully and the report files are ready to
	// be downloaded.
	IsReportReady bool `json:"isReportReady,omitempty"`

	// Kind: Identifies this as a Report resource. Value: the fixed string
	// doubleclicksearch#report.
	Kind string `json:"kind,omitempty"`

	// Request: The request that created the report. Optional fields not
	// specified in the original request are filled with default values.
	Request *ReportRequest `json:"request,omitempty"`

	// RowCount: The number of report rows generated by the report, not
	// including headers.
	RowCount int64 `json:"rowCount,omitempty"`

	// Rows: Synchronous report only. Generated report rows.
	Rows []*ReportRow `json:"rows,omitempty"`

	// StatisticsCurrencyCode: The currency code of all monetary values
	// produced in the report, including values that are set by users (e.g.,
	// keyword bid settings) and metrics (e.g., cost and revenue). The
	// currency code of a report is determined by the statisticsCurrency
	// field of the report request.
	StatisticsCurrencyCode string `json:"statisticsCurrencyCode,omitempty"`

	// StatisticsTimeZone: If all statistics of the report are sourced from
	// the same time zone, this would be it. Otherwise the field is unset.
	StatisticsTimeZone string `json:"statisticsTimeZone,omitempty"`
}

type ReportFiles struct {
	// ByteCount: The size of this report file in bytes.
	ByteCount int64 `json:"byteCount,omitempty,string"`

	// Url: Use this url to download the report file.
	Url string `json:"url,omitempty"`
}

type ReportApiColumnSpec struct {
	// ColumnName: Name of a DoubleClick Search column to include in the
	// report.
	ColumnName string `json:"columnName,omitempty"`

	// CustomDimensionName: Segments a report by a custom dimension. The
	// report must be scoped to an advertiser or lower, and the custom
	// dimension must already be set up in DoubleClick Search. The custom
	// dimension name, which appears in DoubleClick Search, is case
	// sensitive.
	// If used in a conversion report, returns the value of the
	// specified custom dimension for the given conversion, if set. This
	// column does not segment the conversion report.
	CustomDimensionName string `json:"customDimensionName,omitempty"`

	// CustomMetricName: Name of a custom metric to include in the report.
	// The report must be scoped to an advertiser or lower, and the custom
	// metric must already be set up in DoubleClick Search. The custom
	// metric name, which appears in DoubleClick Search, is case sensitive.
	CustomMetricName string `json:"customMetricName,omitempty"`

	// EndDate: Inclusive day in YYYY-MM-DD format. When provided, this
	// overrides the overall time range of the report for this column only.
	// Must be provided together with startDate.
	EndDate string `json:"endDate,omitempty"`

	// GroupByColumn: Synchronous report only. Set to true to group by this
	// column. Defaults to false.
	GroupByColumn bool `json:"groupByColumn,omitempty"`

	// HeaderText: Text used to identify this column in the report output;
	// defaults to columnName or savedColumnName when not specified. This
	// can be used to prevent collisions between DoubleClick Search columns
	// and saved columns with the same name.
	HeaderText string `json:"headerText,omitempty"`

	// PlatformSource: The platform that is used to provide data for the
	// custom dimension. Acceptable values are "Floodlight".
	PlatformSource string `json:"platformSource,omitempty"`

	// SavedColumnName: Name of a saved column to include in the report. The
	// report must be scoped at advertiser or lower, and this saved column
	// must already be created in the DoubleClick Search UI.
	SavedColumnName string `json:"savedColumnName,omitempty"`

	// StartDate: Inclusive date in YYYY-MM-DD format. When provided, this
	// overrides the overall time range of the report for this column only.
	// Must be provided together with endDate.
	StartDate string `json:"startDate,omitempty"`
}

type ReportRequest struct {
	// Columns: The columns to include in the report. This includes both
	// DoubleClick Search columns and saved columns. For DoubleClick Search
	// columns, only the columnName parameter is required. For saved columns
	// only the savedColumnName parameter is required. Both columnName and
	// savedColumnName cannot be set in the same stanza.
	Columns []*ReportApiColumnSpec `json:"columns,omitempty"`

	// DownloadFormat: Format that the report should be returned in.
	// Currently csv or tsv is supported.
	DownloadFormat string `json:"downloadFormat,omitempty"`

	// Filters: A list of filters to be applied to the report.
	Filters []*ReportRequestFilters `json:"filters,omitempty"`

	// IncludeDeletedEntities: Determines if removed entities should be
	// included in the report. Defaults to false. Deprecated, please use
	// includeRemovedEntities instead.
	IncludeDeletedEntities bool `json:"includeDeletedEntities,omitempty"`

	// IncludeRemovedEntities: Determines if removed entities should be
	// included in the report. Defaults to false.
	IncludeRemovedEntities bool `json:"includeRemovedEntities,omitempty"`

	// MaxRowsPerFile: Asynchronous report only. The maximum number of rows
	// per report file. A large report is split into many files based on
	// this field. Acceptable values are 1000000 to 100000000, inclusive.
	MaxRowsPerFile int64 `json:"maxRowsPerFile,omitempty"`

	// OrderBy: Synchronous report only. A list of columns and directions
	// defining sorting to be performed on the report rows.
	OrderBy []*ReportRequestOrderBy `json:"orderBy,omitempty"`

	// ReportScope: The reportScope is a set of IDs that are used to
	// determine which subset of entities will be returned in the report.
	// The full lineage of IDs from the lowest scoped level desired up
	// through agency is required.
	ReportScope *ReportRequestReportScope `json:"reportScope,omitempty"`

	// ReportType: Determines the type of rows that are returned in the
	// report. For example, if you specify reportType: keyword, each row in
	// the report will contain data about a keyword. See the Types of
	// Reports reference for the columns that are available for each type.
	ReportType string `json:"reportType,omitempty"`

	// RowCount: Synchronous report only. The maxinum number of rows to
	// return; additional rows are dropped. Acceptable values are 0 to
	// 10000, inclusive. Defaults to 10000.
	RowCount int64 `json:"rowCount,omitempty"`

	// StartRow: Synchronous report only. Zero-based index of the first row
	// to return. Acceptable values are 0 to 50000, inclusive. Defaults to
	// 0.
	StartRow int64 `json:"startRow,omitempty"`

	// StatisticsCurrency: Specifies the currency in which monetary will be
	// returned. Possible values are: usd, agency (valid if the report is
	// scoped to agency or lower), advertiser (valid if the report is scoped
	// to * advertiser or lower), or account (valid if the report is scoped
	// to engine account or lower).
	StatisticsCurrency string `json:"statisticsCurrency,omitempty"`

	// TimeRange: If metrics are requested in a report, this argument will
	// be used to restrict the metrics to a specific time range.
	TimeRange *ReportRequestTimeRange `json:"timeRange,omitempty"`

	// VerifySingleTimeZone: If true, the report would only be created if
	// all the requested stat data are sourced from a single timezone.
	// Defaults to false.
	VerifySingleTimeZone bool `json:"verifySingleTimeZone,omitempty"`
}

type ReportRequestFilters struct {
	// Column: Column to perform the filter on. This can be a DoubleClick
	// Search column or a saved column.
	Column *ReportApiColumnSpec `json:"column,omitempty"`

	// Operator: Operator to use in the filter. See the filter reference for
	// a list of available operators.
	Operator string `json:"operator,omitempty"`

	// Values: A list of values to filter the column value against.
	Values []interface{} `json:"values,omitempty"`
}

type ReportRequestOrderBy struct {
	// Column: Column to perform the sort on. This can be a DoubleClick
	// Search-defined column or a saved column.
	Column *ReportApiColumnSpec `json:"column,omitempty"`

	// SortOrder: The sort direction, which is either ascending or
	// descending.
	SortOrder string `json:"sortOrder,omitempty"`
}

type ReportRequestReportScope struct {
	// AdGroupId: DS ad group ID.
	AdGroupId int64 `json:"adGroupId,omitempty,string"`

	// AdId: DS ad ID.
	AdId int64 `json:"adId,omitempty,string"`

	// AdvertiserId: DS advertiser ID.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AgencyId: DS agency ID.
	AgencyId int64 `json:"agencyId,omitempty,string"`

	// CampaignId: DS campaign ID.
	CampaignId int64 `json:"campaignId,omitempty,string"`

	// EngineAccountId: DS engine account ID.
	EngineAccountId int64 `json:"engineAccountId,omitempty,string"`

	// KeywordId: DS keyword ID.
	KeywordId int64 `json:"keywordId,omitempty,string"`
}

type ReportRequestTimeRange struct {
	// ChangedAttributesSinceTimestamp: Inclusive UTC timestamp in RFC
	// format, e.g., 2013-07-16T10:16:23.555Z. See additional references on
	// how changed attribute reports work.
	ChangedAttributesSinceTimestamp string `json:"changedAttributesSinceTimestamp,omitempty"`

	// ChangedMetricsSinceTimestamp: Inclusive UTC timestamp in RFC format,
	// e.g., 2013-07-16T10:16:23.555Z. See additional references on how
	// changed metrics reports work.
	ChangedMetricsSinceTimestamp string `json:"changedMetricsSinceTimestamp,omitempty"`

	// EndDate: Inclusive date in YYYY-MM-DD format.
	EndDate string `json:"endDate,omitempty"`

	// StartDate: Inclusive date in YYYY-MM-DD format.
	StartDate string `json:"startDate,omitempty"`
}

type ReportRow struct {
}

type SavedColumn struct {
	// Kind: Identifies this as a SavedColumn resource. Value: the fixed
	// string doubleclicksearch#savedColumn.
	Kind string `json:"kind,omitempty"`

	// SavedColumnName: The name of the saved column.
	SavedColumnName string `json:"savedColumnName,omitempty"`

	// Type: The type of data this saved column will produce.
	Type string `json:"type,omitempty"`
}

type SavedColumnList struct {
	// Items: The saved columns being requested.
	Items []*SavedColumn `json:"items,omitempty"`

	// Kind: Identifies this as a SavedColumnList resource. Value: the fixed
	// string doubleclicksearch#savedColumnList.
	Kind string `json:"kind,omitempty"`
}

type UpdateAvailabilityRequest struct {
	// Availabilities: The availabilities being requested.
	Availabilities []*Availability `json:"availabilities,omitempty"`
}

type UpdateAvailabilityResponse struct {
	// Availabilities: The availabilities being returned.
	Availabilities []*Availability `json:"availabilities,omitempty"`
}

// method id "doubleclicksearch.conversion.get":

type ConversionGetCall struct {
	s               *Service
	agencyId        int64
	advertiserId    int64
	engineAccountId int64
	endDate         int64
	rowCount        int64
	startDate       int64
	startRow        int64
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Get: Retrieves a list of conversions from a DoubleClick Search engine
// account.

func (r *ConversionService) Get(agencyId int64, advertiserId int64, engineAccountId int64, endDate int64, rowCount int64, startDate int64, startRow int64) *ConversionGetCall {
	return &ConversionGetCall{
		s:               r.s,
		agencyId:        agencyId,
		advertiserId:    advertiserId,
		engineAccountId: engineAccountId,
		endDate:         endDate,
		rowCount:        rowCount,
		startDate:       startDate,
		startRow:        startRow,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "agency/{agencyId}/advertiser/{advertiserId}/engine/{engineAccountId}/conversion",
	}
}

// AdGroupId sets the optional parameter "adGroupId": Numeric ID of the
// ad group.
func (c *ConversionGetCall) AdGroupId(adGroupId int64) *ConversionGetCall {
	c.params_.Set("adGroupId", fmt.Sprintf("%v", adGroupId))
	return c
}

// AdId sets the optional parameter "adId": Numeric ID of the ad.
func (c *ConversionGetCall) AdId(adId int64) *ConversionGetCall {
	c.params_.Set("adId", fmt.Sprintf("%v", adId))
	return c
}

// CampaignId sets the optional parameter "campaignId": Numeric ID of
// the campaign.
func (c *ConversionGetCall) CampaignId(campaignId int64) *ConversionGetCall {
	c.params_.Set("campaignId", fmt.Sprintf("%v", campaignId))
	return c
}

// CriterionId sets the optional parameter "criterionId": Numeric ID of
// the criterion.
func (c *ConversionGetCall) CriterionId(criterionId int64) *ConversionGetCall {
	c.params_.Set("criterionId", fmt.Sprintf("%v", criterionId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ConversionGetCall) Fields(s ...googleapi.Field) *ConversionGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ConversionGetCall) Do() (*ConversionList, error) {
	var returnValue *ConversionList
	c.params_.Set("endDate", fmt.Sprintf("%v", c.endDate))
	c.params_.Set("rowCount", fmt.Sprintf("%v", c.rowCount))
	c.params_.Set("startDate", fmt.Sprintf("%v", c.startDate))
	c.params_.Set("startRow", fmt.Sprintf("%v", c.startRow))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"agencyId":        strconv.FormatInt(c.agencyId, 10),
		"advertiserId":    strconv.FormatInt(c.advertiserId, 10),
		"engineAccountId": strconv.FormatInt(c.engineAccountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of conversions from a DoubleClick Search engine account.",
	//   "httpMethod": "GET",
	//   "id": "doubleclicksearch.conversion.get",
	//   "parameterOrder": [
	//     "agencyId",
	//     "advertiserId",
	//     "engineAccountId",
	//     "endDate",
	//     "rowCount",
	//     "startDate",
	//     "startRow"
	//   ],
	//   "parameters": {
	//     "adGroupId": {
	//       "description": "Numeric ID of the ad group.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "adId": {
	//       "description": "Numeric ID of the ad.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "advertiserId": {
	//       "description": "Numeric ID of the advertiser.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "agencyId": {
	//       "description": "Numeric ID of the agency.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "campaignId": {
	//       "description": "Numeric ID of the campaign.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "criterionId": {
	//       "description": "Numeric ID of the criterion.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Last date (inclusive) on which to retrieve conversions. Format is yyyymmdd.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "99991231",
	//       "minimum": "20091101",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "engineAccountId": {
	//       "description": "Numeric ID of the engine account.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rowCount": {
	//       "description": "The number of conversions to return per call.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "startDate": {
	//       "description": "First date (inclusive) on which to retrieve conversions. Format is yyyymmdd.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "99991231",
	//       "minimum": "20091101",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "startRow": {
	//       "description": "The 0-based starting index for retrieving conversions results.",
	//       "format": "uint32",
	//       "location": "query",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "agency/{agencyId}/advertiser/{advertiserId}/engine/{engineAccountId}/conversion",
	//   "response": {
	//     "$ref": "ConversionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.conversion.insert":

type ConversionInsertCall struct {
	s              *Service
	conversionlist *ConversionList
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Insert: Inserts a batch of new conversions into DoubleClick Search.

func (r *ConversionService) Insert(conversionlist *ConversionList) *ConversionInsertCall {
	return &ConversionInsertCall{
		s:              r.s,
		conversionlist: conversionlist,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "conversion",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ConversionInsertCall) Fields(s ...googleapi.Field) *ConversionInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ConversionInsertCall) Do() (*ConversionList, error) {
	var returnValue *ConversionList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.conversionlist,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a batch of new conversions into DoubleClick Search.",
	//   "httpMethod": "POST",
	//   "id": "doubleclicksearch.conversion.insert",
	//   "path": "conversion",
	//   "request": {
	//     "$ref": "ConversionList"
	//   },
	//   "response": {
	//     "$ref": "ConversionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.conversion.patch":

type ConversionPatchCall struct {
	s               *Service
	advertiserId    int64
	agencyId        int64
	endDate         int64
	engineAccountId int64
	rowCount        int64
	startDate       int64
	startRow        int64
	conversionlist  *ConversionList
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Patch: Updates a batch of conversions in DoubleClick Search. This
// method supports patch semantics.

func (r *ConversionService) Patch(advertiserId int64, agencyId int64, endDate int64, engineAccountId int64, rowCount int64, startDate int64, startRow int64, conversionlist *ConversionList) *ConversionPatchCall {
	return &ConversionPatchCall{
		s:               r.s,
		advertiserId:    advertiserId,
		agencyId:        agencyId,
		endDate:         endDate,
		engineAccountId: engineAccountId,
		rowCount:        rowCount,
		startDate:       startDate,
		startRow:        startRow,
		conversionlist:  conversionlist,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "conversion",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ConversionPatchCall) Fields(s ...googleapi.Field) *ConversionPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ConversionPatchCall) Do() (*ConversionList, error) {
	var returnValue *ConversionList
	c.params_.Set("advertiserId", fmt.Sprintf("%v", c.advertiserId))
	c.params_.Set("agencyId", fmt.Sprintf("%v", c.agencyId))
	c.params_.Set("endDate", fmt.Sprintf("%v", c.endDate))
	c.params_.Set("engineAccountId", fmt.Sprintf("%v", c.engineAccountId))
	c.params_.Set("rowCount", fmt.Sprintf("%v", c.rowCount))
	c.params_.Set("startDate", fmt.Sprintf("%v", c.startDate))
	c.params_.Set("startRow", fmt.Sprintf("%v", c.startRow))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.conversionlist,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a batch of conversions in DoubleClick Search. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "doubleclicksearch.conversion.patch",
	//   "parameterOrder": [
	//     "advertiserId",
	//     "agencyId",
	//     "endDate",
	//     "engineAccountId",
	//     "rowCount",
	//     "startDate",
	//     "startRow"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "Numeric ID of the advertiser.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "agencyId": {
	//       "description": "Numeric ID of the agency.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "Last date (inclusive) on which to retrieve conversions. Format is yyyymmdd.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "99991231",
	//       "minimum": "20091101",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "engineAccountId": {
	//       "description": "Numeric ID of the engine account.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rowCount": {
	//       "description": "The number of conversions to return per call.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "startDate": {
	//       "description": "First date (inclusive) on which to retrieve conversions. Format is yyyymmdd.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "99991231",
	//       "minimum": "20091101",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "startRow": {
	//       "description": "The 0-based starting index for retrieving conversions results.",
	//       "format": "uint32",
	//       "location": "query",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "conversion",
	//   "request": {
	//     "$ref": "ConversionList"
	//   },
	//   "response": {
	//     "$ref": "ConversionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.conversion.update":

type ConversionUpdateCall struct {
	s              *Service
	conversionlist *ConversionList
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Update: Updates a batch of conversions in DoubleClick Search.

func (r *ConversionService) Update(conversionlist *ConversionList) *ConversionUpdateCall {
	return &ConversionUpdateCall{
		s:              r.s,
		conversionlist: conversionlist,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "conversion",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ConversionUpdateCall) Fields(s ...googleapi.Field) *ConversionUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ConversionUpdateCall) Do() (*ConversionList, error) {
	var returnValue *ConversionList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.conversionlist,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a batch of conversions in DoubleClick Search.",
	//   "httpMethod": "PUT",
	//   "id": "doubleclicksearch.conversion.update",
	//   "path": "conversion",
	//   "request": {
	//     "$ref": "ConversionList"
	//   },
	//   "response": {
	//     "$ref": "ConversionList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.conversion.updateAvailability":

type ConversionUpdateAvailabilityCall struct {
	s                         *Service
	updateavailabilityrequest *UpdateAvailabilityRequest
	caller_                   googleapi.Caller
	params_                   url.Values
	pathTemplate_             string
}

// UpdateAvailability: Updates the availabilities of a batch of
// floodlight activities in DoubleClick Search.

func (r *ConversionService) UpdateAvailability(updateavailabilityrequest *UpdateAvailabilityRequest) *ConversionUpdateAvailabilityCall {
	return &ConversionUpdateAvailabilityCall{
		s: r.s,
		updateavailabilityrequest: updateavailabilityrequest,
		caller_:                   googleapi.JSONCall{},
		params_:                   make(map[string][]string),
		pathTemplate_:             "conversion/updateAvailability",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ConversionUpdateAvailabilityCall) Fields(s ...googleapi.Field) *ConversionUpdateAvailabilityCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ConversionUpdateAvailabilityCall) Do() (*UpdateAvailabilityResponse, error) {
	var returnValue *UpdateAvailabilityResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.updateavailabilityrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates the availabilities of a batch of floodlight activities in DoubleClick Search.",
	//   "httpMethod": "POST",
	//   "id": "doubleclicksearch.conversion.updateAvailability",
	//   "path": "conversion/updateAvailability",
	//   "request": {
	//     "$ref": "UpdateAvailabilityRequest",
	//     "parameterName": "empty"
	//   },
	//   "response": {
	//     "$ref": "UpdateAvailabilityResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.reports.generate":

type ReportsGenerateCall struct {
	s             *Service
	reportrequest *ReportRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generate: Generates and returns a report immediately.

func (r *ReportsService) Generate(reportrequest *ReportRequest) *ReportsGenerateCall {
	return &ReportsGenerateCall{
		s:             r.s,
		reportrequest: reportrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reports/generate",
	}
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
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.reportrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generates and returns a report immediately.",
	//   "httpMethod": "POST",
	//   "id": "doubleclicksearch.reports.generate",
	//   "path": "reports/generate",
	//   "request": {
	//     "$ref": "ReportRequest",
	//     "parameterName": "reportRequest"
	//   },
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.reports.get":

type ReportsGetCall struct {
	s             *Service
	reportId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Polls for the status of a report request.

func (r *ReportsService) Get(reportId string) *ReportsGetCall {
	return &ReportsGetCall{
		s:             r.s,
		reportId:      reportId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reports/{reportId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsGetCall) Fields(s ...googleapi.Field) *ReportsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsGetCall) Do() (*Report, error) {
	var returnValue *Report
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"reportId": c.reportId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Polls for the status of a report request.",
	//   "httpMethod": "GET",
	//   "id": "doubleclicksearch.reports.get",
	//   "parameterOrder": [
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "reportId": {
	//       "description": "ID of the report request being polled.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "reports/{reportId}",
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.reports.getFile":

type ReportsGetFileCall struct {
	s              *Service
	reportId       string
	reportFragment int64
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// GetFile: Downloads a report file encoded in UTF-8.

func (r *ReportsService) GetFile(reportId string, reportFragment int64) *ReportsGetFileCall {
	return &ReportsGetFileCall{
		s:              r.s,
		reportId:       reportId,
		reportFragment: reportFragment,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "reports/{reportId}/files/{reportFragment}",
	}
}

func (c *ReportsGetFileCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"reportId":       c.reportId,
		"reportFragment": strconv.FormatInt(c.reportFragment, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *ReportsGetFileCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"reportId":       c.reportId,
		"reportFragment": strconv.FormatInt(c.reportFragment, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Downloads a report file encoded in UTF-8.",
	//   "httpMethod": "GET",
	//   "id": "doubleclicksearch.reports.getFile",
	//   "parameterOrder": [
	//     "reportId",
	//     "reportFragment"
	//   ],
	//   "parameters": {
	//     "reportFragment": {
	//       "description": "The index of the report fragment to download.",
	//       "format": "int32",
	//       "location": "path",
	//       "minimum": "0",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "reportId": {
	//       "description": "ID of the report.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "reports/{reportId}/files/{reportFragment}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "doubleclicksearch.reports.request":

type ReportsRequestCall struct {
	s             *Service
	reportrequest *ReportRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Request: Inserts a report request into the reporting system.

func (r *ReportsService) Request(reportrequest *ReportRequest) *ReportsRequestCall {
	return &ReportsRequestCall{
		s:             r.s,
		reportrequest: reportrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reports",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsRequestCall) Fields(s ...googleapi.Field) *ReportsRequestCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsRequestCall) Do() (*Report, error) {
	var returnValue *Report
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.reportrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a report request into the reporting system.",
	//   "httpMethod": "POST",
	//   "id": "doubleclicksearch.reports.request",
	//   "path": "reports",
	//   "request": {
	//     "$ref": "ReportRequest",
	//     "parameterName": "reportRequest"
	//   },
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}

// method id "doubleclicksearch.savedColumns.list":

type SavedColumnsListCall struct {
	s             *Service
	agencyId      int64
	advertiserId  int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve the list of saved columns for a specified advertiser.

func (r *SavedColumnsService) List(agencyId int64, advertiserId int64) *SavedColumnsListCall {
	return &SavedColumnsListCall{
		s:             r.s,
		agencyId:      agencyId,
		advertiserId:  advertiserId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "agency/{agencyId}/advertiser/{advertiserId}/savedcolumns",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SavedColumnsListCall) Fields(s ...googleapi.Field) *SavedColumnsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SavedColumnsListCall) Do() (*SavedColumnList, error) {
	var returnValue *SavedColumnList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"agencyId":     strconv.FormatInt(c.agencyId, 10),
		"advertiserId": strconv.FormatInt(c.advertiserId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve the list of saved columns for a specified advertiser.",
	//   "httpMethod": "GET",
	//   "id": "doubleclicksearch.savedColumns.list",
	//   "parameterOrder": [
	//     "agencyId",
	//     "advertiserId"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "DS ID of the advertiser.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "agencyId": {
	//       "description": "DS ID of the agency.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "agency/{agencyId}/advertiser/{advertiserId}/savedcolumns",
	//   "response": {
	//     "$ref": "SavedColumnList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/doubleclicksearch"
	//   ]
	// }

}
