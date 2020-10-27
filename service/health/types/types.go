// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about an entity that is affected by a Health event.
type AffectedEntity struct {

	// The 12-digit AWS account number that contains the affected entity.
	AwsAccountId *string

	// The unique identifier for the entity. Format:
	// arn:aws:health:entity-region:aws-account:entity/entity-id . Example:
	// arn:aws:health:us-east-1:111222333444:entity/AVh5GGT7ul1arKr1sE1K
	EntityArn *string

	// The URL of the affected entity.
	EntityUrl *string

	// The ID of the affected entity.
	EntityValue *string

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string

	// The most recent time that the entity was updated.
	LastUpdatedTime *time.Time

	// The most recent status of the entity affected by the event. The possible values
	// are IMPAIRED, UNIMPAIRED, and UNKNOWN.
	StatusCode EntityStatusCode

	// A map of entity tags attached to the affected entity. Currently, the tags
	// property isn't supported.
	Tags map[string]*string
}

// A range of dates and times that is used by the EventFilter
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
// and EntityFilter
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
// objects. If from is set and to is set: match items where the timestamp
// (startTime, endTime, or lastUpdatedTime) is between from and to inclusive. If
// from is set and to is not set: match items where the timestamp value is equal to
// or after from. If from is not set and to is set: match items where the timestamp
// value is equal to or before to.
type DateTimeRange struct {

	// The starting date and time of a time range.
	From *time.Time

	// The ending date and time of a time range.
	To *time.Time
}

// The number of entities that are affected by one or more events. Returned by the
// DescribeEntityAggregates
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEntityAggregates.html)
// operation.
type EntityAggregate struct {

	// The number of entities that match the criteria for the specified events.
	Count *int32

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
}

// The values to use to filter results from the EntityFilter
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
// operation.
type EntityFilter struct {

	// A list of event ARNs (unique identifiers). For example:
	// "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	//
	// This member is required.
	EventArns []*string

	// A list of entity ARNs (unique identifiers).
	EntityArns []*string

	// A list of IDs for affected entities.
	EntityValues []*string

	// A list of the most recent dates and times that the entity was updated.
	LastUpdatedTimes []*DateTimeRange

	// A list of entity status codes (IMPAIRED, UNIMPAIRED, or UNKNOWN).
	StatusCodes []EntityStatusCode

	// A map of entity tags attached to the affected entity. Currently, the tags
	// property isn't supported.
	Tags []map[string]*string
}

// Summary information about an AWS Health event. AWS Health events can be public
// or account-specific:
//
//     * Public events might be service events that are not
// specific to an AWS account. For example, if there is an issue with an AWS
// Region, AWS Health provides information about the event, even if you don't use
// services or resources in that Region.
//
//     * Account-specific events are
// specific to either your AWS account or an account in your organization. For
// example, if there's an issue with Amazon Elastic Compute Cloud in a Region that
// you use, AWS Health provides information about the event and the affected
// resources in the account.
//
// You can determine if an event is public or
// account-specific by using the eventScopeCode parameter. For more information,
// see eventScopeCode
// (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html#AWSHealth-Type-Event-eventScopeCode).
type Event struct {

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	Arn *string

	// The AWS Availability Zone of the event. For example, us-east-1a.
	AvailabilityZone *string

	// The date and time that the event ended.
	EndTime *time.Time

	// This parameter specifies if the AWS Health event is a public AWS service event
	// or an account-specific event.
	//
	//     * If the eventScopeCode value is PUBLIC, then
	// the affectedAccounts value is always empty.
	//
	//     * If the eventScopeCode value
	// is ACCOUNT_SPECIFIC, then the affectedAccounts value lists the affected AWS
	// accounts in your organization. For example, if an event affects a service such
	// as Amazon Elastic Compute Cloud and you have AWS accounts that use that service,
	// those account IDs appear in the response.
	//
	//     * If the eventScopeCode value is
	// NONE, then the eventArn that you specified in the request is invalid or doesn't
	// exist.
	EventScopeCode EventScopeCode

	// The category of the event. Possible values are issue, scheduledChange, and
	// accountNotification.
	EventTypeCategory EventTypeCategory

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	EventTypeCode *string

	// The most recent date and time that the event was updated.
	LastUpdatedTime *time.Time

	// The AWS region name of the event.
	Region *string

	// The AWS service that is affected by the event. For example, EC2, RDS.
	Service *string

	// The date and time that the event began.
	StartTime *time.Time

	// The most recent status of the event. Possible values are open, closed, and
	// upcoming.
	StatusCode EventStatusCode
}

