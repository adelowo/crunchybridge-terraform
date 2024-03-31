package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type AccountDataSource struct {
}

func NewAccountDatasource() datasource.DataSource {
	return &AccountDataSource{}
}

func (a *AccountDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "crunchybridge_account"
}

func (a *AccountDataSource) Schema(ctx context.Context,
	req datasource.SchemaRequest,
	resp *datasource.SchemaResponse) {

	resp.Schema = schema.Schema{
		MarkdownDescription: "Datasource for retrieving the authenticated account resource data",
	}
}

func (a *AccountDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {

}

func (a *AccountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}
