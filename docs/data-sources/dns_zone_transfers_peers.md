---
page_title: "cloudflare_dns_zone_transfers_peers Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_dns_zone_transfers_peers (Data Source)



## Example Usage

```terraform
data "cloudflare_dns_zone_transfers_peers" "example_dns_zone_transfers_peers" {
  account_id = "01a7362d577a6c3019a474fd6f485823"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)

### Optional

- `max_items` (Number) Max items to fetch, default: 1000

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `id` (String)
- `ip` (String) IPv4/IPv6 address of primary or secondary nameserver, depending on what zone this peer is linked to. For primary zones this IP defines the IP of the secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary zones this IP defines the IP of the primary nameserver Cloudflare will send AXFR/IXFR requests to.
- `ixfr_enable` (Boolean) Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary zones.
- `name` (String) The name of the peer.
- `port` (Number) DNS port of primary or secondary nameserver, depending on what zone this peer is linked to.
- `tsig_id` (String) TSIG authentication will be used for zone transfer if configured.

