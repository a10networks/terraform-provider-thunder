package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCrlSrcipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_crl_srcip_stats`: Statistics for the object crl-srcip\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbCrlSrcipStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sessions_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Sessions allocated",
						},
						"sessions_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Sessions freed",
						},
						"out_of_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Out of sessions",
						},
						"too_many_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "Too many sessions consumed",
						},
						"called": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold check count",
						},
						"permitted": {
							Type: schema.TypeInt, Optional: true, Description: "Honor threshold  count",
						},
						"threshold_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold exceeded count",
						},
						"lockout_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Lockout drops",
						},
						"log_msg_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Log messages sent",
						},
					},
				},
			},
		},
	}
}

func resourceSlbCrlSrcipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCrlSrcipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCrlSrcipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbCrlSrcipStatsStats := setObjectSlbCrlSrcipStatsStats(res)
		d.Set("stats", SlbCrlSrcipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbCrlSrcipStatsStats(ret edpt.DataSlbCrlSrcipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sessions_alloc":    ret.DtSlbCrlSrcipStats.Stats.Sessions_alloc,
			"sessions_freed":    ret.DtSlbCrlSrcipStats.Stats.Sessions_freed,
			"out_of_sessions":   ret.DtSlbCrlSrcipStats.Stats.Out_of_sessions,
			"too_many_sessions": ret.DtSlbCrlSrcipStats.Stats.Too_many_sessions,
			"called":            ret.DtSlbCrlSrcipStats.Stats.Called,
			"permitted":         ret.DtSlbCrlSrcipStats.Stats.Permitted,
			"threshold_exceed":  ret.DtSlbCrlSrcipStats.Stats.Threshold_exceed,
			"lockout_drop":      ret.DtSlbCrlSrcipStats.Stats.Lockout_drop,
			"log_msg_sent":      ret.DtSlbCrlSrcipStats.Stats.Log_msg_sent,
		},
	}
}

func getObjectSlbCrlSrcipStatsStats(d []interface{}) edpt.SlbCrlSrcipStatsStats {

	count1 := len(d)
	var ret edpt.SlbCrlSrcipStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Sessions_alloc = in["sessions_alloc"].(int)
		ret.Sessions_freed = in["sessions_freed"].(int)
		ret.Out_of_sessions = in["out_of_sessions"].(int)
		ret.Too_many_sessions = in["too_many_sessions"].(int)
		ret.Called = in["called"].(int)
		ret.Permitted = in["permitted"].(int)
		ret.Threshold_exceed = in["threshold_exceed"].(int)
		ret.Lockout_drop = in["lockout_drop"].(int)
		ret.Log_msg_sent = in["log_msg_sent"].(int)
	}
	return ret
}

func dataToEndpointSlbCrlSrcipStats(d *schema.ResourceData) edpt.SlbCrlSrcipStats {
	var ret edpt.SlbCrlSrcipStats

	ret.Stats = getObjectSlbCrlSrcipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
