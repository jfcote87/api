// Package admin provides access to the Admin Directory API.
//
// See https://developers.google.com/admin-sdk/directory/
//
// Usage example:
//
//   import "github.com/jfcote87/api/admin/directory_v1"
//   ...
//   adminService, err := admin.New(oauthHttpClient)
package admin

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

const apiId = "admin:directory_v1"
const apiName = "admin"
const apiVersion = "directory_v1"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/admin/directory/v1/"}

// OAuth2 scopes used by this API.
const (
	// View and manage your Chrome OS devices' metadata
	AdminDirectoryDeviceChromeosScope = "https://www.googleapis.com/auth/admin.directory.device.chromeos"

	// View your Chrome OS devices' metadata
	AdminDirectoryDeviceChromeosReadonlyScope = "https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly"

	// View and manage your mobile devices' metadata
	AdminDirectoryDeviceMobileScope = "https://www.googleapis.com/auth/admin.directory.device.mobile"

	// Manage your mobile devices by performing administrative tasks
	AdminDirectoryDeviceMobileActionScope = "https://www.googleapis.com/auth/admin.directory.device.mobile.action"

	// View your mobile devices' metadata
	AdminDirectoryDeviceMobileReadonlyScope = "https://www.googleapis.com/auth/admin.directory.device.mobile.readonly"

	// View and manage the provisioning of groups on your domain
	AdminDirectoryGroupScope = "https://www.googleapis.com/auth/admin.directory.group"

	// View and manage group subscriptions on your domain
	AdminDirectoryGroupMemberScope = "https://www.googleapis.com/auth/admin.directory.group.member"

	// View group subscriptions on your domain
	AdminDirectoryGroupMemberReadonlyScope = "https://www.googleapis.com/auth/admin.directory.group.member.readonly"

	// View groups on your domain
	AdminDirectoryGroupReadonlyScope = "https://www.googleapis.com/auth/admin.directory.group.readonly"

	// View and manage notifications received on your domain
	AdminDirectoryNotificationsScope = "https://www.googleapis.com/auth/admin.directory.notifications"

	// View and manage organization units on your domain
	AdminDirectoryOrgunitScope = "https://www.googleapis.com/auth/admin.directory.orgunit"

	// View organization units on your domain
	AdminDirectoryOrgunitReadonlyScope = "https://www.googleapis.com/auth/admin.directory.orgunit.readonly"

	// View and manage the provisioning of users on your domain
	AdminDirectoryUserScope = "https://www.googleapis.com/auth/admin.directory.user"

	// View and manage user aliases on your domain
	AdminDirectoryUserAliasScope = "https://www.googleapis.com/auth/admin.directory.user.alias"

	// View user aliases on your domain
	AdminDirectoryUserAliasReadonlyScope = "https://www.googleapis.com/auth/admin.directory.user.alias.readonly"

	// View users on your domain
	AdminDirectoryUserReadonlyScope = "https://www.googleapis.com/auth/admin.directory.user.readonly"

	// Manage data access permissions for users on your domain
	AdminDirectoryUserSecurityScope = "https://www.googleapis.com/auth/admin.directory.user.security"

	// View and manage the provisioning of user schemas on your domain
	AdminDirectoryUserschemaScope = "https://www.googleapis.com/auth/admin.directory.userschema"

	// View user schemas on your domain
	AdminDirectoryUserschemaReadonlyScope = "https://www.googleapis.com/auth/admin.directory.userschema.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.Asps = NewAspsService(s)
	s.Channels = NewChannelsService(s)
	s.Chromeosdevices = NewChromeosdevicesService(s)
	s.Groups = NewGroupsService(s)
	s.Members = NewMembersService(s)
	s.Mobiledevices = NewMobiledevicesService(s)
	s.Notifications = NewNotificationsService(s)
	s.Orgunits = NewOrgunitsService(s)
	s.Schemas = NewSchemasService(s)
	s.Tokens = NewTokensService(s)
	s.Users = NewUsersService(s)
	s.VerificationCodes = NewVerificationCodesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	Asps *AspsService

	Channels *ChannelsService

	Chromeosdevices *ChromeosdevicesService

	Groups *GroupsService

	Members *MembersService

	Mobiledevices *MobiledevicesService

	Notifications *NotificationsService

	Orgunits *OrgunitsService

	Schemas *SchemasService

	Tokens *TokensService

	Users *UsersService

	VerificationCodes *VerificationCodesService
}

func NewAspsService(s *Service) *AspsService {
	rs := &AspsService{s: s}
	return rs
}

type AspsService struct {
	s *Service
}

func NewChannelsService(s *Service) *ChannelsService {
	rs := &ChannelsService{s: s}
	return rs
}

type ChannelsService struct {
	s *Service
}

func NewChromeosdevicesService(s *Service) *ChromeosdevicesService {
	rs := &ChromeosdevicesService{s: s}
	return rs
}

type ChromeosdevicesService struct {
	s *Service
}

func NewGroupsService(s *Service) *GroupsService {
	rs := &GroupsService{s: s}
	rs.Aliases = NewGroupsAliasesService(s)
	return rs
}

type GroupsService struct {
	s *Service

	Aliases *GroupsAliasesService
}

func NewGroupsAliasesService(s *Service) *GroupsAliasesService {
	rs := &GroupsAliasesService{s: s}
	return rs
}

type GroupsAliasesService struct {
	s *Service
}

func NewMembersService(s *Service) *MembersService {
	rs := &MembersService{s: s}
	return rs
}

type MembersService struct {
	s *Service
}

func NewMobiledevicesService(s *Service) *MobiledevicesService {
	rs := &MobiledevicesService{s: s}
	return rs
}

type MobiledevicesService struct {
	s *Service
}

func NewNotificationsService(s *Service) *NotificationsService {
	rs := &NotificationsService{s: s}
	return rs
}

type NotificationsService struct {
	s *Service
}

func NewOrgunitsService(s *Service) *OrgunitsService {
	rs := &OrgunitsService{s: s}
	return rs
}

type OrgunitsService struct {
	s *Service
}

func NewSchemasService(s *Service) *SchemasService {
	rs := &SchemasService{s: s}
	return rs
}

type SchemasService struct {
	s *Service
}

func NewTokensService(s *Service) *TokensService {
	rs := &TokensService{s: s}
	return rs
}

type TokensService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	rs.Aliases = NewUsersAliasesService(s)
	rs.Photos = NewUsersPhotosService(s)
	return rs
}

type UsersService struct {
	s *Service

	Aliases *UsersAliasesService

	Photos *UsersPhotosService
}

func NewUsersAliasesService(s *Service) *UsersAliasesService {
	rs := &UsersAliasesService{s: s}
	return rs
}

type UsersAliasesService struct {
	s *Service
}

func NewUsersPhotosService(s *Service) *UsersPhotosService {
	rs := &UsersPhotosService{s: s}
	return rs
}

type UsersPhotosService struct {
	s *Service
}

func NewVerificationCodesService(s *Service) *VerificationCodesService {
	rs := &VerificationCodesService{s: s}
	return rs
}

type VerificationCodesService struct {
	s *Service
}

type Alias struct {
	// Alias: A alias email
	Alias string `json:"alias,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: Unique id of the group (Read-only) Unique id of the user
	// (Read-only)
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// PrimaryEmail: Group's primary email (Read-only) User's primary email
	// (Read-only)
	PrimaryEmail string `json:"primaryEmail,omitempty"`
}

type Aliases struct {
	// Aliases: List of alias objects.
	Aliases []*Alias `json:"aliases,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`
}

type Asp struct {
	// CodeId: The unique ID of the ASP.
	CodeId int64 `json:"codeId,omitempty"`

	// CreationTime: The time when the ASP was created. Expressed in Unix
	// time format.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// Etag: ETag of the ASP.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of the API resource. This is always
	// admin#directory#asp.
	Kind string `json:"kind,omitempty"`

	// LastTimeUsed: The time when the ASP was last used. Expressed in Unix
	// time format.
	LastTimeUsed int64 `json:"lastTimeUsed,omitempty,string"`

	// Name: The name of the application that the user, represented by their
	// userId, entered when the ASP was created.
	Name string `json:"name,omitempty"`

	// UserKey: The unique ID of the user who issued the ASP.
	UserKey string `json:"userKey,omitempty"`
}

type Asps struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: A list of ASP resources.
	Items []*Asp `json:"items,omitempty"`

	// Kind: The type of the API resource. This is always
	// admin#directory#aspList.
	Kind string `json:"kind,omitempty"`
}

type Channel struct {
	// Address: The address where notifications are delivered for this
	// channel.
	Address string `json:"address,omitempty"`

	// Expiration: Date and time of notification channel expiration,
	// expressed as a Unix timestamp, in milliseconds. Optional.
	Expiration int64 `json:"expiration,omitempty,string"`

	// Id: A UUID or similar unique string that identifies this channel.
	Id string `json:"id,omitempty"`

	// Kind: Identifies this as a notification channel used to watch for
	// changes to a resource. Value: the fixed string "api#channel".
	Kind string `json:"kind,omitempty"`

	// Params: Additional parameters controlling delivery channel behavior.
	// Optional.
	Params map[string]string `json:"params,omitempty"`

	// Payload: A Boolean value to indicate whether payload is wanted.
	// Optional.
	Payload bool `json:"payload,omitempty"`

	// ResourceId: An opaque ID that identifies the resource being watched
	// on this channel. Stable across different API versions.
	ResourceId string `json:"resourceId,omitempty"`

	// ResourceUri: A version-specific identifier for the watched resource.
	ResourceUri string `json:"resourceUri,omitempty"`

	// Token: An arbitrary string delivered to the target address with each
	// notification delivered over this channel. Optional.
	Token string `json:"token,omitempty"`

	// Type: The type of delivery mechanism used for this channel.
	Type string `json:"type,omitempty"`
}

type ChromeOsDevice struct {
	// ActiveTimeRanges: List of active time ranges (Read-only)
	ActiveTimeRanges []*ChromeOsDeviceActiveTimeRanges `json:"activeTimeRanges,omitempty"`

	// AnnotatedLocation: Address or location of the device as noted by the
	// administrator
	AnnotatedLocation string `json:"annotatedLocation,omitempty"`

	// AnnotatedUser: User of the device
	AnnotatedUser string `json:"annotatedUser,omitempty"`

	// BootMode: Chromebook boot mode (Read-only)
	BootMode string `json:"bootMode,omitempty"`

	// DeviceId: Unique identifier of Chrome OS Device (Read-only)
	DeviceId string `json:"deviceId,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// EthernetMacAddress: Chromebook Mac Address on ethernet network
	// interface (Read-only)
	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"`

	// FirmwareVersion: Chromebook firmware version (Read-only)
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// LastEnrollmentTime: Date and time the device was last enrolled
	// (Read-only)
	LastEnrollmentTime string `json:"lastEnrollmentTime,omitempty"`

	// LastSync: Date and time the device was last synchronized with the
	// policy settings in the Google Apps administrator control panel
	// (Read-only)
	LastSync string `json:"lastSync,omitempty"`

	// MacAddress: Chromebook Mac Address on wifi network interface
	// (Read-only)
	MacAddress string `json:"macAddress,omitempty"`

	// Meid: Mobile Equipment identifier for the 3G mobile card in the
	// Chromebook (Read-only)
	Meid string `json:"meid,omitempty"`

	// Model: Chromebook Model (Read-only)
	Model string `json:"model,omitempty"`

	// Notes: Notes added by the administrator
	Notes string `json:"notes,omitempty"`

	// OrderNumber: Chromebook order number (Read-only)
	OrderNumber string `json:"orderNumber,omitempty"`

	// OrgUnitPath: OrgUnit of the device
	OrgUnitPath string `json:"orgUnitPath,omitempty"`

	// OsVersion: Chromebook Os Version (Read-only)
	OsVersion string `json:"osVersion,omitempty"`

	// PlatformVersion: Chromebook platform version (Read-only)
	PlatformVersion string `json:"platformVersion,omitempty"`

	// RecentUsers: List of recent device users, in descending order by last
	// login time (Read-only)
	RecentUsers []*ChromeOsDeviceRecentUsers `json:"recentUsers,omitempty"`

	// SerialNumber: Chromebook serial number (Read-only)
	SerialNumber string `json:"serialNumber,omitempty"`

	// Status: status of the device (Read-only)
	Status string `json:"status,omitempty"`

	// SupportEndDate: Final date the device will be supported (Read-only)
	SupportEndDate string `json:"supportEndDate,omitempty"`

	// WillAutoRenew: Will Chromebook auto renew after support end date
	// (Read-only)
	WillAutoRenew bool `json:"willAutoRenew,omitempty"`
}

type ChromeOsDeviceActiveTimeRanges struct {
	// ActiveTime: Duration in milliseconds
	ActiveTime int64 `json:"activeTime,omitempty"`

	// Date: Date of usage
	Date string `json:"date,omitempty"`
}

type ChromeOsDeviceRecentUsers struct {
	// Email: Email address of the user. Present only if the user type is
	// managed
	Email string `json:"email,omitempty"`

	// Type: The type of the user
	Type string `json:"type,omitempty"`
}

type ChromeOsDevices struct {
	// Chromeosdevices: List of Chrome OS Device objects.
	Chromeosdevices []*ChromeOsDevice `json:"chromeosdevices,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access next page of this result.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Group struct {
	// AdminCreated: Is the group created by admin (Read-only) *
	AdminCreated bool `json:"adminCreated,omitempty"`

	// Aliases: List of aliases (Read-only)
	Aliases []string `json:"aliases,omitempty"`

	// Description: Description of the group
	Description string `json:"description,omitempty"`

	// DirectMembersCount: Group direct members count
	DirectMembersCount int64 `json:"directMembersCount,omitempty,string"`

	// Email: Email of Group
	Email string `json:"email,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: Unique identifier of Group (Read-only)
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Name: Group name
	Name string `json:"name,omitempty"`

	// NonEditableAliases: List of non editable aliases (Read-only)
	NonEditableAliases []string `json:"nonEditableAliases,omitempty"`
}

