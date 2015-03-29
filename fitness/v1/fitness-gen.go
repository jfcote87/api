// Package fitness provides access to the Fitness.
//
// See https://developers.google.com/fit/rest/
//
// Usage example:
//
//   import "github.com/jfcote87/api/fitness/v1"
//   ...
//   fitnessService, err := fitness.New(oauthHttpClient)
package fitness

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

const apiId = "fitness:v1"
const apiName = "fitness"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/fitness/v1/users/"}

// OAuth2 scopes used by this API.
const (
	// View your activity information in Google Fit
	FitnessActivityReadScope = "https://www.googleapis.com/auth/fitness.activity.read"

	// View and store your activity information in Google Fit
	FitnessActivityWriteScope = "https://www.googleapis.com/auth/fitness.activity.write"

	// View body sensor information in Google Fit
	FitnessBodyReadScope = "https://www.googleapis.com/auth/fitness.body.read"

	// View and store body sensor data in Google Fit
	FitnessBodyWriteScope = "https://www.googleapis.com/auth/fitness.body.write"

	// View your stored location data in Google Fit
	FitnessLocationReadScope = "https://www.googleapis.com/auth/fitness.location.read"

	// View and store your location data in Google Fit
	FitnessLocationWriteScope = "https://www.googleapis.com/auth/fitness.location.write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Users *UsersService
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	rs.DataSources = NewUsersDataSourcesService(s)
	rs.Sessions = NewUsersSessionsService(s)
	return rs
}

type UsersService struct {
	s *Service

	DataSources *UsersDataSourcesService

	Sessions *UsersSessionsService
}

func NewUsersDataSourcesService(s *Service) *UsersDataSourcesService {
	rs := &UsersDataSourcesService{s: s}
	rs.Datasets = NewUsersDataSourcesDatasetsService(s)
	return rs
}

type UsersDataSourcesService struct {
	s *Service

	Datasets *UsersDataSourcesDatasetsService
}

func NewUsersDataSourcesDatasetsService(s *Service) *UsersDataSourcesDatasetsService {
	rs := &UsersDataSourcesDatasetsService{s: s}
	return rs
}

type UsersDataSourcesDatasetsService struct {
	s *Service
}

func NewUsersSessionsService(s *Service) *UsersSessionsService {
	rs := &UsersSessionsService{s: s}
	return rs
}

type UsersSessionsService struct {
	s *Service
}

type Application struct {
	// DetailsUrl: An optional URI that can be used to link back to the
	// application.
	DetailsUrl string `json:"detailsUrl,omitempty"`

	// Name: The name of this application. This is required for REST
	// clients, but we do not enforce uniqueness of this name. It is
	// provided as a matter of convenience for other developers who would
	// like to identify which REST created an Application or Data Source.
	Name string `json:"name,omitempty"`

	// PackageName: Package name for this application. This is used as a
	// unique identifier when created by Android applications, but cannot be
	// specified by REST clients. REST clients will have their developer
	// project number reflected into the Data Source data stream IDs,
	// instead of the packageName.
	PackageName string `json:"packageName,omitempty"`

	// Version: Version of the application. You should update this field
	// whenever the application changes in a way that affects the
	// computation of the data.
	Version string `json:"version,omitempty"`
}

type DataPoint struct {
	// ComputationTimeMillis: Used for version checking during
	// transformation; that is, a datapoint can only replace another
	// datapoint that has an older computation time stamp.
	ComputationTimeMillis int64 `json:"computationTimeMillis,omitempty,string"`

	// DataTypeName: The data type defining the format of the values in this
	// data point.
	DataTypeName string `json:"dataTypeName,omitempty"`

	// EndTimeNanos: The end time of the interval represented by this data
	// point, in nanoseconds since epoch.
	EndTimeNanos int64 `json:"endTimeNanos,omitempty,string"`

	// ModifiedTimeMillis: Indicates the last time this data point was
	// modified. Useful only in contexts where we are listing the data
	// changes, rather than representing the current state of the data.
	ModifiedTimeMillis int64 `json:"modifiedTimeMillis,omitempty,string"`

	// OriginDataSourceId: If the data point is contained in a dataset for a
	// derived data source, this field will be populated with the data
	// source stream ID that created the data point originally.
	OriginDataSourceId string `json:"originDataSourceId,omitempty"`

	// RawTimestampNanos: The raw timestamp from the original SensorEvent.
	RawTimestampNanos int64 `json:"rawTimestampNanos,omitempty,string"`

	// StartTimeNanos: The start time of the interval represented by this
	// data point, in nanoseconds since epoch.
	StartTimeNanos int64 `json:"startTimeNanos,omitempty,string"`

	// Value: Values of each data type field for the data point. It is
	// expected that each value corresponding to a data type field will
	// occur in the same order that the field is listed with in the data
	// type specified in a data source.
	//
	// Only one of integer and floating
	// point fields will be populated, depending on the format enum value
	// within data source's type field.
	Value []*Value `json:"value,omitempty"`
}

