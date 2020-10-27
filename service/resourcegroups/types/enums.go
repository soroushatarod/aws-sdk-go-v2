// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type GroupConfigurationStatus string

// Enum values for GroupConfigurationStatus
const (
	GroupConfigurationStatusUpdating        GroupConfigurationStatus = "UPDATING"
	GroupConfigurationStatusUpdate_complete GroupConfigurationStatus = "UPDATE_COMPLETE"
	GroupConfigurationStatusUpdate_failed   GroupConfigurationStatus = "UPDATE_FAILED"
)

// Values returns all known values for GroupConfigurationStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GroupConfigurationStatus) Values() []GroupConfigurationStatus {
	return []GroupConfigurationStatus{
		"UPDATING",
		"UPDATE_COMPLETE",
		"UPDATE_FAILED",
	}
}

type GroupFilterName string

// Enum values for GroupFilterName
const (
	GroupFilterNameResourcetype      GroupFilterName = "resource-type"
	GroupFilterNameConfigurationtype GroupFilterName = "configuration-type"
)

// Values returns all known values for GroupFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GroupFilterName) Values() []GroupFilterName {
	return []GroupFilterName{
		"resource-type",
		"configuration-type",
	}
}

type QueryErrorCode string

// Enum values for QueryErrorCode
const (
	QueryErrorCodeCloudformation_stack_inactive     QueryErrorCode = "CLOUDFORMATION_STACK_INACTIVE"
	QueryErrorCodeCloudformation_stack_not_existing QueryErrorCode = "CLOUDFORMATION_STACK_NOT_EXISTING"
)

// Values returns all known values for QueryErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QueryErrorCode) Values() []QueryErrorCode {
	return []QueryErrorCode{
		"CLOUDFORMATION_STACK_INACTIVE",
		"CLOUDFORMATION_STACK_NOT_EXISTING",
	}
}

type QueryType string

// Enum values for QueryType
const (
	QueryTypeTag_filters_1_0          QueryType = "TAG_FILTERS_1_0"
	QueryTypeCloudformation_stack_1_0 QueryType = "CLOUDFORMATION_STACK_1_0"
)

// Values returns all known values for QueryType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (QueryType) Values() []QueryType {
	return []QueryType{
		"TAG_FILTERS_1_0",
		"CLOUDFORMATION_STACK_1_0",
	}
}

type ResourceFilterName string

// Enum values for ResourceFilterName
const (
	ResourceFilterNameResourcetype ResourceFilterName = "resource-type"
)

// Values returns all known values for ResourceFilterName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceFilterName) Values() []ResourceFilterName {
	return []ResourceFilterName{
		"resource-type",
	}
}
