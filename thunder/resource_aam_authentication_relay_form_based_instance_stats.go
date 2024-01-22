package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayFormBasedInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_form_based_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayFormBasedInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify form-based authentication relay name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"invalid_srv_rsp": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Server Response",
						},
						"post_fail": {
							Type: schema.TypeInt, Optional: true, Description: "POST Failed",
						},
						"invalid_cred": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Credential",
						},
						"bad_req": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Request",
						},
						"not_fnd": {
							Type: schema.TypeInt, Optional: true, Description: "Not Found",
						},
						"error": {
							Type: schema.TypeInt, Optional: true, Description: "Internal Server Error",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Other Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelayFormBasedInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayFormBasedInstanceStatsStats := setObjectAamAuthenticationRelayFormBasedInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationRelayFormBasedInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayFormBasedInstanceStatsStats(ret edpt.DataAamAuthenticationRelayFormBasedInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request":         ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Request,
			"invalid_srv_rsp": ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Invalid_srv_rsp,
			"post_fail":       ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Post_fail,
			"invalid_cred":    ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Invalid_cred,
			"bad_req":         ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Bad_req,
			"not_fnd":         ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Not_fnd,
			"error":           ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Error,
			"other_error":     ret.DtAamAuthenticationRelayFormBasedInstanceStats.Stats.Other_error,
		},
	}
}

func getObjectAamAuthenticationRelayFormBasedInstanceStatsStats(d []interface{}) edpt.AamAuthenticationRelayFormBasedInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayFormBasedInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = in["request"].(int)
		ret.Invalid_srv_rsp = in["invalid_srv_rsp"].(int)
		ret.Post_fail = in["post_fail"].(int)
		ret.Invalid_cred = in["invalid_cred"].(int)
		ret.Bad_req = in["bad_req"].(int)
		ret.Not_fnd = in["not_fnd"].(int)
		ret.Error = in["error"].(int)
		ret.Other_error = in["other_error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayFormBasedInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationRelayFormBasedInstanceStats {
	var ret edpt.AamAuthenticationRelayFormBasedInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelayFormBasedInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
