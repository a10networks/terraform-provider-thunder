package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"requests": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication Request",
						},
						"responses": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication Response",
						},
						"misses": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication Request Missed",
						},
						"ocsp_stapling_requests_to_a10authd": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Request",
						},
						"ocsp_stapling_responses_from_a10authd": {
							Type: schema.TypeInt, Optional: true, Description: "Total OCSP Stapling Response",
						},
						"opened_socket": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Socket Opened",
						},
						"open_socket_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Open Socket Failed",
						},
						"connect": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Connection",
						},
						"connect_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Connect Failed",
						},
						"created_timer": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Timer Created",
						},
						"create_timer_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Timer Creation Failed",
						},
						"total_request": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request Received by A10 Auth Service",
						},
						"get_socket_option_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM Get Socket Option Failed",
						},
						"aflex_authz_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization success number in aFleX",
						},
						"aflex_authz_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization failure number in aFleX",
						},
						"authn_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication success number",
						},
						"authn_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication failure number",
						},
						"authz_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization success number",
						},
						"authz_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization failure number",
						},
						"active_session": {
							Type: schema.TypeInt, Optional: true, Description: "Total Active Auth-Sessions",
						},
						"active_user": {
							Type: schema.TypeInt, Optional: true, Description: "Total Active Users",
						},
						"dns_resolve_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total AAM DNS resolve failed",
						},
						"domain_wlist_match": {
							Type: schema.TypeInt, Optional: true, Description: "Total DOMAIN WHITELIST match number",
						},
						"domain_wlist_unmatch": {
							Type: schema.TypeInt, Optional: true, Description: "Total DOMAIN WHITELIST unmatch number",
						},
						"auth_ctx_num": {
							Type: schema.TypeInt, Optional: true, Description: "Total Auth Contexts",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationGlobalStatsStats := setObjectAamAuthenticationGlobalStatsStats(res)
		d.Set("stats", AamAuthenticationGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationGlobalStatsStats(ret edpt.DataAamAuthenticationGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"requests":                              ret.DtAamAuthenticationGlobalStats.Stats.Requests,
			"responses":                             ret.DtAamAuthenticationGlobalStats.Stats.Responses,
			"misses":                                ret.DtAamAuthenticationGlobalStats.Stats.Misses,
			"ocsp_stapling_requests_to_a10authd":    ret.DtAamAuthenticationGlobalStats.Stats.OcspStaplingRequestsToA10authd,
			"ocsp_stapling_responses_from_a10authd": ret.DtAamAuthenticationGlobalStats.Stats.OcspStaplingResponsesFromA10authd,
			"opened_socket":                         ret.DtAamAuthenticationGlobalStats.Stats.OpenedSocket,
			"open_socket_failed":                    ret.DtAamAuthenticationGlobalStats.Stats.OpenSocketFailed,
			"connect":                               ret.DtAamAuthenticationGlobalStats.Stats.Connect,
			"connect_failed":                        ret.DtAamAuthenticationGlobalStats.Stats.ConnectFailed,
			"created_timer":                         ret.DtAamAuthenticationGlobalStats.Stats.CreatedTimer,
			"create_timer_failed":                   ret.DtAamAuthenticationGlobalStats.Stats.CreateTimerFailed,
			"total_request":                         ret.DtAamAuthenticationGlobalStats.Stats.TotalRequest,
			"get_socket_option_failed":              ret.DtAamAuthenticationGlobalStats.Stats.GetSocketOptionFailed,
			"aflex_authz_succ":                      ret.DtAamAuthenticationGlobalStats.Stats.AflexAuthzSucc,
			"aflex_authz_fail":                      ret.DtAamAuthenticationGlobalStats.Stats.AflexAuthzFail,
			"authn_success":                         ret.DtAamAuthenticationGlobalStats.Stats.AuthnSuccess,
			"authn_failure":                         ret.DtAamAuthenticationGlobalStats.Stats.AuthnFailure,
			"authz_success":                         ret.DtAamAuthenticationGlobalStats.Stats.AuthzSuccess,
			"authz_failure":                         ret.DtAamAuthenticationGlobalStats.Stats.AuthzFailure,
			"active_session":                        ret.DtAamAuthenticationGlobalStats.Stats.ActiveSession,
			"active_user":                           ret.DtAamAuthenticationGlobalStats.Stats.ActiveUser,
			"dns_resolve_failed":                    ret.DtAamAuthenticationGlobalStats.Stats.DnsResolveFailed,
			"domain_wlist_match":                    ret.DtAamAuthenticationGlobalStats.Stats.DomainWlistMatch,
			"domain_wlist_unmatch":                  ret.DtAamAuthenticationGlobalStats.Stats.DomainWlistUnmatch,
			"auth_ctx_num":                          ret.DtAamAuthenticationGlobalStats.Stats.Auth_ctx_num,
		},
	}
}

func getObjectAamAuthenticationGlobalStatsStats(d []interface{}) edpt.AamAuthenticationGlobalStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Requests = in["requests"].(int)
		ret.Responses = in["responses"].(int)
		ret.Misses = in["misses"].(int)
		ret.OcspStaplingRequestsToA10authd = in["ocsp_stapling_requests_to_a10authd"].(int)
		ret.OcspStaplingResponsesFromA10authd = in["ocsp_stapling_responses_from_a10authd"].(int)
		ret.OpenedSocket = in["opened_socket"].(int)
		ret.OpenSocketFailed = in["open_socket_failed"].(int)
		ret.Connect = in["connect"].(int)
		ret.ConnectFailed = in["connect_failed"].(int)
		ret.CreatedTimer = in["created_timer"].(int)
		ret.CreateTimerFailed = in["create_timer_failed"].(int)
		ret.TotalRequest = in["total_request"].(int)
		ret.GetSocketOptionFailed = in["get_socket_option_failed"].(int)
		ret.AflexAuthzSucc = in["aflex_authz_succ"].(int)
		ret.AflexAuthzFail = in["aflex_authz_fail"].(int)
		ret.AuthnSuccess = in["authn_success"].(int)
		ret.AuthnFailure = in["authn_failure"].(int)
		ret.AuthzSuccess = in["authz_success"].(int)
		ret.AuthzFailure = in["authz_failure"].(int)
		ret.ActiveSession = in["active_session"].(int)
		ret.ActiveUser = in["active_user"].(int)
		ret.DnsResolveFailed = in["dns_resolve_failed"].(int)
		ret.DomainWlistMatch = in["domain_wlist_match"].(int)
		ret.DomainWlistUnmatch = in["domain_wlist_unmatch"].(int)
		ret.Auth_ctx_num = in["auth_ctx_num"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationGlobalStats(d *schema.ResourceData) edpt.AamAuthenticationGlobalStats {
	var ret edpt.AamAuthenticationGlobalStats

	ret.Stats = getObjectAamAuthenticationGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
