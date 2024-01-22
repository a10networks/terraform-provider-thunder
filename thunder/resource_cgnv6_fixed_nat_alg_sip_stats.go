package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatAlgSipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_alg_sip_stats`: Statistics for the object sip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatAlgSipStatsRead,

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
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatAlgSipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatAlgSipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatAlgSipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatAlgSipStatsStats := setObjectCgnv6FixedNatAlgSipStatsStats(res)
		d.Set("stats", Cgnv6FixedNatAlgSipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatAlgSipStatsStats(ret edpt.DataCgnv6FixedNatAlgSipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"method_register":  ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodRegister,
			"method_invite":    ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodInvite,
			"method_ack":       ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodAck,
			"method_cancel":    ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodCancel,
			"method_bye":       ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodBye,
			"method_options":   ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodOptions,
			"method_prack":     ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodPrack,
			"method_subscribe": ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodSubscribe,
			"method_notify":    ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodNotify,
			"method_publish":   ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodPublish,
			"method_info":      ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodInfo,
			"method_refer":     ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodRefer,
			"method_message":   ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodMessage,
			"method_update":    ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodUpdate,
			"method_unknown":   ret.DtCgnv6FixedNatAlgSipStats.Stats.MethodUnknown,
		},
	}
}

func getObjectCgnv6FixedNatAlgSipStatsStats(d []interface{}) edpt.Cgnv6FixedNatAlgSipStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatAlgSipStatsStats
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
	}
	return ret
}

func dataToEndpointCgnv6FixedNatAlgSipStats(d *schema.ResourceData) edpt.Cgnv6FixedNatAlgSipStats {
	var ret edpt.Cgnv6FixedNatAlgSipStats

	ret.Stats = getObjectCgnv6FixedNatAlgSipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
