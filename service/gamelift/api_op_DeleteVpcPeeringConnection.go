// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a VPC peering connection. To delete the connection, you must have a
// valid authorization for the VPC peering connection that you want to delete. You
// can check for an authorization by calling DescribeVpcPeeringAuthorizations or
// request a new one using CreateVpcPeeringAuthorization. Once a valid
// authorization exists, call this operation from the AWS account that is used to
// manage the Amazon GameLift fleets. Identify the connection to delete by the
// connection ID and fleet ID. If successful, the connection is removed.
//
//     *
// CreateVpcPeeringAuthorization
//
//     * DescribeVpcPeeringAuthorizations
//
//     *
// DeleteVpcPeeringAuthorization
//
//     * CreateVpcPeeringConnection
//
//     *
// DescribeVpcPeeringConnections
//
//     * DeleteVpcPeeringConnection
func (c *Client) DeleteVpcPeeringConnection(ctx context.Context, params *DeleteVpcPeeringConnectionInput, optFns ...func(*Options)) (*DeleteVpcPeeringConnectionOutput, error) {
	if params == nil {
		params = &DeleteVpcPeeringConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVpcPeeringConnection", params, optFns, addOperationDeleteVpcPeeringConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteVpcPeeringConnectionInput struct {

	// A unique identifier for a fleet. This fleet specified must match the fleet
	// referenced in the VPC peering connection record. You can use either the fleet ID
	// or ARN value.
	//
	// This member is required.
	FleetId *string

	// A unique identifier for a VPC peering connection. This value is included in the
	// VpcPeeringConnection object, which can be retrieved by calling
	// DescribeVpcPeeringConnections.
	//
	// This member is required.
	VpcPeeringConnectionId *string
}

type DeleteVpcPeeringConnectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteVpcPeeringConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteVpcPeeringConnection{}, middleware.After)
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
	addOpDeleteVpcPeeringConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpcPeeringConnection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteVpcPeeringConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteVpcPeeringConnection",
	}
}