// The values used to filter results from the DescribeEventDetailsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
// and DescribeAffectedEntitiesForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
// operations.
type EventAccountFilter struct {

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	//
	// This member is required.
	EventArn *string

	// The 12-digit AWS account numbers that contains the affected entities.
	AwsAccountId *string
}

// The number of events of each issue type. Returned by the DescribeEventAggregates
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventAggregates.html)
// operation.
type EventAggregate struct {

	// The issue type for the associated count.
	AggregateValue *string

	// The number of events of the associated issue type.
	Count *int32
}

// The detailed description of the event. Included in the information returned by
// the DescribeEventDetails
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
// operation.
type EventDescription struct {

	// The most recent description of the event.
	LatestDescription *string
}

// Detailed information about an event. A combination of an Event
// (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html) object,
// an EventDescription
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EventDescription.html)
// object, and additional metadata about the event. Returned by the
// DescribeEventDetails
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
// operation.
type EventDetails struct {

	// Summary information about the event.
	Event *Event

	// The most recent description of the event.
	EventDescription *EventDescription

	// Additional metadata about the event.
	EventMetadata map[string]*string
}

// Error information returned when a DescribeEventDetails
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
// operation cannot find a specified event.
type EventDetailsErrorItem struct {

	// A message that describes the error.
	ErrorMessage *string

	// The name of the error.
	ErrorName *string

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
}

// The values to use to filter results from the DescribeEvents
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEvents.html)
// and DescribeEventAggregates
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventAggregates.html)
// operations.
type EventFilter struct {

	// A list of AWS availability zones.
	AvailabilityZones []*string

	// A list of dates and times that the event ended.
	EndTimes []*DateTimeRange

	// A list of entity ARNs (unique identifiers).
	EntityArns []*string

	// A list of entity identifiers, such as EC2 instance IDs (i-34ab692e) or EBS
	// volumes (vol-426ab23e).
	EntityValues []*string

	// A list of event ARNs (unique identifiers). For example:
	// "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	EventArns []*string

	// A list of event status codes.
	EventStatusCodes []EventStatusCode

	// A list of event type category codes (issue, scheduledChange, or
	// accountNotification).
	EventTypeCategories []EventTypeCategory

	// A list of unique identifiers for event types. For example,
	// "AWS_EC2_SYSTEM_MAINTENANCE_EVENT","AWS_RDS_MAINTENANCE_SCHEDULED".
	EventTypeCodes []*string

	// A list of dates and times that the event was last updated.
	LastUpdatedTimes []*DateTimeRange

	// A list of AWS regions.
	Regions []*string

	// The AWS services associated with the event. For example, EC2, RDS.
	Services []*string

	// A list of dates and times that the event began.
	StartTimes []*DateTimeRange

	// A map of entity tags attached to the affected entity. Currently, the tags
	// property isn't supported.
	Tags []map[string]*string
}

// Metadata about a type of event that is reported by AWS Health. Data consists of
// the category (for example, issue), the service (for example, EC2), and the event
// type code (for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT).
type EventType struct {

	// A list of event type category codes (issue, scheduledChange, or
	// accountNotification).
	Category EventTypeCategory

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	Code *string

	// The AWS service that is affected by the event. For example, EC2, RDS.
	Service *string
}

// The values to use to filter results from the DescribeEventTypes
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventTypes.html)
// operation.
type EventTypeFilter struct {

	// A list of event type category codes (issue, scheduledChange, or
	// accountNotification).
	EventTypeCategories []EventTypeCategory

	// A list of event type codes.
	EventTypeCodes []*string

	// The AWS services associated with the event. For example, EC2, RDS.
	Services []*string
}

// Error information returned when a DescribeAffectedEntitiesForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
// operation cannot find or process a specific entity.
type OrganizationAffectedEntitiesErrorItem struct {

	// The 12-digit AWS account numbers that contains the affected entities.
	AwsAccountId *string

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION.
	// For example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	ErrorMessage *string

	// The name of the error.
	ErrorName *string

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
}

// Summary information about an event, returned by the
// DescribeEventsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventsForOrganization.html)
// operation.
type OrganizationEvent struct {

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	Arn *string

	// The date and time that the event ended.
	EndTime *time.Time

	// This parameter specifies if the AWS Health event is a public AWS service event
	// or an account-specific event.
	//
	//     * If the eventScopeCode value is PUBLIC, then
	// the affectedAccounts value is always empty.
	//
	//     * If the eventScopeCode value
	// is ACCOUNT_SPECIFIC, then the affectedAccounts value lists the affected AWS
	// accounts in your organization. For example, if an event affects a service such
	// as Amazon Elastic Compute Cloud and you have AWS accounts that use that service,
	// those account IDs appear in the response.
	//
	//     * If the eventScopeCode value is
	// NONE, then the eventArn that you specified in the request is invalid or doesn't
	// exist.
	EventScopeCode EventScopeCode

	// The category of the event type.
	EventTypeCategory EventTypeCategory

	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION.
	// For example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT.
	EventTypeCode *string

	// The most recent date and time that the event was updated.
	LastUpdatedTime *time.Time

	// The AWS Region name of the event.
	Region *string

	// The AWS service that is affected by the event. For example, EC2, RDS.
	Service *string

	// The date and time that the event began.
	StartTime *time.Time

	// The most recent status of the event. Possible values are open, closed, and
	// upcoming.
	StatusCode EventStatusCode
}

