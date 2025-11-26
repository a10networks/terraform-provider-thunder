package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSubNetworkSubNetworkV4Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_sub_network_sub_network_v4_stats`: Statistics for the object sub-network-v4\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectSubNetworkSubNetworkV4StatsRead,

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
				Type: schema.TypeString, Required: true, Description: "IPv4 Subnet/host, supported prefix range is from 24 to 32",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}

func resourceDdosNetworkObjectSubNetworkSubNetworkV4StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV4StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectSubNetworkSubNetworkV4StatsStats := setObjectDdosNetworkObjectSubNetworkSubNetworkV4StatsStats(res)
		d.Set("stats", DdosNetworkObjectSubNetworkSubNetworkV4StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectSubNetworkSubNetworkV4StatsStats(ret edpt.DataDdosNetworkObjectSubNetworkSubNetworkV4Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_rate": ret.DtDdosNetworkObjectSubNetworkSubNetworkV4Stats.Stats.Packet_rate,
			"bit_rate":    ret.DtDdosNetworkObjectSubNetworkSubNetworkV4Stats.Stats.Bit_rate,
		},
	}
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4StatsStats(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4StatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4Stats(d *schema.ResourceData) edpt.DdosNetworkObjectSubNetworkSubNetworkV4Stats {
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4Stats

	ret.Stats = getObjectDdosNetworkObjectSubNetworkSubNetworkV4StatsStats(d.Get("stats").([]interface{}))

	ret.SubnetIpAddr = d.Get("subnet_ip_addr").(string)

	ret.ObjectName = d.Get("object_name").(string)
	return ret
}
