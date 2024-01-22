package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_ip_port_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceIpPortStatsRead,

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
			"node_name": {
				Type: schema.TypeString, Required: true, Description: "NodeName",
			},
		},
	}
}

func resourceGslbServiceIpPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceIpPortStatsStats := setObjectGslbServiceIpPortStatsStats(res)
		d.Set("stats", GslbServiceIpPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServiceIpPortStatsStats(ret edpt.DataGslbServiceIpPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"active":  ret.DtGslbServiceIpPortStats.Stats.Active,
			"current": ret.DtGslbServiceIpPortStats.Stats.Current,
		},
	}
}

func getObjectGslbServiceIpPortStatsStats(d []interface{}) edpt.GslbServiceIpPortStatsStats {

	count1 := len(d)
	var ret edpt.GslbServiceIpPortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Active = in["active"].(int)
		ret.Current = in["current"].(int)
	}
	return ret
}

func dataToEndpointGslbServiceIpPortStats(d *schema.ResourceData) edpt.GslbServiceIpPortStats {
	var ret edpt.GslbServiceIpPortStats

	ret.PortNum = d.Get("port_num").(int)

	ret.PortProto = d.Get("port_proto").(string)

	ret.Stats = getObjectGslbServiceIpPortStatsStats(d.Get("stats").([]interface{}))

	ret.NodeName = d.Get("node_name").(string)
	return ret
}
