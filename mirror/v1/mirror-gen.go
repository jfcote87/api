// Package mirror provides access to the Google Mirror API.
//
// See https://developers.google.com/glass
//
// Usage example:
//
//   import "github.com/jfcote87/api/mirror/v1"
//   ...
//   mirrorService, err := mirror.New(oauthHttpClient)
package mirror

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

const apiId = "mirror:v1"
const apiName = "mirror"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/mirror/v1/"}

// OAuth2 scopes used by this API.
const (
	// View your location
	GlassLocationScope = "https://www.googleapis.com/auth/glass.location"

	// View and manage your Glass timeline
	GlassTimelineScope = "https://www.googleapis.com/auth/glass.timeline"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Accounts = NewAccountsService(s)
	s.Contacts = NewContactsService(s)
	s.Locations = NewLocationsService(s)
	s.Settings = NewSettingsService(s)
	s.Subscriptions = NewSubscriptionsService(s)
	s.Timeline = NewTimelineService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Accounts *AccountsService

	Contacts *ContactsService

	Locations *LocationsService

	Settings *SettingsService

	Subscriptions *SubscriptionsService

	Timeline *TimelineService
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	return rs
}

type AccountsService struct {
	s *Service
}

func NewContactsService(s *Service) *ContactsService {
	rs := &ContactsService{s: s}
	return rs
}

type ContactsService struct {
	s *Service
}

func NewLocationsService(s *Service) *LocationsService {
	rs := &LocationsService{s: s}
	return rs
}

type LocationsService struct {
	s *Service
}

func NewSettingsService(s *Service) *SettingsService {
	rs := &SettingsService{s: s}
	return rs
}

type SettingsService struct {
	s *Service
}

func NewSubscriptionsService(s *Service) *SubscriptionsService {
	rs := &SubscriptionsService{s: s}
	return rs
}

type SubscriptionsService struct {
	s *Service
}

func NewTimelineService(s *Service) *TimelineService {
	rs := &TimelineService{s: s}
	rs.Attachments = NewTimelineAttachmentsService(s)
	return rs
}

type TimelineService struct {
	s *Service

	Attachments *TimelineAttachmentsService
}

func NewTimelineAttachmentsService(s *Service) *TimelineAttachmentsService {
	rs := &TimelineAttachmentsService{s: s}
	return rs
}

type TimelineAttachmentsService struct {
	s *Service
}

type Account struct {
	AuthTokens []*AuthToken `json:"authTokens,omitempty"`

	Features []string `json:"features,omitempty"`

	Password string `json:"password,omitempty"`

	UserData []*UserData `json:"userData,omitempty"`
}

type Attachment struct {
	// ContentType: The MIME type of the attachment.
	ContentType string `json:"contentType,omitempty"`

	// ContentUrl: The URL for the content.
	ContentUrl string `json:"contentUrl,omitempty"`

	// Id: The ID of the attachment.
	Id string `json:"id,omitempty"`

	// IsProcessingContent: Indicates that the contentUrl is not available
	// because the attachment content is still being processed. If the
	// caller wishes to retrieve the content, it should try again later.
	IsProcessingContent bool `json:"isProcessingContent,omitempty"`
}

type AttachmentsListResponse struct {
	// Items: The list of attachments.
	Items []*Attachment `json:"items,omitempty"`

	// Kind: The type of resource. This is always mirror#attachmentsList.
	Kind string `json:"kind,omitempty"`
}

type AuthToken struct {
	AuthToken string `json:"authToken,omitempty"`

	Type string `json:"type,omitempty"`
}

type Command struct {
	// Type: The type of operation this command corresponds to. Allowed
	// values are:
	// - TAKE_A_NOTE - Shares a timeline item with the
	// transcription of user speech from the "Take a note" voice menu
	// command.
	// - POST_AN_UPDATE - Shares a timeline item with the
	// transcription of user speech from the "Post an update" voice menu
	// command.
	Type string `json:"type,omitempty"`
}

type Contact struct {
	// AcceptCommands: A list of voice menu commands that a contact can
	// handle. Glass shows up to three contacts for each voice menu command.
	// If there are more than that, the three contacts with the highest
	// priority are shown for that particular command.
	AcceptCommands []*Command `json:"acceptCommands,omitempty"`

	// AcceptTypes: A list of MIME types that a contact supports. The
	// contact will be shown to the user if any of its acceptTypes matches
	// any of the types of the attachments on the item. If no acceptTypes
	// are given, the contact will be shown for all items.
	AcceptTypes []string `json:"acceptTypes,omitempty"`

	// DisplayName: The name to display for this contact.
	DisplayName string `json:"displayName,omitempty"`

	// Id: An ID for this contact. This is generated by the application and
	// is treated as an opaque token.
	Id string `json:"id,omitempty"`

	// ImageUrls: Set of image URLs to display for a contact. Most contacts
	// will have a single image, but a "group" contact may include up to 8
	// image URLs and they will be resized and cropped into a mosaic on the
	// client.
	ImageUrls []string `json:"imageUrls,omitempty"`

	// Kind: The type of resource. This is always mirror#contact.
	Kind string `json:"kind,omitempty"`

	// PhoneNumber: Primary phone number for the contact. This can be a
	// fully-qualified number, with country calling code and area code, or a
	// local number.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Priority: Priority for the contact to determine ordering in a list of
	// contacts. Contacts with higher priorities will be shown before ones
	// with lower priorities.
	Priority int64 `json:"priority,omitempty"`

	// SharingFeatures: A list of sharing features that a contact can
	// handle. Allowed values are:
	// - ADD_CAPTION
	SharingFeatures []string `json:"sharingFeatures,omitempty"`

	// Source: The ID of the application that created this contact. This is
	// populated by the API
	Source string `json:"source,omitempty"`

	// SpeakableName: Name of this contact as it should be pronounced. If
	// this contact's name must be spoken as part of a voice disambiguation
	// menu, this name is used as the expected pronunciation. This is useful
	// for contact names with unpronounceable characters or whose display
	// spelling is otherwise not phonetic.
	SpeakableName string `json:"speakableName,omitempty"`

	// Type: The type for this contact. This is used for sorting in UIs.
	// Allowed values are:
	// - INDIVIDUAL - Represents a single person. This
	// is the default.
	// - GROUP - Represents more than a single person.
	Type string `json:"type,omitempty"`
}

