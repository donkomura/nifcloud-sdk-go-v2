// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package nasiface provides an interface to enable mocking the NIFCLOUD NAS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package nasiface

import (
	"github.com/aokumasan/nifcloud-sdk-go-v2/service/nas"
)

// ClientAPI provides an interface to enable mocking the
// nas.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // nas.
//    func myFunc(svc nasiface.ClientAPI) bool {
//        // Make svc.AuthorizeNASSecurityGroupIngress request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := nas.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        nasiface.ClientPI
//    }
//    func (m *mockClientClient) AuthorizeNASSecurityGroupIngress(input *nas.AuthorizeNASSecurityGroupIngressInput) (*nas.AuthorizeNASSecurityGroupIngressOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AuthorizeNASSecurityGroupIngressRequest(*nas.AuthorizeNASSecurityGroupIngressInput) nas.AuthorizeNASSecurityGroupIngressRequest

	ClearNASSessionRequest(*nas.ClearNASSessionInput) nas.ClearNASSessionRequest

	CreateNASInstanceRequest(*nas.CreateNASInstanceInput) nas.CreateNASInstanceRequest

	CreateNASSecurityGroupRequest(*nas.CreateNASSecurityGroupInput) nas.CreateNASSecurityGroupRequest

	DeleteNASInstanceRequest(*nas.DeleteNASInstanceInput) nas.DeleteNASInstanceRequest

	DeleteNASSecurityGroupRequest(*nas.DeleteNASSecurityGroupInput) nas.DeleteNASSecurityGroupRequest

	DescribeNASInstancesRequest(*nas.DescribeNASInstancesInput) nas.DescribeNASInstancesRequest

	DescribeNASSecurityGroupsRequest(*nas.DescribeNASSecurityGroupsInput) nas.DescribeNASSecurityGroupsRequest

	GetMetricStatisticsRequest(*nas.GetMetricStatisticsInput) nas.GetMetricStatisticsRequest

	ModifyNASInstanceRequest(*nas.ModifyNASInstanceInput) nas.ModifyNASInstanceRequest

	ModifyNASSecurityGroupRequest(*nas.ModifyNASSecurityGroupInput) nas.ModifyNASSecurityGroupRequest

	RevokeNASSecurityGroupIngressRequest(*nas.RevokeNASSecurityGroupIngressInput) nas.RevokeNASSecurityGroupIngressRequest
}

var _ ClientAPI = (*nas.Client)(nil)
