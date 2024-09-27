// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteWANResultEnvelope struct {
	Result MagicTransitSiteWANModel `json:"result"`
}

type MagicTransitSiteWANModel struct {
	AccountID        types.String                                                       `tfsdk:"account_id" path:"account_id,required"`
	SiteID           types.String                                                       `tfsdk:"site_id" path:"site_id,required"`
	WANID            types.String                                                       `tfsdk:"wan_id" path:"wan_id,optional"`
	Physport         types.Int64                                                        `tfsdk:"physport" json:"physport,required"`
	VlanTag          types.Int64                                                        `tfsdk:"vlan_tag" json:"vlan_tag,required"`
	Name             types.String                                                       `tfsdk:"name" json:"name,computed_optional"`
	Priority         types.Int64                                                        `tfsdk:"priority" json:"priority,computed_optional"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteWANStaticAddressingModel] `tfsdk:"static_addressing" json:"static_addressing,computed_optional"`
	ID               types.String                                                       `tfsdk:"id" json:"id,computed"`
}

type MagicTransitSiteWANStaticAddressingModel struct {
	Address          types.String `tfsdk:"address" json:"address,required"`
	GatewayAddress   types.String `tfsdk:"gateway_address" json:"gateway_address,required"`
	SecondaryAddress types.String `tfsdk:"secondary_address" json:"secondary_address,computed_optional"`
}