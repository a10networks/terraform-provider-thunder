package thunder

//Thunder resource SlbVirtualServerPort

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPort() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbVirtualServerPortCreate,
		UpdateContext: resourceSlbVirtualServerPortUpdate,
		ReadContext:   resourceSlbVirtualServerPortRead,
		DeleteContext: resourceSlbVirtualServerPortDelete,
		Schema: map[string]*schema.Schema{
			"ha_conn_mirror": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_doh_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"precedence": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_dblb_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port_translation": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ip_map_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_reqmod_icap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"acl_name_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_name_src_nat_pool_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v_acl_name_src_nat_pool_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared_partition_pool_name": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name_seq_num_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v_shared_partition_pool_name": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name_src_nat_pool": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v_acl_name_seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_name_seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v_acl_name_src_nat_pool": {
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
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_imap_pop3_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"use_default_if_no_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_connection_reuse": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"cpu_compute": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_tcp_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_tcp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_destination_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"substitute_source_mac": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_dynamic_service_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_connection_reuse_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"when_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_client_ssl_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_persist_destination_ip_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_external_service_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persist_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_http_policy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_rcv_hop_for_resp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scaleout_bucket_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ignore_global": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"optimization_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"req_fail": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_dest_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_smpp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_diameter": {
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
			"template_ssli": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"memory_compute": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_policy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"view": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_on_server_selection_fail": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"waf_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ipinip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_auto_up_on_aflex": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gslb_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_dns_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_ssl_sid": {
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
			"template_smpp_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_sip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_dblb": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_server_ssl_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_client_ssl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"support_http2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"p_template_sip_shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_client_ssh": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_tcp_proxy_template": {
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
			"shared_partition_persist_ssl_sid_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"def_selection_if_pref_failed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_udp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"syn_cookie": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_doh": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"alternate_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alternate_port_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_persist_source_ip_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_cache": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_cookie_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rtp_sip_call_id_match": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_persist_cookie_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_file_inspection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_ftp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"serv_sel_fail": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_udp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_virtual_port_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_mqtt": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_client_ssl_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_fix_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_persist_source_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_dynamic_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"gtp_session_lb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_virtual_port_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_cgnv6": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_persist_cookie": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_smtp_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_virtual_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"trunk_fwd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_udp_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_http_policy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"snat_on_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_connection_reuse_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_tcp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"acl_id_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v_acl_id_seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_id_seq_num": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_id_src_nat_pool": {
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
						"acl_id_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v_acl_id_src_nat_pool_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"acl_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_id_src_nat_pool_shared": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v_shared_partition_pool_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"shared_partition_pool_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"v_acl_id_seq_num_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"shared_partition_dblb_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_http_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_external_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"on_syn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_smtp_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_sip_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_ssl_sid_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"force_routing_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_http_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_policy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_scaleout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"when_down_protocol2": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_fix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_smtp": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"redirect_to_https": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alt_protocol2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"alt_protocol1": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"message_switching": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_imap_pop3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"scaleout_device_group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_persist_source_ip_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l7_hardware_assist": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_fix_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_tcp_proxy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_cache_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"use_alternate_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_tcp_proxy_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"trunk_rev": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"eth_fwd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pool_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_respmod_icap": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"range": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_external_service_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_smpp_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_dynamic_service_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_server_ssh": {
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
			"template_http_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_server_ssl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_diameter_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_server_ssl_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_destination_ip_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_cache_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_tcp_proxy_client": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_tcp_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_http": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"expand": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_doh_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"skip_rev_hash": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_diameter_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"clientip_sticky_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"secs": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_imap_pop3_template": {
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
			"eth_rev": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbVirtualServerPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbVirtualServerPort (Inside resourceSlbVirtualServerPortCreate) ")
		name2 := d.Get("port-number").(int)
		name3 := d.Get("protocol").(string)
		name1 := d.Get("name").(string)
		data := dataToSlbVirtualServerPort(d)
		logger.Println("[INFO] received formatted data from method data to SlbVirtualServerPort --")
		d.SetId(name1 + "," + strconv.Itoa(name2) + "," + name3)
		err := go_thunder.PostSlbVirtualServerPort(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbVirtualServerPortRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbVirtualServerPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbVirtualServerPort (Inside resourceSlbVirtualServerPortRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Fetching service Read virtual server port")
		data, err := go_thunder.GetSlbVirtualServerPort(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found virtual server port")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbVirtualServerPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Modifying SlbVirtualServerPort   (Inside resourceSlbVirtualServerPortUpdate) ")
		data := dataToSlbVirtualServerPort(d)
		logger.Println("[INFO] received formatted data from method data to SlbVirtualServerPort ")
		err := go_thunder.PutSlbVirtualServerPort(client.Token, name1, name2, name3, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbVirtualServerPortRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbVirtualServerPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Deleting instance (Inside resourceSlbVirtualServerPortDelete) ")
		err := go_thunder.DeleteSlbVirtualServerPort(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToSlbVirtualServerPort(d *schema.ResourceData) go_thunder.SlbVirtualServerPort {
	var vc go_thunder.SlbVirtualServerPort
	var c go_thunder.SlbVirtualServerPortInstance
	c.HaConnMirror = d.Get("ha_conn_mirror").(int)
	c.Protocol = d.Get("protocol").(string)
	c.TemplateDohShared = d.Get("template_doh_shared").(string)
	c.Precedence = d.Get("precedence").(int)
	c.TemplateDblbShared = d.Get("template_dblb_shared").(string)
	c.PortTranslation = d.Get("port_translation").(int)
	c.IPMapList = d.Get("ip_map_list").(string)
	c.TemplateReqmodIcap = d.Get("template_reqmod_icap").(string)

	AclNameListCount := d.Get("acl_name_list.#").(int)
	c.AclNameSrcNatPoolShared = make([]go_thunder.SlbVirtualServerAclNameList, 0, AclNameListCount)

	for i := 0; i < AclNameListCount; i++ {
		var obj1 go_thunder.SlbVirtualServerAclNameList
		prefix := fmt.Sprintf("acl_name_list.%d.", i)
		obj1.AclNameSrcNatPoolShared = d.Get(prefix + "acl_name_src_nat_pool_shared").(string)
		obj1.VAclNameSrcNatPoolShared = d.Get(prefix + "v_acl_name_src_nat_pool_shared").(string)
		obj1.SharedPartitionPoolName = d.Get(prefix + "shared_partition_pool_name").(int)
		obj1.AclNameSeqNumShared = d.Get(prefix + "acl_name_seq_num_shared").(int)
		obj1.AclName = d.Get(prefix + "acl_name").(string)
		obj1.VSharedPartitionPoolName = d.Get(prefix + "v_shared_partition_pool_name").(int)
		obj1.AclNameSrcNatPool = d.Get(prefix + "acl_name_src_nat_pool").(string)
		obj1.VAclNameSeqNum = d.Get(prefix + "v_acl_name_seq_num").(int)
		obj1.AclNameShared = d.Get(prefix + "acl_name_shared").(string)
		obj1.AclNameSeqNum = d.Get(prefix + "acl_name_seq_num").(int)
		obj1.VAclNameSrcNatPool = d.Get(prefix + "v_acl_name_src_nat_pool").(string)
		obj1.VAclNameSeqNumShared = d.Get(prefix + "v_acl_name_seq_num_shared").(int)
		c.AclNameSrcNatPoolShared = append(c.AclNameSrcNatPoolShared, obj1)
	}

	c.StatsDataAction = d.Get("stats_data_action").(string)
	c.TemplateImapPop3Shared = d.Get("template_imap_pop3_shared").(string)
	c.UseDefaultIfNoServer = d.Get("use_default_if_no_server").(int)
	c.TemplateConnectionReuse = d.Get("template_connection_reuse").(string)
	c.CPUCompute = d.Get("cpu_compute").(int)
	c.TemplateTCPShared = d.Get("template_tcp_shared").(string)
	c.TemplateTCP = d.Get("template_tcp").(string)
	c.TemplatePersistDestinationIP = d.Get("template_persist_destination_ip").(string)
	c.SubstituteSourceMac = d.Get("substitute_source_mac").(int)
	c.SharedPartitionDynamicServiceTemplate = d.Get("shared_partition_dynamic_service_template").(int)
	c.SharedPartitionConnectionReuseTemplate = d.Get("shared_partition_connection_reuse_template").(int)
	c.WhenDown = d.Get("when_down").(int)
	c.TemplateClientSslShared = d.Get("template_client_ssl_shared").(string)
	c.SharedPartitionPersistDestinationIPTemplate = d.Get("shared_partition_persist_destination_ip_template").(int)
	c.SharedPartitionExternalServiceTemplate = d.Get("shared_partition_external_service_template").(int)
	c.PersistType = d.Get("persist_type").(string)
	c.SharedPartitionHTTPPolicyTemplate = d.Get("shared_partition_http_policy_template").(int)
	c.UseRcvHopForResp = d.Get("use_rcv_hop_for_resp").(int)
	c.ScaleoutBucketCount = d.Get("scaleout_bucket_count").(int)
	c.IgnoreGlobal = d.Get("ignore_global").(int)
	c.OptimizationLevel = d.Get("optimization_level").(string)
	c.ReqFail = d.Get("req_fail").(int)
	c.NoDestNat = d.Get("no_dest_nat").(int)
	c.Name = d.Get("name").(string)
	c.TemplateSmpp = d.Get("template_smpp").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.TemplateDiameter = d.Get("template_diameter").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.SlbVirtualServerSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj2 go_thunder.SlbVirtualServerSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj2)
	}

	c.TemplateSsli = d.Get("template_ssli").(string)
	c.MemoryCompute = d.Get("memory_compute").(int)
	c.SharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	c.TemplatePolicy = d.Get("template_policy").(string)
	c.View = d.Get("view").(int)
	c.ResetOnServerSelectionFail = d.Get("reset_on_server_selection_fail").(int)
	c.WafTemplate = d.Get("waf_template").(string)
	c.Ipinip = d.Get("ipinip").(int)
	c.NoAutoUpOnAflex = d.Get("no_auto_up_on_aflex").(int)
	c.Rate = d.Get("rate").(int)
	c.GslbEnable = d.Get("gslb_enable").(int)
	c.TemplateDNSShared = d.Get("template_dns_shared").(string)
	c.TemplatePersistSslSid = d.Get("template_persist_ssl_sid").(string)
	c.TemplateDNS = d.Get("template_dns").(string)
	c.SharedPartitionDNSTemplate = d.Get("shared_partition_dns_template").(int)
	c.TemplateSmppShared = d.Get("template_smpp_shared").(string)
	c.TemplateSip = d.Get("template_sip").(string)
	c.TemplateDblb = d.Get("template_dblb").(string)
	c.SharedPartitionServerSslTemplate = d.Get("shared_partition_server_ssl_template").(int)
	c.TemplateClientSsl = d.Get("template_client_ssl").(string)
	c.SupportHTTP2 = d.Get("support_http2").(int)
	c.PTemplateSipShared = d.Get("p_template_sip_shared").(int)
	c.TemplateClientSSH = d.Get("template_client_ssh").(string)
	c.SharedPartitionTCPProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	c.EnablePlayeridCheck = d.Get("enable_playerid_check").(int)
	c.ServiceGroup = d.Get("service_group").(string)
	c.SharedPartitionPersistSslSidTemplate = d.Get("shared_partition_persist_ssl_sid_template").(int)
	c.DefSelectionIfPrefFailed = d.Get("def_selection_if_pref_failed").(string)
	c.SharedPartitionUDP = d.Get("shared_partition_udp").(int)
	c.SynCookie = d.Get("syn_cookie").(int)
	c.TemplateDoh = d.Get("template_doh").(string)
	c.AlternatePort = d.Get("alternate_port").(int)
	c.AlternatePortNumber = d.Get("alternate_port_number").(int)
	c.TemplatePersistSourceIPShared = d.Get("template_persist_source_ip_shared").(string)
	c.TemplateCache = d.Get("template_cache").(string)
	c.TemplatePersistCookieShared = d.Get("template_persist_cookie_shared").(string)
	c.RtpSipCallIDMatch = d.Get("rtp_sip_call_id_match").(int)
	c.SharedPartitionPersistCookieTemplate = d.Get("shared_partition_persist_cookie_template").(int)
	c.TemplateFileInspection = d.Get("template_file_inspection").(string)
	c.TemplateFtp = d.Get("template_ftp").(string)
	c.ServSelFail = d.Get("serv_sel_fail").(int)
	c.TemplateUDP = d.Get("template_udp").(string)
	c.TemplateVirtualPortShared = d.Get("template_virtual_port_shared").(string)
	c.TemplateMqtt = d.Get("template_mqtt").(string)
	c.Action = d.Get("action").(string)
	c.SharedPartitionClientSslTemplate = d.Get("shared_partition_client_ssl_template").(int)
	c.NoLogging = d.Get("no_logging").(int)
	c.SharedPartitionFixTemplate = d.Get("shared_partition_fix_template").(int)
	c.TemplatePersistSourceIP = d.Get("template_persist_source_ip").(string)
	c.TemplateDynamicService = d.Get("template_dynamic_service").(string)
	c.GtpSessionLb = d.Get("gtp_session_lb").(int)
	c.SharedPartitionVirtualPortTemplate = d.Get("shared_partition_virtual_port_template").(int)
	c.UseCgnv6 = d.Get("use_cgnv6").(int)
	c.TemplatePersistCookie = d.Get("template_persist_cookie").(string)
	c.SharedPartitionSMTPTemplate = d.Get("shared_partition_smtp_template").(int)
	c.TemplateVirtualPort = d.Get("template_virtual_port").(string)
	c.ConnLimit = d.Get("conn_limit").(int)
	c.TrunkFwd = d.Get("trunk_fwd").(int)
	c.TemplateUDPShared = d.Get("template_udp_shared").(string)
	c.TemplateHTTPPolicyShared = d.Get("template_http_policy_shared").(string)
	c.Pool = d.Get("pool").(string)
	c.SnatOnVip = d.Get("snat_on_vip").(int)
	c.TemplateConnectionReuseShared = d.Get("template_connection_reuse_shared").(string)
	c.SharedPartitionTCP = d.Get("shared_partition_tcp").(int)

	AclIdListCount := d.Get("acl_id_list.#").(int)
	c.VAclIDSeqNum = make([]go_thunder.SlbVirtualServerAclIDList, 0, AclIdListCount)

	for i := 0; i < AclIdListCount; i++ {
		var obj3 go_thunder.SlbVirtualServerAclIDList
		prefix := fmt.Sprintf("acl_id_list.%d.", i)
		obj3.VAclIDSeqNum = d.Get(prefix + "v_acl_id_seq_num").(int)
		obj3.AclIDSeqNum = d.Get(prefix + "acl_id_seq_num").(int)
		obj3.AclIDSrcNatPool = d.Get(prefix + "acl_id_src_nat_pool").(string)
		obj3.AclIDSeqNumShared = d.Get(prefix + "acl_id_seq_num_shared").(int)
		obj3.VAclIDSrcNatPool = d.Get(prefix + "v_acl_id_src_nat_pool").(string)
		obj3.AclIDShared = d.Get(prefix + "acl_id_shared").(int)
		obj3.VAclIDSrcNatPoolShared = d.Get(prefix + "v_acl_id_src_nat_pool_shared").(string)
		obj3.AclID = d.Get(prefix + "acl_id").(int)
		obj3.AclIDSrcNatPoolShared = d.Get(prefix + "acl_id_src_nat_pool_shared").(string)
		obj3.VSharedPartitionPoolID = d.Get(prefix + "v_shared_partition_pool_id").(int)
		obj3.SharedPartitionPoolID = d.Get(prefix + "shared_partition_pool_id").(int)
		obj3.VAclIDSeqNumShared = d.Get(prefix + "v_acl_id_seq_num_shared").(int)
		c.VAclIDSeqNum = append(c.VAclIDSeqNum, obj3)
	}

	c.SharedPartitionDblbTemplate = d.Get("shared_partition_dblb_template").(int)
	c.SharedPartitionHTTPTemplate = d.Get("shared_partition_http_template").(int)
	c.TemplateExternalService = d.Get("template_external_service").(string)
	c.OnSyn = d.Get("on_syn").(int)
	c.TemplateSMTPShared = d.Get("template_smtp_shared").(string)
	c.TemplateSipShared = d.Get("template_sip_shared").(string)
	c.TemplatePersistSslSidShared = d.Get("template_persist_ssl_sid_shared").(string)
	c.ForceRoutingMode = d.Get("force_routing_mode").(int)
	c.TemplateHTTPPolicy = d.Get("template_http_policy").(string)
	c.TemplatePolicyShared = d.Get("template_policy_shared").(string)
	c.TemplateScaleout = d.Get("template_scaleout").(string)
	c.WhenDownProtocol2 = d.Get("when_down_protocol2").(int)
	c.TemplateFix = d.Get("template_fix").(string)
	c.TemplateSMTP = d.Get("template_smtp").(string)
	c.RedirectToHTTPS = d.Get("redirect_to_https").(int)
	c.AltProtocol2 = d.Get("alt_protocol2").(string)
	c.AltProtocol1 = d.Get("alt_protocol1").(string)
	c.MessageSwitching = d.Get("message_switching").(int)
	c.TemplateImapPop3 = d.Get("template_imap_pop3").(string)
	c.ScaleoutDeviceGroup = d.Get("scaleout_device_group").(int)
	c.SharedPartitionPersistSourceIPTemplate = d.Get("shared_partition_persist_source_ip_template").(int)
	c.L7HardwareAssist = d.Get("l7_hardware_assist").(int)
	c.TemplateFixShared = d.Get("template_fix_shared").(string)
	c.TemplateTCPProxyShared = d.Get("template_tcp_proxy_shared").(string)
	c.SharedPartitionCacheTemplate = d.Get("shared_partition_cache_template").(int)
	c.UseAlternatePort = d.Get("use_alternate_port").(int)
	c.TemplateTCPProxyServer = d.Get("template_tcp_proxy_server").(string)
	c.TrunkRev = d.Get("trunk_rev").(int)
	c.EthFwd = d.Get("eth_fwd").(int)
	c.PoolShared = d.Get("pool_shared").(string)
	c.TemplateRespmodIcap = d.Get("template_respmod_icap").(string)
	c.Range = d.Get("range").(int)
	c.Reset = d.Get("reset").(int)
	c.TemplateExternalServiceShared = d.Get("template_external_service_shared").(string)
	c.Auto = d.Get("auto").(int)
	c.SharedPartitionSmppTemplate = d.Get("shared_partition_smpp_template").(int)
	c.TemplateDynamicServiceShared = d.Get("template_dynamic_service_shared").(string)
	c.TemplateServerSSH = d.Get("template_server_ssh").(string)

	AflexScriptsCount := d.Get("aflex_scripts.#").(int)
	c.Aflex = make([]go_thunder.SlbVirtualServerAflexScripts, 0, AflexScriptsCount)

	for i := 0; i < AflexScriptsCount; i++ {
		var obj4 go_thunder.SlbVirtualServerAflexScripts
		prefix := fmt.Sprintf("aflex_scripts.%d.", i)
		obj4.Aflex = d.Get(prefix + "aflex").(string)
		obj4.AflexShared = d.Get(prefix + "aflex_shared").(string)
		c.Aflex = append(c.Aflex, obj4)
	}

	c.TemplateHTTPShared = d.Get("template_http_shared").(string)
	c.TemplateServerSsl = d.Get("template_server_ssl").(string)
	c.SharedPartitionDiameterTemplate = d.Get("shared_partition_diameter_template").(int)
	c.TemplateServerSslShared = d.Get("template_server_ssl_shared").(string)
	c.TemplatePersistDestinationIPShared = d.Get("template_persist_destination_ip_shared").(string)
	c.TemplateCacheShared = d.Get("template_cache_shared").(string)
	c.PortNumber = d.Get("port_number").(int)
	c.TemplateTCPProxyClient = d.Get("template_tcp_proxy_client").(string)
	c.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	c.TemplateTCPProxy = d.Get("template_tcp_proxy").(string)
	c.ExtendedStats = d.Get("extended_stats").(int)
	c.TemplateHTTP = d.Get("template_http").(string)
	c.Expand = d.Get("expand").(int)
	c.SharedPartitionDohTemplate = d.Get("shared_partition_doh_template").(int)
	c.SkipRevHash = d.Get("skip_rev_hash").(int)
	c.TemplateDiameterShared = d.Get("template_diameter_shared").(string)
	c.ClientipStickyNat = d.Get("clientip_sticky_nat").(int)
	c.Secs = d.Get("secs").(int)
	c.SharedPartitionImapPop3Template = d.Get("shared_partition_imap_pop3_template").(int)

	var obj5 go_thunder.SlbVirtualServerAuthCfg
	prefix := "auth_cfg.0."
	obj5.AaaPolicy = d.Get(prefix + "aaa_policy").(string)

	c.AaaPolicy = obj5

	c.EthRev = d.Get("eth_rev").(int)

	vc.HaConnMirror = c
	return vc
}