type ContactsListResponse struct {
	// Items: Contact list.
	Items []*Contact `json:"items,omitempty"`

	// Kind: The type of resource. This is always mirror#contacts.
	Kind string `json:"kind,omitempty"`
}

type Location struct {
	// Accuracy: The accuracy of the location fix in meters.
	Accuracy float64 `json:"accuracy,omitempty"`

	// Address: The full address of the location.
	Address string `json:"address,omitempty"`

	// DisplayName: The name to be displayed. This may be a business name or
	// a user-defined place, such as "Home".
	DisplayName string `json:"displayName,omitempty"`

	// Id: The ID of the location.
	Id string `json:"id,omitempty"`

	// Kind: The type of resource. This is always mirror#location.
	Kind string `json:"kind,omitempty"`

	// Latitude: The latitude, in degrees.
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude, in degrees.
	Longitude float64 `json:"longitude,omitempty"`

	// Timestamp: The time at which this location was captured, formatted
	// according to RFC 3339.
	Timestamp string `json:"timestamp,omitempty"`
}

type LocationsListResponse struct {
	// Items: The list of locations.
	Items []*Location `json:"items,omitempty"`

	// Kind: The type of resource. This is always mirror#locationsList.
	Kind string `json:"kind,omitempty"`
}

type MenuItem struct {
	// Action: Controls the behavior when the user picks the menu option.
	// Allowed values are:
	// - CUSTOM - Custom action set by the service.
	// When the user selects this menuItem, the API triggers a notification
	// to your callbackUrl with the userActions.type set to CUSTOM and the
	// userActions.payload set to the ID of this menu item. This is the
	// default value.
	// - Built-in actions:
	// - REPLY - Initiate a reply to
	// the timeline item using the voice recording UI. The creator attribute
	// must be set in the timeline item for this menu to be available.
	// -
	// REPLY_ALL - Same behavior as REPLY. The original timeline item's
	// recipients will be added to the reply item.
	// - DELETE - Delete the
	// timeline item.
	// - SHARE - Share the timeline item with the available
	// contacts.
	// - READ_ALOUD - Read the timeline item's speakableText
	// aloud; if this field is not set, read the text field; if none of
	// those fields are set, this menu item is ignored.
	// - GET_MEDIA_INPUT -
	// Allow users to provide media payloads to Glassware from a menu item
	// (currently, only transcribed text from voice input is supported).
	// Subscribe to notifications when users invoke this menu item to
	// receive the timeline item ID. Retrieve the media from the timeline
	// item in the payload property.
	// - VOICE_CALL - Initiate a phone call
	// using the timeline item's creator.phoneNumber attribute as recipient.
	//
	// - NAVIGATE - Navigate to the timeline item's location.
	// -
	// TOGGLE_PINNED - Toggle the isPinned state of the timeline item.
	// -
	// OPEN_URI - Open the payload of the menu item in the browser.
	// -
	// PLAY_VIDEO - Open the payload of the menu item in the Glass video
	// player.
	// - SEND_MESSAGE - Initiate sending a message to the timeline
	// item's creator:
	// - If the creator.phoneNumber is set and Glass is
	// connected to an Android phone, the message is an SMS.
	// - Otherwise,
	// if the creator.email is set, the message is an email.
	Action string `json:"action,omitempty"`

	// Contextual_command: The ContextualMenus.Command associated with this
	// MenuItem (e.g. READ_ALOUD). The voice label for this command will be
	// displayed in the voice menu and the touch label will be displayed in
	// the touch menu. Note that the default menu value's display name will
	// be overriden if you specify this property. Values that do not
	// correspond to a ContextualMenus.Command name will be ignored.
	Contextual_command string `json:"contextual_command,omitempty"`

	// Id: The ID for this menu item. This is generated by the application
	// and is treated as an opaque token.
	Id string `json:"id,omitempty"`

	// Payload: A generic payload whose meaning changes depending on this
	// MenuItem's action.
	// - When the action is OPEN_URI, the payload is
	// the URL of the website to view.
	// - When the action is PLAY_VIDEO, the
	// payload is the streaming URL of the video
	// - When the action is
	// GET_MEDIA_INPUT, the payload is the text transcription of a user's
	// speech input
	Payload string `json:"payload,omitempty"`

	// RemoveWhenSelected: If set to true on a CUSTOM menu item, that item
	// will be removed from the menu after it is selected.
	RemoveWhenSelected bool `json:"removeWhenSelected,omitempty"`

	// Values: For CUSTOM items, a list of values controlling the appearance
	// of the menu item in each of its states. A value for the DEFAULT state
	// must be provided. If the PENDING or CONFIRMED states are missing,
	// they will not be shown.
	Values []*MenuValue `json:"values,omitempty"`
}

