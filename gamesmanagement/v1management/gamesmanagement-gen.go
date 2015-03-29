// Package gamesmanagement provides access to the Google Play Game Services Management API.
//
// See https://developers.google.com/games/services
//
// Usage example:
//
//   import "github.com/jfcote87/api/gamesmanagement/v1management"
//   ...
//   gamesmanagementService, err := gamesmanagement.New(oauthHttpClient)
package gamesmanagement

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

const apiId = "gamesManagement:v1management"
const apiName = "gamesManagement"
const apiVersion = "v1management"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/games/v1management/"}

// OAuth2 scopes used by this API.
const (
	// Share your Google+ profile information and view and manage your game
	// activity
	GamesScope = "https://www.googleapis.com/auth/games"

	// Know your basic profile info and list of people in your circles.
	PlusLoginScope = "https://www.googleapis.com/auth/plus.login"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Achievements = NewAchievementsService(s)
	s.Applications = NewApplicationsService(s)
	s.Events = NewEventsService(s)
	s.Players = NewPlayersService(s)
	s.Quests = NewQuestsService(s)
	s.Rooms = NewRoomsService(s)
	s.Scores = NewScoresService(s)
	s.TurnBasedMatches = NewTurnBasedMatchesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Achievements *AchievementsService

	Applications *ApplicationsService

	Events *EventsService

	Players *PlayersService

	Quests *QuestsService

	Rooms *RoomsService

	Scores *ScoresService

	TurnBasedMatches *TurnBasedMatchesService
}

func NewAchievementsService(s *Service) *AchievementsService {
	rs := &AchievementsService{s: s}
	return rs
}

type AchievementsService struct {
	s *Service
}

func NewApplicationsService(s *Service) *ApplicationsService {
	rs := &ApplicationsService{s: s}
	return rs
}

type ApplicationsService struct {
	s *Service
}

func NewEventsService(s *Service) *EventsService {
	rs := &EventsService{s: s}
	return rs
}

type EventsService struct {
	s *Service
}

func NewPlayersService(s *Service) *PlayersService {
	rs := &PlayersService{s: s}
	return rs
}

type PlayersService struct {
	s *Service
}

func NewQuestsService(s *Service) *QuestsService {
	rs := &QuestsService{s: s}
	return rs
}

type QuestsService struct {
	s *Service
}

func NewRoomsService(s *Service) *RoomsService {
	rs := &RoomsService{s: s}
	return rs
}

type RoomsService struct {
	s *Service
}

func NewScoresService(s *Service) *ScoresService {
	rs := &ScoresService{s: s}
	return rs
}

type ScoresService struct {
	s *Service
}

func NewTurnBasedMatchesService(s *Service) *TurnBasedMatchesService {
	rs := &TurnBasedMatchesService{s: s}
	return rs
}

type TurnBasedMatchesService struct {
	s *Service
}

type AchievementResetAllResponse struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#achievementResetAllResponse.
	Kind string `json:"kind,omitempty"`

	// Results: The achievement reset results.
	Results []*AchievementResetResponse `json:"results,omitempty"`
}

type AchievementResetMultipleForAllRequest struct {
	// Achievement_ids: The IDs of achievements to reset.
	Achievement_ids []string `json:"achievement_ids,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string
	// gamesManagement#achievementResetMultipleForAllRequest.
	Kind string `json:"kind,omitempty"`
}

type AchievementResetResponse struct {
	// CurrentState: The current state of the achievement. This is the same
	// as the initial state of the achievement.
	// Possible values are:
	// -
	// "HIDDEN"- Achievement is hidden.
	// - "REVEALED" - Achievement is
	// revealed.
	// - "UNLOCKED" - Achievement is unlocked.
	CurrentState string `json:"currentState,omitempty"`

	// DefinitionId: The ID of an achievement for which player state has
	// been updated.
	DefinitionId string `json:"definitionId,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#achievementResetResponse.
	Kind string `json:"kind,omitempty"`

	// UpdateOccurred: Flag to indicate if the requested update actually
	// occurred.
	UpdateOccurred bool `json:"updateOccurred,omitempty"`
}

