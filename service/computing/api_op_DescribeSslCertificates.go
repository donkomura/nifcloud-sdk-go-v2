// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeSslCertificatesInput struct {
	_ struct{} `type:"structure"`

	Fqdn []string `locationName:"Fqdn" type:"list"`

	FqdnId []string `locationName:"FqdnId" type:"list"`
}

// String returns the string representation
func (s DescribeSslCertificatesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeSslCertificatesOutput struct {
	_ struct{} `type:"structure"`

	CertsSet []CertsSetItem `locationName:"certsSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DescribeSslCertificatesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeSslCertificates = "DescribeSslCertificates"

// DescribeSslCertificatesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeSslCertificatesRequest.
//    req := client.DescribeSslCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeSslCertificates
func (c *Client) DescribeSslCertificatesRequest(input *DescribeSslCertificatesInput) DescribeSslCertificatesRequest {
	op := &aws.Operation{
		Name:       opDescribeSslCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeSslCertificatesInput{}
	}

	req := c.newRequest(op, input, &DescribeSslCertificatesOutput{})
	return DescribeSslCertificatesRequest{Request: req, Input: input, Copy: c.DescribeSslCertificatesRequest}
}

// DescribeSslCertificatesRequest is the request type for the
// DescribeSslCertificates API operation.
type DescribeSslCertificatesRequest struct {
	*aws.Request
	Input *DescribeSslCertificatesInput
	Copy  func(*DescribeSslCertificatesInput) DescribeSslCertificatesRequest
}

// Send marshals and sends the DescribeSslCertificates API request.
func (r DescribeSslCertificatesRequest) Send(ctx context.Context) (*DescribeSslCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSslCertificatesResponse{
		DescribeSslCertificatesOutput: r.Request.Data.(*DescribeSslCertificatesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSslCertificatesResponse is the response type for the
// DescribeSslCertificates API operation.
type DescribeSslCertificatesResponse struct {
	*DescribeSslCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSslCertificates request.
func (r *DescribeSslCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
