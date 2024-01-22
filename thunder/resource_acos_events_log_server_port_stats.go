package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogServerPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_acos_events_log_server_port_stats`: Statistics for the object port\n\n__PLACEHOLDER__",
		ReadContext: resourceAcosEventsLogServerPortStatsRead,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceAcosEventsLogServerPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogServerPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogServerPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AcosEventsLogServerPortStatsStats := setObjectAcosEventsLogServerPortStatsStats(res)
		d.Set("stats", AcosEventsLogServerPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAcosEventsLogServerPortStatsStats(ret edpt.DataAcosEventsLogServerPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectAcosEventsLogServerPortStatsStats(d []interface{}) edpt.AcosEventsLogServerPortStatsStats {

	var ret edpt.AcosEventsLogServerPortStatsStats
	return ret
}

func dataToEndpointAcosEventsLogServerPortStats(d *schema.ResourceData) edpt.AcosEventsLogServerPortStats {
	var ret edpt.AcosEventsLogServerPortStats

	ret.PortNumber = d.Get("port_number").(int)

	ret.Protocol = d.Get("protocol").(string)

	ret.Stats = getObjectAcosEventsLogServerPortStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
