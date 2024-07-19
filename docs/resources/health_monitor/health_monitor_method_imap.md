---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_health_monitor_method_imap Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_health_monitor_method_imap: IMAP type
  PLACEHOLDER
---

# thunder_health_monitor_method_imap (Resource)

`thunder_health_monitor_method_imap`: IMAP type

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_imap" "thunder_health_monitor_method_imap" {

  name                 = "tf_test"
  imap                 = 1
  imap_cram_md5        = 1
  imap_login           = 1
  imap_password        = 1
  imap_password_string = "a10net"
  imap_plain           = 1
  imap_port            = 14321
  imap_username        = "a11"
  pwd_auth             = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `imap` (Number) IMAP type
- `imap_cram_md5` (Number) Challenge-response authentication mechanism
- `imap_login` (Number) Simple login
- `imap_password` (Number) Specify the user password
- `imap_password_string` (String) Configure password, '' means empty password
- `imap_plain` (Number) Plain text
- `imap_port` (Number) Specify the IMAP port, default is 143 (Port Number (default 143))
- `imap_username` (String) Specify the username
- `pwd_auth` (Number) Specify the Authentication method
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

