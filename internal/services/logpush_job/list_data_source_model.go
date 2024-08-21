// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_job

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/logpush"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushJobsResultListDataSourceEnvelope struct {
	Result *[]*LogpushJobsResultDataSourceModel `json:"result,computed"`
}

type LogpushJobsDataSourceModel struct {
	AccountID types.String                         `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                         `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                          `tfsdk:"max_items"`
	Result    *[]*LogpushJobsResultDataSourceModel `tfsdk:"result"`
}

func (m *LogpushJobsDataSourceModel) toListParams() (params logpush.JobListParams, diags diag.Diagnostics) {
	params = logpush.JobListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type LogpushJobsResultDataSourceModel struct {
	ID                       types.Int64                              `tfsdk:"id" json:"id"`
	Dataset                  types.String                             `tfsdk:"dataset" json:"dataset"`
	DestinationConf          types.String                             `tfsdk:"destination_conf" json:"destination_conf"`
	Enabled                  types.Bool                               `tfsdk:"enabled" json:"enabled"`
	ErrorMessage             timetypes.RFC3339                        `tfsdk:"error_message" json:"error_message"`
	Frequency                types.String                             `tfsdk:"frequency" json:"frequency,computed"`
	Kind                     types.String                             `tfsdk:"kind" json:"kind"`
	LastComplete             timetypes.RFC3339                        `tfsdk:"last_complete" json:"last_complete"`
	LastError                timetypes.RFC3339                        `tfsdk:"last_error" json:"last_error"`
	LogpullOptions           types.String                             `tfsdk:"logpull_options" json:"logpull_options"`
	MaxUploadBytes           types.Int64                              `tfsdk:"max_upload_bytes" json:"max_upload_bytes"`
	MaxUploadIntervalSeconds types.Int64                              `tfsdk:"max_upload_interval_seconds" json:"max_upload_interval_seconds,computed"`
	MaxUploadRecords         types.Int64                              `tfsdk:"max_upload_records" json:"max_upload_records,computed"`
	Name                     types.String                             `tfsdk:"name" json:"name"`
	OutputOptions            *LogpushJobsOutputOptionsDataSourceModel `tfsdk:"output_options" json:"output_options"`
}

type LogpushJobsOutputOptionsDataSourceModel struct {
	BatchPrefix     types.String    `tfsdk:"batch_prefix" json:"batch_prefix,computed"`
	BatchSuffix     types.String    `tfsdk:"batch_suffix" json:"batch_suffix,computed"`
	Cve2021_4428    types.Bool      `tfsdk:"cve_2021_4428" json:"CVE-2021-4428,computed"`
	FieldDelimiter  types.String    `tfsdk:"field_delimiter" json:"field_delimiter,computed"`
	FieldNames      *[]types.String `tfsdk:"field_names" json:"field_names"`
	OutputType      types.String    `tfsdk:"output_type" json:"output_type,computed"`
	RecordDelimiter types.String    `tfsdk:"record_delimiter" json:"record_delimiter,computed"`
	RecordPrefix    types.String    `tfsdk:"record_prefix" json:"record_prefix,computed"`
	RecordSuffix    types.String    `tfsdk:"record_suffix" json:"record_suffix,computed"`
	RecordTemplate  types.String    `tfsdk:"record_template" json:"record_template,computed"`
	SampleRate      types.Float64   `tfsdk:"sample_rate" json:"sample_rate,computed"`
	TimestampFormat types.String    `tfsdk:"timestamp_format" json:"timestamp_format,computed"`
}
