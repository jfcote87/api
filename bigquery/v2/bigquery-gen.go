// Package bigquery provides access to the BigQuery API.
//
// See https://developers.google.com/bigquery/docs/overview
//
// Usage example:
//
//   import "github.com/jfcote87/api/bigquery/v2"
//   ...
//   bigqueryService, err := bigquery.New(oauthHttpClient)
package bigquery

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

const apiId = "bigquery:v2"
const apiName = "bigquery"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/bigquery/v2/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your data in Google BigQuery
	BigqueryScope = "https://www.googleapis.com/auth/bigquery"

	// Insert data into Google BigQuery
	BigqueryInsertdataScope = "https://www.googleapis.com/auth/bigquery.insertdata"

	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage your data and permissions in Google Cloud Storage
	DevstorageFull_controlScope = "https://www.googleapis.com/auth/devstorage.full_control"

	// View your data in Google Cloud Storage
	DevstorageRead_onlyScope = "https://www.googleapis.com/auth/devstorage.read_only"

	// Manage your data in Google Cloud Storage
	DevstorageRead_writeScope = "https://www.googleapis.com/auth/devstorage.read_write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Datasets = NewDatasetsService(s)
	s.Jobs = NewJobsService(s)
	s.Projects = NewProjectsService(s)
	s.Tabledata = NewTabledataService(s)
	s.Tables = NewTablesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Datasets *DatasetsService

	Jobs *JobsService

	Projects *ProjectsService

	Tabledata *TabledataService

	Tables *TablesService
}

func NewDatasetsService(s *Service) *DatasetsService {
	rs := &DatasetsService{s: s}
	return rs
}

type DatasetsService struct {
	s *Service
}

func NewJobsService(s *Service) *JobsService {
	rs := &JobsService{s: s}
	return rs
}

type JobsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

func NewTabledataService(s *Service) *TabledataService {
	rs := &TabledataService{s: s}
	return rs
}

type TabledataService struct {
	s *Service
}

func NewTablesService(s *Service) *TablesService {
	rs := &TablesService{s: s}
	return rs
}

type TablesService struct {
	s *Service
}

type CsvOptions struct {
	// AllowJaggedRows: [Optional] Indicates if BigQuery should accept rows
	// that are missing trailing optional columns. If true, BigQuery treats
	// missing trailing columns as null values. If false, records with
	// missing trailing columns are treated as bad records, and if there are
	// too many bad records, an invalid error is returned in the job result.
	// The default value is false.
	AllowJaggedRows bool `json:"allowJaggedRows,omitempty"`

	// AllowQuotedNewlines: [Optional] Indicates if BigQuery should allow
	// quoted data sections that contain newline characters in a CSV file.
	// The default value is false.
	AllowQuotedNewlines bool `json:"allowQuotedNewlines,omitempty"`

	// Encoding: [Optional] The character encoding of the data. The
	// supported values are UTF-8 or ISO-8859-1. The default value is UTF-8.
	// BigQuery decodes the data after the raw, binary data has been split
	// using the values of the quote and fieldDelimiter properties.
	Encoding string `json:"encoding,omitempty"`

	// FieldDelimiter: [Optional] The separator for fields in a CSV file.
	// BigQuery converts the string to ISO-8859-1 encoding, and then uses
	// the first byte of the encoded string to split the data in its raw,
	// binary state. BigQuery also supports the escape sequence "\t" to
	// specify a tab separator. The default value is a comma (',').
	FieldDelimiter string `json:"fieldDelimiter,omitempty"`

	// Quote: [Optional] The value that is used to quote data sections in a
	// CSV file. BigQuery converts the string to ISO-8859-1 encoding, and
	// then uses the first byte of the encoded string to split the data in
	// its raw, binary state. The default value is a double-quote ('"'). If
	// your data does not contain quoted sections, set the property value to
	// an empty string. If your data contains quoted newline characters, you
	// must also set the allowQuotedNewlines property to true.
	Quote string `json:"quote,omitempty"`

	// SkipLeadingRows: [Optional] The number of rows at the top of a CSV
	// file that BigQuery will skip when reading the data. The default value
	// is 0. This property is useful if you have header rows in the file
	// that should be skipped.
	SkipLeadingRows int64 `json:"skipLeadingRows,omitempty"`
}

type Dataset struct {
	// Access: [Optional] An array of objects that define dataset access for
	// one or more entities. You can set this property when inserting or
	// updating a dataset in order to control who is allowed to access the
	// data. If unspecified at dataset creation time, BigQuery adds default
	// dataset access for the following entities: access.specialGroup:
	// projectReaders; access.role: READER; access.specialGroup:
	// projectWriters; access.role: WRITER; access.specialGroup:
	// projectOwners; access.role: OWNER; access.userByEmail: [dataset
	// creator email]; access.role: OWNER;
	Access []*DatasetAccess `json:"access,omitempty"`

	// CreationTime: [Output-only] The time when this dataset was created,
	// in milliseconds since the epoch.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// DatasetReference: [Required] A reference that identifies the dataset.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// DefaultTableExpirationMs: [Experimental] The default lifetime of all
	// tables in the dataset, in milliseconds. The minimum value is 3600000
	// milliseconds (one hour). Once this property is set, all newly-created
	// tables in the dataset will have an expirationTime property set to the
	// creation time plus the value in this property, and changing the value
	// will only affect new tables, not existing ones. When the
	// expirationTime for a given table is reached, that table will be
	// deleted automatically. If a table's expirationTime is modified or
	// removed before the table expires, or if you provide an explicit
	// expirationTime when creating a table, that value takes precedence
	// over the default expiration time indicated by this property.
	DefaultTableExpirationMs int64 `json:"defaultTableExpirationMs,omitempty,string"`

	// Description: [Optional] A user-friendly description of the dataset.
	Description string `json:"description,omitempty"`

	// Etag: [Output-only] A hash of the resource.
	Etag string `json:"etag,omitempty"`

	// FriendlyName: [Optional] A descriptive name for the dataset.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Id: [Output-only] The fully-qualified unique name of the dataset in
	// the format projectId:datasetId. The dataset name without the project
	// name is given in the datasetId field. When creating a new dataset,
	// leave this field blank, and instead specify the datasetId field.
	Id string `json:"id,omitempty"`

	// Kind: [Output-only] The resource type.
	Kind string `json:"kind,omitempty"`

	// LastModifiedTime: [Output-only] The date when this dataset or any of
	// its tables was last modified, in milliseconds since the epoch.
	LastModifiedTime int64 `json:"lastModifiedTime,omitempty,string"`

	// SelfLink: [Output-only] A URL that can be used to access the resource
	// again. You can use this URL in Get or Update requests to the
	// resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type DatasetAccess struct {
	// Domain: [Pick one] A domain to grant access to. Any users signed in
	// with the domain specified will be granted the specified access.
	// Example: "example.com".
	Domain string `json:"domain,omitempty"`

	// GroupByEmail: [Pick one] An email address of a Google Group to grant
	// access to.
	GroupByEmail string `json:"groupByEmail,omitempty"`

	// Role: [Required] Describes the rights granted to the user specified
	// by the other member of the access object. The following string values
	// are supported: READER, WRITER, OWNER.
	Role string `json:"role,omitempty"`

	// SpecialGroup: [Pick one] A special group to grant access to. Possible
	// values include: projectOwners: Owners of the enclosing project.
	// projectReaders: Readers of the enclosing project. projectWriters:
	// Writers of the enclosing project. allAuthenticatedUsers: All
	// authenticated BigQuery users.
	SpecialGroup string `json:"specialGroup,omitempty"`

	// UserByEmail: [Pick one] An email address of a user to grant access
	// to. For example: fred@example.com.
	UserByEmail string `json:"userByEmail,omitempty"`

	// View: [Pick one] A view from a different dataset to grant access to.
	// Queries executed against that view will have read access to tables in
	// this dataset. The role field is not required when this field is set.
	// If that view is updated by any user, access to the view needs to be
	// granted again via an update operation.
	View *TableReference `json:"view,omitempty"`
}

