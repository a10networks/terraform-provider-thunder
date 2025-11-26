package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIpServiceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_service_ip_service_stats`: Statistics for the object service\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbServiceIpServiceStatsRead,

		Schema: map[string]*schema.Schema{
			"label": {
				Type: schema.TypeString, Required: true, Description: "Service Label",
			},
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

func resourceGslbServiceIpServiceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpServiceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIpServiceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbServiceIpServiceStatsStats := setObjectGslbServiceIpServiceStatsStats(res)
		d.Set("stats", GslbServiceIpServiceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbServiceIpServiceStatsStats(ret edpt.DataGslbServiceIpServiceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"active":  ret.DtGslbServiceIpServiceStats.Stats.Active,
			"current": ret.DtGslbServiceIpServiceStats.Stats.Current,
		},
	}
}

func getObjectGslbServiceIpServiceStatsStats(d []interface{}) edpt.GslbServiceIpServiceStatsStats {

	count1 := len(d)
	var ret edpt.GslbServiceIpServiceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Active = in["active"].(int)
		ret.Current = in["current"].(int)
	}
	return ret
}

func dataToEndpointGslbServiceIpServiceStats(d *schema.ResourceData) edpt.GslbServiceIpServiceStats {
	var ret edpt.GslbServiceIpServiceStats

	ret.Label = d.Get("label").(string)

	ret.PortNum = d.Get("port_num").(int)

	ret.PortProto = d.Get("port_proto").(string)

	ret.Stats = getObjectGslbServiceIpServiceStatsStats(d.Get("stats").([]interface{}))

	ret.NodeName = d.Get("node_name").(string)
	return ret
}
