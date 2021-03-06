// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDeregisterInstancesFromSeparateInstanceRuleInput struct {
	_ struct{} `type:"structure"`

	InstanceId []string `locationName:"InstanceId" type:"list"`

	InstanceUniqueId []string `locationName:"InstanceUniqueId" type:"list"`

	SeparateInstanceRuleName *string `locationName:"SeparateInstanceRuleName" type:"string"`
}

// String returns the string representation
func (s NiftyDeregisterInstancesFromSeparateInstanceRuleInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeregisterInstancesFromSeparateInstanceRuleOutput struct {
	_ struct{} `type:"structure"`

	InstancesSet []InstancesSetItem `locationName:"instancesSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyDeregisterInstancesFromSeparateInstanceRuleOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeregisterInstancesFromSeparateInstanceRule = "NiftyDeregisterInstancesFromSeparateInstanceRule"

// NiftyDeregisterInstancesFromSeparateInstanceRuleRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeregisterInstancesFromSeparateInstanceRuleRequest.
//    req := client.NiftyDeregisterInstancesFromSeparateInstanceRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDeregisterInstancesFromSeparateInstanceRule
func (c *Client) NiftyDeregisterInstancesFromSeparateInstanceRuleRequest(input *NiftyDeregisterInstancesFromSeparateInstanceRuleInput) NiftyDeregisterInstancesFromSeparateInstanceRuleRequest {
	op := &aws.Operation{
		Name:       opNiftyDeregisterInstancesFromSeparateInstanceRule,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeregisterInstancesFromSeparateInstanceRuleInput{}
	}

	req := c.newRequest(op, input, &NiftyDeregisterInstancesFromSeparateInstanceRuleOutput{})
	return NiftyDeregisterInstancesFromSeparateInstanceRuleRequest{Request: req, Input: input, Copy: c.NiftyDeregisterInstancesFromSeparateInstanceRuleRequest}
}

// NiftyDeregisterInstancesFromSeparateInstanceRuleRequest is the request type for the
// NiftyDeregisterInstancesFromSeparateInstanceRule API operation.
type NiftyDeregisterInstancesFromSeparateInstanceRuleRequest struct {
	*aws.Request
	Input *NiftyDeregisterInstancesFromSeparateInstanceRuleInput
	Copy  func(*NiftyDeregisterInstancesFromSeparateInstanceRuleInput) NiftyDeregisterInstancesFromSeparateInstanceRuleRequest
}

// Send marshals and sends the NiftyDeregisterInstancesFromSeparateInstanceRule API request.
func (r NiftyDeregisterInstancesFromSeparateInstanceRuleRequest) Send(ctx context.Context) (*NiftyDeregisterInstancesFromSeparateInstanceRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeregisterInstancesFromSeparateInstanceRuleResponse{
		NiftyDeregisterInstancesFromSeparateInstanceRuleOutput: r.Request.Data.(*NiftyDeregisterInstancesFromSeparateInstanceRuleOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeregisterInstancesFromSeparateInstanceRuleResponse is the response type for the
// NiftyDeregisterInstancesFromSeparateInstanceRule API operation.
type NiftyDeregisterInstancesFromSeparateInstanceRuleResponse struct {
	*NiftyDeregisterInstancesFromSeparateInstanceRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeregisterInstancesFromSeparateInstanceRule request.
func (r *NiftyDeregisterInstancesFromSeparateInstanceRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