type DatasetList struct {
	// Datasets: An array of the dataset resources in the project. Each
	// resource contains basic information. For full information about a
	// particular dataset resource, use the Datasets: get method. This
	// property is omitted when there are no datasets in the project.
	Datasets []*DatasetListDatasets `json:"datasets,omitempty"`

	// Etag: A hash value of the results page. You can use this property to
	// determine if the page has changed since the last request.
	Etag string `json:"etag,omitempty"`

	// Kind: The list type. This property always returns the value
	// "bigquery#datasetList".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token that can be used to request the next results
	// page. This property is omitted on the final results page.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type DatasetListDatasets struct {
	// DatasetReference: The dataset reference. Use this property to access
	// specific parts of the dataset's ID, such as project ID or dataset ID.
	DatasetReference *DatasetReference `json:"datasetReference,omitempty"`

	// FriendlyName: A descriptive name for the dataset, if one exists.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Id: The fully-qualified, unique, opaque ID of the dataset.
	Id string `json:"id,omitempty"`

	// Kind: The resource type. This property always returns the value
	// "bigquery#dataset".
	Kind string `json:"kind,omitempty"`
}

type DatasetReference struct {
	// DatasetId: [Required] A unique ID for this dataset, without the
	// project name. The ID must contain only letters (a-z, A-Z), numbers
	// (0-9), or underscores (_). The maximum length is 1,024 characters.
	DatasetId string `json:"datasetId,omitempty"`

	// ProjectId: [Optional] The ID of the project containing this dataset.
	ProjectId string `json:"projectId,omitempty"`
}

type ErrorProto struct {
	// DebugInfo: Debugging information. This property is internal to Google
	// and should not be used.
	DebugInfo string `json:"debugInfo,omitempty"`

	// Location: Specifies where the error occurred, if present.
	Location string `json:"location,omitempty"`

	// Message: A human-readable description of the error.
	Message string `json:"message,omitempty"`

	// Reason: A short error code that summarizes the error.
	Reason string `json:"reason,omitempty"`
}

