// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon Chime SDK meeting in the specified media Region with no
// initial attendees. For more information about specifying media Regions, see
// Amazon Chime SDK Media Regions
// (https://docs.aws.amazon.com/chime/latest/dg/chime-sdk-meetings-regions.html) in
// the Amazon Chime Developer Guide. For more information about the Amazon Chime
// SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
func (c *Client) CreateMeeting(ctx context.Context, params *CreateMeetingInput, optFns ...func(*Options)) (*CreateMeetingOutput, error) {
	if params == nil {
		params = &CreateMeetingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMeeting", params, optFns, addOperationCreateMeetingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMeetingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMeetingInput struct {

	// The unique identifier for the client request. Use a different token for
	// different meetings.
	//
	// This member is required.
	ClientRequestToken *string

	// The external meeting ID.
	ExternalMeetingId *string

	// The Region in which to create the meeting. Default: us-east-1. Available values:
	// af-south-1, ap-northeast-1, ap-northeast-2, ap-south-1, ap-southeast-1,
	// ap-southeast-2, ca-central-1, eu-central-1, eu-north-1, eu-south-1, eu-west-1,
	// eu-west-2, eu-west-3, sa-east-1, us-east-1, us-east-2, us-west-1, us-west-2.
	MediaRegion *string

	// Reserved.
	MeetingHostId *string

	// The configuration for resource targets to receive notifications when meeting and
	// attendee events occur.
	NotificationsConfiguration *types.MeetingNotificationConfiguration

	// The tag key-value pairs.
	Tags []*types.Tag
}

type CreateMeetingOutput struct {

	// The meeting information, including the meeting ID and MediaPlacement.
	Meeting *types.Meeting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMeetingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMeeting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMeeting{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addIdempotencyToken_opCreateMeetingMiddleware(stack, options)
	addOpCreateMeetingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMeeting(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateMeeting struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMeeting) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMeeting) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMeetingInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMeetingInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMeetingMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateMeeting{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMeeting(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateMeeting",
	}
}
