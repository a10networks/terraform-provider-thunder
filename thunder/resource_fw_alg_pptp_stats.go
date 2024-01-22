package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwAlgPptpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_alg_pptp_stats`: Statistics for the object pptp\n\n__PLACEHOLDER__",
		ReadContext: resourceFwAlgPptpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"calls_established": {
							Type: schema.TypeInt, Optional: true, Description: "Calls Established",
						},
						"call_req_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Call ID Mismatch on Call Request",
						},
						"call_reply_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Call ID Mismatch on Call Reply",
						},
						"gre_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Session Created",
						},
						"gre_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Session Freed",
						},
					},
				},
			},
		},
	}
}

func resourceFwAlgPptpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAlgPptpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwAlgPptpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwAlgPptpStatsStats := setObjectFwAlgPptpStatsStats(res)
		d.Set("stats", FwAlgPptpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwAlgPptpStatsStats(ret edpt.DataFwAlgPptpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"calls_established":               ret.DtFwAlgPptpStats.Stats.CallsEstablished,
			"call_req_pns_call_id_mismatch":   ret.DtFwAlgPptpStats.Stats.CallReqPnsCallIdMismatch,
			"call_reply_pns_call_id_mismatch": ret.DtFwAlgPptpStats.Stats.CallReplyPnsCallIdMismatch,
			"gre_session_created":             ret.DtFwAlgPptpStats.Stats.GreSessionCreated,
			"gre_session_freed":               ret.DtFwAlgPptpStats.Stats.GreSessionFreed,
		},
	}
}

func getObjectFwAlgPptpStatsStats(d []interface{}) edpt.FwAlgPptpStatsStats {

	count1 := len(d)
	var ret edpt.FwAlgPptpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CallsEstablished = in["calls_established"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
		ret.GreSessionCreated = in["gre_session_created"].(int)
		ret.GreSessionFreed = in["gre_session_freed"].(int)
	}
	return ret
}

func dataToEndpointFwAlgPptpStats(d *schema.ResourceData) edpt.FwAlgPptpStats {
	var ret edpt.FwAlgPptpStats

	ret.Stats = getObjectFwAlgPptpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
