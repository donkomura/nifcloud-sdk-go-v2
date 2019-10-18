// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyReplaceNatRuleInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"Description" type:"string"`

	Destination *RequestDestinationStruct `locationName:"Destination" type:"structure"`

	InboundInterface *RequestInboundInterfaceStruct `locationName:"InboundInterface" type:"structure"`

	NatTableId *string `locationName:"NatTableId" type:"string"`

	NatType *string `locationName:"NatType" type:"string"`

	OutboundInterface *RequestOutboundInterfaceStruct `locationName:"OutboundInterface" type:"structure"`

	Protocol *string `locationName:"Protocol" type:"string"`

	RuleNumber *string `locationName:"RuleNumber" type:"string"`

	Source *RequestSourceStruct `locationName:"Source" type:"structure"`

	Translation *RequestTranslationStruct `locationName:"Translation" type:"structure"`
}

// String returns the string representation
func (s NiftyReplaceNatRuleInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyReplaceNatRuleOutput struct {
	_ struct{} `type:"structure"`

	NatRule *NatRule `locationName:"natRule" type:"structure"`

	NatTableId *string `locationName:"natTableId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyReplaceNatRuleOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyReplaceNatRule = "NiftyReplaceNatRule"

// NiftyReplaceNatRuleRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyReplaceNatRuleRequest.
//    req := client.NiftyReplaceNatRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyReplaceNatRule
func (c *Client) NiftyReplaceNatRuleRequest(input *NiftyReplaceNatRuleInput) NiftyReplaceNatRuleRequest {
	op := &aws.Operation{
		Name:       opNiftyReplaceNatRule,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyReplaceNatRuleInput{}
	}

	req := c.newRequest(op, input, &NiftyReplaceNatRuleOutput{})
	return NiftyReplaceNatRuleRequest{Request: req, Input: input, Copy: c.NiftyReplaceNatRuleRequest}
}

// NiftyReplaceNatRuleRequest is the request type for the
// NiftyReplaceNatRule API operation.
type NiftyReplaceNatRuleRequest struct {
	*aws.Request
	Input *NiftyReplaceNatRuleInput
	Copy  func(*NiftyReplaceNatRuleInput) NiftyReplaceNatRuleRequest
}

// Send marshals and sends the NiftyReplaceNatRule API request.
func (r NiftyReplaceNatRuleRequest) Send(ctx context.Context) (*NiftyReplaceNatRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyReplaceNatRuleResponse{
		NiftyReplaceNatRuleOutput: r.Request.Data.(*NiftyReplaceNatRuleOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyReplaceNatRuleResponse is the response type for the
// NiftyReplaceNatRule API operation.
type NiftyReplaceNatRuleResponse struct {
	*NiftyReplaceNatRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyReplaceNatRule request.
func (r *NiftyReplaceNatRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
