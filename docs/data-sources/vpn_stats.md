---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_vpn_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_vpn_stats: Statistics for the object vpn
  PLACEHOLDER
---

# thunder_vpn_stats (Data Source)

`thunder_vpn_stats`: Statistics for the object vpn

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vpn_stats" "thunder_vpn_stats" {

}
output "get_vpn_stats" {
  value = ["${data.thunder_vpn_stats.thunder_vpn_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `error` (Block List, Max: 1) (see [below for nested schema](#nestedblock--error))
- `ike_gateway_list` (Block List) (see [below for nested schema](#nestedblock--ike_gateway_list))
- `ike_stats_global` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ike_stats_global))
- `ipsec_list` (Block List) (see [below for nested schema](#nestedblock--ipsec_list))
- `ipsec_sa_stats_list` (Block List) (see [below for nested schema](#nestedblock--ipsec_sa_stats_list))
- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--error"></a>
### Nested Schema for `error`

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--error--stats))

<a id="nestedblock--error--stats"></a>
### Nested Schema for `error.stats`

Optional:

- `ah_not_supported_with_gcm_gmac_sha2` (Number)
- `bad_auth_type` (Number)
- `bad_checksum` (Number)
- `bad_encrypt_type` (Number)
- `bad_encrypt_type_ctr_gcm` (Number)
- `bad_esp_next_header` (Number)
- `bad_frag_size_configuration` (Number)
- `bad_fragment_size` (Number)
- `bad_gre_header` (Number)
- `bad_gre_protocol` (Number)
- `bad_inline_data` (Number)
- `bad_ip_payload_type` (Number)
- `bad_ip_version` (Number)
- `bad_ipcomp_configuration` (Number)
- `bad_ipsec_auth` (Number)
- `bad_ipsec_context` (Number)
- `bad_ipsec_context_direction` (Number)
- `bad_ipsec_context_flag_mismatch` (Number)
- `bad_ipsec_padding` (Number)
- `bad_ipsec_protocol` (Number)
- `bad_ipsec_spi` (Number)
- `bad_ipsec_unknown` (Number)
- `bad_len` (Number)
- `bad_min_frag_size_auth_sha384_512` (Number)
- `bad_opcode` (Number)
- `bad_selector_match` (Number)
- `bad_sg_write_len` (Number)
- `bad_srtp_auth_tag` (Number)
- `dsiv_incorrect_param` (Number)
- `dummy_payload` (Number)
- `error_ipv6_decrypt_rh_segs_left_error` (Number)
- `error_ipv6_extension_header_bad` (Number)
- `ipcomp_payload` (Number)
- `ipv6_extension_headers_too_big` (Number)
- `ipv6_hop_by_hop_error` (Number)
- `ipv6_outbound_rh_copy_addr_error` (Number)
- `ipv6_rh_length_error` (Number)
- `tfc_padding_with_prefrag_not_supported` (Number)



<a id="nestedblock--ike_gateway_list"></a>
### Nested Schema for `ike_gateway_list`

Required:

- `name` (String) IKE-gateway name

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ike_gateway_list--stats))

<a id="nestedblock--ike_gateway_list--stats"></a>
### Nested Schema for `ike_gateway_list.stats`

Optional:

- `ike_current_version` (Number) IKE version
- `v1_child_sa_invalid_spi` (Number) Invalid SPI for Child SAs
- `v1_in_aggressive_req` (Number) Incoming Aggressive Request
- `v1_in_aggressive_rsp` (Number) Incoming Aggressive Response
- `v1_in_auth_only_req` (Number) Incoming Auth Only Request
- `v1_in_auth_only_rsp` (Number) Incoming Auth Only Response
- `v1_in_id_prot_req` (Number) Incoming ID Protection Request
- `v1_in_id_prot_rsp` (Number) Incoming ID Protection Response
- `v1_in_info_v1_req` (Number) Incoming Info Request
- `v1_in_info_v1_rsp` (Number) Incoming Info Response
- `v1_in_new_group_mode_req` (Number) Incoming New Group Mode Request
- `v1_in_new_group_mode_rsp` (Number) Incoming New Group Mode Response
- `v1_in_quick_mode_req` (Number) Incoming Quick Mode Request
- `v1_in_quick_mode_rsp` (Number) Incoming Quick Mode Response
- `v1_in_transaction_req` (Number) Incoming Transaction Request
- `v1_in_transaction_rsp` (Number) Incoming Transaction Response
- `v1_out_aggressive_req` (Number) Outgoing Aggressive Request
- `v1_out_aggressive_rsp` (Number) Outgoing Aggressive Response
- `v1_out_auth_only_req` (Number) Outgoing Auth Only Request
- `v1_out_auth_only_rsp` (Number) Outgoing Auth Only Response
- `v1_out_id_prot_req` (Number) Outgoing ID Protection Request
- `v1_out_id_prot_rsp` (Number) Outgoing ID Protection Response
- `v1_out_info_v1_req` (Number) Outgoing Info Request
- `v1_out_info_v1_rsp` (Number) Outgoing Info Response
- `v1_out_new_group_mode_req` (Number) Outgoing New Group Mode Request
- `v1_out_new_group_mode_rsp` (Number) Outgoing New Group Mode Response
- `v1_out_quick_mode_req` (Number) Outgoing Quick Mode Request
- `v1_out_quick_mode_rsp` (Number) Outgoing Quick Mode Response
- `v1_out_transaction_req` (Number) Outgoing Transaction Request
- `v1_out_transaction_rsp` (Number) Outgoing Transaction Response
- `v2_child_sa_invalid_spi` (Number) Invalid SPI for Child SAs
- `v2_child_sa_rekey` (Number) Child SA Rekey
- `v2_in_auth_req` (Number) Incoming Auth Request
- `v2_in_auth_rsp` (Number) Incoming Auth Response
- `v2_in_create_child_req` (Number) Incoming Create Child Request
- `v2_in_create_child_rsp` (Number) Incoming Create Child Response
- `v2_in_info_req` (Number) Incoming Info Request
- `v2_in_info_rsp` (Number) Incoming Info Response
- `v2_in_init_req` (Number) Incoming Init Request
- `v2_in_init_rsp` (Number) Incoming Init Response
- `v2_in_invalid` (Number) Incoming Invalid
- `v2_in_invalid_spi` (Number) Incoming Invalid SPI
- `v2_init_rekey` (Number) Initiate Rekey
- `v2_out_auth_req` (Number) Outgoing Auth Request
- `v2_out_auth_rsp` (Number) Outgoing Auth Response
- `v2_out_create_child_req` (Number) Outgoing Create Child Request
- `v2_out_create_child_rsp` (Number) Outgoing Create Child Response
- `v2_out_info_req` (Number) Outgoing Info Request
- `v2_out_info_rsp` (Number) Outgoing Info Response
- `v2_out_init_req` (Number) Outgoing Init Request
- `v2_out_init_rsp` (Number) Outgoing Init Response
- `v2_rsp_rekey` (Number) Respond Rekey



<a id="nestedblock--ike_stats_global"></a>
### Nested Schema for `ike_stats_global`

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ike_stats_global--stats))

