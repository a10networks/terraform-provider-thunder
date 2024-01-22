package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOneGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_one_to_one_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6OneToOneGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_map_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total One-to-One Address Mapping Allocated",
						},
						"total_map_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total One-to-One Address Mapping Freed",
						},
						"map_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "One-to-One Address Mapping Allocation Failure",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6OneToOneGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOneGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOneGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6OneToOneGlobalStatsStats := setObjectCgnv6OneToOneGlobalStatsStats(res)
		d.Set("stats", Cgnv6OneToOneGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6OneToOneGlobalStatsStats(ret edpt.DataCgnv6OneToOneGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_map_allocated": ret.DtCgnv6OneToOneGlobalStats.Stats.TotalMapAllocated,
			"total_map_freed":     ret.DtCgnv6OneToOneGlobalStats.Stats.TotalMapFreed,
			"map_alloc_failure":   ret.DtCgnv6OneToOneGlobalStats.Stats.MapAllocFailure,
		},
	}
}

func getObjectCgnv6OneToOneGlobalStatsStats(d []interface{}) edpt.Cgnv6OneToOneGlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6OneToOneGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalMapAllocated = in["total_map_allocated"].(int)
		ret.TotalMapFreed = in["total_map_freed"].(int)
		ret.MapAllocFailure = in["map_alloc_failure"].(int)
	}
	return ret
}

func dataToEndpointCgnv6OneToOneGlobalStats(d *schema.ResourceData) edpt.Cgnv6OneToOneGlobalStats {
	var ret edpt.Cgnv6OneToOneGlobalStats

	ret.Stats = getObjectCgnv6OneToOneGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
