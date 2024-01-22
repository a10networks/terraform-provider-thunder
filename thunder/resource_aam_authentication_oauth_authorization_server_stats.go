package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationOauthAuthorizationServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_oauth_authorization_server_stats`: Statistics for the object authorization-server\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationOauthAuthorizationServerStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify authorization server object name",
			},
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
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationOauthAuthorizationServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationOauthAuthorizationServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationOauthAuthorizationServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationOauthAuthorizationServerStatsStats := setObjectAamAuthenticationOauthAuthorizationServerStatsStats(res)
		d.Set("stats", AamAuthenticationOauthAuthorizationServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationOauthAuthorizationServerStatsStats(ret edpt.DataAamAuthenticationOauthAuthorizationServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"auth_req":    ret.DtAamAuthenticationOauthAuthorizationServerStats.Stats.AuthReq,
			"auth_succ":   ret.DtAamAuthenticationOauthAuthorizationServerStats.Stats.AuthSucc,
			"auth_fail":   ret.DtAamAuthenticationOauthAuthorizationServerStats.Stats.AuthFail,
			"auth_error":  ret.DtAamAuthenticationOauthAuthorizationServerStats.Stats.AuthError,
			"other_error": ret.DtAamAuthenticationOauthAuthorizationServerStats.Stats.OtherError,
		},
	}
}

func getObjectAamAuthenticationOauthAuthorizationServerStatsStats(d []interface{}) edpt.AamAuthenticationOauthAuthorizationServerStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationOauthAuthorizationServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthReq = in["auth_req"].(int)
		ret.AuthSucc = in["auth_succ"].(int)
		ret.AuthFail = in["auth_fail"].(int)
		ret.AuthError = in["auth_error"].(int)
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationOauthAuthorizationServerStats(d *schema.ResourceData) edpt.AamAuthenticationOauthAuthorizationServerStats {
	var ret edpt.AamAuthenticationOauthAuthorizationServerStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationOauthAuthorizationServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
