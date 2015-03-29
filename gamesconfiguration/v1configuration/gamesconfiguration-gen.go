// Package gamesconfiguration provides access to the Google Play Game Services Publishing API.
//
// See https://developers.google.com/games/services
//
// Usage example:
//
//   import "github.com/jfcote87/api/gamesconfiguration/v1configuration"
//   ...
//   gamesconfigurationService, err := gamesconfiguration.New(oauthHttpClient)
package gamesconfiguration

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

const apiId = "gamesConfiguration:v1configuration"
const apiName = "gamesConfiguration"
const apiVersion = "v1configuration"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/games/v1configuration/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Play Android Developer account
	AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.AchievementConfigurations = NewAchievementConfigurationsService(s)
	s.ImageConfigurations = NewImageConfigurationsService(s)
	s.LeaderboardConfigurations = NewLeaderboardConfigurationsService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	AchievementConfigurations *AchievementConfigurationsService

	ImageConfigurations *ImageConfigurationsService

	LeaderboardConfigurations *LeaderboardConfigurationsService
}

func NewAchievementConfigurationsService(s *Service) *AchievementConfigurationsService {
	rs := &AchievementConfigurationsService{s: s}
	return rs
}

type AchievementConfigurationsService struct {
	s *Service
}

func NewImageConfigurationsService(s *Service) *ImageConfigurationsService {
	rs := &ImageConfigurationsService{s: s}
	return rs
}

type ImageConfigurationsService struct {
	s *Service
}

func NewLeaderboardConfigurationsService(s *Service) *LeaderboardConfigurationsService {
	rs := &LeaderboardConfigurationsService{s: s}
	return rs
}

type LeaderboardConfigurationsService struct {
	s *Service
}

type AchievementConfiguration struct {
	// AchievementType: The type of the achievement.
	// Possible values are:
	//
	// - "STANDARD" - Achievement is either locked or unlocked.
	// -
	// "INCREMENTAL" - Achievement is incremental.
	AchievementType string `json:"achievementType,omitempty"`

	// Draft: The draft data of the achievement.
	Draft *AchievementConfigurationDetail `json:"draft,omitempty"`

	// Id: The ID of the achievement.
	Id string `json:"id,omitempty"`

	// InitialState: The initial state of the achievement.
	// Possible values
	// are:
	// - "HIDDEN" - Achievement is hidden.
	// - "REVEALED" -
	// Achievement is revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	InitialState string `json:"initialState,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#achievementConfiguration.
	Kind string `json:"kind,omitempty"`

	// Published: The read-only published data of the achievement.
	Published *AchievementConfigurationDetail `json:"published,omitempty"`

	// StepsToUnlock: Steps to unlock. Only applicable to incremental
	// achievements.
	StepsToUnlock int64 `json:"stepsToUnlock,omitempty"`

	// Token: The token for this resource.
	Token string `json:"token,omitempty"`
}

type AchievementConfigurationDetail struct {
	// Description: Localized strings for the achievement description.
	Description *LocalizedStringBundle `json:"description,omitempty"`

	// IconUrl: The icon url of this achievement. Writes to this field are
	// ignored.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#achievementConfigurationDetail.
	Kind string `json:"kind,omitempty"`

	// Name: Localized strings for the achievement name.
	Name *LocalizedStringBundle `json:"name,omitempty"`

	// PointValue: Point value for the achievement.
	PointValue int64 `json:"pointValue,omitempty"`

	// SortRank: The sort rank of this achievement. Writes to this field are
	// ignored.
	SortRank int64 `json:"sortRank,omitempty"`
}

