// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessEndpointType string

// Enum values for AccessEndpointType
const (
	AccessEndpointTypeStreaming AccessEndpointType = "STREAMING"
)

// Values returns all known values for AccessEndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccessEndpointType) Values() []AccessEndpointType {
	return []AccessEndpointType{
		"STREAMING",
	}
}

type Action string

// Enum values for Action
const (
	ActionClipboard_copy_from_local_device Action = "CLIPBOARD_COPY_FROM_LOCAL_DEVICE"
	ActionClipboard_copy_to_local_device   Action = "CLIPBOARD_COPY_TO_LOCAL_DEVICE"
	ActionFile_upload                      Action = "FILE_UPLOAD"
	ActionFile_download                    Action = "FILE_DOWNLOAD"
	ActionPrinting_to_local_device         Action = "PRINTING_TO_LOCAL_DEVICE"
)

// Values returns all known values for Action. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Action) Values() []Action {
	return []Action{
		"CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
		"CLIPBOARD_COPY_TO_LOCAL_DEVICE",
		"FILE_UPLOAD",
		"FILE_DOWNLOAD",
		"PRINTING_TO_LOCAL_DEVICE",
	}
}

type AuthenticationType string

// Enum values for AuthenticationType
const (
	AuthenticationTypeApi      AuthenticationType = "API"
	AuthenticationTypeSaml     AuthenticationType = "SAML"
	AuthenticationTypeUserpool AuthenticationType = "USERPOOL"
)

// Values returns all known values for AuthenticationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		"API",
		"SAML",
		"USERPOOL",
	}
}

type FleetAttribute string

// Enum values for FleetAttribute
const (
	FleetAttributeVpc_configuration                    FleetAttribute = "VPC_CONFIGURATION"
	FleetAttributeVpc_configuration_security_group_ids FleetAttribute = "VPC_CONFIGURATION_SECURITY_GROUP_IDS"
	FleetAttributeDomain_join_info                     FleetAttribute = "DOMAIN_JOIN_INFO"
	FleetAttributeIam_role_arn                         FleetAttribute = "IAM_ROLE_ARN"
)

// Values returns all known values for FleetAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FleetAttribute) Values() []FleetAttribute {
	return []FleetAttribute{
		"VPC_CONFIGURATION",
		"VPC_CONFIGURATION_SECURITY_GROUP_IDS",
		"DOMAIN_JOIN_INFO",
		"IAM_ROLE_ARN",
	}
}

type FleetErrorCode string

