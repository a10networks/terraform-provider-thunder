package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPoolStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_nat_pool_stats`: Statistics for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceIpNatPoolStatsRead,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name or pool group",
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

func resourceIpNatPoolStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpNatPoolStatsStats := setObjectIpNatPoolStatsStats(res)
		d.Set("stats", IpNatPoolStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpNatPoolStatsStats(ret edpt.DataIpNatPoolStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_usage":  ret.DtIpNatPoolStats.Stats.PortUsage,
			"total_used":  ret.DtIpNatPoolStats.Stats.TotalUsed,
			"total_freed": ret.DtIpNatPoolStats.Stats.TotalFreed,
			"failed":      ret.DtIpNatPoolStats.Stats.Failed,
		},
	}
}

func getObjectIpNatPoolStatsStats(d []interface{}) edpt.IpNatPoolStatsStats {

	count1 := len(d)
	var ret edpt.IpNatPoolStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortUsage = in["port_usage"].(int)
		ret.TotalUsed = in["total_used"].(int)
		ret.TotalFreed = in["total_freed"].(int)
		ret.Failed = in["failed"].(int)
	}
	return ret
}

func dataToEndpointIpNatPoolStats(d *schema.ResourceData) edpt.IpNatPoolStats {
	var ret edpt.IpNatPoolStats

	ret.PoolName = d.Get("pool_name").(string)

	ret.Stats = getObjectIpNatPoolStatsStats(d.Get("stats").([]interface{}))
	return ret
}
