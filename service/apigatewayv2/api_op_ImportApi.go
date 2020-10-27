// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Imports an API.
func (c *Client) ImportApi(ctx context.Context, params *ImportApiInput, optFns ...func(*Options)) (*ImportApiOutput, error) {
	if params == nil {
		params = &ImportApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportApi", params, optFns, addOperationImportApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ImportApiInput struct {

	// The OpenAPI definition. Supported only for HTTP APIs.
	//
	// This member is required.
	Body *string

	// Specifies how to interpret the base path of the API during import. Valid values
	// are ignore, prepend, and split. The default value is ignore. To learn more, see
	// Set the OpenAPI basePath Property
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html).
	// Supported only for HTTP APIs.
	Basepath *string

	// Specifies whether to rollback the API creation when a warning is encountered. By
	// default, API creation continues if a warning is encountered.
	FailOnWarnings *bool
}

type ImportApiOutput struct {

	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The
	// stage name is typically appended to this URI to form a complete path to a
	// deployed API stage.
	ApiEndpoint *string

	// Specifies whether an API is managed by API Gateway. You can't update or delete a
	// managed API by using API Gateway. A managed API can be deleted only through the
	// tooling or service that created it.
	ApiGatewayManaged *bool

	// The API ID.
	ApiId *string

	// An API key selection expression. Supported only for WebSocket APIs. See API Key
	// Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *types.Cors

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The description of the API.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint *bool

	// Avoid validating models when creating a deployment. Supported only for WebSocket
	// APIs.
	DisableSchemaValidation *bool

	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	ImportInfo []*string

	// The name of the API.
	Name *string

	// The API protocol.
	ProtocolType types.ProtocolType

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string

	// A collection of tags associated with the API.
	Tags map[string]*string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportApi{}, middleware.After)
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
	addOpImportApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opImportApi(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opImportApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ImportApi",
	}
}
