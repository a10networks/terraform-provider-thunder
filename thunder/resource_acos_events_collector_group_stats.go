package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsCollectorGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_acos_events_collector_group_stats`: Statistics for the object collector-group\n\n__PLACEHOLDER__",
		ReadContext: resourceAcosEventsCollectorGroupStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify log server group name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msgs_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Number of log messages sent",
						},
					},
				},
			},
		},
	}
}

func resourceAcosEventsCollectorGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsCollectorGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsCollectorGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AcosEventsCollectorGroupStatsStats := setObjectAcosEventsCollectorGroupStatsStats(res)
		d.Set("stats", AcosEventsCollectorGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAcosEventsCollectorGroupStatsStats(ret edpt.DataAcosEventsCollectorGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"msgs_sent": ret.DtAcosEventsCollectorGroupStats.Stats.Msgs_sent,
		},
	}
}

func getObjectAcosEventsCollectorGroupStatsStats(d []interface{}) edpt.AcosEventsCollectorGroupStatsStats {

	count1 := len(d)
	var ret edpt.AcosEventsCollectorGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Msgs_sent = in["msgs_sent"].(int)
	}
	return ret
}

func dataToEndpointAcosEventsCollectorGroupStats(d *schema.ResourceData) edpt.AcosEventsCollectorGroupStats {
	var ret edpt.AcosEventsCollectorGroupStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAcosEventsCollectorGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
