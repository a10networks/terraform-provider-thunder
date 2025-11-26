package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_network_object_stats`: Statistics for the object network-object\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosNetworkObjectStatsRead,

		Schema: map[string]*schema.Schema{
			"ip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ip_addr": {
							Type: schema.TypeString, Required: true, Description: "IP Subnet, supported prefix range is from 8 to 32",
						},
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
				},
			},
			"ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Required: true, Description: "IPV6 Subnet, supported prefix range is from 40 to 64",
						},
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
				},
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry Learned",
						},
						"subnet_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry Aged",
						},
						"subnet_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry Create Failures",
						},
						"ip_learned": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry Learned",
						},
						"ip_aged": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry Aged",
						},
						"ip_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry Create Failures",
						},
						"service_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry Learned",
						},
						"service_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry Aged",
						},
						"service_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry Create Failures",
						},
						"packet_rate": {
							Type: schema.TypeInt, Optional: true, Description: "PPS",
						},
						"bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "B(bits)PS",
						},
						"topk_allocate_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Topk Allocate Failures",
						},
						"sport_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Source Port Entry Learned",
						},
						"sport_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Source Port Entry Aged",
						},
						"sport_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source Port Entry Create Failures",
						},
						"agent_group_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Agent Group Entry Learned",
						},
						"agent_group_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Agent Group Entry Aged",
						},
						"agent_group_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Agent Group Entry Create Failures",
						},
						"duplicate_sample_pkt_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Duplicate Sample Packet Received",
						},
					},
				},
			},
		},
	}
}

func resourceDdosNetworkObjectStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosNetworkObjectStatsIpList := setSliceDdosNetworkObjectStatsIpList(res)
		d.Set("ip_list", DdosNetworkObjectStatsIpList)
		DdosNetworkObjectStatsIpv6List := setSliceDdosNetworkObjectStatsIpv6List(res)
		d.Set("ipv6_list", DdosNetworkObjectStatsIpv6List)
		DdosNetworkObjectStatsStats := setObjectDdosNetworkObjectStatsStats(res)
		d.Set("stats", DdosNetworkObjectStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceDdosNetworkObjectStatsIpList(d edpt.DataDdosNetworkObjectStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosNetworkObjectStats.IpList {
		in := make(map[string]interface{})
		in["subnet_ip_addr"] = item.SubnetIpAddr
		in["stats"] = setObjectDdosNetworkObjectStatsIpListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectDdosNetworkObjectStatsIpListStats(d edpt.DdosNetworkObjectStatsIpListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["packet_rate"] = d.Packet_rate

	in["bit_rate"] = d.Bit_rate
	result = append(result, in)
	return result
}

func setSliceDdosNetworkObjectStatsIpv6List(d edpt.DataDdosNetworkObjectStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtDdosNetworkObjectStats.Ipv6List {
		in := make(map[string]interface{})
		in["subnet_ipv6_addr"] = item.SubnetIpv6Addr
		in["stats"] = setObjectDdosNetworkObjectStatsIpv6ListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectDdosNetworkObjectStatsIpv6ListStats(d edpt.DdosNetworkObjectStatsIpv6ListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["packet_rate"] = d.Packet_rate

	in["bit_rate"] = d.Bit_rate
	result = append(result, in)
	return result
}

func setObjectDdosNetworkObjectStatsStats(ret edpt.DataDdosNetworkObjectStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"subnet_learned":           ret.DtDdosNetworkObjectStats.Stats.Subnet_learned,
			"subnet_aged":              ret.DtDdosNetworkObjectStats.Stats.Subnet_aged,
			"subnet_create_fail":       ret.DtDdosNetworkObjectStats.Stats.Subnet_create_fail,
			"ip_learned":               ret.DtDdosNetworkObjectStats.Stats.Ip_learned,
			"ip_aged":                  ret.DtDdosNetworkObjectStats.Stats.Ip_aged,
			"ip_create_fail":           ret.DtDdosNetworkObjectStats.Stats.Ip_create_fail,
			"service_learned":          ret.DtDdosNetworkObjectStats.Stats.Service_learned,
			"service_aged":             ret.DtDdosNetworkObjectStats.Stats.Service_aged,
			"service_create_fail":      ret.DtDdosNetworkObjectStats.Stats.Service_create_fail,
			"packet_rate":              ret.DtDdosNetworkObjectStats.Stats.Packet_rate,
			"bit_rate":                 ret.DtDdosNetworkObjectStats.Stats.Bit_rate,
			"topk_allocate_fail":       ret.DtDdosNetworkObjectStats.Stats.Topk_allocate_fail,
			"sport_learned":            ret.DtDdosNetworkObjectStats.Stats.Sport_learned,
			"sport_aged":               ret.DtDdosNetworkObjectStats.Stats.Sport_aged,
			"sport_create_fail":        ret.DtDdosNetworkObjectStats.Stats.Sport_create_fail,
			"agent_group_learned":      ret.DtDdosNetworkObjectStats.Stats.Agent_group_learned,
			"agent_group_aged":         ret.DtDdosNetworkObjectStats.Stats.Agent_group_aged,
			"agent_group_create_fail":  ret.DtDdosNetworkObjectStats.Stats.Agent_group_create_fail,
			"duplicate_sample_pkt_rcv": ret.DtDdosNetworkObjectStats.Stats.Duplicate_sample_pkt_rcv,
		},
	}
}

func getSliceDdosNetworkObjectStatsIpList(d []interface{}) []edpt.DdosNetworkObjectStatsIpList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectStatsIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectStatsIpList
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		oi.Stats = getObjectDdosNetworkObjectStatsIpListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectStatsIpListStats(d []interface{}) edpt.DdosNetworkObjectStatsIpListStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectStatsIpListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectStatsIpv6List(d []interface{}) []edpt.DdosNetworkObjectStatsIpv6List {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectStatsIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectStatsIpv6List
		oi.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		oi.Stats = getObjectDdosNetworkObjectStatsIpv6ListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectStatsIpv6ListStats(d []interface{}) edpt.DdosNetworkObjectStatsIpv6ListStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectStatsIpv6ListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectStatsStats(d []interface{}) edpt.DdosNetworkObjectStatsStats {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Subnet_learned = in["subnet_learned"].(int)
		ret.Subnet_aged = in["subnet_aged"].(int)
		ret.Subnet_create_fail = in["subnet_create_fail"].(int)
		ret.Ip_learned = in["ip_learned"].(int)
		ret.Ip_aged = in["ip_aged"].(int)
		ret.Ip_create_fail = in["ip_create_fail"].(int)
		ret.Service_learned = in["service_learned"].(int)
		ret.Service_aged = in["service_aged"].(int)
		ret.Service_create_fail = in["service_create_fail"].(int)
		ret.Packet_rate = in["packet_rate"].(int)
		ret.Bit_rate = in["bit_rate"].(int)
		ret.Topk_allocate_fail = in["topk_allocate_fail"].(int)
		ret.Sport_learned = in["sport_learned"].(int)
		ret.Sport_aged = in["sport_aged"].(int)
		ret.Sport_create_fail = in["sport_create_fail"].(int)
		ret.Agent_group_learned = in["agent_group_learned"].(int)
		ret.Agent_group_aged = in["agent_group_aged"].(int)
		ret.Agent_group_create_fail = in["agent_group_create_fail"].(int)
		ret.Duplicate_sample_pkt_rcv = in["duplicate_sample_pkt_rcv"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectStats(d *schema.ResourceData) edpt.DdosNetworkObjectStats {
	var ret edpt.DdosNetworkObjectStats

	ret.IpList = getSliceDdosNetworkObjectStatsIpList(d.Get("ip_list").([]interface{}))

	ret.Ipv6List = getSliceDdosNetworkObjectStatsIpv6List(d.Get("ipv6_list").([]interface{}))

	ret.ObjectName = d.Get("object_name").(string)

	ret.Stats = getObjectDdosNetworkObjectStatsStats(d.Get("stats").([]interface{}))
	return ret
}
