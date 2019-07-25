// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectChildrenRequest
type ListObjectChildrenInput struct {
	_ struct{} `type:"structure"`

	// Represents the manner and timing in which the successful write or update
	// of an object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The reference that identifies the object for which child objects are being
	// listed.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListObjectChildrenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectChildrenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectChildrenInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectChildrenInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-consistency-level", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectChildrenResponse
type ListObjectChildrenOutput struct {
	_ struct{} `type:"structure"`

	// Children structure, which is a map with key as the LinkName and ObjectIdentifier
	// as the value.
	Children map[string]string `type:"map"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListObjectChildrenOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectChildrenOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Children != nil {
		v := s.Children

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Children", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListObjectChildren = "ListObjectChildren"

// ListObjectChildrenRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns a paginated list of child objects that are associated with a given
// object.
//
//    // Example sending a request using ListObjectChildrenRequest.
//    req := client.ListObjectChildrenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectChildren
func (c *Client) ListObjectChildrenRequest(input *ListObjectChildrenInput) ListObjectChildrenRequest {
	op := &aws.Operation{
		Name:       opListObjectChildren,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/children",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListObjectChildrenInput{}
	}

	req := c.newRequest(op, input, &ListObjectChildrenOutput{})
	return ListObjectChildrenRequest{Request: req, Input: input, Copy: c.ListObjectChildrenRequest}
}

// ListObjectChildrenRequest is the request type for the
// ListObjectChildren API operation.
type ListObjectChildrenRequest struct {
	*aws.Request
	Input *ListObjectChildrenInput
	Copy  func(*ListObjectChildrenInput) ListObjectChildrenRequest
}

// Send marshals and sends the ListObjectChildren API request.
func (r ListObjectChildrenRequest) Send(ctx context.Context) (*ListObjectChildrenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectChildrenResponse{
		ListObjectChildrenOutput: r.Request.Data.(*ListObjectChildrenOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectChildrenRequestPaginator returns a paginator for ListObjectChildren.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectChildrenRequest(input)
//   p := clouddirectory.NewListObjectChildrenRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectChildrenPaginator(req ListObjectChildrenRequest) ListObjectChildrenPaginator {
	return ListObjectChildrenPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListObjectChildrenInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListObjectChildrenPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectChildrenPaginator struct {
	aws.Pager
}

func (p *ListObjectChildrenPaginator) CurrentPage() *ListObjectChildrenOutput {
	return p.Pager.CurrentPage().(*ListObjectChildrenOutput)
}

// ListObjectChildrenResponse is the response type for the
// ListObjectChildren API operation.
type ListObjectChildrenResponse struct {
	*ListObjectChildrenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectChildren request.
func (r *ListObjectChildrenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}