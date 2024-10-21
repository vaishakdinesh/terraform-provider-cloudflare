---
page_title: "cloudflare_zero_trust_access_applications Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_access_applications (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `max_items` (Number) Max items to fetch, default: 1000
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