type MenuValue struct {
	// DisplayName: The name to display for the menu item. If you specify
	// this property for a built-in menu item, the default contextual voice
	// command for that menu item is not shown.
	DisplayName string `json:"displayName,omitempty"`

	// IconUrl: URL of an icon to display with the menu item.
	IconUrl string `json:"iconUrl,omitempty"`

	// State: The state that this value applies to. Allowed values are:
	// -
	// DEFAULT - Default value shown when displayed in the menuItems list.
	//
	// - PENDING - Value shown when the menuItem has been selected by the
	// user but can still be cancelled.
	// - CONFIRMED - Value shown when the
	// menuItem has been selected by the user and can no longer be
	// cancelled.
	State string `json:"state,omitempty"`
}

type Notification struct {
	// Collection: The collection that generated the notification.
	Collection string `json:"collection,omitempty"`

	// ItemId: The ID of the item that generated the notification.
	ItemId string `json:"itemId,omitempty"`

	// Operation: The type of operation that generated the notification.
	Operation string `json:"operation,omitempty"`

	// UserActions: A list of actions taken by the user that triggered the
	// notification.
	UserActions []*UserAction `json:"userActions,omitempty"`

	// UserToken: The user token provided by the service when it subscribed
	// for notifications.
	UserToken string `json:"userToken,omitempty"`

	// VerifyToken: The secret verify token provided by the service when it
	// subscribed for notifications.
	VerifyToken string `json:"verifyToken,omitempty"`
}

type NotificationConfig struct {
	// DeliveryTime: The time at which the notification should be delivered.
	DeliveryTime string `json:"deliveryTime,omitempty"`

	// Level: Describes how important the notification is. Allowed values
	// are:
	// - DEFAULT - Notifications of default importance. A chime will
	// be played to alert users.
	Level string `json:"level,omitempty"`
}

type Setting struct {
	// Id: The setting's ID. The following IDs are valid:
	// - locale - The
	// key to the user’s language/locale (BCP 47 identifier) that
	// Glassware should use to render localized content.
	// - timezone - The
	// key to the user’s current time zone region as defined in the tz
	// database. Example: America/Los_Angeles.
	Id string `json:"id,omitempty"`

	// Kind: The type of resource. This is always mirror#setting.
	Kind string `json:"kind,omitempty"`

	// Value: The setting value, as a string.
	Value string `json:"value,omitempty"`
}

type Subscription struct {
	// CallbackUrl: The URL where notifications should be delivered (must
	// start with https://).
	CallbackUrl string `json:"callbackUrl,omitempty"`

	// Collection: The collection to subscribe to. Allowed values are:
	// -
	// timeline - Changes in the timeline including insertion, deletion, and
	// updates.
	// - locations - Location updates.
	// - settings - Settings
	// updates.
	Collection string `json:"collection,omitempty"`

	// Id: The ID of the subscription.
	Id string `json:"id,omitempty"`

	// Kind: The type of resource. This is always mirror#subscription.
	Kind string `json:"kind,omitempty"`

	// Notification: Container object for notifications. This is not
	// populated in the Subscription resource.
	Notification *Notification `json:"notification,omitempty"`

	// Operation: A list of operations that should be subscribed to. An
	// empty list indicates that all operations on the collection should be
	// subscribed to. Allowed values are:
	// - UPDATE - The item has been
	// updated.
	// - INSERT - A new item has been inserted.
	// - DELETE - The
	// item has been deleted.
	// - MENU_ACTION - A custom menu item has been
	// triggered by the user.
	Operation []string `json:"operation,omitempty"`

	// Updated: The time at which this subscription was last modified,
	// formatted according to RFC 3339.
	Updated string `json:"updated,omitempty"`

	// UserToken: An opaque token sent to the subscriber in notifications so
	// that it can determine the ID of the user.
	UserToken string `json:"userToken,omitempty"`

	// VerifyToken: A secret token sent to the subscriber in notifications
	// so that it can verify that the notification was generated by Google.
	VerifyToken string `json:"verifyToken,omitempty"`
}

type SubscriptionsListResponse struct {
	// Items: The list of subscriptions.
	Items []*Subscription `json:"items,omitempty"`

	// Kind: The type of resource. This is always mirror#subscriptionsList.
	Kind string `json:"kind,omitempty"`
}

