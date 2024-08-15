// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANGRETunnelResultDataSourceEnvelope struct {
	Result MagicWANGRETunnelDataSourceModel `json:"result,computed"`
}

type MagicWANGRETunnelDataSourceModel struct {
	AccountID   types.String                               `tfsdk:"account_id" path:"account_id"`
	GRETunnelID types.String                               `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id"`
	GRETunnel   *MagicWANGRETunnelGRETunnelDataSourceModel `tfsdk:"gre_tunnel" json:"gre_tunnel"`
}

type MagicWANGRETunnelGRETunnelDataSourceModel struct {
	CloudflareGREEndpoint types.String                                          `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                          `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                          `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                          `tfsdk:"name" json:"name,computed"`
	ID                    types.String                                          `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                                     `tfsdk:"created_on" json:"created_on,computed"`
	Description           types.String                                          `tfsdk:"description" json:"description"`
	HealthCheck           *MagicWANGRETunnelGRETunnelHealthCheckDataSourceModel `tfsdk:"health_check" json:"health_check"`
	ModifiedOn            timetypes.RFC3339                                     `tfsdk:"modified_on" json:"modified_on,computed"`
	Mtu                   types.Int64                                           `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                           `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelGRETunnelHealthCheckDataSourceModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}