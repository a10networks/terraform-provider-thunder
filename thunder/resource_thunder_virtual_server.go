package thunder

//Thunder resource VirtualServer

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVirtualServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVirtualServerCreate,
		UpdateContext: resourceVirtualServerUpdate,
		ReadContext:   resourceVirtualServerRead,
		DeleteContext: resourceVirtualServerDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"netmask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_acl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipv6_acl_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"acl_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_id_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"acl_name_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"use_if_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"enable_disable_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"redistribution_flagged": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vport_disable_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"suppress_internal_loopback": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"arp_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_policy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_policy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_virtual_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_vs_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_virtual_server_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_scaleout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vrid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_vip_adv": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ha_dynamic": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"redistribute_route_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"migrate_vip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_data_cpu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"target_floating_ipv4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"target_floating_ipv6": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cancel_migration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"finish_migration": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"range": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"proxy_layer": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"optimization_level": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"support_http2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip_only_lb": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"conn_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reset": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"no_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_alternate_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alternate_port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alt_protocol1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"serv_sel_fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"when_down": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"alt_protocol2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"req_fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"when_down_protocol2": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"l7_service_chain": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"def_selection_if_pref_failed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ha_conn_mirror": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"on_syn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"skip_rev_hash": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"message_switching": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"force_routing_mode": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"one_server_conn": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"secs": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reset_on_server_selection_fail": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"clientip_sticky_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"extended_stats": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gslb_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"view": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"snat_on_vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"stats_data_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"syn_cookie": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"expand": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"attack_detection": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_id_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_id_src_nat_pool": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_id_seq_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"shared_partition_pool_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_id_src_nat_pool_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_id_seq_num_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_acl_id_src_nat_pool": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"v_acl_id_seq_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_shared_partition_pool_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_acl_id_src_nat_pool_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"v_acl_id_seq_num_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name_src_nat_pool": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_name_seq_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"shared_partition_pool_name": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name_src_nat_pool_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"acl_name_seq_num_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_acl_name_src_nat_pool": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"v_acl_name_seq_num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_shared_partition_pool_name": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_acl_name_src_nat_pool_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"v_acl_name_seq_num_shared": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"template_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_policy_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_policy_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"aflex_scripts": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aflex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"aflex_shared": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"no_auto_up_on_aflex": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enable_scaleout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"scaleout_bucket_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"scaleout_device_group": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pool": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_pool": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"pool_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"auto": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"precedence": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip_smart_rr": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_cgnv6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"enable_playerid_check": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipinip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip_map_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"rtp_sip_call_id_match": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"use_rcv_hop_for_resp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"persist_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"use_rcv_hop_group": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"reselection": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"eth_fwd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"trunk_fwd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"eth_rev": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"trunk_rev": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_sip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"p_template_sip_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_sip_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_smpp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_smpp_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_smpp_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dblb": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_dblb_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_dblb_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_connection_reuse": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_connection_reuse_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_connection_reuse_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dns": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_dns_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_dns_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_dynamic_service": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_dynamic_service_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_dynamic_service_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_source_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_source_ip_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_source_ip_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_destination_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_destination_ip_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_destination_ip_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_ssl_sid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_ssl_sid_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_ssl_sid_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_persist_cookie": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_persist_cookie_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_persist_cookie_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_imap_pop3": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_imap_pop3_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_imap_pop3_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_smtp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_smtp_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_smtp_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_mqtt": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_http": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_http_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_http_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_http_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_http_policy_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_http_policy_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"redirect_to_https": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_external_service": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_external_service_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_external_service_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_reqmod_icap": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_respmod_icap": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_doh": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_doh_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_doh_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_server_ssl_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_server_ssl_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_client_ssl": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_client_ssl_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_client_ssl_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_server_ssh": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_client_ssh": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_udp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_udp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_udp_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_tcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_tcp_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_virtual_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_virtual_port_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_virtual_port_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_ftp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_diameter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_diameter_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_diameter_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_cache": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_cache_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_cache_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_ram_cache": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_fix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_fix_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_fix_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"waf_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_ssli": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy_client": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_tcp_proxy_template": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_tcp_proxy_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"use_default_if_no_server": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"template_scaleout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"no_dest_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_translation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l7_hardware_assist": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auth_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aaa_policy": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"cpu_compute": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"memory_compute": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"substitute_source_mac": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ignore_global": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"aflex_table_entry_syn_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"aflex_table_entry_syn_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"gtp_session_lb": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reply_acme_challenge": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"resolve_web_cat_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"packet_capture_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVirtualServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating VirtualServer (Inside resourceVirtualServerCreate) ")
		name1 := d.Get("name").(string)
		data := dataToVirtualServer(d)
		logger.Println("[INFO] received formatted data from method data to VirtualServer --")
		d.SetId(name1)
		err := go_thunder.PostVirtualServer(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceVirtualServerRead(ctx, d, meta)

	}
	return diags
}

func resourceVirtualServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading VirtualServer (Inside resourceVirtualServerRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetVirtualServer(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceVirtualServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying VirtualServer   (Inside resourceVirtualServerUpdate) ")
		data := dataToVirtualServer(d)
		logger.Println("[INFO] received formatted data from method data to VirtualServer ")
		err := go_thunder.PutVirtualServer(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceVirtualServerRead(ctx, d, meta)

	}
	return diags
}

func resourceVirtualServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceVirtualServerDelete) " + name1)
		err := go_thunder.DeleteVirtualServer(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToVirtualServer(d *schema.ResourceData) go_thunder.VirtualServer {
	var vc go_thunder.VirtualServer
	var c go_thunder.VirtualServerInstance
	c.VirtualServerInstanceName = d.Get("name").(string)
	c.VirtualServerInstanceIpv6Address = d.Get("ipv6_address").(string)
	c.VirtualServerInstanceIPAddress = d.Get("ip_address").(string)
	c.VirtualServerInstanceNetmask = d.Get("netmask").(string)
	c.VirtualServerInstanceIpv6Acl = d.Get("ipv6_acl").(string)
	c.VirtualServerInstanceIpv6AclShared = d.Get("ipv6_acl_shared").(string)
	c.VirtualServerInstanceAclID = d.Get("acl_id").(int)
	c.VirtualServerInstanceAclName = d.Get("acl_name").(string)
	c.VirtualServerInstanceAclIDShared = d.Get("acl_id_shared").(int)
	c.VirtualServerInstanceAclNameShared = d.Get("acl_name_shared").(string)
	c.VirtualServerInstanceUseIfIP = d.Get("use_if_ip").(int)
	c.VirtualServerInstanceEthernet = d.Get("ethernet").(int)
	c.VirtualServerInstanceDescription = d.Get("description").(string)
	c.VirtualServerInstanceEnableDisableAction = d.Get("enable_disable_action").(string)
	c.VirtualServerInstanceRedistributionFlagged = d.Get("redistribution_flagged").(int)
	c.VirtualServerInstanceVportDisableAction = d.Get("vport_disable_action").(string)
	c.VirtualServerInstanceSuppressInternalLoopback = d.Get("suppress_internal_loopback").(int)
	c.VirtualServerInstanceArpDisable = d.Get("arp_disable").(int)
	c.VirtualServerInstanceTemplatePolicy = d.Get("template_policy").(string)
	c.VirtualServerInstanceSharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	c.VirtualServerInstanceTemplatePolicyShared = d.Get("template_policy_shared").(string)
	c.VirtualServerInstanceTemplateVirtualServer = d.Get("template_virtual_server").(string)
	c.VirtualServerInstanceSharedPartitionVsTemplate = d.Get("shared_partition_vs_template").(int)
	c.VirtualServerInstanceTemplateVirtualServerShared = d.Get("template_virtual_server_shared").(string)
	c.VirtualServerInstanceTemplateLogging = d.Get("template_logging").(string)
	c.VirtualServerInstanceTemplateScaleout = d.Get("template_scaleout").(string)
	c.VirtualServerInstanceStatsDataAction = d.Get("stats_data_action").(string)
	c.VirtualServerInstanceExtendedStats = d.Get("extended_stats").(int)
	c.VirtualServerInstanceVrid = d.Get("vrid").(int)
	c.VirtualServerInstanceDisableVipAdv = d.Get("disable_vip_adv").(int)
	c.VirtualServerInstanceHaDynamic = d.Get("ha_dynamic").(int)
	c.VirtualServerInstanceRedistributeRouteMap = d.Get("redistribute_route_map").(string)
	c.VirtualServerInstanceUserTag = d.Get("user_tag").(string)

	var obj1 go_thunder.VirtualServerInstanceMigrateVip
	prefix1 := "migrate_vip.0."
	obj1.VirtualServerInstanceMigrateVipTargetDataCPU = d.Get(prefix1 + "target_data_cpu").(int)
	obj1.VirtualServerInstanceMigrateVipTargetFloatingIpv4 = d.Get(prefix1 + "target_floating_ipv4").(string)
	obj1.VirtualServerInstanceMigrateVipTargetFloatingIpv6 = d.Get(prefix1 + "target_floating_ipv6").(string)
	obj1.VirtualServerInstanceMigrateVipCancelMigration = d.Get(prefix1 + "cancel_migration").(int)
	obj1.VirtualServerInstanceMigrateVipFinishMigration = d.Get(prefix1 + "finish_migration").(int)

	c.VirtualServerInstanceMigrateVipTargetDataCPU = obj1

	VirtualServerInstancePortListCount := d.Get("port_list.#").(int)
	c.VirtualServerInstancePortListPortNumber = make([]go_thunder.VirtualServerInstancePortList, 0, VirtualServerInstancePortListCount)

	for i := 0; i < VirtualServerInstancePortListCount; i++ {
		var obj2 go_thunder.VirtualServerInstancePortList
		prefix2 := fmt.Sprintf("port_list.%d.", i)
		obj2.VirtualServerInstancePortListPortNumber = d.Get(prefix2 + "port_number").(int)
		obj2.VirtualServerInstancePortListProtocol = d.Get(prefix2 + "protocol").(string)
		obj2.VirtualServerInstancePortListRange = d.Get(prefix2 + "range").(int)
		obj2.VirtualServerInstancePortListAlternatePort = d.Get(prefix2 + "alternate_port").(int)
		obj2.VirtualServerInstancePortListProxyLayer = d.Get(prefix2 + "proxy_layer").(string)
		obj2.VirtualServerInstancePortListOptimizationLevel = d.Get(prefix2 + "optimization_level").(string)
		obj2.VirtualServerInstancePortListSupportHTTP2 = d.Get(prefix2 + "support_http2").(int)
		obj2.VirtualServerInstancePortListIPOnlyLb = d.Get(prefix2 + "ip_only_lb").(int)
		obj2.VirtualServerInstancePortListName = d.Get(prefix2 + "name").(string)
		obj2.VirtualServerInstancePortListConnLimit = d.Get(prefix2 + "conn_limit").(int)
		obj2.VirtualServerInstancePortListReset = d.Get(prefix2 + "reset").(int)
		obj2.VirtualServerInstancePortListNoLogging = d.Get(prefix2 + "no_logging").(int)
		obj2.VirtualServerInstancePortListUseAlternatePort = d.Get(prefix2 + "use_alternate_port").(int)
		obj2.VirtualServerInstancePortListAlternatePortNumber = d.Get(prefix2 + "alternate_port_number").(int)
		obj2.VirtualServerInstancePortListAltProtocol1 = d.Get(prefix2 + "alt_protocol1").(string)
		obj2.VirtualServerInstancePortListServSelFail = d.Get(prefix2 + "serv_sel_fail").(int)
		obj2.VirtualServerInstancePortListWhenDown = d.Get(prefix2 + "when_down").(int)
		obj2.VirtualServerInstancePortListAltProtocol2 = d.Get(prefix2 + "alt_protocol2").(string)
		obj2.VirtualServerInstancePortListReqFail = d.Get(prefix2 + "req_fail").(int)
		obj2.VirtualServerInstancePortListWhenDownProtocol2 = d.Get(prefix2 + "when_down_protocol2").(int)
		obj2.VirtualServerInstancePortListAction = d.Get(prefix2 + "action").(string)
		obj2.VirtualServerInstancePortListL7ServiceChain = d.Get(prefix2 + "l7_service_chain").(int)
		obj2.VirtualServerInstancePortListDefSelectionIfPrefFailed = d.Get(prefix2 + "def_selection_if_pref_failed").(string)
		obj2.VirtualServerInstancePortListHaConnMirror = d.Get(prefix2 + "ha_conn_mirror").(int)
		obj2.VirtualServerInstancePortListOnSyn = d.Get(prefix2 + "on_syn").(int)
		obj2.VirtualServerInstancePortListSkipRevHash = d.Get(prefix2 + "skip_rev_hash").(int)
		obj2.VirtualServerInstancePortListMessageSwitching = d.Get(prefix2 + "message_switching").(int)
		obj2.VirtualServerInstancePortListForceRoutingMode = d.Get(prefix2 + "force_routing_mode").(int)
		obj2.VirtualServerInstancePortListOneServerConn = d.Get(prefix2 + "one_server_conn").(int)
		obj2.VirtualServerInstancePortListRate = d.Get(prefix2 + "rate").(int)
		obj2.VirtualServerInstancePortListSecs = d.Get(prefix2 + "secs").(int)
		obj2.VirtualServerInstancePortListResetOnServerSelectionFail = d.Get(prefix2 + "reset_on_server_selection_fail").(int)
		obj2.VirtualServerInstancePortListClientipStickyNat = d.Get(prefix2 + "clientip_sticky_nat").(int)
		obj2.VirtualServerInstancePortListExtendedStats = d.Get(prefix2 + "extended_stats").(int)
		obj2.VirtualServerInstancePortListGslbEnable = d.Get(prefix2 + "gslb_enable").(int)
		obj2.VirtualServerInstancePortListView = d.Get(prefix2 + "view").(int)
		obj2.VirtualServerInstancePortListSnatOnVip = d.Get(prefix2 + "snat_on_vip").(int)
		obj2.VirtualServerInstancePortListStatsDataAction = d.Get(prefix2 + "stats_data_action").(string)
		obj2.VirtualServerInstancePortListSynCookie = d.Get(prefix2 + "syn_cookie").(int)
		obj2.VirtualServerInstancePortListExpand = d.Get(prefix2 + "expand").(int)
		obj2.VirtualServerInstancePortListAttackDetection = d.Get(prefix2 + "attack_detection").(int)

		VirtualServerInstancePortListAclListCount := d.Get(prefix2 + "acl_list.#").(int)
		obj2.VirtualServerInstancePortListAclListAclID = make([]go_thunder.VirtualServerInstancePortListAclList, 0, VirtualServerInstancePortListAclListCount)

		for i := 0; i < VirtualServerInstancePortListAclListCount; i++ {
			var obj2_1 go_thunder.VirtualServerInstancePortListAclList
			prefix2_1 := prefix2 + fmt.Sprintf("acl_list.%d.", i)
			obj2_1.VirtualServerInstancePortListAclListAclID = d.Get(prefix2_1 + "acl_id").(int)
			obj2_1.VirtualServerInstancePortListAclListAclName = d.Get(prefix2_1 + "acl_name").(string)
			obj2_1.VirtualServerInstancePortListAclListAclIDShared = d.Get(prefix2_1 + "acl_id_shared").(int)
			obj2_1.VirtualServerInstancePortListAclListAclNameShared = d.Get(prefix2_1 + "acl_name_shared").(string)
			obj2_1.VirtualServerInstancePortListAclListAclIDSrcNatPool = d.Get(prefix2_1 + "acl_id_src_nat_pool").(string)
			obj2_1.VirtualServerInstancePortListAclListAclIDSeqNum = d.Get(prefix2_1 + "acl_id_seq_num").(int)
			obj2_1.VirtualServerInstancePortListAclListSharedPartitionPoolID = d.Get(prefix2_1 + "shared_partition_pool_id").(int)
			obj2_1.VirtualServerInstancePortListAclListAclIDSrcNatPoolShared = d.Get(prefix2_1 + "acl_id_src_nat_pool_shared").(string)
			obj2_1.VirtualServerInstancePortListAclListAclIDSeqNumShared = d.Get(prefix2_1 + "acl_id_seq_num_shared").(int)
			obj2_1.VirtualServerInstancePortListAclListVAclIDSrcNatPool = d.Get(prefix2_1 + "v_acl_id_src_nat_pool").(string)
			obj2_1.VirtualServerInstancePortListAclListVAclIDSeqNum = d.Get(prefix2_1 + "v_acl_id_seq_num").(int)
			obj2_1.VirtualServerInstancePortListAclListVSharedPartitionPoolID = d.Get(prefix2_1 + "v_shared_partition_pool_id").(int)
			obj2_1.VirtualServerInstancePortListAclListVAclIDSrcNatPoolShared = d.Get(prefix2_1 + "v_acl_id_src_nat_pool_shared").(string)
			obj2_1.VirtualServerInstancePortListAclListVAclIDSeqNumShared = d.Get(prefix2_1 + "v_acl_id_seq_num_shared").(int)
			obj2_1.VirtualServerInstancePortListAclListAclNameSrcNatPool = d.Get(prefix2_1 + "acl_name_src_nat_pool").(string)
			obj2_1.VirtualServerInstancePortListAclListAclNameSeqNum = d.Get(prefix2_1 + "acl_name_seq_num").(int)
			obj2_1.VirtualServerInstancePortListAclListSharedPartitionPoolName = d.Get(prefix2_1 + "shared_partition_pool_name").(int)
			obj2_1.VirtualServerInstancePortListAclListAclNameSrcNatPoolShared = d.Get(prefix2_1 + "acl_name_src_nat_pool_shared").(string)
			obj2_1.VirtualServerInstancePortListAclListAclNameSeqNumShared = d.Get(prefix2_1 + "acl_name_seq_num_shared").(int)
			obj2_1.VirtualServerInstancePortListAclListVAclNameSrcNatPool = d.Get(prefix2_1 + "v_acl_name_src_nat_pool").(string)
			obj2_1.VirtualServerInstancePortListAclListVAclNameSeqNum = d.Get(prefix2_1 + "v_acl_name_seq_num").(int)
			obj2_1.VirtualServerInstancePortListAclListVSharedPartitionPoolName = d.Get(prefix2_1 + "v_shared_partition_pool_name").(int)
			obj2_1.VirtualServerInstancePortListAclListVAclNameSrcNatPoolShared = d.Get(prefix2_1 + "v_acl_name_src_nat_pool_shared").(string)
			obj2_1.VirtualServerInstancePortListAclListVAclNameSeqNumShared = d.Get(prefix2_1 + "v_acl_name_seq_num_shared").(int)
			obj2.VirtualServerInstancePortListAclListAclID = append(obj2.VirtualServerInstancePortListAclListAclID, obj2_1)
		}

		obj2.VirtualServerInstancePortListTemplatePolicy = d.Get(prefix2 + "template_policy").(string)
		obj2.VirtualServerInstancePortListSharedPartitionPolicyTemplate = d.Get(prefix2 + "shared_partition_policy_template").(int)
		obj2.VirtualServerInstancePortListTemplatePolicyShared = d.Get(prefix2 + "template_policy_shared").(string)

		VirtualServerInstancePortListAflexScriptsCount := d.Get(prefix2 + "aflex_scripts.#").(int)
		obj2.VirtualServerInstancePortListAflexScriptsAflex = make([]go_thunder.VirtualServerInstancePortListAflexScripts, 0, VirtualServerInstancePortListAflexScriptsCount)

		for i := 0; i < VirtualServerInstancePortListAflexScriptsCount; i++ {
			var obj2_2 go_thunder.VirtualServerInstancePortListAflexScripts
			prefix2_2 := prefix2 + fmt.Sprintf("aflex_scripts.%d.", i)
			obj2_2.VirtualServerInstancePortListAflexScriptsAflex = d.Get(prefix2_2 + "aflex").(string)
			obj2_2.VirtualServerInstancePortListAflexScriptsAflexShared = d.Get(prefix2_2 + "aflex_shared").(string)
			obj2.VirtualServerInstancePortListAflexScriptsAflex = append(obj2.VirtualServerInstancePortListAflexScriptsAflex, obj2_2)
		}

		obj2.VirtualServerInstancePortListNoAutoUpOnAflex = d.Get(prefix2 + "no_auto_up_on_aflex").(int)
		obj2.VirtualServerInstancePortListEnableScaleout = d.Get(prefix2 + "enable_scaleout").(int)
		obj2.VirtualServerInstancePortListScaleoutBucketCount = d.Get(prefix2 + "scaleout_bucket_count").(int)
		obj2.VirtualServerInstancePortListScaleoutDeviceGroup = d.Get(prefix2 + "scaleout_device_group").(int)
		obj2.VirtualServerInstancePortListPool = d.Get(prefix2 + "pool").(string)
		obj2.VirtualServerInstancePortListSharedPartitionPool = d.Get(prefix2 + "shared_partition_pool").(int)
		obj2.VirtualServerInstancePortListPoolShared = d.Get(prefix2 + "pool_shared").(string)
		obj2.VirtualServerInstancePortListAuto = d.Get(prefix2 + "auto").(int)
		obj2.VirtualServerInstancePortListPrecedence = d.Get(prefix2 + "precedence").(int)
		obj2.VirtualServerInstancePortListIPSmartRr = d.Get(prefix2 + "ip_smart_rr").(int)
		obj2.VirtualServerInstancePortListUseCgnv6 = d.Get(prefix2 + "use_cgnv6").(int)
		obj2.VirtualServerInstancePortListEnablePlayeridCheck = d.Get(prefix2 + "enable_playerid_check").(int)
		obj2.VirtualServerInstancePortListServiceGroup = d.Get(prefix2 + "service_group").(string)
		obj2.VirtualServerInstancePortListIpinip = d.Get(prefix2 + "ipinip").(int)
		obj2.VirtualServerInstancePortListIPMapList = d.Get(prefix2 + "ip_map_list").(string)
		obj2.VirtualServerInstancePortListRtpSipCallIDMatch = d.Get(prefix2 + "rtp_sip_call_id_match").(int)
		obj2.VirtualServerInstancePortListUseRcvHopForResp = d.Get(prefix2 + "use_rcv_hop_for_resp").(int)
		obj2.VirtualServerInstancePortListPersistType = d.Get(prefix2 + "persist_type").(string)
		obj2.VirtualServerInstancePortListUseRcvHopGroup = d.Get(prefix2 + "use_rcv_hop_group").(int)
		obj2.VirtualServerInstancePortListServerGroup = d.Get(prefix2 + "server_group").(string)
		obj2.VirtualServerInstancePortListReselection = d.Get(prefix2 + "reselection").(string)
		obj2.VirtualServerInstancePortListEthFwd = d.Get(prefix2 + "eth_fwd").(int)
		obj2.VirtualServerInstancePortListTrunkFwd = d.Get(prefix2 + "trunk_fwd").(int)
		obj2.VirtualServerInstancePortListEthRev = d.Get(prefix2 + "eth_rev").(int)
		obj2.VirtualServerInstancePortListTrunkRev = d.Get(prefix2 + "trunk_rev").(int)
		obj2.VirtualServerInstancePortListTemplateSip = d.Get(prefix2 + "template_sip").(string)
		obj2.VirtualServerInstancePortListPTemplateSipShared = d.Get(prefix2 + "p_template_sip_shared").(int)
		obj2.VirtualServerInstancePortListTemplateSipShared = d.Get(prefix2 + "template_sip_shared").(string)
		obj2.VirtualServerInstancePortListTemplateSmpp = d.Get(prefix2 + "template_smpp").(string)
		obj2.VirtualServerInstancePortListSharedPartitionSmppTemplate = d.Get(prefix2 + "shared_partition_smpp_template").(int)
		obj2.VirtualServerInstancePortListTemplateSmppShared = d.Get(prefix2 + "template_smpp_shared").(string)
		obj2.VirtualServerInstancePortListTemplateDblb = d.Get(prefix2 + "template_dblb").(string)
		obj2.VirtualServerInstancePortListSharedPartitionDblbTemplate = d.Get(prefix2 + "shared_partition_dblb_template").(int)
		obj2.VirtualServerInstancePortListTemplateDblbShared = d.Get(prefix2 + "template_dblb_shared").(string)
		obj2.VirtualServerInstancePortListTemplateConnectionReuse = d.Get(prefix2 + "template_connection_reuse").(string)
		obj2.VirtualServerInstancePortListSharedPartitionConnectionReuseTemplate = d.Get(prefix2 + "shared_partition_connection_reuse_template").(int)
		obj2.VirtualServerInstancePortListTemplateConnectionReuseShared = d.Get(prefix2 + "template_connection_reuse_shared").(string)
		obj2.VirtualServerInstancePortListTemplateDNS = d.Get(prefix2 + "template_dns").(string)
		obj2.VirtualServerInstancePortListSharedPartitionDNSTemplate = d.Get(prefix2 + "shared_partition_dns_template").(int)
		obj2.VirtualServerInstancePortListTemplateDNSShared = d.Get(prefix2 + "template_dns_shared").(string)
		obj2.VirtualServerInstancePortListTemplateDynamicService = d.Get(prefix2 + "template_dynamic_service").(string)
		obj2.VirtualServerInstancePortListSharedPartitionDynamicServiceTemplate = d.Get(prefix2 + "shared_partition_dynamic_service_template").(int)
		obj2.VirtualServerInstancePortListTemplateDynamicServiceShared = d.Get(prefix2 + "template_dynamic_service_shared").(string)
		obj2.VirtualServerInstancePortListTemplatePersistSourceIP = d.Get(prefix2 + "template_persist_source_ip").(string)
		obj2.VirtualServerInstancePortListSharedPartitionPersistSourceIPTemplate = d.Get(prefix2 + "shared_partition_persist_source_ip_template").(int)
		obj2.VirtualServerInstancePortListTemplatePersistSourceIPShared = d.Get(prefix2 + "template_persist_source_ip_shared").(string)
		obj2.VirtualServerInstancePortListTemplatePersistDestinationIP = d.Get(prefix2 + "template_persist_destination_ip").(string)
		obj2.VirtualServerInstancePortListSharedPartitionPersistDestinationIPTemplate = d.Get(prefix2 + "shared_partition_persist_destination_ip_template").(int)
		obj2.VirtualServerInstancePortListTemplatePersistDestinationIPShared = d.Get(prefix2 + "template_persist_destination_ip_shared").(string)
		obj2.VirtualServerInstancePortListTemplatePersistSslSid = d.Get(prefix2 + "template_persist_ssl_sid").(string)
		obj2.VirtualServerInstancePortListSharedPartitionPersistSslSidTemplate = d.Get(prefix2 + "shared_partition_persist_ssl_sid_template").(int)
		obj2.VirtualServerInstancePortListTemplatePersistSslSidShared = d.Get(prefix2 + "template_persist_ssl_sid_shared").(string)
		obj2.VirtualServerInstancePortListTemplatePersistCookie = d.Get(prefix2 + "template_persist_cookie").(string)
		obj2.VirtualServerInstancePortListSharedPartitionPersistCookieTemplate = d.Get(prefix2 + "shared_partition_persist_cookie_template").(int)
		obj2.VirtualServerInstancePortListTemplatePersistCookieShared = d.Get(prefix2 + "template_persist_cookie_shared").(string)
		obj2.VirtualServerInstancePortListTemplateImapPop3 = d.Get(prefix2 + "template_imap_pop3").(string)
		obj2.VirtualServerInstancePortListSharedPartitionImapPop3Template = d.Get(prefix2 + "shared_partition_imap_pop3_template").(int)
		obj2.VirtualServerInstancePortListTemplateImapPop3Shared = d.Get(prefix2 + "template_imap_pop3_shared").(string)
		obj2.VirtualServerInstancePortListTemplateSMTP = d.Get(prefix2 + "template_smtp").(string)
		obj2.VirtualServerInstancePortListSharedPartitionSMTPTemplate = d.Get(prefix2 + "shared_partition_smtp_template").(int)
		obj2.VirtualServerInstancePortListTemplateSMTPShared = d.Get(prefix2 + "template_smtp_shared").(string)
		obj2.VirtualServerInstancePortListTemplateMqtt = d.Get(prefix2 + "template_mqtt").(string)
		obj2.VirtualServerInstancePortListTemplateHTTP = d.Get(prefix2 + "template_http").(string)
		obj2.VirtualServerInstancePortListSharedPartitionHTTPTemplate = d.Get(prefix2 + "shared_partition_http_template").(int)
		obj2.VirtualServerInstancePortListTemplateHTTPShared = d.Get(prefix2 + "template_http_shared").(string)
		obj2.VirtualServerInstancePortListTemplateHTTPPolicy = d.Get(prefix2 + "template_http_policy").(string)
		obj2.VirtualServerInstancePortListSharedPartitionHTTPPolicyTemplate = d.Get(prefix2 + "shared_partition_http_policy_template").(int)
		obj2.VirtualServerInstancePortListTemplateHTTPPolicyShared = d.Get(prefix2 + "template_http_policy_shared").(string)
		obj2.VirtualServerInstancePortListRedirectToHTTPS = d.Get(prefix2 + "redirect_to_https").(int)
		obj2.VirtualServerInstancePortListTemplateExternalService = d.Get(prefix2 + "template_external_service").(string)
		obj2.VirtualServerInstancePortListSharedPartitionExternalServiceTemplate = d.Get(prefix2 + "shared_partition_external_service_template").(int)
		obj2.VirtualServerInstancePortListTemplateExternalServiceShared = d.Get(prefix2 + "template_external_service_shared").(string)
		obj2.VirtualServerInstancePortListTemplateReqmodIcap = d.Get(prefix2 + "template_reqmod_icap").(string)
		obj2.VirtualServerInstancePortListTemplateRespmodIcap = d.Get(prefix2 + "template_respmod_icap").(string)
		obj2.VirtualServerInstancePortListTemplateDoh = d.Get(prefix2 + "template_doh").(string)
		obj2.VirtualServerInstancePortListSharedPartitionDohTemplate = d.Get(prefix2 + "shared_partition_doh_template").(int)
		obj2.VirtualServerInstancePortListTemplateDohShared = d.Get(prefix2 + "template_doh_shared").(string)
		obj2.VirtualServerInstancePortListTemplateServerSsl = d.Get(prefix2 + "template_server_ssl").(string)
		obj2.VirtualServerInstancePortListSharedPartitionServerSslTemplate = d.Get(prefix2 + "shared_partition_server_ssl_template").(int)
		obj2.VirtualServerInstancePortListTemplateServerSslShared = d.Get(prefix2 + "template_server_ssl_shared").(string)
		obj2.VirtualServerInstancePortListTemplateClientSsl = d.Get(prefix2 + "template_client_ssl").(string)
		obj2.VirtualServerInstancePortListSharedPartitionClientSslTemplate = d.Get(prefix2 + "shared_partition_client_ssl_template").(int)
		obj2.VirtualServerInstancePortListTemplateClientSslShared = d.Get(prefix2 + "template_client_ssl_shared").(string)
		obj2.VirtualServerInstancePortListTemplateServerSSH = d.Get(prefix2 + "template_server_ssh").(string)
		obj2.VirtualServerInstancePortListTemplateClientSSH = d.Get(prefix2 + "template_client_ssh").(string)
		obj2.VirtualServerInstancePortListTemplateUDP = d.Get(prefix2 + "template_udp").(string)
		obj2.VirtualServerInstancePortListSharedPartitionUDP = d.Get(prefix2 + "shared_partition_udp").(int)
		obj2.VirtualServerInstancePortListTemplateUDPShared = d.Get(prefix2 + "template_udp_shared").(string)
		obj2.VirtualServerInstancePortListTemplateTCP = d.Get(prefix2 + "template_tcp").(string)
		obj2.VirtualServerInstancePortListSharedPartitionTCP = d.Get(prefix2 + "shared_partition_tcp").(int)
		obj2.VirtualServerInstancePortListTemplateTCPShared = d.Get(prefix2 + "template_tcp_shared").(string)
		obj2.VirtualServerInstancePortListTemplateVirtualPort = d.Get(prefix2 + "template_virtual_port").(string)
		obj2.VirtualServerInstancePortListSharedPartitionVirtualPortTemplate = d.Get(prefix2 + "shared_partition_virtual_port_template").(int)
		obj2.VirtualServerInstancePortListTemplateVirtualPortShared = d.Get(prefix2 + "template_virtual_port_shared").(string)
		obj2.VirtualServerInstancePortListTemplateFtp = d.Get(prefix2 + "template_ftp").(string)
		obj2.VirtualServerInstancePortListTemplateDiameter = d.Get(prefix2 + "template_diameter").(string)
		obj2.VirtualServerInstancePortListSharedPartitionDiameterTemplate = d.Get(prefix2 + "shared_partition_diameter_template").(int)
		obj2.VirtualServerInstancePortListTemplateDiameterShared = d.Get(prefix2 + "template_diameter_shared").(string)
		obj2.VirtualServerInstancePortListTemplateCache = d.Get(prefix2 + "template_cache").(string)
		obj2.VirtualServerInstancePortListSharedPartitionCacheTemplate = d.Get(prefix2 + "shared_partition_cache_template").(int)
		obj2.VirtualServerInstancePortListTemplateCacheShared = d.Get(prefix2 + "template_cache_shared").(string)
		obj2.VirtualServerInstancePortListTemplateRAMCache = d.Get(prefix2 + "template_ram_cache").(string)
		obj2.VirtualServerInstancePortListTemplateFix = d.Get(prefix2 + "template_fix").(string)
		obj2.VirtualServerInstancePortListSharedPartitionFixTemplate = d.Get(prefix2 + "shared_partition_fix_template").(int)
		obj2.VirtualServerInstancePortListTemplateFixShared = d.Get(prefix2 + "template_fix_shared").(string)
		obj2.VirtualServerInstancePortListWafTemplate = d.Get(prefix2 + "waf_template").(string)
		obj2.VirtualServerInstancePortListTemplateSsli = d.Get(prefix2 + "template_ssli").(string)
		obj2.VirtualServerInstancePortListTemplateTCPProxyClient = d.Get(prefix2 + "template_tcp_proxy_client").(string)
		obj2.VirtualServerInstancePortListTemplateTCPProxyServer = d.Get(prefix2 + "template_tcp_proxy_server").(string)
		obj2.VirtualServerInstancePortListTemplateTCPProxy = d.Get(prefix2 + "template_tcp_proxy").(string)
		obj2.VirtualServerInstancePortListSharedPartitionTCPProxyTemplate = d.Get(prefix2 + "shared_partition_tcp_proxy_template").(int)
		obj2.VirtualServerInstancePortListTemplateTCPProxyShared = d.Get(prefix2 + "template_tcp_proxy_shared").(string)
		obj2.VirtualServerInstancePortListUseDefaultIfNoServer = d.Get(prefix2 + "use_default_if_no_server").(int)
		obj2.VirtualServerInstancePortListTemplateScaleout = d.Get(prefix2 + "template_scaleout").(string)
		obj2.VirtualServerInstancePortListNoDestNat = d.Get(prefix2 + "no_dest_nat").(int)
		obj2.VirtualServerInstancePortListPortTranslation = d.Get(prefix2 + "port_translation").(int)
		obj2.VirtualServerInstancePortListL7HardwareAssist = d.Get(prefix2 + "l7_hardware_assist").(int)

		var obj2_3 go_thunder.VirtualServerInstancePortListAuthCfg
		prefix2_3 := prefix2 + "auth_cfg.0."
		obj2_3.VirtualServerInstancePortListAuthCfgAaaPolicy = d.Get(prefix2_3 + "aaa_policy").(string)

		obj2.VirtualServerInstancePortListAuthCfgAaaPolicy = obj2_3

		obj2.VirtualServerInstancePortListCPUCompute = d.Get(prefix2 + "cpu_compute").(int)
		obj2.VirtualServerInstancePortListMemoryCompute = d.Get(prefix2 + "memory_compute").(int)
		obj2.VirtualServerInstancePortListSubstituteSourceMac = d.Get(prefix2 + "substitute_source_mac").(int)
		obj2.VirtualServerInstancePortListIgnoreGlobal = d.Get(prefix2 + "ignore_global").(int)
		obj2.VirtualServerInstancePortListAflexTableEntrySynDisable = d.Get(prefix2 + "aflex_table_entry_syn_disable").(int)
		obj2.VirtualServerInstancePortListAflexTableEntrySynEnable = d.Get(prefix2 + "aflex_table_entry_syn_enable").(int)
		obj2.VirtualServerInstancePortListGtpSessionLb = d.Get(prefix2 + "gtp_session_lb").(int)
		obj2.VirtualServerInstancePortListReplyAcmeChallenge = d.Get(prefix2 + "reply_acme_challenge").(int)
		obj2.VirtualServerInstancePortListResolveWebCatList = d.Get(prefix2 + "resolve_web_cat_list").(string)
		obj2.VirtualServerInstancePortListUserTag = d.Get(prefix2 + "user_tag").(string)

		VirtualServerInstancePortListSamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
		obj2.VirtualServerInstancePortListSamplingEnableCounters1 = make([]go_thunder.VirtualServerInstancePortListSamplingEnable, 0, VirtualServerInstancePortListSamplingEnableCount)

		for i := 0; i < VirtualServerInstancePortListSamplingEnableCount; i++ {
			var obj2_4 go_thunder.VirtualServerInstancePortListSamplingEnable
			prefix2_4 := prefix2 + fmt.Sprintf("sampling_enable.%d.", i)
			obj2_4.VirtualServerInstancePortListSamplingEnableCounters1 = d.Get(prefix2_4 + "counters1").(string)
			obj2.VirtualServerInstancePortListSamplingEnableCounters1 = append(obj2.VirtualServerInstancePortListSamplingEnableCounters1, obj2_4)
		}

		obj2.VirtualServerInstancePortListPacketCaptureTemplate = d.Get(prefix2 + "packet_capture_template").(string)
		c.VirtualServerInstancePortListPortNumber = append(c.VirtualServerInstancePortListPortNumber, obj2)
	}

	vc.VirtualServerInstanceName = c
	return vc
}
