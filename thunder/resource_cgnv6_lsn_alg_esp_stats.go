package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnAlgEspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_alg_esp_stats`: Statistics for the object esp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnAlgEspStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ESP Sessions Created",
						},
						"nat_ip_conflict": {
							Type: schema.TypeInt, Optional: true, Description: "NAT IP Conflict",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnAlgEspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnAlgEspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnAlgEspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnAlgEspStatsStats := setObjectCgnv6LsnAlgEspStatsStats(res)
		d.Set("stats", Cgnv6LsnAlgEspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnAlgEspStatsStats(ret edpt.DataCgnv6LsnAlgEspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtCgnv6LsnAlgEspStats.Stats.SessionCreated,
			"nat_ip_conflict": ret.DtCgnv6LsnAlgEspStats.Stats.NatIpConflict,
		},
	}
}

func getObjectCgnv6LsnAlgEspStatsStats(d []interface{}) edpt.Cgnv6LsnAlgEspStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnAlgEspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
		ret.NatIpConflict = in["nat_ip_conflict"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnAlgEspStats(d *schema.ResourceData) edpt.Cgnv6LsnAlgEspStats {
	var ret edpt.Cgnv6LsnAlgEspStats

	ret.Stats = getObjectCgnv6LsnAlgEspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
