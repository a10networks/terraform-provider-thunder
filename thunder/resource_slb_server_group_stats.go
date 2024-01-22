package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServerGroupStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_server_group_stats`: Statistics for the object server-group\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServerGroupStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "server-group name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Current established connections",
						},
					},
				},
			},
		},
	}
}

func resourceSlbServerGroupStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerGroupStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerGroupStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServerGroupStatsStats := setObjectSlbServerGroupStatsStats(res)
		d.Set("stats", SlbServerGroupStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbServerGroupStatsStats(ret edpt.DataSlbServerGroupStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dummy_conn": ret.DtSlbServerGroupStats.Stats.DummyConn,
		},
	}
}

func getObjectSlbServerGroupStatsStats(d []interface{}) edpt.SlbServerGroupStatsStats {

	count1 := len(d)
	var ret edpt.SlbServerGroupStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DummyConn = in["dummy_conn"].(int)
	}
	return ret
}

func dataToEndpointSlbServerGroupStats(d *schema.ResourceData) edpt.SlbServerGroupStats {
	var ret edpt.SlbServerGroupStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbServerGroupStatsStats(d.Get("stats").([]interface{}))
	return ret
}