type Groups struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Groups: List of group objects.
	Groups []*Group `json:"groups,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access next page of this result.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Member struct {
	// Email: Email of member (Read-only)
	Email string `json:"email,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Id: Unique identifier of customer member (Read-only) Unique
	// identifier of group (Read-only) Unique identifier of member
	// (Read-only)
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Role: Role of member
	Role string `json:"role,omitempty"`

	// Type: Type of member (Immutable)
	Type string `json:"type,omitempty"`
}

type Members struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Members: List of member objects.
	Members []*Member `json:"members,omitempty"`

	// NextPageToken: Token used to access next page of this result.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type MobileDevice struct {
	// Applications: List of applications installed on Mobile Device
	Applications []*MobileDeviceApplications `json:"applications,omitempty"`

	// BasebandVersion: Mobile Device Baseband version (Read-only)
	BasebandVersion string `json:"basebandVersion,omitempty"`

	// BuildNumber: Mobile Device Build number (Read-only)
	BuildNumber string `json:"buildNumber,omitempty"`

	// DefaultLanguage: The default locale used on the Mobile Device
	// (Read-only)
	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	// DeviceCompromisedStatus: Mobile Device compromised status (Read-only)
	DeviceCompromisedStatus string `json:"deviceCompromisedStatus,omitempty"`

	// DeviceId: Mobile Device serial number (Read-only)
	DeviceId string `json:"deviceId,omitempty"`

	// Email: List of owner user's email addresses (Read-only)
	Email []string `json:"email,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// FirstSync: Date and time the device was first synchronized with the
	// policy settings in the Google Apps administrator control panel
	// (Read-only)
	FirstSync string `json:"firstSync,omitempty"`

	// HardwareId: Mobile Device Hardware Id (Read-only)
	HardwareId string `json:"hardwareId,omitempty"`

	// Imei: Mobile Device IMEI number (Read-only)
	Imei string `json:"imei,omitempty"`

	// KernelVersion: Mobile Device Kernel version (Read-only)
	KernelVersion string `json:"kernelVersion,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// LastSync: Date and time the device was last synchronized with the
	// policy settings in the Google Apps administrator control panel
	// (Read-only)
	LastSync string `json:"lastSync,omitempty"`

	// ManagedAccountIsOnOwnerProfile: Boolean indicating if this account is
	// on owner/primary profile or not (Read-only)
	ManagedAccountIsOnOwnerProfile bool `json:"managedAccountIsOnOwnerProfile,omitempty"`

	// Meid: Mobile Device MEID number (Read-only)
	Meid string `json:"meid,omitempty"`

	// Model: Name of the model of the device
	Model string `json:"model,omitempty"`

	// Name: List of owner user's names (Read-only)
	Name []string `json:"name,omitempty"`

	// NetworkOperator: Mobile Device mobile or network operator (if
	// available) (Read-only)
	NetworkOperator string `json:"networkOperator,omitempty"`

	// Os: Name of the mobile operating system
	Os string `json:"os,omitempty"`

	// ResourceId: Unique identifier of Mobile Device (Read-only)
	ResourceId string `json:"resourceId,omitempty"`

	// SerialNumber: Mobile Device SSN or Serial Number (Read-only)
	SerialNumber string `json:"serialNumber,omitempty"`

	// Status: Status of the device (Read-only)
	Status string `json:"status,omitempty"`

	// Type: The type of device (Read-only)
	Type string `json:"type,omitempty"`

	// UserAgent: Mobile Device user agent
	UserAgent string `json:"userAgent,omitempty"`

	// WifiMacAddress: Mobile Device WiFi MAC address (Read-only)
	WifiMacAddress string `json:"wifiMacAddress,omitempty"`
}

type MobileDeviceApplications struct {
	// DisplayName: Display name of application
	DisplayName string `json:"displayName,omitempty"`

	// PackageName: Package name of application
	PackageName string `json:"packageName,omitempty"`

	// Permission: List of Permissions for application
	Permission []string `json:"permission,omitempty"`

	// VersionCode: Version code of application
	VersionCode int64 `json:"versionCode,omitempty"`

	// VersionName: Version name of application
	VersionName string `json:"versionName,omitempty"`
}

type MobileDeviceAction struct {
	// Action: Action to be taken on the Mobile Device
	Action string `json:"action,omitempty"`
}

type MobileDevices struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Mobiledevices: List of Mobile Device objects.
	Mobiledevices []*MobileDevice `json:"mobiledevices,omitempty"`

	// NextPageToken: Token used to access next page of this result.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Notification struct {
	// Body: Body of the notification (Read-only)
	Body string `json:"body,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// FromAddress: Address from which the notification is received
	// (Read-only)
	FromAddress string `json:"fromAddress,omitempty"`

	// IsUnread: Boolean indicating whether the notification is unread or
	// not.
	IsUnread bool `json:"isUnread,omitempty"`

	// Kind: The type of the resource.
	Kind string `json:"kind,omitempty"`

	NotificationId string `json:"notificationId,omitempty"`

	// SendTime: Time at which notification was sent (Read-only)
	SendTime string `json:"sendTime,omitempty"`

	// Subject: Subject of the notification (Read-only)
	Subject string `json:"subject,omitempty"`
}

type Notifications struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: List of notifications in this page.
	Items []*Notification `json:"items,omitempty"`

	// Kind: The type of the resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token for fetching the next page of notifications.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// UnreadNotificationsCount: Number of unread notification for the
	// domain.
	UnreadNotificationsCount int64 `json:"unreadNotificationsCount,omitempty"`
}

type OrgUnit struct {
	// BlockInheritance: Should block inheritance
	BlockInheritance bool `json:"blockInheritance,omitempty"`

	// Description: Description of OrgUnit
	Description string `json:"description,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Name: Name of OrgUnit
	Name string `json:"name,omitempty"`

	// OrgUnitPath: Path of OrgUnit
	OrgUnitPath string `json:"orgUnitPath,omitempty"`

	// ParentOrgUnitPath: Path of parent OrgUnit
	ParentOrgUnitPath string `json:"parentOrgUnitPath,omitempty"`
}

type OrgUnits struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// OrganizationUnits: List of user objects.
	OrganizationUnits []*OrgUnit `json:"organizationUnits,omitempty"`
}

type Schema struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Fields: Fields of Schema
	Fields []*SchemaFieldSpec `json:"fields,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// SchemaId: Unique identifier of Schema (Read-only)
	SchemaId string `json:"schemaId,omitempty"`

	// SchemaName: Schema name
	SchemaName string `json:"schemaName,omitempty"`
}

type SchemaFieldSpec struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// FieldId: Unique identifier of Field (Read-only)
	FieldId string `json:"fieldId,omitempty"`

	// FieldName: Name of the field.
	FieldName string `json:"fieldName,omitempty"`

	// FieldType: Type of the field.
	FieldType string `json:"fieldType,omitempty"`

	// Indexed: Boolean specifying whether the field is indexed or not.
	Indexed bool `json:"indexed,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// MultiValued: Boolean specifying whether this is a multi-valued field
	// or not.
	MultiValued bool `json:"multiValued,omitempty"`

	// NumericIndexingSpec: Indexing spec for a numeric field. By default,
	// only exact match queries will be supported for numeric fields.
	// Setting the numericIndexingSpec allows range queries to be supported.
	NumericIndexingSpec *SchemaFieldSpecNumericIndexingSpec `json:"numericIndexingSpec,omitempty"`

	// ReadAccessType: Read ACLs on the field specifying who can view values
	// of this field. Valid values are "ALL_DOMAIN_USERS" and
	// "ADMINS_AND_SELF".
	ReadAccessType string `json:"readAccessType,omitempty"`
}

type SchemaFieldSpecNumericIndexingSpec struct {
	// MaxValue: Maximum value of this field. This is meant to be indicative
	// rather than enforced. Values outside this range will still be
	// indexed, but search may not be as performant.
	MaxValue float64 `json:"maxValue,omitempty"`

	// MinValue: Minimum value of this field. This is meant to be indicative
	// rather than enforced. Values outside this range will still be
	// indexed, but search may not be as performant.
	MinValue float64 `json:"minValue,omitempty"`
}

type Schemas struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// Schemas: List of UserSchema objects.
	Schemas []*Schema `json:"schemas,omitempty"`
}

type Token struct {
	// Anonymous: Whether the application is registered with Google. The
	// value is true if the application has an anonymous Client ID.
	Anonymous bool `json:"anonymous,omitempty"`

	// ClientId: The Client ID of the application the token is issued to.
	ClientId string `json:"clientId,omitempty"`

	// DisplayText: The displayable name of the application the token is
	// issued to.
	DisplayText string `json:"displayText,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of the API resource. This is always
	// admin#directory#token.
	Kind string `json:"kind,omitempty"`

	// NativeApp: Whether the token is issued to an installed application.
	// The value is true if the application is installed to a desktop or
	// mobile device.
	NativeApp bool `json:"nativeApp,omitempty"`

	// Scopes: A list of authorization scopes the application is granted.
	Scopes []string `json:"scopes,omitempty"`

	// UserKey: The unique ID of the user that issued the token.
	UserKey string `json:"userKey,omitempty"`
}

type Tokens struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: A list of Token resources.
	Items []*Token `json:"items,omitempty"`

	// Kind: The type of the API resource. This is always
	// admin#directory#tokenList.
	Kind string `json:"kind,omitempty"`
}

type User struct {
	Addresses interface{} `json:"addresses,omitempty"`

	// AgreedToTerms: Indicates if user has agreed to terms (Read-only)
	AgreedToTerms bool `json:"agreedToTerms,omitempty"`

	// Aliases: List of aliases (Read-only)
	Aliases []string `json:"aliases,omitempty"`

	// ChangePasswordAtNextLogin: Boolean indicating if the user should
	// change password in next login
	ChangePasswordAtNextLogin bool `json:"changePasswordAtNextLogin,omitempty"`

	// CreationTime: User's Google account creation time. (Read-only)
	CreationTime string `json:"creationTime,omitempty"`

	// CustomSchemas: Custom fields of the user.
	CustomSchemas map[string]UserCustomProperties `json:"customSchemas,omitempty"`

	// CustomerId: CustomerId of User (Read-only)
	CustomerId string `json:"customerId,omitempty"`

	DeletionTime string `json:"deletionTime,omitempty"`

	Emails interface{} `json:"emails,omitempty"`

	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	ExternalIds interface{} `json:"externalIds,omitempty"`

	// HashFunction: Hash function name for password. Supported are MD5,
	// SHA-1 and crypt
	HashFunction string `json:"hashFunction,omitempty"`

	// Id: Unique identifier of User (Read-only)
	Id string `json:"id,omitempty"`

	Ims interface{} `json:"ims,omitempty"`

	// IncludeInGlobalAddressList: Boolean indicating if user is included in
	// Global Address List
	IncludeInGlobalAddressList bool `json:"includeInGlobalAddressList,omitempty"`

	// IpWhitelisted: Boolean indicating if ip is whitelisted
	IpWhitelisted bool `json:"ipWhitelisted,omitempty"`

	// IsAdmin: Boolean indicating if the user is admin (Read-only)
	IsAdmin bool `json:"isAdmin,omitempty"`

	// IsDelegatedAdmin: Boolean indicating if the user is delegated admin
	// (Read-only)
	IsDelegatedAdmin bool `json:"isDelegatedAdmin,omitempty"`

	// IsMailboxSetup: Is mailbox setup (Read-only)
	IsMailboxSetup bool `json:"isMailboxSetup,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// LastLoginTime: User's last login time. (Read-only)
	LastLoginTime string `json:"lastLoginTime,omitempty"`

	// Name: User's name
	Name *UserName `json:"name,omitempty"`

	// NonEditableAliases: List of non editable aliases (Read-only)
	NonEditableAliases []string `json:"nonEditableAliases,omitempty"`

	Notes interface{} `json:"notes,omitempty"`

	// OrgUnitPath: OrgUnit of User
	OrgUnitPath string `json:"orgUnitPath,omitempty"`

	Organizations interface{} `json:"organizations,omitempty"`

	// Password: User's password
	Password string `json:"password,omitempty"`

	Phones interface{} `json:"phones,omitempty"`

	// PrimaryEmail: username of User
	PrimaryEmail string `json:"primaryEmail,omitempty"`

	Relations interface{} `json:"relations,omitempty"`

	// Suspended: Indicates if user is suspended
	Suspended bool `json:"suspended,omitempty"`

	// SuspensionReason: Suspension reason if user is suspended (Read-only)
	SuspensionReason string `json:"suspensionReason,omitempty"`

	// ThumbnailPhotoUrl: Photo Url of the user (Read-only)
	ThumbnailPhotoUrl string `json:"thumbnailPhotoUrl,omitempty"`

	Websites interface{} `json:"websites,omitempty"`
}

type UserAbout struct {
	// ContentType: About entry can have a type which indicates the content
	// type. It can either be plain or html. By default, notes contents are
	// assumed to contain plain text.
	ContentType string `json:"contentType,omitempty"`

	// Value: Actual value of notes.
	Value string `json:"value,omitempty"`
}

type UserAddress struct {
	// Country: Country.
	Country string `json:"country,omitempty"`

	// CountryCode: Country code.
	CountryCode string `json:"countryCode,omitempty"`

	// CustomType: Custom type.
	CustomType string `json:"customType,omitempty"`

	// ExtendedAddress: Extended Address.
	ExtendedAddress string `json:"extendedAddress,omitempty"`

	// Formatted: Formatted address.
	Formatted string `json:"formatted,omitempty"`

	// Locality: Locality.
	Locality string `json:"locality,omitempty"`

	// PoBox: Other parts of address.
	PoBox string `json:"poBox,omitempty"`

	// PostalCode: Postal code.
	PostalCode string `json:"postalCode,omitempty"`

	// Primary: If this is user's primary address. Only one entry could be
	// marked as primary.
	Primary bool `json:"primary,omitempty"`

	// Region: Region.
	Region string `json:"region,omitempty"`

	// SourceIsStructured: User supplied address was structured. Structured
	// addresses are NOT supported at this time. You might be able to write
	// structured addresses, but any values will eventually be clobbered.
	SourceIsStructured bool `json:"sourceIsStructured,omitempty"`

	// StreetAddress: Street.
	StreetAddress string `json:"streetAddress,omitempty"`

	// Type: Each entry can have a type which indicates standard values of
	// that entry. For example address could be of home, work etc. In
	// addition to the standard type, an entry can have a custom type and
	// can take any value. Such type should have the CUSTOM value as type
	// and also have a customType value.
	Type string `json:"type,omitempty"`
}

