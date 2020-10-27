// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or removes permission settings for the specified snapshot. You may add or
// remove specified AWS account IDs from a snapshot's list of create volume
// permissions, but you cannot do both in a single operation. If you need to both
// add and remove account IDs for a snapshot, you must use multiple operations. You
// can make up to 500 modifications to a snapshot in a single operation. Encrypted
// snapshots and snapshots with AWS Marketplace product codes cannot be made
// public. Snapshots encrypted with your default CMK cannot be shared with other
// accounts. For more information about modifying snapshot permissions, see Sharing
// snapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modifying-snapshot-permissions.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifySnapshotAttribute(ctx context.Context, params *ModifySnapshotAttributeInput, optFns ...func(*Options)) (*ModifySnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifySnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifySnapshotAttribute", params, optFns, addOperationModifySnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifySnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifySnapshotAttributeInput struct {

	// The ID of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// The snapshot attribute to modify. Only volume creation permissions can be
	// modified.
	Attribute types.SnapshotAttributeName

	// A JSON representation of the snapshot attribute modification.
	CreateVolumePermission *types.CreateVolumePermissionModifications

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The group to modify for the snapshot.
	GroupNames []*string

	// The type of operation to perform to the attribute.
	OperationType types.OperationType

	// The account ID to modify for the snapshot.
	UserIds []*string
}

type ModifySnapshotAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifySnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifySnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifySnapshotAttribute{}, middleware.After)
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
	addOpModifySnapshotAttributeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifySnapshotAttribute(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifySnapshotAttribute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifySnapshotAttribute",
	}
}
