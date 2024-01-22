package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geoloc_stats`: Statistics for the object geoloc\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceSystemGeolocStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocStatsStats := setObjectSystemGeolocStatsStats(res)
		d.Set("stats", SystemGeolocStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocStatsStats(ret edpt.DataSystemGeolocStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectSystemGeolocStatsStats(d []interface{}) edpt.SystemGeolocStatsStats {

	var ret edpt.SystemGeolocStatsStats
	return ret
}

func dataToEndpointSystemGeolocStats(d *schema.ResourceData) edpt.SystemGeolocStats {
	var ret edpt.SystemGeolocStats

	ret.Stats = getObjectSystemGeolocStatsStats(d.Get("stats").([]interface{}))
	return ret
}