type UserCustomProperties struct {
}

type UserEmail struct {
	// Address: Email id of the user.
	Address string `json:"address,omitempty"`

	// CustomType: Custom Type.
	CustomType string `json:"customType,omitempty"`

	// Primary: If this is user's primary email. Only one entry could be
	// marked as primary.
	Primary bool `json:"primary,omitempty"`

	// Type: Each entry can have a type which indicates standard types of
	// that entry. For example email could be of home, work etc. In addition
	// to the standard type, an entry can have a custom type and can take
	// any value Such types should have the CUSTOM value as type and also
	// have a customType value.
	Type string `json:"type,omitempty"`
}

type UserExternalId struct {
	// CustomType: Custom type.
	CustomType string `json:"customType,omitempty"`

	// Type: The type of the Id.
	Type string `json:"type,omitempty"`

	// Value: The value of the id.
	Value string `json:"value,omitempty"`
}

type UserIm struct {
	// CustomProtocol: Custom protocol.
	CustomProtocol string `json:"customProtocol,omitempty"`

	// CustomType: Custom type.
	CustomType string `json:"customType,omitempty"`

	// Im: Instant messenger id.
	Im string `json:"im,omitempty"`

	// Primary: If this is user's primary im. Only one entry could be marked
	// as primary.
	Primary bool `json:"primary,omitempty"`

	// Protocol: Protocol used in the instant messenger. It should be one of
	// the values from ImProtocolTypes map. Similar to type, it can take a
	// CUSTOM value and specify the custom name in customProtocol field.
	Protocol string `json:"protocol,omitempty"`

	// Type: Each entry can have a type which indicates standard types of
	// that entry. For example instant messengers could be of home, work
	// etc. In addition to the standard type, an entry can have a custom
	// type and can take any value. Such types should have the CUSTOM value
	// as type and also have a customType value.
	Type string `json:"type,omitempty"`
}

type UserMakeAdmin struct {
	// Status: Boolean indicating new admin status of the user
	Status bool `json:"status,omitempty"`
}

type UserName struct {
	// FamilyName: Last Name
	FamilyName string `json:"familyName,omitempty"`

	// FullName: Full Name
	FullName string `json:"fullName,omitempty"`

	// GivenName: First Name
	GivenName string `json:"givenName,omitempty"`
}

type UserOrganization struct {
	// CostCenter: The cost center of the users department.
	CostCenter string `json:"costCenter,omitempty"`

	// CustomType: Custom type.
	CustomType string `json:"customType,omitempty"`

	// Department: Department within the organization.
	Department string `json:"department,omitempty"`

	// Description: Description of the organization.
	Description string `json:"description,omitempty"`

	// Domain: The domain to which the organization belongs to.
	Domain string `json:"domain,omitempty"`

	// Location: Location of the organization. This need not be fully
	// qualified address.
	Location string `json:"location,omitempty"`

	// Name: Name of the organization
	Name string `json:"name,omitempty"`

	// Primary: If it user's primary organization.
	Primary bool `json:"primary,omitempty"`

	// Symbol: Symbol of the organization.
	Symbol string `json:"symbol,omitempty"`

	// Title: Title (designation) of the user in the organization.
	Title string `json:"title,omitempty"`

	// Type: Each entry can have a type which indicates standard types of
	// that entry. For example organization could be of school, work etc. In
	// addition to the standard type, an entry can have a custom type and
	// can give it any name. Such types should have the CUSTOM value as type
	// and also have a CustomType value.
	Type string `json:"type,omitempty"`
}

type UserPhone struct {
	// CustomType: Custom Type.
	CustomType string `json:"customType,omitempty"`

	// Primary: If this is user's primary phone or not.
	Primary bool `json:"primary,omitempty"`

	// Type: Each entry can have a type which indicates standard types of
	// that entry. For example phone could be of home_fax, work, mobile etc.
	// In addition to the standard type, an entry can have a custom type and
	// can give it any name. Such types should have the CUSTOM value as type
	// and also have a customType value.
	Type string `json:"type,omitempty"`

	// Value: Phone number.
	Value string `json:"value,omitempty"`
}

