package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkIpEntityStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_ip_entity_stats`: Statistics for the object network-ip-entity\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkIpEntityStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_rate": {
							Type: schema.TypeInt, Optional: true, Description: "PPS",
						},
						"bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "B(bits)PS",
						},
					},
				},
			},
		},
	}
}

func resourceDdosNetworkIpEntityStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkIpEntityStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkIpEntityStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkIpEntityStatsStats := setObjectDdosNetworkIpEntityStatsStats(res)
		d.Set("stats", DdosNetworkIpEntityStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkIpEntityStatsStats(ret edpt.DataDdosNetworkIpEntityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_rate": ret.DtDdosNetworkIpEntityStats.Stats.Packet_rate,
			"bit_rate":    ret.DtDdosNetworkIpEntityStats.Stats.Bit_rate,
		},
	}
}

func getObjectDdosNetworkIpEntityStatsStats(d []interface{}) edpt.DdosNetworkIpEntityStatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkIpEntityStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkIpEntityStats(d *schema.ResourceData) edpt.DdosNetworkIpEntityStats {
	var ret edpt.DdosNetworkIpEntityStats

	ret.Stats = getObjectDdosNetworkIpEntityStatsStats(d.Get("stats").([]interface{}))
	return ret
}
