// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codebuild

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	codebuild_sdkv2 "github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceProject,
			TypeName: "aws_codebuild_project",
			Name:     "Project",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  ResourceReportGroup,
			TypeName: "aws_codebuild_report_group",
			Name:     "Report Group",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_codebuild_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  ResourceSourceCredential,
			TypeName: "aws_codebuild_source_credential",
		},
		{
			Factory:  ResourceWebhook,
			TypeName: "aws_codebuild_webhook",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeBuild
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*codebuild_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return codebuild_sdkv2.NewFromConfig(cfg, func(o *codebuild_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
