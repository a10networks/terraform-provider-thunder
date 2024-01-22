package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtectStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_protect_stats`: Statistics for the object protect\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosProtectStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rr_sess_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS RR Session Lookup",
						},
						"dcmsg_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS DCMSG Pkt",
						},
						"runtime_enable": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS Runtime Enable",
						},
						"runtime_disable": {
							Type: schema.TypeInt, Optional: true, Description: "DDOS Runtime Disable",
						},
					},
				},
			},
		},
	}
}

func resourceDdosProtectStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtectStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosProtectStatsStats := setObjectDdosProtectStatsStats(res)
		d.Set("stats", DdosProtectStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosProtectStatsStats(ret edpt.DataDdosProtectStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rr_sess_lookup":  ret.DtDdosProtectStats.Stats.Rr_sess_lookup,
			"dcmsg_pkt":       ret.DtDdosProtectStats.Stats.Dcmsg_pkt,
			"runtime_enable":  ret.DtDdosProtectStats.Stats.Runtime_enable,
			"runtime_disable": ret.DtDdosProtectStats.Stats.Runtime_disable,
		},
	}
}

func getObjectDdosProtectStatsStats(d []interface{}) edpt.DdosProtectStatsStats {

	count1 := len(d)
	var ret edpt.DdosProtectStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rr_sess_lookup = in["rr_sess_lookup"].(int)
		ret.Dcmsg_pkt = in["dcmsg_pkt"].(int)
		ret.Runtime_enable = in["runtime_enable"].(int)
		ret.Runtime_disable = in["runtime_disable"].(int)
	}
	return ret
}

func dataToEndpointDdosProtectStats(d *schema.ResourceData) edpt.DdosProtectStats {
	var ret edpt.DdosProtectStats

	ret.Stats = getObjectDdosProtectStatsStats(d.Get("stats").([]interface{}))
	return ret
}
