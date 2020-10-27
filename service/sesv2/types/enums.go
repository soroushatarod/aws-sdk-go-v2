// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BehaviorOnMxFailure string

// Enum values for BehaviorOnMxFailure
const (
	BehaviorOnMxFailureUse_default_value BehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	BehaviorOnMxFailureReject_message    BehaviorOnMxFailure = "REJECT_MESSAGE"
)

// Values returns all known values for BehaviorOnMxFailure. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BehaviorOnMxFailure) Values() []BehaviorOnMxFailure {
	return []BehaviorOnMxFailure{
		"USE_DEFAULT_VALUE",
		"REJECT_MESSAGE",
	}
}

type BulkEmailStatus string

// Enum values for BulkEmailStatus
const (
	BulkEmailStatusSuccess                          BulkEmailStatus = "SUCCESS"
	BulkEmailStatusMessage_rejected                 BulkEmailStatus = "MESSAGE_REJECTED"
	BulkEmailStatusMail_from_domain_not_verified    BulkEmailStatus = "MAIL_FROM_DOMAIN_NOT_VERIFIED"
	BulkEmailStatusConfiguration_set_not_found      BulkEmailStatus = "CONFIGURATION_SET_NOT_FOUND"
	BulkEmailStatusTemplate_not_found               BulkEmailStatus = "TEMPLATE_NOT_FOUND"
	BulkEmailStatusAccount_suspended                BulkEmailStatus = "ACCOUNT_SUSPENDED"
	BulkEmailStatusAccount_throttled                BulkEmailStatus = "ACCOUNT_THROTTLED"
	BulkEmailStatusAccount_daily_quota_exceeded     BulkEmailStatus = "ACCOUNT_DAILY_QUOTA_EXCEEDED"
	BulkEmailStatusInvalid_sending_pool_name        BulkEmailStatus = "INVALID_SENDING_POOL_NAME"
	BulkEmailStatusAccount_sending_paused           BulkEmailStatus = "ACCOUNT_SENDING_PAUSED"
	BulkEmailStatusConfiguration_set_sending_paused BulkEmailStatus = "CONFIGURATION_SET_SENDING_PAUSED"
	BulkEmailStatusInvalid_parameter                BulkEmailStatus = "INVALID_PARAMETER"
	BulkEmailStatusTransient_failure                BulkEmailStatus = "TRANSIENT_FAILURE"
	BulkEmailStatusFailed                           BulkEmailStatus = "FAILED"
)

// Values returns all known values for BulkEmailStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BulkEmailStatus) Values() []BulkEmailStatus {
	return []BulkEmailStatus{
		"SUCCESS",
		"MESSAGE_REJECTED",
		"MAIL_FROM_DOMAIN_NOT_VERIFIED",
		"CONFIGURATION_SET_NOT_FOUND",
		"TEMPLATE_NOT_FOUND",
		"ACCOUNT_SUSPENDED",
		"ACCOUNT_THROTTLED",
		"ACCOUNT_DAILY_QUOTA_EXCEEDED",
		"INVALID_SENDING_POOL_NAME",
		"ACCOUNT_SENDING_PAUSED",
		"CONFIGURATION_SET_SENDING_PAUSED",
		"INVALID_PARAMETER",
		"TRANSIENT_FAILURE",
		"FAILED",
	}
}

type ContactLanguage string

// Enum values for ContactLanguage
const (
	ContactLanguageEn ContactLanguage = "EN"
	ContactLanguageJa ContactLanguage = "JA"
)

// Values returns all known values for ContactLanguage. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContactLanguage) Values() []ContactLanguage {
	return []ContactLanguage{
		"EN",
		"JA",
	}
}

type DataFormat string

// Enum values for DataFormat
const (
	DataFormatCsv  DataFormat = "CSV"
	DataFormatJson DataFormat = "JSON"
)

// Values returns all known values for DataFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DataFormat) Values() []DataFormat {
	return []DataFormat{
		"CSV",
		"JSON",
	}
}

type DeliverabilityDashboardAccountStatus string

