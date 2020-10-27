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

// Lists Amazon QuickSight analyses that exist in the specified AWS account.
func (c *Client) ListAnalyses(ctx context.Context, params *ListAnalysesInput, optFns ...func(*Options)) (*ListAnalysesOutput, error) {
	if params == nil {
		params = &ListAnalysesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnalyses", params, optFns, addOperationListAnalysesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnalysesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnalysesInput struct {

	// The ID of the AWS account that contains the analyses.
	//
	// This member is required.
	AwsAccountId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string
}

type ListAnalysesOutput struct {

	// Metadata describing each of the analyses that are listed.
	AnalysisSummaryList []*types.AnalysisSummary

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAnalysesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnalyses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnalyses{}, middleware.After)
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
	addOpListAnalysesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAnalyses(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAnalyses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListAnalyses",
	}
}
