package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationOauthGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_oauth_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationOauthGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_req": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_succ": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"auth_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_req": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_succ": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"relay_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationOauthGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationOauthGlobalStatsStats := setObjectAamAuthenticationOauthGlobalStatsStats(res)
		d.Set("stats", AamAuthenticationOauthGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationOauthGlobalStatsStats(ret edpt.DataAamAuthenticationOauthGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"auth_req":    ret.DtAamAuthenticationOauthGlobalStats.Stats.AuthReq,
			"auth_succ":   ret.DtAamAuthenticationOauthGlobalStats.Stats.AuthSucc,
			"auth_fail":   ret.DtAamAuthenticationOauthGlobalStats.Stats.AuthFail,
			"auth_error":  ret.DtAamAuthenticationOauthGlobalStats.Stats.AuthError,
			"relay_req":   ret.DtAamAuthenticationOauthGlobalStats.Stats.RelayReq,
			"relay_succ":  ret.DtAamAuthenticationOauthGlobalStats.Stats.RelaySucc,
			"relay_fail":  ret.DtAamAuthenticationOauthGlobalStats.Stats.RelayFail,
			"other_error": ret.DtAamAuthenticationOauthGlobalStats.Stats.OtherError,
		},
	}
}

func getObjectAamAuthenticationOauthGlobalStatsStats(d []interface{}) edpt.AamAuthenticationOauthGlobalStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationOauthGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthReq = in["auth_req"].(int)
		ret.AuthSucc = in["auth_succ"].(int)
		ret.AuthFail = in["auth_fail"].(int)
		ret.AuthError = in["auth_error"].(int)
		ret.RelayReq = in["relay_req"].(int)
		ret.RelaySucc = in["relay_succ"].(int)
		ret.RelayFail = in["relay_fail"].(int)
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationOauthGlobalStats(d *schema.ResourceData) edpt.AamAuthenticationOauthGlobalStats {
	var ret edpt.AamAuthenticationOauthGlobalStats

	ret.Stats = getObjectAamAuthenticationOauthGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
