package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerLdapInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_ldap_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerLdapInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify LDAP authentication server name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_bind_success": {
							Type: schema.TypeInt, Optional: true, Description: "Admin Bind Success",
						},
						"admin_bind_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Admin Bind Failure",
						},
						"bind_success": {
							Type: schema.TypeInt, Optional: true, Description: "User Bind Success",
						},
						"bind_failure": {
							Type: schema.TypeInt, Optional: true, Description: "User Bind Failure",
						},
						"search_success": {
							Type: schema.TypeInt, Optional: true, Description: "Search Success",
						},
						"search_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Search Failure",
						},
						"authorize_success": {
							Type: schema.TypeInt, Optional: true, Description: "Authorization Success",
						},
						"authorize_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Authorization Failure",
						},
						"timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Timeout",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Other Error",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"ssl_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TLS/SSL Session Created",
						},
						"ssl_session_failure": {
							Type: schema.TypeInt, Optional: true, Description: "TLS/SSL Session Failure",
						},
						"pw_expiry": {
							Type: schema.TypeInt, Optional: true, Description: "Password expiry",
						},
						"pw_change_success": {
							Type: schema.TypeInt, Optional: true, Description: "Password change success",
						},
						"pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Password change failure",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerLdapInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerLdapInstanceStatsStats := setObjectAamAuthenticationServerLdapInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationServerLdapInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerLdapInstanceStatsStats(ret edpt.DataAamAuthenticationServerLdapInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"admin_bind_success":  ret.DtAamAuthenticationServerLdapInstanceStats.Stats.AdminBindSuccess,
			"admin_bind_failure":  ret.DtAamAuthenticationServerLdapInstanceStats.Stats.AdminBindFailure,
			"bind_success":        ret.DtAamAuthenticationServerLdapInstanceStats.Stats.BindSuccess,
			"bind_failure":        ret.DtAamAuthenticationServerLdapInstanceStats.Stats.BindFailure,
			"search_success":      ret.DtAamAuthenticationServerLdapInstanceStats.Stats.SearchSuccess,
			"search_failure":      ret.DtAamAuthenticationServerLdapInstanceStats.Stats.SearchFailure,
			"authorize_success":   ret.DtAamAuthenticationServerLdapInstanceStats.Stats.AuthorizeSuccess,
			"authorize_failure":   ret.DtAamAuthenticationServerLdapInstanceStats.Stats.AuthorizeFailure,
			"timeout_error":       ret.DtAamAuthenticationServerLdapInstanceStats.Stats.TimeoutError,
			"other_error":         ret.DtAamAuthenticationServerLdapInstanceStats.Stats.OtherError,
			"request":             ret.DtAamAuthenticationServerLdapInstanceStats.Stats.Request,
			"ssl_session_created": ret.DtAamAuthenticationServerLdapInstanceStats.Stats.SslSessionCreated,
			"ssl_session_failure": ret.DtAamAuthenticationServerLdapInstanceStats.Stats.SslSessionFailure,
			"pw_expiry":           ret.DtAamAuthenticationServerLdapInstanceStats.Stats.Pw_expiry,
			"pw_change_success":   ret.DtAamAuthenticationServerLdapInstanceStats.Stats.Pw_change_success,
			"pw_change_failure":   ret.DtAamAuthenticationServerLdapInstanceStats.Stats.Pw_change_failure,
		},
	}
}

func getObjectAamAuthenticationServerLdapInstanceStatsStats(d []interface{}) edpt.AamAuthenticationServerLdapInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AdminBindSuccess = in["admin_bind_success"].(int)
		ret.AdminBindFailure = in["admin_bind_failure"].(int)
		ret.BindSuccess = in["bind_success"].(int)
		ret.BindFailure = in["bind_failure"].(int)
		ret.SearchSuccess = in["search_success"].(int)
		ret.SearchFailure = in["search_failure"].(int)
		ret.AuthorizeSuccess = in["authorize_success"].(int)
		ret.AuthorizeFailure = in["authorize_failure"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.Request = in["request"].(int)
		ret.SslSessionCreated = in["ssl_session_created"].(int)
		ret.SslSessionFailure = in["ssl_session_failure"].(int)
		ret.Pw_expiry = in["pw_expiry"].(int)
		ret.Pw_change_success = in["pw_change_success"].(int)
		ret.Pw_change_failure = in["pw_change_failure"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerLdapInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationServerLdapInstanceStats {
	var ret edpt.AamAuthenticationServerLdapInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationServerLdapInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
