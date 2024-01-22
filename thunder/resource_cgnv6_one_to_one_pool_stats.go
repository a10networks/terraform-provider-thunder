package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOnePoolStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_one_to_one_pool_stats`: Statistics for the object pool\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6OneToOnePoolStatsRead,

		Schema: map[string]*schema.Schema{
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name or pool group",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_address": {
							Type: schema.TypeInt, Optional: true, Description: "Total Address",
						},
						"used_address": {
							Type: schema.TypeInt, Optional: true, Description: "Used Address",
						},
						"free_address": {
							Type: schema.TypeInt, Optional: true, Description: "Free Address",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6OneToOnePoolStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePoolStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6OneToOnePoolStatsStats := setObjectCgnv6OneToOnePoolStatsStats(res)
		d.Set("stats", Cgnv6OneToOnePoolStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6OneToOnePoolStatsStats(ret edpt.DataCgnv6OneToOnePoolStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_address": ret.DtCgnv6OneToOnePoolStats.Stats.TotalAddress,
			"used_address":  ret.DtCgnv6OneToOnePoolStats.Stats.UsedAddress,
			"free_address":  ret.DtCgnv6OneToOnePoolStats.Stats.FreeAddress,
		},
	}
}

func getObjectCgnv6OneToOnePoolStatsStats(d []interface{}) edpt.Cgnv6OneToOnePoolStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOnePoolStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalAddress = in["total_address"].(int)
		ret.UsedAddress = in["used_address"].(int)
		ret.FreeAddress = in["free_address"].(int)
	}
	return ret
}

func dataToEndpointCgnv6OneToOnePoolStats(d *schema.ResourceData) edpt.Cgnv6OneToOnePoolStats {
	var ret edpt.Cgnv6OneToOnePoolStats

	ret.PoolName = d.Get("pool_name").(string)

	ret.Stats = getObjectCgnv6OneToOnePoolStatsStats(d.Get("stats").([]interface{}))
	return ret
}