// Enum values for FleetErrorCode
const (
	FleetErrorCodeIam_service_role_missing_eni_describe_action             FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_ENI_DESCRIBE_ACTION"
	FleetErrorCodeIam_service_role_missing_eni_create_action               FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_ENI_CREATE_ACTION"
	FleetErrorCodeIam_service_role_missing_eni_delete_action               FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_ENI_DELETE_ACTION"
	FleetErrorCodeNetwork_interface_limit_exceeded                         FleetErrorCode = "NETWORK_INTERFACE_LIMIT_EXCEEDED"
	FleetErrorCodeInternal_service_error                                   FleetErrorCode = "INTERNAL_SERVICE_ERROR"
	FleetErrorCodeIam_service_role_is_missing                              FleetErrorCode = "IAM_SERVICE_ROLE_IS_MISSING"
	FleetErrorCodeMachine_role_is_missing                                  FleetErrorCode = "MACHINE_ROLE_IS_MISSING"
	FleetErrorCodeSts_disabled_in_region                                   FleetErrorCode = "STS_DISABLED_IN_REGION"
	FleetErrorCodeSubnet_has_insufficient_ip_addresses                     FleetErrorCode = "SUBNET_HAS_INSUFFICIENT_IP_ADDRESSES"
	FleetErrorCodeIam_service_role_missing_describe_subnet_action          FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_DESCRIBE_SUBNET_ACTION"
	FleetErrorCodeSubnet_not_found                                         FleetErrorCode = "SUBNET_NOT_FOUND"
	FleetErrorCodeImage_not_found                                          FleetErrorCode = "IMAGE_NOT_FOUND"
	FleetErrorCodeInvalid_subnet_configuration                             FleetErrorCode = "INVALID_SUBNET_CONFIGURATION"
	FleetErrorCodeSecurity_groups_not_found                                FleetErrorCode = "SECURITY_GROUPS_NOT_FOUND"
	FleetErrorCodeIgw_not_attached                                         FleetErrorCode = "IGW_NOT_ATTACHED"
	FleetErrorCodeIam_service_role_missing_describe_security_groups_action FleetErrorCode = "IAM_SERVICE_ROLE_MISSING_DESCRIBE_SECURITY_GROUPS_ACTION"
	FleetErrorCodeDomain_join_error_file_not_found                         FleetErrorCode = "DOMAIN_JOIN_ERROR_FILE_NOT_FOUND"
	FleetErrorCodeDomain_join_error_access_denied                          FleetErrorCode = "DOMAIN_JOIN_ERROR_ACCESS_DENIED"
	FleetErrorCodeDomain_join_error_logon_failure                          FleetErrorCode = "DOMAIN_JOIN_ERROR_LOGON_FAILURE"
	FleetErrorCodeDomain_join_error_invalid_parameter                      FleetErrorCode = "DOMAIN_JOIN_ERROR_INVALID_PARAMETER"
	FleetErrorCodeDomain_join_error_more_data                              FleetErrorCode = "DOMAIN_JOIN_ERROR_MORE_DATA"
	FleetErrorCodeDomain_join_error_no_such_domain                         FleetErrorCode = "DOMAIN_JOIN_ERROR_NO_SUCH_DOMAIN"
	FleetErrorCodeDomain_join_error_not_supported                          FleetErrorCode = "DOMAIN_JOIN_ERROR_NOT_SUPPORTED"
	FleetErrorCodeDomain_join_nerr_invalid_workgroup_name                  FleetErrorCode = "DOMAIN_JOIN_NERR_INVALID_WORKGROUP_NAME"
	FleetErrorCodeDomain_join_nerr_workstation_not_started                 FleetErrorCode = "DOMAIN_JOIN_NERR_WORKSTATION_NOT_STARTED"
	FleetErrorCodeDomain_join_error_ds_machine_account_quota_exceeded      FleetErrorCode = "DOMAIN_JOIN_ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED"
	FleetErrorCodeDomain_join_nerr_password_expired                        FleetErrorCode = "DOMAIN_JOIN_NERR_PASSWORD_EXPIRED"
	FleetErrorCodeDomain_join_internal_service_error                       FleetErrorCode = "DOMAIN_JOIN_INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for FleetErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FleetErrorCode) Values() []FleetErrorCode {
	return []FleetErrorCode{
		"IAM_SERVICE_ROLE_MISSING_ENI_DESCRIBE_ACTION",
		"IAM_SERVICE_ROLE_MISSING_ENI_CREATE_ACTION",
		"IAM_SERVICE_ROLE_MISSING_ENI_DELETE_ACTION",
		"NETWORK_INTERFACE_LIMIT_EXCEEDED",
		"INTERNAL_SERVICE_ERROR",
		"IAM_SERVICE_ROLE_IS_MISSING",
		"MACHINE_ROLE_IS_MISSING",
		"STS_DISABLED_IN_REGION",
		"SUBNET_HAS_INSUFFICIENT_IP_ADDRESSES",
		"IAM_SERVICE_ROLE_MISSING_DESCRIBE_SUBNET_ACTION",
		"SUBNET_NOT_FOUND",
		"IMAGE_NOT_FOUND",
		"INVALID_SUBNET_CONFIGURATION",
		"SECURITY_GROUPS_NOT_FOUND",
		"IGW_NOT_ATTACHED",
		"IAM_SERVICE_ROLE_MISSING_DESCRIBE_SECURITY_GROUPS_ACTION",
		"DOMAIN_JOIN_ERROR_FILE_NOT_FOUND",
		"DOMAIN_JOIN_ERROR_ACCESS_DENIED",
		"DOMAIN_JOIN_ERROR_LOGON_FAILURE",
		"DOMAIN_JOIN_ERROR_INVALID_PARAMETER",
		"DOMAIN_JOIN_ERROR_MORE_DATA",
		"DOMAIN_JOIN_ERROR_NO_SUCH_DOMAIN",
		"DOMAIN_JOIN_ERROR_NOT_SUPPORTED",
		"DOMAIN_JOIN_NERR_INVALID_WORKGROUP_NAME",
		"DOMAIN_JOIN_NERR_WORKSTATION_NOT_STARTED",
		"DOMAIN_JOIN_ERROR_DS_MACHINE_ACCOUNT_QUOTA_EXCEEDED",
		"DOMAIN_JOIN_NERR_PASSWORD_EXPIRED",
		"DOMAIN_JOIN_INTERNAL_SERVICE_ERROR",
	}
}

