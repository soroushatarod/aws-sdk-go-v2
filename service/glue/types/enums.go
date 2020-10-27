// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CatalogEncryptionMode string

// Enum values for CatalogEncryptionMode
const (
	CatalogEncryptionModeDisabled CatalogEncryptionMode = "DISABLED"
	CatalogEncryptionModeSsekms   CatalogEncryptionMode = "SSE-KMS"
)

// Values returns all known values for CatalogEncryptionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CatalogEncryptionMode) Values() []CatalogEncryptionMode {
	return []CatalogEncryptionMode{
		"DISABLED",
		"SSE-KMS",
	}
}

type CloudWatchEncryptionMode string

// Enum values for CloudWatchEncryptionMode
const (
	CloudWatchEncryptionModeDisabled CloudWatchEncryptionMode = "DISABLED"
	CloudWatchEncryptionModeSsekms   CloudWatchEncryptionMode = "SSE-KMS"
)

// Values returns all known values for CloudWatchEncryptionMode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CloudWatchEncryptionMode) Values() []CloudWatchEncryptionMode {
	return []CloudWatchEncryptionMode{
		"DISABLED",
		"SSE-KMS",
	}
}

type ColumnStatisticsType string

// Enum values for ColumnStatisticsType
const (
	ColumnStatisticsTypeBoolean ColumnStatisticsType = "BOOLEAN"
	ColumnStatisticsTypeDate    ColumnStatisticsType = "DATE"
	ColumnStatisticsTypeDecimal ColumnStatisticsType = "DECIMAL"
	ColumnStatisticsTypeDouble  ColumnStatisticsType = "DOUBLE"
	ColumnStatisticsTypeLong    ColumnStatisticsType = "LONG"
	ColumnStatisticsTypeString  ColumnStatisticsType = "STRING"
	ColumnStatisticsTypeBinary  ColumnStatisticsType = "BINARY"
)

// Values returns all known values for ColumnStatisticsType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ColumnStatisticsType) Values() []ColumnStatisticsType {
	return []ColumnStatisticsType{
		"BOOLEAN",
		"DATE",
		"DECIMAL",
		"DOUBLE",
		"LONG",
		"STRING",
		"BINARY",
	}
}

type Comparator string

// Enum values for Comparator
const (
	ComparatorEquals              Comparator = "EQUALS"
	ComparatorGreater_than        Comparator = "GREATER_THAN"
	ComparatorLess_than           Comparator = "LESS_THAN"
	ComparatorGreater_than_equals Comparator = "GREATER_THAN_EQUALS"
	ComparatorLess_than_equals    Comparator = "LESS_THAN_EQUALS"
)

// Values returns all known values for Comparator. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Comparator) Values() []Comparator {
	return []Comparator{
		"EQUALS",
		"GREATER_THAN",
		"LESS_THAN",
		"GREATER_THAN_EQUALS",
		"LESS_THAN_EQUALS",
	}
}

type ConnectionPropertyKey string

