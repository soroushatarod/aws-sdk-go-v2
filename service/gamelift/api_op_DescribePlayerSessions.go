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

// Retrieves properties for one or more player sessions. This operation can be used
// in several ways: (1) provide a PlayerSessionId to request properties for a
// specific player session; (2) provide a GameSessionId to request properties for
// all player sessions in the specified game session; (3) provide a PlayerId to
// request properties for all player sessions of a specified player. To get game
// session record(s), specify only one of the following: a player session ID, a
// game session ID, or a player ID. You can filter this request by player session
// status. Use the pagination parameters to retrieve results as a set of sequential
// pages. If successful, a PlayerSession object is returned for each session
// matching the request. Available in Amazon GameLift Local.
//
//     *
// CreatePlayerSession
//
//     * CreatePlayerSessions
//
//     * DescribePlayerSessions
//
//
// * Game session placements
//
//         * StartGameSessionPlacement
//
//         *
// DescribeGameSessionPlacement
//
//         * StopGameSessionPlacement
func (c *Client) DescribePlayerSessions(ctx context.Context, params *DescribePlayerSessionsInput, optFns ...func(*Options)) (*DescribePlayerSessionsOutput, error) {
	if params == nil {
		params = &DescribePlayerSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePlayerSessions", params, optFns, addOperationDescribePlayerSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePlayerSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribePlayerSessionsInput struct {

	// A unique identifier for the game session to retrieve player sessions for.
	GameSessionId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. If a player session ID is specified,
	// this parameter is ignored.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value. If a player session ID is
	// specified, this parameter is ignored.
	NextToken *string

	// A unique identifier for a player to retrieve player sessions for.
	PlayerId *string

	// A unique identifier for a player session to retrieve.
	PlayerSessionId *string

	// Player session status to filter results on. Possible player session statuses
	// include the following:
	//
	//     * RESERVED -- The player session request has been
	// received, but the player has not yet connected to the server process and/or been
	// validated.
	//
	//     * ACTIVE -- The player has been validated by the server process
	// and is currently connected.
	//
	//     * COMPLETED -- The player connection has been
	// dropped.
	//
	//     * TIMEDOUT -- A player session request was received, but the
	// player did not connect and/or was not validated within the timeout limit (60
	// seconds).
	PlayerSessionStatusFilter *string
}

// Represents the returned data in response to a request operation.
type DescribePlayerSessionsOutput struct {

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// A collection of objects containing properties for each player session that
	// matches the request.
	PlayerSessions []*types.PlayerSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePlayerSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePlayerSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePlayerSessions{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlayerSessions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePlayerSessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribePlayerSessions",
	}
}
