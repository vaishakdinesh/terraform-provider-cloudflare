// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/pages"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type PagesDomainsDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &PagesDomainsDataSource{}

func NewPagesDomainsDataSource() datasource.DataSource {
	return &PagesDomainsDataSource{}
}

func (d *PagesDomainsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pages_domains"
}

func (d *PagesDomainsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *PagesDomainsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PagesDomainsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	items := &[]*PagesDomainsResultDataSourceModel{}
	env := PagesDomainsResultListDataSourceEnvelope{items}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []*PagesDomainsResultDataSourceModel{}

	page, err := d.client.Pages.Projects.Domains.List(
		ctx,
		data.ProjectName.ValueString(),
		pages.ProjectDomainListParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Result) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, *items...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:maxItems]
	data.Result = &acc

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}