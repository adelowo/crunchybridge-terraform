package provider

import (
	"context"

	gocrunchybridge "github.com/adelowo/go-crunchybridge"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CrunchybridgeProvider{}
var _ provider.ProviderWithFunctions = &CrunchybridgeProvider{}

// CrunchybridgeProvider defines the provider implementation.
type CrunchybridgeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CrunchybridgeProviderModel describes the provider data model.
type CrunchybridgeProviderModel struct {
	Secret    types.String `tfsdk:"secret"`
	UserAgent types.String `tfsdk:"user_agent"`
}

func (p *CrunchybridgeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "crunchybridge"
	resp.Version = p.version
}

func (p *CrunchybridgeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"secret": schema.StringAttribute{
				MarkdownDescription: "Your Crunchybridge secret key",
				Required:            true,
				Sensitive:           true,
			},
			"user_agent": schema.StringAttribute{
				MarkdownDescription: "Custom useragent",
				Optional:            true,
			},
		},
	}
}

func (p *CrunchybridgeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CrunchybridgeProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var opts []gocrunchybridge.Option

	apiKey := gocrunchybridge.APIKey(data.Secret.String())

	opts = append(opts, gocrunchybridge.WithAPIKey(apiKey))

	if isStringEmpty(data.UserAgent.String()) {
		if !data.UserAgent.IsNull() && !data.UserAgent.IsUnknown() {
			opts = append(opts, gocrunchybridge.WithUserAgent(data.UserAgent.String()))
		}
	}

	client, err := gocrunchybridge.New(opts...)
	if err != nil {
		resp.Diagnostics.AddError("could not set up sdk", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CrunchybridgeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *CrunchybridgeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccountDatasource,
	}
}

func (p *CrunchybridgeProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CrunchybridgeProvider{
			version: version,
		}
	}
}
