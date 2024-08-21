// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_secret

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersSecretsResultListDataSourceEnvelope struct {
	Result *[]*WorkersSecretsResultDataSourceModel `json:"result,computed"`
}

type WorkersSecretsDataSourceModel struct {
	AccountID         types.String                            `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String                            `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	ScriptName        types.String                            `tfsdk:"script_name" path:"script_name"`
	MaxItems          types.Int64                             `tfsdk:"max_items"`
	Result            *[]*WorkersSecretsResultDataSourceModel `tfsdk:"result"`
}

func (m *WorkersSecretsDataSourceModel) toListParams() (params workers_for_platforms.DispatchNamespaceScriptSecretListParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceScriptSecretListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersSecretsResultDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	Type types.String `tfsdk:"type" json:"type"`
}
