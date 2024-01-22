package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerRadiusInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_radius_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerRadiusInstanceStatsRead,

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
	}
}

func resourceAamAuthenticationServerRadiusInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadiusInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerRadiusInstanceStatsStats := setObjectAamAuthenticationServerRadiusInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationServerRadiusInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerRadiusInstanceStatsStats(ret edpt.DataAamAuthenticationServerRadiusInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"authen_success":          ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Authen_success,
			"authen_failure":          ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Authen_failure,
			"authorize_success":       ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Authorize_success,
			"authorize_failure":       ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Authorize_failure,
			"access_challenge":        ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Access_challenge,
			"timeout_error":           ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Timeout_error,
			"other_error":             ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Other_error,
			"request":                 ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.Request,
			"accounting_request_sent": ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.AccountingRequestSent,
			"accounting_success":      ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.AccountingSuccess,
			"accounting_failure":      ret.DtAamAuthenticationServerRadiusInstanceStats.Stats.AccountingFailure,
		},
	}
}

func getObjectAamAuthenticationServerRadiusInstanceStatsStats(d []interface{}) edpt.AamAuthenticationServerRadiusInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadiusInstanceStatsStats
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

func dataToEndpointAamAuthenticationServerRadiusInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationServerRadiusInstanceStats {
	var ret edpt.AamAuthenticationServerRadiusInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationServerRadiusInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