// Detailed information about an event. A combination of an Event
// (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html) object,
// an EventDescription
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EventDescription.html)
// object, and additional metadata about the event. Returned by the
// DescribeEventDetailsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
// operation.
type OrganizationEventDetails struct {

	// The 12-digit AWS account numbers that contains the affected entities.
	AwsAccountId *string

	// Summary information about an AWS Health event. AWS Health events can be public
	// or account-specific:
	//
	//     * Public events might be service events that are not
	// specific to an AWS account. For example, if there is an issue with an AWS
	// Region, AWS Health provides information about the event, even if you don't use
	// services or resources in that Region.
	//
	//     * Account-specific events are
	// specific to either your AWS account or an account in your organization. For
	// example, if there's an issue with Amazon Elastic Compute Cloud in a Region that
	// you use, AWS Health provides information about the event and the affected
	// resources in the account.
	//
	// You can determine if an event is public or
	// account-specific by using the eventScopeCode parameter. For more information,
	// see eventScopeCode
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html#AWSHealth-Type-Event-eventScopeCode).
	Event *Event

	// The detailed description of the event. Included in the information returned by
	// the DescribeEventDetails
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
	// operation.
	EventDescription *EventDescription

	// Additional metadata about the event.
	EventMetadata map[string]*string
}

// Error information returned when a DescribeEventDetailsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
// operation cannot find a specified event.
type OrganizationEventDetailsErrorItem struct {

	// Error information returned when a DescribeEventDetailsForOrganization
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
	// operation cannot find a specified event.
	AwsAccountId *string

	// A message that describes the error.
	ErrorMessage *string

	// The name of the error.
	ErrorName *string

	// The unique identifier for the event. Format:
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID .
	// Example: Example:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
}

// The values to filter results from the DescribeEventsForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventsForOrganization.html)
// operation.
type OrganizationEventFilter struct {

	// A list of 12-digit AWS account numbers that contains the affected entities.
	AwsAccountIds []*string

	// A range of dates and times that is used by the EventFilter
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
	// and EntityFilter
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
	// objects. If from is set and to is set: match items where the timestamp
	// (startTime, endTime, or lastUpdatedTime) is between from and to inclusive. If
	// from is set and to is not set: match items where the timestamp value is equal to
	// or after from. If from is not set and to is set: match items where the timestamp
	// value is equal to or before to.
	EndTime *DateTimeRange

	// A list of entity ARNs (unique identifiers).
	EntityArns []*string

	// A list of entity identifiers, such as EC2 instance IDs (i-34ab692e) or EBS
	// volumes (vol-426ab23e).
	EntityValues []*string

	// A list of event status codes.
	EventStatusCodes []EventStatusCode

	// A list of event type category codes (issue, scheduledChange, or
	// accountNotification).
	EventTypeCategories []EventTypeCategory

	// A list of unique identifiers for event types. For example,
	// "AWS_EC2_SYSTEM_MAINTENANCE_EVENT","AWS_RDS_MAINTENANCE_SCHEDULED".
	EventTypeCodes []*string

	// A range of dates and times that is used by the EventFilter
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
	// and EntityFilter
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
	// objects. If from is set and to is set: match items where the timestamp
	// (startTime, endTime, or lastUpdatedTime) is between from and to inclusive. If
	// from is set and to is not set: match items where the timestamp value is equal to
	// or after from. If from is not set and to is set: match items where the timestamp
	// value is equal to or before to.
	LastUpdatedTime *DateTimeRange

	// A list of AWS Regions.
	Regions []*string

	// The AWS services associated with the event. For example, EC2, RDS.
	Services []*string

	// A range of dates and times that is used by the EventFilter
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
	// and EntityFilter
	// (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
	// objects. If from is set and to is set: match items where the timestamp
	// (startTime, endTime, or lastUpdatedTime) is between from and to inclusive. If
	// from is set and to is not set: match items where the timestamp value is equal to
	// or after from. If from is not set and to is set: match items where the timestamp
	// value is equal to or before to.
	StartTime *DateTimeRange
}
