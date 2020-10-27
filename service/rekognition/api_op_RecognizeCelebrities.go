// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of celebrities recognized in the input image. For more
// information, see Recognizing Celebrities in the Amazon Rekognition Developer
// Guide. RecognizeCelebrities returns the 64 largest faces in the image. It lists
// recognized celebrities in the CelebrityFaces array and unrecognized faces in the
// UnrecognizedFaces array. RecognizeCelebrities doesn't return celebrities whose
// faces aren't among the largest 64 faces in the image. For each celebrity
// recognized, RecognizeCelebrities returns a Celebrity object. The Celebrity
// object contains the celebrity name, ID, URL links to additional information,
// match confidence, and a ComparedFace object that you can use to locate the
// celebrity's face on the image. Amazon Rekognition doesn't retain information
// about which images a celebrity has been recognized in. Your application must
// store this information and use the Celebrity ID property as a unique identifier
// for the celebrity. If you don't store the celebrity name or additional
// information URLs returned by RecognizeCelebrities, you will need the ID to
// identify the celebrity in a call to the GetCelebrityInfo operation. You pass the
// input image either as base64-encoded image bytes or as a reference to an image
// in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, passing image bytes is not supported. The image must be either a PNG
// or JPEG formatted file. For an example, see Recognizing Celebrities in an Image
// in the Amazon Rekognition Developer Guide. This operation requires permissions
// to perform the rekognition:RecognizeCelebrities operation.
func (c *Client) RecognizeCelebrities(ctx context.Context, params *RecognizeCelebritiesInput, optFns ...func(*Options)) (*RecognizeCelebritiesOutput, error) {
	if params == nil {
		params = &RecognizeCelebritiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecognizeCelebrities", params, optFns, addOperationRecognizeCelebritiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RecognizeCelebritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecognizeCelebritiesInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image
}

type RecognizeCelebritiesOutput struct {

	// Details about each celebrity found in the image. Amazon Rekognition can detect a
	// maximum of 64 celebrities in an image.
	CelebrityFaces []*types.Celebrity

	// The orientation of the input image (counterclockwise direction). If your
	// application displays the image, you can use this value to correct the
	// orientation. The bounding box coordinates returned in CelebrityFaces and
	// UnrecognizedFaces represent face locations before the image orientation is
	// corrected. If the input image is in .jpeg format, it might contain exchangeable
	// image (Exif) metadata that includes the image's orientation. If so, and the Exif
	// metadata for the input image populates the orientation field, the value of
	// OrientationCorrection is null. The CelebrityFaces and UnrecognizedFaces bounding
	// box coordinates represent face locations after Exif metadata is used to correct
	// the image orientation. Images in .png format don't contain Exif metadata.
	OrientationCorrection types.OrientationCorrection

	// Details about each unrecognized face in the image.
	UnrecognizedFaces []*types.ComparedFace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRecognizeCelebritiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRecognizeCelebrities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRecognizeCelebrities{}, middleware.After)
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
	addOpRecognizeCelebritiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRecognizeCelebrities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRecognizeCelebrities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "RecognizeCelebrities",
	}
}
