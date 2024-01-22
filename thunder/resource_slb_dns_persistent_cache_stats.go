package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsPersistentCacheStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dns_persistent_cache_stats`: Statistics for the object dns-persistent-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDnsPersistentCacheStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Total persistent cache entry",
						},
						"entry_saved": {
							Type: schema.TypeInt, Optional: true, Description: "Total saved cache entry",
						},
						"entry_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Total deleted cache entry",
						},
						"database_busy": {
							Type: schema.TypeInt, Optional: true, Description: "Database busy",
						},
						"database_error": {
							Type: schema.TypeInt, Optional: true, Description: "Database error",
						},
					},
				},
			},
		},
	}
}

func resourceSlbDnsPersistentCacheStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsPersistentCacheStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsPersistentCacheStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDnsPersistentCacheStatsStats := setObjectSlbDnsPersistentCacheStatsStats(res)
		d.Set("stats", SlbDnsPersistentCacheStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDnsPersistentCacheStatsStats(ret edpt.DataSlbDnsPersistentCacheStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_entry":    ret.DtSlbDnsPersistentCacheStats.Stats.Total_entry,
			"entry_saved":    ret.DtSlbDnsPersistentCacheStats.Stats.Entry_saved,
			"entry_deleted":  ret.DtSlbDnsPersistentCacheStats.Stats.Entry_deleted,
			"database_busy":  ret.DtSlbDnsPersistentCacheStats.Stats.Database_busy,
			"database_error": ret.DtSlbDnsPersistentCacheStats.Stats.Database_error,
		},
	}
}

func getObjectSlbDnsPersistentCacheStatsStats(d []interface{}) edpt.SlbDnsPersistentCacheStatsStats {

	count1 := len(d)
	var ret edpt.SlbDnsPersistentCacheStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_entry = in["total_entry"].(int)
		ret.Entry_saved = in["entry_saved"].(int)
		ret.Entry_deleted = in["entry_deleted"].(int)
		ret.Database_busy = in["database_busy"].(int)
		ret.Database_error = in["database_error"].(int)
	}
	return ret
}

func dataToEndpointSlbDnsPersistentCacheStats(d *schema.ResourceData) edpt.SlbDnsPersistentCacheStats {
	var ret edpt.SlbDnsPersistentCacheStats

	ret.Stats = getObjectSlbDnsPersistentCacheStatsStats(d.Get("stats").([]interface{}))
	return ret
}
