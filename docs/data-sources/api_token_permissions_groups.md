---
page_title: "cloudflare_api_token_permissions_groups Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_api_token_permissions_groups (Data Source)



## Example Usage

```terraform
data "cloudflare_api_token_permissions_groups" "example_api_token_permissions_groups" {
  account_id = "eb78d65290b24279ba6f44721b3ea3c4"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account identifier tag.

### Read-Only

- `id` (String) Public ID.
- `name` (String) Permission Group Name
- `scopes` (List of String) Resources to which the Permission Group is scoped