type ExternalDataConfiguration struct {
	// Compression: [Optional] The compression type of the data source.
	// Possible values include GZIP and NONE. The default value is NONE.
	Compression string `json:"compression,omitempty"`

	// CsvOptions: Additional properties to set if sourceFormat is set to
	// CSV.
	CsvOptions *CsvOptions `json:"csvOptions,omitempty"`

	// IgnoreUnknownValues: [Optional] Indicates if BigQuery should allow
	// extra values that are not represented in the table schema. If true,
	// the extra values are ignored. If false, records with extra columns
	// are treated as bad records, and if there are too many bad records, an
	// invalid error is returned in the job result. The default value is
	// false. The sourceFormat property determines what BigQuery treats as
	// an extra value: CSV: Trailing columns
	IgnoreUnknownValues bool `json:"ignoreUnknownValues,omitempty"`

	// MaxBadRecords: [Optional] The maximum number of bad records that
	// BigQuery can ignore when reading data. If the number of bad records
	// exceeds this value, an invalid error is returned in the job result.
	// The default value is 0, which requires that all records are valid.
	MaxBadRecords int64 `json:"maxBadRecords,omitempty"`

	// Schema: [Required] The schema for the data.
	Schema *TableSchema `json:"schema,omitempty"`

	// SourceFormat: [Optional] The data format. External data sources must
	// be in CSV format. The default value is CSV.
	SourceFormat string `json:"sourceFormat,omitempty"`

	// SourceUris: [Required] The fully-qualified URIs that point to your
	// data in Google Cloud Storage. Each URI can contain one '*' wildcard
	// character and it must come after the 'bucket' name. CSV limits
	// related to load jobs apply to external data sources, plus an
	// additional limit of 10 GB maximum size across all URIs.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type GetQueryResultsResponse struct {
	// CacheHit: Whether the query result was fetched from the query cache.
	CacheHit bool `json:"cacheHit,omitempty"`

	// Etag: A hash of this response.
	Etag string `json:"etag,omitempty"`

	// JobComplete: Whether the query has completed or not. If rows or
	// totalRows are present, this will always be true. If this is false,
	// totalRows will not be available.
	JobComplete bool `json:"jobComplete,omitempty"`

	// JobReference: Reference to the BigQuery Job that was created to run
	// the query. This field will be present even if the original request
	// timed out, in which case GetQueryResults can be used to read the
	// results once the query has completed. Since this API only returns the
	// first page of results, subsequent pages can be fetched via the same
	// mechanism (GetQueryResults).
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`

	// PageToken: A token used for paging results.
	PageToken string `json:"pageToken,omitempty"`

	// Rows: An object with as many results as can be contained within the
	// maximum permitted reply size. To get any additional rows, you can
	// call GetQueryResults and specify the jobReference returned above.
	// Present only when the query completes successfully.
	Rows []*TableRow `json:"rows,omitempty"`

	// Schema: The schema of the results. Present only when the query
	// completes successfully.
	Schema *TableSchema `json:"schema,omitempty"`

	// TotalBytesProcessed: The total number of bytes processed for this
	// query.
	TotalBytesProcessed int64 `json:"totalBytesProcessed,omitempty,string"`

	// TotalRows: The total number of rows in the complete query result set,
	// which can be more than the number of rows in this single page of
	// results. Present only when the query completes successfully.
	TotalRows uint64 `json:"totalRows,omitempty,string"`
}

type Job struct {
	// Configuration: [Required] Describes the job configuration.
	Configuration *JobConfiguration `json:"configuration,omitempty"`

	// Etag: [Output-only] A hash of this resource.
	Etag string `json:"etag,omitempty"`

	// Id: [Output-only] Opaque ID field of the job
	Id string `json:"id,omitempty"`

	// JobReference: [Optional] Reference describing the unique-per-user
	// name of the job.
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Kind: [Output-only] The type of the resource.
	Kind string `json:"kind,omitempty"`

	// SelfLink: [Output-only] A URL that can be used to access this
	// resource again.
	SelfLink string `json:"selfLink,omitempty"`

	// Statistics: [Output-only] Information about the job, including
	// starting time and ending time of the job.
	Statistics *JobStatistics `json:"statistics,omitempty"`

	// Status: [Output-only] The status of this job. Examine this value when
	// polling an asynchronous job to see if the job is complete.
	Status *JobStatus `json:"status,omitempty"`

	// User_email: [Output-only] Email address of the user who ran the job.
	User_email string `json:"user_email,omitempty"`
}

type JobConfiguration struct {
	// Copy: [Pick one] Copies a table.
	Copy *JobConfigurationTableCopy `json:"copy,omitempty"`

	// DryRun: [Optional] If set, don't actually run this job. A valid query
	// will return a mostly empty response with some processing statistics,
	// while an invalid query will return the same error it would if it
	// wasn't a dry run. Behavior of non-query jobs is undefined.
	DryRun bool `json:"dryRun,omitempty"`

	// Extract: [Pick one] Configures an extract job.
	Extract *JobConfigurationExtract `json:"extract,omitempty"`

	// Link: [Pick one] Configures a link job.
	Link *JobConfigurationLink `json:"link,omitempty"`

	// Load: [Pick one] Configures a load job.
	Load *JobConfigurationLoad `json:"load,omitempty"`

	// Query: [Pick one] Configures a query job.
	Query *JobConfigurationQuery `json:"query,omitempty"`
}

type JobConfigurationExtract struct {
	// Compression: [Optional] The compression type to use for exported
	// files. Possible values include GZIP and NONE. The default value is
	// NONE.
	Compression string `json:"compression,omitempty"`

	// DestinationFormat: [Optional] The exported file format. Possible
	// values include CSV, NEWLINE_DELIMITED_JSON and AVRO. The default
	// value is CSV. Tables with nested or repeated fields cannot be
	// exported as CSV.
	DestinationFormat string `json:"destinationFormat,omitempty"`

	// DestinationUri: [Pick one] DEPRECATED: Use destinationUris instead,
	// passing only one URI as necessary. The fully-qualified Google Cloud
	// Storage URI where the extracted table should be written.
	DestinationUri string `json:"destinationUri,omitempty"`

	// DestinationUris: [Pick one] A list of fully-qualified Google Cloud
	// Storage URIs where the extracted table should be written.
	DestinationUris []string `json:"destinationUris,omitempty"`

	// FieldDelimiter: [Optional] Delimiter to use between fields in the
	// exported data. Default is ','
	FieldDelimiter string `json:"fieldDelimiter,omitempty"`

	// PrintHeader: [Optional] Whether to print out a header row in the
	// results. Default is true.
	PrintHeader bool `json:"printHeader,omitempty"`

	// SourceTable: [Required] A reference to the table being exported.
	SourceTable *TableReference `json:"sourceTable,omitempty"`
}

type JobConfigurationLink struct {
	// CreateDisposition: [Optional] Specifies whether the job is allowed to
	// create new tables. The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	// table. CREATE_NEVER: The table must already exist. If it does not, a
	// 'notFound' error is returned in the job result. The default value is
	// CREATE_IF_NEEDED. Creation, truncation and append actions occur as
	// one atomic update upon job completion.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DestinationTable: [Required] The destination table of the link job.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// SourceUri: [Required] URI of source table to link.
	SourceUri []string `json:"sourceUri,omitempty"`

	// WriteDisposition: [Optional] Specifies the action that occurs if the
	// destination table already exists. The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the
	// table data. WRITE_APPEND: If the table already exists, BigQuery
	// appends the data to the table. WRITE_EMPTY: If the table already
	// exists and contains data, a 'duplicate' error is returned in the job
	// result. The default value is WRITE_EMPTY. Each action is atomic and
	// only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update
	// upon job completion.
	WriteDisposition string `json:"writeDisposition,omitempty"`
}

type JobConfigurationLoad struct {
	// AllowJaggedRows: [Optional] Accept rows that are missing trailing
	// optional columns. The missing values are treated as nulls. If false,
	// records with missing trailing columns are treated as bad records, and
	// if there are too many bad records, an invalid error is returned in
	// the job result. The default value is false. Only applicable to CSV,
	// ignored for other formats.
	AllowJaggedRows bool `json:"allowJaggedRows,omitempty"`

	// AllowQuotedNewlines: Indicates if BigQuery should allow quoted data
	// sections that contain newline characters in a CSV file. The default
	// value is false.
	AllowQuotedNewlines bool `json:"allowQuotedNewlines,omitempty"`

	// CreateDisposition: [Optional] Specifies whether the job is allowed to
	// create new tables. The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	// table. CREATE_NEVER: The table must already exist. If it does not, a
	// 'notFound' error is returned in the job result. The default value is
	// CREATE_IF_NEEDED. Creation, truncation and append actions occur as
	// one atomic update upon job completion.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DestinationTable: [Required] The destination table to load the data
	// into.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// Encoding: [Optional] The character encoding of the data. The
	// supported values are UTF-8 or ISO-8859-1. The default value is UTF-8.
	// BigQuery decodes the data after the raw, binary data has been split
	// using the values of the quote and fieldDelimiter properties.
	Encoding string `json:"encoding,omitempty"`

	// FieldDelimiter: [Optional] The separator for fields in a CSV file.
	// BigQuery converts the string to ISO-8859-1 encoding, and then uses
	// the first byte of the encoded string to split the data in its raw,
	// binary state. BigQuery also supports the escape sequence "\t" to
	// specify a tab separator. The default value is a comma (',').
	FieldDelimiter string `json:"fieldDelimiter,omitempty"`

	// IgnoreUnknownValues: [Optional] Indicates if BigQuery should allow
	// extra values that are not represented in the table schema. If true,
	// the extra values are ignored. If false, records with extra columns
	// are treated as bad records, and if there are too many bad records, an
	// invalid error is returned in the job result. The default value is
	// false. The sourceFormat property determines what BigQuery treats as
	// an extra value: CSV: Trailing columns JSON: Named values that don't
	// match any column names
	IgnoreUnknownValues bool `json:"ignoreUnknownValues,omitempty"`

	// MaxBadRecords: [Optional] The maximum number of bad records that
	// BigQuery can ignore when running the job. If the number of bad
	// records exceeds this value, an invalid error is returned in the job
	// result. The default value is 0, which requires that all records are
	// valid.
	MaxBadRecords int64 `json:"maxBadRecords,omitempty"`

	// ProjectionFields: [Experimental] Names(case-sensitive) of properties
	// to keep when importing data. If this is populated, only the specified
	// properties will be imported for each entity. Currently, this is only
	// supported for DATASTORE_BACKUP imports and only top level properties
	// are supported. If any specified property is not found in the
	// Datastore 'Kind' being imported, that is an error. Note: This feature
	// is experimental and can change in the future.
	ProjectionFields []string `json:"projectionFields,omitempty"`

	// Quote: [Optional] The value that is used to quote data sections in a
	// CSV file. BigQuery converts the string to ISO-8859-1 encoding, and
	// then uses the first byte of the encoded string to split the data in
	// its raw, binary state. The default value is a double-quote ('"'). If
	// your data does not contain quoted sections, set the property value to
	// an empty string. If your data contains quoted newline characters, you
	// must also set the allowQuotedNewlines property to true.
	Quote string `json:"quote,omitempty"`

	// Schema: [Optional] The schema for the destination table. The schema
	// can be omitted if the destination table already exists or if the
	// schema can be inferred from the loaded data.
	Schema *TableSchema `json:"schema,omitempty"`

	// SchemaInline: [Deprecated] The inline schema. For CSV schemas,
	// specify as "Field1:Type1[,Field2:Type2]*". For example, "foo:STRING,
	// bar:INTEGER, baz:FLOAT".
	SchemaInline string `json:"schemaInline,omitempty"`

	// SchemaInlineFormat: [Deprecated] The format of the schemaInline
	// property.
	SchemaInlineFormat string `json:"schemaInlineFormat,omitempty"`

	// SkipLeadingRows: [Optional] The number of rows at the top of a CSV
	// file that BigQuery will skip when loading the data. The default value
	// is 0. This property is useful if you have header rows in the file
	// that should be skipped.
	SkipLeadingRows int64 `json:"skipLeadingRows,omitempty"`

	// SourceFormat: [Optional] The format of the data files. For CSV files,
	// specify "CSV". For datastore backups, specify "DATASTORE_BACKUP". For
	// newline-delimited JSON, specify "NEWLINE_DELIMITED_JSON". The default
	// value is CSV.
	SourceFormat string `json:"sourceFormat,omitempty"`

	// SourceUris: [Required] The fully-qualified URIs that point to your
	// data in Google Cloud Storage. Each URI can contain one '*' wildcard
	// character and it must come after the 'bucket' name.
	SourceUris []string `json:"sourceUris,omitempty"`

	// WriteDisposition: [Optional] Specifies the action that occurs if the
	// destination table already exists. The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the
	// table data. WRITE_APPEND: If the table already exists, BigQuery
	// appends the data to the table. WRITE_EMPTY: If the table already
	// exists and contains data, a 'duplicate' error is returned in the job
	// result. The default value is WRITE_EMPTY. Each action is atomic and
	// only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update
	// upon job completion.
	WriteDisposition string `json:"writeDisposition,omitempty"`
}

type JobConfigurationQuery struct {
	// AllowLargeResults: If true, allows the query to produce arbitrarily
	// large result tables at a slight cost in performance. Requires
	// destinationTable to be set.
	AllowLargeResults bool `json:"allowLargeResults,omitempty"`

	// CreateDisposition: [Optional] Specifies whether the job is allowed to
	// create new tables. The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	// table. CREATE_NEVER: The table must already exist. If it does not, a
	// 'notFound' error is returned in the job result. The default value is
	// CREATE_IF_NEEDED. Creation, truncation and append actions occur as
	// one atomic update upon job completion.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DefaultDataset: [Optional] Specifies the default dataset to use for
	// unqualified table names in the query.
	DefaultDataset *DatasetReference `json:"defaultDataset,omitempty"`

	// DestinationTable: [Optional] Describes the table where the query
	// results should be stored. If not present, a new table will be created
	// to store the results.
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// FlattenResults: [Experimental] Flattens all nested and repeated
	// fields in the query results. The default value is true.
	// allowLargeResults must be true if this is set to false.
	FlattenResults bool `json:"flattenResults,omitempty"`

	// PreserveNulls: [Deprecated] This property is deprecated.
	PreserveNulls bool `json:"preserveNulls,omitempty"`

	// Priority: [Optional] Specifies a priority for the query. Possible
	// values include INTERACTIVE and BATCH. The default value is
	// INTERACTIVE.
	Priority string `json:"priority,omitempty"`

	// Query: [Required] BigQuery SQL query to execute.
	Query string `json:"query,omitempty"`

	// TableDefinitions: [Experimental] If querying an external data source
	// outside of BigQuery, describes the data format, location and other
	// properties of the data source. By defining these properties, the data
	// source can then be queried as if it were a standard BigQuery table.
	TableDefinitions map[string]ExternalDataConfiguration `json:"tableDefinitions,omitempty"`

	// UseQueryCache: [Optional] Whether to look for the result in the query
	// cache. The query cache is a best-effort cache that will be flushed
	// whenever tables in the query are modified. Moreover, the query cache
	// is only available when a query does not have a destination table
	// specified.
	UseQueryCache bool `json:"useQueryCache,omitempty"`

	// WriteDisposition: [Optional] Specifies the action that occurs if the
	// destination table already exists. The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the
	// table data. WRITE_APPEND: If the table already exists, BigQuery
	// appends the data to the table. WRITE_EMPTY: If the table already
	// exists and contains data, a 'duplicate' error is returned in the job
	// result. The default value is WRITE_EMPTY. Each action is atomic and
	// only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update
	// upon job completion.
	WriteDisposition string `json:"writeDisposition,omitempty"`
}

type JobConfigurationTableCopy struct {
	// CreateDisposition: [Optional] Specifies whether the job is allowed to
	// create new tables. The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the
	// table. CREATE_NEVER: The table must already exist. If it does not, a
	// 'notFound' error is returned in the job result. The default value is
	// CREATE_IF_NEEDED. Creation, truncation and append actions occur as
	// one atomic update upon job completion.
	CreateDisposition string `json:"createDisposition,omitempty"`

	// DestinationTable: [Required] The destination table
	DestinationTable *TableReference `json:"destinationTable,omitempty"`

	// SourceTable: [Pick one] Source table to copy.
	SourceTable *TableReference `json:"sourceTable,omitempty"`

	// SourceTables: [Pick one] Source tables to copy.
	SourceTables []*TableReference `json:"sourceTables,omitempty"`

	// WriteDisposition: [Optional] Specifies the action that occurs if the
	// destination table already exists. The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the
	// table data. WRITE_APPEND: If the table already exists, BigQuery
	// appends the data to the table. WRITE_EMPTY: If the table already
	// exists and contains data, a 'duplicate' error is returned in the job
	// result. The default value is WRITE_EMPTY. Each action is atomic and
	// only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update
	// upon job completion.
	WriteDisposition string `json:"writeDisposition,omitempty"`
}

type JobList struct {
	// Etag: A hash of this page of results.
	Etag string `json:"etag,omitempty"`

	// Jobs: List of jobs that were requested.
	Jobs []*JobListJobs `json:"jobs,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to request the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of jobs in this collection.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type JobListJobs struct {
	// Configuration: [Full-projection-only] Specifies the job
	// configuration.
	Configuration *JobConfiguration `json:"configuration,omitempty"`

	// ErrorResult: A result object that will be present only if the job has
	// failed.
	ErrorResult *ErrorProto `json:"errorResult,omitempty"`

	// Id: Unique opaque ID of the job.
	Id string `json:"id,omitempty"`

	// JobReference: Job reference uniquely identifying the job.
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// State: Running state of the job. When the state is DONE, errorResult
	// can be checked to determine whether the job succeeded or failed.
	State string `json:"state,omitempty"`

	// Statistics: [Output-only] Information about the job, including
	// starting time and ending time of the job.
	Statistics *JobStatistics `json:"statistics,omitempty"`

	// Status: [Full-projection-only] Describes the state of the job.
	Status *JobStatus `json:"status,omitempty"`

	// User_email: [Full-projection-only] Email address of the user who ran
	// the job.
	User_email string `json:"user_email,omitempty"`
}

