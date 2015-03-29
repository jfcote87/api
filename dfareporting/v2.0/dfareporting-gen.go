// Package dfareporting provides access to the DCM/DFA Reporting And Trafficking API.
//
// See https://developers.google.com/doubleclick-advertisers/reporting/
//
// Usage example:
//
//   import "github.com/jfcote87/api/dfareporting/v2.0"
//   ...
//   dfareportingService, err := dfareporting.New(oauthHttpClient)
package dfareporting

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

const apiId = "dfareporting:v2.0"
const apiName = "dfareporting"
const apiVersion = "v2.0"

var baseURL *url.URL = &url.URL{Scheme: "https", Host: "www.googleapis.com", Path: "/dfareporting/v2.0/"}

// OAuth2 scopes used by this API.
const (
	// View and manage DoubleClick for Advertisers reports
	DfareportingScope = "https://www.googleapis.com/auth/dfareporting"

	// View and manage your DoubleClick Campaign Manager's (DCM) display ad
	// campaigns
	DfatraffickingScope = "https://www.googleapis.com/auth/dfatrafficking"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client}
	s.AccountActiveAdSummaries = NewAccountActiveAdSummariesService(s)
	s.AccountPermissionGroups = NewAccountPermissionGroupsService(s)
	s.AccountPermissions = NewAccountPermissionsService(s)
	s.AccountUserProfiles = NewAccountUserProfilesService(s)
	s.Accounts = NewAccountsService(s)
	s.Ads = NewAdsService(s)
	s.AdvertiserGroups = NewAdvertiserGroupsService(s)
	s.Advertisers = NewAdvertisersService(s)
	s.Browsers = NewBrowsersService(s)
	s.CampaignCreativeAssociations = NewCampaignCreativeAssociationsService(s)
	s.Campaigns = NewCampaignsService(s)
	s.ChangeLogs = NewChangeLogsService(s)
	s.Cities = NewCitiesService(s)
	s.ConnectionTypes = NewConnectionTypesService(s)
	s.ContentCategories = NewContentCategoriesService(s)
	s.Countries = NewCountriesService(s)
	s.CreativeAssets = NewCreativeAssetsService(s)
	s.CreativeFieldValues = NewCreativeFieldValuesService(s)
	s.CreativeFields = NewCreativeFieldsService(s)
	s.CreativeGroups = NewCreativeGroupsService(s)
	s.Creatives = NewCreativesService(s)
	s.DimensionValues = NewDimensionValuesService(s)
	s.DirectorySiteContacts = NewDirectorySiteContactsService(s)
	s.DirectorySites = NewDirectorySitesService(s)
	s.EventTags = NewEventTagsService(s)
	s.Files = NewFilesService(s)
	s.FloodlightActivities = NewFloodlightActivitiesService(s)
	s.FloodlightActivityGroups = NewFloodlightActivityGroupsService(s)
	s.FloodlightConfigurations = NewFloodlightConfigurationsService(s)
	s.LandingPages = NewLandingPagesService(s)
	s.Metros = NewMetrosService(s)
	s.MobileCarriers = NewMobileCarriersService(s)
	s.OperatingSystemVersions = NewOperatingSystemVersionsService(s)
	s.OperatingSystems = NewOperatingSystemsService(s)
	s.PlacementGroups = NewPlacementGroupsService(s)
	s.PlacementStrategies = NewPlacementStrategiesService(s)
	s.Placements = NewPlacementsService(s)
	s.PlatformTypes = NewPlatformTypesService(s)
	s.PostalCodes = NewPostalCodesService(s)
	s.Regions = NewRegionsService(s)
	s.Reports = NewReportsService(s)
	s.Sites = NewSitesService(s)
	s.Sizes = NewSizesService(s)
	s.Subaccounts = NewSubaccountsService(s)
	s.UserProfiles = NewUserProfilesService(s)
	s.UserRolePermissionGroups = NewUserRolePermissionGroupsService(s)
	s.UserRolePermissions = NewUserRolePermissionsService(s)
	s.UserRoles = NewUserRolesService(s)
	return s, nil
}

type Service struct {
	client *http.Client

	AccountActiveAdSummaries *AccountActiveAdSummariesService

	AccountPermissionGroups *AccountPermissionGroupsService

	AccountPermissions *AccountPermissionsService

	AccountUserProfiles *AccountUserProfilesService

	Accounts *AccountsService

	Ads *AdsService

	AdvertiserGroups *AdvertiserGroupsService

	Advertisers *AdvertisersService

	Browsers *BrowsersService

	CampaignCreativeAssociations *CampaignCreativeAssociationsService

	Campaigns *CampaignsService

	ChangeLogs *ChangeLogsService

	Cities *CitiesService

	ConnectionTypes *ConnectionTypesService

	ContentCategories *ContentCategoriesService

	Countries *CountriesService

	CreativeAssets *CreativeAssetsService

	CreativeFieldValues *CreativeFieldValuesService

	CreativeFields *CreativeFieldsService

	CreativeGroups *CreativeGroupsService

	Creatives *CreativesService

	DimensionValues *DimensionValuesService

	DirectorySiteContacts *DirectorySiteContactsService

	DirectorySites *DirectorySitesService

	EventTags *EventTagsService

	Files *FilesService

	FloodlightActivities *FloodlightActivitiesService

	FloodlightActivityGroups *FloodlightActivityGroupsService

	FloodlightConfigurations *FloodlightConfigurationsService

	LandingPages *LandingPagesService

	Metros *MetrosService

	MobileCarriers *MobileCarriersService

	OperatingSystemVersions *OperatingSystemVersionsService

	OperatingSystems *OperatingSystemsService

	PlacementGroups *PlacementGroupsService

	PlacementStrategies *PlacementStrategiesService

	Placements *PlacementsService

	PlatformTypes *PlatformTypesService

	PostalCodes *PostalCodesService

	Regions *RegionsService

	Reports *ReportsService

	Sites *SitesService

	Sizes *SizesService

	Subaccounts *SubaccountsService

	UserProfiles *UserProfilesService

	UserRolePermissionGroups *UserRolePermissionGroupsService

	UserRolePermissions *UserRolePermissionsService

	UserRoles *UserRolesService
}

func NewAccountActiveAdSummariesService(s *Service) *AccountActiveAdSummariesService {
	rs := &AccountActiveAdSummariesService{s: s}
	return rs
}

type AccountActiveAdSummariesService struct {
	s *Service
}

func NewAccountPermissionGroupsService(s *Service) *AccountPermissionGroupsService {
	rs := &AccountPermissionGroupsService{s: s}
	return rs
}

type AccountPermissionGroupsService struct {
	s *Service
}

func NewAccountPermissionsService(s *Service) *AccountPermissionsService {
	rs := &AccountPermissionsService{s: s}
	return rs
}

type AccountPermissionsService struct {
	s *Service
}

func NewAccountUserProfilesService(s *Service) *AccountUserProfilesService {
	rs := &AccountUserProfilesService{s: s}
	return rs
}

type AccountUserProfilesService struct {
	s *Service
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	return rs
}

type AccountsService struct {
	s *Service
}

func NewAdsService(s *Service) *AdsService {
	rs := &AdsService{s: s}
	return rs
}

type AdsService struct {
	s *Service
}

func NewAdvertiserGroupsService(s *Service) *AdvertiserGroupsService {
	rs := &AdvertiserGroupsService{s: s}
	return rs
}

type AdvertiserGroupsService struct {
	s *Service
}

func NewAdvertisersService(s *Service) *AdvertisersService {
	rs := &AdvertisersService{s: s}
	return rs
}

type AdvertisersService struct {
	s *Service
}

func NewBrowsersService(s *Service) *BrowsersService {
	rs := &BrowsersService{s: s}
	return rs
}

type BrowsersService struct {
	s *Service
}

func NewCampaignCreativeAssociationsService(s *Service) *CampaignCreativeAssociationsService {
	rs := &CampaignCreativeAssociationsService{s: s}
	return rs
}

type CampaignCreativeAssociationsService struct {
	s *Service
}

func NewCampaignsService(s *Service) *CampaignsService {
	rs := &CampaignsService{s: s}
	return rs
}

type CampaignsService struct {
	s *Service
}

func NewChangeLogsService(s *Service) *ChangeLogsService {
	rs := &ChangeLogsService{s: s}
	return rs
}

type ChangeLogsService struct {
	s *Service
}

func NewCitiesService(s *Service) *CitiesService {
	rs := &CitiesService{s: s}
	return rs
}

type CitiesService struct {
	s *Service
}

func NewConnectionTypesService(s *Service) *ConnectionTypesService {
	rs := &ConnectionTypesService{s: s}
	return rs
}

type ConnectionTypesService struct {
	s *Service
}

func NewContentCategoriesService(s *Service) *ContentCategoriesService {
	rs := &ContentCategoriesService{s: s}
	return rs
}

type ContentCategoriesService struct {
	s *Service
}

func NewCountriesService(s *Service) *CountriesService {
	rs := &CountriesService{s: s}
	return rs
}

type CountriesService struct {
	s *Service
}

func NewCreativeAssetsService(s *Service) *CreativeAssetsService {
	rs := &CreativeAssetsService{s: s}
	return rs
}

type CreativeAssetsService struct {
	s *Service
}

func NewCreativeFieldValuesService(s *Service) *CreativeFieldValuesService {
	rs := &CreativeFieldValuesService{s: s}
	return rs
}

type CreativeFieldValuesService struct {
	s *Service
}

func NewCreativeFieldsService(s *Service) *CreativeFieldsService {
	rs := &CreativeFieldsService{s: s}
	return rs
}

type CreativeFieldsService struct {
	s *Service
}

func NewCreativeGroupsService(s *Service) *CreativeGroupsService {
	rs := &CreativeGroupsService{s: s}
	return rs
}

type CreativeGroupsService struct {
	s *Service
}

func NewCreativesService(s *Service) *CreativesService {
	rs := &CreativesService{s: s}
	return rs
}

type CreativesService struct {
	s *Service
}

func NewDimensionValuesService(s *Service) *DimensionValuesService {
	rs := &DimensionValuesService{s: s}
	return rs
}

type DimensionValuesService struct {
	s *Service
}

func NewDirectorySiteContactsService(s *Service) *DirectorySiteContactsService {
	rs := &DirectorySiteContactsService{s: s}
	return rs
}

type DirectorySiteContactsService struct {
	s *Service
}

func NewDirectorySitesService(s *Service) *DirectorySitesService {
	rs := &DirectorySitesService{s: s}
	return rs
}

type DirectorySitesService struct {
	s *Service
}

func NewEventTagsService(s *Service) *EventTagsService {
	rs := &EventTagsService{s: s}
	return rs
}

type EventTagsService struct {
	s *Service
}

func NewFilesService(s *Service) *FilesService {
	rs := &FilesService{s: s}
	return rs
}

type FilesService struct {
	s *Service
}

func NewFloodlightActivitiesService(s *Service) *FloodlightActivitiesService {
	rs := &FloodlightActivitiesService{s: s}
	return rs
}

type FloodlightActivitiesService struct {
	s *Service
}

func NewFloodlightActivityGroupsService(s *Service) *FloodlightActivityGroupsService {
	rs := &FloodlightActivityGroupsService{s: s}
	return rs
}

type FloodlightActivityGroupsService struct {
	s *Service
}

func NewFloodlightConfigurationsService(s *Service) *FloodlightConfigurationsService {
	rs := &FloodlightConfigurationsService{s: s}
	return rs
}

type FloodlightConfigurationsService struct {
	s *Service
}

func NewLandingPagesService(s *Service) *LandingPagesService {
	rs := &LandingPagesService{s: s}
	return rs
}

type LandingPagesService struct {
	s *Service
}

func NewMetrosService(s *Service) *MetrosService {
	rs := &MetrosService{s: s}
	return rs
}

type MetrosService struct {
	s *Service
}

func NewMobileCarriersService(s *Service) *MobileCarriersService {
	rs := &MobileCarriersService{s: s}
	return rs
}

type MobileCarriersService struct {
	s *Service
}

func NewOperatingSystemVersionsService(s *Service) *OperatingSystemVersionsService {
	rs := &OperatingSystemVersionsService{s: s}
	return rs
}

type OperatingSystemVersionsService struct {
	s *Service
}

func NewOperatingSystemsService(s *Service) *OperatingSystemsService {
	rs := &OperatingSystemsService{s: s}
	return rs
}

type OperatingSystemsService struct {
	s *Service
}

func NewPlacementGroupsService(s *Service) *PlacementGroupsService {
	rs := &PlacementGroupsService{s: s}
	return rs
}

type PlacementGroupsService struct {
	s *Service
}

func NewPlacementStrategiesService(s *Service) *PlacementStrategiesService {
	rs := &PlacementStrategiesService{s: s}
	return rs
}

type PlacementStrategiesService struct {
	s *Service
}

func NewPlacementsService(s *Service) *PlacementsService {
	rs := &PlacementsService{s: s}
	return rs
}

type PlacementsService struct {
	s *Service
}

func NewPlatformTypesService(s *Service) *PlatformTypesService {
	rs := &PlatformTypesService{s: s}
	return rs
}

type PlatformTypesService struct {
	s *Service
}

func NewPostalCodesService(s *Service) *PostalCodesService {
	rs := &PostalCodesService{s: s}
	return rs
}

type PostalCodesService struct {
	s *Service
}

func NewRegionsService(s *Service) *RegionsService {
	rs := &RegionsService{s: s}
	return rs
}

type RegionsService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	rs.CompatibleFields = NewReportsCompatibleFieldsService(s)
	rs.Files = NewReportsFilesService(s)
	return rs
}

type ReportsService struct {
	s *Service

	CompatibleFields *ReportsCompatibleFieldsService

	Files *ReportsFilesService
}

func NewReportsCompatibleFieldsService(s *Service) *ReportsCompatibleFieldsService {
	rs := &ReportsCompatibleFieldsService{s: s}
	return rs
}

type ReportsCompatibleFieldsService struct {
	s *Service
}

func NewReportsFilesService(s *Service) *ReportsFilesService {
	rs := &ReportsFilesService{s: s}
	return rs
}

type ReportsFilesService struct {
	s *Service
}

func NewSitesService(s *Service) *SitesService {
	rs := &SitesService{s: s}
	return rs
}

type SitesService struct {
	s *Service
}

func NewSizesService(s *Service) *SizesService {
	rs := &SizesService{s: s}
	return rs
}

type SizesService struct {
	s *Service
}

func NewSubaccountsService(s *Service) *SubaccountsService {
	rs := &SubaccountsService{s: s}
	return rs
}

type SubaccountsService struct {
	s *Service
}

func NewUserProfilesService(s *Service) *UserProfilesService {
	rs := &UserProfilesService{s: s}
	return rs
}

type UserProfilesService struct {
	s *Service
}

func NewUserRolePermissionGroupsService(s *Service) *UserRolePermissionGroupsService {
	rs := &UserRolePermissionGroupsService{s: s}
	return rs
}

type UserRolePermissionGroupsService struct {
	s *Service
}

func NewUserRolePermissionsService(s *Service) *UserRolePermissionsService {
	rs := &UserRolePermissionsService{s: s}
	return rs
}

type UserRolePermissionsService struct {
	s *Service
}

func NewUserRolesService(s *Service) *UserRolesService {
	rs := &UserRolesService{s: s}
	return rs
}

type UserRolesService struct {
	s *Service
}

type Account struct {
	// AccountPermissionIds: Account permissions assigned to this account.
	AccountPermissionIds googleapi.Int64s `json:"accountPermissionIds,omitempty"`

	// AccountProfile: Profile for this account. This is a read-only field
	// that can be left blank.
	AccountProfile string `json:"accountProfile,omitempty"`

	// Active: Whether this account is active.
	Active bool `json:"active,omitempty"`

	// ActiveAdsLimitTier: Maximum number of active ads allowed for this
	// account.
	ActiveAdsLimitTier string `json:"activeAdsLimitTier,omitempty"`

	// ActiveViewOptOut: Whether to serve creatives with Active View tags.
	// If disabled, viewability data will not be available for any
	// impressions.
	ActiveViewOptOut bool `json:"activeViewOptOut,omitempty"`

	// AvailablePermissionIds: User role permissions available to the user
	// roles of this account.
	AvailablePermissionIds googleapi.Int64s `json:"availablePermissionIds,omitempty"`

	// ComscoreVceEnabled: Whether campaigns created in this account will be
	// enabled for comScore vCE by default.
	ComscoreVceEnabled bool `json:"comscoreVceEnabled,omitempty"`

	// CountryId: ID of the country associated with this account.
	CountryId int64 `json:"countryId,omitempty,string"`

	// CurrencyId: ID of currency associated with this account. This is a
	// required field.
	// Acceptable values are:
	// - "1" for USD
	// - "2" for GBP
	//
	// - "3" for ESP
	// - "4" for SEK
	// - "5" for CAD
	// - "6" for JPY
	// - "7"
	// for DEM
	// - "8" for AUD
	// - "9" for FRF
	// - "10" for ITL
	// - "11" for DKK
	//
	// - "12" for NOK
	// - "13" for FIM
	// - "14" for ZAR
	// - "15" for IEP
	// -
	// "16" for NLG
	// - "17" for EUR
	// - "18" for KRW
	// - "19" for TWD
	// - "20"
	// for SGD
	// - "21" for CNY
	// - "22" for HKD
	// - "23" for NZD
	// - "24" for
	// MYR
	// - "25" for BRL
	// - "26" for PTE
	// - "27" for MXP
	// - "28" for CLP
	//
	// - "29" for TRY
	// - "30" for ARS
	// - "31" for PEN
	// - "32" for ILS
	// -
	// "33" for CHF
	// - "34" for VEF
	// - "35" for COP
	// - "36" for GTQ
	CurrencyId int64 `json:"currencyId,omitempty,string"`

	// DefaultCreativeSizeId: Default placement dimensions for this account.
	DefaultCreativeSizeId int64 `json:"defaultCreativeSizeId,omitempty,string"`

	// Description: Description of this account.
	Description string `json:"description,omitempty"`

	// Id: ID of this account. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#account".
	Kind string `json:"kind,omitempty"`

	// Locale: Locale of this account.
	// Acceptable values are:
	// - "cs"
	// (Czech)
	// - "de" (German)
	// - "en" (English)
	// - "en-GB" (English United
	// Kingdom)
	// - "es" (Spanish)
	// - "fr" (French)
	// - "it" (Italian)
	// - "ja"
	// (Japanese)
	// - "ko" (Korean)
	// - "pl" (Polish)
	// - "pt-BR" (Portuguese
	// Brazil)
	// - "ru" (Russian)
	// - "sv" (Swedish)
	// - "tr" (Turkish)
	// -
	// "zh-CN" (Chinese Simplified)
	// - "zh-TW" (Chinese Traditional)
	Locale string `json:"locale,omitempty"`

	// MaximumImageSize: Maximum image size allowed for this account.
	MaximumImageSize int64 `json:"maximumImageSize,omitempty,string"`

	// Name: Name of this account. This is a required field, and must be
	// less than 128 characters long and be globally unique.
	Name string `json:"name,omitempty"`

	// NielsenOcrEnabled: Whether campaigns created in this account will be
	// enabled for Nielsen OCR reach ratings by default.
	NielsenOcrEnabled bool `json:"nielsenOcrEnabled,omitempty"`

	// ReportsConfiguration: Reporting configuration of this account.
	ReportsConfiguration *ReportsConfiguration `json:"reportsConfiguration,omitempty"`

	// TeaserSizeLimit: File size limit in kilobytes of Rich Media teaser
	// creatives. Must be between 1 and 10240.
	TeaserSizeLimit int64 `json:"teaserSizeLimit,omitempty,string"`
}

type AccountActiveAdSummary struct {
	// AccountId: ID of the account.
	AccountId int64 `json:"accountId,omitempty,string"`

	// ActiveAds: Ads that have been activated for the account
	ActiveAds int64 `json:"activeAds,omitempty,string"`

	// ActiveAdsLimitTier: Maximum number of active ads allowed for the
	// account.
	ActiveAdsLimitTier string `json:"activeAdsLimitTier,omitempty"`

	// AvailableAds: Ads that can be activated for the account.
	AvailableAds int64 `json:"availableAds,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountActiveAdSummary".
	Kind string `json:"kind,omitempty"`
}

type AccountPermission struct {
	// AccountProfiles: Account profiles associated with this account
	// permission.
	//
	// Possible values are:
	// - "ACCOUNT_PROFILE_BASIC"
	// -
	// "ACCOUNT_PROFILE_STANDARD"
	AccountProfiles []string `json:"accountProfiles,omitempty"`

	// Id: ID of this account permission.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountPermission".
	Kind string `json:"kind,omitempty"`

	// Level: Administrative level required to enable this account
	// permission.
	Level string `json:"level,omitempty"`

	// Name: Name of this account permission.
	Name string `json:"name,omitempty"`

	// PermissionGroupId: Permission group of this account permission.
	PermissionGroupId int64 `json:"permissionGroupId,omitempty,string"`
}

type AccountPermissionGroup struct {
	// Id: ID of this account permission group.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountPermissionGroup".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this account permission group.
	Name string `json:"name,omitempty"`
}

type AccountPermissionGroupsListResponse struct {
	// AccountPermissionGroups: Account permission group collection
	AccountPermissionGroups []*AccountPermissionGroup `json:"accountPermissionGroups,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountPermissionGroupsListResponse".
	Kind string `json:"kind,omitempty"`
}

type AccountPermissionsListResponse struct {
	// AccountPermissions: Account permission collection
	AccountPermissions []*AccountPermission `json:"accountPermissions,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountPermissionsListResponse".
	Kind string `json:"kind,omitempty"`
}

type AccountUserProfile struct {
	// AccountId: Account ID of the user profile. This is a read-only field
	// that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Active: Whether this user profile is active. This defaults to false,
	// and must be set true on insert for the user profile to be usable.
	Active bool `json:"active,omitempty"`

	// AdvertiserFilter: Filter that describes which advertisers are visible
	// to the user profile.
	AdvertiserFilter *ObjectFilter `json:"advertiserFilter,omitempty"`

	// CampaignFilter: Filter that describes which campaigns are visible to
	// the user profile.
	CampaignFilter *ObjectFilter `json:"campaignFilter,omitempty"`

	// Comments: Comments for this user profile.
	Comments string `json:"comments,omitempty"`

	// Email: Email of the user profile. The email addresss must be linked
	// to a Google Account. This field is required on insertion and is
	// read-only after insertion.
	Email string `json:"email,omitempty"`

	// Id: ID of the user profile. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountUserProfile".
	Kind string `json:"kind,omitempty"`

	// Locale: Locale of the user profile. This is a required
	// field.
	// Acceptable values are:
	// - "cs" (Czech)
	// - "de" (German)
	// -
	// "en" (English)
	// - "en-GB" (English United Kingdom)
	// - "es" (Spanish)
	//
	// - "fr" (French)
	// - "it" (Italian)
	// - "ja" (Japanese)
	// - "ko"
	// (Korean)
	// - "pl" (Polish)
	// - "pt-BR" (Portuguese Brazil)
	// - "ru"
	// (Russian)
	// - "sv" (Swedish)
	// - "tr" (Turkish)
	// - "zh-CN" (Chinese
	// Simplified)
	// - "zh-TW" (Chinese Traditional)
	Locale string `json:"locale,omitempty"`

	// Name: Name of the user profile. This is a required field. Must be
	// less than 64 characters long, must be globally unique, and cannot
	// contain whitespace or any of the following characters: "&;"#%,".
	Name string `json:"name,omitempty"`

	// SiteFilter: Filter that describes which sites are visible to the user
	// profile.
	SiteFilter *ObjectFilter `json:"siteFilter,omitempty"`

	// SubaccountId: Subaccount ID of the user profile. This is a read-only
	// field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TraffickerType: Trafficker type of this user profile.
	TraffickerType string `json:"traffickerType,omitempty"`

	// UserAccessType: User type of the user profile. This is a read-only
	// field that can be left blank.
	UserAccessType string `json:"userAccessType,omitempty"`

	// UserRoleFilter: Filter that describes which user roles are visible to
	// the user profile.
	UserRoleFilter *ObjectFilter `json:"userRoleFilter,omitempty"`

	// UserRoleId: User role ID of the user profile. This is a required
	// field.
	UserRoleId int64 `json:"userRoleId,omitempty,string"`
}

type AccountUserProfilesListResponse struct {
	// AccountUserProfiles: Account user profile collection
	AccountUserProfiles []*AccountUserProfile `json:"accountUserProfiles,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountUserProfilesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AccountsListResponse struct {
	// Accounts: Account collection
	Accounts []*Account `json:"accounts,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#accountsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Activities struct {
	// Filters: List of activity filters. The dimension values need to be
	// all either of type "dfa:activity" or "dfa:activityGroup".
	Filters []*DimensionValue `json:"filters,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#activities.
	Kind string `json:"kind,omitempty"`

	// MetricNames: List of names of floodlight activity metrics.
	MetricNames []string `json:"metricNames,omitempty"`
}

type Ad struct {
	// AccountId: Account ID of this ad. This is a read-only field that can
	// be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Active: Whether this ad is active.
	Active bool `json:"active,omitempty"`

	// AdvertiserId: Advertiser ID of this ad. This is a required field on
	// insertion.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// Archived: Whether this ad is archived.
	Archived bool `json:"archived,omitempty"`

	// AudienceSegmentId: Audience segment ID that is being targeted for
	// this ad. Applicable when type is AD_SERVING_STANDARD_AD.
	AudienceSegmentId int64 `json:"audienceSegmentId,omitempty,string"`

	// CampaignId: Campaign ID of this ad. This is a required field on
	// insertion.
	CampaignId int64 `json:"campaignId,omitempty,string"`

	// CampaignIdDimensionValue: Dimension value for the ID of the campaign.
	// This is a read-only, auto-generated field.
	CampaignIdDimensionValue *DimensionValue `json:"campaignIdDimensionValue,omitempty"`

	// ClickThroughUrl: Click-through URL for this ad. This is a required
	// field on insertion. Applicable when type is AD_SERVING_CLICK_TRACKER.
	ClickThroughUrl *ClickThroughUrl `json:"clickThroughUrl,omitempty"`

	// ClickThroughUrlSuffixProperties: Click-through URL suffix properties
	// for this ad. Applies to the URL in the ad or (if overriding ad
	// properties) the URL in the creative.
	ClickThroughUrlSuffixProperties *ClickThroughUrlSuffixProperties `json:"clickThroughUrlSuffixProperties,omitempty"`

	// Comments: Comments for this ad.
	Comments string `json:"comments,omitempty"`

	// Compatibility: Compatibility of this ad. Applicable when type is
	// AD_SERVING_DEFAULT_AD. WEB and WEB_INTERSTITIAL refer to rendering
	// either on desktop or on mobile devices for regular or interstitial
	// ads, respectively. APP and APP_INTERSTITIAL are for rendering in
	// mobile apps. IN_STREAM_VIDEO refers to rendering an in-stream video
	// ads developed with the VAST standard.
	Compatibility string `json:"compatibility,omitempty"`

	// CreateInfo: Information about the creation of this ad.This is a
	// read-only field.
	CreateInfo *LastModifiedInfo `json:"createInfo,omitempty"`

	// CreativeGroupAssignments: Creative group assignments for this ad.
	// Applicable when type is AD_SERVING_CLICK_TRACKER. Only one assignment
	// per creative group number is allowed for a maximum of two
	// assignments.
	CreativeGroupAssignments []*CreativeGroupAssignment `json:"creativeGroupAssignments,omitempty"`

	// CreativeRotation: Creative rotation for this ad. Applicable when type
	// is AD_SERVING_DEFAULT_AD, AD_SERVING_STANDARD_AD, or
	// AD_SERVING_TRACKING. When type is AD_SERVING_DEFAULT_AD, this field
	// should have exactly one creativeAssignment.
	CreativeRotation *CreativeRotation `json:"creativeRotation,omitempty"`

	// DayPartTargeting: Time and day targeting information for this ad.
	// Applicable when type is AD_SERVING_STANDARD_AD.
	DayPartTargeting *DayPartTargeting `json:"dayPartTargeting,omitempty"`

	// DefaultClickThroughEventTagProperties: Default click-through event
	// tag properties for this ad.
	DefaultClickThroughEventTagProperties *DefaultClickThroughEventTagProperties `json:"defaultClickThroughEventTagProperties,omitempty"`

	// DeliverySchedule: Delivery schedule information for this ad.
	// Applicable when type is AD_SERVING_STANDARD_AD or
	// AD_SERVING_TRACKING. This field along with subfields priority and
	// impressionRatio are required on insertion when type is
	// AD_SERVING_STANDARD_AD.
	DeliverySchedule *DeliverySchedule `json:"deliverySchedule,omitempty"`

	// DynamicClickTracker: Whether this ad is a dynamic click tracker.
	// Applicable when type is AD_SERVING_CLICK_TRACKER. This is a required
	// field on insert, and is read-only after insert.
	DynamicClickTracker bool `json:"dynamicClickTracker,omitempty"`

	// EndTime: Date and time that this ad should stop serving. Must be
	// later than the start time. This is a required field on insertion.
	EndTime string `json:"endTime,omitempty"`

	// EventTagOverrides: Event tag overrides for this ad.
	EventTagOverrides []*EventTagOverride `json:"eventTagOverrides,omitempty"`

	// GeoTargeting: Geographical targeting information for this
	// ad.Applicable when type is AD_SERVING_STANDARD_AD.
	GeoTargeting *GeoTargeting `json:"geoTargeting,omitempty"`

	// Id: ID of this ad. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this ad. This is a
	// read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// KeyValueTargetingExpression: Key-value targeting information for this
	// ad. Applicable when type is AD_SERVING_STANDARD_AD.
	KeyValueTargetingExpression *KeyValueTargetingExpression `json:"keyValueTargetingExpression,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#ad".
	Kind string `json:"kind,omitempty"`

	// LastModifiedInfo: Information about the most recent modification of
	// this ad. This is a read-only field.
	LastModifiedInfo *LastModifiedInfo `json:"lastModifiedInfo,omitempty"`

	// Name: Name of this ad. This is a required field and must be less than
	// 256 characters long.
	Name string `json:"name,omitempty"`

	// PlacementAssignments: Placement assignments for this ad.
	PlacementAssignments []*PlacementAssignment `json:"placementAssignments,omitempty"`

	// Remarketing_list_expression: Applicable when type is
	// AD_SERVING_STANDARD_AD. Remarketing list targeting expression for
	// this ad.
	Remarketing_list_expression *ListTargetingExpression `json:"remarketing_list_expression,omitempty"`

	// Size: Size of this ad. Applicable when type is AD_SERVING_DEFAULT_AD.
	Size *Size `json:"size,omitempty"`

	// SslCompliant: Whether this ad is ssl compliant. This is a read-only
	// field that is auto-generated when the ad is inserted or updated.
	SslCompliant bool `json:"sslCompliant,omitempty"`

	// SslRequired: Whether this ad requires ssl. This is a read-only field
	// that is auto-generated when the ad is inserted or updated.
	SslRequired bool `json:"sslRequired,omitempty"`

	// StartTime: Date and time that this ad should start serving. If
	// creating an ad, this field must be a time in the future. This is a
	// required field on insertion.
	StartTime string `json:"startTime,omitempty"`

	// SubaccountId: Subaccount ID of this ad. This is a read-only field
	// that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TechnologyTargeting: Technology platform targeting information for
	// this ad. Applicable when type is AD_SERVING_STANDARD_AD.
	TechnologyTargeting *TechnologyTargeting `json:"technologyTargeting,omitempty"`

	// Type: Type of ad. This is a required field on insertion. Note that
	// default ads (AD_SERVING_DEFAULT_AD) cannot be created directly (see
	// Creative resource).
	Type string `json:"type,omitempty"`
}

type AdsListResponse struct {
	// Ads: Ad collection
	Ads []*Ad `json:"ads,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#adsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Advertiser struct {
	// AccountId: Account ID of this advertiser.This is a read-only field
	// that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserGroupId: ID of the advertiser group this advertiser belongs
	// to. You can group advertisers for reporting purposes, allowing you to
	// see aggregated information for all advertisers in each group.
	AdvertiserGroupId int64 `json:"advertiserGroupId,omitempty,string"`

	// ClickThroughUrlSuffix: Suffix added to click-through URL of ad
	// creative associations under this advertiser. Must be less than 129
	// characters long.
	ClickThroughUrlSuffix string `json:"clickThroughUrlSuffix,omitempty"`

	// DefaultClickThroughEventTagId: ID of the click-through event tag to
	// apply by default to the landing pages of this advertiser's campaigns.
	DefaultClickThroughEventTagId int64 `json:"defaultClickThroughEventTagId,omitempty,string"`

	// DefaultEmail: Default email address used in sender field for tag
	// emails.
	DefaultEmail string `json:"defaultEmail,omitempty"`

	// FloodlightConfigurationId: Floodlight configuration ID of this
	// advertiser. The floodlight configuration ID will be created
	// automatically, so on insert this field should be left blank. This
	// field can be set to another advertiser's floodlight configuration ID
	// in order to share that advertiser's floodlight configuration with
	// this advertiser, so long as:
	// - This advertiser's original floodlight
	// configuration is not already associated with floodlight activities or
	// floodlight activity groups.
	// - This advertiser's original floodlight
	// configuration is not already shared with another advertiser.
	FloodlightConfigurationId int64 `json:"floodlightConfigurationId,omitempty,string"`

	// FloodlightConfigurationIdDimensionValue: Dimension value for the ID
	// of the floodlight configuration. This is a read-only, auto-generated
	// field.
	FloodlightConfigurationIdDimensionValue *DimensionValue `json:"floodlightConfigurationIdDimensionValue,omitempty"`

	// Id: ID of this advertiser. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this advertiser. This
	// is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#advertiser".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this advertiser. This is a required field and must be
	// less than 256 characters long and unique among advertisers of the
	// same account.
	Name string `json:"name,omitempty"`

	// Status: Status of this advertiser.
	Status string `json:"status,omitempty"`

	// SubaccountId: Subaccount ID of this advertiser.This is a read-only
	// field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`
}

type AdvertiserGroup struct {
	// AccountId: Account ID of this advertiser group. This is a read-only
	// field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Id: ID of this advertiser group. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#advertiserGroup".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this advertiser group. This is a required field and
	// must be less than 256 characters long and unique among advertiser
	// groups of the same account.
	Name string `json:"name,omitempty"`
}

type AdvertiserGroupsListResponse struct {
	// AdvertiserGroups: Advertiser group collection
	AdvertiserGroups []*AdvertiserGroup `json:"advertiserGroups,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#advertiserGroupsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AdvertisersListResponse struct {
	// Advertisers: Advertiser collection
	Advertisers []*Advertiser `json:"advertisers,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#advertisersListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AudienceSegment struct {
	// Allocation: Weight allocated to this segment. Must be between 1 and
	// 1000. The weight assigned will be understood in proportion to the
	// weights assigned to other segments in the same segment group.
	Allocation int64 `json:"allocation,omitempty"`

	// Id: ID of this audience segment. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Name: Name of this audience segment. This is a required field and
	// must be less than 65 characters long.
	Name string `json:"name,omitempty"`
}

type AudienceSegmentGroup struct {
	// AudienceSegments: Audience segments assigned to this group. The
	// number of segments must be between 2 and 100.
	AudienceSegments []*AudienceSegment `json:"audienceSegments,omitempty"`

	// Id: ID of this audience segment group. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Name: Name of this audience segment group. This is a required field
	// and must be less than 65 characters long.
	Name string `json:"name,omitempty"`
}

type Browser struct {
	// BrowserVersionId: ID referring to this grouping of browser and
	// version numbers. This is the ID used for targeting.
	BrowserVersionId int64 `json:"browserVersionId,omitempty,string"`

	// DartId: DART ID of this browser. This is the ID used when generating
	// reports.
	DartId int64 `json:"dartId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#browser".
	Kind string `json:"kind,omitempty"`

	// MajorVersion: Major version number (leftmost number) of this browser.
	// For example, for Chrome 5.0.376.86 beta, this field should be set to
	// 5. An asterisk (*) may be used to target any version number, and a
	// question mark (?) may be used to target cases where the version
	// number cannot be identified. For example, Chrome *.* targets any
	// version of Chrome: 1.2, 2.5, 3.5, and so on. Chrome 3.* targets
	// Chrome 3.1, 3.5, but not 4.0. Firefox ?.? targets cases where the ad
	// server knows the browser is Firefox but can't tell which version it
	// is.
	MajorVersion string `json:"majorVersion,omitempty"`

	// MinorVersion: Minor version number (number after first dot on left)
	// of this browser. For example, for Chrome 5.0.375.86 beta, this field
	// should be set to 0. An asterisk (*) may be used to target any version
	// number, and a question mark (?) may be used to target cases where the
	// version number cannot be identified. For example, Chrome *.* targets
	// any version of Chrome: 1.2, 2.5, 3.5, and so on. Chrome 3.* targets
	// Chrome 3.1, 3.5, but not 4.0. Firefox ?.? targets cases where the ad
	// server knows the browser is Firefox but can't tell which version it
	// is.
	MinorVersion string `json:"minorVersion,omitempty"`

	// Name: Name of this browser.
	Name string `json:"name,omitempty"`
}

type BrowsersListResponse struct {
	// Browsers: Browser collection
	Browsers []*Browser `json:"browsers,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#browsersListResponse".
	Kind string `json:"kind,omitempty"`
}

type Campaign struct {
	// AccountId: Account ID of this campaign. This is a read-only field
	// that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdditionalCreativeOptimizationConfigurations: Additional creative
	// optimization configurations for the campaign.
	AdditionalCreativeOptimizationConfigurations []*CreativeOptimizationConfiguration `json:"additionalCreativeOptimizationConfigurations,omitempty"`

	// AdvertiserGroupId: Advertiser group ID of the associated advertiser.
	AdvertiserGroupId int64 `json:"advertiserGroupId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this campaign. This is a required
	// field.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the advertiser ID of
	// this campaign. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// Archived: Whether this campaign has been archived.
	Archived bool `json:"archived,omitempty"`

	// AudienceSegmentGroups: Audience segment groups assigned to this
	// campaign. Cannot have more than 300 segment groups.
	AudienceSegmentGroups []*AudienceSegmentGroup `json:"audienceSegmentGroups,omitempty"`

	// BillingInvoiceCode: Billing invoice code included in the DCM client
	// billing invoices associated with the campaign.
	BillingInvoiceCode string `json:"billingInvoiceCode,omitempty"`

	// ClickThroughUrlSuffixProperties: Click-through URL suffix override
	// properties for this campaign.
	ClickThroughUrlSuffixProperties *ClickThroughUrlSuffixProperties `json:"clickThroughUrlSuffixProperties,omitempty"`

	// Comment: Arbitrary comments about this campaign. Must be less than
	// 256 characters long.
	Comment string `json:"comment,omitempty"`

	// ComscoreVceEnabled: Whether comScore vCE reports are enabled for this
	// campaign.
	ComscoreVceEnabled bool `json:"comscoreVceEnabled,omitempty"`

	// CreateInfo: Information about the creation of this campaign. This is
	// a read-only field.
	CreateInfo *LastModifiedInfo `json:"createInfo,omitempty"`

	// CreativeGroupIds: List of creative group IDs that are assigned to the
	// campaign.
	CreativeGroupIds googleapi.Int64s `json:"creativeGroupIds,omitempty"`

	// CreativeOptimizationConfiguration: Creative optimization
	// configuration for the campaign.
	CreativeOptimizationConfiguration *CreativeOptimizationConfiguration `json:"creativeOptimizationConfiguration,omitempty"`

	// DefaultClickThroughEventTagProperties: Click-through event tag ID
	// override properties for this campaign.
	DefaultClickThroughEventTagProperties *DefaultClickThroughEventTagProperties `json:"defaultClickThroughEventTagProperties,omitempty"`

	// EndDate: Date on which the campaign will stop running. On insert, the
	// end date must be today or a future date. The end date must be later
	// than or be the same as the start date. If, for example, you set
	// 6/25/2015 as both the start and end dates, the effective campaign run
	// date is just that day only, 6/25/2015. The hours, minutes, and
	// seconds of the end date should not be set, as doing so will result in
	// an error. This is a required field.
	EndDate string `json:"endDate,omitempty"`

	// EventTagOverrides: Overrides that can be used to activate or
	// deactivate advertiser event tags.
	EventTagOverrides []*EventTagOverride `json:"eventTagOverrides,omitempty"`

	// ExternalId: External ID for this campaign.
	ExternalId string `json:"externalId,omitempty"`

	// Id: ID of this campaign. This is a read-only auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this campaign. This
	// is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#campaign".
	Kind string `json:"kind,omitempty"`

	// LastModifiedInfo: Information about the most recent modification of
	// this campaign. This is a read-only field.
	LastModifiedInfo *LastModifiedInfo `json:"lastModifiedInfo,omitempty"`

	// LookbackConfiguration: Lookback window settings for the campaign.
	LookbackConfiguration *LookbackConfiguration `json:"lookbackConfiguration,omitempty"`

	// Name: Name of this campaign. This is a required field and must be
	// less than 256 characters long and unique among campaigns of the same
	// advertiser.
	Name string `json:"name,omitempty"`

	// NielsenOcrEnabled: Whether Nielsen reports are enabled for this
	// campaign.
	NielsenOcrEnabled bool `json:"nielsenOcrEnabled,omitempty"`

	// StartDate: Date on which the campaign starts running. The start date
	// can be any date. The hours, minutes, and seconds of the start date
	// should not be set, as doing so will result in an error. This is a
	// required field.
	StartDate string `json:"startDate,omitempty"`

	// SubaccountId: Subaccount ID of this campaign. This is a read-only
	// field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TraffickerEmails: Campaign trafficker contact emails.
	TraffickerEmails []string `json:"traffickerEmails,omitempty"`
}

type CampaignCreativeAssociation struct {
	// CreativeId: ID of the creative associated with the campaign. This is
	// a required field.
	CreativeId int64 `json:"creativeId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#campaignCreativeAssociation".
	Kind string `json:"kind,omitempty"`
}

type CampaignCreativeAssociationsListResponse struct {
	// CampaignCreativeAssociations: Campaign creative association
	// collection
	CampaignCreativeAssociations []*CampaignCreativeAssociation `json:"campaignCreativeAssociations,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#campaignCreativeAssociationsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type CampaignsListResponse struct {
	// Campaigns: Campaign collection
	Campaigns []*Campaign `json:"campaigns,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#campaignsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ChangeLog struct {
	// AccountId: Account ID of the modified object.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Action: Action which caused the change.
	Action string `json:"action,omitempty"`

	// ChangeTime: Time when the object was modified.
	ChangeTime string `json:"changeTime,omitempty"`

	// FieldName: Field name of the object which changed.
	FieldName string `json:"fieldName,omitempty"`

	// Id: ID of this change log.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#changeLog".
	Kind string `json:"kind,omitempty"`

	// NewValue: New value of the object field.
	NewValue string `json:"newValue,omitempty"`

	// ObjectId: ID of the object of this change log. The object could be a
	// campaign, placement, ad, or other type.
	ObjectId int64 `json:"objectId,omitempty,string"`

	// ObjectType: Object type of the change log.
	ObjectType string `json:"objectType,omitempty"`

	// OldValue: Old value of the object field.
	OldValue string `json:"oldValue,omitempty"`

	// SubaccountId: Subaccount ID of the modified object.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TransactionId: Transaction ID of this change log. When a single API
	// call results in many changes, each change will have a separate ID in
	// the change log but will share the same transactionId.
	TransactionId int64 `json:"transactionId,omitempty,string"`

	// UserProfileId: ID of the user who modified the object.
	UserProfileId int64 `json:"userProfileId,omitempty,string"`

	// UserProfileName: User profile name of the user who modified the
	// object.
	UserProfileName string `json:"userProfileName,omitempty"`
}

type ChangeLogsListResponse struct {
	// ChangeLogs: Change log collection
	ChangeLogs []*ChangeLog `json:"changeLogs,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#changeLogsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type CitiesListResponse struct {
	// Cities: City collection
	Cities []*City `json:"cities,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#citiesListResponse".
	Kind string `json:"kind,omitempty"`
}

type City struct {
	// CountryCode: Country code of the country to which this city belongs.
	CountryCode string `json:"countryCode,omitempty"`

	// CountryDartId: DART ID of the country to which this city belongs.
	CountryDartId int64 `json:"countryDartId,omitempty,string"`

	// DartId: DART ID of this city. This is the ID used for targeting and
	// generating reports.
	DartId int64 `json:"dartId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#city".
	Kind string `json:"kind,omitempty"`

	// MetroCode: Metro region code of the metro region (DMA) to which this
	// city belongs.
	MetroCode string `json:"metroCode,omitempty"`

	// MetroDmaId: ID of the metro region (DMA) to which this city belongs.
	MetroDmaId int64 `json:"metroDmaId,omitempty,string"`

	// Name: Name of this city.
	Name string `json:"name,omitempty"`

	// RegionCode: Region code of the region to which this city belongs.
	RegionCode string `json:"regionCode,omitempty"`

	// RegionDartId: DART ID of the region to which this city belongs.
	RegionDartId int64 `json:"regionDartId,omitempty,string"`
}

type ClickTag struct {
	// EventName: Advertiser event name associated with the click tag. This
	// field is used by ENHANCED_BANNER, ENHANCED_IMAGE, and HTML5_BANNER
	// creatives.
	EventName string `json:"eventName,omitempty"`

	// Name: Parameter name for the specified click tag. For ENHANCED_IMAGE
	// creative assets, this field must match the value of the creative
	// asset's creativeAssetId.name field.
	Name string `json:"name,omitempty"`

	// Value: Parameter value for the specified click tag. This field
	// contains a click-through url.
	Value string `json:"value,omitempty"`
}

type ClickThroughUrl struct {
	// CustomClickThroughUrl: Custom click-through URL. Applicable if the
	// defaultLandingPage field is set to false and the landingPageId field
	// is left unset.
	CustomClickThroughUrl string `json:"customClickThroughUrl,omitempty"`

	// DefaultLandingPage: Whether the campaign default landing page is
	// used.
	DefaultLandingPage bool `json:"defaultLandingPage,omitempty"`

	// LandingPageId: ID of the landing page for the click-through URL.
	// Applicable if the defaultLandingPage field is set to false.
	LandingPageId int64 `json:"landingPageId,omitempty,string"`
}

type ClickThroughUrlSuffixProperties struct {
	// ClickThroughUrlSuffix: Click-through URL suffix to apply to all ads
	// in this entity's scope. Must be less than 128 characters long.
	ClickThroughUrlSuffix string `json:"clickThroughUrlSuffix,omitempty"`

	// OverrideInheritedSuffix: Whether this entity should override the
	// inherited click-through URL suffix with its own defined value.
	OverrideInheritedSuffix bool `json:"overrideInheritedSuffix,omitempty"`
}

type CompanionClickThroughOverride struct {
	// ClickThroughUrl: Click-through URL of this companion click-through
	// override.
	ClickThroughUrl *ClickThroughUrl `json:"clickThroughUrl,omitempty"`

	// CreativeId: ID of the creative for this companion click-through
	// override.
	CreativeId int64 `json:"creativeId,omitempty,string"`
}

type CompatibleFields struct {
	// CrossDimensionReachReportCompatibleFields: Contains items that are
	// compatible to be selected for a report of type
	// "CROSS_DIMENSION_REACH".
	CrossDimensionReachReportCompatibleFields *CrossDimensionReachReportCompatibleFields `json:"crossDimensionReachReportCompatibleFields,omitempty"`

	// FloodlightReportCompatibleFields: Contains items that are compatible
	// to be selected for a report of type "FLOODLIGHT".
	FloodlightReportCompatibleFields *FloodlightReportCompatibleFields `json:"floodlightReportCompatibleFields,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#compatibleFields.
	Kind string `json:"kind,omitempty"`

	// PathToConversionReportCompatibleFields: Contains items that are
	// compatible to be selected for a report of type "PATH_TO_CONVERSION".
	PathToConversionReportCompatibleFields *PathToConversionReportCompatibleFields `json:"pathToConversionReportCompatibleFields,omitempty"`

	// ReachReportCompatibleFields: Contains items that are compatible to be
	// selected for a report of type "REACH".
	ReachReportCompatibleFields *ReachReportCompatibleFields `json:"reachReportCompatibleFields,omitempty"`

	// ReportCompatibleFields: Contains items that are compatible to be
	// selected for a report of type "STANDARD".
	ReportCompatibleFields *ReportCompatibleFields `json:"reportCompatibleFields,omitempty"`
}

type ConnectionType struct {
	// Id: ID of this connection type.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#connectionType".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this connection type.
	Name string `json:"name,omitempty"`
}

type ConnectionTypesListResponse struct {
	// ConnectionTypes: Connection Type Collection.
	ConnectionTypes []*ConnectionType `json:"connectionTypes,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#connectionTypesListResponse".
	Kind string `json:"kind,omitempty"`
}

type ContentCategoriesListResponse struct {
	// ContentCategories: Content category collection
	ContentCategories []*ContentCategory `json:"contentCategories,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#contentCategoriesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ContentCategory struct {
	// AccountId: Account ID of this content category. This is a read-only
	// field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Description: Description of this content category.
	Description string `json:"description,omitempty"`

	// Id: ID of this content category. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#contentCategory".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this content category. This is a required field and
	// must be less than 256 characters long and unique among content
	// categories of the same account.
	Name string `json:"name,omitempty"`
}

type CountriesListResponse struct {
	// Countries: Country collection
	Countries []*Country `json:"countries,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#countriesListResponse".
	Kind string `json:"kind,omitempty"`
}

type Country struct {
	// CountryCode: Country code.
	CountryCode string `json:"countryCode,omitempty"`

	// DartId: DART ID of this country. This is the ID used for targeting
	// and generating reports.
	DartId int64 `json:"dartId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#country".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this country.
	Name string `json:"name,omitempty"`

	// SslEnabled: Whether ad serving supports secure servers in this
	// country.
	SslEnabled bool `json:"sslEnabled,omitempty"`
}

type Creative struct {
	// AccountId: Account ID of this creative. This field, if left unset,
	// will be auto-generated for both insert and update operations.
	// Applicable to all creative types.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Active: Whether the creative is active. Applicable to all creative
	// types.
	Active bool `json:"active,omitempty"`

	// AdParameters: Ad parameters user for VPAID creative. This is a
	// read-only field. Applicable to the following creative types: all
	// VPAID.
	AdParameters string `json:"adParameters,omitempty"`

	// AdTagKeys: Keywords for a Rich Media creative. Keywords let you
	// customize the creative settings of a Rich Media ad running on your
	// site without having to contact the advertiser. You can use keywords
	// to dynamically change the look or functionality of a creative.
	// Applicable to the following creative types: all RICH_MEDIA, and all
	// VPAID.
	AdTagKeys []string `json:"adTagKeys,omitempty"`

	// AdvertiserId: Advertiser ID of this creative. This is a required
	// field. Applicable to all creative types.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AllowScriptAccess: Whether script access is allowed for this
	// creative. This is a read-only and deprecated field which will
	// automatically be set to true on update. Applicable to the following
	// creative types: FLASH_INPAGE.
	AllowScriptAccess bool `json:"allowScriptAccess,omitempty"`

	// Archived: Whether the creative is archived. Applicable to all
	// creative types.
	Archived bool `json:"archived,omitempty"`

	// ArtworkType: Type of artwork used for the creative. This is a
	// read-only field. Applicable to the following creative types: all
	// RICH_MEDIA, and all VPAID.
	ArtworkType string `json:"artworkType,omitempty"`

	// AuthoringTool: Authoring tool for HTML5 banner creatives. This is a
	// read-only field. Applicable to the following creative types:
	// HTML5_BANNER.
	AuthoringTool string `json:"authoringTool,omitempty"`

	// Auto_advance_images: Whether images are automatically advanced for
	// enhanced image creatives. Applicable to the following creative types:
	// ENHANCED_IMAGE.
	Auto_advance_images bool `json:"auto_advance_images,omitempty"`

	// BackgroundColor: The 6-character HTML color code, beginning with #,
	// for the background of the window area where the Flash file is
	// displayed. Default is white. Applicable to the following creative
	// types: FLASH_INPAGE.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// BackupImageClickThroughUrl: Click-through URL for backup image.
	// Applicable to the following creative types: ENHANCED_BANNER,
	// FLASH_INPAGE, and HTML5_BANNER.
	BackupImageClickThroughUrl string `json:"backupImageClickThroughUrl,omitempty"`

	// BackupImageFeatures: List of feature dependencies that will cause a
	// backup image to be served if the browser that serves the ad does not
	// support them. Feature dependencies are features that a browser must
	// be able to support in order to render your HTML5 creative asset
	// correctly. This field is initially auto-generated to contain all
	// features detected by DCM for all the assets of this creative and can
	// then be modified by the client. To reset this field, copy over all
	// the creativeAssets' detected features. Applicable to the following
	// creative types: ENHANCED_BANNER and HTML5_BANNER.
	BackupImageFeatures []string `json:"backupImageFeatures,omitempty"`

	// BackupImageReportingLabel: Reporting label used for HTML5 banner
	// backup image. Applicable to the following creative types:
	// ENHANCED_BANNER.
	BackupImageReportingLabel string `json:"backupImageReportingLabel,omitempty"`

	// BackupImageTargetWindow: Target window for backup image. Applicable
	// to the following creative types: ENHANCED_BANNER, FLASH_INPAGE, and
	// HTML5_BANNER.
	BackupImageTargetWindow *TargetWindow `json:"backupImageTargetWindow,omitempty"`

	// ClickTags: Click tags of the creative. For ENHANCED_BANNER,
	// FLASH_INPAGE, and HTML5_BANNER creatives, this is a subset of
	// detected click tags for the assets associated with this creative.
	// After creating a flash asset, detected click tags will be returned in
	// the creativeAssetMetadata. When inserting the creative, populate the
	// creative clickTags field using the creativeAssetMetadata.clickTags
	// field. For ENHANCED_IMAGE creatives, there should be exactly one
	// entry in this list for each image creative asset. A click tag is
	// matched with a corresponding creative asset by matching the
	// clickTag.name field with the creativeAsset.assetIdentifier.name
	// field. Applicable to the following creative types: ENHANCED_BANNER,
	// ENHANCED_IMAGE, FLASH_INPAGE, HTML5_BANNER.
	ClickTags []*ClickTag `json:"clickTags,omitempty"`

	// CommercialId: Industry standard ID assigned to creative for reach and
	// frequency. Applicable to the following creative types: INSTREAM_VIDEO
	// and all VPAID.
	CommercialId string `json:"commercialId,omitempty"`

	// CompanionCreatives: List of companion creatives assigned to an
	// in-Stream videocreative. Acceptable values include IDs of existing
	// flash and image creatives. Applicable to the following creative
	// types: INSTREAM_VIDEO and all VPAID.
	CompanionCreatives googleapi.Int64s `json:"companionCreatives,omitempty"`

	// Compatibility: Compatibilities associated with this creative. This is
	// a read-only field. WEB and WEB_INTERSTITIAL refer to rendering either
	// on desktop or on mobile devices for regular or interstitial ads,
	// respectively. APP and APP_INTERSTITIAL are for rendering in mobile
	// apps. IN_STREAM_VIDEO refers to rendering in in-stream video ads
	// developed with the VAST standard. Applicable to all creative
	// types.
	//
	// Acceptable values are:
	// - "APP"
	// - "APP_INTERSTITIAL"
	// -
	// "IN_STREAM_VIDEO"
	// - "WEB"
	// - "WEB_INTERSTITIAL"
	Compatibility []string `json:"compatibility,omitempty"`

	// CounterCustomEvents: List of counter events configured for the
	// creative. Applicable to the following creative types: all RICH_MEDIA,
	// and all VPAID.
	CounterCustomEvents []*CreativeCustomEvent `json:"counterCustomEvents,omitempty"`

	// CreativeAssets: Assets associated with a creative. Applicable to all
	// but the following creative types: INTERNAL_REDIRECT,
	// INTERSTITIAL_INTERNAL_REDIRECT, and REDIRECT
	CreativeAssets []*CreativeAsset `json:"creativeAssets,omitempty"`

	// CreativeFieldAssignments: Creative field assignments for this
	// creative. Applicable to all creative types.
	CreativeFieldAssignments []*CreativeFieldAssignment `json:"creativeFieldAssignments,omitempty"`

	// CustomKeyValues: Custom key-values for a Rich Media creative.
	// Key-values let you customize the creative settings of a Rich Media ad
	// running on your site without having to contact the advertiser. You
	// can use key-values to dynamically change the look or functionality of
	// a creative. Applicable to the following creative types: all
	// RICH_MEDIA, and all VPAID.
	CustomKeyValues []string `json:"customKeyValues,omitempty"`

	// ExitCustomEvents: List of exit events configured for the creative.
	// Applicable to the following creative types: all RICH_MEDIA, and all
	// VPAID.
	ExitCustomEvents []*CreativeCustomEvent `json:"exitCustomEvents,omitempty"`

	// FsCommand: OpenWindow FSCommand of this creative. This lets the SWF
	// file communicate with either Flash Player or the program hosting
	// Flash Player, such as a web browser. This is only triggered if
	// allowScriptAccess field is true. Applicable to the following creative
	// types: FLASH_INPAGE.
	FsCommand *FsCommand `json:"fsCommand,omitempty"`

	// HtmlCode: HTML code for the creative. This is a required field when
	// applicable. This field is ignored if htmlCodeLocked is false.
	// Applicable to the following creative types: all CUSTOM, FLASH_INPAGE,
	// and HTML5_BANNER, and all RICH_MEDIA.
	HtmlCode string `json:"htmlCode,omitempty"`

	// HtmlCodeLocked: Whether HTML code is DCM-generated or manually
	// entered. Set to true to ignore changes to htmlCode. Applicable to the
	// following creative types: FLASH_INPAGE and HTML5_BANNER.
	HtmlCodeLocked bool `json:"htmlCodeLocked,omitempty"`

	// Id: ID of this creative. This is a read-only, auto-generated field.
	// Applicable to all creative types.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this creative. This
	// is a read-only field. Applicable to all creative types.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creative".
	Kind string `json:"kind,omitempty"`

	// LastModifiedInfo: Creative last modification information. This is a
	// read-only field. Applicable to all creative types.
	LastModifiedInfo *LastModifiedInfo `json:"lastModifiedInfo,omitempty"`

	// LatestTraffickedCreativeId: Latest Studio trafficked creative ID
	// associated with rich media and VPAID creatives. This is a read-only
	// field. Applicable to the following creative types: all RICH_MEDIA,
	// and all VPAID.
	LatestTraffickedCreativeId int64 `json:"latestTraffickedCreativeId,omitempty,string"`

	// Name: Name of the creative. This is a required field and must be less
	// than 256 characters long. Applicable to all creative types.
	Name string `json:"name,omitempty"`

	// OverrideCss: Override CSS value for rich media creatives. Applicable
	// to the following creative types: all RICH_MEDIA.
	OverrideCss string `json:"overrideCss,omitempty"`

	// RedirectUrl: URL of hosted image or another ad tag. This is a
	// required field when applicable. Applicable to the following creative
	// types: INTERNAL_REDIRECT, INTERSTITIAL_INTERNAL_REDIRECT, and
	// REDIRECT
	RedirectUrl string `json:"redirectUrl,omitempty"`

	// RenderingId: ID of current rendering version. This is a read-only
	// field. Applicable to all creative types.
	RenderingId int64 `json:"renderingId,omitempty,string"`

	// RenderingIdDimensionValue: Dimension value for the rendering ID of
	// this creative. This is a read-only field. Applicable to all creative
	// types.
	RenderingIdDimensionValue *DimensionValue `json:"renderingIdDimensionValue,omitempty"`

	// RequiredFlashPluginVersion: The minimum required Flash plugin version
	// for this creative. For example, 11.2.202.235. This is a read-only
	// field. Applicable to the following creative types: all RICH_MEDIA,
	// and all VPAID.
	RequiredFlashPluginVersion string `json:"requiredFlashPluginVersion,omitempty"`

	// RequiredFlashVersion: The internal Flash version for this creative as
	// calculated by DoubleClick Studio. This is a read-only field.
	// Applicable to the following creative types: FLASH_INPAGE,
	// ENHANCED_BANNER, all RICH_MEDIA, and all VPAID.
	RequiredFlashVersion int64 `json:"requiredFlashVersion,omitempty"`

	// Size: Size associated with this creative. When inserting or updating
	// a creative either the size ID field or size width and height fields
	// can be used. This is a required field when applicable; however for
	// IMAGE and FLASH_INPAGE creatives, if left blank, this field will be
	// automatically set using the actual size of the associated image
	// assets. Applicable to the following creative types: ENHANCED_BANNER,
	// ENHANCED_IMAGE, FLASH_INPAGE, HTML5_BANNER, IMAGE, and all
	// RICH_MEDIA.
	Size *Size `json:"size,omitempty"`

	// Skippable: Whether the user can choose to skip the creative.
	// Applicable to the following creative types: INSTREAM_VIDEO.
	Skippable bool `json:"skippable,omitempty"`

	// SslCompliant: Whether the creative is SSL-compliant. This is a
	// read-only field. Applicable to all creative types.
	SslCompliant bool `json:"sslCompliant,omitempty"`

	// StudioAdvertiserId: Studio advertiser ID associated with rich media
	// and VPAID creatives. This is a read-only field. Applicable to the
	// following creative types: all RICH_MEDIA, and all VPAID.
	StudioAdvertiserId int64 `json:"studioAdvertiserId,omitempty,string"`

	// StudioCreativeId: Studio creative ID associated with rich media and
	// VPAID creatives. This is a read-only field. Applicable to the
	// following creative types: all RICH_MEDIA, and all VPAID.
	StudioCreativeId int64 `json:"studioCreativeId,omitempty,string"`

	// StudioTraffickedCreativeId: Studio trafficked creative ID associated
	// with rich media and VPAID creatives. This is a read-only field.
	// Applicable to the following creative types: all RICH_MEDIA, and all
	// VPAID.
	StudioTraffickedCreativeId int64 `json:"studioTraffickedCreativeId,omitempty,string"`

	// SubaccountId: Subaccount ID of this creative. This field, if left
	// unset, will be auto-generated for both insert and update operations.
	// Applicable to all creative types.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// ThirdPartyBackupImageImpressionsUrl: Third-party URL used to record
	// backup image impressions. Applicable to the following creative types:
	// all RICH_MEDIA
	ThirdPartyBackupImageImpressionsUrl string `json:"thirdPartyBackupImageImpressionsUrl,omitempty"`

	// ThirdPartyRichMediaImpressionsUrl: Third-party URL used to record
	// rich media impressions. Applicable to the following creative types:
	// all RICH_MEDIA
	ThirdPartyRichMediaImpressionsUrl string `json:"thirdPartyRichMediaImpressionsUrl,omitempty"`

	// ThirdPartyUrls: Third-party URLs for tracking in-stream video
	// creative events. Applicable to the following creative types:
	// INSTREAM_VIDEO and all VPAID.
	ThirdPartyUrls []*ThirdPartyTrackingUrl `json:"thirdPartyUrls,omitempty"`

	// TimerCustomEvents: List of timer events configured for the creative.
	// Applicable to the following creative types: all RICH_MEDIA, and all
	// VPAID.
	TimerCustomEvents []*CreativeCustomEvent `json:"timerCustomEvents,omitempty"`

	// TotalFileSize: Combined size of all creative assets. This is a
	// read-only field. Applicable to the following creative types: all
	// RICH_MEDIA, and all VPAID.
	TotalFileSize int64 `json:"totalFileSize,omitempty,string"`

	// Type: Type of this creative.This is a required field. Applicable to
	// all creative types.
	Type string `json:"type,omitempty"`

	// Version: The version number helps you keep track of multiple versions
	// of your creative in your reports. The version number will always be
	// auto-generated during insert operations to start at 1. For tracking
	// creatives the version cannot be incremented and will always remain at
	// 1. For all other creative types the version can be incremented only
	// by 1 during update operations. In addition, the version will be
	// automatically incremented by 1 when undergoing Rich Media creative
	// merging. Applicable to all creative types.
	Version int64 `json:"version,omitempty"`

	// VideoDescription: Description of the video ad. Applicable to the
	// following creative types: INSTREAM_VIDEO and all VPAID.
	VideoDescription string `json:"videoDescription,omitempty"`

	// VideoDuration: Creative video duration in seconds. This is a
	// read-only field. Applicable to the following creative types:
	// INSTREAM_VIDEO, all RICH_MEDIA, and all VPAID.
	VideoDuration float64 `json:"videoDuration,omitempty"`
}

type CreativeAsset struct {
	// ActionScript3: Whether ActionScript3 is enabled for the flash asset.
	// This is a read-only field. Applicable to the following creative
	// types: FLASH_INPAGE and ENHANCED_BANNER.
	ActionScript3 bool `json:"actionScript3,omitempty"`

	// Active: Whether the video asset is active. This is a read-only field
	// for VPAID_NON_LINEAR assets. Applicable to the following creative
	// types: INSTREAM_VIDEO and all VPAID.
	Active bool `json:"active,omitempty"`

	// Alignment: Possible alignments for an asset. This is a read-only
	// field. Applicable to the following creative types:
	// RICH_MEDIA_MULTI_FLOATING.
	Alignment string `json:"alignment,omitempty"`

	// ArtworkType: Artwork type of rich media creative. This is a read-only
	// field. Applicable to the following creative types: all RICH_MEDIA.
	ArtworkType string `json:"artworkType,omitempty"`

	// AssetIdentifier: Identifier of this asset. This is the same
	// identifier returned during creative asset insert operation. This is a
	// required field. Applicable to all but the following creative types:
	// all REDIRECT and TRACKING_TEXT.
	AssetIdentifier *CreativeAssetId `json:"assetIdentifier,omitempty"`

	// BackupImageExit: Exit event configured for the backup image.
	// Applicable to the following creative types: all RICH_MEDIA.
	BackupImageExit *CreativeCustomEvent `json:"backupImageExit,omitempty"`

	// BitRate: Detected bit-rate for video asset. This is a read-only
	// field. Applicable to the following creative types: INSTREAM_VIDEO and
	// all VPAID.
	BitRate int64 `json:"bitRate,omitempty"`

	// ChildAssetType: Rich media child asset type. This is a read-only
	// field. Applicable to the following creative types: all VPAID.
	ChildAssetType string `json:"childAssetType,omitempty"`

	// CollapsedSize: Size of an asset when collapsed. This is a read-only
	// field. Applicable to the following creative types: all RICH_MEDIA and
	// all VPAID. Additionally, applicable to assets whose displayType is
	// ASSET_DISPLAY_TYPE_EXPANDING or ASSET_DISPLAY_TYPE_PEEL_DOWN.
	CollapsedSize *Size `json:"collapsedSize,omitempty"`

	// CustomStartTimeValue: Custom start time in seconds for making the
	// asset visible. Applicable to the following creative types: all
	// RICH_MEDIA.
	CustomStartTimeValue int64 `json:"customStartTimeValue,omitempty"`

	// DetectedFeatures: List of feature dependencies for the creative asset
	// that are detected by DCM. Feature dependencies are features that a
	// browser must be able to support in order to render your HTML5
	// creative correctly. This is a read-only, auto-generated field.
	// Applicable to the following creative types: ENHANCED_BANNER and
	// HTML5_BANNER.
	DetectedFeatures []string `json:"detectedFeatures,omitempty"`

	// DisplayType: Type of rich media asset. This is a read-only field.
	// Applicable to the following creative types: all RICH_MEDIA.
	DisplayType string `json:"displayType,omitempty"`

	// Duration: Duration in seconds for which an asset will be displayed.
	// Applicable to the following creative types: INSTREAM_VIDEO and
	// VPAID_LINEAR.
	Duration int64 `json:"duration,omitempty"`

	// DurationType: Duration type for which an asset will be displayed.
	// Applicable to the following creative types: all RICH_MEDIA.
	DurationType string `json:"durationType,omitempty"`

	// ExpandedDimension: Detected expanded dimension for video asset. This
	// is a read-only field. Applicable to the following creative types:
	// INSTREAM_VIDEO and all VPAID.
	ExpandedDimension *Size `json:"expandedDimension,omitempty"`

	// FileSize: File size associated with this creative asset. This is a
	// read-only field. Applicable to all but the following creative types:
	// all REDIRECT and TRACKING_TEXT.
	FileSize int64 `json:"fileSize,omitempty,string"`

	// FlashVersion: Flash version of the asset. This is a read-only field.
	// Applicable to the following creative types: FLASH_INPAGE,
	// ENHANCED_BANNER, all RICH_MEDIA, and all VPAID.
	FlashVersion int64 `json:"flashVersion,omitempty"`

	// HideFlashObjects: Whether to hide Flash objects flag for an asset.
	// Applicable to the following creative types: all RICH_MEDIA.
	HideFlashObjects bool `json:"hideFlashObjects,omitempty"`

	// HideSelectionBoxes: Whether to hide selection boxes flag for an
	// asset. Applicable to the following creative types: all RICH_MEDIA.
	HideSelectionBoxes bool `json:"hideSelectionBoxes,omitempty"`

	// HorizontallyLocked: Whether the asset is horizontally locked. This is
	// a read-only field. Applicable to the following creative types: all
	// RICH_MEDIA.
	HorizontallyLocked bool `json:"horizontallyLocked,omitempty"`

	// Id: Numeric ID of this creative asset. This is a required field and
	// should not be modified. Applicable to all but the following creative
	// types: all REDIRECT and TRACKING_TEXT.
	Id int64 `json:"id,omitempty,string"`

	// MimeType: Detected MIME type for video asset. This is a read-only
	// field. Applicable to the following creative types: INSTREAM_VIDEO and
	// all VPAID.
	MimeType string `json:"mimeType,omitempty"`

	// Offset: Offset position for an asset in collapsed mode. This is a
	// read-only field. Applicable to the following creative types: all
	// RICH_MEDIA and all VPAID. Additionally, only applicable to assets
	// whose displayType is ASSET_DISPLAY_TYPE_EXPANDING or
	// ASSET_DISPLAY_TYPE_PEEL_DOWN.
	Offset *OffsetPosition `json:"offset,omitempty"`

	// OriginalBackup: Whether the backup asset is original or changed by
	// the user in DCM. Applicable to the following creative types: all
	// RICH_MEDIA.
	OriginalBackup bool `json:"originalBackup,omitempty"`

	// Position: Offset position for an asset. Applicable to the following
	// creative types: all RICH_MEDIA.
	Position *OffsetPosition `json:"position,omitempty"`

	// PositionLeftUnit: Offset left unit for an asset. This is a read-only
	// field. Applicable to the following creative types: all RICH_MEDIA.
	PositionLeftUnit string `json:"positionLeftUnit,omitempty"`

	// PositionTopUnit: Offset top unit for an asset. This is a read-only
	// field if the asset displayType is ASSET_DISPLAY_TYPE_OVERLAY.
	// Applicable to the following creative types: all RICH_MEDIA.
	PositionTopUnit string `json:"positionTopUnit,omitempty"`

	// ProgressiveServingUrl: Progressive URL for video asset. This is a
	// read-only field. Applicable to the following creative types:
	// INSTREAM_VIDEO and all VPAID.
	ProgressiveServingUrl string `json:"progressiveServingUrl,omitempty"`

	// Pushdown: Whether the asset pushes down other content. Applicable to
	// the following creative types: all RICH_MEDIA. Additionally, only
	// applicable when the asset offsets are 0, the collapsedSize.width
	// matches size.width, and the collapsedSize.height is less than
	// size.height.
	Pushdown bool `json:"pushdown,omitempty"`

	// PushdownDuration: Pushdown duration in seconds for an asset. Must be
	// between 0 and 9.99. Applicable to the following creative types: all
	// RICH_MEDIA.Additionally, only applicable when the asset pushdown
	// field is true, the offsets are 0, the collapsedSize.width matches
	// size.width, and the collapsedSize.height is less than size.height.
	PushdownDuration float64 `json:"pushdownDuration,omitempty"`

	// Role: Role of the asset in relation to creative. Applicable to all
	// but the following creative types: all REDIRECT and TRACKING_TEXT.
	// This is a required field.
	// PRIMARY applies to ENHANCED_BANNER,
	// FLASH_INPAGE, HTML5_BANNER, IMAGE, IMAGE_GALLERY, all RICH_MEDIA
	// (which may contain multiple primary assets), and all VPAID
	// creatives.
	// BACKUP_IMAGE applies to ENHANCED_BANNER, FLASH_INPAGE,
	// HTML5_BANNER, all RICH_MEDIA, and all VPAID
	// creatives.
	// ADDITIONAL_IMAGE and ADDITIONAL_FLASH apply to
	// FLASH_INPAGE creatives.
	// OTHER refers to assets from sources other
	// than DCM, such as Studio uploaded assets, applicable to all
	// RICH_MEDIA and all VPAID creatives.
	// PARENT_VIDEO refers to videos
	// uploaded by the user in DCM and is applicable to INSTREAM_VIDEO and
	// VPAID_LINEAR creatives.
	// TRANSCODED_VIDEO refers to videos transcoded
	// by DCM from PARENT_VIDEO assets and is applicable to INSTREAM_VIDEO
	// and VPAID_LINEAR creatives.
	// ALTERNATE_VIDEO refers to the DCM
	// representation of child asset videos from Studio, and is applicable
	// to VPAID_LINEAR creatives. These cannot be added or removed within
	// DCM.
	// For VPAID_LINEAR creatives, PARENT_VIDEO, TRANSCODED_VIDEO and
	// ALTERNATE_VIDEO assets that are marked active serve as backup in case
	// the VPAID creative cannot be served. Only PARENT_VIDEO assets can be
	// added or removed for an INSTREAM_VIDEO or VPAID_LINEAR creative.
	Role string `json:"role,omitempty"`

	// Size: Size associated with this creative asset. This is a required
	// field when applicable; however for IMAGE and FLASH_INPAGE creatives,
	// if left blank, this field will be automatically set using the actual
	// size of the associated image asset. Applicable to the following
	// creative types: ENHANCED_BANNER, ENHANCED_IMAGE, FLASH_INPAGE,
	// HTML5_BANNER, IMAGE, and all RICH_MEDIA.
	Size *Size `json:"size,omitempty"`

	// SslCompliant: Whether the asset is SSL-compliant. This is a read-only
	// field. Applicable to all but the following creative types: all
	// REDIRECT and TRACKING_TEXT.
	SslCompliant bool `json:"sslCompliant,omitempty"`

	// StartTimeType: Initial wait time type before making the asset
	// visible. Applicable to the following creative types: all RICH_MEDIA.
	StartTimeType string `json:"startTimeType,omitempty"`

	// StreamingServingUrl: Streaming URL for video asset. This is a
	// read-only field. Applicable to the following creative types:
	// INSTREAM_VIDEO and all VPAID.
	StreamingServingUrl string `json:"streamingServingUrl,omitempty"`

	// Transparency: Whether the asset is transparent. Applicable to the
	// following creative types: all RICH_MEDIA. Additionally, only
	// applicable to HTML5 assets.
	Transparency bool `json:"transparency,omitempty"`

	// VerticallyLocked: Whether the asset is vertically locked. This is a
	// read-only field. Applicable to the following creative types: all
	// RICH_MEDIA.
	VerticallyLocked bool `json:"verticallyLocked,omitempty"`

	// VideoDuration: Detected video duration for video asset. This is a
	// read-only field. Applicable to the following creative types:
	// INSTREAM_VIDEO and all VPAID.
	VideoDuration float64 `json:"videoDuration,omitempty"`

	// WindowMode: Window mode options for flash assets. Applicable to the
	// following creative types: FLASH_INPAGE, RICH_MEDIA_EXPANDING,
	// RICH_MEDIA_IM_EXPAND, RICH_MEDIA_INPAGE, and
	// RICH_MEDIA_INPAGE_FLOATING.
	WindowMode string `json:"windowMode,omitempty"`

	// ZIndex: zIndex value of an asset. This is a read-only field.
	// Applicable to the following creative types: all
	// RICH_MEDIA.Additionally, only applicable to assets whose displayType
	// is NOT one of the following types: ASSET_DISPLAY_TYPE_INPAGE or
	// ASSET_DISPLAY_TYPE_OVERLAY.
	ZIndex int64 `json:"zIndex,omitempty"`

	// ZipFilename: File name of zip file. This is a read-only field.
	// Applicable to the following creative types: HTML5_BANNER.
	ZipFilename string `json:"zipFilename,omitempty"`

	// ZipFilesize: Size of zip file. This is a read-only field. Applicable
	// to the following creative types: HTML5_BANNER.
	ZipFilesize string `json:"zipFilesize,omitempty"`
}

type CreativeAssetId struct {
	// Name: Name of the creative asset. This is a required field while
	// inserting an asset. After insertion, this assetIdentifier is used to
	// identify the uploaded asset. Characters in the name must be
	// alphanumeric or one of the following: ".-_ ". Spaces are allowed.
	Name string `json:"name,omitempty"`

	// Type: Type of asset to upload. This is a required field. IMAGE is
	// solely used for IMAGE creatives. Other image assets should use
	// HTML_IMAGE.
	Type string `json:"type,omitempty"`
}

type CreativeAssetMetadata struct {
	// AssetIdentifier: ID of the creative asset. This is a required field.
	AssetIdentifier *CreativeAssetId `json:"assetIdentifier,omitempty"`

	// ClickTags: List of detected click tags for assets. This is a
	// read-only auto-generated field.
	ClickTags []*ClickTag `json:"clickTags,omitempty"`

	// DetectedFeatures: List of feature dependencies for the creative asset
	// that are detected by DCM. Feature dependencies are features that a
	// browser must be able to support in order to render your HTML5
	// creative correctly. This is a read-only, auto-generated field.
	DetectedFeatures []string `json:"detectedFeatures,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeAssetMetadata".
	Kind string `json:"kind,omitempty"`

	// WarnedValidationRules: Rules validated during code generation that
	// generated a warning. This is a read-only, auto-generated
	// field.
	//
	// Possible values are:
	// - "CLICK_TAG_NON_TOP_LEVEL"
	// -
	// "CLICK_TAG_MISSING"
	// - "CLICK_TAG_MORE_THAN_ONE"
	// -
	// "CLICK_TAG_INVALID"
	// - "ORPHANED_ASSET"
	// - "PRIMARY_HTML_MISSING"
	// -
	// "EXTERNAL_FILE_REFERENCED"
	// - "MRAID_REFERENCED"
	// -
	// "ADMOB_REFERENCED"
	// - "FILE_TYPE_INVALID"
	// - "ZIP_INVALID"
	// -
	// "LINKED_FILE_NOT_FOUND"
	// - "MAX_FLASH_VERSION_11"
	// -
	// "NOT_SSL_COMPLIANT"
	// - "FILE_DETAIL_EMPTY"
	// - "ASSET_INVALID"
	// -
	// "GWD_PROPERTIES_INVALID"
	// - "ENABLER_UNSUPPORTED_METHOD_DCM"
	// -
	// "ASSET_FORMAT_UNSUPPORTED_DCM"
	// - "COMPONENT_UNSUPPORTED_DCM"
	// -
	// "HTML5_FEATURE_UNSUPPORTED' "
	WarnedValidationRules []string `json:"warnedValidationRules,omitempty"`
}

type CreativeAssignment struct {
	// Active: Whether this creative assignment is active. When true, the
	// creative will be included in the ad's rotation.
	Active bool `json:"active,omitempty"`

	// ApplyEventTags: Whether applicable event tags should fire when this
	// creative assignment is rendered. If this value is unset when the ad
	// is inserted or updated, it will default to true for all creative
	// types EXCEPT for INTERNAL_REDIRECT, INTERSTITIAL_INTERNAL_REDIRECT,
	// and INSTREAM_VIDEO.
	ApplyEventTags bool `json:"applyEventTags,omitempty"`

	// ClickThroughUrl: Click-through URL of the creative assignment.
	ClickThroughUrl *ClickThroughUrl `json:"clickThroughUrl,omitempty"`

	// CompanionCreativeOverrides: Companion creative overrides for this
	// creative assignment. Applicable to video ads.
	CompanionCreativeOverrides []*CompanionClickThroughOverride `json:"companionCreativeOverrides,omitempty"`

	// CreativeGroupAssignments: Creative group assignments for this
	// creative assignment. Only one assignment per creative group number is
	// allowed for a maximum of two assignments.
	CreativeGroupAssignments []*CreativeGroupAssignment `json:"creativeGroupAssignments,omitempty"`

	// CreativeId: ID of the creative to be assigned. This is a required
	// field.
	CreativeId int64 `json:"creativeId,omitempty,string"`

	// CreativeIdDimensionValue: Dimension value for the ID of the creative.
	// This is a read-only, auto-generated field.
	CreativeIdDimensionValue *DimensionValue `json:"creativeIdDimensionValue,omitempty"`

	// EndTime: Date and time that the assigned creative should stop
	// serving. Must be later than the start time.
	EndTime string `json:"endTime,omitempty"`

	// RichMediaExitOverrides: Rich media exit overrides for this creative
	// assignment.
	// Applicable when the creative type is any of the
	// following:
	// - RICH_MEDIA_INPAGE
	// - RICH_MEDIA_INPAGE_FLOATING
	// -
	// RICH_MEDIA_IM_EXPAND
	// - RICH_MEDIA_EXPANDING
	// -
	// RICH_MEDIA_INTERSTITIAL_FLOAT
	// - RICH_MEDIA_MOBILE_IN_APP
	// -
	// RICH_MEDIA_MULTI_FLOATING
	// - RICH_MEDIA_PEEL_DOWN
	// - ADVANCED_BANNER
	// -
	// VPAID_LINEAR
	// - VPAID_NON_LINEAR
	RichMediaExitOverrides []*RichMediaExitOverride `json:"richMediaExitOverrides,omitempty"`

	// Sequence: Sequence number of the creative assignment, applicable when
	// the rotation type is CREATIVE_ROTATION_TYPE_SEQUENTIAL.
	Sequence int64 `json:"sequence,omitempty"`

	// SslCompliant: Whether the creative to be assigned is SSL-compliant.
	// This is a read-only field that is auto-generated when the ad is
	// inserted or updated.
	SslCompliant bool `json:"sslCompliant,omitempty"`

	// StartTime: Date and time that the assigned creative should start
	// serving.
	StartTime string `json:"startTime,omitempty"`

	// Weight: Weight of the creative assignment, applicable when the
	// rotation type is CREATIVE_ROTATION_TYPE_RANDOM.
	Weight int64 `json:"weight,omitempty"`
}

type CreativeCustomEvent struct {
	// Active: Whether the event is active.
	Active bool `json:"active,omitempty"`

	// AdvertiserCustomEventName: User-entered name for the event.
	AdvertiserCustomEventName string `json:"advertiserCustomEventName,omitempty"`

	// AdvertiserCustomEventType: Type of the event. This is a read-only
	// field.
	AdvertiserCustomEventType string `json:"advertiserCustomEventType,omitempty"`

	// ArtworkLabel: Artwork label column, used to link events in DCM back
	// to events in Studio. This is a required field and should not be
	// modified after insertion.
	ArtworkLabel string `json:"artworkLabel,omitempty"`

	// ArtworkType: Artwork type used by the creative.This is a read-only
	// field.
	ArtworkType string `json:"artworkType,omitempty"`

	// ExitUrl: Exit URL of the event. This field is used only for exit
	// events.
	ExitUrl string `json:"exitUrl,omitempty"`

	// Id: ID of this event. This is a required field and should not be
	// modified after insertion.
	Id int64 `json:"id,omitempty,string"`

	// PopupWindowProperties: Properties for rich media popup windows. This
	// field is used only for exit events.
	PopupWindowProperties *PopupWindowProperties `json:"popupWindowProperties,omitempty"`

	// TargetType: Target type used by the event.
	TargetType string `json:"targetType,omitempty"`

	// VideoReportingId: Reporting ID, used to differentiate multiple videos
	// in a single creative.
	VideoReportingId string `json:"videoReportingId,omitempty"`
}

type CreativeField struct {
	// AccountId: Account ID of this creative field. This is a read-only
	// field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this creative field. This is a
	// required field on insertion.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// Id: ID of this creative field. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeField".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this creative field. This is a required field and must
	// be less than 256 characters long and unique among creative fields of
	// the same advertiser.
	Name string `json:"name,omitempty"`

	// SubaccountId: Subaccount ID of this creative field. This is a
	// read-only field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`
}

type CreativeFieldAssignment struct {
	// CreativeFieldId: ID of the creative field.
	CreativeFieldId int64 `json:"creativeFieldId,omitempty,string"`

	// CreativeFieldValueId: ID of the creative field value.
	CreativeFieldValueId int64 `json:"creativeFieldValueId,omitempty,string"`
}

type CreativeFieldValue struct {
	// Id: ID of this creative field value. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeFieldValue".
	Kind string `json:"kind,omitempty"`

	// Value: Value of this creative field value. It needs to be less than
	// 256 characters in length and unique per creative field.
	Value string `json:"value,omitempty"`
}

type CreativeFieldValuesListResponse struct {
	// CreativeFieldValues: Creative field value collection
	CreativeFieldValues []*CreativeFieldValue `json:"creativeFieldValues,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeFieldValuesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type CreativeFieldsListResponse struct {
	// CreativeFields: Creative field collection
	CreativeFields []*CreativeField `json:"creativeFields,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeFieldsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type CreativeGroup struct {
	// AccountId: Account ID of this creative group. This is a read-only
	// field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this creative group. This is a
	// required field on insertion.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// GroupNumber: Subgroup of the creative group. Assign your creative
	// groups to one of the following subgroups in order to filter or manage
	// them more easily. This field is required on insertion and is
	// read-only after insertion.
	// Acceptable values are:
	// - 1
	// - 2
	GroupNumber int64 `json:"groupNumber,omitempty"`

	// Id: ID of this creative group. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeGroup".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this creative group. This is a required field and must
	// be less than 256 characters long and unique among creative groups of
	// the same advertiser.
	Name string `json:"name,omitempty"`

	// SubaccountId: Subaccount ID of this creative group. This is a
	// read-only field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`
}

type CreativeGroupAssignment struct {
	// CreativeGroupId: ID of the creative group to be assigned.
	CreativeGroupId int64 `json:"creativeGroupId,omitempty,string"`

	// CreativeGroupNumber: Creative group number of the creative group
	// assignment.
	CreativeGroupNumber string `json:"creativeGroupNumber,omitempty"`
}

type CreativeGroupsListResponse struct {
	// CreativeGroups: Creative group collection
	CreativeGroups []*CreativeGroup `json:"creativeGroups,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativeGroupsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type CreativeOptimizationConfiguration struct {
	// Id: ID of this creative optimization config. This field is
	// auto-generated when the campaign is inserted or updated. It can be
	// null for existing campaigns.
	Id int64 `json:"id,omitempty,string"`

	// Name: Name of this creative optimization config. This is a required
	// field and must be less than 129 characters long.
	Name string `json:"name,omitempty"`

	// OptimizationActivitys: List of optimization activities associated
	// with this configuration.
	OptimizationActivitys []*OptimizationActivity `json:"optimizationActivitys,omitempty"`

	// OptimizationModel: Optimization model for this configuration.
	OptimizationModel string `json:"optimizationModel,omitempty"`
}

type CreativeRotation struct {
	// CreativeAssignments: Creative assignments in this creative rotation.
	CreativeAssignments []*CreativeAssignment `json:"creativeAssignments,omitempty"`

	// CreativeOptimizationConfigurationId: Creative optimization
	// configuration that is used by this ad. It should refer to one of the
	// existing optimization configurations in the ad's campaign. If it is
	// unset or set to 0, then the campaign's default optimization
	// configuration will be used for this ad.
	CreativeOptimizationConfigurationId int64 `json:"creativeOptimizationConfigurationId,omitempty,string"`

	// Type: Type of creative rotation. Can be used to specify whether to
	// use sequential or random rotation.
	Type string `json:"type,omitempty"`

	// WeightCalculationStrategy: Strategy for calculating weights. Used
	// with CREATIVE_ROTATION_TYPE_RANDOM.
	WeightCalculationStrategy string `json:"weightCalculationStrategy,omitempty"`
}

type CreativeSettings struct {
	// IFrameFooter: Header text for iFrames for this site. Must be less
	// than or equal to 2000 characters long.
	IFrameFooter string `json:"iFrameFooter,omitempty"`

	// IFrameHeader: Header text for iFrames for this site. Must be less
	// than or equal to 2000 characters long.
	IFrameHeader string `json:"iFrameHeader,omitempty"`
}

type CreativesListResponse struct {
	// Creatives: Creative collection
	Creatives []*Creative `json:"creatives,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#creativesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type CrossDimensionReachReportCompatibleFields struct {
	// Breakdown: Dimensions which are compatible to be selected in the
	// "breakdown" section of the report.
	Breakdown []*Dimension `json:"breakdown,omitempty"`

	// DimensionFilters: Dimensions which are compatible to be selected in
	// the "dimensionFilters" section of the report.
	DimensionFilters []*Dimension `json:"dimensionFilters,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#crossDimensionReachReportCompatibleFields.
	Kind string `json:"kind,omitempty"`

	// Metrics: Metrics which are compatible to be selected in the
	// "metricNames" section of the report.
	Metrics []*Metric `json:"metrics,omitempty"`

	// OverlapMetrics: Metrics which are compatible to be selected in the
	// "overlapMetricNames" section of the report.
	OverlapMetrics []*Metric `json:"overlapMetrics,omitempty"`
}

type CustomRichMediaEvents struct {
	// FilteredEventIds: List of custom rich media event IDs. Dimension
	// values must be all of type dfa:richMediaEventTypeIdAndName.
	FilteredEventIds []*DimensionValue `json:"filteredEventIds,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#customRichMediaEvents.
	Kind string `json:"kind,omitempty"`
}

type DateRange struct {
	// EndDate: The end date of the date range, inclusive. A string of the
	// format: "yyyy-MM-dd".
	EndDate string `json:"endDate,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#dateRange.
	Kind string `json:"kind,omitempty"`

	// RelativeDateRange: The date range relative to the date of when the
	// report is run.
	RelativeDateRange string `json:"relativeDateRange,omitempty"`

	// StartDate: The start date of the date range, inclusive. A string of
	// the format: "yyyy-MM-dd".
	StartDate string `json:"startDate,omitempty"`
}

type DayPartTargeting struct {
	// DaysOfWeek: Days of the week when the ad will serve.
	//
	// Acceptable
	// values are:
	// - "SUNDAY"
	// - "MONDAY"
	// - "TUESDAY"
	// - "WEDNESDAY"
	// -
	// "THURSDAY"
	// - "FRIDAY"
	// - "SATURDAY"
	DaysOfWeek []string `json:"daysOfWeek,omitempty"`

	// HoursOfDay: Hours of the day when the ad will serve. Must be an
	// integer between 0 and 23 (inclusive), where 0 is midnight to 1 AM,
	// and 23 is 11 PM to midnight. Can be specified with days of week, in
	// which case the ad would serve during these hours on the specified
	// days. For example, if Monday, Wednesday, Friday are the days of week
	// specified and 9-10am, 3-5pm (hours 9, 15, and 16) is specified, the
	// ad would serve Monday, Wednesdays, and Fridays at 9-10am and 3-5pm.
	HoursOfDay []int64 `json:"hoursOfDay,omitempty"`

	// UserLocalTime: Whether or not to use the user's local time. If false,
	// the America/New York time zone applies.
	UserLocalTime bool `json:"userLocalTime,omitempty"`
}

type DefaultClickThroughEventTagProperties struct {
	// DefaultClickThroughEventTagId: ID of the click-through event tag to
	// apply to all ads in this entity's scope.
	DefaultClickThroughEventTagId int64 `json:"defaultClickThroughEventTagId,omitempty,string"`

	// OverrideInheritedEventTag: Whether this entity should override the
	// inherited default click-through event tag with its own defined value.
	OverrideInheritedEventTag bool `json:"overrideInheritedEventTag,omitempty"`
}

type DeliverySchedule struct {
	// FrequencyCap: Limit on the number of times an individual user can be
	// served the ad within a specified period of time.
	FrequencyCap *FrequencyCap `json:"frequencyCap,omitempty"`

	// HardCutoff: Whether or not hard cutoff is enabled. If true, the ad
	// will not serve after the end date and time. Otherwise the ad will
	// continue to be served until it has reached its delivery goals.
	HardCutoff bool `json:"hardCutoff,omitempty"`

	// ImpressionRatio: Impression ratio for this ad. This ratio determines
	// how often each ad is served relative to the others. For example, if
	// ad A has an impression ratio of 1 and ad B has an impression ratio of
	// 3, then DCM will serve ad B three times as often as ad A. Must be
	// between 1 and 10.
	ImpressionRatio int64 `json:"impressionRatio,omitempty,string"`

	// Priority: Serving priority of an ad, with respect to other ads. The
	// lower the priority number, the greater the priority with which it is
	// served.
	Priority string `json:"priority,omitempty"`
}

type DfpSettings struct {
	// Dfp_network_code: DFP network code for this directory site.
	Dfp_network_code string `json:"dfp_network_code,omitempty"`

	// Dfp_network_name: DFP network name for this directory site.
	Dfp_network_name string `json:"dfp_network_name,omitempty"`

	// ProgrammaticPlacementAccepted: Whether this directory site accepts
	// programmatic placements.
	ProgrammaticPlacementAccepted bool `json:"programmaticPlacementAccepted,omitempty"`

	// PubPaidPlacementAccepted: Whether this directory site accepts
	// publisher-paid tags.
	PubPaidPlacementAccepted bool `json:"pubPaidPlacementAccepted,omitempty"`

	// PublisherPortalOnly: Whether this directory site is available only
	// via Publisher Portal.
	PublisherPortalOnly bool `json:"publisherPortalOnly,omitempty"`
}

type Dimension struct {
	// Kind: The kind of resource this is, in this case
	// dfareporting#dimension.
	Kind string `json:"kind,omitempty"`

	// Name: The dimension name, e.g. dfa:advertiser
	Name string `json:"name,omitempty"`
}

type DimensionFilter struct {
	// DimensionName: The name of the dimension to filter.
	DimensionName string `json:"dimensionName,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#dimensionFilter.
	Kind string `json:"kind,omitempty"`

	// Value: The value of the dimension to filter.
	Value string `json:"value,omitempty"`
}

type DimensionValue struct {
	// DimensionName: The name of the dimension.
	DimensionName string `json:"dimensionName,omitempty"`

	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Id: The ID associated with the value if available.
	Id string `json:"id,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#dimensionValue.
	Kind string `json:"kind,omitempty"`

	// MatchType: Determines how the 'value' field is matched when
	// filtering. If not specified, defaults to EXACT. If set to
	// WILDCARD_EXPRESSION, '*' is allowed as a placeholder for variable
	// length character sequences, and it can be escaped with a backslash.
	// Note, only paid search dimensions ('dfa:paidSearch*') allow a
	// matchType other than EXACT.
	MatchType string `json:"matchType,omitempty"`

	// Value: The value of the dimension.
	Value string `json:"value,omitempty"`
}

type DimensionValueList struct {
	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The dimension values returned in this response.
	Items []*DimensionValue `json:"items,omitempty"`

	// Kind: The kind of list this is, in this case
	// dfareporting#dimensionValueList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through dimension
	// values. To retrieve the next page of results, set the next request's
	// "pageToken" to the value of this field. The page token is only valid
	// for a limited amount of time and should not be persisted.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type DimensionValueRequest struct {
	// DimensionName: The name of the dimension for which values should be
	// requested.
	DimensionName string `json:"dimensionName,omitempty"`

	// EndDate: The end date of the date range for which to retrieve
	// dimension values. A string of the format "yyyy-MM-dd".
	EndDate string `json:"endDate,omitempty"`

	// Filters: The list of filters by which to filter values. The filters
	// are ANDed.
	Filters []*DimensionFilter `json:"filters,omitempty"`

	// Kind: The kind of request this is, in this case
	// dfareporting#dimensionValueRequest.
	Kind string `json:"kind,omitempty"`

	// StartDate: The start date of the date range for which to retrieve
	// dimension values. A string of the format "yyyy-MM-dd".
	StartDate string `json:"startDate,omitempty"`
}

type DirectorySite struct {
	// Active: Whether this directory site is active.
	Active bool `json:"active,omitempty"`

	// ContactAssignments: Directory site contacts.
	ContactAssignments []*DirectorySiteContactAssignment `json:"contactAssignments,omitempty"`

	// CountryId: Country ID of this directory site.
	CountryId int64 `json:"countryId,omitempty,string"`

	// CurrencyId: Currency ID of this directory site.
	// Possible values are:
	//
	// - "1" for USD
	// - "2" for GBP
	// - "3" for ESP
	// - "4" for SEK
	// - "5"
	// for CAD
	// - "6" for JPY
	// - "7" for DEM
	// - "8" for AUD
	// - "9" for FRF
	//
	// - "10" for ITL
	// - "11" for DKK
	// - "12" for NOK
	// - "13" for FIM
	// -
	// "14" for ZAR
	// - "15" for IEP
	// - "16" for NLG
	// - "17" for EUR
	// - "18"
	// for KRW
	// - "19" for TWD
	// - "20" for SGD
	// - "21" for CNY
	// - "22" for
	// HKD
	// - "23" for NZD
	// - "24" for MYR
	// - "25" for BRL
	// - "26" for PTE
	//
	// - "27" for MXP
	// - "28" for CLP
	// - "29" for TRY
	// - "30" for ARS
	// -
	// "31" for PEN
	// - "32" for ILS
	// - "33" for CHF
	// - "34" for VEF
	// - "35"
	// for COP
	// - "36" for GTQ
	CurrencyId int64 `json:"currencyId,omitempty,string"`

	// Description: Description of this directory site.
	Description string `json:"description,omitempty"`

	// Id: ID of this directory site. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this directory site.
	// This is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// InpageTagFormats: Tag types for regular placements.
	//
	// Acceptable
	// values are:
	// - "STANDARD"
	// - "IFRAME_JAVASCRIPT_INPAGE"
	// -
	// "INTERNAL_REDIRECT_INPAGE"
	// - "JAVASCRIPT_INPAGE"
	InpageTagFormats []string `json:"inpageTagFormats,omitempty"`

	// InterstitialTagFormats: Tag types for interstitial
	// placements.
	//
	// Acceptable values are:
	// -
	// "IFRAME_JAVASCRIPT_INTERSTITIAL"
	// - "INTERNAL_REDIRECT_INTERSTITIAL"
	// -
	// "JAVASCRIPT_INTERSTITIAL"
	InterstitialTagFormats []string `json:"interstitialTagFormats,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#directorySite".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this directory site.
	Name string `json:"name,omitempty"`

	// ParentId: Parent directory site ID.
	ParentId int64 `json:"parentId,omitempty,string"`

	// Settings: Directory site settings.
	Settings *DirectorySiteSettings `json:"settings,omitempty"`

	// Url: URL of this directory site.
	Url string `json:"url,omitempty"`
}

type DirectorySiteContact struct {
	// Email: Email address of this directory site contact.
	Email string `json:"email,omitempty"`

	// FirstName: First name of this directory site contact.
	FirstName string `json:"firstName,omitempty"`

	// Id: ID of this directory site contact. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#directorySiteContact".
	Kind string `json:"kind,omitempty"`

	// LastName: Last name of this directory site contact.
	LastName string `json:"lastName,omitempty"`

	// Role: Directory site contact role.
	Role string `json:"role,omitempty"`

	// Type: Directory site contact type.
	Type string `json:"type,omitempty"`
}

type DirectorySiteContactAssignment struct {
	// ContactId: ID of this directory site contact. This is a read-only,
	// auto-generated field.
	ContactId int64 `json:"contactId,omitempty,string"`

	// Visibility: Visibility of this directory site contact assignment.
	// When set to PUBLIC this contact assignment is visible to all account
	// and agency users; when set to PRIVATE it is visible only to the site.
	Visibility string `json:"visibility,omitempty"`
}

type DirectorySiteContactsListResponse struct {
	// DirectorySiteContacts: Directory site contact collection
	DirectorySiteContacts []*DirectorySiteContact `json:"directorySiteContacts,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#directorySiteContactsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type DirectorySiteSettings struct {
	// ActiveViewOptOut: Whether this directory site has disabled active
	// view creatives.
	ActiveViewOptOut bool `json:"activeViewOptOut,omitempty"`

	// Dfp_settings: Directory site DFP settings.
	Dfp_settings *DfpSettings `json:"dfp_settings,omitempty"`

	// Instream_video_placement_accepted: Whether this site accepts
	// in-stream video ads.
	Instream_video_placement_accepted bool `json:"instream_video_placement_accepted,omitempty"`

	// InterstitialPlacementAccepted: Whether this site accepts interstitial
	// ads.
	InterstitialPlacementAccepted bool `json:"interstitialPlacementAccepted,omitempty"`

	// NielsenOcrOptOut: Whether this directory site has disabled Nielsen
	// OCR reach ratings.
	NielsenOcrOptOut bool `json:"nielsenOcrOptOut,omitempty"`

	// VerificationTagOptOut: Whether this directory site has disabled
	// generation of Verification tags.
	VerificationTagOptOut bool `json:"verificationTagOptOut,omitempty"`

	// VideoActiveViewOptOut: Whether this directory site has disabled
	// active view for in-stream video creatives.
	VideoActiveViewOptOut bool `json:"videoActiveViewOptOut,omitempty"`
}

type DirectorySitesListResponse struct {
	// DirectorySites: Directory site collection
	DirectorySites []*DirectorySite `json:"directorySites,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#directorySitesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type EventTag struct {
	// AccountId: Account ID of this event tag. This is a read-only field
	// that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this event tag. This field or the
	// campaignId field is required on insertion.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// CampaignId: Campaign ID of this event tag. This field or the
	// advertiserId field is required on insertion.
	CampaignId int64 `json:"campaignId,omitempty,string"`

	// CampaignIdDimensionValue: Dimension value for the ID of the campaign.
	// This is a read-only, auto-generated field.
	CampaignIdDimensionValue *DimensionValue `json:"campaignIdDimensionValue,omitempty"`

	// EnabledByDefault: Whether this event tag should be automatically
	// enabled for all of the advertiser's campaigns and ads.
	EnabledByDefault bool `json:"enabledByDefault,omitempty"`

	// Id: ID of this event tag. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#eventTag".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this event tag. This is a required field and must be
	// less than 256 characters long.
	Name string `json:"name,omitempty"`

	// SiteFilterType: Site filter type for this event tag. If no type is
	// specified then the event tag will be applied to all sites.
	SiteFilterType string `json:"siteFilterType,omitempty"`

	// SiteIds: Filter list of site IDs associated with this event tag. The
	// siteFilterType determines whether this is a whitelist or blacklist
	// filter.
	SiteIds googleapi.Int64s `json:"siteIds,omitempty"`

	// SslCompliant: Whether this tag is SSL-compliant or not.
	SslCompliant bool `json:"sslCompliant,omitempty"`

	// Status: Status of this event tag. Must be ENABLED for this event tag
	// to fire. This is a required field.
	Status string `json:"status,omitempty"`

	// SubaccountId: Subaccount ID of this event tag. This is a read-only
	// field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// Type: Event tag type. Can be used to specify whether to use a
	// third-party pixel, a third-party JavaScript URL, or a third-party
	// click-through URL for either impression or click tracking. This is a
	// required field.
	Type string `json:"type,omitempty"`

	// Url: Payload URL for this event tag. The URL on a click-through event
	// tag should have a landing page URL appended to the end of it. This
	// field is required on insertion.
	Url string `json:"url,omitempty"`

	// UrlEscapeLevels: Number of times the landing page URL should be
	// URL-escaped before being appended to the click-through event tag URL.
	// Only applies to click-through event tags as specified by the event
	// tag type.
	UrlEscapeLevels int64 `json:"urlEscapeLevels,omitempty"`
}

type EventTagOverride struct {
	// Enabled: Whether this override is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// Id: ID of this event tag override. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`
}

type EventTagsListResponse struct {
	// EventTags: Event tag collection
	EventTags []*EventTag `json:"eventTags,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#eventTagsListResponse".
	Kind string `json:"kind,omitempty"`
}

type File struct {
	// DateRange: The date range for which the file has report data. The
	// date range will always be the absolute date range for which the
	// report is run.
	DateRange *DateRange `json:"dateRange,omitempty"`

	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// FileName: The filename of the file.
	FileName string `json:"fileName,omitempty"`

	// Format: The output format of the report. Only available once the file
	// is available.
	Format string `json:"format,omitempty"`

	// Id: The unique ID of this report file.
	Id int64 `json:"id,omitempty,string"`

	// Kind: The kind of resource this is, in this case dfareporting#file.
	Kind string `json:"kind,omitempty"`

	// LastModifiedTime: The timestamp in milliseconds since epoch when this
	// file was last modified.
	LastModifiedTime int64 `json:"lastModifiedTime,omitempty,string"`

	// ReportId: The ID of the report this file was generated from.
	ReportId int64 `json:"reportId,omitempty,string"`

	// Status: The status of the report file.
	Status string `json:"status,omitempty"`

	// Urls: The URLs where the completed report file can be downloaded.
	Urls *FileUrls `json:"urls,omitempty"`
}

type FileUrls struct {
	// ApiUrl: The URL for downloading the report data through the API.
	ApiUrl string `json:"apiUrl,omitempty"`

	// BrowserUrl: The URL for downloading the report data through a
	// browser.
	BrowserUrl string `json:"browserUrl,omitempty"`
}

type FileList struct {
	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The files returned in this response.
	Items []*File `json:"items,omitempty"`

	// Kind: The kind of list this is, in this case dfareporting#fileList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through files. To
	// retrieve the next page of results, set the next request's "pageToken"
	// to the value of this field. The page token is only valid for a
	// limited amount of time and should not be persisted.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type FloodlightActivitiesGenerateTagResponse struct {
	// FloodlightActivityTag: Generated tag for this floodlight activity.
	FloodlightActivityTag string `json:"floodlightActivityTag,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightActivitiesGenerateTagResponse".
	Kind string `json:"kind,omitempty"`
}

type FloodlightActivitiesListResponse struct {
	// FloodlightActivities: Floodlight activity collection
	FloodlightActivities []*FloodlightActivity `json:"floodlightActivities,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightActivitiesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type FloodlightActivity struct {
	// AccountId: Account ID of this floodlight activity. This is a
	// read-only field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this floodlight activity. If this
	// field is left blank, the value will be copied over either from the
	// activity group's advertiser or the existing activity's advertiser.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// CacheBustingType: Code type used for cache busting in the generated
	// tag.
	CacheBustingType string `json:"cacheBustingType,omitempty"`

	// CountingMethod: Counting method for conversions for this floodlight
	// activity. This is a required field.
	CountingMethod string `json:"countingMethod,omitempty"`

	// DefaultTags: Dynamic floodlight tags.
	DefaultTags []*FloodlightActivityDynamicTag `json:"defaultTags,omitempty"`

	// ExpectedUrl: URL where this tag will be deployed. If specified, must
	// be less than 256 characters long.
	ExpectedUrl string `json:"expectedUrl,omitempty"`

	// FloodlightActivityGroupId: Floodlight activity group ID of this
	// floodlight activity. This is a required field.
	FloodlightActivityGroupId int64 `json:"floodlightActivityGroupId,omitempty,string"`

	// FloodlightActivityGroupName: Name of the associated floodlight
	// activity group. This is a read-only field.
	FloodlightActivityGroupName string `json:"floodlightActivityGroupName,omitempty"`

	// FloodlightActivityGroupTagString: Tag string of the associated
	// floodlight activity group. This is a read-only field.
	FloodlightActivityGroupTagString string `json:"floodlightActivityGroupTagString,omitempty"`

	// FloodlightActivityGroupType: Type of the associated floodlight
	// activity group. This is a read-only field.
	FloodlightActivityGroupType string `json:"floodlightActivityGroupType,omitempty"`

	// FloodlightConfigurationId: Floodlight configuration ID of this
	// floodlight activity. If this field is left blank, the value will be
	// copied over either from the activity group's floodlight configuration
	// or from the existing activity's floodlight configuration.
	FloodlightConfigurationId int64 `json:"floodlightConfigurationId,omitempty,string"`

	// FloodlightConfigurationIdDimensionValue: Dimension value for the ID
	// of the floodlight configuration. This is a read-only, auto-generated
	// field.
	FloodlightConfigurationIdDimensionValue *DimensionValue `json:"floodlightConfigurationIdDimensionValue,omitempty"`

	// Hidden: Whether this activity is archived.
	Hidden bool `json:"hidden,omitempty"`

	// Id: ID of this floodlight activity. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this floodlight
	// activity. This is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// ImageTagEnabled: Whether the image tag is enabled for this activity.
	ImageTagEnabled bool `json:"imageTagEnabled,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightActivity".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this floodlight activity. This is a required field.
	// Must be less than 129 characters long and cannot contain quotes.
	Name string `json:"name,omitempty"`

	// Notes: General notes or implementation instructions for the tag.
	Notes string `json:"notes,omitempty"`

	// PublisherTags: Publisher dynamic floodlight tags.
	PublisherTags []*FloodlightActivityPublisherDynamicTag `json:"publisherTags,omitempty"`

	// Secure: Whether this tag should use SSL.
	Secure bool `json:"secure,omitempty"`

	// SslCompliant: Whether the floodlight activity is SSL-compliant. This
	// is a read-only field, its value detected by the system from the
	// floodlight tags.
	SslCompliant bool `json:"sslCompliant,omitempty"`

	// SslRequired: Whether this floodlight activity must be SSL-compliant.
	SslRequired bool `json:"sslRequired,omitempty"`

	// SubaccountId: Subaccount ID of this floodlight activity. This is a
	// read-only field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TagFormat: Tag format type for the floodlight activity. If left
	// blank, the tag format will default to HTML.
	TagFormat string `json:"tagFormat,omitempty"`

	// TagString: Value of the cat= paramter in the floodlight tag, which
	// the ad servers use to identify the activity. This is optional: if
	// empty, a new tag string will be generated for you. This string must
	// be 1 to 8 characters long, with valid characters being
	// [a-z][A-Z][0-9][-][ _ ]. This tag string must also be unique among
	// activities of the same activity group. This field is read-only after
	// insertion.
	TagString string `json:"tagString,omitempty"`

	// UserDefinedVariableTypes: List of the user-defined variables used by
	// this conversion tag. These map to the "u[1-20]=" in the tags. Each of
	// these can have a user defined type.
	// Acceptable values are:
	// - "U1"
	// -
	// "U2"
	// - "U3"
	// - "U4"
	// - "U5"
	// - "U6"
	// - "U7"
	// - "U8"
	// - "U9"
	// - "U10"
	// -
	// "U11"
	// - "U12"
	// - "U13"
	// - "U14"
	// - "U15"
	// - "U16"
	// - "U17"
	// - "U18"
	// -
	// "U19"
	// - "U20"
	UserDefinedVariableTypes []string `json:"userDefinedVariableTypes,omitempty"`
}

type FloodlightActivityDynamicTag struct {
	// Id: ID of this dynamic tag. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Name: Name of this tag.
	Name string `json:"name,omitempty"`

	// Tag: Tag code.
	Tag string `json:"tag,omitempty"`
}

type FloodlightActivityGroup struct {
	// AccountId: Account ID of this floodlight activity group. This is a
	// read-only field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this floodlight activity group. If
	// this field is left blank, the value will be copied over either from
	// the floodlight configuration's advertiser or from the existing
	// activity group's advertiser.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// FloodlightConfigurationId: Floodlight configuration ID of this
	// floodlight activity group. This is a required field.
	FloodlightConfigurationId int64 `json:"floodlightConfigurationId,omitempty,string"`

	// FloodlightConfigurationIdDimensionValue: Dimension value for the ID
	// of the floodlight configuration. This is a read-only, auto-generated
	// field.
	FloodlightConfigurationIdDimensionValue *DimensionValue `json:"floodlightConfigurationIdDimensionValue,omitempty"`

	// Id: ID of this floodlight activity group. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this floodlight
	// activity group. This is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightActivityGroup".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this floodlight activity group. This is a required
	// field. Must be less than 65 characters long and cannot contain
	// quotes.
	Name string `json:"name,omitempty"`

	// SubaccountId: Subaccount ID of this floodlight activity group. This
	// is a read-only field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TagString: Value of the type= parameter in the floodlight tag, which
	// the ad servers use to identify the activity group that the activity
	// belongs to. This is optional: if empty, a new tag string will be
	// generated for you. This string must be 1 to 8 characters long, with
	// valid characters being [a-z][A-Z][0-9][-][ _ ]. This tag string must
	// also be unique among activity groups of the same floodlight
	// configuration. This field is read-only after insertion.
	TagString string `json:"tagString,omitempty"`

	// Type: Type of the floodlight activity group. This is a required field
	// that is read-only after insertion.
	Type string `json:"type,omitempty"`
}

type FloodlightActivityGroupsListResponse struct {
	// FloodlightActivityGroups: Floodlight activity group collection
	FloodlightActivityGroups []*FloodlightActivityGroup `json:"floodlightActivityGroups,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightActivityGroupsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type FloodlightActivityPublisherDynamicTag struct {
	// ClickThrough: Whether this tag is applicable only for click-throughs.
	ClickThrough bool `json:"clickThrough,omitempty"`

	// DirectorySiteId: Directory site ID of this dynamic tag. This is a
	// write-only field that can be used as an alternative to the siteId
	// field. When this resource is retrieved, only the siteId field will be
	// populated.
	DirectorySiteId int64 `json:"directorySiteId,omitempty,string"`

	// DynamicTag: Dynamic floodlight tag.
	DynamicTag *FloodlightActivityDynamicTag `json:"dynamicTag,omitempty"`

	// SiteId: Site ID of this dynamic tag.
	SiteId int64 `json:"siteId,omitempty,string"`

	// SiteIdDimensionValue: Dimension value for the ID of the site. This is
	// a read-only, auto-generated field.
	SiteIdDimensionValue *DimensionValue `json:"siteIdDimensionValue,omitempty"`

	// ViewThrough: Whether this tag is applicable only for view-throughs.
	ViewThrough bool `json:"viewThrough,omitempty"`
}

type FloodlightConfiguration struct {
	// AccountId: Account ID of this floodlight configuration. This is a
	// read-only field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of the parent advertiser of this
	// floodlight configuration.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// AnalyticsDataSharingEnabled: Whether advertiser data is shared with
	// Google Analytics.
	AnalyticsDataSharingEnabled bool `json:"analyticsDataSharingEnabled,omitempty"`

	// ExposureToConversionEnabled: Whether the exposure-to-conversion
	// report is enabled. This report shows detailed pathway information on
	// up to 10 of the most recent ad exposures seen by a user before
	// converting.
	ExposureToConversionEnabled bool `json:"exposureToConversionEnabled,omitempty"`

	// FirstDayOfWeek: Day that will be counted as the first day of the week
	// in reports. This is a required field.
	FirstDayOfWeek string `json:"firstDayOfWeek,omitempty"`

	// Id: ID of this floodlight configuration. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this floodlight
	// configuration. This is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightConfiguration".
	Kind string `json:"kind,omitempty"`

	// LookbackConfiguration: Lookback window settings for this floodlight
	// configuration.
	LookbackConfiguration *LookbackConfiguration `json:"lookbackConfiguration,omitempty"`

	// NaturalSearchConversionAttributionOption: Types of attribution
	// options for natural search conversions.
	NaturalSearchConversionAttributionOption string `json:"naturalSearchConversionAttributionOption,omitempty"`

	// OmnitureSettings: Settings for DCM Omniture integration.
	OmnitureSettings *OmnitureSettings `json:"omnitureSettings,omitempty"`

	// SslRequired: Whether floodlight activities owned by this
	// configuration are required to be SSL-compliant.
	SslRequired bool `json:"sslRequired,omitempty"`

	// StandardVariableTypes: List of standard variables enabled for this
	// configuration.
	//
	// Acceptable values are:
	// - "ORD"
	// - "NUM"
	StandardVariableTypes []string `json:"standardVariableTypes,omitempty"`

	// SubaccountId: Subaccount ID of this floodlight configuration. This is
	// a read-only field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TagSettings: Configuration settings for dynamic and image floodlight
	// tags.
	TagSettings *TagSettings `json:"tagSettings,omitempty"`

	// UserDefinedVariableConfigurations: List of user defined variables
	// enabled for this configuration.
	UserDefinedVariableConfigurations []*UserDefinedVariableConfiguration `json:"userDefinedVariableConfigurations,omitempty"`
}

type FloodlightConfigurationsListResponse struct {
	// FloodlightConfigurations: Floodlight configuration collection
	FloodlightConfigurations []*FloodlightConfiguration `json:"floodlightConfigurations,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#floodlightConfigurationsListResponse".
	Kind string `json:"kind,omitempty"`
}

type FloodlightReportCompatibleFields struct {
	// DimensionFilters: Dimensions which are compatible to be selected in
	// the "dimensionFilters" section of the report.
	DimensionFilters []*Dimension `json:"dimensionFilters,omitempty"`

	// Dimensions: Dimensions which are compatible to be selected in the
	// "dimensions" section of the report.
	Dimensions []*Dimension `json:"dimensions,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#floodlightReportCompatibleFields.
	Kind string `json:"kind,omitempty"`

	// Metrics: Metrics which are compatible to be selected in the
	// "metricNames" section of the report.
	Metrics []*Metric `json:"metrics,omitempty"`
}

type FrequencyCap struct {
	// Duration: Duration of time, in seconds, for this frequency cap. The
	// maximum duration is 90 days in seconds, or 7,776,000.
	Duration int64 `json:"duration,omitempty,string"`

	// Impressions: Number of times an individual user can be served the ad
	// within the specified duration. The maximum allowed is 15.
	Impressions int64 `json:"impressions,omitempty,string"`
}

type FsCommand struct {
	// Left: Distance from the left of the browser.Applicable when
	// positionOption is DISTANCE_FROM_TOP_LEFT_CORNER.
	Left int64 `json:"left,omitempty"`

	// PositionOption: Position in the browser where the window will open.
	PositionOption string `json:"positionOption,omitempty"`

	// Top: Distance from the top of the browser. Applicable when
	// positionOption is DISTANCE_FROM_TOP_LEFT_CORNER.
	Top int64 `json:"top,omitempty"`

	// WindowHeight: Height of the window.
	WindowHeight int64 `json:"windowHeight,omitempty"`

	// WindowWidth: Width of the window.
	WindowWidth int64 `json:"windowWidth,omitempty"`
}

type GeoTargeting struct {
	// Cities: Cities to be targeted. For each city only dartId is required.
	// The other fields are populated automatically when the ad is inserted
	// or updated. If targeting a city, do not target or exclude the country
	// of the city, and do not target the metro or region of the city.
	Cities []*City `json:"cities,omitempty"`

	// Countries: Countries to be targeted or excluded from targeting,
	// depending on the setting of the excludeCountries field. For each
	// country only dartId is required. The other fields are populated
	// automatically when the ad is inserted or updated. If targeting or
	// excluding a country, do not target regions, cities, metros, or postal
	// codes in the same country.
	Countries []*Country `json:"countries,omitempty"`

	// ExcludeCountries: Whether or not to exclude the countries in the
	// countries field from targeting. If false, the countries field refers
	// to countries which will be targeted by the ad.
	ExcludeCountries bool `json:"excludeCountries,omitempty"`

	// Metros: Metros to be targeted. For each metro only dmaId is required.
	// The other fields are populated automatically when the ad is inserted
	// or updated. If targeting a metro, do not target or exclude the
	// country of the metro.
	Metros []*Metro `json:"metros,omitempty"`

	// PostalCodes: Postal codes to be targeted. For each postal code only
	// id is required. The other fields are populated automatically when the
	// ad is inserted or updated. If targeting a postal code, do not target
	// or exclude the country of the postal code.
	PostalCodes []*PostalCode `json:"postalCodes,omitempty"`

	// Regions: Regions to be targeted. For each region only dartId is
	// required. The other fields are populated automatically when the ad is
	// inserted or updated. If targeting a region, do not target or exclude
	// the country of the region.
	Regions []*Region `json:"regions,omitempty"`
}

type KeyValueTargetingExpression struct {
	// Expression: Keyword expression being targeted by the ad.
	Expression string `json:"expression,omitempty"`
}

type LandingPage struct {
	// Default: Whether or not this landing page will be assigned to any ads
	// or creatives that do not have a landing page assigned explicitly.
	// Only one default landing page is allowed per campaign.
	Default bool `json:"default,omitempty"`

	// Id: ID of this landing page. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#landingPage".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this landing page. This is a required field. Must be
	// less than 256 characters long, and must be unique among landing pages
	// of the same campaign.
	Name string `json:"name,omitempty"`

	// Url: URL of this landing page. This is a required field.
	Url string `json:"url,omitempty"`
}

type LandingPagesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#landingPagesListResponse".
	Kind string `json:"kind,omitempty"`

	// LandingPages: Landing page collection
	LandingPages []*LandingPage `json:"landingPages,omitempty"`
}

type LastModifiedInfo struct {
	// Time: Timestamp of the last change in milliseconds since epoch.
	Time int64 `json:"time,omitempty,string"`
}

type ListTargetingExpression struct {
	// Expression: Expression describing which lists are being targeted by
	// the ad.
	Expression string `json:"expression,omitempty"`
}

type LookbackConfiguration struct {
	// ClickDuration: Lookback window, in days, from the last time a given
	// user clicked on one of your ads. If you enter 0, clicks will not be
	// considered as triggering events for floodlight tracking. If you leave
	// this field blank, the default value for your account will be used.
	ClickDuration int64 `json:"clickDuration,omitempty"`

	// PostImpressionActivitiesDuration: Lookback window, in days, from the
	// last time a given user viewed one of your ads. If you enter 0,
	// impressions will not be considered as triggering events for
	// floodlight tracking. If you leave this field blank, the default value
	// for your account will be used.
	PostImpressionActivitiesDuration int64 `json:"postImpressionActivitiesDuration,omitempty"`
}

type Metric struct {
	// Kind: The kind of resource this is, in this case dfareporting#metric.
	Kind string `json:"kind,omitempty"`

	// Name: The metric name, e.g. dfa:impressions
	Name string `json:"name,omitempty"`
}

type Metro struct {
	// CountryCode: Country code of the country to which this metro region
	// belongs.
	CountryCode string `json:"countryCode,omitempty"`

	// CountryDartId: DART ID of the country to which this metro region
	// belongs.
	CountryDartId int64 `json:"countryDartId,omitempty,string"`

	// DartId: DART ID of this metro region.
	DartId int64 `json:"dartId,omitempty,string"`

	// DmaId: DMA ID of this metro region. This is the ID used for targeting
	// and generating reports, and is equivalent to metro_code.
	DmaId int64 `json:"dmaId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#metro".
	Kind string `json:"kind,omitempty"`

	// MetroCode: Metro code of this metro region. This is equivalent to
	// dma_id.
	MetroCode string `json:"metroCode,omitempty"`

	// Name: Name of this metro region.
	Name string `json:"name,omitempty"`
}

type MetrosListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#metrosListResponse".
	Kind string `json:"kind,omitempty"`

	// Metros: Metro collection
	Metros []*Metro `json:"metros,omitempty"`
}

type MobileCarrier struct {
	// CountryCode: Country code of the country to which this mobile carrier
	// belongs.
	CountryCode string `json:"countryCode,omitempty"`

	// CountryDartId: DART ID of the country to which this mobile carrier
	// belongs.
	CountryDartId int64 `json:"countryDartId,omitempty,string"`

	// Id: ID of this mobile carrier.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#mobileCarrier".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this mobile carrier.
	Name string `json:"name,omitempty"`
}

type MobileCarriersListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#mobileCarriersListResponse".
	Kind string `json:"kind,omitempty"`

	// MobileCarriers: Mobile carrier collection
	MobileCarriers []*MobileCarrier `json:"mobileCarriers,omitempty"`
}

type ObjectFilter struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#objectFilter".
	Kind string `json:"kind,omitempty"`

	// ObjectIds: Applicable when status is ASSIGNED. The user has access to
	// objects with these object IDs.
	ObjectIds googleapi.Int64s `json:"objectIds,omitempty"`

	// Status: Status of the filter. NONE means the user has access to none
	// of the objects. ALL means the user has access to all objects.
	// ASSIGNED means the user has access to the objects with IDs in the
	// objectIds list.
	Status string `json:"status,omitempty"`
}

type OffsetPosition struct {
	// Left: Offset distance from left side of an asset or a window.
	Left int64 `json:"left,omitempty"`

	// Top: Offset distance from top side of an asset or a window.
	Top int64 `json:"top,omitempty"`
}

type OmnitureSettings struct {
	// OmnitureCostDataEnabled: Whether placement cost data will be sent to
	// Omniture. This property can be enabled only if
	// omnitureIntegrationEnabled is true.
	OmnitureCostDataEnabled bool `json:"omnitureCostDataEnabled,omitempty"`

	// OmnitureIntegrationEnabled: Whether Omniture integration is enabled.
	// This property can be enabled only when the "Advanced Ad Serving"
	// account setting is enabled.
	OmnitureIntegrationEnabled bool `json:"omnitureIntegrationEnabled,omitempty"`
}

type OperatingSystem struct {
	// DartId: DART ID of this operating system. This is the ID used for
	// targeting.
	DartId int64 `json:"dartId,omitempty,string"`

	// Desktop: Whether this operating system is for desktop.
	Desktop bool `json:"desktop,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#operatingSystem".
	Kind string `json:"kind,omitempty"`

	// Mobile: Whether this operating system is for mobile.
	Mobile bool `json:"mobile,omitempty"`

	// Name: Name of this operating system.
	Name string `json:"name,omitempty"`
}

type OperatingSystemVersion struct {
	// Id: ID of this operating system version.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#operatingSystemVersion".
	Kind string `json:"kind,omitempty"`

	// MajorVersion: Major version (leftmost number) of this operating
	// system version.
	MajorVersion string `json:"majorVersion,omitempty"`

	// MinorVersion: Minor version (number after first dot on the left) of
	// this operating system version.
	MinorVersion string `json:"minorVersion,omitempty"`

	// Name: Name of this operating system version.
	Name string `json:"name,omitempty"`

	// OperatingSystem: Operating system of this operating system version.
	OperatingSystem *OperatingSystem `json:"operatingSystem,omitempty"`
}

type OperatingSystemVersionsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#operatingSystemVersionsListResponse".
	Kind string `json:"kind,omitempty"`

	// OperatingSystemVersions: Operating system version collection
	OperatingSystemVersions []*OperatingSystemVersion `json:"operatingSystemVersions,omitempty"`
}

type OperatingSystemsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#operatingSystemsListResponse".
	Kind string `json:"kind,omitempty"`

	// OperatingSystems: Operating system collection
	OperatingSystems []*OperatingSystem `json:"operatingSystems,omitempty"`
}

type OptimizationActivity struct {
	// FloodlightActivityId: Floodlight activity ID of this optimization
	// activity. This is a required field.
	FloodlightActivityId int64 `json:"floodlightActivityId,omitempty,string"`

	// FloodlightActivityIdDimensionValue: Dimension value for the ID of the
	// floodlight activity. This is a read-only, auto-generated field.
	FloodlightActivityIdDimensionValue *DimensionValue `json:"floodlightActivityIdDimensionValue,omitempty"`

	// Weight: Weight associated with this optimization. Must be greater
	// than 1. The weight assigned will be understood in proportion to the
	// weights assigned to the other optimization activities.
	Weight int64 `json:"weight,omitempty"`
}

type PathToConversionReportCompatibleFields struct {
	// ConversionDimensions: Conversion dimensions which are compatible to
	// be selected in the "conversionDimensions" section of the report.
	ConversionDimensions []*Dimension `json:"conversionDimensions,omitempty"`

	// CustomFloodlightVariables: Custom floodlight variables which are
	// compatible to be selected in the "customFloodlightVariables" section
	// of the report.
	CustomFloodlightVariables []*Dimension `json:"customFloodlightVariables,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#pathToConversionReportCompatibleFields.
	Kind string `json:"kind,omitempty"`

	// Metrics: Metrics which are compatible to be selected in the
	// "metricNames" section of the report.
	Metrics []*Metric `json:"metrics,omitempty"`

	// PerInteractionDimensions: Per-interaction dimensions which are
	// compatible to be selected in the "perInteractionDimensions" section
	// of the report.
	PerInteractionDimensions []*Dimension `json:"perInteractionDimensions,omitempty"`
}

type Placement struct {
	// AccountId: Account ID of this placement. This field can be left
	// blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this placement. This field can be left
	// blank.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// Archived: Whether this placement is archived.
	Archived bool `json:"archived,omitempty"`

	// CampaignId: Campaign ID of this placement. This field is a required
	// field on insertion.
	CampaignId int64 `json:"campaignId,omitempty,string"`

	// CampaignIdDimensionValue: Dimension value for the ID of the campaign.
	// This is a read-only, auto-generated field.
	CampaignIdDimensionValue *DimensionValue `json:"campaignIdDimensionValue,omitempty"`

	// Comment: Comments for this placement.
	Comment string `json:"comment,omitempty"`

	// Compatibility: Placement compatibility. WEB and WEB_INTERSTITIAL
	// refer to rendering either on desktop or on mobile devices for regular
	// or interstitial ads, respectively. APP and APP_INTERSTITIAL are for
	// rendering in mobile apps.IN_STREAM_VIDEO refers to rendering in
	// in-stream video ads developed with the VAST standard. This field is
	// required on insertion.
	Compatibility string `json:"compatibility,omitempty"`

	// ContentCategoryId: ID of the content category assigned to this
	// placement.
	ContentCategoryId int64 `json:"contentCategoryId,omitempty,string"`

	// CreateInfo: Information about the creation of this placement. This is
	// a read-only field.
	CreateInfo *LastModifiedInfo `json:"createInfo,omitempty"`

	// DirectorySiteId: Directory site ID of this placement. On insert, you
	// must set either this field or the siteId field to specify the site
	// associated with this placement. This is a required field that is
	// read-only after insertion.
	DirectorySiteId int64 `json:"directorySiteId,omitempty,string"`

	// DirectorySiteIdDimensionValue: Dimension value for the ID of the
	// directory site. This is a read-only, auto-generated field.
	DirectorySiteIdDimensionValue *DimensionValue `json:"directorySiteIdDimensionValue,omitempty"`

	// ExternalId: External ID for this placement.
	ExternalId string `json:"externalId,omitempty"`

	// Id: ID of this placement. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this placement. This
	// is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// KeyName: Key name of this placement. This is a read-only,
	// auto-generated field.
	KeyName string `json:"keyName,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placement".
	Kind string `json:"kind,omitempty"`

	// LastModifiedInfo: Information about the most recent modification of
	// this placement. This is a read-only field.
	LastModifiedInfo *LastModifiedInfo `json:"lastModifiedInfo,omitempty"`

	// LookbackConfiguration: Lookback window settings for this placement.
	LookbackConfiguration *LookbackConfiguration `json:"lookbackConfiguration,omitempty"`

	// Name: Name of this placement.This is a required field and must be
	// less than 256 characters long.
	Name string `json:"name,omitempty"`

	// PaymentApproved: Whether payment was approved for this placement.
	// This is a read-only field relevant only to publisher-paid placements.
	PaymentApproved bool `json:"paymentApproved,omitempty"`

	// PaymentSource: Payment source for this placement. This is a required
	// field that is read-only after insertion.
	PaymentSource string `json:"paymentSource,omitempty"`

	// PlacementGroupId: ID of this placement's group, if applicable.
	PlacementGroupId int64 `json:"placementGroupId,omitempty,string"`

	// PlacementGroupIdDimensionValue: Dimension value for the ID of the
	// placement group. This is a read-only, auto-generated field.
	PlacementGroupIdDimensionValue *DimensionValue `json:"placementGroupIdDimensionValue,omitempty"`

	// PlacementStrategyId: ID of the placement strategy assigned to this
	// placement.
	PlacementStrategyId int64 `json:"placementStrategyId,omitempty,string"`

	// PricingSchedule: Pricing schedule of this placement. This field is
	// required on insertion, specifically subfields startDate, endDate and
	// pricingType.
	PricingSchedule *PricingSchedule `json:"pricingSchedule,omitempty"`

	// Primary: Whether this placement is the primary placement of a
	// roadblock (placement group). You cannot change this field from true
	// to false. Setting this field to true will automatically set the
	// primary field on the original primary placement of the roadblock to
	// false, and it will automatically set the roadblock's
	// primaryPlacementId field to the ID of this placement.
	Primary bool `json:"primary,omitempty"`

	// PublisherUpdateInfo: Information about the last publisher update.
	// This is a read-only field.
	PublisherUpdateInfo *LastModifiedInfo `json:"publisherUpdateInfo,omitempty"`

	// SiteId: Site ID associated with this placement. On insert, you must
	// set either this field or the directorySiteId field to specify the
	// site associated with this placement. This is a required field that is
	// read-only after insertion.
	SiteId int64 `json:"siteId,omitempty,string"`

	// SiteIdDimensionValue: Dimension value for the ID of the site. This is
	// a read-only, auto-generated field.
	SiteIdDimensionValue *DimensionValue `json:"siteIdDimensionValue,omitempty"`

	// Size: Size associated with this placement. When inserting or updating
	// a placement, only the size ID field is used. This field is required
	// on insertion.
	Size *Size `json:"size,omitempty"`

	// SslRequired: Whether creatives assigned to this placement must be
	// SSL-compliant.
	SslRequired bool `json:"sslRequired,omitempty"`

	// Status: Third-party placement status.
	Status string `json:"status,omitempty"`

	// SubaccountId: Subaccount ID of this placement. This field can be left
	// blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`

	// TagFormats: Tag formats to generate for this placement. This field is
	// required on insertion.
	// Acceptable values are:
	// -
	// "PLACEMENT_TAG_STANDARD"
	// - "PLACEMENT_TAG_IFRAME_JAVASCRIPT"
	// -
	// "PLACEMENT_TAG_IFRAME_ILAYER"
	// - "PLACEMENT_TAG_INTERNAL_REDIRECT"
	// -
	// "PLACEMENT_TAG_JAVASCRIPT"
	// -
	// "PLACEMENT_TAG_INTERSTITIAL_IFRAME_JAVASCRIPT"
	// -
	// "PLACEMENT_TAG_INTERSTITIAL_INTERNAL_REDIRECT"
	// -
	// "PLACEMENT_TAG_INTERSTITIAL_JAVASCRIPT"
	// -
	// "PLACEMENT_TAG_CLICK_COMMANDS"
	// -
	// "PLACEMENT_TAG_INSTREAM_VIDEO_PREFETCH"
	// - "PLACEMENT_TAG_TRACKING"
	// -
	// "PLACEMENT_TAG_TRACKING_IFRAME"
	// - "PLACEMENT_TAG_TRACKING_JAVASCRIPT"
	TagFormats []string `json:"tagFormats,omitempty"`

	// TagSetting: Tag settings for this placement.
	TagSetting *TagSetting `json:"tagSetting,omitempty"`
}

type PlacementAssignment struct {
	// Active: Whether this placement assignment is active. When true, the
	// placement will be included in the ad's rotation.
	Active bool `json:"active,omitempty"`

	// PlacementId: ID of the placement to be assigned. This is a required
	// field.
	PlacementId int64 `json:"placementId,omitempty,string"`

	// PlacementIdDimensionValue: Dimension value for the ID of the
	// placement. This is a read-only, auto-generated field.
	PlacementIdDimensionValue *DimensionValue `json:"placementIdDimensionValue,omitempty"`

	// SslRequired: Whether the placement to be assigned requires SSL. This
	// is a read-only field that is auto-generated when the ad is inserted
	// or updated.
	SslRequired bool `json:"sslRequired,omitempty"`
}

type PlacementGroup struct {
	// AccountId: Account ID of this placement group. This is a read-only
	// field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AdvertiserId: Advertiser ID of this placement group. This is a
	// required field on insertion.
	AdvertiserId int64 `json:"advertiserId,omitempty,string"`

	// AdvertiserIdDimensionValue: Dimension value for the ID of the
	// advertiser. This is a read-only, auto-generated field.
	AdvertiserIdDimensionValue *DimensionValue `json:"advertiserIdDimensionValue,omitempty"`

	// Archived: Whether this placement group is archived.
	Archived bool `json:"archived,omitempty"`

	// CampaignId: Campaign ID of this placement group. This field is
	// required on insertion.
	CampaignId int64 `json:"campaignId,omitempty,string"`

	// CampaignIdDimensionValue: Dimension value for the ID of the campaign.
	// This is a read-only, auto-generated field.
	CampaignIdDimensionValue *DimensionValue `json:"campaignIdDimensionValue,omitempty"`

	// ChildPlacementIds: IDs of placements which are assigned to this
	// placement group. This is a read-only, auto-generated field.
	ChildPlacementIds googleapi.Int64s `json:"childPlacementIds,omitempty"`

	// Comment: Comments for this placement group.
	Comment string `json:"comment,omitempty"`

	// ContentCategoryId: ID of the content category assigned to this
	// placement group.
	ContentCategoryId int64 `json:"contentCategoryId,omitempty,string"`

	// CreateInfo: Information about the creation of this placement group.
	// This is a read-only field.
	CreateInfo *LastModifiedInfo `json:"createInfo,omitempty"`

	// DirectorySiteId: Directory site ID associated with this placement
	// group. On insert, you must set either this field or the site_id field
	// to specify the site associated with this placement group. This is a
	// required field that is read-only after insertion.
	DirectorySiteId int64 `json:"directorySiteId,omitempty,string"`

	// DirectorySiteIdDimensionValue: Dimension value for the ID of the
	// directory site. This is a read-only, auto-generated field.
	DirectorySiteIdDimensionValue *DimensionValue `json:"directorySiteIdDimensionValue,omitempty"`

	// ExternalId: External ID for this placement.
	ExternalId string `json:"externalId,omitempty"`

	// Id: ID of this placement group. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this placement group.
	// This is a read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placementGroup".
	Kind string `json:"kind,omitempty"`

	// LastModifiedInfo: Information about the most recent modification of
	// this placement group. This is a read-only field.
	LastModifiedInfo *LastModifiedInfo `json:"lastModifiedInfo,omitempty"`

	// Name: Name of this placement group. This is a required field and must
	// be less than 256 characters long.
	Name string `json:"name,omitempty"`

	// PlacementGroupType: Type of this placement group. A package is a
	// simple group of placements that acts as a single pricing point for a
	// group of tags. A roadblock is a group of placements that not only
	// acts as a single pricing point, but also assumes that all the tags in
	// it will be served at the same time. A roadblock requires one of its
	// assigned placements to be marked as primary for reporting. This field
	// is required on insertion.
	PlacementGroupType string `json:"placementGroupType,omitempty"`

	// PlacementStrategyId: ID of the placement strategy assigned to this
	// placement group.
	PlacementStrategyId int64 `json:"placementStrategyId,omitempty,string"`

	// PricingSchedule: Pricing schedule of this placement group. This field
	// is required on insertion.
	PricingSchedule *PricingSchedule `json:"pricingSchedule,omitempty"`

	// PrimaryPlacementId: ID of the primary placement, used to calculate
	// the media cost of a roadblock (placement group). Modifying this field
	// will automatically modify the primary field on all affected roadblock
	// child placements.
	PrimaryPlacementId int64 `json:"primaryPlacementId,omitempty,string"`

	// PrimaryPlacementIdDimensionValue: Dimension value for the ID of the
	// primary placement. This is a read-only, auto-generated field.
	PrimaryPlacementIdDimensionValue *DimensionValue `json:"primaryPlacementIdDimensionValue,omitempty"`

	// ProgrammaticSetting: Settings for a programmatic placement.
	ProgrammaticSetting *ProgrammaticSetting `json:"programmaticSetting,omitempty"`

	// SiteId: Site ID associated with this placement group. On insert, you
	// must set either this field or the directorySiteId field to specify
	// the site associated with this placement group. This is a required
	// field that is read-only after insertion.
	SiteId int64 `json:"siteId,omitempty,string"`

	// SiteIdDimensionValue: Dimension value for the ID of the site. This is
	// a read-only, auto-generated field.
	SiteIdDimensionValue *DimensionValue `json:"siteIdDimensionValue,omitempty"`

	// SubaccountId: Subaccount ID of this placement group. This is a
	// read-only field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`
}

type PlacementGroupsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placementGroupsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PlacementGroups: Placement group collection
	PlacementGroups []*PlacementGroup `json:"placementGroups,omitempty"`
}

type PlacementStrategiesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placementStrategiesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// PlacementStrategies: Placement strategy collection
	PlacementStrategies []*PlacementStrategy `json:"placementStrategies,omitempty"`
}

type PlacementStrategy struct {
	// AccountId: Account ID of this placement strategy.This is a read-only
	// field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Id: ID of this placement strategy. This is a read-only,
	// auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placementStrategy".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this placement strategy. This is a required field. Must
	// be less than 256 characters long and unique among placement
	// strategies of the same account.
	Name string `json:"name,omitempty"`
}

type PlacementTag struct {
	// PlacementId: Placement ID
	PlacementId int64 `json:"placementId,omitempty,string"`

	// TagDatas: Tags generated for this placement.
	TagDatas []*TagData `json:"tagDatas,omitempty"`
}

type PlacementsGenerateTagsResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placementsGenerateTagsResponse".
	Kind string `json:"kind,omitempty"`

	// PlacementTags: Set of generated tags for the specified placements.
	PlacementTags []*PlacementTag `json:"placementTags,omitempty"`
}

type PlacementsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#placementsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Placements: Placement collection
	Placements []*Placement `json:"placements,omitempty"`
}

type PlatformType struct {
	// Id: ID of this platform type.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#platformType".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this platform type.
	Name string `json:"name,omitempty"`
}

type PlatformTypesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#platformTypesListResponse".
	Kind string `json:"kind,omitempty"`

	// PlatformTypes: Platform type collection
	PlatformTypes []*PlatformType `json:"platformTypes,omitempty"`
}

type PopupWindowProperties struct {
	// Dimension: Popup dimension for a creative. This is a read-only field.
	// Applicable to the following creative types: all RICH_MEDIA and all
	// VPAID
	Dimension *Size `json:"dimension,omitempty"`

	// Offset: Upper-left corner coordinates of the popup window. Applicable
	// if positionType is COORDINATES.
	Offset *OffsetPosition `json:"offset,omitempty"`

	// PositionType: Popup window position either centered or at specific
	// coordinate.
	PositionType string `json:"positionType,omitempty"`

	// ShowAddressBar: Whether to display the browser address bar.
	ShowAddressBar bool `json:"showAddressBar,omitempty"`

	// ShowMenuBar: Whether to display the browser menu bar.
	ShowMenuBar bool `json:"showMenuBar,omitempty"`

	// ShowScrollBar: Whether to display the browser scroll bar.
	ShowScrollBar bool `json:"showScrollBar,omitempty"`

	// ShowStatusBar: Whether to display the browser status bar.
	ShowStatusBar bool `json:"showStatusBar,omitempty"`

	// ShowToolBar: Whether to display the browser tool bar.
	ShowToolBar bool `json:"showToolBar,omitempty"`

	// Title: Title of popup window.
	Title string `json:"title,omitempty"`
}

type PostalCode struct {
	// CountryCode: Country code of the country to which this postal code
	// belongs.
	CountryCode string `json:"countryCode,omitempty"`

	// CountryDartId: DART ID of the country to which this postal code
	// belongs.
	CountryDartId int64 `json:"countryDartId,omitempty,string"`

	// Id: ID of this postal code.
	Id string `json:"id,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#postalCode".
	Kind string `json:"kind,omitempty"`
}

type PostalCodesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#postalCodesListResponse".
	Kind string `json:"kind,omitempty"`

	// PostalCodes: Postal code collection
	PostalCodes []*PostalCode `json:"postalCodes,omitempty"`
}

type PricingSchedule struct {
	// CapCostOption: Placement cap cost option.
	CapCostOption string `json:"capCostOption,omitempty"`

	// DisregardOverdelivery: Whether cap costs are ignored by ad serving.
	DisregardOverdelivery bool `json:"disregardOverdelivery,omitempty"`

	// EndDate: Placement end date. This date must be later than or be the
	// same day as the placement start date, but not later than the campaign
	// end date. If, for example, you set 6/25/2015 as both the start and
	// end dates, the effective placement date is just that day only,
	// 6/25/2015. The hours, minutes, and seconds of the end date should not
	// be set, as doing so will result in an error. This field is required
	// on insertion.
	EndDate string `json:"endDate,omitempty"`

	// Flighted: Whether this placement is flighted. If true, pricing
	// periods will be computed automatically.
	Flighted bool `json:"flighted,omitempty"`

	// FloodlightActivityId: Floodlight activity ID associated with this
	// placement. This field should be set when placement pricing type is
	// set to PRICING_TYPE_CPA.
	FloodlightActivityId int64 `json:"floodlightActivityId,omitempty,string"`

	// PricingPeriods: Pricing periods for this placement.
	PricingPeriods []*PricingSchedulePricingPeriod `json:"pricingPeriods,omitempty"`

	// PricingType: Placement pricing type. This field is required on
	// insertion.
	PricingType string `json:"pricingType,omitempty"`

	// StartDate: Placement start date. This date cannot be earlier than the
	// campaign start date. The hours, minutes, and seconds of the start
	// date should not be set, as doing so will result in an error. This
	// field is required on insertion.
	StartDate string `json:"startDate,omitempty"`

	// TestingStartDate: Testing start date of this placement. The hours,
	// minutes, and seconds of the start date should not be set, as doing so
	// will result in an error.
	TestingStartDate string `json:"testingStartDate,omitempty"`
}

type PricingSchedulePricingPeriod struct {
	// EndDate: Pricing period end date. This date must be later than or be
	// the same day as the pricing period start date, but not later than the
	// placement end date. The period end date can be the same date as the
	// period start date. If, for example, you set 6/25/2015 as both the
	// start and end dates, the effective pricing period date is just that
	// day only, 6/25/2015. The hours, minutes, and seconds of the end date
	// should not be set, as doing so will result in an error.
	EndDate string `json:"endDate,omitempty"`

	// PricingComment: Comments for this pricing period.
	PricingComment string `json:"pricingComment,omitempty"`

	// RateOrCostNanos: Rate or cost of this pricing period.
	RateOrCostNanos int64 `json:"rateOrCostNanos,omitempty,string"`

	// StartDate: Pricing period start date. This date cannot be earlier
	// than the placement start date. The hours, minutes, and seconds of the
	// start date should not be set, as doing so will result in an error.
	StartDate string `json:"startDate,omitempty"`

	// Units: Units of this pricing period.
	Units int64 `json:"units,omitempty,string"`
}

type ProgrammaticSetting struct {
	// AdxDealIds: Adx deal IDs assigned to the placement.
	AdxDealIds googleapi.Int64s `json:"adxDealIds,omitempty"`

	// InsertionOrderId: Insertion order ID.
	InsertionOrderId string `json:"insertionOrderId,omitempty"`

	// InsertionOrderIdStatus: Whether insertion order ID has been placed in
	// DFP. This is a read-only field.
	InsertionOrderIdStatus bool `json:"insertionOrderIdStatus,omitempty"`

	// MediaCostNanos: Media cost for the programmatic placement.
	MediaCostNanos int64 `json:"mediaCostNanos,omitempty,string"`

	// Programmatic: Whether programmatic is enabled.
	Programmatic bool `json:"programmatic,omitempty"`

	// TraffickerEmails: Trafficker emails assigned to the placement.
	TraffickerEmails []string `json:"traffickerEmails,omitempty"`
}

type ReachReportCompatibleFields struct {
	// DimensionFilters: Dimensions which are compatible to be selected in
	// the "dimensionFilters" section of the report.
	DimensionFilters []*Dimension `json:"dimensionFilters,omitempty"`

	// Dimensions: Dimensions which are compatible to be selected in the
	// "dimensions" section of the report.
	Dimensions []*Dimension `json:"dimensions,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#reachReportCompatibleFields.
	Kind string `json:"kind,omitempty"`

	// Metrics: Metrics which are compatible to be selected in the
	// "metricNames" section of the report.
	Metrics []*Metric `json:"metrics,omitempty"`

	// PivotedActivityMetrics: Metrics which are compatible to be selected
	// as activity metrics to pivot on in the "activities" section of the
	// report.
	PivotedActivityMetrics []*Metric `json:"pivotedActivityMetrics,omitempty"`

	// ReachByFrequencyMetrics: Metrics which are compatible to be selected
	// in the "reachByFrequencyMetricNames" section of the report.
	ReachByFrequencyMetrics []*Metric `json:"reachByFrequencyMetrics,omitempty"`
}

type Recipient struct {
	// DeliveryType: The delivery type for the recipient.
	DeliveryType string `json:"deliveryType,omitempty"`

	// Email: The email address of the recipient.
	Email string `json:"email,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#recipient.
	Kind string `json:"kind,omitempty"`
}

type Region struct {
	// CountryCode: Country code of the country to which this region
	// belongs.
	CountryCode string `json:"countryCode,omitempty"`

	// CountryDartId: DART ID of the country to which this region belongs.
	CountryDartId int64 `json:"countryDartId,omitempty,string"`

	// DartId: DART ID of this region.
	DartId int64 `json:"dartId,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#region".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this region.
	Name string `json:"name,omitempty"`

	// RegionCode: Region code.
	RegionCode string `json:"regionCode,omitempty"`
}

type RegionsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#regionsListResponse".
	Kind string `json:"kind,omitempty"`

	// Regions: Region Collection.
	Regions []*Region `json:"regions,omitempty"`
}

type Report struct {
	// AccountId: The account ID to which this report belongs.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Criteria: The report criteria for a report of type "STANDARD".
	Criteria *ReportCriteria `json:"criteria,omitempty"`

	// CrossDimensionReachCriteria: The report criteria for a report of type
	// "CROSS_DIMENSION_REACH".
	CrossDimensionReachCriteria *ReportCrossDimensionReachCriteria `json:"crossDimensionReachCriteria,omitempty"`

	// Delivery: The report's email delivery settings.
	Delivery *ReportDelivery `json:"delivery,omitempty"`

	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// FileName: The filename used when generating report files for this
	// report.
	FileName string `json:"fileName,omitempty"`

	// FloodlightCriteria: The report criteria for a report of type
	// "FLOODLIGHT".
	FloodlightCriteria *ReportFloodlightCriteria `json:"floodlightCriteria,omitempty"`

	// Format: The output format of the report. If not specified, default
	// format is "CSV". Note that the actual format in the completed report
	// file might differ if for instance the report's size exceeds the
	// format's capabilities. "CSV" will then be the fallback format.
	Format string `json:"format,omitempty"`

	// Id: The unique ID identifying this report resource.
	Id int64 `json:"id,omitempty,string"`

	// Kind: The kind of resource this is, in this case dfareporting#report.
	Kind string `json:"kind,omitempty"`

	// LastModifiedTime: The timestamp (in milliseconds since epoch) of when
	// this report was last modified.
	LastModifiedTime uint64 `json:"lastModifiedTime,omitempty,string"`

	// Name: The name of the report.
	Name string `json:"name,omitempty"`

	// OwnerProfileId: The user profile id of the owner of this report.
	OwnerProfileId int64 `json:"ownerProfileId,omitempty,string"`

	// PathToConversionCriteria: The report criteria for a report of type
	// "PATH_TO_CONVERSION".
	PathToConversionCriteria *ReportPathToConversionCriteria `json:"pathToConversionCriteria,omitempty"`

	// ReachCriteria: The report criteria for a report of type "REACH".
	ReachCriteria *ReportReachCriteria `json:"reachCriteria,omitempty"`

	// Schedule: The report's schedule. Can only be set if the report's
	// 'dateRange' is a relative date range and the relative date range is
	// not "TODAY".
	Schedule *ReportSchedule `json:"schedule,omitempty"`

	// SubAccountId: The subaccount ID to which this report belongs if
	// applicable.
	SubAccountId int64 `json:"subAccountId,omitempty,string"`

	// Type: The type of the report.
	Type string `json:"type,omitempty"`
}

type ReportCriteria struct {
	// Activities: Activity group.
	Activities *Activities `json:"activities,omitempty"`

	// CustomRichMediaEvents: Custom Rich Media Events group.
	CustomRichMediaEvents *CustomRichMediaEvents `json:"customRichMediaEvents,omitempty"`

	// DateRange: The date range for which this report should be run.
	DateRange *DateRange `json:"dateRange,omitempty"`

	// DimensionFilters: The list of filters on which dimensions are
	// filtered.
	// Filters for different dimensions are ANDed, filters for the
	// same dimension are grouped together and ORed.
	DimensionFilters []*DimensionValue `json:"dimensionFilters,omitempty"`

	// Dimensions: The list of standard dimensions the report should
	// include.
	Dimensions []*SortedDimension `json:"dimensions,omitempty"`

	// MetricNames: The list of names of metrics the report should include.
	MetricNames []string `json:"metricNames,omitempty"`
}

type ReportCrossDimensionReachCriteria struct {
	// Breakdown: The list of dimensions the report should include.
	Breakdown []*SortedDimension `json:"breakdown,omitempty"`

	// DateRange: The date range this report should be run for.
	DateRange *DateRange `json:"dateRange,omitempty"`

	// Dimension: The dimension option.
	Dimension string `json:"dimension,omitempty"`

	// DimensionFilters: The list of filters on which dimensions are
	// filtered.
	DimensionFilters []*DimensionValue `json:"dimensionFilters,omitempty"`

	// MetricNames: The list of names of metrics the report should include.
	MetricNames []string `json:"metricNames,omitempty"`

	// OverlapMetricNames: The list of names of overlap metrics the report
	// should include.
	OverlapMetricNames []string `json:"overlapMetricNames,omitempty"`

	// Pivoted: Whether the report is pivoted or not. Defaults to true.
	Pivoted bool `json:"pivoted,omitempty"`
}

type ReportDelivery struct {
	// EmailOwner: Whether the report should be emailed to the report owner.
	EmailOwner bool `json:"emailOwner,omitempty"`

	// EmailOwnerDeliveryType: The type of delivery for the owner to
	// receive, if enabled.
	EmailOwnerDeliveryType string `json:"emailOwnerDeliveryType,omitempty"`

	// Message: The message to be sent with each email.
	Message string `json:"message,omitempty"`

	// Recipients: The list of recipients to which to email the report.
	Recipients []*Recipient `json:"recipients,omitempty"`
}

type ReportFloodlightCriteria struct {
	// CustomRichMediaEvents: The list of custom rich media events to
	// include.
	CustomRichMediaEvents []*DimensionValue `json:"customRichMediaEvents,omitempty"`

	// DateRange: The date range this report should be run for.
	DateRange *DateRange `json:"dateRange,omitempty"`

	// DimensionFilters: The list of filters on which dimensions are
	// filtered.
	// Filters for different dimensions are ANDed, filters for the
	// same dimension are grouped together and ORed.
	DimensionFilters []*DimensionValue `json:"dimensionFilters,omitempty"`

	// Dimensions: The list of dimensions the report should include.
	Dimensions []*SortedDimension `json:"dimensions,omitempty"`

	// FloodlightConfigId: The floodlight ID for which to show data in this
	// report. All advertisers associated with that ID will automatically be
	// added. The dimension of the value needs to be
	// 'dfa:floodlightConfigId'.
	FloodlightConfigId *DimensionValue `json:"floodlightConfigId,omitempty"`

	// MetricNames: The list of names of metrics the report should include.
	MetricNames []string `json:"metricNames,omitempty"`

	// ReportProperties: The properties of the report.
	ReportProperties *ReportFloodlightCriteriaReportProperties `json:"reportProperties,omitempty"`
}

type ReportFloodlightCriteriaReportProperties struct {
	// IncludeAttributedIPConversions: Include conversions that have no
	// cookie, but do have an exposure path.
	IncludeAttributedIPConversions bool `json:"includeAttributedIPConversions,omitempty"`

	// IncludeUnattributedCookieConversions: Include conversions of users
	// with a DoubleClick cookie but without an exposure. That means the
	// user did not click or see an ad from the advertiser within the
	// Floodlight group, or that the interaction happened outside the
	// lookback window.
	IncludeUnattributedCookieConversions bool `json:"includeUnattributedCookieConversions,omitempty"`

	// IncludeUnattributedIPConversions: Include conversions that have no
	// associated cookies and no exposures. It’s therefore impossible to
	// know how the user was exposed to your ads during the lookback window
	// prior to a conversion.
	IncludeUnattributedIPConversions bool `json:"includeUnattributedIPConversions,omitempty"`
}

type ReportPathToConversionCriteria struct {
	// ActivityFilters: The list of 'dfa:activity' values to filter on.
	ActivityFilters []*DimensionValue `json:"activityFilters,omitempty"`

	// ConversionDimensions: The list of conversion dimensions the report
	// should include.
	ConversionDimensions []*SortedDimension `json:"conversionDimensions,omitempty"`

	// CustomFloodlightVariables: The list of custom floodlight variables
	// the report should include.
	CustomFloodlightVariables []*SortedDimension `json:"customFloodlightVariables,omitempty"`

	// CustomRichMediaEvents: The list of custom rich media events to
	// include.
	CustomRichMediaEvents []*DimensionValue `json:"customRichMediaEvents,omitempty"`

	// DateRange: The date range this report should be run for.
	DateRange *DateRange `json:"dateRange,omitempty"`

	// FloodlightConfigId: The floodlight ID for which to show data in this
	// report. All advertisers associated with that ID will automatically be
	// added. The dimension of the value needs to be
	// 'dfa:floodlightConfigId'.
	FloodlightConfigId *DimensionValue `json:"floodlightConfigId,omitempty"`

	// MetricNames: The list of names of metrics the report should include.
	MetricNames []string `json:"metricNames,omitempty"`

	// PerInteractionDimensions: The list of per interaction dimensions the
	// report should include.
	PerInteractionDimensions []*SortedDimension `json:"perInteractionDimensions,omitempty"`

	// ReportProperties: The properties of the report.
	ReportProperties *ReportPathToConversionCriteriaReportProperties `json:"reportProperties,omitempty"`
}

type ReportPathToConversionCriteriaReportProperties struct {
	// ClicksLookbackWindow: DFA checks to see if a click interaction
	// occurred within the specified period of time before a conversion. By
	// default the value is pulled from Floodlight or you can manually enter
	// a custom value. Valid values: 1-90.
	ClicksLookbackWindow int64 `json:"clicksLookbackWindow,omitempty"`

	// ImpressionsLookbackWindow: DFA checks to see if an impression
	// interaction occurred within the specified period of time before a
	// conversion. By default the value is pulled from Floodlight or you can
	// manually enter a custom value. Valid values: 1-90.
	ImpressionsLookbackWindow int64 `json:"impressionsLookbackWindow,omitempty"`

	// IncludeAttributedIPConversions: Deprecated: has no effect.
	IncludeAttributedIPConversions bool `json:"includeAttributedIPConversions,omitempty"`

	// IncludeUnattributedCookieConversions: Include conversions of users
	// with a DoubleClick cookie but without an exposure. That means the
	// user did not click or see an ad from the advertiser within the
	// Floodlight group, or that the interaction happened outside the
	// lookback window.
	IncludeUnattributedCookieConversions bool `json:"includeUnattributedCookieConversions,omitempty"`

	// IncludeUnattributedIPConversions: Include conversions that have no
	// associated cookies and no exposures. It’s therefore impossible to
	// know how the user was exposed to your ads during the lookback window
	// prior to a conversion.
	IncludeUnattributedIPConversions bool `json:"includeUnattributedIPConversions,omitempty"`

	// MaximumClickInteractions: The maximum number of click interactions to
	// include in the report. Advertisers currently paying for E2C reports
	// get up to 200 (100 clicks, 100 impressions). If another advertiser in
	// your network is paying for E2C, you can have up to 5 total exposures
	// per report.
	MaximumClickInteractions int64 `json:"maximumClickInteractions,omitempty"`

	// MaximumImpressionInteractions: The maximum number of click
	// interactions to include in the report. Advertisers currently paying
	// for E2C reports get up to 200 (100 clicks, 100 impressions). If
	// another advertiser in your network is paying for E2C, you can have up
	// to 5 total exposures per report.
	MaximumImpressionInteractions int64 `json:"maximumImpressionInteractions,omitempty"`

	// MaximumInteractionGap: The maximum amount of time that can take place
	// between interactions (clicks or impressions) by the same user. Valid
	// values: 1-90.
	MaximumInteractionGap int64 `json:"maximumInteractionGap,omitempty"`

	// PivotOnInteractionPath: Enable pivoting on interaction path.
	PivotOnInteractionPath bool `json:"pivotOnInteractionPath,omitempty"`
}

type ReportReachCriteria struct {
	// Activities: Activity group.
	Activities *Activities `json:"activities,omitempty"`

	// CustomRichMediaEvents: Custom Rich Media Events group.
	CustomRichMediaEvents *CustomRichMediaEvents `json:"customRichMediaEvents,omitempty"`

	// DateRange: The date range this report should be run for.
	DateRange *DateRange `json:"dateRange,omitempty"`

	// DimensionFilters: The list of filters on which dimensions are
	// filtered.
	// Filters for different dimensions are ANDed, filters for the
	// same dimension are grouped together and ORed.
	DimensionFilters []*DimensionValue `json:"dimensionFilters,omitempty"`

	// Dimensions: The list of dimensions the report should include.
	Dimensions []*SortedDimension `json:"dimensions,omitempty"`

	// EnableAllDimensionCombinations: Whether to enable all reach dimension
	// combinations in the report. Defaults to false. If enabled, the date
	// range of the report should be within the last three months.
	EnableAllDimensionCombinations bool `json:"enableAllDimensionCombinations,omitempty"`

	// MetricNames: The list of names of metrics the report should include.
	MetricNames []string `json:"metricNames,omitempty"`

	// ReachByFrequencyMetricNames: The list of names of  Reach By Frequency
	// metrics the report should include.
	ReachByFrequencyMetricNames []string `json:"reachByFrequencyMetricNames,omitempty"`
}

type ReportSchedule struct {
	// Active: Whether the schedule is active or not. Must be set to either
	// true or false.
	Active bool `json:"active,omitempty"`

	// Every: Defines every how many days, weeks or months the report should
	// be run. Needs to be set when "repeats" is either "DAILY", "WEEKLY" or
	// "MONTHLY".
	Every int64 `json:"every,omitempty"`

	// ExpirationDate: The expiration date when the scheduled report stops
	// running.
	ExpirationDate string `json:"expirationDate,omitempty"`

	// Repeats: The interval for which the report is repeated. Note:
	// -
	// "DAILY" also requires field "every" to be set.
	// - "WEEKLY" also
	// requires fields "every" and "repeatsOnWeekDays" to be set.
	// -
	// "MONTHLY" also requires fields "every" and "runsOnDayOfMonth" to be
	// set.
	Repeats string `json:"repeats,omitempty"`

	// RepeatsOnWeekDays: List of week days "WEEKLY" on which scheduled
	// reports should run.
	RepeatsOnWeekDays []string `json:"repeatsOnWeekDays,omitempty"`

	// RunsOnDayOfMonth: Enum to define for "MONTHLY" scheduled reports
	// whether reports should be repeated on the same day of the month as
	// "startDate" or the same day of the week of the month.
	// Example: If
	// 'startDate' is Monday, April 2nd 2012 (2012-04-02), "DAY_OF_MONTH"
	// would run subsequent reports on the 2nd of every Month, and
	// "WEEK_OF_MONTH" would run subsequent reports on the first Monday of
	// the month.
	RunsOnDayOfMonth string `json:"runsOnDayOfMonth,omitempty"`

	// StartDate: Start date of date range for which scheduled reports
	// should be run.
	StartDate string `json:"startDate,omitempty"`
}

type ReportCompatibleFields struct {
	// DimensionFilters: Dimensions which are compatible to be selected in
	// the "dimensionFilters" section of the report.
	DimensionFilters []*Dimension `json:"dimensionFilters,omitempty"`

	// Dimensions: Dimensions which are compatible to be selected in the
	// "dimensions" section of the report.
	Dimensions []*Dimension `json:"dimensions,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#reportCompatibleFields.
	Kind string `json:"kind,omitempty"`

	// Metrics: Metrics which are compatible to be selected in the
	// "metricNames" section of the report.
	Metrics []*Metric `json:"metrics,omitempty"`

	// PivotedActivityMetrics: Metrics which are compatible to be selected
	// as activity metrics to pivot on in the "activities" section of the
	// report.
	PivotedActivityMetrics []*Metric `json:"pivotedActivityMetrics,omitempty"`
}

type ReportList struct {
	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The reports returned in this response.
	Items []*Report `json:"items,omitempty"`

	// Kind: The kind of list this is, in this case dfareporting#reportList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Continuation token used to page through reports. To
	// retrieve the next page of results, set the next request's "pageToken"
	// to the value of this field. The page token is only valid for a
	// limited amount of time and should not be persisted.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ReportsConfiguration struct {
	// ExposureToConversionEnabled: Whether the exposure to conversion
	// report is enabled. This report shows detailed pathway information on
	// up to 10 of the most recent ad exposures seen by a user before
	// converting.
	ExposureToConversionEnabled bool `json:"exposureToConversionEnabled,omitempty"`

	// LookbackConfiguration: Default lookback windows for new advertisers
	// in this account.
	LookbackConfiguration *LookbackConfiguration `json:"lookbackConfiguration,omitempty"`

	// ReportGenerationTimeZoneId: Report generation time zone ID of this
	// account. This is a required field that can only be changed by a
	// superuser.
	// Acceptable values are:
	//
	// - "1" for "America/New_York"
	// -
	// "2" for "Europe/London"
	// - "3" for "Europe/Paris"
	// - "4" for
	// "Africa/Johannesburg"
	// - "5" for "Asia/Jerusalem"
	// - "6" for
	// "Asia/Shanghai"
	// - "7" for "Asia/Hong_Kong"
	// - "8" for "Asia/Tokyo"
	//
	// - "9" for "Australia/Sydney"
	// - "10" for "Asia/Dubai"
	// - "11" for
	// "America/Los_Angeles"
	// - "12" for "Pacific/Auckland"
	// - "13" for
	// "America/Sao_Paulo"
	ReportGenerationTimeZoneId int64 `json:"reportGenerationTimeZoneId,omitempty,string"`
}

type RichMediaExitOverride struct {
	// CustomExitUrl: Click-through URL to override the default exit URL.
	// Applicable if the useCustomExitUrl field is set to true.
	CustomExitUrl string `json:"customExitUrl,omitempty"`

	// ExitId: ID for the override to refer to a specific exit in the
	// creative.
	ExitId int64 `json:"exitId,omitempty,string"`

	// UseCustomExitUrl: Whether to use the custom exit URL.
	UseCustomExitUrl bool `json:"useCustomExitUrl,omitempty"`
}

type Site struct {
	// AccountId: Account ID of this site. This is a read-only field that
	// can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// Approved: Whether this site is approved.
	Approved bool `json:"approved,omitempty"`

	// DirectorySiteId: Directory site associated with this site. This is a
	// required field that is read-only after insertion.
	DirectorySiteId int64 `json:"directorySiteId,omitempty,string"`

	// DirectorySiteIdDimensionValue: Dimension value for the ID of the
	// directory site. This is a read-only, auto-generated field.
	DirectorySiteIdDimensionValue *DimensionValue `json:"directorySiteIdDimensionValue,omitempty"`

	// Id: ID of this site. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// IdDimensionValue: Dimension value for the ID of this site. This is a
	// read-only, auto-generated field.
	IdDimensionValue *DimensionValue `json:"idDimensionValue,omitempty"`

	// KeyName: Key name of this site. This is a read-only, auto-generated
	// field.
	KeyName string `json:"keyName,omitempty"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#site".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this site.This is a required field. Must be less than
	// 128 characters long. If this site is under a subaccount, the name
	// must be unique among sites of the same subaccount. Otherwise, this
	// site is a top-level site, and the name must be unique among top-level
	// sites of the same account.
	Name string `json:"name,omitempty"`

	// SiteContacts: Site contacts.
	SiteContacts []*SiteContact `json:"siteContacts,omitempty"`

	// SiteSettings: Site-wide settings.
	SiteSettings *SiteSettings `json:"siteSettings,omitempty"`

	// SubaccountId: Subaccount ID of this site. This is a read-only field
	// that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`
}

type SiteContact struct {
	// ContactType: Site contact type.
	ContactType string `json:"contactType,omitempty"`

	// Email: Email address of this site contact. This is a required field.
	Email string `json:"email,omitempty"`

	// FirstName: First name of this site contact.
	FirstName string `json:"firstName,omitempty"`

	// Id: ID of this site contact. This is a read-only, auto-generated
	// field.
	Id int64 `json:"id,omitempty,string"`

	// LastName: Last name of this site contact.
	LastName string `json:"lastName,omitempty"`
}

type SiteSettings struct {
	// ActiveViewOptOut: Whether active view creatives are disabled for this
	// site.
	ActiveViewOptOut bool `json:"activeViewOptOut,omitempty"`

	// CreativeSettings: Site-wide creative settings.
	CreativeSettings *CreativeSettings `json:"creativeSettings,omitempty"`

	// DisableBrandSafeAds: Whether brand safe ads are disabled for this
	// site.
	DisableBrandSafeAds bool `json:"disableBrandSafeAds,omitempty"`

	// DisableNewCookie: Whether new cookies are disabled for this site.
	DisableNewCookie bool `json:"disableNewCookie,omitempty"`

	// LookbackConfiguration: Lookback window settings for this site.
	LookbackConfiguration *LookbackConfiguration `json:"lookbackConfiguration,omitempty"`

	// TagSetting: Configuration settings for dynamic and image floodlight
	// tags.
	TagSetting *TagSetting `json:"tagSetting,omitempty"`
}

type SitesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#sitesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Sites: Site collection
	Sites []*Site `json:"sites,omitempty"`
}

type Size struct {
	// Height: Height of this size.
	Height int64 `json:"height,omitempty"`

	// Iab: IAB standard size. This is a read-only, auto-generated field.
	Iab bool `json:"iab,omitempty"`

	// Id: ID of this size. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#size".
	Kind string `json:"kind,omitempty"`

	// Width: Width of this size.
	Width int64 `json:"width,omitempty"`
}

type SizesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#sizesListResponse".
	Kind string `json:"kind,omitempty"`

	// Sizes: Size collection
	Sizes []*Size `json:"sizes,omitempty"`
}

type SortedDimension struct {
	// Kind: The kind of resource this is, in this case
	// dfareporting#sortedDimension.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the dimension.
	Name string `json:"name,omitempty"`

	// SortOrder: An optional sort order for the dimension column.
	SortOrder string `json:"sortOrder,omitempty"`
}

type Subaccount struct {
	// AccountId: ID of the account that contains this subaccount. This is a
	// read-only field that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AvailablePermissionIds: IDs of the available user role permissions
	// for this subaccount.
	AvailablePermissionIds googleapi.Int64s `json:"availablePermissionIds,omitempty"`

	// Id: ID of this subaccount. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#subaccount".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this subaccount. This is a required field. Must be less
	// than 128 characters long and be unique among subaccounts of the same
	// account.
	Name string `json:"name,omitempty"`
}

type SubaccountsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#subaccountsListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Subaccounts: Subaccount collection
	Subaccounts []*Subaccount `json:"subaccounts,omitempty"`
}

type TagData struct {
	// AdId: Ad associated with this placement tag.
	AdId int64 `json:"adId,omitempty,string"`

	// ClickTag: Tag string to record a click.
	ClickTag string `json:"clickTag,omitempty"`

	// CreativeId: Creative associated with this placement tag.
	CreativeId int64 `json:"creativeId,omitempty,string"`

	// Format: TagData tag format of this tag.
	Format string `json:"format,omitempty"`

	// ImpressionTag: Tag string for serving an ad.
	ImpressionTag string `json:"impressionTag,omitempty"`
}

type TagSetting struct {
	// AdditionalKeyValues: Additional key-values to be included in tags.
	// Each key-value pair must be of the form key=value, and pairs must be
	// separated by a semicolon (;). Keys and values must not contain
	// commas. For example, id=2;color=red is a valid value for this field.
	AdditionalKeyValues string `json:"additionalKeyValues,omitempty"`

	// IncludeClickThroughUrls: Whether static landing page URLs should be
	// included in the tags. This setting applies only to placements.
	IncludeClickThroughUrls bool `json:"includeClickThroughUrls,omitempty"`

	// IncludeClickTracking: Whether click-tracking string should be
	// included in the tags.
	IncludeClickTracking bool `json:"includeClickTracking,omitempty"`

	// KeywordOption: Option specifying how keywords are embedded in ad
	// tags. This setting can be used to specify whether keyword
	// placeholders are inserted in placement tags for this site. Publishers
	// can then add keywords to those placeholders.
	KeywordOption string `json:"keywordOption,omitempty"`
}

type TagSettings struct {
	// DynamicTagEnabled: Whether dynamic floodlight tags are enabled.
	DynamicTagEnabled bool `json:"dynamicTagEnabled,omitempty"`

	// ImageTagEnabled: Whether image tags are enabled.
	ImageTagEnabled bool `json:"imageTagEnabled,omitempty"`
}

type TargetWindow struct {
	// CustomHtml: User-entered value.
	CustomHtml string `json:"customHtml,omitempty"`

	// TargetWindowOption: Type of browser window for which the backup image
	// of the flash creative can be displayed.
	TargetWindowOption string `json:"targetWindowOption,omitempty"`
}

type TechnologyTargeting struct {
	// Browsers: Browsers that this ad targets. For each browser either set
	// browserVersionId or dartId along with the version numbers. If both
	// are specified, only browserVersionId will be used.The other fields
	// are populated automatically when the ad is inserted or updated.
	Browsers []*Browser `json:"browsers,omitempty"`

	// ConnectionTypes: Connection types that this ad targets. For each
	// connection type only id is required.The other fields are populated
	// automatically when the ad is inserted or updated.
	ConnectionTypes []*ConnectionType `json:"connectionTypes,omitempty"`

	// MobileCarriers: Mobile carriers that this ad targets. For each mobile
	// carrier only id is required, and the other fields are populated
	// automatically when the ad is inserted or updated. If targeting a
	// mobile carrier, do not set targeting for any zip codes.
	MobileCarriers []*MobileCarrier `json:"mobileCarriers,omitempty"`

	// OperatingSystemVersions: Operating system versions that this ad
	// targets. To target all versions, use operatingSystems. For each
	// operating system version, only id is required. The other fields are
	// populated automatically when the ad is inserted or updated. If
	// targeting an operating system version, do not set targeting for the
	// corresponding operating system in operatingSystems.
	OperatingSystemVersions []*OperatingSystemVersion `json:"operatingSystemVersions,omitempty"`

	// OperatingSystems: Operating systems that this ad targets. To target
	// specific versions, use operatingSystemVersions. For each operating
	// system only dartId is required. The other fields are populated
	// automatically when the ad is inserted or updated. If targeting an
	// operating system, do not set targeting for operating system versions
	// for the same operating system.
	OperatingSystems []*OperatingSystem `json:"operatingSystems,omitempty"`

	// PlatformTypes: Platform types that this ad targets. For example,
	// desktop, mobile, or tablet. For each platform type, only id is
	// required, and the other fields are populated automatically when the
	// ad is inserted or updated.
	PlatformTypes []*PlatformType `json:"platformTypes,omitempty"`
}

type ThirdPartyTrackingUrl struct {
	// ThirdPartyUrlType: Third-party URL type for in-stream video
	// creatives.
	ThirdPartyUrlType string `json:"thirdPartyUrlType,omitempty"`

	// Url: URL for the specified third-party URL type.
	Url string `json:"url,omitempty"`
}

type UserDefinedVariableConfiguration struct {
	// DataType: Data type for the variable. This is a required field.
	DataType string `json:"dataType,omitempty"`

	// ReportName: User-friendly name for the variable which will appear in
	// reports. This is a required field, must be less than 64 characters
	// long, and cannot contain the following characters: ""<>".
	ReportName string `json:"reportName,omitempty"`

	// VariableType: Variable name in the tag. This is a required field.
	VariableType string `json:"variableType,omitempty"`
}

type UserProfile struct {
	// AccountId: The account ID to which this profile belongs.
	AccountId int64 `json:"accountId,omitempty,string"`

	// AccountName: The account name this profile belongs to.
	AccountName string `json:"accountName,omitempty"`

	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Kind: The kind of resource this is, in this case
	// dfareporting#userProfile.
	Kind string `json:"kind,omitempty"`

	// ProfileId: The unique ID of the user profile.
	ProfileId int64 `json:"profileId,omitempty,string"`

	// SubAccountId: The sub account ID this profile belongs to if
	// applicable.
	SubAccountId int64 `json:"subAccountId,omitempty,string"`

	// SubAccountName: The sub account name this profile belongs to if
	// applicable.
	SubAccountName string `json:"subAccountName,omitempty"`

	// UserName: The user name.
	UserName string `json:"userName,omitempty"`
}

type UserProfileList struct {
	// Etag: The eTag of this response for caching purposes.
	Etag string `json:"etag,omitempty"`

	// Items: The user profiles returned in this response.
	Items []*UserProfile `json:"items,omitempty"`

	// Kind: The kind of list this is, in this case
	// dfareporting#userProfileList.
	Kind string `json:"kind,omitempty"`
}

type UserRole struct {
	// AccountId: Account ID of this user role. This is a read-only field
	// that can be left blank.
	AccountId int64 `json:"accountId,omitempty,string"`

	// DefaultUserRole: Whether this is a default user role. Default user
	// roles are created by the system for the account/subaccount and cannot
	// be modified or deleted. Each default user role comes with a basic set
	// of preassigned permissions.
	DefaultUserRole bool `json:"defaultUserRole,omitempty"`

	// Id: ID of this user role. This is a read-only, auto-generated field.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#userRole".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this user role. This is a required field. Must be less
	// than 256 characters long. If this user role is under a subaccount,
	// the name must be unique among sites of the same subaccount.
	// Otherwise, this user role is a top-level user role, and the name must
	// be unique among top-level user roles of the same account.
	Name string `json:"name,omitempty"`

	// ParentUserRoleId: ID of the user role that this user role is based on
	// or copied from. This is a required field.
	ParentUserRoleId int64 `json:"parentUserRoleId,omitempty,string"`

	// Permissions: List of permissions associated with this user role.
	Permissions []*UserRolePermission `json:"permissions,omitempty"`

	// SubaccountId: Subaccount ID of this user role. This is a read-only
	// field that can be left blank.
	SubaccountId int64 `json:"subaccountId,omitempty,string"`
}

type UserRolePermission struct {
	// Availability: Levels of availability for a user role permission.
	Availability string `json:"availability,omitempty"`

	// Id: ID of this user role permission.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#userRolePermission".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this user role permission.
	Name string `json:"name,omitempty"`

	// PermissionGroupId: ID of the permission group that this user role
	// permission belongs to.
	PermissionGroupId int64 `json:"permissionGroupId,omitempty,string"`
}

type UserRolePermissionGroup struct {
	// Id: ID of this user role permission.
	Id int64 `json:"id,omitempty,string"`

	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#userRolePermissionGroup".
	Kind string `json:"kind,omitempty"`

	// Name: Name of this user role permission group.
	Name string `json:"name,omitempty"`
}

type UserRolePermissionGroupsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#userRolePermissionGroupsListResponse".
	Kind string `json:"kind,omitempty"`

	// UserRolePermissionGroups: User role permission group collection
	UserRolePermissionGroups []*UserRolePermissionGroup `json:"userRolePermissionGroups,omitempty"`
}

type UserRolePermissionsListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#userRolePermissionsListResponse".
	Kind string `json:"kind,omitempty"`

	// UserRolePermissions: User role permission collection
	UserRolePermissions []*UserRolePermission `json:"userRolePermissions,omitempty"`
}

type UserRolesListResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "dfareporting#userRolesListResponse".
	Kind string `json:"kind,omitempty"`

	// NextPageToken: Pagination token to be used for the next list
	// operation.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// UserRoles: User role collection
	UserRoles []*UserRole `json:"userRoles,omitempty"`
}

// method id "dfareporting.accountActiveAdSummaries.get":

type AccountActiveAdSummariesGetCall struct {
	s                *Service
	profileId        int64
	summaryAccountId int64
	caller_          googleapi.Caller
	params_          url.Values
	pathTemplate_    string
}

// Get: Gets the account's active ad summary by account ID.

func (r *AccountActiveAdSummariesService) Get(profileId int64, summaryAccountId int64) *AccountActiveAdSummariesGetCall {
	return &AccountActiveAdSummariesGetCall{
		s:                r.s,
		profileId:        profileId,
		summaryAccountId: summaryAccountId,
		caller_:          googleapi.JSONCall{},
		params_:          make(map[string][]string),
		pathTemplate_:    "userprofiles/{profileId}/accountActiveAdSummaries/{summaryAccountId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountActiveAdSummariesGetCall) Fields(s ...googleapi.Field) *AccountActiveAdSummariesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountActiveAdSummariesGetCall) Do() (*AccountActiveAdSummary, error) {
	var returnValue *AccountActiveAdSummary
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":        strconv.FormatInt(c.profileId, 10),
		"summaryAccountId": strconv.FormatInt(c.summaryAccountId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets the account's active ad summary by account ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountActiveAdSummaries.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "summaryAccountId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "summaryAccountId": {
	//       "description": "Account ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountActiveAdSummaries/{summaryAccountId}",
	//   "response": {
	//     "$ref": "AccountActiveAdSummary"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountPermissionGroups.get":

type AccountPermissionGroupsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one account permission group by ID.

func (r *AccountPermissionGroupsService) Get(profileId int64, id int64) *AccountPermissionGroupsGetCall {
	return &AccountPermissionGroupsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accountPermissionGroups/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountPermissionGroupsGetCall) Fields(s ...googleapi.Field) *AccountPermissionGroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountPermissionGroupsGetCall) Do() (*AccountPermissionGroup, error) {
	var returnValue *AccountPermissionGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one account permission group by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountPermissionGroups.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Account permission group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountPermissionGroups/{id}",
	//   "response": {
	//     "$ref": "AccountPermissionGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountPermissionGroups.list":

type AccountPermissionGroupsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of account permission groups.

func (r *AccountPermissionGroupsService) List(profileId int64) *AccountPermissionGroupsListCall {
	return &AccountPermissionGroupsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accountPermissionGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountPermissionGroupsListCall) Fields(s ...googleapi.Field) *AccountPermissionGroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountPermissionGroupsListCall) Do() (*AccountPermissionGroupsListResponse, error) {
	var returnValue *AccountPermissionGroupsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the list of account permission groups.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountPermissionGroups.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountPermissionGroups",
	//   "response": {
	//     "$ref": "AccountPermissionGroupsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountPermissions.get":

type AccountPermissionsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one account permission by ID.

func (r *AccountPermissionsService) Get(profileId int64, id int64) *AccountPermissionsGetCall {
	return &AccountPermissionsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accountPermissions/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountPermissionsGetCall) Fields(s ...googleapi.Field) *AccountPermissionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountPermissionsGetCall) Do() (*AccountPermission, error) {
	var returnValue *AccountPermission
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one account permission by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountPermissions.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Account permission ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountPermissions/{id}",
	//   "response": {
	//     "$ref": "AccountPermission"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountPermissions.list":

type AccountPermissionsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of account permissions.

func (r *AccountPermissionsService) List(profileId int64) *AccountPermissionsListCall {
	return &AccountPermissionsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accountPermissions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountPermissionsListCall) Fields(s ...googleapi.Field) *AccountPermissionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountPermissionsListCall) Do() (*AccountPermissionsListResponse, error) {
	var returnValue *AccountPermissionsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the list of account permissions.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountPermissions.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountPermissions",
	//   "response": {
	//     "$ref": "AccountPermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountUserProfiles.get":

type AccountUserProfilesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one account user profile by ID.

func (r *AccountUserProfilesService) Get(profileId int64, id int64) *AccountUserProfilesGetCall {
	return &AccountUserProfilesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accountUserProfiles/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountUserProfilesGetCall) Fields(s ...googleapi.Field) *AccountUserProfilesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountUserProfilesGetCall) Do() (*AccountUserProfile, error) {
	var returnValue *AccountUserProfile
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one account user profile by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountUserProfiles.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountUserProfiles/{id}",
	//   "response": {
	//     "$ref": "AccountUserProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountUserProfiles.list":

type AccountUserProfilesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of account user profiles, possibly filtered.

func (r *AccountUserProfilesService) List(profileId int64) *AccountUserProfilesListCall {
	return &AccountUserProfilesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accountUserProfiles",
	}
}

// Active sets the optional parameter "active": Select only active user
// profiles.
func (c *AccountUserProfilesListCall) Active(active bool) *AccountUserProfilesListCall {
	c.params_.Set("active", fmt.Sprintf("%v", active))
	return c
}

// Ids sets the optional parameter "ids": Select only user profiles with
// these IDs.
func (c *AccountUserProfilesListCall) Ids(ids ...int64) *AccountUserProfilesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *AccountUserProfilesListCall) MaxResults(maxResults int64) *AccountUserProfilesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *AccountUserProfilesListCall) PageToken(pageToken string) *AccountUserProfilesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name, ID or email. Wildcards (*) are
// allowed. For example, "user profile*2015" will return objects with
// names like "user profile June 2015", "user profile April 2015" or
// simply "user profile 2015". Most of the searches also add wildcards
// implicitly at the start and the end of the search string. For
// example, a search string of "user profile" will match objects with
// name "my user profile", "user profile 2015" or simply "user profile".
func (c *AccountUserProfilesListCall) SearchString(searchString string) *AccountUserProfilesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *AccountUserProfilesListCall) SortField(sortField string) *AccountUserProfilesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *AccountUserProfilesListCall) SortOrder(sortOrder string) *AccountUserProfilesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// SubaccountId sets the optional parameter "subaccountId": Select only
// user profiles with the specified subaccount ID.
func (c *AccountUserProfilesListCall) SubaccountId(subaccountId int64) *AccountUserProfilesListCall {
	c.params_.Set("subaccountId", fmt.Sprintf("%v", subaccountId))
	return c
}

// UserRoleId sets the optional parameter "userRoleId": Select only user
// profiles with the specified user role ID.
func (c *AccountUserProfilesListCall) UserRoleId(userRoleId int64) *AccountUserProfilesListCall {
	c.params_.Set("userRoleId", fmt.Sprintf("%v", userRoleId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountUserProfilesListCall) Fields(s ...googleapi.Field) *AccountUserProfilesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountUserProfilesListCall) Do() (*AccountUserProfilesListResponse, error) {
	var returnValue *AccountUserProfilesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of account user profiles, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accountUserProfiles.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "active": {
	//       "description": "Select only active user profiles.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ids": {
	//       "description": "Select only user profiles with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name, ID or email. Wildcards (*) are allowed. For example, \"user profile*2015\" will return objects with names like \"user profile June 2015\", \"user profile April 2015\" or simply \"user profile 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"user profile\" will match objects with name \"my user profile\", \"user profile 2015\" or simply \"user profile\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "subaccountId": {
	//       "description": "Select only user profiles with the specified subaccount ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userRoleId": {
	//       "description": "Select only user profiles with the specified user role ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountUserProfiles",
	//   "response": {
	//     "$ref": "AccountUserProfilesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountUserProfiles.patch":

type AccountUserProfilesPatchCall struct {
	s                  *Service
	profileId          int64
	id                 int64
	accountuserprofile *AccountUserProfile
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Patch: Updates an existing account user profile. This method supports
// patch semantics.

func (r *AccountUserProfilesService) Patch(profileId int64, id int64, accountuserprofile *AccountUserProfile) *AccountUserProfilesPatchCall {
	return &AccountUserProfilesPatchCall{
		s:                  r.s,
		profileId:          profileId,
		id:                 id,
		accountuserprofile: accountuserprofile,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/accountUserProfiles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountUserProfilesPatchCall) Fields(s ...googleapi.Field) *AccountUserProfilesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountUserProfilesPatchCall) Do() (*AccountUserProfile, error) {
	var returnValue *AccountUserProfile
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountuserprofile,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing account user profile. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.accountUserProfiles.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User profile ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountUserProfiles",
	//   "request": {
	//     "$ref": "AccountUserProfile"
	//   },
	//   "response": {
	//     "$ref": "AccountUserProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accountUserProfiles.update":

type AccountUserProfilesUpdateCall struct {
	s                  *Service
	profileId          int64
	accountuserprofile *AccountUserProfile
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Update: Updates an existing account user profile.

func (r *AccountUserProfilesService) Update(profileId int64, accountuserprofile *AccountUserProfile) *AccountUserProfilesUpdateCall {
	return &AccountUserProfilesUpdateCall{
		s:                  r.s,
		profileId:          profileId,
		accountuserprofile: accountuserprofile,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/accountUserProfiles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountUserProfilesUpdateCall) Fields(s ...googleapi.Field) *AccountUserProfilesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountUserProfilesUpdateCall) Do() (*AccountUserProfile, error) {
	var returnValue *AccountUserProfile
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.accountuserprofile,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing account user profile.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.accountUserProfiles.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accountUserProfiles",
	//   "request": {
	//     "$ref": "AccountUserProfile"
	//   },
	//   "response": {
	//     "$ref": "AccountUserProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accounts.get":

type AccountsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one account by ID.

func (r *AccountsService) Get(profileId int64, id int64) *AccountsGetCall {
	return &AccountsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accounts/{id}",
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
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one account by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accounts.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Account ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accounts/{id}",
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accounts.list":

type AccountsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of accounts, possibly filtered.

func (r *AccountsService) List(profileId int64) *AccountsListCall {
	return &AccountsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accounts",
	}
}

// Active sets the optional parameter "active": Select only active
// accounts. Don't set this field to select both active and non-active
// accounts.
func (c *AccountsListCall) Active(active bool) *AccountsListCall {
	c.params_.Set("active", fmt.Sprintf("%v", active))
	return c
}

// Ids sets the optional parameter "ids": Select only accounts with
// these IDs.
func (c *AccountsListCall) Ids(ids ...int64) *AccountsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *AccountsListCall) MaxResults(maxResults int64) *AccountsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *AccountsListCall) PageToken(pageToken string) *AccountsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "account*2015" will return objects with names like "account
// June 2015", "account April 2015" or simply "account 2015". Most of
// the searches also add wildcards implicitly at the start and the end
// of the search string. For example, a search string of "account" will
// match objects with name "my account", "account 2015" or simply
// "account".
func (c *AccountsListCall) SearchString(searchString string) *AccountsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *AccountsListCall) SortField(sortField string) *AccountsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *AccountsListCall) SortOrder(sortOrder string) *AccountsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsListCall) Fields(s ...googleapi.Field) *AccountsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsListCall) Do() (*AccountsListResponse, error) {
	var returnValue *AccountsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the list of accounts, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.accounts.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "active": {
	//       "description": "Select only active accounts. Don't set this field to select both active and non-active accounts.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ids": {
	//       "description": "Select only accounts with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"account*2015\" will return objects with names like \"account June 2015\", \"account April 2015\" or simply \"account 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"account\" will match objects with name \"my account\", \"account 2015\" or simply \"account\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accounts",
	//   "response": {
	//     "$ref": "AccountsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accounts.patch":

type AccountsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing account. This method supports patch
// semantics.

func (r *AccountsService) Patch(profileId int64, id int64, account *Account) *AccountsPatchCall {
	return &AccountsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accounts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsPatchCall) Fields(s ...googleapi.Field) *AccountsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AccountsPatchCall) Do() (*Account, error) {
	var returnValue *Account
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.account,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing account. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.accounts.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Account ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accounts",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.accounts.update":

type AccountsUpdateCall struct {
	s             *Service
	profileId     int64
	account       *Account
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing account.

func (r *AccountsService) Update(profileId int64, account *Account) *AccountsUpdateCall {
	return &AccountsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		account:       account,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/accounts",
	}
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
		"profileId": strconv.FormatInt(c.profileId, 10),
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
	//   "description": "Updates an existing account.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.accounts.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/accounts",
	//   "request": {
	//     "$ref": "Account"
	//   },
	//   "response": {
	//     "$ref": "Account"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.ads.get":

type AdsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one ad by ID.

func (r *AdsService) Get(profileId int64, id int64) *AdsGetCall {
	return &AdsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/ads/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdsGetCall) Fields(s ...googleapi.Field) *AdsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdsGetCall) Do() (*Ad, error) {
	var returnValue *Ad
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one ad by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.ads.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Ad ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/ads/{id}",
	//   "response": {
	//     "$ref": "Ad"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.ads.insert":

type AdsInsertCall struct {
	s             *Service
	profileId     int64
	ad            *Ad
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new ad.

func (r *AdsService) Insert(profileId int64, ad *Ad) *AdsInsertCall {
	return &AdsInsertCall{
		s:             r.s,
		profileId:     profileId,
		ad:            ad,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/ads",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdsInsertCall) Fields(s ...googleapi.Field) *AdsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdsInsertCall) Do() (*Ad, error) {
	var returnValue *Ad
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.ad,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new ad.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.ads.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/ads",
	//   "request": {
	//     "$ref": "Ad"
	//   },
	//   "response": {
	//     "$ref": "Ad"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.ads.list":

type AdsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of ads, possibly filtered.

func (r *AdsService) List(profileId int64) *AdsListCall {
	return &AdsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/ads",
	}
}

// Active sets the optional parameter "active": Select only active ads.
func (c *AdsListCall) Active(active bool) *AdsListCall {
	c.params_.Set("active", fmt.Sprintf("%v", active))
	return c
}

// AdvertiserId sets the optional parameter "advertiserId": Select only
// ads with this advertiser ID.
func (c *AdsListCall) AdvertiserId(advertiserId int64) *AdsListCall {
	c.params_.Set("advertiserId", fmt.Sprintf("%v", advertiserId))
	return c
}

// Archived sets the optional parameter "archived": Select only archived
// ads.
func (c *AdsListCall) Archived(archived bool) *AdsListCall {
	c.params_.Set("archived", fmt.Sprintf("%v", archived))
	return c
}

// AudienceSegmentIds sets the optional parameter "audienceSegmentIds":
// Select only ads with these audience segment IDs.
func (c *AdsListCall) AudienceSegmentIds(audienceSegmentIds ...int64) *AdsListCall {
	for _, v_ := range audienceSegmentIds {
		c.params_.Add("audienceSegmentIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// CampaignIds sets the optional parameter "campaignIds": Select only
// ads with these campaign IDs.
func (c *AdsListCall) CampaignIds(campaignIds ...int64) *AdsListCall {
	for _, v_ := range campaignIds {
		c.params_.Add("campaignIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Compatibility sets the optional parameter "compatibility": Select
// default ads with the specified compatibility. Applicable when type is
// AD_SERVING_DEFAULT_AD. WEB and WEB_INTERSTITIAL refer to rendering
// either on desktop or on mobile devices for regular or interstitial
// ads, respectively. APP and APP_INTERSTITIAL are for rendering in
// mobile apps. IN_STREAM_VIDEO refers to rendering an in-stream video
// ads developed with the VAST standard.
func (c *AdsListCall) Compatibility(compatibility string) *AdsListCall {
	c.params_.Set("compatibility", fmt.Sprintf("%v", compatibility))
	return c
}

// CreativeIds sets the optional parameter "creativeIds": Select only
// ads with these creative IDs assigned.
func (c *AdsListCall) CreativeIds(creativeIds ...int64) *AdsListCall {
	for _, v_ := range creativeIds {
		c.params_.Add("creativeIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// CreativeOptimizationConfigurationIds sets the optional parameter
// "creativeOptimizationConfigurationIds": Select only ads with these
// creative optimization configuration IDs.
func (c *AdsListCall) CreativeOptimizationConfigurationIds(creativeOptimizationConfigurationIds ...int64) *AdsListCall {
	for _, v_ := range creativeOptimizationConfigurationIds {
		c.params_.Add("creativeOptimizationConfigurationIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// CreativeType sets the optional parameter "creativeType": Select only
// ads with the specified creativeType.
func (c *AdsListCall) CreativeType(creativeType string) *AdsListCall {
	c.params_.Set("creativeType", fmt.Sprintf("%v", creativeType))
	return c
}

// DynamicClickTracker sets the optional parameter
// "dynamicClickTracker": Select only dynamic click trackers. Applicable
// when type is AD_SERVING_CLICK_TRACKER. If true, select dynamic click
// trackers. If false, select static click trackers. Leave unset to
// select both.
func (c *AdsListCall) DynamicClickTracker(dynamicClickTracker bool) *AdsListCall {
	c.params_.Set("dynamicClickTracker", fmt.Sprintf("%v", dynamicClickTracker))
	return c
}

// Ids sets the optional parameter "ids": Select only ads with these
// IDs.
func (c *AdsListCall) Ids(ids ...int64) *AdsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// LandingPageIds sets the optional parameter "landingPageIds": Select
// only ads with these landing page IDs.
func (c *AdsListCall) LandingPageIds(landingPageIds ...int64) *AdsListCall {
	for _, v_ := range landingPageIds {
		c.params_.Add("landingPageIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *AdsListCall) MaxResults(maxResults int64) *AdsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OverriddenEventTagId sets the optional parameter
// "overriddenEventTagId": Select only ads with this event tag override
// ID.
func (c *AdsListCall) OverriddenEventTagId(overriddenEventTagId int64) *AdsListCall {
	c.params_.Set("overriddenEventTagId", fmt.Sprintf("%v", overriddenEventTagId))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *AdsListCall) PageToken(pageToken string) *AdsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// PlacementIds sets the optional parameter "placementIds": Select only
// ads with these placement IDs assigned.
func (c *AdsListCall) PlacementIds(placementIds ...int64) *AdsListCall {
	for _, v_ := range placementIds {
		c.params_.Add("placementIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// RemarketingListIds sets the optional parameter "remarketingListIds":
// Select only ads whose list targeting expression use these remarketing
// list IDs.
func (c *AdsListCall) RemarketingListIds(remarketingListIds ...int64) *AdsListCall {
	for _, v_ := range remarketingListIds {
		c.params_.Add("remarketingListIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "ad*2015" will return objects with names like "ad June
// 2015", "ad April 2015" or simply "ad 2015". Most of the searches also
// add wildcards implicitly at the start and the end of the search
// string. For example, a search string of "ad" will match objects with
// name "my ad", "ad 2015" or simply "ad".
func (c *AdsListCall) SearchString(searchString string) *AdsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SizeIds sets the optional parameter "sizeIds": Select only ads with
// these size IDs.
func (c *AdsListCall) SizeIds(sizeIds ...int64) *AdsListCall {
	for _, v_ := range sizeIds {
		c.params_.Add("sizeIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *AdsListCall) SortField(sortField string) *AdsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *AdsListCall) SortOrder(sortOrder string) *AdsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// SslCompliant sets the optional parameter "sslCompliant": Select only
// ads that are SSL-compliant.
func (c *AdsListCall) SslCompliant(sslCompliant bool) *AdsListCall {
	c.params_.Set("sslCompliant", fmt.Sprintf("%v", sslCompliant))
	return c
}

// SslRequired sets the optional parameter "sslRequired": Select only
// ads that require SSL.
func (c *AdsListCall) SslRequired(sslRequired bool) *AdsListCall {
	c.params_.Set("sslRequired", fmt.Sprintf("%v", sslRequired))
	return c
}

// Type sets the optional parameter "type": Select only ads with these
// types.
func (c *AdsListCall) Type(type_ ...string) *AdsListCall {
	c.params_["type"] = type_
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdsListCall) Fields(s ...googleapi.Field) *AdsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdsListCall) Do() (*AdsListResponse, error) {
	var returnValue *AdsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of ads, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.ads.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "active": {
	//       "description": "Select only active ads.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "advertiserId": {
	//       "description": "Select only ads with this advertiser ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "archived": {
	//       "description": "Select only archived ads.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "audienceSegmentIds": {
	//       "description": "Select only ads with these audience segment IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "campaignIds": {
	//       "description": "Select only ads with these campaign IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "compatibility": {
	//       "description": "Select default ads with the specified compatibility. Applicable when type is AD_SERVING_DEFAULT_AD. WEB and WEB_INTERSTITIAL refer to rendering either on desktop or on mobile devices for regular or interstitial ads, respectively. APP and APP_INTERSTITIAL are for rendering in mobile apps. IN_STREAM_VIDEO refers to rendering an in-stream video ads developed with the VAST standard.",
	//       "enum": [
	//         "APP",
	//         "APP_INTERSTITIAL",
	//         "IN_STREAM_VIDEO",
	//         "WEB",
	//         "WEB_INTERSTITIAL"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creativeIds": {
	//       "description": "Select only ads with these creative IDs assigned.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "creativeOptimizationConfigurationIds": {
	//       "description": "Select only ads with these creative optimization configuration IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "creativeType": {
	//       "description": "Select only ads with the specified creativeType.",
	//       "enum": [
	//         "BRAND_SAFE_DEFAULT_INSTREAM_VIDEO",
	//         "CUSTOM_INPAGE",
	//         "CUSTOM_INTERSTITIAL",
	//         "ENHANCED_BANNER",
	//         "ENHANCED_IMAGE",
	//         "FLASH_INPAGE",
	//         "HTML5_BANNER",
	//         "IMAGE",
	//         "INSTREAM_VIDEO",
	//         "INTERNAL_REDIRECT",
	//         "INTERSTITIAL_INTERNAL_REDIRECT",
	//         "REDIRECT",
	//         "RICH_MEDIA_EXPANDING",
	//         "RICH_MEDIA_IM_EXPAND",
	//         "RICH_MEDIA_INPAGE",
	//         "RICH_MEDIA_INPAGE_FLOATING",
	//         "RICH_MEDIA_INTERSTITIAL_FLOAT",
	//         "RICH_MEDIA_MOBILE_IN_APP",
	//         "RICH_MEDIA_MULTI_FLOATING",
	//         "RICH_MEDIA_PEEL_DOWN",
	//         "TRACKING_TEXT",
	//         "VPAID_LINEAR",
	//         "VPAID_NON_LINEAR"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dynamicClickTracker": {
	//       "description": "Select only dynamic click trackers. Applicable when type is AD_SERVING_CLICK_TRACKER. If true, select dynamic click trackers. If false, select static click trackers. Leave unset to select both.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ids": {
	//       "description": "Select only ads with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "landingPageIds": {
	//       "description": "Select only ads with these landing page IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "overriddenEventTagId": {
	//       "description": "Select only ads with this event tag override ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "placementIds": {
	//       "description": "Select only ads with these placement IDs assigned.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "remarketingListIds": {
	//       "description": "Select only ads whose list targeting expression use these remarketing list IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"ad*2015\" will return objects with names like \"ad June 2015\", \"ad April 2015\" or simply \"ad 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"ad\" will match objects with name \"my ad\", \"ad 2015\" or simply \"ad\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sizeIds": {
	//       "description": "Select only ads with these size IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sslCompliant": {
	//       "description": "Select only ads that are SSL-compliant.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "sslRequired": {
	//       "description": "Select only ads that require SSL.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "type": {
	//       "description": "Select only ads with these types.",
	//       "enum": [
	//         "AD_SERVING_CLICK_TRACKER",
	//         "AD_SERVING_DEFAULT_AD",
	//         "AD_SERVING_STANDARD_AD",
	//         "AD_SERVING_TRACKING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/ads",
	//   "response": {
	//     "$ref": "AdsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.ads.patch":

type AdsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	ad            *Ad
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing ad. This method supports patch semantics.

func (r *AdsService) Patch(profileId int64, id int64, ad *Ad) *AdsPatchCall {
	return &AdsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		ad:            ad,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/ads",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdsPatchCall) Fields(s ...googleapi.Field) *AdsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdsPatchCall) Do() (*Ad, error) {
	var returnValue *Ad
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.ad,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing ad. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.ads.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Ad ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/ads",
	//   "request": {
	//     "$ref": "Ad"
	//   },
	//   "response": {
	//     "$ref": "Ad"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.ads.update":

type AdsUpdateCall struct {
	s             *Service
	profileId     int64
	ad            *Ad
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing ad.

func (r *AdsService) Update(profileId int64, ad *Ad) *AdsUpdateCall {
	return &AdsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		ad:            ad,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/ads",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdsUpdateCall) Fields(s ...googleapi.Field) *AdsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdsUpdateCall) Do() (*Ad, error) {
	var returnValue *Ad
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.ad,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing ad.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.ads.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/ads",
	//   "request": {
	//     "$ref": "Ad"
	//   },
	//   "response": {
	//     "$ref": "Ad"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertiserGroups.delete":

type AdvertiserGroupsDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing advertiser group.

func (r *AdvertiserGroupsService) Delete(profileId int64, id int64) *AdvertiserGroupsDeleteCall {
	return &AdvertiserGroupsDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertiserGroups/{id}",
	}
}

func (c *AdvertiserGroupsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing advertiser group.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.advertiserGroups.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Advertiser group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertiserGroups/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertiserGroups.get":

type AdvertiserGroupsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one advertiser group by ID.

func (r *AdvertiserGroupsService) Get(profileId int64, id int64) *AdvertiserGroupsGetCall {
	return &AdvertiserGroupsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertiserGroups/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertiserGroupsGetCall) Fields(s ...googleapi.Field) *AdvertiserGroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertiserGroupsGetCall) Do() (*AdvertiserGroup, error) {
	var returnValue *AdvertiserGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one advertiser group by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.advertiserGroups.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Advertiser group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertiserGroups/{id}",
	//   "response": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertiserGroups.insert":

type AdvertiserGroupsInsertCall struct {
	s               *Service
	profileId       int64
	advertisergroup *AdvertiserGroup
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Insert: Inserts a new advertiser group.

func (r *AdvertiserGroupsService) Insert(profileId int64, advertisergroup *AdvertiserGroup) *AdvertiserGroupsInsertCall {
	return &AdvertiserGroupsInsertCall{
		s:               r.s,
		profileId:       profileId,
		advertisergroup: advertisergroup,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/advertiserGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertiserGroupsInsertCall) Fields(s ...googleapi.Field) *AdvertiserGroupsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertiserGroupsInsertCall) Do() (*AdvertiserGroup, error) {
	var returnValue *AdvertiserGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.advertisergroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new advertiser group.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.advertiserGroups.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertiserGroups",
	//   "request": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "response": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertiserGroups.list":

type AdvertiserGroupsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of advertiser groups, possibly filtered.

func (r *AdvertiserGroupsService) List(profileId int64) *AdvertiserGroupsListCall {
	return &AdvertiserGroupsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertiserGroups",
	}
}

// Ids sets the optional parameter "ids": Select only advertiser groups
// with these IDs.
func (c *AdvertiserGroupsListCall) Ids(ids ...int64) *AdvertiserGroupsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *AdvertiserGroupsListCall) MaxResults(maxResults int64) *AdvertiserGroupsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *AdvertiserGroupsListCall) PageToken(pageToken string) *AdvertiserGroupsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "advertiser*2015" will return objects with names like
// "advertiser group June 2015", "advertiser group April 2015" or simply
// "advertiser group 2015". Most of the searches also add wildcards
// implicitly at the start and the end of the search string. For
// example, a search string of "advertisergroup" will match objects with
// name "my advertisergroup", "advertisergroup 2015" or simply
// "advertisergroup".
func (c *AdvertiserGroupsListCall) SearchString(searchString string) *AdvertiserGroupsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *AdvertiserGroupsListCall) SortField(sortField string) *AdvertiserGroupsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *AdvertiserGroupsListCall) SortOrder(sortOrder string) *AdvertiserGroupsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertiserGroupsListCall) Fields(s ...googleapi.Field) *AdvertiserGroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertiserGroupsListCall) Do() (*AdvertiserGroupsListResponse, error) {
	var returnValue *AdvertiserGroupsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of advertiser groups, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.advertiserGroups.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "ids": {
	//       "description": "Select only advertiser groups with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"advertiser*2015\" will return objects with names like \"advertiser group June 2015\", \"advertiser group April 2015\" or simply \"advertiser group 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"advertisergroup\" will match objects with name \"my advertisergroup\", \"advertisergroup 2015\" or simply \"advertisergroup\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertiserGroups",
	//   "response": {
	//     "$ref": "AdvertiserGroupsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertiserGroups.patch":

type AdvertiserGroupsPatchCall struct {
	s               *Service
	profileId       int64
	id              int64
	advertisergroup *AdvertiserGroup
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Patch: Updates an existing advertiser group. This method supports
// patch semantics.

func (r *AdvertiserGroupsService) Patch(profileId int64, id int64, advertisergroup *AdvertiserGroup) *AdvertiserGroupsPatchCall {
	return &AdvertiserGroupsPatchCall{
		s:               r.s,
		profileId:       profileId,
		id:              id,
		advertisergroup: advertisergroup,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/advertiserGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertiserGroupsPatchCall) Fields(s ...googleapi.Field) *AdvertiserGroupsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertiserGroupsPatchCall) Do() (*AdvertiserGroup, error) {
	var returnValue *AdvertiserGroup
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.advertisergroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing advertiser group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.advertiserGroups.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Advertiser group ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertiserGroups",
	//   "request": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "response": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertiserGroups.update":

type AdvertiserGroupsUpdateCall struct {
	s               *Service
	profileId       int64
	advertisergroup *AdvertiserGroup
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Update: Updates an existing advertiser group.

func (r *AdvertiserGroupsService) Update(profileId int64, advertisergroup *AdvertiserGroup) *AdvertiserGroupsUpdateCall {
	return &AdvertiserGroupsUpdateCall{
		s:               r.s,
		profileId:       profileId,
		advertisergroup: advertisergroup,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/advertiserGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertiserGroupsUpdateCall) Fields(s ...googleapi.Field) *AdvertiserGroupsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertiserGroupsUpdateCall) Do() (*AdvertiserGroup, error) {
	var returnValue *AdvertiserGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.advertisergroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing advertiser group.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.advertiserGroups.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertiserGroups",
	//   "request": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "response": {
	//     "$ref": "AdvertiserGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertisers.get":

type AdvertisersGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one advertiser by ID.

func (r *AdvertisersService) Get(profileId int64, id int64) *AdvertisersGetCall {
	return &AdvertisersGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertisers/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertisersGetCall) Fields(s ...googleapi.Field) *AdvertisersGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertisersGetCall) Do() (*Advertiser, error) {
	var returnValue *Advertiser
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one advertiser by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.advertisers.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Advertiser ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertisers/{id}",
	//   "response": {
	//     "$ref": "Advertiser"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertisers.insert":

type AdvertisersInsertCall struct {
	s             *Service
	profileId     int64
	advertiser    *Advertiser
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new advertiser.

func (r *AdvertisersService) Insert(profileId int64, advertiser *Advertiser) *AdvertisersInsertCall {
	return &AdvertisersInsertCall{
		s:             r.s,
		profileId:     profileId,
		advertiser:    advertiser,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertisers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertisersInsertCall) Fields(s ...googleapi.Field) *AdvertisersInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertisersInsertCall) Do() (*Advertiser, error) {
	var returnValue *Advertiser
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.advertiser,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new advertiser.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.advertisers.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertisers",
	//   "request": {
	//     "$ref": "Advertiser"
	//   },
	//   "response": {
	//     "$ref": "Advertiser"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertisers.list":

type AdvertisersListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of advertisers, possibly filtered.

func (r *AdvertisersService) List(profileId int64) *AdvertisersListCall {
	return &AdvertisersListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertisers",
	}
}

// AdvertiserGroupIds sets the optional parameter "advertiserGroupIds":
// Select only advertisers with these advertiser group IDs.
func (c *AdvertisersListCall) AdvertiserGroupIds(advertiserGroupIds ...int64) *AdvertisersListCall {
	for _, v_ := range advertiserGroupIds {
		c.params_.Add("advertiserGroupIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// FloodlightConfigurationIds sets the optional parameter
// "floodlightConfigurationIds": Select only advertisers with these
// floodlight configuration IDs.
func (c *AdvertisersListCall) FloodlightConfigurationIds(floodlightConfigurationIds ...int64) *AdvertisersListCall {
	for _, v_ := range floodlightConfigurationIds {
		c.params_.Add("floodlightConfigurationIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only advertisers with
// these IDs.
func (c *AdvertisersListCall) Ids(ids ...int64) *AdvertisersListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// IncludeAdvertisersWithoutGroupsOnly sets the optional parameter
// "includeAdvertisersWithoutGroupsOnly": Select only advertisers which
// do not belong to any advertiser group.
func (c *AdvertisersListCall) IncludeAdvertisersWithoutGroupsOnly(includeAdvertisersWithoutGroupsOnly bool) *AdvertisersListCall {
	c.params_.Set("includeAdvertisersWithoutGroupsOnly", fmt.Sprintf("%v", includeAdvertisersWithoutGroupsOnly))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *AdvertisersListCall) MaxResults(maxResults int64) *AdvertisersListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OnlyParent sets the optional parameter "onlyParent": Select only
// advertisers which use another advertiser's floodlight configuration.
func (c *AdvertisersListCall) OnlyParent(onlyParent bool) *AdvertisersListCall {
	c.params_.Set("onlyParent", fmt.Sprintf("%v", onlyParent))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *AdvertisersListCall) PageToken(pageToken string) *AdvertisersListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "advertiser*2015" will return objects with names like
// "advertiser June 2015", "advertiser April 2015" or simply "advertiser
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "advertiser" will match objects with name "my advertiser",
// "advertiser 2015" or simply "advertiser".
func (c *AdvertisersListCall) SearchString(searchString string) *AdvertisersListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *AdvertisersListCall) SortField(sortField string) *AdvertisersListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *AdvertisersListCall) SortOrder(sortOrder string) *AdvertisersListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Status sets the optional parameter "status": Select only advertisers
// with the specified status.
func (c *AdvertisersListCall) Status(status string) *AdvertisersListCall {
	c.params_.Set("status", fmt.Sprintf("%v", status))
	return c
}

// SubaccountId sets the optional parameter "subaccountId": Select only
// advertisers with these subaccount IDs.
func (c *AdvertisersListCall) SubaccountId(subaccountId int64) *AdvertisersListCall {
	c.params_.Set("subaccountId", fmt.Sprintf("%v", subaccountId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertisersListCall) Fields(s ...googleapi.Field) *AdvertisersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertisersListCall) Do() (*AdvertisersListResponse, error) {
	var returnValue *AdvertisersListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of advertisers, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.advertisers.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserGroupIds": {
	//       "description": "Select only advertisers with these advertiser group IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "floodlightConfigurationIds": {
	//       "description": "Select only advertisers with these floodlight configuration IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only advertisers with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "includeAdvertisersWithoutGroupsOnly": {
	//       "description": "Select only advertisers which do not belong to any advertiser group.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "onlyParent": {
	//       "description": "Select only advertisers which use another advertiser's floodlight configuration.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"advertiser*2015\" will return objects with names like \"advertiser June 2015\", \"advertiser April 2015\" or simply \"advertiser 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"advertiser\" will match objects with name \"my advertiser\", \"advertiser 2015\" or simply \"advertiser\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "status": {
	//       "description": "Select only advertisers with the specified status.",
	//       "enum": [
	//         "APPROVED",
	//         "ON_HOLD"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "subaccountId": {
	//       "description": "Select only advertisers with these subaccount IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertisers",
	//   "response": {
	//     "$ref": "AdvertisersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertisers.patch":

type AdvertisersPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	advertiser    *Advertiser
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing advertiser. This method supports patch
// semantics.

func (r *AdvertisersService) Patch(profileId int64, id int64, advertiser *Advertiser) *AdvertisersPatchCall {
	return &AdvertisersPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		advertiser:    advertiser,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertisers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertisersPatchCall) Fields(s ...googleapi.Field) *AdvertisersPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertisersPatchCall) Do() (*Advertiser, error) {
	var returnValue *Advertiser
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.advertiser,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing advertiser. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.advertisers.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Advertiser ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertisers",
	//   "request": {
	//     "$ref": "Advertiser"
	//   },
	//   "response": {
	//     "$ref": "Advertiser"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.advertisers.update":

type AdvertisersUpdateCall struct {
	s             *Service
	profileId     int64
	advertiser    *Advertiser
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing advertiser.

func (r *AdvertisersService) Update(profileId int64, advertiser *Advertiser) *AdvertisersUpdateCall {
	return &AdvertisersUpdateCall{
		s:             r.s,
		profileId:     profileId,
		advertiser:    advertiser,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/advertisers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AdvertisersUpdateCall) Fields(s ...googleapi.Field) *AdvertisersUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *AdvertisersUpdateCall) Do() (*Advertiser, error) {
	var returnValue *Advertiser
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.advertiser,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing advertiser.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.advertisers.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/advertisers",
	//   "request": {
	//     "$ref": "Advertiser"
	//   },
	//   "response": {
	//     "$ref": "Advertiser"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.browsers.list":

type BrowsersListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of browsers.

func (r *BrowsersService) List(profileId int64) *BrowsersListCall {
	return &BrowsersListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/browsers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BrowsersListCall) Fields(s ...googleapi.Field) *BrowsersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *BrowsersListCall) Do() (*BrowsersListResponse, error) {
	var returnValue *BrowsersListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of browsers.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.browsers.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/browsers",
	//   "response": {
	//     "$ref": "BrowsersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaignCreativeAssociations.insert":

type CampaignCreativeAssociationsInsertCall struct {
	s                           *Service
	profileId                   int64
	campaignId                  int64
	campaigncreativeassociation *CampaignCreativeAssociation
	caller_                     googleapi.Caller
	params_                     url.Values
	pathTemplate_               string
}

// Insert: Associates a creative with the specified campaign. This
// method creates a default ad with dimensions matching the creative in
// the campaign if such a default ad does not exist already.

func (r *CampaignCreativeAssociationsService) Insert(profileId int64, campaignId int64, campaigncreativeassociation *CampaignCreativeAssociation) *CampaignCreativeAssociationsInsertCall {
	return &CampaignCreativeAssociationsInsertCall{
		s:                           r.s,
		profileId:                   profileId,
		campaignId:                  campaignId,
		campaigncreativeassociation: campaigncreativeassociation,
		caller_:                     googleapi.JSONCall{},
		params_:                     make(map[string][]string),
		pathTemplate_:               "userprofiles/{profileId}/campaigns/{campaignId}/campaignCreativeAssociations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignCreativeAssociationsInsertCall) Fields(s ...googleapi.Field) *CampaignCreativeAssociationsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignCreativeAssociationsInsertCall) Do() (*CampaignCreativeAssociation, error) {
	var returnValue *CampaignCreativeAssociation
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.campaigncreativeassociation,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Associates a creative with the specified campaign. This method creates a default ad with dimensions matching the creative in the campaign if such a default ad does not exist already.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.campaignCreativeAssociations.insert",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Campaign ID in this association.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/campaignCreativeAssociations",
	//   "request": {
	//     "$ref": "CampaignCreativeAssociation"
	//   },
	//   "response": {
	//     "$ref": "CampaignCreativeAssociation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaignCreativeAssociations.list":

type CampaignCreativeAssociationsListCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of creative IDs associated with the
// specified campaign.

func (r *CampaignCreativeAssociationsService) List(profileId int64, campaignId int64) *CampaignCreativeAssociationsListCall {
	return &CampaignCreativeAssociationsListCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/campaignCreativeAssociations",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *CampaignCreativeAssociationsListCall) MaxResults(maxResults int64) *CampaignCreativeAssociationsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *CampaignCreativeAssociationsListCall) PageToken(pageToken string) *CampaignCreativeAssociationsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *CampaignCreativeAssociationsListCall) SortOrder(sortOrder string) *CampaignCreativeAssociationsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignCreativeAssociationsListCall) Fields(s ...googleapi.Field) *CampaignCreativeAssociationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignCreativeAssociationsListCall) Do() (*CampaignCreativeAssociationsListResponse, error) {
	var returnValue *CampaignCreativeAssociationsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the list of creative IDs associated with the specified campaign.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.campaignCreativeAssociations.list",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Campaign ID in this association.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/campaignCreativeAssociations",
	//   "response": {
	//     "$ref": "CampaignCreativeAssociationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaigns.get":

type CampaignsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one campaign by ID.

func (r *CampaignsService) Get(profileId int64, id int64) *CampaignsGetCall {
	return &CampaignsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsGetCall) Fields(s ...googleapi.Field) *CampaignsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignsGetCall) Do() (*Campaign, error) {
	var returnValue *Campaign
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one campaign by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.campaigns.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{id}",
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaigns.insert":

type CampaignsInsertCall struct {
	s                      *Service
	profileId              int64
	defaultLandingPageName string
	defaultLandingPageUrl  string
	campaign               *Campaign
	caller_                googleapi.Caller
	params_                url.Values
	pathTemplate_          string
}

// Insert: Inserts a new campaign.

func (r *CampaignsService) Insert(profileId int64, defaultLandingPageName string, defaultLandingPageUrl string, campaign *Campaign) *CampaignsInsertCall {
	return &CampaignsInsertCall{
		s:                      r.s,
		profileId:              profileId,
		defaultLandingPageName: defaultLandingPageName,
		defaultLandingPageUrl:  defaultLandingPageUrl,
		campaign:               campaign,
		caller_:                googleapi.JSONCall{},
		params_:                make(map[string][]string),
		pathTemplate_:          "userprofiles/{profileId}/campaigns",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsInsertCall) Fields(s ...googleapi.Field) *CampaignsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignsInsertCall) Do() (*Campaign, error) {
	var returnValue *Campaign
	c.params_.Set("defaultLandingPageName", fmt.Sprintf("%v", c.defaultLandingPageName))
	c.params_.Set("defaultLandingPageUrl", fmt.Sprintf("%v", c.defaultLandingPageUrl))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.campaign,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new campaign.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.campaigns.insert",
	//   "parameterOrder": [
	//     "profileId",
	//     "defaultLandingPageName",
	//     "defaultLandingPageUrl"
	//   ],
	//   "parameters": {
	//     "defaultLandingPageName": {
	//       "description": "Default landing page name for this new campaign. Must be less than 256 characters long.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "defaultLandingPageUrl": {
	//       "description": "Default landing page URL for this new campaign.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns",
	//   "request": {
	//     "$ref": "Campaign"
	//   },
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaigns.list":

type CampaignsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of campaigns, possibly filtered.

func (r *CampaignsService) List(profileId int64) *CampaignsListCall {
	return &CampaignsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns",
	}
}

// AdvertiserGroupIds sets the optional parameter "advertiserGroupIds":
// Select only campaigns whose advertisers belong to these advertiser
// groups.
func (c *CampaignsListCall) AdvertiserGroupIds(advertiserGroupIds ...int64) *CampaignsListCall {
	for _, v_ := range advertiserGroupIds {
		c.params_.Add("advertiserGroupIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// AdvertiserIds sets the optional parameter "advertiserIds": Select
// only campaigns that belong to these advertisers.
func (c *CampaignsListCall) AdvertiserIds(advertiserIds ...int64) *CampaignsListCall {
	for _, v_ := range advertiserIds {
		c.params_.Add("advertiserIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Archived sets the optional parameter "archived": Select only archived
// campaigns. Don't set this field to select both archived and
// non-archived campaigns.
func (c *CampaignsListCall) Archived(archived bool) *CampaignsListCall {
	c.params_.Set("archived", fmt.Sprintf("%v", archived))
	return c
}

// AtLeastOneOptimizationActivity sets the optional parameter
// "atLeastOneOptimizationActivity": Select only campaigns that have at
// least one optimization activity.
func (c *CampaignsListCall) AtLeastOneOptimizationActivity(atLeastOneOptimizationActivity bool) *CampaignsListCall {
	c.params_.Set("atLeastOneOptimizationActivity", fmt.Sprintf("%v", atLeastOneOptimizationActivity))
	return c
}

// ExcludedIds sets the optional parameter "excludedIds": Exclude
// campaigns with these IDs.
func (c *CampaignsListCall) ExcludedIds(excludedIds ...int64) *CampaignsListCall {
	for _, v_ := range excludedIds {
		c.params_.Add("excludedIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only campaigns with
// these IDs.
func (c *CampaignsListCall) Ids(ids ...int64) *CampaignsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *CampaignsListCall) MaxResults(maxResults int64) *CampaignsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OverriddenEventTagId sets the optional parameter
// "overriddenEventTagId": Select only campaigns that have overridden
// this event tag ID.
func (c *CampaignsListCall) OverriddenEventTagId(overriddenEventTagId int64) *CampaignsListCall {
	c.params_.Set("overriddenEventTagId", fmt.Sprintf("%v", overriddenEventTagId))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *CampaignsListCall) PageToken(pageToken string) *CampaignsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for campaigns by name or ID. Wildcards (*) are allowed. For
// example, "campaign*2015" will return campaigns with names like
// "campaign June 2015", "campaign April 2015" or simply "campaign
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "campaign" will match campaigns with name "my campaign", "campaign
// 2015" or simply "campaign".
func (c *CampaignsListCall) SearchString(searchString string) *CampaignsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *CampaignsListCall) SortField(sortField string) *CampaignsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *CampaignsListCall) SortOrder(sortOrder string) *CampaignsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// SubaccountId sets the optional parameter "subaccountId": Select only
// campaigns that belong to this subaccount.
func (c *CampaignsListCall) SubaccountId(subaccountId int64) *CampaignsListCall {
	c.params_.Set("subaccountId", fmt.Sprintf("%v", subaccountId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsListCall) Fields(s ...googleapi.Field) *CampaignsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignsListCall) Do() (*CampaignsListResponse, error) {
	var returnValue *CampaignsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of campaigns, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.campaigns.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserGroupIds": {
	//       "description": "Select only campaigns whose advertisers belong to these advertiser groups.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "advertiserIds": {
	//       "description": "Select only campaigns that belong to these advertisers.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "archived": {
	//       "description": "Select only archived campaigns. Don't set this field to select both archived and non-archived campaigns.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "atLeastOneOptimizationActivity": {
	//       "description": "Select only campaigns that have at least one optimization activity.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "excludedIds": {
	//       "description": "Exclude campaigns with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only campaigns with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "overriddenEventTagId": {
	//       "description": "Select only campaigns that have overridden this event tag ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for campaigns by name or ID. Wildcards (*) are allowed. For example, \"campaign*2015\" will return campaigns with names like \"campaign June 2015\", \"campaign April 2015\" or simply \"campaign 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"campaign\" will match campaigns with name \"my campaign\", \"campaign 2015\" or simply \"campaign\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "subaccountId": {
	//       "description": "Select only campaigns that belong to this subaccount.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns",
	//   "response": {
	//     "$ref": "CampaignsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaigns.patch":

type CampaignsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	campaign      *Campaign
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing campaign. This method supports patch
// semantics.

func (r *CampaignsService) Patch(profileId int64, id int64, campaign *Campaign) *CampaignsPatchCall {
	return &CampaignsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		campaign:      campaign,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsPatchCall) Fields(s ...googleapi.Field) *CampaignsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignsPatchCall) Do() (*Campaign, error) {
	var returnValue *Campaign
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.campaign,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing campaign. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.campaigns.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Campaign ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns",
	//   "request": {
	//     "$ref": "Campaign"
	//   },
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.campaigns.update":

type CampaignsUpdateCall struct {
	s             *Service
	profileId     int64
	campaign      *Campaign
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing campaign.

func (r *CampaignsService) Update(profileId int64, campaign *Campaign) *CampaignsUpdateCall {
	return &CampaignsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		campaign:      campaign,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CampaignsUpdateCall) Fields(s ...googleapi.Field) *CampaignsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CampaignsUpdateCall) Do() (*Campaign, error) {
	var returnValue *Campaign
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.campaign,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing campaign.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.campaigns.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns",
	//   "request": {
	//     "$ref": "Campaign"
	//   },
	//   "response": {
	//     "$ref": "Campaign"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.changeLogs.get":

type ChangeLogsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one change log by ID.

func (r *ChangeLogsService) Get(profileId int64, id int64) *ChangeLogsGetCall {
	return &ChangeLogsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/changeLogs/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChangeLogsGetCall) Fields(s ...googleapi.Field) *ChangeLogsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChangeLogsGetCall) Do() (*ChangeLog, error) {
	var returnValue *ChangeLog
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one change log by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.changeLogs.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Change log ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/changeLogs/{id}",
	//   "response": {
	//     "$ref": "ChangeLog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.changeLogs.list":

type ChangeLogsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of change logs.

func (r *ChangeLogsService) List(profileId int64) *ChangeLogsListCall {
	return &ChangeLogsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/changeLogs",
	}
}

// Action sets the optional parameter "action": Select only change logs
// with the specified action.
func (c *ChangeLogsListCall) Action(action string) *ChangeLogsListCall {
	c.params_.Set("action", fmt.Sprintf("%v", action))
	return c
}

// Ids sets the optional parameter "ids": Select only change logs with
// these IDs.
func (c *ChangeLogsListCall) Ids(ids ...int64) *ChangeLogsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxChangeTime sets the optional parameter "maxChangeTime": Select
// only change logs whose change time is before the specified
// maxChangeTime.The time should be formatted as an RFC3339 date/time
// string. For example, for 10:54 PM on July 18th, 2015, in the
// America/New York time zone, the format is
// "2015-07-18T22:54:00-04:00". In other words, the year, month, day,
// the letter T, the hour (24-hour clock system), minute, second, and
// then the time zone offset.
func (c *ChangeLogsListCall) MaxChangeTime(maxChangeTime string) *ChangeLogsListCall {
	c.params_.Set("maxChangeTime", fmt.Sprintf("%v", maxChangeTime))
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *ChangeLogsListCall) MaxResults(maxResults int64) *ChangeLogsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// MinChangeTime sets the optional parameter "minChangeTime": Select
// only change logs whose change time is before the specified
// minChangeTime.The time should be formatted as an RFC3339 date/time
// string. For example, for 10:54 PM on July 18th, 2015, in the
// America/New York time zone, the format is
// "2015-07-18T22:54:00-04:00". In other words, the year, month, day,
// the letter T, the hour (24-hour clock system), minute, second, and
// then the time zone offset.
func (c *ChangeLogsListCall) MinChangeTime(minChangeTime string) *ChangeLogsListCall {
	c.params_.Set("minChangeTime", fmt.Sprintf("%v", minChangeTime))
	return c
}

// ObjectIds sets the optional parameter "objectIds": Select only change
// logs with these object IDs.
func (c *ChangeLogsListCall) ObjectIds(objectIds ...int64) *ChangeLogsListCall {
	for _, v_ := range objectIds {
		c.params_.Add("objectIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// ObjectType sets the optional parameter "objectType": Select only
// change logs with the specified object type.
func (c *ChangeLogsListCall) ObjectType(objectType string) *ChangeLogsListCall {
	c.params_.Set("objectType", fmt.Sprintf("%v", objectType))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *ChangeLogsListCall) PageToken(pageToken string) *ChangeLogsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Select only
// change logs whose object ID, user name, old or new values match the
// search string.
func (c *ChangeLogsListCall) SearchString(searchString string) *ChangeLogsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// UserProfileIds sets the optional parameter "userProfileIds": Select
// only change logs with these user profile IDs.
func (c *ChangeLogsListCall) UserProfileIds(userProfileIds ...int64) *ChangeLogsListCall {
	for _, v_ := range userProfileIds {
		c.params_.Add("userProfileIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChangeLogsListCall) Fields(s ...googleapi.Field) *ChangeLogsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ChangeLogsListCall) Do() (*ChangeLogsListResponse, error) {
	var returnValue *ChangeLogsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of change logs.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.changeLogs.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "action": {
	//       "description": "Select only change logs with the specified action.",
	//       "enum": [
	//         "ACTION_ADD",
	//         "ACTION_ASSIGN",
	//         "ACTION_ASSOCIATE",
	//         "ACTION_CREATE",
	//         "ACTION_DELETE",
	//         "ACTION_DISABLE",
	//         "ACTION_EMAIL_TAGS",
	//         "ACTION_ENABLE",
	//         "ACTION_LINK",
	//         "ACTION_MARK_AS_DEFAULT",
	//         "ACTION_PUSH",
	//         "ACTION_REMOVE",
	//         "ACTION_SEND",
	//         "ACTION_UNASSIGN",
	//         "ACTION_UNLINK",
	//         "ACTION_UPDATE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only change logs with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxChangeTime": {
	//       "description": "Select only change logs whose change time is before the specified maxChangeTime.The time should be formatted as an RFC3339 date/time string. For example, for 10:54 PM on July 18th, 2015, in the America/New York time zone, the format is \"2015-07-18T22:54:00-04:00\". In other words, the year, month, day, the letter T, the hour (24-hour clock system), minute, second, and then the time zone offset.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "minChangeTime": {
	//       "description": "Select only change logs whose change time is before the specified minChangeTime.The time should be formatted as an RFC3339 date/time string. For example, for 10:54 PM on July 18th, 2015, in the America/New York time zone, the format is \"2015-07-18T22:54:00-04:00\". In other words, the year, month, day, the letter T, the hour (24-hour clock system), minute, second, and then the time zone offset.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "objectIds": {
	//       "description": "Select only change logs with these object IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "objectType": {
	//       "description": "Select only change logs with the specified object type.",
	//       "enum": [
	//         "OBJECT_ACCOUNT",
	//         "OBJECT_ACCOUNT_BILLING_FEATURE",
	//         "OBJECT_AD",
	//         "OBJECT_ADVERTISER",
	//         "OBJECT_ADVERTISER_GROUP",
	//         "OBJECT_BILLING_ACCOUNT_GROUP",
	//         "OBJECT_BILLING_FEATURE",
	//         "OBJECT_BILLING_MINIMUM_FEE",
	//         "OBJECT_BILLING_PROFILE",
	//         "OBJECT_CAMPAIGN",
	//         "OBJECT_CONTENT_CATEGORY",
	//         "OBJECT_CREATIVE",
	//         "OBJECT_CREATIVE_ASSET",
	//         "OBJECT_CREATIVE_BUNDLE",
	//         "OBJECT_CREATIVE_FIELD",
	//         "OBJECT_CREATIVE_GROUP",
	//         "OBJECT_DFA_SITE",
	//         "OBJECT_EVENT_TAG",
	//         "OBJECT_FLOODLIGHT_ACTIVITY_GROUP",
	//         "OBJECT_FLOODLIGHT_ACTVITY",
	//         "OBJECT_FLOODLIGHT_CONFIGURATION",
	//         "OBJECT_INSTREAM_CREATIVE",
	//         "OBJECT_LANDING_PAGE",
	//         "OBJECT_MEDIA_ORDER",
	//         "OBJECT_PLACEMENT",
	//         "OBJECT_PLACEMENT_STRATEGY",
	//         "OBJECT_PROVIDED_LIST_CLIENT",
	//         "OBJECT_RATE_CARD",
	//         "OBJECT_REMARKETING_LIST",
	//         "OBJECT_RICHMEDIA_CREATIVE",
	//         "OBJECT_SD_SITE",
	//         "OBJECT_SIZE",
	//         "OBJECT_SUBACCOUNT",
	//         "OBJECT_USER_PROFILE",
	//         "OBJECT_USER_PROFILE_FILTER",
	//         "OBJECT_USER_ROLE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Select only change logs whose object ID, user name, old or new values match the search string.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userProfileIds": {
	//       "description": "Select only change logs with these user profile IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/changeLogs",
	//   "response": {
	//     "$ref": "ChangeLogsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.cities.list":

type CitiesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of cities, possibly filtered.

func (r *CitiesService) List(profileId int64) *CitiesListCall {
	return &CitiesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/cities",
	}
}

// CountryDartIds sets the optional parameter "countryDartIds": Select
// only cities from these countries.
func (c *CitiesListCall) CountryDartIds(countryDartIds ...int64) *CitiesListCall {
	for _, v_ := range countryDartIds {
		c.params_.Add("countryDartIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// DartIds sets the optional parameter "dartIds": Select only cities
// with these DART IDs.
func (c *CitiesListCall) DartIds(dartIds ...int64) *CitiesListCall {
	for _, v_ := range dartIds {
		c.params_.Add("dartIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// NamePrefix sets the optional parameter "namePrefix": Select only
// cities with names starting with this prefix.
func (c *CitiesListCall) NamePrefix(namePrefix string) *CitiesListCall {
	c.params_.Set("namePrefix", fmt.Sprintf("%v", namePrefix))
	return c
}

// RegionDartIds sets the optional parameter "regionDartIds": Select
// only cities from these regions.
func (c *CitiesListCall) RegionDartIds(regionDartIds ...int64) *CitiesListCall {
	for _, v_ := range regionDartIds {
		c.params_.Add("regionDartIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CitiesListCall) Fields(s ...googleapi.Field) *CitiesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CitiesListCall) Do() (*CitiesListResponse, error) {
	var returnValue *CitiesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of cities, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.cities.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "countryDartIds": {
	//       "description": "Select only cities from these countries.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "dartIds": {
	//       "description": "Select only cities with these DART IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "namePrefix": {
	//       "description": "Select only cities with names starting with this prefix.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "regionDartIds": {
	//       "description": "Select only cities from these regions.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/cities",
	//   "response": {
	//     "$ref": "CitiesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.connectionTypes.list":

type ConnectionTypesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of connection types.

func (r *ConnectionTypesService) List(profileId int64) *ConnectionTypesListCall {
	return &ConnectionTypesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/connectionTypes",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ConnectionTypesListCall) Fields(s ...googleapi.Field) *ConnectionTypesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ConnectionTypesListCall) Do() (*ConnectionTypesListResponse, error) {
	var returnValue *ConnectionTypesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of connection types.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.connectionTypes.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/connectionTypes",
	//   "response": {
	//     "$ref": "ConnectionTypesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.contentCategories.delete":

type ContentCategoriesDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing content category.

func (r *ContentCategoriesService) Delete(profileId int64, id int64) *ContentCategoriesDeleteCall {
	return &ContentCategoriesDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/contentCategories/{id}",
	}
}

func (c *ContentCategoriesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing content category.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.contentCategories.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Content category ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/contentCategories/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.contentCategories.get":

type ContentCategoriesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one content category by ID.

func (r *ContentCategoriesService) Get(profileId int64, id int64) *ContentCategoriesGetCall {
	return &ContentCategoriesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/contentCategories/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentCategoriesGetCall) Fields(s ...googleapi.Field) *ContentCategoriesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContentCategoriesGetCall) Do() (*ContentCategory, error) {
	var returnValue *ContentCategory
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one content category by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.contentCategories.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Content category ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/contentCategories/{id}",
	//   "response": {
	//     "$ref": "ContentCategory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.contentCategories.insert":

type ContentCategoriesInsertCall struct {
	s               *Service
	profileId       int64
	contentcategory *ContentCategory
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Insert: Inserts a new content category.

func (r *ContentCategoriesService) Insert(profileId int64, contentcategory *ContentCategory) *ContentCategoriesInsertCall {
	return &ContentCategoriesInsertCall{
		s:               r.s,
		profileId:       profileId,
		contentcategory: contentcategory,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/contentCategories",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentCategoriesInsertCall) Fields(s ...googleapi.Field) *ContentCategoriesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContentCategoriesInsertCall) Do() (*ContentCategory, error) {
	var returnValue *ContentCategory
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.contentcategory,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new content category.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.contentCategories.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/contentCategories",
	//   "request": {
	//     "$ref": "ContentCategory"
	//   },
	//   "response": {
	//     "$ref": "ContentCategory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.contentCategories.list":

type ContentCategoriesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of content categories, possibly filtered.

func (r *ContentCategoriesService) List(profileId int64) *ContentCategoriesListCall {
	return &ContentCategoriesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/contentCategories",
	}
}

// Ids sets the optional parameter "ids": Select only content categories
// with these IDs.
func (c *ContentCategoriesListCall) Ids(ids ...int64) *ContentCategoriesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *ContentCategoriesListCall) MaxResults(maxResults int64) *ContentCategoriesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *ContentCategoriesListCall) PageToken(pageToken string) *ContentCategoriesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "contentcategory*2015" will return objects with names like
// "contentcategory June 2015", "contentcategory April 2015" or simply
// "contentcategory 2015". Most of the searches also add wildcards
// implicitly at the start and the end of the search string. For
// example, a search string of "contentcategory" will match objects with
// name "my contentcategory", "contentcategory 2015" or simply
// "contentcategory".
func (c *ContentCategoriesListCall) SearchString(searchString string) *ContentCategoriesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *ContentCategoriesListCall) SortField(sortField string) *ContentCategoriesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *ContentCategoriesListCall) SortOrder(sortOrder string) *ContentCategoriesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentCategoriesListCall) Fields(s ...googleapi.Field) *ContentCategoriesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContentCategoriesListCall) Do() (*ContentCategoriesListResponse, error) {
	var returnValue *ContentCategoriesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of content categories, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.contentCategories.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "ids": {
	//       "description": "Select only content categories with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"contentcategory*2015\" will return objects with names like \"contentcategory June 2015\", \"contentcategory April 2015\" or simply \"contentcategory 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"contentcategory\" will match objects with name \"my contentcategory\", \"contentcategory 2015\" or simply \"contentcategory\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/contentCategories",
	//   "response": {
	//     "$ref": "ContentCategoriesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.contentCategories.patch":

type ContentCategoriesPatchCall struct {
	s               *Service
	profileId       int64
	id              int64
	contentcategory *ContentCategory
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Patch: Updates an existing content category. This method supports
// patch semantics.

func (r *ContentCategoriesService) Patch(profileId int64, id int64, contentcategory *ContentCategory) *ContentCategoriesPatchCall {
	return &ContentCategoriesPatchCall{
		s:               r.s,
		profileId:       profileId,
		id:              id,
		contentcategory: contentcategory,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/contentCategories",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentCategoriesPatchCall) Fields(s ...googleapi.Field) *ContentCategoriesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContentCategoriesPatchCall) Do() (*ContentCategory, error) {
	var returnValue *ContentCategory
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.contentcategory,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing content category. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.contentCategories.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Content category ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/contentCategories",
	//   "request": {
	//     "$ref": "ContentCategory"
	//   },
	//   "response": {
	//     "$ref": "ContentCategory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.contentCategories.update":

type ContentCategoriesUpdateCall struct {
	s               *Service
	profileId       int64
	contentcategory *ContentCategory
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Update: Updates an existing content category.

func (r *ContentCategoriesService) Update(profileId int64, contentcategory *ContentCategory) *ContentCategoriesUpdateCall {
	return &ContentCategoriesUpdateCall{
		s:               r.s,
		profileId:       profileId,
		contentcategory: contentcategory,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/contentCategories",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ContentCategoriesUpdateCall) Fields(s ...googleapi.Field) *ContentCategoriesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ContentCategoriesUpdateCall) Do() (*ContentCategory, error) {
	var returnValue *ContentCategory
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.contentcategory,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing content category.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.contentCategories.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/contentCategories",
	//   "request": {
	//     "$ref": "ContentCategory"
	//   },
	//   "response": {
	//     "$ref": "ContentCategory"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.countries.get":

type CountriesGetCall struct {
	s             *Service
	profileId     int64
	dartId        int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one country by ID.

func (r *CountriesService) Get(profileId int64, dartId int64) *CountriesGetCall {
	return &CountriesGetCall{
		s:             r.s,
		profileId:     profileId,
		dartId:        dartId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/countries/{dartId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CountriesGetCall) Fields(s ...googleapi.Field) *CountriesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CountriesGetCall) Do() (*Country, error) {
	var returnValue *Country
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"dartId":    strconv.FormatInt(c.dartId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one country by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.countries.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "dartId"
	//   ],
	//   "parameters": {
	//     "dartId": {
	//       "description": "Country DART ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/countries/{dartId}",
	//   "response": {
	//     "$ref": "Country"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.countries.list":

type CountriesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of countries.

func (r *CountriesService) List(profileId int64) *CountriesListCall {
	return &CountriesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/countries",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CountriesListCall) Fields(s ...googleapi.Field) *CountriesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CountriesListCall) Do() (*CountriesListResponse, error) {
	var returnValue *CountriesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of countries.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.countries.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/countries",
	//   "response": {
	//     "$ref": "CountriesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeAssets.insert":

type CreativeAssetsInsertCall struct {
	s                     *Service
	profileId             int64
	advertiserId          int64
	creativeassetmetadata *CreativeAssetMetadata
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
	context_              context.Context
	callback_             googleapi.ProgressUpdater
}

// Insert: Inserts a new creative asset.

func (r *CreativeAssetsService) Insert(profileId int64, advertiserId int64, creativeassetmetadata *CreativeAssetMetadata) *CreativeAssetsInsertCall {
	return &CreativeAssetsInsertCall{
		s:                     r.s,
		profileId:             profileId,
		advertiserId:          advertiserId,
		creativeassetmetadata: creativeassetmetadata,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets",
		context_:              context.TODO(),
	}
}

// MediaUpload takes a context and UploadCaller interface
func (c *CreativeAssetsInsertCall) Upload(ctx context.Context, u googleapi.UploadCaller) *CreativeAssetsInsertCall {
	c.caller_ = u
	c.context_ = ctx
	switch u.(type) {
	case *googleapi.MediaUpload:
		c.pathTemplate_ = "/upload/dfareporting/v2.0/userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets"
	case *googleapi.ResumableUpload:
		c.pathTemplate_ = "/resumable/upload/dfareporting/v2.0/userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets"
	}
	return c
}

// Media specifies the media to upload in a single chunk.
// At most one of Media and ResumableMedia may be set.
// The mime type type will be auto-detected unless r is a googleapi.ContentTyper as well.
func (c *CreativeAssetsInsertCall) Media(r io.Reader) *CreativeAssetsInsertCall {
	c.caller_ = &googleapi.MediaUpload{
		Media: r,
	}
	c.pathTemplate_ = "/upload/dfareporting/v2.0/userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be cancelled with ctx.
// At most one of Media and ResumableMedia may be set.
// mediaType identifies the MIME media type of the upload, such as "image/png".
// If mediaType is "", it will be auto-detected.
func (c *CreativeAssetsInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *CreativeAssetsInsertCall {
	c.caller_ = &googleapi.ResumableUpload{
		Media:         io.NewSectionReader(r, 0, size),
		MediaType:     mediaType,
		ContentLength: size,
		Callback:      c.callback_,
	}
	c.pathTemplate_ = "/resumable/upload/dfareporting/v2.0/userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets"
	c.context_ = ctx
	return c
}

// ProgressUpdater provides a callback function that will be called after every chunk.
// It should be a low-latency function in order to not slow down the upload operation.
// This should only be called when using ResumableMedia (as opposed to Media).
func (c *CreativeAssetsInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *CreativeAssetsInsertCall {
	c.callback_ = pu
	if rx, ok := c.caller_.(*googleapi.ResumableUpload); ok {
		rx.Callback = pu
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeAssetsInsertCall) Fields(s ...googleapi.Field) *CreativeAssetsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeAssetsInsertCall) Do() (*CreativeAssetMetadata, error) {
	var returnValue *CreativeAssetMetadata
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":    strconv.FormatInt(c.profileId, 10),
		"advertiserId": strconv.FormatInt(c.advertiserId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativeassetmetadata,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(c.context_, c.s.client, call)
	// {
	//   "description": "Inserts a new creative asset.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.creativeAssets.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "100MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/dfareporting/v2.0/userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/dfareporting/v2.0/userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "profileId",
	//     "advertiserId"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "Advertiser ID of this creative. This is a required field.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeAssets/{advertiserId}/creativeAssets",
	//   "request": {
	//     "$ref": "CreativeAssetMetadata"
	//   },
	//   "response": {
	//     "$ref": "CreativeAssetMetadata"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "dfareporting.creativeFieldValues.delete":

type CreativeFieldValuesDeleteCall struct {
	s               *Service
	profileId       int64
	creativeFieldId int64
	id              int64
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Delete: Deletes an existing creative field value.

func (r *CreativeFieldValuesService) Delete(profileId int64, creativeFieldId int64, id int64) *CreativeFieldValuesDeleteCall {
	return &CreativeFieldValuesDeleteCall{
		s:               r.s,
		profileId:       profileId,
		creativeFieldId: creativeFieldId,
		id:              id,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues/{id}",
	}
}

func (c *CreativeFieldValuesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":       strconv.FormatInt(c.profileId, 10),
		"creativeFieldId": strconv.FormatInt(c.creativeFieldId, 10),
		"id":              strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing creative field value.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.creativeFieldValues.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "creativeFieldId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "creativeFieldId": {
	//       "description": "Creative field ID for this creative field value.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Creative Field Value ID",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFieldValues.get":

type CreativeFieldValuesGetCall struct {
	s               *Service
	profileId       int64
	creativeFieldId int64
	id              int64
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// Get: Gets one creative field value by ID.

func (r *CreativeFieldValuesService) Get(profileId int64, creativeFieldId int64, id int64) *CreativeFieldValuesGetCall {
	return &CreativeFieldValuesGetCall{
		s:               r.s,
		profileId:       profileId,
		creativeFieldId: creativeFieldId,
		id:              id,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldValuesGetCall) Fields(s ...googleapi.Field) *CreativeFieldValuesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldValuesGetCall) Do() (*CreativeFieldValue, error) {
	var returnValue *CreativeFieldValue
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":       strconv.FormatInt(c.profileId, 10),
		"creativeFieldId": strconv.FormatInt(c.creativeFieldId, 10),
		"id":              strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one creative field value by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creativeFieldValues.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "creativeFieldId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "creativeFieldId": {
	//       "description": "Creative field ID for this creative field value.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Creative Field Value ID",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues/{id}",
	//   "response": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFieldValues.insert":

type CreativeFieldValuesInsertCall struct {
	s                  *Service
	profileId          int64
	creativeFieldId    int64
	creativefieldvalue *CreativeFieldValue
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Insert: Inserts a new creative field value.

func (r *CreativeFieldValuesService) Insert(profileId int64, creativeFieldId int64, creativefieldvalue *CreativeFieldValue) *CreativeFieldValuesInsertCall {
	return &CreativeFieldValuesInsertCall{
		s:                  r.s,
		profileId:          profileId,
		creativeFieldId:    creativeFieldId,
		creativefieldvalue: creativefieldvalue,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldValuesInsertCall) Fields(s ...googleapi.Field) *CreativeFieldValuesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldValuesInsertCall) Do() (*CreativeFieldValue, error) {
	var returnValue *CreativeFieldValue
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":       strconv.FormatInt(c.profileId, 10),
		"creativeFieldId": strconv.FormatInt(c.creativeFieldId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativefieldvalue,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new creative field value.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.creativeFieldValues.insert",
	//   "parameterOrder": [
	//     "profileId",
	//     "creativeFieldId"
	//   ],
	//   "parameters": {
	//     "creativeFieldId": {
	//       "description": "Creative field ID for this creative field value.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	//   "request": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "response": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFieldValues.list":

type CreativeFieldValuesListCall struct {
	s               *Service
	profileId       int64
	creativeFieldId int64
	caller_         googleapi.Caller
	params_         url.Values
	pathTemplate_   string
}

// List: Retrieves a list of creative field values, possibly filtered.

func (r *CreativeFieldValuesService) List(profileId int64, creativeFieldId int64) *CreativeFieldValuesListCall {
	return &CreativeFieldValuesListCall{
		s:               r.s,
		profileId:       profileId,
		creativeFieldId: creativeFieldId,
		caller_:         googleapi.JSONCall{},
		params_:         make(map[string][]string),
		pathTemplate_:   "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	}
}

// Ids sets the optional parameter "ids": Select only creative field
// values with these IDs.
func (c *CreativeFieldValuesListCall) Ids(ids ...int64) *CreativeFieldValuesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *CreativeFieldValuesListCall) MaxResults(maxResults int64) *CreativeFieldValuesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *CreativeFieldValuesListCall) PageToken(pageToken string) *CreativeFieldValuesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for creative field values by their values. Wildcards (e.g.
// *) are not allowed.
func (c *CreativeFieldValuesListCall) SearchString(searchString string) *CreativeFieldValuesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *CreativeFieldValuesListCall) SortField(sortField string) *CreativeFieldValuesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *CreativeFieldValuesListCall) SortOrder(sortOrder string) *CreativeFieldValuesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldValuesListCall) Fields(s ...googleapi.Field) *CreativeFieldValuesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldValuesListCall) Do() (*CreativeFieldValuesListResponse, error) {
	var returnValue *CreativeFieldValuesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":       strconv.FormatInt(c.profileId, 10),
		"creativeFieldId": strconv.FormatInt(c.creativeFieldId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of creative field values, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creativeFieldValues.list",
	//   "parameterOrder": [
	//     "profileId",
	//     "creativeFieldId"
	//   ],
	//   "parameters": {
	//     "creativeFieldId": {
	//       "description": "Creative field ID for this creative field value.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only creative field values with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for creative field values by their values. Wildcards (e.g. *) are not allowed.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "VALUE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	//   "response": {
	//     "$ref": "CreativeFieldValuesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFieldValues.patch":

type CreativeFieldValuesPatchCall struct {
	s                  *Service
	profileId          int64
	creativeFieldId    int64
	id                 int64
	creativefieldvalue *CreativeFieldValue
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Patch: Updates an existing creative field value. This method supports
// patch semantics.

func (r *CreativeFieldValuesService) Patch(profileId int64, creativeFieldId int64, id int64, creativefieldvalue *CreativeFieldValue) *CreativeFieldValuesPatchCall {
	return &CreativeFieldValuesPatchCall{
		s:                  r.s,
		profileId:          profileId,
		creativeFieldId:    creativeFieldId,
		id:                 id,
		creativefieldvalue: creativefieldvalue,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldValuesPatchCall) Fields(s ...googleapi.Field) *CreativeFieldValuesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldValuesPatchCall) Do() (*CreativeFieldValue, error) {
	var returnValue *CreativeFieldValue
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":       strconv.FormatInt(c.profileId, 10),
		"creativeFieldId": strconv.FormatInt(c.creativeFieldId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativefieldvalue,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative field value. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.creativeFieldValues.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "creativeFieldId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "creativeFieldId": {
	//       "description": "Creative field ID for this creative field value.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Creative Field Value ID",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	//   "request": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "response": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFieldValues.update":

type CreativeFieldValuesUpdateCall struct {
	s                  *Service
	profileId          int64
	creativeFieldId    int64
	creativefieldvalue *CreativeFieldValue
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Update: Updates an existing creative field value.

func (r *CreativeFieldValuesService) Update(profileId int64, creativeFieldId int64, creativefieldvalue *CreativeFieldValue) *CreativeFieldValuesUpdateCall {
	return &CreativeFieldValuesUpdateCall{
		s:                  r.s,
		profileId:          profileId,
		creativeFieldId:    creativeFieldId,
		creativefieldvalue: creativefieldvalue,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldValuesUpdateCall) Fields(s ...googleapi.Field) *CreativeFieldValuesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldValuesUpdateCall) Do() (*CreativeFieldValue, error) {
	var returnValue *CreativeFieldValue
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":       strconv.FormatInt(c.profileId, 10),
		"creativeFieldId": strconv.FormatInt(c.creativeFieldId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativefieldvalue,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative field value.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.creativeFieldValues.update",
	//   "parameterOrder": [
	//     "profileId",
	//     "creativeFieldId"
	//   ],
	//   "parameters": {
	//     "creativeFieldId": {
	//       "description": "Creative field ID for this creative field value.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{creativeFieldId}/creativeFieldValues",
	//   "request": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "response": {
	//     "$ref": "CreativeFieldValue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFields.delete":

type CreativeFieldsDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing creative field.

func (r *CreativeFieldsService) Delete(profileId int64, id int64) *CreativeFieldsDeleteCall {
	return &CreativeFieldsDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeFields/{id}",
	}
}

func (c *CreativeFieldsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing creative field.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.creativeFields.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative Field ID",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFields.get":

type CreativeFieldsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one creative field by ID.

func (r *CreativeFieldsService) Get(profileId int64, id int64) *CreativeFieldsGetCall {
	return &CreativeFieldsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeFields/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldsGetCall) Fields(s ...googleapi.Field) *CreativeFieldsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldsGetCall) Do() (*CreativeField, error) {
	var returnValue *CreativeField
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one creative field by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creativeFields.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative Field ID",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields/{id}",
	//   "response": {
	//     "$ref": "CreativeField"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFields.insert":

type CreativeFieldsInsertCall struct {
	s             *Service
	profileId     int64
	creativefield *CreativeField
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new creative field.

func (r *CreativeFieldsService) Insert(profileId int64, creativefield *CreativeField) *CreativeFieldsInsertCall {
	return &CreativeFieldsInsertCall{
		s:             r.s,
		profileId:     profileId,
		creativefield: creativefield,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeFields",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldsInsertCall) Fields(s ...googleapi.Field) *CreativeFieldsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldsInsertCall) Do() (*CreativeField, error) {
	var returnValue *CreativeField
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativefield,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new creative field.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.creativeFields.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields",
	//   "request": {
	//     "$ref": "CreativeField"
	//   },
	//   "response": {
	//     "$ref": "CreativeField"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFields.list":

type CreativeFieldsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of creative fields, possibly filtered.

func (r *CreativeFieldsService) List(profileId int64) *CreativeFieldsListCall {
	return &CreativeFieldsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeFields",
	}
}

// AdvertiserIds sets the optional parameter "advertiserIds": Select
// only creative fields that belong to these advertisers.
func (c *CreativeFieldsListCall) AdvertiserIds(advertiserIds ...int64) *CreativeFieldsListCall {
	for _, v_ := range advertiserIds {
		c.params_.Add("advertiserIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only creative fields
// with these IDs.
func (c *CreativeFieldsListCall) Ids(ids ...int64) *CreativeFieldsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *CreativeFieldsListCall) MaxResults(maxResults int64) *CreativeFieldsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *CreativeFieldsListCall) PageToken(pageToken string) *CreativeFieldsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for creative fields by name or ID. Wildcards (*) are
// allowed. For example, "creativefield*2015" will return creative
// fields with names like "creativefield June 2015", "creativefield
// April 2015" or simply "creativefield 2015". Most of the searches also
// add wild-cards implicitly at the start and the end of the search
// string. For example, a search string of "creativefield" will match
// creative fields with the name "my creativefield", "creativefield
// 2015" or simply "creativefield".
func (c *CreativeFieldsListCall) SearchString(searchString string) *CreativeFieldsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *CreativeFieldsListCall) SortField(sortField string) *CreativeFieldsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *CreativeFieldsListCall) SortOrder(sortOrder string) *CreativeFieldsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldsListCall) Fields(s ...googleapi.Field) *CreativeFieldsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldsListCall) Do() (*CreativeFieldsListResponse, error) {
	var returnValue *CreativeFieldsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of creative fields, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creativeFields.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserIds": {
	//       "description": "Select only creative fields that belong to these advertisers.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only creative fields with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for creative fields by name or ID. Wildcards (*) are allowed. For example, \"creativefield*2015\" will return creative fields with names like \"creativefield June 2015\", \"creativefield April 2015\" or simply \"creativefield 2015\". Most of the searches also add wild-cards implicitly at the start and the end of the search string. For example, a search string of \"creativefield\" will match creative fields with the name \"my creativefield\", \"creativefield 2015\" or simply \"creativefield\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields",
	//   "response": {
	//     "$ref": "CreativeFieldsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFields.patch":

type CreativeFieldsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	creativefield *CreativeField
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing creative field. This method supports patch
// semantics.

func (r *CreativeFieldsService) Patch(profileId int64, id int64, creativefield *CreativeField) *CreativeFieldsPatchCall {
	return &CreativeFieldsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		creativefield: creativefield,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeFields",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldsPatchCall) Fields(s ...googleapi.Field) *CreativeFieldsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldsPatchCall) Do() (*CreativeField, error) {
	var returnValue *CreativeField
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativefield,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative field. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.creativeFields.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative Field ID",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields",
	//   "request": {
	//     "$ref": "CreativeField"
	//   },
	//   "response": {
	//     "$ref": "CreativeField"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeFields.update":

type CreativeFieldsUpdateCall struct {
	s             *Service
	profileId     int64
	creativefield *CreativeField
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing creative field.

func (r *CreativeFieldsService) Update(profileId int64, creativefield *CreativeField) *CreativeFieldsUpdateCall {
	return &CreativeFieldsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		creativefield: creativefield,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeFields",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeFieldsUpdateCall) Fields(s ...googleapi.Field) *CreativeFieldsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeFieldsUpdateCall) Do() (*CreativeField, error) {
	var returnValue *CreativeField
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativefield,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative field.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.creativeFields.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeFields",
	//   "request": {
	//     "$ref": "CreativeField"
	//   },
	//   "response": {
	//     "$ref": "CreativeField"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeGroups.get":

type CreativeGroupsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one creative group by ID.

func (r *CreativeGroupsService) Get(profileId int64, id int64) *CreativeGroupsGetCall {
	return &CreativeGroupsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeGroups/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeGroupsGetCall) Fields(s ...googleapi.Field) *CreativeGroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeGroupsGetCall) Do() (*CreativeGroup, error) {
	var returnValue *CreativeGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one creative group by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creativeGroups.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeGroups/{id}",
	//   "response": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeGroups.insert":

type CreativeGroupsInsertCall struct {
	s             *Service
	profileId     int64
	creativegroup *CreativeGroup
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new creative group.

func (r *CreativeGroupsService) Insert(profileId int64, creativegroup *CreativeGroup) *CreativeGroupsInsertCall {
	return &CreativeGroupsInsertCall{
		s:             r.s,
		profileId:     profileId,
		creativegroup: creativegroup,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeGroupsInsertCall) Fields(s ...googleapi.Field) *CreativeGroupsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeGroupsInsertCall) Do() (*CreativeGroup, error) {
	var returnValue *CreativeGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativegroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new creative group.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.creativeGroups.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeGroups",
	//   "request": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "response": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeGroups.list":

type CreativeGroupsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of creative groups, possibly filtered.

func (r *CreativeGroupsService) List(profileId int64) *CreativeGroupsListCall {
	return &CreativeGroupsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeGroups",
	}
}

// AdvertiserIds sets the optional parameter "advertiserIds": Select
// only creative groups that belong to these advertisers.
func (c *CreativeGroupsListCall) AdvertiserIds(advertiserIds ...int64) *CreativeGroupsListCall {
	for _, v_ := range advertiserIds {
		c.params_.Add("advertiserIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// GroupNumber sets the optional parameter "groupNumber": Select only
// creative groups that belong to this subgroup.
func (c *CreativeGroupsListCall) GroupNumber(groupNumber int64) *CreativeGroupsListCall {
	c.params_.Set("groupNumber", fmt.Sprintf("%v", groupNumber))
	return c
}

// Ids sets the optional parameter "ids": Select only creative groups
// with these IDs.
func (c *CreativeGroupsListCall) Ids(ids ...int64) *CreativeGroupsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *CreativeGroupsListCall) MaxResults(maxResults int64) *CreativeGroupsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *CreativeGroupsListCall) PageToken(pageToken string) *CreativeGroupsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for creative groups by name or ID. Wildcards (*) are
// allowed. For example, "creativegroup*2015" will return creative
// groups with names like "creativegroup June 2015", "creativegroup
// April 2015" or simply "creativegroup 2015". Most of the searches also
// add wild-cards implicitly at the start and the end of the search
// string. For example, a search string of "creativegroup" will match
// creative groups with the name "my creativegroup", "creativegroup
// 2015" or simply "creativegroup".
func (c *CreativeGroupsListCall) SearchString(searchString string) *CreativeGroupsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *CreativeGroupsListCall) SortField(sortField string) *CreativeGroupsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *CreativeGroupsListCall) SortOrder(sortOrder string) *CreativeGroupsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeGroupsListCall) Fields(s ...googleapi.Field) *CreativeGroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeGroupsListCall) Do() (*CreativeGroupsListResponse, error) {
	var returnValue *CreativeGroupsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of creative groups, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creativeGroups.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserIds": {
	//       "description": "Select only creative groups that belong to these advertisers.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "groupNumber": {
	//       "description": "Select only creative groups that belong to this subgroup.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "ids": {
	//       "description": "Select only creative groups with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for creative groups by name or ID. Wildcards (*) are allowed. For example, \"creativegroup*2015\" will return creative groups with names like \"creativegroup June 2015\", \"creativegroup April 2015\" or simply \"creativegroup 2015\". Most of the searches also add wild-cards implicitly at the start and the end of the search string. For example, a search string of \"creativegroup\" will match creative groups with the name \"my creativegroup\", \"creativegroup 2015\" or simply \"creativegroup\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeGroups",
	//   "response": {
	//     "$ref": "CreativeGroupsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeGroups.patch":

type CreativeGroupsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	creativegroup *CreativeGroup
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing creative group. This method supports patch
// semantics.

func (r *CreativeGroupsService) Patch(profileId int64, id int64, creativegroup *CreativeGroup) *CreativeGroupsPatchCall {
	return &CreativeGroupsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		creativegroup: creativegroup,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeGroupsPatchCall) Fields(s ...googleapi.Field) *CreativeGroupsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeGroupsPatchCall) Do() (*CreativeGroup, error) {
	var returnValue *CreativeGroup
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativegroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.creativeGroups.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative group ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeGroups",
	//   "request": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "response": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creativeGroups.update":

type CreativeGroupsUpdateCall struct {
	s             *Service
	profileId     int64
	creativegroup *CreativeGroup
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing creative group.

func (r *CreativeGroupsService) Update(profileId int64, creativegroup *CreativeGroup) *CreativeGroupsUpdateCall {
	return &CreativeGroupsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		creativegroup: creativegroup,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creativeGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativeGroupsUpdateCall) Fields(s ...googleapi.Field) *CreativeGroupsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativeGroupsUpdateCall) Do() (*CreativeGroup, error) {
	var returnValue *CreativeGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.creativegroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative group.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.creativeGroups.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creativeGroups",
	//   "request": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "response": {
	//     "$ref": "CreativeGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creatives.get":

type CreativesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one creative by ID.

func (r *CreativesService) Get(profileId int64, id int64) *CreativesGetCall {
	return &CreativesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creatives/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesGetCall) Fields(s ...googleapi.Field) *CreativesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesGetCall) Do() (*Creative, error) {
	var returnValue *Creative
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one creative by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creatives.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creatives/{id}",
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creatives.insert":

type CreativesInsertCall struct {
	s             *Service
	profileId     int64
	creative      *Creative
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new creative.

func (r *CreativesService) Insert(profileId int64, creative *Creative) *CreativesInsertCall {
	return &CreativesInsertCall{
		s:             r.s,
		profileId:     profileId,
		creative:      creative,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creatives",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesInsertCall) Fields(s ...googleapi.Field) *CreativesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesInsertCall) Do() (*Creative, error) {
	var returnValue *Creative
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.creative,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new creative.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.creatives.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creatives",
	//   "request": {
	//     "$ref": "Creative"
	//   },
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creatives.list":

type CreativesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of creatives, possibly filtered.

func (r *CreativesService) List(profileId int64) *CreativesListCall {
	return &CreativesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creatives",
	}
}

// Active sets the optional parameter "active": Select only active
// creatives. Leave blank to select active and inactive creatives.
func (c *CreativesListCall) Active(active bool) *CreativesListCall {
	c.params_.Set("active", fmt.Sprintf("%v", active))
	return c
}

// AdvertiserId sets the optional parameter "advertiserId": Select only
// creatives with this advertiser ID.
func (c *CreativesListCall) AdvertiserId(advertiserId int64) *CreativesListCall {
	c.params_.Set("advertiserId", fmt.Sprintf("%v", advertiserId))
	return c
}

// Archived sets the optional parameter "archived": Select only archived
// creatives. Leave blank to select archived and unarchived creatives.
func (c *CreativesListCall) Archived(archived bool) *CreativesListCall {
	c.params_.Set("archived", fmt.Sprintf("%v", archived))
	return c
}

// CampaignId sets the optional parameter "campaignId": Select only
// creatives with this campaign ID.
func (c *CreativesListCall) CampaignId(campaignId int64) *CreativesListCall {
	c.params_.Set("campaignId", fmt.Sprintf("%v", campaignId))
	return c
}

// CompanionCreativeIds sets the optional parameter
// "companionCreativeIds": Select only in-stream video creatives with
// these companion IDs.
func (c *CreativesListCall) CompanionCreativeIds(companionCreativeIds ...int64) *CreativesListCall {
	for _, v_ := range companionCreativeIds {
		c.params_.Add("companionCreativeIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// CreativeFieldIds sets the optional parameter "creativeFieldIds":
// Select only creatives with these creative field IDs.
func (c *CreativesListCall) CreativeFieldIds(creativeFieldIds ...int64) *CreativesListCall {
	for _, v_ := range creativeFieldIds {
		c.params_.Add("creativeFieldIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only creatives with
// these IDs.
func (c *CreativesListCall) Ids(ids ...int64) *CreativesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *CreativesListCall) MaxResults(maxResults int64) *CreativesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *CreativesListCall) PageToken(pageToken string) *CreativesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// RenderingIds sets the optional parameter "renderingIds": Select only
// creatives with these rendering IDs.
func (c *CreativesListCall) RenderingIds(renderingIds ...int64) *CreativesListCall {
	for _, v_ := range renderingIds {
		c.params_.Add("renderingIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "creative*2015" will return objects with names like
// "creative June 2015", "creative April 2015" or simply "creative
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "creative" will match objects with name "my creative", "creative
// 2015" or simply "creative".
func (c *CreativesListCall) SearchString(searchString string) *CreativesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SizeIds sets the optional parameter "sizeIds": Select only creatives
// with these size IDs.
func (c *CreativesListCall) SizeIds(sizeIds ...int64) *CreativesListCall {
	for _, v_ := range sizeIds {
		c.params_.Add("sizeIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *CreativesListCall) SortField(sortField string) *CreativesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *CreativesListCall) SortOrder(sortOrder string) *CreativesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// StudioCreativeId sets the optional parameter "studioCreativeId":
// Select only creatives corresponding to this Studio creative ID.
func (c *CreativesListCall) StudioCreativeId(studioCreativeId int64) *CreativesListCall {
	c.params_.Set("studioCreativeId", fmt.Sprintf("%v", studioCreativeId))
	return c
}

// Types sets the optional parameter "types": Select only creatives with
// these creative types.
func (c *CreativesListCall) Types(types ...string) *CreativesListCall {
	c.params_["types"] = types
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesListCall) Fields(s ...googleapi.Field) *CreativesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesListCall) Do() (*CreativesListResponse, error) {
	var returnValue *CreativesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of creatives, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.creatives.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "active": {
	//       "description": "Select only active creatives. Leave blank to select active and inactive creatives.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "advertiserId": {
	//       "description": "Select only creatives with this advertiser ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "archived": {
	//       "description": "Select only archived creatives. Leave blank to select archived and unarchived creatives.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "campaignId": {
	//       "description": "Select only creatives with this campaign ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "companionCreativeIds": {
	//       "description": "Select only in-stream video creatives with these companion IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "creativeFieldIds": {
	//       "description": "Select only creatives with these creative field IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only creatives with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "renderingIds": {
	//       "description": "Select only creatives with these rendering IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"creative*2015\" will return objects with names like \"creative June 2015\", \"creative April 2015\" or simply \"creative 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"creative\" will match objects with name \"my creative\", \"creative 2015\" or simply \"creative\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sizeIds": {
	//       "description": "Select only creatives with these size IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "studioCreativeId": {
	//       "description": "Select only creatives corresponding to this Studio creative ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "types": {
	//       "description": "Select only creatives with these creative types.",
	//       "enum": [
	//         "BRAND_SAFE_DEFAULT_INSTREAM_VIDEO",
	//         "CUSTOM_INPAGE",
	//         "CUSTOM_INTERSTITIAL",
	//         "ENHANCED_BANNER",
	//         "ENHANCED_IMAGE",
	//         "FLASH_INPAGE",
	//         "HTML5_BANNER",
	//         "IMAGE",
	//         "INSTREAM_VIDEO",
	//         "INTERNAL_REDIRECT",
	//         "INTERSTITIAL_INTERNAL_REDIRECT",
	//         "REDIRECT",
	//         "RICH_MEDIA_EXPANDING",
	//         "RICH_MEDIA_IM_EXPAND",
	//         "RICH_MEDIA_INPAGE",
	//         "RICH_MEDIA_INPAGE_FLOATING",
	//         "RICH_MEDIA_INTERSTITIAL_FLOAT",
	//         "RICH_MEDIA_MOBILE_IN_APP",
	//         "RICH_MEDIA_MULTI_FLOATING",
	//         "RICH_MEDIA_PEEL_DOWN",
	//         "TRACKING_TEXT",
	//         "VPAID_LINEAR",
	//         "VPAID_NON_LINEAR"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creatives",
	//   "response": {
	//     "$ref": "CreativesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creatives.patch":

type CreativesPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	creative      *Creative
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing creative. This method supports patch
// semantics.

func (r *CreativesService) Patch(profileId int64, id int64, creative *Creative) *CreativesPatchCall {
	return &CreativesPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		creative:      creative,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creatives",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesPatchCall) Fields(s ...googleapi.Field) *CreativesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesPatchCall) Do() (*Creative, error) {
	var returnValue *Creative
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.creative,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.creatives.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Creative ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creatives",
	//   "request": {
	//     "$ref": "Creative"
	//   },
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.creatives.update":

type CreativesUpdateCall struct {
	s             *Service
	profileId     int64
	creative      *Creative
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing creative.

func (r *CreativesService) Update(profileId int64, creative *Creative) *CreativesUpdateCall {
	return &CreativesUpdateCall{
		s:             r.s,
		profileId:     profileId,
		creative:      creative,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/creatives",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CreativesUpdateCall) Fields(s ...googleapi.Field) *CreativesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *CreativesUpdateCall) Do() (*Creative, error) {
	var returnValue *Creative
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.creative,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing creative.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.creatives.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/creatives",
	//   "request": {
	//     "$ref": "Creative"
	//   },
	//   "response": {
	//     "$ref": "Creative"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.dimensionValues.query":

type DimensionValuesQueryCall struct {
	s                     *Service
	profileId             int64
	dimensionvaluerequest *DimensionValueRequest
	caller_               googleapi.Caller
	params_               url.Values
	pathTemplate_         string
}

// Query: Retrieves list of report dimension values for a list of
// filters.

func (r *DimensionValuesService) Query(profileId int64, dimensionvaluerequest *DimensionValueRequest) *DimensionValuesQueryCall {
	return &DimensionValuesQueryCall{
		s:                     r.s,
		profileId:             profileId,
		dimensionvaluerequest: dimensionvaluerequest,
		caller_:               googleapi.JSONCall{},
		params_:               make(map[string][]string),
		pathTemplate_:         "userprofiles/{profileId}/dimensionvalues/query",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *DimensionValuesQueryCall) MaxResults(maxResults int64) *DimensionValuesQueryCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The value of the
// nextToken from the previous result page.
func (c *DimensionValuesQueryCall) PageToken(pageToken string) *DimensionValuesQueryCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DimensionValuesQueryCall) Fields(s ...googleapi.Field) *DimensionValuesQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DimensionValuesQueryCall) Do() (*DimensionValueList, error) {
	var returnValue *DimensionValueList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.dimensionvaluerequest,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves list of report dimension values for a list of filters.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.dimensionValues.query",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of the nextToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/dimensionvalues/query",
	//   "request": {
	//     "$ref": "DimensionValueRequest"
	//   },
	//   "response": {
	//     "$ref": "DimensionValueList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.directorySiteContacts.get":

type DirectorySiteContactsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one directory site contact by ID.

func (r *DirectorySiteContactsService) Get(profileId int64, id int64) *DirectorySiteContactsGetCall {
	return &DirectorySiteContactsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/directorySiteContacts/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DirectorySiteContactsGetCall) Fields(s ...googleapi.Field) *DirectorySiteContactsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DirectorySiteContactsGetCall) Do() (*DirectorySiteContact, error) {
	var returnValue *DirectorySiteContact
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one directory site contact by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.directorySiteContacts.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Directory site contact ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/directorySiteContacts/{id}",
	//   "response": {
	//     "$ref": "DirectorySiteContact"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.directorySiteContacts.list":

type DirectorySiteContactsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of directory site contacts, possibly filtered.

func (r *DirectorySiteContactsService) List(profileId int64) *DirectorySiteContactsListCall {
	return &DirectorySiteContactsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/directorySiteContacts",
	}
}

// DirectorySiteIds sets the optional parameter "directorySiteIds":
// Select only directory site contacts with these directory site IDs.
// This is a required field.
func (c *DirectorySiteContactsListCall) DirectorySiteIds(directorySiteIds ...int64) *DirectorySiteContactsListCall {
	for _, v_ := range directorySiteIds {
		c.params_.Add("directorySiteIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only directory site
// contacts with these IDs.
func (c *DirectorySiteContactsListCall) Ids(ids ...int64) *DirectorySiteContactsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *DirectorySiteContactsListCall) MaxResults(maxResults int64) *DirectorySiteContactsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *DirectorySiteContactsListCall) PageToken(pageToken string) *DirectorySiteContactsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name, ID or email. Wildcards (*) are
// allowed. For example, "directory site contact*2015" will return
// objects with names like "directory site contact June 2015",
// "directory site contact April 2015" or simply "directory site contact
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "directory site contact" will match objects with name "my
// directory site contact", "directory site contact 2015" or simply
// "directory site contact".
func (c *DirectorySiteContactsListCall) SearchString(searchString string) *DirectorySiteContactsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *DirectorySiteContactsListCall) SortField(sortField string) *DirectorySiteContactsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *DirectorySiteContactsListCall) SortOrder(sortOrder string) *DirectorySiteContactsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DirectorySiteContactsListCall) Fields(s ...googleapi.Field) *DirectorySiteContactsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DirectorySiteContactsListCall) Do() (*DirectorySiteContactsListResponse, error) {
	var returnValue *DirectorySiteContactsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of directory site contacts, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.directorySiteContacts.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "directorySiteIds": {
	//       "description": "Select only directory site contacts with these directory site IDs. This is a required field.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only directory site contacts with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name, ID or email. Wildcards (*) are allowed. For example, \"directory site contact*2015\" will return objects with names like \"directory site contact June 2015\", \"directory site contact April 2015\" or simply \"directory site contact 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"directory site contact\" will match objects with name \"my directory site contact\", \"directory site contact 2015\" or simply \"directory site contact\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/directorySiteContacts",
	//   "response": {
	//     "$ref": "DirectorySiteContactsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.directorySites.get":

type DirectorySitesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one directory site by ID.

func (r *DirectorySitesService) Get(profileId int64, id int64) *DirectorySitesGetCall {
	return &DirectorySitesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/directorySites/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DirectorySitesGetCall) Fields(s ...googleapi.Field) *DirectorySitesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DirectorySitesGetCall) Do() (*DirectorySite, error) {
	var returnValue *DirectorySite
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one directory site by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.directorySites.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Directory site ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/directorySites/{id}",
	//   "response": {
	//     "$ref": "DirectorySite"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.directorySites.list":

type DirectorySitesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of directory sites, possibly filtered.

func (r *DirectorySitesService) List(profileId int64) *DirectorySitesListCall {
	return &DirectorySitesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/directorySites",
	}
}

// AcceptsInStreamVideoPlacements sets the optional parameter
// "acceptsInStreamVideoPlacements": This search filter is no longer
// supported and will have no effect on the results returned.
func (c *DirectorySitesListCall) AcceptsInStreamVideoPlacements(acceptsInStreamVideoPlacements bool) *DirectorySitesListCall {
	c.params_.Set("acceptsInStreamVideoPlacements", fmt.Sprintf("%v", acceptsInStreamVideoPlacements))
	return c
}

// AcceptsInterstitialPlacements sets the optional parameter
// "acceptsInterstitialPlacements": This search filter is no longer
// supported and will have no effect on the results returned.
func (c *DirectorySitesListCall) AcceptsInterstitialPlacements(acceptsInterstitialPlacements bool) *DirectorySitesListCall {
	c.params_.Set("acceptsInterstitialPlacements", fmt.Sprintf("%v", acceptsInterstitialPlacements))
	return c
}

// AcceptsPublisherPaidPlacements sets the optional parameter
// "acceptsPublisherPaidPlacements": Select only directory sites that
// accept publisher paid placements. This field can be left blank.
func (c *DirectorySitesListCall) AcceptsPublisherPaidPlacements(acceptsPublisherPaidPlacements bool) *DirectorySitesListCall {
	c.params_.Set("acceptsPublisherPaidPlacements", fmt.Sprintf("%v", acceptsPublisherPaidPlacements))
	return c
}

// Active sets the optional parameter "active": Select only active
// directory sites. Leave blank to retrieve both active and inactive
// directory sites.
func (c *DirectorySitesListCall) Active(active bool) *DirectorySitesListCall {
	c.params_.Set("active", fmt.Sprintf("%v", active))
	return c
}

// CountryId sets the optional parameter "countryId": Select only
// directory sites with this country ID.
func (c *DirectorySitesListCall) CountryId(countryId int64) *DirectorySitesListCall {
	c.params_.Set("countryId", fmt.Sprintf("%v", countryId))
	return c
}

// Dfp_network_code sets the optional parameter "dfp_network_code":
// Select only directory sites with this DFP network code.
func (c *DirectorySitesListCall) Dfp_network_code(dfp_network_code string) *DirectorySitesListCall {
	c.params_.Set("dfp_network_code", fmt.Sprintf("%v", dfp_network_code))
	return c
}

// Ids sets the optional parameter "ids": Select only directory sites
// with these IDs.
func (c *DirectorySitesListCall) Ids(ids ...int64) *DirectorySitesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *DirectorySitesListCall) MaxResults(maxResults int64) *DirectorySitesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *DirectorySitesListCall) PageToken(pageToken string) *DirectorySitesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ParentId sets the optional parameter "parentId": Select only
// directory sites with this parent ID.
func (c *DirectorySitesListCall) ParentId(parentId int64) *DirectorySitesListCall {
	c.params_.Set("parentId", fmt.Sprintf("%v", parentId))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name, ID or URL. Wildcards (*) are allowed.
// For example, "directory site*2015" will return objects with names
// like "directory site June 2015", "directory site April 2015" or
// simply "directory site 2015". Most of the searches also add wildcards
// implicitly at the start and the end of the search string. For
// example, a search string of "directory site" will match objects with
// name "my directory site", "directory site 2015" or simply "directory
// site".
func (c *DirectorySitesListCall) SearchString(searchString string) *DirectorySitesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *DirectorySitesListCall) SortField(sortField string) *DirectorySitesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *DirectorySitesListCall) SortOrder(sortOrder string) *DirectorySitesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DirectorySitesListCall) Fields(s ...googleapi.Field) *DirectorySitesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *DirectorySitesListCall) Do() (*DirectorySitesListResponse, error) {
	var returnValue *DirectorySitesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of directory sites, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.directorySites.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "acceptsInStreamVideoPlacements": {
	//       "description": "This search filter is no longer supported and will have no effect on the results returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "acceptsInterstitialPlacements": {
	//       "description": "This search filter is no longer supported and will have no effect on the results returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "acceptsPublisherPaidPlacements": {
	//       "description": "Select only directory sites that accept publisher paid placements. This field can be left blank.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "active": {
	//       "description": "Select only active directory sites. Leave blank to retrieve both active and inactive directory sites.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "countryId": {
	//       "description": "Select only directory sites with this country ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dfp_network_code": {
	//       "description": "Select only directory sites with this DFP network code.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only directory sites with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parentId": {
	//       "description": "Select only directory sites with this parent ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name, ID or URL. Wildcards (*) are allowed. For example, \"directory site*2015\" will return objects with names like \"directory site June 2015\", \"directory site April 2015\" or simply \"directory site 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"directory site\" will match objects with name \"my directory site\", \"directory site 2015\" or simply \"directory site\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/directorySites",
	//   "response": {
	//     "$ref": "DirectorySitesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.eventTags.delete":

type EventTagsDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing event tag.

func (r *EventTagsService) Delete(profileId int64, id int64) *EventTagsDeleteCall {
	return &EventTagsDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/eventTags/{id}",
	}
}

func (c *EventTagsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing event tag.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.eventTags.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Event tag ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/eventTags/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.eventTags.get":

type EventTagsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one event tag by ID.

func (r *EventTagsService) Get(profileId int64, id int64) *EventTagsGetCall {
	return &EventTagsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/eventTags/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventTagsGetCall) Fields(s ...googleapi.Field) *EventTagsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventTagsGetCall) Do() (*EventTag, error) {
	var returnValue *EventTag
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one event tag by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.eventTags.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Event tag ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/eventTags/{id}",
	//   "response": {
	//     "$ref": "EventTag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.eventTags.insert":

type EventTagsInsertCall struct {
	s             *Service
	profileId     int64
	eventtag      *EventTag
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new event tag.

func (r *EventTagsService) Insert(profileId int64, eventtag *EventTag) *EventTagsInsertCall {
	return &EventTagsInsertCall{
		s:             r.s,
		profileId:     profileId,
		eventtag:      eventtag,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/eventTags",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventTagsInsertCall) Fields(s ...googleapi.Field) *EventTagsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventTagsInsertCall) Do() (*EventTag, error) {
	var returnValue *EventTag
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.eventtag,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new event tag.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.eventTags.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/eventTags",
	//   "request": {
	//     "$ref": "EventTag"
	//   },
	//   "response": {
	//     "$ref": "EventTag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.eventTags.list":

type EventTagsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of event tags, possibly filtered.

func (r *EventTagsService) List(profileId int64) *EventTagsListCall {
	return &EventTagsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/eventTags",
	}
}

// AdId sets the optional parameter "adId": Select only event tags that
// belong to this ad.
func (c *EventTagsListCall) AdId(adId int64) *EventTagsListCall {
	c.params_.Set("adId", fmt.Sprintf("%v", adId))
	return c
}

// AdvertiserId sets the optional parameter "advertiserId": Select only
// event tags that belong to this advertiser.
func (c *EventTagsListCall) AdvertiserId(advertiserId int64) *EventTagsListCall {
	c.params_.Set("advertiserId", fmt.Sprintf("%v", advertiserId))
	return c
}

// CampaignId sets the optional parameter "campaignId": Select only
// event tags that belong to this campaign.
func (c *EventTagsListCall) CampaignId(campaignId int64) *EventTagsListCall {
	c.params_.Set("campaignId", fmt.Sprintf("%v", campaignId))
	return c
}

// DefinitionsOnly sets the optional parameter "definitionsOnly":
// Examine only the specified ad or campaign or advertiser's event tags
// for matching selector criteria. When set to false, the parent
// advertiser and parent campaign is examined as well. In addition, when
// set to false, the status field is examined as well along with the
// enabledByDefault field.
func (c *EventTagsListCall) DefinitionsOnly(definitionsOnly bool) *EventTagsListCall {
	c.params_.Set("definitionsOnly", fmt.Sprintf("%v", definitionsOnly))
	return c
}

// Enabled sets the optional parameter "enabled": Select only enabled
// event tags. When definitionsOnly is set to true, only the specified
// advertiser or campaign's event tags' enabledByDefault field is
// examined. When definitionsOnly is set to false, the specified ad or
// specified campaign's parent advertiser's or parent campaign's event
// tags' enabledByDefault and status fields are examined as well.
func (c *EventTagsListCall) Enabled(enabled bool) *EventTagsListCall {
	c.params_.Set("enabled", fmt.Sprintf("%v", enabled))
	return c
}

// EventTagTypes sets the optional parameter "eventTagTypes": Select
// only event tags with the specified event tag types. Event tag types
// can be used to specify whether to use a third-party pixel, a
// third-party JavaScript URL, or a third-party click-through URL for
// either impression or click tracking.
func (c *EventTagsListCall) EventTagTypes(eventTagTypes ...string) *EventTagsListCall {
	c.params_["eventTagTypes"] = eventTagTypes
	return c
}

// Ids sets the optional parameter "ids": Select only event tags with
// these IDs.
func (c *EventTagsListCall) Ids(ids ...int64) *EventTagsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "eventtag*2015" will return objects with names like
// "eventtag June 2015", "eventtag April 2015" or simply "eventtag
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "eventtag" will match objects with name "my eventtag", "eventtag
// 2015" or simply "eventtag".
func (c *EventTagsListCall) SearchString(searchString string) *EventTagsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *EventTagsListCall) SortField(sortField string) *EventTagsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *EventTagsListCall) SortOrder(sortOrder string) *EventTagsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventTagsListCall) Fields(s ...googleapi.Field) *EventTagsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventTagsListCall) Do() (*EventTagsListResponse, error) {
	var returnValue *EventTagsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of event tags, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.eventTags.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "adId": {
	//       "description": "Select only event tags that belong to this ad.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "advertiserId": {
	//       "description": "Select only event tags that belong to this advertiser.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "campaignId": {
	//       "description": "Select only event tags that belong to this campaign.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "definitionsOnly": {
	//       "description": "Examine only the specified ad or campaign or advertiser's event tags for matching selector criteria. When set to false, the parent advertiser and parent campaign is examined as well. In addition, when set to false, the status field is examined as well along with the enabledByDefault field.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "enabled": {
	//       "description": "Select only enabled event tags. When definitionsOnly is set to true, only the specified advertiser or campaign's event tags' enabledByDefault field is examined. When definitionsOnly is set to false, the specified ad or specified campaign's parent advertiser's or parent campaign's event tags' enabledByDefault and status fields are examined as well.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "eventTagTypes": {
	//       "description": "Select only event tags with the specified event tag types. Event tag types can be used to specify whether to use a third-party pixel, a third-party JavaScript URL, or a third-party click-through URL for either impression or click tracking.",
	//       "enum": [
	//         "CLICK_THROUGH_EVENT_TAG",
	//         "IMPRESSION_IMAGE_EVENT_TAG",
	//         "IMPRESSION_JAVASCRIPT_EVENT_TAG"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only event tags with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"eventtag*2015\" will return objects with names like \"eventtag June 2015\", \"eventtag April 2015\" or simply \"eventtag 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"eventtag\" will match objects with name \"my eventtag\", \"eventtag 2015\" or simply \"eventtag\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/eventTags",
	//   "response": {
	//     "$ref": "EventTagsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.eventTags.patch":

type EventTagsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	eventtag      *EventTag
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing event tag. This method supports patch
// semantics.

func (r *EventTagsService) Patch(profileId int64, id int64, eventtag *EventTag) *EventTagsPatchCall {
	return &EventTagsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		eventtag:      eventtag,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/eventTags",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventTagsPatchCall) Fields(s ...googleapi.Field) *EventTagsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventTagsPatchCall) Do() (*EventTag, error) {
	var returnValue *EventTag
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.eventtag,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing event tag. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.eventTags.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Event tag ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/eventTags",
	//   "request": {
	//     "$ref": "EventTag"
	//   },
	//   "response": {
	//     "$ref": "EventTag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.eventTags.update":

type EventTagsUpdateCall struct {
	s             *Service
	profileId     int64
	eventtag      *EventTag
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing event tag.

func (r *EventTagsService) Update(profileId int64, eventtag *EventTag) *EventTagsUpdateCall {
	return &EventTagsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		eventtag:      eventtag,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/eventTags",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EventTagsUpdateCall) Fields(s ...googleapi.Field) *EventTagsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *EventTagsUpdateCall) Do() (*EventTag, error) {
	var returnValue *EventTag
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.eventtag,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing event tag.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.eventTags.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/eventTags",
	//   "request": {
	//     "$ref": "EventTag"
	//   },
	//   "response": {
	//     "$ref": "EventTag"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.files.get":

type FilesGetCall struct {
	s             *Service
	reportId      int64
	fileId        int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a report file by its report ID and file ID.

func (r *FilesService) Get(reportId int64, fileId int64) *FilesGetCall {
	return &FilesGetCall{
		s:             r.s,
		reportId:      reportId,
		fileId:        fileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "reports/{reportId}/files/{fileId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FilesGetCall) Fields(s ...googleapi.Field) *FilesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FilesGetCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"reportId": strconv.FormatInt(c.reportId, 10),
		"fileId":   strconv.FormatInt(c.fileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *FilesGetCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"reportId": strconv.FormatInt(c.reportId, 10),
		"fileId":   strconv.FormatInt(c.fileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a report file by its report ID and file ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.files.get",
	//   "parameterOrder": [
	//     "reportId",
	//     "fileId"
	//   ],
	//   "parameters": {
	//     "fileId": {
	//       "description": "The ID of the report file.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "reports/{reportId}/files/{fileId}",
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "dfareporting.files.list":

type FilesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists files for a user profile.

func (r *FilesService) List(profileId int64) *FilesListCall {
	return &FilesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/files",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *FilesListCall) MaxResults(maxResults int64) *FilesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The value of the
// nextToken from the previous result page.
func (c *FilesListCall) PageToken(pageToken string) *FilesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Scope sets the optional parameter "scope": The scope that defines
// which results are returned, default is 'MINE'.
func (c *FilesListCall) Scope(scope string) *FilesListCall {
	c.params_.Set("scope", fmt.Sprintf("%v", scope))
	return c
}

// SortField sets the optional parameter "sortField": The field by which
// to sort the list.
func (c *FilesListCall) SortField(sortField string) *FilesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is 'DESCENDING'.
func (c *FilesListCall) SortOrder(sortOrder string) *FilesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FilesListCall) Fields(s ...googleapi.Field) *FilesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FilesListCall) Do() (*FileList, error) {
	var returnValue *FileList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists files for a user profile.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.files.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of the nextToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "The DFA profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "default": "MINE",
	//       "description": "The scope that defines which results are returned, default is 'MINE'.",
	//       "enum": [
	//         "ALL",
	//         "MINE",
	//         "SHARED_WITH_ME"
	//       ],
	//       "enumDescriptions": [
	//         "All files in account.",
	//         "My files.",
	//         "Files shared with me."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "default": "LAST_MODIFIED_TIME",
	//       "description": "The field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "LAST_MODIFIED_TIME"
	//       ],
	//       "enumDescriptions": [
	//         "Sort by file ID.",
	//         "Sort by 'lastmodifiedAt' field."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "default": "DESCENDING",
	//       "description": "Order of sorted results, default is 'DESCENDING'.",
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
	//   "path": "userprofiles/{profileId}/files",
	//   "response": {
	//     "$ref": "FileList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.delete":

type FloodlightActivitiesDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing floodlight activity.

func (r *FloodlightActivitiesService) Delete(profileId int64, id int64) *FloodlightActivitiesDeleteCall {
	return &FloodlightActivitiesDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivities/{id}",
	}
}

func (c *FloodlightActivitiesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing floodlight activity.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.floodlightActivities.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight activity ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.generatetag":

type FloodlightActivitiesGeneratetagCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generatetag: Generates a tag for a floodlight activity.

func (r *FloodlightActivitiesService) Generatetag(profileId int64) *FloodlightActivitiesGeneratetagCall {
	return &FloodlightActivitiesGeneratetagCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivities/generatetag",
	}
}

// FloodlightActivityId sets the optional parameter
// "floodlightActivityId": Floodlight activity ID for which we want to
// generate a tag.
func (c *FloodlightActivitiesGeneratetagCall) FloodlightActivityId(floodlightActivityId int64) *FloodlightActivitiesGeneratetagCall {
	c.params_.Set("floodlightActivityId", fmt.Sprintf("%v", floodlightActivityId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivitiesGeneratetagCall) Fields(s ...googleapi.Field) *FloodlightActivitiesGeneratetagCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivitiesGeneratetagCall) Do() (*FloodlightActivitiesGenerateTagResponse, error) {
	var returnValue *FloodlightActivitiesGenerateTagResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generates a tag for a floodlight activity.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.floodlightActivities.generatetag",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "floodlightActivityId": {
	//       "description": "Floodlight activity ID for which we want to generate a tag.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities/generatetag",
	//   "response": {
	//     "$ref": "FloodlightActivitiesGenerateTagResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.get":

type FloodlightActivitiesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one floodlight activity by ID.

func (r *FloodlightActivitiesService) Get(profileId int64, id int64) *FloodlightActivitiesGetCall {
	return &FloodlightActivitiesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivities/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivitiesGetCall) Fields(s ...googleapi.Field) *FloodlightActivitiesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivitiesGetCall) Do() (*FloodlightActivity, error) {
	var returnValue *FloodlightActivity
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one floodlight activity by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.floodlightActivities.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight activity ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities/{id}",
	//   "response": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.insert":

type FloodlightActivitiesInsertCall struct {
	s                  *Service
	profileId          int64
	floodlightactivity *FloodlightActivity
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Insert: Inserts a new floodlight activity.

func (r *FloodlightActivitiesService) Insert(profileId int64, floodlightactivity *FloodlightActivity) *FloodlightActivitiesInsertCall {
	return &FloodlightActivitiesInsertCall{
		s:                  r.s,
		profileId:          profileId,
		floodlightactivity: floodlightactivity,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/floodlightActivities",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivitiesInsertCall) Fields(s ...googleapi.Field) *FloodlightActivitiesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivitiesInsertCall) Do() (*FloodlightActivity, error) {
	var returnValue *FloodlightActivity
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightactivity,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new floodlight activity.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.floodlightActivities.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities",
	//   "request": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "response": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.list":

type FloodlightActivitiesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of floodlight activities, possibly filtered.

func (r *FloodlightActivitiesService) List(profileId int64) *FloodlightActivitiesListCall {
	return &FloodlightActivitiesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivities",
	}
}

// AdvertiserId sets the optional parameter "advertiserId": Select only
// floodlight activities for the specified advertiser ID. Must specify
// either ids, advertiserId, or floodlightConfigurationId for a
// non-empty result.
func (c *FloodlightActivitiesListCall) AdvertiserId(advertiserId int64) *FloodlightActivitiesListCall {
	c.params_.Set("advertiserId", fmt.Sprintf("%v", advertiserId))
	return c
}

// FloodlightActivityGroupIds sets the optional parameter
// "floodlightActivityGroupIds": Select only floodlight activities with
// the specified floodlight activity group IDs.
func (c *FloodlightActivitiesListCall) FloodlightActivityGroupIds(floodlightActivityGroupIds ...int64) *FloodlightActivitiesListCall {
	for _, v_ := range floodlightActivityGroupIds {
		c.params_.Add("floodlightActivityGroupIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// FloodlightActivityGroupName sets the optional parameter
// "floodlightActivityGroupName": Select only floodlight activities with
// the specified floodlight activity group name.
func (c *FloodlightActivitiesListCall) FloodlightActivityGroupName(floodlightActivityGroupName string) *FloodlightActivitiesListCall {
	c.params_.Set("floodlightActivityGroupName", fmt.Sprintf("%v", floodlightActivityGroupName))
	return c
}

// FloodlightActivityGroupTagString sets the optional parameter
// "floodlightActivityGroupTagString": Select only floodlight activities
// with the specified floodlight activity group tag string.
func (c *FloodlightActivitiesListCall) FloodlightActivityGroupTagString(floodlightActivityGroupTagString string) *FloodlightActivitiesListCall {
	c.params_.Set("floodlightActivityGroupTagString", fmt.Sprintf("%v", floodlightActivityGroupTagString))
	return c
}

// FloodlightActivityGroupType sets the optional parameter
// "floodlightActivityGroupType": Select only floodlight activities with
// the specified floodlight activity group type.
func (c *FloodlightActivitiesListCall) FloodlightActivityGroupType(floodlightActivityGroupType string) *FloodlightActivitiesListCall {
	c.params_.Set("floodlightActivityGroupType", fmt.Sprintf("%v", floodlightActivityGroupType))
	return c
}

// FloodlightConfigurationId sets the optional parameter
// "floodlightConfigurationId": Select only floodlight activities for
// the specified floodlight configuration ID. Must specify either ids,
// advertiserId, or floodlightConfigurationId for a non-empty result.
func (c *FloodlightActivitiesListCall) FloodlightConfigurationId(floodlightConfigurationId int64) *FloodlightActivitiesListCall {
	c.params_.Set("floodlightConfigurationId", fmt.Sprintf("%v", floodlightConfigurationId))
	return c
}

// Ids sets the optional parameter "ids": Select only floodlight
// activities with the specified IDs. Must specify either ids,
// advertiserId, or floodlightConfigurationId for a non-empty result.
func (c *FloodlightActivitiesListCall) Ids(ids ...int64) *FloodlightActivitiesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *FloodlightActivitiesListCall) MaxResults(maxResults int64) *FloodlightActivitiesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *FloodlightActivitiesListCall) PageToken(pageToken string) *FloodlightActivitiesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "floodlightactivity*2015" will return objects with names
// like "floodlightactivity June 2015", "floodlightactivity April 2015"
// or simply "floodlightactivity 2015". Most of the searches also add
// wildcards implicitly at the start and the end of the search string.
// For example, a search string of "floodlightactivity" will match
// objects with name "my floodlightactivity activity",
// "floodlightactivity 2015" or simply "floodlightactivity".
func (c *FloodlightActivitiesListCall) SearchString(searchString string) *FloodlightActivitiesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *FloodlightActivitiesListCall) SortField(sortField string) *FloodlightActivitiesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *FloodlightActivitiesListCall) SortOrder(sortOrder string) *FloodlightActivitiesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// TagString sets the optional parameter "tagString": Select only
// floodlight activities with the specified tag string.
func (c *FloodlightActivitiesListCall) TagString(tagString string) *FloodlightActivitiesListCall {
	c.params_.Set("tagString", fmt.Sprintf("%v", tagString))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivitiesListCall) Fields(s ...googleapi.Field) *FloodlightActivitiesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivitiesListCall) Do() (*FloodlightActivitiesListResponse, error) {
	var returnValue *FloodlightActivitiesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of floodlight activities, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.floodlightActivities.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "Select only floodlight activities for the specified advertiser ID. Must specify either ids, advertiserId, or floodlightConfigurationId for a non-empty result.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "floodlightActivityGroupIds": {
	//       "description": "Select only floodlight activities with the specified floodlight activity group IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "floodlightActivityGroupName": {
	//       "description": "Select only floodlight activities with the specified floodlight activity group name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "floodlightActivityGroupTagString": {
	//       "description": "Select only floodlight activities with the specified floodlight activity group tag string.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "floodlightActivityGroupType": {
	//       "description": "Select only floodlight activities with the specified floodlight activity group type.",
	//       "enum": [
	//         "COUNTER",
	//         "SALE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "floodlightConfigurationId": {
	//       "description": "Select only floodlight activities for the specified floodlight configuration ID. Must specify either ids, advertiserId, or floodlightConfigurationId for a non-empty result.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only floodlight activities with the specified IDs. Must specify either ids, advertiserId, or floodlightConfigurationId for a non-empty result.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"floodlightactivity*2015\" will return objects with names like \"floodlightactivity June 2015\", \"floodlightactivity April 2015\" or simply \"floodlightactivity 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"floodlightactivity\" will match objects with name \"my floodlightactivity activity\", \"floodlightactivity 2015\" or simply \"floodlightactivity\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tagString": {
	//       "description": "Select only floodlight activities with the specified tag string.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities",
	//   "response": {
	//     "$ref": "FloodlightActivitiesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.patch":

type FloodlightActivitiesPatchCall struct {
	s                  *Service
	profileId          int64
	id                 int64
	floodlightactivity *FloodlightActivity
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Patch: Updates an existing floodlight activity. This method supports
// patch semantics.

func (r *FloodlightActivitiesService) Patch(profileId int64, id int64, floodlightactivity *FloodlightActivity) *FloodlightActivitiesPatchCall {
	return &FloodlightActivitiesPatchCall{
		s:                  r.s,
		profileId:          profileId,
		id:                 id,
		floodlightactivity: floodlightactivity,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/floodlightActivities",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivitiesPatchCall) Fields(s ...googleapi.Field) *FloodlightActivitiesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivitiesPatchCall) Do() (*FloodlightActivity, error) {
	var returnValue *FloodlightActivity
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightactivity,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing floodlight activity. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.floodlightActivities.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight activity ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities",
	//   "request": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "response": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivities.update":

type FloodlightActivitiesUpdateCall struct {
	s                  *Service
	profileId          int64
	floodlightactivity *FloodlightActivity
	caller_            googleapi.Caller
	params_            url.Values
	pathTemplate_      string
}

// Update: Updates an existing floodlight activity.

func (r *FloodlightActivitiesService) Update(profileId int64, floodlightactivity *FloodlightActivity) *FloodlightActivitiesUpdateCall {
	return &FloodlightActivitiesUpdateCall{
		s:                  r.s,
		profileId:          profileId,
		floodlightactivity: floodlightactivity,
		caller_:            googleapi.JSONCall{},
		params_:            make(map[string][]string),
		pathTemplate_:      "userprofiles/{profileId}/floodlightActivities",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivitiesUpdateCall) Fields(s ...googleapi.Field) *FloodlightActivitiesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivitiesUpdateCall) Do() (*FloodlightActivity, error) {
	var returnValue *FloodlightActivity
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightactivity,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing floodlight activity.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.floodlightActivities.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivities",
	//   "request": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "response": {
	//     "$ref": "FloodlightActivity"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivityGroups.delete":

type FloodlightActivityGroupsDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing floodlight activity group.

func (r *FloodlightActivityGroupsService) Delete(profileId int64, id int64) *FloodlightActivityGroupsDeleteCall {
	return &FloodlightActivityGroupsDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivityGroups/{id}",
	}
}

func (c *FloodlightActivityGroupsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing floodlight activity group.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.floodlightActivityGroups.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight activity Group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivityGroups/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivityGroups.get":

type FloodlightActivityGroupsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one floodlight activity group by ID.

func (r *FloodlightActivityGroupsService) Get(profileId int64, id int64) *FloodlightActivityGroupsGetCall {
	return &FloodlightActivityGroupsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivityGroups/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivityGroupsGetCall) Fields(s ...googleapi.Field) *FloodlightActivityGroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivityGroupsGetCall) Do() (*FloodlightActivityGroup, error) {
	var returnValue *FloodlightActivityGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one floodlight activity group by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.floodlightActivityGroups.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight activity Group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivityGroups/{id}",
	//   "response": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivityGroups.insert":

type FloodlightActivityGroupsInsertCall struct {
	s                       *Service
	profileId               int64
	floodlightactivitygroup *FloodlightActivityGroup
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Insert: Inserts a new floodlight activity group.

func (r *FloodlightActivityGroupsService) Insert(profileId int64, floodlightactivitygroup *FloodlightActivityGroup) *FloodlightActivityGroupsInsertCall {
	return &FloodlightActivityGroupsInsertCall{
		s:                       r.s,
		profileId:               profileId,
		floodlightactivitygroup: floodlightactivitygroup,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "userprofiles/{profileId}/floodlightActivityGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivityGroupsInsertCall) Fields(s ...googleapi.Field) *FloodlightActivityGroupsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivityGroupsInsertCall) Do() (*FloodlightActivityGroup, error) {
	var returnValue *FloodlightActivityGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightactivitygroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new floodlight activity group.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.floodlightActivityGroups.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivityGroups",
	//   "request": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "response": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivityGroups.list":

type FloodlightActivityGroupsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of floodlight activity groups, possibly
// filtered.

func (r *FloodlightActivityGroupsService) List(profileId int64) *FloodlightActivityGroupsListCall {
	return &FloodlightActivityGroupsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightActivityGroups",
	}
}

// AdvertiserId sets the optional parameter "advertiserId": Select only
// floodlight activity groups with the specified advertiser ID. Must
// specify either ids, advertiserId, or floodlightConfigurationId for a
// non-empty result.
func (c *FloodlightActivityGroupsListCall) AdvertiserId(advertiserId int64) *FloodlightActivityGroupsListCall {
	c.params_.Set("advertiserId", fmt.Sprintf("%v", advertiserId))
	return c
}

// FloodlightConfigurationId sets the optional parameter
// "floodlightConfigurationId": Select only floodlight activity groups
// with the specified floodlight configuration ID. Must specify either
// ids, advertiserId, or floodlightConfigurationId for a non-empty
// result.
func (c *FloodlightActivityGroupsListCall) FloodlightConfigurationId(floodlightConfigurationId int64) *FloodlightActivityGroupsListCall {
	c.params_.Set("floodlightConfigurationId", fmt.Sprintf("%v", floodlightConfigurationId))
	return c
}

// Ids sets the optional parameter "ids": Select only floodlight
// activity groups with the specified IDs. Must specify either ids,
// advertiserId, or floodlightConfigurationId for a non-empty result.
func (c *FloodlightActivityGroupsListCall) Ids(ids ...int64) *FloodlightActivityGroupsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *FloodlightActivityGroupsListCall) MaxResults(maxResults int64) *FloodlightActivityGroupsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *FloodlightActivityGroupsListCall) PageToken(pageToken string) *FloodlightActivityGroupsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "floodlightactivitygroup*2015" will return objects with
// names like "floodlightactivitygroup June 2015",
// "floodlightactivitygroup April 2015" or simply
// "floodlightactivitygroup 2015". Most of the searches also add
// wildcards implicitly at the start and the end of the search string.
// For example, a search string of "floodlightactivitygroup" will match
// objects with name "my floodlightactivitygroup activity",
// "floodlightactivitygroup 2015" or simply "floodlightactivitygroup".
func (c *FloodlightActivityGroupsListCall) SearchString(searchString string) *FloodlightActivityGroupsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *FloodlightActivityGroupsListCall) SortField(sortField string) *FloodlightActivityGroupsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *FloodlightActivityGroupsListCall) SortOrder(sortOrder string) *FloodlightActivityGroupsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Type sets the optional parameter "type": Select only floodlight
// activity groups with the specified floodlight activity group type.
func (c *FloodlightActivityGroupsListCall) Type(type_ string) *FloodlightActivityGroupsListCall {
	c.params_.Set("type", fmt.Sprintf("%v", type_))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivityGroupsListCall) Fields(s ...googleapi.Field) *FloodlightActivityGroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivityGroupsListCall) Do() (*FloodlightActivityGroupsListResponse, error) {
	var returnValue *FloodlightActivityGroupsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of floodlight activity groups, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.floodlightActivityGroups.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserId": {
	//       "description": "Select only floodlight activity groups with the specified advertiser ID. Must specify either ids, advertiserId, or floodlightConfigurationId for a non-empty result.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "floodlightConfigurationId": {
	//       "description": "Select only floodlight activity groups with the specified floodlight configuration ID. Must specify either ids, advertiserId, or floodlightConfigurationId for a non-empty result.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only floodlight activity groups with the specified IDs. Must specify either ids, advertiserId, or floodlightConfigurationId for a non-empty result.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"floodlightactivitygroup*2015\" will return objects with names like \"floodlightactivitygroup June 2015\", \"floodlightactivitygroup April 2015\" or simply \"floodlightactivitygroup 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"floodlightactivitygroup\" will match objects with name \"my floodlightactivitygroup activity\", \"floodlightactivitygroup 2015\" or simply \"floodlightactivitygroup\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "Select only floodlight activity groups with the specified floodlight activity group type.",
	//       "enum": [
	//         "COUNTER",
	//         "SALE"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivityGroups",
	//   "response": {
	//     "$ref": "FloodlightActivityGroupsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivityGroups.patch":

type FloodlightActivityGroupsPatchCall struct {
	s                       *Service
	profileId               int64
	id                      int64
	floodlightactivitygroup *FloodlightActivityGroup
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Patch: Updates an existing floodlight activity group. This method
// supports patch semantics.

func (r *FloodlightActivityGroupsService) Patch(profileId int64, id int64, floodlightactivitygroup *FloodlightActivityGroup) *FloodlightActivityGroupsPatchCall {
	return &FloodlightActivityGroupsPatchCall{
		s:         r.s,
		profileId: profileId,
		id:        id,
		floodlightactivitygroup: floodlightactivitygroup,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "userprofiles/{profileId}/floodlightActivityGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivityGroupsPatchCall) Fields(s ...googleapi.Field) *FloodlightActivityGroupsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivityGroupsPatchCall) Do() (*FloodlightActivityGroup, error) {
	var returnValue *FloodlightActivityGroup
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightactivitygroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing floodlight activity group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.floodlightActivityGroups.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight activity Group ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivityGroups",
	//   "request": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "response": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightActivityGroups.update":

type FloodlightActivityGroupsUpdateCall struct {
	s                       *Service
	profileId               int64
	floodlightactivitygroup *FloodlightActivityGroup
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Update: Updates an existing floodlight activity group.

func (r *FloodlightActivityGroupsService) Update(profileId int64, floodlightactivitygroup *FloodlightActivityGroup) *FloodlightActivityGroupsUpdateCall {
	return &FloodlightActivityGroupsUpdateCall{
		s:                       r.s,
		profileId:               profileId,
		floodlightactivitygroup: floodlightactivitygroup,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "userprofiles/{profileId}/floodlightActivityGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightActivityGroupsUpdateCall) Fields(s ...googleapi.Field) *FloodlightActivityGroupsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightActivityGroupsUpdateCall) Do() (*FloodlightActivityGroup, error) {
	var returnValue *FloodlightActivityGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightactivitygroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing floodlight activity group.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.floodlightActivityGroups.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightActivityGroups",
	//   "request": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "response": {
	//     "$ref": "FloodlightActivityGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightConfigurations.get":

type FloodlightConfigurationsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one floodlight configuration by ID.

func (r *FloodlightConfigurationsService) Get(profileId int64, id int64) *FloodlightConfigurationsGetCall {
	return &FloodlightConfigurationsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightConfigurations/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightConfigurationsGetCall) Fields(s ...googleapi.Field) *FloodlightConfigurationsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightConfigurationsGetCall) Do() (*FloodlightConfiguration, error) {
	var returnValue *FloodlightConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one floodlight configuration by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.floodlightConfigurations.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight configuration ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightConfigurations/{id}",
	//   "response": {
	//     "$ref": "FloodlightConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightConfigurations.list":

type FloodlightConfigurationsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of floodlight configurations, possibly
// filtered.

func (r *FloodlightConfigurationsService) List(profileId int64) *FloodlightConfigurationsListCall {
	return &FloodlightConfigurationsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/floodlightConfigurations",
	}
}

// Ids sets the optional parameter "ids": Set of IDs of floodlight
// configurations to retrieve. Required field; otherwise an empty list
// will be returned.
func (c *FloodlightConfigurationsListCall) Ids(ids ...int64) *FloodlightConfigurationsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightConfigurationsListCall) Fields(s ...googleapi.Field) *FloodlightConfigurationsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightConfigurationsListCall) Do() (*FloodlightConfigurationsListResponse, error) {
	var returnValue *FloodlightConfigurationsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of floodlight configurations, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.floodlightConfigurations.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "ids": {
	//       "description": "Set of IDs of floodlight configurations to retrieve. Required field; otherwise an empty list will be returned.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightConfigurations",
	//   "response": {
	//     "$ref": "FloodlightConfigurationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightConfigurations.patch":

type FloodlightConfigurationsPatchCall struct {
	s                       *Service
	profileId               int64
	id                      int64
	floodlightconfiguration *FloodlightConfiguration
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Patch: Updates an existing floodlight configuration. This method
// supports patch semantics.

func (r *FloodlightConfigurationsService) Patch(profileId int64, id int64, floodlightconfiguration *FloodlightConfiguration) *FloodlightConfigurationsPatchCall {
	return &FloodlightConfigurationsPatchCall{
		s:         r.s,
		profileId: profileId,
		id:        id,
		floodlightconfiguration: floodlightconfiguration,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "userprofiles/{profileId}/floodlightConfigurations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightConfigurationsPatchCall) Fields(s ...googleapi.Field) *FloodlightConfigurationsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightConfigurationsPatchCall) Do() (*FloodlightConfiguration, error) {
	var returnValue *FloodlightConfiguration
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing floodlight configuration. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.floodlightConfigurations.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Floodlight configuration ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightConfigurations",
	//   "request": {
	//     "$ref": "FloodlightConfiguration"
	//   },
	//   "response": {
	//     "$ref": "FloodlightConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.floodlightConfigurations.update":

type FloodlightConfigurationsUpdateCall struct {
	s                       *Service
	profileId               int64
	floodlightconfiguration *FloodlightConfiguration
	caller_                 googleapi.Caller
	params_                 url.Values
	pathTemplate_           string
}

// Update: Updates an existing floodlight configuration.

func (r *FloodlightConfigurationsService) Update(profileId int64, floodlightconfiguration *FloodlightConfiguration) *FloodlightConfigurationsUpdateCall {
	return &FloodlightConfigurationsUpdateCall{
		s:                       r.s,
		profileId:               profileId,
		floodlightconfiguration: floodlightconfiguration,
		caller_:                 googleapi.JSONCall{},
		params_:                 make(map[string][]string),
		pathTemplate_:           "userprofiles/{profileId}/floodlightConfigurations",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FloodlightConfigurationsUpdateCall) Fields(s ...googleapi.Field) *FloodlightConfigurationsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *FloodlightConfigurationsUpdateCall) Do() (*FloodlightConfiguration, error) {
	var returnValue *FloodlightConfiguration
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.floodlightconfiguration,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing floodlight configuration.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.floodlightConfigurations.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/floodlightConfigurations",
	//   "request": {
	//     "$ref": "FloodlightConfiguration"
	//   },
	//   "response": {
	//     "$ref": "FloodlightConfiguration"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.landingPages.delete":

type LandingPagesDeleteCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing campaign landing page.

func (r *LandingPagesService) Delete(profileId int64, campaignId int64, id int64) *LandingPagesDeleteCall {
	return &LandingPagesDeleteCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/landingPages/{id}",
	}
}

func (c *LandingPagesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
		"id":         strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing campaign landing page.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.landingPages.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Landing page campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Landing page ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/landingPages/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.landingPages.get":

type LandingPagesGetCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one campaign landing page by ID.

func (r *LandingPagesService) Get(profileId int64, campaignId int64, id int64) *LandingPagesGetCall {
	return &LandingPagesGetCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/landingPages/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LandingPagesGetCall) Fields(s ...googleapi.Field) *LandingPagesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LandingPagesGetCall) Do() (*LandingPage, error) {
	var returnValue *LandingPage
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
		"id":         strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one campaign landing page by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.landingPages.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Landing page campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Landing page ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/landingPages/{id}",
	//   "response": {
	//     "$ref": "LandingPage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.landingPages.insert":

type LandingPagesInsertCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	landingpage   *LandingPage
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new landing page for the specified campaign.

func (r *LandingPagesService) Insert(profileId int64, campaignId int64, landingpage *LandingPage) *LandingPagesInsertCall {
	return &LandingPagesInsertCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		landingpage:   landingpage,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LandingPagesInsertCall) Fields(s ...googleapi.Field) *LandingPagesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LandingPagesInsertCall) Do() (*LandingPage, error) {
	var returnValue *LandingPage
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.landingpage,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new landing page for the specified campaign.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.landingPages.insert",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Landing page campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	//   "request": {
	//     "$ref": "LandingPage"
	//   },
	//   "response": {
	//     "$ref": "LandingPage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.landingPages.list":

type LandingPagesListCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves the list of landing pages for the specified campaign.

func (r *LandingPagesService) List(profileId int64, campaignId int64) *LandingPagesListCall {
	return &LandingPagesListCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LandingPagesListCall) Fields(s ...googleapi.Field) *LandingPagesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LandingPagesListCall) Do() (*LandingPagesListResponse, error) {
	var returnValue *LandingPagesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves the list of landing pages for the specified campaign.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.landingPages.list",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Landing page campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	//   "response": {
	//     "$ref": "LandingPagesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.landingPages.patch":

type LandingPagesPatchCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	id            int64
	landingpage   *LandingPage
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing campaign landing page. This method
// supports patch semantics.

func (r *LandingPagesService) Patch(profileId int64, campaignId int64, id int64, landingpage *LandingPage) *LandingPagesPatchCall {
	return &LandingPagesPatchCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		id:            id,
		landingpage:   landingpage,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LandingPagesPatchCall) Fields(s ...googleapi.Field) *LandingPagesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LandingPagesPatchCall) Do() (*LandingPage, error) {
	var returnValue *LandingPage
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.landingpage,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing campaign landing page. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.landingPages.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Landing page campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "Landing page ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	//   "request": {
	//     "$ref": "LandingPage"
	//   },
	//   "response": {
	//     "$ref": "LandingPage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.landingPages.update":

type LandingPagesUpdateCall struct {
	s             *Service
	profileId     int64
	campaignId    int64
	landingpage   *LandingPage
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing campaign landing page.

func (r *LandingPagesService) Update(profileId int64, campaignId int64, landingpage *LandingPage) *LandingPagesUpdateCall {
	return &LandingPagesUpdateCall{
		s:             r.s,
		profileId:     profileId,
		campaignId:    campaignId,
		landingpage:   landingpage,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LandingPagesUpdateCall) Fields(s ...googleapi.Field) *LandingPagesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *LandingPagesUpdateCall) Do() (*LandingPage, error) {
	var returnValue *LandingPage
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId":  strconv.FormatInt(c.profileId, 10),
		"campaignId": strconv.FormatInt(c.campaignId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.landingpage,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing campaign landing page.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.landingPages.update",
	//   "parameterOrder": [
	//     "profileId",
	//     "campaignId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Landing page campaign ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/campaigns/{campaignId}/landingPages",
	//   "request": {
	//     "$ref": "LandingPage"
	//   },
	//   "response": {
	//     "$ref": "LandingPage"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.metros.list":

type MetrosListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of metros.

func (r *MetrosService) List(profileId int64) *MetrosListCall {
	return &MetrosListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/metros",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MetrosListCall) Fields(s ...googleapi.Field) *MetrosListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MetrosListCall) Do() (*MetrosListResponse, error) {
	var returnValue *MetrosListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of metros.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.metros.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/metros",
	//   "response": {
	//     "$ref": "MetrosListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.mobileCarriers.list":

type MobileCarriersListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of mobile carriers.

func (r *MobileCarriersService) List(profileId int64) *MobileCarriersListCall {
	return &MobileCarriersListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/mobileCarriers",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MobileCarriersListCall) Fields(s ...googleapi.Field) *MobileCarriersListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *MobileCarriersListCall) Do() (*MobileCarriersListResponse, error) {
	var returnValue *MobileCarriersListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of mobile carriers.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.mobileCarriers.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/mobileCarriers",
	//   "response": {
	//     "$ref": "MobileCarriersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.operatingSystemVersions.list":

type OperatingSystemVersionsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of operating system versions.

func (r *OperatingSystemVersionsService) List(profileId int64) *OperatingSystemVersionsListCall {
	return &OperatingSystemVersionsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/operatingSystemVersions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperatingSystemVersionsListCall) Fields(s ...googleapi.Field) *OperatingSystemVersionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OperatingSystemVersionsListCall) Do() (*OperatingSystemVersionsListResponse, error) {
	var returnValue *OperatingSystemVersionsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of operating system versions.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.operatingSystemVersions.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/operatingSystemVersions",
	//   "response": {
	//     "$ref": "OperatingSystemVersionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.operatingSystems.list":

type OperatingSystemsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of operating systems.

func (r *OperatingSystemsService) List(profileId int64) *OperatingSystemsListCall {
	return &OperatingSystemsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/operatingSystems",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperatingSystemsListCall) Fields(s ...googleapi.Field) *OperatingSystemsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *OperatingSystemsListCall) Do() (*OperatingSystemsListResponse, error) {
	var returnValue *OperatingSystemsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of operating systems.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.operatingSystems.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/operatingSystems",
	//   "response": {
	//     "$ref": "OperatingSystemsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementGroups.get":

type PlacementGroupsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one placement group by ID.

func (r *PlacementGroupsService) Get(profileId int64, id int64) *PlacementGroupsGetCall {
	return &PlacementGroupsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placementGroups/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementGroupsGetCall) Fields(s ...googleapi.Field) *PlacementGroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementGroupsGetCall) Do() (*PlacementGroup, error) {
	var returnValue *PlacementGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one placement group by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.placementGroups.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementGroups/{id}",
	//   "response": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementGroups.insert":

type PlacementGroupsInsertCall struct {
	s              *Service
	profileId      int64
	placementgroup *PlacementGroup
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Insert: Inserts a new placement group.

func (r *PlacementGroupsService) Insert(profileId int64, placementgroup *PlacementGroup) *PlacementGroupsInsertCall {
	return &PlacementGroupsInsertCall{
		s:              r.s,
		profileId:      profileId,
		placementgroup: placementgroup,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "userprofiles/{profileId}/placementGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementGroupsInsertCall) Fields(s ...googleapi.Field) *PlacementGroupsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementGroupsInsertCall) Do() (*PlacementGroup, error) {
	var returnValue *PlacementGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.placementgroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new placement group.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.placementGroups.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementGroups",
	//   "request": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "response": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementGroups.list":

type PlacementGroupsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of placement groups, possibly filtered.

func (r *PlacementGroupsService) List(profileId int64) *PlacementGroupsListCall {
	return &PlacementGroupsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placementGroups",
	}
}

// AdvertiserIds sets the optional parameter "advertiserIds": Select
// only placement groups that belong to these advertisers.
func (c *PlacementGroupsListCall) AdvertiserIds(advertiserIds ...int64) *PlacementGroupsListCall {
	for _, v_ := range advertiserIds {
		c.params_.Add("advertiserIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Archived sets the optional parameter "archived": Select only archived
// placements. Don't set this field to select both archived and
// non-archived placements.
func (c *PlacementGroupsListCall) Archived(archived bool) *PlacementGroupsListCall {
	c.params_.Set("archived", fmt.Sprintf("%v", archived))
	return c
}

// CampaignIds sets the optional parameter "campaignIds": Select only
// placement groups that belong to these campaigns.
func (c *PlacementGroupsListCall) CampaignIds(campaignIds ...int64) *PlacementGroupsListCall {
	for _, v_ := range campaignIds {
		c.params_.Add("campaignIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// ContentCategoryIds sets the optional parameter "contentCategoryIds":
// Select only placement groups that are associated with these content
// categories.
func (c *PlacementGroupsListCall) ContentCategoryIds(contentCategoryIds ...int64) *PlacementGroupsListCall {
	for _, v_ := range contentCategoryIds {
		c.params_.Add("contentCategoryIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// DirectorySiteIds sets the optional parameter "directorySiteIds":
// Select only placement groups that are associated with these directory
// sites.
func (c *PlacementGroupsListCall) DirectorySiteIds(directorySiteIds ...int64) *PlacementGroupsListCall {
	for _, v_ := range directorySiteIds {
		c.params_.Add("directorySiteIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only placement groups
// with these IDs.
func (c *PlacementGroupsListCall) Ids(ids ...int64) *PlacementGroupsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *PlacementGroupsListCall) MaxResults(maxResults int64) *PlacementGroupsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *PlacementGroupsListCall) PageToken(pageToken string) *PlacementGroupsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// PlacementGroupType sets the optional parameter "placementGroupType":
// Select only placement groups belonging with this group type. A
// package is a simple group of placements that acts as a single pricing
// point for a group of tags. A roadblock is a group of placements that
// not only acts as a single pricing point but also assumes that all the
// tags in it will be served at the same time. A roadblock requires one
// of its assigned placements to be marked as primary for reporting.
func (c *PlacementGroupsListCall) PlacementGroupType(placementGroupType string) *PlacementGroupsListCall {
	c.params_.Set("placementGroupType", fmt.Sprintf("%v", placementGroupType))
	return c
}

// PlacementStrategyIds sets the optional parameter
// "placementStrategyIds": Select only placement groups that are
// associated with these placement strategies.
func (c *PlacementGroupsListCall) PlacementStrategyIds(placementStrategyIds ...int64) *PlacementGroupsListCall {
	for _, v_ := range placementStrategyIds {
		c.params_.Add("placementStrategyIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// PricingTypes sets the optional parameter "pricingTypes": Select only
// placement groups with these pricing types.
func (c *PlacementGroupsListCall) PricingTypes(pricingTypes ...string) *PlacementGroupsListCall {
	c.params_["pricingTypes"] = pricingTypes
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for placement groups by name or ID. Wildcards (*) are
// allowed. For example, "placement*2015" will return placement groups
// with names like "placement group June 2015", "placement group May
// 2015" or simply "placements 2015". Most of the searches also add
// wildcards implicitly at the start and the end of the search string.
// For example, a search string of "placementgroup" will match placement
// groups with name "my placementgroup", "placementgroup 2015" or simply
// "placementgroup".
func (c *PlacementGroupsListCall) SearchString(searchString string) *PlacementGroupsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SiteIds sets the optional parameter "siteIds": Select only placement
// groups that are associated with these sites.
func (c *PlacementGroupsListCall) SiteIds(siteIds ...int64) *PlacementGroupsListCall {
	for _, v_ := range siteIds {
		c.params_.Add("siteIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *PlacementGroupsListCall) SortField(sortField string) *PlacementGroupsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *PlacementGroupsListCall) SortOrder(sortOrder string) *PlacementGroupsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementGroupsListCall) Fields(s ...googleapi.Field) *PlacementGroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementGroupsListCall) Do() (*PlacementGroupsListResponse, error) {
	var returnValue *PlacementGroupsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of placement groups, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.placementGroups.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserIds": {
	//       "description": "Select only placement groups that belong to these advertisers.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "archived": {
	//       "description": "Select only archived placements. Don't set this field to select both archived and non-archived placements.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "campaignIds": {
	//       "description": "Select only placement groups that belong to these campaigns.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "contentCategoryIds": {
	//       "description": "Select only placement groups that are associated with these content categories.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "directorySiteIds": {
	//       "description": "Select only placement groups that are associated with these directory sites.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only placement groups with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "placementGroupType": {
	//       "description": "Select only placement groups belonging with this group type. A package is a simple group of placements that acts as a single pricing point for a group of tags. A roadblock is a group of placements that not only acts as a single pricing point but also assumes that all the tags in it will be served at the same time. A roadblock requires one of its assigned placements to be marked as primary for reporting.",
	//       "enum": [
	//         "PLACEMENT_PACKAGE",
	//         "PLACEMENT_ROADBLOCK"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "placementStrategyIds": {
	//       "description": "Select only placement groups that are associated with these placement strategies.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "pricingTypes": {
	//       "description": "Select only placement groups with these pricing types.",
	//       "enum": [
	//         "PRICING_TYPE_CPA",
	//         "PRICING_TYPE_CPC",
	//         "PRICING_TYPE_CPM",
	//         "PRICING_TYPE_FLAT_RATE_CLICKS",
	//         "PRICING_TYPE_FLAT_RATE_IMPRESSIONS"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for placement groups by name or ID. Wildcards (*) are allowed. For example, \"placement*2015\" will return placement groups with names like \"placement group June 2015\", \"placement group May 2015\" or simply \"placements 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"placementgroup\" will match placement groups with name \"my placementgroup\", \"placementgroup 2015\" or simply \"placementgroup\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "siteIds": {
	//       "description": "Select only placement groups that are associated with these sites.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementGroups",
	//   "response": {
	//     "$ref": "PlacementGroupsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementGroups.patch":

type PlacementGroupsPatchCall struct {
	s              *Service
	profileId      int64
	id             int64
	placementgroup *PlacementGroup
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Patch: Updates an existing placement group. This method supports
// patch semantics.

func (r *PlacementGroupsService) Patch(profileId int64, id int64, placementgroup *PlacementGroup) *PlacementGroupsPatchCall {
	return &PlacementGroupsPatchCall{
		s:              r.s,
		profileId:      profileId,
		id:             id,
		placementgroup: placementgroup,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "userprofiles/{profileId}/placementGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementGroupsPatchCall) Fields(s ...googleapi.Field) *PlacementGroupsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementGroupsPatchCall) Do() (*PlacementGroup, error) {
	var returnValue *PlacementGroup
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.placementgroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing placement group. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.placementGroups.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement group ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementGroups",
	//   "request": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "response": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementGroups.update":

type PlacementGroupsUpdateCall struct {
	s              *Service
	profileId      int64
	placementgroup *PlacementGroup
	caller_        googleapi.Caller
	params_        url.Values
	pathTemplate_  string
}

// Update: Updates an existing placement group.

func (r *PlacementGroupsService) Update(profileId int64, placementgroup *PlacementGroup) *PlacementGroupsUpdateCall {
	return &PlacementGroupsUpdateCall{
		s:              r.s,
		profileId:      profileId,
		placementgroup: placementgroup,
		caller_:        googleapi.JSONCall{},
		params_:        make(map[string][]string),
		pathTemplate_:  "userprofiles/{profileId}/placementGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementGroupsUpdateCall) Fields(s ...googleapi.Field) *PlacementGroupsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementGroupsUpdateCall) Do() (*PlacementGroup, error) {
	var returnValue *PlacementGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.placementgroup,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing placement group.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.placementGroups.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementGroups",
	//   "request": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "response": {
	//     "$ref": "PlacementGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementStrategies.delete":

type PlacementStrategiesDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing placement strategy.

func (r *PlacementStrategiesService) Delete(profileId int64, id int64) *PlacementStrategiesDeleteCall {
	return &PlacementStrategiesDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placementStrategies/{id}",
	}
}

func (c *PlacementStrategiesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing placement strategy.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.placementStrategies.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement strategy ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementStrategies/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementStrategies.get":

type PlacementStrategiesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one placement strategy by ID.

func (r *PlacementStrategiesService) Get(profileId int64, id int64) *PlacementStrategiesGetCall {
	return &PlacementStrategiesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placementStrategies/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementStrategiesGetCall) Fields(s ...googleapi.Field) *PlacementStrategiesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementStrategiesGetCall) Do() (*PlacementStrategy, error) {
	var returnValue *PlacementStrategy
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one placement strategy by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.placementStrategies.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement strategy ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementStrategies/{id}",
	//   "response": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementStrategies.insert":

type PlacementStrategiesInsertCall struct {
	s                 *Service
	profileId         int64
	placementstrategy *PlacementStrategy
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// Insert: Inserts a new placement strategy.

func (r *PlacementStrategiesService) Insert(profileId int64, placementstrategy *PlacementStrategy) *PlacementStrategiesInsertCall {
	return &PlacementStrategiesInsertCall{
		s:                 r.s,
		profileId:         profileId,
		placementstrategy: placementstrategy,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "userprofiles/{profileId}/placementStrategies",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementStrategiesInsertCall) Fields(s ...googleapi.Field) *PlacementStrategiesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementStrategiesInsertCall) Do() (*PlacementStrategy, error) {
	var returnValue *PlacementStrategy
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.placementstrategy,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new placement strategy.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.placementStrategies.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementStrategies",
	//   "request": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "response": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementStrategies.list":

type PlacementStrategiesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of placement strategies, possibly filtered.

func (r *PlacementStrategiesService) List(profileId int64) *PlacementStrategiesListCall {
	return &PlacementStrategiesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placementStrategies",
	}
}

// Ids sets the optional parameter "ids": Select only placement
// strategies with these IDs.
func (c *PlacementStrategiesListCall) Ids(ids ...int64) *PlacementStrategiesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *PlacementStrategiesListCall) MaxResults(maxResults int64) *PlacementStrategiesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *PlacementStrategiesListCall) PageToken(pageToken string) *PlacementStrategiesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "placementstrategy*2015" will return objects with names like
// "placementstrategy June 2015", "placementstrategy April 2015" or
// simply "placementstrategy 2015". Most of the searches also add
// wildcards implicitly at the start and the end of the search string.
// For example, a search string of "placementstrategy" will match
// objects with name "my placementstrategy", "placementstrategy 2015" or
// simply "placementstrategy".
func (c *PlacementStrategiesListCall) SearchString(searchString string) *PlacementStrategiesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *PlacementStrategiesListCall) SortField(sortField string) *PlacementStrategiesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *PlacementStrategiesListCall) SortOrder(sortOrder string) *PlacementStrategiesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementStrategiesListCall) Fields(s ...googleapi.Field) *PlacementStrategiesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementStrategiesListCall) Do() (*PlacementStrategiesListResponse, error) {
	var returnValue *PlacementStrategiesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of placement strategies, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.placementStrategies.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "ids": {
	//       "description": "Select only placement strategies with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"placementstrategy*2015\" will return objects with names like \"placementstrategy June 2015\", \"placementstrategy April 2015\" or simply \"placementstrategy 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"placementstrategy\" will match objects with name \"my placementstrategy\", \"placementstrategy 2015\" or simply \"placementstrategy\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementStrategies",
	//   "response": {
	//     "$ref": "PlacementStrategiesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementStrategies.patch":

type PlacementStrategiesPatchCall struct {
	s                 *Service
	profileId         int64
	id                int64
	placementstrategy *PlacementStrategy
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// Patch: Updates an existing placement strategy. This method supports
// patch semantics.

func (r *PlacementStrategiesService) Patch(profileId int64, id int64, placementstrategy *PlacementStrategy) *PlacementStrategiesPatchCall {
	return &PlacementStrategiesPatchCall{
		s:                 r.s,
		profileId:         profileId,
		id:                id,
		placementstrategy: placementstrategy,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "userprofiles/{profileId}/placementStrategies",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementStrategiesPatchCall) Fields(s ...googleapi.Field) *PlacementStrategiesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementStrategiesPatchCall) Do() (*PlacementStrategy, error) {
	var returnValue *PlacementStrategy
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.placementstrategy,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing placement strategy. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.placementStrategies.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement strategy ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementStrategies",
	//   "request": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "response": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placementStrategies.update":

type PlacementStrategiesUpdateCall struct {
	s                 *Service
	profileId         int64
	placementstrategy *PlacementStrategy
	caller_           googleapi.Caller
	params_           url.Values
	pathTemplate_     string
}

// Update: Updates an existing placement strategy.

func (r *PlacementStrategiesService) Update(profileId int64, placementstrategy *PlacementStrategy) *PlacementStrategiesUpdateCall {
	return &PlacementStrategiesUpdateCall{
		s:                 r.s,
		profileId:         profileId,
		placementstrategy: placementstrategy,
		caller_:           googleapi.JSONCall{},
		params_:           make(map[string][]string),
		pathTemplate_:     "userprofiles/{profileId}/placementStrategies",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementStrategiesUpdateCall) Fields(s ...googleapi.Field) *PlacementStrategiesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementStrategiesUpdateCall) Do() (*PlacementStrategy, error) {
	var returnValue *PlacementStrategy
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.placementstrategy,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing placement strategy.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.placementStrategies.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placementStrategies",
	//   "request": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "response": {
	//     "$ref": "PlacementStrategy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placements.generatetags":

type PlacementsGeneratetagsCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Generatetags: Generates tags for a placement.

func (r *PlacementsService) Generatetags(profileId int64) *PlacementsGeneratetagsCall {
	return &PlacementsGeneratetagsCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placements/generatetags",
	}
}

// CampaignId sets the optional parameter "campaignId": Generate
// placements belonging to this campaign. This is a required field.
func (c *PlacementsGeneratetagsCall) CampaignId(campaignId int64) *PlacementsGeneratetagsCall {
	c.params_.Set("campaignId", fmt.Sprintf("%v", campaignId))
	return c
}

// PlacementIds sets the optional parameter "placementIds": Generate
// tags for these placements.
func (c *PlacementsGeneratetagsCall) PlacementIds(placementIds ...int64) *PlacementsGeneratetagsCall {
	for _, v_ := range placementIds {
		c.params_.Add("placementIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// TagFormats sets the optional parameter "tagFormats": Tag formats to
// generate for these placements.
func (c *PlacementsGeneratetagsCall) TagFormats(tagFormats ...string) *PlacementsGeneratetagsCall {
	c.params_["tagFormats"] = tagFormats
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementsGeneratetagsCall) Fields(s ...googleapi.Field) *PlacementsGeneratetagsCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementsGeneratetagsCall) Do() (*PlacementsGenerateTagsResponse, error) {
	var returnValue *PlacementsGenerateTagsResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Generates tags for a placement.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.placements.generatetags",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "campaignId": {
	//       "description": "Generate placements belonging to this campaign. This is a required field.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "placementIds": {
	//       "description": "Generate tags for these placements.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "tagFormats": {
	//       "description": "Tag formats to generate for these placements.",
	//       "enum": [
	//         "PLACEMENT_TAG_CLICK_COMMANDS",
	//         "PLACEMENT_TAG_IFRAME_ILAYER",
	//         "PLACEMENT_TAG_IFRAME_JAVASCRIPT",
	//         "PLACEMENT_TAG_INSTREAM_VIDEO_PREFETCH",
	//         "PLACEMENT_TAG_INTERNAL_REDIRECT",
	//         "PLACEMENT_TAG_INTERSTITIAL_IFRAME_JAVASCRIPT",
	//         "PLACEMENT_TAG_INTERSTITIAL_INTERNAL_REDIRECT",
	//         "PLACEMENT_TAG_INTERSTITIAL_JAVASCRIPT",
	//         "PLACEMENT_TAG_JAVASCRIPT",
	//         "PLACEMENT_TAG_STANDARD",
	//         "PLACEMENT_TAG_TRACKING",
	//         "PLACEMENT_TAG_TRACKING_IFRAME",
	//         "PLACEMENT_TAG_TRACKING_JAVASCRIPT"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placements/generatetags",
	//   "response": {
	//     "$ref": "PlacementsGenerateTagsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placements.get":

type PlacementsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one placement by ID.

func (r *PlacementsService) Get(profileId int64, id int64) *PlacementsGetCall {
	return &PlacementsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placements/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementsGetCall) Fields(s ...googleapi.Field) *PlacementsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementsGetCall) Do() (*Placement, error) {
	var returnValue *Placement
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one placement by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.placements.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placements/{id}",
	//   "response": {
	//     "$ref": "Placement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placements.insert":

type PlacementsInsertCall struct {
	s             *Service
	profileId     int64
	placement     *Placement
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new placement.

func (r *PlacementsService) Insert(profileId int64, placement *Placement) *PlacementsInsertCall {
	return &PlacementsInsertCall{
		s:             r.s,
		profileId:     profileId,
		placement:     placement,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placements",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementsInsertCall) Fields(s ...googleapi.Field) *PlacementsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementsInsertCall) Do() (*Placement, error) {
	var returnValue *Placement
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.placement,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new placement.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.placements.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placements",
	//   "request": {
	//     "$ref": "Placement"
	//   },
	//   "response": {
	//     "$ref": "Placement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placements.list":

type PlacementsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of placements, possibly filtered.

func (r *PlacementsService) List(profileId int64) *PlacementsListCall {
	return &PlacementsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placements",
	}
}

// AdvertiserIds sets the optional parameter "advertiserIds": Select
// only placements that belong to these advertisers.
func (c *PlacementsListCall) AdvertiserIds(advertiserIds ...int64) *PlacementsListCall {
	for _, v_ := range advertiserIds {
		c.params_.Add("advertiserIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Archived sets the optional parameter "archived": Select only archived
// placements. Don't set this field to select both archived and
// non-archived placements.
func (c *PlacementsListCall) Archived(archived bool) *PlacementsListCall {
	c.params_.Set("archived", fmt.Sprintf("%v", archived))
	return c
}

// CampaignIds sets the optional parameter "campaignIds": Select only
// placements that belong to these campaigns.
func (c *PlacementsListCall) CampaignIds(campaignIds ...int64) *PlacementsListCall {
	for _, v_ := range campaignIds {
		c.params_.Add("campaignIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Compatibilities sets the optional parameter "compatibilities": Select
// only placements that are associated with these compatibilities. WEB
// and WEB_INTERSTITIAL refer to rendering either on desktop or on
// mobile devices for regular or interstitial ads respectively. APP and
// APP_INTERSTITIAL are for rendering in mobile apps.IN_STREAM_VIDEO
// refers to rendering in in-stream video ads developed with the VAST
// standard.
func (c *PlacementsListCall) Compatibilities(compatibilities ...string) *PlacementsListCall {
	c.params_["compatibilities"] = compatibilities
	return c
}

// ContentCategoryIds sets the optional parameter "contentCategoryIds":
// Select only placements that are associated with these content
// categories.
func (c *PlacementsListCall) ContentCategoryIds(contentCategoryIds ...int64) *PlacementsListCall {
	for _, v_ := range contentCategoryIds {
		c.params_.Add("contentCategoryIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// DirectorySiteIds sets the optional parameter "directorySiteIds":
// Select only placements that are associated with these directory
// sites.
func (c *PlacementsListCall) DirectorySiteIds(directorySiteIds ...int64) *PlacementsListCall {
	for _, v_ := range directorySiteIds {
		c.params_.Add("directorySiteIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// GroupIds sets the optional parameter "groupIds": Select only
// placements that belong to these placement groups.
func (c *PlacementsListCall) GroupIds(groupIds ...int64) *PlacementsListCall {
	for _, v_ := range groupIds {
		c.params_.Add("groupIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only placements with
// these IDs.
func (c *PlacementsListCall) Ids(ids ...int64) *PlacementsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *PlacementsListCall) MaxResults(maxResults int64) *PlacementsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *PlacementsListCall) PageToken(pageToken string) *PlacementsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// PaymentSource sets the optional parameter "paymentSource": Select
// only placements with this payment source.
func (c *PlacementsListCall) PaymentSource(paymentSource string) *PlacementsListCall {
	c.params_.Set("paymentSource", fmt.Sprintf("%v", paymentSource))
	return c
}

// PlacementStrategyIds sets the optional parameter
// "placementStrategyIds": Select only placements that are associated
// with these placement strategies.
func (c *PlacementsListCall) PlacementStrategyIds(placementStrategyIds ...int64) *PlacementsListCall {
	for _, v_ := range placementStrategyIds {
		c.params_.Add("placementStrategyIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// PricingTypes sets the optional parameter "pricingTypes": Select only
// placements with these pricing types.
func (c *PlacementsListCall) PricingTypes(pricingTypes ...string) *PlacementsListCall {
	c.params_["pricingTypes"] = pricingTypes
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for placements by name or ID. Wildcards (*) are allowed.
// For example, "placement*2015" will return placements with names like
// "placement June 2015", "placement May 2015" or simply "placements
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "placement" will match placements with name "my placement",
// "placement 2015" or simply "placement".
func (c *PlacementsListCall) SearchString(searchString string) *PlacementsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SiteIds sets the optional parameter "siteIds": Select only placements
// that are associated with these sites.
func (c *PlacementsListCall) SiteIds(siteIds ...int64) *PlacementsListCall {
	for _, v_ := range siteIds {
		c.params_.Add("siteIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SizeIds sets the optional parameter "sizeIds": Select only placements
// that are associated with these sizes.
func (c *PlacementsListCall) SizeIds(sizeIds ...int64) *PlacementsListCall {
	for _, v_ := range sizeIds {
		c.params_.Add("sizeIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *PlacementsListCall) SortField(sortField string) *PlacementsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *PlacementsListCall) SortOrder(sortOrder string) *PlacementsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementsListCall) Fields(s ...googleapi.Field) *PlacementsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementsListCall) Do() (*PlacementsListResponse, error) {
	var returnValue *PlacementsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of placements, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.placements.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "advertiserIds": {
	//       "description": "Select only placements that belong to these advertisers.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "archived": {
	//       "description": "Select only archived placements. Don't set this field to select both archived and non-archived placements.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "campaignIds": {
	//       "description": "Select only placements that belong to these campaigns.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "compatibilities": {
	//       "description": "Select only placements that are associated with these compatibilities. WEB and WEB_INTERSTITIAL refer to rendering either on desktop or on mobile devices for regular or interstitial ads respectively. APP and APP_INTERSTITIAL are for rendering in mobile apps.IN_STREAM_VIDEO refers to rendering in in-stream video ads developed with the VAST standard.",
	//       "enum": [
	//         "APP",
	//         "APP_INTERSTITIAL",
	//         "IN_STREAM_VIDEO",
	//         "WEB",
	//         "WEB_INTERSTITIAL"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "contentCategoryIds": {
	//       "description": "Select only placements that are associated with these content categories.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "directorySiteIds": {
	//       "description": "Select only placements that are associated with these directory sites.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "groupIds": {
	//       "description": "Select only placements that belong to these placement groups.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only placements with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "paymentSource": {
	//       "description": "Select only placements with this payment source.",
	//       "enum": [
	//         "PLACEMENT_AGENCY_PAID",
	//         "PLACEMENT_PUBLISHER_PAID"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "placementStrategyIds": {
	//       "description": "Select only placements that are associated with these placement strategies.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "pricingTypes": {
	//       "description": "Select only placements with these pricing types.",
	//       "enum": [
	//         "PRICING_TYPE_CPA",
	//         "PRICING_TYPE_CPC",
	//         "PRICING_TYPE_CPM",
	//         "PRICING_TYPE_FLAT_RATE_CLICKS",
	//         "PRICING_TYPE_FLAT_RATE_IMPRESSIONS"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         "",
	//         "",
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for placements by name or ID. Wildcards (*) are allowed. For example, \"placement*2015\" will return placements with names like \"placement June 2015\", \"placement May 2015\" or simply \"placements 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"placement\" will match placements with name \"my placement\", \"placement 2015\" or simply \"placement\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "siteIds": {
	//       "description": "Select only placements that are associated with these sites.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sizeIds": {
	//       "description": "Select only placements that are associated with these sizes.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placements",
	//   "response": {
	//     "$ref": "PlacementsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placements.patch":

type PlacementsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	placement     *Placement
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing placement. This method supports patch
// semantics.

func (r *PlacementsService) Patch(profileId int64, id int64, placement *Placement) *PlacementsPatchCall {
	return &PlacementsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		placement:     placement,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placements",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementsPatchCall) Fields(s ...googleapi.Field) *PlacementsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementsPatchCall) Do() (*Placement, error) {
	var returnValue *Placement
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.placement,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing placement. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.placements.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Placement ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placements",
	//   "request": {
	//     "$ref": "Placement"
	//   },
	//   "response": {
	//     "$ref": "Placement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.placements.update":

type PlacementsUpdateCall struct {
	s             *Service
	profileId     int64
	placement     *Placement
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing placement.

func (r *PlacementsService) Update(profileId int64, placement *Placement) *PlacementsUpdateCall {
	return &PlacementsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		placement:     placement,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/placements",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlacementsUpdateCall) Fields(s ...googleapi.Field) *PlacementsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlacementsUpdateCall) Do() (*Placement, error) {
	var returnValue *Placement
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.placement,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing placement.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.placements.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/placements",
	//   "request": {
	//     "$ref": "Placement"
	//   },
	//   "response": {
	//     "$ref": "Placement"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.platformTypes.list":

type PlatformTypesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of platform types.

func (r *PlatformTypesService) List(profileId int64) *PlatformTypesListCall {
	return &PlatformTypesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/platformTypes",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PlatformTypesListCall) Fields(s ...googleapi.Field) *PlatformTypesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PlatformTypesListCall) Do() (*PlatformTypesListResponse, error) {
	var returnValue *PlatformTypesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of platform types.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.platformTypes.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/platformTypes",
	//   "response": {
	//     "$ref": "PlatformTypesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.postalCodes.list":

type PostalCodesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of postal codes.

func (r *PostalCodesService) List(profileId int64) *PostalCodesListCall {
	return &PostalCodesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/postalCodes",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *PostalCodesListCall) Fields(s ...googleapi.Field) *PostalCodesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *PostalCodesListCall) Do() (*PostalCodesListResponse, error) {
	var returnValue *PostalCodesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of postal codes.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.postalCodes.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/postalCodes",
	//   "response": {
	//     "$ref": "PostalCodesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.regions.list":

type RegionsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of regions.

func (r *RegionsService) List(profileId int64) *RegionsListCall {
	return &RegionsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/regions",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RegionsListCall) Fields(s ...googleapi.Field) *RegionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *RegionsListCall) Do() (*RegionsListResponse, error) {
	var returnValue *RegionsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of regions.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.regions.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/regions",
	//   "response": {
	//     "$ref": "RegionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.reports.delete":

type ReportsDeleteCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes a report by its ID.

func (r *ReportsService) Delete(profileId int64, reportId int64) *ReportsDeleteCall {
	return &ReportsDeleteCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}",
	}
}

func (c *ReportsDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes a report by its ID.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.reports.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/{reportId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.get":

type ReportsGetCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a report by its ID.

func (r *ReportsService) Get(profileId int64, reportId int64) *ReportsGetCall {
	return &ReportsGetCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}",
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
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a report by its ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.reports.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/{reportId}",
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.insert":

type ReportsInsertCall struct {
	s             *Service
	profileId     int64
	report        *Report
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Creates a report.

func (r *ReportsService) Insert(profileId int64, report *Report) *ReportsInsertCall {
	return &ReportsInsertCall{
		s:             r.s,
		profileId:     profileId,
		report:        report,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsInsertCall) Fields(s ...googleapi.Field) *ReportsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsInsertCall) Do() (*Report, error) {
	var returnValue *Report
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.report,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Creates a report.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.reports.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports",
	//   "request": {
	//     "$ref": "Report"
	//   },
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.list":

type ReportsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves list of reports.

func (r *ReportsService) List(profileId int64) *ReportsListCall {
	return &ReportsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *ReportsListCall) MaxResults(maxResults int64) *ReportsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The value of the
// nextToken from the previous result page.
func (c *ReportsListCall) PageToken(pageToken string) *ReportsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Scope sets the optional parameter "scope": The scope that defines
// which results are returned, default is 'MINE'.
func (c *ReportsListCall) Scope(scope string) *ReportsListCall {
	c.params_.Set("scope", fmt.Sprintf("%v", scope))
	return c
}

// SortField sets the optional parameter "sortField": The field by which
// to sort the list.
func (c *ReportsListCall) SortField(sortField string) *ReportsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is 'DESCENDING'.
func (c *ReportsListCall) SortOrder(sortOrder string) *ReportsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsListCall) Fields(s ...googleapi.Field) *ReportsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsListCall) Do() (*ReportList, error) {
	var returnValue *ReportList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves list of reports.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.reports.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of the nextToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "scope": {
	//       "default": "MINE",
	//       "description": "The scope that defines which results are returned, default is 'MINE'.",
	//       "enum": [
	//         "ALL",
	//         "MINE"
	//       ],
	//       "enumDescriptions": [
	//         "All reports in account.",
	//         "My reports."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "default": "LAST_MODIFIED_TIME",
	//       "description": "The field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "LAST_MODIFIED_TIME",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "Sort by report ID.",
	//         "Sort by 'lastModifiedTime' field.",
	//         "Sort by name of reports."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "default": "DESCENDING",
	//       "description": "Order of sorted results, default is 'DESCENDING'.",
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
	//   "path": "userprofiles/{profileId}/reports",
	//   "response": {
	//     "$ref": "ReportList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.patch":

type ReportsPatchCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	report        *Report
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates a report. This method supports patch semantics.

func (r *ReportsService) Patch(profileId int64, reportId int64, report *Report) *ReportsPatchCall {
	return &ReportsPatchCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		report:        report,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsPatchCall) Fields(s ...googleapi.Field) *ReportsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsPatchCall) Do() (*Report, error) {
	var returnValue *Report
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.report,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a report. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.reports.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/{reportId}",
	//   "request": {
	//     "$ref": "Report"
	//   },
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.run":

type ReportsRunCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Run: Runs a report.

func (r *ReportsService) Run(profileId int64, reportId int64) *ReportsRunCall {
	return &ReportsRunCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}/run",
	}
}

// Synchronous sets the optional parameter "synchronous": If set and
// true, tries to run the report synchronously.
func (c *ReportsRunCall) Synchronous(synchronous bool) *ReportsRunCall {
	c.params_.Set("synchronous", fmt.Sprintf("%v", synchronous))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsRunCall) Fields(s ...googleapi.Field) *ReportsRunCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsRunCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
	})
	call := &googleapi.Call{
		Method: "POST",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Runs a report.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.reports.run",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "synchronous": {
	//       "description": "If set and true, tries to run the report synchronously.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/{reportId}/run",
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.update":

type ReportsUpdateCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	report        *Report
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates a report.

func (r *ReportsService) Update(profileId int64, reportId int64, report *Report) *ReportsUpdateCall {
	return &ReportsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		report:        report,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsUpdateCall) Fields(s ...googleapi.Field) *ReportsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsUpdateCall) Do() (*Report, error) {
	var returnValue *Report
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.report,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates a report.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.reports.update",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/{reportId}",
	//   "request": {
	//     "$ref": "Report"
	//   },
	//   "response": {
	//     "$ref": "Report"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.compatibleFields.query":

type ReportsCompatibleFieldsQueryCall struct {
	s             *Service
	profileId     int64
	report        *Report
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Query: Returns the fields that are compatible to be selected in the
// respective sections of a report criteria, given the fields already
// selected in the input report and user permissions.

func (r *ReportsCompatibleFieldsService) Query(profileId int64, report *Report) *ReportsCompatibleFieldsQueryCall {
	return &ReportsCompatibleFieldsQueryCall{
		s:             r.s,
		profileId:     profileId,
		report:        report,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/compatiblefields/query",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsCompatibleFieldsQueryCall) Fields(s ...googleapi.Field) *ReportsCompatibleFieldsQueryCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsCompatibleFieldsQueryCall) Do() (*CompatibleFields, error) {
	var returnValue *CompatibleFields
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.report,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Returns the fields that are compatible to be selected in the respective sections of a report criteria, given the fields already selected in the input report and user permissions.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.reports.compatibleFields.query",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The DFA user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/compatiblefields/query",
	//   "request": {
	//     "$ref": "Report"
	//   },
	//   "response": {
	//     "$ref": "CompatibleFields"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.reports.files.get":

type ReportsFilesGetCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	fileId        int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Retrieves a report file.

func (r *ReportsFilesService) Get(profileId int64, reportId int64, fileId int64) *ReportsFilesGetCall {
	return &ReportsFilesGetCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		fileId:        fileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}/files/{fileId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsFilesGetCall) Fields(s ...googleapi.Field) *ReportsFilesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsFilesGetCall) GetResponse() (*http.Response, error) {
	var res *http.Response
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
		"fileId":    strconv.FormatInt(c.fileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &res,
	}

	return res, c.caller_.Do(googleapi.NoContext, c.s.client, call)
}

func (c *ReportsFilesGetCall) Do() (*File, error) {
	var returnValue *File
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
		"fileId":    strconv.FormatInt(c.fileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a report file.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.reports.files.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId",
	//     "fileId"
	//   ],
	//   "parameters": {
	//     "fileId": {
	//       "description": "The ID of the report file.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "The DFA profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/reports/{reportId}/files/{fileId}",
	//   "response": {
	//     "$ref": "File"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "dfareporting.reports.files.list":

type ReportsFilesListCall struct {
	s             *Service
	profileId     int64
	reportId      int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Lists files for a report.

func (r *ReportsFilesService) List(profileId int64, reportId int64) *ReportsFilesListCall {
	return &ReportsFilesListCall{
		s:             r.s,
		profileId:     profileId,
		reportId:      reportId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/reports/{reportId}/files",
	}
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *ReportsFilesListCall) MaxResults(maxResults int64) *ReportsFilesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The value of the
// nextToken from the previous result page.
func (c *ReportsFilesListCall) PageToken(pageToken string) *ReportsFilesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SortField sets the optional parameter "sortField": The field by which
// to sort the list.
func (c *ReportsFilesListCall) SortField(sortField string) *ReportsFilesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is 'DESCENDING'.
func (c *ReportsFilesListCall) SortOrder(sortOrder string) *ReportsFilesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsFilesListCall) Fields(s ...googleapi.Field) *ReportsFilesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *ReportsFilesListCall) Do() (*FileList, error) {
	var returnValue *FileList
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"reportId":  strconv.FormatInt(c.reportId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Lists files for a report.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.reports.files.list",
	//   "parameterOrder": [
	//     "profileId",
	//     "reportId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "maximum": "10",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of the nextToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "The DFA profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "reportId": {
	//       "description": "The ID of the parent report.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "default": "LAST_MODIFIED_TIME",
	//       "description": "The field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "LAST_MODIFIED_TIME"
	//       ],
	//       "enumDescriptions": [
	//         "Sort by file ID.",
	//         "Sort by 'lastmodifiedAt' field."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "default": "DESCENDING",
	//       "description": "Order of sorted results, default is 'DESCENDING'.",
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
	//   "path": "userprofiles/{profileId}/reports/{reportId}/files",
	//   "response": {
	//     "$ref": "FileList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting"
	//   ]
	// }

}

// method id "dfareporting.sites.get":

type SitesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one site by ID.

func (r *SitesService) Get(profileId int64, id int64) *SitesGetCall {
	return &SitesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sites/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesGetCall) Fields(s ...googleapi.Field) *SitesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesGetCall) Do() (*Site, error) {
	var returnValue *Site
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one site by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.sites.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Site ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sites/{id}",
	//   "response": {
	//     "$ref": "Site"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sites.insert":

type SitesInsertCall struct {
	s             *Service
	profileId     int64
	site          *Site
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new site.

func (r *SitesService) Insert(profileId int64, site *Site) *SitesInsertCall {
	return &SitesInsertCall{
		s:             r.s,
		profileId:     profileId,
		site:          site,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sites",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesInsertCall) Fields(s ...googleapi.Field) *SitesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesInsertCall) Do() (*Site, error) {
	var returnValue *Site
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.site,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new site.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.sites.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sites",
	//   "request": {
	//     "$ref": "Site"
	//   },
	//   "response": {
	//     "$ref": "Site"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sites.list":

type SitesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of sites, possibly filtered.

func (r *SitesService) List(profileId int64) *SitesListCall {
	return &SitesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sites",
	}
}

// AcceptsInStreamVideoPlacements sets the optional parameter
// "acceptsInStreamVideoPlacements": This search filter is no longer
// supported and will have no effect on the results returned.
func (c *SitesListCall) AcceptsInStreamVideoPlacements(acceptsInStreamVideoPlacements bool) *SitesListCall {
	c.params_.Set("acceptsInStreamVideoPlacements", fmt.Sprintf("%v", acceptsInStreamVideoPlacements))
	return c
}

// AcceptsInterstitialPlacements sets the optional parameter
// "acceptsInterstitialPlacements": This search filter is no longer
// supported and will have no effect on the results returned.
func (c *SitesListCall) AcceptsInterstitialPlacements(acceptsInterstitialPlacements bool) *SitesListCall {
	c.params_.Set("acceptsInterstitialPlacements", fmt.Sprintf("%v", acceptsInterstitialPlacements))
	return c
}

// AcceptsPublisherPaidPlacements sets the optional parameter
// "acceptsPublisherPaidPlacements": Select only sites that accept
// publisher paid placements.
func (c *SitesListCall) AcceptsPublisherPaidPlacements(acceptsPublisherPaidPlacements bool) *SitesListCall {
	c.params_.Set("acceptsPublisherPaidPlacements", fmt.Sprintf("%v", acceptsPublisherPaidPlacements))
	return c
}

// AdWordsSite sets the optional parameter "adWordsSite": Select only
// AdWords sites.
func (c *SitesListCall) AdWordsSite(adWordsSite bool) *SitesListCall {
	c.params_.Set("adWordsSite", fmt.Sprintf("%v", adWordsSite))
	return c
}

// Approved sets the optional parameter "approved": Select only approved
// sites.
func (c *SitesListCall) Approved(approved bool) *SitesListCall {
	c.params_.Set("approved", fmt.Sprintf("%v", approved))
	return c
}

// CampaignIds sets the optional parameter "campaignIds": Select only
// sites with these campaign IDs.
func (c *SitesListCall) CampaignIds(campaignIds ...int64) *SitesListCall {
	for _, v_ := range campaignIds {
		c.params_.Add("campaignIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// DirectorySiteIds sets the optional parameter "directorySiteIds":
// Select only sites with these directory site IDs.
func (c *SitesListCall) DirectorySiteIds(directorySiteIds ...int64) *SitesListCall {
	for _, v_ := range directorySiteIds {
		c.params_.Add("directorySiteIds", fmt.Sprintf("%v", v_))
	}
	return c
}

// Ids sets the optional parameter "ids": Select only sites with these
// IDs.
func (c *SitesListCall) Ids(ids ...int64) *SitesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *SitesListCall) MaxResults(maxResults int64) *SitesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *SitesListCall) PageToken(pageToken string) *SitesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name, ID or keyName. Wildcards (*) are
// allowed. For example, "site*2015" will return objects with names like
// "site June 2015", "site April 2015" or simply "site 2015". Most of
// the searches also add wildcards implicitly at the start and the end
// of the search string. For example, a search string of "site" will
// match objects with name "my site", "site 2015" or simply "site".
func (c *SitesListCall) SearchString(searchString string) *SitesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *SitesListCall) SortField(sortField string) *SitesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *SitesListCall) SortOrder(sortOrder string) *SitesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// SubaccountId sets the optional parameter "subaccountId": Select only
// sites with this subaccount ID.
func (c *SitesListCall) SubaccountId(subaccountId int64) *SitesListCall {
	c.params_.Set("subaccountId", fmt.Sprintf("%v", subaccountId))
	return c
}

// UnmappedSite sets the optional parameter "unmappedSite": Select only
// sites that have not been mapped to a directory site.
func (c *SitesListCall) UnmappedSite(unmappedSite bool) *SitesListCall {
	c.params_.Set("unmappedSite", fmt.Sprintf("%v", unmappedSite))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesListCall) Fields(s ...googleapi.Field) *SitesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesListCall) Do() (*SitesListResponse, error) {
	var returnValue *SitesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of sites, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.sites.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "acceptsInStreamVideoPlacements": {
	//       "description": "This search filter is no longer supported and will have no effect on the results returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "acceptsInterstitialPlacements": {
	//       "description": "This search filter is no longer supported and will have no effect on the results returned.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "acceptsPublisherPaidPlacements": {
	//       "description": "Select only sites that accept publisher paid placements.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "adWordsSite": {
	//       "description": "Select only AdWords sites.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "approved": {
	//       "description": "Select only approved sites.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "campaignIds": {
	//       "description": "Select only sites with these campaign IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "directorySiteIds": {
	//       "description": "Select only sites with these directory site IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Select only sites with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name, ID or keyName. Wildcards (*) are allowed. For example, \"site*2015\" will return objects with names like \"site June 2015\", \"site April 2015\" or simply \"site 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"site\" will match objects with name \"my site\", \"site 2015\" or simply \"site\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "subaccountId": {
	//       "description": "Select only sites with this subaccount ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "unmappedSite": {
	//       "description": "Select only sites that have not been mapped to a directory site.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sites",
	//   "response": {
	//     "$ref": "SitesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sites.patch":

type SitesPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	site          *Site
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing site. This method supports patch
// semantics.

func (r *SitesService) Patch(profileId int64, id int64, site *Site) *SitesPatchCall {
	return &SitesPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		site:          site,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sites",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesPatchCall) Fields(s ...googleapi.Field) *SitesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesPatchCall) Do() (*Site, error) {
	var returnValue *Site
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.site,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing site. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.sites.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Site ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sites",
	//   "request": {
	//     "$ref": "Site"
	//   },
	//   "response": {
	//     "$ref": "Site"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sites.update":

type SitesUpdateCall struct {
	s             *Service
	profileId     int64
	site          *Site
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing site.

func (r *SitesService) Update(profileId int64, site *Site) *SitesUpdateCall {
	return &SitesUpdateCall{
		s:             r.s,
		profileId:     profileId,
		site:          site,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sites",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SitesUpdateCall) Fields(s ...googleapi.Field) *SitesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SitesUpdateCall) Do() (*Site, error) {
	var returnValue *Site
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.site,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing site.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.sites.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sites",
	//   "request": {
	//     "$ref": "Site"
	//   },
	//   "response": {
	//     "$ref": "Site"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sizes.get":

type SizesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one size by ID.

func (r *SizesService) Get(profileId int64, id int64) *SizesGetCall {
	return &SizesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sizes/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SizesGetCall) Fields(s ...googleapi.Field) *SizesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SizesGetCall) Do() (*Size, error) {
	var returnValue *Size
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one size by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.sizes.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Size ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sizes/{id}",
	//   "response": {
	//     "$ref": "Size"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sizes.insert":

type SizesInsertCall struct {
	s             *Service
	profileId     int64
	size          *Size
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new size.

func (r *SizesService) Insert(profileId int64, size *Size) *SizesInsertCall {
	return &SizesInsertCall{
		s:             r.s,
		profileId:     profileId,
		size:          size,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sizes",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SizesInsertCall) Fields(s ...googleapi.Field) *SizesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SizesInsertCall) Do() (*Size, error) {
	var returnValue *Size
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.size,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new size.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.sizes.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sizes",
	//   "request": {
	//     "$ref": "Size"
	//   },
	//   "response": {
	//     "$ref": "Size"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.sizes.list":

type SizesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of sizes, possibly filtered.

func (r *SizesService) List(profileId int64) *SizesListCall {
	return &SizesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/sizes",
	}
}

// Height sets the optional parameter "height": Select only sizes with
// this height.
func (c *SizesListCall) Height(height int64) *SizesListCall {
	c.params_.Set("height", fmt.Sprintf("%v", height))
	return c
}

// IabStandard sets the optional parameter "iabStandard": Select only
// IAB standard sizes.
func (c *SizesListCall) IabStandard(iabStandard bool) *SizesListCall {
	c.params_.Set("iabStandard", fmt.Sprintf("%v", iabStandard))
	return c
}

// Ids sets the optional parameter "ids": Select only sizes with these
// IDs.
func (c *SizesListCall) Ids(ids ...int64) *SizesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// Width sets the optional parameter "width": Select only sizes with
// this width.
func (c *SizesListCall) Width(width int64) *SizesListCall {
	c.params_.Set("width", fmt.Sprintf("%v", width))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SizesListCall) Fields(s ...googleapi.Field) *SizesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SizesListCall) Do() (*SizesListResponse, error) {
	var returnValue *SizesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of sizes, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.sizes.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "height": {
	//       "description": "Select only sizes with this height.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "iabStandard": {
	//       "description": "Select only IAB standard sizes.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ids": {
	//       "description": "Select only sizes with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "width": {
	//       "description": "Select only sizes with this width.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/sizes",
	//   "response": {
	//     "$ref": "SizesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.subaccounts.get":

type SubaccountsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one subaccount by ID.

func (r *SubaccountsService) Get(profileId int64, id int64) *SubaccountsGetCall {
	return &SubaccountsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/subaccounts/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubaccountsGetCall) Fields(s ...googleapi.Field) *SubaccountsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubaccountsGetCall) Do() (*Subaccount, error) {
	var returnValue *Subaccount
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one subaccount by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.subaccounts.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Subaccount ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/subaccounts/{id}",
	//   "response": {
	//     "$ref": "Subaccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.subaccounts.insert":

type SubaccountsInsertCall struct {
	s             *Service
	profileId     int64
	subaccount    *Subaccount
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new subaccount.

func (r *SubaccountsService) Insert(profileId int64, subaccount *Subaccount) *SubaccountsInsertCall {
	return &SubaccountsInsertCall{
		s:             r.s,
		profileId:     profileId,
		subaccount:    subaccount,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/subaccounts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubaccountsInsertCall) Fields(s ...googleapi.Field) *SubaccountsInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubaccountsInsertCall) Do() (*Subaccount, error) {
	var returnValue *Subaccount
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.subaccount,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new subaccount.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.subaccounts.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/subaccounts",
	//   "request": {
	//     "$ref": "Subaccount"
	//   },
	//   "response": {
	//     "$ref": "Subaccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.subaccounts.list":

type SubaccountsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Gets a list of subaccounts, possibly filtered.

func (r *SubaccountsService) List(profileId int64) *SubaccountsListCall {
	return &SubaccountsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/subaccounts",
	}
}

// Ids sets the optional parameter "ids": Select only subaccounts with
// these IDs.
func (c *SubaccountsListCall) Ids(ids ...int64) *SubaccountsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *SubaccountsListCall) MaxResults(maxResults int64) *SubaccountsListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *SubaccountsListCall) PageToken(pageToken string) *SubaccountsListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "subaccount*2015" will return objects with names like
// "subaccount June 2015", "subaccount April 2015" or simply "subaccount
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "subaccount" will match objects with name "my subaccount",
// "subaccount 2015" or simply "subaccount".
func (c *SubaccountsListCall) SearchString(searchString string) *SubaccountsListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *SubaccountsListCall) SortField(sortField string) *SubaccountsListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *SubaccountsListCall) SortOrder(sortOrder string) *SubaccountsListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubaccountsListCall) Fields(s ...googleapi.Field) *SubaccountsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubaccountsListCall) Do() (*SubaccountsListResponse, error) {
	var returnValue *SubaccountsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of subaccounts, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.subaccounts.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "ids": {
	//       "description": "Select only subaccounts with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"subaccount*2015\" will return objects with names like \"subaccount June 2015\", \"subaccount April 2015\" or simply \"subaccount 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"subaccount\" will match objects with name \"my subaccount\", \"subaccount 2015\" or simply \"subaccount\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/subaccounts",
	//   "response": {
	//     "$ref": "SubaccountsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.subaccounts.patch":

type SubaccountsPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	subaccount    *Subaccount
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing subaccount. This method supports patch
// semantics.

func (r *SubaccountsService) Patch(profileId int64, id int64, subaccount *Subaccount) *SubaccountsPatchCall {
	return &SubaccountsPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		subaccount:    subaccount,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/subaccounts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubaccountsPatchCall) Fields(s ...googleapi.Field) *SubaccountsPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubaccountsPatchCall) Do() (*Subaccount, error) {
	var returnValue *Subaccount
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.subaccount,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing subaccount. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.subaccounts.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "Subaccount ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/subaccounts",
	//   "request": {
	//     "$ref": "Subaccount"
	//   },
	//   "response": {
	//     "$ref": "Subaccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.subaccounts.update":

type SubaccountsUpdateCall struct {
	s             *Service
	profileId     int64
	subaccount    *Subaccount
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing subaccount.

func (r *SubaccountsService) Update(profileId int64, subaccount *Subaccount) *SubaccountsUpdateCall {
	return &SubaccountsUpdateCall{
		s:             r.s,
		profileId:     profileId,
		subaccount:    subaccount,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/subaccounts",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubaccountsUpdateCall) Fields(s ...googleapi.Field) *SubaccountsUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *SubaccountsUpdateCall) Do() (*Subaccount, error) {
	var returnValue *Subaccount
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.subaccount,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing subaccount.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.subaccounts.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/subaccounts",
	//   "request": {
	//     "$ref": "Subaccount"
	//   },
	//   "response": {
	//     "$ref": "Subaccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userProfiles.get":

type UserProfilesGetCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one user profile by ID.

func (r *UserProfilesService) Get(profileId int64) *UserProfilesGetCall {
	return &UserProfilesGetCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserProfilesGetCall) Fields(s ...googleapi.Field) *UserProfilesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserProfilesGetCall) Do() (*UserProfile, error) {
	var returnValue *UserProfile
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one user profile by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userProfiles.get",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "The user profile ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}",
	//   "response": {
	//     "$ref": "UserProfile"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting",
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userProfiles.list":

type UserProfilesListCall struct {
	s             *Service
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves list of user profiles for a user.

func (r *UserProfilesService) List() *UserProfilesListCall {
	return &UserProfilesListCall{
		s:             r.s,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserProfilesListCall) Fields(s ...googleapi.Field) *UserProfilesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserProfilesListCall) Do() (*UserProfileList, error) {
	var returnValue *UserProfileList
	u := googleapi.Expand(baseURL, c.pathTemplate_, nil)
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves list of user profiles for a user.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userProfiles.list",
	//   "path": "userprofiles",
	//   "response": {
	//     "$ref": "UserProfileList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfareporting",
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRolePermissionGroups.get":

type UserRolePermissionGroupsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one user role permission group by ID.

func (r *UserRolePermissionGroupsService) Get(profileId int64, id int64) *UserRolePermissionGroupsGetCall {
	return &UserRolePermissionGroupsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRolePermissionGroups/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolePermissionGroupsGetCall) Fields(s ...googleapi.Field) *UserRolePermissionGroupsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolePermissionGroupsGetCall) Do() (*UserRolePermissionGroup, error) {
	var returnValue *UserRolePermissionGroup
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one user role permission group by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userRolePermissionGroups.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User role permission group ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRolePermissionGroups/{id}",
	//   "response": {
	//     "$ref": "UserRolePermissionGroup"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRolePermissionGroups.list":

type UserRolePermissionGroupsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Gets a list of all supported user role permission groups.

func (r *UserRolePermissionGroupsService) List(profileId int64) *UserRolePermissionGroupsListCall {
	return &UserRolePermissionGroupsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRolePermissionGroups",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolePermissionGroupsListCall) Fields(s ...googleapi.Field) *UserRolePermissionGroupsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolePermissionGroupsListCall) Do() (*UserRolePermissionGroupsListResponse, error) {
	var returnValue *UserRolePermissionGroupsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of all supported user role permission groups.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userRolePermissionGroups.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRolePermissionGroups",
	//   "response": {
	//     "$ref": "UserRolePermissionGroupsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRolePermissions.get":

type UserRolePermissionsGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one user role permission by ID.

func (r *UserRolePermissionsService) Get(profileId int64, id int64) *UserRolePermissionsGetCall {
	return &UserRolePermissionsGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRolePermissions/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolePermissionsGetCall) Fields(s ...googleapi.Field) *UserRolePermissionsGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolePermissionsGetCall) Do() (*UserRolePermission, error) {
	var returnValue *UserRolePermission
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one user role permission by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userRolePermissions.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User role permission ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRolePermissions/{id}",
	//   "response": {
	//     "$ref": "UserRolePermission"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRolePermissions.list":

type UserRolePermissionsListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Gets a list of user role permissions, possibly filtered.

func (r *UserRolePermissionsService) List(profileId int64) *UserRolePermissionsListCall {
	return &UserRolePermissionsListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRolePermissions",
	}
}

// Ids sets the optional parameter "ids": Select only user role
// permissions with these IDs.
func (c *UserRolePermissionsListCall) Ids(ids ...int64) *UserRolePermissionsListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolePermissionsListCall) Fields(s ...googleapi.Field) *UserRolePermissionsListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolePermissionsListCall) Do() (*UserRolePermissionsListResponse, error) {
	var returnValue *UserRolePermissionsListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets a list of user role permissions, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userRolePermissions.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "ids": {
	//       "description": "Select only user role permissions with these IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRolePermissions",
	//   "response": {
	//     "$ref": "UserRolePermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRoles.delete":

type UserRolesDeleteCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Delete: Deletes an existing user role.

func (r *UserRolesService) Delete(profileId int64, id int64) *UserRolesDeleteCall {
	return &UserRolesDeleteCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRoles/{id}",
	}
}

func (c *UserRolesDeleteCall) Do() error {
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "DELETE",
		URL:    u,
		Params: c.params_,
	}

	return c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Deletes an existing user role.",
	//   "httpMethod": "DELETE",
	//   "id": "dfareporting.userRoles.delete",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User role ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRoles/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRoles.get":

type UserRolesGetCall struct {
	s             *Service
	profileId     int64
	id            int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Get: Gets one user role by ID.

func (r *UserRolesService) Get(profileId int64, id int64) *UserRolesGetCall {
	return &UserRolesGetCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRoles/{id}",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolesGetCall) Fields(s ...googleapi.Field) *UserRolesGetCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolesGetCall) Do() (*UserRole, error) {
	var returnValue *UserRole
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
		"id":        strconv.FormatInt(c.id, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Gets one user role by ID.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userRoles.get",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User role ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRoles/{id}",
	//   "response": {
	//     "$ref": "UserRole"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRoles.insert":

type UserRolesInsertCall struct {
	s             *Service
	profileId     int64
	userrole      *UserRole
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Insert: Inserts a new user role.

func (r *UserRolesService) Insert(profileId int64, userrole *UserRole) *UserRolesInsertCall {
	return &UserRolesInsertCall{
		s:             r.s,
		profileId:     profileId,
		userrole:      userrole,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRoles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolesInsertCall) Fields(s ...googleapi.Field) *UserRolesInsertCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolesInsertCall) Do() (*UserRole, error) {
	var returnValue *UserRole
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "POST",
		URL:     u,
		Params:  c.params_,
		Payload: c.userrole,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Inserts a new user role.",
	//   "httpMethod": "POST",
	//   "id": "dfareporting.userRoles.insert",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRoles",
	//   "request": {
	//     "$ref": "UserRole"
	//   },
	//   "response": {
	//     "$ref": "UserRole"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRoles.list":

type UserRolesListCall struct {
	s             *Service
	profileId     int64
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// List: Retrieves a list of user roles, possibly filtered.

func (r *UserRolesService) List(profileId int64) *UserRolesListCall {
	return &UserRolesListCall{
		s:             r.s,
		profileId:     profileId,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRoles",
	}
}

// AccountUserRoleOnly sets the optional parameter
// "accountUserRoleOnly": Select only account level user roles not
// associated with any specific subaccount.
func (c *UserRolesListCall) AccountUserRoleOnly(accountUserRoleOnly bool) *UserRolesListCall {
	c.params_.Set("accountUserRoleOnly", fmt.Sprintf("%v", accountUserRoleOnly))
	return c
}

// Ids sets the optional parameter "ids": Select only user roles with
// the specified IDs.
func (c *UserRolesListCall) Ids(ids ...int64) *UserRolesListCall {
	for _, v_ := range ids {
		c.params_.Add("ids", fmt.Sprintf("%v", v_))
	}
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of results to return.
func (c *UserRolesListCall) MaxResults(maxResults int64) *UserRolesListCall {
	c.params_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Value of the
// nextPageToken from the previous result page.
func (c *UserRolesListCall) PageToken(pageToken string) *UserRolesListCall {
	c.params_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// SearchString sets the optional parameter "searchString": Allows
// searching for objects by name or ID. Wildcards (*) are allowed. For
// example, "userrole*2015" will return objects with names like
// "userrole June 2015", "userrole April 2015" or simply "userrole
// 2015". Most of the searches also add wildcards implicitly at the
// start and the end of the search string. For example, a search string
// of "userrole" will match objects with name "my userrole", "userrole
// 2015" or simply "userrole".
func (c *UserRolesListCall) SearchString(searchString string) *UserRolesListCall {
	c.params_.Set("searchString", fmt.Sprintf("%v", searchString))
	return c
}

// SortField sets the optional parameter "sortField": Field by which to
// sort the list.
func (c *UserRolesListCall) SortField(sortField string) *UserRolesListCall {
	c.params_.Set("sortField", fmt.Sprintf("%v", sortField))
	return c
}

// SortOrder sets the optional parameter "sortOrder": Order of sorted
// results, default is ASCENDING.
func (c *UserRolesListCall) SortOrder(sortOrder string) *UserRolesListCall {
	c.params_.Set("sortOrder", fmt.Sprintf("%v", sortOrder))
	return c
}

// SubaccountId sets the optional parameter "subaccountId": Select only
// user roles that belong to this subaccount.
func (c *UserRolesListCall) SubaccountId(subaccountId int64) *UserRolesListCall {
	c.params_.Set("subaccountId", fmt.Sprintf("%v", subaccountId))
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolesListCall) Fields(s ...googleapi.Field) *UserRolesListCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolesListCall) Do() (*UserRolesListResponse, error) {
	var returnValue *UserRolesListResponse
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method: "GET",
		URL:    u,
		Params: c.params_,
		Result: &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Retrieves a list of user roles, possibly filtered.",
	//   "httpMethod": "GET",
	//   "id": "dfareporting.userRoles.list",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "accountUserRoleOnly": {
	//       "description": "Select only account level user roles not associated with any specific subaccount.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "ids": {
	//       "description": "Select only user roles with the specified IDs.",
	//       "format": "int64",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Value of the nextPageToken from the previous result page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "searchString": {
	//       "description": "Allows searching for objects by name or ID. Wildcards (*) are allowed. For example, \"userrole*2015\" will return objects with names like \"userrole June 2015\", \"userrole April 2015\" or simply \"userrole 2015\". Most of the searches also add wildcards implicitly at the start and the end of the search string. For example, a search string of \"userrole\" will match objects with name \"my userrole\", \"userrole 2015\" or simply \"userrole\".",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortField": {
	//       "description": "Field by which to sort the list.",
	//       "enum": [
	//         "ID",
	//         "NAME"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sortOrder": {
	//       "description": "Order of sorted results, default is ASCENDING.",
	//       "enum": [
	//         "ASCENDING",
	//         "DESCENDING"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "subaccountId": {
	//       "description": "Select only user roles that belong to this subaccount.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRoles",
	//   "response": {
	//     "$ref": "UserRolesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRoles.patch":

type UserRolesPatchCall struct {
	s             *Service
	profileId     int64
	id            int64
	userrole      *UserRole
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Patch: Updates an existing user role. This method supports patch
// semantics.

func (r *UserRolesService) Patch(profileId int64, id int64, userrole *UserRole) *UserRolesPatchCall {
	return &UserRolesPatchCall{
		s:             r.s,
		profileId:     profileId,
		id:            id,
		userrole:      userrole,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRoles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolesPatchCall) Fields(s ...googleapi.Field) *UserRolesPatchCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolesPatchCall) Do() (*UserRole, error) {
	var returnValue *UserRole
	c.params_.Set("id", fmt.Sprintf("%v", c.id))
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PATCH",
		URL:     u,
		Params:  c.params_,
		Payload: c.userrole,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing user role. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "dfareporting.userRoles.patch",
	//   "parameterOrder": [
	//     "profileId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "User role ID.",
	//       "format": "int64",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRoles",
	//   "request": {
	//     "$ref": "UserRole"
	//   },
	//   "response": {
	//     "$ref": "UserRole"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}

// method id "dfareporting.userRoles.update":

type UserRolesUpdateCall struct {
	s             *Service
	profileId     int64
	userrole      *UserRole
	caller_       googleapi.Caller
	params_       url.Values
	pathTemplate_ string
}

// Update: Updates an existing user role.

func (r *UserRolesService) Update(profileId int64, userrole *UserRole) *UserRolesUpdateCall {
	return &UserRolesUpdateCall{
		s:             r.s,
		profileId:     profileId,
		userrole:      userrole,
		caller_:       googleapi.JSONCall{},
		params_:       make(map[string][]string),
		pathTemplate_: "userprofiles/{profileId}/userRoles",
	}
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserRolesUpdateCall) Fields(s ...googleapi.Field) *UserRolesUpdateCall {
	c.params_.Set("fields", googleapi.CombineFields(s))
	return c
}

func (c *UserRolesUpdateCall) Do() (*UserRole, error) {
	var returnValue *UserRole
	u := googleapi.Expand(baseURL, c.pathTemplate_, map[string]string{
		"profileId": strconv.FormatInt(c.profileId, 10),
	})
	call := &googleapi.Call{
		Method:  "PUT",
		URL:     u,
		Params:  c.params_,
		Payload: c.userrole,
		Result:  &returnValue,
	}

	return returnValue, c.caller_.Do(googleapi.NoContext, c.s.client, call)
	// {
	//   "description": "Updates an existing user role.",
	//   "httpMethod": "PUT",
	//   "id": "dfareporting.userRoles.update",
	//   "parameterOrder": [
	//     "profileId"
	//   ],
	//   "parameters": {
	//     "profileId": {
	//       "description": "User profile ID associated with this request.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "userprofiles/{profileId}/userRoles",
	//   "request": {
	//     "$ref": "UserRole"
	//   },
	//   "response": {
	//     "$ref": "UserRole"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/dfatrafficking"
	//   ]
	// }

}