type DataSource struct {
	// Application: Information about an application which feeds sensor data
	// into the platform.
	Application *Application `json:"application,omitempty"`

	// DataStreamId: A unique identifier for the data stream produced by
	// this data source. The identifier includes:
	//
	//
	// - The physical device's
	// manufacturer, model, and serial number (UID).
	// - The application's
	// package name or name. Package name is used when the data source was
	// created by an Android application. The developer project number is
	// used when the data source was created by a REST client.
	// - The data
	// source's type.
	// - The data source's stream name.  Note that not all
	// attributes of the data source are used as part of the stream
	// identifier. In particular, the version of the hardware/the
	// application isn't used. This allows us to preserve the same stream
	// through version updates. This also means that two DataSource objects
	// may represent the same data stream even if they're not equal.
	//
	// The
	// exact format of the data stream ID created by an Android application
	// is:
	// type:dataType.name:application.packageName:device.manufacturer:device.
	// model:device.uid:dataStreamName
	//
	// The exact format of the data stream
	// ID created by a REST client is: type:dataType.name:developer project
	// number:device.manufacturer:device.model:device.uid:dataStreamName
	//
	//
	// When any of the optional fields that comprise of the data stream ID
	// are blank, they will be omitted from the data stream ID. The minnimum
	// viable data stream ID would be: type:dataType.name:developer project
	// number
	//
	// Finally, the developer project number is obfuscated when read
	// by any REST or Android client that did not create the data source.
	// Only the data source creator will see the developer project number in
	// clear and normal form.
	DataStreamId string `json:"dataStreamId,omitempty"`

	// DataStreamName: The stream name uniquely identifies this particular
	// data source among other data sources of the same type from the same
	// underlying producer. Setting the stream name is optional, but should
	// be done whenever an application exposes two streams for the same data
	// type, or when a device has two equivalent sensors.
	DataStreamName string `json:"dataStreamName,omitempty"`

	// DataType: The data type defines the schema for a stream of data being
	// collected by, inserted into, or queried from the Fitness API.
	DataType *DataType `json:"dataType,omitempty"`

	// Device: Representation of an integrated device (such as a phone or a
	// wearable) that can hold sensors.
	Device *Device `json:"device,omitempty"`

	// Name: An end-user visible name for this data source.
	Name string `json:"name,omitempty"`

	// Type: A constant describing the type of this data source. Indicates
	// whether this data source produces raw or derived data.
	Type string `json:"type,omitempty"`
}

type DataType struct {
	// Field: A field represents one dimension of a data type.
	Field []*DataTypeField `json:"field,omitempty"`

	// Name: Each data type has a unique, namespaced, name. All data types
	// in the com.google namespace are shared as part of the platform.
	Name string `json:"name,omitempty"`
}

type DataTypeField struct {
	// Format: The different supported formats for each field in a data
	// type.
	Format string `json:"format,omitempty"`

	// Name: Defines the name and format of data. Unlike data type names,
	// field names are not namespaced, and only need to be unique within the
	// data type.
	Name string `json:"name,omitempty"`

	Optional bool `json:"optional,omitempty"`
}

