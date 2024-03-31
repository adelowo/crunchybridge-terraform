package provider

import (
	"context"
	"fmt"

	gocrunchybridge "github.com/adelowo/go-crunchybridge"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type AccountDataSource struct {
	client *gocrunchybridge.Client
}

type AccountModel struct {
	ID            types.String `tfsdk:"id"`
	DefaultTeamID types.String `tfsdk:"default_team_id"`
	Email         types.String `tfsdk:"email"`
	MFAEnabled    types.Bool   `tfsdk:"mfa_enabled"`
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
		Attributes:          map[string]schema.Attribute{},
	}
}

func (a *AccountDataSource) Configure(ctx context.Context,
	req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {

	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*gocrunchybridge.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Provider not correctly configured",
			"Please configure the provider setting your secret key",
		)

		return
	}

	a.client = client
}

func (a *AccountDataSource) Read(ctx context.Context,
	req datasource.ReadRequest, resp *datasource.ReadResponse) {

	tflog.Trace(ctx, "Fectching your account details")

	account, err := a.client.Account.User(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid secret key",
			fmt.Errorf("Your authenticated account could not be retrieved. Please verify your secret key..%v", err).Error(),
		)
		return
	}

	data := AccountModel{
		ID:            types.StringValue(account.ID),
		DefaultTeamID: types.StringValue(account.DefaultTeamID),
		Email:         types.StringValue(account.Email),
		MFAEnabled:    types.BoolValue(account.MultiFactorEnabled),
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