// Enum values for ConnectionPropertyKey
const (
	ConnectionPropertyKeyHost                              ConnectionPropertyKey = "HOST"
	ConnectionPropertyKeyPort                              ConnectionPropertyKey = "PORT"
	ConnectionPropertyKeyUser_name                         ConnectionPropertyKey = "USERNAME"
	ConnectionPropertyKeyPassword                          ConnectionPropertyKey = "PASSWORD"
	ConnectionPropertyKeyEncrypted_password                ConnectionPropertyKey = "ENCRYPTED_PASSWORD"
	ConnectionPropertyKeyJdbc_driver_jar_uri               ConnectionPropertyKey = "JDBC_DRIVER_JAR_URI"
	ConnectionPropertyKeyJdbc_driver_class_name            ConnectionPropertyKey = "JDBC_DRIVER_CLASS_NAME"
	ConnectionPropertyKeyJdbc_engine                       ConnectionPropertyKey = "JDBC_ENGINE"
	ConnectionPropertyKeyJdbc_engine_version               ConnectionPropertyKey = "JDBC_ENGINE_VERSION"
	ConnectionPropertyKeyConfig_files                      ConnectionPropertyKey = "CONFIG_FILES"
	ConnectionPropertyKeyInstance_id                       ConnectionPropertyKey = "INSTANCE_ID"
	ConnectionPropertyKeyJdbc_connection_url               ConnectionPropertyKey = "JDBC_CONNECTION_URL"
	ConnectionPropertyKeyJdbc_enforce_ssl                  ConnectionPropertyKey = "JDBC_ENFORCE_SSL"
	ConnectionPropertyKeyCustom_jdbc_cert                  ConnectionPropertyKey = "CUSTOM_JDBC_CERT"
	ConnectionPropertyKeySkip_custom_jdbc_cert_validation  ConnectionPropertyKey = "SKIP_CUSTOM_JDBC_CERT_VALIDATION"
	ConnectionPropertyKeyCustom_jdbc_cert_string           ConnectionPropertyKey = "CUSTOM_JDBC_CERT_STRING"
	ConnectionPropertyKeyConnection_url                    ConnectionPropertyKey = "CONNECTION_URL"
	ConnectionPropertyKeyKafka_bootstrap_servers           ConnectionPropertyKey = "KAFKA_BOOTSTRAP_SERVERS"
	ConnectionPropertyKeyKafka_ssl_enabled                 ConnectionPropertyKey = "KAFKA_SSL_ENABLED"
	ConnectionPropertyKeyKafka_custom_cert                 ConnectionPropertyKey = "KAFKA_CUSTOM_CERT"
	ConnectionPropertyKeyKafka_skip_custom_cert_validation ConnectionPropertyKey = "KAFKA_SKIP_CUSTOM_CERT_VALIDATION"
)

// Values returns all known values for ConnectionPropertyKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionPropertyKey) Values() []ConnectionPropertyKey {
	return []ConnectionPropertyKey{
		"HOST",
		"PORT",
		"USERNAME",
		"PASSWORD",
		"ENCRYPTED_PASSWORD",
		"JDBC_DRIVER_JAR_URI",
		"JDBC_DRIVER_CLASS_NAME",
		"JDBC_ENGINE",
		"JDBC_ENGINE_VERSION",
		"CONFIG_FILES",
		"INSTANCE_ID",
		"JDBC_CONNECTION_URL",
		"JDBC_ENFORCE_SSL",
		"CUSTOM_JDBC_CERT",
		"SKIP_CUSTOM_JDBC_CERT_VALIDATION",
		"CUSTOM_JDBC_CERT_STRING",
		"CONNECTION_URL",
		"KAFKA_BOOTSTRAP_SERVERS",
		"KAFKA_SSL_ENABLED",
		"KAFKA_CUSTOM_CERT",
		"KAFKA_SKIP_CUSTOM_CERT_VALIDATION",
	}
}

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeJdbc    ConnectionType = "JDBC"
	ConnectionTypeSftp    ConnectionType = "SFTP"
	ConnectionTypeMongodb ConnectionType = "MONGODB"
	ConnectionTypeKafka   ConnectionType = "KAFKA"
	ConnectionTypeNetwork ConnectionType = "NETWORK"
)

// Values returns all known values for ConnectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionType) Values() []ConnectionType {
	return []ConnectionType{
		"JDBC",
		"SFTP",
		"MONGODB",
		"KAFKA",
		"NETWORK",
	}
}

type CrawlerState string

// Enum values for CrawlerState
const (
	CrawlerStateReady    CrawlerState = "READY"
	CrawlerStateRunning  CrawlerState = "RUNNING"
	CrawlerStateStopping CrawlerState = "STOPPING"
)

// Values returns all known values for CrawlerState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CrawlerState) Values() []CrawlerState {
	return []CrawlerState{
		"READY",
		"RUNNING",
		"STOPPING",
	}
}

type CrawlState string

// Enum values for CrawlState
const (
	CrawlStateRunning    CrawlState = "RUNNING"
	CrawlStateCancelling CrawlState = "CANCELLING"
	CrawlStateCancelled  CrawlState = "CANCELLED"
	CrawlStateSucceeded  CrawlState = "SUCCEEDED"
	CrawlStateFailed     CrawlState = "FAILED"
)