// Enum values for DeliverabilityDashboardAccountStatus
const (
	DeliverabilityDashboardAccountStatusActive             DeliverabilityDashboardAccountStatus = "ACTIVE"
	DeliverabilityDashboardAccountStatusPending_expiration DeliverabilityDashboardAccountStatus = "PENDING_EXPIRATION"
	DeliverabilityDashboardAccountStatusDisabled           DeliverabilityDashboardAccountStatus = "DISABLED"
)

// Values returns all known values for DeliverabilityDashboardAccountStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DeliverabilityDashboardAccountStatus) Values() []DeliverabilityDashboardAccountStatus {
	return []DeliverabilityDashboardAccountStatus{
		"ACTIVE",
		"PENDING_EXPIRATION",
		"DISABLED",
	}
}

type DeliverabilityTestStatus string

// Enum values for DeliverabilityTestStatus
const (
	DeliverabilityTestStatusIn_progress DeliverabilityTestStatus = "IN_PROGRESS"
	DeliverabilityTestStatusCompleted   DeliverabilityTestStatus = "COMPLETED"
)

// Values returns all known values for DeliverabilityTestStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliverabilityTestStatus) Values() []DeliverabilityTestStatus {
	return []DeliverabilityTestStatus{
		"IN_PROGRESS",
		"COMPLETED",
	}
}

type DimensionValueSource string

// Enum values for DimensionValueSource
const (
	DimensionValueSourceMessage_tag  DimensionValueSource = "MESSAGE_TAG"
	DimensionValueSourceEmail_header DimensionValueSource = "EMAIL_HEADER"
	DimensionValueSourceLink_tag     DimensionValueSource = "LINK_TAG"
)

// Values returns all known values for DimensionValueSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DimensionValueSource) Values() []DimensionValueSource {
	return []DimensionValueSource{
		"MESSAGE_TAG",
		"EMAIL_HEADER",
		"LINK_TAG",
	}
}

type DkimSigningAttributesOrigin string

// Enum values for DkimSigningAttributesOrigin
const (
	DkimSigningAttributesOriginAws_ses  DkimSigningAttributesOrigin = "AWS_SES"
	DkimSigningAttributesOriginExternal DkimSigningAttributesOrigin = "EXTERNAL"
)

// Values returns all known values for DkimSigningAttributesOrigin. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DkimSigningAttributesOrigin) Values() []DkimSigningAttributesOrigin {
	return []DkimSigningAttributesOrigin{
		"AWS_SES",
		"EXTERNAL",
	}
}

type DkimStatus string

// Enum values for DkimStatus
const (
	DkimStatusPending           DkimStatus = "PENDING"
	DkimStatusSuccess           DkimStatus = "SUCCESS"
	DkimStatusFailed            DkimStatus = "FAILED"
	DkimStatusTemporary_failure DkimStatus = "TEMPORARY_FAILURE"
	DkimStatusNot_started       DkimStatus = "NOT_STARTED"
)

// Values returns all known values for DkimStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DkimStatus) Values() []DkimStatus {
	return []DkimStatus{
		"PENDING",
		"SUCCESS",
		"FAILED",
		"TEMPORARY_FAILURE",
		"NOT_STARTED",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeSend              EventType = "SEND"
	EventTypeReject            EventType = "REJECT"
	EventTypeBounce            EventType = "BOUNCE"
	EventTypeComplaint         EventType = "COMPLAINT"
	EventTypeDelivery          EventType = "DELIVERY"
	EventTypeOpen              EventType = "OPEN"
	EventTypeClick             EventType = "CLICK"
	EventTypeRendering_failure EventType = "RENDERING_FAILURE"
	EventTypeDelivery_delay    EventType = "DELIVERY_DELAY"
)

// Values returns all known values for EventType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"SEND",
		"REJECT",
		"BOUNCE",
		"COMPLAINT",
		"DELIVERY",
		"OPEN",
		"CLICK",
		"RENDERING_FAILURE",
		"DELIVERY_DELAY",
	}
}

type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeEmail_address  IdentityType = "EMAIL_ADDRESS"
	IdentityTypeDomain         IdentityType = "DOMAIN"
	IdentityTypeManaged_domain IdentityType = "MANAGED_DOMAIN"
)

