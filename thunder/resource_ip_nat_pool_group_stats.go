package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPoolGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ip_nat_pool_group_stats`: Statistics for the object pool-group\n\n__PLACEHOLDER__",
		ReadContext: resourceIpNatPoolGroupStatsRead,

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

func resourceIpNatPoolGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpNatPoolGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpNatPoolGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpNatPoolGroupStatsStats := setObjectIpNatPoolGroupStatsStats(res)
		d.Set("stats", IpNatPoolGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpNatPoolGroupStatsStats(ret edpt.DataIpNatPoolGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"failed": ret.DtIpNatPoolGroupStats.Stats.Failed,
		},
	}
}

func getObjectIpNatPoolGroupStatsStats(d []interface{}) edpt.IpNatPoolGroupStatsStats {

	count1 := len(d)
	var ret edpt.IpNatPoolGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Failed = in["failed"].(int)
	}
	return ret
}

func dataToEndpointIpNatPoolGroupStats(d *schema.ResourceData) edpt.IpNatPoolGroupStats {
	var ret edpt.IpNatPoolGroupStats

	ret.PoolGroupName = d.Get("pool_group_name").(string)

	ret.Stats = getObjectIpNatPoolGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