type TimelineItem struct {
	// Attachments: A list of media attachments associated with this item.
	// As a convenience, you can refer to attachments in your HTML payloads
	// with the attachment or cid scheme. For example:
	// - attachment: <img
	// src="attachment:attachment_index"> where attachment_index is the
	// 0-based index of this array.
	// - cid: <img src="cid:attachment_id">
	// where attachment_id is the ID of the attachment.
	Attachments []*Attachment `json:"attachments,omitempty"`

	// BundleId: The bundle ID for this item. Services can specify a
	// bundleId to group many items together. They appear under a single
	// top-level item on the device.
	BundleId string `json:"bundleId,omitempty"`

	// CanonicalUrl: A canonical URL pointing to the canonical/high quality
	// version of the data represented by the timeline item.
	CanonicalUrl string `json:"canonicalUrl,omitempty"`

	// Created: The time at which this item was created, formatted according
	// to RFC 3339.
	Created string `json:"created,omitempty"`

	// Creator: The user or group that created this item.
	Creator *Contact `json:"creator,omitempty"`

	// DisplayTime: The time that should be displayed when this item is
	// viewed in the timeline, formatted according to RFC 3339. This user's
	// timeline is sorted chronologically on display time, so this will also
	// determine where the item is displayed in the timeline. If not set by
	// the service, the display time defaults to the updated time.
	DisplayTime string `json:"displayTime,omitempty"`

	// Etag: ETag for this item.
	Etag string `json:"etag,omitempty"`

	// Html: HTML content for this item. If both text and html are provided
	// for an item, the html will be rendered in the timeline.
	// Allowed HTML
	// elements - You can use these elements in your timeline cards.
	//
	// -
	// Headers: h1, h2, h3, h4, h5, h6
	// - Images: img
	// - Lists: li, ol, ul
	//
	// - HTML5 semantics: article, aside, details, figure, figcaption,
	// footer, header, nav, section, summary, time
	// - Structural:
	// blockquote, br, div, hr, p, span
	// - Style: b, big, center, em, i, u,
	// s, small, strike, strong, style, sub, sup
	// - Tables: table, tbody,
	// td, tfoot, th, thead, tr
	// Blocked HTML elements: These elements and
	// their contents are removed from HTML payloads.
	//
	// - Document headers:
	// head, title
	// - Embeds: audio, embed, object, source, video
	// - Frames:
	// frame, frameset
	// - Scripting: applet, script
	// Other elements: Any
	// elements that aren't listed are removed, but their contents are
	// preserved.
	Html string `json:"html,omitempty"`

	// Id: The ID of the timeline item. This is unique within a user's
	// timeline.
	Id string `json:"id,omitempty"`

	// InReplyTo: If this item was generated as a reply to another item,
	// this field will be set to the ID of the item being replied to. This
	// can be used to attach a reply to the appropriate conversation or
	// post.
	InReplyTo string `json:"inReplyTo,omitempty"`

	// IsBundleCover: Whether this item is a bundle cover.
	//
	// If an item is
	// marked as a bundle cover, it will be the entry point to the bundle of
	// items that have the same bundleId as that item. It will be shown only
	// on the main timeline — not within the opened bundle.
	//
	// On the main
	// timeline, items that are shown are:
	// - Items that have isBundleCover
	// set to true
	// - Items that do not have a bundleId  In a bundle
	// sub-timeline, items that are shown are:
	// - Items that have the
	// bundleId in question AND isBundleCover set to false
	IsBundleCover bool `json:"isBundleCover,omitempty"`

	// IsDeleted: When true, indicates this item is deleted, and only the ID
	// property is set.
	IsDeleted bool `json:"isDeleted,omitempty"`

	// IsPinned: When true, indicates this item is pinned, which means it's
	// grouped alongside "active" items like navigation and hangouts, on the
	// opposite side of the home screen from historical (non-pinned)
	// timeline items. You can allow the user to toggle the value of this
	// property with the TOGGLE_PINNED built-in menu item.
	IsPinned bool `json:"isPinned,omitempty"`

	// Kind: The type of resource. This is always mirror#timelineItem.
	Kind string `json:"kind,omitempty"`

	// Location: The geographic location associated with this item.
	Location *Location `json:"location,omitempty"`

	// MenuItems: A list of menu items that will be presented to the user
	// when this item is selected in the timeline.
	MenuItems []*MenuItem `json:"menuItems,omitempty"`

	// Notification: Controls how notifications for this item are presented
	// on the device. If this is missing, no notification will be generated.
	Notification *NotificationConfig `json:"notification,omitempty"`

	// PinScore: For pinned items, this determines the order in which the
	// item is displayed in the timeline, with a higher score appearing
	// closer to the clock. Note: setting this field is currently not
	// supported.
	PinScore int64 `json:"pinScore,omitempty"`

	// Recipients: A list of users or groups that this item has been shared
	// with.
	Recipients []*Contact `json:"recipients,omitempty"`

	// SelfLink: A URL that can be used to retrieve this item.
	SelfLink string `json:"selfLink,omitempty"`

	// SourceItemId: Opaque string you can use to map a timeline item to
	// data in your own service.
	SourceItemId string `json:"sourceItemId,omitempty"`

	// SpeakableText: The speakable version of the content of this item.
	// Along with the READ_ALOUD menu item, use this field to provide text
	// that would be clearer when read aloud, or to provide extended
	// information to what is displayed visually on Glass.
	//
	// Glassware should
	// also specify the speakableType field, which will be spoken before
	// this text in cases where the additional context is useful, for
	// example when the user requests that the item be read aloud following
	// a notification.
	SpeakableText string `json:"speakableText,omitempty"`

	// SpeakableType: A speakable description of the type of this item. This
	// will be announced to the user prior to reading the content of the
	// item in cases where the additional context is useful, for example
	// when the user requests that the item be read aloud following a
	// notification.
	//
	// This should be a short, simple noun phrase such as
	// "Email", "Text message", or "Daily Planet News Update".
	//
	// Glassware
	// are encouraged to populate this field for every timeline item, even
	// if the item does not contain speakableText or text so that the user
	// can learn the type of the item without looking at the screen.
	SpeakableType string `json:"speakableType,omitempty"`

	// Text: Text content of this item.
	Text string `json:"text,omitempty"`

	// Title: The title of this item.
	Title string `json:"title,omitempty"`

	// Updated: The time at which this item was last modified, formatted
	// according to RFC 3339.
	Updated string `json:"updated,omitempty"`
}