<a id="nestedblock--ike_stats_global--stats"></a>
### Nested Schema for `ike_stats_global.stats`

Optional:

- `v1_in_aggressive_req` (Number) Incoming Aggressive Request
- `v1_in_aggressive_rsp` (Number) Incoming Aggressive Response
- `v1_in_auth_only_req` (Number) Incoming Auth Only Request
- `v1_in_auth_only_rsp` (Number) Incoming Auth Only Response
- `v1_in_id_prot_req` (Number) Incoming ID Protection Request
- `v1_in_id_prot_rsp` (Number) Incoming ID Protection Response
- `v1_in_info_v1_req` (Number) Incoming Info Request
- `v1_in_info_v1_rsp` (Number) Incoming Info Response
- `v1_in_new_group_mode_req` (Number) Incoming New Group Mode Request
- `v1_in_new_group_mode_rsp` (Number) Incoming New Group Mode Response
- `v1_in_quick_mode_req` (Number) Incoming Quick Mode Request
- `v1_in_quick_mode_rsp` (Number) Incoming Quick Mode Response
- `v1_in_transaction_req` (Number) Incoming Transaction Request
- `v1_in_transaction_rsp` (Number) Incoming Transaction Response
- `v1_out_aggressive_req` (Number) Outgoing Aggressive Request
- `v1_out_aggressive_rsp` (Number) Outgoing Aggressive Response
- `v1_out_auth_only_req` (Number) Outgoing Auth Only Request
- `v1_out_auth_only_rsp` (Number) Outgoing Auth Only Response
- `v1_out_id_prot_req` (Number) Outgoing ID Protection Request
- `v1_out_id_prot_rsp` (Number) Outgoing ID Protection Response
- `v1_out_info_v1_req` (Number) Outgoing Info Request
- `v1_out_info_v1_rsp` (Number) Outgoing Info Response
- `v1_out_new_group_mode_req` (Number) Outgoing New Group Mode Request
- `v1_out_new_group_mode_rsp` (Number) Outgoing New Group Mode Response
- `v1_out_quick_mode_req` (Number) Outgoing Quick Mode Request
- `v1_out_quick_mode_rsp` (Number) Outgoing Quick Mode Response
- `v1_out_transaction_req` (Number) Outgoing Transaction Request
- `v1_out_transaction_rsp` (Number) Outgoing Transaction Response
- `v2_child_sa_rekey` (Number) Child SA Rekey
- `v2_in_auth_req` (Number) Incoming Auth Request
- `v2_in_auth_rsp` (Number) Incoming Auth Response
- `v2_in_create_child_req` (Number) Incoming Create Child Request
- `v2_in_create_child_rsp` (Number) Incoming Create Child Response
- `v2_in_info_req` (Number) Incoming Info Request
- `v2_in_info_rsp` (Number) Incoming Info Response
- `v2_in_init_req` (Number) Incoming Init Request
- `v2_in_init_rsp` (Number) Incoming Init Response
- `v2_in_invalid` (Number) Incoming Invalid
- `v2_in_invalid_spi` (Number) Incoming Invalid SPI
- `v2_init_rekey` (Number) Initiate Rekey
- `v2_out_auth_req` (Number) Outgoing Auth Request
- `v2_out_auth_rsp` (Number) Outgoing Auth Response
- `v2_out_create_child_req` (Number) Outgoing Create Child Request
- `v2_out_create_child_rsp` (Number) Outgoing Create Child Response
- `v2_out_info_req` (Number) Outgoing Info Request
- `v2_out_info_rsp` (Number) Outgoing Info Response
- `v2_out_init_req` (Number) Outgoing Init Request
- `v2_out_init_rsp` (Number) Outgoing Init Response
- `v2_rsp_rekey` (Number) Respond Rekey



