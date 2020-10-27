// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeImmediate       ChangeType = "IMMEDIATE"
	ChangeTypeRequires_reboot ChangeType = "REQUIRES_REBOOT"
)

// Values returns all known values for ChangeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{
		"IMMEDIATE",
		"REQUIRES_REBOOT",
	}
}

type IsModifiable string

// Enum values for IsModifiable
const (
	IsModifiableTrue        IsModifiable = "TRUE"
	IsModifiableFalse       IsModifiable = "FALSE"
	IsModifiableConditional IsModifiable = "CONDITIONAL"
)

// Values returns all known values for IsModifiable. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IsModifiable) Values() []IsModifiable {
	return []IsModifiable{
		"TRUE",
		"FALSE",
		"CONDITIONAL",
	}
}

type ParameterType string

// Enum values for ParameterType
const (
	ParameterTypeDefault            ParameterType = "DEFAULT"
	ParameterTypeNode_type_specific ParameterType = "NODE_TYPE_SPECIFIC"
)

// Values returns all known values for ParameterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParameterType) Values() []ParameterType {
	return []ParameterType{
		"DEFAULT",
		"NODE_TYPE_SPECIFIC",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeCluster         SourceType = "CLUSTER"
	SourceTypeParameter_group SourceType = "PARAMETER_GROUP"
	SourceTypeSubnet_group    SourceType = "SUBNET_GROUP"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"CLUSTER",
		"PARAMETER_GROUP",
		"SUBNET_GROUP",
	}
}

type SSEStatus string

// Enum values for SSEStatus
const (
	SSEStatusEnabling  SSEStatus = "ENABLING"
	SSEStatusEnabled   SSEStatus = "ENABLED"
	SSEStatusDisabling SSEStatus = "DISABLING"
	SSEStatusDisabled  SSEStatus = "DISABLED"
)

// Values returns all known values for SSEStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SSEStatus) Values() []SSEStatus {
	return []SSEStatus{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
	}
}
