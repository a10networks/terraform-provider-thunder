package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSubNetworkSubNetworkV6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_sub_network_sub_network_v6_stats`: Statistics for the object sub-network-v6\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectSubNetworkSubNetworkV6StatsRead,

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
			"subnet_ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "IPv6 Subnet/host, supported prefix range is from 56 to 64",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}

func resourceDdosNetworkObjectSubNetworkSubNetworkV6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectSubNetworkSubNetworkV6StatsStats := setObjectDdosNetworkObjectSubNetworkSubNetworkV6StatsStats(res)
		d.Set("stats", DdosNetworkObjectSubNetworkSubNetworkV6StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectSubNetworkSubNetworkV6StatsStats(ret edpt.DataDdosNetworkObjectSubNetworkSubNetworkV6Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_rate": ret.DtDdosNetworkObjectSubNetworkSubNetworkV6Stats.Stats.Packet_rate,
			"bit_rate":    ret.DtDdosNetworkObjectSubNetworkSubNetworkV6Stats.Stats.Bit_rate,
		},
	}
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV6StatsStats(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV6StatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6Stats(d *schema.ResourceData) edpt.DdosNetworkObjectSubNetworkSubNetworkV6Stats {
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6Stats

	ret.Stats = getObjectDdosNetworkObjectSubNetworkSubNetworkV6StatsStats(d.Get("stats").([]interface{}))

	ret.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)

	ret.ObjectName = d.Get("object_name").(string)
	return ret
}
