// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeSpotInstanceRequests.
type DescribeSpotInstanceRequestsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * availability-zone-group - The Availability Zone group.
	//
	//    * create-time - The time stamp when the Spot Instance request was created.
	//
	//    * fault-code - The fault code related to the request.
	//
	//    * fault-message - The fault message related to the request.
	//
	//    * instance-id - The ID of the instance that fulfilled the request.
	//
	//    * launch-group - The Spot Instance launch group.
	//
	//    * launch.block-device-mapping.delete-on-termination - Indicates whether
	//    the EBS volume is deleted on instance termination.
	//
	//    * launch.block-device-mapping.device-name - The device name for the volume
	//    in the block device mapping (for example, /dev/sdh or xvdh).
	//
	//    * launch.block-device-mapping.snapshot-id - The ID of the snapshot for
	//    the EBS volume.
	//
	//    * launch.block-device-mapping.volume-size - The size of the EBS volume,
	//    in GiB.
	//
	//    * launch.block-device-mapping.volume-type - The type of EBS volume: gp2
	//    for General Purpose SSD, io1 for Provisioned IOPS SSD, st1 for Throughput
	//    Optimized HDD, sc1for Cold HDD, or standard for Magnetic.
	//
	//    * launch.group-id - The ID of the security group for the instance.
	//
	//    * launch.group-name - The name of the security group for the instance.
	//
	//    * launch.image-id - The ID of the AMI.
	//
	//    * launch.instance-type - The type of instance (for example, m3.medium).
	//
	//    * launch.kernel-id - The kernel ID.
	//
	//    * launch.key-name - The name of the key pair the instance launched with.
	//
	//    * launch.monitoring-enabled - Whether detailed monitoring is enabled for
	//    the Spot Instance.
	//
	//    * launch.ramdisk-id - The RAM disk ID.
	//
	//    * launched-availability-zone - The Availability Zone in which the request
	//    is launched.
	//
	//    * network-interface.addresses.primary - Indicates whether the IP address
	//    is the primary private IP address.
	//
	//    * network-interface.delete-on-termination - Indicates whether the network
	//    interface is deleted when the instance is terminated.
	//
	//    * network-interface.description - A description of the network interface.
	//
	//    * network-interface.device-index - The index of the device for the network
	//    interface attachment on the instance.
	//
	//    * network-interface.group-id - The ID of the security group associated
	//    with the network interface.
	//
	//    * network-interface.network-interface-id - The ID of the network interface.
	//
	//    * network-interface.private-ip-address - The primary private IP address
	//    of the network interface.
	//
	//    * network-interface.subnet-id - The ID of the subnet for the instance.
	//
	//    * product-description - The product description associated with the instance
	//    (Linux/UNIX | Windows).
	//
	//    * spot-instance-request-id - The Spot Instance request ID.
	//
	//    * spot-price - The maximum hourly price for any Spot Instance launched
	//    to fulfill the request.
	//
	//    * state - The state of the Spot Instance request (open | active | closed
	//    | cancelled | failed). Spot request status information can help you track
	//    your Amazon EC2 Spot Instance requests. For more information, see Spot
	//    Request Status (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-bid-status.html)
	//    in the Amazon EC2 User Guide for Linux Instances.
	//
	//    * status-code - The short code describing the most recent evaluation of
	//    your Spot Instance request.
	//
	//    * status-message - The message explaining the status of the Spot Instance
	//    request.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * type - The type of Spot Instance request (one-time | persistent).
	//
	//    * valid-from - The start date of the request.
	//
	//    * valid-until - The end date of the request.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. Specify a value
	// between 5 and 1000. To retrieve the remaining results, make another call
	// with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token to request the next set of results. This value is null when there
	// are no more results to return.
	NextToken *string `type:"string"`

	// One or more Spot Instance request IDs.
	SpotInstanceRequestIds []string `locationName:"SpotInstanceRequestId" locationNameList:"SpotInstanceRequestId" type:"list"`
}

// String returns the string representation
func (s DescribeSpotInstanceRequestsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeSpotInstanceRequests.
type DescribeSpotInstanceRequestsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next set of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// One or more Spot Instance requests.
	SpotInstanceRequests []SpotInstanceRequest `locationName:"spotInstanceRequestSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSpotInstanceRequestsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSpotInstanceRequests = "DescribeSpotInstanceRequests"

// DescribeSpotInstanceRequestsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified Spot Instance requests.
//
// You can use DescribeSpotInstanceRequests to find a running Spot Instance
// by examining the response. If the status of the Spot Instance is fulfilled,
// the instance ID appears in the response and contains the identifier of the
// instance. Alternatively, you can use DescribeInstances (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances)
// with a filter to look for instances where the instance lifecycle is spot.
//
// We recommend that you set MaxResults to a value between 5 and 1000 to limit
// the number of results returned. This paginates the output, which makes the
// list more manageable and returns the results faster. If the list of results
// exceeds your MaxResults value, then that number of results is returned along
// with a NextToken value that can be passed to a subsequent DescribeSpotInstanceRequests
// request to retrieve the remaining results.
//
// Spot Instance requests are deleted four hours after they are canceled and
// their instances are terminated.
//
//    // Example sending a request using DescribeSpotInstanceRequestsRequest.
//    req := client.DescribeSpotInstanceRequestsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSpotInstanceRequests
func (c *Client) DescribeSpotInstanceRequestsRequest(input *DescribeSpotInstanceRequestsInput) DescribeSpotInstanceRequestsRequest {
	op := &aws.Operation{
		Name:       opDescribeSpotInstanceRequests,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeSpotInstanceRequestsInput{}
	}

	req := c.newRequest(op, input, &DescribeSpotInstanceRequestsOutput{})

	return DescribeSpotInstanceRequestsRequest{Request: req, Input: input, Copy: c.DescribeSpotInstanceRequestsRequest}
}

// DescribeSpotInstanceRequestsRequest is the request type for the
// DescribeSpotInstanceRequests API operation.
type DescribeSpotInstanceRequestsRequest struct {
	*aws.Request
	Input *DescribeSpotInstanceRequestsInput
	Copy  func(*DescribeSpotInstanceRequestsInput) DescribeSpotInstanceRequestsRequest
}

// Send marshals and sends the DescribeSpotInstanceRequests API request.
func (r DescribeSpotInstanceRequestsRequest) Send(ctx context.Context) (*DescribeSpotInstanceRequestsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSpotInstanceRequestsResponse{
		DescribeSpotInstanceRequestsOutput: r.Request.Data.(*DescribeSpotInstanceRequestsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeSpotInstanceRequestsRequestPaginator returns a paginator for DescribeSpotInstanceRequests.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeSpotInstanceRequestsRequest(input)
//   p := ec2.NewDescribeSpotInstanceRequestsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeSpotInstanceRequestsPaginator(req DescribeSpotInstanceRequestsRequest) DescribeSpotInstanceRequestsPaginator {
	return DescribeSpotInstanceRequestsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeSpotInstanceRequestsInput
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

// DescribeSpotInstanceRequestsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeSpotInstanceRequestsPaginator struct {
	aws.Pager
}

func (p *DescribeSpotInstanceRequestsPaginator) CurrentPage() *DescribeSpotInstanceRequestsOutput {
	return p.Pager.CurrentPage().(*DescribeSpotInstanceRequestsOutput)
}

// DescribeSpotInstanceRequestsResponse is the response type for the
// DescribeSpotInstanceRequests API operation.
type DescribeSpotInstanceRequestsResponse struct {
	*DescribeSpotInstanceRequestsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSpotInstanceRequests request.
func (r *DescribeSpotInstanceRequestsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
