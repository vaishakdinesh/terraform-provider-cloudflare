// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRulesResultListDataSourceEnvelope struct {
	Result *[]*EmailRoutingRulesItemsDataSourceModel `json:"result,computed"`
}

type EmailRoutingRulesDataSourceModel struct {
	ZoneIdentifier types.String                              `tfsdk:"zone_identifier" path:"zone_identifier"`
	Enabled        types.Bool                                `tfsdk:"enabled" query:"enabled"`
	Page           types.Float64                             `tfsdk:"page" query:"page"`
	PerPage        types.Float64                             `tfsdk:"per_page" query:"per_page"`
	MaxItems       types.Int64                               `tfsdk:"max_items"`
	Items          *[]*EmailRoutingRulesItemsDataSourceModel `tfsdk:"items"`
}

type EmailRoutingRulesItemsDataSourceModel struct {
	ID       types.String                                      `tfsdk:"id" json:"id,computed"`
	Actions  *[]*EmailRoutingRulesItemsActionsDataSourceModel  `tfsdk:"actions" json:"actions"`
	Enabled  types.Bool                                        `tfsdk:"enabled" json:"enabled,computed"`
	Matchers *[]*EmailRoutingRulesItemsMatchersDataSourceModel `tfsdk:"matchers" json:"matchers"`
	Name     types.String                                      `tfsdk:"name" json:"name"`
	Priority types.Float64                                     `tfsdk:"priority" json:"priority,computed"`
	Tag      types.String                                      `tfsdk:"tag" json:"tag,computed"`
}

type EmailRoutingRulesItemsActionsDataSourceModel struct {
	Type  types.String    `tfsdk:"type" json:"type,computed"`
	Value *[]types.String `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRulesItemsMatchersDataSourceModel struct {
	Field types.String `tfsdk:"field" json:"field,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}