// Values returns all known values for CrawlState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CrawlState) Values() []CrawlState {
	return []CrawlState{
		"RUNNING",
		"CANCELLING",
		"CANCELLED",
		"SUCCEEDED",
		"FAILED",
	}
}

type CsvHeaderOption string

// Enum values for CsvHeaderOption
const (
	CsvHeaderOptionUnknown CsvHeaderOption = "UNKNOWN"
	CsvHeaderOptionPresent CsvHeaderOption = "PRESENT"
	CsvHeaderOptionAbsent  CsvHeaderOption = "ABSENT"
)

// Values returns all known values for CsvHeaderOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CsvHeaderOption) Values() []CsvHeaderOption {
	return []CsvHeaderOption{
		"UNKNOWN",
		"PRESENT",
		"ABSENT",
	}
}

type DeleteBehavior string

// Enum values for DeleteBehavior
const (
	DeleteBehaviorLog                   DeleteBehavior = "LOG"
	DeleteBehaviorDelete_from_database  DeleteBehavior = "DELETE_FROM_DATABASE"
	DeleteBehaviorDeprecate_in_database DeleteBehavior = "DEPRECATE_IN_DATABASE"
)

// Values returns all known values for DeleteBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeleteBehavior) Values() []DeleteBehavior {
	return []DeleteBehavior{
		"LOG",
		"DELETE_FROM_DATABASE",
		"DEPRECATE_IN_DATABASE",
	}
}

type EnableHybridValues string

// Enum values for EnableHybridValues
const (
	EnableHybridValuesTrue  EnableHybridValues = "TRUE"
	EnableHybridValuesFalse EnableHybridValues = "FALSE"
)

// Values returns all known values for EnableHybridValues. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnableHybridValues) Values() []EnableHybridValues {
	return []EnableHybridValues{
		"TRUE",
		"FALSE",
	}
}

type ExistCondition string

// Enum values for ExistCondition
const (
	ExistConditionMust_exist ExistCondition = "MUST_EXIST"
	ExistConditionNot_exist  ExistCondition = "NOT_EXIST"
	ExistConditionNone       ExistCondition = "NONE"
)

// Values returns all known values for ExistCondition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExistCondition) Values() []ExistCondition {
	return []ExistCondition{
		"MUST_EXIST",
		"NOT_EXIST",
		"NONE",
	}
}

type JobBookmarksEncryptionMode string

// Enum values for JobBookmarksEncryptionMode
const (
	JobBookmarksEncryptionModeDisabled JobBookmarksEncryptionMode = "DISABLED"
	JobBookmarksEncryptionModeCsekms   JobBookmarksEncryptionMode = "CSE-KMS"
)

// Values returns all known values for JobBookmarksEncryptionMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobBookmarksEncryptionMode) Values() []JobBookmarksEncryptionMode {
	return []JobBookmarksEncryptionMode{
		"DISABLED",
		"CSE-KMS",
	}
}

type JobRunState string

// Enum values for JobRunState
const (
	JobRunStateStarting  JobRunState = "STARTING"
	JobRunStateRunning   JobRunState = "RUNNING"
	JobRunStateStopping  JobRunState = "STOPPING"
	JobRunStateStopped   JobRunState = "STOPPED"
	JobRunStateSucceeded JobRunState = "SUCCEEDED"
	JobRunStateFailed    JobRunState = "FAILED"
	JobRunStateTimeout   JobRunState = "TIMEOUT"
)

// Values returns all known values for JobRunState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobRunState) Values() []JobRunState {
	return []JobRunState{
		"STARTING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"SUCCEEDED",
		"FAILED",
		"TIMEOUT",
	}
}

type Language string

// Enum values for Language
const (
	LanguagePython Language = "PYTHON"
	LanguageScala  Language = "SCALA"
)

// Values returns all known values for Language. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Language) Values() []Language {
	return []Language{
		"PYTHON",
		"SCALA",
	}
}

type LastCrawlStatus string

// Enum values for LastCrawlStatus
const (
	LastCrawlStatusSucceeded LastCrawlStatus = "SUCCEEDED"
	LastCrawlStatusCancelled LastCrawlStatus = "CANCELLED"
	LastCrawlStatusFailed    LastCrawlStatus = "FAILED"
)

