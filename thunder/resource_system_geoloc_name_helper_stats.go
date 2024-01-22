package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocNameHelperStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geoloc_name_helper_stats`: Statistics for the object geoloc-name-helper\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocNameHelperStatsRead,

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

func resourceSystemGeolocNameHelperStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocNameHelperStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocNameHelperStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocNameHelperStatsStats := setObjectSystemGeolocNameHelperStatsStats(res)
		d.Set("stats", SystemGeolocNameHelperStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocNameHelperStatsStats(ret edpt.DataSystemGeolocNameHelperStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectSystemGeolocNameHelperStatsStats(d []interface{}) edpt.SystemGeolocNameHelperStatsStats {

	var ret edpt.SystemGeolocNameHelperStatsStats
	return ret
}

func dataToEndpointSystemGeolocNameHelperStats(d *schema.ResourceData) edpt.SystemGeolocNameHelperStats {
	var ret edpt.SystemGeolocNameHelperStats

	ret.Stats = getObjectSystemGeolocNameHelperStatsStats(d.Get("stats").([]interface{}))
	return ret
}
