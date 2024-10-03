---
page_title: "cloudflare_zero_trust_device_posture_rule Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_device_posture_rule (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String)
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `rule_id` (String) API UUID.

### Read-Only

- `description` (String) The description of the device posture rule.
- `expiration` (String) Sets the expiration time for a posture check result. If empty, the result remains valid until it is overwritten by new data from the WARP client.
- `id` (String) API UUID.
- `input` (Attributes) The value to be checked against. (see [below for nested schema](#nestedatt--input))
- `match` (Attributes List) The conditions that the client must match to run the rule. (see [below for nested schema](#nestedatt--match))
- `name` (String) The name of the device posture rule.
- `schedule` (String) Polling frequency for the WARP client posture check. Default: `5m` (poll every five minutes). Minimum: `1m`.
- `type` (String) The type of device posture rule.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String)


<a id="nestedatt--input"></a>
### Nested Schema for `input`

Read-Only:

- `active_threats` (Number) The Number of active threats.
- `certificate_id` (String) UUID of Cloudflare managed certificate.
- `check_disks` (List of String) List of volume names to be checked for encryption.
- `check_private_key` (Boolean) Confirm the certificate was not imported from another device. We recommend keeping this enabled unless the certificate was deployed without a private key.
- `cn` (String) Common Name that is protected by the certificate
- `compliance_status` (String) Compliance Status
- `connection_id` (String) Posture Integration ID.
- `count_operator` (String) Count Operator
- `domain` (String) Domain
- `eid_last_seen` (String) For more details on eid last seen, refer to the Tanium documentation.
- `enabled` (Boolean) Enabled
- `exists` (Boolean) Whether or not file exists
- `extended_key_usage` (List of String) List of values indicating purposes for which the certificate public key can be used
- `id` (String) List ID.
- `infected` (Boolean) Whether device is infected.
- `is_active` (Boolean) Whether device is active.
- `issue_count` (String) The Number of Issues.
- `last_seen` (String) For more details on last seen, please refer to the Crowdstrike documentation.
- `locations` (Attributes) (see [below for nested schema](#nestedatt--input--locations))
- `network_status` (String) Network status of device.
- `operating_system` (String) Operating system
- `operator` (String) operator
- `os` (String) Os Version
- `os_distro_name` (String) Operating System Distribution Name (linux only)
- `os_distro_revision` (String) Version of OS Distribution (linux only)
- `os_version_extra` (String) Additional version data. For Mac or iOS, the Product Verison Extra. For Linux, the kernel release version. (Mac, iOS, and Linux only)
- `overall` (String) overall
- `path` (String) File path.
- `require_all` (Boolean) Whether to check all disks for encryption.
- `risk_level` (String) For more details on risk level, refer to the Tanium documentation.
- `score` (Number) A value between 0-100 assigned to devices set by the 3rd party posture provider.
- `score_operator` (String) Score Operator
- `sensor_config` (String) SensorConfig
- `sha256` (String) SHA-256.
- `state` (String) For more details on state, please refer to the Crowdstrike documentation.
- `thumbprint` (String) Signing certificate thumbprint.
- `total_score` (Number) For more details on total score, refer to the Tanium documentation.
- `version` (String) Version of OS
- `version_operator` (String) Version Operator

<a id="nestedatt--input--locations"></a>
### Nested Schema for `input.locations`

Read-Only:

- `paths` (List of String) List of paths to check for client certificate on linux.
- `trust_stores` (List of String) List of trust stores to check for client certificate.



<a id="nestedatt--match"></a>
### Nested Schema for `match`

Read-Only:

- `platform` (String)

