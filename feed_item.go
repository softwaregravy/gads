package gads

type FeedItemService struct {
	Auth
}

func NewFeedItemService(auth *Auth) *FeedItemService {
	return &FeedItemService{Auth: *auth}
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItem.Status
// ENABLED, REMOVED, UNKNOWN
type FeedItemStatus string

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemDevicePreference
// Represents a FeedItem device preference
type FeedItemDevicePreference struct {
	DevicePreference int64 `xml:"https://adwords.google.com/api/adwords/cm/v201802 devicePreference,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemScheduling
// Represents a collection of FeedItem schedules specifying all time intervals for which the feed item may serve.
// Any time range not covered by the specified FeedItemSchedules will prevent the feed item from serving during those times.
type FeedItemScheduling struct {
	feedItemSchedules int64 `xml:"https://adwords.google.com/api/adwords/cm/v201802 feedItemSchedules,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemSchedule
// Represents a FeedItem schedule, which specifies a time interval on a given day when the feed item may serve.
// The FeedItemSchedule times are in the account's time zone.
type FeedItemSchedule struct {
	DayOfWeek   DayOfWeek    `xml:"https://adwords.google.com/api/adwords/cm/v201802 dayOfWeek,omitempty"`
	StartHour   int          `xml:"https://adwords.google.com/api/adwords/cm/v201802 startHour,omitempty"`
	StartMinute MinuteOfHour `xml:"https://adwords.google.com/api/adwords/cm/v201802 startMinute,omitempty"`
	EndHour     int          `xml:"https://adwords.google.com/api/adwords/cm/v201802 endHour,omitempty"`
	EndMinute   MinuteOfHour `xml:"https://adwords.google.com/api/adwords/cm/v201802 endMinute,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemCampaignTargeting
// Specifies the campaign the request context must match in order for the feed item to be considered eligible for serving (aka the targeted campaign).
// E.g., if the below campaign targeting is set to campaignId = X, then the feed item can only serve under campaign X.
type FeedItemCampaignTargeting struct {
	TargetingCampaignId int64 `xml:"https://adwords.google.com/api/adwords/cm/v201802 TargetingCampaignId,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemAdGroupTargeting
// Specifies the adgroup the request context must match in order for the feed item to be considered eligible for serving (aka the targeted adgroup).
// E.g., if the below adgroup targeting is set to adgroup = X, then the feed item can only serve under adgroup X.
type FeedItemAdGroupTargeting struct {
	TargetingAdGroupId int64 `xml:"https://adwords.google.com/api/adwords/cm/v201802 TargetingAdGroupId,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemGeoRestriction
// Represents a FeedItem geo restriction.
type FeedItemGeoRestriction struct {
	GeoRestriction GeoRestriction `xml:"https://adwords.google.com/api/adwords/cm/v201802 geoRestriction,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemPolicyData
// Contains offline-validation and approval results for a given FeedItem and FeedMapping.
// Each validation data indicates any issues found on the feed item when used in the context of the feed mapping.
type FeedItemPolicyData struct {
	PolicyData
	PlaceholderType           int                               `xml:"https://adwords.google.com/api/adwords/cm/v201802 placeholderType,omitempty"`
	FeedMappingId             int64                             `xml:"https://adwords.google.com/api/adwords/cm/v201802 feedMappingId,omitempty"`
	ValidationStatus          FeedItemValidationStatus          `xml:"https://adwords.google.com/api/adwords/cm/v201802 validationStatus,omitempty"`
	ApprovalStatus            FeedItemApprovalStatus            `xml:"https://adwords.google.com/api/adwords/cm/v201802 approvalStatus,omitempty"`
	ValidationErrors          []FeedItemAttributeError          `xml:"https://adwords.google.com/api/adwords/cm/v201802 validationErrors,omitempty"`
	QualityApprovalStatus     FeedItemQualityApprovalStatus     `xml:"https://adwords.google.com/api/adwords/cm/v201802 qualityApprovalStatus,omitempty"`
	QualityDisapprovalReasons FeedItemQualityDisapprovalReasons `xml:"https://adwords.google.com/api/adwords/cm/v201802 qualityDisapprovalReasons,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemValidationStatus
// Validation status of a FeedItem
// UNCHECKED, ERROR, VALID
type FeedItemValidationStatus string

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemApprovalStatus
// Feed item approval status
// UNCHECKED, APPROVED, DISAPPROVED
type FeedItemApprovalStatus string

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemAttributeError
// Contains validation error details for a set of feed attributes
type FeedItemAttributeError struct {
	FeedAttributeIds    []int64 `xml:"https://adwords.google.com/api/adwords/cm/v201802 feedAttributeIds,omitempty"`
	ValidationErrorCode int     `xml:"https://adwords.google.com/api/adwords/cm/v201802 validationErrorCode,omitempty"`
	ErrorInformation    string  `xml:"https://adwords.google.com/api/adwords/cm/v201802 errorInformation,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemQualityApprovalStatus
// Feed item quality evaluation approval status
// UNCHECKED, APPROVED, DISAPPROVED
type FeedItemQualityApprovalStatus string

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.FeedItemQualityDisapprovalReasons
// Feed item quality evaluation disapproval reasons.
type FeedItemQualityDisapprovalReasons string

// https://developers.google.com/adwords/api/docs/reference/v201802/AdGroupExtensionSettingService.CallConversionType
// Conversion type for a call extension.
type CallConversionType struct {
	ConversionTypeId int64 `xml:"https://adwords.google.com/api/adwords/cm/v201802 conversionTypeId,omitempty"`
}
