// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRuleResultDataSourceEnvelope struct {
	Result UserAgentBlockingRuleDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRuleResultListDataSourceEnvelope struct {
	Result *[]*UserAgentBlockingRuleDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRuleDataSourceModel struct {
	ZoneIdentifier types.String                                       `tfsdk:"zone_identifier" path:"zone_identifier"`
	ID             types.String                                       `tfsdk:"id" path:"id"`
	Configuration  *UserAgentBlockingRuleConfigurationDataSourceModel `tfsdk:"configuration" json:"configuration"`
	Description    types.String                                       `tfsdk:"description" json:"description"`
	Mode           types.String                                       `tfsdk:"mode" json:"mode"`
	Paused         types.Bool                                         `tfsdk:"paused" json:"paused"`
	FindOneBy      *UserAgentBlockingRuleFindOneByDataSourceModel     `tfsdk:"find_one_by"`
}

type UserAgentBlockingRuleConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}

type UserAgentBlockingRuleFindOneByDataSourceModel struct {
	ZoneIdentifier    types.String  `tfsdk:"zone_identifier" path:"zone_identifier"`
	Description       types.String  `tfsdk:"description" query:"description"`
	DescriptionSearch types.String  `tfsdk:"description_search" query:"description_search"`
	Page              types.Float64 `tfsdk:"page" query:"page"`
	PerPage           types.Float64 `tfsdk:"per_page" query:"per_page"`
	UASearch          types.String  `tfsdk:"ua_search" query:"ua_search"`
}