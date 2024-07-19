---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_rate: Configure stats to triggers packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl_trigger_stats_rate" {

  name                       = "test"
  duration                   = 60
  krb_other_error            = 1
  krb_pw_change_failure      = 1
  krb_pw_expiry              = 1
  krb_timeout_error          = 1
  krb_validate_kdc_failure   = 1
  ntlm_auth_failure          = 1
  ntlm_other_error           = 1
  ntlm_prepare_req_error     = 1
  ntlm_proto_nego_failure    = 1
  ntlm_session_setup_failure = 1
  ntlm_timeout_error         = 1
  threshold_exceeded_by      = 5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `krb_other_error` (Number) Enable automatic packet-capture for Kerberos Other Error
- `krb_pw_change_failure` (Number) Enable automatic packet-capture for Kerberos password change failure
- `krb_pw_expiry` (Number) Enable automatic packet-capture for Kerberos password expiry
- `krb_timeout_error` (Number) Enable automatic packet-capture for Kerberos Timeout
- `krb_validate_kdc_failure` (Number) Enable automatic packet-capture for Kerberos KDC Validation Failure
- `ntlm_auth_failure` (Number) Enable automatic packet-capture for NTLM Authentication Failure
- `ntlm_other_error` (Number) Enable automatic packet-capture for NTLM Other Error
- `ntlm_prepare_req_error` (Number) Enable automatic packet-capture for NTLM Prepare Request Error
- `ntlm_proto_nego_failure` (Number) Enable automatic packet-capture for NTLM Protocol Negotiation Failure
- `ntlm_session_setup_failure` (Number) Enable automatic packet-capture for NTLM Session Setup Failure
- `ntlm_timeout_error` (Number) Enable automatic packet-capture for NTLM Timeout
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

