// Package tagmanager provides access to the Tag Manager API.
//
// See https://developers.google.com/tag-manager/api/v1/
//
// Usage example:
//
//   import "github.com/jfcote87/api/tagmanager/v1"
//   ...
//   tagmanagerService, err := tagmanager.New(oauthHttpClient)
package tagmanager

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

const apiId = "tagmanager:v1"
const apiName = "tagmanager"
const apiVersion = "v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/tagmanager/v1/"}

// OAuth2 scopes used by this API.
const (
	// Delete your Google Tag Manager containers
	TagmanagerDeleteContainersScope = "https://www.googleapis.com/auth/tagmanager.delete.containers"

	// Manage your Google Tag Manager containers
	TagmanagerEditContainersScope = "https://www.googleapis.com/auth/tagmanager.edit.containers"

	// Manage your Google Tag Manager container versions
	TagmanagerEditContainerversionsScope = "https://www.googleapis.com/auth/tagmanager.edit.containerversions"

	// Manage your Google Tag Manager accounts
	TagmanagerManageAccountsScope = "https://www.googleapis.com/auth/tagmanager.manage.accounts"

	// Manage user permissions of your Google Tag Manager data
	TagmanagerManageUsersScope = "https://www.googleapis.com/auth/tagmanager.manage.users"

	// Publish your Google Tag Manager containers
	TagmanagerPublishScope = "https://www.googleapis.com/auth/tagmanager.publish"

	// View your Google Tag Manager containers
	TagmanagerReadonlyScope = "https://www.googleapis.com/auth/tagmanager.readonly"
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
	rs.Containers = NewAccountsContainersService(s)
	rs.Permissions = NewAccountsPermissionsService(s)
	return rs
}

type AccountsService struct {
	s *Service

	Containers *AccountsContainersService

	Permissions *AccountsPermissionsService
}

func NewAccountsContainersService(s *Service) *AccountsContainersService {
	rs := &AccountsContainersService{s: s}
	rs.Macros = NewAccountsContainersMacrosService(s)
	rs.Rules = NewAccountsContainersRulesService(s)
	rs.Tags = NewAccountsContainersTagsService(s)
	rs.Triggers = NewAccountsContainersTriggersService(s)
	rs.Variables = NewAccountsContainersVariablesService(s)
	rs.Versions = NewAccountsContainersVersionsService(s)
	return rs
}

type AccountsContainersService struct {
	s *Service

	Macros *AccountsContainersMacrosService

	Rules *AccountsContainersRulesService

	Tags *AccountsContainersTagsService

	Triggers *AccountsContainersTriggersService

	Variables *AccountsContainersVariablesService

	Versions *AccountsContainersVersionsService
}

func NewAccountsContainersMacrosService(s *Service) *AccountsContainersMacrosService {
	rs := &AccountsContainersMacrosService{s: s}
	return rs
}

type AccountsContainersMacrosService struct {
	s *Service
}

func NewAccountsContainersRulesService(s *Service) *AccountsContainersRulesService {
	rs := &AccountsContainersRulesService{s: s}
	return rs
}

type AccountsContainersRulesService struct {
	s *Service
}

func NewAccountsContainersTagsService(s *Service) *AccountsContainersTagsService {
	rs := &AccountsContainersTagsService{s: s}
	return rs
}

type AccountsContainersTagsService struct {
	s *Service
}

func NewAccountsContainersTriggersService(s *Service) *AccountsContainersTriggersService {
	rs := &AccountsContainersTriggersService{s: s}
	return rs
}

type AccountsContainersTriggersService struct {
	s *Service
}

func NewAccountsContainersVariablesService(s *Service) *AccountsContainersVariablesService {
	rs := &AccountsContainersVariablesService{s: s}
	return rs
}

type AccountsContainersVariablesService struct {
	s *Service
}

func NewAccountsContainersVersionsService(s *Service) *AccountsContainersVersionsService {
	rs := &AccountsContainersVersionsService{s: s}
	return rs
}

type AccountsContainersVersionsService struct {
	s *Service
}

func NewAccountsPermissionsService(s *Service) *AccountsPermissionsService {
	rs := &AccountsPermissionsService{s: s}
	return rs
}

type AccountsPermissionsService struct {
	s *Service
}

type Account struct {
	// AccountId: The Account ID uniquely identifies the GTM Account.
	AccountId string `json:"accountId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Account as computed at
	// storage time. This value is recomputed whenever the account is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Account display name.
	Name string `json:"name,omitempty"`

	// ShareData: Whether the account shares data anonymously with Google
	// and others.
	ShareData bool `json:"shareData,omitempty"`
}

type AccountAccess struct {
	// Permission: List of Account permissions. Valid account permissions
	// are read and manage.
	Permission []string `json:"permission,omitempty"`
}

type Condition struct {
	// Parameter: A list of named parameters (key/value), depending on the
	// condition's type. Notes:
	// - For binary operators, include parameters
	// named arg0 and arg1 for specifying the left and right operands,
	// respectively.
	// - At this time, the left operand (arg0) must be a
	// reference to a macro.
	// - For case-insensitive Regex matching, include
	// a boolean parameter named ignore_case that is set to true. If not
	// specified or set to any other value, the matching will be case
	// sensitive.
	// - To negate an operator, include a boolean parameter
	// named negate boolean parameter that is set to true.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// Type: The type of operator for this condition.
	Type string `json:"type,omitempty"`
}

type Container struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: The Container ID uniquely identifies the GTM Container.
	ContainerId string `json:"containerId,omitempty"`

	// DomainName: Optional list of domain names associated with the
	// Container.
	DomainName []string `json:"domainName,omitempty"`

	// EnabledBuiltInVariable: List of enabled built-in variables. Valid
	// values include: pageUrl, pageHostname, pagePath, referrer, event,
	// clickElement, clickClasses, clickId, clickTarget, clickUrl,
	// clickText, formElement, formClasses, formId, formTarget, formUrl,
	// formText, errorMessage, errorUrl, errorLine, newHistoryFragment,
	// oldHistoryFragment, newHistoryState, oldHistoryState, historySource,
	// containerVersion, debugMode, randomNumber, containerId.
	EnabledBuiltInVariable []string `json:"enabledBuiltInVariable,omitempty"`

	// Fingerprint: The fingerprint of the GTM Container as computed at
	// storage time. This value is recomputed whenever the account is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Container display name.
	Name string `json:"name,omitempty"`

	// Notes: Container Notes.
	Notes string `json:"notes,omitempty"`

	// PublicId: Container Public ID.
	PublicId string `json:"publicId,omitempty"`

	// TimeZoneCountryId: Container Country ID.
	TimeZoneCountryId string `json:"timeZoneCountryId,omitempty"`

	// TimeZoneId: Container Time Zone ID.
	TimeZoneId string `json:"timeZoneId,omitempty"`

	// UsageContext: List of Usage Contexts for the Container. Valid values
	// include: web, android, ios.
	UsageContext []string `json:"usageContext,omitempty"`
}

type ContainerAccess struct {
	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Permission: List of Container permissions. Valid container
	// permissions are: read, edit, delete, publish.
	Permission []string `json:"permission,omitempty"`
}

type ContainerVersion struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// Container: The container that this version was taken from.
	Container *Container `json:"container,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// ContainerVersionId: The Container Version ID uniquely identifies the
	// GTM Container Version.
	ContainerVersionId string `json:"containerVersionId,omitempty"`

	// Deleted: A value of true indicates this container version has been
	// deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Fingerprint: The fingerprint of the GTM Container Version as computed
	// at storage time. This value is recomputed whenever the container
	// version is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Macro: The macros in the container that this version was taken from.
	Macro []*Macro `json:"macro,omitempty"`

	// Name: Container version display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this container version in the
	// container.
	Notes string `json:"notes,omitempty"`

	// Rule: The rules in the container that this version was taken from.
	Rule []*Rule `json:"rule,omitempty"`

	// Tag: The tags in the container that this version was taken from.
	Tag []*Tag `json:"tag,omitempty"`

	// Trigger: The triggers in the container that this version was taken
	// from.
	Trigger []*Trigger `json:"trigger,omitempty"`

	// Variable: The variables in the container that this version was taken
	// from.
	Variable []*Variable `json:"variable,omitempty"`
}

