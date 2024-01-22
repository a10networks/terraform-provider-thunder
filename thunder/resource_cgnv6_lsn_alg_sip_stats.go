package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgSipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_sip_stats`: Statistics for the object sip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgSipStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method_register": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method REGISTER",
						},
						"method_invite": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method INVITE",
						},
						"method_ack": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method ACK",
						},
						"method_cancel": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method CANCEL",
						},
						"method_bye": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method BYE",
						},
						"method_options": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method OPTIONS",
						},
						"method_prack": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method PRACK",
						},
						"method_subscribe": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method SUBSCRIBE",
						},
						"method_notify": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method NOTIFY",
						},
						"method_publish": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method PUBLISH",
						},
						"method_info": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method INFO",
						},
						"method_refer": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method REFER",
						},
						"method_message": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method MESSAGE",
						},
						"method_update": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method UPDATE",
						},
						"method_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Method UNKNOWN",
						},
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Message Parse Error",
						},
						"tcp_out_of_order_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-of-Order Drop",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnAlgSipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgSipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgSipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgSipStatsStats := setObjectCgnv6LsnAlgSipStatsStats(res)
		d.Set("stats", Cgnv6LsnAlgSipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgSipStatsStats(ret edpt.DataCgnv6LsnAlgSipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"method_register":       ret.DtCgnv6LsnAlgSipStats.Stats.MethodRegister,
			"method_invite":         ret.DtCgnv6LsnAlgSipStats.Stats.MethodInvite,
			"method_ack":            ret.DtCgnv6LsnAlgSipStats.Stats.MethodAck,
			"method_cancel":         ret.DtCgnv6LsnAlgSipStats.Stats.MethodCancel,
			"method_bye":            ret.DtCgnv6LsnAlgSipStats.Stats.MethodBye,
			"method_options":        ret.DtCgnv6LsnAlgSipStats.Stats.MethodOptions,
			"method_prack":          ret.DtCgnv6LsnAlgSipStats.Stats.MethodPrack,
			"method_subscribe":      ret.DtCgnv6LsnAlgSipStats.Stats.MethodSubscribe,
			"method_notify":         ret.DtCgnv6LsnAlgSipStats.Stats.MethodNotify,
			"method_publish":        ret.DtCgnv6LsnAlgSipStats.Stats.MethodPublish,
			"method_info":           ret.DtCgnv6LsnAlgSipStats.Stats.MethodInfo,
			"method_refer":          ret.DtCgnv6LsnAlgSipStats.Stats.MethodRefer,
			"method_message":        ret.DtCgnv6LsnAlgSipStats.Stats.MethodMessage,
			"method_update":         ret.DtCgnv6LsnAlgSipStats.Stats.MethodUpdate,
			"method_unknown":        ret.DtCgnv6LsnAlgSipStats.Stats.MethodUnknown,
			"parse_error":           ret.DtCgnv6LsnAlgSipStats.Stats.ParseError,
			"tcp_out_of_order_drop": ret.DtCgnv6LsnAlgSipStats.Stats.TcpOutOfOrderDrop,
		},
	}
}

func getObjectCgnv6LsnAlgSipStatsStats(d []interface{}) edpt.Cgnv6LsnAlgSipStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgSipStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MethodRegister = in["method_register"].(int)
		ret.MethodInvite = in["method_invite"].(int)
		ret.MethodAck = in["method_ack"].(int)
		ret.MethodCancel = in["method_cancel"].(int)
		ret.MethodBye = in["method_bye"].(int)
		ret.MethodOptions = in["method_options"].(int)
		ret.MethodPrack = in["method_prack"].(int)
		ret.MethodSubscribe = in["method_subscribe"].(int)
		ret.MethodNotify = in["method_notify"].(int)
		ret.MethodPublish = in["method_publish"].(int)
		ret.MethodInfo = in["method_info"].(int)
		ret.MethodRefer = in["method_refer"].(int)
		ret.MethodMessage = in["method_message"].(int)
		ret.MethodUpdate = in["method_update"].(int)
		ret.MethodUnknown = in["method_unknown"].(int)
		ret.ParseError = in["parse_error"].(int)
		ret.TcpOutOfOrderDrop = in["tcp_out_of_order_drop"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgSipStats(d *schema.ResourceData) edpt.Cgnv6LsnAlgSipStats {
	var ret edpt.Cgnv6LsnAlgSipStats

	ret.Stats = getObjectCgnv6LsnAlgSipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
