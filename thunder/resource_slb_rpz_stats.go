package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRpzStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rpz_stats`: Statistics for the object rpz\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRpzStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set_bw_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total RPZ Set Class-list Error",
						},
						"parse_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total RPZ Parse Error",
						},
					},
				},
			},
		},
	}
}

func resourceSlbRpzStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpzStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRpzStatsStats := setObjectSlbRpzStatsStats(res)
		d.Set("stats", SlbRpzStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRpzStatsStats(ret edpt.DataSlbRpzStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"set_bw_error": ret.DtSlbRpzStats.Stats.Set_bw_error,
			"parse_error":  ret.DtSlbRpzStats.Stats.Parse_error,
		},
	}
}

func getObjectSlbRpzStatsStats(d []interface{}) edpt.SlbRpzStatsStats {

	count1 := len(d)
	var ret edpt.SlbRpzStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Set_bw_error = in["set_bw_error"].(int)
		ret.Parse_error = in["parse_error"].(int)
	}
	return ret
}

func dataToEndpointSlbRpzStats(d *schema.ResourceData) edpt.SlbRpzStats {
	var ret edpt.SlbRpzStats

	ret.Stats = getObjectSlbRpzStatsStats(d.Get("stats").([]interface{}))
	return ret
}