type ContainerVersionHeader struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// ContainerVersionId: The Container Version ID uniquely identifies the
	// GTM Container Version.
	ContainerVersionId string `json:"containerVersionId,omitempty"`

	// Deleted: A value of true indicates this container version has been
	// deleted.
	Deleted bool `json:"deleted,omitempty"`

	// Name: Container version display name.
	Name string `json:"name,omitempty"`

	// NumMacros: Number of macros in the container version.
	NumMacros string `json:"numMacros,omitempty"`

	// NumRules: Number of rules in the container version.
	NumRules string `json:"numRules,omitempty"`

	// NumTags: Number of tags in the container version.
	NumTags string `json:"numTags,omitempty"`

	// NumTriggers: Number of triggers in the container version.
	NumTriggers string `json:"numTriggers,omitempty"`

	// NumVariables: Number of variables in the container version.
	NumVariables string `json:"numVariables,omitempty"`
}

type CreateContainerVersionRequestVersionOptions struct {
	// Name: The name of the container version to be created.
	Name string `json:"name,omitempty"`

	// Notes: The notes of the container version to be created.
	Notes string `json:"notes,omitempty"`

	// QuickPreview: The creation of this version may be for quick preview
	// and shouldn't be saved.
	QuickPreview bool `json:"quickPreview,omitempty"`
}

type CreateContainerVersionResponse struct {
	// CompilerError: Compiler errors or not.
	CompilerError bool `json:"compilerError,omitempty"`

	// ContainerVersion: The container version created.
	ContainerVersion *ContainerVersion `json:"containerVersion,omitempty"`
}

type ListAccountUsersResponse struct {
	// UserAccess: All GTM AccountUsers of a GTM Account.
	UserAccess []*UserAccess `json:"userAccess,omitempty"`
}

type ListAccountsResponse struct {
	// Accounts: List of GTM Accounts that a user has access to.
	Accounts []*Account `json:"accounts,omitempty"`
}

type ListContainerVersionsResponse struct {
	// ContainerVersion: All versions of a GTM Container.
	ContainerVersion []*ContainerVersion `json:"containerVersion,omitempty"`

	// ContainerVersionHeader: All container version headers of a GTM
	// Container.
	ContainerVersionHeader []*ContainerVersionHeader `json:"containerVersionHeader,omitempty"`
}

type ListContainersResponse struct {
	// Containers: All Containers of a GTM Account.
	Containers []*Container `json:"containers,omitempty"`
}

type ListMacrosResponse struct {
	// Macros: All GTM Macros of a GTM Container.
	Macros []*Macro `json:"macros,omitempty"`
}

type ListRulesResponse struct {
	// Rules: All GTM Rules of a GTM Container.
	Rules []*Rule `json:"rules,omitempty"`
}

type ListTagsResponse struct {
	// Tags: All GTM Tags of a GTM Container.
	Tags []*Tag `json:"tags,omitempty"`
}

type ListTriggersResponse struct {
	// Triggers: All GTM Triggers of a GTM Container.
	Triggers []*Trigger `json:"triggers,omitempty"`
}

type ListVariablesResponse struct {
	// Variables: All GTM Variables of a GTM Container.
	Variables []*Variable `json:"variables,omitempty"`
}

type Macro struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// DisablingRuleId: For mobile containers only: A list of rule IDs for
	// disabling conditional macros; the macro is enabled if one of the
	// enabling rules is true while all the disabling rules are false.
	// Treated as an unordered set.
	DisablingRuleId []string `json:"disablingRuleId,omitempty"`

	// EnablingRuleId: For mobile containers only: A list of rule IDs for
	// enabling conditional macros; the macro is enabled if one of the
	// enabling rules is true while all the disabling rules are false.
	// Treated as an unordered set.
	EnablingRuleId []string `json:"enablingRuleId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Macro as computed at storage
	// time. This value is recomputed whenever the macro is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// MacroId: The Macro ID uniquely identifies the GTM Macro.
	MacroId string `json:"macroId,omitempty"`

	// Name: Macro display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this macro in the container.
	Notes string `json:"notes,omitempty"`

	// Parameter: The macro's parameters.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// ScheduleEndMs: The end timestamp in milliseconds to schedule a macro.
	ScheduleEndMs int64 `json:"scheduleEndMs,omitempty,string"`

	// ScheduleStartMs: The start timestamp in milliseconds to schedule a
	// macro.
	ScheduleStartMs int64 `json:"scheduleStartMs,omitempty,string"`

	// Type: GTM Macro Type.
	Type string `json:"type,omitempty"`
}

type Parameter struct {
	// Key: The named key that uniquely identifies a parameter. Required for
	// top-level parameters, as well as map values. Ignored for list values.
	Key string `json:"key,omitempty"`

	// List: This list parameter's parameters (keys will be ignored).
	List []*Parameter `json:"list,omitempty"`

	// Map: This map parameter's parameters (must have keys; keys must be
	// unique).
	Map []*Parameter `json:"map,omitempty"`

	// Type: The parameter type. Valid values are:
	// - boolean: The value
	// represents a boolean, represented as 'true' or 'false'
	// - integer:
	// The value represents a 64-bit signed integer value, in base 10
	// -
	// list: A list of parameters should be specified
	// - map: A map of
	// parameters should be specified
	// - template: The value represents any
	// text; this can include macro references (even macro references that
	// might return non-string types)
	Type string `json:"type,omitempty"`

	// Value: A parameter's value (may contain macro references such as
	// "{{myMacro}}") as appropriate to the specified type.
	Value string `json:"value,omitempty"`
}

type PublishContainerVersionResponse struct {
	// CompilerError: Compiler errors or not.
	CompilerError bool `json:"compilerError,omitempty"`

	// ContainerVersion: The container version created.
	ContainerVersion *ContainerVersion `json:"containerVersion,omitempty"`
}

type Rule struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// Condition: The list of conditions that make up this rule (implicit
	// AND between them).
	Condition []*Condition `json:"condition,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Rule as computed at storage
	// time. This value is recomputed whenever the rule is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Rule display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this rule in the container.
	Notes string `json:"notes,omitempty"`

	// RuleId: The Rule ID uniquely identifies the GTM Rule.
	RuleId string `json:"ruleId,omitempty"`
}

