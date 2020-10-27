// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes an existing IAM policy assignment, as specified by the assignment
// name.
func (c *Client) DescribeIAMPolicyAssignment(ctx context.Context, params *DescribeIAMPolicyAssignmentInput, optFns ...func(*Options)) (*DescribeIAMPolicyAssignmentOutput, error) {
	if params == nil {
		params = &DescribeIAMPolicyAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIAMPolicyAssignment", params, optFns, addOperationDescribeIAMPolicyAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIAMPolicyAssignmentInput struct {

	// The name of the assignment.
	//
	// This member is required.
	AssignmentName *string

	// The ID of the AWS account that contains the assignment that you want to
	// describe.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that contains the assignment.
	//
	// This member is required.
	Namespace *string
}

type DescribeIAMPolicyAssignmentOutput struct {

	// Information describing the IAM policy assignment.
	IAMPolicyAssignment *types.IAMPolicyAssignment

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeIAMPolicyAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIAMPolicyAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIAMPolicyAssignment{}, middleware.After)
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
	addOpDescribeIAMPolicyAssignmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIAMPolicyAssignment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeIAMPolicyAssignment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeIAMPolicyAssignment",
	}
}
