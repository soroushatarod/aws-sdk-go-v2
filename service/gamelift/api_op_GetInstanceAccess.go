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

// Requests remote access to a fleet instance. Remote access is useful for
// debugging, gathering benchmarking data, or observing activity in real time. To
// remotely access an instance, you need credentials that match the operating
// system of the instance. For a Windows instance, Amazon GameLift returns a user
// name and password as strings for use with a Windows Remote Desktop client. For a
// Linux instance, Amazon GameLift returns a user name and RSA private key, also as
// strings, for use with an SSH client. The private key must be saved in the proper
// format to a .pem file before using. If you're making this request using the AWS
// CLI, saving the secret can be handled as part of the GetInstanceAccess request,
// as shown in one of the examples for this operation. To request access to a
// specific instance, specify the IDs of both the instance and the fleet it belongs
// to. You can retrieve a fleet's instance IDs by calling DescribeInstances. If
// successful, an InstanceAccess object is returned that contains the instance's IP
// address and a set of credentials. Learn more Remotely Access Fleet Instances
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html)Debug
// Fleet Issues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html)
// Related operations
//
//     * DescribeInstances
//
//     * GetInstanceAccess
func (c *Client) GetInstanceAccess(ctx context.Context, params *GetInstanceAccessInput, optFns ...func(*Options)) (*GetInstanceAccessOutput, error) {
	if params == nil {
		params = &GetInstanceAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstanceAccess", params, optFns, addOperationGetInstanceAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstanceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type GetInstanceAccessInput struct {

	// A unique identifier for a fleet that contains the instance you want access to.
	// You can use either the fleet ID or ARN value. The fleet can be in any of the
	// following statuses: ACTIVATING, ACTIVE, or ERROR. Fleets with an ERROR status
	// may be accessible for a short time before they are deleted.
	//
	// This member is required.
	FleetId *string

	// A unique identifier for an instance you want to get access to. You can access an
	// instance in any status.
	//
	// This member is required.
	InstanceId *string
}

// Represents the returned data in response to a request operation.
type GetInstanceAccessOutput struct {

	// The connection information for a fleet instance, including IP address and access
	// credentials.
	InstanceAccess *types.InstanceAccess

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetInstanceAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstanceAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstanceAccess{}, middleware.After)
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
	addOpGetInstanceAccessValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstanceAccess(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetInstanceAccess(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "GetInstanceAccess",
	}
}
