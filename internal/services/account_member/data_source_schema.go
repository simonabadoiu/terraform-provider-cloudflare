// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &AccountMemberDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"member_id": schema.StringAttribute{
				Description: "Membership identifier tag.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Membership identifier tag.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "A member's status in the account.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
				},
			},
			"user": schema.SingleNestedAttribute{
				Description: "Details of the user associated to the membership.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AccountMemberUserDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description: "The contact email address of the user.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
					"first_name": schema.StringAttribute{
						Description: "User's first name",
						Computed:    true,
						Optional:    true,
					},
					"last_name": schema.StringAttribute{
						Description: "User's last name",
						Computed:    true,
						Optional:    true,
					},
					"two_factor_authentication_enabled": schema.BoolAttribute{
						Description: "Indicates whether two-factor authentication is enabled for the user account. Does not apply to API authentication.",
						Computed:    true,
					},
				},
			},
			"policies": schema.ListNestedAttribute{
				Description: "Access policy for the membership",
				Computed:    true,
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Policy identifier.",
							Computed:    true,
						},
						"access": schema.StringAttribute{
							Description: "Allow or deny operations against the resources.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("allow", "deny"),
							},
						},
						"permission_groups": schema.ListNestedAttribute{
							Description: "A set of permission groups that are specified to the policy.",
							Computed:    true,
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Identifier of the group.",
										Computed:    true,
									},
									"meta": schema.StringAttribute{
										Description: "Attributes associated to the permission group.",
										Computed:    true,
										Optional:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"name": schema.StringAttribute{
										Description: "Name of the group.",
										Computed:    true,
									},
								},
							},
						},
						"resource_groups": schema.ListNestedAttribute{
							Description: "A list of resource groups that the policy applies to.",
							Computed:    true,
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Identifier of the group.",
										Computed:    true,
									},
									"scope": schema.ListNestedAttribute{
										Description: "The scope associated to the resource group",
										Computed:    true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
													Computed:    true,
												},
												"objects": schema.ListNestedAttribute{
													Description: "A list of scope objects for additional context.",
													Computed:    true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description: "This is a combination of pre-defined resource name and identifier (like Zone ID etc.)",
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"meta": schema.StringAttribute{
										Description: "Attributes associated to the resource group.",
										Computed:    true,
										Optional:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"name": schema.StringAttribute{
										Description: "Name of the resource group.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"roles": schema.ListNestedAttribute{
				Description: "Roles assigned to this Member.",
				Computed:    true,
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Role identifier tag.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Description of role's permissions.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Role Name.",
							Computed:    true,
						},
						"permissions": schema.ListAttribute{
							Description: "Access permissions for this User.",
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
					"direction": schema.StringAttribute{
						Description: "Direction to order results.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Field to order results by.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"user.first_name",
								"user.last_name",
								"user.email",
								"status",
							),
						},
					},
					"status": schema.StringAttribute{
						Description: "A member's status in the account.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"accepted",
								"pending",
								"rejected",
							),
						},
					},
				},
			},
		},
	}
}

func (d *AccountMemberDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccountMemberDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
