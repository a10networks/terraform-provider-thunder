package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerLdapStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_ldap_stats`: Statistics for the object ldap\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerLdapStatsRead,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_bind_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Admin Bind Success",
						},
						"admin_bind_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Admin Bind Failure",
						},
						"bind_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total User Bind Success",
						},
						"bind_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total User Bind Failure",
						},
						"search_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Search Success",
						},
						"search_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Search Failure",
						},
						"authorize_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization Success",
						},
						"authorize_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization Failure",
						},
						"timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Timeout",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Error",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request",
						},
						"request_normal": {
							Type: schema.TypeInt, Optional: true, Description: "Total Normal Request",
						},
						"request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total Dropped Request",
						},
						"response_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Success Response",
						},
						"response_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Failure Response",
						},
						"response_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Error Response",
						},
						"response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Total Timeout Response",
						},
						"response_other": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Response",
						},
						"job_start_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Job Start Error",
						},
						"polling_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Polling Control Error",
						},
						"ssl_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TLS/SSL Session Created",
						},
						"ssl_session_failure": {
							Type: schema.TypeInt, Optional: true, Description: "TLS/SSL Session Failure",
						},
						"ldaps_idle_conn_num": {
							Type: schema.TypeInt, Optional: true, Description: "LDAPS Idle Connection Number",
						},
						"ldaps_inuse_conn_num": {
							Type: schema.TypeInt, Optional: true, Description: "LDAPS In-use Connection Number",
						},
						"pw_expiry": {
							Type: schema.TypeInt, Optional: true, Description: "Total Password expiry",
						},
						"pw_change_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total password change success",
						},
						"pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total password change failure",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerLdapStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerLdapStatsInstanceList := setSliceAamAuthenticationServerLdapStatsInstanceList(res)
		d.Set("instance_list", AamAuthenticationServerLdapStatsInstanceList)
		AamAuthenticationServerLdapStatsStats := setObjectAamAuthenticationServerLdapStatsStats(res)
		d.Set("stats", AamAuthenticationServerLdapStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAuthenticationServerLdapStatsInstanceList(d edpt.DataAamAuthenticationServerLdapStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAuthenticationServerLdapStats.InstanceList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectAamAuthenticationServerLdapStatsInstanceListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationServerLdapStatsInstanceListStats(d edpt.AamAuthenticationServerLdapStatsInstanceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["admin_bind_success"] = d.AdminBindSuccess

	in["admin_bind_failure"] = d.AdminBindFailure

	in["bind_success"] = d.BindSuccess

	in["bind_failure"] = d.BindFailure

	in["search_success"] = d.SearchSuccess

	in["search_failure"] = d.SearchFailure

	in["authorize_success"] = d.AuthorizeSuccess

	in["authorize_failure"] = d.AuthorizeFailure

	in["timeout_error"] = d.TimeoutError

	in["other_error"] = d.OtherError

	in["request"] = d.Request

	in["ssl_session_created"] = d.SslSessionCreated

	in["ssl_session_failure"] = d.SslSessionFailure

	in["pw_expiry"] = d.Pw_expiry

	in["pw_change_success"] = d.Pw_change_success

	in["pw_change_failure"] = d.Pw_change_failure
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationServerLdapStatsStats(ret edpt.DataAamAuthenticationServerLdapStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"admin_bind_success":    ret.DtAamAuthenticationServerLdapStats.Stats.AdminBindSuccess,
			"admin_bind_failure":    ret.DtAamAuthenticationServerLdapStats.Stats.AdminBindFailure,
			"bind_success":          ret.DtAamAuthenticationServerLdapStats.Stats.BindSuccess,
			"bind_failure":          ret.DtAamAuthenticationServerLdapStats.Stats.BindFailure,
			"search_success":        ret.DtAamAuthenticationServerLdapStats.Stats.SearchSuccess,
			"search_failure":        ret.DtAamAuthenticationServerLdapStats.Stats.SearchFailure,
			"authorize_success":     ret.DtAamAuthenticationServerLdapStats.Stats.AuthorizeSuccess,
			"authorize_failure":     ret.DtAamAuthenticationServerLdapStats.Stats.AuthorizeFailure,
			"timeout_error":         ret.DtAamAuthenticationServerLdapStats.Stats.TimeoutError,
			"other_error":           ret.DtAamAuthenticationServerLdapStats.Stats.OtherError,
			"request":               ret.DtAamAuthenticationServerLdapStats.Stats.Request,
			"request_normal":        ret.DtAamAuthenticationServerLdapStats.Stats.RequestNormal,
			"request_dropped":       ret.DtAamAuthenticationServerLdapStats.Stats.RequestDropped,
			"response_success":      ret.DtAamAuthenticationServerLdapStats.Stats.ResponseSuccess,
			"response_failure":      ret.DtAamAuthenticationServerLdapStats.Stats.ResponseFailure,
			"response_error":        ret.DtAamAuthenticationServerLdapStats.Stats.ResponseError,
			"response_timeout":      ret.DtAamAuthenticationServerLdapStats.Stats.ResponseTimeout,
			"response_other":        ret.DtAamAuthenticationServerLdapStats.Stats.ResponseOther,
			"job_start_error":       ret.DtAamAuthenticationServerLdapStats.Stats.JobStartError,
			"polling_control_error": ret.DtAamAuthenticationServerLdapStats.Stats.PollingControlError,
			"ssl_session_created":   ret.DtAamAuthenticationServerLdapStats.Stats.SslSessionCreated,
			"ssl_session_failure":   ret.DtAamAuthenticationServerLdapStats.Stats.SslSessionFailure,
			"ldaps_idle_conn_num":   ret.DtAamAuthenticationServerLdapStats.Stats.LdapsIdleConnNum,
			"ldaps_inuse_conn_num":  ret.DtAamAuthenticationServerLdapStats.Stats.LdapsInuseConnNum,
			"pw_expiry":             ret.DtAamAuthenticationServerLdapStats.Stats.PwExpiry,
			"pw_change_success":     ret.DtAamAuthenticationServerLdapStats.Stats.PwChangeSuccess,
			"pw_change_failure":     ret.DtAamAuthenticationServerLdapStats.Stats.PwChangeFailure,
		},
	}
}

func getSliceAamAuthenticationServerLdapStatsInstanceList(d []interface{}) []edpt.AamAuthenticationServerLdapStatsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapStatsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapStatsInstanceList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectAamAuthenticationServerLdapStatsInstanceListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerLdapStatsInstanceListStats(d []interface{}) edpt.AamAuthenticationServerLdapStatsInstanceListStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapStatsInstanceListStats
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

func getObjectAamAuthenticationServerLdapStatsStats(d []interface{}) edpt.AamAuthenticationServerLdapStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapStatsStats
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
		ret.RequestNormal = in["request_normal"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseSuccess = in["response_success"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.ResponseOther = in["response_other"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
		ret.SslSessionCreated = in["ssl_session_created"].(int)
		ret.SslSessionFailure = in["ssl_session_failure"].(int)
		ret.LdapsIdleConnNum = in["ldaps_idle_conn_num"].(int)
		ret.LdapsInuseConnNum = in["ldaps_inuse_conn_num"].(int)
		ret.PwExpiry = in["pw_expiry"].(int)
		ret.PwChangeSuccess = in["pw_change_success"].(int)
		ret.PwChangeFailure = in["pw_change_failure"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerLdapStats(d *schema.ResourceData) edpt.AamAuthenticationServerLdapStats {
	var ret edpt.AamAuthenticationServerLdapStats

	ret.InstanceList = getSliceAamAuthenticationServerLdapStatsInstanceList(d.Get("instance_list").([]interface{}))

	ret.Stats = getObjectAamAuthenticationServerLdapStatsStats(d.Get("stats").([]interface{}))
	return ret
}
