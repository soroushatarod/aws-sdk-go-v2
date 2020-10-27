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

// Lists the Amazon QuickSight groups that an Amazon QuickSight user is a member
// of.
func (c *Client) ListUserGroups(ctx context.Context, params *ListUserGroupsInput, optFns ...func(*Options)) (*ListUserGroupsOutput, error) {
	if params == nil {
		params = &ListUserGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUserGroups", params, optFns, addOperationListUserGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUserGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUserGroupsInput struct {

	// The AWS account ID that the user is in. Currently, you use the ID for the AWS
	// account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace. Currently, you should set this to default.
	//
	// This member is required.
	Namespace *string

	// The Amazon QuickSight user name that you want to list group memberships for.
	//
	// This member is required.
	UserName *string

	// The maximum number of results to return from this request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string
}

type ListUserGroupsOutput struct {

	// The list of groups the user is a member of.
	GroupList []*types.Group

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListUserGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUserGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUserGroups{}, middleware.After)
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
	addOpListUserGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUserGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListUserGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListUserGroups",
	}
}
