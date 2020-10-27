// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the list of users that belong to the user group.
func (c *Client) ModifyUserGroup(ctx context.Context, params *ModifyUserGroupInput, optFns ...func(*Options)) (*ModifyUserGroupOutput, error) {
	if params == nil {
		params = &ModifyUserGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyUserGroup", params, optFns, addOperationModifyUserGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyUserGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyUserGroupInput struct {

	// The ID of the user group.
	//
	// This member is required.
	UserGroupId *string

	// The list of user IDs to add to the user group.
	UserIdsToAdd []*string

	// The list of user IDs to remove from the user group.
	UserIdsToRemove []*string
}

type ModifyUserGroupOutput struct {

	// The Amazon Resource Name (ARN) of the user group.
	ARN *string

	// Must be Redis.
	Engine *string

	// A list of updates being applied to the user groups.
	PendingChanges *types.UserGroupPendingChanges

	// A list of replication groups that the user group can access.
	ReplicationGroups []*string

	// Indicates user group status. Can be "creating", "active", "modifying",
	// "deleting".
	Status *string

	// The ID of the user group.
	UserGroupId *string

	// The list of user IDs that belong to the user group.
	UserIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyUserGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyUserGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyUserGroup{}, middleware.After)
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
	addOpModifyUserGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyUserGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyUserGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyUserGroup",
	}
}
