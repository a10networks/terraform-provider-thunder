package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryStatisticsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_statistics_stats`: Statistics for the object statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryStatisticsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cloud_cache_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cloud_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rtu_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"lookup_latency": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"db_mem": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rtu_cache_mem": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"lookup_cache_mem": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryStatisticsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryStatisticsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryStatisticsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryStatisticsStatsStats := setObjectWebCategoryStatisticsStatsStats(res)
		d.Set("stats", WebCategoryStatisticsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryStatisticsStatsStats(ret edpt.DataWebCategoryStatisticsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"db_lookup":          ret.DtWebCategoryStatisticsStats.Stats.DbLookup,
			"cloud_cache_lookup": ret.DtWebCategoryStatisticsStats.Stats.CloudCacheLookup,
			"cloud_lookup":       ret.DtWebCategoryStatisticsStats.Stats.CloudLookup,
			"rtu_lookup":         ret.DtWebCategoryStatisticsStats.Stats.RtuLookup,
			"lookup_latency":     ret.DtWebCategoryStatisticsStats.Stats.LookupLatency,
			"db_mem":             ret.DtWebCategoryStatisticsStats.Stats.DbMem,
			"rtu_cache_mem":      ret.DtWebCategoryStatisticsStats.Stats.RtuCacheMem,
			"lookup_cache_mem":   ret.DtWebCategoryStatisticsStats.Stats.LookupCacheMem,
		},
	}
}

func getObjectWebCategoryStatisticsStatsStats(d []interface{}) edpt.WebCategoryStatisticsStatsStats {

	count1 := len(d)
	var ret edpt.WebCategoryStatisticsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DbLookup = in["db_lookup"].(int)
		ret.CloudCacheLookup = in["cloud_cache_lookup"].(int)
		ret.CloudLookup = in["cloud_lookup"].(int)
		ret.RtuLookup = in["rtu_lookup"].(int)
		ret.LookupLatency = in["lookup_latency"].(int)
		ret.DbMem = in["db_mem"].(int)
		ret.RtuCacheMem = in["rtu_cache_mem"].(int)
		ret.LookupCacheMem = in["lookup_cache_mem"].(int)
	}
	return ret
}

func dataToEndpointWebCategoryStatisticsStats(d *schema.ResourceData) edpt.WebCategoryStatisticsStats {
	var ret edpt.WebCategoryStatisticsStats

	ret.Stats = getObjectWebCategoryStatisticsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
