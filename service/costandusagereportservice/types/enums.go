// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdditionalArtifact string

// Enum values for AdditionalArtifact
const (
	AdditionalArtifactRedshift   AdditionalArtifact = "REDSHIFT"
	AdditionalArtifactQuicksight AdditionalArtifact = "QUICKSIGHT"
	AdditionalArtifactAthena     AdditionalArtifact = "ATHENA"
)

// Values returns all known values for AdditionalArtifact. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdditionalArtifact) Values() []AdditionalArtifact {
	return []AdditionalArtifact{
		"REDSHIFT",
		"QUICKSIGHT",
		"ATHENA",
	}
}

type AWSRegion string

// Enum values for AWSRegion
const (
	AWSRegionCape_town           AWSRegion = "af-south-1"
	AWSRegionHong_kong           AWSRegion = "ap-east-1"
	AWSRegionMumbai              AWSRegion = "ap-south-1"
	AWSRegionSingapore           AWSRegion = "ap-southeast-1"
	AWSRegionSydney              AWSRegion = "ap-southeast-2"
	AWSRegionTokyo               AWSRegion = "ap-northeast-1"
	AWSRegionSeoul               AWSRegion = "ap-northeast-2"
	AWSRegionOsaka               AWSRegion = "ap-northeast-3"
	AWSRegionCanada_central      AWSRegion = "ca-central-1"
	AWSRegionFrankfurt           AWSRegion = "eu-central-1"
	AWSRegionIreland             AWSRegion = "eu-west-1"
	AWSRegionLondon              AWSRegion = "eu-west-2"
	AWSRegionParis               AWSRegion = "eu-west-3"
	AWSRegionStockholm           AWSRegion = "eu-north-1"
	AWSRegionMilano              AWSRegion = "eu-south-1"
	AWSRegionBahrain             AWSRegion = "me-south-1"
	AWSRegionSao_paulo           AWSRegion = "sa-east-1"
	AWSRegionUs_standard         AWSRegion = "us-east-1"
	AWSRegionOhio                AWSRegion = "us-east-2"
	AWSRegionNorthern_california AWSRegion = "us-west-1"
	AWSRegionOregon              AWSRegion = "us-west-2"
	AWSRegionBeijing             AWSRegion = "cn-north-1"
	AWSRegionNingxia             AWSRegion = "cn-northwest-1"
)

// Values returns all known values for AWSRegion. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AWSRegion) Values() []AWSRegion {
	return []AWSRegion{
		"af-south-1",
		"ap-east-1",
		"ap-south-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-north-1",
		"eu-south-1",
		"me-south-1",
		"sa-east-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"cn-north-1",
		"cn-northwest-1",
	}
}

type CompressionFormat string

// Enum values for CompressionFormat
const (
	CompressionFormatZip     CompressionFormat = "ZIP"
	CompressionFormatGzip    CompressionFormat = "GZIP"
	CompressionFormatParquet CompressionFormat = "Parquet"
)

// Values returns all known values for CompressionFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CompressionFormat) Values() []CompressionFormat {
	return []CompressionFormat{
		"ZIP",
		"GZIP",
		"Parquet",
	}
}

type ReportFormat string

// Enum values for ReportFormat
const (
	ReportFormatCsv     ReportFormat = "textORcsv"
	ReportFormatParquet ReportFormat = "Parquet"
)

// Values returns all known values for ReportFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportFormat) Values() []ReportFormat {
	return []ReportFormat{
		"textORcsv",
		"Parquet",
	}
}

type ReportVersioning string

// Enum values for ReportVersioning
const (
	ReportVersioningCreate_new_report ReportVersioning = "CREATE_NEW_REPORT"
	ReportVersioningOverwrite_report  ReportVersioning = "OVERWRITE_REPORT"
)

// Values returns all known values for ReportVersioning. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportVersioning) Values() []ReportVersioning {
	return []ReportVersioning{
		"CREATE_NEW_REPORT",
		"OVERWRITE_REPORT",
	}
}

type SchemaElement string

// Enum values for SchemaElement
const (
	SchemaElementResources SchemaElement = "RESOURCES"
)

// Values returns all known values for SchemaElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SchemaElement) Values() []SchemaElement {
	return []SchemaElement{
		"RESOURCES",
	}
}

type TimeUnit string

// Enum values for TimeUnit
const (
	TimeUnitHourly  TimeUnit = "HOURLY"
	TimeUnitDaily   TimeUnit = "DAILY"
	TimeUnitMonthly TimeUnit = "MONTHLY"
)

// Values returns all known values for TimeUnit. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TimeUnit) Values() []TimeUnit {
	return []TimeUnit{
		"HOURLY",
		"DAILY",
		"MONTHLY",
	}
}
