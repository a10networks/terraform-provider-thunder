package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geoloc_list_stats`: Statistics for the object geoloc-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocListStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of Geolocation list",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hit_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_geoloc": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_active": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemGeolocListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocListStatsStats := setObjectSystemGeolocListStatsStats(res)
		d.Set("stats", SystemGeolocListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocListStatsStats(ret edpt.DataSystemGeolocListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hit_count":    ret.DtSystemGeolocListStats.Stats.HitCount,
			"total_geoloc": ret.DtSystemGeolocListStats.Stats.TotalGeoloc,
			"total_active": ret.DtSystemGeolocListStats.Stats.TotalActive,
		},
	}
}

func getObjectSystemGeolocListStatsStats(d []interface{}) edpt.SystemGeolocListStatsStats {

	count1 := len(d)
	var ret edpt.SystemGeolocListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCount = in["hit_count"].(int)
		ret.TotalGeoloc = in["total_geoloc"].(int)
		ret.TotalActive = in["total_active"].(int)
	}
	return ret
}

func dataToEndpointSystemGeolocListStats(d *schema.ResourceData) edpt.SystemGeolocListStats {
	var ret edpt.SystemGeolocListStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSystemGeolocListStatsStats(d.Get("stats").([]interface{}))
	return ret
}
