package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePartitionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_partition_stats`: Statistics for the object partition\n\n__PLACEHOLDER__",
		ReadContext: resourcePartitionStatsRead,

		Schema: map[string]*schema.Schema{
			"partition_name": {
				Type: schema.TypeString, Required: true, Description: "Object partition name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourcePartitionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePartitionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPartitionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		PartitionStatsStats := setObjectPartitionStatsStats(res)
		d.Set("stats", PartitionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectPartitionStatsStats(ret edpt.DataPartitionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectPartitionStatsStats(d []interface{}) edpt.PartitionStatsStats {

	var ret edpt.PartitionStatsStats
	return ret
}

func dataToEndpointPartitionStats(d *schema.ResourceData) edpt.PartitionStats {
	var ret edpt.PartitionStats

	ret.PartitionName = d.Get("partition_name").(string)

	ret.Stats = getObjectPartitionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
