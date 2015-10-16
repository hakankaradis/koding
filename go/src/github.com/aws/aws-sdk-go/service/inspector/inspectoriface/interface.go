// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package inspectoriface provides an interface for the Amazon Inspector.
package inspectoriface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/inspector"
)

// InspectorAPI is the interface type for inspector.Inspector.
type InspectorAPI interface {
	AddAttributesToFindingsRequest(*inspector.AddAttributesToFindingsInput) (*request.Request, *inspector.AddAttributesToFindingsOutput)

	AddAttributesToFindings(*inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error)

	AttachAssessmentAndRulesPackageRequest(*inspector.AttachAssessmentAndRulesPackageInput) (*request.Request, *inspector.AttachAssessmentAndRulesPackageOutput)

	AttachAssessmentAndRulesPackage(*inspector.AttachAssessmentAndRulesPackageInput) (*inspector.AttachAssessmentAndRulesPackageOutput, error)

	CreateApplicationRequest(*inspector.CreateApplicationInput) (*request.Request, *inspector.CreateApplicationOutput)

	CreateApplication(*inspector.CreateApplicationInput) (*inspector.CreateApplicationOutput, error)

	CreateAssessmentRequest(*inspector.CreateAssessmentInput) (*request.Request, *inspector.CreateAssessmentOutput)

	CreateAssessment(*inspector.CreateAssessmentInput) (*inspector.CreateAssessmentOutput, error)

	CreateResourceGroupRequest(*inspector.CreateResourceGroupInput) (*request.Request, *inspector.CreateResourceGroupOutput)

	CreateResourceGroup(*inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error)

	DeleteApplicationRequest(*inspector.DeleteApplicationInput) (*request.Request, *inspector.DeleteApplicationOutput)

	DeleteApplication(*inspector.DeleteApplicationInput) (*inspector.DeleteApplicationOutput, error)

	DeleteAssessmentRequest(*inspector.DeleteAssessmentInput) (*request.Request, *inspector.DeleteAssessmentOutput)

	DeleteAssessment(*inspector.DeleteAssessmentInput) (*inspector.DeleteAssessmentOutput, error)

	DeleteRunRequest(*inspector.DeleteRunInput) (*request.Request, *inspector.DeleteRunOutput)

	DeleteRun(*inspector.DeleteRunInput) (*inspector.DeleteRunOutput, error)

	DescribeApplicationRequest(*inspector.DescribeApplicationInput) (*request.Request, *inspector.DescribeApplicationOutput)

	DescribeApplication(*inspector.DescribeApplicationInput) (*inspector.DescribeApplicationOutput, error)

	DescribeAssessmentRequest(*inspector.DescribeAssessmentInput) (*request.Request, *inspector.DescribeAssessmentOutput)

	DescribeAssessment(*inspector.DescribeAssessmentInput) (*inspector.DescribeAssessmentOutput, error)

	DescribeCrossAccountAccessRoleRequest(*inspector.DescribeCrossAccountAccessRoleInput) (*request.Request, *inspector.DescribeCrossAccountAccessRoleOutput)

	DescribeCrossAccountAccessRole(*inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error)

	DescribeFindingRequest(*inspector.DescribeFindingInput) (*request.Request, *inspector.DescribeFindingOutput)

	DescribeFinding(*inspector.DescribeFindingInput) (*inspector.DescribeFindingOutput, error)

	DescribeResourceGroupRequest(*inspector.DescribeResourceGroupInput) (*request.Request, *inspector.DescribeResourceGroupOutput)

	DescribeResourceGroup(*inspector.DescribeResourceGroupInput) (*inspector.DescribeResourceGroupOutput, error)

	DescribeRulesPackageRequest(*inspector.DescribeRulesPackageInput) (*request.Request, *inspector.DescribeRulesPackageOutput)

	DescribeRulesPackage(*inspector.DescribeRulesPackageInput) (*inspector.DescribeRulesPackageOutput, error)

	DescribeRunRequest(*inspector.DescribeRunInput) (*request.Request, *inspector.DescribeRunOutput)

	DescribeRun(*inspector.DescribeRunInput) (*inspector.DescribeRunOutput, error)

	DetachAssessmentAndRulesPackageRequest(*inspector.DetachAssessmentAndRulesPackageInput) (*request.Request, *inspector.DetachAssessmentAndRulesPackageOutput)

	DetachAssessmentAndRulesPackage(*inspector.DetachAssessmentAndRulesPackageInput) (*inspector.DetachAssessmentAndRulesPackageOutput, error)

	GetAssessmentTelemetryRequest(*inspector.GetAssessmentTelemetryInput) (*request.Request, *inspector.GetAssessmentTelemetryOutput)

	GetAssessmentTelemetry(*inspector.GetAssessmentTelemetryInput) (*inspector.GetAssessmentTelemetryOutput, error)

	ListApplicationsRequest(*inspector.ListApplicationsInput) (*request.Request, *inspector.ListApplicationsOutput)

	ListApplications(*inspector.ListApplicationsInput) (*inspector.ListApplicationsOutput, error)

	ListAssessmentAgentsRequest(*inspector.ListAssessmentAgentsInput) (*request.Request, *inspector.ListAssessmentAgentsOutput)

	ListAssessmentAgents(*inspector.ListAssessmentAgentsInput) (*inspector.ListAssessmentAgentsOutput, error)

	ListAssessmentsRequest(*inspector.ListAssessmentsInput) (*request.Request, *inspector.ListAssessmentsOutput)

	ListAssessments(*inspector.ListAssessmentsInput) (*inspector.ListAssessmentsOutput, error)

	ListAttachedAssessmentsRequest(*inspector.ListAttachedAssessmentsInput) (*request.Request, *inspector.ListAttachedAssessmentsOutput)

	ListAttachedAssessments(*inspector.ListAttachedAssessmentsInput) (*inspector.ListAttachedAssessmentsOutput, error)

	ListAttachedRulesPackagesRequest(*inspector.ListAttachedRulesPackagesInput) (*request.Request, *inspector.ListAttachedRulesPackagesOutput)

	ListAttachedRulesPackages(*inspector.ListAttachedRulesPackagesInput) (*inspector.ListAttachedRulesPackagesOutput, error)

	ListFindingsRequest(*inspector.ListFindingsInput) (*request.Request, *inspector.ListFindingsOutput)

	ListFindings(*inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error)

	ListRulesPackagesRequest(*inspector.ListRulesPackagesInput) (*request.Request, *inspector.ListRulesPackagesOutput)

	ListRulesPackages(*inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error)

	ListRunsRequest(*inspector.ListRunsInput) (*request.Request, *inspector.ListRunsOutput)

	ListRuns(*inspector.ListRunsInput) (*inspector.ListRunsOutput, error)

	ListTagsForResourceRequest(*inspector.ListTagsForResourceInput) (*request.Request, *inspector.ListTagsForResourceOutput)

	ListTagsForResource(*inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error)

	LocalizeTextRequest(*inspector.LocalizeTextInput) (*request.Request, *inspector.LocalizeTextOutput)

	LocalizeText(*inspector.LocalizeTextInput) (*inspector.LocalizeTextOutput, error)

	PreviewAgentsForResourceGroupRequest(*inspector.PreviewAgentsForResourceGroupInput) (*request.Request, *inspector.PreviewAgentsForResourceGroupOutput)

	PreviewAgentsForResourceGroup(*inspector.PreviewAgentsForResourceGroupInput) (*inspector.PreviewAgentsForResourceGroupOutput, error)

	RegisterCrossAccountAccessRoleRequest(*inspector.RegisterCrossAccountAccessRoleInput) (*request.Request, *inspector.RegisterCrossAccountAccessRoleOutput)

	RegisterCrossAccountAccessRole(*inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error)

	RemoveAttributesFromFindingsRequest(*inspector.RemoveAttributesFromFindingsInput) (*request.Request, *inspector.RemoveAttributesFromFindingsOutput)

	RemoveAttributesFromFindings(*inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error)

	RunAssessmentRequest(*inspector.RunAssessmentInput) (*request.Request, *inspector.RunAssessmentOutput)

	RunAssessment(*inspector.RunAssessmentInput) (*inspector.RunAssessmentOutput, error)

	SetTagsForResourceRequest(*inspector.SetTagsForResourceInput) (*request.Request, *inspector.SetTagsForResourceOutput)

	SetTagsForResource(*inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error)

	StartDataCollectionRequest(*inspector.StartDataCollectionInput) (*request.Request, *inspector.StartDataCollectionOutput)

	StartDataCollection(*inspector.StartDataCollectionInput) (*inspector.StartDataCollectionOutput, error)

	StopDataCollectionRequest(*inspector.StopDataCollectionInput) (*request.Request, *inspector.StopDataCollectionOutput)

	StopDataCollection(*inspector.StopDataCollectionInput) (*inspector.StopDataCollectionOutput, error)

	UpdateApplicationRequest(*inspector.UpdateApplicationInput) (*request.Request, *inspector.UpdateApplicationOutput)

	UpdateApplication(*inspector.UpdateApplicationInput) (*inspector.UpdateApplicationOutput, error)

	UpdateAssessmentRequest(*inspector.UpdateAssessmentInput) (*request.Request, *inspector.UpdateAssessmentOutput)

	UpdateAssessment(*inspector.UpdateAssessmentInput) (*inspector.UpdateAssessmentOutput, error)
}
