// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package acmpcaiface provides an interface to enable mocking the AWS Certificate Manager Private Certificate Authority service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package acmpcaiface

import (
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

// ACMPCAAPI provides an interface to enable mocking the
// acmpca.ACMPCA service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Certificate Manager Private Certificate Authority.
//    func myFunc(svc acmpcaiface.ACMPCAAPI) bool {
//        // Make svc.CreateCertificateAuthority request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := acmpca.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockACMPCAClient struct {
//        acmpcaiface.ACMPCAAPI
//    }
//    func (m *mockACMPCAClient) CreateCertificateAuthority(input *acmpca.CreateCertificateAuthorityInput) (*acmpca.CreateCertificateAuthorityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockACMPCAClient{}
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
type ACMPCAAPI interface {
	CreateCertificateAuthorityRequest(*acmpca.CreateCertificateAuthorityInput) acmpca.CreateCertificateAuthorityRequest

	CreateCertificateAuthorityAuditReportRequest(*acmpca.CreateCertificateAuthorityAuditReportInput) acmpca.CreateCertificateAuthorityAuditReportRequest

	DeleteCertificateAuthorityRequest(*acmpca.DeleteCertificateAuthorityInput) acmpca.DeleteCertificateAuthorityRequest

	DescribeCertificateAuthorityRequest(*acmpca.DescribeCertificateAuthorityInput) acmpca.DescribeCertificateAuthorityRequest

	DescribeCertificateAuthorityAuditReportRequest(*acmpca.DescribeCertificateAuthorityAuditReportInput) acmpca.DescribeCertificateAuthorityAuditReportRequest

	GetCertificateRequest(*acmpca.GetCertificateInput) acmpca.GetCertificateRequest

	GetCertificateAuthorityCertificateRequest(*acmpca.GetCertificateAuthorityCertificateInput) acmpca.GetCertificateAuthorityCertificateRequest

	GetCertificateAuthorityCsrRequest(*acmpca.GetCertificateAuthorityCsrInput) acmpca.GetCertificateAuthorityCsrRequest

	ImportCertificateAuthorityCertificateRequest(*acmpca.ImportCertificateAuthorityCertificateInput) acmpca.ImportCertificateAuthorityCertificateRequest

	IssueCertificateRequest(*acmpca.IssueCertificateInput) acmpca.IssueCertificateRequest

	ListCertificateAuthoritiesRequest(*acmpca.ListCertificateAuthoritiesInput) acmpca.ListCertificateAuthoritiesRequest

	ListTagsRequest(*acmpca.ListTagsInput) acmpca.ListTagsRequest

	RevokeCertificateRequest(*acmpca.RevokeCertificateInput) acmpca.RevokeCertificateRequest

	TagCertificateAuthorityRequest(*acmpca.TagCertificateAuthorityInput) acmpca.TagCertificateAuthorityRequest

	UntagCertificateAuthorityRequest(*acmpca.UntagCertificateAuthorityInput) acmpca.UntagCertificateAuthorityRequest

	UpdateCertificateAuthorityRequest(*acmpca.UpdateCertificateAuthorityInput) acmpca.UpdateCertificateAuthorityRequest
}

var _ ACMPCAAPI = (*acmpca.ACMPCA)(nil)
