// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates a Device Defender audit suppression.
func (c *Client) UpdateAuditSuppression(ctx context.Context, params *UpdateAuditSuppressionInput, optFns ...func(*Options)) (*UpdateAuditSuppressionOutput, error) {
	if params == nil {
		params = &UpdateAuditSuppressionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAuditSuppression", params, optFns, addOperationUpdateAuditSuppressionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAuditSuppressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAuditSuppressionInput struct {

	// An audit check name. Checks must be enabled for your account. (Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or use UpdateAccountAuditConfiguration to select which checks
	// are enabled.)
	//
	// This member is required.
	CheckName *string

	// Information that identifies the noncompliant resource.
	//
	// This member is required.
	ResourceIdentifier *types.ResourceIdentifier

	// The description of the audit suppression.
	Description *string

	// The expiration date (epoch timestamp in seconds) that you want the suppression
	// to adhere to.
	ExpirationDate *time.Time

	// Indicates whether a suppression should exist indefinitely or not.
	SuppressIndefinitely *bool
}

type UpdateAuditSuppressionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAuditSuppressionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAuditSuppression{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAuditSuppression{}, middleware.After)
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
	addOpUpdateAuditSuppressionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAuditSuppression(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateAuditSuppression(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateAuditSuppression",
	}
}