// Values returns all known values for LastCrawlStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LastCrawlStatus) Values() []LastCrawlStatus {
	return []LastCrawlStatus{
		"SUCCEEDED",
		"CANCELLED",
		"FAILED",
	}
}

type Logical string

// Enum values for Logical
const (
	LogicalAnd Logical = "AND"
	LogicalAny Logical = "ANY"
)

// Values returns all known values for Logical. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Logical) Values() []Logical {
	return []Logical{
		"AND",
		"ANY",
	}
}

type LogicalOperator string

// Enum values for LogicalOperator
const (
	LogicalOperatorEquals LogicalOperator = "EQUALS"
)

// Values returns all known values for LogicalOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LogicalOperator) Values() []LogicalOperator {
	return []LogicalOperator{
		"EQUALS",
	}
}

type NodeType string

// Enum values for NodeType
const (
	NodeTypeCrawler NodeType = "CRAWLER"
	NodeTypeJob     NodeType = "JOB"
	NodeTypeTrigger NodeType = "TRIGGER"
)

// Values returns all known values for NodeType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (NodeType) Values() []NodeType {
	return []NodeType{
		"CRAWLER",
		"JOB",
		"TRIGGER",
	}
}

type PartitionIndexStatus string

// Enum values for PartitionIndexStatus
const (
	PartitionIndexStatusActive PartitionIndexStatus = "ACTIVE"
)

// Values returns all known values for PartitionIndexStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PartitionIndexStatus) Values() []PartitionIndexStatus {
	return []PartitionIndexStatus{
		"ACTIVE",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionAll                  Permission = "ALL"
	PermissionSelect               Permission = "SELECT"
	PermissionAlter                Permission = "ALTER"
	PermissionDrop                 Permission = "DROP"
	PermissionDelete               Permission = "DELETE"
	PermissionInsert               Permission = "INSERT"
	PermissionCreate_database      Permission = "CREATE_DATABASE"
	PermissionCreate_table         Permission = "CREATE_TABLE"
	PermissionData_location_access Permission = "DATA_LOCATION_ACCESS"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"ALL",
		"SELECT",
		"ALTER",
		"DROP",
		"DELETE",
		"INSERT",
		"CREATE_DATABASE",
		"CREATE_TABLE",
		"DATA_LOCATION_ACCESS",
	}
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeUser  PrincipalType = "USER"
	PrincipalTypeRole  PrincipalType = "ROLE"
	PrincipalTypeGroup PrincipalType = "GROUP"
)

// Values returns all known values for PrincipalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		"USER",
		"ROLE",
		"GROUP",
	}
}

type RecrawlBehavior string

// Enum values for RecrawlBehavior
const (
	RecrawlBehaviorCrawl_everything       RecrawlBehavior = "CRAWL_EVERYTHING"
	RecrawlBehaviorCrawl_new_folders_only RecrawlBehavior = "CRAWL_NEW_FOLDERS_ONLY"
)

// Values returns all known values for RecrawlBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecrawlBehavior) Values() []RecrawlBehavior {
	return []RecrawlBehavior{
		"CRAWL_EVERYTHING",
		"CRAWL_NEW_FOLDERS_ONLY",
	}
}

type ResourceShareType string

// Enum values for ResourceShareType
const (
	ResourceShareTypeForeign ResourceShareType = "FOREIGN"
	ResourceShareTypeAll     ResourceShareType = "ALL"
)

// Values returns all known values for ResourceShareType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceShareType) Values() []ResourceShareType {
	return []ResourceShareType{
		"FOREIGN",
		"ALL",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeJar     ResourceType = "JAR"
	ResourceTypeFile    ResourceType = "FILE"
	ResourceTypeArchive ResourceType = "ARCHIVE"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"JAR",
		"FILE",
		"ARCHIVE",
	}
}

type S3EncryptionMode string

// Enum values for S3EncryptionMode
const (
	S3EncryptionModeDisabled S3EncryptionMode = "DISABLED"
	S3EncryptionModeSsekms   S3EncryptionMode = "SSE-KMS"
	S3EncryptionModeSses3    S3EncryptionMode = "SSE-S3"
)

