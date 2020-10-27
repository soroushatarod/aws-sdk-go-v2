// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new custom language model. Use Amazon S3 prefixes to provide the
// location of your input files. The time it takes to create your model depends on
// the size of your training data.
func (c *Client) CreateLanguageModel(ctx context.Context, params *CreateLanguageModelInput, optFns ...func(*Options)) (*CreateLanguageModelOutput, error) {
	if params == nil {
		params = &CreateLanguageModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLanguageModel", params, optFns, addOperationCreateLanguageModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLanguageModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLanguageModelInput struct {

	// The Amazon Transcribe standard language model, or base model used to create your
	// custom language model. If you want to use your custom language model to
	// transcribe audio with a sample rate of 16 kHz or greater, choose Wideband. If
	// you want to use your custom language model to transcribe audio with a sample
	// rate that is less than 16 kHz, choose Narrowband.
	//
	// This member is required.
	BaseModelName types.BaseModelName

	// Contains the data access role and the Amazon S3 prefixes to read the required
	// input files to create a custom language model.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input text you're using to train your custom language model.
	//
	// This member is required.
	LanguageCode types.CLMLanguageCode

	// The name you choose for your custom language model when you create it.
	//
	// This member is required.
	ModelName *string
}

type CreateLanguageModelOutput struct {

	// The Amazon Transcribe standard language model, or base model you've used to
	// create a custom language model.
	BaseModelName types.BaseModelName

	// The data access role and Amazon S3 prefixes you've chosen to create your custom
	// language model.
	InputDataConfig *types.InputDataConfig

	// The language code of the text you've used to create a custom language model.
	LanguageCode types.CLMLanguageCode

	// The name you've chosen for your custom language model.
	ModelName *string

	// The status of the custom language model. When the status is COMPLETED the model
	// is ready to use.
	ModelStatus types.ModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLanguageModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLanguageModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLanguageModel{}, middleware.After)
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
	addOpCreateLanguageModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLanguageModel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateLanguageModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateLanguageModel",
	}
}
