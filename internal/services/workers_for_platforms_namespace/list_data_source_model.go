// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsNamespacesResultListDataSourceEnvelope struct {
	Result *[]*WorkersForPlatformsNamespacesItemsDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespacesDataSourceModel struct {
	AccountID types.String                                          `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                           `tfsdk:"max_items"`
	Items     *[]*WorkersForPlatformsNamespacesItemsDataSourceModel `tfsdk:"items"`
}

type WorkersForPlatformsNamespacesItemsDataSourceModel struct {
	CreatedBy     types.String      `tfsdk:"created_by" json:"created_by"`
	CreatedOn     timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedBy    types.String      `tfsdk:"modified_by" json:"modified_by"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed"`
	NamespaceID   types.String      `tfsdk:"namespace_id" json:"namespace_id"`
	NamespaceName types.String      `tfsdk:"namespace_name" json:"namespace_name"`
	ScriptCount   types.Int64       `tfsdk:"script_count" json:"script_count"`
}