type Tag struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// BlockingRuleId: Blocking rule IDs. If any of the listed rules
	// evaluate to true, the tag will not fire.
	BlockingRuleId []string `json:"blockingRuleId,omitempty"`

	// BlockingTriggerId: Blocking trigger IDs. If any of the listed
	// triggers evaluate to true, the tag will not fire.
	BlockingTriggerId []string `json:"blockingTriggerId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Tag as computed at storage
	// time. This value is recomputed whenever the tag is modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// FiringRuleId: Firing rule IDs. A tag will fire when any of the listed
	// rules are true and all of its blockingRuleIds (if any specified) are
	// false.
	FiringRuleId []string `json:"firingRuleId,omitempty"`

	// FiringTriggerId: Firing trigger IDs. A tag will fire when any of the
	// listed triggers are true and all of its blockingTriggerIds (if any
	// specified) are false.
	FiringTriggerId []string `json:"firingTriggerId,omitempty"`

	// LiveOnly: If set to true, this tag will only fire in the live
	// environment (e.g. not in preview or debug mode).
	LiveOnly bool `json:"liveOnly,omitempty"`

	// Name: Tag display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this tag in the container.
	Notes string `json:"notes,omitempty"`

	// Parameter: The tag's parameters.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// Priority: User defined numeric priority of the tag. Tags are fired
	// asynchronously in order of priority. Tags with higher numeric value
	// fire first. A tag's priority can be a positive or negative value. The
	// default value is 0.
	Priority *Parameter `json:"priority,omitempty"`

	// ScheduleEndMs: The end timestamp in milliseconds to schedule a tag.
	ScheduleEndMs int64 `json:"scheduleEndMs,omitempty,string"`

	// ScheduleStartMs: The start timestamp in milliseconds to schedule a
	// tag.
	ScheduleStartMs int64 `json:"scheduleStartMs,omitempty,string"`

	// TagId: The Tag ID uniquely identifies the GTM Tag.
	TagId string `json:"tagId,omitempty"`

	// Type: GTM Tag Type.
	Type string `json:"type,omitempty"`
}

type Trigger struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// AutoEventFilter: Used in the case of auto event tracking.
	AutoEventFilter []*Condition `json:"autoEventFilter,omitempty"`

	// CheckValidation: Whether or not we should only fire tags if the form
	// submit or link click event is not cancelled by some other event
	// handler (e.g. because of validation). Only valid for Form Submission
	// and Link Click triggers.
	CheckValidation *Parameter `json:"checkValidation,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// CustomEventFilter: Used in the case of custom event, which is fired
	// iff all Conditions are true.
	CustomEventFilter []*Condition `json:"customEventFilter,omitempty"`

	// EnableAllVideos: Reloads the videos in the page that don't already
	// have the YT API enabled. If false, only capture events from videos
	// that already have the API enabled. Only valid for YouTube triggers.
	EnableAllVideos *Parameter `json:"enableAllVideos,omitempty"`

	// EventName: Name of the GTM event that is fired. Only valid for Timer
	// triggers.
	EventName *Parameter `json:"eventName,omitempty"`

	// Filter: The trigger will only fire iff all Conditions are true.
	Filter []*Condition `json:"filter,omitempty"`

	// Fingerprint: The fingerprint of the GTM Trigger as computed at
	// storage time. This value is recomputed whenever the trigger is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Interval: Time between triggering recurring Timer Events (in
	// milliseconds). Only valid for Timer triggers.
	Interval *Parameter `json:"interval,omitempty"`

	// Limit: Limit of the number of GTM events this Timer Trigger will
	// fire. If no limit is set, we will continue to fire GTM events until
	// the user leaves the page. Only valid for Timer triggers.
	Limit *Parameter `json:"limit,omitempty"`

	// Name: Trigger display name.
	Name string `json:"name,omitempty"`

	// TriggerId: The Trigger ID uniquely identifies the GTM Trigger.
	TriggerId string `json:"triggerId,omitempty"`

	// Type: Defines the data layer event that causes this trigger.
	Type string `json:"type,omitempty"`

	// UniqueTriggerId: Globally unique id of the trigger that
	// auto-generates this (a Form Submit, Link Click or Timer listener) if
	// any. Used to make incompatible auto-events work together with trigger
	// filtering based on trigger ids. This value is populated during output
	// generation since the tags implied by triggers don't exist until then.
	// Only valid for Form Submit, Link Click and Timer triggers.
	UniqueTriggerId *Parameter `json:"uniqueTriggerId,omitempty"`

	// VideoPercentageList: List of integer percentage values. The trigger
	// will fire as each percentage is reached in any instrumented videos.
	// Only valid for YouTube triggers.
	VideoPercentageList *Parameter `json:"videoPercentageList,omitempty"`

	// WaitForTags: Whether or not we should delay the form submissions or
	// link opening until all of the tags have fired (by preventing the
	// default action and later simulating the default action). Only valid
	// for Form Submission and Link Click triggers.
	WaitForTags *Parameter `json:"waitForTags,omitempty"`

	// WaitForTagsTimeout: How long to wait (in milliseconds) for tags to
	// fire when 'waits_for_tags' above evaluates to true. Only valid for
	// Form Submission and Link Click triggers.
	WaitForTagsTimeout *Parameter `json:"waitForTagsTimeout,omitempty"`
}

type UserAccess struct {
	// AccountAccess: GTM Account access permissions.
	AccountAccess *AccountAccess `json:"accountAccess,omitempty"`

	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerAccess: GTM Container access permissions.
	ContainerAccess []*ContainerAccess `json:"containerAccess,omitempty"`

	// EmailAddress: User's email address.
	EmailAddress string `json:"emailAddress,omitempty"`

	// PermissionId: Account Permission ID.
	PermissionId string `json:"permissionId,omitempty"`
}

