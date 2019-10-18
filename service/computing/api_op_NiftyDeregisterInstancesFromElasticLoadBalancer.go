// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDeregisterInstancesFromElasticLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	ElasticLoadBalancerId *string `locationName:"ElasticLoadBalancerId" type:"string"`

	ElasticLoadBalancerName *string `locationName:"ElasticLoadBalancerName" type:"string"`

	ElasticLoadBalancerPort *int64 `locationName:"ElasticLoadBalancerPort" type:"integer"`

	InstancePort *int64 `locationName:"InstancePort" type:"integer"`

	Instances []RequestInstancesStruct `locationName:"Instances" locationNameList:"member" type:"list"`

	Protocol *string `locationName:"Protocol" type:"string"`
}

// String returns the string representation
func (s NiftyDeregisterInstancesFromElasticLoadBalancerInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeregisterInstancesFromElasticLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s NiftyDeregisterInstancesFromElasticLoadBalancerOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeregisterInstancesFromElasticLoadBalancer = "NiftyDeregisterInstancesFromElasticLoadBalancer"

// NiftyDeregisterInstancesFromElasticLoadBalancerRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeregisterInstancesFromElasticLoadBalancerRequest.
//    req := client.NiftyDeregisterInstancesFromElasticLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDeregisterInstancesFromElasticLoadBalancer
func (c *Client) NiftyDeregisterInstancesFromElasticLoadBalancerRequest(input *NiftyDeregisterInstancesFromElasticLoadBalancerInput) NiftyDeregisterInstancesFromElasticLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opNiftyDeregisterInstancesFromElasticLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeregisterInstancesFromElasticLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &NiftyDeregisterInstancesFromElasticLoadBalancerOutput{})
	return NiftyDeregisterInstancesFromElasticLoadBalancerRequest{Request: req, Input: input, Copy: c.NiftyDeregisterInstancesFromElasticLoadBalancerRequest}
}

// NiftyDeregisterInstancesFromElasticLoadBalancerRequest is the request type for the
// NiftyDeregisterInstancesFromElasticLoadBalancer API operation.
type NiftyDeregisterInstancesFromElasticLoadBalancerRequest struct {
	*aws.Request
	Input *NiftyDeregisterInstancesFromElasticLoadBalancerInput
	Copy  func(*NiftyDeregisterInstancesFromElasticLoadBalancerInput) NiftyDeregisterInstancesFromElasticLoadBalancerRequest
}

// Send marshals and sends the NiftyDeregisterInstancesFromElasticLoadBalancer API request.
func (r NiftyDeregisterInstancesFromElasticLoadBalancerRequest) Send(ctx context.Context) (*NiftyDeregisterInstancesFromElasticLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeregisterInstancesFromElasticLoadBalancerResponse{
		NiftyDeregisterInstancesFromElasticLoadBalancerOutput: r.Request.Data.(*NiftyDeregisterInstancesFromElasticLoadBalancerOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeregisterInstancesFromElasticLoadBalancerResponse is the response type for the
// NiftyDeregisterInstancesFromElasticLoadBalancer API operation.
type NiftyDeregisterInstancesFromElasticLoadBalancerResponse struct {
	*NiftyDeregisterInstancesFromElasticLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeregisterInstancesFromElasticLoadBalancer request.
func (r *NiftyDeregisterInstancesFromElasticLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}