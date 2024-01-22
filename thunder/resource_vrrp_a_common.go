package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpACommon() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_common`: HA VRRP-A Global Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpACommonCreate,
		UpdateContext: resourceVrrpACommonUpdate,
		ReadContext:   resourceVrrpACommonRead,
		DeleteContext: resourceVrrpACommonDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable vrrp-a; 'disable': disable vrrp-a;",
			},
			"arp_retry": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Number of additional gratuitous ARPs sent out after HA failover (1-255, default is 4)",
			},
			"dead_timer": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "VRRP-A dead timer in terms of how many hello messages missed, default is 5 (2-255, default is 5)",
			},
			"device_id": {
				Type: schema.TypeInt, Optional: true, Description: "Unique ID for each VRRP-A box (Device-id number)",
			},
			"disable_default_vrid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable default vrid",
			},
			"forward_l4_packet_on_standby": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enables Layer 2/3 forwarding of Layer 4 traffic on the Standby ACOS device",
			},
			"get_ready_time": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "set get ready time after ax starting up (60-1200, in unit of 100millisec, default is 60)",
			},
			"hello_interval": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "VRRP-A Hello Interval (1-255, in unit of 100millisec, default is 2)",
			},
			"hop_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64, Description: "VRRP-A packet IPv6 header hop-limit (hop-limit, default is 64)",
			},
			"hostid_append_to_vrid": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostid_append_to_vrid_default": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "hostid append to vrid default",
						},
						"hostid_append_to_vrid_value": {
							Type: schema.TypeInt, Optional: true, Description: "hostid append to vrid num",
						},
					},
				},
			},
			"inline_mode_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inline_mode": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Layer 2 Inline Hot Standby Mode",
						},
						"preferred_port": {
							Type: schema.TypeInt, Optional: true, Description: "Preferred ethernet Port",
						},
						"preferred_trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Preferred trunk Port",
						},
					},
				},
			},
			"preemption_delay": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Delay before changing state from Active to Standby (1-255, in unit of 100millisec, default is 60)",
			},
			"restart_time": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Time between restarting ports on standby system after transition",
			},
			"set_id": {
				Type: schema.TypeInt, Optional: true, Description: "Set-ID for HA configuration (Set id from 1 to 15)",
			},
			"track_event_delay": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Delay before changing state after up/down event (Units of 100 milliseconds (default 30))",
			},
			"ttl": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "VRRP-A packet IPv4 header TTL (TTL, default is 128)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpACommonCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpACommonCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpACommon(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpACommonRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpACommonUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpACommonUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpACommon(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpACommonRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpACommonDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpACommonDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpACommon(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpACommonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpACommonRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpACommon(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpACommonHostidAppendToVrid(d []interface{}) edpt.VrrpACommonHostidAppendToVrid {

	count1 := len(d)
	var ret edpt.VrrpACommonHostidAppendToVrid
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostidAppendToVridDefault = in["hostid_append_to_vrid_default"].(int)
		ret.HostidAppendToVridValue = in["hostid_append_to_vrid_value"].(int)
	}
	return ret
}

func getObjectVrrpACommonInlineModeCfg(d []interface{}) edpt.VrrpACommonInlineModeCfg {

	count1 := len(d)
	var ret edpt.VrrpACommonInlineModeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InlineMode = in["inline_mode"].(int)
		ret.PreferredPort = in["preferred_port"].(int)
		ret.PreferredTrunk = in["preferred_trunk"].(int)
	}
	return ret
}

func dataToEndpointVrrpACommon(d *schema.ResourceData) edpt.VrrpACommon {
	var ret edpt.VrrpACommon
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ArpRetry = d.Get("arp_retry").(int)
	ret.Inst.DeadTimer = d.Get("dead_timer").(int)
	ret.Inst.DeviceId = d.Get("device_id").(int)
	ret.Inst.DisableDefaultVrid = d.Get("disable_default_vrid").(int)
	ret.Inst.ForwardL4PacketOnStandby = d.Get("forward_l4_packet_on_standby").(int)
	ret.Inst.GetReadyTime = d.Get("get_ready_time").(int)
	ret.Inst.HelloInterval = d.Get("hello_interval").(int)
	ret.Inst.HopLimit = d.Get("hop_limit").(int)
	ret.Inst.HostidAppendToVrid = getObjectVrrpACommonHostidAppendToVrid(d.Get("hostid_append_to_vrid").([]interface{}))
	ret.Inst.InlineModeCfg = getObjectVrrpACommonInlineModeCfg(d.Get("inline_mode_cfg").([]interface{}))
	ret.Inst.PreemptionDelay = d.Get("preemption_delay").(int)
	ret.Inst.RestartTime = d.Get("restart_time").(int)
	ret.Inst.SetId = d.Get("set_id").(int)
	ret.Inst.TrackEventDelay = d.Get("track_event_delay").(int)
	ret.Inst.Ttl = d.Get("ttl").(int)
	//omit uuid
	return ret
}
