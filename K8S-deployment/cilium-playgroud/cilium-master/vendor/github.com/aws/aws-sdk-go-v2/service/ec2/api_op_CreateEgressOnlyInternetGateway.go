// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateEgressOnlyInternetGatewayInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the VPC for which to create the egress-only internet gateway.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateEgressOnlyInternetGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEgressOnlyInternetGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEgressOnlyInternetGatewayInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateEgressOnlyInternetGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the egress-only internet gateway.
	EgressOnlyInternetGateway *EgressOnlyInternetGateway `locationName:"egressOnlyInternetGateway" type:"structure"`
}

// String returns the string representation
func (s CreateEgressOnlyInternetGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateEgressOnlyInternetGateway = "CreateEgressOnlyInternetGateway"

// CreateEgressOnlyInternetGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// [IPv6 only] Creates an egress-only internet gateway for your VPC. An egress-only
// internet gateway is used to enable outbound communication over IPv6 from
// instances in your VPC to the internet, and prevents hosts outside of your
// VPC from initiating an IPv6 connection with your instance.
//
//    // Example sending a request using CreateEgressOnlyInternetGatewayRequest.
//    req := client.CreateEgressOnlyInternetGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateEgressOnlyInternetGateway
func (c *Client) CreateEgressOnlyInternetGatewayRequest(input *CreateEgressOnlyInternetGatewayInput) CreateEgressOnlyInternetGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateEgressOnlyInternetGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEgressOnlyInternetGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateEgressOnlyInternetGatewayOutput{})

	return CreateEgressOnlyInternetGatewayRequest{Request: req, Input: input, Copy: c.CreateEgressOnlyInternetGatewayRequest}
}

// CreateEgressOnlyInternetGatewayRequest is the request type for the
// CreateEgressOnlyInternetGateway API operation.
type CreateEgressOnlyInternetGatewayRequest struct {
	*aws.Request
	Input *CreateEgressOnlyInternetGatewayInput
	Copy  func(*CreateEgressOnlyInternetGatewayInput) CreateEgressOnlyInternetGatewayRequest
}

// Send marshals and sends the CreateEgressOnlyInternetGateway API request.
func (r CreateEgressOnlyInternetGatewayRequest) Send(ctx context.Context) (*CreateEgressOnlyInternetGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEgressOnlyInternetGatewayResponse{
		CreateEgressOnlyInternetGatewayOutput: r.Request.Data.(*CreateEgressOnlyInternetGatewayOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEgressOnlyInternetGatewayResponse is the response type for the
// CreateEgressOnlyInternetGateway API operation.
type CreateEgressOnlyInternetGatewayResponse struct {
	*CreateEgressOnlyInternetGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEgressOnlyInternetGateway request.
func (r *CreateEgressOnlyInternetGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
