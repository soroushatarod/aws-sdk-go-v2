// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ResourceOwner string

// Enum values for ResourceOwner
const (
	ResourceOwnerSelf           ResourceOwner = "SELF"
	ResourceOwnerOther_accounts ResourceOwner = "OTHER-ACCOUNTS"
)

// Values returns all known values for ResourceOwner. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceOwner) Values() []ResourceOwner {
	return []ResourceOwner{
		"SELF",
		"OTHER-ACCOUNTS",
	}
}

type ResourceShareAssociationStatus string

// Enum values for ResourceShareAssociationStatus
const (
	ResourceShareAssociationStatusAssociating    ResourceShareAssociationStatus = "ASSOCIATING"
	ResourceShareAssociationStatusAssociated     ResourceShareAssociationStatus = "ASSOCIATED"
	ResourceShareAssociationStatusFailed         ResourceShareAssociationStatus = "FAILED"
	ResourceShareAssociationStatusDisassociating ResourceShareAssociationStatus = "DISASSOCIATING"
	ResourceShareAssociationStatusDisassociated  ResourceShareAssociationStatus = "DISASSOCIATED"
)

// Values returns all known values for ResourceShareAssociationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResourceShareAssociationStatus) Values() []ResourceShareAssociationStatus {
	return []ResourceShareAssociationStatus{
		"ASSOCIATING",
		"ASSOCIATED",
		"FAILED",
		"DISASSOCIATING",
		"DISASSOCIATED",
	}
}

type ResourceShareAssociationType string

// Enum values for ResourceShareAssociationType
const (
	ResourceShareAssociationTypePrincipal ResourceShareAssociationType = "PRINCIPAL"
	ResourceShareAssociationTypeResource  ResourceShareAssociationType = "RESOURCE"
)

// Values returns all known values for ResourceShareAssociationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceShareAssociationType) Values() []ResourceShareAssociationType {
	return []ResourceShareAssociationType{
		"PRINCIPAL",
		"RESOURCE",
	}
}

type ResourceShareFeatureSet string

// Enum values for ResourceShareFeatureSet
const (
	ResourceShareFeatureSetCreated_from_policy   ResourceShareFeatureSet = "CREATED_FROM_POLICY"
	ResourceShareFeatureSetPromoting_to_standard ResourceShareFeatureSet = "PROMOTING_TO_STANDARD"
	ResourceShareFeatureSetStandard              ResourceShareFeatureSet = "STANDARD"
)

// Values returns all known values for ResourceShareFeatureSet. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceShareFeatureSet) Values() []ResourceShareFeatureSet {
	return []ResourceShareFeatureSet{
		"CREATED_FROM_POLICY",
		"PROMOTING_TO_STANDARD",
		"STANDARD",
	}
}

type ResourceShareInvitationStatus string

// Enum values for ResourceShareInvitationStatus
const (
	ResourceShareInvitationStatusPending  ResourceShareInvitationStatus = "PENDING"
	ResourceShareInvitationStatusAccepted ResourceShareInvitationStatus = "ACCEPTED"
	ResourceShareInvitationStatusRejected ResourceShareInvitationStatus = "REJECTED"
	ResourceShareInvitationStatusExpired  ResourceShareInvitationStatus = "EXPIRED"
)

// Values returns all known values for ResourceShareInvitationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResourceShareInvitationStatus) Values() []ResourceShareInvitationStatus {
	return []ResourceShareInvitationStatus{
		"PENDING",
		"ACCEPTED",
		"REJECTED",
		"EXPIRED",
	}
}

type ResourceShareStatus string

// Enum values for ResourceShareStatus
const (
	ResourceShareStatusPending  ResourceShareStatus = "PENDING"
	ResourceShareStatusActive   ResourceShareStatus = "ACTIVE"
	ResourceShareStatusFailed   ResourceShareStatus = "FAILED"
	ResourceShareStatusDeleting ResourceShareStatus = "DELETING"
	ResourceShareStatusDeleted  ResourceShareStatus = "DELETED"
)

// Values returns all known values for ResourceShareStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceShareStatus) Values() []ResourceShareStatus {
	return []ResourceShareStatus{
		"PENDING",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusAvailable                   ResourceStatus = "AVAILABLE"
	ResourceStatusZonal_resource_inaccessible ResourceStatus = "ZONAL_RESOURCE_INACCESSIBLE"
	ResourceStatusLimit_exceeded              ResourceStatus = "LIMIT_EXCEEDED"
	ResourceStatusUnavailable                 ResourceStatus = "UNAVAILABLE"
	ResourceStatusPending                     ResourceStatus = "PENDING"
)

// Values returns all known values for ResourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStatus) Values() []ResourceStatus {
	return []ResourceStatus{
		"AVAILABLE",
		"ZONAL_RESOURCE_INACCESSIBLE",
		"LIMIT_EXCEEDED",
		"UNAVAILABLE",
		"PENDING",
	}
}
