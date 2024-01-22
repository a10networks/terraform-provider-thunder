package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationCaptchaInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_captcha_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationCaptchaInstanceStatsRead,

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
	}
}

func resourceAamAuthenticationCaptchaInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationCaptchaInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationCaptchaInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationCaptchaInstanceStatsStats := setObjectAamAuthenticationCaptchaInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationCaptchaInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationCaptchaInstanceStatsStats(ret edpt.DataAamAuthenticationCaptchaInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request":       ret.DtAamAuthenticationCaptchaInstanceStats.Stats.Request,
			"verify_succ":   ret.DtAamAuthenticationCaptchaInstanceStats.Stats.VerifySucc,
			"parse_fail":    ret.DtAamAuthenticationCaptchaInstanceStats.Stats.ParseFail,
			"json_fail":     ret.DtAamAuthenticationCaptchaInstanceStats.Stats.JsonFail,
			"attr_fail":     ret.DtAamAuthenticationCaptchaInstanceStats.Stats.AttrFail,
			"timeout_error": ret.DtAamAuthenticationCaptchaInstanceStats.Stats.TimeoutError,
			"other_error":   ret.DtAamAuthenticationCaptchaInstanceStats.Stats.OtherError,
		},
	}
}

func getObjectAamAuthenticationCaptchaInstanceStatsStats(d []interface{}) edpt.AamAuthenticationCaptchaInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationCaptchaInstanceStatsStats
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

func dataToEndpointAamAuthenticationCaptchaInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationCaptchaInstanceStats {
	var ret edpt.AamAuthenticationCaptchaInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationCaptchaInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
