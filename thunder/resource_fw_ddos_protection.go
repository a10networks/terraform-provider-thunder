package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwDdosProtection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_ddos_protection`: Configure FW DDoS Protection\n\n__PLACEHOLDER__",
		CreateContext: resourceFwDdosProtectionCreate,
		UpdateContext: resourceFwDdosProtectionUpdate,
		ReadContext:   resourceFwDdosProtectionRead,
		DeleteContext: resourceFwDdosProtectionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action_type": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Log, and drop all packets (default); 'redistribute-route': Log, Drop, and Notify upstream router to reroute the packets;",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map name",
						},
						"expiration": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "To specify time in minutes to revert the action (Expiration time, in minutes (default is 5 mins))",
						},
						"expiration_route": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "To specify time in minutes to revert the action (Expiration time, in minutes (default is 60 mins))",
						},
						"timer_multiply_max": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "To specify max value of timer multiplier for attacks lasted long time (Max value of timer multiplier (default is 6))",
						},
						"remove_wait_timer": {
							Type: schema.TypeInt, Optional: true, Default: 300, Description: "Max time to wait before removing IP from blackhole (Max value in seconds (default 300))",
						},
					},
				},
			},
			"dynamic_blacklist": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dynamic_blacklist_action": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable protection against volumetric attacks using dynamic blacklist; 'disable': Disable protection against volumetric attacks using dynamic blacklist;",
						},
						"dir": {
							Type: schema.TypeString, Optional: true, Default: "both", Description: "'inbound': enable in inbound direction; 'outbound': enable in outbound direction; 'both': enable in both directions;",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Timeout value (in seconds) for dynamic blacklist (Timeout value (in seconds) for dynamic blacklist(default is 5 seconds))",
						},
						"cpu_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Core-level CPU usage threshold for dynamic blacklist creation (Core-level CPU usage threshold for dynamic blacklist creation (default is 60))",
						},
					},
				},
			},
			"logging": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging_action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable FW DDoS protection logging; 'disable': Disable both local & remote FW DDoS protection logging;",
						},
						"enable_action": {
							Type: schema.TypeString, Optional: true, Default: "local", Description: "'local': Enable local logs only; 'remote': Enable logging to remote server & IPFIX; 'both': Enable both local & remote logs;",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'ddos_entries_too_many': Too many DDOS entries; 'ddos_entry_added': DDOS entry added; 'ddos_entry_removed': DDOS entry removed; 'ddos_entry_added_to_bgp': DDoS Entry added to BGP; 'ddos_entry_removed_from_bgp': DDoS Entry Removed from BGP; 'ddos_entry_add_to_bgp_failure': DDoS Entry BGP add failures; 'ddos_entry_remove_from_bgp_failure': DDOS entry BGP remove failures; 'ddos_packet_dropped': DDOS Packet Drop;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwDdosProtectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwDdosProtectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwDdosProtection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwDdosProtectionRead(ctx, d, meta)
	}
	return diags
}

func resourceFwDdosProtectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwDdosProtectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwDdosProtection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwDdosProtectionRead(ctx, d, meta)
	}
	return diags
}
func resourceFwDdosProtectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwDdosProtectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwDdosProtection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwDdosProtectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwDdosProtectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwDdosProtection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFwDdosProtectionAction(d []interface{}) edpt.FwDdosProtectionAction {

	count1 := len(d)
	var ret edpt.FwDdosProtectionAction
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

func getObjectFwDdosProtectionDynamicBlacklist(d []interface{}) edpt.FwDdosProtectionDynamicBlacklist {

	count1 := len(d)
	var ret edpt.FwDdosProtectionDynamicBlacklist
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DynamicBlacklistAction = in["dynamic_blacklist_action"].(string)
		ret.Dir = in["dir"].(string)
		ret.Timeout = in["timeout"].(int)
		ret.CpuThreshold = in["cpu_threshold"].(int)
	}
	return ret
}

func getObjectFwDdosProtectionLogging(d []interface{}) edpt.FwDdosProtectionLogging {

	count1 := len(d)
	var ret edpt.FwDdosProtectionLogging
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LoggingAction = in["logging_action"].(string)
		ret.EnableAction = in["enable_action"].(string)
	}
	return ret
}

func getSliceFwDdosProtectionSamplingEnable(d []interface{}) []edpt.FwDdosProtectionSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwDdosProtectionSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwDdosProtectionSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwDdosProtection(d *schema.ResourceData) edpt.FwDdosProtection {
	var ret edpt.FwDdosProtection
	ret.Inst.Action = getObjectFwDdosProtectionAction(d.Get("action").([]interface{}))
	ret.Inst.DynamicBlacklist = getObjectFwDdosProtectionDynamicBlacklist(d.Get("dynamic_blacklist").([]interface{}))
	ret.Inst.Logging = getObjectFwDdosProtectionLogging(d.Get("logging").([]interface{}))
	ret.Inst.SamplingEnable = getSliceFwDdosProtectionSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
