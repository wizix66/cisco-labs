// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DeleteVpcEndpoints.
type DeleteVpcEndpointsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more VPC endpoint IDs.
	//
	// VpcEndpointIds is a required field
	VpcEndpointIds []string `locationName:"VpcEndpointId" locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteVpcEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVpcEndpointsInput"}

	if s.VpcEndpointIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcEndpointIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DeleteVpcEndpoints.
type DeleteVpcEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the VPC endpoints that were not successfully deleted.
	Unsuccessful []UnsuccessfulItem `locationName:"unsuccessful" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeleteVpcEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVpcEndpoints = "DeleteVpcEndpoints"

// DeleteVpcEndpointsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes one or more specified VPC endpoints. Deleting a gateway endpoint
// also deletes the endpoint routes in the route tables that were associated
// with the endpoint. Deleting an interface endpoint deletes the endpoint network
// interfaces.
//
//    // Example sending a request using DeleteVpcEndpointsRequest.
//    req := client.DeleteVpcEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteVpcEndpoints
func (c *Client) DeleteVpcEndpointsRequest(input *DeleteVpcEndpointsInput) DeleteVpcEndpointsRequest {
	op := &aws.Operation{
		Name:       opDeleteVpcEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcEndpointsInput{}
	}

	req := c.newRequest(op, input, &DeleteVpcEndpointsOutput{})

	return DeleteVpcEndpointsRequest{Request: req, Input: input, Copy: c.DeleteVpcEndpointsRequest}
}

// DeleteVpcEndpointsRequest is the request type for the
// DeleteVpcEndpoints API operation.
type DeleteVpcEndpointsRequest struct {
	*aws.Request
	Input *DeleteVpcEndpointsInput
	Copy  func(*DeleteVpcEndpointsInput) DeleteVpcEndpointsRequest
}

// Send marshals and sends the DeleteVpcEndpoints API request.
func (r DeleteVpcEndpointsRequest) Send(ctx context.Context) (*DeleteVpcEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVpcEndpointsResponse{
		DeleteVpcEndpointsOutput: r.Request.Data.(*DeleteVpcEndpointsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVpcEndpointsResponse is the response type for the
// DeleteVpcEndpoints API operation.
type DeleteVpcEndpointsResponse struct {
	*DeleteVpcEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVpcEndpoints request.
func (r *DeleteVpcEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
