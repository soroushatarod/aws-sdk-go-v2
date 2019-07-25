// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Provides options for deleting a vault from Amazon S3 Glacier.
type DeleteVaultInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVaultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVaultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVaultInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVaultInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVaultOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVaultOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVaultOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVault = "DeleteVault"

// DeleteVaultRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation deletes a vault. Amazon S3 Glacier will delete a vault only
// if there are no archives in the vault as of the last inventory and there
// have been no writes to the vault since the last inventory. If either of these
// conditions is not satisfied, the vault deletion fails (that is, the vault
// is not removed) and Amazon S3 Glacier returns an error. You can use DescribeVault
// to return the number of archives in a vault, and you can use Initiate a Job
// (POST jobs) (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html)
// to initiate a new inventory retrieval for a vault. The inventory contains
// the archive IDs you use to delete archives using Delete Archive (DELETE archive)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-delete.html).
//
// This operation is idempotent.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Deleting a Vault
// in Amazon Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/deleting-vaults.html)
// and Delete Vault (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-delete.html)
// in the Amazon S3 Glacier Developer Guide.
//
//    // Example sending a request using DeleteVaultRequest.
//    req := client.DeleteVaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteVaultRequest(input *DeleteVaultInput) DeleteVaultRequest {
	op := &aws.Operation{
		Name:       opDeleteVault,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{accountId}/vaults/{vaultName}",
	}

	if input == nil {
		input = &DeleteVaultInput{}
	}

	req := c.newRequest(op, input, &DeleteVaultOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVaultRequest{Request: req, Input: input, Copy: c.DeleteVaultRequest}
}

// DeleteVaultRequest is the request type for the
// DeleteVault API operation.
type DeleteVaultRequest struct {
	*aws.Request
	Input *DeleteVaultInput
	Copy  func(*DeleteVaultInput) DeleteVaultRequest
}

// Send marshals and sends the DeleteVault API request.
func (r DeleteVaultRequest) Send(ctx context.Context) (*DeleteVaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVaultResponse{
		DeleteVaultOutput: r.Request.Data.(*DeleteVaultOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVaultResponse is the response type for the
// DeleteVault API operation.
type DeleteVaultResponse struct {
	*DeleteVaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVault request.
func (r *DeleteVaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}