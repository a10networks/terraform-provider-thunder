package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgPptpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_pptp_stats`: Statistics for the object pptp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgPptpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"calls_established": {
							Type: schema.TypeInt, Optional: true, Description: "Calls Established",
						},
						"mismatched_pns_call_id": {
							Type: schema.TypeInt, Optional: true, Description: "Mismatched PNS Call ID",
						},
						"gre_sessions_created": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Sessions Created",
						},
						"gre_sessions_freed": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Sessions Freed",
						},
						"no_gre_session_match": {
							Type: schema.TypeInt, Optional: true, Description: "No Matching GRE Session",
						},
						"call_req_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Call ID Mismatch on Call Request",
						},
						"call_reply_pns_call_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Call ID Mismatch on Call Reply",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnAlgPptpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgPptpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgPptpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgPptpStatsStats := setObjectCgnv6LsnAlgPptpStatsStats(res)
		d.Set("stats", Cgnv6LsnAlgPptpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgPptpStatsStats(ret edpt.DataCgnv6LsnAlgPptpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"calls_established":               ret.DtCgnv6LsnAlgPptpStats.Stats.CallsEstablished,
			"mismatched_pns_call_id":          ret.DtCgnv6LsnAlgPptpStats.Stats.MismatchedPnsCallId,
			"gre_sessions_created":            ret.DtCgnv6LsnAlgPptpStats.Stats.GreSessionsCreated,
			"gre_sessions_freed":              ret.DtCgnv6LsnAlgPptpStats.Stats.GreSessionsFreed,
			"no_gre_session_match":            ret.DtCgnv6LsnAlgPptpStats.Stats.NoGreSessionMatch,
			"call_req_pns_call_id_mismatch":   ret.DtCgnv6LsnAlgPptpStats.Stats.CallReqPnsCallIdMismatch,
			"call_reply_pns_call_id_mismatch": ret.DtCgnv6LsnAlgPptpStats.Stats.CallReplyPnsCallIdMismatch,
		},
	}
}

func getObjectCgnv6LsnAlgPptpStatsStats(d []interface{}) edpt.Cgnv6LsnAlgPptpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgPptpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CallsEstablished = in["calls_established"].(int)
		ret.MismatchedPnsCallId = in["mismatched_pns_call_id"].(int)
		ret.GreSessionsCreated = in["gre_sessions_created"].(int)
		ret.GreSessionsFreed = in["gre_sessions_freed"].(int)
		ret.NoGreSessionMatch = in["no_gre_session_match"].(int)
		ret.CallReqPnsCallIdMismatch = in["call_req_pns_call_id_mismatch"].(int)
		ret.CallReplyPnsCallIdMismatch = in["call_reply_pns_call_id_mismatch"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgPptpStats(d *schema.ResourceData) edpt.Cgnv6LsnAlgPptpStats {
	var ret edpt.Cgnv6LsnAlgPptpStats

	ret.Stats = getObjectCgnv6LsnAlgPptpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
