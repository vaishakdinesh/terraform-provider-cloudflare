---
page_title: "cloudflare_magic_transit_sites Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_magic_transit_sites (Data Source)



## Example Usage

```terraform
data "cloudflare_magic_transit_sites" "example_magic_transit_sites" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  connector_identifier = "023e105f4ecef8ad9ca31a8372d0c353"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Identifier

### Optional

- `connector_identifier` (String) Identifier
- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `connector_id` (String) Magic Connector identifier tag.
- `description` (String)
- `ha_mode` (Boolean) Site high availability mode. If set to true, the site can have two connectors and runs in high availability mode.
- `id` (String) Identifier
- `location` (Attributes) Location of site in latitude and longitude. (see [below for nested schema](#nestedatt--result--location))
- `name` (String) The name of the site.
- `secondary_connector_id` (String) Magic Connector identifier tag. Used when high availability mode is on.

<a id="nestedatt--result--location"></a>
### Nested Schema for `result.location`

Read-Only:

- `lat` (String) Latitude
- `lon` (String) Longitude

