// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the detailed information about a table from metadata in the cluster.
// The information includes its columns. A token is returned to page through the
// column list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
//     * AWS Secrets Manager - specify the
// Amazon Resource Name (ARN) of the secret and the cluster identifier that matches
// the cluster in the secret.
//
//     * Temporary credentials - specify the cluster
// identifier, the database name, and the database user name. Permission to call
// the redshift:GetClusterCredentials operation is required to use this method.
func (c *Client) DescribeTable(ctx context.Context, params *DescribeTableInput, optFns ...func(*Options)) (*DescribeTableOutput, error) {
	if params == nil {
		params = &DescribeTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTable", params, optFns, addOperationDescribeTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTableInput struct {

	// The cluster identifier. This parameter is required when authenticating using
	// either AWS Secrets Manager or temporary credentials.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database. This parameter is required when authenticating using
	// temporary credentials.
	Database *string

	// The database user name. This parameter is required when authenticating using
	// temporary credentials.
	DbUser *string

	// The maximum number of tables to return in the response. If more tables exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults *int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The schema that contains the table. If no schema is specified, then matching
	// tables for all schemas are returned.
	Schema *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using AWS Secrets Manager.
	SecretArn *string

	// The table name. If no table is specified, then all tables for all matching
	// schemas are returned. If no table and no schema is specified, then all tables
	// for all schemas in the database are returned
	Table *string
}

type DescribeTableOutput struct {

	// A list of columns in the table.
	ColumnList []*types.ColumnMetadata

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The table name.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTable{}, middleware.After)
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
	addOpDescribeTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "DescribeTable",
	}
}