type EventsResetMultipleForAllRequest struct {
	// Event_ids: The IDs of events to reset.
	Event_ids []string `json:"event_ids,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#eventsResetMultipleForAllRequest.
	Kind string `json:"kind,omitempty"`
}

type GamesPlayedResource struct {
	// AutoMatched: True if the player was auto-matched with the currently
	// authenticated user.
	AutoMatched bool `json:"autoMatched,omitempty"`

	// TimeMillis: The last time the player played the game in milliseconds
	// since the epoch in UTC.
	TimeMillis int64 `json:"timeMillis,omitempty,string"`
}

type GamesPlayerExperienceInfoResource struct {
	// CurrentExperiencePoints: The current number of experience points for
	// the player.
	CurrentExperiencePoints int64 `json:"currentExperiencePoints,omitempty,string"`

	// CurrentLevel: The current level of the player.
	CurrentLevel *GamesPlayerLevelResource `json:"currentLevel,omitempty"`

	// LastLevelUpTimestampMillis: The timestamp when the player was leveled
	// up, in millis since Unix epoch UTC.
	LastLevelUpTimestampMillis int64 `json:"lastLevelUpTimestampMillis,omitempty,string"`

	// NextLevel: The next level of the player. If the current level is the
	// maximum level, this should be same as the current level.
	NextLevel *GamesPlayerLevelResource `json:"nextLevel,omitempty"`
}

type GamesPlayerLevelResource struct {
	// Level: The level for the user.
	Level int64 `json:"level,omitempty"`

	// MaxExperiencePoints: The maximum experience points for this level.
	MaxExperiencePoints int64 `json:"maxExperiencePoints,omitempty,string"`

	// MinExperiencePoints: The minimum experience points for this level.
	MinExperiencePoints int64 `json:"minExperiencePoints,omitempty,string"`
}

type HiddenPlayer struct {
	// HiddenTimeMillis: The time this player was hidden.
	HiddenTimeMillis int64 `json:"hiddenTimeMillis,omitempty,string"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#hiddenPlayer.
	Kind string `json:"kind,omitempty"`

	// Player: The player information.
	Player *Player `json:"player,omitempty"`
}

type HiddenPlayerList struct {
	// Items: The players.
	Items []*HiddenPlayer `json:"items,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#hiddenPlayerList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Player struct {
	// AvatarImageUrl: The base URL for the image that represents the
	// player.
	AvatarImageUrl string `json:"avatarImageUrl,omitempty"`

	// DisplayName: The name to display for the player.
	DisplayName string `json:"displayName,omitempty"`

	// ExperienceInfo: An object to represent Play Game experience
	// information for the player.
	ExperienceInfo *GamesPlayerExperienceInfoResource `json:"experienceInfo,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#player.
	Kind string `json:"kind,omitempty"`

	// LastPlayedWith: Details about the last time this player played a
	// multiplayer game with the currently authenticated player. Populated
	// for PLAYED_WITH player collection members.
	LastPlayedWith *GamesPlayedResource `json:"lastPlayedWith,omitempty"`

	// Name: An object representation of the individual components of the
	// player's name. For some players, these fields may not be present.
	Name *PlayerName `json:"name,omitempty"`

	// PlayerId: The ID of the player.
	PlayerId string `json:"playerId,omitempty"`

	// Title: The player's title rewarded for their game activities.
	Title string `json:"title,omitempty"`
}

type PlayerName struct {
	// FamilyName: The family name of this player. In some places, this is
	// known as the last name.
	FamilyName string `json:"familyName,omitempty"`

	// GivenName: The given name of this player. In some places, this is
	// known as the first name.
	GivenName string `json:"givenName,omitempty"`
}

type PlayerScoreResetAllResponse struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#playerScoreResetResponse.
	Kind string `json:"kind,omitempty"`

	// Results: The leaderboard reset results.
	Results []*PlayerScoreResetResponse `json:"results,omitempty"`
}

