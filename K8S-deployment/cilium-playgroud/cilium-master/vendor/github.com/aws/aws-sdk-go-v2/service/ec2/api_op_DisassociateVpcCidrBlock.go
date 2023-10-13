// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateVpcCidrBlockInput struct {
	_ struct{} `type:"structure"`

	// The association ID for the CIDR block.
	//
	// AssociationId is a required field
	AssociationId *string `locationName:"associationId" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateVpcCidrBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateVpcCidrBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateVpcCidrBlockInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateVpcCidrBlockOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IPv4 CIDR block association.
	CidrBlockAssociation *VpcCidrBlockAssociation `locationName:"cidrBlockAssociation" type:"structure"`

	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *VpcIpv6CidrBlockAssociation `locationName:"ipv6CidrBlockAssociation" type:"structure"`

	// The ID of the VPC.
	VpcId *string `locationName:"vpcId" type:"string"`
}

// String returns the string representation
func (s DisassociateVpcCidrBlockOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateVpcCidrBlock = "DisassociateVpcCidrBlock"

// DisassociateVpcCidrBlockRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates a CIDR block from a VPC. To disassociate the CIDR block, you
// must specify its association ID. You can get the association ID by using
// DescribeVpcs. You must detach or delete all gateways and resources that are
// associated with the CIDR block before you can disassociate it.
//
// You cannot disassociate the CIDR block with which you originally created
// the VPC (the primary CIDR block).
//
//    // Example sending a request using DisassociateVpcCidrBlockRequest.
//    req := client.DisassociateVpcCidrBlockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateVpcCidrBlock
func (c *Client) DisassociateVpcCidrBlockRequest(input *DisassociateVpcCidrBlockInput) DisassociateVpcCidrBlockRequest {
	op := &aws.Operation{
		Name:       opDisassociateVpcCidrBlock,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateVpcCidrBlockInput{}
	}

	req := c.newRequest(op, input, &DisassociateVpcCidrBlockOutput{})

	return DisassociateVpcCidrBlockRequest{Request: req, Input: input, Copy: c.DisassociateVpcCidrBlockRequest}
}

// DisassociateVpcCidrBlockRequest is the request type for the
// DisassociateVpcCidrBlock API operation.
type DisassociateVpcCidrBlockRequest struct {
	*aws.Request
	Input *DisassociateVpcCidrBlockInput
	Copy  func(*DisassociateVpcCidrBlockInput) DisassociateVpcCidrBlockRequest
}

// Send marshals and sends the DisassociateVpcCidrBlock API request.
func (r DisassociateVpcCidrBlockRequest) Send(ctx context.Context) (*DisassociateVpcCidrBlockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateVpcCidrBlockResponse{
		DisassociateVpcCidrBlockOutput: r.Request.Data.(*DisassociateVpcCidrBlockOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateVpcCidrBlockResponse is the response type for the
// DisassociateVpcCidrBlock API operation.
type DisassociateVpcCidrBlockResponse struct {
	*DisassociateVpcCidrBlockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateVpcCidrBlock request.
func (r *DisassociateVpcCidrBlockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
