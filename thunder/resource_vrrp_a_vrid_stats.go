package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vrrp_a_vrid_stats`: Statistics for the object vrid\n\n__PLACEHOLDER__",
		ReadContext: resourceVrrpAVridStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"associated_vip_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of vips associated to vrid",
						},
						"associated_vport_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of vports associated to vrid",
						},
						"associated_natpool_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of nat pools associated to vrid",
						},
					},
				},
			},
			"vrid_val": {
				Type: schema.TypeInt, Required: true, Description: "Specify ha VRRP-A vrid",
			},
		},
	}
}

func resourceVrrpAVridStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VrrpAVridStatsStats := setObjectVrrpAVridStatsStats(res)
		d.Set("stats", VrrpAVridStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVrrpAVridStatsStats(ret edpt.DataVrrpAVridStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"associated_vip_count":     ret.DtVrrpAVridStats.Stats.Associated_vip_count,
			"associated_vport_count":   ret.DtVrrpAVridStats.Stats.Associated_vport_count,
			"associated_natpool_count": ret.DtVrrpAVridStats.Stats.Associated_natpool_count,
		},
	}
}

func getObjectVrrpAVridStatsStats(d []interface{}) edpt.VrrpAVridStatsStats {

	count1 := len(d)
	var ret edpt.VrrpAVridStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Associated_vip_count = in["associated_vip_count"].(int)
		ret.Associated_vport_count = in["associated_vport_count"].(int)
		ret.Associated_natpool_count = in["associated_natpool_count"].(int)
	}
	return ret
}

func dataToEndpointVrrpAVridStats(d *schema.ResourceData) edpt.VrrpAVridStats {
	var ret edpt.VrrpAVridStats

	ret.Stats = getObjectVrrpAVridStatsStats(d.Get("stats").([]interface{}))

	ret.VridVal = d.Get("vrid_val").(int)
	return ret
}
