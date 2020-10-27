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

// Creates and starts a new SPICE ingestion on a dataset Any ingestions operating
// on tagged datasets inherit the same tags automatically for use in access
// control. For an example, see How do I create an IAM policy to control access to
// Amazon EC2 resources using tags?
// (https://aws.amazon.com/premiumsupport/knowledge-center/iam-ec2-resource-tags/)
// in the AWS Knowledge Center. Tags are visible on the tagged dataset, but not on
// the ingestion resource.
func (c *Client) CreateIngestion(ctx context.Context, params *CreateIngestionInput, optFns ...func(*Options)) (*CreateIngestionOutput, error) {
	if params == nil {
		params = &CreateIngestionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIngestion", params, optFns, addOperationCreateIngestionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIngestionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIngestionInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the dataset used in the ingestion.
	//
	// This member is required.
	DataSetId *string

	// An ID for the ingestion.
	//
	// This member is required.
	IngestionId *string
}

type CreateIngestionOutput struct {

	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string

	// An ID for the ingestion.
	IngestionId *string

	// The ingestion status.
	IngestionStatus types.IngestionStatus

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateIngestionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIngestion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIngestion{}, middleware.After)
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
	addOpCreateIngestionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIngestion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateIngestion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateIngestion",
	}
}
