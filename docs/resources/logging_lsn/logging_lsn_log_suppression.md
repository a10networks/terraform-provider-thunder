---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_logging_lsn_log_suppression Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_logging_lsn_log_suppression: Set LSN system log generation parameters
  PLACEHOLDER
---

# thunder_logging_lsn_log_suppression (Resource)

`thunder_logging_lsn_log_suppression`: Set LSN system log generation parameters

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_lsn_log_suppression" "thunder_logging_lsn_log_suppression" {

  count1 = 50
  time   = 20

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `count1` (Number) Configure log suppression count (default: 100)
- `time` (Number) Log generation timeout(default: 30 seconds)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

