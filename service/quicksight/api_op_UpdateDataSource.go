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

// Updates a data source.
func (c *Client) UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) {
	if params == nil {
		params = &UpdateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSource", params, optFns, addOperationUpdateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourceInput struct {

	// The AWS account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	//
	// This member is required.
	DataSourceId *string

	// A display name for the data source.
	//
	// This member is required.
	Name *string

	// The credentials that QuickSight that uses to connect to your underlying source.
	// Currently, only credentials based on user name and password are supported.
	Credentials *types.DataSourceCredentials

	// The parameters that QuickSight uses to connect to your underlying source.
	DataSourceParameters *types.DataSourceParameters

	// Secure Socket Layer (SSL) properties that apply when QuickSight connects to your
	// underlying source.
	SslProperties *types.SslProperties

	// Use this parameter only when you want QuickSight to use a VPC connection when
	// connecting to your underlying source.
	VpcConnectionProperties *types.VpcConnectionProperties
}

type UpdateDataSourceOutput struct {

	// The Amazon Resource Name (ARN) of the data source.
	Arn *string

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// The update status of the data source's last update.
	UpdateStatus types.ResourceStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataSource{}, middleware.After)
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
	addOpUpdateDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDataSource",
	}
}
