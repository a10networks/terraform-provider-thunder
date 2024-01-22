package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DdosProtection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ddos_protection`: Configure CGNV6 DDoS Protection\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DdosProtectionCreate,
		UpdateContext: resourceCgnv6DdosProtectionUpdate,
		ReadContext:   resourceCgnv6DdosProtectionRead,
		DeleteContext: resourceCgnv6DdosProtectionDelete,

		Schema: map[string]*schema.Schema{
			"disable_nat_ip_by_bgp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"enable_action": {
				Type: schema.TypeString, Optional: true, Default: "local", Description: "'local': Enable local logs only; 'remote': Enable logging to remote server & IPFIX; 'both': Enable both local & remote logs;",
			},
			"ip_entries": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"l4_entries": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"logging_action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable CGN DDoS protection logging; 'disable': Disable both local & remote CGN DDoS protection logging;",
			},
			"max_hw_entries": {
				Type: schema.TypeInt, Optional: true, Default: 262144, Description: "Configure maximum HW entries",
			},
			"packets_per_second": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeInt, Optional: true, Default: 3000000, Description: "Configure packets-per-second threshold per IP(default 3000000)",
						},
						"action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_type": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'log': Log the event only; 'drop': Log, and drop all packets (default); 'redistribute-route': Log, Drop, and Notify upstream router to reroute the packets;",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map name",
									},
									"expiration": {
										Type: schema.TypeInt, Optional: true, Default: 3600, Description: "To specify time to revert the action after pps is decreased to below threshold (Expiration time, in minutes (default is 3600 seconds))",
									},
									"expiration_route": {
										Type: schema.TypeInt, Optional: true, Default: 3600, Description: "To specify time to revert the action after pps is decreased to below threshold (Expiration time, in seconds (default is 3600 seconds))",
									},
									"timer_multiply_max": {
										Type: schema.TypeInt, Optional: true, Default: 6, Description: "To specify max value of timer multiplier for attacks lasted long time (Max value of timer multiplier (default is 6))",
									},
									"remove_wait_timer": {
										Type: schema.TypeInt, Optional: true, Default: 300, Description: "Time after which IP will be removed from blackhole",
									},
								},
							},
						},
						"tcp": {
							Type: schema.TypeInt, Optional: true, Default: 3000, Description: "Configure packets-per-second threshold per TCP port (default: 3000)",
						},
						"tcp_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_action_type": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'log': Log the event only; 'drop': Log, and drop all packets (default);",
									},
									"tcp_expiration": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "To specify time to revert the action after pps is decreased to below threshold (Expiration time, in seconds (default is 30 seconds))",
									},
								},
							},
						},
						"udp": {
							Type: schema.TypeInt, Optional: true, Default: 3000, Description: "Configure packets-per-second threshold per UDP port (default: 3000)",
						},
						"udp_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_action_type": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'log': Log the event only; 'drop': Log, and drop all packets (default);",
									},
									"udp_expiration": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "To specify time to revert the action after pps is decreased to below threshold (Expiration time, in seconds (default is 30 seconds))",
									},
								},
							},
						},
						"other": {
							Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Configure packets-per-second threshold for other L4 protocols(default 10000)",
						},
						"other_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"other_action_type": {
										Type: schema.TypeString, Optional: true, Default: "drop", Description: "'log': Log the event only; 'drop': Log, and drop all packets (default);",
									},
									"other_expiration": {
										Type: schema.TypeInt, Optional: true, Default: 30, Description: "To specify time to revert the action after pps is decreased to below threshold (Expiration time, in seconds (default is 30 seconds))",
									},
								},
							},
						},
						"include_existing_session": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Count traffic associated with existing session into the packets-per-second (Default: Disabled)",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'l3_entry_added': L3 Entry Added; 'l3_entry_deleted': L3 Entry Deleted; 'l3_entry_added_to_bgp': L3 Entry added to BGP; 'l3_entry_removed_from_bgp': Entry removed from BGP; 'l3_entry_added_to_hw': L3 Entry added to HW; 'l3_entry_removed_from_hw': L3 Entry removed from HW; 'l3_entry_too_many': L3 Too many entries; 'l3_entry_match_drop': L3 Entry match drop; 'l3_entry_match_drop_hw': L3 HW entry match drop; 'l3_entry_drop_max_hw_exceeded': L3 Entry Drop due to HW Limit Exceeded; 'l4_entry_added': L4 Entry added; 'l4_entry_deleted': L4 Entry deleted; 'l4_entry_added_to_hw': L4 Entry added to HW; 'l4_entry_removed_from_hw': L4 Entry removed from HW; 'l4_hw_out_of_entries': HW out of L4 entries; 'l4_entry_match_drop': L4 Entry match drop; 'l4_entry_match_drop_hw': L4 HW Entry match drop; 'l4_entry_drop_max_hw_exceeded': L4 Entry Drop due to HW Limit Exceeded; 'l4_entry_list_alloc': L4 Entry list alloc; 'l4_entry_list_free': L4 Entry list free; 'l4_entry_list_alloc_failure': L4 Entry list alloc failures; 'ip_node_alloc': Node alloc; 'ip_node_free': Node free; 'ip_node_alloc_failure': Node alloc failures; 'ip_port_block_alloc': Port block alloc; 'ip_port_block_free': Port block free; 'ip_port_block_alloc_failure': Port block alloc failure; 'ip_other_block_alloc': Other block alloc; 'ip_other_block_free': Other block free; 'ip_other_block_alloc_failure': Other block alloc failure; 'entry_added_shadow': Entry added shadow; 'entry_invalidated': Entry invalidated; 'l3_entry_add_to_bgp_failure': L3 Entry BGP add failures; 'l3_entry_remove_from_bgp_failure': L3 entry BGP remove failures; 'l3_entry_add_to_hw_failure': L3 entry HW add failure; 'syn_cookie_syn_ack_sent': SYN cookie SYN ACK sent; 'syn_cookie_verification_passed': SYN cookie verification passed; 'syn_cookie_verification_failed': SYN cookie verification failed; 'syn_cookie_conn_setup_failed': SYN cookie connection setup failed;",
						},
					},
				},
			},
			"syn_cookie": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"syn_cookie_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CGNv6 Syn-Cookie Protection",
						},
						"syn_cookie_on_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "on-threshold for Syn-cookie (Decimal number)",
						},
						"syn_cookie_on_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 120, Description: "on-timeout for Syn-cookie (Timeout in seconds, default is 120 seconds (2 minutes))",
						},
					},
				},
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable CGNV6 NAT pool DDoS protection (default); 'disable': Disable CGNV6 NAT pool DDoS protection;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone": {
				Type: schema.TypeString, Optional: true, Description: "Disable NAT IP based on DDoS zone name set in BGP",
			},
		},
	}
}
func resourceCgnv6DdosProtectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DdosProtectionRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DdosProtectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DdosProtectionRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DdosProtectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DdosProtectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DdosProtectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DdosProtection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6DdosProtectionDisableNatIpByBgp81(d []interface{}) edpt.Cgnv6DdosProtectionDisableNatIpByBgp81 {

	var ret edpt.Cgnv6DdosProtectionDisableNatIpByBgp81
	return ret
}

func getObjectCgnv6DdosProtectionIpEntries82(d []interface{}) edpt.Cgnv6DdosProtectionIpEntries82 {

	var ret edpt.Cgnv6DdosProtectionIpEntries82
	return ret
}

func getObjectCgnv6DdosProtectionL4Entries83(d []interface{}) edpt.Cgnv6DdosProtectionL4Entries83 {

	var ret edpt.Cgnv6DdosProtectionL4Entries83
	return ret
}

func getObjectCgnv6DdosProtectionPacketsPerSecond(d []interface{}) edpt.Cgnv6DdosProtectionPacketsPerSecond {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionPacketsPerSecond
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(int)
		ret.Action = getObjectCgnv6DdosProtectionPacketsPerSecondAction(in["action"].([]interface{}))
		ret.Tcp = in["tcp"].(int)
		ret.TcpAction = getObjectCgnv6DdosProtectionPacketsPerSecondTcpAction(in["tcp_action"].([]interface{}))
		ret.Udp = in["udp"].(int)
		ret.UdpAction = getObjectCgnv6DdosProtectionPacketsPerSecondUdpAction(in["udp_action"].([]interface{}))
		ret.Other = in["other"].(int)
		ret.OtherAction = getObjectCgnv6DdosProtectionPacketsPerSecondOtherAction(in["other_action"].([]interface{}))
		ret.IncludeExistingSession = in["include_existing_session"].(int)
	}
	return ret
}

func getObjectCgnv6DdosProtectionPacketsPerSecondAction(d []interface{}) edpt.Cgnv6DdosProtectionPacketsPerSecondAction {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionPacketsPerSecondAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActionType = in["action_type"].(string)
		ret.RouteMap = in["route_map"].(string)
		ret.Expiration = in["expiration"].(int)
		ret.ExpirationRoute = in["expiration_route"].(int)
		ret.TimerMultiplyMax = in["timer_multiply_max"].(int)
		ret.RemoveWaitTimer = in["remove_wait_timer"].(int)
	}
	return ret
}

func getObjectCgnv6DdosProtectionPacketsPerSecondTcpAction(d []interface{}) edpt.Cgnv6DdosProtectionPacketsPerSecondTcpAction {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionPacketsPerSecondTcpAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpActionType = in["tcp_action_type"].(string)
		ret.TcpExpiration = in["tcp_expiration"].(int)
	}
	return ret
}

func getObjectCgnv6DdosProtectionPacketsPerSecondUdpAction(d []interface{}) edpt.Cgnv6DdosProtectionPacketsPerSecondUdpAction {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionPacketsPerSecondUdpAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpActionType = in["udp_action_type"].(string)
		ret.UdpExpiration = in["udp_expiration"].(int)
	}
	return ret
}

func getObjectCgnv6DdosProtectionPacketsPerSecondOtherAction(d []interface{}) edpt.Cgnv6DdosProtectionPacketsPerSecondOtherAction {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionPacketsPerSecondOtherAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OtherActionType = in["other_action_type"].(string)
		ret.OtherExpiration = in["other_expiration"].(int)
	}
	return ret
}

func getSliceCgnv6DdosProtectionSamplingEnable(d []interface{}) []edpt.Cgnv6DdosProtectionSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6DdosProtectionSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6DdosProtectionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6DdosProtectionSynCookie(d []interface{}) edpt.Cgnv6DdosProtectionSynCookie {

	count1 := len(d)
	var ret edpt.Cgnv6DdosProtectionSynCookie
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SynCookieEnable = in["syn_cookie_enable"].(int)
		ret.SynCookieOnThreshold = in["syn_cookie_on_threshold"].(int)
		ret.SynCookieOnTimeout = in["syn_cookie_on_timeout"].(int)
	}
	return ret
}

func dataToEndpointCgnv6DdosProtection(d *schema.ResourceData) edpt.Cgnv6DdosProtection {
	var ret edpt.Cgnv6DdosProtection
	ret.Inst.DisableNatIpByBgp = getObjectCgnv6DdosProtectionDisableNatIpByBgp81(d.Get("disable_nat_ip_by_bgp").([]interface{}))
	ret.Inst.EnableAction = d.Get("enable_action").(string)
	ret.Inst.IpEntries = getObjectCgnv6DdosProtectionIpEntries82(d.Get("ip_entries").([]interface{}))
	ret.Inst.L4Entries = getObjectCgnv6DdosProtectionL4Entries83(d.Get("l4_entries").([]interface{}))
	ret.Inst.LoggingAction = d.Get("logging_action").(string)
	ret.Inst.MaxHwEntries = d.Get("max_hw_entries").(int)
	ret.Inst.PacketsPerSecond = getObjectCgnv6DdosProtectionPacketsPerSecond(d.Get("packets_per_second").([]interface{}))
	ret.Inst.SamplingEnable = getSliceCgnv6DdosProtectionSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SynCookie = getObjectCgnv6DdosProtectionSynCookie(d.Get("syn_cookie").([]interface{}))
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	ret.Inst.Zone = d.Get("zone").(string)
	return ret
}
