package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_global`: Stateful Firewall Configuration (default:disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallGlobalCreate,
		UpdateContext: resourceCgnv6StatefulFirewallGlobalUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallGlobalRead,
		DeleteContext: resourceCgnv6StatefulFirewallGlobalDelete,

		Schema: map[string]*schema.Schema{
			"respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default: off)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'tcp_packet_process': TCP Packet Process; 'udp_packet_process': UDP Packet Process; 'other_packet_process': Other Packet Process; 'packet_inbound_deny': Inbound Packet Denied; 'packet_process_failure': Packet Error Drop; 'outbound_session_created': Outbound Session Created; 'outbound_session_freed': Outbound Session Freed; 'inbound_session_created': Inbound Session Created; 'inbound_session_freed': Inbound Session Freed; 'tcp_session_created': TCP Session Created; 'tcp_session_freed': TCP Session Freed; 'udp_session_created': UDP Session Created; 'udp_session_freed': UDP Session Freed; 'other_session_created': Other Session Created; 'other_session_freed': Other Session Freed; 'session_creation_failure': Session Creation Failure; 'no_fwd_route': No Forward Route; 'no_rev_route': No Reverse Route; 'packet_standby_drop': Standby Drop; 'tcp_fullcone_created': TCP Full-cone Created; 'tcp_fullcone_freed': TCP Full-cone Freed; 'udp_fullcone_created': UDP Full-cone Created; 'udp_fullcone_freed': UDP Full-cone Freed; 'fullcone_creation_failure': Full-Cone Creation Failure; 'eif_process': Endpnt-Independent Filter Matched; 'one_arm_drop': One-Arm Drop; 'no_class_list_match': No Class-List Match Drop; 'outbound_session_created_shadow': Outbound Session Created Shadow; 'outbound_session_freed_shadow': Outbound Session Freed Shadow; 'inbound_session_created_shadow': Inbound Session Created Shadow; 'inbound_session_freed_shadow': Inbound Session Freed Shadow; 'tcp_session_created_shadow': TCP Session Created Shadow; 'tcp_session_freed_shadow': TCP Session Freed Shadow; 'udp_session_created_shadow': UDP Session Created Shadow; 'udp_session_freed_shadow': UDP Session Freed Shadow; 'other_session_created_shadow': Other Session Created Shadow; 'other_session_freed_shadow': Other Session Freed Shadow; 'session_creation_failure_shadow': Session Creation Failure Shadow; 'bad_session_freed': Bad Session Proto on Free; 'ctl_mem_alloc': Memory Alloc; 'ctl_mem_free': Memory Free; 'tcp_fullcone_created_shadow': TCP Full-cone Created Shadow; 'tcp_fullcone_freed_shadow': TCP Full-cone Freed Shadow; 'udp_fullcone_created_shadow': UDP Full-cone Created Shadow; 'udp_fullcone_freed_shadow': UDP Full-cone Freed Shadow; 'fullcone_in_del_q': Full-cone Found in Delete Queue; 'fullcone_overflow_eim': EIM Overflow; 'fullcone_overflow_eif': EIF Overflow; 'fullcone_free_found': Full-cone Free Found From Conn; 'fullcone_free_retry_lookup': Full-cone Retry Look-up; 'fullcone_free_not_found': Full-cone Free Not Found; 'eif_limit_exceeded': EIF Limit Exceeded; 'eif_disable_drop': EIF Disable Drop; 'eif_process_failure': EIF Process Failure; 'eif_filtered': EIF Filtered; 'ha_standby_session_created': HA Standby Session Created; 'ha_standby_session_eim': HA Standby Session EIM; 'ha_standby_session_eif': HA Standby Session EIF;",
						},
					},
				},
			},
			"stateful_firewall_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable stateful firewall;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6StatefulFirewallGlobalSamplingEnable(d []interface{}) []edpt.Cgnv6StatefulFirewallGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6StatefulFirewallGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6StatefulFirewallGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallGlobal(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallGlobal {
	var ret edpt.Cgnv6StatefulFirewallGlobal
	ret.Inst.RespondToUserMac = d.Get("respond_to_user_mac").(int)
	ret.Inst.SamplingEnable = getSliceCgnv6StatefulFirewallGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.StatefulFirewallValue = d.Get("stateful_firewall_value").(string)
	//omit uuid
	return ret
}
