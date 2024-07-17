---
page_title: "cloudflare_mtls_certificate Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_mtls_certificate (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Identifier
- `ca` (Boolean) Indicates whether the certificate is a CA or leaf certificate.
- `certificates` (String) The uploaded root CA certificate.
- `find_one_by` (Attributes) (see [below for nested schema](#nestedatt--find_one_by))
- `mtls_certificate_id` (String) Identifier
- `name` (String) Optional unique name for the certificate. Only used for human readability.
- `uploaded_on` (String) This is the time the certificate was uploaded.

### Read-Only

- `expires_on` (String) When the certificate expires.
- `id` (String) Identifier
- `issuer` (String) The certificate authority that issued the certificate.
- `serial_number` (String) The certificate serial number.
- `signature` (String) The type of hash used for the certificate.

<a id="nestedatt--find_one_by"></a>
### Nested Schema for `find_one_by`

Required:

- `account_id` (String) Identifier

