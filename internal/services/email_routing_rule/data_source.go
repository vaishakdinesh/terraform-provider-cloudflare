// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type EmailRoutingRuleDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &EmailRoutingRuleDataSource{}

func NewEmailRoutingRuleDataSource() datasource.DataSource {
	return &EmailRoutingRuleDataSource{}
}

func (d *EmailRoutingRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_email_routing_rule"
}

func (r *EmailRoutingRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EmailRoutingRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *EmailRoutingRuleDataSource

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
