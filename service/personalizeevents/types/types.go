// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Represents user interaction event information sent using the PutEvents API.
type Event struct {

	// The type of event, such as click or download. This property corresponds to the
	// EVENT_TYPE field of your Interactions schema and depends on the types of events
	// you are tracking.
	//
	// This member is required.
	EventType *string

	// The timestamp (in Unix time) on the client side when the event occurred.
	//
	// This member is required.
	SentAt *time.Time

	// An ID associated with the event. If an event ID is not provided, Amazon
	// Personalize generates a unique ID for the event. An event ID is not used as an
	// input to the model. Amazon Personalize uses the event ID to distinquish unique
	// events. Any subsequent events after the first with the same event ID are not
	// used in model training.
	EventId *string

	// The event value that corresponds to the EVENT_VALUE field of the Interactions
	// schema.
	EventValue *float32

	// A list of item IDs that represents the sequence of items you have shown the
	// user. For example, ["itemId1", "itemId2", "itemId3"].
	Impression []*string

	// The item ID key that corresponds to the ITEM_ID field of the Interactions
	// schema.
	ItemId *string

	// A string map of event-specific data that you might choose to record. For
	// example, if a user rates a movie on your site, other than movie ID (itemId) and
	// rating (eventValue) , you might also send the number of movie ratings made by
	// the user. Each item in the map consists of a key-value pair. For example,
	// {"numberOfRatings": "12"} The keys use camel case names that match the fields in
	// the Interactions schema. In the above example, the numberOfRatings would match
	// the 'NUMBER_OF_RATINGS' field defined in the Interactions schema.
	// This value conforms to the media type: application/json
	Properties *string

	// The ID of the recommendation.
	RecommendationId *string
}

// Represents item metadata added to an Items dataset using the PutItems API.
type Item struct {

	// The ID associated with the item.
	//
	// This member is required.
	ItemId *string

	// A string map of item-specific metadata. Each element in the map consists of a
	// key-value pair. For example, {"numberOfRatings": "12"} The keys use camel case
	// names that match the fields in the Items schema. In the above example, the
	// numberOfRatings would match the 'NUMBER_OF_RATINGS' field defined in the Items
	// schema.
	// This value conforms to the media type: application/json
	Properties *string
}

// Represents user metadata added to a Users dataset using the PutUsers API.
type User struct {

	// The ID associated with the user.
	//
	// This member is required.
	UserId *string

	// A string map of user-specific metadata. Each element in the map consists of a
	// key-value pair. For example, {"numberOfVideosWatched": "45"} The keys use camel
	// case names that match the fields in the Users schema. In the above example, the
	// numberOfVideosWatched would match the 'NUMBER_OF_VIDEOS_WATCHED' field defined
	// in the Users schema.
	// This value conforms to the media type: application/json
	Properties *string
}
