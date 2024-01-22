package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatNatGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_nat_nat_global_stats`: Statistics for the object nat-global\n\n__PLACEHOLDER__",
		ReadContext: resourceIpNatNatGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceIpNatNatGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatNatGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatNatGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpNatNatGlobalStatsStats := setObjectIpNatNatGlobalStatsStats(res)
		d.Set("stats", IpNatNatGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpNatNatGlobalStatsStats(ret edpt.DataIpNatNatGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectIpNatNatGlobalStatsStats(d []interface{}) edpt.IpNatNatGlobalStatsStats {

	var ret edpt.IpNatNatGlobalStatsStats
	return ret
}

func dataToEndpointIpNatNatGlobalStats(d *schema.ResourceData) edpt.IpNatNatGlobalStats {
	var ret edpt.IpNatNatGlobalStats

	ret.Stats = getObjectIpNatNatGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
