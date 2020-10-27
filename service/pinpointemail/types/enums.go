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
