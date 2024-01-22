package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64AlgEspStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat64_alg_esp_stats`: Statistics for the object esp\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nat64AlgEspStatsRead,

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

func resourceCgnv6Nat64AlgEspStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64AlgEspStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64AlgEspStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nat64AlgEspStatsStats := setObjectCgnv6Nat64AlgEspStatsStats(res)
		d.Set("stats", Cgnv6Nat64AlgEspStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nat64AlgEspStatsStats(ret edpt.DataCgnv6Nat64AlgEspStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_created": ret.DtCgnv6Nat64AlgEspStats.Stats.SessionCreated,
			"nat_ip_conflict": ret.DtCgnv6Nat64AlgEspStats.Stats.NatIpConflict,
		},
	}
}

func getObjectCgnv6Nat64AlgEspStatsStats(d []interface{}) edpt.Cgnv6Nat64AlgEspStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Nat64AlgEspStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionCreated = in["session_created"].(int)
		ret.NatIpConflict = in["nat_ip_conflict"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Nat64AlgEspStats(d *schema.ResourceData) edpt.Cgnv6Nat64AlgEspStats {
	var ret edpt.Cgnv6Nat64AlgEspStats

	ret.Stats = getObjectCgnv6Nat64AlgEspStatsStats(d.Get("stats").([]interface{}))
	return ret
}
