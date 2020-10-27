// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type PeriodTriggersElement string

// Enum values for PeriodTriggersElement
const (
	PeriodTriggersElementAds PeriodTriggersElement = "ADS"
)

// Values returns all known values for PeriodTriggersElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PeriodTriggersElement) Values() []PeriodTriggersElement {
	return []PeriodTriggersElement{
		"ADS",
	}
}

type AdMarkers string

// Enum values for AdMarkers
const (
	AdMarkersNone            AdMarkers = "NONE"
	AdMarkersScte35_enhanced AdMarkers = "SCTE35_ENHANCED"
	AdMarkersPassthrough     AdMarkers = "PASSTHROUGH"
)

// Values returns all known values for AdMarkers. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AdMarkers) Values() []AdMarkers {
	return []AdMarkers{
		"NONE",
		"SCTE35_ENHANCED",
		"PASSTHROUGH",
	}
}

type EncryptionMethod string

// Enum values for EncryptionMethod
const (
	EncryptionMethodAes_128    EncryptionMethod = "AES_128"
	EncryptionMethodSample_aes EncryptionMethod = "SAMPLE_AES"
)

// Values returns all known values for EncryptionMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionMethod) Values() []EncryptionMethod {
	return []EncryptionMethod{
		"AES_128",
		"SAMPLE_AES",
	}
}

type ManifestLayout string

// Enum values for ManifestLayout
const (
	ManifestLayoutFull    ManifestLayout = "FULL"
	ManifestLayoutCompact ManifestLayout = "COMPACT"
)

// Values returns all known values for ManifestLayout. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ManifestLayout) Values() []ManifestLayout {
	return []ManifestLayout{
		"FULL",
		"COMPACT",
	}
}

type Profile string

// Enum values for Profile
const (
	ProfileNone      Profile = "NONE"
	ProfileHbbtv_1_5 Profile = "HBBTV_1_5"
)

// Values returns all known values for Profile. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Profile) Values() []Profile {
	return []Profile{
		"NONE",
		"HBBTV_1_5",
	}
}

type SegmentTemplateFormat string

// Enum values for SegmentTemplateFormat
const (
	SegmentTemplateFormatNumber_with_timeline SegmentTemplateFormat = "NUMBER_WITH_TIMELINE"
	SegmentTemplateFormatTime_with_timeline   SegmentTemplateFormat = "TIME_WITH_TIMELINE"
	SegmentTemplateFormatNumber_with_duration SegmentTemplateFormat = "NUMBER_WITH_DURATION"
)

// Values returns all known values for SegmentTemplateFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SegmentTemplateFormat) Values() []SegmentTemplateFormat {
	return []SegmentTemplateFormat{
		"NUMBER_WITH_TIMELINE",
		"TIME_WITH_TIMELINE",
		"NUMBER_WITH_DURATION",
	}
}

type StreamOrder string

// Enum values for StreamOrder
const (
	StreamOrderOriginal                 StreamOrder = "ORIGINAL"
	StreamOrderVideo_bitrate_ascending  StreamOrder = "VIDEO_BITRATE_ASCENDING"
	StreamOrderVideo_bitrate_descending StreamOrder = "VIDEO_BITRATE_DESCENDING"
)

// Values returns all known values for StreamOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StreamOrder) Values() []StreamOrder {
	return []StreamOrder{
		"ORIGINAL",
		"VIDEO_BITRATE_ASCENDING",
		"VIDEO_BITRATE_DESCENDING",
	}
}