type TimelineListResponse struct {
	// Items: Items in the timeline.
	Items []*TimelineItem `json:"items,omitempty"`

	// Kind: The type of resource. This is always mirror#timeline.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The next page token. Provide this as the pageToken
	// parameter in the request to retrieve the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type UserAction struct {
	// Payload: An optional payload for the action.
	//
	// For actions of type
	// CUSTOM, this is the ID of the custom menu item that was selected.
	Payload string `json:"payload,omitempty"`

	// Type: The type of action. The value of this can be:
	// - SHARE - the
	// user shared an item.
	// - REPLY - the user replied to an item.
	// -
	// REPLY_ALL - the user replied to all recipients of an item.
	// - CUSTOM
	// - the user selected a custom menu item on the timeline item.
	// -
	// DELETE - the user deleted the item.
	// - PIN - the user pinned the
	// item.
	// - UNPIN - the user unpinned the item.
	// - LAUNCH - the user
	// initiated a voice command.  In the future, additional types may be
	// added. UserActions with unrecognized types should be ignored.
	Type string `json:"type,omitempty"`
}

type UserData struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

// method id "mirror.accounts.insert":

type AccountsInsertCall struct {
	s             *Service
	userToken     string
	accountType   string
	accountName   string
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new account for a user

func (r *AccountsService) Insert(userToken string, accountType string, accountName string, account *Account) *AccountsInsertCall {
	return &AccountsInsertCall{
		s:             r.s,
		userToken:     userToken,
		accountType:   accountType,
		accountName:   accountName,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{userToken}/{accountType}/{accountName}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsInsertCall) Fields(s ...googleapi.Field) *AccountsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsInsertCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userToken":   c.userToken,
		"accountType": c.accountType,
		"accountName": c.accountName,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.account,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new account for a user",
	//   "httpMethod": "POST",
	//   "id": "mirror.accounts.insert",
	//   "parameterOrder": [
	//     "userToken",
	//     "accountType",
	//     "accountName"
	//   ],
	//   "parameters": {
	//     "accountName": {
	//       "description": "The name of the account to be passed to the Android Account Manager.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "accountType": {
	//       "description": "Account type to be passed to Android Account Manager.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userToken": {
	//       "description": "The ID for the user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{userToken}/{accountType}/{accountName}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   }
	// }

}

// method id "mirror.contacts.delete":

type ContactsDeleteCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a contact.

func (r *ContactsService) Delete(id string) *ContactsDeleteCall {
	return &ContactsDeleteCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "contacts/{id}",
	}
}

