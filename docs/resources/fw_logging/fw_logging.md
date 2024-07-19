---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_logging Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_logging: Bind a logging template to firewall
  PLACEHOLDER
---

# thunder_fw_logging (Resource)

`thunder_fw_logging`: Bind a logging template to firewall

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_logging" "thunder_fw_logging" {
  gtp {
    sampling_enable {
      counters1 = "all"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `gtp` (Block List, Max: 1) (see [below for nested schema](#nestedblock--gtp))
- `name` (String) Logging Template Name
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--gtp"></a>
### Nested Schema for `gtp`

Optional:

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--gtp--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--gtp--sampling_enable"></a>
### Nested Schema for `gtp.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'log_type_gtp_invalid_teid': Log Event Type GTP Invalid TEID; 'log_gtp_type_reserved_ie_present': Log Event Type GTP Reserved IE Present; 'log_type_gtp_mandatory_ie_missing': Log Event Type GTP Mandatory IE Missing; 'log_type_gtp_mandatory_ie_inside_grouped_ie_missing': Log Event Type GTP Mandatory IE Missing Inside Grouped IE; 'log_type_gtp_msisdn_filtering': Log Event Type GTP MSISDN Filtering; 'log_type_gtp_out_of_order_ie': Log Event Type GTP Out of Order IE V1; 'log_type_gtp_out_of_state_ie': Log Event Type GTP Out of State IE; 'log_type_enduser_ip_spoofed': Log Event Type GTP Enduser IP Spoofed; 'log_type_crosslayer_correlation': Log Event GTP Crosslayer Correlation; 'log_type_message_not_supported': Log Event GTP Reserved Message Found; 'log_type_out_of_state': Log Event GTP Out of State Message; 'log_type_max_msg_length': Log Event GTP Message Length Exceeded Max; 'log_type_gtp_message_filtering': Log Event Type GTP Message Filtering; 'log_type_gtp_apn_filtering': Log Event Type GTP Apn Filtering; 'log_type_gtp_rat_type_filtering': Log Event GTP RAT Type Filtering; 'log_type_country_code_mismatch': Log Event GTP Country Code Mismatch; 'log_type_gtp_in_gtp_filtering': Log Event GTP in GTP Filtering; 'log_type_gtp_node_restart': Log Event GTP SGW/PGW restarted; 'log_type_gtp_seq_num_mismatch': Log Event GTP Response Sequence number Mismatch; 'log_type_gtp_rate_limit_periodic': Log Event GTP Rate Limit Periodic; 'log_type_gtp_invalid_message_length': Log Event GTP Invalid message length across layers; 'log_type_gtp_hdr_invalid_protocol_flag': Log Event GTP Protocol flag in header; 'log_type_gtp_hdr_invalid_spare_bits': Log Event GTP invalid spare bits in header; 'log_type_gtp_hdr_invalid_piggy_flag': Log Event GTP invalid piggyback flag in header; 'log_type_gtp_invalid_version': Log Event invalid GTP version; 'log_type_gtp_invalid_ports': Log Event mismatch of GTP message and ports;



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'log_message_sent': Log Packet Sent; 'log_type_reset': Log Event Type Reset; 'log_type_deny': Log Event Type Deny; 'log_type_session_closed': Log Event Type Session Close; 'log_type_session_opened': Log Event Type Session Open; 'rule_not_logged': Firewall Rule Not Logged; 'log-dropped': Log Packets Dropped; 'tcp-session-created': TCP Session Created; 'tcp-session-deleted': TCP Session Deleted; 'udp-session-created': UDP Session Created; 'udp-session-deleted': UDP Session Deleted; 'icmp-session-deleted': ICMP Session Deleted; 'icmp-session-created': ICMP Session Created; 'icmpv6-session-deleted': ICMPV6 Session Deleted; 'icmpv6-session-created': ICMPV6 Session Created; 'other-session-deleted': Other Session Deleted; 'other-session-created': Other Session Created; 'http-request-logged': HTTP Request Logged; 'http-logging-invalid-format': HTTP Logging Invalid Format Error; 'dcmsg_permit': Dcmsg Permit; 'alg_override_permit': Alg Override Permit; 'template_error': Template Error; 'ipv4-frag-applied': IPv4 Fragmentation Applied; 'ipv4-frag-failed': IPv4 Fragmentation Failed; 'ipv6-frag-applied': IPv6 Fragmentation Applied; 'ipv6-frag-failed': IPv6 Fragmentation Failed; 'out-of-buffers': Out of Buffers; 'add-msg-failed': Add Message to Buffer Failed; 'tcp-logging-conn-established': TCP Logging Conn Established; 'tcp-logging-conn-create-failed': TCP Logging Conn Create Failed; 'tcp-logging-conn-dropped': TCP Logging Conn Dropped; 'log-message-too-long': Log message too long; 'http-out-of-order-dropped': HTTP out-of-order dropped; 'http-alloc-failed': HTTP Request Info Allocation Failed; 'sctp-session-created': SCTP Session Created; 'sctp-session-deleted': SCTP Session Deleted; 'log_type_sctp_inner_proto_filter': Log Event Type SCTP Inner Proto Filter; 'tcp-logging-port-allocated': TCP Logging Port Allocated; 'tcp-logging-port-freed': TCP Logging Port Freed; 'tcp-logging-port-allocation-failed': TCP Logging Port Allocation Failed; 'iddos-blackhole-entry-create': iDDoS IP Entry Created; 'iddos-blackhole-entry-delete': iDDoS IP Entry Deleted; 'session-limit-exceeded': Session Limit Exceeded;

