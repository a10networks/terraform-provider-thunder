package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgPptpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_stateful_firewall_alg_pptp_stats`: Statistics for the object pptp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6StatefulFirewallAlgPptpStatsRead,

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

func resourceCgnv6StatefulFirewallAlgPptpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgPptpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgPptpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6StatefulFirewallAlgPptpStatsStats := setObjectCgnv6StatefulFirewallAlgPptpStatsStats(res)
		d.Set("stats", Cgnv6StatefulFirewallAlgPptpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6StatefulFirewallAlgPptpStatsStats(ret edpt.DataCgnv6StatefulFirewallAlgPptpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"calls_established":               ret.DtCgnv6StatefulFirewallAlgPptpStats.Stats.CallsEstablished,
			"call_req_pns_call_id_mismatch":   ret.DtCgnv6StatefulFirewallAlgPptpStats.Stats.CallReqPnsCallIdMismatch,
			"call_reply_pns_call_id_mismatch": ret.DtCgnv6StatefulFirewallAlgPptpStats.Stats.CallReplyPnsCallIdMismatch,
			"gre_session_created":             ret.DtCgnv6StatefulFirewallAlgPptpStats.Stats.GreSessionCreated,
			"gre_session_freed":               ret.DtCgnv6StatefulFirewallAlgPptpStats.Stats.GreSessionFreed,
		},
	}
}

func getObjectCgnv6StatefulFirewallAlgPptpStatsStats(d []interface{}) edpt.Cgnv6StatefulFirewallAlgPptpStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6StatefulFirewallAlgPptpStatsStats
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

func dataToEndpointCgnv6StatefulFirewallAlgPptpStats(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgPptpStats {
	var ret edpt.Cgnv6StatefulFirewallAlgPptpStats

	ret.Stats = getObjectCgnv6StatefulFirewallAlgPptpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
