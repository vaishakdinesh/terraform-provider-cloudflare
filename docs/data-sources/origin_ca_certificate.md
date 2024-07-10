---
page_title: "cloudflare_origin_ca_certificate Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_origin_ca_certificate (Data Source)



## Example Usage

```terraform
data "cloudflare_origin_ca_certificate" "example" {
  id = "REPLACE_ME"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `certificate` (String) The Origin CA certificate. Will be newline-encoded.
- `certificate_id` (String) Identifier
- `csr` (String) The Certificate Signing Request (CSR). Must be newline-encoded.
- `expires_on` (String) When the certificate will expire.
- `find_one_by` (Attributes) (see [below for nested schema](#nestedatt--find_one_by))
- `hostnames` (String) Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.
- `id` (String) Identifier
- `request_type` (String) Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa), or "keyless-certificate" (for Keyless SSL servers).
- `requested_validity` (Number) The number of days for which the certificate should be valid.

<a id="nestedatt--find_one_by"></a>
### Nested Schema for `find_one_by`

Optional:

- `zone_id` (String) Identifier

