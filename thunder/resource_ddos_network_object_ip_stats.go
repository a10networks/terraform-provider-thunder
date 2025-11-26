package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectIpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_ip_stats`: Statistics for the object ip\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectIpStatsRead,

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
			"subnet_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "IP Subnet, supported prefix range is from 8 to 32",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}

func resourceDdosNetworkObjectIpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectIpStatsStats := setObjectDdosNetworkObjectIpStatsStats(res)
		d.Set("stats", DdosNetworkObjectIpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectIpStatsStats(ret edpt.DataDdosNetworkObjectIpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_rate": ret.DtDdosNetworkObjectIpStats.Stats.Packet_rate,
			"bit_rate":    ret.DtDdosNetworkObjectIpStats.Stats.Bit_rate,
		},
	}
}

func getObjectDdosNetworkObjectIpStatsStats(d []interface{}) edpt.DdosNetworkObjectIpStatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectIpStats(d *schema.ResourceData) edpt.DdosNetworkObjectIpStats {
	var ret edpt.DdosNetworkObjectIpStats

	ret.Stats = getObjectDdosNetworkObjectIpStatsStats(d.Get("stats").([]interface{}))

	ret.SubnetIpAddr = d.Get("subnet_ip_addr").(string)

	ret.ObjectName = d.Get("object_name").(string)
	return ret
}
