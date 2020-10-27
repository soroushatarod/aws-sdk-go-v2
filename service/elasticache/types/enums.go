// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthenticationType string

// Enum values for AuthenticationType
const (
	AuthenticationTypePassword    AuthenticationType = "password"
	AuthenticationTypeNo_password AuthenticationType = "no-password"
)

// Values returns all known values for AuthenticationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		"password",
		"no-password",
	}
}

type AuthTokenUpdateStatus string

// Enum values for AuthTokenUpdateStatus
const (
	AuthTokenUpdateStatusSetting  AuthTokenUpdateStatus = "SETTING"
	AuthTokenUpdateStatusRotating AuthTokenUpdateStatus = "ROTATING"
)

// Values returns all known values for AuthTokenUpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthTokenUpdateStatus) Values() []AuthTokenUpdateStatus {
	return []AuthTokenUpdateStatus{
		"SETTING",
		"ROTATING",
	}
}

type AuthTokenUpdateStrategyType string

// Enum values for AuthTokenUpdateStrategyType
const (
	AuthTokenUpdateStrategyTypeSet    AuthTokenUpdateStrategyType = "SET"
	AuthTokenUpdateStrategyTypeRotate AuthTokenUpdateStrategyType = "ROTATE"
	AuthTokenUpdateStrategyTypeDelete AuthTokenUpdateStrategyType = "DELETE"
)

// Values returns all known values for AuthTokenUpdateStrategyType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AuthTokenUpdateStrategyType) Values() []AuthTokenUpdateStrategyType {
	return []AuthTokenUpdateStrategyType{
		"SET",
		"ROTATE",
		"DELETE",
	}
}

type AutomaticFailoverStatus string

// Enum values for AutomaticFailoverStatus
const (
	AutomaticFailoverStatusEnabled   AutomaticFailoverStatus = "enabled"
	AutomaticFailoverStatusDisabled  AutomaticFailoverStatus = "disabled"
	AutomaticFailoverStatusEnabling  AutomaticFailoverStatus = "enabling"
	AutomaticFailoverStatusDisabling AutomaticFailoverStatus = "disabling"
)

// Values returns all known values for AutomaticFailoverStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutomaticFailoverStatus) Values() []AutomaticFailoverStatus {
	return []AutomaticFailoverStatus{
		"enabled",
		"disabled",
		"enabling",
		"disabling",
	}
}

type AZMode string

// Enum values for AZMode
const (
	AZModeSingle_az AZMode = "single-az"
	AZModeCross_az  AZMode = "cross-az"
)

// Values returns all known values for AZMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AZMode) Values() []AZMode {
	return []AZMode{
		"single-az",
		"cross-az",
	}
}

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeImmediate      ChangeType = "immediate"
	ChangeTypeRequiresReboot ChangeType = "requires-reboot"
)

// Values returns all known values for ChangeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{
		"immediate",
		"requires-reboot",
	}
}

type MultiAZStatus string

// Enum values for MultiAZStatus
const (
	MultiAZStatusEnabled  MultiAZStatus = "enabled"
	MultiAZStatusDisabled MultiAZStatus = "disabled"
)

// Values returns all known values for MultiAZStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MultiAZStatus) Values() []MultiAZStatus {
	return []MultiAZStatus{
		"enabled",
		"disabled",
	}
}

type NodeUpdateInitiatedBy string

// Enum values for NodeUpdateInitiatedBy
const (
	NodeUpdateInitiatedBySystem   NodeUpdateInitiatedBy = "system"
	NodeUpdateInitiatedByCustomer NodeUpdateInitiatedBy = "customer"
)

// Values returns all known values for NodeUpdateInitiatedBy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NodeUpdateInitiatedBy) Values() []NodeUpdateInitiatedBy {
	return []NodeUpdateInitiatedBy{
		"system",
		"customer",
	}
}

type NodeUpdateStatus string

// Enum values for NodeUpdateStatus
const (
	NodeUpdateStatusNot_applied      NodeUpdateStatus = "not-applied"
	NodeUpdateStatusWaiting_to_start NodeUpdateStatus = "waiting-to-start"
	NodeUpdateStatusIn_progress      NodeUpdateStatus = "in-progress"
	NodeUpdateStatusStopping         NodeUpdateStatus = "stopping"
	NodeUpdateStatusStopped          NodeUpdateStatus = "stopped"
	NodeUpdateStatusComplete         NodeUpdateStatus = "complete"
)

// Values returns all known values for NodeUpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NodeUpdateStatus) Values() []NodeUpdateStatus {
	return []NodeUpdateStatus{
		"not-applied",
		"waiting-to-start",
		"in-progress",
		"stopping",
		"stopped",
		"complete",
	}
}

