// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package servicecatalogiface provides an interface to enable mocking the AWS Service Catalog service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicecatalogiface

import (
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

// ServiceCatalogAPI provides an interface to enable mocking the
// servicecatalog.ServiceCatalog service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Service Catalog.
//    func myFunc(svc servicecatalogiface.ServiceCatalogAPI) bool {
//        // Make svc.AcceptPortfolioShare request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := servicecatalog.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServiceCatalogClient struct {
//        servicecatalogiface.ServiceCatalogAPI
//    }
//    func (m *mockServiceCatalogClient) AcceptPortfolioShare(input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServiceCatalogClient{}
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
type ServiceCatalogAPI interface {
	AcceptPortfolioShareRequest(*servicecatalog.AcceptPortfolioShareInput) servicecatalog.AcceptPortfolioShareRequest

	AssociatePrincipalWithPortfolioRequest(*servicecatalog.AssociatePrincipalWithPortfolioInput) servicecatalog.AssociatePrincipalWithPortfolioRequest

	AssociateProductWithPortfolioRequest(*servicecatalog.AssociateProductWithPortfolioInput) servicecatalog.AssociateProductWithPortfolioRequest

	AssociateTagOptionWithResourceRequest(*servicecatalog.AssociateTagOptionWithResourceInput) servicecatalog.AssociateTagOptionWithResourceRequest

	CopyProductRequest(*servicecatalog.CopyProductInput) servicecatalog.CopyProductRequest

	CreateConstraintRequest(*servicecatalog.CreateConstraintInput) servicecatalog.CreateConstraintRequest

	CreatePortfolioRequest(*servicecatalog.CreatePortfolioInput) servicecatalog.CreatePortfolioRequest

	CreatePortfolioShareRequest(*servicecatalog.CreatePortfolioShareInput) servicecatalog.CreatePortfolioShareRequest

	CreateProductRequest(*servicecatalog.CreateProductInput) servicecatalog.CreateProductRequest

	CreateProvisionedProductPlanRequest(*servicecatalog.CreateProvisionedProductPlanInput) servicecatalog.CreateProvisionedProductPlanRequest

	CreateProvisioningArtifactRequest(*servicecatalog.CreateProvisioningArtifactInput) servicecatalog.CreateProvisioningArtifactRequest

	CreateTagOptionRequest(*servicecatalog.CreateTagOptionInput) servicecatalog.CreateTagOptionRequest

	DeleteConstraintRequest(*servicecatalog.DeleteConstraintInput) servicecatalog.DeleteConstraintRequest

	DeletePortfolioRequest(*servicecatalog.DeletePortfolioInput) servicecatalog.DeletePortfolioRequest

	DeletePortfolioShareRequest(*servicecatalog.DeletePortfolioShareInput) servicecatalog.DeletePortfolioShareRequest

	DeleteProductRequest(*servicecatalog.DeleteProductInput) servicecatalog.DeleteProductRequest

	DeleteProvisionedProductPlanRequest(*servicecatalog.DeleteProvisionedProductPlanInput) servicecatalog.DeleteProvisionedProductPlanRequest

	DeleteProvisioningArtifactRequest(*servicecatalog.DeleteProvisioningArtifactInput) servicecatalog.DeleteProvisioningArtifactRequest

	DeleteTagOptionRequest(*servicecatalog.DeleteTagOptionInput) servicecatalog.DeleteTagOptionRequest

	DescribeConstraintRequest(*servicecatalog.DescribeConstraintInput) servicecatalog.DescribeConstraintRequest

	DescribeCopyProductStatusRequest(*servicecatalog.DescribeCopyProductStatusInput) servicecatalog.DescribeCopyProductStatusRequest

	DescribePortfolioRequest(*servicecatalog.DescribePortfolioInput) servicecatalog.DescribePortfolioRequest

	DescribeProductRequest(*servicecatalog.DescribeProductInput) servicecatalog.DescribeProductRequest

	DescribeProductAsAdminRequest(*servicecatalog.DescribeProductAsAdminInput) servicecatalog.DescribeProductAsAdminRequest

	DescribeProductViewRequest(*servicecatalog.DescribeProductViewInput) servicecatalog.DescribeProductViewRequest

	DescribeProvisionedProductRequest(*servicecatalog.DescribeProvisionedProductInput) servicecatalog.DescribeProvisionedProductRequest

	DescribeProvisionedProductPlanRequest(*servicecatalog.DescribeProvisionedProductPlanInput) servicecatalog.DescribeProvisionedProductPlanRequest

	DescribeProvisioningArtifactRequest(*servicecatalog.DescribeProvisioningArtifactInput) servicecatalog.DescribeProvisioningArtifactRequest

	DescribeProvisioningParametersRequest(*servicecatalog.DescribeProvisioningParametersInput) servicecatalog.DescribeProvisioningParametersRequest

	DescribeRecordRequest(*servicecatalog.DescribeRecordInput) servicecatalog.DescribeRecordRequest

	DescribeTagOptionRequest(*servicecatalog.DescribeTagOptionInput) servicecatalog.DescribeTagOptionRequest

	DisassociatePrincipalFromPortfolioRequest(*servicecatalog.DisassociatePrincipalFromPortfolioInput) servicecatalog.DisassociatePrincipalFromPortfolioRequest

	DisassociateProductFromPortfolioRequest(*servicecatalog.DisassociateProductFromPortfolioInput) servicecatalog.DisassociateProductFromPortfolioRequest

	DisassociateTagOptionFromResourceRequest(*servicecatalog.DisassociateTagOptionFromResourceInput) servicecatalog.DisassociateTagOptionFromResourceRequest

	ExecuteProvisionedProductPlanRequest(*servicecatalog.ExecuteProvisionedProductPlanInput) servicecatalog.ExecuteProvisionedProductPlanRequest

	ListAcceptedPortfolioSharesRequest(*servicecatalog.ListAcceptedPortfolioSharesInput) servicecatalog.ListAcceptedPortfolioSharesRequest

	ListConstraintsForPortfolioRequest(*servicecatalog.ListConstraintsForPortfolioInput) servicecatalog.ListConstraintsForPortfolioRequest

	ListLaunchPathsRequest(*servicecatalog.ListLaunchPathsInput) servicecatalog.ListLaunchPathsRequest

	ListPortfolioAccessRequest(*servicecatalog.ListPortfolioAccessInput) servicecatalog.ListPortfolioAccessRequest

	ListPortfoliosRequest(*servicecatalog.ListPortfoliosInput) servicecatalog.ListPortfoliosRequest

	ListPortfoliosForProductRequest(*servicecatalog.ListPortfoliosForProductInput) servicecatalog.ListPortfoliosForProductRequest

	ListPrincipalsForPortfolioRequest(*servicecatalog.ListPrincipalsForPortfolioInput) servicecatalog.ListPrincipalsForPortfolioRequest

	ListProvisionedProductPlansRequest(*servicecatalog.ListProvisionedProductPlansInput) servicecatalog.ListProvisionedProductPlansRequest

	ListProvisioningArtifactsRequest(*servicecatalog.ListProvisioningArtifactsInput) servicecatalog.ListProvisioningArtifactsRequest

	ListRecordHistoryRequest(*servicecatalog.ListRecordHistoryInput) servicecatalog.ListRecordHistoryRequest

	ListResourcesForTagOptionRequest(*servicecatalog.ListResourcesForTagOptionInput) servicecatalog.ListResourcesForTagOptionRequest

	ListTagOptionsRequest(*servicecatalog.ListTagOptionsInput) servicecatalog.ListTagOptionsRequest

	ProvisionProductRequest(*servicecatalog.ProvisionProductInput) servicecatalog.ProvisionProductRequest

	RejectPortfolioShareRequest(*servicecatalog.RejectPortfolioShareInput) servicecatalog.RejectPortfolioShareRequest

	ScanProvisionedProductsRequest(*servicecatalog.ScanProvisionedProductsInput) servicecatalog.ScanProvisionedProductsRequest

	SearchProductsRequest(*servicecatalog.SearchProductsInput) servicecatalog.SearchProductsRequest

	SearchProductsAsAdminRequest(*servicecatalog.SearchProductsAsAdminInput) servicecatalog.SearchProductsAsAdminRequest

	SearchProvisionedProductsRequest(*servicecatalog.SearchProvisionedProductsInput) servicecatalog.SearchProvisionedProductsRequest

	TerminateProvisionedProductRequest(*servicecatalog.TerminateProvisionedProductInput) servicecatalog.TerminateProvisionedProductRequest

	UpdateConstraintRequest(*servicecatalog.UpdateConstraintInput) servicecatalog.UpdateConstraintRequest

	UpdatePortfolioRequest(*servicecatalog.UpdatePortfolioInput) servicecatalog.UpdatePortfolioRequest

	UpdateProductRequest(*servicecatalog.UpdateProductInput) servicecatalog.UpdateProductRequest

	UpdateProvisionedProductRequest(*servicecatalog.UpdateProvisionedProductInput) servicecatalog.UpdateProvisionedProductRequest

	UpdateProvisioningArtifactRequest(*servicecatalog.UpdateProvisioningArtifactInput) servicecatalog.UpdateProvisioningArtifactRequest

	UpdateTagOptionRequest(*servicecatalog.UpdateTagOptionInput) servicecatalog.UpdateTagOptionRequest
}

var _ ServiceCatalogAPI = (*servicecatalog.ServiceCatalog)(nil)
