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

// This operation is used with the Amazon GameLift FleetIQ solution and game server
// groups. Temporarily stops activity on a game server group without terminating
// instances or the game server group. You can restart activity by calling
// ResumeGameServerGroup. You can suspend the following activity:
//
//     * Instance
// type replacement - This activity evaluates the current game hosting viability of
// all Spot instance types that are defined for the game server group. It updates
// the Auto Scaling group to remove nonviable Spot Instance types, which have a
// higher chance of game server interruptions. It then balances capacity across the
// remaining viable Spot Instance types. When this activity is suspended, the Auto
// Scaling group continues with its current balance, regardless of viability.
// Instance protection, utilization metrics, and capacity scaling activities
// continue to be active.
//
// To suspend activity, specify a game server group ARN and
// the type of activity to be suspended. If successful, a GameServerGroup object is
// returned showing that the activity is listed in SuspendedActions. Learn more
// GameLift FleetIQ Guide
// (https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html)
// Related operations
//
//     * CreateGameServerGroup
//
//     * ListGameServerGroups
//
//
// * DescribeGameServerGroup
//
//     * UpdateGameServerGroup
//
//     *
// DeleteGameServerGroup
//
//     * ResumeGameServerGroup
//
//     *
// SuspendGameServerGroup
//
//     * DescribeGameServerInstances
func (c *Client) SuspendGameServerGroup(ctx context.Context, params *SuspendGameServerGroupInput, optFns ...func(*Options)) (*SuspendGameServerGroupOutput, error) {
	if params == nil {
		params = &SuspendGameServerGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SuspendGameServerGroup", params, optFns, addOperationSuspendGameServerGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SuspendGameServerGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SuspendGameServerGroupInput struct {

	// A unique identifier for the game server group. Use either the GameServerGroup
	// name or ARN value.
	//
	// This member is required.
	GameServerGroupName *string

	// The activity to suspend for this game server group.
	//
	// This member is required.
	SuspendActions []types.GameServerGroupAction
}

type SuspendGameServerGroupOutput struct {

	// An object that describes the game server group resource, with the
	// SuspendedActions property updated to reflect the suspended activity.
	GameServerGroup *types.GameServerGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSuspendGameServerGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSuspendGameServerGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSuspendGameServerGroup{}, middleware.After)
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
	addOpSuspendGameServerGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSuspendGameServerGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSuspendGameServerGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "SuspendGameServerGroup",
	}
}