type JobReference struct {
	// JobId: [Required] The ID of the job. The ID must contain only letters
	// (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The
	// maximum length is 1,024 characters.
	JobId string `json:"jobId,omitempty"`

	// ProjectId: [Required] The ID of the project containing this job.
	ProjectId string `json:"projectId,omitempty"`
}

type JobStatistics struct {
	// CreationTime: [Output-only] Creation time of this job, in
	// milliseconds since the epoch. This field will be present on all jobs.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// EndTime: [Output-only] End time of this job, in milliseconds since
	// the epoch. This field will be present whenever a job is in the DONE
	// state.
	EndTime int64 `json:"endTime,omitempty,string"`

	// Extract: [Output-only] Statistics for an extract job.
	Extract *JobStatistics4 `json:"extract,omitempty"`

	// Load: [Output-only] Statistics for a load job.
	Load *JobStatistics3 `json:"load,omitempty"`

	// Query: [Output-only] Statistics for a query job.
	Query *JobStatistics2 `json:"query,omitempty"`

	// StartTime: [Output-only] Start time of this job, in milliseconds
	// since the epoch. This field will be present when the job transitions
	// from the PENDING state to either RUNNING or DONE.
	StartTime int64 `json:"startTime,omitempty,string"`

	// TotalBytesProcessed: [Output-only] [Deprecated] Use the bytes
	// processed in the query statistics instead.
	TotalBytesProcessed int64 `json:"totalBytesProcessed,omitempty,string"`
}

type JobStatistics2 struct {
	// CacheHit: [Output-only] Whether the query result was fetched from the
	// query cache.
	CacheHit bool `json:"cacheHit,omitempty"`

	// TotalBytesProcessed: [Output-only] Total bytes processed for this
	// job.
	TotalBytesProcessed int64 `json:"totalBytesProcessed,omitempty,string"`
}

type JobStatistics3 struct {
	// InputFileBytes: [Output-only] Number of bytes of source data in a
	// joad job.
	InputFileBytes int64 `json:"inputFileBytes,omitempty,string"`

	// InputFiles: [Output-only] Number of source files in a load job.
	InputFiles int64 `json:"inputFiles,omitempty,string"`

	// OutputBytes: [Output-only] Size of the loaded data in bytes. Note
	// that while an import job is in the running state, this value may
	// change.
	OutputBytes int64 `json:"outputBytes,omitempty,string"`

	// OutputRows: [Output-only] Number of rows imported in a load job. Note
	// that while an import job is in the running state, this value may
	// change.
	OutputRows int64 `json:"outputRows,omitempty,string"`
}

