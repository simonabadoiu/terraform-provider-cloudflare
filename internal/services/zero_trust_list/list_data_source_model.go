// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustListsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustListsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustListsDataSourceModel struct {
	AccountID types.String                                                      `tfsdk:"account_id" path:"account_id"`
	Type      types.String                                                      `tfsdk:"type" query:"type"`
	MaxItems  types.Int64                                                       `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustListsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustListsDataSourceModel) toListParams() (params zero_trust.GatewayListListParams, diags diag.Diagnostics) {
	params = zero_trust.GatewayListListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Type.IsNull() {
		params.Type = cloudflare.F(zero_trust.GatewayListListParamsType(m.Type.ValueString()))
	}

	return
}

type ZeroTrustListsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed_optional"`
	ListCount   types.Float64     `tfsdk:"list_count" json:"count,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String      `tfsdk:"description" json:"description,computed_optional"`
	Name        types.String      `tfsdk:"name" json:"name,computed_optional"`
	Type        types.String      `tfsdk:"type" json:"type,computed_optional"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}