type Variable struct {
	// AccountId: GTM Account ID.
	AccountId string `json:"accountId,omitempty"`

	// ContainerId: GTM Container ID.
	ContainerId string `json:"containerId,omitempty"`

	// DisablingTriggerId: For mobile containers only: A list of trigger IDs
	// for disabling conditional variables; the variable is enabled if one
	// of the enabling trigger is true while all the disabling trigger are
	// false. Treated as an unordered set.
	DisablingTriggerId []string `json:"disablingTriggerId,omitempty"`

	// EnablingTriggerId: For mobile containers only: A list of trigger IDs
	// for enabling conditional variables; the variable is enabled if one of
	// the enabling triggers is true while all the disabling triggers are
	// false. Treated as an unordered set.
	EnablingTriggerId []string `json:"enablingTriggerId,omitempty"`

	// Fingerprint: The fingerprint of the GTM Variable as computed at
	// storage time. This value is recomputed whenever the variable is
	// modified.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Name: Variable display name.
	Name string `json:"name,omitempty"`

	// Notes: User notes on how to apply this variable in the container.
	Notes string `json:"notes,omitempty"`

	// Parameter: The variable's parameters.
	Parameter []*Parameter `json:"parameter,omitempty"`

	// ScheduleEndMs: The end timestamp in milliseconds to schedule a
	// variable.
	ScheduleEndMs int64 `json:"scheduleEndMs,omitempty,string"`

	// ScheduleStartMs: The start timestamp in milliseconds to schedule a
	// variable.
	ScheduleStartMs int64 `json:"scheduleStartMs,omitempty,string"`

	// Type: GTM Variable Type.
	Type string `json:"type,omitempty"`

	// VariableId: The Variable ID uniquely identifies the GTM Variable.
	VariableId string `json:"variableId,omitempty"`
}

// method id "tagmanager.accounts.get":

type AccountsGetCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a GTM Account.

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
	//   "description": "Gets a GTM Account.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.get",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
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
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.manage.accounts",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.list":

type AccountsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all GTM Accounts that a user has access to.