type JobStatistics4 struct {
	// DestinationUriFileCounts: [Experimental] Number of files per
	// destination URI or URI pattern specified in the extract
	// configuration. These values will be in the same order as the URIs
	// specified in the 'destinationUris' field.
	DestinationUriFileCounts googleapi.Int64s `json:"destinationUriFileCounts,omitempty"`
}

type JobStatus struct {
	// ErrorResult: [Output-only] Final error result of the job. If present,
	// indicates that the job has completed and was unsuccessful.
	ErrorResult *ErrorProto `json:"errorResult,omitempty"`

	// Errors: [Output-only] All errors encountered during the running of
	// the job. Errors here do not necessarily mean that the job has
	// completed or was unsuccessful.
	Errors []*ErrorProto `json:"errors,omitempty"`

	// State: [Output-only] Running state of the job.
	State string `json:"state,omitempty"`
}

type JsonValue interface{}

type ProjectList struct {
	// Etag: A hash of the page of results
	Etag string `json:"etag,omitempty"`

	// Kind: The type of list.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to request the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Projects: Projects to which you have at least READ access.
	Projects []*ProjectListProjects `json:"projects,omitempty"`

	// TotalItems: The total number of projects in the list.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type ProjectListProjects struct {
	// FriendlyName: A descriptive name for this project.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Id: An opaque ID of this project.
	Id string `json:"id,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// NumericId: The numeric ID of this project.
	NumericId uint64 `json:"numericId,omitempty,string"`

	// ProjectReference: A unique reference to this project.
	ProjectReference *ProjectReference `json:"projectReference,omitempty"`
}

type ProjectReference struct {
	// ProjectId: [Required] ID of the project. Can be either the numeric ID
	// or the assigned ID of the project.
	ProjectId string `json:"projectId,omitempty"`
}

type QueryRequest struct {
	// DefaultDataset: [Optional] Specifies the default datasetId and
	// projectId to assume for any unqualified table names in the query. If
	// not set, all table names in the query string must be qualified in the
	// format 'datasetId.tableId'.
	DefaultDataset *DatasetReference `json:"defaultDataset,omitempty"`

	// DryRun: [Optional] If set, don't actually run this job. A valid query
	// will return a mostly empty response with some processing statistics,
	// while an invalid query will return the same error it would if it
	// wasn't a dry run.
	DryRun bool `json:"dryRun,omitempty"`

	// Kind: The resource type of the request.
	Kind string `json:"kind,omitempty"`

	// MaxResults: [Optional] The maximum number of rows of data to return
	// per page of results. Setting this flag to a small value such as 1000
	// and then paging through results might improve reliability when the
	// query result set is large. In addition to this limit, responses are
	// also limited to 10 MB. By default, there is no maximum row count, and
	// only the byte limit applies.
	MaxResults int64 `json:"maxResults,omitempty"`

	// PreserveNulls: [Deprecated] This property is deprecated.
	PreserveNulls bool `json:"preserveNulls,omitempty"`

	// Query: [Required] A query string, following the BigQuery query
	// syntax, of the query to execute. Example: "SELECT count(f1) FROM
	// [myProjectId:myDatasetId.myTableId]".
	Query string `json:"query,omitempty"`

	// TimeoutMs: [Optional] How long to wait for the query to complete, in
	// milliseconds, before the request times out and returns. Note that
	// this is only a timeout for the request, not the query. If the query
	// takes longer to run than the timeout value, the call returns without
	// any results and with the 'jobComplete' flag set to false. You can
	// call GetQueryResults() to wait for the query to complete and read the
	// results. The default value is 10000 milliseconds (10 seconds).
	TimeoutMs int64 `json:"timeoutMs,omitempty"`

	// UseQueryCache: [Optional] Whether to look for the result in the query
	// cache. The query cache is a best-effort cache that will be flushed
	// whenever tables in the query are modified. The default value is true.
	UseQueryCache bool `json:"useQueryCache,omitempty"`
}

type QueryResponse struct {
	// CacheHit: Whether the query result was fetched from the query cache.
	CacheHit bool `json:"cacheHit,omitempty"`

	// JobComplete: Whether the query has completed or not. If rows or
	// totalRows are present, this will always be true. If this is false,
	// totalRows will not be available.
	JobComplete bool `json:"jobComplete,omitempty"`

	// JobReference: Reference to the Job that was created to run the query.
	// This field will be present even if the original request timed out, in
	// which case GetQueryResults can be used to read the results once the
	// query has completed. Since this API only returns the first page of
	// results, subsequent pages can be fetched via the same mechanism
	// (GetQueryResults).
	JobReference *JobReference `json:"jobReference,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// PageToken: A token used for paging results.
	PageToken string `json:"pageToken,omitempty"`

	// Rows: An object with as many results as can be contained within the
	// maximum permitted reply size. To get any additional rows, you can
	// call GetQueryResults and specify the jobReference returned above.
	Rows []*TableRow `json:"rows,omitempty"`

	// Schema: The schema of the results. Present only when the query
	// completes successfully.
	Schema *TableSchema `json:"schema,omitempty"`

	// TotalBytesProcessed: The total number of bytes processed for this
	// query. If this query was a dry run, this is the number of bytes that
	// would be processed if the query were run.
	TotalBytesProcessed int64 `json:"totalBytesProcessed,omitempty,string"`

	// TotalRows: The total number of rows in the complete query result set,
	// which can be more than the number of rows in this single page of
	// results.
	TotalRows uint64 `json:"totalRows,omitempty,string"`
}

type Table struct {
	// CreationTime: [Output-only] The time when this table was created, in
	// milliseconds since the epoch.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// Description: [Optional] A user-friendly description of this table.
	Description string `json:"description,omitempty"`

	// Etag: [Output-only] A hash of this resource.
	Etag string `json:"etag,omitempty"`

	// ExpirationTime: [Optional] The time when this table expires, in
	// milliseconds since the epoch. If not present, the table will persist
	// indefinitely. Expired tables will be deleted and their storage
	// reclaimed.
	ExpirationTime int64 `json:"expirationTime,omitempty,string"`

	// FriendlyName: [Optional] A descriptive name for this table.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Id: [Output-only] An opaque ID uniquely identifying the table.
	Id string `json:"id,omitempty"`

	// Kind: [Output-only] The type of the resource.
	Kind string `json:"kind,omitempty"`

	// LastModifiedTime: [Output-only] The time when this table was last
	// modified, in milliseconds since the epoch.
	LastModifiedTime uint64 `json:"lastModifiedTime,omitempty,string"`

	// NumBytes: [Output-only] The size of the table in bytes. This property
	// is unavailable for tables that are actively receiving streaming
	// inserts.
	NumBytes int64 `json:"numBytes,omitempty,string"`

	// NumRows: [Output-only] The number of rows of data in this table. This
	// property is unavailable for tables that are actively receiving
	// streaming inserts.
	NumRows uint64 `json:"numRows,omitempty,string"`

	// Schema: [Optional] Describes the schema of this table.
	Schema *TableSchema `json:"schema,omitempty"`

	// SelfLink: [Output-only] A URL that can be used to access this
	// resource again.
	SelfLink string `json:"selfLink,omitempty"`

	// TableReference: [Required] Reference describing the ID of this table.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// Type: [Output-only] Describes the table type. The following values
	// are supported: TABLE: A normal BigQuery table. VIEW: A virtual table
	// defined by a SQL query. The default value is TABLE.
	Type string `json:"type,omitempty"`

	// View: [Optional] The view definition.
	View *ViewDefinition `json:"view,omitempty"`
}

type TableCell struct {
	V interface{} `json:"v,omitempty"`
}

