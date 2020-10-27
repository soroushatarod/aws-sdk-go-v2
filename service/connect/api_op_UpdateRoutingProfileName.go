// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the name and description of a routing profile. The request accepts the
// following data in JSON format. At least Name or Description must be provided.
func (c *Client) UpdateRoutingProfileName(ctx context.Context, params *UpdateRoutingProfileNameInput, optFns ...func(*Options)) (*UpdateRoutingProfileNameOutput, error) {
	if params == nil {
		params = &UpdateRoutingProfileNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoutingProfileName", params, optFns, addOperationUpdateRoutingProfileNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoutingProfileNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoutingProfileNameInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier of the routing profile.
	//
	// This member is required.
	RoutingProfileId *string

	// The description of the routing profile. Must not be more than 250 characters.
	Description *string

	// The name of the routing profile. Must not be more than 127 characters.
	Name *string
}

type UpdateRoutingProfileNameOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRoutingProfileNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoutingProfileName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoutingProfileName{}, middleware.After)
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
	addOpUpdateRoutingProfileNameValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoutingProfileName(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRoutingProfileName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateRoutingProfileName",
	}
}