type PlayerScoreResetResponse struct {
	// DefinitionId: The ID of an leaderboard for which player state has
	// been updated.
	DefinitionId string `json:"definitionId,omitempty"`

	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#playerScoreResetResponse.
	Kind string `json:"kind,omitempty"`

	// ResetScoreTimeSpans: The time spans of the updated score.
	// Possible
	// values are:
	// - "ALL_TIME" - The score is an all-time score.
	// -
	// "WEEKLY" - The score is a weekly score.
	// - "DAILY" - The score is a
	// daily score.
	ResetScoreTimeSpans []string `json:"resetScoreTimeSpans,omitempty"`
}

type QuestsResetMultipleForAllRequest struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#questsResetMultipleForAllRequest.
	Kind string `json:"kind,omitempty"`

	// Quest_ids: The IDs of quests to reset.
	Quest_ids []string `json:"quest_ids,omitempty"`
}

type ScoresResetMultipleForAllRequest struct {
	// Kind: Uniquely identifies the type of this resource. Value is always
	// the fixed string gamesManagement#scoresResetMultipleForAllRequest.
	Kind string `json:"kind,omitempty"`

	// Leaderboard_ids: The IDs of leaderboards to reset.
	Leaderboard_ids []string `json:"leaderboard_ids,omitempty"`
}

// method id "gamesManagement.achievements.reset":

