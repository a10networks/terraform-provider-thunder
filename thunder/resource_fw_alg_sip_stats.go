package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgSipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_alg_sip_stats`: Statistics for the object sip\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAlgSipStatsRead,

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
						"method_options": {
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

func resourceFwAlgSipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgSipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgSipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAlgSipStatsStats := setObjectFwAlgSipStatsStats(res)
		d.Set("stats", FwAlgSipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAlgSipStatsStats(ret edpt.DataFwAlgSipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stat_request":     ret.DtFwAlgSipStats.Stats.StatRequest,
			"stat_response":    ret.DtFwAlgSipStats.Stats.StatResponse,
			"method_register":  ret.DtFwAlgSipStats.Stats.MethodRegister,
			"method_invite":    ret.DtFwAlgSipStats.Stats.MethodInvite,
			"method_ack":       ret.DtFwAlgSipStats.Stats.MethodAck,
			"method_cancel":    ret.DtFwAlgSipStats.Stats.MethodCancel,
			"method_bye":       ret.DtFwAlgSipStats.Stats.MethodBye,
			"method_options":   ret.DtFwAlgSipStats.Stats.MethodOptions,
			"method_prack":     ret.DtFwAlgSipStats.Stats.MethodPrack,
			"method_subscribe": ret.DtFwAlgSipStats.Stats.MethodSubscribe,
			"method_notify":    ret.DtFwAlgSipStats.Stats.MethodNotify,
			"method_publish":   ret.DtFwAlgSipStats.Stats.MethodPublish,
			"method_info":      ret.DtFwAlgSipStats.Stats.MethodInfo,
			"method_refer":     ret.DtFwAlgSipStats.Stats.MethodRefer,
			"method_message":   ret.DtFwAlgSipStats.Stats.MethodMessage,
			"method_update":    ret.DtFwAlgSipStats.Stats.MethodUpdate,
			"method_unknown":   ret.DtFwAlgSipStats.Stats.MethodUnknown,
		},
	}
}

func getObjectFwAlgSipStatsStats(d []interface{}) edpt.FwAlgSipStatsStats {

	count1 := len(d)
	var ret edpt.FwAlgSipStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatRequest = in["stat_request"].(int)
		ret.StatResponse = in["stat_response"].(int)
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

func dataToEndpointFwAlgSipStats(d *schema.ResourceData) edpt.FwAlgSipStats {
	var ret edpt.FwAlgSipStats

	ret.Stats = getObjectFwAlgSipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