type FleetState string

// Enum values for FleetState
const (
	FleetStateStarting FleetState = "STARTING"
	FleetStateRunning  FleetState = "RUNNING"
	FleetStateStopping FleetState = "STOPPING"
	FleetStateStopped  FleetState = "STOPPED"
)

// Values returns all known values for FleetState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FleetState) Values() []FleetState {
	return []FleetState{
		"STARTING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
	}
}

type FleetType string

// Enum values for FleetType
const (
	FleetTypeAlways_on FleetType = "ALWAYS_ON"
	FleetTypeOn_demand FleetType = "ON_DEMAND"
)

// Values returns all known values for FleetType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (FleetType) Values() []FleetType {
	return []FleetType{
		"ALWAYS_ON",
		"ON_DEMAND",
	}
}

type ImageBuilderState string

// Enum values for ImageBuilderState
const (
	ImageBuilderStatePending        ImageBuilderState = "PENDING"
	ImageBuilderStateUpdating_agent ImageBuilderState = "UPDATING_AGENT"
	ImageBuilderStateRunning        ImageBuilderState = "RUNNING"
	ImageBuilderStateStopping       ImageBuilderState = "STOPPING"
	ImageBuilderStateStopped        ImageBuilderState = "STOPPED"
	ImageBuilderStateRebooting      ImageBuilderState = "REBOOTING"
	ImageBuilderStateSnapshotting   ImageBuilderState = "SNAPSHOTTING"
	ImageBuilderStateDeleting       ImageBuilderState = "DELETING"
	ImageBuilderStateFailed         ImageBuilderState = "FAILED"
)

// Values returns all known values for ImageBuilderState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageBuilderState) Values() []ImageBuilderState {
	return []ImageBuilderState{
		"PENDING",
		"UPDATING_AGENT",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"REBOOTING",
		"SNAPSHOTTING",
		"DELETING",
		"FAILED",
	}
}

type ImageBuilderStateChangeReasonCode string

// Enum values for ImageBuilderStateChangeReasonCode
const (
	ImageBuilderStateChangeReasonCodeInternal_error    ImageBuilderStateChangeReasonCode = "INTERNAL_ERROR"
	ImageBuilderStateChangeReasonCodeImage_unavailable ImageBuilderStateChangeReasonCode = "IMAGE_UNAVAILABLE"
)

// Values returns all known values for ImageBuilderStateChangeReasonCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ImageBuilderStateChangeReasonCode) Values() []ImageBuilderStateChangeReasonCode {
	return []ImageBuilderStateChangeReasonCode{
		"INTERNAL_ERROR",
		"IMAGE_UNAVAILABLE",
	}
}

type ImageState string

// Enum values for ImageState
const (
	ImageStatePending   ImageState = "PENDING"
	ImageStateAvailable ImageState = "AVAILABLE"
	ImageStateFailed    ImageState = "FAILED"
	ImageStateCopying   ImageState = "COPYING"
	ImageStateDeleting  ImageState = "DELETING"
)

// Values returns all known values for ImageState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ImageState) Values() []ImageState {
	return []ImageState{
		"PENDING",
		"AVAILABLE",
		"FAILED",
		"COPYING",
		"DELETING",
	}
}

type ImageStateChangeReasonCode string

// Enum values for ImageStateChangeReasonCode
const (
	ImageStateChangeReasonCodeInternal_error              ImageStateChangeReasonCode = "INTERNAL_ERROR"
	ImageStateChangeReasonCodeImage_builder_not_available ImageStateChangeReasonCode = "IMAGE_BUILDER_NOT_AVAILABLE"
	ImageStateChangeReasonCodeImage_copy_failure          ImageStateChangeReasonCode = "IMAGE_COPY_FAILURE"
)

