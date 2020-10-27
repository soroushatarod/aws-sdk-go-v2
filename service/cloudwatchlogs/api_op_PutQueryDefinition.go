// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a query definition for CloudWatch Logs Insights. For more
// information, see Analyzing Log Data with CloudWatch Logs Insights
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html).
// To update a query definition, specify its queryDefinitionId in your request. The
// values of name, queryString, and logGroupNames are changed to the values that
// you specify in your update operation. No current values are retained from the
// current query definition. For example, if you update a current query definition
// that includes log groups, and you don't specify the logGroupNames parameter in
// your update operation, the query definition changes to contain no log groups.
// You must have the logs:PutQueryDefinition permission to be able to perform this
// operation.
func (c *Client) PutQueryDefinition(ctx context.Context, params *PutQueryDefinitionInput, optFns ...func(*Options)) (*PutQueryDefinitionOutput, error) {
	if params == nil {
		params = &PutQueryDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutQueryDefinition", params, optFns, addOperationPutQueryDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutQueryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutQueryDefinitionInput struct {

	// A name for the query definition. If you are saving a lot of query definitions,
	// we recommend that you name them so that you can easily find the ones you want by
	// using the first part of the name as a filter in the queryDefinitionNamePrefix
	// parameter of DescribeQueryDefinitions
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeQueryDefinitions.html).
	//
	// This member is required.
	Name *string

	// The query string to use for this definition. For more information, see
	// CloudWatch Logs Insights Query Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	//
	// This member is required.
	QueryString *string

	// Use this parameter to include specific log groups as part of your query
	// definition. If you are updating a query definition and you omit this parameter,
	// then the updated definition will contain no log groups.
	LogGroupNames []*string

	// If you are updating a query definition, use this parameter to specify the ID of
	// the query definition that you want to update. You can use
	// DescribeQueryDefinitions
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeQueryDefinitions.html)
	// to retrieve the IDs of your saved query definitions. If you are creating a query
	// definition, do not specify this parameter. CloudWatch generates a unique ID for
	// the new query definition and include it in the response to this operation.
	QueryDefinitionId *string
}

type PutQueryDefinitionOutput struct {

	// The ID of the query definition.
	QueryDefinitionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutQueryDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutQueryDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutQueryDefinition{}, middleware.After)
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
	addOpPutQueryDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutQueryDefinition(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutQueryDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutQueryDefinition",
	}
}
