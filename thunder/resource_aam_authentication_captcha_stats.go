package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationCaptchaStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_captcha_stats`: Statistics for the object captcha\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationCaptchaStatsRead,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify captcha profile name",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"request": {
										Type: schema.TypeInt, Optional: true, Description: "Total Request",
									},
									"verify_succ": {
										Type: schema.TypeInt, Optional: true, Description: "Total Verification Success Response",
									},
									"parse_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Total JSON Response Parse Failure",
									},
									"json_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Total Failure JSON Response",
									},
									"attr_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Total Attibute Check Failure",
									},
									"timeout_error": {
										Type: schema.TypeInt, Optional: true, Description: "Total Timeout",
									},
									"other_error": {
										Type: schema.TypeInt, Optional: true, Description: "Total Other Error",
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
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request",
						},
						"verify_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Total Verification Success Response",
						},
						"parse_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total JSON Response Parse Failure",
						},
						"json_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total Failure JSON Response",
						},
						"attr_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total Attribute Check Failure",
						},
						"timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Timeout",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Error",
						},
						"job_start_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Job Start Error",
						},
						"polling_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Polling Control Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationCaptchaStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptchaStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationCaptchaStatsInstanceList := setSliceAamAuthenticationCaptchaStatsInstanceList(res)
		d.Set("instance_list", AamAuthenticationCaptchaStatsInstanceList)
		AamAuthenticationCaptchaStatsStats := setObjectAamAuthenticationCaptchaStatsStats(res)
		d.Set("stats", AamAuthenticationCaptchaStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAuthenticationCaptchaStatsInstanceList(d edpt.DataAamAuthenticationCaptchaStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAuthenticationCaptchaStats.InstanceList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectAamAuthenticationCaptchaStatsInstanceListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationCaptchaStatsInstanceListStats(d edpt.AamAuthenticationCaptchaStatsInstanceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["request"] = d.Request

	in["verify_succ"] = d.VerifySucc

	in["parse_fail"] = d.ParseFail

	in["json_fail"] = d.JsonFail

	in["attr_fail"] = d.AttrFail

	in["timeout_error"] = d.TimeoutError

	in["other_error"] = d.OtherError
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationCaptchaStatsStats(ret edpt.DataAamAuthenticationCaptchaStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_normal":        ret.DtAamAuthenticationCaptchaStats.Stats.RequestNormal,
			"request_dropped":       ret.DtAamAuthenticationCaptchaStats.Stats.RequestDropped,
			"response_success":      ret.DtAamAuthenticationCaptchaStats.Stats.ResponseSuccess,
			"response_failure":      ret.DtAamAuthenticationCaptchaStats.Stats.ResponseFailure,
			"response_error":        ret.DtAamAuthenticationCaptchaStats.Stats.ResponseError,
			"response_timeout":      ret.DtAamAuthenticationCaptchaStats.Stats.ResponseTimeout,
			"response_other":        ret.DtAamAuthenticationCaptchaStats.Stats.ResponseOther,
			"request":               ret.DtAamAuthenticationCaptchaStats.Stats.Request,
			"verify_succ":           ret.DtAamAuthenticationCaptchaStats.Stats.VerifySucc,
			"parse_fail":            ret.DtAamAuthenticationCaptchaStats.Stats.ParseFail,
			"json_fail":             ret.DtAamAuthenticationCaptchaStats.Stats.JsonFail,
			"attr_fail":             ret.DtAamAuthenticationCaptchaStats.Stats.AttrFail,
			"timeout_error":         ret.DtAamAuthenticationCaptchaStats.Stats.TimeoutError,
			"other_error":           ret.DtAamAuthenticationCaptchaStats.Stats.OtherError,
			"job_start_error":       ret.DtAamAuthenticationCaptchaStats.Stats.JobStartError,
			"polling_control_error": ret.DtAamAuthenticationCaptchaStats.Stats.PollingControlError,
		},
	}
}

func getSliceAamAuthenticationCaptchaStatsInstanceList(d []interface{}) []edpt.AamAuthenticationCaptchaStatsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationCaptchaStatsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationCaptchaStatsInstanceList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectAamAuthenticationCaptchaStatsInstanceListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationCaptchaStatsInstanceListStats(d []interface{}) edpt.AamAuthenticationCaptchaStatsInstanceListStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationCaptchaStatsInstanceListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = in["request"].(int)
		ret.VerifySucc = in["verify_succ"].(int)
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.AttrFail = in["attr_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func getObjectAamAuthenticationCaptchaStatsStats(d []interface{}) edpt.AamAuthenticationCaptchaStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationCaptchaStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestNormal = in["request_normal"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseSuccess = in["response_success"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.ResponseOther = in["response_other"].(int)
		ret.Request = in["request"].(int)
		ret.VerifySucc = in["verify_succ"].(int)
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.AttrFail = in["attr_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		ret.JobStartError = in["job_start_error"].(int)
		ret.PollingControlError = in["polling_control_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationCaptchaStats(d *schema.ResourceData) edpt.AamAuthenticationCaptchaStats {
	var ret edpt.AamAuthenticationCaptchaStats

	ret.InstanceList = getSliceAamAuthenticationCaptchaStatsInstanceList(d.Get("instance_list").([]interface{}))

	ret.Stats = getObjectAamAuthenticationCaptchaStatsStats(d.Get("stats").([]interface{}))
	return ret
}
