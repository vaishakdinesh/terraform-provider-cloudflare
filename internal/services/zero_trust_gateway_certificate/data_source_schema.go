// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewayCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"certificate_id": schema.StringAttribute{
				Description: "Certificate UUID tag.",
				Optional:    true,
			},
			"binding_status": schema.StringAttribute{
				Description: "The deployment status of the certificate on Cloudflare's edge. Certificates in the 'available' (previously called 'active') state may be used for Gateway TLS interception.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending_deployment",
						"available",
						"pending_deletion",
						"inactive",
					),
				},
			},
			"certificate": schema.StringAttribute{
				Description: "The CA certificate",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"expires_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"fingerprint": schema.StringAttribute{
				Description: "The SHA256 fingerprint of the certificate.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Certificate UUID tag.",
				Computed:    true,
			},
			"in_use": schema.BoolAttribute{
				Description: "Use this certificate for Gateway TLS interception",
				Computed:    true,
			},
			"issuer_org": schema.StringAttribute{
				Description: "The organization that issued the certificate.",
				Computed:    true,
			},
			"issuer_raw": schema.StringAttribute{
				Description: "The entire issuer field of the certificate.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of certificate, either BYO-PKI (custom) or Gateway-managed.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("custom", "gateway_managed"),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"uploaded_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustGatewayCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustGatewayCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("certificate_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("certificate_id")),
	}
}