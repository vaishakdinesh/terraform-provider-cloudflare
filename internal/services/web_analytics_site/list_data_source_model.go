// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_site

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WebAnalyticsSitesResultListDataSourceEnvelope struct {
	Result *[]*WebAnalyticsSitesItemsDataSourceModel `json:"result,computed"`
}

type WebAnalyticsSitesDataSourceModel struct {
	AccountID types.String                              `tfsdk:"account_id" path:"account_id"`
	OrderBy   types.String                              `tfsdk:"order_by" query:"order_by"`
	Page      types.Float64                             `tfsdk:"page" query:"page"`
	PerPage   types.Float64                             `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                               `tfsdk:"max_items"`
	Items     *[]*WebAnalyticsSitesItemsDataSourceModel `tfsdk:"items"`
}

type WebAnalyticsSitesItemsDataSourceModel struct {
	AutoInstall types.Bool                                     `tfsdk:"auto_install" json:"auto_install"`
	Created     timetypes.RFC3339                              `tfsdk:"created" json:"created,computed"`
	Rules       *[]*WebAnalyticsSitesItemsRulesDataSourceModel `tfsdk:"rules" json:"rules"`
	Ruleset     *WebAnalyticsSitesItemsRulesetDataSourceModel  `tfsdk:"ruleset" json:"ruleset"`
	SiteTag     types.String                                   `tfsdk:"site_tag" json:"site_tag"`
	SiteToken   types.String                                   `tfsdk:"site_token" json:"site_token"`
	Snippet     types.String                                   `tfsdk:"snippet" json:"snippet"`
}

type WebAnalyticsSitesItemsRulesDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed"`
	Host      types.String      `tfsdk:"host" json:"host"`
	Inclusive types.Bool        `tfsdk:"inclusive" json:"inclusive"`
	IsPaused  types.Bool        `tfsdk:"is_paused" json:"is_paused"`
	Paths     *[]types.String   `tfsdk:"paths" json:"paths"`
	Priority  types.Float64     `tfsdk:"priority" json:"priority"`
}

type WebAnalyticsSitesItemsRulesetDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled"`
	ZoneName types.String `tfsdk:"zone_name" json:"zone_name"`
	ZoneTag  types.String `tfsdk:"zone_tag" json:"zone_tag"`
}