// Values returns all known values for S3EncryptionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (S3EncryptionMode) Values() []S3EncryptionMode {
	return []S3EncryptionMode{
		"DISABLED",
		"SSE-KMS",
		"SSE-S3",
	}
}

type ScheduleState string

// Enum values for ScheduleState
const (
	ScheduleStateScheduled     ScheduleState = "SCHEDULED"
	ScheduleStateNot_scheduled ScheduleState = "NOT_SCHEDULED"
	ScheduleStateTransitioning ScheduleState = "TRANSITIONING"
)

// Values returns all known values for ScheduleState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScheduleState) Values() []ScheduleState {
	return []ScheduleState{
		"SCHEDULED",
		"NOT_SCHEDULED",
		"TRANSITIONING",
	}
}

type Sort string

// Enum values for Sort
const (
	SortAscending  Sort = "ASC"
	SortDescending Sort = "DESC"
)

// Values returns all known values for Sort. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Sort) Values() []Sort {
	return []Sort{
		"ASC",
		"DESC",
	}
}

type SortDirectionType string

// Enum values for SortDirectionType
const (
	SortDirectionTypeDescending SortDirectionType = "DESCENDING"
	SortDirectionTypeAscending  SortDirectionType = "ASCENDING"
)

// Values returns all known values for SortDirectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SortDirectionType) Values() []SortDirectionType {
	return []SortDirectionType{
		"DESCENDING",
		"ASCENDING",
	}
}

type TaskRunSortColumnType string

// Enum values for TaskRunSortColumnType
const (
	TaskRunSortColumnTypeTask_run_type TaskRunSortColumnType = "TASK_RUN_TYPE"
	TaskRunSortColumnTypeStatus        TaskRunSortColumnType = "STATUS"
	TaskRunSortColumnTypeStarted       TaskRunSortColumnType = "STARTED"
)

// Values returns all known values for TaskRunSortColumnType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskRunSortColumnType) Values() []TaskRunSortColumnType {
	return []TaskRunSortColumnType{
		"TASK_RUN_TYPE",
		"STATUS",
		"STARTED",
	}
}

type TaskStatusType string

// Enum values for TaskStatusType
const (
	TaskStatusTypeStarting  TaskStatusType = "STARTING"
	TaskStatusTypeRunning   TaskStatusType = "RUNNING"
	TaskStatusTypeStopping  TaskStatusType = "STOPPING"
	TaskStatusTypeStopped   TaskStatusType = "STOPPED"
	TaskStatusTypeSucceeded TaskStatusType = "SUCCEEDED"
	TaskStatusTypeFailed    TaskStatusType = "FAILED"
	TaskStatusTypeTimeout   TaskStatusType = "TIMEOUT"
)

// Values returns all known values for TaskStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskStatusType) Values() []TaskStatusType {
	return []TaskStatusType{
		"STARTING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"SUCCEEDED",
		"FAILED",
		"TIMEOUT",
	}
}

type TaskType string

// Enum values for TaskType
const (
	TaskTypeEvaluation              TaskType = "EVALUATION"
	TaskTypeLabeling_set_generation TaskType = "LABELING_SET_GENERATION"
	TaskTypeImport_labels           TaskType = "IMPORT_LABELS"
	TaskTypeExport_labels           TaskType = "EXPORT_LABELS"
	TaskTypeFind_matches            TaskType = "FIND_MATCHES"
)

// Values returns all known values for TaskType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TaskType) Values() []TaskType {
	return []TaskType{
		"EVALUATION",
		"LABELING_SET_GENERATION",
		"IMPORT_LABELS",
		"EXPORT_LABELS",
		"FIND_MATCHES",
	}
}

type TransformSortColumnType string

// Enum values for TransformSortColumnType
const (
	TransformSortColumnTypeName           TransformSortColumnType = "NAME"
	TransformSortColumnTypeTransform_type TransformSortColumnType = "TRANSFORM_TYPE"
	TransformSortColumnTypeStatus         TransformSortColumnType = "STATUS"
	TransformSortColumnTypeCreated        TransformSortColumnType = "CREATED"
	TransformSortColumnTypeLast_modified  TransformSortColumnType = "LAST_MODIFIED"
)

