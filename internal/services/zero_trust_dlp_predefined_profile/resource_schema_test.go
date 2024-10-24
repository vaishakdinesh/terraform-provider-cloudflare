// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_predefined_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZeroTrustDLPPredefinedProfileModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_dlp_predefined_profile.ZeroTrustDLPPredefinedProfileModel)(nil)
	schema := zero_trust_dlp_predefined_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}