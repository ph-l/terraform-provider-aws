// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	bedrock_sdkv2 "github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceCustomModel,
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "model_arn",
			},
		},
		{
			Factory: newDataSourceCustomModels,
		},
		{
			Factory: newDataSourceFoundationModel,
			Name:    "Foundation Model",
		},
		{
			Factory: newFoundationModelsDataSource,
			Name:    "Foundation Models",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceCustomModel,
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "model_arn",
			},
		},
		{
			Factory: newResourceModelInvocationLoggingConfiguration,
			Name:    "Model Invocation Logging Configuration",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Bedrock
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*bedrock_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return bedrock_sdkv2.NewFromConfig(cfg, func(o *bedrock_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
