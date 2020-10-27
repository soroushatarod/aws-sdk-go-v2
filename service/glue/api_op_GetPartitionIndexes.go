// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the partition indexes associated with a table.
func (c *Client) GetPartitionIndexes(ctx context.Context, params *GetPartitionIndexesInput, optFns ...func(*Options)) (*GetPartitionIndexesOutput, error) {
	if params == nil {
		params = &GetPartitionIndexesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPartitionIndexes", params, optFns, addOperationGetPartitionIndexesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPartitionIndexesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPartitionIndexesInput struct {

	// Specifies the name of a database from which you want to retrieve partition
	// indexes.
	//
	// This member is required.
	DatabaseName *string

	// Specifies the name of a table for which you want to retrieve the partition
	// indexes.
	//
	// This member is required.
	TableName *string

	// The catalog ID where the table resides.
	CatalogId *string

	// A continuation token, included if this is a continuation call.
	NextToken *string
}

type GetPartitionIndexesOutput struct {

	// A continuation token, present if the current list segment is not the last.
	NextToken *string

	// A list of index descriptors.
	PartitionIndexDescriptorList []*types.PartitionIndexDescriptor

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPartitionIndexesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPartitionIndexes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPartitionIndexes{}, middleware.After)
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
	addOpGetPartitionIndexesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPartitionIndexes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPartitionIndexes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetPartitionIndexes",
	}
}