type AchievementConfigurationListResponse struct {
	// Items: The achievement configurations.
	Items []*AchievementConfiguration `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#achievementConfigurationListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type GamesNumberAffixConfiguration struct {
	// Few: When the language requires special treatment of "small" numbers
	// (as with 2, 3, and 4 in Czech; or numbers ending 2, 3, or 4 but not
	// 12, 13, or 14 in Polish).
	Few *LocalizedStringBundle `json:"few,omitempty"`

	// Many: When the language requires special treatment of "large" numbers
	// (as with numbers ending 11-99 in Maltese).
	Many *LocalizedStringBundle `json:"many,omitempty"`

	// One: When the language requires special treatment of numbers like one
	// (as with the number 1 in English and most other languages; in
	// Russian, any number ending in 1 but not ending in 11 is in this
	// class).
	One *LocalizedStringBundle `json:"one,omitempty"`

	// Other: When the language does not require special treatment of the
	// given quantity (as with all numbers in Chinese, or 42 in English).
	Other *LocalizedStringBundle `json:"other,omitempty"`

	// Two: When the language requires special treatment of numbers like two
	// (as with 2 in Welsh, or 102 in Slovenian).
	Two *LocalizedStringBundle `json:"two,omitempty"`

	// Zero: When the language requires special treatment of the number 0
	// (as in Arabic).
	Zero *LocalizedStringBundle `json:"zero,omitempty"`
}

type GamesNumberFormatConfiguration struct {
	// CurrencyCode: The curreny code string. Only used for CURRENCY format
	// type.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// NumDecimalPlaces: The number of decimal places for number. Only used
	// for NUMERIC format type.
	NumDecimalPlaces int64 `json:"numDecimalPlaces,omitempty"`

	// NumberFormatType: The formatting for the number.
	// Possible values are:
	//
	// - "NUMERIC" - Numbers are formatted to have no digits or a fixed
	// number of digits after the decimal point according to locale. An
	// optional custom unit can be added.
	// - "TIME_DURATION" - Numbers are
	// formatted to hours, minutes and seconds.
	// - "CURRENCY" - Numbers are
	// formatted to currency according to locale.
	NumberFormatType string `json:"numberFormatType,omitempty"`

	// Suffix: An optional suffix for the NUMERIC format type. These strings
	// follow the same  plural rules as all Android string resources.
	Suffix *GamesNumberAffixConfiguration `json:"suffix,omitempty"`
}

type ImageConfiguration struct {
	// ImageType: The image type for the image.
	ImageType string `json:"imageType,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#imageConfiguration.
	Kind string `json:"kind,omitempty"`

	// ResourceId: The resource ID of resource which the image belongs to.
	ResourceId string `json:"resourceId,omitempty"`

	// Url: The url for this image.
	Url string `json:"url,omitempty"`
}

type LeaderboardConfiguration struct {
	// Draft: The draft data of the leaderboard.
	Draft *LeaderboardConfigurationDetail `json:"draft,omitempty"`

	// Id: The ID of the leaderboard.
	Id string `json:"id,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#leaderboardConfiguration.
	Kind string `json:"kind,omitempty"`

	// Published: The read-only published data of the leaderboard.
	Published *LeaderboardConfigurationDetail `json:"published,omitempty"`

	// ScoreMax: Maximum score that can be posted to this leaderboard.
	ScoreMax int64 `json:"scoreMax,omitempty,string"`

	// ScoreMin: Minimum score that can be posted to this leaderboard.
	ScoreMin int64 `json:"scoreMin,omitempty,string"`

	// ScoreOrder: The type of the leaderboard.
	// Possible values are:
	// -
	// "LARGER_IS_BETTER" - Larger scores posted are ranked higher.
	// -
	// "SMALLER_IS_BETTER" - Smaller scores posted are ranked higher.
	ScoreOrder string `json:"scoreOrder,omitempty"`

	// Token: The token for this resource.
	Token string `json:"token,omitempty"`
}

type LeaderboardConfigurationDetail struct {
	// IconUrl: The icon url of this leaderboard. Writes to this field are
	// ignored.
	IconUrl string `json:"iconUrl,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#leaderboardConfigurationDetail.
	Kind string `json:"kind,omitempty"`

	// Name: Localized strings for the leaderboard name.
	Name *LocalizedStringBundle `json:"name,omitempty"`

	// ScoreFormat: The score formatting for the leaderboard.
	ScoreFormat *GamesNumberFormatConfiguration `json:"scoreFormat,omitempty"`

	// SortRank: The sort rank of this leaderboard. Writes to this field are
	// ignored.
	SortRank int64 `json:"sortRank,omitempty"`
}

