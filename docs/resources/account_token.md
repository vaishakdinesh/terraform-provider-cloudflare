---
page_title: "cloudflare_account_token Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_account_token (Resource)



## Example Usage

```terraform
resource "cloudflare_account_token" "example_account_token" {
  account_id = "eb78d65290b24279ba6f44721b3ea3c4"
  name = "readonly token"
  policies = [{
    id = "f267e341f3dd4697bd3b9f71dd96247f"
    effect = "allow"
    permission_groups = [{
      id = "c8fed203ed3043cba015a93ad1616f1f"
      meta = {
        key = "key"
        value = "value"
      }
      name = "Zone Read"
    }, {
      id = "82e64a83756745bbbb1c9c2701bf816b"
      meta = {
        key = "key"
        value = "value"
      }
      name = "Magic Network Monitoring"
    }]
    resources = {
      "com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43" = "*"
      "com.cloudflare.api.account.zone.eb78d65290b24279ba6f44721b3ea3c4" = "*"
    }
  }]
  condition = {
    request_ip = {
      in = ["123.123.123.0/24", "2606:4700::/32"]
      not_in = ["123.123.123.100/24", "2606:4700:4700::/48"]
    }
  }
  expires_on = "2020-01-01T00:00:00Z"
  not_before = "2018-07-01T05:20:00Z"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Account identifier tag.
- `name` (String) Token name.
- `policies` (Attributes List) List of access policies assigned to the token. (see [below for nested schema](#nestedatt--policies))

### Optional

- `condition` (Attributes) (see [below for nested schema](#nestedatt--condition))
- `expires_on` (String) The expiration time on or after which the JWT MUST NOT be accepted for processing.
- `not_before` (String) The time before which the token MUST NOT be accepted for processing.
- `status` (String) Status of the token.

### Read-Only

- `id` (String) Token identifier tag.
- `issued_on` (String) The time on which the token was created.
- `last_used_on` (String) Last time the token was used.
- `modified_on` (String) Last time the token was modified.
- `value` (String) The token value.

<a id="nestedatt--policies"></a>
### Nested Schema for `policies`

Required:

- `effect` (String) Allow or deny operations against the resources.
- `permission_groups` (Attributes List) A set of permission groups that are specified to the policy. (see [below for nested schema](#nestedatt--policies--permission_groups))
- `resources` (Map of String) A list of resource names that the policy applies to.

Read-Only:

- `id` (String) Policy identifier.

<a id="nestedatt--policies--permission_groups"></a>
### Nested Schema for `policies.permission_groups`

Optional:

- `meta` (Attributes) Attributes associated to the permission group. (see [below for nested schema](#nestedatt--policies--permission_groups--meta))

Read-Only:

- `id` (String) Identifier of the group.
- `name` (String) Name of the group.

<a id="nestedatt--policies--permission_groups--meta"></a>
### Nested Schema for `policies.permission_groups.meta`

Optional:

- `key` (String)
- `value` (String)




<a id="nestedatt--condition"></a>
### Nested Schema for `condition`

Optional:

- `request_ip` (Attributes) Client IP restrictions. (see [below for nested schema](#nestedatt--condition--request_ip))

<a id="nestedatt--condition--request_ip"></a>
### Nested Schema for `condition.request_ip`

Optional:

- `in` (List of String) List of IPv4/IPv6 CIDR addresses.
- `not_in` (List of String) List of IPv4/IPv6 CIDR addresses.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_account_token.example '<account_id>/<token_id>'
```