// Values returns all known values for ImageStateChangeReasonCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ImageStateChangeReasonCode) Values() []ImageStateChangeReasonCode {
	return []ImageStateChangeReasonCode{
		"INTERNAL_ERROR",
		"IMAGE_BUILDER_NOT_AVAILABLE",
		"IMAGE_COPY_FAILURE",
	}
}

type MessageAction string

// Enum values for MessageAction
const (
	MessageActionSuppress MessageAction = "SUPPRESS"
	MessageActionResend   MessageAction = "RESEND"
)

// Values returns all known values for MessageAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MessageAction) Values() []MessageAction {
	return []MessageAction{
		"SUPPRESS",
		"RESEND",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionEnabled  Permission = "ENABLED"
	PermissionDisabled Permission = "DISABLED"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"ENABLED",
		"DISABLED",
	}
}

type PlatformType string

// Enum values for PlatformType
const (
	PlatformTypeWindows             PlatformType = "WINDOWS"
	PlatformTypeWindows_server_2016 PlatformType = "WINDOWS_SERVER_2016"
	PlatformTypeWindows_server_2019 PlatformType = "WINDOWS_SERVER_2019"
)

// Values returns all known values for PlatformType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PlatformType) Values() []PlatformType {
	return []PlatformType{
		"WINDOWS",
		"WINDOWS_SERVER_2016",
		"WINDOWS_SERVER_2019",
	}
}

type SessionConnectionState string

// Enum values for SessionConnectionState
const (
	SessionConnectionStateConnected     SessionConnectionState = "CONNECTED"
	SessionConnectionStateNot_connected SessionConnectionState = "NOT_CONNECTED"
)

// Values returns all known values for SessionConnectionState. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SessionConnectionState) Values() []SessionConnectionState {
	return []SessionConnectionState{
		"CONNECTED",
		"NOT_CONNECTED",
	}
}

type SessionState string

// Enum values for SessionState
const (
	SessionStateActive  SessionState = "ACTIVE"
	SessionStatePending SessionState = "PENDING"
	SessionStateExpired SessionState = "EXPIRED"
)

// Values returns all known values for SessionState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SessionState) Values() []SessionState {
	return []SessionState{
		"ACTIVE",
		"PENDING",
		"EXPIRED",
	}
}

type StackAttribute string

// Enum values for StackAttribute
const (
	StackAttributeStorage_connectors             StackAttribute = "STORAGE_CONNECTORS"
	StackAttributeStorage_connector_homefolders  StackAttribute = "STORAGE_CONNECTOR_HOMEFOLDERS"
	StackAttributeStorage_connector_google_drive StackAttribute = "STORAGE_CONNECTOR_GOOGLE_DRIVE"
	StackAttributeStorage_connector_one_drive    StackAttribute = "STORAGE_CONNECTOR_ONE_DRIVE"
	StackAttributeRedirect_url                   StackAttribute = "REDIRECT_URL"
	StackAttributeFeedback_url                   StackAttribute = "FEEDBACK_URL"
	StackAttributeTheme_name                     StackAttribute = "THEME_NAME"
	StackAttributeUser_settings                  StackAttribute = "USER_SETTINGS"
	StackAttributeEmbed_host_domains             StackAttribute = "EMBED_HOST_DOMAINS"
	StackAttributeIam_role_arn                   StackAttribute = "IAM_ROLE_ARN"
	StackAttributeAccess_endpoints               StackAttribute = "ACCESS_ENDPOINTS"
)

// Values returns all known values for StackAttribute. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackAttribute) Values() []StackAttribute {
	return []StackAttribute{
		"STORAGE_CONNECTORS",
		"STORAGE_CONNECTOR_HOMEFOLDERS",
		"STORAGE_CONNECTOR_GOOGLE_DRIVE",
		"STORAGE_CONNECTOR_ONE_DRIVE",
		"REDIRECT_URL",
		"FEEDBACK_URL",
		"THEME_NAME",
		"USER_SETTINGS",
		"EMBED_HOST_DOMAINS",
		"IAM_ROLE_ARN",
		"ACCESS_ENDPOINTS",
	}
}

