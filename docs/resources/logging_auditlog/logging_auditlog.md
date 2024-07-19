---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_logging_auditlog Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_logging_auditlog: Set send auditlog to syslog host
  PLACEHOLDER
---

# thunder_logging_auditlog (Resource)

`thunder_logging_auditlog`: Set send auditlog to syslog host

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_auditlog" "thunder_logging_auditlog" {
  audit_facility = "local0"
  host4          = "15"
  port           = 514
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `audit_facility` (String) 'local0': Local use; 'local1': Local use; 'local2': Local use; 'local3': Local use; 'local4': Local use; 'local5': Local use; 'local6': Local use; 'local7': Local use;  (Configure the facility of auditlog)
- `host4` (String) Configure the auditlog host
- `host6` (String) Configure the auditlog host
- `partition_name` (String) Select partition name for logging
- `port` (Number) Set remote audit log port number(Default 514)
- `shared` (Number) Select shared partition
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

