// Package fusiontables provides access to the Fusion Tables API.
//
// See https://developers.google.com/fusiontables
//
// Usage example:
//
//   import "github.com/jfcote87/api/fusiontables/v2"
//   ...
//   fusiontablesService, err := fusiontables.New(oauthHttpClient)
package fusiontables

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

const apiId = "fusiontables:v2"
const apiName = "fusiontables"
const apiVersion = "v2"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/fusiontables/v2/"}

// OAuth2 scopes used by this API.
const (
	// Manage your Fusion Tables
	FusiontablesScope = "https://www.googleapis.com/auth/fusiontables"

	// View your Fusion Tables
	FusiontablesReadonlyScope = "https://www.googleapis.com/auth/fusiontables.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Column = NewColumnService(s)
	s.Query = NewQueryService(s)
	s.Style = NewStyleService(s)
	s.Table = NewTableService(s)
	s.Task = NewTaskService(s)
	s.Template = NewTemplateService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Column *ColumnService

	Query *QueryService

	Style *StyleService

	Table *TableService

	Task *TaskService

	Template *TemplateService
}

func NewColumnService(s *Service) *ColumnService {
	rs := &ColumnService{s: s}
	return rs
}

type ColumnService struct {
	s *Service
}

func NewQueryService(s *Service) *QueryService {
	rs := &QueryService{s: s}
	return rs
}

type QueryService struct {
	s *Service
}

func NewStyleService(s *Service) *StyleService {
	rs := &StyleService{s: s}
	return rs
}

type StyleService struct {
	s *Service
}

func NewTableService(s *Service) *TableService {
	rs := &TableService{s: s}
	return rs
}

type TableService struct {
	s *Service
}

func NewTaskService(s *Service) *TaskService {
	rs := &TaskService{s: s}
	return rs
}

type TaskService struct {
	s *Service
}

func NewTemplateService(s *Service) *TemplateService {
	rs := &TemplateService{s: s}
	return rs
}

type TemplateService struct {
	s *Service
}

type Bucket struct {
	// Color: Color of line or the interior of a polygon in #RRGGBB format.
	Color string `json:"color,omitempty"`

	// Icon: Icon name used for a point.
	Icon string `json:"icon,omitempty"`

	// Max: Maximum value in the selected column for a row to be styled
	// according to the bucket color, opacity, icon, or weight.
	Max float64 `json:"max,omitempty"`

	// Min: Minimum value in the selected column for a row to be styled
	// according to the bucket color, opacity, icon, or weight.
	Min float64 `json:"min,omitempty"`

	// Opacity: Opacity of the color: 0.0 (transparent) to 1.0 (opaque).
	Opacity float64 `json:"opacity,omitempty"`

	// Weight: Width of a line (in pixels).
	Weight int64 `json:"weight,omitempty"`
}

type Column struct {
	// BaseColumn: Identifier of the base column. If present, this column is
	// derived from the specified base column.
	BaseColumn *ColumnBaseColumn `json:"baseColumn,omitempty"`

	// ColumnId: Identifier for the column.
	ColumnId int64 `json:"columnId,omitempty"`

	// ColumnJsonSchema: JSON schema for interpreting JSON in this column.
	ColumnJsonSchema string `json:"columnJsonSchema,omitempty"`

	// ColumnPropertiesJson: JSON object containing custom column
	// properties.
	ColumnPropertiesJson string `json:"columnPropertiesJson,omitempty"`

	// Description: Column description.
	Description string `json:"description,omitempty"`

	// FormatPattern: Format pattern.
	// Acceptable values are
	// DT_DATE_MEDIUMe.g Dec 24, 2008 DT_DATE_SHORTfor example 12/24/08
	// DT_DATE_TIME_MEDIUMfor example Dec 24, 2008 8:30:45 PM
	// DT_DATE_TIME_SHORTfor example 12/24/08 8:30 PM
	// DT_DAY_MONTH_2_DIGIT_YEARfor example 24/12/08
	// DT_DAY_MONTH_2_DIGIT_YEAR_TIMEfor example 24/12/08 20:30
	// DT_DAY_MONTH_2_DIGIT_YEAR_TIME_MERIDIANfor example 24/12/08 8:30 PM
	// DT_DAY_MONTH_4_DIGIT_YEARfor example 24/12/2008
	// DT_DAY_MONTH_4_DIGIT_YEAR_TIMEfor example 24/12/2008 20:30
	// DT_DAY_MONTH_4_DIGIT_YEAR_TIME_MERIDIANfor example 24/12/2008 8:30 PM
	// DT_ISO_YEAR_MONTH_DAYfor example 2008-12-24
	// DT_ISO_YEAR_MONTH_DAY_TIMEfor example 2008-12-24 20:30:45
	// DT_MONTH_DAY_4_DIGIT_YEARfor example 12/24/2008 DT_TIME_LONGfor
	// example 8:30:45 PM UTC-6 DT_TIME_MEDIUMfor example 8:30:45 PM
	// DT_TIME_SHORTfor example 8:30 PM DT_YEAR_ONLYfor example 2008
	// HIGHLIGHT_UNTYPED_CELLSHighlight cell data that does not match the
	// data type NONENo formatting (default) NUMBER_CURRENCYfor example
	// $1234.56 NUMBER_DEFAULTfor example 1,234.56 NUMBER_INTEGERfor example
	// 1235 NUMBER_NO_SEPARATORfor example 1234.56 NUMBER_PERCENTfor example
	// 123,456% NUMBER_SCIENTIFICfor example 1E3
	// STRING_EIGHT_LINE_IMAGEDisplays thumbnail images as tall as eight
	// lines of text STRING_FOUR_LINE_IMAGEDisplays thumbnail images as tall
	// as four lines of text STRING_JSON_TEXTAllows JSON editing of text in
	// UI STRING_LINKTreats cell as a link (must start with http:// or
	// https://) STRING_ONE_LINE_IMAGEDisplays thumbnail images as tall as
	// one line of text STRING_VIDEO_OR_MAPDisplay a video or map thumbnail
	FormatPattern string `json:"formatPattern,omitempty"`

	// GraphPredicate: Column graph predicate.
	// Used to map table to graph
	// data model (subject,predicate,object)
	// See W3C Graph-based Data Model.
	GraphPredicate string `json:"graphPredicate,omitempty"`

	// Kind: The kind of item this is. For a column, this is always
	// fusiontables#column.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the column.
	Name string `json:"name,omitempty"`

	// Type: Type of the column.
	Type string `json:"type,omitempty"`

	// ValidValues: List of valid values used to validate data and supply a
	// drop-down list of values in the web application.
	ValidValues []string `json:"validValues,omitempty"`

	// ValidateData: If true, data entered via the web application is
	// validated.
	ValidateData bool `json:"validateData,omitempty"`
}

