package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayOauthStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_oauth_stats`: Statistics for the object oauth\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayOauthStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify oauth authentication relay name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relay_req": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_succ": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelayOauthStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayOauthStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayOauthStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayOauthStatsStats := setObjectAamAuthenticationRelayOauthStatsStats(res)
		d.Set("stats", AamAuthenticationRelayOauthStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayOauthStatsStats(ret edpt.DataAamAuthenticationRelayOauthStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"relay_req":  ret.DtAamAuthenticationRelayOauthStats.Stats.RelayReq,
			"relay_succ": ret.DtAamAuthenticationRelayOauthStats.Stats.RelaySucc,
			"relay_fail": ret.DtAamAuthenticationRelayOauthStats.Stats.RelayFail,
		},
	}
}

func getObjectAamAuthenticationRelayOauthStatsStats(d []interface{}) edpt.AamAuthenticationRelayOauthStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayOauthStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RelayReq = in["relay_req"].(int)
		ret.RelaySucc = in["relay_succ"].(int)
		ret.RelayFail = in["relay_fail"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayOauthStats(d *schema.ResourceData) edpt.AamAuthenticationRelayOauthStats {
	var ret edpt.AamAuthenticationRelayOauthStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelayOauthStatsStats(d.Get("stats").([]interface{}))
	return ret
}
