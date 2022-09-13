// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CopyImage.
type CopyImageInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	ClientToken *string `type:"string"`

	// A description for the new AMI in the destination Region.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Specifies whether the destination snapshots of the copied image should be
	// encrypted. You can encrypt a copy of an unencrypted snapshot, but you cannot
	// create an unencrypted copy of an encrypted snapshot. The default CMK for
	// EBS is used unless you specify a non-default AWS Key Management Service (AWS
	// KMS) CMK using KmsKeyId. For more information, see Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// An identifier for the symmetric AWS Key Management Service (AWS KMS) customer
	// master key (CMK) to use when creating the encrypted volume. This parameter
	// is only required if you want to use a non-default CMK; if this parameter
	// is not specified, the default CMK for EBS is used. If a KmsKeyId is specified,
	// the Encrypted flag must also be set.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// AWS parses KmsKeyId asynchronously, meaning that the action you call may
	// appear to complete even though you provided an invalid identifier. This action
	// will eventually report failure.
	//
	// The specified CMK must exist in the Region that the snapshot is being copied
	// to.
	//
	// Amazon EBS does not support asymmetric CMKs.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The name of the new AMI in the destination Region.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The ID of the AMI to copy.
	//
	// SourceImageId is a required field
	SourceImageId *string `type:"string" required:"true"`

	// The name of the Region that contains the AMI to copy.
	//
	// SourceRegion is a required field
	SourceRegion *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyImageInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SourceImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceImageId"))
	}

	if s.SourceRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceRegion"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CopyImage.
type CopyImageOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the new AMI.
	ImageId *string `locationName:"imageId" type:"string"`
}

// String returns the string representation
func (s CopyImageOutput) String() string {
	return awsutil.Prettify(s)
}

const opCopyImage = "CopyImage"

// CopyImageRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Initiates the copy of an AMI from the specified source Region to the current
// Region. You specify the destination Region by using its endpoint when making
// the request.
//
// Copies of encrypted backing snapshots for the AMI are encrypted. Copies of
// unencrypted backing snapshots remain unencrypted, unless you set Encrypted
// during the copy operation. You cannot create an unencrypted copy of an encrypted
// backing snapshot.
//
// For more information about the prerequisites and limits when copying an AMI,
// see Copying an AMI (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/CopyingAMIs.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CopyImageRequest.
//    req := client.CopyImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CopyImage
func (c *Client) CopyImageRequest(input *CopyImageInput) CopyImageRequest {
	op := &aws.Operation{
		Name:       opCopyImage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyImageInput{}
	}

	req := c.newRequest(op, input, &CopyImageOutput{})

	return CopyImageRequest{Request: req, Input: input, Copy: c.CopyImageRequest}
}

// CopyImageRequest is the request type for the
// CopyImage API operation.
type CopyImageRequest struct {
	*aws.Request
	Input *CopyImageInput
	Copy  func(*CopyImageInput) CopyImageRequest
}

// Send marshals and sends the CopyImage API request.
func (r CopyImageRequest) Send(ctx context.Context) (*CopyImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyImageResponse{
		CopyImageOutput: r.Request.Data.(*CopyImageOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyImageResponse is the response type for the
// CopyImage API operation.
type CopyImageResponse struct {
	*CopyImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyImage request.
func (r *CopyImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
