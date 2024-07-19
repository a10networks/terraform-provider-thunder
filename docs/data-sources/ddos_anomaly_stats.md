---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_anomaly_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_anomaly_stats: Statistics for the object anomaly
  PLACEHOLDER
---

# thunder_ddos_anomaly_stats (Data Source)

`thunder_ddos_anomaly_stats`: Statistics for the object anomaly

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_anomaly_stats" "thunder_ddos_anomaly_stats" {
}
output "get_ddos_anomaly_stats" {
  value = ["${data.thunder_ddos_anomaly_stats.thunder_ddos_anomaly_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `bad_ip_csum` (Number) IPv4 Invalid Checksum
- `bad_ip_flags` (Number) IPv4 Invalid Flags
- `bad_ip_frag_off` (Number) IPv4 Invalid Fragment Offset
- `bad_ip_hdr_len` (Number) IPv4 Invalid Header Length
- `bad_ip_pl_len` (Number) IP Invalid Payload Length
- `bad_ip_ttl` (Number) IP Invalid TTL
- `empty_frag` (Number) IPv4 Empty Fragment
- `icmp_pod` (Number) ICMP Ping Of Death
- `ip_frag` (Number) IP Fragment
- `ipv4_opt` (Number) IPv4 Option
- `land_attack` (Number) IP Land Attack
- `micro_frag` (Number) IPv4 Micro Fragment
- `no_ip_payload` (Number) IP Payload None
- `oversize_ip_pl` (Number) IP Payload Too Large
- `runt_ip_hdr` (Number) IP Runt Header
- `runt_tcpudp_hdr` (Number) TCPUDP Runt Header
- `tcp_bad_csum` (Number) TCP Invalid Checksum
- `tcp_bad_ip_len` (Number) TCP Invalid IPv4 Length
- `tcp_bad_urg_off` (Number) TCP Invalid Urgent Offset
- `tcp_frag_header` (Number) TCP Fragment Header
- `tcp_null_flags` (Number) TCP Null Flags
- `tcp_null_scan` (Number) TCP Null Scan
- `tcp_opt_overflow` (Number) TCP Option Error
- `tcp_short_hdr` (Number) TCP Short Header
- `tcp_syn_fin` (Number) TCP SYN&FIN
- `tcp_syn_frag` (Number) TCP SYN Fragment
- `tcp_xmas_flags` (Number) TCP XMAS Flags
- `tcp_xmas_scan` (Number) TCP XMAS Scan
- `tun_mismatch` (Number) IP Tunnel Mismatch
- `udp_bad_csum` (Number) UDP Invalid Checksum
- `udp_kb_frag` (Number) UDP KB Fragment
- `udp_port_lb` (Number) UDP Port LB
- `udp_short_hdr` (Number) UDP Short Header
- `udp_short_leng` (Number) UDP Invalid Length