type AchievementsResetCall struct {
	s             *Service
	achievementId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reset: Resets the achievement with the given ID for the currently
// authenticated player. This method is only accessible to whitelisted
// tester accounts for your application.

func (r *AchievementsService) Reset(achievementId string) *AchievementsResetCall {
	return &AchievementsResetCall{
		s:             r.s,
		achievementId: achievementId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/{achievementId}/reset",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementsResetCall) Fields(s ...googleapi.Field) *AchievementsResetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementsResetCall) Do() (*AchievementResetResponse, error) {
	var returnValue *AchievementResetResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"achievementId": c.achievementId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets the achievement with the given ID for the currently authenticated player. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.achievements.reset",
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
	//   "path": "achievements/{achievementId}/reset",
	//   "response": {
	//     "$ref": "AchievementResetResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.achievements.resetAll":

type AchievementsResetAllCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAll: Resets all achievements for the currently authenticated
// player for your application. This method is only accessible to
// whitelisted tester accounts for your application.

func (r *AchievementsService) ResetAll() *AchievementsResetAllCall {
	return &AchievementsResetAllCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/reset",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AchievementsResetAllCall) Fields(s ...googleapi.Field) *AchievementsResetAllCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AchievementsResetAllCall) Do() (*AchievementResetAllResponse, error) {
	var returnValue *AchievementResetAllResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all achievements for the currently authenticated player for your application. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.achievements.resetAll",
	//   "path": "achievements/reset",
	//   "response": {
	//     "$ref": "AchievementResetAllResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.achievements.resetAllForAllPlayers":

type AchievementsResetAllForAllPlayersCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAllForAllPlayers: Resets all draft achievements for all players.
// This method is only available to user accounts for your developer
// console.

func (r *AchievementsService) ResetAllForAllPlayers() *AchievementsResetAllForAllPlayersCall {
	return &AchievementsResetAllForAllPlayersCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/resetAllForAllPlayers",
	}
}

func (c *AchievementsResetAllForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all draft achievements for all players. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.achievements.resetAllForAllPlayers",
	//   "path": "achievements/resetAllForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.achievements.resetForAllPlayers":

type AchievementsResetForAllPlayersCall struct {
	s             *Service
	achievementId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetForAllPlayers: Resets the achievement with the given ID for all
// players. This method is only available to user accounts for your
// developer console. Only draft achievements can be reset.

func (r *AchievementsService) ResetForAllPlayers(achievementId string) *AchievementsResetForAllPlayersCall {
	return &AchievementsResetForAllPlayersCall{
		s:             r.s,
		achievementId: achievementId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/{achievementId}/resetForAllPlayers",
	}
}

func (c *AchievementsResetForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"achievementId": c.achievementId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets the achievement with the given ID for all players. This method is only available to user accounts for your developer console. Only draft achievements can be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.achievements.resetForAllPlayers",
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
	//   "path": "achievements/{achievementId}/resetForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.achievements.resetMultipleForAllPlayers":

type AchievementsResetMultipleForAllPlayersCall struct {
	s                                     *Service
	achievementresetmultipleforallrequest *AchievementResetMultipleForAllRequest
	caller_                               googleapi.Caller
	params_                               url.Values
	pathTemplate_                         string
}

// ResetMultipleForAllPlayers: Resets achievements with the given IDs
// for all players. This method is only available to user accounts for
// your developer console. Only draft achievements may be reset.

func (r *AchievementsService) ResetMultipleForAllPlayers(achievementresetmultipleforallrequest *AchievementResetMultipleForAllRequest) *AchievementsResetMultipleForAllPlayersCall {
	return &AchievementsResetMultipleForAllPlayersCall{
		s: r.s,
		achievementresetmultipleforallrequest: achievementresetmultipleforallrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "achievements/resetMultipleForAllPlayers",
	}
}

func (c *AchievementsResetMultipleForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.achievementresetmultipleforallrequest,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets achievements with the given IDs for all players. This method is only available to user accounts for your developer console. Only draft achievements may be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.achievements.resetMultipleForAllPlayers",
	//   "path": "achievements/resetMultipleForAllPlayers",
	//   "request": {
	//     "$ref": "AchievementResetMultipleForAllRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.applications.listHidden":

type ApplicationsListHiddenCall struct {
	s             *Service
	applicationId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ListHidden: Get the list of players hidden from the given
// application. This method is only available to user accounts for your
// developer console.

func (r *ApplicationsService) ListHidden(applicationId string) *ApplicationsListHiddenCall {
	return &ApplicationsListHiddenCall{
		s:             r.s,
		applicationId: applicationId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "applications/{applicationId}/players/hidden",
	}
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of player resources to return in the response, used for
// paging. For any response, the actual number of player resources
// returned may be less than the specified maxResults.
func (c *ApplicationsListHiddenCall) MaxResults(maxResults int64) *ApplicationsListHiddenCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token returned
// by the previous request.
func (c *ApplicationsListHiddenCall) PageToken(pageToken string) *ApplicationsListHiddenCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ApplicationsListHiddenCall) Fields(s ...googleapi.Field) *ApplicationsListHiddenCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ApplicationsListHiddenCall) Do() (*HiddenPlayerList, error) {
	var returnValue *HiddenPlayerList
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
	//   "description": "Get the list of players hidden from the given application. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "GET",
	//   "id": "gamesManagement.applications.listHidden",
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
	//       "description": "The maximum number of player resources to return in the response, used for paging. For any response, the actual number of player resources returned may be less than the specified maxResults.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "15",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token returned by the previous request.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/players/hidden",
	//   "response": {
	//     "$ref": "HiddenPlayerList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.events.reset":

type EventsResetCall struct {
	s             *Service
	eventId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reset: Resets all player progress on the event with the given ID for
// the currently authenticated player. This method is only accessible to
// whitelisted tester accounts for your application. All quests for this
// player that use the event will also be reset.

func (r *EventsService) Reset(eventId string) *EventsResetCall {
	return &EventsResetCall{
		s:             r.s,
		eventId:       eventId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "events/{eventId}/reset",
	}
}

func (c *EventsResetCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"eventId": c.eventId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all player progress on the event with the given ID for the currently authenticated player. This method is only accessible to whitelisted tester accounts for your application. All quests for this player that use the event will also be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.events.reset",
	//   "parameterOrder": [
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "eventId": {
	//       "description": "The ID of the event.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "events/{eventId}/reset",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.events.resetAll":

type EventsResetAllCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAll: Resets all player progress on all events for the currently
// authenticated player. This method is only accessible to whitelisted
// tester accounts for your application. All quests for this player will
// also be reset.

func (r *EventsService) ResetAll() *EventsResetAllCall {
	return &EventsResetAllCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "events/reset",
	}
}

func (c *EventsResetAllCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all player progress on all events for the currently authenticated player. This method is only accessible to whitelisted tester accounts for your application. All quests for this player will also be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.events.resetAll",
	//   "path": "events/reset",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.events.resetAllForAllPlayers":

type EventsResetAllForAllPlayersCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAllForAllPlayers: Resets all draft events for all players. This
// method is only available to user accounts for your developer console.
// All quests that use any of these events will also be reset.

func (r *EventsService) ResetAllForAllPlayers() *EventsResetAllForAllPlayersCall {
	return &EventsResetAllForAllPlayersCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "events/resetAllForAllPlayers",
	}
}

func (c *EventsResetAllForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all draft events for all players. This method is only available to user accounts for your developer console. All quests that use any of these events will also be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.events.resetAllForAllPlayers",
	//   "path": "events/resetAllForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.events.resetForAllPlayers":

type EventsResetForAllPlayersCall struct {
	s             *Service
	eventId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetForAllPlayers: Resets the event with the given ID for all
// players. This method is only available to user accounts for your
// developer console. Only draft events can be reset. All quests that
// use the event will also be reset.

func (r *EventsService) ResetForAllPlayers(eventId string) *EventsResetForAllPlayersCall {
	return &EventsResetForAllPlayersCall{
		s:             r.s,
		eventId:       eventId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "events/{eventId}/resetForAllPlayers",
	}
}

func (c *EventsResetForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"eventId": c.eventId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets the event with the given ID for all players. This method is only available to user accounts for your developer console. Only draft events can be reset. All quests that use the event will also be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.events.resetForAllPlayers",
	//   "parameterOrder": [
	//     "eventId"
	//   ],
	//   "parameters": {
	//     "eventId": {
	//       "description": "The ID of the event.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "events/{eventId}/resetForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.events.resetMultipleForAllPlayers":

type EventsResetMultipleForAllPlayersCall struct {
	s                                *Service
	eventsresetmultipleforallrequest *EventsResetMultipleForAllRequest
	caller_                          googleapi.Caller
	params_                          url.Values
	pathTemplate_                    string
}

// ResetMultipleForAllPlayers: Resets events with the given IDs for all
// players. This method is only available to user accounts for your
// developer console. Only draft events may be reset. All quests that
// use any of the events will also be reset.

func (r *EventsService) ResetMultipleForAllPlayers(eventsresetmultipleforallrequest *EventsResetMultipleForAllRequest) *EventsResetMultipleForAllPlayersCall {
	return &EventsResetMultipleForAllPlayersCall{
		s: r.s,
		eventsresetmultipleforallrequest: eventsresetmultipleforallrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "events/resetMultipleForAllPlayers",
	}
}

func (c *EventsResetMultipleForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.eventsresetmultipleforallrequest,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets events with the given IDs for all players. This method is only available to user accounts for your developer console. Only draft events may be reset. All quests that use any of the events will also be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.events.resetMultipleForAllPlayers",
	//   "path": "events/resetMultipleForAllPlayers",
	//   "request": {
	//     "$ref": "EventsResetMultipleForAllRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.players.hide":

type PlayersHideCall struct {
	s             *Service
	applicationId string
	playerId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Hide: Hide the given player's leaderboard scores from the given
// application. This method is only available to user accounts for your
// developer console.

func (r *PlayersService) Hide(applicationId string, playerId string) *PlayersHideCall {
	return &PlayersHideCall{
		s:             r.s,
		applicationId: applicationId,
		playerId:      playerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "applications/{applicationId}/players/hidden/{playerId}",
	}
}

func (c *PlayersHideCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"applicationId": c.applicationId,
		"playerId":      c.playerId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Hide the given player's leaderboard scores from the given application. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.players.hide",
	//   "parameterOrder": [
	//     "applicationId",
	//     "playerId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "playerId": {
	//       "description": "A player ID. A value of me may be used in place of the authenticated player's ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/players/hidden/{playerId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.players.unhide":

type PlayersUnhideCall struct {
	s             *Service
	applicationId string
	playerId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Unhide: Unhide the given player's leaderboard scores from the given
// application. This method is only available to user accounts for your
// developer console.

func (r *PlayersService) Unhide(applicationId string, playerId string) *PlayersUnhideCall {
	return &PlayersUnhideCall{
		s:             r.s,
		applicationId: applicationId,
		playerId:      playerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "applications/{applicationId}/players/hidden/{playerId}",
	}
}

func (c *PlayersUnhideCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"applicationId": c.applicationId,
		"playerId":      c.playerId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Unhide the given player's leaderboard scores from the given application. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "DELETE",
	//   "id": "gamesManagement.players.unhide",
	//   "parameterOrder": [
	//     "applicationId",
	//     "playerId"
	//   ],
	//   "parameters": {
	//     "applicationId": {
	//       "description": "The application ID from the Google Play developer console.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "playerId": {
	//       "description": "A player ID. A value of me may be used in place of the authenticated player's ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "applications/{applicationId}/players/hidden/{playerId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.quests.reset":

type QuestsResetCall struct {
	s             *Service
	questId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reset: Resets all player progress on the quest with the given ID for
// the currently authenticated player. This method is only accessible to
// whitelisted tester accounts for your application.

func (r *QuestsService) Reset(questId string) *QuestsResetCall {
	return &QuestsResetCall{
		s:             r.s,
		questId:       questId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "quests/{questId}/reset",
	}
}

func (c *QuestsResetCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"questId": c.questId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all player progress on the quest with the given ID for the currently authenticated player. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.quests.reset",
	//   "parameterOrder": [
	//     "questId"
	//   ],
	//   "parameters": {
	//     "questId": {
	//       "description": "The ID of the quest.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "quests/{questId}/reset",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.quests.resetAll":

type QuestsResetAllCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAll: Resets all player progress on all quests for the currently
// authenticated player. This method is only accessible to whitelisted
// tester accounts for your application.

func (r *QuestsService) ResetAll() *QuestsResetAllCall {
	return &QuestsResetAllCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "quests/reset",
	}
}

func (c *QuestsResetAllCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all player progress on all quests for the currently authenticated player. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.quests.resetAll",
	//   "path": "quests/reset",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.quests.resetAllForAllPlayers":

type QuestsResetAllForAllPlayersCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAllForAllPlayers: Resets all draft quests for all players. This
// method is only available to user accounts for your developer console.

func (r *QuestsService) ResetAllForAllPlayers() *QuestsResetAllForAllPlayersCall {
	return &QuestsResetAllForAllPlayersCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "quests/resetAllForAllPlayers",
	}
}

func (c *QuestsResetAllForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all draft quests for all players. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.quests.resetAllForAllPlayers",
	//   "path": "quests/resetAllForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.quests.resetForAllPlayers":

type QuestsResetForAllPlayersCall struct {
	s             *Service
	questId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetForAllPlayers: Resets all player progress on the quest with the
// given ID for all players. This method is only available to user
// accounts for your developer console. Only draft quests can be reset.

func (r *QuestsService) ResetForAllPlayers(questId string) *QuestsResetForAllPlayersCall {
	return &QuestsResetForAllPlayersCall{
		s:             r.s,
		questId:       questId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "quests/{questId}/resetForAllPlayers",
	}
}

func (c *QuestsResetForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"questId": c.questId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all player progress on the quest with the given ID for all players. This method is only available to user accounts for your developer console. Only draft quests can be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.quests.resetForAllPlayers",
	//   "parameterOrder": [
	//     "questId"
	//   ],
	//   "parameters": {
	//     "questId": {
	//       "description": "The ID of the quest.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "quests/{questId}/resetForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.quests.resetMultipleForAllPlayers":

type QuestsResetMultipleForAllPlayersCall struct {
	s                                *Service
	questsresetmultipleforallrequest *QuestsResetMultipleForAllRequest
	caller_                          googleapi.Caller
	params_                          url.Values
	pathTemplate_                    string
}

// ResetMultipleForAllPlayers: Resets quests with the given IDs for all
// players. This method is only available to user accounts for your
// developer console. Only draft quests may be reset.

func (r *QuestsService) ResetMultipleForAllPlayers(questsresetmultipleforallrequest *QuestsResetMultipleForAllRequest) *QuestsResetMultipleForAllPlayersCall {
	return &QuestsResetMultipleForAllPlayersCall{
		s: r.s,
		questsresetmultipleforallrequest: questsresetmultipleforallrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "quests/resetMultipleForAllPlayers",
	}
}

func (c *QuestsResetMultipleForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.questsresetmultipleforallrequest,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets quests with the given IDs for all players. This method is only available to user accounts for your developer console. Only draft quests may be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.quests.resetMultipleForAllPlayers",
	//   "path": "quests/resetMultipleForAllPlayers",
	//   "request": {
	//     "$ref": "QuestsResetMultipleForAllRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.rooms.reset":

type RoomsResetCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reset: Reset all rooms for the currently authenticated player for
// your application. This method is only accessible to whitelisted
// tester accounts for your application.

func (r *RoomsService) Reset() *RoomsResetCall {
	return &RoomsResetCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "rooms/reset",
	}
}

func (c *RoomsResetCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Reset all rooms for the currently authenticated player for your application. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.rooms.reset",
	//   "path": "rooms/reset",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.rooms.resetForAllPlayers":

type RoomsResetForAllPlayersCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetForAllPlayers: Deletes rooms where the only room participants
// are from whitelisted tester accounts for your application. This
// method is only available to user accounts for your developer console.

func (r *RoomsService) ResetForAllPlayers() *RoomsResetForAllPlayersCall {
	return &RoomsResetForAllPlayersCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "rooms/resetForAllPlayers",
	}
}

func (c *RoomsResetForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes rooms where the only room participants are from whitelisted tester accounts for your application. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.rooms.resetForAllPlayers",
	//   "path": "rooms/resetForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.scores.reset":

type ScoresResetCall struct {
	s             *Service
	leaderboardId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reset: Resets scores for the leaderboard with the given ID for the
// currently authenticated player. This method is only accessible to
// whitelisted tester accounts for your application.

func (r *ScoresService) Reset(leaderboardId string) *ScoresResetCall {
	return &ScoresResetCall{
		s:             r.s,
		leaderboardId: leaderboardId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "leaderboards/{leaderboardId}/scores/reset",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ScoresResetCall) Fields(s ...googleapi.Field) *ScoresResetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ScoresResetCall) Do() (*PlayerScoreResetResponse, error) {
	var returnValue *PlayerScoreResetResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"leaderboardId": c.leaderboardId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets scores for the leaderboard with the given ID for the currently authenticated player. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.scores.reset",
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
	//   "path": "leaderboards/{leaderboardId}/scores/reset",
	//   "response": {
	//     "$ref": "PlayerScoreResetResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.scores.resetAll":

type ScoresResetAllCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAll: Resets all scores for all leaderboards for the currently
// authenticated players. This method is only accessible to whitelisted
// tester accounts for your application.

func (r *ScoresService) ResetAll() *ScoresResetAllCall {
	return &ScoresResetAllCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "scores/reset",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ScoresResetAllCall) Fields(s ...googleapi.Field) *ScoresResetAllCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ScoresResetAllCall) Do() (*PlayerScoreResetAllResponse, error) {
	var returnValue *PlayerScoreResetAllResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets all scores for all leaderboards for the currently authenticated players. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.scores.resetAll",
	//   "path": "scores/reset",
	//   "response": {
	//     "$ref": "PlayerScoreResetAllResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.scores.resetAllForAllPlayers":

type ScoresResetAllForAllPlayersCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetAllForAllPlayers: Resets scores for all draft leaderboards for
// all players. This method is only available to user accounts for your
// developer console.

func (r *ScoresService) ResetAllForAllPlayers() *ScoresResetAllForAllPlayersCall {
	return &ScoresResetAllForAllPlayersCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "scores/resetAllForAllPlayers",
	}
}

func (c *ScoresResetAllForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets scores for all draft leaderboards for all players. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.scores.resetAllForAllPlayers",
	//   "path": "scores/resetAllForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.scores.resetForAllPlayers":

type ScoresResetForAllPlayersCall struct {
	s             *Service
	leaderboardId string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetForAllPlayers: Resets scores for the leaderboard with the given
// ID for all players. This method is only available to user accounts
// for your developer console. Only draft leaderboards can be reset.

func (r *ScoresService) ResetForAllPlayers(leaderboardId string) *ScoresResetForAllPlayersCall {
	return &ScoresResetForAllPlayersCall{
		s:             r.s,
		leaderboardId: leaderboardId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "leaderboards/{leaderboardId}/scores/resetForAllPlayers",
	}
}

func (c *ScoresResetForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"leaderboardId": c.leaderboardId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets scores for the leaderboard with the given ID for all players. This method is only available to user accounts for your developer console. Only draft leaderboards can be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.scores.resetForAllPlayers",
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
	//   "path": "leaderboards/{leaderboardId}/scores/resetForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.scores.resetMultipleForAllPlayers":

type ScoresResetMultipleForAllPlayersCall struct {
	s                                *Service
	scoresresetmultipleforallrequest *ScoresResetMultipleForAllRequest
	caller_                          googleapi.Caller
	params_                          url.Values
	pathTemplate_                    string
}

// ResetMultipleForAllPlayers: Resets scores for the leaderboards with
// the given IDs for all players. This method is only available to user
// accounts for your developer console. Only draft leaderboards may be
// reset.

func (r *ScoresService) ResetMultipleForAllPlayers(scoresresetmultipleforallrequest *ScoresResetMultipleForAllRequest) *ScoresResetMultipleForAllPlayersCall {
	return &ScoresResetMultipleForAllPlayersCall{
		s: r.s,
		scoresresetmultipleforallrequest: scoresresetmultipleforallrequest,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "scores/resetMultipleForAllPlayers",
	}
}

func (c *ScoresResetMultipleForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.scoresresetmultipleforallrequest,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Resets scores for the leaderboards with the given IDs for all players. This method is only available to user accounts for your developer console. Only draft leaderboards may be reset.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.scores.resetMultipleForAllPlayers",
	//   "path": "scores/resetMultipleForAllPlayers",
	//   "request": {
	//     "$ref": "ScoresResetMultipleForAllRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.turnBasedMatches.reset":

type TurnBasedMatchesResetCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Reset: Reset all turn-based match data for a user. This method is
// only accessible to whitelisted tester accounts for your application.

func (r *TurnBasedMatchesService) Reset() *TurnBasedMatchesResetCall {
	return &TurnBasedMatchesResetCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "turnbasedmatches/reset",
	}
}

func (c *TurnBasedMatchesResetCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Reset all turn-based match data for a user. This method is only accessible to whitelisted tester accounts for your application.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.turnBasedMatches.reset",
	//   "path": "turnbasedmatches/reset",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}

// method id "gamesManagement.turnBasedMatches.resetForAllPlayers":

type TurnBasedMatchesResetForAllPlayersCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// ResetForAllPlayers: Deletes turn-based matches where the only match
// participants are from whitelisted tester accounts for your
// application. This method is only available to user accounts for your
// developer console.

func (r *TurnBasedMatchesService) ResetForAllPlayers() *TurnBasedMatchesResetForAllPlayersCall {
	return &TurnBasedMatchesResetForAllPlayersCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "turnbasedmatches/resetForAllPlayers",
	}
}

func (c *TurnBasedMatchesResetForAllPlayersCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes turn-based matches where the only match participants are from whitelisted tester accounts for your application. This method is only available to user accounts for your developer console.",
	//   "httpMethod": "POST",
	//   "id": "gamesManagement.turnBasedMatches.resetForAllPlayers",
	//   "path": "turnbasedmatches/resetForAllPlayers",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/games",
	//     "https://www.googleapis.com/auth/plus.login"
	//   ]
	// }

}
