// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/certificate_pack"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCertificatePackDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*certificate_pack.CertificatePackDataSourceModel)(nil)
	schema := certificate_pack.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}