type OutpostMode string

// Enum values for OutpostMode
const (
	OutpostModeSingle_outpost OutpostMode = "single-outpost"
	OutpostModeCross_outpost  OutpostMode = "cross-outpost"
)

// Values returns all known values for OutpostMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutpostMode) Values() []OutpostMode {
	return []OutpostMode{
		"single-outpost",
		"cross-outpost",
	}
}

type PendingAutomaticFailoverStatus string

// Enum values for PendingAutomaticFailoverStatus
const (
	PendingAutomaticFailoverStatusEnabled  PendingAutomaticFailoverStatus = "enabled"
	PendingAutomaticFailoverStatusDisabled PendingAutomaticFailoverStatus = "disabled"
)

// Values returns all known values for PendingAutomaticFailoverStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (PendingAutomaticFailoverStatus) Values() []PendingAutomaticFailoverStatus {
	return []PendingAutomaticFailoverStatus{
		"enabled",
		"disabled",
	}
}

type ServiceUpdateSeverity string

// Enum values for ServiceUpdateSeverity
const (
	ServiceUpdateSeverityCritical  ServiceUpdateSeverity = "critical"
	ServiceUpdateSeverityImportant ServiceUpdateSeverity = "important"
	ServiceUpdateSeverityMedium    ServiceUpdateSeverity = "medium"
	ServiceUpdateSeverityLow       ServiceUpdateSeverity = "low"
)

// Values returns all known values for ServiceUpdateSeverity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceUpdateSeverity) Values() []ServiceUpdateSeverity {
	return []ServiceUpdateSeverity{
		"critical",
		"important",
		"medium",
		"low",
	}
}

type ServiceUpdateStatus string

// Enum values for ServiceUpdateStatus
const (
	ServiceUpdateStatusAvailable ServiceUpdateStatus = "available"
	ServiceUpdateStatusCancelled ServiceUpdateStatus = "cancelled"
	ServiceUpdateStatusExpired   ServiceUpdateStatus = "expired"
)

// Values returns all known values for ServiceUpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceUpdateStatus) Values() []ServiceUpdateStatus {
	return []ServiceUpdateStatus{
		"available",
		"cancelled",
		"expired",
	}
}

type ServiceUpdateType string

// Enum values for ServiceUpdateType
const (
	ServiceUpdateTypeSecurity_update ServiceUpdateType = "security-update"
)

// Values returns all known values for ServiceUpdateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceUpdateType) Values() []ServiceUpdateType {
	return []ServiceUpdateType{
		"security-update",
	}
}

type SlaMet string

// Enum values for SlaMet
const (
	SlaMetYes SlaMet = "yes"
	SlaMetNo  SlaMet = "no"
	SlaMetNa  SlaMet = "n/a"
)

// Values returns all known values for SlaMet. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SlaMet) Values() []SlaMet {
	return []SlaMet{
		"yes",
		"no",
		"n/a",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeCacheCluster        SourceType = "cache-cluster"
	SourceTypeCacheParameterGroup SourceType = "cache-parameter-group"
	SourceTypeCacheSecurityGroup  SourceType = "cache-security-group"
	SourceTypeCacheSubnetGroup    SourceType = "cache-subnet-group"
	SourceTypeReplicationGroup    SourceType = "replication-group"
	SourceTypeUser                SourceType = "user"
	SourceTypeUserGroup           SourceType = "user-group"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"cache-cluster",
		"cache-parameter-group",
		"cache-security-group",
		"cache-subnet-group",
		"replication-group",
		"user",
		"user-group",
	}
}

type UpdateActionStatus string

// Enum values for UpdateActionStatus
const (
	UpdateActionStatusNot_applied      UpdateActionStatus = "not-applied"
	UpdateActionStatusWaiting_to_start UpdateActionStatus = "waiting-to-start"
	UpdateActionStatusIn_progress      UpdateActionStatus = "in-progress"
	UpdateActionStatusStopping         UpdateActionStatus = "stopping"
	UpdateActionStatusStopped          UpdateActionStatus = "stopped"
	UpdateActionStatusComplete         UpdateActionStatus = "complete"
	UpdateActionStatusScheduling       UpdateActionStatus = "scheduling"
	UpdateActionStatusScheduled        UpdateActionStatus = "scheduled"
	UpdateActionStatusNot_applicable   UpdateActionStatus = "not-applicable"
)

// Values returns all known values for UpdateActionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateActionStatus) Values() []UpdateActionStatus {
	return []UpdateActionStatus{
		"not-applied",
		"waiting-to-start",
		"in-progress",
		"stopping",
		"stopped",
		"complete",
		"scheduling",
		"scheduled",
		"not-applicable",
	}
}