<a id="nestedblock--ipsec_list"></a>
### Nested Schema for `ipsec_list`

Required:

- `name` (String) IPsec name

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipsec_list--stats))

<a id="nestedblock--ipsec_list--stats"></a>
### Nested Schema for `ipsec_list.stats`

Optional:

- `anti_replay_num` (Number) Anti-Replay Failure
- `bytes_decrypted` (Number) Decrypted Bytes
- `bytes_encrypted` (Number) Encrypted Bytes
- `cavium_bytes_decrypted` (Number) CAVIUM Decrypted Bytes
- `cavium_bytes_encrypted` (Number) CAVIUM Encrypted Bytes
- `cavium_packets_decrypted` (Number) CAVIUM Decrypted Packets
- `cavium_packets_encrypted` (Number) CAVIUM Encrypted Packets
- `frag_after_encap_frag_packets` (Number) Frag-after-encap Fragment Generated
- `frag_received` (Number) Fragment Received
- `invalid_tunnel_id` (Number) Packet dropped: Invalid tunnel ID
- `no_next_hop` (Number) Packet dropped: No next hop
- `no_tunnel_found` (Number) Packet dropped: No tunnel found
- `packets_decrypted` (Number) Decrypted Packets
- `packets_encrypted` (Number) Encrypted Packets
- `packets_err_encryption` (Number) Encryption Error
- `packets_err_icv_check` (Number) ICV Check Error
- `packets_err_inactive` (Number) Inactive Error
- `packets_err_lifetime_lifebytes` (Number) Lifetime Lifebytes Error
- `packets_err_nh_check` (Number) Next Header Check Error
- `packets_err_pad_check` (Number) Pad Check Error
- `packets_err_pkt_sanity` (Number) Packets Sanity Error
- `pkt_fail_prep_to_send` (Number) Packet dropped: Failed in prepare to send
- `pkt_fail_to_send` (Number) Packet dropped: Failed to send
- `prefrag_error` (Number) Pre-frag Error
- `prefrag_success` (Number) Pre-frag Success
- `qat_bytes_decrypted` (Number) QAT Decrypted Bytes
- `qat_bytes_encrypted` (Number) QAT Encrypted Bytes
- `qat_packets_decrypted` (Number) QAT Decrypted Packets
- `qat_packets_encrypted` (Number) QAT Encrypted Packets
- `rekey_num` (Number) Rekey Times
- `sequence_num` (Number) Sequence Number
- `sequence_num_rollover` (Number) Sequence Number Rollover
- `tunnel_intf_down` (Number) Packet dropped: Tunnel Interface Down



