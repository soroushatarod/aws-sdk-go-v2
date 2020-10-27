// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a game session placement that is in PENDING status. To stop a placement,
// provide the placement ID values. If successful, the placement is moved to
// CANCELLED status.
//
//     * CreateGameSession
//
//     * DescribeGameSessions
//
//     *
// DescribeGameSessionDetails
//
//     * SearchGameSessions
//
//     * UpdateGameSession
//
//
// * GetGameSessionLogUrl
//
//     * Game session placements
//
//         *
// StartGameSessionPlacement
//
//         * DescribeGameSessionPlacement
//
//         *
// StopGameSessionPlacement
func (c *Client) StopGameSessionPlacement(ctx context.Context, params *StopGameSessionPlacementInput, optFns ...func(*Options)) (*StopGameSessionPlacementOutput, error) {
	if params == nil {
		params = &StopGameSessionPlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopGameSessionPlacement", params, optFns, addOperationStopGameSessionPlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type StopGameSessionPlacementInput struct {

	// A unique identifier for a game session placement to cancel.
	//
	// This member is required.
	PlacementId *string
}

// Represents the returned data in response to a request operation.
type StopGameSessionPlacementOutput struct {

	// Object that describes the canceled game session placement, with CANCELLED status
	// and an end time stamp.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopGameSessionPlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopGameSessionPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopGameSessionPlacement{}, middleware.After)
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
	addOpStopGameSessionPlacementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopGameSessionPlacement(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopGameSessionPlacement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "StopGameSessionPlacement",
	}
}
