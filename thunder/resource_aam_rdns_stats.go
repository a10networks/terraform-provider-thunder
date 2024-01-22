package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamRdnsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_rdns_stats`: Statistics for the object rdns\n\n__PLACEHOLDER__",
		ReadContext: resourceAamRdnsStatsRead,

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

func resourceAamRdnsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamRdnsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamRdnsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamRdnsStatsStats := setObjectAamRdnsStatsStats(res)
		d.Set("stats", AamRdnsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamRdnsStatsStats(ret edpt.DataAamRdnsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_normal":   ret.DtAamRdnsStats.Stats.RequestNormal,
			"request_dropped":  ret.DtAamRdnsStats.Stats.RequestDropped,
			"response_success": ret.DtAamRdnsStats.Stats.ResponseSuccess,
			"response_failure": ret.DtAamRdnsStats.Stats.ResponseFailure,
			"response_error":   ret.DtAamRdnsStats.Stats.ResponseError,
			"response_timeout": ret.DtAamRdnsStats.Stats.ResponseTimeout,
			"response_other":   ret.DtAamRdnsStats.Stats.ResponseOther,
		},
	}
}

func getObjectAamRdnsStatsStats(d []interface{}) edpt.AamRdnsStatsStats {

	count1 := len(d)
	var ret edpt.AamRdnsStatsStats
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

func dataToEndpointAamRdnsStats(d *schema.ResourceData) edpt.AamRdnsStats {
	var ret edpt.AamRdnsStats

	ret.Stats = getObjectAamRdnsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
