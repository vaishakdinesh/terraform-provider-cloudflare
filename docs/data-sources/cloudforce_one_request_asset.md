---
page_title: "cloudflare_cloudforce_one_request_asset Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_cloudforce_one_request_asset (Data Source)



## Example Usage

```terraform
data "cloudflare_cloudforce_one_request_asset" "example_cloudforce_one_request_asset" {
  account_identifier = "023e105f4ecef8ad9ca31a8372d0c353"
  request_identifier = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  asset_identifer = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_identifier` (String) Identifier
- `asset_identifer` (String) UUID
- `request_identifier` (String) UUID

### Read-Only

- `created` (String) Asset creation time
- `description` (String) Asset description
- `file_type` (String) Asset file type
- `id` (Number) Asset ID
- `name` (String) Asset name


