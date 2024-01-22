package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationAccountStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_account_stats`: Statistics for the object account\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationAccountStatsRead,

		Schema: map[string]*schema.Schema{
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
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationAccountStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccountStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationAccountStatsStats := setObjectAamAuthenticationAccountStatsStats(res)
		d.Set("stats", AamAuthenticationAccountStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationAccountStatsStats(ret edpt.DataAamAuthenticationAccountStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_normal":   ret.DtAamAuthenticationAccountStats.Stats.RequestNormal,
			"request_dropped":  ret.DtAamAuthenticationAccountStats.Stats.RequestDropped,
			"response_success": ret.DtAamAuthenticationAccountStats.Stats.ResponseSuccess,
			"response_failure": ret.DtAamAuthenticationAccountStats.Stats.ResponseFailure,
			"response_error":   ret.DtAamAuthenticationAccountStats.Stats.ResponseError,
			"response_timeout": ret.DtAamAuthenticationAccountStats.Stats.ResponseTimeout,
			"response_other":   ret.DtAamAuthenticationAccountStats.Stats.ResponseOther,
		},
	}
}

func getObjectAamAuthenticationAccountStatsStats(d []interface{}) edpt.AamAuthenticationAccountStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationAccountStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestNormal = in["request_normal"].(int)
		ret.RequestDropped = in["request_dropped"].(int)
		ret.ResponseSuccess = in["response_success"].(int)
		ret.ResponseFailure = in["response_failure"].(int)
		ret.ResponseError = in["response_error"].(int)
		ret.ResponseTimeout = in["response_timeout"].(int)
		ret.ResponseOther = in["response_other"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationAccountStats(d *schema.ResourceData) edpt.AamAuthenticationAccountStats {
	var ret edpt.AamAuthenticationAccountStats

	ret.Stats = getObjectAamAuthenticationAccountStatsStats(d.Get("stats").([]interface{}))
	return ret
}