func (c *ContactsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a contact.",
	//   "httpMethod": "DELETE",
	//   "id": "mirror.contacts.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the contact.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "contacts/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.contacts.get":

type ContactsGetCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a single contact by ID.

func (r *ContactsService) Get(id string) *ContactsGetCall {
	return &ContactsGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "contacts/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContactsGetCall) Fields(s ...googleapi.Field) *ContactsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContactsGetCall) Do() (*Contact, error) {
	var returnValue *Contact
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a single contact by ID.",
	//   "httpMethod": "GET",
	//   "id": "mirror.contacts.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the contact.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "contacts/{id}",
	//   "response": {
	//     "$ref": "Contact"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.contacts.insert":

type ContactsInsertCall struct {
	s             *Service
	contact       *Contact
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new contact.

func (r *ContactsService) Insert(contact *Contact) *ContactsInsertCall {
	return &ContactsInsertCall{
		s:             r.s,
		contact:       contact,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "contacts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContactsInsertCall) Fields(s ...googleapi.Field) *ContactsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContactsInsertCall) Do() (*Contact, error) {
	var returnValue *Contact
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.contact,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new contact.",
	//   "httpMethod": "POST",
	//   "id": "mirror.contacts.insert",
	//   "path": "contacts",
	//   "request": {
	//     "$ref": "Contact"
	//   },
	//   "response": {
	//     "$ref": "Contact"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.contacts.list":

type ContactsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of contacts for the authenticated user.

func (r *ContactsService) List() *ContactsListCall {
	return &ContactsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "contacts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContactsListCall) Fields(s ...googleapi.Field) *ContactsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContactsListCall) Do() (*ContactsListResponse, error) {
	var returnValue *ContactsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of contacts for the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "mirror.contacts.list",
	//   "path": "contacts",
	//   "response": {
	//     "$ref": "ContactsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.contacts.patch":

type ContactsPatchCall struct {
	s             *Service
	id            string
	contact       *Contact
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a contact in place. This method supports patch
// semantics.

func (r *ContactsService) Patch(id string, contact *Contact) *ContactsPatchCall {
	return &ContactsPatchCall{
		s:             r.s,
		id:            id,
		contact:       contact,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "contacts/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContactsPatchCall) Fields(s ...googleapi.Field) *ContactsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContactsPatchCall) Do() (*Contact, error) {
	var returnValue *Contact
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.contact,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a contact in place. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "mirror.contacts.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the contact.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "contacts/{id}",
	//   "request": {
	//     "$ref": "Contact"
	//   },
	//   "response": {
	//     "$ref": "Contact"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.contacts.update":

type ContactsUpdateCall struct {
	s             *Service
	id            string
	contact       *Contact
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a contact in place.

func (r *ContactsService) Update(id string, contact *Contact) *ContactsUpdateCall {
	return &ContactsUpdateCall{
		s:             r.s,
		id:            id,
		contact:       contact,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "contacts/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContactsUpdateCall) Fields(s ...googleapi.Field) *ContactsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContactsUpdateCall) Do() (*Contact, error) {
	var returnValue *Contact
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.contact,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a contact in place.",
	//   "httpMethod": "PUT",
	//   "id": "mirror.contacts.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the contact.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "contacts/{id}",
	//   "request": {
	//     "$ref": "Contact"
	//   },
	//   "response": {
	//     "$ref": "Contact"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.locations.get":

type LocationsGetCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a single location by ID.

func (r *LocationsService) Get(id string) *LocationsGetCall {
	return &LocationsGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "locations/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LocationsGetCall) Fields(s ...googleapi.Field) *LocationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LocationsGetCall) Do() (*Location, error) {
	var returnValue *Location
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a single location by ID.",
	//   "httpMethod": "GET",
	//   "id": "mirror.locations.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the location or latest for the last known location.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "locations/{id}",
	//   "response": {
	//     "$ref": "Location"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.locations.list":

type LocationsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of locations for the user.

func (r *LocationsService) List() *LocationsListCall {
	return &LocationsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "locations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LocationsListCall) Fields(s ...googleapi.Field) *LocationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LocationsListCall) Do() (*LocationsListResponse, error) {
	var returnValue *LocationsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of locations for the user.",
	//   "httpMethod": "GET",
	//   "id": "mirror.locations.list",
	//   "path": "locations",
	//   "response": {
	//     "$ref": "LocationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.settings.get":

type SettingsGetCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a single setting by ID.

func (r *SettingsService) Get(id string) *SettingsGetCall {
	return &SettingsGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "settings/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SettingsGetCall) Fields(s ...googleapi.Field) *SettingsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SettingsGetCall) Do() (*Setting, error) {
	var returnValue *Setting
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a single setting by ID.",
	//   "httpMethod": "GET",
	//   "id": "mirror.settings.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the setting. The following IDs are valid: \n- locale - The key to the user’s language/locale (BCP 47 identifier) that Glassware should use to render localized content. \n- timezone - The key to the user’s current time zone region as defined in the tz database. Example: America/Los_Angeles.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "settings/{id}",
	//   "response": {
	//     "$ref": "Setting"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.subscriptions.delete":

type SubscriptionsDeleteCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a subscription.

func (r *SubscriptionsService) Delete(id string) *SubscriptionsDeleteCall {
	return &SubscriptionsDeleteCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "subscriptions/{id}",
	}
}

func (c *SubscriptionsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a subscription.",
	//   "httpMethod": "DELETE",
	//   "id": "mirror.subscriptions.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the subscription.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.subscriptions.insert":

type SubscriptionsInsertCall struct {
	s             *Service
	subscription  *Subscription
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a new subscription.

func (r *SubscriptionsService) Insert(subscription *Subscription) *SubscriptionsInsertCall {
	return &SubscriptionsInsertCall{
		s:             r.s,
		subscription:  subscription,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "subscriptions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsInsertCall) Fields(s ...googleapi.Field) *SubscriptionsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsInsertCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.subscription,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a new subscription.",
	//   "httpMethod": "POST",
	//   "id": "mirror.subscriptions.insert",
	//   "path": "subscriptions",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.subscriptions.list":

type SubscriptionsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of subscriptions for the authenticated user
// and service.

func (r *SubscriptionsService) List() *SubscriptionsListCall {
	return &SubscriptionsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "subscriptions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsListCall) Fields(s ...googleapi.Field) *SubscriptionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsListCall) Do() (*SubscriptionsListResponse, error) {
	var returnValue *SubscriptionsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of subscriptions for the authenticated user and service.",
	//   "httpMethod": "GET",
	//   "id": "mirror.subscriptions.list",
	//   "path": "subscriptions",
	//   "response": {
	//     "$ref": "SubscriptionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.subscriptions.update":

type SubscriptionsUpdateCall struct {
	s             *Service
	id            string
	subscription  *Subscription
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing subscription in place.

func (r *SubscriptionsService) Update(id string, subscription *Subscription) *SubscriptionsUpdateCall {
	return &SubscriptionsUpdateCall{
		s:             r.s,
		id:            id,
		subscription:  subscription,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "subscriptions/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsUpdateCall) Fields(s ...googleapi.Field) *SubscriptionsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubscriptionsUpdateCall) Do() (*Subscription, error) {
	var returnValue *Subscription
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.subscription,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing subscription in place.",
	//   "httpMethod": "PUT",
	//   "id": "mirror.subscriptions.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the subscription.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions/{id}",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.timeline.delete":

type TimelineDeleteCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a timeline item.

func (r *TimelineService) Delete(id string) *TimelineDeleteCall {
	return &TimelineDeleteCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{id}",
	}
}

func (c *TimelineDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a timeline item.",
	//   "httpMethod": "DELETE",
	//   "id": "mirror.timeline.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the timeline item.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.timeline.get":

type TimelineGetCall struct {
	s             *Service
	id            string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a single timeline item by ID.

func (r *TimelineService) Get(id string) *TimelineGetCall {
	return &TimelineGetCall{
		s:             r.s,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineGetCall) Fields(s ...googleapi.Field) *TimelineGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineGetCall) Do() (*TimelineItem, error) {
	var returnValue *TimelineItem
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a single timeline item by ID.",
	//   "httpMethod": "GET",
	//   "id": "mirror.timeline.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the timeline item.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{id}",
	//   "response": {
	//     "$ref": "TimelineItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.timeline.insert":

type TimelineInsertCall struct {
	s             *Service
	timelineitem  *TimelineItem
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Inserts a new item into the timeline.

func (r *TimelineService) Insert(timelineitem *TimelineItem) *TimelineInsertCall {
	return &TimelineInsertCall{
		s:             r.s,
		timelineitem:  timelineitem,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *TimelineInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *TimelineInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/mirror/v1/timeline"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/mirror/v1/timeline"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *TimelineInsertCall) Media(r io.Reader) *TimelineInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/mirror/v1/timeline"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TimelineInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TimelineInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/mirror/v1/timeline"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TimelineInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TimelineInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineInsertCall) Fields(s ...googleapi.Field) *TimelineInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineInsertCall) Do() (*TimelineItem, error) {
	var returnValue *TimelineItem
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.timelineitem,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Inserts a new item into the timeline.",
	//   "httpMethod": "POST",
	//   "id": "mirror.timeline.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "audio/*",
	//       "image/*",
	//       "video/*"
	//     ],
	//     "maxSize": "10MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/mirror/v1/timeline"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/mirror/v1/timeline"
	//       }
	//     }
	//   },
	//   "path": "timeline",
	//   "request": {
	//     "$ref": "TimelineItem"
	//   },
	//   "response": {
	//     "$ref": "TimelineItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "mirror.timeline.list":

type TimelineListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of timeline items for the authenticated user.

func (r *TimelineService) List() *TimelineListCall {
	return &TimelineListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline",
	}
}

// BundleId sets the optional parameter "bundleId": If provided, only
// items with the given bundleId will be returned.
func (c *TimelineListCall) BundleId(bundleId string) *TimelineListCall {
	c.params_.Set("bundleId", fmt.Sprintf("%v", bundleId))
	return c
}

// IncludeDeleted sets the optional parameter "includeDeleted": If true,
// tombstone records for deleted items will be returned.
func (c *TimelineListCall) IncludeDeleted(includeDeleted bool) *TimelineListCall {
	c.params_.Set("includeDeleted", fmt.Sprintf("%v", includeDeleted))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in the response, used for paging.
func (c *TimelineListCall) MaxResults(maxResults int64) *TimelineListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Controls the order in
// which timeline items are returned.
func (c *TimelineListCall) OrderBy(orderBy string) *TimelineListCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token for the page
// of results to return.
func (c *TimelineListCall) PageToken(pageToken string) *TimelineListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// PinnedOnly sets the optional parameter "pinnedOnly": If true, only
// pinned items will be returned.
func (c *TimelineListCall) PinnedOnly(pinnedOnly bool) *TimelineListCall {
	c.params_.Set("pinnedOnly", fmt.Sprintf("%v", pinnedOnly))
	return c
}

// SourceItemId sets the optional parameter "sourceItemId": If provided,
// only items with the given sourceItemId will be returned.
func (c *TimelineListCall) SourceItemId(sourceItemId string) *TimelineListCall {
	c.params_.Set("sourceItemId", fmt.Sprintf("%v", sourceItemId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineListCall) Fields(s ...googleapi.Field) *TimelineListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineListCall) Do() (*TimelineListResponse, error) {
	var returnValue *TimelineListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of timeline items for the authenticated user.",
	//   "httpMethod": "GET",
	//   "id": "mirror.timeline.list",
	//   "parameters": {
	//     "bundleId": {
	//       "description": "If provided, only items with the given bundleId will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "includeDeleted": {
	//       "description": "If true, tombstone records for deleted items will be returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in the response, used for paging.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Controls the order in which timeline items are returned.",
	//       "enum": [
	//         "displayTime",
	//         "writeTime"
	//       ],
	//       "enumDescriptions": [
	//         "Results will be ordered by displayTime (default). This is the same ordering as is used in the timeline on the device.",
	//         "Results will be ordered by the time at which they were last written to the data store."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token for the page of results to return.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pinnedOnly": {
	//       "description": "If true, only pinned items will be returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sourceItemId": {
	//       "description": "If provided, only items with the given sourceItemId will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline",
	//   "response": {
	//     "$ref": "TimelineListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.timeline.patch":

type TimelinePatchCall struct {
	s             *Service
	id            string
	timelineitem  *TimelineItem
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a timeline item in place. This method supports patch
// semantics.

func (r *TimelineService) Patch(id string, timelineitem *TimelineItem) *TimelinePatchCall {
	return &TimelinePatchCall{
		s:             r.s,
		id:            id,
		timelineitem:  timelineitem,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelinePatchCall) Fields(s ...googleapi.Field) *TimelinePatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelinePatchCall) Do() (*TimelineItem, error) {
	var returnValue *TimelineItem
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.timelineitem,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a timeline item in place. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "mirror.timeline.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the timeline item.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{id}",
	//   "request": {
	//     "$ref": "TimelineItem"
	//   },
	//   "response": {
	//     "$ref": "TimelineItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.timeline.update":

type TimelineUpdateCall struct {
	s             *Service
	id            string
	timelineitem  *TimelineItem
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Update: Updates a timeline item in place.

func (r *TimelineService) Update(id string, timelineitem *TimelineItem) *TimelineUpdateCall {
	return &TimelineUpdateCall{
		s:             r.s,
		id:            id,
		timelineitem:  timelineitem,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{id}",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *TimelineUpdateCall) Upload(ctx context.Context, u googleapi.UploadCaller) *TimelineUpdateCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/mirror/v1/timeline/{id}"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/mirror/v1/timeline/{id}"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *TimelineUpdateCall) Media(r io.Reader) *TimelineUpdateCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/mirror/v1/timeline/{id}"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TimelineUpdateCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TimelineUpdateCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/mirror/v1/timeline/{id}"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TimelineUpdateCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TimelineUpdateCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineUpdateCall) Fields(s ...googleapi.Field) *TimelineUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineUpdateCall) Do() (*TimelineItem, error) {
	var returnValue *TimelineItem
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"id": c.id,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.timelineitem,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Updates a timeline item in place.",
	//   "httpMethod": "PUT",
	//   "id": "mirror.timeline.update",
	//   "mediaUpload": {
	//     "accept": [
	//       "audio/*",
	//       "image/*",
	//       "video/*"
	//     ],
	//     "maxSize": "10MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/mirror/v1/timeline/{id}"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/mirror/v1/timeline/{id}"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the timeline item.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{id}",
	//   "request": {
	//     "$ref": "TimelineItem"
	//   },
	//   "response": {
	//     "$ref": "TimelineItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.location",
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "mirror.timeline.attachments.delete":

type TimelineAttachmentsDeleteCall struct {
	s             *Service
	itemId        string
	attachmentId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an attachment from a timeline item.

func (r *TimelineAttachmentsService) Delete(itemId string, attachmentId string) *TimelineAttachmentsDeleteCall {
	return &TimelineAttachmentsDeleteCall{
		s:             r.s,
		itemId:        itemId,
		attachmentId:  attachmentId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{itemId}/attachments/{attachmentId}",
	}
}

func (c *TimelineAttachmentsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"itemId":       c.itemId,
		"attachmentId": c.attachmentId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an attachment from a timeline item.",
	//   "httpMethod": "DELETE",
	//   "id": "mirror.timeline.attachments.delete",
	//   "parameterOrder": [
	//     "itemId",
	//     "attachmentId"
	//   ],
	//   "parameters": {
	//     "attachmentId": {
	//       "description": "The ID of the attachment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "itemId": {
	//       "description": "The ID of the timeline item the attachment belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{itemId}/attachments/{attachmentId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}

// method id "mirror.timeline.attachments.get":

type TimelineAttachmentsGetCall struct {
	s             *Service
	itemId        string
	attachmentId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves an attachment on a timeline item by item ID and
// attachment ID.

func (r *TimelineAttachmentsService) Get(itemId string, attachmentId string) *TimelineAttachmentsGetCall {
	return &TimelineAttachmentsGetCall{
		s:             r.s,
		itemId:        itemId,
		attachmentId:  attachmentId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{itemId}/attachments/{attachmentId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineAttachmentsGetCall) Fields(s ...googleapi.Field) *TimelineAttachmentsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineAttachmentsGetCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"itemId":       c.itemId,
		"attachmentId": c.attachmentId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *TimelineAttachmentsGetCall) Do() (*Attachment, error) {
	var returnValue *Attachment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"itemId":       c.itemId,
		"attachmentId": c.attachmentId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves an attachment on a timeline item by item ID and attachment ID.",
	//   "httpMethod": "GET",
	//   "id": "mirror.timeline.attachments.get",
	//   "parameterOrder": [
	//     "itemId",
	//     "attachmentId"
	//   ],
	//   "parameters": {
	//     "attachmentId": {
	//       "description": "The ID of the attachment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "itemId": {
	//       "description": "The ID of the timeline item the attachment belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{itemId}/attachments/{attachmentId}",
	//   "response": {
	//     "$ref": "Attachment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "mirror.timeline.attachments.insert":

type TimelineAttachmentsInsertCall struct {
	s             *Service
	itemId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
	context_      context.Context
	callback_     googleapi.ProgressUpdater
}

// Insert: Adds a new attachment to a timeline item.

func (r *TimelineAttachmentsService) Insert(itemId string) *TimelineAttachmentsInsertCall {
	return &TimelineAttachmentsInsertCall{
		s:             r.s,
		itemId:        itemId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{itemId}/attachments",
		context_:      context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *TimelineAttachmentsInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *TimelineAttachmentsInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/mirror/v1/timeline/{itemId}/attachments"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/mirror/v1/timeline/{itemId}/attachments"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *TimelineAttachmentsInsertCall) Media(r io.Reader) *TimelineAttachmentsInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/mirror/v1/timeline/{itemId}/attachments"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *TimelineAttachmentsInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TimelineAttachmentsInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/mirror/v1/timeline/{itemId}/attachments"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *TimelineAttachmentsInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TimelineAttachmentsInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineAttachmentsInsertCall) Fields(s ...googleapi.Field) *TimelineAttachmentsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineAttachmentsInsertCall) Do() (*Attachment, error) {
	var returnValue *Attachment
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"itemId": c.itemId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Adds a new attachment to a timeline item.",
	//   "httpMethod": "POST",
	//   "id": "mirror.timeline.attachments.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "audio/*",
	//       "image/*",
	//       "video/*"
	//     ],
	//     "maxSize": "10MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/mirror/v1/timeline/{itemId}/attachments"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/mirror/v1/timeline/{itemId}/attachments"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "itemId"
	//   ],
	//   "parameters": {
	//     "itemId": {
	//       "description": "The ID of the timeline item the attachment belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{itemId}/attachments",
	//   "response": {
	//     "$ref": "Attachment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "mirror.timeline.attachments.list":

type TimelineAttachmentsListCall struct {
	s             *Service
	itemId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns a list of attachments for a timeline item.

func (r *TimelineAttachmentsService) List(itemId string) *TimelineAttachmentsListCall {
	return &TimelineAttachmentsListCall{
		s:             r.s,
		itemId:        itemId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "timeline/{itemId}/attachments",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TimelineAttachmentsListCall) Fields(s ...googleapi.Field) *TimelineAttachmentsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TimelineAttachmentsListCall) Do() (*AttachmentsListResponse, error) {
	var returnValue *AttachmentsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"itemId": c.itemId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns a list of attachments for a timeline item.",
	//   "httpMethod": "GET",
	//   "id": "mirror.timeline.attachments.list",
	//   "parameterOrder": [
	//     "itemId"
	//   ],
	//   "parameters": {
	//     "itemId": {
	//       "description": "The ID of the timeline item whose attachments should be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "timeline/{itemId}/attachments",
	//   "response": {
	//     "$ref": "AttachmentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/glass.timeline"
	//   ]
	// }

}