type StackErrorCode string

// Enum values for StackErrorCode
const (
	StackErrorCodeStorage_connector_error StackErrorCode = "STORAGE_CONNECTOR_ERROR"
	StackErrorCodeInternal_service_error  StackErrorCode = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for StackErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StackErrorCode) Values() []StackErrorCode {
	return []StackErrorCode{
		"STORAGE_CONNECTOR_ERROR",
		"INTERNAL_SERVICE_ERROR",
	}
}

type StorageConnectorType string

// Enum values for StorageConnectorType
const (
	StorageConnectorTypeHomefolders  StorageConnectorType = "HOMEFOLDERS"
	StorageConnectorTypeGoogle_drive StorageConnectorType = "GOOGLE_DRIVE"
	StorageConnectorTypeOne_drive    StorageConnectorType = "ONE_DRIVE"
)

// Values returns all known values for StorageConnectorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StorageConnectorType) Values() []StorageConnectorType {
	return []StorageConnectorType{
		"HOMEFOLDERS",
		"GOOGLE_DRIVE",
		"ONE_DRIVE",
	}
}

type StreamView string

// Enum values for StreamView
const (
	StreamViewApp     StreamView = "APP"
	StreamViewDesktop StreamView = "DESKTOP"
)

// Values returns all known values for StreamView. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StreamView) Values() []StreamView {
	return []StreamView{
		"APP",
		"DESKTOP",
	}
}

type UsageReportExecutionErrorCode string

// Enum values for UsageReportExecutionErrorCode
const (
	UsageReportExecutionErrorCodeResource_not_found     UsageReportExecutionErrorCode = "RESOURCE_NOT_FOUND"
	UsageReportExecutionErrorCodeAccess_denied          UsageReportExecutionErrorCode = "ACCESS_DENIED"
	UsageReportExecutionErrorCodeInternal_service_error UsageReportExecutionErrorCode = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for UsageReportExecutionErrorCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (UsageReportExecutionErrorCode) Values() []UsageReportExecutionErrorCode {
	return []UsageReportExecutionErrorCode{
		"RESOURCE_NOT_FOUND",
		"ACCESS_DENIED",
		"INTERNAL_SERVICE_ERROR",
	}
}

type UsageReportSchedule string

// Enum values for UsageReportSchedule
const (
	UsageReportScheduleDaily UsageReportSchedule = "DAILY"
)

// Values returns all known values for UsageReportSchedule. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageReportSchedule) Values() []UsageReportSchedule {
	return []UsageReportSchedule{
		"DAILY",
	}
}

type UserStackAssociationErrorCode string

// Enum values for UserStackAssociationErrorCode
const (
	UserStackAssociationErrorCodeStack_not_found     UserStackAssociationErrorCode = "STACK_NOT_FOUND"
	UserStackAssociationErrorCodeUser_name_not_found UserStackAssociationErrorCode = "USER_NAME_NOT_FOUND"
	UserStackAssociationErrorCodeDirectory_not_found UserStackAssociationErrorCode = "DIRECTORY_NOT_FOUND"
	UserStackAssociationErrorCodeInternal_error      UserStackAssociationErrorCode = "INTERNAL_ERROR"
)

// Values returns all known values for UserStackAssociationErrorCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (UserStackAssociationErrorCode) Values() []UserStackAssociationErrorCode {
	return []UserStackAssociationErrorCode{
		"STACK_NOT_FOUND",
		"USER_NAME_NOT_FOUND",
		"DIRECTORY_NOT_FOUND",
		"INTERNAL_ERROR",
	}
}

type VisibilityType string

// Enum values for VisibilityType
const (
	VisibilityTypePublic  VisibilityType = "PUBLIC"
	VisibilityTypePrivate VisibilityType = "PRIVATE"
	VisibilityTypeShared  VisibilityType = "SHARED"
)

// Values returns all known values for VisibilityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VisibilityType) Values() []VisibilityType {
	return []VisibilityType{
		"PUBLIC",
		"PRIVATE",
		"SHARED",
	}
}
