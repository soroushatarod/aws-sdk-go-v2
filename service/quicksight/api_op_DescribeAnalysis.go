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

// Provides a summary of the metadata for an analysis.
func (c *Client) DescribeAnalysis(ctx context.Context, params *DescribeAnalysisInput, optFns ...func(*Options)) (*DescribeAnalysisOutput, error) {
	if params == nil {
		params = &DescribeAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAnalysis", params, optFns, addOperationDescribeAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAnalysisInput struct {

	// The ID of the analysis that you're describing. The ID is part of the URL of the
	// analysis.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the AWS account that contains the analysis. You must be using the AWS
	// account that the analysis is in.
	//
	// This member is required.
	AwsAccountId *string
}

type DescribeAnalysisOutput struct {

	// A metadata structure that contains summary information for the analysis that
	// you're describing.
	Analysis *types.Analysis

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAnalysis{}, middleware.After)
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
	addOpDescribeAnalysisValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnalysis(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAnalysis(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAnalysis",
	}
}
