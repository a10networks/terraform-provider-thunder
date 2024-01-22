package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerRadiusStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_radius_stats`: Statistics for the object radius\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerRadiusStatsRead,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify RADIUS authentication server name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authen_success": {
										Type: schema.TypeInt, Optional: true, Description: "Authentication Success",
									},
									"authen_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Authentication Failure",
									},
									"authorize_success": {
										Type: schema.TypeInt, Optional: true, Description: "Authorization Success",
									},
									"authorize_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Authorization Failure",
									},
									"access_challenge": {
										Type: schema.TypeInt, Optional: true, Description: "Access-Challenge Message Receive",
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
									"accounting_request_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Accounting-Request Sent",
									},
									"accounting_success": {
										Type: schema.TypeInt, Optional: true, Description: "Accounting Success",
									},
									"accounting_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Accounting Failure",
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
						"authen_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication Success",
						},
						"authen_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authentication Failure",
						},
						"authorize_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization Success",
						},
						"authorize_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Authorization Failure",
						},
						"access_challenge": {
							Type: schema.TypeInt, Optional: true, Description: "Total Access-Challenge Message Receive",
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
						"accounting_request_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Accounting-Request Sent",
						},
						"accounting_success": {
							Type: schema.TypeInt, Optional: true, Description: "Accounting Success",
						},
						"accounting_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Accounting Failure",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerRadiusStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadiusStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerRadiusStatsInstanceList := setSliceAamAuthenticationServerRadiusStatsInstanceList(res)
		d.Set("instance_list", AamAuthenticationServerRadiusStatsInstanceList)
		AamAuthenticationServerRadiusStatsStats := setObjectAamAuthenticationServerRadiusStatsStats(res)
		d.Set("stats", AamAuthenticationServerRadiusStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAuthenticationServerRadiusStatsInstanceList(d edpt.DataAamAuthenticationServerRadiusStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAuthenticationServerRadiusStats.InstanceList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectAamAuthenticationServerRadiusStatsInstanceListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationServerRadiusStatsInstanceListStats(d edpt.AamAuthenticationServerRadiusStatsInstanceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["authen_success"] = d.Authen_success

	in["authen_failure"] = d.Authen_failure

	in["authorize_success"] = d.Authorize_success

	in["authorize_failure"] = d.Authorize_failure

	in["access_challenge"] = d.Access_challenge

	in["timeout_error"] = d.Timeout_error

	in["other_error"] = d.Other_error

	in["request"] = d.Request

	in["accounting_request_sent"] = d.AccountingRequestSent

	in["accounting_success"] = d.AccountingSuccess

	in["accounting_failure"] = d.AccountingFailure
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationServerRadiusStatsStats(ret edpt.DataAamAuthenticationServerRadiusStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"authen_success":          ret.DtAamAuthenticationServerRadiusStats.Stats.Authen_success,
			"authen_failure":          ret.DtAamAuthenticationServerRadiusStats.Stats.Authen_failure,
			"authorize_success":       ret.DtAamAuthenticationServerRadiusStats.Stats.Authorize_success,
			"authorize_failure":       ret.DtAamAuthenticationServerRadiusStats.Stats.Authorize_failure,
			"access_challenge":        ret.DtAamAuthenticationServerRadiusStats.Stats.Access_challenge,
			"timeout_error":           ret.DtAamAuthenticationServerRadiusStats.Stats.Timeout_error,
			"other_error":             ret.DtAamAuthenticationServerRadiusStats.Stats.Other_error,
			"request":                 ret.DtAamAuthenticationServerRadiusStats.Stats.Request,
			"request_normal":          ret.DtAamAuthenticationServerRadiusStats.Stats.RequestNormal,
			"request_dropped":         ret.DtAamAuthenticationServerRadiusStats.Stats.RequestDropped,
			"response_success":        ret.DtAamAuthenticationServerRadiusStats.Stats.ResponseSuccess,
			"response_failure":        ret.DtAamAuthenticationServerRadiusStats.Stats.ResponseFailure,
			"response_error":          ret.DtAamAuthenticationServerRadiusStats.Stats.ResponseError,
			"response_timeout":        ret.DtAamAuthenticationServerRadiusStats.Stats.ResponseTimeout,
			"response_other":          ret.DtAamAuthenticationServerRadiusStats.Stats.ResponseOther,
			"job_start_error":         ret.DtAamAuthenticationServerRadiusStats.Stats.JobStartError,
			"polling_control_error":   ret.DtAamAuthenticationServerRadiusStats.Stats.PollingControlError,
			"accounting_request_sent": ret.DtAamAuthenticationServerRadiusStats.Stats.AccountingRequestSent,
			"accounting_success":      ret.DtAamAuthenticationServerRadiusStats.Stats.AccountingSuccess,
			"accounting_failure":      ret.DtAamAuthenticationServerRadiusStats.Stats.AccountingFailure,
		},
	}
}

func getSliceAamAuthenticationServerRadiusStatsInstanceList(d []interface{}) []edpt.AamAuthenticationServerRadiusStatsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusStatsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusStatsInstanceList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectAamAuthenticationServerRadiusStatsInstanceListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerRadiusStatsInstanceListStats(d []interface{}) edpt.AamAuthenticationServerRadiusStatsInstanceListStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadiusStatsInstanceListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authen_success = in["authen_success"].(int)
		ret.Authen_failure = in["authen_failure"].(int)
		ret.Authorize_success = in["authorize_success"].(int)
		ret.Authorize_failure = in["authorize_failure"].(int)
		ret.Access_challenge = in["access_challenge"].(int)
		ret.Timeout_error = in["timeout_error"].(int)
		ret.Other_error = in["other_error"].(int)
		ret.Request = in["request"].(int)
		ret.AccountingRequestSent = in["accounting_request_sent"].(int)
		ret.AccountingSuccess = in["accounting_success"].(int)
		ret.AccountingFailure = in["accounting_failure"].(int)
	}
	return ret
}

func getObjectAamAuthenticationServerRadiusStatsStats(d []interface{}) edpt.AamAuthenticationServerRadiusStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadiusStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authen_success = in["authen_success"].(int)
		ret.Authen_failure = in["authen_failure"].(int)
		ret.Authorize_success = in["authorize_success"].(int)
		ret.Authorize_failure = in["authorize_failure"].(int)
		ret.Access_challenge = in["access_challenge"].(int)
		ret.Timeout_error = in["timeout_error"].(int)
		ret.Other_error = in["other_error"].(int)
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
		ret.AccountingRequestSent = in["accounting_request_sent"].(int)
		ret.AccountingSuccess = in["accounting_success"].(int)
		ret.AccountingFailure = in["accounting_failure"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerRadiusStats(d *schema.ResourceData) edpt.AamAuthenticationServerRadiusStats {
	var ret edpt.AamAuthenticationServerRadiusStats

	ret.InstanceList = getSliceAamAuthenticationServerRadiusStatsInstanceList(d.Get("instance_list").([]interface{}))

	ret.Stats = getObjectAamAuthenticationServerRadiusStatsStats(d.Get("stats").([]interface{}))
	return ret
}
