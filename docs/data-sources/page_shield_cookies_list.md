---
page_title: "cloudflare_page_shield_cookies_list Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_page_shield_cookies_list (Data Source)



## Example Usage

```terraform
data "cloudflare_page_shield_cookies_list" "example_page_shield_cookies_list" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  domain = "example.com"
  export = "csv"
  hosts = "blog.cloudflare.com,www.example*,*cloudflare.com"
  http_only = true
  name = "session_id"
  order_by = "first_seen_at"
  page = "2"
  page_url = "example.com/page,*/checkout,example.com/*,*checkout*"
  path = "/"
  per_page = 100
  same_site = "lax"
  secure = true
  type = "first_party"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) Identifier

### Optional

- `direction` (String) The direction used to sort returned cookies.'
- `domain` (String) Filters the returned cookies that match the specified domain attribute
- `export` (String) Export the list of cookies as a file.
- `hosts` (String) Includes cookies that match one or more URL-encoded hostnames separated by commas.

Wildcards are supported at the start and end of each hostname to support starts with, ends with
and contains. If no wildcards are used, results will be filtered by exact match
- `http_only` (Boolean) Filters the returned cookies that are set with HttpOnly
- `max_items` (Number) Max items to fetch, default: 1000
- `name` (String) Filters the returned cookies that match the specified name.
Wildcards are supported at the start and end to support starts with, ends with
and contains. e.g. session*
- `order_by` (String) The field used to sort returned cookies.
- `page` (String) The current page number of the paginated results.

We additionally support a special value "all". When "all" is used, the API will return all the cookies
with the applied filters in a single page. This feature is best-effort and it may only work for zones with 
a low number of cookies
- `page_url` (String) Includes connections that match one or more page URLs (separated by commas) where they were last seen

Wildcards are supported at the start and end of each page URL to support starts with, ends with
and contains. If no wildcards are used, results will be filtered by exact match
- `path` (String) Filters the returned cookies that match the specified path attribute
- `per_page` (Number) The number of results per page.
- `same_site` (String) Filters the returned cookies that match the specified same_site attribute
- `secure` (Boolean) Filters the returned cookies that are set with Secure
- `type` (String) Filters the returned cookies that match the specified type attribute

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `domain_attribute` (String)
- `expires_attribute` (String)
- `first_seen_at` (String)
- `host` (String)
- `http_only_attribute` (Boolean)
- `id` (String) Identifier
- `last_seen_at` (String)
- `max_age_attribute` (Number)
- `name` (String)
- `page_urls` (List of String)
- `path_attribute` (String)
- `same_site_attribute` (String)
- `secure_attribute` (Boolean)
- `type` (String)

