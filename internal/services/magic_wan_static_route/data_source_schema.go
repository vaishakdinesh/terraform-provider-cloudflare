// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_static_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &MagicWANStaticRouteDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"route_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"route": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"nexthop": schema.StringAttribute{
						Description: "The next-hop IP Address for the static route.",
						Computed:    true,
					},
					"prefix": schema.StringAttribute{
						Description: "IP Prefix in Classless Inter-Domain Routing format.",
						Computed:    true,
					},
					"priority": schema.Int64Attribute{
						Description: "Priority of the static route.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
						Optional:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "When the route was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"description": schema.StringAttribute{
						Description: "An optional human provided description of the static route.",
						Computed:    true,
						Optional:    true,
					},
					"modified_on": schema.StringAttribute{
						Description: "When the route was last modified.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"scope": schema.SingleNestedAttribute{
						Description: "Used only for ECMP routes.",
						Computed:    true,
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"colo_names": schema.ListAttribute{
								Description: "List of colo names for the ECMP scope.",
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"colo_regions": schema.ListAttribute{
								Description: "List of colo regions for the ECMP scope.",
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"weight": schema.Int64Attribute{
						Description: "Optional weight of the ECMP scope - if provided.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *MagicWANStaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicWANStaticRouteDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}