// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of entities that have been affected by one or more events for one
// or more accounts in your organization in AWS Organizations, based on the filter
// criteria. Entities can refer to individual customer resources, groups of
// customer resources, or any other construct, depending on the AWS service. At
// least one event Amazon Resource Name (ARN) and account ID are required. Results
// are sorted by the lastUpdatedTime of the entity, starting with the most recent.
// Before you can call this operation, you must first enable AWS Health to work
// with AWS Organizations. To do this, call the
// EnableHealthServiceAccessForOrganization
// (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// operation from your organization's master account. This API operation uses
// pagination. Specify the nextToken parameter in the next request to return more
// results.
func (c *Client) DescribeAffectedEntitiesForOrganization(ctx context.Context, params *DescribeAffectedEntitiesForOrganizationInput, optFns ...func(*Options)) (*DescribeAffectedEntitiesForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeAffectedEntitiesForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAffectedEntitiesForOrganization", params, optFns, addOperationDescribeAffectedEntitiesForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAffectedEntitiesForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAffectedEntitiesForOrganizationInput struct {

	// A JSON set of elements including the awsAccountId and the eventArn.
	//
	// This member is required.
	OrganizationEntityFilters []*types.EventAccountFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	// The maximum number of items to return in one batch, between 10 and 100,
	// inclusive.
	MaxResults *int32

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string
}

type DescribeAffectedEntitiesForOrganizationOutput struct {

	// A JSON set of elements including the awsAccountId and its entityArn, entityValue
	// and its entityArn, lastUpdatedTime, and statusCode.
	Entities []*types.AffectedEntity

	// A JSON set of elements of the failed response, including the awsAccountId,
	// errorMessage, errorName, and eventArn.
	FailedSet []*types.OrganizationAffectedEntitiesErrorItem

	// If the results of a search are large, only a portion of the results are
	// returned, and a nextToken pagination token is returned in the response. To
	// retrieve the next batch of results, reissue the search request and include the
	// returned token. When all results have been returned, the response does not
	// contain a pagination token value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAffectedEntitiesForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAffectedEntitiesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAffectedEntitiesForOrganization{}, middleware.After)
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
	addOpDescribeAffectedEntitiesForOrganizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAffectedEntitiesForOrganization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAffectedEntitiesForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeAffectedEntitiesForOrganization",
	}
}
