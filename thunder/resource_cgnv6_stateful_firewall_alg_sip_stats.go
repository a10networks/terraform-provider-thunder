package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgSipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_stateful_firewall_alg_sip_stats`: Statistics for the object sip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6StatefulFirewallAlgSipStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stat_request": {
							Type: schema.TypeInt, Optional: true, Description: "Request Received",
						},
						"stat_response": {
							Type: schema.TypeInt, Optional: true, Description: "Response Received",
						},
						"method_register": {
							Type: schema.TypeInt, Optional: true, Description: "Method REGISTER",
						},
						"method_invite": {
							Type: schema.TypeInt, Optional: true, Description: "Method INVITE",
						},
						"method_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Method ACK",
						},
						"method_cancel": {
							Type: schema.TypeInt, Optional: true, Description: "Method CANCEL",
						},
						"method_bye": {
							Type: schema.TypeInt, Optional: true, Description: "Method BYE",
						},
						"method_port_config": {
							Type: schema.TypeInt, Optional: true, Description: "Method OPTIONS",
						},
						"method_prack": {
							Type: schema.TypeInt, Optional: true, Description: "Method PRACK",
						},
						"method_subscribe": {
							Type: schema.TypeInt, Optional: true, Description: "Method SUBSCRIBE",
						},
						"method_notify": {
							Type: schema.TypeInt, Optional: true, Description: "Method NOTIFY",
						},
						"method_publish": {
							Type: schema.TypeInt, Optional: true, Description: "Method PUBLISH",
						},
						"method_info": {
							Type: schema.TypeInt, Optional: true, Description: "Method INFO",
						},
						"method_refer": {
							Type: schema.TypeInt, Optional: true, Description: "Method REFER",
						},
						"method_message": {
							Type: schema.TypeInt, Optional: true, Description: "Method MESSAGE",
						},
						"method_update": {
							Type: schema.TypeInt, Optional: true, Description: "Method UPDATE",
						},
						"method_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Method Unknown",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6StatefulFirewallAlgSipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgSipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgSipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6StatefulFirewallAlgSipStatsStats := setObjectCgnv6StatefulFirewallAlgSipStatsStats(res)
		d.Set("stats", Cgnv6StatefulFirewallAlgSipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6StatefulFirewallAlgSipStatsStats(ret edpt.DataCgnv6StatefulFirewallAlgSipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stat_request":       ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.StatRequest,
			"stat_response":      ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.StatResponse,
			"method_register":    ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodRegister,
			"method_invite":      ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodInvite,
			"method_ack":         ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodAck,
			"method_cancel":      ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodCancel,
			"method_bye":         ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodBye,
			"method_port_config": ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodPortConfig,
			"method_prack":       ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodPrack,
			"method_subscribe":   ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodSubscribe,
			"method_notify":      ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodNotify,
			"method_publish":     ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodPublish,
			"method_info":        ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodInfo,
			"method_refer":       ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodRefer,
			"method_message":     ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodMessage,
			"method_update":      ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodUpdate,
			"method_unknown":     ret.DtCgnv6StatefulFirewallAlgSipStats.Stats.MethodUnknown,
		},
	}
}

func getObjectCgnv6StatefulFirewallAlgSipStatsStats(d []interface{}) edpt.Cgnv6StatefulFirewallAlgSipStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6StatefulFirewallAlgSipStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatRequest = in["stat_request"].(int)
		ret.StatResponse = in["stat_response"].(int)
		ret.MethodRegister = in["method_register"].(int)
		ret.MethodInvite = in["method_invite"].(int)
		ret.MethodAck = in["method_ack"].(int)
		ret.MethodCancel = in["method_cancel"].(int)
		ret.MethodBye = in["method_bye"].(int)
		ret.MethodPortConfig = in["method_port_config"].(int)
		ret.MethodPrack = in["method_prack"].(int)
		ret.MethodSubscribe = in["method_subscribe"].(int)
		ret.MethodNotify = in["method_notify"].(int)
		ret.MethodPublish = in["method_publish"].(int)
		ret.MethodInfo = in["method_info"].(int)
		ret.MethodRefer = in["method_refer"].(int)
		ret.MethodMessage = in["method_message"].(int)
		ret.MethodUpdate = in["method_update"].(int)
		ret.MethodUnknown = in["method_unknown"].(int)
	}
	return ret
}

func dataToEndpointCgnv6StatefulFirewallAlgSipStats(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgSipStats {
	var ret edpt.Cgnv6StatefulFirewallAlgSipStats

	ret.Stats = getObjectCgnv6StatefulFirewallAlgSipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
