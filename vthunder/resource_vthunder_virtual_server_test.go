package vthunder

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"reflect"
	"testing"
)

var NAME_VS = "vs9"
var SG = "sg9"

var TEST_VS_RESOURCE = `
resource "vthunder_virtual_server" "vs9" {
ha_dynamic=1
name="` + NAME_VS + `"
ip_address="10.0.9.7"
port_list {
auto=1
protocol="tcp"
port_number=8080
snat_on_vip=1
service_group="` + SG + `"
} 
}
`

//Acceptance test
func TestAccVthunderVirtualServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckVirtualServerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_VS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckVirtualServerExists(NAME_VS, true),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "name", NAME_VS),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "ha_dynamic", "1"),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "ip_address", "10.0.9.7"),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "port_list.0.auto", "1"),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "port_list.0.service_group", SG),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "port_list.0.snat_on_vip", "1"),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "port_list.0.port_number", "8080"),
					resource.TestCheckResourceAttr("vthunder_virtual_server.vs9", "port_list.0.protocol", "tcp"),
				),
			},
		},
	})
}

func TestAccVthunderVirtualServer_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckVirtualServerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_VS_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckVirtualServerExists(NAME_VS, true),
				),
				ResourceName:      NAME_VS,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}

//Unit test for utility method to create Virtual Server structure
func TestDataToVs(t *testing.T) {

	StatsDataAction := "enable"
	Ipv6ACLShared := "v6"
	ACLName := "name"
	EnableDisableAction := "enable"
	HaDynamic := 1
	RedistributeRouteMap := "map"
	ACLNameShared := "name"
	IPAddress := "1.1.1.3"
	UseIfIP := 1
	UUID := "uuid"
	Vrid := 1
	DisableVipAdv := 1
	TemplateVirtualServer := "template"
	ArpDisable := 1
	Description := "description"
	RedistributionFlagged := 1
	Netmask := "255.255.0.0"
	ACLID := 1
	Ipv6ACL := "ipv6"
	TemplateLogging := "enable"
	ExtendedStats := 1
	Name := "name"
	TemplateScaleout := "scaleout"
	TemplatePolicy := "policy"
	UserTag := "tag"
	TemplatePolicyShared := "policyshared"
	Ethernet := 1
	SharedPartitionPolicyTemplate := 1
	ACLIDShared := 1

	resourceSchema := map[string]*schema.Schema{
		"stats_data_action": &schema.Schema{
			Type: schema.TypeString,
		},
		"ipv6_acl_shared": &schema.Schema{
			Type: schema.TypeString,
		},
		"acl_name": &schema.Schema{
			Type: schema.TypeString,
		},
		"enable_disable_action": &schema.Schema{
			Type: schema.TypeString,
		},
		"ha_dynamic": &schema.Schema{
			Type: schema.TypeInt,
		},
		"redistribute_route_map": &schema.Schema{
			Type: schema.TypeString,
		},
		"acl_name_shared": &schema.Schema{
			Type: schema.TypeString,
		},
		"ip_address": &schema.Schema{
			Type: schema.TypeString,
		},
		"use_if_ip": &schema.Schema{
			Type: schema.TypeInt,
		},
		"uuid": &schema.Schema{
			Type: schema.TypeString,
		},
		"vrid": &schema.Schema{
			Type: schema.TypeInt,
		},
		"disable_vip_adv": &schema.Schema{
			Type: schema.TypeInt,
		},
		"template_virtual_server": &schema.Schema{
			Type: schema.TypeString,
		},
		"arp_disable": &schema.Schema{
			Type: schema.TypeInt,
		},
		"description": &schema.Schema{
			Type: schema.TypeString,
		},
		"redistribution_flagged": &schema.Schema{
			Type: schema.TypeInt,
		},
		"netmask": &schema.Schema{
			Type: schema.TypeString,
		},
		"acl_id": &schema.Schema{
			Type: schema.TypeInt,
		},
		"ipv6_acl": &schema.Schema{
			Type: schema.TypeString,
		},
		"template_logging": &schema.Schema{
			Type: schema.TypeString,
		},
		"extended_stats": &schema.Schema{
			Type: schema.TypeInt,
		},
		"name": &schema.Schema{
			Type: schema.TypeString,
		},
		"template_scaleout": &schema.Schema{
			Type: schema.TypeString,
		},
		"template_policy": &schema.Schema{
			Type: schema.TypeString,
		},
		"user_tag": &schema.Schema{
			Type: schema.TypeString,
		},
		"template_policy_shared": &schema.Schema{
			Type: schema.TypeString,
		},
		"ethernet": &schema.Schema{
			Type: schema.TypeInt,
		},
		"shared_partition_policy_template": &schema.Schema{
			Type: schema.TypeInt,
		},
		"acl_id_shared": &schema.Schema{
			Type: schema.TypeInt,
		},
		"port_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"extended_stats": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_file_inspection": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"use_default_if_no_server": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_tcp_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"use_alternate_port": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_http_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"when_down_protocol2": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"serv_sel_fail": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"support_http2": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"pool_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"action": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_persist_destination_ip_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_udp": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"alt_protocol1": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"alt_protocol2": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_tcp_proxy_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_policy_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"ipv6_acl": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"alternate_port": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_server_ssl_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"when_down": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"pool": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"scaleout_device_group": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"trunk_fwd": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"message_switching": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_udp": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"port_number": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"service_group": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_pool": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_persist_ssl_sid_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"no_auto_up_on_aflex": {
						Type:        schema.TypeInt,
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
					"reset": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_external_service_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"trunk_rev": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"ipinip": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_http_policy_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_respmod_icap": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_diameter_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"scaleout_bucket_count": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_persist_source_ip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_virtual_port_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"snat_on_vip": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"redirect_to_https": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_cache_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"range": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_smtp": {
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
					"ha_conn_mirror": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_ftp": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"on_syn": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"alternate_port_number": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_connection_reuse_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_external_service_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"persist_type": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"force_routing_mode": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_dynamic_service_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_reqmod_icap": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"no_logging": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_dns_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"clientip_sticky_nat": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"gslb_enable": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_client_ssl_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_persist_cookie_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"use_cgnv6": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"conn_limit": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"req_fail": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"secs": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"use_rcv_hop_for_resp": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_persist_cookie_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_tcp_proxy_client": {
						Type:        schema.TypeString,
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
					"template_dynamic_service": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_fix": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"auto": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_http": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_virtual_port": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_persist_destination_ip_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"ip_map_list": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"waf_template": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"acl_id_list": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"acl_id_src_nat_pool": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
								"acl_id_src_nat_pool_shared": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
								"acl_id": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "",
								},
								"acl_id_seq_num_shared": {
									Type:        schema.TypeInt,
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
							},
						},
					},
					"uuid": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"precedence": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"protocol": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"view": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"user_tag": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"def_selection_if_pref_failed": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_persist_destination_ip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_dns": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_policy": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_dblb": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"acl_name_list": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
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
								"acl_name_src_nat_pool_shared": {
									Type:        schema.TypeString,
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
								"acl_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "",
								},
							},
						},
					},
					"shared_partition_tcp_proxy_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_imap_pop3": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"no_dest_nat": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_persist_cookie": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"reset_on_server_selection_fail": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"rtp_sip_call_id_match": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_udp_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"expand": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"skip_rev_hash": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"enable_playerid_check": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_dynamic_service_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_dns_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_virtual_port_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_external_service": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_scaleout": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_http_policy_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_sip": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_diameter": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_tcp": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_client_ssl_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"stats_data_action": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_smpp": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_connection_reuse_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_ssli": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"rate": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_connection_reuse": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"syn_cookie": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"optimization_level": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_tcp_proxy_server": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_persist_source_ip_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_policy_template": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_tcp_proxy": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"shared_partition_http_template": {
						Type:        schema.TypeInt,
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
					"eth_fwd": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"template_http_policy": {
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
					"template_cache_shared": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_client_ssl": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_server_ssh": {
						Type:        schema.TypeString,
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
					"template_client_ssh": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template_server_ssl": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"l7_hardware_assist": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"eth_rev": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"shared_partition_tcp": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
					"port_translation": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
				},
			},
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
					"target_floating_template_policy": {
						Type:        schema.TypeString,
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
					"cancel_migration": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "",
					},
				},
			},
		},
	}
	resourceDataMap := map[string]interface{}{
		"stats_data_action":                StatsDataAction,
		"ipv6_acl_shared":                  Ipv6ACLShared,
		"acl_name":                         ACLName,
		"enable_disable_action":            EnableDisableAction,
		"ha_dynamic":                       HaDynamic,
		"redistribute_route_map":           RedistributeRouteMap,
		"acl_name_shared":                  ACLNameShared,
		"ip_address":                       IPAddress,
		"use_if_ip":                        UseIfIP,
		"uuid":                             UUID,
		"vrid":                             Vrid,
		"disable_vip_adv":                  DisableVipAdv,
		"template_virtual_server":          TemplateVirtualServer,
		"arp_disable":                      ArpDisable,
		"description":                      Description,
		"redistribution_flagged":           RedistributionFlagged,
		"netmask":                          Netmask,
		"acl_id":                           ACLID,
		"ipv6_acl":                         Ipv6ACL,
		"template_logging":                 TemplateLogging,
		"extended_stats":                   ExtendedStats,
		"name":                             Name,
		"template_scaleout":                TemplateScaleout,
		"template_policy":                  TemplatePolicy,
		"user_tag":                         UserTag,
		"template_policy_shared":           TemplatePolicyShared,
		"ethernet":                         Ethernet,
		"shared_partition_policy_template": SharedPartitionPolicyTemplate,
		"acl_id_shared":                    ACLIDShared,
		"migrate_vip":                      map[string]interface{}{},
		"port_list":                        map[string]interface{}{},
	}
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	var v go_vthunder.VirtalServerInstanceMain

	var vsMain go_vthunder.VirtualServerMain

	vsMain.StatsDataAction = StatsDataAction
	vsMain.Ipv6ACLShared = Ipv6ACLShared
	vsMain.ACLName = ACLName
	vsMain.EnableDisableAction = EnableDisableAction
	vsMain.HaDynamic = HaDynamic
	vsMain.RedistributeRouteMap = RedistributeRouteMap
	vsMain.ACLNameShared = ACLNameShared
	vsMain.IPAddress = IPAddress
	vsMain.UseIfIP = UseIfIP
	vsMain.UUID = UUID
	vsMain.Vrid = Vrid
	vsMain.DisableVipAdv = DisableVipAdv
	vsMain.TemplateVirtualServer = TemplateVirtualServer
	vsMain.ArpDisable = ArpDisable
	vsMain.Description = Description
	vsMain.RedistributionFlagged = RedistributionFlagged
	vsMain.Netmask = Netmask
	vsMain.ACLID = ACLID
	vsMain.Ipv6ACL = Ipv6ACL
	vsMain.TemplateLogging = TemplateLogging
	vsMain.ExtendedStats = ExtendedStats
	vsMain.Name = Name
	vsMain.TemplateScaleout = TemplateScaleout
	vsMain.TemplatePolicy = TemplatePolicy
	vsMain.UserTag = UserTag
	vsMain.TemplatePolicyShared = TemplatePolicyShared
	vsMain.Ethernet = Ethernet
	vsMain.SharedPartitionPolicyTemplate = SharedPartitionPolicyTemplate
	vsMain.ACLIDShared = ACLIDShared
	vsMain.Protocol = make([]go_vthunder.PortList, 0, 1)

	v.Name = vsMain

	cases := []struct {
		input  *schema.ResourceData
		output go_vthunder.VirtalServerInstanceMain
	}{
		{
			input:  resourceLocalData,
			output: v,
		},
	}

	for _, tc := range cases {
		output := dataToVs("name", resourceLocalData)
		if !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("Got:\n\n%#v\n\nExpected:\n\n%#v", output, tc.output)
		}
	}

}

func testCheckVirtualServerExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		client := testAccProvider.Meta().(vThunder)
		vs, err := go_vthunder.GetVS(client.Token, name, client.Host)

		if err != nil {
			return err
		}
		if exists && vs == nil {
			return fmt.Errorf("virtual server %s was not created.", name)
		}
		if !exists && vs != nil {
			return fmt.Errorf("virtual server %s still exists.", name)
		}
		return nil
	}
}

func testCheckVirtualServerDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(vThunder)
	for _, rs := range s.RootModule().Resources {

		name := rs.Primary.ID
		device, err := go_vthunder.GetVS(client.Token, name, client.Host)
		if err != nil {
			return err
		}
		if device == nil {
			return fmt.Errorf("virtual server %s not destroyed.", name)
		}
	}
	return nil
}
