package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_acos_events_log_server_stats`: Statistics for the object log-server\n\n__PLACEHOLDER__",
		ReadContext: resourceAcosEventsLogServerStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Name",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceAcosEventsLogServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AcosEventsLogServerStatsPortList := setSliceAcosEventsLogServerStatsPortList(res)
		d.Set("port_list", AcosEventsLogServerStatsPortList)
		AcosEventsLogServerStatsStats := setObjectAcosEventsLogServerStatsStats(res)
		d.Set("stats", AcosEventsLogServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAcosEventsLogServerStatsPortList(d edpt.DataAcosEventsLogServerStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAcosEventsLogServerStats.PortList {
		in := make(map[string]interface{})
		in["port_number"] = item.PortNumber
		in["protocol"] = item.Protocol
		in["stats"] = setObjectAcosEventsLogServerStatsPortListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAcosEventsLogServerStatsPortListStats(d edpt.AcosEventsLogServerStatsPortListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setObjectAcosEventsLogServerStatsStats(ret edpt.DataAcosEventsLogServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getSliceAcosEventsLogServerStatsPortList(d []interface{}) []edpt.AcosEventsLogServerStatsPortList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLogServerStatsPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLogServerStatsPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Stats = getObjectAcosEventsLogServerStatsPortListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAcosEventsLogServerStatsPortListStats(d []interface{}) edpt.AcosEventsLogServerStatsPortListStats {

	var ret edpt.AcosEventsLogServerStatsPortListStats
	return ret
}

func getObjectAcosEventsLogServerStatsStats(d []interface{}) edpt.AcosEventsLogServerStatsStats {

	var ret edpt.AcosEventsLogServerStatsStats
	return ret
}

func dataToEndpointAcosEventsLogServerStats(d *schema.ResourceData) edpt.AcosEventsLogServerStats {
	var ret edpt.AcosEventsLogServerStats

	ret.Name = d.Get("name").(string)

	ret.PortList = getSliceAcosEventsLogServerStatsPortList(d.Get("port_list").([]interface{}))

	ret.Stats = getObjectAcosEventsLogServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
