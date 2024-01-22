package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatPoolStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_nat_pool_stats`: Statistics for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6NatPoolStatsRead,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_freed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"failed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpv6NatPoolStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6NatPoolStatsStats := setObjectIpv6NatPoolStatsStats(res)
		d.Set("stats", Ipv6NatPoolStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6NatPoolStatsStats(ret edpt.DataIpv6NatPoolStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_usage":  ret.DtIpv6NatPoolStats.Stats.PortUsage,
			"total_used":  ret.DtIpv6NatPoolStats.Stats.TotalUsed,
			"total_freed": ret.DtIpv6NatPoolStats.Stats.TotalFreed,
			"failed":      ret.DtIpv6NatPoolStats.Stats.Failed,
		},
	}
}

func getObjectIpv6NatPoolStatsStats(d []interface{}) edpt.Ipv6NatPoolStatsStats {

	count1 := len(d)
	var ret edpt.Ipv6NatPoolStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortUsage = in["port_usage"].(int)
		ret.TotalUsed = in["total_used"].(int)
		ret.TotalFreed = in["total_freed"].(int)
		ret.Failed = in["failed"].(int)
	}
	return ret
}

func dataToEndpointIpv6NatPoolStats(d *schema.ResourceData) edpt.Ipv6NatPoolStats {
	var ret edpt.Ipv6NatPoolStats

	ret.PoolName = d.Get("pool_name").(string)

	ret.Stats = getObjectIpv6NatPoolStatsStats(d.Get("stats").([]interface{}))
	return ret
}