type UserPhoto struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Height: Height in pixels of the photo
	Height int64 `json:"height,omitempty"`

	// Id: Unique identifier of User (Read-only)
	Id string `json:"id,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// MimeType: Mime Type of the photo
	MimeType string `json:"mimeType,omitempty"`

	// PhotoData: Base64 encoded photo data
	PhotoData string `json:"photoData,omitempty"`

	// PrimaryEmail: Primary email of User (Read-only)
	PrimaryEmail string `json:"primaryEmail,omitempty"`

	// Width: Width in pixels of the photo
	Width int64 `json:"width,omitempty"`
}

type UserRelation struct {
	// CustomType: Custom Type.
	CustomType string `json:"customType,omitempty"`

	// Type: The relation of the user. Some of the possible values are
	// mother, father, sister, brother, manager, assistant, partner.
	Type string `json:"type,omitempty"`

	// Value: The name of the relation.
	Value string `json:"value,omitempty"`
}

type UserUndelete struct {
	// OrgUnitPath: OrgUnit of User
	OrgUnitPath string `json:"orgUnitPath,omitempty"`
}

type UserWebsite struct {
	// CustomType: Custom Type.
	CustomType string `json:"customType,omitempty"`

	// Primary: If this is user's primary website or not.
	Primary bool `json:"primary,omitempty"`

	// Type: Each entry can have a type which indicates standard types of
	// that entry. For example website could be of home, work, blog etc. In
	// addition to the standard type, an entry can have a custom type and
	// can give it any name. Such types should have the CUSTOM value as type
	// and also have a customType value.
	Type string `json:"type,omitempty"`

	// Value: Website.
	Value string `json:"value,omitempty"`
}

type Users struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: Kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Token used to access next page of this result.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Trigger_event: Event that triggered this response (only used in case
	// of Push Response)
	Trigger_event string `json:"trigger_event,omitempty"`

	// Users: List of user objects.
	Users []*User `json:"users,omitempty"`
}

type VerificationCode struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Kind: The type of the resource. This is always
	// admin#directory#verificationCode.
	Kind string `json:"kind,omitempty"`

	// UserId: The obfuscated unique ID of the user.
	UserId string `json:"userId,omitempty"`

	// VerificationCode: A current verification code for the user.
	// Invalidated or used verification codes are not returned as part of
	// the result.
	VerificationCode string `json:"verificationCode,omitempty"`
}

type VerificationCodes struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: A list of verification code resources.
	Items []*VerificationCode `json:"items,omitempty"`

	// Kind: The type of the resource. This is always
	// admin#directory#verificationCodesList.
	Kind string `json:"kind,omitempty"`
}

// method id "directory.asps.delete":

type AspsDeleteCall struct {
	s             *Service
	userKey       string
	codeId        int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete an ASP issued by a user.

func (r *AspsService) Delete(userKey string, codeId int64) *AspsDeleteCall {
	return &AspsDeleteCall{
		s:             r.s,
		userKey:       userKey,
		codeId:        codeId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/asps/{codeId}",
	}
}

func (c *AspsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
		"codeId":  strconv.FormatInt(c.codeId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete an ASP issued by a user.",
	//   "httpMethod": "DELETE",
	//   "id": "directory.asps.delete",
	//   "parameterOrder": [
	//     "userKey",
	//     "codeId"
	//   ],
	//   "parameters": {
	//     "codeId": {
	//       "description": "The unique ID of the ASP to be deleted.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/asps/{codeId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.asps.get":

type AspsGetCall struct {
	s             *Service
	userKey       string
	codeId        int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about an ASP issued by a user.

func (r *AspsService) Get(userKey string, codeId int64) *AspsGetCall {
	return &AspsGetCall{
		s:             r.s,
		userKey:       userKey,
		codeId:        codeId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/asps/{codeId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AspsGetCall) Fields(s ...googleapi.Field) *AspsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AspsGetCall) Do() (*Asp, error) {
	var returnValue *Asp
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
		"codeId":  strconv.FormatInt(c.codeId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get information about an ASP issued by a user.",
	//   "httpMethod": "GET",
	//   "id": "directory.asps.get",
	//   "parameterOrder": [
	//     "userKey",
	//     "codeId"
	//   ],
	//   "parameters": {
	//     "codeId": {
	//       "description": "The unique ID of the ASP.",
	//       "format": "int32",
	//       "location": "path",
	//       "required": true,
	//       "type": "integer"
	//     },
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/asps/{codeId}",
	//   "response": {
	//     "$ref": "Asp"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.asps.list":

type AspsListCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List the ASPs issued by a user.

func (r *AspsService) List(userKey string) *AspsListCall {
	return &AspsListCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/asps",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AspsListCall) Fields(s ...googleapi.Field) *AspsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AspsListCall) Do() (*Asps, error) {
	var returnValue *Asps
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List the ASPs issued by a user.",
	//   "httpMethod": "GET",
	//   "id": "directory.asps.list",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/asps",
	//   "response": {
	//     "$ref": "Asps"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "admin.channels.stop":

type ChannelsStopCall struct {
	s             *Service
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Stop: Stop watching resources through this channel

func (r *ChannelsService) Stop(channel *Channel) *ChannelsStopCall {
	return &ChannelsStopCall{
		s:             r.s,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "/admin/directory_v1/channels/stop",
	}
}

func (c *ChannelsStopCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Stop watching resources through this channel",
	//   "httpMethod": "POST",
	//   "id": "admin.channels.stop",
	//   "path": "/admin/directory_v1/channels/stop",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias.readonly",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ]
	// }

}

// method id "directory.chromeosdevices.get":

type ChromeosdevicesGetCall struct {
	s             *Service
	customerId    string
	deviceId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve Chrome OS Device

func (r *ChromeosdevicesService) Get(customerId string, deviceId string) *ChromeosdevicesGetCall {
	return &ChromeosdevicesGetCall{
		s:             r.s,
		customerId:    customerId,
		deviceId:      deviceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/devices/chromeos/{deviceId}",
	}
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *ChromeosdevicesGetCall) Projection(projection string) *ChromeosdevicesGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChromeosdevicesGetCall) Fields(s ...googleapi.Field) *ChromeosdevicesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChromeosdevicesGetCall) Do() (*ChromeOsDevice, error) {
	var returnValue *ChromeOsDevice
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"deviceId":   c.deviceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve Chrome OS Device",
	//   "httpMethod": "GET",
	//   "id": "directory.chromeosdevices.get",
	//   "parameterOrder": [
	//     "customerId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deviceId": {
	//       "description": "Immutable id of Chrome OS Device",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Includes only the basic metadata fields (e.g., deviceId, serialNumber, status, and user)",
	//         "Includes all metadata fields"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/chromeos/{deviceId}",
	//   "response": {
	//     "$ref": "ChromeOsDevice"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.chromeos",
	//     "https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly"
	//   ]
	// }

}

// method id "directory.chromeosdevices.list":

type ChromeosdevicesListCall struct {
	s             *Service
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve all Chrome OS Devices of a customer (paginated)

func (r *ChromeosdevicesService) List(customerId string) *ChromeosdevicesListCall {
	return &ChromeosdevicesListCall{
		s:             r.s,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/devices/chromeos",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 100
func (c *ChromeosdevicesListCall) MaxResults(maxResults int64) *ChromeosdevicesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Column to use for
// sorting results
func (c *ChromeosdevicesListCall) OrderBy(orderBy string) *ChromeosdevicesListCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *ChromeosdevicesListCall) PageToken(pageToken string) *ChromeosdevicesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *ChromeosdevicesListCall) Projection(projection string) *ChromeosdevicesListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Query sets the optional parameter "query": Search string in the
// format given at
// http://support.google.com/chromeos/a/bin/answer.py?hl=en&answer=169833
// 3
func (c *ChromeosdevicesListCall) Query(query string) *ChromeosdevicesListCall {
	c.params_.Set("query", fmt.Sprintf("%v", query))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Whether to return
// results in ascending or descending order. Only of use when orderBy is
// also used
func (c *ChromeosdevicesListCall) SortOrder(sortOrder string) *ChromeosdevicesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChromeosdevicesListCall) Fields(s ...googleapi.Field) *ChromeosdevicesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChromeosdevicesListCall) Do() (*ChromeOsDevices, error) {
	var returnValue *ChromeOsDevices
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve all Chrome OS Devices of a customer (paginated)",
	//   "httpMethod": "GET",
	//   "id": "directory.chromeosdevices.list",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 100",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Column to use for sorting results",
	//       "enum": [
	//         "annotatedLocation",
	//         "annotatedUser",
	//         "lastSync",
	//         "notes",
	//         "serialNumber",
	//         "status",
	//         "supportEndDate"
	//       ],
	//       "enumDescriptions": [
	//         "Chromebook location as annotated by the administrator.",
	//         "Chromebook user as annotated by administrator.",
	//         "Chromebook last sync.",
	//         "Chromebook notes as annotated by the administrator.",
	//         "Chromebook Serial Number.",
	//         "Chromebook status.",
	//         "Chromebook support end date."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Includes only the basic metadata fields (e.g., deviceId, serialNumber, status, and user)",
	//         "Includes all metadata fields"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Search string in the format given at http://support.google.com/chromeos/a/bin/answer.py?hl=en\u0026answer=1698333",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Whether to return results in ascending or descending order. Only of use when orderBy is also used",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "Ascending order.",
	//         "Descending order."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/chromeos",
	//   "response": {
	//     "$ref": "ChromeOsDevices"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.chromeos",
	//     "https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly"
	//   ]
	// }

}

// method id "directory.chromeosdevices.patch":

type ChromeosdevicesPatchCall struct {
	s              *Service
	customerId     string
	deviceId       string
	chromeosdevice *ChromeOsDevice
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Patch: Update Chrome OS Device. This method supports patch semantics.

func (r *ChromeosdevicesService) Patch(customerId string, deviceId string, chromeosdevice *ChromeOsDevice) *ChromeosdevicesPatchCall {
	return &ChromeosdevicesPatchCall{
		s:              r.s,
		customerId:     customerId,
		deviceId:       deviceId,
		chromeosdevice: chromeosdevice,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customer/{customerId}/devices/chromeos/{deviceId}",
	}
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *ChromeosdevicesPatchCall) Projection(projection string) *ChromeosdevicesPatchCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChromeosdevicesPatchCall) Fields(s ...googleapi.Field) *ChromeosdevicesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChromeosdevicesPatchCall) Do() (*ChromeOsDevice, error) {
	var returnValue *ChromeOsDevice
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"deviceId":   c.deviceId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.chromeosdevice,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update Chrome OS Device. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.chromeosdevices.patch",
	//   "parameterOrder": [
	//     "customerId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deviceId": {
	//       "description": "Immutable id of Chrome OS Device",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Includes only the basic metadata fields (e.g., deviceId, serialNumber, status, and user)",
	//         "Includes all metadata fields"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/chromeos/{deviceId}",
	//   "request": {
	//     "$ref": "ChromeOsDevice"
	//   },
	//   "response": {
	//     "$ref": "ChromeOsDevice"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.chromeos"
	//   ]
	// }

}

// method id "directory.chromeosdevices.update":

type ChromeosdevicesUpdateCall struct {
	s              *Service
	customerId     string
	deviceId       string
	chromeosdevice *ChromeOsDevice
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Update: Update Chrome OS Device

func (r *ChromeosdevicesService) Update(customerId string, deviceId string, chromeosdevice *ChromeOsDevice) *ChromeosdevicesUpdateCall {
	return &ChromeosdevicesUpdateCall{
		s:              r.s,
		customerId:     customerId,
		deviceId:       deviceId,
		chromeosdevice: chromeosdevice,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customer/{customerId}/devices/chromeos/{deviceId}",
	}
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *ChromeosdevicesUpdateCall) Projection(projection string) *ChromeosdevicesUpdateCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChromeosdevicesUpdateCall) Fields(s ...googleapi.Field) *ChromeosdevicesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChromeosdevicesUpdateCall) Do() (*ChromeOsDevice, error) {
	var returnValue *ChromeOsDevice
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"deviceId":   c.deviceId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.chromeosdevice,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update Chrome OS Device",
	//   "httpMethod": "PUT",
	//   "id": "directory.chromeosdevices.update",
	//   "parameterOrder": [
	//     "customerId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deviceId": {
	//       "description": "Immutable id of Chrome OS Device",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Includes only the basic metadata fields (e.g., deviceId, serialNumber, status, and user)",
	//         "Includes all metadata fields"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/chromeos/{deviceId}",
	//   "request": {
	//     "$ref": "ChromeOsDevice"
	//   },
	//   "response": {
	//     "$ref": "ChromeOsDevice"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.chromeos"
	//   ]
	// }

}

// method id "directory.groups.delete":

type GroupsDeleteCall struct {
	s             *Service
	groupKey      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete Group

func (r *GroupsService) Delete(groupKey string) *GroupsDeleteCall {
	return &GroupsDeleteCall{
		s:             r.s,
		groupKey:      groupKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}",
	}
}

func (c *GroupsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete Group",
	//   "httpMethod": "DELETE",
	//   "id": "directory.groups.delete",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group"
	//   ]
	// }

}

// method id "directory.groups.get":

type GroupsGetCall struct {
	s             *Service
	groupKey      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve Group

func (r *GroupsService) Get(groupKey string) *GroupsGetCall {
	return &GroupsGetCall{
		s:             r.s,
		groupKey:      groupKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsGetCall) Fields(s ...googleapi.Field) *GroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsGetCall) Do() (*Group, error) {
	var returnValue *Group
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve Group",
	//   "httpMethod": "GET",
	//   "id": "directory.groups.get",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}",
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.readonly"
	//   ]
	// }

}

// method id "directory.groups.insert":

type GroupsInsertCall struct {
	s             *Service
	group         *Group
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Create Group

func (r *GroupsService) Insert(group *Group) *GroupsInsertCall {
	return &GroupsInsertCall{
		s:             r.s,
		group:         group,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsInsertCall) Fields(s ...googleapi.Field) *GroupsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsInsertCall) Do() (*Group, error) {
	var returnValue *Group
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.group,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create Group",
	//   "httpMethod": "POST",
	//   "id": "directory.groups.insert",
	//   "path": "groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group"
	//   ]
	// }

}

// method id "directory.groups.list":

type GroupsListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve all groups in a domain (paginated)

func (r *GroupsService) List() *GroupsListCall {
	return &GroupsListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups",
	}
}

// Customer sets the optional parameter "customer": Immutable id of the
// Google Apps account. In case of multi-domain, to fetch all groups for
// a customer, fill this field instead of domain.
func (c *GroupsListCall) Customer(customer string) *GroupsListCall {
	c.params_.Set("customer", fmt.Sprintf("%v", customer))
	return c
}

// Domain sets the optional parameter "domain": Name of the domain. Fill
// this field to get groups from only this domain. To return all groups
// in a multi-domain fill customer field instead.
func (c *GroupsListCall) Domain(domain string) *GroupsListCall {
	c.params_.Set("domain", fmt.Sprintf("%v", domain))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 200
func (c *GroupsListCall) MaxResults(maxResults int64) *GroupsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *GroupsListCall) PageToken(pageToken string) *GroupsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// UserKey sets the optional parameter "userKey": Email or immutable Id
// of the user if only those groups are to be listed, the given user is
// a member of. If Id, it should match with id of user object
func (c *GroupsListCall) UserKey(userKey string) *GroupsListCall {
	c.params_.Set("userKey", fmt.Sprintf("%v", userKey))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsListCall) Fields(s ...googleapi.Field) *GroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsListCall) Do() (*Groups, error) {
	var returnValue *Groups
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve all groups in a domain (paginated)",
	//   "httpMethod": "GET",
	//   "id": "directory.groups.list",
	//   "parameters": {
	//     "customer": {
	//       "description": "Immutable id of the Google Apps account. In case of multi-domain, to fetch all groups for a customer, fill this field instead of domain.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "domain": {
	//       "description": "Name of the domain. Fill this field to get groups from only this domain. To return all groups in a multi-domain fill customer field instead.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 200",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Email or immutable Id of the user if only those groups are to be listed, the given user is a member of. If Id, it should match with id of user object",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups",
	//   "response": {
	//     "$ref": "Groups"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.readonly"
	//   ]
	// }

}

// method id "directory.groups.patch":

type GroupsPatchCall struct {
	s             *Service
	groupKey      string
	group         *Group
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update Group. This method supports patch semantics.

func (r *GroupsService) Patch(groupKey string, group *Group) *GroupsPatchCall {
	return &GroupsPatchCall{
		s:             r.s,
		groupKey:      groupKey,
		group:         group,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsPatchCall) Fields(s ...googleapi.Field) *GroupsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsPatchCall) Do() (*Group, error) {
	var returnValue *Group
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.group,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update Group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.groups.patch",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group. If Id, it should match with id of group object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group"
	//   ]
	// }

}

// method id "directory.groups.update":

type GroupsUpdateCall struct {
	s             *Service
	groupKey      string
	group         *Group
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update Group

func (r *GroupsService) Update(groupKey string, group *Group) *GroupsUpdateCall {
	return &GroupsUpdateCall{
		s:             r.s,
		groupKey:      groupKey,
		group:         group,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsUpdateCall) Fields(s ...googleapi.Field) *GroupsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsUpdateCall) Do() (*Group, error) {
	var returnValue *Group
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.group,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update Group",
	//   "httpMethod": "PUT",
	//   "id": "directory.groups.update",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group. If Id, it should match with id of group object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group"
	//   ]
	// }

}

// method id "directory.groups.aliases.delete":

type GroupsAliasesDeleteCall struct {
	s             *Service
	groupKey      string
	alias         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Remove a alias for the group

func (r *GroupsAliasesService) Delete(groupKey string, alias string) *GroupsAliasesDeleteCall {
	return &GroupsAliasesDeleteCall{
		s:             r.s,
		groupKey:      groupKey,
		alias:         alias,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/aliases/{alias}",
	}
}

func (c *GroupsAliasesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
		"alias":    c.alias,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Remove a alias for the group",
	//   "httpMethod": "DELETE",
	//   "id": "directory.groups.aliases.delete",
	//   "parameterOrder": [
	//     "groupKey",
	//     "alias"
	//   ],
	//   "parameters": {
	//     "alias": {
	//       "description": "The alias to be removed",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/aliases/{alias}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group"
	//   ]
	// }

}

// method id "directory.groups.aliases.insert":

type GroupsAliasesInsertCall struct {
	s             *Service
	groupKey      string
	alias         *Alias
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Add a alias for the group

func (r *GroupsAliasesService) Insert(groupKey string, alias *Alias) *GroupsAliasesInsertCall {
	return &GroupsAliasesInsertCall{
		s:             r.s,
		groupKey:      groupKey,
		alias:         alias,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/aliases",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsAliasesInsertCall) Fields(s ...googleapi.Field) *GroupsAliasesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsAliasesInsertCall) Do() (*Alias, error) {
	var returnValue *Alias
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.alias,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add a alias for the group",
	//   "httpMethod": "POST",
	//   "id": "directory.groups.aliases.insert",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/aliases",
	//   "request": {
	//     "$ref": "Alias"
	//   },
	//   "response": {
	//     "$ref": "Alias"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group"
	//   ]
	// }

}

// method id "directory.groups.aliases.list":

type GroupsAliasesListCall struct {
	s             *Service
	groupKey      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all aliases for a group

func (r *GroupsAliasesService) List(groupKey string) *GroupsAliasesListCall {
	return &GroupsAliasesListCall{
		s:             r.s,
		groupKey:      groupKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/aliases",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsAliasesListCall) Fields(s ...googleapi.Field) *GroupsAliasesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *GroupsAliasesListCall) Do() (*Aliases, error) {
	var returnValue *Aliases
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all aliases for a group",
	//   "httpMethod": "GET",
	//   "id": "directory.groups.aliases.list",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/aliases",
	//   "response": {
	//     "$ref": "Aliases"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "directory.members.delete":

type MembersDeleteCall struct {
	s             *Service
	groupKey      string
	memberKey     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Remove membership.

func (r *MembersService) Delete(groupKey string, memberKey string) *MembersDeleteCall {
	return &MembersDeleteCall{
		s:             r.s,
		groupKey:      groupKey,
		memberKey:     memberKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/members/{memberKey}",
	}
}

func (c *MembersDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey":  c.groupKey,
		"memberKey": c.memberKey,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Remove membership.",
	//   "httpMethod": "DELETE",
	//   "id": "directory.members.delete",
	//   "parameterOrder": [
	//     "groupKey",
	//     "memberKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "memberKey": {
	//       "description": "Email or immutable Id of the member",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/members/{memberKey}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.member"
	//   ]
	// }

}

// method id "directory.members.get":

type MembersGetCall struct {
	s             *Service
	groupKey      string
	memberKey     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve Group Member

func (r *MembersService) Get(groupKey string, memberKey string) *MembersGetCall {
	return &MembersGetCall{
		s:             r.s,
		groupKey:      groupKey,
		memberKey:     memberKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/members/{memberKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MembersGetCall) Fields(s ...googleapi.Field) *MembersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MembersGetCall) Do() (*Member, error) {
	var returnValue *Member
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey":  c.groupKey,
		"memberKey": c.memberKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve Group Member",
	//   "httpMethod": "GET",
	//   "id": "directory.members.get",
	//   "parameterOrder": [
	//     "groupKey",
	//     "memberKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "memberKey": {
	//       "description": "Email or immutable Id of the member",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/members/{memberKey}",
	//   "response": {
	//     "$ref": "Member"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.member",
	//     "https://www.googleapis.com/auth/admin.directory.group.member.readonly",
	//     "https://www.googleapis.com/auth/admin.directory.group.readonly"
	//   ]
	// }

}

// method id "directory.members.insert":

type MembersInsertCall struct {
	s             *Service
	groupKey      string
	member        *Member
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Add user to the specified group.

func (r *MembersService) Insert(groupKey string, member *Member) *MembersInsertCall {
	return &MembersInsertCall{
		s:             r.s,
		groupKey:      groupKey,
		member:        member,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/members",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MembersInsertCall) Fields(s ...googleapi.Field) *MembersInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MembersInsertCall) Do() (*Member, error) {
	var returnValue *Member
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.member,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add user to the specified group.",
	//   "httpMethod": "POST",
	//   "id": "directory.members.insert",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/members",
	//   "request": {
	//     "$ref": "Member"
	//   },
	//   "response": {
	//     "$ref": "Member"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.member"
	//   ]
	// }

}

// method id "directory.members.list":

type MembersListCall struct {
	s             *Service
	groupKey      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve all members in a group (paginated)

func (r *MembersService) List(groupKey string) *MembersListCall {
	return &MembersListCall{
		s:             r.s,
		groupKey:      groupKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/members",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 200
func (c *MembersListCall) MaxResults(maxResults int64) *MembersListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *MembersListCall) PageToken(pageToken string) *MembersListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Roles sets the optional parameter "roles": Comma separated role
// values to filter list results on.
func (c *MembersListCall) Roles(roles string) *MembersListCall {
	c.params_.Set("roles", fmt.Sprintf("%v", roles))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MembersListCall) Fields(s ...googleapi.Field) *MembersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MembersListCall) Do() (*Members, error) {
	var returnValue *Members
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey": c.groupKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve all members in a group (paginated)",
	//   "httpMethod": "GET",
	//   "id": "directory.members.list",
	//   "parameterOrder": [
	//     "groupKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 200",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "roles": {
	//       "description": "Comma separated role values to filter list results on.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/members",
	//   "response": {
	//     "$ref": "Members"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.member",
	//     "https://www.googleapis.com/auth/admin.directory.group.member.readonly",
	//     "https://www.googleapis.com/auth/admin.directory.group.readonly"
	//   ]
	// }

}

// method id "directory.members.patch":

type MembersPatchCall struct {
	s             *Service
	groupKey      string
	memberKey     string
	member        *Member
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update membership of a user in the specified group. This
// method supports patch semantics.

func (r *MembersService) Patch(groupKey string, memberKey string, member *Member) *MembersPatchCall {
	return &MembersPatchCall{
		s:             r.s,
		groupKey:      groupKey,
		memberKey:     memberKey,
		member:        member,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/members/{memberKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MembersPatchCall) Fields(s ...googleapi.Field) *MembersPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MembersPatchCall) Do() (*Member, error) {
	var returnValue *Member
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey":  c.groupKey,
		"memberKey": c.memberKey,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.member,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update membership of a user in the specified group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.members.patch",
	//   "parameterOrder": [
	//     "groupKey",
	//     "memberKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group. If Id, it should match with id of group object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "memberKey": {
	//       "description": "Email or immutable Id of the user. If Id, it should match with id of member object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/members/{memberKey}",
	//   "request": {
	//     "$ref": "Member"
	//   },
	//   "response": {
	//     "$ref": "Member"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.member"
	//   ]
	// }

}

// method id "directory.members.update":

type MembersUpdateCall struct {
	s             *Service
	groupKey      string
	memberKey     string
	member        *Member
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update membership of a user in the specified group.

func (r *MembersService) Update(groupKey string, memberKey string, member *Member) *MembersUpdateCall {
	return &MembersUpdateCall{
		s:             r.s,
		groupKey:      groupKey,
		memberKey:     memberKey,
		member:        member,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "groups/{groupKey}/members/{memberKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MembersUpdateCall) Fields(s ...googleapi.Field) *MembersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MembersUpdateCall) Do() (*Member, error) {
	var returnValue *Member
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"groupKey":  c.groupKey,
		"memberKey": c.memberKey,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.member,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update membership of a user in the specified group.",
	//   "httpMethod": "PUT",
	//   "id": "directory.members.update",
	//   "parameterOrder": [
	//     "groupKey",
	//     "memberKey"
	//   ],
	//   "parameters": {
	//     "groupKey": {
	//       "description": "Email or immutable Id of the group. If Id, it should match with id of group object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "memberKey": {
	//       "description": "Email or immutable Id of the user. If Id, it should match with id of member object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "groups/{groupKey}/members/{memberKey}",
	//   "request": {
	//     "$ref": "Member"
	//   },
	//   "response": {
	//     "$ref": "Member"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.group",
	//     "https://www.googleapis.com/auth/admin.directory.group.member"
	//   ]
	// }

}

// method id "directory.mobiledevices.action":

type MobiledevicesActionCall struct {
	s                  *Service
	customerId         string
	resourceId         string
	mobiledeviceaction *MobileDeviceAction
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Action: Take action on Mobile Device

func (r *MobiledevicesService) Action(customerId string, resourceId string, mobiledeviceaction *MobileDeviceAction) *MobiledevicesActionCall {
	return &MobiledevicesActionCall{
		s:                  r.s,
		customerId:         customerId,
		resourceId:         resourceId,
		mobiledeviceaction: mobiledeviceaction,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "customer/{customerId}/devices/mobile/{resourceId}/action",
	}
}

func (c *MobiledevicesActionCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"resourceId": c.resourceId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.mobiledeviceaction,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Take action on Mobile Device",
	//   "httpMethod": "POST",
	//   "id": "directory.mobiledevices.action",
	//   "parameterOrder": [
	//     "customerId",
	//     "resourceId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceId": {
	//       "description": "Immutable id of Mobile Device",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/mobile/{resourceId}/action",
	//   "request": {
	//     "$ref": "MobileDeviceAction"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile",
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile.action"
	//   ]
	// }

}

// method id "directory.mobiledevices.delete":

type MobiledevicesDeleteCall struct {
	s             *Service
	customerId    string
	resourceId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete Mobile Device

func (r *MobiledevicesService) Delete(customerId string, resourceId string) *MobiledevicesDeleteCall {
	return &MobiledevicesDeleteCall{
		s:             r.s,
		customerId:    customerId,
		resourceId:    resourceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/devices/mobile/{resourceId}",
	}
}

func (c *MobiledevicesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"resourceId": c.resourceId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete Mobile Device",
	//   "httpMethod": "DELETE",
	//   "id": "directory.mobiledevices.delete",
	//   "parameterOrder": [
	//     "customerId",
	//     "resourceId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceId": {
	//       "description": "Immutable id of Mobile Device",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/mobile/{resourceId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile"
	//   ]
	// }

}

// method id "directory.mobiledevices.get":

type MobiledevicesGetCall struct {
	s             *Service
	customerId    string
	resourceId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve Mobile Device

func (r *MobiledevicesService) Get(customerId string, resourceId string) *MobiledevicesGetCall {
	return &MobiledevicesGetCall{
		s:             r.s,
		customerId:    customerId,
		resourceId:    resourceId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/devices/mobile/{resourceId}",
	}
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *MobiledevicesGetCall) Projection(projection string) *MobiledevicesGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MobiledevicesGetCall) Fields(s ...googleapi.Field) *MobiledevicesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MobiledevicesGetCall) Do() (*MobileDevice, error) {
	var returnValue *MobileDevice
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"resourceId": c.resourceId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve Mobile Device",
	//   "httpMethod": "GET",
	//   "id": "directory.mobiledevices.get",
	//   "parameterOrder": [
	//     "customerId",
	//     "resourceId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Includes only the basic metadata fields (e.g., deviceId, model, status, type, and status)",
	//         "Includes all metadata fields"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceId": {
	//       "description": "Immutable id of Mobile Device",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/mobile/{resourceId}",
	//   "response": {
	//     "$ref": "MobileDevice"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile",
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile.action",
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile.readonly"
	//   ]
	// }

}

// method id "directory.mobiledevices.list":

type MobiledevicesListCall struct {
	s             *Service
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve all Mobile Devices of a customer (paginated)

func (r *MobiledevicesService) List(customerId string) *MobiledevicesListCall {
	return &MobiledevicesListCall{
		s:             r.s,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/devices/mobile",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 100
func (c *MobiledevicesListCall) MaxResults(maxResults int64) *MobiledevicesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Column to use for
// sorting results
func (c *MobiledevicesListCall) OrderBy(orderBy string) *MobiledevicesListCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *MobiledevicesListCall) PageToken(pageToken string) *MobiledevicesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Projection sets the optional parameter "projection": Restrict
// information returned to a set of selected fields.
func (c *MobiledevicesListCall) Projection(projection string) *MobiledevicesListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Query sets the optional parameter "query": Search string in the
// format given at
// http://support.google.com/a/bin/answer.py?hl=en&answer=1408863#search
func (c *MobiledevicesListCall) Query(query string) *MobiledevicesListCall {
	c.params_.Set("query", fmt.Sprintf("%v", query))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Whether to return
// results in ascending or descending order. Only of use when orderBy is
// also used
func (c *MobiledevicesListCall) SortOrder(sortOrder string) *MobiledevicesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MobiledevicesListCall) Fields(s ...googleapi.Field) *MobiledevicesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MobiledevicesListCall) Do() (*MobileDevices, error) {
	var returnValue *MobileDevices
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve all Mobile Devices of a customer (paginated)",
	//   "httpMethod": "GET",
	//   "id": "directory.mobiledevices.list",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 100",
	//       "format": "int32",
	//       "location": "query",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Column to use for sorting results",
	//       "enum": [
	//         "deviceId",
	//         "email",
	//         "lastSync",
	//         "model",
	//         "name",
	//         "os",
	//         "status",
	//         "type"
	//       ],
	//       "enumDescriptions": [
	//         "Mobile Device serial number.",
	//         "Owner user email.",
	//         "Last policy settings sync date time of the device.",
	//         "Mobile Device model.",
	//         "Owner user name.",
	//         "Mobile operating system.",
	//         "Status of the device.",
	//         "Type of the device."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "description": "Restrict information returned to a set of selected fields.",
	//       "enum": [
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "enumDescriptions": [
	//         "Includes only the basic metadata fields (e.g., deviceId, model, status, type, and status)",
	//         "Includes all metadata fields"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Search string in the format given at http://support.google.com/a/bin/answer.py?hl=en\u0026answer=1408863#search",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Whether to return results in ascending or descending order. Only of use when orderBy is also used",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "Ascending order.",
	//         "Descending order."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/devices/mobile",
	//   "response": {
	//     "$ref": "MobileDevices"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile",
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile.action",
	//     "https://www.googleapis.com/auth/admin.directory.device.mobile.readonly"
	//   ]
	// }

}

// method id "directory.notifications.delete":

type NotificationsDeleteCall struct {
	s              *Service
	customer       string
	notificationId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Delete: Deletes a notification

func (r *NotificationsService) Delete(customer string, notificationId string) *NotificationsDeleteCall {
	return &NotificationsDeleteCall{
		s:              r.s,
		customer:       customer,
		notificationId: notificationId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customer/{customer}/notifications/{notificationId}",
	}
}

func (c *NotificationsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customer":       c.customer,
		"notificationId": c.notificationId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a notification",
	//   "httpMethod": "DELETE",
	//   "id": "directory.notifications.delete",
	//   "parameterOrder": [
	//     "customer",
	//     "notificationId"
	//   ],
	//   "parameters": {
	//     "customer": {
	//       "description": "The unique ID for the customer's Google account. The customerId is also returned as part of the Users resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "notificationId": {
	//       "description": "The unique ID of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customer}/notifications/{notificationId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.notifications"
	//   ]
	// }

}

// method id "directory.notifications.get":

type NotificationsGetCall struct {
	s              *Service
	customer       string
	notificationId string
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Get: Retrieves a notification.

func (r *NotificationsService) Get(customer string, notificationId string) *NotificationsGetCall {
	return &NotificationsGetCall{
		s:              r.s,
		customer:       customer,
		notificationId: notificationId,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customer/{customer}/notifications/{notificationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsGetCall) Fields(s ...googleapi.Field) *NotificationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NotificationsGetCall) Do() (*Notification, error) {
	var returnValue *Notification
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customer":       c.customer,
		"notificationId": c.notificationId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a notification.",
	//   "httpMethod": "GET",
	//   "id": "directory.notifications.get",
	//   "parameterOrder": [
	//     "customer",
	//     "notificationId"
	//   ],
	//   "parameters": {
	//     "customer": {
	//       "description": "The unique ID for the customer's Google account. The customerId is also returned as part of the Users resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "notificationId": {
	//       "description": "The unique ID of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customer}/notifications/{notificationId}",
	//   "response": {
	//     "$ref": "Notification"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.notifications"
	//   ]
	// }

}

// method id "directory.notifications.list":

type NotificationsListCall struct {
	s             *Service
	customer      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of notifications.

func (r *NotificationsService) List(customer string) *NotificationsListCall {
	return &NotificationsListCall{
		s:             r.s,
		customer:      customer,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customer}/notifications",
	}
}

// Language sets the optional parameter "language": The ISO 639-1 code
// of the language notifications are returned in. The default is English
// (en).
func (c *NotificationsListCall) Language(language string) *NotificationsListCall {
	c.params_.Set("language", fmt.Sprintf("%v", language))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of notifications to return per page. The default is 100.
func (c *NotificationsListCall) MaxResults(maxResults int64) *NotificationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The token to
// specify the page of results to retrieve.
func (c *NotificationsListCall) PageToken(pageToken string) *NotificationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsListCall) Fields(s ...googleapi.Field) *NotificationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NotificationsListCall) Do() (*Notifications, error) {
	var returnValue *Notifications
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customer": c.customer,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of notifications.",
	//   "httpMethod": "GET",
	//   "id": "directory.notifications.list",
	//   "parameterOrder": [
	//     "customer"
	//   ],
	//   "parameters": {
	//     "customer": {
	//       "description": "The unique ID for the customer's Google account.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "language": {
	//       "description": "The ISO 639-1 code of the language notifications are returned in. The default is English (en).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of notifications to return per page. The default is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The token to specify the page of results to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customer}/notifications",
	//   "response": {
	//     "$ref": "Notifications"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.notifications"
	//   ]
	// }

}

// method id "directory.notifications.patch":

type NotificationsPatchCall struct {
	s              *Service
	customer       string
	notificationId string
	notification   *Notification
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Patch: Updates a notification. This method supports patch semantics.

func (r *NotificationsService) Patch(customer string, notificationId string, notification *Notification) *NotificationsPatchCall {
	return &NotificationsPatchCall{
		s:              r.s,
		customer:       customer,
		notificationId: notificationId,
		notification:   notification,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customer/{customer}/notifications/{notificationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsPatchCall) Fields(s ...googleapi.Field) *NotificationsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NotificationsPatchCall) Do() (*Notification, error) {
	var returnValue *Notification
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customer":       c.customer,
		"notificationId": c.notificationId,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.notification,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a notification. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.notifications.patch",
	//   "parameterOrder": [
	//     "customer",
	//     "notificationId"
	//   ],
	//   "parameters": {
	//     "customer": {
	//       "description": "The unique ID for the customer's Google account.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "notificationId": {
	//       "description": "The unique ID of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customer}/notifications/{notificationId}",
	//   "request": {
	//     "$ref": "Notification"
	//   },
	//   "response": {
	//     "$ref": "Notification"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.notifications"
	//   ]
	// }

}

// method id "directory.notifications.update":

type NotificationsUpdateCall struct {
	s              *Service
	customer       string
	notificationId string
	notification   *Notification
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Update: Updates a notification.

func (r *NotificationsService) Update(customer string, notificationId string, notification *Notification) *NotificationsUpdateCall {
	return &NotificationsUpdateCall{
		s:              r.s,
		customer:       customer,
		notificationId: notificationId,
		notification:   notification,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "customer/{customer}/notifications/{notificationId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *NotificationsUpdateCall) Fields(s ...googleapi.Field) *NotificationsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *NotificationsUpdateCall) Do() (*Notification, error) {
	var returnValue *Notification
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customer":       c.customer,
		"notificationId": c.notificationId,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.notification,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a notification.",
	//   "httpMethod": "PUT",
	//   "id": "directory.notifications.update",
	//   "parameterOrder": [
	//     "customer",
	//     "notificationId"
	//   ],
	//   "parameters": {
	//     "customer": {
	//       "description": "The unique ID for the customer's Google account.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "notificationId": {
	//       "description": "The unique ID of the notification.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customer}/notifications/{notificationId}",
	//   "request": {
	//     "$ref": "Notification"
	//   },
	//   "response": {
	//     "$ref": "Notification"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.notifications"
	//   ]
	// }

}

// method id "directory.orgunits.delete":

type OrgunitsDeleteCall struct {
	s             *Service
	customerId    string
	orgUnitPath   []string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Remove Organization Unit

func (r *OrgunitsService) Delete(customerId string, orgUnitPath []string) *OrgunitsDeleteCall {
	return &OrgunitsDeleteCall{
		s:             r.s,
		customerId:    customerId,
		orgUnitPath:   orgUnitPath,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/orgunits{/orgUnitPath*}",
	}
}

func (c *OrgunitsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":  c.customerId,
		"orgUnitPath": c.orgUnitPath[0],
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Remove Organization Unit",
	//   "httpMethod": "DELETE",
	//   "id": "directory.orgunits.delete",
	//   "parameterOrder": [
	//     "customerId",
	//     "orgUnitPath"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orgUnitPath": {
	//       "description": "Full path of the organization unit",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/orgunits{/orgUnitPath*}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.orgunit"
	//   ]
	// }

}

// method id "directory.orgunits.get":

type OrgunitsGetCall struct {
	s             *Service
	customerId    string
	orgUnitPath   []string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve Organization Unit

func (r *OrgunitsService) Get(customerId string, orgUnitPath []string) *OrgunitsGetCall {
	return &OrgunitsGetCall{
		s:             r.s,
		customerId:    customerId,
		orgUnitPath:   orgUnitPath,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/orgunits{/orgUnitPath*}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrgunitsGetCall) Fields(s ...googleapi.Field) *OrgunitsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OrgunitsGetCall) Do() (*OrgUnit, error) {
	var returnValue *OrgUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":  c.customerId,
		"orgUnitPath": c.orgUnitPath[0],
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve Organization Unit",
	//   "httpMethod": "GET",
	//   "id": "directory.orgunits.get",
	//   "parameterOrder": [
	//     "customerId",
	//     "orgUnitPath"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orgUnitPath": {
	//       "description": "Full path of the organization unit",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/orgunits{/orgUnitPath*}",
	//   "response": {
	//     "$ref": "OrgUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.orgunit",
	//     "https://www.googleapis.com/auth/admin.directory.orgunit.readonly"
	//   ]
	// }

}

// method id "directory.orgunits.insert":

type OrgunitsInsertCall struct {
	s             *Service
	customerId    string
	orgunit       *OrgUnit
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Add Organization Unit

func (r *OrgunitsService) Insert(customerId string, orgunit *OrgUnit) *OrgunitsInsertCall {
	return &OrgunitsInsertCall{
		s:             r.s,
		customerId:    customerId,
		orgunit:       orgunit,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/orgunits",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrgunitsInsertCall) Fields(s ...googleapi.Field) *OrgunitsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OrgunitsInsertCall) Do() (*OrgUnit, error) {
	var returnValue *OrgUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.orgunit,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add Organization Unit",
	//   "httpMethod": "POST",
	//   "id": "directory.orgunits.insert",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/orgunits",
	//   "request": {
	//     "$ref": "OrgUnit"
	//   },
	//   "response": {
	//     "$ref": "OrgUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.orgunit"
	//   ]
	// }

}

// method id "directory.orgunits.list":

type OrgunitsListCall struct {
	s             *Service
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve all Organization Units

func (r *OrgunitsService) List(customerId string) *OrgunitsListCall {
	return &OrgunitsListCall{
		s:             r.s,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/orgunits",
	}
}

// OrgUnitPath sets the optional parameter "orgUnitPath": the
// URL-encoded organization unit
func (c *OrgunitsListCall) OrgUnitPath(orgUnitPath string) *OrgunitsListCall {
	c.params_.Set("orgUnitPath", fmt.Sprintf("%v", orgUnitPath))
	return c
}

// Type sets the optional parameter "type": Whether to return all
// sub-organizations or just immediate children
func (c *OrgunitsListCall) Type(type_ string) *OrgunitsListCall {
	c.params_.Set("type", fmt.Sprintf("%v", type_))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrgunitsListCall) Fields(s ...googleapi.Field) *OrgunitsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OrgunitsListCall) Do() (*OrgUnits, error) {
	var returnValue *OrgUnits
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve all Organization Units",
	//   "httpMethod": "GET",
	//   "id": "directory.orgunits.list",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orgUnitPath": {
	//       "default": "",
	//       "description": "the URL-encoded organization unit",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Whether to return all sub-organizations or just immediate children",
	//       "enum": [
	//         "all",
	//         "children"
	//       ],
	//       "enumDescriptions": [
	//         "All sub-organization units.",
	//         "Immediate children only (default)."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/orgunits",
	//   "response": {
	//     "$ref": "OrgUnits"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.orgunit",
	//     "https://www.googleapis.com/auth/admin.directory.orgunit.readonly"
	//   ]
	// }

}

// method id "directory.orgunits.patch":

type OrgunitsPatchCall struct {
	s             *Service
	customerId    string
	orgUnitPath   []string
	orgunit       *OrgUnit
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update Organization Unit. This method supports patch
// semantics.

func (r *OrgunitsService) Patch(customerId string, orgUnitPath []string, orgunit *OrgUnit) *OrgunitsPatchCall {
	return &OrgunitsPatchCall{
		s:             r.s,
		customerId:    customerId,
		orgUnitPath:   orgUnitPath,
		orgunit:       orgunit,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/orgunits{/orgUnitPath*}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrgunitsPatchCall) Fields(s ...googleapi.Field) *OrgunitsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OrgunitsPatchCall) Do() (*OrgUnit, error) {
	var returnValue *OrgUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":  c.customerId,
		"orgUnitPath": c.orgUnitPath[0],
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.orgunit,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update Organization Unit. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.orgunits.patch",
	//   "parameterOrder": [
	//     "customerId",
	//     "orgUnitPath"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orgUnitPath": {
	//       "description": "Full path of the organization unit",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/orgunits{/orgUnitPath*}",
	//   "request": {
	//     "$ref": "OrgUnit"
	//   },
	//   "response": {
	//     "$ref": "OrgUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.orgunit"
	//   ]
	// }

}

// method id "directory.orgunits.update":

type OrgunitsUpdateCall struct {
	s             *Service
	customerId    string
	orgUnitPath   []string
	orgunit       *OrgUnit
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update Organization Unit

func (r *OrgunitsService) Update(customerId string, orgUnitPath []string, orgunit *OrgUnit) *OrgunitsUpdateCall {
	return &OrgunitsUpdateCall{
		s:             r.s,
		customerId:    customerId,
		orgUnitPath:   orgUnitPath,
		orgunit:       orgunit,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/orgunits{/orgUnitPath*}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrgunitsUpdateCall) Fields(s ...googleapi.Field) *OrgunitsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OrgunitsUpdateCall) Do() (*OrgUnit, error) {
	var returnValue *OrgUnit
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId":  c.customerId,
		"orgUnitPath": c.orgUnitPath[0],
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.orgunit,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update Organization Unit",
	//   "httpMethod": "PUT",
	//   "id": "directory.orgunits.update",
	//   "parameterOrder": [
	//     "customerId",
	//     "orgUnitPath"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orgUnitPath": {
	//       "description": "Full path of the organization unit",
	//       "location": "path",
	//       "repeated": true,
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/orgunits{/orgUnitPath*}",
	//   "request": {
	//     "$ref": "OrgUnit"
	//   },
	//   "response": {
	//     "$ref": "OrgUnit"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.orgunit"
	//   ]
	// }

}

// method id "directory.schemas.delete":

type SchemasDeleteCall struct {
	s             *Service
	customerId    string
	schemaKey     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete schema

func (r *SchemasService) Delete(customerId string, schemaKey string) *SchemasDeleteCall {
	return &SchemasDeleteCall{
		s:             r.s,
		customerId:    customerId,
		schemaKey:     schemaKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/schemas/{schemaKey}",
	}
}

func (c *SchemasDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"schemaKey":  c.schemaKey,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete schema",
	//   "httpMethod": "DELETE",
	//   "id": "directory.schemas.delete",
	//   "parameterOrder": [
	//     "customerId",
	//     "schemaKey"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "schemaKey": {
	//       "description": "Name or immutable Id of the schema",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/schemas/{schemaKey}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.userschema"
	//   ]
	// }

}

// method id "directory.schemas.get":

type SchemasGetCall struct {
	s             *Service
	customerId    string
	schemaKey     string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve schema

func (r *SchemasService) Get(customerId string, schemaKey string) *SchemasGetCall {
	return &SchemasGetCall{
		s:             r.s,
		customerId:    customerId,
		schemaKey:     schemaKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/schemas/{schemaKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SchemasGetCall) Fields(s ...googleapi.Field) *SchemasGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SchemasGetCall) Do() (*Schema, error) {
	var returnValue *Schema
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"schemaKey":  c.schemaKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve schema",
	//   "httpMethod": "GET",
	//   "id": "directory.schemas.get",
	//   "parameterOrder": [
	//     "customerId",
	//     "schemaKey"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "schemaKey": {
	//       "description": "Name or immutable Id of the schema",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/schemas/{schemaKey}",
	//   "response": {
	//     "$ref": "Schema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.userschema",
	//     "https://www.googleapis.com/auth/admin.directory.userschema.readonly"
	//   ]
	// }

}

// method id "directory.schemas.insert":

type SchemasInsertCall struct {
	s             *Service
	customerId    string
	schema        *Schema
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Create schema.

func (r *SchemasService) Insert(customerId string, schema *Schema) *SchemasInsertCall {
	return &SchemasInsertCall{
		s:             r.s,
		customerId:    customerId,
		schema:        schema,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/schemas",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SchemasInsertCall) Fields(s ...googleapi.Field) *SchemasInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SchemasInsertCall) Do() (*Schema, error) {
	var returnValue *Schema
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.schema,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Create schema.",
	//   "httpMethod": "POST",
	//   "id": "directory.schemas.insert",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/schemas",
	//   "request": {
	//     "$ref": "Schema"
	//   },
	//   "response": {
	//     "$ref": "Schema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.userschema"
	//   ]
	// }

}

// method id "directory.schemas.list":

type SchemasListCall struct {
	s             *Service
	customerId    string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve all schemas for a customer

func (r *SchemasService) List(customerId string) *SchemasListCall {
	return &SchemasListCall{
		s:             r.s,
		customerId:    customerId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/schemas",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SchemasListCall) Fields(s ...googleapi.Field) *SchemasListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SchemasListCall) Do() (*Schemas, error) {
	var returnValue *Schemas
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve all schemas for a customer",
	//   "httpMethod": "GET",
	//   "id": "directory.schemas.list",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/schemas",
	//   "response": {
	//     "$ref": "Schemas"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.userschema",
	//     "https://www.googleapis.com/auth/admin.directory.userschema.readonly"
	//   ]
	// }

}

// method id "directory.schemas.patch":

type SchemasPatchCall struct {
	s             *Service
	customerId    string
	schemaKey     string
	schema        *Schema
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Update schema. This method supports patch semantics.

func (r *SchemasService) Patch(customerId string, schemaKey string, schema *Schema) *SchemasPatchCall {
	return &SchemasPatchCall{
		s:             r.s,
		customerId:    customerId,
		schemaKey:     schemaKey,
		schema:        schema,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/schemas/{schemaKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SchemasPatchCall) Fields(s ...googleapi.Field) *SchemasPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SchemasPatchCall) Do() (*Schema, error) {
	var returnValue *Schema
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"schemaKey":  c.schemaKey,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.schema,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update schema. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.schemas.patch",
	//   "parameterOrder": [
	//     "customerId",
	//     "schemaKey"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "schemaKey": {
	//       "description": "Name or immutable Id of the schema.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/schemas/{schemaKey}",
	//   "request": {
	//     "$ref": "Schema"
	//   },
	//   "response": {
	//     "$ref": "Schema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.userschema"
	//   ]
	// }

}

// method id "directory.schemas.update":

type SchemasUpdateCall struct {
	s             *Service
	customerId    string
	schemaKey     string
	schema        *Schema
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Update schema

func (r *SchemasService) Update(customerId string, schemaKey string, schema *Schema) *SchemasUpdateCall {
	return &SchemasUpdateCall{
		s:             r.s,
		customerId:    customerId,
		schemaKey:     schemaKey,
		schema:        schema,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "customer/{customerId}/schemas/{schemaKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SchemasUpdateCall) Fields(s ...googleapi.Field) *SchemasUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SchemasUpdateCall) Do() (*Schema, error) {
	var returnValue *Schema
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"customerId": c.customerId,
		"schemaKey":  c.schemaKey,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.schema,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Update schema",
	//   "httpMethod": "PUT",
	//   "id": "directory.schemas.update",
	//   "parameterOrder": [
	//     "customerId",
	//     "schemaKey"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Immutable id of the Google Apps account",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "schemaKey": {
	//       "description": "Name or immutable Id of the schema.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customer/{customerId}/schemas/{schemaKey}",
	//   "request": {
	//     "$ref": "Schema"
	//   },
	//   "response": {
	//     "$ref": "Schema"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.userschema"
	//   ]
	// }

}

// method id "directory.tokens.delete":

type TokensDeleteCall struct {
	s             *Service
	userKey       string
	clientId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete all access tokens issued by a user for an application.

func (r *TokensService) Delete(userKey string, clientId string) *TokensDeleteCall {
	return &TokensDeleteCall{
		s:             r.s,
		userKey:       userKey,
		clientId:      clientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/tokens/{clientId}",
	}
}

func (c *TokensDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey":  c.userKey,
		"clientId": c.clientId,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete all access tokens issued by a user for an application.",
	//   "httpMethod": "DELETE",
	//   "id": "directory.tokens.delete",
	//   "parameterOrder": [
	//     "userKey",
	//     "clientId"
	//   ],
	//   "parameters": {
	//     "clientId": {
	//       "description": "The Client ID of the application the token is issued to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/tokens/{clientId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.tokens.get":

type TokensGetCall struct {
	s             *Service
	userKey       string
	clientId      string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Get information about an access token issued by a user.

func (r *TokensService) Get(userKey string, clientId string) *TokensGetCall {
	return &TokensGetCall{
		s:             r.s,
		userKey:       userKey,
		clientId:      clientId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/tokens/{clientId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TokensGetCall) Fields(s ...googleapi.Field) *TokensGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TokensGetCall) Do() (*Token, error) {
	var returnValue *Token
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey":  c.userKey,
		"clientId": c.clientId,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Get information about an access token issued by a user.",
	//   "httpMethod": "GET",
	//   "id": "directory.tokens.get",
	//   "parameterOrder": [
	//     "userKey",
	//     "clientId"
	//   ],
	//   "parameters": {
	//     "clientId": {
	//       "description": "The Client ID of the application the token is issued to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/tokens/{clientId}",
	//   "response": {
	//     "$ref": "Token"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.tokens.list":

type TokensListCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns the set of tokens specified user has issued to 3rd
// party applications.

func (r *TokensService) List(userKey string) *TokensListCall {
	return &TokensListCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/tokens",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TokensListCall) Fields(s ...googleapi.Field) *TokensListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *TokensListCall) Do() (*Tokens, error) {
	var returnValue *Tokens
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the set of tokens specified user has issued to 3rd party applications.",
	//   "httpMethod": "GET",
	//   "id": "directory.tokens.list",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/tokens",
	//   "response": {
	//     "$ref": "Tokens"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.users.delete":

type UsersDeleteCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Delete user

func (r *UsersService) Delete(userKey string) *UsersDeleteCall {
	return &UsersDeleteCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}",
	}
}

func (c *UsersDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Delete user",
	//   "httpMethod": "DELETE",
	//   "id": "directory.users.delete",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.get":

type UsersGetCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: retrieve user

func (r *UsersService) Get(userKey string) *UsersGetCall {
	return &UsersGetCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}",
	}
}

// CustomFieldMask sets the optional parameter "customFieldMask":
// Comma-separated list of schema names. All fields from these schemas
// are fetched. This should only be set when projection=custom.
func (c *UsersGetCall) CustomFieldMask(customFieldMask string) *UsersGetCall {
	c.params_.Set("customFieldMask", fmt.Sprintf("%v", customFieldMask))
	return c
}

// Projection sets the optional parameter "projection": What subset of
// fields to fetch for this user.
func (c *UsersGetCall) Projection(projection string) *UsersGetCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// ViewType sets the optional parameter "viewType": Whether to fetch the
// ADMIN_VIEW or DOMAIN_PUBLIC view of the user.
func (c *UsersGetCall) ViewType(viewType string) *UsersGetCall {
	c.params_.Set("viewType", fmt.Sprintf("%v", viewType))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersGetCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "retrieve user",
	//   "httpMethod": "GET",
	//   "id": "directory.users.get",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "customFieldMask": {
	//       "description": "Comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when projection=custom.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "default": "basic",
	//       "description": "What subset of fields to fetch for this user.",
	//       "enum": [
	//         "basic",
	//         "custom",
	//         "full"
	//       ],
	//       "enumDescriptions": [
	//         "Do not include any custom fields for the user.",
	//         "Include custom fields from schemas mentioned in customFieldMask.",
	//         "Include all fields associated with this user."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "viewType": {
	//       "default": "admin_view",
	//       "description": "Whether to fetch the ADMIN_VIEW or DOMAIN_PUBLIC view of the user.",
	//       "enum": [
	//         "admin_view",
	//         "domain_public"
	//       ],
	//       "enumDescriptions": [
	//         "Fetches the ADMIN_VIEW of the user.",
	//         "Fetches the DOMAIN_PUBLIC view of the user."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}",
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ]
	// }

}

// method id "directory.users.insert":

type UsersInsertCall struct {
	s             *Service
	user          *User
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: create user.

func (r *UsersService) Insert(user *User) *UsersInsertCall {
	return &UsersInsertCall{
		s:             r.s,
		user:          user,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersInsertCall) Fields(s ...googleapi.Field) *UsersInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersInsertCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.user,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "create user.",
	//   "httpMethod": "POST",
	//   "id": "directory.users.insert",
	//   "path": "users",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.list":

type UsersListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieve either deleted users or all users in a domain
// (paginated)

func (r *UsersService) List() *UsersListCall {
	return &UsersListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users",
	}
}

// CustomFieldMask sets the optional parameter "customFieldMask":
// Comma-separated list of schema names. All fields from these schemas
// are fetched. This should only be set when projection=custom.
func (c *UsersListCall) CustomFieldMask(customFieldMask string) *UsersListCall {
	c.params_.Set("customFieldMask", fmt.Sprintf("%v", customFieldMask))
	return c
}

// Customer sets the optional parameter "customer": Immutable id of the
// Google Apps account. In case of multi-domain, to fetch all users for
// a customer, fill this field instead of domain.
func (c *UsersListCall) Customer(customer string) *UsersListCall {
	c.params_.Set("customer", fmt.Sprintf("%v", customer))
	return c
}

// Domain sets the optional parameter "domain": Name of the domain. Fill
// this field to get users from only this domain. To return all users in
// a multi-domain fill customer field instead.
func (c *UsersListCall) Domain(domain string) *UsersListCall {
	c.params_.Set("domain", fmt.Sprintf("%v", domain))
	return c
}

// Event sets the optional parameter "event": Event on which
// subscription is intended (if subscribing)
func (c *UsersListCall) Event(event string) *UsersListCall {
	c.params_.Set("event", fmt.Sprintf("%v", event))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 100. Max allowed is 500
func (c *UsersListCall) MaxResults(maxResults int64) *UsersListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Column to use for
// sorting results
func (c *UsersListCall) OrderBy(orderBy string) *UsersListCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *UsersListCall) PageToken(pageToken string) *UsersListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Projection sets the optional parameter "projection": What subset of
// fields to fetch for this user.
func (c *UsersListCall) Projection(projection string) *UsersListCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Query sets the optional parameter "query": Query string search.
// Should be of the form "". Complete documentation is at
// https://developers.google.com/admin-sdk/directory/v1/guides/search-use
// rs
func (c *UsersListCall) Query(query string) *UsersListCall {
	c.params_.Set("query", fmt.Sprintf("%v", query))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": If set to true
// retrieves the list of deleted users. Default is false
func (c *UsersListCall) ShowDeleted(showDeleted string) *UsersListCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Whether to return
// results in ascending or descending order.
func (c *UsersListCall) SortOrder(sortOrder string) *UsersListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// ViewType sets the optional parameter "viewType": Whether to fetch the
// ADMIN_VIEW or DOMAIN_PUBLIC view of the user.
func (c *UsersListCall) ViewType(viewType string) *UsersListCall {
	c.params_.Set("viewType", fmt.Sprintf("%v", viewType))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListCall) Fields(s ...googleapi.Field) *UsersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersListCall) Do() (*Users, error) {
	var returnValue *Users
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve either deleted users or all users in a domain (paginated)",
	//   "httpMethod": "GET",
	//   "id": "directory.users.list",
	//   "parameters": {
	//     "customFieldMask": {
	//       "description": "Comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when projection=custom.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customer": {
	//       "description": "Immutable id of the Google Apps account. In case of multi-domain, to fetch all users for a customer, fill this field instead of domain.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "domain": {
	//       "description": "Name of the domain. Fill this field to get users from only this domain. To return all users in a multi-domain fill customer field instead.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "event": {
	//       "description": "Event on which subscription is intended (if subscribing)",
	//       "enum": [
	//         "add",
	//         "delete",
	//         "makeAdmin",
	//         "undelete",
	//         "update"
	//       ],
	//       "enumDescriptions": [
	//         "User Created Event",
	//         "User Deleted Event",
	//         "User Admin Status Change Event",
	//         "User Undeleted Event",
	//         "User Updated Event"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 100. Max allowed is 500",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Column to use for sorting results",
	//       "enum": [
	//         "email",
	//         "familyName",
	//         "givenName"
	//       ],
	//       "enumDescriptions": [
	//         "Primary email of the user.",
	//         "User's family name.",
	//         "User's given name."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "default": "basic",
	//       "description": "What subset of fields to fetch for this user.",
	//       "enum": [
	//         "basic",
	//         "custom",
	//         "full"
	//       ],
	//       "enumDescriptions": [
	//         "Do not include any custom fields for the user.",
	//         "Include custom fields from schemas mentioned in customFieldMask.",
	//         "Include all fields associated with this user."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Query string search. Should be of the form \"\". Complete documentation is at https://developers.google.com/admin-sdk/directory/v1/guides/search-users",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "If set to true retrieves the list of deleted users. Default is false",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Whether to return results in ascending or descending order.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "Ascending order.",
	//         "Descending order."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "viewType": {
	//       "default": "admin_view",
	//       "description": "Whether to fetch the ADMIN_VIEW or DOMAIN_PUBLIC view of the user.",
	//       "enum": [
	//         "admin_view",
	//         "domain_public"
	//       ],
	//       "enumDescriptions": [
	//         "Fetches the ADMIN_VIEW of the user.",
	//         "Fetches the DOMAIN_PUBLIC view of the user."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users",
	//   "response": {
	//     "$ref": "Users"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "directory.users.makeAdmin":

type UsersMakeAdminCall struct {
	s             *Service
	userKey       string
	usermakeadmin *UserMakeAdmin
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// MakeAdmin: change admin status of a user

func (r *UsersService) MakeAdmin(userKey string, usermakeadmin *UserMakeAdmin) *UsersMakeAdminCall {
	return &UsersMakeAdminCall{
		s:             r.s,
		userKey:       userKey,
		usermakeadmin: usermakeadmin,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/makeAdmin",
	}
}

func (c *UsersMakeAdminCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.usermakeadmin,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "change admin status of a user",
	//   "httpMethod": "POST",
	//   "id": "directory.users.makeAdmin",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user as admin",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/makeAdmin",
	//   "request": {
	//     "$ref": "UserMakeAdmin"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.patch":

type UsersPatchCall struct {
	s             *Service
	userKey       string
	user          *User
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: update user. This method supports patch semantics.

func (r *UsersService) Patch(userKey string, user *User) *UsersPatchCall {
	return &UsersPatchCall{
		s:             r.s,
		userKey:       userKey,
		user:          user,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersPatchCall) Fields(s ...googleapi.Field) *UsersPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersPatchCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.user,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "update user. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.users.patch",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user. If Id, it should match with id of user object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.undelete":

type UsersUndeleteCall struct {
	s             *Service
	userKey       string
	userundelete  *UserUndelete
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Undelete: Undelete a deleted user

func (r *UsersService) Undelete(userKey string, userundelete *UserUndelete) *UsersUndeleteCall {
	return &UsersUndeleteCall{
		s:             r.s,
		userKey:       userKey,
		userundelete:  userundelete,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/undelete",
	}
}

func (c *UsersUndeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.userundelete,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Undelete a deleted user",
	//   "httpMethod": "POST",
	//   "id": "directory.users.undelete",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "The immutable id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/undelete",
	//   "request": {
	//     "$ref": "UserUndelete"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.update":

type UsersUpdateCall struct {
	s             *Service
	userKey       string
	user          *User
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: update user

func (r *UsersService) Update(userKey string, user *User) *UsersUpdateCall {
	return &UsersUpdateCall{
		s:             r.s,
		userKey:       userKey,
		user:          user,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersUpdateCall) Fields(s ...googleapi.Field) *UsersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersUpdateCall) Do() (*User, error) {
	var returnValue *User
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.user,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "update user",
	//   "httpMethod": "PUT",
	//   "id": "directory.users.update",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user. If Id, it should match with id of user object",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "User"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.watch":

type UsersWatchCall struct {
	s             *Service
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Watch: Watch for changes in users list

func (r *UsersService) Watch(channel *Channel) *UsersWatchCall {
	return &UsersWatchCall{
		s:             r.s,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/watch",
	}
}

// CustomFieldMask sets the optional parameter "customFieldMask":
// Comma-separated list of schema names. All fields from these schemas
// are fetched. This should only be set when projection=custom.
func (c *UsersWatchCall) CustomFieldMask(customFieldMask string) *UsersWatchCall {
	c.params_.Set("customFieldMask", fmt.Sprintf("%v", customFieldMask))
	return c
}

// Customer sets the optional parameter "customer": Immutable id of the
// Google Apps account. In case of multi-domain, to fetch all users for
// a customer, fill this field instead of domain.
func (c *UsersWatchCall) Customer(customer string) *UsersWatchCall {
	c.params_.Set("customer", fmt.Sprintf("%v", customer))
	return c
}

// Domain sets the optional parameter "domain": Name of the domain. Fill
// this field to get users from only this domain. To return all users in
// a multi-domain fill customer field instead.
func (c *UsersWatchCall) Domain(domain string) *UsersWatchCall {
	c.params_.Set("domain", fmt.Sprintf("%v", domain))
	return c
}

// Event sets the optional parameter "event": Event on which
// subscription is intended (if subscribing)
func (c *UsersWatchCall) Event(event string) *UsersWatchCall {
	c.params_.Set("event", fmt.Sprintf("%v", event))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return. Default is 100. Max allowed is 500
func (c *UsersWatchCall) MaxResults(maxResults int64) *UsersWatchCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": Column to use for
// sorting results
func (c *UsersWatchCall) OrderBy(orderBy string) *UsersWatchCall {
	c.params_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *UsersWatchCall) PageToken(pageToken string) *UsersWatchCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Projection sets the optional parameter "projection": What subset of
// fields to fetch for this user.
func (c *UsersWatchCall) Projection(projection string) *UsersWatchCall {
	c.params_.Set("projection", fmt.Sprintf("%v", projection))
	return c
}

// Query sets the optional parameter "query": Query string search.
// Should be of the form "". Complete documentation is at
// https://developers.google.com/admin-sdk/directory/v1/guides/search-use
// rs
func (c *UsersWatchCall) Query(query string) *UsersWatchCall {
	c.params_.Set("query", fmt.Sprintf("%v", query))
	return c
}

// ShowDeleted sets the optional parameter "showDeleted": If set to true
// retrieves the list of deleted users. Default is false
func (c *UsersWatchCall) ShowDeleted(showDeleted string) *UsersWatchCall {
	c.params_.Set("showDeleted", fmt.Sprintf("%v", showDeleted))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Whether to return
// results in ascending or descending order.
func (c *UsersWatchCall) SortOrder(sortOrder string) *UsersWatchCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// ViewType sets the optional parameter "viewType": Whether to fetch the
// ADMIN_VIEW or DOMAIN_PUBLIC view of the user.
func (c *UsersWatchCall) ViewType(viewType string) *UsersWatchCall {
	c.params_.Set("viewType", fmt.Sprintf("%v", viewType))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersWatchCall) Fields(s ...googleapi.Field) *UsersWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Watch for changes in users list",
	//   "httpMethod": "POST",
	//   "id": "directory.users.watch",
	//   "parameters": {
	//     "customFieldMask": {
	//       "description": "Comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when projection=custom.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customer": {
	//       "description": "Immutable id of the Google Apps account. In case of multi-domain, to fetch all users for a customer, fill this field instead of domain.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "domain": {
	//       "description": "Name of the domain. Fill this field to get users from only this domain. To return all users in a multi-domain fill customer field instead.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "event": {
	//       "description": "Event on which subscription is intended (if subscribing)",
	//       "enum": [
	//         "add",
	//         "delete",
	//         "makeAdmin",
	//         "undelete",
	//         "update"
	//       ],
	//       "enumDescriptions": [
	//         "User Created Event",
	//         "User Deleted Event",
	//         "User Admin Status Change Event",
	//         "User Undeleted Event",
	//         "User Updated Event"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return. Default is 100. Max allowed is 500",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "Column to use for sorting results",
	//       "enum": [
	//         "email",
	//         "familyName",
	//         "givenName"
	//       ],
	//       "enumDescriptions": [
	//         "Primary email of the user.",
	//         "User's family name.",
	//         "User's given name."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projection": {
	//       "default": "basic",
	//       "description": "What subset of fields to fetch for this user.",
	//       "enum": [
	//         "basic",
	//         "custom",
	//         "full"
	//       ],
	//       "enumDescriptions": [
	//         "Do not include any custom fields for the user.",
	//         "Include custom fields from schemas mentioned in customFieldMask.",
	//         "Include all fields associated with this user."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Query string search. Should be of the form \"\". Complete documentation is at https://developers.google.com/admin-sdk/directory/v1/guides/search-users",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showDeleted": {
	//       "description": "If set to true retrieves the list of deleted users. Default is false",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Whether to return results in ascending or descending order.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "Ascending order.",
	//         "Descending order."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "viewType": {
	//       "default": "admin_view",
	//       "description": "Whether to fetch the ADMIN_VIEW or DOMAIN_PUBLIC view of the user.",
	//       "enum": [
	//         "admin_view",
	//         "domain_public"
	//       ],
	//       "enumDescriptions": [
	//         "Fetches the ADMIN_VIEW of the user.",
	//         "Fetches the DOMAIN_PUBLIC view of the user."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "directory.users.aliases.delete":

type UsersAliasesDeleteCall struct {
	s             *Service
	userKey       string
	alias         string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Remove a alias for the user

func (r *UsersAliasesService) Delete(userKey string, alias string) *UsersAliasesDeleteCall {
	return &UsersAliasesDeleteCall{
		s:             r.s,
		userKey:       userKey,
		alias:         alias,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/aliases/{alias}",
	}
}

func (c *UsersAliasesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
		"alias":   c.alias,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Remove a alias for the user",
	//   "httpMethod": "DELETE",
	//   "id": "directory.users.aliases.delete",
	//   "parameterOrder": [
	//     "userKey",
	//     "alias"
	//   ],
	//   "parameters": {
	//     "alias": {
	//       "description": "The alias to be removed",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/aliases/{alias}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias"
	//   ]
	// }

}

// method id "directory.users.aliases.insert":

type UsersAliasesInsertCall struct {
	s             *Service
	userKey       string
	alias         *Alias
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Add a alias for the user

func (r *UsersAliasesService) Insert(userKey string, alias *Alias) *UsersAliasesInsertCall {
	return &UsersAliasesInsertCall{
		s:             r.s,
		userKey:       userKey,
		alias:         alias,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/aliases",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersAliasesInsertCall) Fields(s ...googleapi.Field) *UsersAliasesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersAliasesInsertCall) Do() (*Alias, error) {
	var returnValue *Alias
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.alias,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add a alias for the user",
	//   "httpMethod": "POST",
	//   "id": "directory.users.aliases.insert",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/aliases",
	//   "request": {
	//     "$ref": "Alias"
	//   },
	//   "response": {
	//     "$ref": "Alias"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias"
	//   ]
	// }

}

// method id "directory.users.aliases.list":

type UsersAliasesListCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: List all aliases for a user

func (r *UsersAliasesService) List(userKey string) *UsersAliasesListCall {
	return &UsersAliasesListCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/aliases",
	}
}

// Event sets the optional parameter "event": Event on which
// subscription is intended (if subscribing)
func (c *UsersAliasesListCall) Event(event string) *UsersAliasesListCall {
	c.params_.Set("event", fmt.Sprintf("%v", event))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersAliasesListCall) Fields(s ...googleapi.Field) *UsersAliasesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersAliasesListCall) Do() (*Aliases, error) {
	var returnValue *Aliases
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "List all aliases for a user",
	//   "httpMethod": "GET",
	//   "id": "directory.users.aliases.list",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "event": {
	//       "description": "Event on which subscription is intended (if subscribing)",
	//       "enum": [
	//         "add",
	//         "delete"
	//       ],
	//       "enumDescriptions": [
	//         "Alias Created Event",
	//         "Alias Deleted Event"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/aliases",
	//   "response": {
	//     "$ref": "Aliases"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias.readonly",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "directory.users.aliases.watch":

type UsersAliasesWatchCall struct {
	s             *Service
	userKey       string
	channel       *Channel
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Watch: Watch for changes in user aliases list

func (r *UsersAliasesService) Watch(userKey string, channel *Channel) *UsersAliasesWatchCall {
	return &UsersAliasesWatchCall{
		s:             r.s,
		userKey:       userKey,
		channel:       channel,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/aliases/watch",
	}
}

// Event sets the optional parameter "event": Event on which
// subscription is intended (if subscribing)
func (c *UsersAliasesWatchCall) Event(event string) *UsersAliasesWatchCall {
	c.params_.Set("event", fmt.Sprintf("%v", event))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersAliasesWatchCall) Fields(s ...googleapi.Field) *UsersAliasesWatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersAliasesWatchCall) Do() (*Channel, error) {
	var returnValue *Channel
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.channel,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Watch for changes in user aliases list",
	//   "httpMethod": "POST",
	//   "id": "directory.users.aliases.watch",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "event": {
	//       "description": "Event on which subscription is intended (if subscribing)",
	//       "enum": [
	//         "add",
	//         "delete"
	//       ],
	//       "enumDescriptions": [
	//         "Alias Created Event",
	//         "Alias Deleted Event"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/aliases/watch",
	//   "request": {
	//     "$ref": "Channel",
	//     "parameterName": "resource"
	//   },
	//   "response": {
	//     "$ref": "Channel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias",
	//     "https://www.googleapis.com/auth/admin.directory.user.alias.readonly",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ],
	//   "supportsSubscription": true
	// }

}

// method id "directory.users.photos.delete":

type UsersPhotosDeleteCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Remove photos for the user

func (r *UsersPhotosService) Delete(userKey string) *UsersPhotosDeleteCall {
	return &UsersPhotosDeleteCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/photos/thumbnail",
	}
}

func (c *UsersPhotosDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Remove photos for the user",
	//   "httpMethod": "DELETE",
	//   "id": "directory.users.photos.delete",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/photos/thumbnail",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.photos.get":

type UsersPhotosGetCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieve photo of a user

func (r *UsersPhotosService) Get(userKey string) *UsersPhotosGetCall {
	return &UsersPhotosGetCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/photos/thumbnail",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersPhotosGetCall) Fields(s ...googleapi.Field) *UsersPhotosGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersPhotosGetCall) Do() (*UserPhoto, error) {
	var returnValue *UserPhoto
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieve photo of a user",
	//   "httpMethod": "GET",
	//   "id": "directory.users.photos.get",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/photos/thumbnail",
	//   "response": {
	//     "$ref": "UserPhoto"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user",
	//     "https://www.googleapis.com/auth/admin.directory.user.readonly"
	//   ]
	// }

}

// method id "directory.users.photos.patch":

type UsersPhotosPatchCall struct {
	s             *Service
	userKey       string
	userphoto     *UserPhoto
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Add a photo for the user. This method supports patch
// semantics.

func (r *UsersPhotosService) Patch(userKey string, userphoto *UserPhoto) *UsersPhotosPatchCall {
	return &UsersPhotosPatchCall{
		s:             r.s,
		userKey:       userKey,
		userphoto:     userphoto,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/photos/thumbnail",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersPhotosPatchCall) Fields(s ...googleapi.Field) *UsersPhotosPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersPhotosPatchCall) Do() (*UserPhoto, error) {
	var returnValue *UserPhoto
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.userphoto,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add a photo for the user. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "directory.users.photos.patch",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/photos/thumbnail",
	//   "request": {
	//     "$ref": "UserPhoto"
	//   },
	//   "response": {
	//     "$ref": "UserPhoto"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.users.photos.update":

type UsersPhotosUpdateCall struct {
	s             *Service
	userKey       string
	userphoto     *UserPhoto
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Add a photo for the user

func (r *UsersPhotosService) Update(userKey string, userphoto *UserPhoto) *UsersPhotosUpdateCall {
	return &UsersPhotosUpdateCall{
		s:             r.s,
		userKey:       userKey,
		userphoto:     userphoto,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/photos/thumbnail",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersPhotosUpdateCall) Fields(s ...googleapi.Field) *UsersPhotosUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UsersPhotosUpdateCall) Do() (*UserPhoto, error) {
	var returnValue *UserPhoto
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.userphoto,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Add a photo for the user",
	//   "httpMethod": "PUT",
	//   "id": "directory.users.photos.update",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/photos/thumbnail",
	//   "request": {
	//     "$ref": "UserPhoto"
	//   },
	//   "response": {
	//     "$ref": "UserPhoto"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user"
	//   ]
	// }

}

// method id "directory.verificationCodes.generate":

type VerificationCodesGenerateCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generate: Generate new backup verification codes for the user.

func (r *VerificationCodesService) Generate(userKey string) *VerificationCodesGenerateCall {
	return &VerificationCodesGenerateCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/verificationCodes/generate",
	}
}

func (c *VerificationCodesGenerateCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generate new backup verification codes for the user.",
	//   "httpMethod": "POST",
	//   "id": "directory.verificationCodes.generate",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/verificationCodes/generate",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.verificationCodes.invalidate":

type VerificationCodesInvalidateCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Invalidate: Invalidate the current backup verification codes for the
// user.

func (r *VerificationCodesService) Invalidate(userKey string) *VerificationCodesInvalidateCall {
	return &VerificationCodesInvalidateCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/verificationCodes/invalidate",
	}
}

func (c *VerificationCodesInvalidateCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Invalidate the current backup verification codes for the user.",
	//   "httpMethod": "POST",
	//   "id": "directory.verificationCodes.invalidate",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Email or immutable Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/verificationCodes/invalidate",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}

// method id "directory.verificationCodes.list":

type VerificationCodesListCall struct {
	s             *Service
	userKey       string
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Returns the current set of valid backup verification codes for
// the specified user.

func (r *VerificationCodesService) List(userKey string) *VerificationCodesListCall {
	return &VerificationCodesListCall{
		s:             r.s,
		userKey:       userKey,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "users/{userKey}/verificationCodes",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VerificationCodesListCall) Fields(s ...googleapi.Field) *VerificationCodesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *VerificationCodesListCall) Do() (*VerificationCodes, error) {
	var returnValue *VerificationCodes
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"userKey": c.userKey,
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the current set of valid backup verification codes for the specified user.",
	//   "httpMethod": "GET",
	//   "id": "directory.verificationCodes.list",
	//   "parameterOrder": [
	//     "userKey"
	//   ],
	//   "parameters": {
	//     "userKey": {
	//       "description": "Identifies the user in the API request. The value can be the user's primary email address, alias email address, or unique user ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "users/{userKey}/verificationCodes",
	//   "response": {
	//     "$ref": "VerificationCodes"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/admin.directory.user.security"
	//   ]
	// }

}