type ColumnBaseColumn struct {
	// ColumnId: The id of the column in the base table from which this
	// column is derived.
	ColumnId int64 `json:"columnId,omitempty"`

	// TableIndex: Offset to the entry in the list of base tables in the
	// table definition.
	TableIndex int64 `json:"tableIndex,omitempty"`
}

type ColumnList struct {
	// Items: List of all requested columns.
	Items []*Column `json:"items,omitempty"`

	// Kind: The kind of item this is. For a column list, this is always
	// fusiontables#columnList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of columns for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type Geometry struct {
	// Geometries: The list of geometries in this geometry collection.
	Geometries []interface{} `json:"geometries,omitempty"`

	Geometry interface{} `json:"geometry,omitempty"`

	// Type: Type: A collection of geometries.
	Type string `json:"type,omitempty"`
}

type Import struct {
	// Kind: The kind of item this is. For an import, this is always
	// fusiontables#import.
	Kind string `json:"kind,omitempty"`

	// NumRowsReceived: The number of rows received from the import request.
	NumRowsReceived int64 `json:"numRowsReceived,omitempty,string"`
}

type Line struct {
	// Coordinates: The coordinates that define the line.
	Coordinates [][]float64 `json:"coordinates,omitempty"`

	// Type: Type: A line geometry.
	Type string `json:"type,omitempty"`
}

type LineStyle struct {
	// StrokeColor: Color of the line in #RRGGBB format.
	StrokeColor string `json:"strokeColor,omitempty"`

	// StrokeColorStyler: Column-value, gradient or buckets styler that is
	// used to determine the line color and opacity.
	StrokeColorStyler *StyleFunction `json:"strokeColorStyler,omitempty"`

	// StrokeOpacity: Opacity of the line : 0.0 (transparent) to 1.0
	// (opaque).
	StrokeOpacity float64 `json:"strokeOpacity,omitempty"`

	// StrokeWeight: Width of the line in pixels.
	StrokeWeight int64 `json:"strokeWeight,omitempty"`

	// StrokeWeightStyler: Column-value or bucket styler that is used to
	// determine the width of the line.
	StrokeWeightStyler *StyleFunction `json:"strokeWeightStyler,omitempty"`
}

type Point struct {
	// Coordinates: The coordinates that define the point.
	Coordinates []float64 `json:"coordinates,omitempty"`

	// Type: Point: A point geometry.
	Type string `json:"type,omitempty"`
}

type PointStyle struct {
	// IconName: Name of the icon. Use values defined in
	// http://www.google.com/fusiontables/DataSource?dsrcid=308519
	IconName string `json:"iconName,omitempty"`

	// IconStyler: Column or a bucket value from which the icon name is to
	// be determined.
	IconStyler *StyleFunction `json:"iconStyler,omitempty"`
}

type Polygon struct {
	// Coordinates: The coordinates that define the polygon.
	Coordinates [][][]float64 `json:"coordinates,omitempty"`

	// Type: Type: A polygon geometry.
	Type string `json:"type,omitempty"`
}

type PolygonStyle struct {
	// FillColor: Color of the interior of the polygon in #RRGGBB format.
	FillColor string `json:"fillColor,omitempty"`

	// FillColorStyler: Column-value, gradient, or bucket styler that is
	// used to determine the interior color and opacity of the polygon.
	FillColorStyler *StyleFunction `json:"fillColorStyler,omitempty"`

	// FillOpacity: Opacity of the interior of the polygon: 0.0
	// (transparent) to 1.0 (opaque).
	FillOpacity float64 `json:"fillOpacity,omitempty"`

	// StrokeColor: Color of the polygon border in #RRGGBB format.
	StrokeColor string `json:"strokeColor,omitempty"`

	// StrokeColorStyler: Column-value, gradient or buckets styler that is
	// used to determine the border color and opacity.
	StrokeColorStyler *StyleFunction `json:"strokeColorStyler,omitempty"`

	// StrokeOpacity: Opacity of the polygon border: 0.0 (transparent) to
	// 1.0 (opaque).
	StrokeOpacity float64 `json:"strokeOpacity,omitempty"`

	// StrokeWeight: Width of the polyon border in pixels.
	StrokeWeight int64 `json:"strokeWeight,omitempty"`

	// StrokeWeightStyler: Column-value or bucket styler that is used to
	// determine the width of the polygon border.
	StrokeWeightStyler *StyleFunction `json:"strokeWeightStyler,omitempty"`
}

type Sqlresponse struct {
	// Columns: Columns in the table.
	Columns []string `json:"columns,omitempty"`

	// Kind: The kind of item this is. For responses to SQL queries, this is
	// always fusiontables#sqlresponse.
	Kind string `json:"kind,omitempty"`

	// Rows: The rows in the table. For each cell we print out whatever cell
	// value (e.g., numeric, string) exists. Thus it is important that each
	// cell contains only one value.
	Rows [][]interface{} `json:"rows,omitempty"`
}

type StyleFunction struct {
	// Buckets: Bucket function that assigns a style based on the range a
	// column value falls into.
	Buckets []*Bucket `json:"buckets,omitempty"`

	// ColumnName: Name of the column whose value is used in the style.
	ColumnName string `json:"columnName,omitempty"`

	// Gradient: Gradient function that interpolates a range of colors based
	// on column value.
	Gradient *StyleFunctionGradient `json:"gradient,omitempty"`

	// Kind: Stylers can be one of three kinds: "fusiontables#fromColumn if
	// the column value is to be used as is, i.e., the column values can
	// have colors in #RRGGBBAA format or integer line widths or icon names;
	// fusiontables#gradient if the styling of the row is to be based on
	// applying the gradient function on the column value; or
	// fusiontables#buckets if the styling is to based on the bucket into
	// which the the column value falls.
	Kind string `json:"kind,omitempty"`
}

type StyleFunctionGradient struct {
	// Colors: Array with two or more colors.
	Colors []*StyleFunctionGradientColors `json:"colors,omitempty"`

	// Max: Higher-end of the interpolation range: rows with this value will
	// be assigned to colors[n-1].
	Max float64 `json:"max,omitempty"`

	// Min: Lower-end of the interpolation range: rows with this value will
	// be assigned to colors[0].
	Min float64 `json:"min,omitempty"`
}

type StyleFunctionGradientColors struct {
	// Color: Color in #RRGGBB format.
	Color string `json:"color,omitempty"`

	// Opacity: Opacity of the color: 0.0 (transparent) to 1.0 (opaque).
	Opacity float64 `json:"opacity,omitempty"`
}

