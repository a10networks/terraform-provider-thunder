package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCaptchaInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_captcha_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateCaptchaInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify captcha template name",
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

func resourceSlbTemplateCaptchaInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCaptchaInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCaptchaInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateCaptchaInstanceStatsStats := setObjectSlbTemplateCaptchaInstanceStatsStats(res)
		d.Set("stats", SlbTemplateCaptchaInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateCaptchaInstanceStatsStats(ret edpt.DataSlbTemplateCaptchaInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request":       ret.DtSlbTemplateCaptchaInstanceStats.Stats.Request,
			"verify_succ":   ret.DtSlbTemplateCaptchaInstanceStats.Stats.VerifySucc,
			"parse_fail":    ret.DtSlbTemplateCaptchaInstanceStats.Stats.ParseFail,
			"json_fail":     ret.DtSlbTemplateCaptchaInstanceStats.Stats.JsonFail,
			"timeout_error": ret.DtSlbTemplateCaptchaInstanceStats.Stats.TimeoutError,
			"other_error":   ret.DtSlbTemplateCaptchaInstanceStats.Stats.OtherError,
		},
	}
}

func getObjectSlbTemplateCaptchaInstanceStatsStats(d []interface{}) edpt.SlbTemplateCaptchaInstanceStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplateCaptchaInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = in["request"].(int)
		ret.VerifySucc = in["verify_succ"].(int)
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateCaptchaInstanceStats(d *schema.ResourceData) edpt.SlbTemplateCaptchaInstanceStats {
	var ret edpt.SlbTemplateCaptchaInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplateCaptchaInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
