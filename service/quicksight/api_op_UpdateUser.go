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

// Updates an Amazon QuickSight user.
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	if params == nil {
		params = &UpdateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUser", params, optFns, addOperationUpdateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {

	// The ID for the AWS account that the user is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The email address of the user that you want to update.
	//
	// This member is required.
	Email *string

	// The namespace. Currently, you should set this to default.
	//
	// This member is required.
	Namespace *string

	// The Amazon QuickSight role of the user. The role can be one of the following
	// default security cohorts:
	//
	//     * READER: A user who has read-only access to
	// dashboards.
	//
	//     * AUTHOR: A user who can create data sources, datasets,
	// analyses, and dashboards.
	//
	//     * ADMIN: A user who is an author, who can also
	// manage Amazon QuickSight settings.
	//
	// The name of the QuickSight role is invisible
	// to the user except for the console screens dealing with permissions.
	//
	// This member is required.
	Role types.UserRole

	// The Amazon QuickSight user name that you want to update.
	//
	// This member is required.
	UserName *string

	// (Enterprise edition only) The name of the custom permissions profile that you
	// want to assign to this user. Customized permissions allows you to control a
	// user's access by restricting access the following operations:
	//
	//     * Create and
	// update data sources
	//
	//     * Create and update datasets
	//
	//     * Create and update
	// email reports
	//
	//     * Subscribe to email reports
	//
	// A set of custom permissions
	// includes any combination of these restrictions. Currently, you need to create
	// the profile names for custom permission sets by using the QuickSight console.
	// Then, you use the RegisterUser API operation to assign the named set of
	// permissions to a QuickSight user. QuickSight custom permissions are applied
	// through IAM policies. Therefore, they override the permissions typically granted
	// by assigning QuickSight users to one of the default security cohorts in
	// QuickSight (admin, author, reader). This feature is available only to QuickSight
	// Enterprise edition subscriptions that use SAML 2.0-Based Federation for Single
	// Sign-On (SSO).
	CustomPermissionsName *string

	// A flag that you use to indicate that you want to remove all custom permissions
	// from this user. Using this parameter resets the user to the state it was in
	// before a custom permissions profile was applied. This parameter defaults to NULL
	// and it doesn't accept any other value.
	UnapplyCustomPermissions *bool
}

type UpdateUserOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// The Amazon QuickSight user.
	User *types.User

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUser{}, middleware.After)
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
	addOpUpdateUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateUser",
	}
}
