package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpSynCookieStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_tcp_syn_cookie_stats`: Statistics for the object syn-cookie\n\n__PLACEHOLDER__",
		ReadContext: resourceFwTcpSynCookieStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"syn_ack_sent": {
							Type: schema.TypeInt, Optional: true, Description: "SYN cookie SYN ACK sent",
						},
						"verification_passed": {
							Type: schema.TypeInt, Optional: true, Description: "SYN cookie verification passed",
						},
						"verification_failed": {
							Type: schema.TypeInt, Optional: true, Description: "SYN cookie verification failed",
						},
					},
				},
			},
		},
	}
}

func resourceFwTcpSynCookieStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpSynCookieStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpSynCookieStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwTcpSynCookieStatsStats := setObjectFwTcpSynCookieStatsStats(res)
		d.Set("stats", FwTcpSynCookieStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwTcpSynCookieStatsStats(ret edpt.DataFwTcpSynCookieStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"syn_ack_sent":        ret.DtFwTcpSynCookieStats.Stats.Syn_ack_sent,
			"verification_passed": ret.DtFwTcpSynCookieStats.Stats.Verification_passed,
			"verification_failed": ret.DtFwTcpSynCookieStats.Stats.Verification_failed,
		},
	}
}

func getObjectFwTcpSynCookieStatsStats(d []interface{}) edpt.FwTcpSynCookieStatsStats {

	count1 := len(d)
	var ret edpt.FwTcpSynCookieStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn_ack_sent = in["syn_ack_sent"].(int)
		ret.Verification_passed = in["verification_passed"].(int)
		ret.Verification_failed = in["verification_failed"].(int)
	}
	return ret
}

func dataToEndpointFwTcpSynCookieStats(d *schema.ResourceData) edpt.FwTcpSynCookieStats {
	var ret edpt.FwTcpSynCookieStats

	ret.Stats = getObjectFwTcpSynCookieStatsStats(d.Get("stats").([]interface{}))
	return ret
}
