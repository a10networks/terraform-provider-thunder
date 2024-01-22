package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatPoolGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_nat_pool_group_stats`: Statistics for the object pool-group\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6NatPoolGroupStatsRead,

		Schema: map[string]*schema.Schema{
			"pool_group_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool group name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"failed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpv6NatPoolGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatPoolGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatPoolGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6NatPoolGroupStatsStats := setObjectIpv6NatPoolGroupStatsStats(res)
		d.Set("stats", Ipv6NatPoolGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6NatPoolGroupStatsStats(ret edpt.DataIpv6NatPoolGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"failed": ret.DtIpv6NatPoolGroupStats.Stats.Failed,
		},
	}
}

func getObjectIpv6NatPoolGroupStatsStats(d []interface{}) edpt.Ipv6NatPoolGroupStatsStats {

	count1 := len(d)
	var ret edpt.Ipv6NatPoolGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Failed = in["failed"].(int)
	}
	return ret
}

func dataToEndpointIpv6NatPoolGroupStats(d *schema.ResourceData) edpt.Ipv6NatPoolGroupStats {
	var ret edpt.Ipv6NatPoolGroupStats

	ret.PoolGroupName = d.Get("pool_group_name").(string)

	ret.Stats = getObjectIpv6NatPoolGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
