// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type TurnstileWidgetDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &TurnstileWidgetDataSource{}

func NewTurnstileWidgetDataSource() datasource.DataSource {
	return &TurnstileWidgetDataSource{}
}

func (d *TurnstileWidgetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_turnstile_widget"
}

func (d *TurnstileWidgetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *TurnstileWidgetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TurnstileWidgetDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Filter == nil {
		params, diags := data.toReadParams()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		res := new(http.Response)
		env := TurnstileWidgetResultDataSourceEnvelope{*data}
		_, err := d.client.Challenges.Widgets.Get(
			ctx,
			data.Sitekey.ValueString(),
			params,
			option.WithResponseBodyInto(&res),
			option.WithMiddleware(logging.Middleware(ctx)),
		)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}
		bytes, _ := io.ReadAll(res.Body)
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
			return
		}
		data = &env.Result
	} else {
		params, diags := data.toListParams()
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		items := &[]*TurnstileWidgetDataSourceModel{}
		env := TurnstileWidgetResultListDataSourceEnvelope{items}

		page, err := d.client.Challenges.Widgets.List(ctx, params)
		if err != nil {
			resp.Diagnostics.AddError("failed to make http request", err.Error())
			return
		}

		bytes := []byte(page.JSON.RawJSON())
		err = apijson.Unmarshal(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}

		if count := len(*items); count != 1 {
			resp.Diagnostics.AddError("failed to find exactly one result", fmt.Sprint(count)+" found")
			return
		}
		data = (*items)[0]
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
