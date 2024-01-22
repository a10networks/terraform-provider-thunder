package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_oper`: Operational Status for the object vpn\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnOperRead,

		Schema: map[string]*schema.Schema{
			"crl": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crl_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"subject": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"issuer": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"updates": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"serial": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"revoked": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"storage_type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"total_crls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"default": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_dh_group": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_auth_method": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_encryption": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_hash": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ike_lifetime": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ike_nat_traversal": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_local_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_remote_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_dpd_interval": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_dh_group": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_encryption": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_hash": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_priority": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_lifetime": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_lifebytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_traffic_selector": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_local_subnet": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_local_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_local_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_remote_subnet": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_remote_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_remote_protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_anti_replay_window": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"errordump": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec_error_dump_path": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"group_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"group_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipsec_sa_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_gateway_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"role": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"is_new_group": {
													Type: schema.TypeInt, Optional: true, Description: "a value of 1 indicates new group",
												},
												"grp_member_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"remote_ip_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_id_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"brief_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sa_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"initiator_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"responder_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"encryption": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hash": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"sign_hash": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lifetime": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"nat_traversal": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"remote_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dh_group": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"fragment_message_generated": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"fragment_message_received": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"fragmentation_error": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"fragment_reassemble_error": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ike_sa": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_sa_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"initiator_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"responder_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"encryption": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hash": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lifetime": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"nat_traversal": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ike_sa_brief": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_brief_remote_gw": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ike_sa_brief_remote_gw_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_brief_remote_gw_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_brief_remote_gw_lifetime": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_brief_remote_gw_status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ike_sa_clients": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_local_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_remote_gw": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ike_sa_clients_remote_gw_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_clients_remote_gw_remote_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_clients_remote_gw_user_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_clients_remote_gw_idle_time": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_clients_remote_gw_session_time": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_sa_clients_remote_gw_bytes": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ike_stats_by_gw": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gateway_name_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_ip_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_id_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"display_all_filter": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ike_stats_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_id": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_version": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"v1_in_id_prot_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_id_prot_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_id_prot_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_id_prot_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_auth_only_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_auth_only_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_auth_only_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_auth_only_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_aggressive_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_aggressive_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_aggressive_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_aggressive_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_info_v1_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_info_v1_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_info_v1_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_info_v1_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_transaction_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_transaction_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_transaction_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_transaction_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_quick_mode_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_quick_mode_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_quick_mode_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_quick_mode_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_new_group_mode_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_in_new_group_mode_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_new_group_mode_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_out_new_group_mode_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v1_child_sa_invalid_spi": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_init_rekey": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_rsp_rekey": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_child_sa_rekey": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_invalid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_invalid_spi": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_init_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_init_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_init_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_init_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_auth_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_auth_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_auth_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_auth_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_create_child_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_create_child_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_create_child_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_create_child_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_info_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_in_info_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_info_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_out_info_rsp": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"v2_child_sa_invalid_spi": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"remote_ts_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"remote_ts_v6_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"in_spi_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"out_spi_filter": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sa_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"sa_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ts_proto": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"local_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"peer_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"peer_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"local_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"encryption_algorithm": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hash_algorithm": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lifetime": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"lifebytes": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dh_group": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"nat_traversal": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"anti_replay": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"packets_encrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_decrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"anti_replay_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"rekey_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_inactive": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_encryption": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_pad_check": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_pkt_sanity": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_icv_check": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_lifetime_lifebytes": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
												},
												"bytes_encrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"bytes_decrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"prefrag_success": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"prefrag_error": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cavium_bytes_encrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cavium_bytes_decrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cavium_packets_encrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cavium_packets_decrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"qat_bytes_encrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"qat_bytes_decrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"qat_packets_encrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"qat_packets_decrypted": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"tunnel_intf_down": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"pkt_fail_prep_to_send": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"no_next_hop": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"invalid_tunnel_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"no_tunnel_found": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"pkt_fail_to_send": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"frag_after_encap_frag_packets": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"frag_received": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sequence_num": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"sequence_num_rollover": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"packets_err_nh_check": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"enforce_ts_encap_drop": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"enforce_ts_decap_drop": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ipsec_sa": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec_sa_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipsec_sa_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ike_gateway_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_ts": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_ts": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"in_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"out_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"encryption": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hash": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lifetime": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"lifebytes": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ipsec_sa_by_gw": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_gateway_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"local_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"peer_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_sa_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipsec_sa_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"local_ts": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"remote_ts": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"in_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"out_spi": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"protocol": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"encryption": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"hash": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"lifetime": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"lifebytes": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"ipsec_sa_clients": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipsec_clients": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipsec_clients_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"sa_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"local_ts": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"in_spi": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"out_spi": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"lifetime": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"lifebytes": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpn_log_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vpn_log_data": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"vpn_log_offset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vpn_log_over": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"follow": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"from_start": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"num_lines": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"ocsp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ocsp_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"subject": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"issuer": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"validity": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"certificate_status": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"total_ocsps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_gateway_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ike_sa_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_sa_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ipsec_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"num_hardware_devices": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"crypto_cores_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"crypto_cores_assigned_to_ipsec": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"crypto_mem": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_partition_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_gateway_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ike_sa_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_sa_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_stateless": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipsec_mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipsec_hardware_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"num_hardware_devices": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ike_hardware_accelerate": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"crypto_cores_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_cores_assigned_to_ipsec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_mem": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_req_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_enqueue_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_sg_buff_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_bad_pointer": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_bad_ctx_pointer": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_req_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_state_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"crypto_hw_err_time_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_time_out_state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"crypto_hw_err_buff_alloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"passthrough_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vpn_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"passthrough": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cpu_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"standby_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"partition_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"signed_auth_flag": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"all_partitions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"shared": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"specific_partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVpnOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnOperCrl := setObjectVpnOperCrl(res)
		d.Set("crl", VpnOperCrl)
		VpnOperDefault := setObjectVpnOperDefault(res)
		d.Set("default", VpnOperDefault)
		VpnOperErrordump := setObjectVpnOperErrordump(res)
		d.Set("errordump", VpnOperErrordump)
		VpnOperGroupList := setObjectVpnOperGroupList(res)
		d.Set("group_list", VpnOperGroupList)
		VpnOperIkeGatewayList := setSliceVpnOperIkeGatewayList(res)
		d.Set("ike_gateway_list", VpnOperIkeGatewayList)
		VpnOperIkeSa := setObjectVpnOperIkeSa(res)
		d.Set("ike_sa", VpnOperIkeSa)
		VpnOperIkeSaBrief := setObjectVpnOperIkeSaBrief(res)
		d.Set("ike_sa_brief", VpnOperIkeSaBrief)
		VpnOperIkeSaClients := setObjectVpnOperIkeSaClients(res)
		d.Set("ike_sa_clients", VpnOperIkeSaClients)
		VpnOperIkeStatsByGw := setObjectVpnOperIkeStatsByGw(res)
		d.Set("ike_stats_by_gw", VpnOperIkeStatsByGw)
		VpnOperIpsecList := setSliceVpnOperIpsecList(res)
		d.Set("ipsec_list", VpnOperIpsecList)
		VpnOperIpsecSa := setObjectVpnOperIpsecSa(res)
		d.Set("ipsec_sa", VpnOperIpsecSa)
		VpnOperIpsecSaByGw := setObjectVpnOperIpsecSaByGw(res)
		d.Set("ipsec_sa_by_gw", VpnOperIpsecSaByGw)
		VpnOperIpsecSaClients := setObjectVpnOperIpsecSaClients(res)
		d.Set("ipsec_sa_clients", VpnOperIpsecSaClients)
		VpnOperLog := setObjectVpnOperLog(res)
		d.Set("log", VpnOperLog)
		VpnOperOcsp := setObjectVpnOperOcsp(res)
		d.Set("ocsp", VpnOperOcsp)
		VpnOperOper := setObjectVpnOperOper(res)
		d.Set("oper", VpnOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnOperCrl(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperCrlOper(ret.DtVpnOper.Crl.Oper),
		},
	}
}

func setObjectVpnOperCrlOper(d edpt.VpnOperCrlOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["crl_list"] = setSliceVpnOperCrlOperCrlList(d.CrlList)

	in["total_crls"] = d.TotalCrls
	result = append(result, in)
	return result
}

func setSliceVpnOperCrlOperCrlList(d []edpt.VpnOperCrlOperCrlList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["subject"] = item.Subject
		in["issuer"] = item.Issuer
		in["updates"] = item.Updates
		in["serial"] = item.Serial
		in["revoked"] = item.Revoked
		in["storage_type"] = item.StorageType
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperDefault(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperDefaultOper(ret.DtVpnOper.Default.Oper),
		},
	}
}

func setObjectVpnOperDefaultOper(d edpt.VpnOperDefaultOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ike_version"] = d.IkeVersion

	in["ike_mode"] = d.IkeMode

	in["ike_dh_group"] = d.IkeDhGroup

	in["ike_auth_method"] = d.IkeAuthMethod

	in["ike_encryption"] = d.IkeEncryption

	in["ike_hash"] = d.IkeHash

	in["ike_priority"] = d.IkePriority

	in["ike_lifetime"] = d.IkeLifetime

	in["ike_nat_traversal"] = d.IkeNatTraversal

	in["ike_local_address"] = d.IkeLocalAddress

	in["ike_remote_address"] = d.IkeRemoteAddress

	in["ike_dpd_interval"] = d.IkeDpdInterval

	in["ipsec_mode"] = d.IpsecMode

	in["ipsec_protocol"] = d.IpsecProtocol

	in["ipsec_dh_group"] = d.IpsecDhGroup

	in["ipsec_encryption"] = d.IpsecEncryption

	in["ipsec_hash"] = d.IpsecHash

	in["ipsec_priority"] = d.IpsecPriority

	in["ipsec_lifetime"] = d.IpsecLifetime

	in["ipsec_lifebytes"] = d.IpsecLifebytes

	in["ipsec_traffic_selector"] = d.IpsecTrafficSelector

	in["ipsec_local_subnet"] = d.IpsecLocalSubnet

	in["ipsec_local_port"] = d.IpsecLocalPort

	in["ipsec_local_protocol"] = d.IpsecLocalProtocol

	in["ipsec_remote_subnet"] = d.IpsecRemoteSubnet

	in["ipsec_remote_port"] = d.IpsecRemotePort

	in["ipsec_remote_protocol"] = d.IpsecRemoteProtocol

	in["ipsec_anti_replay_window"] = d.IpsecAntiReplayWindow
	result = append(result, in)
	return result
}

func setObjectVpnOperErrordump(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperErrordumpOper(ret.DtVpnOper.Errordump.Oper),
		},
	}
}

func setObjectVpnOperErrordumpOper(d edpt.VpnOperErrordumpOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipsec_error_dump_path"] = d.IpsecErrorDumpPath
	result = append(result, in)
	return result
}

func setObjectVpnOperGroupList(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperGroupListOper(ret.DtVpnOper.GroupList.Oper),
		},
	}
}

func setObjectVpnOperGroupListOper(d edpt.VpnOperGroupListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["group_name"] = d.GroupName
	in["group_list"] = setSliceVpnOperGroupListOperGroupList(d.GroupList)
	result = append(result, in)
	return result
}

func setSliceVpnOperGroupListOperGroupList(d []edpt.VpnOperGroupListOperGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["ipsec_sa_name"] = item.IpsecSaName
		in["ike_gateway_name"] = item.IkeGatewayName
		in["priority"] = item.Priority
		in["status"] = item.Status
		in["role"] = item.Role
		in["is_new_group"] = item.IsNewGroup
		in["grp_member_count"] = item.GrpMemberCount
		result = append(result, in)
	}
	return result
}

func setSliceVpnOperIkeGatewayList(d edpt.DataVpnOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVpnOper.IkeGatewayList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["oper"] = setObjectVpnOperIkeGatewayListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIkeGatewayListOper(d edpt.VpnOperIkeGatewayListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["remote_ip_filter"] = d.RemoteIpFilter

	in["remote_id_filter"] = d.RemoteIdFilter

	in["brief_filter"] = d.BriefFilter
	in["sa_list"] = setSliceVpnOperIkeGatewayListOperSaList(d.SaList)
	result = append(result, in)
	return result
}

func setSliceVpnOperIkeGatewayListOperSaList(d []edpt.VpnOperIkeGatewayListOperSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["initiator_spi"] = item.InitiatorSpi
		in["responder_spi"] = item.ResponderSpi
		in["local_ip"] = item.LocalIp
		in["remote_ip"] = item.RemoteIp
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["sign_hash"] = item.SignHash
		in["lifetime"] = item.Lifetime
		in["status"] = item.Status
		in["nat_traversal"] = item.NatTraversal
		in["remote_id"] = item.RemoteId
		in["dh_group"] = item.DhGroup
		in["fragment_message_generated"] = item.FragmentMessageGenerated
		in["fragment_message_received"] = item.FragmentMessageReceived
		in["fragmentation_error"] = item.FragmentationError
		in["fragment_reassemble_error"] = item.FragmentReassembleError
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIkeSa(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIkeSaOper(ret.DtVpnOper.IkeSa.Oper),
		},
	}
}

func setObjectVpnOperIkeSaOper(d edpt.VpnOperIkeSaOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ike_sa_list"] = setSliceVpnOperIkeSaOperIkeSaList(d.IkeSaList)
	result = append(result, in)
	return result
}

func setSliceVpnOperIkeSaOperIkeSaList(d []edpt.VpnOperIkeSaOperIkeSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["initiator_spi"] = item.InitiatorSpi
		in["responder_spi"] = item.ResponderSpi
		in["local_ip"] = item.LocalIp
		in["remote_ip"] = item.RemoteIp
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["lifetime"] = item.Lifetime
		in["status"] = item.Status
		in["nat_traversal"] = item.NatTraversal
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIkeSaBrief(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIkeSaBriefOper(ret.DtVpnOper.IkeSaBrief.Oper),
		},
	}
}

func setObjectVpnOperIkeSaBriefOper(d edpt.VpnOperIkeSaBriefOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["name"] = d.Name

	in["local_ip"] = d.LocalIp
	in["ike_sa_brief_remote_gw"] = setSliceVpnOperIkeSaBriefOperIkeSaBriefRemoteGw(d.IkeSaBriefRemoteGw)
	result = append(result, in)
	return result
}

func setSliceVpnOperIkeSaBriefOperIkeSaBriefRemoteGw(d []edpt.VpnOperIkeSaBriefOperIkeSaBriefRemoteGw) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ike_sa_brief_remote_gw_ip"] = item.IkeSaBriefRemoteGwIp
		in["ike_sa_brief_remote_gw_id"] = item.IkeSaBriefRemoteGwId
		in["ike_sa_brief_remote_gw_lifetime"] = item.IkeSaBriefRemoteGwLifetime
		in["ike_sa_brief_remote_gw_status"] = item.IkeSaBriefRemoteGwStatus
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIkeSaClients(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIkeSaClientsOper(ret.DtVpnOper.IkeSaClients.Oper),
		},
	}
}

func setObjectVpnOperIkeSaClientsOper(d edpt.VpnOperIkeSaClientsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["name"] = d.Name

	in["ike_sa_clients_local_ip"] = d.IkeSaClientsLocalIp
	in["ike_sa_clients_remote_gw"] = setSliceVpnOperIkeSaClientsOperIkeSaClientsRemoteGw(d.IkeSaClientsRemoteGw)
	result = append(result, in)
	return result
}

func setSliceVpnOperIkeSaClientsOperIkeSaClientsRemoteGw(d []edpt.VpnOperIkeSaClientsOperIkeSaClientsRemoteGw) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ike_sa_clients_remote_gw_ip"] = item.IkeSaClientsRemoteGwIp
		in["ike_sa_clients_remote_gw_remote_id"] = item.IkeSaClientsRemoteGwRemoteId
		in["ike_sa_clients_remote_gw_user_id"] = item.IkeSaClientsRemoteGwUserId
		in["ike_sa_clients_remote_gw_idle_time"] = item.IkeSaClientsRemoteGwIdleTime
		in["ike_sa_clients_remote_gw_session_time"] = item.IkeSaClientsRemoteGwSessionTime
		in["ike_sa_clients_remote_gw_bytes"] = item.IkeSaClientsRemoteGwBytes
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIkeStatsByGw(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIkeStatsByGwOper(ret.DtVpnOper.IkeStatsByGw.Oper),
		},
	}
}

func setObjectVpnOperIkeStatsByGwOper(d edpt.VpnOperIkeStatsByGwOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["gateway_name_filter"] = d.GatewayNameFilter

	in["remote_ip_filter"] = d.RemoteIpFilter

	in["remote_id_filter"] = d.RemoteIdFilter

	in["display_all_filter"] = d.DisplayAllFilter
	in["ike_stats_list"] = setSliceVpnOperIkeStatsByGwOperIkeStatsList(d.IkeStatsList)
	result = append(result, in)
	return result
}

func setSliceVpnOperIkeStatsByGwOperIkeStatsList(d []edpt.VpnOperIkeStatsByGwOperIkeStatsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["remote_id"] = item.RemoteId
		in["remote_ip"] = item.RemoteIp
		in["ike_version"] = item.IkeVersion
		in["v1_in_id_prot_req"] = item.V1InIdProtReq
		in["v1_in_id_prot_rsp"] = item.V1InIdProtRsp
		in["v1_out_id_prot_req"] = item.V1OutIdProtReq
		in["v1_out_id_prot_rsp"] = item.V1OutIdProtRsp
		in["v1_in_auth_only_req"] = item.V1InAuthOnlyReq
		in["v1_in_auth_only_rsp"] = item.V1InAuthOnlyRsp
		in["v1_out_auth_only_req"] = item.V1OutAuthOnlyReq
		in["v1_out_auth_only_rsp"] = item.V1OutAuthOnlyRsp
		in["v1_in_aggressive_req"] = item.V1InAggressiveReq
		in["v1_in_aggressive_rsp"] = item.V1InAggressiveRsp
		in["v1_out_aggressive_req"] = item.V1OutAggressiveReq
		in["v1_out_aggressive_rsp"] = item.V1OutAggressiveRsp
		in["v1_in_info_v1_req"] = item.V1InInfoV1Req
		in["v1_in_info_v1_rsp"] = item.V1InInfoV1Rsp
		in["v1_out_info_v1_req"] = item.V1OutInfoV1Req
		in["v1_out_info_v1_rsp"] = item.V1OutInfoV1Rsp
		in["v1_in_transaction_req"] = item.V1InTransactionReq
		in["v1_in_transaction_rsp"] = item.V1InTransactionRsp
		in["v1_out_transaction_req"] = item.V1OutTransactionReq
		in["v1_out_transaction_rsp"] = item.V1OutTransactionRsp
		in["v1_in_quick_mode_req"] = item.V1InQuickModeReq
		in["v1_in_quick_mode_rsp"] = item.V1InQuickModeRsp
		in["v1_out_quick_mode_req"] = item.V1OutQuickModeReq
		in["v1_out_quick_mode_rsp"] = item.V1OutQuickModeRsp
		in["v1_in_new_group_mode_req"] = item.V1InNewGroupModeReq
		in["v1_in_new_group_mode_rsp"] = item.V1InNewGroupModeRsp
		in["v1_out_new_group_mode_req"] = item.V1OutNewGroupModeReq
		in["v1_out_new_group_mode_rsp"] = item.V1OutNewGroupModeRsp
		in["v1_child_sa_invalid_spi"] = item.V1ChildSaInvalidSpi
		in["v2_init_rekey"] = item.V2InitRekey
		in["v2_rsp_rekey"] = item.V2RspRekey
		in["v2_child_sa_rekey"] = item.V2ChildSaRekey
		in["v2_in_invalid"] = item.V2InInvalid
		in["v2_in_invalid_spi"] = item.V2InInvalidSpi
		in["v2_in_init_req"] = item.V2InInitReq
		in["v2_in_init_rsp"] = item.V2InInitRsp
		in["v2_out_init_req"] = item.V2OutInitReq
		in["v2_out_init_rsp"] = item.V2OutInitRsp
		in["v2_in_auth_req"] = item.V2InAuthReq
		in["v2_in_auth_rsp"] = item.V2InAuthRsp
		in["v2_out_auth_req"] = item.V2OutAuthReq
		in["v2_out_auth_rsp"] = item.V2OutAuthRsp
		in["v2_in_create_child_req"] = item.V2InCreateChildReq
		in["v2_in_create_child_rsp"] = item.V2InCreateChildRsp
		in["v2_out_create_child_req"] = item.V2OutCreateChildReq
		in["v2_out_create_child_rsp"] = item.V2OutCreateChildRsp
		in["v2_in_info_req"] = item.V2InInfoReq
		in["v2_in_info_rsp"] = item.V2InInfoRsp
		in["v2_out_info_req"] = item.V2OutInfoReq
		in["v2_out_info_rsp"] = item.V2OutInfoRsp
		in["v2_child_sa_invalid_spi"] = item.V2ChildSaInvalidSpi
		result = append(result, in)
	}
	return result
}

func setSliceVpnOperIpsecList(d edpt.DataVpnOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVpnOper.IpsecList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["oper"] = setObjectVpnOperIpsecListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIpsecListOper(d edpt.VpnOperIpsecListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["remote_ts_filter"] = d.RemoteTsFilter

	in["remote_ts_v6_filter"] = d.RemoteTsV6Filter

	in["in_spi_filter"] = d.InSpiFilter

	in["out_spi_filter"] = d.OutSpiFilter
	in["sa_list"] = setSliceVpnOperIpsecListOperSaList(d.SaList)
	result = append(result, in)
	return result
}

func setSliceVpnOperIpsecListOperSaList(d []edpt.VpnOperIpsecListOperSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["status"] = item.Status
		in["sa_index"] = item.SaIndex
		in["ts_proto"] = item.TsProto
		in["local_ip"] = item.LocalIp
		in["local_port"] = item.LocalPort
		in["peer_ip"] = item.PeerIp
		in["peer_port"] = item.PeerPort
		in["local_spi"] = item.LocalSpi
		in["remote_spi"] = item.RemoteSpi
		in["protocol"] = item.Protocol
		in["mode"] = item.Mode
		in["encryption_algorithm"] = item.EncryptionAlgorithm
		in["hash_algorithm"] = item.HashAlgorithm
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		in["dh_group"] = item.DhGroup
		in["nat_traversal"] = item.NatTraversal
		in["anti_replay"] = item.AntiReplay
		in["packets_encrypted"] = item.PacketsEncrypted
		in["packets_decrypted"] = item.PacketsDecrypted
		in["anti_replay_num"] = item.AntiReplayNum
		in["rekey_num"] = item.RekeyNum
		in["packets_err_inactive"] = item.PacketsErrInactive
		in["packets_err_encryption"] = item.PacketsErrEncryption
		in["packets_err_pad_check"] = item.PacketsErrPadCheck
		in["packets_err_pkt_sanity"] = item.PacketsErrPktSanity
		in["packets_err_icv_check"] = item.PacketsErrIcvCheck
		in["packets_err_lifetime_lifebytes"] = setObjectVpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes(item.PacketsErrLifetimeLifebytes)
		in["bytes_encrypted"] = item.BytesEncrypted
		in["bytes_decrypted"] = item.BytesDecrypted
		in["prefrag_success"] = item.PrefragSuccess
		in["prefrag_error"] = item.PrefragError
		in["cavium_bytes_encrypted"] = item.CaviumBytesEncrypted
		in["cavium_bytes_decrypted"] = item.CaviumBytesDecrypted
		in["cavium_packets_encrypted"] = item.CaviumPacketsEncrypted
		in["cavium_packets_decrypted"] = item.CaviumPacketsDecrypted
		in["qat_bytes_encrypted"] = item.QatBytesEncrypted
		in["qat_bytes_decrypted"] = item.QatBytesDecrypted
		in["qat_packets_encrypted"] = item.QatPacketsEncrypted
		in["qat_packets_decrypted"] = item.QatPacketsDecrypted
		in["tunnel_intf_down"] = item.TunnelIntfDown
		in["pkt_fail_prep_to_send"] = item.PktFailPrepToSend
		in["no_next_hop"] = item.NoNextHop
		in["invalid_tunnel_id"] = item.InvalidTunnelId
		in["no_tunnel_found"] = item.NoTunnelFound
		in["pkt_fail_to_send"] = item.PktFailToSend
		in["frag_after_encap_frag_packets"] = item.FragAfterEncapFragPackets
		in["frag_received"] = item.FragReceived
		in["sequence_num"] = item.SequenceNum
		in["sequence_num_rollover"] = item.SequenceNumRollover
		in["packets_err_nh_check"] = item.PacketsErrNhCheck
		in["enforce_ts_encap_drop"] = item.EnforceTsEncapDrop
		in["enforce_ts_decap_drop"] = item.EnforceTsDecapDrop
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes(d edpt.VpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setObjectVpnOperIpsecSa(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIpsecSaOper(ret.DtVpnOper.IpsecSa.Oper),
		},
	}
}

func setObjectVpnOperIpsecSaOper(d edpt.VpnOperIpsecSaOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ipsec_sa_list"] = setSliceVpnOperIpsecSaOperIpsecSaList(d.IpsecSaList)
	result = append(result, in)
	return result
}

func setSliceVpnOperIpsecSaOperIpsecSaList(d []edpt.VpnOperIpsecSaOperIpsecSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipsec_sa_name"] = item.IpsecSaName
		in["ike_gateway_name"] = item.IkeGatewayName
		in["local_ts"] = item.LocalTs
		in["remote_ts"] = item.RemoteTs
		in["in_spi"] = item.InSpi
		in["out_spi"] = item.OutSpi
		in["protocol"] = item.Protocol
		in["mode"] = item.Mode
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIpsecSaByGw(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIpsecSaByGwOper(ret.DtVpnOper.IpsecSaByGw.Oper),
		},
	}
}

func setObjectVpnOperIpsecSaByGwOper(d edpt.VpnOperIpsecSaByGwOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ike_gateway_name"] = d.IkeGatewayName

	in["local_ip"] = d.LocalIp

	in["peer_ip"] = d.PeerIp
	in["ipsec_sa_list"] = setSliceVpnOperIpsecSaByGwOperIpsecSaList(d.IpsecSaList)
	result = append(result, in)
	return result
}

func setSliceVpnOperIpsecSaByGwOperIpsecSaList(d []edpt.VpnOperIpsecSaByGwOperIpsecSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipsec_sa_name"] = item.IpsecSaName
		in["local_ts"] = item.LocalTs
		in["remote_ts"] = item.RemoteTs
		in["in_spi"] = item.InSpi
		in["out_spi"] = item.OutSpi
		in["protocol"] = item.Protocol
		in["mode"] = item.Mode
		in["encryption"] = item.Encryption
		in["hash"] = item.Hash
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperIpsecSaClients(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperIpsecSaClientsOper(ret.DtVpnOper.IpsecSaClients.Oper),
		},
	}
}

func setObjectVpnOperIpsecSaClientsOper(d edpt.VpnOperIpsecSaClientsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ipsec_clients"] = setSliceVpnOperIpsecSaClientsOperIpsecClients(d.IpsecClients)
	result = append(result, in)
	return result
}

func setSliceVpnOperIpsecSaClientsOperIpsecClients(d []edpt.VpnOperIpsecSaClientsOperIpsecClients) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ipsec_clients_ip"] = item.IpsecClientsIp
		in["sa_list"] = setSliceVpnOperIpsecSaClientsOperIpsecClientsSaList(item.SaList)
		result = append(result, in)
	}
	return result
}

func setSliceVpnOperIpsecSaClientsOperIpsecClientsSaList(d []edpt.VpnOperIpsecSaClientsOperIpsecClientsSaList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["local_ts"] = item.LocalTs
		in["in_spi"] = item.InSpi
		in["out_spi"] = item.OutSpi
		in["lifetime"] = item.Lifetime
		in["lifebytes"] = item.Lifebytes
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperLog(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperLogOper(ret.DtVpnOper.Log.Oper),
		},
	}
}

func setObjectVpnOperLogOper(d edpt.VpnOperLogOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["vpn_log_list"] = setSliceVpnOperLogOperVpnLogList(d.VpnLogList)

	in["vpn_log_offset"] = d.VpnLogOffset

	in["vpn_log_over"] = d.VpnLogOver

	in["follow"] = d.Follow

	in["from_start"] = d.FromStart

	in["num_lines"] = d.NumLines
	result = append(result, in)
	return result
}

func setSliceVpnOperLogOperVpnLogList(d []edpt.VpnOperLogOperVpnLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vpn_log_data"] = item.VpnLogData
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperOcsp(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVpnOperOcspOper(ret.DtVpnOper.Ocsp.Oper),
		},
	}
}

func setObjectVpnOperOcspOper(d edpt.VpnOperOcspOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["ocsp_list"] = setSliceVpnOperOcspOperOcspList(d.OcspList)

	in["total_ocsps"] = d.TotalOcsps
	result = append(result, in)
	return result
}

func setSliceVpnOperOcspOperOcspList(d []edpt.VpnOperOcspOperOcspList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["subject"] = item.Subject
		in["issuer"] = item.Issuer
		in["validity"] = item.Validity
		in["certificate_status"] = item.CertificateStatus
		result = append(result, in)
	}
	return result
}

func setObjectVpnOperOper(ret edpt.DataVpnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ike_gateway_total":              ret.DtVpnOper.Oper.IkeGatewayTotal,
			"ipsec_total":                    ret.DtVpnOper.Oper.IpsecTotal,
			"ike_sa_total":                   ret.DtVpnOper.Oper.IkeSaTotal,
			"ipsec_sa_total":                 ret.DtVpnOper.Oper.IpsecSaTotal,
			"ipsec_mode":                     ret.DtVpnOper.Oper.IpsecMode,
			"num_hardware_devices":           ret.DtVpnOper.Oper.NumHardwareDevices,
			"crypto_cores_total":             ret.DtVpnOper.Oper.CryptoCoresTotal,
			"crypto_cores_assigned_to_ipsec": ret.DtVpnOper.Oper.CryptoCoresAssignedToIpsec,
			"crypto_mem":                     ret.DtVpnOper.Oper.CryptoMem,
			"all_partition_list":             setSliceVpnOperOperAllPartitionList(ret.DtVpnOper.Oper.AllPartitionList),
			"all_partitions":                 ret.DtVpnOper.Oper.AllPartitions,
			"shared":                         ret.DtVpnOper.Oper.Shared,
			"specific_partition":             ret.DtVpnOper.Oper.SpecificPartition,
		},
	}
}

func setSliceVpnOperOperAllPartitionList(d []edpt.VpnOperOperAllPartitionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ike_gateway_total"] = item.IkeGatewayTotal
		in["ipsec_total"] = item.IpsecTotal
		in["ike_sa_total"] = item.IkeSaTotal
		in["ipsec_sa_total"] = item.IpsecSaTotal
		in["ipsec_stateless"] = item.IpsecStateless
		in["ipsec_mode"] = item.IpsecMode
		in["ipsec_hardware_type"] = item.IpsecHardwareType
		in["num_hardware_devices"] = item.NumHardwareDevices
		in["ike_hardware_accelerate"] = item.IkeHardwareAccelerate
		in["crypto_cores_total"] = item.CryptoCoresTotal
		in["crypto_cores_assigned_to_ipsec"] = item.CryptoCoresAssignedToIpsec
		in["crypto_mem"] = item.CryptoMem
		in["crypto_hw_err"] = item.CryptoHwErr
		in["crypto_hw_err_req_alloc_fail"] = item.CryptoHwErrReqAllocFail
		in["crypto_hw_err_enqueue_fail"] = item.CryptoHwErrEnqueueFail
		in["crypto_hw_err_sg_buff_alloc_fail"] = item.CryptoHwErrSgBuffAllocFail
		in["crypto_hw_err_bad_pointer"] = item.CryptoHwErrBadPointer
		in["crypto_hw_err_bad_ctx_pointer"] = item.CryptoHwErrBadCtxPointer
		in["crypto_hw_err_req_error"] = item.CryptoHwErrReqError
		in["crypto_hw_err_state_error"] = item.CryptoHwErrStateError
		in["crypto_hw_err_state"] = item.CryptoHwErrState
		in["crypto_hw_err_time_out"] = item.CryptoHwErrTimeOut
		in["crypto_hw_err_time_out_state"] = item.CryptoHwErrTimeOutState
		in["crypto_hw_err_buff_alloc_error"] = item.CryptoHwErrBuffAllocError
		in["passthrough_total"] = item.PassthroughTotal
		in["vpn_list"] = setSliceVpnOperOperAllPartitionListVpnList(item.VpnList)
		in["standby_drop"] = item.StandbyDrop
		in["partition_name"] = item.PartitionName
		in["signed_auth_flag"] = item.SignedAuthFlag
		result = append(result, in)
	}
	return result
}

func setSliceVpnOperOperAllPartitionListVpnList(d []edpt.VpnOperOperAllPartitionListVpnList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["passthrough"] = item.Passthrough
		in["cpu_id"] = item.CpuId
		result = append(result, in)
	}
	return result
}

func getObjectVpnOperCrl(d []interface{}) edpt.VpnOperCrl {

	count1 := len(d)
	var ret edpt.VpnOperCrl
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperCrlOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperCrlOper(d []interface{}) edpt.VpnOperCrlOper {

	count1 := len(d)
	var ret edpt.VpnOperCrlOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CrlList = getSliceVpnOperCrlOperCrlList(in["crl_list"].([]interface{}))
		ret.TotalCrls = in["total_crls"].(int)
	}
	return ret
}

func getSliceVpnOperCrlOperCrlList(d []interface{}) []edpt.VpnOperCrlOperCrlList {

	count1 := len(d)
	ret := make([]edpt.VpnOperCrlOperCrlList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperCrlOperCrlList
		oi.Subject = in["subject"].(string)
		oi.Issuer = in["issuer"].(string)
		oi.Updates = in["updates"].(string)
		oi.Serial = in["serial"].(string)
		oi.Revoked = in["revoked"].(string)
		oi.StorageType = in["storage_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperDefault(d []interface{}) edpt.VpnOperDefault {

	count1 := len(d)
	var ret edpt.VpnOperDefault
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperDefaultOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperDefaultOper(d []interface{}) edpt.VpnOperDefaultOper {

	count1 := len(d)
	var ret edpt.VpnOperDefaultOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeVersion = in["ike_version"].(string)
		ret.IkeMode = in["ike_mode"].(string)
		ret.IkeDhGroup = in["ike_dh_group"].(string)
		ret.IkeAuthMethod = in["ike_auth_method"].(string)
		ret.IkeEncryption = in["ike_encryption"].(string)
		ret.IkeHash = in["ike_hash"].(string)
		ret.IkePriority = in["ike_priority"].(int)
		ret.IkeLifetime = in["ike_lifetime"].(int)
		ret.IkeNatTraversal = in["ike_nat_traversal"].(string)
		ret.IkeLocalAddress = in["ike_local_address"].(string)
		ret.IkeRemoteAddress = in["ike_remote_address"].(string)
		ret.IkeDpdInterval = in["ike_dpd_interval"].(int)
		ret.IpsecMode = in["ipsec_mode"].(string)
		ret.IpsecProtocol = in["ipsec_protocol"].(string)
		ret.IpsecDhGroup = in["ipsec_dh_group"].(string)
		ret.IpsecEncryption = in["ipsec_encryption"].(string)
		ret.IpsecHash = in["ipsec_hash"].(string)
		ret.IpsecPriority = in["ipsec_priority"].(int)
		ret.IpsecLifetime = in["ipsec_lifetime"].(int)
		ret.IpsecLifebytes = in["ipsec_lifebytes"].(int)
		ret.IpsecTrafficSelector = in["ipsec_traffic_selector"].(string)
		ret.IpsecLocalSubnet = in["ipsec_local_subnet"].(string)
		ret.IpsecLocalPort = in["ipsec_local_port"].(int)
		ret.IpsecLocalProtocol = in["ipsec_local_protocol"].(int)
		ret.IpsecRemoteSubnet = in["ipsec_remote_subnet"].(string)
		ret.IpsecRemotePort = in["ipsec_remote_port"].(int)
		ret.IpsecRemoteProtocol = in["ipsec_remote_protocol"].(int)
		ret.IpsecAntiReplayWindow = in["ipsec_anti_replay_window"].(int)
	}
	return ret
}

func getObjectVpnOperErrordump(d []interface{}) edpt.VpnOperErrordump {

	count1 := len(d)
	var ret edpt.VpnOperErrordump
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperErrordumpOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperErrordumpOper(d []interface{}) edpt.VpnOperErrordumpOper {

	count1 := len(d)
	var ret edpt.VpnOperErrordumpOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpsecErrorDumpPath = in["ipsec_error_dump_path"].(string)
	}
	return ret
}

func getObjectVpnOperGroupList(d []interface{}) edpt.VpnOperGroupList {

	count1 := len(d)
	var ret edpt.VpnOperGroupList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperGroupListOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperGroupListOper(d []interface{}) edpt.VpnOperGroupListOper {

	count1 := len(d)
	var ret edpt.VpnOperGroupListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupName = in["group_name"].(string)
		ret.GroupList = getSliceVpnOperGroupListOperGroupList(in["group_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperGroupListOperGroupList(d []interface{}) []edpt.VpnOperGroupListOperGroupList {

	count1 := len(d)
	ret := make([]edpt.VpnOperGroupListOperGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperGroupListOperGroupList
		oi.Name = in["name"].(string)
		oi.IpsecSaName = in["ipsec_sa_name"].(string)
		oi.IkeGatewayName = in["ike_gateway_name"].(string)
		oi.Priority = in["priority"].(int)
		oi.Status = in["status"].(string)
		oi.Role = in["role"].(string)
		oi.IsNewGroup = in["is_new_group"].(int)
		oi.GrpMemberCount = in["grp_member_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnOperIkeGatewayList(d []interface{}) []edpt.VpnOperIkeGatewayList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIkeGatewayList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIkeGatewayList
		oi.Name = in["name"].(string)
		oi.Oper = getObjectVpnOperIkeGatewayListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIkeGatewayListOper(d []interface{}) edpt.VpnOperIkeGatewayListOper {

	count1 := len(d)
	var ret edpt.VpnOperIkeGatewayListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteIpFilter = in["remote_ip_filter"].(string)
		ret.RemoteIdFilter = in["remote_id_filter"].(string)
		ret.BriefFilter = in["brief_filter"].(string)
		ret.SaList = getSliceVpnOperIkeGatewayListOperSaList(in["sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIkeGatewayListOperSaList(d []interface{}) []edpt.VpnOperIkeGatewayListOperSaList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIkeGatewayListOperSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIkeGatewayListOperSaList
		oi.InitiatorSpi = in["initiator_spi"].(string)
		oi.ResponderSpi = in["responder_spi"].(string)
		oi.LocalIp = in["local_ip"].(string)
		oi.RemoteIp = in["remote_ip"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.SignHash = in["sign_hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Status = in["status"].(string)
		oi.NatTraversal = in["nat_traversal"].(int)
		oi.RemoteId = in["remote_id"].(string)
		oi.DhGroup = in["dh_group"].(int)
		oi.FragmentMessageGenerated = in["fragment_message_generated"].(int)
		oi.FragmentMessageReceived = in["fragment_message_received"].(int)
		oi.FragmentationError = in["fragmentation_error"].(int)
		oi.FragmentReassembleError = in["fragment_reassemble_error"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIkeSa(d []interface{}) edpt.VpnOperIkeSa {

	count1 := len(d)
	var ret edpt.VpnOperIkeSa
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIkeSaOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIkeSaOper(d []interface{}) edpt.VpnOperIkeSaOper {

	count1 := len(d)
	var ret edpt.VpnOperIkeSaOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeSaList = getSliceVpnOperIkeSaOperIkeSaList(in["ike_sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIkeSaOperIkeSaList(d []interface{}) []edpt.VpnOperIkeSaOperIkeSaList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIkeSaOperIkeSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIkeSaOperIkeSaList
		oi.Name = in["name"].(string)
		oi.InitiatorSpi = in["initiator_spi"].(string)
		oi.ResponderSpi = in["responder_spi"].(string)
		oi.LocalIp = in["local_ip"].(string)
		oi.RemoteIp = in["remote_ip"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Status = in["status"].(string)
		oi.NatTraversal = in["nat_traversal"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIkeSaBrief(d []interface{}) edpt.VpnOperIkeSaBrief {

	count1 := len(d)
	var ret edpt.VpnOperIkeSaBrief
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIkeSaBriefOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIkeSaBriefOper(d []interface{}) edpt.VpnOperIkeSaBriefOper {

	count1 := len(d)
	var ret edpt.VpnOperIkeSaBriefOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.IkeSaBriefRemoteGw = getSliceVpnOperIkeSaBriefOperIkeSaBriefRemoteGw(in["ike_sa_brief_remote_gw"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIkeSaBriefOperIkeSaBriefRemoteGw(d []interface{}) []edpt.VpnOperIkeSaBriefOperIkeSaBriefRemoteGw {

	count1 := len(d)
	ret := make([]edpt.VpnOperIkeSaBriefOperIkeSaBriefRemoteGw, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIkeSaBriefOperIkeSaBriefRemoteGw
		oi.IkeSaBriefRemoteGwIp = in["ike_sa_brief_remote_gw_ip"].(string)
		oi.IkeSaBriefRemoteGwId = in["ike_sa_brief_remote_gw_id"].(string)
		oi.IkeSaBriefRemoteGwLifetime = in["ike_sa_brief_remote_gw_lifetime"].(string)
		oi.IkeSaBriefRemoteGwStatus = in["ike_sa_brief_remote_gw_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIkeSaClients(d []interface{}) edpt.VpnOperIkeSaClients {

	count1 := len(d)
	var ret edpt.VpnOperIkeSaClients
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIkeSaClientsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIkeSaClientsOper(d []interface{}) edpt.VpnOperIkeSaClientsOper {

	count1 := len(d)
	var ret edpt.VpnOperIkeSaClientsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.IkeSaClientsLocalIp = in["ike_sa_clients_local_ip"].(string)
		ret.IkeSaClientsRemoteGw = getSliceVpnOperIkeSaClientsOperIkeSaClientsRemoteGw(in["ike_sa_clients_remote_gw"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIkeSaClientsOperIkeSaClientsRemoteGw(d []interface{}) []edpt.VpnOperIkeSaClientsOperIkeSaClientsRemoteGw {

	count1 := len(d)
	ret := make([]edpt.VpnOperIkeSaClientsOperIkeSaClientsRemoteGw, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIkeSaClientsOperIkeSaClientsRemoteGw
		oi.IkeSaClientsRemoteGwIp = in["ike_sa_clients_remote_gw_ip"].(string)
		oi.IkeSaClientsRemoteGwRemoteId = in["ike_sa_clients_remote_gw_remote_id"].(string)
		oi.IkeSaClientsRemoteGwUserId = in["ike_sa_clients_remote_gw_user_id"].(string)
		oi.IkeSaClientsRemoteGwIdleTime = in["ike_sa_clients_remote_gw_idle_time"].(string)
		oi.IkeSaClientsRemoteGwSessionTime = in["ike_sa_clients_remote_gw_session_time"].(string)
		oi.IkeSaClientsRemoteGwBytes = in["ike_sa_clients_remote_gw_bytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIkeStatsByGw(d []interface{}) edpt.VpnOperIkeStatsByGw {

	count1 := len(d)
	var ret edpt.VpnOperIkeStatsByGw
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIkeStatsByGwOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIkeStatsByGwOper(d []interface{}) edpt.VpnOperIkeStatsByGwOper {

	count1 := len(d)
	var ret edpt.VpnOperIkeStatsByGwOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GatewayNameFilter = in["gateway_name_filter"].(string)
		ret.RemoteIpFilter = in["remote_ip_filter"].(string)
		ret.RemoteIdFilter = in["remote_id_filter"].(string)
		ret.DisplayAllFilter = in["display_all_filter"].(int)
		ret.IkeStatsList = getSliceVpnOperIkeStatsByGwOperIkeStatsList(in["ike_stats_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIkeStatsByGwOperIkeStatsList(d []interface{}) []edpt.VpnOperIkeStatsByGwOperIkeStatsList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIkeStatsByGwOperIkeStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIkeStatsByGwOperIkeStatsList
		oi.Name = in["name"].(string)
		oi.RemoteId = in["remote_id"].(string)
		oi.RemoteIp = in["remote_ip"].(string)
		oi.IkeVersion = in["ike_version"].(string)
		oi.V1InIdProtReq = in["v1_in_id_prot_req"].(int)
		oi.V1InIdProtRsp = in["v1_in_id_prot_rsp"].(int)
		oi.V1OutIdProtReq = in["v1_out_id_prot_req"].(int)
		oi.V1OutIdProtRsp = in["v1_out_id_prot_rsp"].(int)
		oi.V1InAuthOnlyReq = in["v1_in_auth_only_req"].(int)
		oi.V1InAuthOnlyRsp = in["v1_in_auth_only_rsp"].(int)
		oi.V1OutAuthOnlyReq = in["v1_out_auth_only_req"].(int)
		oi.V1OutAuthOnlyRsp = in["v1_out_auth_only_rsp"].(int)
		oi.V1InAggressiveReq = in["v1_in_aggressive_req"].(int)
		oi.V1InAggressiveRsp = in["v1_in_aggressive_rsp"].(int)
		oi.V1OutAggressiveReq = in["v1_out_aggressive_req"].(int)
		oi.V1OutAggressiveRsp = in["v1_out_aggressive_rsp"].(int)
		oi.V1InInfoV1Req = in["v1_in_info_v1_req"].(int)
		oi.V1InInfoV1Rsp = in["v1_in_info_v1_rsp"].(int)
		oi.V1OutInfoV1Req = in["v1_out_info_v1_req"].(int)
		oi.V1OutInfoV1Rsp = in["v1_out_info_v1_rsp"].(int)
		oi.V1InTransactionReq = in["v1_in_transaction_req"].(int)
		oi.V1InTransactionRsp = in["v1_in_transaction_rsp"].(int)
		oi.V1OutTransactionReq = in["v1_out_transaction_req"].(int)
		oi.V1OutTransactionRsp = in["v1_out_transaction_rsp"].(int)
		oi.V1InQuickModeReq = in["v1_in_quick_mode_req"].(int)
		oi.V1InQuickModeRsp = in["v1_in_quick_mode_rsp"].(int)
		oi.V1OutQuickModeReq = in["v1_out_quick_mode_req"].(int)
		oi.V1OutQuickModeRsp = in["v1_out_quick_mode_rsp"].(int)
		oi.V1InNewGroupModeReq = in["v1_in_new_group_mode_req"].(int)
		oi.V1InNewGroupModeRsp = in["v1_in_new_group_mode_rsp"].(int)
		oi.V1OutNewGroupModeReq = in["v1_out_new_group_mode_req"].(int)
		oi.V1OutNewGroupModeRsp = in["v1_out_new_group_mode_rsp"].(int)
		oi.V1ChildSaInvalidSpi = in["v1_child_sa_invalid_spi"].(int)
		oi.V2InitRekey = in["v2_init_rekey"].(int)
		oi.V2RspRekey = in["v2_rsp_rekey"].(int)
		oi.V2ChildSaRekey = in["v2_child_sa_rekey"].(int)
		oi.V2InInvalid = in["v2_in_invalid"].(int)
		oi.V2InInvalidSpi = in["v2_in_invalid_spi"].(int)
		oi.V2InInitReq = in["v2_in_init_req"].(int)
		oi.V2InInitRsp = in["v2_in_init_rsp"].(int)
		oi.V2OutInitReq = in["v2_out_init_req"].(int)
		oi.V2OutInitRsp = in["v2_out_init_rsp"].(int)
		oi.V2InAuthReq = in["v2_in_auth_req"].(int)
		oi.V2InAuthRsp = in["v2_in_auth_rsp"].(int)
		oi.V2OutAuthReq = in["v2_out_auth_req"].(int)
		oi.V2OutAuthRsp = in["v2_out_auth_rsp"].(int)
		oi.V2InCreateChildReq = in["v2_in_create_child_req"].(int)
		oi.V2InCreateChildRsp = in["v2_in_create_child_rsp"].(int)
		oi.V2OutCreateChildReq = in["v2_out_create_child_req"].(int)
		oi.V2OutCreateChildRsp = in["v2_out_create_child_rsp"].(int)
		oi.V2InInfoReq = in["v2_in_info_req"].(int)
		oi.V2InInfoRsp = in["v2_in_info_rsp"].(int)
		oi.V2OutInfoReq = in["v2_out_info_req"].(int)
		oi.V2OutInfoRsp = in["v2_out_info_rsp"].(int)
		oi.V2ChildSaInvalidSpi = in["v2_child_sa_invalid_spi"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnOperIpsecList(d []interface{}) []edpt.VpnOperIpsecList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIpsecList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIpsecList
		oi.Name = in["name"].(string)
		oi.Oper = getObjectVpnOperIpsecListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIpsecListOper(d []interface{}) edpt.VpnOperIpsecListOper {

	count1 := len(d)
	var ret edpt.VpnOperIpsecListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RemoteTsFilter = in["remote_ts_filter"].(string)
		ret.RemoteTsV6Filter = in["remote_ts_v6_filter"].(string)
		ret.InSpiFilter = in["in_spi_filter"].(string)
		ret.OutSpiFilter = in["out_spi_filter"].(string)
		ret.SaList = getSliceVpnOperIpsecListOperSaList(in["sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIpsecListOperSaList(d []interface{}) []edpt.VpnOperIpsecListOperSaList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIpsecListOperSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIpsecListOperSaList
		oi.Status = in["status"].(string)
		oi.SaIndex = in["sa_index"].(int)
		oi.TsProto = in["ts_proto"].(int)
		oi.LocalIp = in["local_ip"].(string)
		oi.LocalPort = in["local_port"].(int)
		oi.PeerIp = in["peer_ip"].(string)
		oi.PeerPort = in["peer_port"].(int)
		oi.LocalSpi = in["local_spi"].(string)
		oi.RemoteSpi = in["remote_spi"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Mode = in["mode"].(string)
		oi.EncryptionAlgorithm = in["encryption_algorithm"].(string)
		oi.HashAlgorithm = in["hash_algorithm"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Lifebytes = in["lifebytes"].(string)
		oi.DhGroup = in["dh_group"].(int)
		oi.NatTraversal = in["nat_traversal"].(int)
		oi.AntiReplay = in["anti_replay"].(string)
		oi.PacketsEncrypted = in["packets_encrypted"].(int)
		oi.PacketsDecrypted = in["packets_decrypted"].(int)
		oi.AntiReplayNum = in["anti_replay_num"].(int)
		oi.RekeyNum = in["rekey_num"].(int)
		oi.PacketsErrInactive = in["packets_err_inactive"].(int)
		oi.PacketsErrEncryption = in["packets_err_encryption"].(int)
		oi.PacketsErrPadCheck = in["packets_err_pad_check"].(int)
		oi.PacketsErrPktSanity = in["packets_err_pkt_sanity"].(int)
		oi.PacketsErrIcvCheck = in["packets_err_icv_check"].(int)
		oi.PacketsErrLifetimeLifebytes = getObjectVpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes(in["packets_err_lifetime_lifebytes"].([]interface{}))
		oi.BytesEncrypted = in["bytes_encrypted"].(int)
		oi.BytesDecrypted = in["bytes_decrypted"].(int)
		oi.PrefragSuccess = in["prefrag_success"].(int)
		oi.PrefragError = in["prefrag_error"].(int)
		oi.CaviumBytesEncrypted = in["cavium_bytes_encrypted"].(int)
		oi.CaviumBytesDecrypted = in["cavium_bytes_decrypted"].(int)
		oi.CaviumPacketsEncrypted = in["cavium_packets_encrypted"].(int)
		oi.CaviumPacketsDecrypted = in["cavium_packets_decrypted"].(int)
		oi.QatBytesEncrypted = in["qat_bytes_encrypted"].(int)
		oi.QatBytesDecrypted = in["qat_bytes_decrypted"].(int)
		oi.QatPacketsEncrypted = in["qat_packets_encrypted"].(int)
		oi.QatPacketsDecrypted = in["qat_packets_decrypted"].(int)
		oi.TunnelIntfDown = in["tunnel_intf_down"].(int)
		oi.PktFailPrepToSend = in["pkt_fail_prep_to_send"].(int)
		oi.NoNextHop = in["no_next_hop"].(int)
		oi.InvalidTunnelId = in["invalid_tunnel_id"].(int)
		oi.NoTunnelFound = in["no_tunnel_found"].(int)
		oi.PktFailToSend = in["pkt_fail_to_send"].(int)
		oi.FragAfterEncapFragPackets = in["frag_after_encap_frag_packets"].(int)
		oi.FragReceived = in["frag_received"].(int)
		oi.SequenceNum = in["sequence_num"].(int)
		oi.SequenceNumRollover = in["sequence_num_rollover"].(int)
		oi.PacketsErrNhCheck = in["packets_err_nh_check"].(int)
		oi.EnforceTsEncapDrop = in["enforce_ts_encap_drop"].(int)
		oi.EnforceTsDecapDrop = in["enforce_ts_decap_drop"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes(d []interface{}) edpt.VpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes {

	var ret edpt.VpnOperIpsecListOperSaListPacketsErrLifetimeLifebytes
	return ret
}

func getObjectVpnOperIpsecSa(d []interface{}) edpt.VpnOperIpsecSa {

	count1 := len(d)
	var ret edpt.VpnOperIpsecSa
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIpsecSaOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIpsecSaOper(d []interface{}) edpt.VpnOperIpsecSaOper {

	count1 := len(d)
	var ret edpt.VpnOperIpsecSaOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpsecSaList = getSliceVpnOperIpsecSaOperIpsecSaList(in["ipsec_sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIpsecSaOperIpsecSaList(d []interface{}) []edpt.VpnOperIpsecSaOperIpsecSaList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIpsecSaOperIpsecSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIpsecSaOperIpsecSaList
		oi.IpsecSaName = in["ipsec_sa_name"].(string)
		oi.IkeGatewayName = in["ike_gateway_name"].(string)
		oi.LocalTs = in["local_ts"].(string)
		oi.RemoteTs = in["remote_ts"].(string)
		oi.InSpi = in["in_spi"].(string)
		oi.OutSpi = in["out_spi"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Mode = in["mode"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Lifebytes = in["lifebytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIpsecSaByGw(d []interface{}) edpt.VpnOperIpsecSaByGw {

	count1 := len(d)
	var ret edpt.VpnOperIpsecSaByGw
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIpsecSaByGwOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIpsecSaByGwOper(d []interface{}) edpt.VpnOperIpsecSaByGwOper {

	count1 := len(d)
	var ret edpt.VpnOperIpsecSaByGwOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeGatewayName = in["ike_gateway_name"].(string)
		ret.LocalIp = in["local_ip"].(string)
		ret.PeerIp = in["peer_ip"].(string)
		ret.IpsecSaList = getSliceVpnOperIpsecSaByGwOperIpsecSaList(in["ipsec_sa_list"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIpsecSaByGwOperIpsecSaList(d []interface{}) []edpt.VpnOperIpsecSaByGwOperIpsecSaList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIpsecSaByGwOperIpsecSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIpsecSaByGwOperIpsecSaList
		oi.IpsecSaName = in["ipsec_sa_name"].(string)
		oi.LocalTs = in["local_ts"].(string)
		oi.RemoteTs = in["remote_ts"].(string)
		oi.InSpi = in["in_spi"].(string)
		oi.OutSpi = in["out_spi"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.Mode = in["mode"].(string)
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Lifetime = in["lifetime"].(int)
		oi.Lifebytes = in["lifebytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperIpsecSaClients(d []interface{}) edpt.VpnOperIpsecSaClients {

	count1 := len(d)
	var ret edpt.VpnOperIpsecSaClients
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperIpsecSaClientsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperIpsecSaClientsOper(d []interface{}) edpt.VpnOperIpsecSaClientsOper {

	count1 := len(d)
	var ret edpt.VpnOperIpsecSaClientsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpsecClients = getSliceVpnOperIpsecSaClientsOperIpsecClients(in["ipsec_clients"].([]interface{}))
	}
	return ret
}

func getSliceVpnOperIpsecSaClientsOperIpsecClients(d []interface{}) []edpt.VpnOperIpsecSaClientsOperIpsecClients {

	count1 := len(d)
	ret := make([]edpt.VpnOperIpsecSaClientsOperIpsecClients, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIpsecSaClientsOperIpsecClients
		oi.IpsecClientsIp = in["ipsec_clients_ip"].(string)
		oi.SaList = getSliceVpnOperIpsecSaClientsOperIpsecClientsSaList(in["sa_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnOperIpsecSaClientsOperIpsecClientsSaList(d []interface{}) []edpt.VpnOperIpsecSaClientsOperIpsecClientsSaList {

	count1 := len(d)
	ret := make([]edpt.VpnOperIpsecSaClientsOperIpsecClientsSaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperIpsecSaClientsOperIpsecClientsSaList
		oi.Name = in["name"].(string)
		oi.LocalTs = in["local_ts"].(string)
		oi.InSpi = in["in_spi"].(string)
		oi.OutSpi = in["out_spi"].(string)
		oi.Lifetime = in["lifetime"].(string)
		oi.Lifebytes = in["lifebytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperLog(d []interface{}) edpt.VpnOperLog {

	count1 := len(d)
	var ret edpt.VpnOperLog
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperLogOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperLogOper(d []interface{}) edpt.VpnOperLogOper {

	count1 := len(d)
	var ret edpt.VpnOperLogOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VpnLogList = getSliceVpnOperLogOperVpnLogList(in["vpn_log_list"].([]interface{}))
		ret.VpnLogOffset = in["vpn_log_offset"].(int)
		ret.VpnLogOver = in["vpn_log_over"].(int)
		ret.Follow = in["follow"].(int)
		ret.FromStart = in["from_start"].(int)
		ret.NumLines = in["num_lines"].(int)
	}
	return ret
}

func getSliceVpnOperLogOperVpnLogList(d []interface{}) []edpt.VpnOperLogOperVpnLogList {

	count1 := len(d)
	ret := make([]edpt.VpnOperLogOperVpnLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperLogOperVpnLogList
		oi.VpnLogData = in["vpn_log_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperOcsp(d []interface{}) edpt.VpnOperOcsp {

	count1 := len(d)
	var ret edpt.VpnOperOcsp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVpnOperOcspOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVpnOperOcspOper(d []interface{}) edpt.VpnOperOcspOper {

	count1 := len(d)
	var ret edpt.VpnOperOcspOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OcspList = getSliceVpnOperOcspOperOcspList(in["ocsp_list"].([]interface{}))
		ret.TotalOcsps = in["total_ocsps"].(int)
	}
	return ret
}

func getSliceVpnOperOcspOperOcspList(d []interface{}) []edpt.VpnOperOcspOperOcspList {

	count1 := len(d)
	ret := make([]edpt.VpnOperOcspOperOcspList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperOcspOperOcspList
		oi.Subject = in["subject"].(string)
		oi.Issuer = in["issuer"].(string)
		oi.Validity = in["validity"].(string)
		oi.CertificateStatus = in["certificate_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnOperOper(d []interface{}) edpt.VpnOperOper {

	count1 := len(d)
	var ret edpt.VpnOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeGatewayTotal = in["ike_gateway_total"].(int)
		ret.IpsecTotal = in["ipsec_total"].(int)
		ret.IkeSaTotal = in["ike_sa_total"].(int)
		ret.IpsecSaTotal = in["ipsec_sa_total"].(int)
		ret.IpsecMode = in["ipsec_mode"].(string)
		ret.NumHardwareDevices = in["num_hardware_devices"].(int)
		ret.CryptoCoresTotal = in["crypto_cores_total"].(int)
		ret.CryptoCoresAssignedToIpsec = in["crypto_cores_assigned_to_ipsec"].(int)
		ret.CryptoMem = in["crypto_mem"].(int)
		ret.AllPartitionList = getSliceVpnOperOperAllPartitionList(in["all_partition_list"].([]interface{}))
		ret.AllPartitions = in["all_partitions"].(int)
		ret.Shared = in["shared"].(int)
		ret.SpecificPartition = in["specific_partition"].(string)
	}
	return ret
}

func getSliceVpnOperOperAllPartitionList(d []interface{}) []edpt.VpnOperOperAllPartitionList {

	count1 := len(d)
	ret := make([]edpt.VpnOperOperAllPartitionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperOperAllPartitionList
		oi.IkeGatewayTotal = in["ike_gateway_total"].(int)
		oi.IpsecTotal = in["ipsec_total"].(int)
		oi.IkeSaTotal = in["ike_sa_total"].(int)
		oi.IpsecSaTotal = in["ipsec_sa_total"].(int)
		oi.IpsecStateless = in["ipsec_stateless"].(int)
		oi.IpsecMode = in["ipsec_mode"].(string)
		oi.IpsecHardwareType = in["ipsec_hardware_type"].(string)
		oi.NumHardwareDevices = in["num_hardware_devices"].(int)
		oi.IkeHardwareAccelerate = in["ike_hardware_accelerate"].(string)
		oi.CryptoCoresTotal = in["crypto_cores_total"].(int)
		oi.CryptoCoresAssignedToIpsec = in["crypto_cores_assigned_to_ipsec"].(int)
		oi.CryptoMem = in["crypto_mem"].(int)
		oi.CryptoHwErr = in["crypto_hw_err"].(int)
		oi.CryptoHwErrReqAllocFail = in["crypto_hw_err_req_alloc_fail"].(int)
		oi.CryptoHwErrEnqueueFail = in["crypto_hw_err_enqueue_fail"].(int)
		oi.CryptoHwErrSgBuffAllocFail = in["crypto_hw_err_sg_buff_alloc_fail"].(int)
		oi.CryptoHwErrBadPointer = in["crypto_hw_err_bad_pointer"].(int)
		oi.CryptoHwErrBadCtxPointer = in["crypto_hw_err_bad_ctx_pointer"].(int)
		oi.CryptoHwErrReqError = in["crypto_hw_err_req_error"].(int)
		oi.CryptoHwErrStateError = in["crypto_hw_err_state_error"].(int)
		oi.CryptoHwErrState = in["crypto_hw_err_state"].(string)
		oi.CryptoHwErrTimeOut = in["crypto_hw_err_time_out"].(int)
		oi.CryptoHwErrTimeOutState = in["crypto_hw_err_time_out_state"].(int)
		oi.CryptoHwErrBuffAllocError = in["crypto_hw_err_buff_alloc_error"].(int)
		oi.PassthroughTotal = in["passthrough_total"].(int)
		oi.VpnList = getSliceVpnOperOperAllPartitionListVpnList(in["vpn_list"].([]interface{}))
		oi.StandbyDrop = in["standby_drop"].(int)
		oi.PartitionName = in["partition_name"].(string)
		oi.SignedAuthFlag = in["signed_auth_flag"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVpnOperOperAllPartitionListVpnList(d []interface{}) []edpt.VpnOperOperAllPartitionListVpnList {

	count1 := len(d)
	ret := make([]edpt.VpnOperOperAllPartitionListVpnList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnOperOperAllPartitionListVpnList
		oi.Passthrough = in["passthrough"].(int)
		oi.CpuId = in["cpu_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnOper(d *schema.ResourceData) edpt.VpnOper {
	var ret edpt.VpnOper

	ret.Crl = getObjectVpnOperCrl(d.Get("crl").([]interface{}))

	ret.Default = getObjectVpnOperDefault(d.Get("default").([]interface{}))

	ret.Errordump = getObjectVpnOperErrordump(d.Get("errordump").([]interface{}))

	ret.GroupList = getObjectVpnOperGroupList(d.Get("group_list").([]interface{}))

	ret.IkeGatewayList = getSliceVpnOperIkeGatewayList(d.Get("ike_gateway_list").([]interface{}))

	ret.IkeSa = getObjectVpnOperIkeSa(d.Get("ike_sa").([]interface{}))

	ret.IkeSaBrief = getObjectVpnOperIkeSaBrief(d.Get("ike_sa_brief").([]interface{}))

	ret.IkeSaClients = getObjectVpnOperIkeSaClients(d.Get("ike_sa_clients").([]interface{}))

	ret.IkeStatsByGw = getObjectVpnOperIkeStatsByGw(d.Get("ike_stats_by_gw").([]interface{}))

	ret.IpsecList = getSliceVpnOperIpsecList(d.Get("ipsec_list").([]interface{}))

	ret.IpsecSa = getObjectVpnOperIpsecSa(d.Get("ipsec_sa").([]interface{}))

	ret.IpsecSaByGw = getObjectVpnOperIpsecSaByGw(d.Get("ipsec_sa_by_gw").([]interface{}))

	ret.IpsecSaClients = getObjectVpnOperIpsecSaClients(d.Get("ipsec_sa_clients").([]interface{}))

	ret.Log = getObjectVpnOperLog(d.Get("log").([]interface{}))

	ret.Ocsp = getObjectVpnOperOcsp(d.Get("ocsp").([]interface{}))

	ret.Oper = getObjectVpnOperOper(d.Get("oper").([]interface{}))
	return ret
}
