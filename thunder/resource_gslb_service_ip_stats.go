package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_ip_stats`: Statistics for the object service-ip\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceIpStatsRead,

		Schema: map[string]*schema.Schema{
			"node_name": {
				Type: schema.TypeString, Required: true, Description: "Service-IP Name",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"port_proto": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"active": {
										Type: schema.TypeInt, Optional: true, Description: "Active Servers",
									},
									"current": {
										Type: schema.TypeInt, Optional: true, Description: "Current Connections",
									},
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of times the service IP has been selected",
						},
						"recent": {
							Type: schema.TypeInt, Optional: true, Description: "Recent hits",
						},
					},
				},
			},
		},
	}
}

func resourceGslbServiceIpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceIpStatsPortList := setSliceGslbServiceIpStatsPortList(res)
		d.Set("port_list", GslbServiceIpStatsPortList)
		GslbServiceIpStatsStats := setObjectGslbServiceIpStatsStats(res)
		d.Set("stats", GslbServiceIpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceGslbServiceIpStatsPortList(d edpt.DataGslbServiceIpStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtGslbServiceIpStats.PortList {
		in := make(map[string]interface{})
		in["port_num"] = item.PortNum
		in["port_proto"] = item.PortProto
		in["stats"] = setObjectGslbServiceIpStatsPortListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectGslbServiceIpStatsPortListStats(d edpt.GslbServiceIpStatsPortListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["active"] = d.Active

	in["current"] = d.Current
	result = append(result, in)
	return result
}

func setObjectGslbServiceIpStatsStats(ret edpt.DataGslbServiceIpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits":   ret.DtGslbServiceIpStats.Stats.Hits,
			"recent": ret.DtGslbServiceIpStats.Stats.Recent,
		},
	}
}

func getSliceGslbServiceIpStatsPortList(d []interface{}) []edpt.GslbServiceIpStatsPortList {

	count1 := len(d)
	ret := make([]edpt.GslbServiceIpStatsPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceIpStatsPortList
		oi.PortNum = in["port_num"].(int)
		oi.PortProto = in["port_proto"].(string)
		oi.Stats = getObjectGslbServiceIpStatsPortListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGslbServiceIpStatsPortListStats(d []interface{}) edpt.GslbServiceIpStatsPortListStats {

	count1 := len(d)
	var ret edpt.GslbServiceIpStatsPortListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Active = in["active"].(int)
		ret.Current = in["current"].(int)
	}
	return ret
}

func getObjectGslbServiceIpStatsStats(d []interface{}) edpt.GslbServiceIpStatsStats {

	count1 := len(d)
	var ret edpt.GslbServiceIpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
		ret.Recent = in["recent"].(int)
	}
	return ret
}

func dataToEndpointGslbServiceIpStats(d *schema.ResourceData) edpt.GslbServiceIpStats {
	var ret edpt.GslbServiceIpStats

	ret.NodeName = d.Get("node_name").(string)

	ret.PortList = getSliceGslbServiceIpStatsPortList(d.Get("port_list").([]interface{}))

	ret.Stats = getObjectGslbServiceIpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
