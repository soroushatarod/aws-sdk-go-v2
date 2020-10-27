// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ComponentFormat string

// Enum values for ComponentFormat
const (
	ComponentFormatShell ComponentFormat = "SHELL"
)

// Values returns all known values for ComponentFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentFormat) Values() []ComponentFormat {
	return []ComponentFormat{
		"SHELL",
	}
}

type ComponentType string

// Enum values for ComponentType
const (
	ComponentTypeBuild ComponentType = "BUILD"
	ComponentTypeTest  ComponentType = "TEST"
)

// Values returns all known values for ComponentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComponentType) Values() []ComponentType {
	return []ComponentType{
		"BUILD",
		"TEST",
	}
}

type EbsVolumeType string

// Enum values for EbsVolumeType
const (
	EbsVolumeTypeStandard EbsVolumeType = "standard"
	EbsVolumeTypeIo1      EbsVolumeType = "io1"
	EbsVolumeTypeIo2      EbsVolumeType = "io2"
	EbsVolumeTypeGp2      EbsVolumeType = "gp2"
	EbsVolumeTypeSc1      EbsVolumeType = "sc1"
	EbsVolumeTypeSt1      EbsVolumeType = "st1"
)

// Values returns all known values for EbsVolumeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EbsVolumeType) Values() []EbsVolumeType {
	return []EbsVolumeType{
		"standard",
		"io1",
		"io2",
		"gp2",
		"sc1",
		"st1",
	}
}

type ImageStatus string

// Enum values for ImageStatus
const (
	ImageStatusPending      ImageStatus = "PENDING"
	ImageStatusCreating     ImageStatus = "CREATING"
	ImageStatusBuilding     ImageStatus = "BUILDING"
	ImageStatusTesting      ImageStatus = "TESTING"
	ImageStatusDistributing ImageStatus = "DISTRIBUTING"
	ImageStatusIntegrating  ImageStatus = "INTEGRATING"
	ImageStatusAvailable    ImageStatus = "AVAILABLE"
	ImageStatusCancelled    ImageStatus = "CANCELLED"
	ImageStatusFailed       ImageStatus = "FAILED"
	ImageStatusDeprecated   ImageStatus = "DEPRECATED"
	ImageStatusDeleted      ImageStatus = "DELETED"
)

// Values returns all known values for ImageStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImageStatus) Values() []ImageStatus {
	return []ImageStatus{
		"PENDING",
		"CREATING",
		"BUILDING",
		"TESTING",
		"DISTRIBUTING",
		"INTEGRATING",
		"AVAILABLE",
		"CANCELLED",
		"FAILED",
		"DEPRECATED",
		"DELETED",
	}
}

type Ownership string

// Enum values for Ownership
const (
	OwnershipSelf   Ownership = "Self"
	OwnershipShared Ownership = "Shared"
	OwnershipAmazon Ownership = "Amazon"
)

// Values returns all known values for Ownership. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Ownership) Values() []Ownership {
	return []Ownership{
		"Self",
		"Shared",
		"Amazon",
	}
}

type PipelineExecutionStartCondition string

// Enum values for PipelineExecutionStartCondition
const (
	PipelineExecutionStartConditionExpression_match_only                             PipelineExecutionStartCondition = "EXPRESSION_MATCH_ONLY"
	PipelineExecutionStartConditionExpression_match_and_dependency_updates_available PipelineExecutionStartCondition = "EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE"
)

// Values returns all known values for PipelineExecutionStartCondition. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (PipelineExecutionStartCondition) Values() []PipelineExecutionStartCondition {
	return []PipelineExecutionStartCondition{
		"EXPRESSION_MATCH_ONLY",
		"EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE",
	}
}

type PipelineStatus string

// Enum values for PipelineStatus
const (
	PipelineStatusDisabled PipelineStatus = "DISABLED"
	PipelineStatusEnabled  PipelineStatus = "ENABLED"
)

// Values returns all known values for PipelineStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PipelineStatus) Values() []PipelineStatus {
	return []PipelineStatus{
		"DISABLED",
		"ENABLED",
	}
}

type Platform string

// Enum values for Platform
const (
	PlatformWindows Platform = "Windows"
	PlatformLinux   Platform = "Linux"
)

// Values returns all known values for Platform. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Platform) Values() []Platform {
	return []Platform{
		"Windows",
		"Linux",
	}
}
