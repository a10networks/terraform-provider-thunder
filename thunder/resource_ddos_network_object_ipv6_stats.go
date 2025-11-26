package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectIpv6Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_ipv6_stats`: Statistics for the object ipv6\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectIpv6StatsRead,

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
				Type: schema.TypeString, Required: true, Description: "IPV6 Subnet, supported prefix range is from 40 to 64",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}

func resourceDdosNetworkObjectIpv6StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectIpv6StatsStats := setObjectDdosNetworkObjectIpv6StatsStats(res)
		d.Set("stats", DdosNetworkObjectIpv6StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosNetworkObjectIpv6StatsStats(ret edpt.DataDdosNetworkObjectIpv6Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_rate": ret.DtDdosNetworkObjectIpv6Stats.Stats.Packet_rate,
			"bit_rate":    ret.DtDdosNetworkObjectIpv6Stats.Stats.Bit_rate,
		},
	}
}

func getObjectDdosNetworkObjectIpv6StatsStats(d []interface{}) edpt.DdosNetworkObjectIpv6StatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectIpv6Stats(d *schema.ResourceData) edpt.DdosNetworkObjectIpv6Stats {
	var ret edpt.DdosNetworkObjectIpv6Stats

	ret.Stats = getObjectDdosNetworkObjectIpv6StatsStats(d.Get("stats").([]interface{}))

	ret.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)

	ret.ObjectName = d.Get("object_name").(string)
	return ret
}