type TableDataInsertAllRequest struct {
	// IgnoreUnknownValues: [Optional] Accept rows that contain values that
	// do not match the schema. The unknown values are ignored. Default is
	// false, which treats unknown values as errors.
	IgnoreUnknownValues bool `json:"ignoreUnknownValues,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`

	// Rows: The rows to insert.
	Rows []*TableDataInsertAllRequestRows `json:"rows,omitempty"`

	// SkipInvalidRows: [Optional] Insert all valid rows of a request, even
	// if invalid rows exist. The default value is false, which causes the
	// entire request to fail if any invalid rows exist.
	SkipInvalidRows bool `json:"skipInvalidRows,omitempty"`
}

type TableDataInsertAllRequestRows struct {
	// InsertId: [Optional] A unique ID for each row. BigQuery uses this
	// property to detect duplicate insertion requests on a best-effort
	// basis.
	InsertId string `json:"insertId,omitempty"`

	// Json: [Required] A JSON object that contains a row of data. The
	// object's properties and values must match the destination table's
	// schema.
	Json map[string]JsonValue `json:"json,omitempty"`
}

type TableDataInsertAllResponse struct {
	// InsertErrors: An array of errors for rows that were not inserted.
	InsertErrors []*TableDataInsertAllResponseInsertErrors `json:"insertErrors,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`
}

type TableDataInsertAllResponseInsertErrors struct {
	// Errors: Error information for the row indicated by the index
	// property.
	Errors []*ErrorProto `json:"errors,omitempty"`

	// Index: The index of the row that error applies to.
	Index int64 `json:"index,omitempty"`
}

type TableDataList struct {
	// Etag: A hash of this page of results.
	Etag string `json:"etag,omitempty"`

	// Kind: The resource type of the response.
	Kind string `json:"kind,omitempty"`

	// PageToken: A token used for paging results. Providing this token
	// instead of the startIndex parameter can help you retrieve stable
	// results when an underlying table is changing.
	PageToken string `json:"pageToken,omitempty"`

	// Rows: Rows of results.
	Rows []*TableRow `json:"rows,omitempty"`

	// TotalRows: The total number of rows in the complete table.
	TotalRows int64 `json:"totalRows,omitempty,string"`
}

type TableFieldSchema struct {
	// Description: [Optional] The field description. The maximum length is
	// 16K characters.
	Description string `json:"description,omitempty"`

	// Fields: [Optional] Describes the nested schema fields if the type
	// property is set to RECORD.
	Fields []*TableFieldSchema `json:"fields,omitempty"`

	// Mode: [Optional] The field mode. Possible values include NULLABLE,
	// REQUIRED and REPEATED. The default value is NULLABLE.
	Mode string `json:"mode,omitempty"`

	// Name: [Required] The field name. The name must contain only letters
	// (a-z, A-Z), numbers (0-9), or underscores (_), and must start with a
	// letter or underscore. The maximum length is 128 characters.
	Name string `json:"name,omitempty"`

	// Type: [Required] The field data type. Possible values include STRING,
	// INTEGER, FLOAT, BOOLEAN, TIMESTAMP or RECORD (where RECORD indicates
	// that the field contains a nested schema).
	Type string `json:"type,omitempty"`
}

type TableList struct {
	// Etag: A hash of this page of results.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of list.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token to request the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Tables: Tables in the requested dataset.
	Tables []*TableListTables `json:"tables,omitempty"`

	// TotalItems: The total number of tables in the dataset.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type TableListTables struct {
	// FriendlyName: The user-friendly name for this table.
	FriendlyName string `json:"friendlyName,omitempty"`

	// Id: An opaque ID of the table
	Id string `json:"id,omitempty"`

	// Kind: The resource type.
	Kind string `json:"kind,omitempty"`

	// TableReference: A reference uniquely identifying the table.
	TableReference *TableReference `json:"tableReference,omitempty"`

	// Type: The type of table. Possible values are: TABLE, VIEW.
	Type string `json:"type,omitempty"`
}

type TableReference struct {
	// DatasetId: [Required] The ID of the dataset containing this table.
	DatasetId string `json:"datasetId,omitempty"`

	// ProjectId: [Required] The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty"`

	// TableId: [Required] The ID of the table. The ID must contain only
	// letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 1,024 characters.
	TableId string `json:"tableId,omitempty"`
}

type TableRow struct {
	F []*TableCell `json:"f,omitempty"`
}

type TableSchema struct {
	// Fields: Describes the fields in a table.
	Fields []*TableFieldSchema `json:"fields,omitempty"`
}

type ViewDefinition struct {
	// Query: [Required] A query that BigQuery executes when the view is
	// referenced.
	Query string `json:"query,omitempty"`
}

// method id "bigquery.datasets.delete":

type DatasetsDeleteCall struct {
	s             *Service
	projectId     string
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the dataset specified by the datasetId value. Before
// you can delete a dataset, you must delete all its tables, either
// manually or by specifying deleteContents. Immediately after deletion,
// you can create another dataset with the same name.

func (r *DatasetsService) Delete(projectId string, datasetId string) *DatasetsDeleteCall {
	return &DatasetsDeleteCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}",
	}
}

// DeleteContents sets the optional parameter "deleteContents": If True,
// delete all the tables in the dataset. If False and the dataset
// contains tables, the request will fail. Default is False
func (c *DatasetsDeleteCall) DeleteContents(deleteContents bool) *DatasetsDeleteCall {
	c.params_.Set("deleteContents", fmt.Sprintf("%v", deleteContents))
	return c
}

