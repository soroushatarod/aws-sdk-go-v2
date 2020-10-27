// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current status of an asynchronous request to create an account.
// This operation can be called only from the organization's management account or
// by a member account that is a delegated administrator for an AWS service.
func (c *Client) DescribeCreateAccountStatus(ctx context.Context, params *DescribeCreateAccountStatusInput, optFns ...func(*Options)) (*DescribeCreateAccountStatusOutput, error) {
	if params == nil {
		params = &DescribeCreateAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCreateAccountStatus", params, optFns, addOperationDescribeCreateAccountStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCreateAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCreateAccountStatusInput struct {

	// Specifies the Id value that uniquely identifies the CreateAccount request. You
	// can get the value from the CreateAccountStatus.Id response in an earlier
	// CreateAccount request, or from the ListCreateAccountStatus operation. The regex
	// pattern (http://wikipedia.org/wiki/regex) for a create account request ID string
	// requires "car-" followed by from 8 to 32 lowercase letters or digits.
	//
	// This member is required.
	CreateAccountRequestId *string
}

type DescribeCreateAccountStatusOutput struct {

	// A structure that contains the current status of an account creation request.
	CreateAccountStatus *types.CreateAccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCreateAccountStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCreateAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCreateAccountStatus{}, middleware.After)
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
	addOpDescribeCreateAccountStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCreateAccountStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeCreateAccountStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeCreateAccountStatus",
	}
}
