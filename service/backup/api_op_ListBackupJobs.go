// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a list of existing backup jobs for an authenticated account.
func (c *Client) ListBackupJobs(ctx context.Context, params *ListBackupJobsInput, optFns ...func(*Options)) (*ListBackupJobsOutput, error) {
	if params == nil {
		params = &ListBackupJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackupJobs", params, optFns, addOperationListBackupJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackupJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBackupJobsInput struct {

	// The account ID to list the jobs from. Returns only backup jobs associated with
	// the specified account ID.
	ByAccountId *string

	// Returns only backup jobs that will be stored in the specified backup vault.
	// Backup vaults are identified by names that are unique to the account used to
	// create them and the AWS Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	ByBackupVaultName *string

	// Returns only backup jobs that were created after the specified date.
	ByCreatedAfter *time.Time

	// Returns only backup jobs that were created before the specified date.
	ByCreatedBefore *time.Time

	// Returns only backup jobs that match the specified resource Amazon Resource Name
	// (ARN).
	ByResourceArn *string

	// Returns only backup jobs for the specified resources:
	//
	//     * DynamoDB for Amazon
	// DynamoDB
	//
	//     * EBS for Amazon Elastic Block Store
	//
	//     * EC2 for Amazon Elastic
	// Compute Cloud
	//
	//     * EFS for Amazon Elastic File System
	//
	//     * RDS for Amazon
	// Relational Database Service
	//
	//     * Storage Gateway for AWS Storage Gateway
	ByResourceType *string

	// Returns only backup jobs that are in the specified state.
	ByState types.BackupJobState

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
}

type ListBackupJobsOutput struct {

	// An array of structures containing metadata about your backup jobs returned in
	// JSON format.
	BackupJobs []*types.BackupJob

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBackupJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackupJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackupJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBackupJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListBackupJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListBackupJobs",
	}
}