// Values returns all known values for TransformSortColumnType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransformSortColumnType) Values() []TransformSortColumnType {
	return []TransformSortColumnType{
		"NAME",
		"TRANSFORM_TYPE",
		"STATUS",
		"CREATED",
		"LAST_MODIFIED",
	}
}

type TransformStatusType string

// Enum values for TransformStatusType
const (
	TransformStatusTypeNot_ready TransformStatusType = "NOT_READY"
	TransformStatusTypeReady     TransformStatusType = "READY"
	TransformStatusTypeDeleting  TransformStatusType = "DELETING"
)

// Values returns all known values for TransformStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransformStatusType) Values() []TransformStatusType {
	return []TransformStatusType{
		"NOT_READY",
		"READY",
		"DELETING",
	}
}

type TransformType string

// Enum values for TransformType
const (
	TransformTypeFind_matches TransformType = "FIND_MATCHES"
)

// Values returns all known values for TransformType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransformType) Values() []TransformType {
	return []TransformType{
		"FIND_MATCHES",
	}
}

type TriggerState string

// Enum values for TriggerState
const (
	TriggerStateCreating     TriggerState = "CREATING"
	TriggerStateCreated      TriggerState = "CREATED"
	TriggerStateActivating   TriggerState = "ACTIVATING"
	TriggerStateActivated    TriggerState = "ACTIVATED"
	TriggerStateDeactivating TriggerState = "DEACTIVATING"
	TriggerStateDeactivated  TriggerState = "DEACTIVATED"
	TriggerStateDeleting     TriggerState = "DELETING"
	TriggerStateUpdating     TriggerState = "UPDATING"
)

// Values returns all known values for TriggerState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TriggerState) Values() []TriggerState {
	return []TriggerState{
		"CREATING",
		"CREATED",
		"ACTIVATING",
		"ACTIVATED",
		"DEACTIVATING",
		"DEACTIVATED",
		"DELETING",
		"UPDATING",
	}
}

type TriggerType string

// Enum values for TriggerType
const (
	TriggerTypeScheduled   TriggerType = "SCHEDULED"
	TriggerTypeConditional TriggerType = "CONDITIONAL"
	TriggerTypeOn_demand   TriggerType = "ON_DEMAND"
)

// Values returns all known values for TriggerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TriggerType) Values() []TriggerType {
	return []TriggerType{
		"SCHEDULED",
		"CONDITIONAL",
		"ON_DEMAND",
	}
}

type UpdateBehavior string

// Enum values for UpdateBehavior
const (
	UpdateBehaviorLog                UpdateBehavior = "LOG"
	UpdateBehaviorUpdate_in_database UpdateBehavior = "UPDATE_IN_DATABASE"
)

// Values returns all known values for UpdateBehavior. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UpdateBehavior) Values() []UpdateBehavior {
	return []UpdateBehavior{
		"LOG",
		"UPDATE_IN_DATABASE",
	}
}

type WorkerType string

// Enum values for WorkerType
const (
	WorkerTypeStandard WorkerType = "Standard"
	WorkerTypeG1x      WorkerType = "G.1X"
	WorkerTypeG2x      WorkerType = "G.2X"
)

// Values returns all known values for WorkerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (WorkerType) Values() []WorkerType {
	return []WorkerType{
		"Standard",
		"G.1X",
		"G.2X",
	}
}

type WorkflowRunStatus string

// Enum values for WorkflowRunStatus
const (
	WorkflowRunStatusRunning   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatusCompleted WorkflowRunStatus = "COMPLETED"
	WorkflowRunStatusStopping  WorkflowRunStatus = "STOPPING"
	WorkflowRunStatusStopped   WorkflowRunStatus = "STOPPED"
	WorkflowRunStatusError     WorkflowRunStatus = "ERROR"
)

// Values returns all known values for WorkflowRunStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkflowRunStatus) Values() []WorkflowRunStatus {
	return []WorkflowRunStatus{
		"RUNNING",
		"COMPLETED",
		"STOPPING",
		"STOPPED",
		"ERROR",
	}
}