<a id="nestedblock--ipsec_sa_stats_list"></a>
### Nested Schema for `ipsec_sa_stats_list`

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ipsec_sa_stats_list--stats))

<a id="nestedblock--ipsec_sa_stats_list--stats"></a>
### Nested Schema for `ipsec_sa_stats_list.stats`

Optional:

- `anti_replay_num` (Number) Anti-Replay Failure
- `bytes_decrypted` (Number) Decrypted Bytes
- `bytes_encrypted` (Number) Encrypted Bytes
- `cavium_bytes_decrypted` (Number) CAVIUM Decrypted Bytes
- `cavium_bytes_encrypted` (Number) CAVIUM Encrypted Bytes
- `cavium_packets_decrypted` (Number) CAVIUM Decrypted Packets
- `cavium_packets_encrypted` (Number) CAVIUM Encrypted Packets
- `frag_after_encap_frag_packets` (Number) Frag-after-encap Fragment Generated
- `frag_received` (Number) Fragment Received
- `invalid_tunnel_id` (Number) Packet dropped: Invalid tunnel ID
- `no_next_hop` (Number) Packet dropped: No next hop
- `no_tunnel_found` (Number) Packet dropped: No tunnel found
- `packets_decrypted` (Number) Decrypted Packets
- `packets_encrypted` (Number) Encrypted Packets
- `packets_err_encryption` (Number) Encryption Error
- `packets_err_icv_check` (Number) ICV Check Error
- `packets_err_inactive` (Number) Inactive Error
- `packets_err_lifetime_lifebytes` (Number) Lifetime Lifebytes Error
- `packets_err_nh_check` (Number) Next Header Check Error
- `packets_err_pad_check` (Number) Pad Check Error
- `packets_err_pkt_sanity` (Number) Packets Sanity Error
- `pkt_fail_prep_to_send` (Number) Packet dropped: Failed in prepare to send
- `pkt_fail_to_send` (Number) Packet dropped: Failed to send
- `prefrag_error` (Number) Pre-frag Error
- `prefrag_success` (Number) Pre-frag Success
- `qat_bytes_decrypted` (Number) QAT Decrypted Bytes
- `qat_bytes_encrypted` (Number) QAT Encrypted Bytes
- `qat_packets_decrypted` (Number) QAT Decrypted Packets
- `qat_packets_encrypted` (Number) QAT Encrypted Packets
- `rekey_num` (Number) Rekey Times
- `sequence_num` (Number) Sequence Number
- `sequence_num_rollover` (Number) Sequence Number Rollover
- `tunnel_intf_down` (Number) Packet dropped: Tunnel Interface Down



<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `ha_standby_drop` (Number)
- `passthrough` (Number)

