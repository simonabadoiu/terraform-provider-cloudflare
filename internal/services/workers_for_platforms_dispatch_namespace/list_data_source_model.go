// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsDispatchNamespacesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersForPlatformsDispatchNamespacesResultDataSourceModel] `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespacesDataSourceModel struct {
	AccountID types.String                                                                             `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkersForPlatformsDispatchNamespacesResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersForPlatformsDispatchNamespacesDataSourceModel) toListParams() (params workers_for_platforms.DispatchNamespaceListParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersForPlatformsDispatchNamespacesResultDataSourceModel struct {
	CreatedBy     types.String      `tfsdk:"created_by" json:"created_by,computed_optional"`
	CreatedOn     timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedBy    types.String      `tfsdk:"modified_by" json:"modified_by,computed_optional"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed"`
	NamespaceID   types.String      `tfsdk:"namespace_id" json:"namespace_id,computed_optional"`
	NamespaceName types.String      `tfsdk:"namespace_name" json:"namespace_name,computed_optional"`
	ScriptCount   types.Int64       `tfsdk:"script_count" json:"script_count,computed_optional"`
}
