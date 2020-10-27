// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a saved CloudWatch Logs Insights query definition. A query definition
// contains details about a saved CloudWatch Logs Insights query. Each
// DeleteQueryDefinition operation can delete one query definition. You must have
// the logs:DeleteQueryDefinition permission to be able to perform this operation.
func (c *Client) DeleteQueryDefinition(ctx context.Context, params *DeleteQueryDefinitionInput, optFns ...func(*Options)) (*DeleteQueryDefinitionOutput, error) {
	if params == nil {
		params = &DeleteQueryDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteQueryDefinition", params, optFns, addOperationDeleteQueryDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteQueryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQueryDefinitionInput struct {

	// The ID of the query definition that you want to delete. You can use
	// DescribeQueryDefinitions
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeQueryDefinitions.html)
	// to retrieve the IDs of your saved query definitions.
	//
	// This member is required.
	QueryDefinitionId *string
}

type DeleteQueryDefinitionOutput struct {

	// A value of TRUE indicates that the operation succeeded. FALSE indicates that the
	// operation failed.
	Success *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteQueryDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteQueryDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteQueryDefinition{}, middleware.After)
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
	addOpDeleteQueryDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueryDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteQueryDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DeleteQueryDefinition",
	}
}
