package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_stats`: Statistics for the object vpn\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnStatsRead,

		Schema: map[string]*schema.Schema{
			"error": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bad_opcode": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_sg_write_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_auth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_padding": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ip_version": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_auth_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_encrypt_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_spi": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_checksum": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_context": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_context_direction": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_context_flag_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipcomp_payload": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_selector_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_fragment_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_inline_data": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_frag_size_configuration": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dummy_payload": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ip_payload_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_min_frag_size_auth_sha384_512": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_esp_next_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_gre_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_gre_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_extension_headers_too_big": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_hop_by_hop_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error_ipv6_decrypt_rh_segs_left_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_rh_length_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv6_outbound_rh_copy_addr_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error_ipv6_extension_header_bad": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_encrypt_type_ctr_gcm": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ah_not_supported_with_gcm_gmac_sha2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tfc_padding_with_prefrag_not_supported": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_srtp_auth_tag": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipcomp_configuration": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dsiv_incorrect_param": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_ipsec_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"ike_gateway_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "IKE-gateway name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v2_init_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "Initiate Rekey",
									},
									"v2_rsp_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "Respond Rekey",
									},
									"v2_child_sa_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "Child SA Rekey",
									},
									"v2_in_invalid": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Invalid",
									},
									"v2_in_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Invalid SPI",
									},
									"v2_in_init_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Init Request",
									},
									"v2_in_init_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Init Response",
									},
									"v2_out_init_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Init Request",
									},
									"v2_out_init_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Init Response",
									},
									"v2_in_auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Request",
									},
									"v2_in_auth_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Response",
									},
									"v2_out_auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Request",
									},
									"v2_out_auth_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Response",
									},
									"v2_in_create_child_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Create Child Request",
									},
									"v2_in_create_child_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Create Child Response",
									},
									"v2_out_create_child_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Create Child Request",
									},
									"v2_out_create_child_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Create Child Response",
									},
									"v2_in_info_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Request",
									},
									"v2_in_info_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Response",
									},
									"v2_out_info_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Request",
									},
									"v2_out_info_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Response",
									},
									"v1_in_id_prot_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming ID Protection Request",
									},
									"v1_in_id_prot_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming ID Protection Response",
									},
									"v1_out_id_prot_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing ID Protection Request",
									},
									"v1_out_id_prot_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing ID Protection Response",
									},
									"v1_in_auth_only_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Only Request",
									},
									"v1_in_auth_only_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Only Response",
									},
									"v1_out_auth_only_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Only Request",
									},
									"v1_out_auth_only_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Only Response",
									},
									"v1_in_aggressive_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Aggressive Request",
									},
									"v1_in_aggressive_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Aggressive Response",
									},
									"v1_out_aggressive_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Aggressive Request",
									},
									"v1_out_aggressive_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Aggressive Response",
									},
									"v1_in_info_v1_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Request",
									},
									"v1_in_info_v1_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Response",
									},
									"v1_out_info_v1_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Request",
									},
									"v1_out_info_v1_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Response",
									},
									"v1_in_transaction_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Transaction Request",
									},
									"v1_in_transaction_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Transaction Response",
									},
									"v1_out_transaction_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Transaction Request",
									},
									"v1_out_transaction_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Transaction Response",
									},
									"v1_in_quick_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Quick Mode Request",
									},
									"v1_in_quick_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Quick Mode Response",
									},
									"v1_out_quick_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Quick Mode Request",
									},
									"v1_out_quick_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Quick Mode Response",
									},
									"v1_in_new_group_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming New Group Mode Request",
									},
									"v1_in_new_group_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming New Group Mode Response",
									},
									"v1_out_new_group_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing New Group Mode Request",
									},
									"v1_out_new_group_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing New Group Mode Response",
									},
									"v1_child_sa_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "Invalid SPI for Child SAs",
									},
									"v2_child_sa_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "Invalid SPI for Child SAs",
									},
									"ike_current_version": {
										Type: schema.TypeInt, Optional: true, Description: "IKE version",
									},
								},
							},
						},
					},
				},
			},
			"ike_stats_global": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v2_init_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "Initiate Rekey",
									},
									"v2_rsp_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "Respond Rekey",
									},
									"v2_child_sa_rekey": {
										Type: schema.TypeInt, Optional: true, Description: "Child SA Rekey",
									},
									"v2_in_invalid": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Invalid",
									},
									"v2_in_invalid_spi": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Invalid SPI",
									},
									"v2_in_init_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Init Request",
									},
									"v2_in_init_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Init Response",
									},
									"v2_out_init_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Init Request",
									},
									"v2_out_init_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Init Response",
									},
									"v2_in_auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Request",
									},
									"v2_in_auth_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Response",
									},
									"v2_out_auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Request",
									},
									"v2_out_auth_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Response",
									},
									"v2_in_create_child_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Create Child Request",
									},
									"v2_in_create_child_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Create Child Response",
									},
									"v2_out_create_child_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Create Child Request",
									},
									"v2_out_create_child_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Create Child Response",
									},
									"v2_in_info_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Request",
									},
									"v2_in_info_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Response",
									},
									"v2_out_info_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Request",
									},
									"v2_out_info_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Response",
									},
									"v1_in_id_prot_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming ID Protection Request",
									},
									"v1_in_id_prot_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming ID Protection Response",
									},
									"v1_out_id_prot_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing ID Protection Request",
									},
									"v1_out_id_prot_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing ID Protection Response",
									},
									"v1_in_auth_only_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Only Request",
									},
									"v1_in_auth_only_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Auth Only Response",
									},
									"v1_out_auth_only_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Only Request",
									},
									"v1_out_auth_only_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Auth Only Response",
									},
									"v1_in_aggressive_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Aggressive Request",
									},
									"v1_in_aggressive_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Aggressive Response",
									},
									"v1_out_aggressive_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Aggressive Request",
									},
									"v1_out_aggressive_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Aggressive Response",
									},
									"v1_in_info_v1_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Request",
									},
									"v1_in_info_v1_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Info Response",
									},
									"v1_out_info_v1_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Request",
									},
									"v1_out_info_v1_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Info Response",
									},
									"v1_in_transaction_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Transaction Request",
									},
									"v1_in_transaction_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Transaction Response",
									},
									"v1_out_transaction_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Transaction Request",
									},
									"v1_out_transaction_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Transaction Response",
									},
									"v1_in_quick_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Quick Mode Request",
									},
									"v1_in_quick_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming Quick Mode Response",
									},
									"v1_out_quick_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Quick Mode Request",
									},
									"v1_out_quick_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing Quick Mode Response",
									},
									"v1_in_new_group_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming New Group Mode Request",
									},
									"v1_in_new_group_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Incoming New Group Mode Response",
									},
									"v1_out_new_group_mode_req": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing New Group Mode Request",
									},
									"v1_out_new_group_mode_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "Outgoing New Group Mode Response",
									},
								},
							},
						},
					},
				},
			},
			"ipsec_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "IPsec name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Encrypted Packets",
									},
									"packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Decrypted Packets",
									},
									"anti_replay_num": {
										Type: schema.TypeInt, Optional: true, Description: "Anti-Replay Failure",
									},
									"rekey_num": {
										Type: schema.TypeInt, Optional: true, Description: "Rekey Times",
									},
									"packets_err_inactive": {
										Type: schema.TypeInt, Optional: true, Description: "Inactive Error",
									},
									"packets_err_encryption": {
										Type: schema.TypeInt, Optional: true, Description: "Encryption Error",
									},
									"packets_err_pad_check": {
										Type: schema.TypeInt, Optional: true, Description: "Pad Check Error",
									},
									"packets_err_pkt_sanity": {
										Type: schema.TypeInt, Optional: true, Description: "Packets Sanity Error",
									},
									"packets_err_icv_check": {
										Type: schema.TypeInt, Optional: true, Description: "ICV Check Error",
									},
									"packets_err_lifetime_lifebytes": {
										Type: schema.TypeInt, Optional: true, Description: "Lifetime Lifebytes Error",
									},
									"bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Encrypted Bytes",
									},
									"bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Decrypted Bytes",
									},
									"prefrag_success": {
										Type: schema.TypeInt, Optional: true, Description: "Pre-frag Success",
									},
									"prefrag_error": {
										Type: schema.TypeInt, Optional: true, Description: "Pre-frag Error",
									},
									"cavium_bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Encrypted Bytes",
									},
									"cavium_bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Decrypted Bytes",
									},
									"cavium_packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Encrypted Packets",
									},
									"cavium_packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Decrypted Packets",
									},
									"qat_bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Encrypted Bytes",
									},
									"qat_bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Decrypted Bytes",
									},
									"qat_packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Encrypted Packets",
									},
									"qat_packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Decrypted Packets",
									},
									"tunnel_intf_down": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Tunnel Interface Down",
									},
									"pkt_fail_prep_to_send": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Failed in prepare to send",
									},
									"no_next_hop": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: No next hop",
									},
									"invalid_tunnel_id": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Invalid tunnel ID",
									},
									"no_tunnel_found": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: No tunnel found",
									},
									"pkt_fail_to_send": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Failed to send",
									},
									"frag_after_encap_frag_packets": {
										Type: schema.TypeInt, Optional: true, Description: "Frag-after-encap Fragment Generated",
									},
									"frag_received": {
										Type: schema.TypeInt, Optional: true, Description: "Fragment Received",
									},
									"sequence_num": {
										Type: schema.TypeInt, Optional: true, Description: "Sequence Number",
									},
									"sequence_num_rollover": {
										Type: schema.TypeInt, Optional: true, Description: "Sequence Number Rollover",
									},
									"packets_err_nh_check": {
										Type: schema.TypeInt, Optional: true, Description: "Next Header Check Error",
									},
								},
							},
						},
					},
				},
			},
			"ipsec_sa_stats_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Encrypted Packets",
									},
									"packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Decrypted Packets",
									},
									"anti_replay_num": {
										Type: schema.TypeInt, Optional: true, Description: "Anti-Replay Failure",
									},
									"rekey_num": {
										Type: schema.TypeInt, Optional: true, Description: "Rekey Times",
									},
									"packets_err_inactive": {
										Type: schema.TypeInt, Optional: true, Description: "Inactive Error",
									},
									"packets_err_encryption": {
										Type: schema.TypeInt, Optional: true, Description: "Encryption Error",
									},
									"packets_err_pad_check": {
										Type: schema.TypeInt, Optional: true, Description: "Pad Check Error",
									},
									"packets_err_pkt_sanity": {
										Type: schema.TypeInt, Optional: true, Description: "Packets Sanity Error",
									},
									"packets_err_icv_check": {
										Type: schema.TypeInt, Optional: true, Description: "ICV Check Error",
									},
									"packets_err_lifetime_lifebytes": {
										Type: schema.TypeInt, Optional: true, Description: "Lifetime Lifebytes Error",
									},
									"bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Encrypted Bytes",
									},
									"bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "Decrypted Bytes",
									},
									"prefrag_success": {
										Type: schema.TypeInt, Optional: true, Description: "Pre-frag Success",
									},
									"prefrag_error": {
										Type: schema.TypeInt, Optional: true, Description: "Pre-frag Error",
									},
									"cavium_bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Encrypted Bytes",
									},
									"cavium_bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Decrypted Bytes",
									},
									"cavium_packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Encrypted Packets",
									},
									"cavium_packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "CAVIUM Decrypted Packets",
									},
									"qat_bytes_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Encrypted Bytes",
									},
									"qat_bytes_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Decrypted Bytes",
									},
									"qat_packets_encrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Encrypted Packets",
									},
									"qat_packets_decrypted": {
										Type: schema.TypeInt, Optional: true, Description: "QAT Decrypted Packets",
									},
									"tunnel_intf_down": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Tunnel Interface Down",
									},
									"pkt_fail_prep_to_send": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Failed in prepare to send",
									},
									"no_next_hop": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: No next hop",
									},
									"invalid_tunnel_id": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Invalid tunnel ID",
									},
									"no_tunnel_found": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: No tunnel found",
									},
									"pkt_fail_to_send": {
										Type: schema.TypeInt, Optional: true, Description: "Packet dropped: Failed to send",
									},
									"frag_after_encap_frag_packets": {
										Type: schema.TypeInt, Optional: true, Description: "Frag-after-encap Fragment Generated",
									},
									"frag_received": {
										Type: schema.TypeInt, Optional: true, Description: "Fragment Received",
									},
									"sequence_num": {
										Type: schema.TypeInt, Optional: true, Description: "Sequence Number",
									},
									"sequence_num_rollover": {
										Type: schema.TypeInt, Optional: true, Description: "Sequence Number Rollover",
									},
									"packets_err_nh_check": {
										Type: schema.TypeInt, Optional: true, Description: "Next Header Check Error",
									},
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"passthrough": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ha_standby_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnStatsError := setObjectVpnStatsError(res)
		d.Set("error", VpnStatsError)
		VpnStatsIkeGatewayList := setSliceVpnStatsIkeGatewayList(res)
		d.Set("ike_gateway_list", VpnStatsIkeGatewayList)
		VpnStatsIkeStatsGlobal := setObjectVpnStatsIkeStatsGlobal(res)
		d.Set("ike_stats_global", VpnStatsIkeStatsGlobal)
		VpnStatsIpsecList := setSliceVpnStatsIpsecList(res)
		d.Set("ipsec_list", VpnStatsIpsecList)
		VpnStatsIpsecSaStatsList := setSliceVpnStatsIpsecSaStatsList(res)
		d.Set("ipsec_sa_stats_list", VpnStatsIpsecSaStatsList)
		VpnStatsStats := setObjectVpnStatsStats(res)
		d.Set("stats", VpnStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnStatsError(ret edpt.DataVpnStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVpnStatsErrorStats(ret.DtVpnStats.Error.Stats),
		},
	}
}

func setObjectVpnStatsErrorStats(d edpt.VpnStatsErrorStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["bad_opcode"] = d.Bad_opcode

	in["bad_sg_write_len"] = d.Bad_sg_write_len

	in["bad_len"] = d.Bad_len

	in["bad_ipsec_protocol"] = d.Bad_ipsec_protocol

	in["bad_ipsec_auth"] = d.Bad_ipsec_auth

	in["bad_ipsec_padding"] = d.Bad_ipsec_padding

	in["bad_ip_version"] = d.Bad_ip_version

	in["bad_auth_type"] = d.Bad_auth_type

	in["bad_encrypt_type"] = d.Bad_encrypt_type

	in["bad_ipsec_spi"] = d.Bad_ipsec_spi

	in["bad_checksum"] = d.Bad_checksum

	in["bad_ipsec_context"] = d.Bad_ipsec_context

	in["bad_ipsec_context_direction"] = d.Bad_ipsec_context_direction

	in["bad_ipsec_context_flag_mismatch"] = d.Bad_ipsec_context_flag_mismatch

	in["ipcomp_payload"] = d.Ipcomp_payload

	in["bad_selector_match"] = d.Bad_selector_match

	in["bad_fragment_size"] = d.Bad_fragment_size

	in["bad_inline_data"] = d.Bad_inline_data

	in["bad_frag_size_configuration"] = d.Bad_frag_size_configuration

	in["dummy_payload"] = d.Dummy_payload

	in["bad_ip_payload_type"] = d.Bad_ip_payload_type

	in["bad_min_frag_size_auth_sha384_512"] = d.Bad_min_frag_size_auth_sha384_512

	in["bad_esp_next_header"] = d.Bad_esp_next_header

	in["bad_gre_header"] = d.Bad_gre_header

	in["bad_gre_protocol"] = d.Bad_gre_protocol

	in["ipv6_extension_headers_too_big"] = d.Ipv6_extension_headers_too_big

	in["ipv6_hop_by_hop_error"] = d.Ipv6_hop_by_hop_error

	in["error_ipv6_decrypt_rh_segs_left_error"] = d.Error_ipv6_decrypt_rh_segs_left_error

	in["ipv6_rh_length_error"] = d.Ipv6_rh_length_error

	in["ipv6_outbound_rh_copy_addr_error"] = d.Ipv6_outbound_rh_copy_addr_error

	in["error_ipv6_extension_header_bad"] = d.Error_ipv6_extension_header_bad

	in["bad_encrypt_type_ctr_gcm"] = d.Bad_encrypt_type_ctr_gcm

	in["ah_not_supported_with_gcm_gmac_sha2"] = d.Ah_not_supported_with_gcm_gmac_sha2

	in["tfc_padding_with_prefrag_not_supported"] = d.Tfc_padding_with_prefrag_not_supported

	in["bad_srtp_auth_tag"] = d.Bad_srtp_auth_tag

	in["bad_ipcomp_configuration"] = d.Bad_ipcomp_configuration

	in["dsiv_incorrect_param"] = d.Dsiv_incorrect_param

	in["bad_ipsec_unknown"] = d.Bad_ipsec_unknown
	result = append(result, in)
	return result
}

func setSliceVpnStatsIkeGatewayList(d edpt.DataVpnStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVpnStats.IkeGatewayList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectVpnStatsIkeGatewayListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectVpnStatsIkeGatewayListStats(d edpt.VpnStatsIkeGatewayListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["v2_init_rekey"] = d.V2InitRekey

	in["v2_rsp_rekey"] = d.V2RspRekey

	in["v2_child_sa_rekey"] = d.V2ChildSaRekey

	in["v2_in_invalid"] = d.V2InInvalid

	in["v2_in_invalid_spi"] = d.V2InInvalidSpi

	in["v2_in_init_req"] = d.V2InInitReq

	in["v2_in_init_rsp"] = d.V2InInitRsp

	in["v2_out_init_req"] = d.V2OutInitReq

	in["v2_out_init_rsp"] = d.V2OutInitRsp

	in["v2_in_auth_req"] = d.V2InAuthReq

	in["v2_in_auth_rsp"] = d.V2InAuthRsp

	in["v2_out_auth_req"] = d.V2OutAuthReq

	in["v2_out_auth_rsp"] = d.V2OutAuthRsp

	in["v2_in_create_child_req"] = d.V2InCreateChildReq

	in["v2_in_create_child_rsp"] = d.V2InCreateChildRsp

	in["v2_out_create_child_req"] = d.V2OutCreateChildReq

	in["v2_out_create_child_rsp"] = d.V2OutCreateChildRsp

	in["v2_in_info_req"] = d.V2InInfoReq

	in["v2_in_info_rsp"] = d.V2InInfoRsp

	in["v2_out_info_req"] = d.V2OutInfoReq

	in["v2_out_info_rsp"] = d.V2OutInfoRsp

	in["v1_in_id_prot_req"] = d.V1InIdProtReq

	in["v1_in_id_prot_rsp"] = d.V1InIdProtRsp

	in["v1_out_id_prot_req"] = d.V1OutIdProtReq

	in["v1_out_id_prot_rsp"] = d.V1OutIdProtRsp

	in["v1_in_auth_only_req"] = d.V1InAuthOnlyReq

	in["v1_in_auth_only_rsp"] = d.V1InAuthOnlyRsp

	in["v1_out_auth_only_req"] = d.V1OutAuthOnlyReq

	in["v1_out_auth_only_rsp"] = d.V1OutAuthOnlyRsp

	in["v1_in_aggressive_req"] = d.V1InAggressiveReq

	in["v1_in_aggressive_rsp"] = d.V1InAggressiveRsp

	in["v1_out_aggressive_req"] = d.V1OutAggressiveReq

	in["v1_out_aggressive_rsp"] = d.V1OutAggressiveRsp

	in["v1_in_info_v1_req"] = d.V1InInfoV1Req

	in["v1_in_info_v1_rsp"] = d.V1InInfoV1Rsp

	in["v1_out_info_v1_req"] = d.V1OutInfoV1Req

	in["v1_out_info_v1_rsp"] = d.V1OutInfoV1Rsp

	in["v1_in_transaction_req"] = d.V1InTransactionReq

	in["v1_in_transaction_rsp"] = d.V1InTransactionRsp

	in["v1_out_transaction_req"] = d.V1OutTransactionReq

	in["v1_out_transaction_rsp"] = d.V1OutTransactionRsp

	in["v1_in_quick_mode_req"] = d.V1InQuickModeReq

	in["v1_in_quick_mode_rsp"] = d.V1InQuickModeRsp

	in["v1_out_quick_mode_req"] = d.V1OutQuickModeReq

	in["v1_out_quick_mode_rsp"] = d.V1OutQuickModeRsp

	in["v1_in_new_group_mode_req"] = d.V1InNewGroupModeReq

	in["v1_in_new_group_mode_rsp"] = d.V1InNewGroupModeRsp

	in["v1_out_new_group_mode_req"] = d.V1OutNewGroupModeReq

	in["v1_out_new_group_mode_rsp"] = d.V1OutNewGroupModeRsp

	in["v1_child_sa_invalid_spi"] = d.V1ChildSaInvalidSpi

	in["v2_child_sa_invalid_spi"] = d.V2ChildSaInvalidSpi

	in["ike_current_version"] = d.IkeCurrentVersion
	result = append(result, in)
	return result
}

func setObjectVpnStatsIkeStatsGlobal(ret edpt.DataVpnStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVpnStatsIkeStatsGlobalStats(ret.DtVpnStats.IkeStatsGlobal.Stats),
		},
	}
}

func setObjectVpnStatsIkeStatsGlobalStats(d edpt.VpnStatsIkeStatsGlobalStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["v2_init_rekey"] = d.V2InitRekey

	in["v2_rsp_rekey"] = d.V2RspRekey

	in["v2_child_sa_rekey"] = d.V2ChildSaRekey

	in["v2_in_invalid"] = d.V2InInvalid

	in["v2_in_invalid_spi"] = d.V2InInvalidSpi

	in["v2_in_init_req"] = d.V2InInitReq

	in["v2_in_init_rsp"] = d.V2InInitRsp

	in["v2_out_init_req"] = d.V2OutInitReq

	in["v2_out_init_rsp"] = d.V2OutInitRsp

	in["v2_in_auth_req"] = d.V2InAuthReq

	in["v2_in_auth_rsp"] = d.V2InAuthRsp

	in["v2_out_auth_req"] = d.V2OutAuthReq

	in["v2_out_auth_rsp"] = d.V2OutAuthRsp

	in["v2_in_create_child_req"] = d.V2InCreateChildReq

	in["v2_in_create_child_rsp"] = d.V2InCreateChildRsp

	in["v2_out_create_child_req"] = d.V2OutCreateChildReq

	in["v2_out_create_child_rsp"] = d.V2OutCreateChildRsp

	in["v2_in_info_req"] = d.V2InInfoReq

	in["v2_in_info_rsp"] = d.V2InInfoRsp

	in["v2_out_info_req"] = d.V2OutInfoReq

	in["v2_out_info_rsp"] = d.V2OutInfoRsp

	in["v1_in_id_prot_req"] = d.V1InIdProtReq

	in["v1_in_id_prot_rsp"] = d.V1InIdProtRsp

	in["v1_out_id_prot_req"] = d.V1OutIdProtReq

	in["v1_out_id_prot_rsp"] = d.V1OutIdProtRsp

	in["v1_in_auth_only_req"] = d.V1InAuthOnlyReq

	in["v1_in_auth_only_rsp"] = d.V1InAuthOnlyRsp

	in["v1_out_auth_only_req"] = d.V1OutAuthOnlyReq

	in["v1_out_auth_only_rsp"] = d.V1OutAuthOnlyRsp

	in["v1_in_aggressive_req"] = d.V1InAggressiveReq

	in["v1_in_aggressive_rsp"] = d.V1InAggressiveRsp

	in["v1_out_aggressive_req"] = d.V1OutAggressiveReq

	in["v1_out_aggressive_rsp"] = d.V1OutAggressiveRsp

	in["v1_in_info_v1_req"] = d.V1InInfoV1Req

	in["v1_in_info_v1_rsp"] = d.V1InInfoV1Rsp

	in["v1_out_info_v1_req"] = d.V1OutInfoV1Req

	in["v1_out_info_v1_rsp"] = d.V1OutInfoV1Rsp

	in["v1_in_transaction_req"] = d.V1InTransactionReq

	in["v1_in_transaction_rsp"] = d.V1InTransactionRsp

	in["v1_out_transaction_req"] = d.V1OutTransactionReq

	in["v1_out_transaction_rsp"] = d.V1OutTransactionRsp

	in["v1_in_quick_mode_req"] = d.V1InQuickModeReq

	in["v1_in_quick_mode_rsp"] = d.V1InQuickModeRsp

	in["v1_out_quick_mode_req"] = d.V1OutQuickModeReq

	in["v1_out_quick_mode_rsp"] = d.V1OutQuickModeRsp

	in["v1_in_new_group_mode_req"] = d.V1InNewGroupModeReq

	in["v1_in_new_group_mode_rsp"] = d.V1InNewGroupModeRsp

	in["v1_out_new_group_mode_req"] = d.V1OutNewGroupModeReq

	in["v1_out_new_group_mode_rsp"] = d.V1OutNewGroupModeRsp
	result = append(result, in)
	return result
}

func setSliceVpnStatsIpsecList(d edpt.DataVpnStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVpnStats.IpsecList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectVpnStatsIpsecListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectVpnStatsIpsecListStats(d edpt.VpnStatsIpsecListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["packets_encrypted"] = d.PacketsEncrypted

	in["packets_decrypted"] = d.PacketsDecrypted

	in["anti_replay_num"] = d.AntiReplayNum

	in["rekey_num"] = d.RekeyNum

	in["packets_err_inactive"] = d.PacketsErrInactive

	in["packets_err_encryption"] = d.PacketsErrEncryption

	in["packets_err_pad_check"] = d.PacketsErrPadCheck

	in["packets_err_pkt_sanity"] = d.PacketsErrPktSanity

	in["packets_err_icv_check"] = d.PacketsErrIcvCheck

	in["packets_err_lifetime_lifebytes"] = d.PacketsErrLifetimeLifebytes

	in["bytes_encrypted"] = d.BytesEncrypted

	in["bytes_decrypted"] = d.BytesDecrypted

	in["prefrag_success"] = d.PrefragSuccess

	in["prefrag_error"] = d.PrefragError

	in["cavium_bytes_encrypted"] = d.CaviumBytesEncrypted

	in["cavium_bytes_decrypted"] = d.CaviumBytesDecrypted

	in["cavium_packets_encrypted"] = d.CaviumPacketsEncrypted

	in["cavium_packets_decrypted"] = d.CaviumPacketsDecrypted

	in["qat_bytes_encrypted"] = d.QatBytesEncrypted

	in["qat_bytes_decrypted"] = d.QatBytesDecrypted

	in["qat_packets_encrypted"] = d.QatPacketsEncrypted

	in["qat_packets_decrypted"] = d.QatPacketsDecrypted

	in["tunnel_intf_down"] = d.TunnelIntfDown

	in["pkt_fail_prep_to_send"] = d.PktFailPrepToSend

	in["no_next_hop"] = d.NoNextHop

	in["invalid_tunnel_id"] = d.InvalidTunnelId

	in["no_tunnel_found"] = d.NoTunnelFound

	in["pkt_fail_to_send"] = d.PktFailToSend

	in["frag_after_encap_frag_packets"] = d.FragAfterEncapFragPackets

	in["frag_received"] = d.FragReceived

	in["sequence_num"] = d.SequenceNum

	in["sequence_num_rollover"] = d.SequenceNumRollover

	in["packets_err_nh_check"] = d.PacketsErrNhCheck
	result = append(result, in)
	return result
}

func setSliceVpnStatsIpsecSaStatsList(d edpt.DataVpnStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVpnStats.IpsecSaStatsList {
		in := make(map[string]interface{})
		in["stats"] = setObjectVpnStatsIpsecSaStatsListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectVpnStatsIpsecSaStatsListStats(d edpt.VpnStatsIpsecSaStatsListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["packets_encrypted"] = d.PacketsEncrypted

	in["packets_decrypted"] = d.PacketsDecrypted

	in["anti_replay_num"] = d.AntiReplayNum

	in["rekey_num"] = d.RekeyNum

	in["packets_err_inactive"] = d.PacketsErrInactive

	in["packets_err_encryption"] = d.PacketsErrEncryption

	in["packets_err_pad_check"] = d.PacketsErrPadCheck

	in["packets_err_pkt_sanity"] = d.PacketsErrPktSanity

	in["packets_err_icv_check"] = d.PacketsErrIcvCheck

	in["packets_err_lifetime_lifebytes"] = d.PacketsErrLifetimeLifebytes

	in["bytes_encrypted"] = d.BytesEncrypted

	in["bytes_decrypted"] = d.BytesDecrypted

	in["prefrag_success"] = d.PrefragSuccess

	in["prefrag_error"] = d.PrefragError

	in["cavium_bytes_encrypted"] = d.CaviumBytesEncrypted

	in["cavium_bytes_decrypted"] = d.CaviumBytesDecrypted

	in["cavium_packets_encrypted"] = d.CaviumPacketsEncrypted

	in["cavium_packets_decrypted"] = d.CaviumPacketsDecrypted

	in["qat_bytes_encrypted"] = d.QatBytesEncrypted

	in["qat_bytes_decrypted"] = d.QatBytesDecrypted

	in["qat_packets_encrypted"] = d.QatPacketsEncrypted

	in["qat_packets_decrypted"] = d.QatPacketsDecrypted

	in["tunnel_intf_down"] = d.TunnelIntfDown

	in["pkt_fail_prep_to_send"] = d.PktFailPrepToSend

	in["no_next_hop"] = d.NoNextHop

	in["invalid_tunnel_id"] = d.InvalidTunnelId

	in["no_tunnel_found"] = d.NoTunnelFound

	in["pkt_fail_to_send"] = d.PktFailToSend

	in["frag_after_encap_frag_packets"] = d.FragAfterEncapFragPackets

	in["frag_received"] = d.FragReceived

	in["sequence_num"] = d.SequenceNum

	in["sequence_num_rollover"] = d.SequenceNumRollover

	in["packets_err_nh_check"] = d.PacketsErrNhCheck
	result = append(result, in)
	return result
}

func setObjectVpnStatsStats(ret edpt.DataVpnStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"passthrough":     ret.DtVpnStats.Stats.Passthrough,
			"ha_standby_drop": ret.DtVpnStats.Stats.HaStandbyDrop,
		},
	}
}

func getObjectVpnStatsError(d []interface{}) edpt.VpnStatsError {

	count1 := len(d)
	var ret edpt.VpnStatsError
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVpnStatsErrorStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVpnStatsErrorStats(d []interface{}) edpt.VpnStatsErrorStats {

	count1 := len(d)
	var ret edpt.VpnStatsErrorStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bad_opcode = in["bad_opcode"].(int)
		ret.Bad_sg_write_len = in["bad_sg_write_len"].(int)
		ret.Bad_len = in["bad_len"].(int)
		ret.Bad_ipsec_protocol = in["bad_ipsec_protocol"].(int)
		ret.Bad_ipsec_auth = in["bad_ipsec_auth"].(int)
		ret.Bad_ipsec_padding = in["bad_ipsec_padding"].(int)
		ret.Bad_ip_version = in["bad_ip_version"].(int)
		ret.Bad_auth_type = in["bad_auth_type"].(int)
		ret.Bad_encrypt_type = in["bad_encrypt_type"].(int)
		ret.Bad_ipsec_spi = in["bad_ipsec_spi"].(int)
		ret.Bad_checksum = in["bad_checksum"].(int)
		ret.Bad_ipsec_context = in["bad_ipsec_context"].(int)
		ret.Bad_ipsec_context_direction = in["bad_ipsec_context_direction"].(int)
		ret.Bad_ipsec_context_flag_mismatch = in["bad_ipsec_context_flag_mismatch"].(int)
		ret.Ipcomp_payload = in["ipcomp_payload"].(int)
		ret.Bad_selector_match = in["bad_selector_match"].(int)
		ret.Bad_fragment_size = in["bad_fragment_size"].(int)
		ret.Bad_inline_data = in["bad_inline_data"].(int)
		ret.Bad_frag_size_configuration = in["bad_frag_size_configuration"].(int)
		ret.Dummy_payload = in["dummy_payload"].(int)
		ret.Bad_ip_payload_type = in["bad_ip_payload_type"].(int)
		ret.Bad_min_frag_size_auth_sha384_512 = in["bad_min_frag_size_auth_sha384_512"].(int)
		ret.Bad_esp_next_header = in["bad_esp_next_header"].(int)
		ret.Bad_gre_header = in["bad_gre_header"].(int)
		ret.Bad_gre_protocol = in["bad_gre_protocol"].(int)
		ret.Ipv6_extension_headers_too_big = in["ipv6_extension_headers_too_big"].(int)
		ret.Ipv6_hop_by_hop_error = in["ipv6_hop_by_hop_error"].(int)
		ret.Error_ipv6_decrypt_rh_segs_left_error = in["error_ipv6_decrypt_rh_segs_left_error"].(int)
		ret.Ipv6_rh_length_error = in["ipv6_rh_length_error"].(int)
		ret.Ipv6_outbound_rh_copy_addr_error = in["ipv6_outbound_rh_copy_addr_error"].(int)
		ret.Error_ipv6_extension_header_bad = in["error_ipv6_extension_header_bad"].(int)
		ret.Bad_encrypt_type_ctr_gcm = in["bad_encrypt_type_ctr_gcm"].(int)
		ret.Ah_not_supported_with_gcm_gmac_sha2 = in["ah_not_supported_with_gcm_gmac_sha2"].(int)
		ret.Tfc_padding_with_prefrag_not_supported = in["tfc_padding_with_prefrag_not_supported"].(int)
		ret.Bad_srtp_auth_tag = in["bad_srtp_auth_tag"].(int)
		ret.Bad_ipcomp_configuration = in["bad_ipcomp_configuration"].(int)
		ret.Dsiv_incorrect_param = in["dsiv_incorrect_param"].(int)
		ret.Bad_ipsec_unknown = in["bad_ipsec_unknown"].(int)
	}
	return ret
}

func getSliceVpnStatsIkeGatewayList(d []interface{}) []edpt.VpnStatsIkeGatewayList {

	count1 := len(d)
	ret := make([]edpt.VpnStatsIkeGatewayList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnStatsIkeGatewayList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectVpnStatsIkeGatewayListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnStatsIkeGatewayListStats(d []interface{}) edpt.VpnStatsIkeGatewayListStats {

	count1 := len(d)
	var ret edpt.VpnStatsIkeGatewayListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V2InitRekey = in["v2_init_rekey"].(int)
		ret.V2RspRekey = in["v2_rsp_rekey"].(int)
		ret.V2ChildSaRekey = in["v2_child_sa_rekey"].(int)
		ret.V2InInvalid = in["v2_in_invalid"].(int)
		ret.V2InInvalidSpi = in["v2_in_invalid_spi"].(int)
		ret.V2InInitReq = in["v2_in_init_req"].(int)
		ret.V2InInitRsp = in["v2_in_init_rsp"].(int)
		ret.V2OutInitReq = in["v2_out_init_req"].(int)
		ret.V2OutInitRsp = in["v2_out_init_rsp"].(int)
		ret.V2InAuthReq = in["v2_in_auth_req"].(int)
		ret.V2InAuthRsp = in["v2_in_auth_rsp"].(int)
		ret.V2OutAuthReq = in["v2_out_auth_req"].(int)
		ret.V2OutAuthRsp = in["v2_out_auth_rsp"].(int)
		ret.V2InCreateChildReq = in["v2_in_create_child_req"].(int)
		ret.V2InCreateChildRsp = in["v2_in_create_child_rsp"].(int)
		ret.V2OutCreateChildReq = in["v2_out_create_child_req"].(int)
		ret.V2OutCreateChildRsp = in["v2_out_create_child_rsp"].(int)
		ret.V2InInfoReq = in["v2_in_info_req"].(int)
		ret.V2InInfoRsp = in["v2_in_info_rsp"].(int)
		ret.V2OutInfoReq = in["v2_out_info_req"].(int)
		ret.V2OutInfoRsp = in["v2_out_info_rsp"].(int)
		ret.V1InIdProtReq = in["v1_in_id_prot_req"].(int)
		ret.V1InIdProtRsp = in["v1_in_id_prot_rsp"].(int)
		ret.V1OutIdProtReq = in["v1_out_id_prot_req"].(int)
		ret.V1OutIdProtRsp = in["v1_out_id_prot_rsp"].(int)
		ret.V1InAuthOnlyReq = in["v1_in_auth_only_req"].(int)
		ret.V1InAuthOnlyRsp = in["v1_in_auth_only_rsp"].(int)
		ret.V1OutAuthOnlyReq = in["v1_out_auth_only_req"].(int)
		ret.V1OutAuthOnlyRsp = in["v1_out_auth_only_rsp"].(int)
		ret.V1InAggressiveReq = in["v1_in_aggressive_req"].(int)
		ret.V1InAggressiveRsp = in["v1_in_aggressive_rsp"].(int)
		ret.V1OutAggressiveReq = in["v1_out_aggressive_req"].(int)
		ret.V1OutAggressiveRsp = in["v1_out_aggressive_rsp"].(int)
		ret.V1InInfoV1Req = in["v1_in_info_v1_req"].(int)
		ret.V1InInfoV1Rsp = in["v1_in_info_v1_rsp"].(int)
		ret.V1OutInfoV1Req = in["v1_out_info_v1_req"].(int)
		ret.V1OutInfoV1Rsp = in["v1_out_info_v1_rsp"].(int)
		ret.V1InTransactionReq = in["v1_in_transaction_req"].(int)
		ret.V1InTransactionRsp = in["v1_in_transaction_rsp"].(int)
		ret.V1OutTransactionReq = in["v1_out_transaction_req"].(int)
		ret.V1OutTransactionRsp = in["v1_out_transaction_rsp"].(int)
		ret.V1InQuickModeReq = in["v1_in_quick_mode_req"].(int)
		ret.V1InQuickModeRsp = in["v1_in_quick_mode_rsp"].(int)
		ret.V1OutQuickModeReq = in["v1_out_quick_mode_req"].(int)
		ret.V1OutQuickModeRsp = in["v1_out_quick_mode_rsp"].(int)
		ret.V1InNewGroupModeReq = in["v1_in_new_group_mode_req"].(int)
		ret.V1InNewGroupModeRsp = in["v1_in_new_group_mode_rsp"].(int)
		ret.V1OutNewGroupModeReq = in["v1_out_new_group_mode_req"].(int)
		ret.V1OutNewGroupModeRsp = in["v1_out_new_group_mode_rsp"].(int)
		ret.V1ChildSaInvalidSpi = in["v1_child_sa_invalid_spi"].(int)
		ret.V2ChildSaInvalidSpi = in["v2_child_sa_invalid_spi"].(int)
		ret.IkeCurrentVersion = in["ike_current_version"].(int)
	}
	return ret
}

func getObjectVpnStatsIkeStatsGlobal(d []interface{}) edpt.VpnStatsIkeStatsGlobal {

	count1 := len(d)
	var ret edpt.VpnStatsIkeStatsGlobal
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVpnStatsIkeStatsGlobalStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVpnStatsIkeStatsGlobalStats(d []interface{}) edpt.VpnStatsIkeStatsGlobalStats {

	count1 := len(d)
	var ret edpt.VpnStatsIkeStatsGlobalStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V2InitRekey = in["v2_init_rekey"].(int)
		ret.V2RspRekey = in["v2_rsp_rekey"].(int)
		ret.V2ChildSaRekey = in["v2_child_sa_rekey"].(int)
		ret.V2InInvalid = in["v2_in_invalid"].(int)
		ret.V2InInvalidSpi = in["v2_in_invalid_spi"].(int)
		ret.V2InInitReq = in["v2_in_init_req"].(int)
		ret.V2InInitRsp = in["v2_in_init_rsp"].(int)
		ret.V2OutInitReq = in["v2_out_init_req"].(int)
		ret.V2OutInitRsp = in["v2_out_init_rsp"].(int)
		ret.V2InAuthReq = in["v2_in_auth_req"].(int)
		ret.V2InAuthRsp = in["v2_in_auth_rsp"].(int)
		ret.V2OutAuthReq = in["v2_out_auth_req"].(int)
		ret.V2OutAuthRsp = in["v2_out_auth_rsp"].(int)
		ret.V2InCreateChildReq = in["v2_in_create_child_req"].(int)
		ret.V2InCreateChildRsp = in["v2_in_create_child_rsp"].(int)
		ret.V2OutCreateChildReq = in["v2_out_create_child_req"].(int)
		ret.V2OutCreateChildRsp = in["v2_out_create_child_rsp"].(int)
		ret.V2InInfoReq = in["v2_in_info_req"].(int)
		ret.V2InInfoRsp = in["v2_in_info_rsp"].(int)
		ret.V2OutInfoReq = in["v2_out_info_req"].(int)
		ret.V2OutInfoRsp = in["v2_out_info_rsp"].(int)
		ret.V1InIdProtReq = in["v1_in_id_prot_req"].(int)
		ret.V1InIdProtRsp = in["v1_in_id_prot_rsp"].(int)
		ret.V1OutIdProtReq = in["v1_out_id_prot_req"].(int)
		ret.V1OutIdProtRsp = in["v1_out_id_prot_rsp"].(int)
		ret.V1InAuthOnlyReq = in["v1_in_auth_only_req"].(int)
		ret.V1InAuthOnlyRsp = in["v1_in_auth_only_rsp"].(int)
		ret.V1OutAuthOnlyReq = in["v1_out_auth_only_req"].(int)
		ret.V1OutAuthOnlyRsp = in["v1_out_auth_only_rsp"].(int)
		ret.V1InAggressiveReq = in["v1_in_aggressive_req"].(int)
		ret.V1InAggressiveRsp = in["v1_in_aggressive_rsp"].(int)
		ret.V1OutAggressiveReq = in["v1_out_aggressive_req"].(int)
		ret.V1OutAggressiveRsp = in["v1_out_aggressive_rsp"].(int)
		ret.V1InInfoV1Req = in["v1_in_info_v1_req"].(int)
		ret.V1InInfoV1Rsp = in["v1_in_info_v1_rsp"].(int)
		ret.V1OutInfoV1Req = in["v1_out_info_v1_req"].(int)
		ret.V1OutInfoV1Rsp = in["v1_out_info_v1_rsp"].(int)
		ret.V1InTransactionReq = in["v1_in_transaction_req"].(int)
		ret.V1InTransactionRsp = in["v1_in_transaction_rsp"].(int)
		ret.V1OutTransactionReq = in["v1_out_transaction_req"].(int)
		ret.V1OutTransactionRsp = in["v1_out_transaction_rsp"].(int)
		ret.V1InQuickModeReq = in["v1_in_quick_mode_req"].(int)
		ret.V1InQuickModeRsp = in["v1_in_quick_mode_rsp"].(int)
		ret.V1OutQuickModeReq = in["v1_out_quick_mode_req"].(int)
		ret.V1OutQuickModeRsp = in["v1_out_quick_mode_rsp"].(int)
		ret.V1InNewGroupModeReq = in["v1_in_new_group_mode_req"].(int)
		ret.V1InNewGroupModeRsp = in["v1_in_new_group_mode_rsp"].(int)
		ret.V1OutNewGroupModeReq = in["v1_out_new_group_mode_req"].(int)
		ret.V1OutNewGroupModeRsp = in["v1_out_new_group_mode_rsp"].(int)
	}
	return ret
}

func getSliceVpnStatsIpsecList(d []interface{}) []edpt.VpnStatsIpsecList {

	count1 := len(d)
	ret := make([]edpt.VpnStatsIpsecList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnStatsIpsecList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectVpnStatsIpsecListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnStatsIpsecListStats(d []interface{}) edpt.VpnStatsIpsecListStats {

	count1 := len(d)
	var ret edpt.VpnStatsIpsecListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketsEncrypted = in["packets_encrypted"].(int)
		ret.PacketsDecrypted = in["packets_decrypted"].(int)
		ret.AntiReplayNum = in["anti_replay_num"].(int)
		ret.RekeyNum = in["rekey_num"].(int)
		ret.PacketsErrInactive = in["packets_err_inactive"].(int)
		ret.PacketsErrEncryption = in["packets_err_encryption"].(int)
		ret.PacketsErrPadCheck = in["packets_err_pad_check"].(int)
		ret.PacketsErrPktSanity = in["packets_err_pkt_sanity"].(int)
		ret.PacketsErrIcvCheck = in["packets_err_icv_check"].(int)
		ret.PacketsErrLifetimeLifebytes = in["packets_err_lifetime_lifebytes"].(int)
		ret.BytesEncrypted = in["bytes_encrypted"].(int)
		ret.BytesDecrypted = in["bytes_decrypted"].(int)
		ret.PrefragSuccess = in["prefrag_success"].(int)
		ret.PrefragError = in["prefrag_error"].(int)
		ret.CaviumBytesEncrypted = in["cavium_bytes_encrypted"].(int)
		ret.CaviumBytesDecrypted = in["cavium_bytes_decrypted"].(int)
		ret.CaviumPacketsEncrypted = in["cavium_packets_encrypted"].(int)
		ret.CaviumPacketsDecrypted = in["cavium_packets_decrypted"].(int)
		ret.QatBytesEncrypted = in["qat_bytes_encrypted"].(int)
		ret.QatBytesDecrypted = in["qat_bytes_decrypted"].(int)
		ret.QatPacketsEncrypted = in["qat_packets_encrypted"].(int)
		ret.QatPacketsDecrypted = in["qat_packets_decrypted"].(int)
		ret.TunnelIntfDown = in["tunnel_intf_down"].(int)
		ret.PktFailPrepToSend = in["pkt_fail_prep_to_send"].(int)
		ret.NoNextHop = in["no_next_hop"].(int)
		ret.InvalidTunnelId = in["invalid_tunnel_id"].(int)
		ret.NoTunnelFound = in["no_tunnel_found"].(int)
		ret.PktFailToSend = in["pkt_fail_to_send"].(int)
		ret.FragAfterEncapFragPackets = in["frag_after_encap_frag_packets"].(int)
		ret.FragReceived = in["frag_received"].(int)
		ret.SequenceNum = in["sequence_num"].(int)
		ret.SequenceNumRollover = in["sequence_num_rollover"].(int)
		ret.PacketsErrNhCheck = in["packets_err_nh_check"].(int)
	}
	return ret
}

func getSliceVpnStatsIpsecSaStatsList(d []interface{}) []edpt.VpnStatsIpsecSaStatsList {

	count1 := len(d)
	ret := make([]edpt.VpnStatsIpsecSaStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnStatsIpsecSaStatsList
		oi.Stats = getObjectVpnStatsIpsecSaStatsListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnStatsIpsecSaStatsListStats(d []interface{}) edpt.VpnStatsIpsecSaStatsListStats {

	count1 := len(d)
	var ret edpt.VpnStatsIpsecSaStatsListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketsEncrypted = in["packets_encrypted"].(int)
		ret.PacketsDecrypted = in["packets_decrypted"].(int)
		ret.AntiReplayNum = in["anti_replay_num"].(int)
		ret.RekeyNum = in["rekey_num"].(int)
		ret.PacketsErrInactive = in["packets_err_inactive"].(int)
		ret.PacketsErrEncryption = in["packets_err_encryption"].(int)
		ret.PacketsErrPadCheck = in["packets_err_pad_check"].(int)
		ret.PacketsErrPktSanity = in["packets_err_pkt_sanity"].(int)
		ret.PacketsErrIcvCheck = in["packets_err_icv_check"].(int)
		ret.PacketsErrLifetimeLifebytes = in["packets_err_lifetime_lifebytes"].(int)
		ret.BytesEncrypted = in["bytes_encrypted"].(int)
		ret.BytesDecrypted = in["bytes_decrypted"].(int)
		ret.PrefragSuccess = in["prefrag_success"].(int)
		ret.PrefragError = in["prefrag_error"].(int)
		ret.CaviumBytesEncrypted = in["cavium_bytes_encrypted"].(int)
		ret.CaviumBytesDecrypted = in["cavium_bytes_decrypted"].(int)
		ret.CaviumPacketsEncrypted = in["cavium_packets_encrypted"].(int)
		ret.CaviumPacketsDecrypted = in["cavium_packets_decrypted"].(int)
		ret.QatBytesEncrypted = in["qat_bytes_encrypted"].(int)
		ret.QatBytesDecrypted = in["qat_bytes_decrypted"].(int)
		ret.QatPacketsEncrypted = in["qat_packets_encrypted"].(int)
		ret.QatPacketsDecrypted = in["qat_packets_decrypted"].(int)
		ret.TunnelIntfDown = in["tunnel_intf_down"].(int)
		ret.PktFailPrepToSend = in["pkt_fail_prep_to_send"].(int)
		ret.NoNextHop = in["no_next_hop"].(int)
		ret.InvalidTunnelId = in["invalid_tunnel_id"].(int)
		ret.NoTunnelFound = in["no_tunnel_found"].(int)
		ret.PktFailToSend = in["pkt_fail_to_send"].(int)
		ret.FragAfterEncapFragPackets = in["frag_after_encap_frag_packets"].(int)
		ret.FragReceived = in["frag_received"].(int)
		ret.SequenceNum = in["sequence_num"].(int)
		ret.SequenceNumRollover = in["sequence_num_rollover"].(int)
		ret.PacketsErrNhCheck = in["packets_err_nh_check"].(int)
	}
	return ret
}

func getObjectVpnStatsStats(d []interface{}) edpt.VpnStatsStats {

	count1 := len(d)
	var ret edpt.VpnStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Passthrough = in["passthrough"].(int)
		ret.HaStandbyDrop = in["ha_standby_drop"].(int)
	}
	return ret
}

func dataToEndpointVpnStats(d *schema.ResourceData) edpt.VpnStats {
	var ret edpt.VpnStats

	ret.Error = getObjectVpnStatsError(d.Get("error").([]interface{}))

	ret.IkeGatewayList = getSliceVpnStatsIkeGatewayList(d.Get("ike_gateway_list").([]interface{}))

	ret.IkeStatsGlobal = getObjectVpnStatsIkeStatsGlobal(d.Get("ike_stats_global").([]interface{}))

	ret.IpsecList = getSliceVpnStatsIpsecList(d.Get("ipsec_list").([]interface{}))

	ret.IpsecSaStatsList = getSliceVpnStatsIpsecSaStatsList(d.Get("ipsec_sa_stats_list").([]interface{}))

	ret.Stats = getObjectVpnStatsStats(d.Get("stats").([]interface{}))
	return ret
}