// Values returns all known values for IdentityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IdentityType) Values() []IdentityType {
	return []IdentityType{
		"EMAIL_ADDRESS",
		"DOMAIN",
		"MANAGED_DOMAIN",
	}
}

type ImportDestinationType string

// Enum values for ImportDestinationType
const (
	ImportDestinationTypeSuppression_list ImportDestinationType = "SUPPRESSION_LIST"
)

// Values returns all known values for ImportDestinationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportDestinationType) Values() []ImportDestinationType {
	return []ImportDestinationType{
		"SUPPRESSION_LIST",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusCreated    JobStatus = "CREATED"
	JobStatusProcessing JobStatus = "PROCESSING"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"CREATED",
		"PROCESSING",
		"COMPLETED",
		"FAILED",
	}
}

type MailFromDomainStatus string

// Enum values for MailFromDomainStatus
const (
	MailFromDomainStatusPending           MailFromDomainStatus = "PENDING"
	MailFromDomainStatusSuccess           MailFromDomainStatus = "SUCCESS"
	MailFromDomainStatusFailed            MailFromDomainStatus = "FAILED"
	MailFromDomainStatusTemporary_failure MailFromDomainStatus = "TEMPORARY_FAILURE"
)

// Values returns all known values for MailFromDomainStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MailFromDomainStatus) Values() []MailFromDomainStatus {
	return []MailFromDomainStatus{
		"PENDING",
		"SUCCESS",
		"FAILED",
		"TEMPORARY_FAILURE",
	}
}

type MailType string

// Enum values for MailType
const (
	MailTypeMarketing     MailType = "MARKETING"
	MailTypeTransactional MailType = "TRANSACTIONAL"
)

// Values returns all known values for MailType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (MailType) Values() []MailType {
	return []MailType{
		"MARKETING",
		"TRANSACTIONAL",
	}
}

type ReviewStatus string

// Enum values for ReviewStatus
const (
	ReviewStatusPending ReviewStatus = "PENDING"
	ReviewStatusFailed  ReviewStatus = "FAILED"
	ReviewStatusGranted ReviewStatus = "GRANTED"
	ReviewStatusDenied  ReviewStatus = "DENIED"
)

// Values returns all known values for ReviewStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReviewStatus) Values() []ReviewStatus {
	return []ReviewStatus{
		"PENDING",
		"FAILED",
		"GRANTED",
		"DENIED",
	}
}

type SuppressionListImportAction string

// Enum values for SuppressionListImportAction
const (
	SuppressionListImportActionDelete SuppressionListImportAction = "DELETE"
	SuppressionListImportActionPut    SuppressionListImportAction = "PUT"
)

// Values returns all known values for SuppressionListImportAction. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionListImportAction) Values() []SuppressionListImportAction {
	return []SuppressionListImportAction{
		"DELETE",
		"PUT",
	}
}

type SuppressionListReason string

// Enum values for SuppressionListReason
const (
	SuppressionListReasonBounce    SuppressionListReason = "BOUNCE"
	SuppressionListReasonComplaint SuppressionListReason = "COMPLAINT"
)

// Values returns all known values for SuppressionListReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionListReason) Values() []SuppressionListReason {
	return []SuppressionListReason{
		"BOUNCE",
		"COMPLAINT",
	}
}

type TlsPolicy string

// Enum values for TlsPolicy
const (
	TlsPolicyRequire  TlsPolicy = "REQUIRE"
	TlsPolicyOptional TlsPolicy = "OPTIONAL"
)

// Values returns all known values for TlsPolicy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TlsPolicy) Values() []TlsPolicy {
	return []TlsPolicy{
		"REQUIRE",
		"OPTIONAL",
	}
}

type WarmupStatus string

// Enum values for WarmupStatus
const (
	WarmupStatusIn_progress WarmupStatus = "IN_PROGRESS"
	WarmupStatusDone        WarmupStatus = "DONE"
)

// Values returns all known values for WarmupStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (WarmupStatus) Values() []WarmupStatus {
	return []WarmupStatus{
		"IN_PROGRESS",
		"DONE",
	}
}