func (r *AccountsService) List() *AccountsListCall {
	return &AccountsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsListCall) Fields(s ...googleapi.Field) *AccountsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsListCall) Do() (*ListAccountsResponse, error) {
	var returnValue *ListAccountsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all GTM Accounts that a user has access to.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.list",
	//   "path": "accounts",
	//   "response": {
	//     "$ref": "ListAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.manage.accounts",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.update":

type AccountsUpdateCall struct {
	s             *Service
	accountId     string
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a GTM Account.

func (r *AccountsService) Update(accountId string, account *Account) *AccountsUpdateCall {
	return &AccountsUpdateCall{
		s:             r.s,
		accountId:     accountId,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the account in
// storage.
func (c *AccountsUpdateCall) Fingerprint(fingerprint string) *AccountsUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsUpdateCall) Fields(s ...googleapi.Field) *AccountsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsUpdateCall) Do() (*Account, error) {
	var returnValue *Account
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.account,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a GTM Account.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.update",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the account in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.accounts"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.create":

type AccountsContainersCreateCall struct {
	s             *Service
	accountId     string
	container     *Container
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a Container.

func (r *AccountsContainersService) Create(accountId string, container *Container) *AccountsContainersCreateCall {
	return &AccountsContainersCreateCall{
		s:             r.s,
		accountId:     accountId,
		container:     container,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersCreateCall) Fields(s ...googleapi.Field) *AccountsContainersCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersCreateCall) Do() (*Container, error) {
	var returnValue *Container
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.container,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a Container.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers",
	//   "request": {
	//     "$ref": "Container"
	//   },
	//   "response": {
	//     "$ref": "Container"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.delete":

type AccountsContainersDeleteCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a Container.

func (r *AccountsContainersService) Delete(accountId string, containerId string) *AccountsContainersDeleteCall {
	return &AccountsContainersDeleteCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}",
	}
}

func (c *AccountsContainersDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a Container.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.delete.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.get":

type AccountsContainersGetCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a Container.

func (r *AccountsContainersService) Get(accountId string, containerId string) *AccountsContainersGetCall {
	return &AccountsContainersGetCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersGetCall) Fields(s ...googleapi.Field) *AccountsContainersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersGetCall) Do() (*Container, error) {
	var returnValue *Container
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}",
	//   "response": {
	//     "$ref": "Container"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.list":

type AccountsContainersListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all Containers that belongs to a GTM Account.

func (r *AccountsContainersService) List(accountId string) *AccountsContainersListCall {
	return &AccountsContainersListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersListCall) Fields(s ...googleapi.Field) *AccountsContainersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersListCall) Do() (*ListContainersResponse, error) {
	var returnValue *ListContainersResponse
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
	//   "description": "Lists all Containers that belongs to a GTM Account.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers",
	//   "response": {
	//     "$ref": "ListContainersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.update":

type AccountsContainersUpdateCall struct {
	s             *Service
	accountId     string
	containerId   string
	container     *Container
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a Container.

func (r *AccountsContainersService) Update(accountId string, containerId string, container *Container) *AccountsContainersUpdateCall {
	return &AccountsContainersUpdateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		container:     container,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the container in
// storage.
func (c *AccountsContainersUpdateCall) Fingerprint(fingerprint string) *AccountsContainersUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersUpdateCall) Do() (*Container, error) {
	var returnValue *Container
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.container,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a Container.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the container in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}",
	//   "request": {
	//     "$ref": "Container"
	//   },
	//   "response": {
	//     "$ref": "Container"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.create":

type AccountsContainersMacrosCreateCall struct {
	s             *Service
	accountId     string
	containerId   string
	macro         *Macro
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a GTM Macro.

func (r *AccountsContainersMacrosService) Create(accountId string, containerId string, macro *Macro) *AccountsContainersMacrosCreateCall {
	return &AccountsContainersMacrosCreateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		macro:         macro,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/macros",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosCreateCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersMacrosCreateCall) Do() (*Macro, error) {
	var returnValue *Macro
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.macro,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a GTM Macro.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.macros.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros",
	//   "request": {
	//     "$ref": "Macro"
	//   },
	//   "response": {
	//     "$ref": "Macro"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.delete":

type AccountsContainersMacrosDeleteCall struct {
	s             *Service
	accountId     string
	containerId   string
	macroId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a GTM Macro.

func (r *AccountsContainersMacrosService) Delete(accountId string, containerId string, macroId string) *AccountsContainersMacrosDeleteCall {
	return &AccountsContainersMacrosDeleteCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		macroId:       macroId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	}
}

func (c *AccountsContainersMacrosDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"macroId":     c.macroId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a GTM Macro.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.macros.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "macroId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "macroId": {
	//       "description": "The GTM Macro ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.get":

type AccountsContainersMacrosGetCall struct {
	s             *Service
	accountId     string
	containerId   string
	macroId       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a GTM Macro.

func (r *AccountsContainersMacrosService) Get(accountId string, containerId string, macroId string) *AccountsContainersMacrosGetCall {
	return &AccountsContainersMacrosGetCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		macroId:       macroId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosGetCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersMacrosGetCall) Do() (*Macro, error) {
	var returnValue *Macro
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"macroId":     c.macroId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a GTM Macro.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.macros.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "macroId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "macroId": {
	//       "description": "The GTM Macro ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	//   "response": {
	//     "$ref": "Macro"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.list":

type AccountsContainersMacrosListCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all GTM Macros of a Container.

func (r *AccountsContainersMacrosService) List(accountId string, containerId string) *AccountsContainersMacrosListCall {
	return &AccountsContainersMacrosListCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/macros",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosListCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersMacrosListCall) Do() (*ListMacrosResponse, error) {
	var returnValue *ListMacrosResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all GTM Macros of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.macros.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros",
	//   "response": {
	//     "$ref": "ListMacrosResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.macros.update":

type AccountsContainersMacrosUpdateCall struct {
	s             *Service
	accountId     string
	containerId   string
	macroId       string
	macro         *Macro
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a GTM Macro.

func (r *AccountsContainersMacrosService) Update(accountId string, containerId string, macroId string, macro *Macro) *AccountsContainersMacrosUpdateCall {
	return &AccountsContainersMacrosUpdateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		macroId:       macroId,
		macro:         macro,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the macro in storage.
func (c *AccountsContainersMacrosUpdateCall) Fingerprint(fingerprint string) *AccountsContainersMacrosUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersMacrosUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersMacrosUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersMacrosUpdateCall) Do() (*Macro, error) {
	var returnValue *Macro
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"macroId":     c.macroId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.macro,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a GTM Macro.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.macros.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "macroId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the macro in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "macroId": {
	//       "description": "The GTM Macro ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/macros/{macroId}",
	//   "request": {
	//     "$ref": "Macro"
	//   },
	//   "response": {
	//     "$ref": "Macro"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.create":

type AccountsContainersRulesCreateCall struct {
	s             *Service
	accountId     string
	containerId   string
	rule          *Rule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a GTM Rule.

func (r *AccountsContainersRulesService) Create(accountId string, containerId string, rule *Rule) *AccountsContainersRulesCreateCall {
	return &AccountsContainersRulesCreateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		rule:          rule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/rules",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesCreateCall) Fields(s ...googleapi.Field) *AccountsContainersRulesCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersRulesCreateCall) Do() (*Rule, error) {
	var returnValue *Rule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.rule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a GTM Rule.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.rules.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules",
	//   "request": {
	//     "$ref": "Rule"
	//   },
	//   "response": {
	//     "$ref": "Rule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.delete":

type AccountsContainersRulesDeleteCall struct {
	s             *Service
	accountId     string
	containerId   string
	ruleId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a GTM Rule.

func (r *AccountsContainersRulesService) Delete(accountId string, containerId string, ruleId string) *AccountsContainersRulesDeleteCall {
	return &AccountsContainersRulesDeleteCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		ruleId:        ruleId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	}
}

func (c *AccountsContainersRulesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"ruleId":      c.ruleId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a GTM Rule.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.rules.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "The GTM Rule ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.get":

type AccountsContainersRulesGetCall struct {
	s             *Service
	accountId     string
	containerId   string
	ruleId        string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a GTM Rule.

func (r *AccountsContainersRulesService) Get(accountId string, containerId string, ruleId string) *AccountsContainersRulesGetCall {
	return &AccountsContainersRulesGetCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		ruleId:        ruleId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesGetCall) Fields(s ...googleapi.Field) *AccountsContainersRulesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersRulesGetCall) Do() (*Rule, error) {
	var returnValue *Rule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"ruleId":      c.ruleId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a GTM Rule.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.rules.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "The GTM Rule ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	//   "response": {
	//     "$ref": "Rule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.list":

type AccountsContainersRulesListCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all GTM Rules of a Container.

func (r *AccountsContainersRulesService) List(accountId string, containerId string) *AccountsContainersRulesListCall {
	return &AccountsContainersRulesListCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/rules",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesListCall) Fields(s ...googleapi.Field) *AccountsContainersRulesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersRulesListCall) Do() (*ListRulesResponse, error) {
	var returnValue *ListRulesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all GTM Rules of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.rules.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules",
	//   "response": {
	//     "$ref": "ListRulesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.rules.update":

type AccountsContainersRulesUpdateCall struct {
	s             *Service
	accountId     string
	containerId   string
	ruleId        string
	rule          *Rule
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a GTM Rule.

func (r *AccountsContainersRulesService) Update(accountId string, containerId string, ruleId string, rule *Rule) *AccountsContainersRulesUpdateCall {
	return &AccountsContainersRulesUpdateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		ruleId:        ruleId,
		rule:          rule,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the rule in storage.
func (c *AccountsContainersRulesUpdateCall) Fingerprint(fingerprint string) *AccountsContainersRulesUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersRulesUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersRulesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersRulesUpdateCall) Do() (*Rule, error) {
	var returnValue *Rule
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"ruleId":      c.ruleId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.rule,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a GTM Rule.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.rules.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "ruleId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the rule in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ruleId": {
	//       "description": "The GTM Rule ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/rules/{ruleId}",
	//   "request": {
	//     "$ref": "Rule"
	//   },
	//   "response": {
	//     "$ref": "Rule"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.create":

type AccountsContainersTagsCreateCall struct {
	s             *Service
	accountId     string
	containerId   string
	tag           *Tag
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a GTM Tag.

func (r *AccountsContainersTagsService) Create(accountId string, containerId string, tag *Tag) *AccountsContainersTagsCreateCall {
	return &AccountsContainersTagsCreateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		tag:           tag,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/tags",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsCreateCall) Fields(s ...googleapi.Field) *AccountsContainersTagsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTagsCreateCall) Do() (*Tag, error) {
	var returnValue *Tag
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.tag,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a GTM Tag.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.tags.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags",
	//   "request": {
	//     "$ref": "Tag"
	//   },
	//   "response": {
	//     "$ref": "Tag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.delete":

type AccountsContainersTagsDeleteCall struct {
	s             *Service
	accountId     string
	containerId   string
	tagId         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a GTM Tag.

func (r *AccountsContainersTagsService) Delete(accountId string, containerId string, tagId string) *AccountsContainersTagsDeleteCall {
	return &AccountsContainersTagsDeleteCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		tagId:         tagId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	}
}

func (c *AccountsContainersTagsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"tagId":       c.tagId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a GTM Tag.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.tags.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "tagId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The GTM Tag ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.get":

type AccountsContainersTagsGetCall struct {
	s             *Service
	accountId     string
	containerId   string
	tagId         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a GTM Tag.

func (r *AccountsContainersTagsService) Get(accountId string, containerId string, tagId string) *AccountsContainersTagsGetCall {
	return &AccountsContainersTagsGetCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		tagId:         tagId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsGetCall) Fields(s ...googleapi.Field) *AccountsContainersTagsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTagsGetCall) Do() (*Tag, error) {
	var returnValue *Tag
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"tagId":       c.tagId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a GTM Tag.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.tags.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "tagId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The GTM Tag ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	//   "response": {
	//     "$ref": "Tag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.list":

type AccountsContainersTagsListCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all GTM Tags of a Container.

func (r *AccountsContainersTagsService) List(accountId string, containerId string) *AccountsContainersTagsListCall {
	return &AccountsContainersTagsListCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/tags",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsListCall) Fields(s ...googleapi.Field) *AccountsContainersTagsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTagsListCall) Do() (*ListTagsResponse, error) {
	var returnValue *ListTagsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all GTM Tags of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.tags.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags",
	//   "response": {
	//     "$ref": "ListTagsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.tags.update":

type AccountsContainersTagsUpdateCall struct {
	s             *Service
	accountId     string
	containerId   string
	tagId         string
	tag           *Tag
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a GTM Tag.

func (r *AccountsContainersTagsService) Update(accountId string, containerId string, tagId string, tag *Tag) *AccountsContainersTagsUpdateCall {
	return &AccountsContainersTagsUpdateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		tagId:         tagId,
		tag:           tag,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the tag in storage.
func (c *AccountsContainersTagsUpdateCall) Fingerprint(fingerprint string) *AccountsContainersTagsUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTagsUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersTagsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTagsUpdateCall) Do() (*Tag, error) {
	var returnValue *Tag
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"tagId":       c.tagId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.tag,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a GTM Tag.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.tags.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "tagId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the tag in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tagId": {
	//       "description": "The GTM Tag ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/tags/{tagId}",
	//   "request": {
	//     "$ref": "Tag"
	//   },
	//   "response": {
	//     "$ref": "Tag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.create":

type AccountsContainersTriggersCreateCall struct {
	s             *Service
	accountId     string
	containerId   string
	trigger       *Trigger
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a GTM Trigger.

func (r *AccountsContainersTriggersService) Create(accountId string, containerId string, trigger *Trigger) *AccountsContainersTriggersCreateCall {
	return &AccountsContainersTriggersCreateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		trigger:       trigger,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/triggers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersCreateCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTriggersCreateCall) Do() (*Trigger, error) {
	var returnValue *Trigger
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.trigger,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a GTM Trigger.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.triggers.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers",
	//   "request": {
	//     "$ref": "Trigger"
	//   },
	//   "response": {
	//     "$ref": "Trigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.delete":

type AccountsContainersTriggersDeleteCall struct {
	s             *Service
	accountId     string
	containerId   string
	triggerId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a GTM Trigger.

func (r *AccountsContainersTriggersService) Delete(accountId string, containerId string, triggerId string) *AccountsContainersTriggersDeleteCall {
	return &AccountsContainersTriggersDeleteCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		triggerId:     triggerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	}
}

func (c *AccountsContainersTriggersDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"triggerId":   c.triggerId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a GTM Trigger.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.triggers.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The GTM Trigger ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.get":

type AccountsContainersTriggersGetCall struct {
	s             *Service
	accountId     string
	containerId   string
	triggerId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a GTM Trigger.

func (r *AccountsContainersTriggersService) Get(accountId string, containerId string, triggerId string) *AccountsContainersTriggersGetCall {
	return &AccountsContainersTriggersGetCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		triggerId:     triggerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersGetCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTriggersGetCall) Do() (*Trigger, error) {
	var returnValue *Trigger
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"triggerId":   c.triggerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a GTM Trigger.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.triggers.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The GTM Trigger ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	//   "response": {
	//     "$ref": "Trigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.list":

type AccountsContainersTriggersListCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all GTM Triggers of a Container.

func (r *AccountsContainersTriggersService) List(accountId string, containerId string) *AccountsContainersTriggersListCall {
	return &AccountsContainersTriggersListCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/triggers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersListCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTriggersListCall) Do() (*ListTriggersResponse, error) {
	var returnValue *ListTriggersResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all GTM Triggers of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.triggers.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers",
	//   "response": {
	//     "$ref": "ListTriggersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.triggers.update":

type AccountsContainersTriggersUpdateCall struct {
	s             *Service
	accountId     string
	containerId   string
	triggerId     string
	trigger       *Trigger
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a GTM Trigger.

func (r *AccountsContainersTriggersService) Update(accountId string, containerId string, triggerId string, trigger *Trigger) *AccountsContainersTriggersUpdateCall {
	return &AccountsContainersTriggersUpdateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		triggerId:     triggerId,
		trigger:       trigger,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the trigger in
// storage.
func (c *AccountsContainersTriggersUpdateCall) Fingerprint(fingerprint string) *AccountsContainersTriggersUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersTriggersUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersTriggersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersTriggersUpdateCall) Do() (*Trigger, error) {
	var returnValue *Trigger
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"triggerId":   c.triggerId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.trigger,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a GTM Trigger.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.triggers.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "triggerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the trigger in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "triggerId": {
	//       "description": "The GTM Trigger ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/triggers/{triggerId}",
	//   "request": {
	//     "$ref": "Trigger"
	//   },
	//   "response": {
	//     "$ref": "Trigger"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.create":

type AccountsContainersVariablesCreateCall struct {
	s             *Service
	accountId     string
	containerId   string
	variable      *Variable
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a GTM Variable.

func (r *AccountsContainersVariablesService) Create(accountId string, containerId string, variable *Variable) *AccountsContainersVariablesCreateCall {
	return &AccountsContainersVariablesCreateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		variable:      variable,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/variables",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesCreateCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVariablesCreateCall) Do() (*Variable, error) {
	var returnValue *Variable
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.variable,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a GTM Variable.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.variables.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables",
	//   "request": {
	//     "$ref": "Variable"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.delete":

type AccountsContainersVariablesDeleteCall struct {
	s             *Service
	accountId     string
	containerId   string
	variableId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a GTM Variable.

func (r *AccountsContainersVariablesService) Delete(accountId string, containerId string, variableId string) *AccountsContainersVariablesDeleteCall {
	return &AccountsContainersVariablesDeleteCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		variableId:    variableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	}
}

func (c *AccountsContainersVariablesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"variableId":  c.variableId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a GTM Variable.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.variables.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "variableId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The GTM Variable ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.get":

type AccountsContainersVariablesGetCall struct {
	s             *Service
	accountId     string
	containerId   string
	variableId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a GTM Variable.

func (r *AccountsContainersVariablesService) Get(accountId string, containerId string, variableId string) *AccountsContainersVariablesGetCall {
	return &AccountsContainersVariablesGetCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		variableId:    variableId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesGetCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVariablesGetCall) Do() (*Variable, error) {
	var returnValue *Variable
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"variableId":  c.variableId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a GTM Variable.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.variables.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "variableId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The GTM Variable ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.list":

type AccountsContainersVariablesListCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all GTM Variables of a Container.

func (r *AccountsContainersVariablesService) List(accountId string, containerId string) *AccountsContainersVariablesListCall {
	return &AccountsContainersVariablesListCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/variables",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesListCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVariablesListCall) Do() (*ListVariablesResponse, error) {
	var returnValue *ListVariablesResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all GTM Variables of a Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.variables.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables",
	//   "response": {
	//     "$ref": "ListVariablesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.variables.update":

type AccountsContainersVariablesUpdateCall struct {
	s             *Service
	accountId     string
	containerId   string
	variableId    string
	variable      *Variable
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a GTM Variable.

func (r *AccountsContainersVariablesService) Update(accountId string, containerId string, variableId string, variable *Variable) *AccountsContainersVariablesUpdateCall {
	return &AccountsContainersVariablesUpdateCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		variableId:    variableId,
		variable:      variable,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the variable in
// storage.
func (c *AccountsContainersVariablesUpdateCall) Fingerprint(fingerprint string) *AccountsContainersVariablesUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVariablesUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersVariablesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVariablesUpdateCall) Do() (*Variable, error) {
	var returnValue *Variable
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
		"variableId":  c.variableId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.variable,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a GTM Variable.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.variables.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "variableId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the variable in storage.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "variableId": {
	//       "description": "The GTM Variable ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/variables/{variableId}",
	//   "request": {
	//     "$ref": "Variable"
	//   },
	//   "response": {
	//     "$ref": "Variable"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.create":

type AccountsContainersVersionsCreateCall struct {
	s                                           *Service
	accountId                                   string
	containerId                                 string
	createcontainerversionrequestversionoptions *CreateContainerVersionRequestVersionOptions
	caller_                                     googleapi.Caller
	params_                                     url.Values
	pathTemplate_                               string
}

// Create: Creates a Container Version.

func (r *AccountsContainersVersionsService) Create(accountId string, containerId string, createcontainerversionrequestversionoptions *CreateContainerVersionRequestVersionOptions) *AccountsContainersVersionsCreateCall {
	return &AccountsContainersVersionsCreateCall{
		s:                                           r.s,
		accountId:                                   accountId,
		containerId:                                 containerId,
		createcontainerversionrequestversionoptions: createcontainerversionrequestversionoptions,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/versions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsCreateCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsCreateCall) Do() (*CreateContainerVersionResponse, error) {
	var returnValue *CreateContainerVersionResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.createcontainerversionrequestversionoptions,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a Container Version.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.create",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions",
	//   "request": {
	//     "$ref": "CreateContainerVersionRequestVersionOptions"
	//   },
	//   "response": {
	//     "$ref": "CreateContainerVersionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.delete":

type AccountsContainersVersionsDeleteCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Delete: Deletes a Container Version.

func (r *AccountsContainersVersionsService) Delete(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsDeleteCall {
	return &AccountsContainersVersionsDeleteCall{
		s:                  r.s,
		accountId:          accountId,
		containerId:        containerId,
		containerVersionId: containerVersionId,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	}
}

func (c *AccountsContainersVersionsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a Container Version.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.containers.versions.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.get":

type AccountsContainersVersionsGetCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Get: Gets a Container Version.

func (r *AccountsContainersVersionsService) Get(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsGetCall {
	return &AccountsContainersVersionsGetCall{
		s:                  r.s,
		accountId:          accountId,
		containerId:        containerId,
		containerVersionId: containerVersionId,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsGetCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsGetCall) Do() (*ContainerVersion, error) {
	var returnValue *ContainerVersion
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a Container Version.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.versions.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID. Specify published to retrieve the currently published version.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.list":

type AccountsContainersVersionsListCall struct {
	s             *Service
	accountId     string
	containerId   string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists all Container Versions of a GTM Container.

func (r *AccountsContainersVersionsService) List(accountId string, containerId string) *AccountsContainersVersionsListCall {
	return &AccountsContainersVersionsListCall{
		s:             r.s,
		accountId:     accountId,
		containerId:   containerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/containers/{containerId}/versions",
	}
}

// Headers sets the optional parameter "headers": Retrieve headers only
// when true.
func (c *AccountsContainersVersionsListCall) Headers(headers bool) *AccountsContainersVersionsListCall {
	c.params_.Set("headers", fmt.Sprintf("%v", headers))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsListCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsListCall) Do() (*ListContainerVersionsResponse, error) {
	var returnValue *ListContainerVersionsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":   c.accountId,
		"containerId": c.containerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists all Container Versions of a GTM Container.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.containers.versions.list",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "headers": {
	//       "default": "false",
	//       "description": "Retrieve headers only when true.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions",
	//   "response": {
	//     "$ref": "ListContainerVersionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers",
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions",
	//     "https://www.googleapis.com/auth/tagmanager.readonly"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.publish":

type AccountsContainersVersionsPublishCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Publish: Publishes a Container Version.

func (r *AccountsContainersVersionsService) Publish(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsPublishCall {
	return &AccountsContainersVersionsPublishCall{
		s:                  r.s,
		accountId:          accountId,
		containerId:        containerId,
		containerVersionId: containerVersionId,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/publish",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the container version
// in storage.
func (c *AccountsContainersVersionsPublishCall) Fingerprint(fingerprint string) *AccountsContainersVersionsPublishCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsPublishCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsPublishCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsPublishCall) Do() (*PublishContainerVersionResponse, error) {
	var returnValue *PublishContainerVersionResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Publishes a Container Version.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.publish",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the container version in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/publish",
	//   "response": {
	//     "$ref": "PublishContainerVersionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.publish"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.restore":

type AccountsContainersVersionsRestoreCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Restore: Restores a Container Version. This will overwrite the
// container's current configuration (including its macros, rules and
// tags). The operation will not have any effect on the version that is
// being served (i.e. the published version).

func (r *AccountsContainersVersionsService) Restore(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsRestoreCall {
	return &AccountsContainersVersionsRestoreCall{
		s:                  r.s,
		accountId:          accountId,
		containerId:        containerId,
		containerVersionId: containerVersionId,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/restore",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsRestoreCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsRestoreCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsRestoreCall) Do() (*ContainerVersion, error) {
	var returnValue *ContainerVersion
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Restores a Container Version. This will overwrite the container's current configuration (including its macros, rules and tags). The operation will not have any effect on the version that is being served (i.e. the published version).",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.restore",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/restore",
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containers"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.undelete":

type AccountsContainersVersionsUndeleteCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Undelete: Undeletes a Container Version.

func (r *AccountsContainersVersionsService) Undelete(accountId string, containerId string, containerVersionId string) *AccountsContainersVersionsUndeleteCall {
	return &AccountsContainersVersionsUndeleteCall{
		s:                  r.s,
		accountId:          accountId,
		containerId:        containerId,
		containerVersionId: containerVersionId,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/undelete",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsUndeleteCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsUndeleteCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsUndeleteCall) Do() (*ContainerVersion, error) {
	var returnValue *ContainerVersion
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Undeletes a Container Version.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.containers.versions.undelete",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}/undelete",
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.containers.versions.update":

type AccountsContainersVersionsUpdateCall struct {
	s                  *Service
	accountId          string
	containerId        string
	containerVersionId string
	containerversion   *ContainerVersion
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Update: Updates a Container Version.

func (r *AccountsContainersVersionsService) Update(accountId string, containerId string, containerVersionId string, containerversion *ContainerVersion) *AccountsContainersVersionsUpdateCall {
	return &AccountsContainersVersionsUpdateCall{
		s:                  r.s,
		accountId:          accountId,
		containerId:        containerId,
		containerVersionId: containerVersionId,
		containerversion:   containerversion,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	}
}

// Fingerprint sets the optional parameter "fingerprint": When provided,
// this fingerprint must match the fingerprint of the container version
// in storage.
func (c *AccountsContainersVersionsUpdateCall) Fingerprint(fingerprint string) *AccountsContainersVersionsUpdateCall {
	c.params_.Set("fingerprint", fmt.Sprintf("%v", fingerprint))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsContainersVersionsUpdateCall) Fields(s ...googleapi.Field) *AccountsContainersVersionsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsContainersVersionsUpdateCall) Do() (*ContainerVersion, error) {
	var returnValue *ContainerVersion
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":          c.accountId,
		"containerId":        c.containerId,
		"containerVersionId": c.containerVersionId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.containerversion,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a Container Version.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.containers.versions.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "containerId",
	//     "containerVersionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerId": {
	//       "description": "The GTM Container ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "containerVersionId": {
	//       "description": "The GTM Container Version ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "fingerprint": {
	//       "description": "When provided, this fingerprint must match the fingerprint of the container version in storage.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/containers/{containerId}/versions/{containerVersionId}",
	//   "request": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "response": {
	//     "$ref": "ContainerVersion"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.edit.containerversions"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.create":

type AccountsPermissionsCreateCall struct {
	s             *Service
	accountId     string
	useraccess    *UserAccess
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Create: Creates a user's Account & Container Permissions.

func (r *AccountsPermissionsService) Create(accountId string, useraccess *UserAccess) *AccountsPermissionsCreateCall {
	return &AccountsPermissionsCreateCall{
		s:             r.s,
		accountId:     accountId,
		useraccess:    useraccess,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/permissions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsCreateCall) Fields(s ...googleapi.Field) *AccountsPermissionsCreateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPermissionsCreateCall) Do() (*UserAccess, error) {
	var returnValue *UserAccess
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId": c.accountId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.useraccess,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a user's Account \u0026 Container Permissions.",
	//   "httpMethod": "POST",
	//   "id": "tagmanager.accounts.permissions.create",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions",
	//   "request": {
	//     "$ref": "UserAccess"
	//   },
	//   "response": {
	//     "$ref": "UserAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.delete":

type AccountsPermissionsDeleteCall struct {
	s             *Service
	accountId     string
	permissionId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Removes a user from the account, revoking access to it and
// all of its containers.

func (r *AccountsPermissionsService) Delete(accountId string, permissionId string) *AccountsPermissionsDeleteCall {
	return &AccountsPermissionsDeleteCall{
		s:             r.s,
		accountId:     accountId,
		permissionId:  permissionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/permissions/{permissionId}",
	}
}

func (c *AccountsPermissionsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":    c.accountId,
		"permissionId": c.permissionId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Removes a user from the account, revoking access to it and all of its containers.",
	//   "httpMethod": "DELETE",
	//   "id": "tagmanager.accounts.permissions.delete",
	//   "parameterOrder": [
	//     "accountId",
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The GTM User ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions/{permissionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.get":

type AccountsPermissionsGetCall struct {
	s             *Service
	accountId     string
	permissionId  string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets a user's Account & Container Permissions.

func (r *AccountsPermissionsService) Get(accountId string, permissionId string) *AccountsPermissionsGetCall {
	return &AccountsPermissionsGetCall{
		s:             r.s,
		accountId:     accountId,
		permissionId:  permissionId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/permissions/{permissionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsGetCall) Fields(s ...googleapi.Field) *AccountsPermissionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPermissionsGetCall) Do() (*UserAccess, error) {
	var returnValue *UserAccess
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":    c.accountId,
		"permissionId": c.permissionId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a user's Account \u0026 Container Permissions.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.permissions.get",
	//   "parameterOrder": [
	//     "accountId",
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The GTM User ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions/{permissionId}",
	//   "response": {
	//     "$ref": "UserAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.list":

type AccountsPermissionsListCall struct {
	s             *Service
	accountId     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all users that have access to the account along with
// Account and Container Permissions granted to each of them.

func (r *AccountsPermissionsService) List(accountId string) *AccountsPermissionsListCall {
	return &AccountsPermissionsListCall{
		s:             r.s,
		accountId:     accountId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/permissions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsListCall) Fields(s ...googleapi.Field) *AccountsPermissionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPermissionsListCall) Do() (*ListAccountUsersResponse, error) {
	var returnValue *ListAccountUsersResponse
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
	//   "description": "List all users that have access to the account along with Account and Container Permissions granted to each of them.",
	//   "httpMethod": "GET",
	//   "id": "tagmanager.accounts.permissions.list",
	//   "parameterOrder": [
	//     "accountId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID. @required tagmanager.accounts.permissions.list",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions",
	//   "response": {
	//     "$ref": "ListAccountUsersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}

// method id "tagmanager.accounts.permissions.update":

type AccountsPermissionsUpdateCall struct {
	s             *Service
	accountId     string
	permissionId  string
	useraccess    *UserAccess
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a user's Account & Container Permissions.

func (r *AccountsPermissionsService) Update(accountId string, permissionId string, useraccess *UserAccess) *AccountsPermissionsUpdateCall {
	return &AccountsPermissionsUpdateCall{
		s:             r.s,
		accountId:     accountId,
		permissionId:  permissionId,
		useraccess:    useraccess,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "accounts/{accountId}/permissions/{permissionId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPermissionsUpdateCall) Fields(s ...googleapi.Field) *AccountsPermissionsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPermissionsUpdateCall) Do() (*UserAccess, error) {
	var returnValue *UserAccess
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"accountId":    c.accountId,
		"permissionId": c.permissionId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.useraccess,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a user's Account \u0026 Container Permissions.",
	//   "httpMethod": "PUT",
	//   "id": "tagmanager.accounts.permissions.update",
	//   "parameterOrder": [
	//     "accountId",
	//     "permissionId"
	//   ],
	//   "parameters": {
	//     "accountId": {
	//       "description": "The GTM Account ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "permissionId": {
	//       "description": "The GTM User ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "accounts/{accountId}/permissions/{permissionId}",
	//   "request": {
	//     "$ref": "UserAccess"
	//   },
	//   "response": {
	//     "$ref": "UserAccess"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/tagmanager.manage.users"
	//   ]
	// }

}