type Dataset struct {
	// DataSourceId: The data stream ID of the data source that created the
	// points in this dataset.
	DataSourceId string `json:"dataSourceId,omitempty"`

	// MaxEndTimeNs: The largest end time of all data points in this
	// possibly partial representation of the dataset. Time is in
	// nanoseconds from epoch. This should also match the first part of the
	// dataset identifier.
	MaxEndTimeNs int64 `json:"maxEndTimeNs,omitempty,string"`

	// MinStartTimeNs: The smallest start time of all data points in this
	// possibly partial representation of the dataset. Time is in
	// nanoseconds from epoch. This should also match the first part of the
	// dataset identifier.
	MinStartTimeNs int64 `json:"minStartTimeNs,omitempty,string"`

	// NextPageToken: This token will be set when a dataset is received in
	// response to a GET request and the dataset is too large to be included
	// in a single response. Provide this value in a subsequent GET request
	// to return the next page of data points within this dataset.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Point: A partial list of data points contained in the dataset,
	// ordered by largest endTimeNanos first. This list is considered
	// complete when retrieving a small dataset and partial when patching a
	// dataset or retrieving a dataset that is too large to include in a
	// single response.
	Point []*DataPoint `json:"point,omitempty"`
}

type Device struct {
	// Manufacturer: Manufacturer of the product/hardware.
	Manufacturer string `json:"manufacturer,omitempty"`

	// Model: End-user visible model name for the device.
	Model string `json:"model,omitempty"`

	// Type: A constant representing the type of the device.
	Type string `json:"type,omitempty"`

	// Uid: The serial number or other unique ID for the hardware. This
	// field is obfuscated when read by any REST or Android client that did
	// not create the data source. Only the data source creator will see the
	// uid field in clear and normal form.
	Uid string `json:"uid,omitempty"`

	// Version: Version string for the device hardware/software.
	Version string `json:"version,omitempty"`
}

type ListDataSourcesResponse struct {
	// DataSource: A previously created data source.
	DataSource []*DataSource `json:"dataSource,omitempty"`
}

