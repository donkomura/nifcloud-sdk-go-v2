// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package hatobaiface provides an interface to enable mocking the NIFCLOUD Hatoba beta service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package hatobaiface

import (
	"github.com/alice02/nifcloud-sdk-go-v2/service/hatoba"
)

// HatobaAPI provides an interface to enable mocking the
// hatoba.Hatoba service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // NIFCLOUD Hatoba beta.
//    func myFunc(svc hatobaiface.HatobaAPI) bool {
//        // Make svc.AuthorizeFirewallGroup request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := hatoba.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockHatobaClient struct {
//        hatobaiface.HatobaAPI
//    }
//    func (m *mockHatobaClient) AuthorizeFirewallGroup(input *hatoba.AuthorizeFirewallGroupInput) (*hatoba.AuthorizeFirewallGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockHatobaClient{}
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
type HatobaAPI interface {
	AuthorizeFirewallGroupRequest(*hatoba.AuthorizeFirewallGroupInput) hatoba.AuthorizeFirewallGroupRequest

	CreateClusterRequest(*hatoba.CreateClusterInput) hatoba.CreateClusterRequest

	CreateFirewallGroupRequest(*hatoba.CreateFirewallGroupInput) hatoba.CreateFirewallGroupRequest

	CreateNodePoolRequest(*hatoba.CreateNodePoolInput) hatoba.CreateNodePoolRequest

	CreateSnapshotRequest(*hatoba.CreateSnapshotInput) hatoba.CreateSnapshotRequest

	DeleteClusterRequest(*hatoba.DeleteClusterInput) hatoba.DeleteClusterRequest

	DeleteClustersRequest(*hatoba.DeleteClustersInput) hatoba.DeleteClustersRequest

	DeleteFirewallGroupRequest(*hatoba.DeleteFirewallGroupInput) hatoba.DeleteFirewallGroupRequest

	DeleteFirewallGroupsRequest(*hatoba.DeleteFirewallGroupsInput) hatoba.DeleteFirewallGroupsRequest

	DeleteNodePoolRequest(*hatoba.DeleteNodePoolInput) hatoba.DeleteNodePoolRequest

	DeleteNodePoolsRequest(*hatoba.DeleteNodePoolsInput) hatoba.DeleteNodePoolsRequest

	DeleteSnapshotRequest(*hatoba.DeleteSnapshotInput) hatoba.DeleteSnapshotRequest

	DeleteSnapshotsRequest(*hatoba.DeleteSnapshotsInput) hatoba.DeleteSnapshotsRequest

	GetClusterRequest(*hatoba.GetClusterInput) hatoba.GetClusterRequest

	GetClusterCredentialsRequest(*hatoba.GetClusterCredentialsInput) hatoba.GetClusterCredentialsRequest

	GetFirewallGroupRequest(*hatoba.GetFirewallGroupInput) hatoba.GetFirewallGroupRequest

	GetNodePoolRequest(*hatoba.GetNodePoolInput) hatoba.GetNodePoolRequest

	GetServerConfigRequest(*hatoba.GetServerConfigInput) hatoba.GetServerConfigRequest

	GetSnapshotRequest(*hatoba.GetSnapshotInput) hatoba.GetSnapshotRequest

	ListClustersRequest(*hatoba.ListClustersInput) hatoba.ListClustersRequest

	ListFirewallGroupsRequest(*hatoba.ListFirewallGroupsInput) hatoba.ListFirewallGroupsRequest

	ListNodePoolsRequest(*hatoba.ListNodePoolsInput) hatoba.ListNodePoolsRequest

	ListSnapshotsRequest(*hatoba.ListSnapshotsInput) hatoba.ListSnapshotsRequest

	RestoreClusterFromSnapshotRequest(*hatoba.RestoreClusterFromSnapshotInput) hatoba.RestoreClusterFromSnapshotRequest

	RevokeFirewallGroupRequest(*hatoba.RevokeFirewallGroupInput) hatoba.RevokeFirewallGroupRequest

	SetNodePoolSizeRequest(*hatoba.SetNodePoolSizeInput) hatoba.SetNodePoolSizeRequest

	UpdateClusterRequest(*hatoba.UpdateClusterInput) hatoba.UpdateClusterRequest

	UpdateFirewallGroupRequest(*hatoba.UpdateFirewallGroupInput) hatoba.UpdateFirewallGroupRequest

	UpdateSnapshotRequest(*hatoba.UpdateSnapshotInput) hatoba.UpdateSnapshotRequest
}

var _ HatobaAPI = (*hatoba.Hatoba)(nil)
