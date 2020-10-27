// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the user from signing in with the specified external (SAML or social)
// identity provider. If the user to disable is a Cognito User Pools native
// username + password user, they are not permitted to use their password to
// sign-in. If the user to disable is a linked external IdP user, any link between
// that user and an existing user is removed. The next time the external user (no
// longer attached to the previously linked DestinationUser) signs in, they must
// create a new user account. See AdminLinkProviderForUser
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminLinkProviderForUser.html).
// This action is enabled only for admin access and requires developer credentials.
// The ProviderName must match the value specified when creating an IdP for the
// pool. To disable a native username + password user, the ProviderName value must
// be Cognito and the ProviderAttributeName must be Cognito_Subject, with the
// ProviderAttributeValue being the name that is used in the user pool for the
// user. The ProviderAttributeName must always be Cognito_Subject for social
// identity providers. The ProviderAttributeValue must always be the exact subject
// that was used when the user was originally linked as a source user. For
// de-linking a SAML identity, there are two scenarios. If the linked identity has
// not yet been used to sign-in, the ProviderAttributeName and
// ProviderAttributeValue must be the same values that were used for the SourceUser
// when the identities were originally linked using  AdminLinkProviderForUser call.
// (If the linking was done with ProviderAttributeName set to Cognito_Subject, the
// same applies here). However, if the user has already signed in, the
// ProviderAttributeName must be Cognito_Subject and ProviderAttributeValue must be
// the subject of the SAML assertion.
func (c *Client) AdminDisableProviderForUser(ctx context.Context, params *AdminDisableProviderForUserInput, optFns ...func(*Options)) (*AdminDisableProviderForUserOutput, error) {
	if params == nil {
		params = &AdminDisableProviderForUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminDisableProviderForUser", params, optFns, addOperationAdminDisableProviderForUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminDisableProviderForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminDisableProviderForUserInput struct {

	// The user to be disabled.
	//
	// This member is required.
	User *types.ProviderUserIdentifierType

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string
}

type AdminDisableProviderForUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminDisableProviderForUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminDisableProviderForUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminDisableProviderForUser{}, middleware.After)
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
	addOpAdminDisableProviderForUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminDisableProviderForUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAdminDisableProviderForUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminDisableProviderForUser",
	}
}
