// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type RegisterCorporateInfoForCertificateInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	AlphabetName1 *string `locationName:"AlphabetName1" type:"string"`

	AlphabetName2 *string `locationName:"AlphabetName2" type:"string"`

	City *string `locationName:"City" type:"string"`

	CorpGrade *string `locationName:"CorpGrade" type:"string"`

	CorpName *string `locationName:"CorpName" type:"string"`

	DivisionName *string `locationName:"DivisionName" type:"string"`

	EmailAddress *string `locationName:"EmailAddress" type:"string"`

	KanaName1 *string `locationName:"KanaName1" type:"string"`

	KanaName2 *string `locationName:"KanaName2" type:"string"`

	Name1 *string `locationName:"Name1" type:"string"`

	Name2 *string `locationName:"Name2" type:"string"`

	PhoneNumber *string `locationName:"PhoneNumber" type:"string"`

	PostName *string `locationName:"PostName" type:"string"`

	Pref *string `locationName:"Pref" type:"string"`

	PresidentName1 *string `locationName:"PresidentName1" type:"string"`

	PresidentName2 *string `locationName:"PresidentName2" type:"string"`

	TdbCode *string `locationName:"TdbCode" type:"string"`

	Zip1 *string `locationName:"Zip1" type:"string"`

	Zip2 *string `locationName:"Zip2" type:"string"`
}

// String returns the string representation
func (s RegisterCorporateInfoForCertificateInput) String() string {
	return nifcloudutil.Prettify(s)
}

type RegisterCorporateInfoForCertificateOutput struct {
	_ struct{} `type:"structure"`

	City *string `locationName:"city" type:"string"`

	CorpGrade *string `locationName:"corpGrade" type:"string"`

	CorpName *string `locationName:"corpName" type:"string"`

	DivisionName *string `locationName:"divisionName" type:"string"`

	KanaName1 *string `locationName:"kanaName1" type:"string"`

	KanaName2 *string `locationName:"kanaName2" type:"string"`

	Name1 *string `locationName:"name1" type:"string"`

	Name2 *string `locationName:"name2" type:"string"`

	PostName *string `locationName:"postName" type:"string"`

	Pref *string `locationName:"pref" type:"string"`

	PresidentName1 *string `locationName:"presidentName1" type:"string"`

	PresidentName2 *string `locationName:"presidentName2" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	TdbCode *string `locationName:"tdbCode" type:"string"`

	Zip1 *string `locationName:"zip1" type:"string"`

	Zip2 *string `locationName:"zip2" type:"string"`
}

// String returns the string representation
func (s RegisterCorporateInfoForCertificateOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opRegisterCorporateInfoForCertificate = "RegisterCorporateInfoForCertificate"

// RegisterCorporateInfoForCertificateRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using RegisterCorporateInfoForCertificateRequest.
//    req := client.RegisterCorporateInfoForCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/RegisterCorporateInfoForCertificate
func (c *Client) RegisterCorporateInfoForCertificateRequest(input *RegisterCorporateInfoForCertificateInput) RegisterCorporateInfoForCertificateRequest {
	op := &aws.Operation{
		Name:       opRegisterCorporateInfoForCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &RegisterCorporateInfoForCertificateInput{}
	}

	req := c.newRequest(op, input, &RegisterCorporateInfoForCertificateOutput{})
	return RegisterCorporateInfoForCertificateRequest{Request: req, Input: input, Copy: c.RegisterCorporateInfoForCertificateRequest}
}

// RegisterCorporateInfoForCertificateRequest is the request type for the
// RegisterCorporateInfoForCertificate API operation.
type RegisterCorporateInfoForCertificateRequest struct {
	*aws.Request
	Input *RegisterCorporateInfoForCertificateInput
	Copy  func(*RegisterCorporateInfoForCertificateInput) RegisterCorporateInfoForCertificateRequest
}

// Send marshals and sends the RegisterCorporateInfoForCertificate API request.
func (r RegisterCorporateInfoForCertificateRequest) Send(ctx context.Context) (*RegisterCorporateInfoForCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterCorporateInfoForCertificateResponse{
		RegisterCorporateInfoForCertificateOutput: r.Request.Data.(*RegisterCorporateInfoForCertificateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterCorporateInfoForCertificateResponse is the response type for the
// RegisterCorporateInfoForCertificate API operation.
type RegisterCorporateInfoForCertificateResponse struct {
	*RegisterCorporateInfoForCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterCorporateInfoForCertificate request.
func (r *RegisterCorporateInfoForCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}