type LeaderboardConfigurationListResponse struct {
	// Items: The leaderboard configurations.
	Items []*LeaderboardConfiguration `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string games#leaderboardConfigurationListResponse.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type LocalizedString struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#localizedString.
	Kind string `json:"kind,omitempty"`

	// Locale: The locale string.
	Locale string `json:"locale,omitempty"`

	// Value: The string value.
	Value string `json:"value,omitempty"`
}

type LocalizedStringBundle struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesConfiguration#localizedStringBundle.
	Kind string `json:"kind,omitempty"`

	// Translations: The locale strings.
	Translations []*LocalizedString `json:"translations,omitempty"`
}

// method id "gamesConfiguration.achievementConfigurations.delete":

type AchievementConfigurationsDeleteCall struct {
	s             *Service
	achievementId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete the achievement configuration with the given ID.

func (r *AchievementConfigurationsService) Delete(achievementId string) *AchievementConfigurationsDeleteCall {
	return &AchievementConfigurationsDeleteCall{
		s:             r.s,
		achievementId: achievementId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/{achievementId}",
	}
}

func (c *AchievementConfigurationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"achievementId": c.achievementId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete the achievement configuration with the given ID.",
	//   "httpMethod": "DELETE",
	//   "id": "gamesConfiguration.achievementConfigurations.delete",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.get":

type AchievementConfigurationsGetCall struct {
	s             *Service
	achievementId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the metadata of the achievement configuration with the
// given ID.

func (r *AchievementConfigurationsService) Get(achievementId string) *AchievementConfigurationsGetCall {
	return &AchievementConfigurationsGetCall{
		s:             r.s,
		achievementId: achievementId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/{achievementId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsGetCall) Fields(s ...googleapi.Field) *AchievementConfigurationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementConfigurationsGetCall) Do() (*AchievementConfiguration, error) {
	var returnValue *AchievementConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"achievementId": c.achievementId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the metadata of the achievement configuration with the given ID.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.achievementConfigurations.get",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.insert":

type AchievementConfigurationsInsertCall struct {
	s                        *Service
	applicationId            string
	achievementconfiguration *AchievementConfiguration
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Insert: Insert a new achievement configuration in this application.

func (r *AchievementConfigurationsService) Insert(applicationId string, achievementconfiguration *AchievementConfiguration) *AchievementConfigurationsInsertCall {
	return &AchievementConfigurationsInsertCall{
		s:                        r.s,
		applicationId:            applicationId,
		achievementconfiguration: achievementconfiguration,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "applications/{applicationId}/achievements",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsInsertCall) Fields(s ...googleapi.Field) *AchievementConfigurationsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementConfigurationsInsertCall) Do() (*AchievementConfiguration, error) {
	var returnValue *AchievementConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"applicationId": c.applicationId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.achievementconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Insert a new achievement configuration in this application.",
	//   "httpMethod": "POST",
	//   "id": "gamesConfiguration.achievementConfigurations.insert",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/achievements",
	//   "request": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.list":

type AchievementConfigurationsListCall struct {
	s             *Service
	applicationId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns a list of the achievement configurations in this
// application.

func (r *AchievementConfigurationsService) List(applicationId string) *AchievementConfigurationsListCall {
	return &AchievementConfigurationsListCall{
		s:             r.s,
		applicationId: applicationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "applications/{applicationId}/achievements",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of resource configurations to return in the response, used for
// paging. For any response, the actual number of resources returned may
// be less than the specified maxResults.
func (c *AchievementConfigurationsListCall) MaxResults(maxResults int64) *AchievementConfigurationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *AchievementConfigurationsListCall) PageToken(pageToken string) *AchievementConfigurationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsListCall) Fields(s ...googleapi.Field) *AchievementConfigurationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementConfigurationsListCall) Do() (*AchievementConfigurationListResponse, error) {
	var returnValue *AchievementConfigurationListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"applicationId": c.applicationId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a list of the achievement configurations in this application.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.achievementConfigurations.list",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of resource configurations to return in the response, used for paging. For any response, the actual number of resources returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "200",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/achievements",
	//   "response": {
	//     "$ref": "AchievementConfigurationListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.patch":

type AchievementConfigurationsPatchCall struct {
	s                        *Service
	achievementId            string
	achievementconfiguration *AchievementConfiguration
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Patch: Update the metadata of the achievement configuration with the
// given ID. This method supports patch semantics.

func (r *AchievementConfigurationsService) Patch(achievementId string, achievementconfiguration *AchievementConfiguration) *AchievementConfigurationsPatchCall {
	return &AchievementConfigurationsPatchCall{
		s:                        r.s,
		achievementId:            achievementId,
		achievementconfiguration: achievementconfiguration,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "achievements/{achievementId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsPatchCall) Fields(s ...googleapi.Field) *AchievementConfigurationsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementConfigurationsPatchCall) Do() (*AchievementConfiguration, error) {
	var returnValue *AchievementConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"achievementId": c.achievementId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.achievementconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the metadata of the achievement configuration with the given ID. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "gamesConfiguration.achievementConfigurations.patch",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "request": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.achievementConfigurations.update":

type AchievementConfigurationsUpdateCall struct {
	s                        *Service
	achievementId            string
	achievementconfiguration *AchievementConfiguration
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Update: Update the metadata of the achievement configuration with the
// given ID.

func (r *AchievementConfigurationsService) Update(achievementId string, achievementconfiguration *AchievementConfiguration) *AchievementConfigurationsUpdateCall {
	return &AchievementConfigurationsUpdateCall{
		s:                        r.s,
		achievementId:            achievementId,
		achievementconfiguration: achievementconfiguration,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "achievements/{achievementId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementConfigurationsUpdateCall) Fields(s ...googleapi.Field) *AchievementConfigurationsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementConfigurationsUpdateCall) Do() (*AchievementConfiguration, error) {
	var returnValue *AchievementConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"achievementId": c.achievementId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.achievementconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the metadata of the achievement configuration with the given ID.",
	//   "httpMethod": "PUT",
	//   "id": "gamesConfiguration.achievementConfigurations.update",
	//   "parameterOrder": [
	//     "achievementId"
	//   ],
	//   "parameters": {
	//     "achievementId": {
	//       "description": "The ID of the achievement used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "achievements/{achievementId}",
	//   "request": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "response": {
	//     "$ref": "AchievementConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.imageConfigurations.upload":

type ImageConfigurationsUploadCall struct {
	s             *Service
	resourceId    string
	imageType     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Upload: Uploads an image for a resource with the given ID and image
// type.

func (r *ImageConfigurationsService) Upload(resourceId string, imageType string) *ImageConfigurationsUploadCall {
	return &ImageConfigurationsUploadCall{
		s:             r.s,
		resourceId:    resourceId,
		imageType:     imageType,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "images/{resourceId}/imageType/{imageType}",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *ImageConfigurationsUploadCall) Upload(ctx context.Context, u googleapi.UploadCaller) *ImageConfigurationsUploadCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *ImageConfigurationsUploadCall) Media(r io.Reader) *ImageConfigurationsUploadCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *ImageConfigurationsUploadCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ImageConfigurationsUploadCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *ImageConfigurationsUploadCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ImageConfigurationsUploadCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ImageConfigurationsUploadCall) Fields(s ...googleapi.Field) *ImageConfigurationsUploadCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ImageConfigurationsUploadCall) Do() (*ImageConfiguration, error) {
	var returnValue *ImageConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"resourceId": c.resourceId,
		"imageType":  c.imageType,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Uploads an image for a resource with the given ID and image type.",
	//   "httpMethod": "POST",
	//   "id": "gamesConfiguration.imageConfigurations.upload",
	//   "mediaUpload": {
	//     "accept": [
	//       "image/*"
	//     ],
	//     "maxSize": "15MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/games/v1configuration/images/{resourceId}/imageType/{imageType}"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "resourceId",
	//     "imageType"
	//   ],
	//   "parameters": {
	//     "imageType": {
	//       "description": "Selects which image in a resource for this method.",
	//       "enum": [
	//         "ACHIEVEMENT_ICON",
	//         "LEADERBOARD_ICON"
	//       ],
	//       "enumDescriptions": [
	//         "The icon image for an achievement resource.",
	//         "The icon image for a leaderboard resource."
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceId": {
	//       "description": "The ID of the resource used by this method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "images/{resourceId}/imageType/{imageType}",
	//   "response": {
	//     "$ref": "ImageConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.delete":

type LeaderboardConfigurationsDeleteCall struct {
	s             *Service
	leaderboardId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete the leaderboard configuration with the given ID.

func (r *LeaderboardConfigurationsService) Delete(leaderboardId string) *LeaderboardConfigurationsDeleteCall {
	return &LeaderboardConfigurationsDeleteCall{
		s:             r.s,
		leaderboardId: leaderboardId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "leaderboards/{leaderboardId}",
	}
}

func (c *LeaderboardConfigurationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"leaderboardId": c.leaderboardId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete the leaderboard configuration with the given ID.",
	//   "httpMethod": "DELETE",
	//   "id": "gamesConfiguration.leaderboardConfigurations.delete",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.get":

type LeaderboardConfigurationsGetCall struct {
	s             *Service
	leaderboardId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves the metadata of the leaderboard configuration with the
// given ID.

func (r *LeaderboardConfigurationsService) Get(leaderboardId string) *LeaderboardConfigurationsGetCall {
	return &LeaderboardConfigurationsGetCall{
		s:             r.s,
		leaderboardId: leaderboardId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "leaderboards/{leaderboardId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsGetCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LeaderboardConfigurationsGetCall) Do() (*LeaderboardConfiguration, error) {
	var returnValue *LeaderboardConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"leaderboardId": c.leaderboardId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the metadata of the leaderboard configuration with the given ID.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.leaderboardConfigurations.get",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.insert":

type LeaderboardConfigurationsInsertCall struct {
	s                        *Service
	applicationId            string
	leaderboardconfiguration *LeaderboardConfiguration
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Insert: Insert a new leaderboard configuration in this application.

func (r *LeaderboardConfigurationsService) Insert(applicationId string, leaderboardconfiguration *LeaderboardConfiguration) *LeaderboardConfigurationsInsertCall {
	return &LeaderboardConfigurationsInsertCall{
		s:                        r.s,
		applicationId:            applicationId,
		leaderboardconfiguration: leaderboardconfiguration,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "applications/{applicationId}/leaderboards",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsInsertCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LeaderboardConfigurationsInsertCall) Do() (*LeaderboardConfiguration, error) {
	var returnValue *LeaderboardConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"applicationId": c.applicationId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.leaderboardconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Insert a new leaderboard configuration in this application.",
	//   "httpMethod": "POST",
	//   "id": "gamesConfiguration.leaderboardConfigurations.insert",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/leaderboards",
	//   "request": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.list":

type LeaderboardConfigurationsListCall struct {
	s             *Service
	applicationId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns a list of the leaderboard configurations in this
// application.

func (r *LeaderboardConfigurationsService) List(applicationId string) *LeaderboardConfigurationsListCall {
	return &LeaderboardConfigurationsListCall{
		s:             r.s,
		applicationId: applicationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "applications/{applicationId}/leaderboards",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of resource configurations to return in the response, used for
// paging. For any response, the actual number of resources returned may
// be less than the specified maxResults.
func (c *LeaderboardConfigurationsListCall) MaxResults(maxResults int64) *LeaderboardConfigurationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *LeaderboardConfigurationsListCall) PageToken(pageToken string) *LeaderboardConfigurationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsListCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LeaderboardConfigurationsListCall) Do() (*LeaderboardConfigurationListResponse, error) {
	var returnValue *LeaderboardConfigurationListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"applicationId": c.applicationId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a list of the leaderboard configurations in this application.",
	//   "httpMethod": "GET",
	//   "id": "gamesConfiguration.leaderboardConfigurations.list",
	//   "parameterOrder": [
	//     "applicationId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of resource configurations to return in the response, used for paging. For any response, the actual number of resources returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "200",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/leaderboards",
	//   "response": {
	//     "$ref": "LeaderboardConfigurationListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.patch":

type LeaderboardConfigurationsPatchCall struct {
	s                        *Service
	leaderboardId            string
	leaderboardconfiguration *LeaderboardConfiguration
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Patch: Update the metadata of the leaderboard configuration with the
// given ID. This method supports patch semantics.

func (r *LeaderboardConfigurationsService) Patch(leaderboardId string, leaderboardconfiguration *LeaderboardConfiguration) *LeaderboardConfigurationsPatchCall {
	return &LeaderboardConfigurationsPatchCall{
		s:                        r.s,
		leaderboardId:            leaderboardId,
		leaderboardconfiguration: leaderboardconfiguration,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "leaderboards/{leaderboardId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsPatchCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LeaderboardConfigurationsPatchCall) Do() (*LeaderboardConfiguration, error) {
	var returnValue *LeaderboardConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"leaderboardId": c.leaderboardId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.leaderboardconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the metadata of the leaderboard configuration with the given ID. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "gamesConfiguration.leaderboardConfigurations.patch",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "request": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}

// method id "gamesConfiguration.leaderboardConfigurations.update":

type LeaderboardConfigurationsUpdateCall struct {
	s                        *Service
	leaderboardId            string
	leaderboardconfiguration *LeaderboardConfiguration
	caller_                  googleapi.Caller
	params_                  url.Values
	pathTemplate_            string
}

// Update: Update the metadata of the leaderboard configuration with the
// given ID.

func (r *LeaderboardConfigurationsService) Update(leaderboardId string, leaderboardconfiguration *LeaderboardConfiguration) *LeaderboardConfigurationsUpdateCall {
	return &LeaderboardConfigurationsUpdateCall{
		s:                        r.s,
		leaderboardId:            leaderboardId,
		leaderboardconfiguration: leaderboardconfiguration,
		caller_:                  googleapi.JSONCall{},
		params_:                  make(map[string][]string),
		pathTemplate_:            "leaderboards/{leaderboardId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeaderboardConfigurationsUpdateCall) Fields(s ...googleapi.Field) *LeaderboardConfigurationsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LeaderboardConfigurationsUpdateCall) Do() (*LeaderboardConfiguration, error) {
	var returnValue *LeaderboardConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"leaderboardId": c.leaderboardId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.leaderboardconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update the metadata of the leaderboard configuration with the given ID.",
	//   "httpMethod": "PUT",
	//   "id": "gamesConfiguration.leaderboardConfigurations.update",
	//   "parameterOrder": [
	//     "leaderboardId"
	//   ],
	//   "parameters": {
	//     "leaderboardId": {
	//       "description": "The ID of the leaderboard.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "leaderboards/{leaderboardId}",
	//   "request": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "response": {
	//     "$ref": "LeaderboardConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ]
	// }

}
