package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAaaPolicyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_aaa_policy_stats`: Statistics for the object aaa-policy\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAaaPolicyStatsRead,

		Schema: map[string]*schema.Schema{
			"aaa_rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": {
							Type: schema.TypeInt, Required: true, Description: "Specify AAA rule index",
						},
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hit_deny": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hit_auth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hit_bypass": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"failure_bypass": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify AAA policy name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"req": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"req_reject": {
							Type: schema.TypeInt, Optional: true, Description: "Request Rejected",
						},
						"req_auth": {
							Type: schema.TypeInt, Optional: true, Description: "Request Matching Authentication Template",
						},
						"req_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Request Bypassed",
						},
						"req_skip": {
							Type: schema.TypeInt, Optional: true, Description: "Request Skipped",
						},
						"error": {
							Type: schema.TypeInt, Optional: true, Description: "Error",
						},
						"failure_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Failure Bypass",
						},
					},
				},
			},
		},
	}
}

func resourceAamAaaPolicyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAaaPolicyStatsAaaRuleList := setSliceAamAaaPolicyStatsAaaRuleList(res)
		d.Set("aaa_rule_list", AamAaaPolicyStatsAaaRuleList)
		AamAaaPolicyStatsStats := setObjectAamAaaPolicyStatsStats(res)
		d.Set("stats", AamAaaPolicyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAaaPolicyStatsAaaRuleList(d edpt.DataAamAaaPolicyStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAaaPolicyStats.AaaRuleList {
		in := make(map[string]interface{})
		in["index"] = item.Index
		in["stats"] = setObjectAamAaaPolicyStatsAaaRuleListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAaaPolicyStatsAaaRuleListStats(d edpt.AamAaaPolicyStatsAaaRuleListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["total_count"] = d.Total_count

	in["hit_deny"] = d.Hit_deny

	in["hit_auth"] = d.Hit_auth

	in["hit_bypass"] = d.Hit_bypass

	in["failure_bypass"] = d.Failure_bypass
	result = append(result, in)
	return result
}

func setObjectAamAaaPolicyStatsStats(ret edpt.DataAamAaaPolicyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"req":            ret.DtAamAaaPolicyStats.Stats.Req,
			"req_reject":     ret.DtAamAaaPolicyStats.Stats.ReqReject,
			"req_auth":       ret.DtAamAaaPolicyStats.Stats.ReqAuth,
			"req_bypass":     ret.DtAamAaaPolicyStats.Stats.ReqBypass,
			"req_skip":       ret.DtAamAaaPolicyStats.Stats.ReqSkip,
			"error":          ret.DtAamAaaPolicyStats.Stats.Error,
			"failure_bypass": ret.DtAamAaaPolicyStats.Stats.FailureBypass,
		},
	}
}

func getSliceAamAaaPolicyStatsAaaRuleList(d []interface{}) []edpt.AamAaaPolicyStatsAaaRuleList {

	count1 := len(d)
	ret := make([]edpt.AamAaaPolicyStatsAaaRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAaaPolicyStatsAaaRuleList
		oi.Index = in["index"].(int)
		oi.Stats = getObjectAamAaaPolicyStatsAaaRuleListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAaaPolicyStatsAaaRuleListStats(d []interface{}) edpt.AamAaaPolicyStatsAaaRuleListStats {

	count1 := len(d)
	var ret edpt.AamAaaPolicyStatsAaaRuleListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_count = in["total_count"].(int)
		ret.Hit_deny = in["hit_deny"].(int)
		ret.Hit_auth = in["hit_auth"].(int)
		ret.Hit_bypass = in["hit_bypass"].(int)
		ret.Failure_bypass = in["failure_bypass"].(int)
	}
	return ret
}

func getObjectAamAaaPolicyStatsStats(d []interface{}) edpt.AamAaaPolicyStatsStats {

	count1 := len(d)
	var ret edpt.AamAaaPolicyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Req = in["req"].(int)
		ret.ReqReject = in["req_reject"].(int)
		ret.ReqAuth = in["req_auth"].(int)
		ret.ReqBypass = in["req_bypass"].(int)
		ret.ReqSkip = in["req_skip"].(int)
		ret.Error = in["error"].(int)
		ret.FailureBypass = in["failure_bypass"].(int)
	}
	return ret
}

func dataToEndpointAamAaaPolicyStats(d *schema.ResourceData) edpt.AamAaaPolicyStats {
	var ret edpt.AamAaaPolicyStats

	ret.AaaRuleList = getSliceAamAaaPolicyStatsAaaRuleList(d.Get("aaa_rule_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAaaPolicyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