type StyleSetting struct {
	// Kind: The kind of item this is. A StyleSetting contains the style
	// definitions for points, lines, and polygons in a table. Since a table
	// can have any one or all of them, a style definition can have point,
	// line and polygon style definitions.
	Kind string `json:"kind,omitempty"`

	// MarkerOptions: Style definition for points in the table.
	MarkerOptions *PointStyle `json:"markerOptions,omitempty"`

	// Name: Optional name for the style setting.
	Name string `json:"name,omitempty"`

	// PolygonOptions: Style definition for polygons in the table.
	PolygonOptions *PolygonStyle `json:"polygonOptions,omitempty"`

	// PolylineOptions: Style definition for lines in the table.
	PolylineOptions *LineStyle `json:"polylineOptions,omitempty"`

	// StyleId: Identifier for the style setting (unique only within
	// tables).
	StyleId int64 `json:"styleId,omitempty"`

	// TableId: Identifier for the table.
	TableId string `json:"tableId,omitempty"`
}

type StyleSettingList struct {
	// Items: All requested style settings.
	Items []*StyleSetting `json:"items,omitempty"`

	// Kind: The kind of item this is. For a style list, this is always
	// fusiontables#styleSettingList .
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more styles left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of styles for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type Table struct {
	// Attribution: Attribution assigned to the table.
	Attribution string `json:"attribution,omitempty"`

	// AttributionLink: Optional link for attribution.
	AttributionLink string `json:"attributionLink,omitempty"`

	// BaseTableIds: Base table identifier if this table is a view or merged
	// table.
	BaseTableIds []string `json:"baseTableIds,omitempty"`

	// ColumnPropertiesJsonSchema: Default JSON schema for validating all
	// JSON column properties.
	ColumnPropertiesJsonSchema string `json:"columnPropertiesJsonSchema,omitempty"`

	// Columns: Columns in the table.
	Columns []*Column `json:"columns,omitempty"`

	// Description: Description assigned to the table.
	Description string `json:"description,omitempty"`

	// IsExportable: Variable for whether table is exportable.
	IsExportable bool `json:"isExportable,omitempty"`

	// Kind: The kind of item this is. For a table, this is always
	// fusiontables#table.
	Kind string `json:"kind,omitempty"`

	// Name: Name assigned to a table.
	Name string `json:"name,omitempty"`

	// Sql: SQL that encodes the table definition for derived tables.
	Sql string `json:"sql,omitempty"`

	// TableId: Encrypted unique alphanumeric identifier for the table.
	TableId string `json:"tableId,omitempty"`

	// TablePropertiesJson: JSON object containing custom table properties.
	TablePropertiesJson string `json:"tablePropertiesJson,omitempty"`

	// TablePropertiesJsonSchema: JSON schema for validating the JSON table
	// properties.
	TablePropertiesJsonSchema string `json:"tablePropertiesJsonSchema,omitempty"`
}

type TableList struct {
	// Items: List of all requested tables.
	Items []*Table `json:"items,omitempty"`

	// Kind: The kind of item this is. For table list, this is always
	// fusiontables#tableList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Task struct {
	// Kind: Type of the resource. This is always "fusiontables#task".
	Kind string `json:"kind,omitempty"`

	// Progress: Task percentage completion.
	Progress string `json:"progress,omitempty"`

	// Started: false while the table is busy with some other task. true if
	// this background task is currently running.
	Started bool `json:"started,omitempty"`

	// TaskId: Identifier for the task.
	TaskId int64 `json:"taskId,omitempty,string"`

	// Type: Type of background task.
	Type string `json:"type,omitempty"`
}

type TaskList struct {
	// Items: List of all requested tasks.
	Items []*Task `json:"items,omitempty"`

	// Kind: Type of the resource. This is always "fusiontables#taskList".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of tasks for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

type Template struct {
	// AutomaticColumnNames: List of columns from which the template is to
	// be automatically constructed. Only one of body or automaticColumns
	// can be specified.
	AutomaticColumnNames []string `json:"automaticColumnNames,omitempty"`

	// Body: Body of the template. It contains HTML with {column_name} to
	// insert values from a particular column. The body is sanitized to
	// remove certain tags, e.g., script. Only one of body or
	// automaticColumns can be specified.
	Body string `json:"body,omitempty"`

	// Kind: The kind of item this is. For a template, this is always
	// fusiontables#template.
	Kind string `json:"kind,omitempty"`

	// Name: Optional name assigned to a template.
	Name string `json:"name,omitempty"`

	// TableId: Identifier for the table for which the template is defined.
	TableId string `json:"tableId,omitempty"`

	// TemplateId: Identifier for the template, unique within the context of
	// a particular table.
	TemplateId int64 `json:"templateId,omitempty"`
}

type TemplateList struct {
	// Items: List of all requested templates.
	Items []*Template `json:"items,omitempty"`

	// Kind: The kind of item this is. For a template list, this is always
	// fusiontables#templateList .
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access the next page of this result. No
	// token is displayed if there are no more pages left.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalItems: Total number of templates for the table.
	TotalItems int64 `json:"totalItems,omitempty"`
}

// method id "fusiontables.column.delete":

type ColumnDeleteCall struct {
	s             *Service
	tableId       string
	columnId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes the specified column.

func (r *ColumnService) Delete(tableId string, columnId string) *ColumnDeleteCall {
	return &ColumnDeleteCall{
		s:             r.s,
		tableId:       tableId,
		columnId:      columnId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/columns/{columnId}",
	}
}

func (c *ColumnDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes the specified column.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.column.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table from which the column is being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.column.get":

type ColumnGetCall struct {
	s             *Service
	tableId       string
	columnId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a specific column by its ID.

func (r *ColumnService) Get(tableId string, columnId string) *ColumnGetCall {
	return &ColumnGetCall{
		s:             r.s,
		tableId:       tableId,
		columnId:      columnId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/columns/{columnId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnGetCall) Fields(s ...googleapi.Field) *ColumnGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ColumnGetCall) Do() (*Column, error) {
	var returnValue *Column
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a specific column by its ID.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.column.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column that is being requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table to which the column belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.column.insert":

type ColumnInsertCall struct {
	s             *Service
	tableId       string
	column        *Column
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Adds a new column to the table.

func (r *ColumnService) Insert(tableId string, column *Column) *ColumnInsertCall {
	return &ColumnInsertCall{
		s:             r.s,
		tableId:       tableId,
		column:        column,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/columns",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnInsertCall) Fields(s ...googleapi.Field) *ColumnInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ColumnInsertCall) Do() (*Column, error) {
	var returnValue *Column
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.column,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds a new column to the table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.column.insert",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table for which a new column is being added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns",
	//   "request": {
	//     "$ref": "Column"
	//   },
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.column.list":

type ColumnListCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of columns.

func (r *ColumnService) List(tableId string) *ColumnListCall {
	return &ColumnListCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/columns",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of columns to return. Default is 5.
func (c *ColumnListCall) MaxResults(maxResults int64) *ColumnListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *ColumnListCall) PageToken(pageToken string) *ColumnListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnListCall) Fields(s ...googleapi.Field) *ColumnListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ColumnListCall) Do() (*ColumnList, error) {
	var returnValue *ColumnList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of columns.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.column.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of columns to return. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table whose columns are being listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns",
	//   "response": {
	//     "$ref": "ColumnList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.column.patch":

type ColumnPatchCall struct {
	s             *Service
	tableId       string
	columnId      string
	column        *Column
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates the name or type of an existing column. This method
// supports patch semantics.

func (r *ColumnService) Patch(tableId string, columnId string, column *Column) *ColumnPatchCall {
	return &ColumnPatchCall{
		s:             r.s,
		tableId:       tableId,
		columnId:      columnId,
		column:        column,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/columns/{columnId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnPatchCall) Fields(s ...googleapi.Field) *ColumnPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ColumnPatchCall) Do() (*Column, error) {
	var returnValue *Column
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.column,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates the name or type of an existing column. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.column.patch",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table for which the column is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "request": {
	//     "$ref": "Column"
	//   },
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.column.update":

type ColumnUpdateCall struct {
	s             *Service
	tableId       string
	columnId      string
	column        *Column
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates the name or type of an existing column.

func (r *ColumnService) Update(tableId string, columnId string, column *Column) *ColumnUpdateCall {
	return &ColumnUpdateCall{
		s:             r.s,
		tableId:       tableId,
		columnId:      columnId,
		column:        column,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/columns/{columnId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ColumnUpdateCall) Fields(s ...googleapi.Field) *ColumnUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ColumnUpdateCall) Do() (*Column, error) {
	var returnValue *Column
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":  c.tableId,
		"columnId": c.columnId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.column,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates the name or type of an existing column.",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.column.update",
	//   "parameterOrder": [
	//     "tableId",
	//     "columnId"
	//   ],
	//   "parameters": {
	//     "columnId": {
	//       "description": "Name or identifier for the column that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table for which the column is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/columns/{columnId}",
	//   "request": {
	//     "$ref": "Column"
	//   },
	//   "response": {
	//     "$ref": "Column"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.query.sql":

type QuerySqlCall struct {
	s             *Service
	sql           string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Sql: Executes a Fusion Tables SQL statement, which can be any of
// -
// SELECT
// - INSERT
// - UPDATE
// - DELETE
// - SHOW
// - DESCRIBE
// - CREATE
// statement.

func (r *QueryService) Sql(sql string) *QuerySqlCall {
	return &QuerySqlCall{
		s:             r.s,
		sql:           sql,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "query",
	}
}

// Hdrs sets the optional parameter "hdrs": Whether column names are
// included in the first row. Default is true.
func (c *QuerySqlCall) Hdrs(hdrs bool) *QuerySqlCall {
	c.params_.Set("hdrs", fmt.Sprintf("%v", hdrs))
	return c
}

// Typed sets the optional parameter "typed": Whether typed values are
// returned in the (JSON) response: numbers for numeric values and
// parsed geometries for KML values. Default is true.
func (c *QuerySqlCall) Typed(typed bool) *QuerySqlCall {
	c.params_.Set("typed", fmt.Sprintf("%v", typed))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySqlCall) Fields(s ...googleapi.Field) *QuerySqlCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *QuerySqlCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	c.params_.Set("sql", fmt.Sprintf("%v", c.sql))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *QuerySqlCall) Do() (*Sqlresponse, error) {
	var returnValue *Sqlresponse
	c.params_.Set("sql", fmt.Sprintf("%v", c.sql))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Executes a Fusion Tables SQL statement, which can be any of \n- SELECT\n- INSERT\n- UPDATE\n- DELETE\n- SHOW\n- DESCRIBE\n- CREATE statement.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.query.sql",
	//   "parameterOrder": [
	//     "sql"
	//   ],
	//   "parameters": {
	//     "hdrs": {
	//       "description": "Whether column names are included in the first row. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sql": {
	//       "description": "A Fusion Tables SQL statement, which can be any of \n- SELECT\n- INSERT\n- UPDATE\n- DELETE\n- SHOW\n- DESCRIBE\n- CREATE",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typed": {
	//       "description": "Whether typed values are returned in the (JSON) response: numbers for numeric values and parsed geometries for KML values. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "query",
	//   "response": {
	//     "$ref": "Sqlresponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "fusiontables.query.sqlGet":

type QuerySqlGetCall struct {
	s             *Service
	sql           string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// SqlGet: Executes a SQL statement which can be any of
// - SELECT
// -
// SHOW
// - DESCRIBE

func (r *QueryService) SqlGet(sql string) *QuerySqlGetCall {
	return &QuerySqlGetCall{
		s:             r.s,
		sql:           sql,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "query",
	}
}

// Hdrs sets the optional parameter "hdrs": Whether column names are
// included (in the first row). Default is true.
func (c *QuerySqlGetCall) Hdrs(hdrs bool) *QuerySqlGetCall {
	c.params_.Set("hdrs", fmt.Sprintf("%v", hdrs))
	return c
}

// Typed sets the optional parameter "typed": Whether typed values are
// returned in the (JSON) response: numbers for numeric values and
// parsed geometries for KML values. Default is true.
func (c *QuerySqlGetCall) Typed(typed bool) *QuerySqlGetCall {
	c.params_.Set("typed", fmt.Sprintf("%v", typed))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QuerySqlGetCall) Fields(s ...googleapi.Field) *QuerySqlGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *QuerySqlGetCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	c.params_.Set("sql", fmt.Sprintf("%v", c.sql))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *QuerySqlGetCall) Do() (*Sqlresponse, error) {
	var returnValue *Sqlresponse
	c.params_.Set("sql", fmt.Sprintf("%v", c.sql))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Executes a SQL statement which can be any of \n- SELECT\n- SHOW\n- DESCRIBE",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.query.sqlGet",
	//   "parameterOrder": [
	//     "sql"
	//   ],
	//   "parameters": {
	//     "hdrs": {
	//       "description": "Whether column names are included (in the first row). Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sql": {
	//       "description": "A SQL statement which can be any of \n- SELECT\n- SHOW\n- DESCRIBE",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "typed": {
	//       "description": "Whether typed values are returned in the (JSON) response: numbers for numeric values and parsed geometries for KML values. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "query",
	//   "response": {
	//     "$ref": "Sqlresponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "fusiontables.style.delete":

type StyleDeleteCall struct {
	s             *Service
	tableId       string
	styleId       int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a style.

func (r *StyleService) Delete(tableId string, styleId int64) *StyleDeleteCall {
	return &StyleDeleteCall{
		s:             r.s,
		tableId:       tableId,
		styleId:       styleId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/styles/{styleId}",
	}
}

func (c *StyleDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a style.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.style.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (within a table) for the style being deleted",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table from which the style is being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.style.get":

type StyleGetCall struct {
	s             *Service
	tableId       string
	styleId       int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a specific style.

func (r *StyleService) Get(tableId string, styleId int64) *StyleGetCall {
	return &StyleGetCall{
		s:             r.s,
		tableId:       tableId,
		styleId:       styleId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/styles/{styleId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleGetCall) Fields(s ...googleapi.Field) *StyleGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StyleGetCall) Do() (*StyleSetting, error) {
	var returnValue *StyleSetting
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a specific style.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.style.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (integer) for a specific style in a table",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table to which the requested style belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.style.insert":

type StyleInsertCall struct {
	s             *Service
	tableId       string
	stylesetting  *StyleSetting
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Adds a new style for the table.

func (r *StyleService) Insert(tableId string, stylesetting *StyleSetting) *StyleInsertCall {
	return &StyleInsertCall{
		s:             r.s,
		tableId:       tableId,
		stylesetting:  stylesetting,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/styles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleInsertCall) Fields(s ...googleapi.Field) *StyleInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StyleInsertCall) Do() (*StyleSetting, error) {
	var returnValue *StyleSetting
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.stylesetting,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Adds a new style for the table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.style.insert",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table for which a new style is being added",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles",
	//   "request": {
	//     "$ref": "StyleSetting"
	//   },
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.style.list":

type StyleListCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of styles.

func (r *StyleService) List(tableId string) *StyleListCall {
	return &StyleListCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/styles",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of styles to return.  Default is 5.
func (c *StyleListCall) MaxResults(maxResults int64) *StyleListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *StyleListCall) PageToken(pageToken string) *StyleListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleListCall) Fields(s ...googleapi.Field) *StyleListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StyleListCall) Do() (*StyleSettingList, error) {
	var returnValue *StyleSettingList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of styles.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.style.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of styles to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Table whose styles are being listed",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles",
	//   "response": {
	//     "$ref": "StyleSettingList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.style.patch":

type StylePatchCall struct {
	s             *Service
	tableId       string
	styleId       int64
	stylesetting  *StyleSetting
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing style. This method supports patch
// semantics.

func (r *StyleService) Patch(tableId string, styleId int64, stylesetting *StyleSetting) *StylePatchCall {
	return &StylePatchCall{
		s:             r.s,
		tableId:       tableId,
		styleId:       styleId,
		stylesetting:  stylesetting,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/styles/{styleId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StylePatchCall) Fields(s ...googleapi.Field) *StylePatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StylePatchCall) Do() (*StyleSetting, error) {
	var returnValue *StyleSetting
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.stylesetting,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing style. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.style.patch",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (within a table) for the style being updated.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose style is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "request": {
	//     "$ref": "StyleSetting"
	//   },
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.style.update":

type StyleUpdateCall struct {
	s             *Service
	tableId       string
	styleId       int64
	stylesetting  *StyleSetting
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing style.

func (r *StyleService) Update(tableId string, styleId int64, stylesetting *StyleSetting) *StyleUpdateCall {
	return &StyleUpdateCall{
		s:             r.s,
		tableId:       tableId,
		styleId:       styleId,
		stylesetting:  stylesetting,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/styles/{styleId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *StyleUpdateCall) Fields(s ...googleapi.Field) *StyleUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *StyleUpdateCall) Do() (*StyleSetting, error) {
	var returnValue *StyleSetting
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
		"styleId": strconv.FormatInt(c.styleId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.stylesetting,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing style.",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.style.update",
	//   "parameterOrder": [
	//     "tableId",
	//     "styleId"
	//   ],
	//   "parameters": {
	//     "styleId": {
	//       "description": "Identifier (within a table) for the style being updated.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose style is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/styles/{styleId}",
	//   "request": {
	//     "$ref": "StyleSetting"
	//   },
	//   "response": {
	//     "$ref": "StyleSetting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.copy":

type TableCopyCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Copy: Copies a table.

func (r *TableService) Copy(tableId string) *TableCopyCall {
	return &TableCopyCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/copy",
	}
}

// CopyPresentation sets the optional parameter "copyPresentation":
// Whether to also copy tabs, styles, and templates. Default is false.
func (c *TableCopyCall) CopyPresentation(copyPresentation bool) *TableCopyCall {
	c.params_.Set("copyPresentation", fmt.Sprintf("%v", copyPresentation))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableCopyCall) Fields(s ...googleapi.Field) *TableCopyCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableCopyCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Copies a table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.copy",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "copyPresentation": {
	//       "description": "Whether to also copy tabs, styles, and templates. Default is false.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "tableId": {
	//       "description": "ID of the table that is being copied.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/copy",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.table.delete":

type TableDeleteCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a table.

func (r *TableService) Delete(tableId string) *TableDeleteCall {
	return &TableDeleteCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}",
	}
}

func (c *TableDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a table.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.table.delete",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "ID of the table to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.get":

type TableGetCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a specific table by its ID.

func (r *TableService) Get(tableId string) *TableGetCall {
	return &TableGetCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableGetCall) Fields(s ...googleapi.Field) *TableGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableGetCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a specific table by its ID.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.table.get",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Identifier for the table being requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.table.importRows":

type TableImportRowsCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// ImportRows: Imports more rows into a table.

func (r *TableService) ImportRows(tableId string) *TableImportRowsCall {
	return &TableImportRowsCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/import",
		context_:      context.TODO(),
	}
}

// Delimiter sets the optional parameter "delimiter": The delimiter used
// to separate cell values. This can only consist of a single character.
// Default is ,.
func (c *TableImportRowsCall) Delimiter(delimiter string) *TableImportRowsCall {
	c.params_.Set("delimiter", fmt.Sprintf("%v", delimiter))
	return c
}

// Encoding sets the optional parameter "encoding": The encoding of the
// content. Default is UTF-8. Use auto-detect if you are unsure of the
// encoding.
func (c *TableImportRowsCall) Encoding(encoding string) *TableImportRowsCall {
	c.params_.Set("encoding", fmt.Sprintf("%v", encoding))
	return c
}

// EndLine sets the optional parameter "endLine": The index of the line
// up to which data will be imported. Default is to import the entire
// file. If endLine is negative, it is an offset from the end of the
// file; the imported content will exclude the last endLine lines.
func (c *TableImportRowsCall) EndLine(endLine int64) *TableImportRowsCall {
	c.params_.Set("endLine", fmt.Sprintf("%v", endLine))
	return c
}

// IsStrict sets the optional parameter "isStrict": Whether the imported
// CSV must have the same number of values for each row. If false, rows
// with fewer values will be padded with empty values. Default is true.
func (c *TableImportRowsCall) IsStrict(isStrict bool) *TableImportRowsCall {
	c.params_.Set("isStrict", fmt.Sprintf("%v", isStrict))
	return c
}

// StartLine sets the optional parameter "startLine": The index of the
// first line from which to start importing, inclusive. Default is 0.
func (c *TableImportRowsCall) StartLine(startLine int64) *TableImportRowsCall {
	c.params_.Set("startLine", fmt.Sprintf("%v", startLine))
	return c
}

// MediaUpload takes a context and UploadCaller interface
func (c *TableImportRowsCall) Upload(ctx context.Context, u googleapi.UploadCaller) *TableImportRowsCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/fusiontables/v2/tables/{tableId}/import"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/fusiontables/v2/tables/{tableId}/import"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *TableImportRowsCall) Media(r io.Reader) *TableImportRowsCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/fusiontables/v2/tables/{tableId}/import"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TableImportRowsCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TableImportRowsCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/fusiontables/v2/tables/{tableId}/import"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TableImportRowsCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TableImportRowsCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableImportRowsCall) Fields(s ...googleapi.Field) *TableImportRowsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableImportRowsCall) Do() (*Import, error) {
	var returnValue *Import
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Imports more rows into a table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.importRows",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "maxSize": "250MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/fusiontables/v2/tables/{tableId}/import"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/fusiontables/v2/tables/{tableId}/import"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "delimiter": {
	//       "description": "The delimiter used to separate cell values. This can only consist of a single character. Default is ,.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "encoding": {
	//       "description": "The encoding of the content. Default is UTF-8. Use auto-detect if you are unsure of the encoding.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endLine": {
	//       "description": "The index of the line up to which data will be imported. Default is to import the entire file. If endLine is negative, it is an offset from the end of the file; the imported content will exclude the last endLine lines.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "isStrict": {
	//       "description": "Whether the imported CSV must have the same number of values for each row. If false, rows with fewer values will be padded with empty values. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "startLine": {
	//       "description": "The index of the first line from which to start importing, inclusive. Default is 0.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "The table into which new rows are being imported.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/import",
	//   "response": {
	//     "$ref": "Import"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "fusiontables.table.importTable":

type TableImportTableCall struct {
	s             *Service
	name          string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// ImportTable: Imports a new table.

func (r *TableService) ImportTable(name string) *TableImportTableCall {
	return &TableImportTableCall{
		s:             r.s,
		name:          name,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/import",
		context_:      context.TODO(),
	}
}

// Delimiter sets the optional parameter "delimiter": The delimiter used
// to separate cell values. This can only consist of a single character.
// Default is ,.
func (c *TableImportTableCall) Delimiter(delimiter string) *TableImportTableCall {
	c.params_.Set("delimiter", fmt.Sprintf("%v", delimiter))
	return c
}

// Encoding sets the optional parameter "encoding": The encoding of the
// content. Default is UTF-8. Use auto-detect if you are unsure of the
// encoding.
func (c *TableImportTableCall) Encoding(encoding string) *TableImportTableCall {
	c.params_.Set("encoding", fmt.Sprintf("%v", encoding))
	return c
}

// MediaUpload takes a context and UploadCaller interface
func (c *TableImportTableCall) Upload(ctx context.Context, u googleapi.UploadCaller) *TableImportTableCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/fusiontables/v2/tables/import"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/fusiontables/v2/tables/import"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *TableImportTableCall) Media(r io.Reader) *TableImportTableCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/fusiontables/v2/tables/import"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TableImportTableCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TableImportTableCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/fusiontables/v2/tables/import"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TableImportTableCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TableImportTableCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableImportTableCall) Fields(s ...googleapi.Field) *TableImportTableCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableImportTableCall) Do() (*Table, error) {
	var returnValue *Table
	c.params_.Set("name", fmt.Sprintf("%v", c.name))
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Imports a new table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.importTable",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "maxSize": "250MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/fusiontables/v2/tables/import"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/fusiontables/v2/tables/import"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "delimiter": {
	//       "description": "The delimiter used to separate cell values. This can only consist of a single character. Default is ,.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "encoding": {
	//       "description": "The encoding of the content. Default is UTF-8. Use auto-detect if you are unsure of the encoding.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name to be assigned to the new table.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/import",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "fusiontables.table.insert":

type TableInsertCall struct {
	s             *Service
	table         *Table
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new table.

func (r *TableService) Insert(table *Table) *TableInsertCall {
	return &TableInsertCall{
		s:             r.s,
		table:         table,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableInsertCall) Fields(s ...googleapi.Field) *TableInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableInsertCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.table,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.insert",
	//   "path": "tables",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.list":

type TableListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of tables a user owns.

func (r *TableService) List() *TableListCall {
	return &TableListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of tables to return. Default is 5.
func (c *TableListCall) MaxResults(maxResults int64) *TableListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *TableListCall) PageToken(pageToken string) *TableListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableListCall) Fields(s ...googleapi.Field) *TableListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableListCall) Do() (*TableList, error) {
	var returnValue *TableList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of tables a user owns.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.table.list",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of tables to return. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables",
	//   "response": {
	//     "$ref": "TableList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.table.patch":

type TablePatchCall struct {
	s             *Service
	tableId       string
	table         *Table
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing table. Unless explicitly requested, only
// the name, description, and attribution will be updated. This method
// supports patch semantics.

func (r *TableService) Patch(tableId string, table *Table) *TablePatchCall {
	return &TablePatchCall{
		s:             r.s,
		tableId:       tableId,
		table:         table,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}",
	}
}

// ReplaceViewDefinition sets the optional parameter
// "replaceViewDefinition": Whether the view definition is also updated.
// The specified view definition replaces the existing one. Only a view
// can be updated with a new definition.
func (c *TablePatchCall) ReplaceViewDefinition(replaceViewDefinition bool) *TablePatchCall {
	c.params_.Set("replaceViewDefinition", fmt.Sprintf("%v", replaceViewDefinition))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablePatchCall) Fields(s ...googleapi.Field) *TablePatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TablePatchCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
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
	//   "description": "Updates an existing table. Unless explicitly requested, only the name, description, and attribution will be updated. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.table.patch",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "replaceViewDefinition": {
	//       "description": "Whether the view definition is also updated. The specified view definition replaces the existing one. Only a view can be updated with a new definition.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "tableId": {
	//       "description": "ID of the table that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.table.replaceRows":

type TableReplaceRowsCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// ReplaceRows: Replaces rows of an existing table. Current rows remain
// visible until all replacement rows are ready.

func (r *TableService) ReplaceRows(tableId string) *TableReplaceRowsCall {
	return &TableReplaceRowsCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/replace",
		context_:      context.TODO(),
	}
}

// Delimiter sets the optional parameter "delimiter": The delimiter used
// to separate cell values. This can only consist of a single character.
// Default is ,.
func (c *TableReplaceRowsCall) Delimiter(delimiter string) *TableReplaceRowsCall {
	c.params_.Set("delimiter", fmt.Sprintf("%v", delimiter))
	return c
}

// Encoding sets the optional parameter "encoding": The encoding of the
// content. Default is UTF-8. Use 'auto-detect' if you are unsure of the
// encoding.
func (c *TableReplaceRowsCall) Encoding(encoding string) *TableReplaceRowsCall {
	c.params_.Set("encoding", fmt.Sprintf("%v", encoding))
	return c
}

// EndLine sets the optional parameter "endLine": The index of the line
// up to which data will be imported. Default is to import the entire
// file. If endLine is negative, it is an offset from the end of the
// file; the imported content will exclude the last endLine lines.
func (c *TableReplaceRowsCall) EndLine(endLine int64) *TableReplaceRowsCall {
	c.params_.Set("endLine", fmt.Sprintf("%v", endLine))
	return c
}

// IsStrict sets the optional parameter "isStrict": Whether the imported
// CSV must have the same number of column values for each row. If true,
// throws an exception if the CSV does not have the same number of
// columns. If false, rows with fewer column values will be padded with
// empty values. Default is true.
func (c *TableReplaceRowsCall) IsStrict(isStrict bool) *TableReplaceRowsCall {
	c.params_.Set("isStrict", fmt.Sprintf("%v", isStrict))
	return c
}

// StartLine sets the optional parameter "startLine": The index of the
// first line from which to start importing, inclusive. Default is 0.
func (c *TableReplaceRowsCall) StartLine(startLine int64) *TableReplaceRowsCall {
	c.params_.Set("startLine", fmt.Sprintf("%v", startLine))
	return c
}

// MediaUpload takes a context and UploadCaller interface
func (c *TableReplaceRowsCall) Upload(ctx context.Context, u googleapi.UploadCaller) *TableReplaceRowsCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/fusiontables/v2/tables/{tableId}/replace"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/fusiontables/v2/tables/{tableId}/replace"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *TableReplaceRowsCall) Media(r io.Reader) *TableReplaceRowsCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/fusiontables/v2/tables/{tableId}/replace"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TableReplaceRowsCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TableReplaceRowsCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/fusiontables/v2/tables/{tableId}/replace"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TableReplaceRowsCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TableReplaceRowsCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableReplaceRowsCall) Fields(s ...googleapi.Field) *TableReplaceRowsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableReplaceRowsCall) Do() (*Task, error) {
	var returnValue *Task
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Replaces rows of an existing table. Current rows remain visible until all replacement rows are ready.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.table.replaceRows",
	//   "mediaUpload": {
	//     "accept": [
	//       "application/octet-stream"
	//     ],
	//     "maxSize": "250MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/fusiontables/v2/tables/{tableId}/replace"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/fusiontables/v2/tables/{tableId}/replace"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "delimiter": {
	//       "description": "The delimiter used to separate cell values. This can only consist of a single character. Default is ,.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "encoding": {
	//       "description": "The encoding of the content. Default is UTF-8. Use 'auto-detect' if you are unsure of the encoding.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endLine": {
	//       "description": "The index of the line up to which data will be imported. Default is to import the entire file. If endLine is negative, it is an offset from the end of the file; the imported content will exclude the last endLine lines.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "isStrict": {
	//       "description": "Whether the imported CSV must have the same number of column values for each row. If true, throws an exception if the CSV does not have the same number of columns. If false, rows with fewer column values will be padded with empty values. Default is true.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "startLine": {
	//       "description": "The index of the first line from which to start importing, inclusive. Default is 0.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose rows will be replaced.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/replace",
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "fusiontables.table.update":

type TableUpdateCall struct {
	s             *Service
	tableId       string
	table         *Table
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing table. Unless explicitly requested, only
// the name, description, and attribution will be updated.

func (r *TableService) Update(tableId string, table *Table) *TableUpdateCall {
	return &TableUpdateCall{
		s:             r.s,
		tableId:       tableId,
		table:         table,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}",
	}
}

// ReplaceViewDefinition sets the optional parameter
// "replaceViewDefinition": Whether the view definition is also updated.
// The specified view definition replaces the existing one. Only a view
// can be updated with a new definition.
func (c *TableUpdateCall) ReplaceViewDefinition(replaceViewDefinition bool) *TableUpdateCall {
	c.params_.Set("replaceViewDefinition", fmt.Sprintf("%v", replaceViewDefinition))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TableUpdateCall) Fields(s ...googleapi.Field) *TableUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TableUpdateCall) Do() (*Table, error) {
	var returnValue *Table
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
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
	//   "description": "Updates an existing table. Unless explicitly requested, only the name, description, and attribution will be updated.",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.table.update",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "replaceViewDefinition": {
	//       "description": "Whether the view definition is also updated. The specified view definition replaces the existing one. Only a view can be updated with a new definition.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "tableId": {
	//       "description": "ID of the table that is being updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.task.delete":

type TaskDeleteCall struct {
	s             *Service
	tableId       string
	taskId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a specific task by its ID, unless that task has
// already started running.

func (r *TaskService) Delete(tableId string, taskId string) *TaskDeleteCall {
	return &TaskDeleteCall{
		s:             r.s,
		tableId:       tableId,
		taskId:        taskId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/tasks/{taskId}",
	}
}

func (c *TaskDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
		"taskId":  c.taskId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a specific task by its ID, unless that task has already started running.",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.task.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "taskId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table from which the task is being deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskId": {
	//       "description": "The identifier of the task to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/tasks/{taskId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.task.get":

type TaskGetCall struct {
	s             *Service
	tableId       string
	taskId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a specific task by its ID.

func (r *TaskService) Get(tableId string, taskId string) *TaskGetCall {
	return &TaskGetCall{
		s:             r.s,
		tableId:       tableId,
		taskId:        taskId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/tasks/{taskId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TaskGetCall) Fields(s ...googleapi.Field) *TaskGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TaskGetCall) Do() (*Task, error) {
	var returnValue *Task
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
		"taskId":  c.taskId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a specific task by its ID.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.task.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "taskId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the task belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskId": {
	//       "description": "The identifier of the task to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/tasks/{taskId}",
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.task.list":

type TaskListCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of tasks.

func (r *TaskService) List(tableId string) *TaskListCall {
	return &TaskListCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/tasks",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of tasks to return. Default is 5.
func (c *TaskListCall) MaxResults(maxResults int64) *TaskListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which result page to return.
func (c *TaskListCall) PageToken(pageToken string) *TaskListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// StartIndex sets the optional parameter "startIndex": Index of the
// first result returned in the current page.
func (c *TaskListCall) StartIndex(startIndex int64) *TaskListCall {
	c.params_.Set("startIndex", fmt.Sprintf("%v", startIndex))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TaskListCall) Fields(s ...googleapi.Field) *TaskListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TaskListCall) Do() (*TaskList, error) {
	var returnValue *TaskList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of tasks.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.task.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of tasks to return. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which result page to return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "Index of the first result returned in the current page.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "tableId": {
	//       "description": "Table whose tasks are being listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/tasks",
	//   "response": {
	//     "$ref": "TaskList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.template.delete":

type TemplateDeleteCall struct {
	s             *Service
	tableId       string
	templateId    int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a template

func (r *TemplateService) Delete(tableId string, templateId int64) *TemplateDeleteCall {
	return &TemplateDeleteCall{
		s:             r.s,
		tableId:       tableId,
		templateId:    templateId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/templates/{templateId}",
	}
}

func (c *TemplateDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a template",
	//   "httpMethod": "DELETE",
	//   "id": "fusiontables.template.delete",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table from which the template is being deleted",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template which is being deleted",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.template.get":

type TemplateGetCall struct {
	s             *Service
	tableId       string
	templateId    int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a specific template by its id

func (r *TemplateService) Get(tableId string, templateId int64) *TemplateGetCall {
	return &TemplateGetCall{
		s:             r.s,
		tableId:       tableId,
		templateId:    templateId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/templates/{templateId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateGetCall) Fields(s ...googleapi.Field) *TemplateGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TemplateGetCall) Do() (*Template, error) {
	var returnValue *Template
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a specific template by its id",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.template.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the template belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template that is being requested",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.template.insert":

type TemplateInsertCall struct {
	s             *Service
	tableId       string
	template      *Template
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new template for the table.

func (r *TemplateService) Insert(tableId string, template *Template) *TemplateInsertCall {
	return &TemplateInsertCall{
		s:             r.s,
		tableId:       tableId,
		template:      template,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/templates",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateInsertCall) Fields(s ...googleapi.Field) *TemplateInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TemplateInsertCall) Do() (*Template, error) {
	var returnValue *Template
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.template,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new template for the table.",
	//   "httpMethod": "POST",
	//   "id": "fusiontables.template.insert",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table for which a new template is being created",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates",
	//   "request": {
	//     "$ref": "Template"
	//   },
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.template.list":

type TemplateListCall struct {
	s             *Service
	tableId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of templates.

func (r *TemplateService) List(tableId string) *TemplateListCall {
	return &TemplateListCall{
		s:             r.s,
		tableId:       tableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/templates",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of templates to return.  Default is 5.
func (c *TemplateListCall) MaxResults(maxResults int64) *TemplateListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Continuation token
// specifying which results page to return.
func (c *TemplateListCall) PageToken(pageToken string) *TemplateListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateListCall) Fields(s ...googleapi.Field) *TemplateListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TemplateListCall) Do() (*TemplateList, error) {
	var returnValue *TemplateList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId": c.tableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of templates.",
	//   "httpMethod": "GET",
	//   "id": "fusiontables.template.list",
	//   "parameterOrder": [
	//     "tableId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of templates to return. Optional. Default is 5.",
	//       "format": "uint32",
	//       "location": "query",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Continuation token specifying which results page to return. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "Identifier for the table whose templates are being requested",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates",
	//   "response": {
	//     "$ref": "TemplateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables",
	//     "https://www.googleapis.com/auth/fusiontables.readonly"
	//   ]
	// }

}

// method id "fusiontables.template.patch":

type TemplatePatchCall struct {
	s             *Service
	tableId       string
	templateId    int64
	template      *Template
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing template. This method supports patch
// semantics.

func (r *TemplateService) Patch(tableId string, templateId int64, template *Template) *TemplatePatchCall {
	return &TemplatePatchCall{
		s:             r.s,
		tableId:       tableId,
		templateId:    templateId,
		template:      template,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/templates/{templateId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplatePatchCall) Fields(s ...googleapi.Field) *TemplatePatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TemplatePatchCall) Do() (*Template, error) {
	var returnValue *Template
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.template,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing template. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fusiontables.template.patch",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the updated template belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template that is being updated",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "request": {
	//     "$ref": "Template"
	//   },
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}

// method id "fusiontables.template.update":

type TemplateUpdateCall struct {
	s             *Service
	tableId       string
	templateId    int64
	template      *Template
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing template

func (r *TemplateService) Update(tableId string, templateId int64, template *Template) *TemplateUpdateCall {
	return &TemplateUpdateCall{
		s:             r.s,
		tableId:       tableId,
		templateId:    templateId,
		template:      template,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "tables/{tableId}/templates/{templateId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TemplateUpdateCall) Fields(s ...googleapi.Field) *TemplateUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TemplateUpdateCall) Do() (*Template, error) {
	var returnValue *Template
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"tableId":    c.tableId,
		"templateId": strconv.FormatInt(c.templateId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.template,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing template",
	//   "httpMethod": "PUT",
	//   "id": "fusiontables.template.update",
	//   "parameterOrder": [
	//     "tableId",
	//     "templateId"
	//   ],
	//   "parameters": {
	//     "tableId": {
	//       "description": "Table to which the updated template belongs",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "templateId": {
	//       "description": "Identifier for the template that is being updated",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     }
	//   },
	//   "path": "tables/{tableId}/templates/{templateId}",
	//   "request": {
	//     "$ref": "Template"
	//   },
	//   "response": {
	//     "$ref": "Template"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fusiontables"
	//   ]
	// }

}