type ListSessionsResponse struct {
	// DeletedSession: If includeDeleted is set to true in the request, this
	// list will contain sessions deleted with original end times that are
	// within the startTime and endTime frame.
	DeletedSession []*Session `json:"deletedSession,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Session: Sessions with an end time that is between startTime and
	// endTime of the request.
	Session []*Session `json:"session,omitempty"`
}

type Session struct {
	// ActivityType: The type of activity this session represents.
	ActivityType int64 `json:"activityType,omitempty"`

	// Application: The application that created the session.
	Application *Application `json:"application,omitempty"`

	// Description: A description for this session.
	Description string `json:"description,omitempty"`

	// EndTimeMillis: An end time, in milliseconds since epoch, inclusive.
	EndTimeMillis int64 `json:"endTimeMillis,omitempty,string"`

	// Id: A client-generated identifier that is unique across all sessions
	// owned by this particular user.
	Id string `json:"id,omitempty"`

	// ModifiedTimeMillis: A timestamp that indicates when the session was
	// last modified.
	ModifiedTimeMillis int64 `json:"modifiedTimeMillis,omitempty,string"`

	// Name: A human readable name of the session.
	Name string `json:"name,omitempty"`

	// StartTimeMillis: A start time, in milliseconds since epoch,
	// inclusive.
	StartTimeMillis int64 `json:"startTimeMillis,omitempty,string"`
}

type Value struct {
	// FpVal: Floating point value. When this is set, intVal must not be
	// set.
	FpVal float64 `json:"fpVal,omitempty"`

	// IntVal: Integer value. When this is set, fpVal must not be set.
	IntVal int64 `json:"intVal,omitempty"`
}

// method id "fitness.users.dataSources.create":

type UsersDataSourcesCreateCall struct {
	s             *Service
	userId        string
	datasource    *DataSource
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a new data source that is unique across all data
// sources belonging to this user. The data stream ID field can be
// omitted and will be generated by the server with the correct format.
// The data stream ID is an ordered combination of some fields from the
// data source. In addition to the data source fields reflected into the
// data source ID, the developer project number that is authenticated
// when creating the data source is included. This developer project
// number is obfuscated when read by any other developer reading public
// data types.

func (r *UsersDataSourcesService) Create(userId string, datasource *DataSource) *UsersDataSourcesCreateCall {
	return &UsersDataSourcesCreateCall{
		s:             r.s,
		userId:        userId,
		datasource:    datasource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesCreateCall) Fields(s ...googleapi.Field) *UsersDataSourcesCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesCreateCall) Do() (*DataSource, error) {
	var returnValue *DataSource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId": c.userId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.datasource,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new data source that is unique across all data sources belonging to this user. The data stream ID field can be omitted and will be generated by the server with the correct format. The data stream ID is an ordered combination of some fields from the data source. In addition to the data source fields reflected into the data source ID, the developer project number that is authenticated when creating the data source is included. This developer project number is obfuscated when read by any other developer reading public data types.",
	//   "httpMethod": "POST",
	//   "id": "fitness.users.dataSources.create",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "Create the data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.get":

type UsersDataSourcesGetCall struct {
	s             *Service
	userId        string
	dataSourceId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns a data source identified by a data stream ID.

func (r *UsersDataSourcesService) Get(userId string, dataSourceId string) *UsersDataSourcesGetCall {
	return &UsersDataSourcesGetCall{
		s:             r.s,
		userId:        userId,
		dataSourceId:  dataSourceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources/{dataSourceId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesGetCall) Fields(s ...googleapi.Field) *UsersDataSourcesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesGetCall) Do() (*DataSource, error) {
	var returnValue *DataSource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a data source identified by a data stream ID.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.dataSources.get",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Retrieve a data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.list":

type UsersDataSourcesListCall struct {
	s             *Service
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all data sources that are visible to the developer, using
// the OAuth scopes provided. The list is not exhaustive: the user may
// have private data sources that are only visible to other developers
// or calls using other scopes.

func (r *UsersDataSourcesService) List(userId string) *UsersDataSourcesListCall {
	return &UsersDataSourcesListCall{
		s:             r.s,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources",
	}
}

// DataTypeName sets the optional parameter "dataTypeName": The names of
// data types to include in the list. If not specified, all data sources
// will be returned.
func (c *UsersDataSourcesListCall) DataTypeName(dataTypeName ...string) *UsersDataSourcesListCall {
	c.params_["dataTypeName"] = dataTypeName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesListCall) Fields(s ...googleapi.Field) *UsersDataSourcesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesListCall) Do() (*ListDataSourcesResponse, error) {
	var returnValue *ListDataSourcesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId": c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all data sources that are visible to the developer, using the OAuth scopes provided. The list is not exhaustive: the user may have private data sources that are only visible to other developers or calls using other scopes.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.dataSources.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "dataTypeName": {
	//       "description": "The names of data types to include in the list. If not specified, all data sources will be returned.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "List data sources for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources",
	//   "response": {
	//     "$ref": "ListDataSourcesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.patch":

type UsersDataSourcesPatchCall struct {
	s             *Service
	userId        string
	dataSourceId  string
	datasource    *DataSource
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a given data source. It is an error to modify the data
// source's data stream ID, data type, type, stream name or device
// information apart from the device version. Changing these fields
// would require a new unique data stream ID and separate data
// source.
//
// Data sources are identified by their data stream ID. This
// method supports patch semantics.

func (r *UsersDataSourcesService) Patch(userId string, dataSourceId string, datasource *DataSource) *UsersDataSourcesPatchCall {
	return &UsersDataSourcesPatchCall{
		s:             r.s,
		userId:        userId,
		dataSourceId:  dataSourceId,
		datasource:    datasource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources/{dataSourceId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesPatchCall) Fields(s ...googleapi.Field) *UsersDataSourcesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesPatchCall) Do() (*DataSource, error) {
	var returnValue *DataSource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.datasource,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a given data source. It is an error to modify the data source's data stream ID, data type, type, stream name or device information apart from the device version. Changing these fields would require a new unique data stream ID and separate data source.\n\nData sources are identified by their data stream ID. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fitness.users.dataSources.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Update the data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.update":

type UsersDataSourcesUpdateCall struct {
	s             *Service
	userId        string
	dataSourceId  string
	datasource    *DataSource
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a given data source. It is an error to modify the
// data source's data stream ID, data type, type, stream name or device
// information apart from the device version. Changing these fields
// would require a new unique data stream ID and separate data
// source.
//
// Data sources are identified by their data stream ID.

func (r *UsersDataSourcesService) Update(userId string, dataSourceId string, datasource *DataSource) *UsersDataSourcesUpdateCall {
	return &UsersDataSourcesUpdateCall{
		s:             r.s,
		userId:        userId,
		dataSourceId:  dataSourceId,
		datasource:    datasource,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources/{dataSourceId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesUpdateCall) Fields(s ...googleapi.Field) *UsersDataSourcesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesUpdateCall) Do() (*DataSource, error) {
	var returnValue *DataSource
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.datasource,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a given data source. It is an error to modify the data source's data stream ID, data type, type, stream name or device information apart from the device version. Changing these fields would require a new unique data stream ID and separate data source.\n\nData sources are identified by their data stream ID.",
	//   "httpMethod": "PUT",
	//   "id": "fitness.users.dataSources.update",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Update the data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.datasets.delete":

type UsersDataSourcesDatasetsDeleteCall struct {
	s             *Service
	userId        string
	dataSourceId  string
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Performs an inclusive delete of all data points whose start
// and end times have any overlap with the time range specified by the
// dataset ID. For most data types, the entire data point will be
// deleted. For data types where the time span represents a consistent
// value (such as com.google.activity.segment), and a data point
// straddles either end point of the dataset, only the overlapping
// portion of the data point will be deleted.

func (r *UsersDataSourcesDatasetsService) Delete(userId string, dataSourceId string, datasetId string) *UsersDataSourcesDatasetsDeleteCall {
	return &UsersDataSourcesDatasetsDeleteCall{
		s:             r.s,
		userId:        userId,
		dataSourceId:  dataSourceId,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	}
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch.
func (c *UsersDataSourcesDatasetsDeleteCall) CurrentTimeMillis(currentTimeMillis int64) *UsersDataSourcesDatasetsDeleteCall {
	c.params_.Set("currentTimeMillis", fmt.Sprintf("%v", currentTimeMillis))
	return c
}

// ModifiedTimeMillis sets the optional parameter "modifiedTimeMillis":
// When the operation was performed on the client.
func (c *UsersDataSourcesDatasetsDeleteCall) ModifiedTimeMillis(modifiedTimeMillis int64) *UsersDataSourcesDatasetsDeleteCall {
	c.params_.Set("modifiedTimeMillis", fmt.Sprintf("%v", modifiedTimeMillis))
	return c
}

func (c *UsersDataSourcesDatasetsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
		"datasetId":    c.datasetId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Performs an inclusive delete of all data points whose start and end times have any overlap with the time range specified by the dataset ID. For most data types, the entire data point will be deleted. For data types where the time span represents a consistent value (such as com.google.activity.segment), and a data point straddles either end point of the dataset, only the overlapping portion of the data point will be deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "fitness.users.dataSources.datasets.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source that created the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "Dataset identifier that is a composite of the minimum data point start time and maximum data point end time represented as nanoseconds from the epoch. The ID is formatted like: \"startTime-endTime\" where startTime and endTime are 64 bit integers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modifiedTimeMillis": {
	//       "description": "When the operation was performed on the client.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Delete a dataset for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.datasets.get":

type UsersDataSourcesDatasetsGetCall struct {
	s             *Service
	userId        string
	dataSourceId  string
	datasetId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Returns a dataset containing all data points whose start and end
// times overlap with the specified range of the dataset minimum start
// time and maximum end time. Specifically, any data point whose start
// time is less than or equal to the dataset end time and whose end time
// is greater than or equal to the dataset start time.

func (r *UsersDataSourcesDatasetsService) Get(userId string, dataSourceId string, datasetId string) *UsersDataSourcesDatasetsGetCall {
	return &UsersDataSourcesDatasetsGetCall{
		s:             r.s,
		userId:        userId,
		dataSourceId:  dataSourceId,
		datasetId:     datasetId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	}
}

// Limit sets the optional parameter "limit": If specified, no more than
// this many data points will be included in the dataset. If the there
// are more data points in the dataset, nextPageToken will be set in the
// dataset response.
func (c *UsersDataSourcesDatasetsGetCall) Limit(limit int64) *UsersDataSourcesDatasetsGetCall {
	c.params_.Set("limit", fmt.Sprintf("%v", limit))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large datasets. To get the next
// page of a dataset, set this parameter to the value of nextPageToken
// from the previous response. Each subsequent call will yield a partial
// dataset with data point end timestamps that are strictly smaller than
// those in the previous partial response.
func (c *UsersDataSourcesDatasetsGetCall) PageToken(pageToken string) *UsersDataSourcesDatasetsGetCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesDatasetsGetCall) Fields(s ...googleapi.Field) *UsersDataSourcesDatasetsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesDatasetsGetCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
		"datasetId":    c.datasetId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a dataset containing all data points whose start and end times overlap with the specified range of the dataset minimum start time and maximum end time. Specifically, any data point whose start time is less than or equal to the dataset end time and whose end time is greater than or equal to the dataset start time.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.dataSources.datasets.get",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source that created the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "Dataset identifier that is a composite of the minimum data point start time and maximum data point end time represented as nanoseconds from the epoch. The ID is formatted like: \"startTime-endTime\" where startTime and endTime are 64 bit integers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "limit": {
	//       "description": "If specified, no more than this many data points will be included in the dataset. If the there are more data points in the dataset, nextPageToken will be set in the dataset response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large datasets. To get the next page of a dataset, set this parameter to the value of nextPageToken from the previous response. Each subsequent call will yield a partial dataset with data point end timestamps that are strictly smaller than those in the previous partial response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Retrieve a dataset for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.datasets.patch":

type UsersDataSourcesDatasetsPatchCall struct {
	s             *Service
	userId        string
	dataSourceId  string
	datasetId     string
	dataset       *Dataset
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Adds data points to a dataset. The dataset need not be
// previously created. All points within the given dataset will be
// returned with subsquent calls to retrieve this dataset. Data points
// can belong to more than one dataset. This method does not use patch
// semantics.

func (r *UsersDataSourcesDatasetsService) Patch(userId string, dataSourceId string, datasetId string, dataset *Dataset) *UsersDataSourcesDatasetsPatchCall {
	return &UsersDataSourcesDatasetsPatchCall{
		s:             r.s,
		userId:        userId,
		dataSourceId:  dataSourceId,
		datasetId:     datasetId,
		dataset:       dataset,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	}
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch. Note that the
// minStartTimeNs and maxEndTimeNs properties in the request body are in
// nanoseconds instead of milliseconds.
func (c *UsersDataSourcesDatasetsPatchCall) CurrentTimeMillis(currentTimeMillis int64) *UsersDataSourcesDatasetsPatchCall {
	c.params_.Set("currentTimeMillis", fmt.Sprintf("%v", currentTimeMillis))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesDatasetsPatchCall) Fields(s ...googleapi.Field) *UsersDataSourcesDatasetsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersDataSourcesDatasetsPatchCall) Do() (*Dataset, error) {
	var returnValue *Dataset
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
		"datasetId":    c.datasetId,
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
	//   "description": "Adds data points to a dataset. The dataset need not be previously created. All points within the given dataset will be returned with subsquent calls to retrieve this dataset. Data points can belong to more than one dataset. This method does not use patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fitness.users.dataSources.datasets.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch. Note that the minStartTimeNs and maxEndTimeNs properties in the request body are in nanoseconds instead of milliseconds.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source that created the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "Dataset identifier that is a composite of the minimum data point start time and maximum data point end time represented as nanoseconds from the epoch. The ID is formatted like: \"startTime-endTime\" where startTime and endTime are 64 bit integers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Patch a dataset for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.sessions.delete":

type UsersSessionsDeleteCall struct {
	s             *Service
	userId        string
	sessionId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a session specified by the given session ID.

func (r *UsersSessionsService) Delete(userId string, sessionId string) *UsersSessionsDeleteCall {
	return &UsersSessionsDeleteCall{
		s:             r.s,
		userId:        userId,
		sessionId:     sessionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/sessions/{sessionId}",
	}
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch.
func (c *UsersSessionsDeleteCall) CurrentTimeMillis(currentTimeMillis int64) *UsersSessionsDeleteCall {
	c.params_.Set("currentTimeMillis", fmt.Sprintf("%v", currentTimeMillis))
	return c
}

func (c *UsersSessionsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":    c.userId,
		"sessionId": c.sessionId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a session specified by the given session ID.",
	//   "httpMethod": "DELETE",
	//   "id": "fitness.users.sessions.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "sessionId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sessionId": {
	//       "description": "The ID of the session to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Delete a session for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/sessions/{sessionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write"
	//   ]
	// }

}

// method id "fitness.users.sessions.list":

type UsersSessionsListCall struct {
	s             *Service
	userId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists sessions previously created.

func (r *UsersSessionsService) List(userId string) *UsersSessionsListCall {
	return &UsersSessionsListCall{
		s:             r.s,
		userId:        userId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/sessions",
	}
}

// EndTime sets the optional parameter "endTime": An RFC3339 timestamp.
// Only sessions ending between the start and end times will be included
// in the response.
func (c *UsersSessionsListCall) EndTime(endTime string) *UsersSessionsListCall {
	c.params_.Set("endTime", fmt.Sprintf("%v", endTime))
	return c
}

// IncludeDeleted sets the optional parameter "includeDeleted": If true,
// deleted sessions will be returned. When set to true, sessions
// returned in this response will only have an ID and will not have any
// other fields.
func (c *UsersSessionsListCall) IncludeDeleted(includeDeleted bool) *UsersSessionsListCall {
	c.params_.Set("includeDeleted", fmt.Sprintf("%v", includeDeleted))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// nextPageToken from the previous response.
func (c *UsersSessionsListCall) PageToken(pageToken string) *UsersSessionsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// StartTime sets the optional parameter "startTime": An RFC3339
// timestamp. Only sessions ending between the start and end times will
// be included in the response.
func (c *UsersSessionsListCall) StartTime(startTime string) *UsersSessionsListCall {
	c.params_.Set("startTime", fmt.Sprintf("%v", startTime))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersSessionsListCall) Fields(s ...googleapi.Field) *UsersSessionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersSessionsListCall) Do() (*ListSessionsResponse, error) {
	var returnValue *ListSessionsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId": c.userId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists sessions previously created.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.sessions.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "description": "An RFC3339 timestamp. Only sessions ending between the start and end times will be included in the response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "includeDeleted": {
	//       "description": "If true, deleted sessions will be returned. When set to true, sessions returned in this response will only have an ID and will not have any other fields.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "An RFC3339 timestamp. Only sessions ending between the start and end times will be included in the response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "List sessions for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/sessions",
	//   "response": {
	//     "$ref": "ListSessionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.sessions.update":

type UsersSessionsUpdateCall struct {
	s             *Service
	userId        string
	sessionId     string
	session       *Session
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates or insert a given session.

func (r *UsersSessionsService) Update(userId string, sessionId string, session *Session) *UsersSessionsUpdateCall {
	return &UsersSessionsUpdateCall{
		s:             r.s,
		userId:        userId,
		sessionId:     sessionId,
		session:       session,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "{userId}/sessions/{sessionId}",
	}
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch.
func (c *UsersSessionsUpdateCall) CurrentTimeMillis(currentTimeMillis int64) *UsersSessionsUpdateCall {
	c.params_.Set("currentTimeMillis", fmt.Sprintf("%v", currentTimeMillis))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersSessionsUpdateCall) Fields(s ...googleapi.Field) *UsersSessionsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersSessionsUpdateCall) Do() (*Session, error) {
	var returnValue *Session
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userId":    c.userId,
		"sessionId": c.sessionId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.session,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates or insert a given session.",
	//   "httpMethod": "PUT",
	//   "id": "fitness.users.sessions.update",
	//   "parameterOrder": [
	//     "userId",
	//     "sessionId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sessionId": {
	//       "description": "The ID of the session to be created.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Create sessions for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/sessions/{sessionId}",
	//   "request": {
	//     "$ref": "Session"
	//   },
	//   "response": {
	//     "$ref": "Session"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write"
	//   ]
	// }

}
