// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeleteNetworkAclInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the network ACL.
	//
	// NetworkAclId is a required field
	NetworkAclId *string `locationName:"networkAclId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkAclInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkAclInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNetworkAclInput"}

	if s.NetworkAclId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNetworkAclOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNetworkAclOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNetworkAcl = "DeleteNetworkAcl"

// DeleteNetworkAclRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified network ACL. You can't delete the ACL if it's associated
// with any subnets. You can't delete the default network ACL.
//
//    // Example sending a request using DeleteNetworkAclRequest.
//    req := client.DeleteNetworkAclRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteNetworkAcl
func (c *Client) DeleteNetworkAclRequest(input *DeleteNetworkAclInput) DeleteNetworkAclRequest {
	op := &aws.Operation{
		Name:       opDeleteNetworkAcl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNetworkAclInput{}
	}

	req := c.newRequest(op, input, &DeleteNetworkAclOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteNetworkAclRequest{Request: req, Input: input, Copy: c.DeleteNetworkAclRequest}
}

// DeleteNetworkAclRequest is the request type for the
// DeleteNetworkAcl API operation.
type DeleteNetworkAclRequest struct {
	*aws.Request
	Input *DeleteNetworkAclInput
	Copy  func(*DeleteNetworkAclInput) DeleteNetworkAclRequest
}

// Send marshals and sends the DeleteNetworkAcl API request.
func (r DeleteNetworkAclRequest) Send(ctx context.Context) (*DeleteNetworkAclResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNetworkAclResponse{
		DeleteNetworkAclOutput: r.Request.Data.(*DeleteNetworkAclOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNetworkAclResponse is the response type for the
// DeleteNetworkAcl API operation.
type DeleteNetworkAclResponse struct {
	*DeleteNetworkAclOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNetworkAcl request.
func (r *DeleteNetworkAclResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
