// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyCreatePrivateLanInput struct {
	_ struct{} `type:"structure"`

	AccountingType *string `locationName:"AccountingType" type:"string"`

	AvailabilityZone *string `locationName:"AvailabilityZone" type:"string"`

	CidrBlock *string `locationName:"CidrBlock" type:"string"`

	Description *string `locationName:"Description" type:"string"`

	PrivateLanName *string `locationName:"PrivateLanName" type:"string"`
}

// String returns the string representation
func (s NiftyCreatePrivateLanInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyCreatePrivateLanOutput struct {
	_ struct{} `type:"structure"`

	PrivateLan *PrivateLan `locationName:"privateLan" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyCreatePrivateLanOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyCreatePrivateLan = "NiftyCreatePrivateLan"

// NiftyCreatePrivateLanRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyCreatePrivateLanRequest.
//    req := client.NiftyCreatePrivateLanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyCreatePrivateLan
func (c *Client) NiftyCreatePrivateLanRequest(input *NiftyCreatePrivateLanInput) NiftyCreatePrivateLanRequest {
	op := &aws.Operation{
		Name:       opNiftyCreatePrivateLan,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyCreatePrivateLanInput{}
	}

	req := c.newRequest(op, input, &NiftyCreatePrivateLanOutput{})
	return NiftyCreatePrivateLanRequest{Request: req, Input: input, Copy: c.NiftyCreatePrivateLanRequest}
}

// NiftyCreatePrivateLanRequest is the request type for the
// NiftyCreatePrivateLan API operation.
type NiftyCreatePrivateLanRequest struct {
	*aws.Request
	Input *NiftyCreatePrivateLanInput
	Copy  func(*NiftyCreatePrivateLanInput) NiftyCreatePrivateLanRequest
}

// Send marshals and sends the NiftyCreatePrivateLan API request.
func (r NiftyCreatePrivateLanRequest) Send(ctx context.Context) (*NiftyCreatePrivateLanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyCreatePrivateLanResponse{
		NiftyCreatePrivateLanOutput: r.Request.Data.(*NiftyCreatePrivateLanOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyCreatePrivateLanResponse is the response type for the
// NiftyCreatePrivateLan API operation.
type NiftyCreatePrivateLanResponse struct {
	*NiftyCreatePrivateLanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyCreatePrivateLan request.
func (r *NiftyCreatePrivateLanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}