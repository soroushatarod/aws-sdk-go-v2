// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing workflow.
func (c *Client) UpdateWorkflow(ctx context.Context, params *UpdateWorkflowInput, optFns ...func(*Options)) (*UpdateWorkflowOutput, error) {
	if params == nil {
		params = &UpdateWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkflow", params, optFns, addOperationUpdateWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkflowInput struct {

	// Name of the workflow to be updated.
	//
	// This member is required.
	Name *string

	// A collection of properties to be used as part of each execution of the workflow.
	DefaultRunProperties map[string]*string

	// The description of the workflow.
	Description *string

	// You can use this parameter to prevent unwanted multiple updates to data, to
	// control costs, or in some cases, to prevent exceeding the maximum number of
	// concurrent runs of any of the component jobs. If you leave this parameter blank,
	// there is no limit to the number of concurrent workflow runs.
	MaxConcurrentRuns *int32
}

type UpdateWorkflowOutput struct {

	// The name of the workflow which was specified in input.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkflow{}, middleware.After)
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
	addOpUpdateWorkflowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkflow(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateWorkflow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateWorkflow",
	}
}
