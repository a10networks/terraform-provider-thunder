package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayWsFederationStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_ws_federation_stats`: Statistics for the object ws-federation\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayWsFederationStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify WS-Federation authentication relay name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"success": {
							Type: schema.TypeInt, Optional: true, Description: "Success",
						},
						"failure": {
							Type: schema.TypeInt, Optional: true, Description: "Failure",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelayWsFederationStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayWsFederationStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayWsFederationStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayWsFederationStatsStats := setObjectAamAuthenticationRelayWsFederationStatsStats(res)
		d.Set("stats", AamAuthenticationRelayWsFederationStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayWsFederationStatsStats(ret edpt.DataAamAuthenticationRelayWsFederationStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request": ret.DtAamAuthenticationRelayWsFederationStats.Stats.Request,
			"success": ret.DtAamAuthenticationRelayWsFederationStats.Stats.Success,
			"failure": ret.DtAamAuthenticationRelayWsFederationStats.Stats.Failure,
		},
	}
}

func getObjectAamAuthenticationRelayWsFederationStatsStats(d []interface{}) edpt.AamAuthenticationRelayWsFederationStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayWsFederationStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = in["request"].(int)
		ret.Success = in["success"].(int)
		ret.Failure = in["failure"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayWsFederationStats(d *schema.ResourceData) edpt.AamAuthenticationRelayWsFederationStats {
	var ret edpt.AamAuthenticationRelayWsFederationStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelayWsFederationStatsStats(d.Get("stats").([]interface{}))
	return ret
}
