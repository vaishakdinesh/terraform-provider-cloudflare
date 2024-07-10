---
page_title: "cloudflare_origin_ca_certificate Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_origin_ca_certificate (Resource)



~> Since [v3.32.0](https://github.com/cloudflare/terraform-provider-cloudflare/releases/tag/v3.32.0)
   all authentication schemes are supported for managing Origin CA certificates.
   Versions prior to v3.32.0 will still need to use [`api_user_service_key`](../index.html#api_user_service_key).

## Example Usage

```terraform
resource "tls_private_key" "example" {
  algorithm = "RSA"
}

resource "tls_cert_request" "example" {
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = ""
    organization = "Terraform Test"
  }
}

resource "cloudflare_origin_ca_certificate" "example" {
  csr                = tls_cert_request.example.cert_request_pem
  hostnames          = ["example.com"]
  request_type       = "origin-rsa"
  requested_validity = 7
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `certificate_id` (String) Identifier
- `csr` (String) The Certificate Signing Request (CSR). Must be newline-encoded.
- `hostnames` (List of String) Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.
- `request_type` (String) Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa), or "keyless-certificate" (for Keyless SSL servers).
- `requested_validity` (Number) The number of days for which the certificate should be valid.

### Read-Only

- `id` (String) Identifier

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_origin_ca_certificate.example <certificate_id>
```