func (c *DatasetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the dataset specified by the datasetId value. Before you can delete a dataset, you must delete all its tables, either manually or by specifying deleteContents. Immediately after deletion, you can create another dataset with the same name.",
	//   "httpMethod": "DELETE",
	//   "id": "bigquery.datasets.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of dataset being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deleteContents": {
	//       "description": "If True, delete all the tables in the dataset. If False and the dataset contains tables, the request will fail. Default is False",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the dataset being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.datasets.get":

type DatasetsGetCall struct {
	s             *Service
	projectId     string
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns the dataset specified by datasetID.

func (r *DatasetsService) Get(projectId string, datasetId string) *DatasetsGetCall {
	return &DatasetsGetCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsGetCall) Fields(s ...googleapi.Field) *DatasetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsGetCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the dataset specified by datasetID.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.datasets.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the requested dataset",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the requested dataset",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.datasets.insert":

type DatasetsInsertCall struct {
	s             *Service
	projectId     string
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new empty dataset.

func (r *DatasetsService) Insert(projectId string, dataset *Dataset) *DatasetsInsertCall {
	return &DatasetsInsertCall{
		s:             r.s,
		projectId:     projectId,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsInsertCall) Fields(s ...googleapi.Field) *DatasetsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsInsertCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.dataset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new empty dataset.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.datasets.insert",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID of the new dataset",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.datasets.list":

type DatasetsListCall struct {
	s             *Service
	projectId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all the datasets in the specified project to which the
// caller has read access; however, a project owner can list (but not
// necessarily get) all datasets in his project.

func (r *DatasetsService) List(projectId string) *DatasetsListCall {
	return &DatasetsListCall{
		s:             r.s,
		projectId:     projectId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets",
	}
}

// All sets the optional parameter "all": Whether to list all datasets,
// including hidden ones
func (c *DatasetsListCall) All(all bool) *DatasetsListCall {
	c.params_.Set("all", fmt.Sprintf("%v", all))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results to return
func (c *DatasetsListCall) MaxResults(maxResults int64) *DatasetsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *DatasetsListCall) PageToken(pageToken string) *DatasetsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsListCall) Fields(s ...googleapi.Field) *DatasetsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsListCall) Do() (*DatasetList, error) {
	var returnValue *DatasetList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all the datasets in the specified project to which the caller has read access; however, a project owner can list (but not necessarily get) all datasets in his project.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.datasets.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "all": {
	//       "description": "Whether to list all datasets, including hidden ones",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the datasets to be listed",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets",
	//   "response": {
	//     "$ref": "DatasetList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.datasets.patch":

type DatasetsPatchCall struct {
	s             *Service
	projectId     string
	datasetId     string
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates information in an existing dataset. The update method
// replaces the entire dataset resource, whereas the patch method only
// replaces fields that are provided in the submitted dataset resource.
// This method supports patch semantics.

func (r *DatasetsService) Patch(projectId string, datasetId string, dataset *Dataset) *DatasetsPatchCall {
	return &DatasetsPatchCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsPatchCall) Fields(s ...googleapi.Field) *DatasetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsPatchCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.dataset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates information in an existing dataset. The update method replaces the entire dataset resource, whereas the patch method only replaces fields that are provided in the submitted dataset resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "bigquery.datasets.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.datasets.update":

type DatasetsUpdateCall struct {
	s             *Service
	projectId     string
	datasetId     string
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates information in an existing dataset. The update method
// replaces the entire dataset resource, whereas the patch method only
// replaces fields that are provided in the submitted dataset resource.

func (r *DatasetsService) Update(projectId string, datasetId string, dataset *Dataset) *DatasetsUpdateCall {
	return &DatasetsUpdateCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsUpdateCall) Fields(s ...googleapi.Field) *DatasetsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DatasetsUpdateCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.dataset,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates information in an existing dataset. The update method replaces the entire dataset resource, whereas the patch method only replaces fields that are provided in the submitted dataset resource.",
	//   "httpMethod": "PUT",
	//   "id": "bigquery.datasets.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the dataset being updated",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.jobs.get":

type JobsGetCall struct {
	s             *Service
	projectId     string
	jobId         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the specified job by ID.

func (r *JobsService) Get(projectId string, jobId string) *JobsGetCall {
	return &JobsGetCall{
		s:             r.s,
		projectId:     projectId,
		jobId:         jobId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/jobs/{jobId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsGetCall) Fields(s ...googleapi.Field) *JobsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsGetCall) Do() (*Job, error) {
	var returnValue *Job
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the specified job by ID.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.jobs.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Job ID of the requested job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the requested job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.jobs.getQueryResults":

type JobsGetQueryResultsCall struct {
	s             *Service
	projectId     string
	jobId         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// GetQueryResults: Retrieves the results of a query job.

func (r *JobsService) GetQueryResults(projectId string, jobId string) *JobsGetQueryResultsCall {
	return &JobsGetQueryResultsCall{
		s:             r.s,
		projectId:     projectId,
		jobId:         jobId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/queries/{jobId}",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to read
func (c *JobsGetQueryResultsCall) MaxResults(maxResults int64) *JobsGetQueryResultsCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *JobsGetQueryResultsCall) PageToken(pageToken string) *JobsGetQueryResultsCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// StartIndex sets the optional parameter "startIndex": Zero-based index
// of the starting row
func (c *JobsGetQueryResultsCall) StartIndex(startIndex uint64) *JobsGetQueryResultsCall {
	c.params_.Set("startIndex", fmt.Sprintf("%v", startIndex))
	return c
}

// TimeoutMs sets the optional parameter "timeoutMs": How long to wait
// for the query to complete, in milliseconds, before returning. Default
// is to return immediately. If the timeout passes before the job
// completes, the request will fail with a TIMEOUT error
func (c *JobsGetQueryResultsCall) TimeoutMs(timeoutMs int64) *JobsGetQueryResultsCall {
	c.params_.Set("timeoutMs", fmt.Sprintf("%v", timeoutMs))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsGetQueryResultsCall) Fields(s ...googleapi.Field) *JobsGetQueryResultsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsGetQueryResultsCall) Do() (*GetQueryResultsResponse, error) {
	var returnValue *GetQueryResultsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the results of a query job.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.jobs.getQueryResults",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Job ID of the query job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to read",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the query job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Zero-based index of the starting row",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "timeoutMs": {
	//       "description": "How long to wait for the query to complete, in milliseconds, before returning. Default is to return immediately. If the timeout passes before the job completes, the request will fail with a TIMEOUT error",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "projects/{projectId}/queries/{jobId}",
	//   "response": {
	//     "$ref": "GetQueryResultsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.jobs.insert":

type JobsInsertCall struct {
	s             *Service
	projectId     string
	job           *Job
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Starts a new asynchronous job.

func (r *JobsService) Insert(projectId string, job *Job) *JobsInsertCall {
	return &JobsInsertCall{
		s:             r.s,
		projectId:     projectId,
		job:           job,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/jobs",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *JobsInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *JobsInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/bigquery/v2/projects/{projectId}/jobs"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/bigquery/v2/projects/{projectId}/jobs"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *JobsInsertCall) Media(r io.Reader) *JobsInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/bigquery/v2/projects/{projectId}/jobs"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *JobsInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *JobsInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/bigquery/v2/projects/{projectId}/jobs"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *JobsInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *JobsInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsInsertCall) Fields(s ...googleapi.Field) *JobsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsInsertCall) Do() (*Job, error) {
	var returnValue *Job
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.job,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Starts a new asynchronous job.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.jobs.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/bigquery/v2/projects/{projectId}/jobs"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/bigquery/v2/projects/{projectId}/jobs"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID of the project that will be billed for the job",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/jobs",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "bigquery.jobs.list":

type JobsListCall struct {
	s             *Service
	projectId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all the Jobs in the specified project that were started
// by the user. The job list returns in reverse chronological order of
// when the jobs were created, starting with the most recent job
// created.

func (r *JobsService) List(projectId string) *JobsListCall {
	return &JobsListCall{
		s:             r.s,
		projectId:     projectId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/jobs",
	}
}

// AllUsers sets the optional parameter "allUsers": Whether to display
// jobs owned by all users in the project. Default false
func (c *JobsListCall) AllUsers(allUsers bool) *JobsListCall {
	c.params_.Set("allUsers", fmt.Sprintf("%v", allUsers))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *JobsListCall) MaxResults(maxResults int64) *JobsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *JobsListCall) PageToken(pageToken string) *JobsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields
func (c *JobsListCall) Projection(projection string) *JobsListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// StateFilter sets the optional parameter "stateFilter": Filter for job
// state
func (c *JobsListCall) StateFilter(stateFilter ...string) *JobsListCall {
	c.params_["stateFilter"] = stateFilter
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsListCall) Fields(s ...googleapi.Field) *JobsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsListCall) Do() (*JobList, error) {
	var returnValue *JobList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all the Jobs in the specified project that were started by the user. The job list returns in reverse chronological order of when the jobs were created, starting with the most recent job created.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.jobs.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "allUsers": {
	//       "description": "Whether to display jobs owned by all users in the project. Default false",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the jobs to list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields",
	//       "enum": [
	//         "full",
	//         "minimal"
	//       ],
	//       "enumDescriptions": [
	//         "Includes all job data",
	//         "Does not include the job configuration"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "stateFilter": {
	//       "description": "Filter for job state",
	//       "enum": [
	//         "done",
	//         "pending",
	//         "running"
	//       ],
	//       "enumDescriptions": [
	//         "Finished jobs",
	//         "Pending jobs",
	//         "Running jobs"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/jobs",
	//   "response": {
	//     "$ref": "JobList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.jobs.query":

type JobsQueryCall struct {
	s             *Service
	projectId     string
	queryrequest  *QueryRequest
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Query: Runs a BigQuery SQL query synchronously and returns query
// results if the query completes within a specified timeout.

func (r *JobsService) Query(projectId string, queryrequest *QueryRequest) *JobsQueryCall {
	return &JobsQueryCall{
		s:             r.s,
		projectId:     projectId,
		queryrequest:  queryrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/queries",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *JobsQueryCall) Fields(s ...googleapi.Field) *JobsQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *JobsQueryCall) Do() (*QueryResponse, error) {
	var returnValue *QueryResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.queryrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Runs a BigQuery SQL query synchronously and returns query results if the query completes within a specified timeout.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.jobs.query",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID of the project billed for the query",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/queries",
	//   "request": {
	//     "$ref": "QueryRequest"
	//   },
	//   "response": {
	//     "$ref": "QueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.projects.list":

type ProjectsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists the projects to which you have at least read access.

func (r *ProjectsService) List() *ProjectsListCall {
	return &ProjectsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *ProjectsListCall) MaxResults(maxResults int64) *ProjectsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *ProjectsListCall) PageToken(pageToken string) *ProjectsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsListCall) Fields(s ...googleapi.Field) *ProjectsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ProjectsListCall) Do() (*ProjectList, error) {
	var returnValue *ProjectList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists the projects to which you have at least read access.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.projects.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects",
	//   "response": {
	//     "$ref": "ProjectList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tabledata.insertAll":

type TabledataInsertAllCall struct {
	s                         *Service
	projectId                 string
	datasetId                 string
	tableId                   string
	tabledatainsertallrequest *TableDataInsertAllRequest
	caller_                   googleapi.Caller
	params_                   url.Values
	pathTemplate_             string
}

// InsertAll: Streams data into BigQuery one record at a time without
// needing to run a load job.

func (r *TabledataService) InsertAll(projectId string, datasetId string, tableId string, tabledatainsertallrequest *TableDataInsertAllRequest) *TabledataInsertAllCall {
	return &TabledataInsertAllCall{
		s:                         r.s,
		projectId:                 projectId,
		datasetId:                 datasetId,
		tableId:                   tableId,
		tabledatainsertallrequest: tabledatainsertallrequest,
		caller_:                   googleapi.JSONCall{},
		params_:                   make(map[string][]string),
		pathTemplate_:             "projects/{projectId}/datasets/{datasetId}/tables/{tableId}/insertAll",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TabledataInsertAllCall) Fields(s ...googleapi.Field) *TabledataInsertAllCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TabledataInsertAllCall) Do() (*TableDataInsertAllResponse, error) {
	var returnValue *TableDataInsertAllResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
		"tableId":   c.tableId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.tabledatainsertallrequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Streams data into BigQuery one record at a time without needing to run a load job.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.tabledata.insertAll",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the destination table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the destination table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the destination table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}/insertAll",
	//   "request": {
	//     "$ref": "TableDataInsertAllRequest"
	//   },
	//   "response": {
	//     "$ref": "TableDataInsertAllResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/bigquery.insertdata",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tabledata.list":

type TabledataListCall struct {
	s             *Service
	projectId     string
	datasetId     string
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves table data from a specified set of rows.

func (r *TabledataService) List(projectId string, datasetId string, tableId string) *TabledataListCall {
	return &TabledataListCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables/{tableId}/data",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *TabledataListCall) MaxResults(maxResults int64) *TabledataListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, identifying the result set
func (c *TabledataListCall) PageToken(pageToken string) *TabledataListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// StartIndex sets the optional parameter "startIndex": Zero-based index
// of the starting row to read
func (c *TabledataListCall) StartIndex(startIndex uint64) *TabledataListCall {
	c.params_.Set("startIndex", fmt.Sprintf("%v", startIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TabledataListCall) Fields(s ...googleapi.Field) *TabledataListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TabledataListCall) Do() (*TableDataList, error) {
	var returnValue *TableDataList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
		"tableId":   c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves table data from a specified set of rows.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.tabledata.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to read",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, identifying the result set",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to read",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Zero-based index of the starting row to read",
	//       "format": "uint64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to read",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}/data",
	//   "response": {
	//     "$ref": "TableDataList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tables.delete":

type TablesDeleteCall struct {
	s             *Service
	projectId     string
	datasetId     string
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the table specified by tableId from the dataset. If
// the table contains data, all the data will be deleted.

func (r *TablesService) Delete(projectId string, datasetId string, tableId string) *TablesDeleteCall {
	return &TablesDeleteCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	}
}

func (c *TablesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
		"tableId":   c.tableId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the table specified by tableId from the dataset. If the table contains data, all the data will be deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "bigquery.tables.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to delete",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tables.get":

type TablesGetCall struct {
	s             *Service
	projectId     string
	datasetId     string
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets the specified table resource by table ID. This method does
// not return the data in the table, it only returns the table resource,
// which describes the structure of this table.

func (r *TablesService) Get(projectId string, datasetId string, tableId string) *TablesGetCall {
	return &TablesGetCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesGetCall) Fields(s ...googleapi.Field) *TablesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TablesGetCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
		"tableId":   c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets the specified table resource by table ID. This method does not return the data in the table, it only returns the table resource, which describes the structure of this table.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.tables.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the requested table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the requested table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the requested table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tables.insert":

type TablesInsertCall struct {
	s             *Service
	projectId     string
	datasetId     string
	table         *Table
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new, empty table in the dataset.

func (r *TablesService) Insert(projectId string, datasetId string, table *Table) *TablesInsertCall {
	return &TablesInsertCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		table:         table,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesInsertCall) Fields(s ...googleapi.Field) *TablesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TablesInsertCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.table,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new, empty table in the dataset.",
	//   "httpMethod": "POST",
	//   "id": "bigquery.tables.insert",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the new table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the new table",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tables.list":

type TablesListCall struct {
	s             *Service
	projectId     string
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all tables in the specified dataset.

func (r *TablesService) List(projectId string, datasetId string) *TablesListCall {
	return &TablesListCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return
func (c *TablesListCall) MaxResults(maxResults int64) *TablesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Page token,
// returned by a previous call, to request the next page of results
func (c *TablesListCall) PageToken(pageToken string) *TablesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesListCall) Fields(s ...googleapi.Field) *TablesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TablesListCall) Do() (*TableList, error) {
	var returnValue *TableList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all tables in the specified dataset.",
	//   "httpMethod": "GET",
	//   "id": "bigquery.tables.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the tables to list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page token, returned by a previous call, to request the next page of results",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the tables to list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables",
	//   "response": {
	//     "$ref": "TableList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tables.patch":

type TablesPatchCall struct {
	s             *Service
	projectId     string
	datasetId     string
	tableId       string
	table         *Table
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates information in an existing table. The update method
// replaces the entire table resource, whereas the patch method only
// replaces fields that are provided in the submitted table resource.
// This method supports patch semantics.

func (r *TablesService) Patch(projectId string, datasetId string, tableId string, table *Table) *TablesPatchCall {
	return &TablesPatchCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		tableId:       tableId,
		table:         table,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesPatchCall) Fields(s ...googleapi.Field) *TablesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TablesPatchCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
		"tableId":   c.tableId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.table,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates information in an existing table. The update method replaces the entire table resource, whereas the patch method only replaces fields that are provided in the submitted table resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "bigquery.tables.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigquery.tables.update":

type TablesUpdateCall struct {
	s             *Service
	projectId     string
	datasetId     string
	tableId       string
	table         *Table
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates information in an existing table. The update method
// replaces the entire table resource, whereas the patch method only
// replaces fields that are provided in the submitted table resource.

func (r *TablesService) Update(projectId string, datasetId string, tableId string, table *Table) *TablesUpdateCall {
	return &TablesUpdateCall{
		s:             r.s,
		projectId:     projectId,
		datasetId:     datasetId,
		tableId:       tableId,
		table:         table,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesUpdateCall) Fields(s ...googleapi.Field) *TablesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TablesUpdateCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"projectId": c.projectId,
		"datasetId": c.datasetId,
		"tableId":   c.tableId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.table,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates information in an existing table. The update method replaces the entire table resource, whereas the patch method only replaces fields that are provided in the submitted table resource.",
	//   "httpMethod": "PUT",
	//   "id": "bigquery.tables.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "datasetId",
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "Dataset ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table ID of the table to update",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/datasets/{datasetId}/tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
