package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelaySamlStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_saml_stats`: Statistics for the object saml\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelaySamlStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify SAML authentication relay name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Request",
						},
						"success": {
							Type: schema.TypeInt, Optional: true, Description: "Success",
						},
						"failure": {
							Type: schema.TypeInt, Optional: true, Description: "Failure",
						},
						"error": {
							Type: schema.TypeInt, Optional: true, Description: "Error",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelaySamlStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelaySamlStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelaySamlStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelaySamlStatsStats := setObjectAamAuthenticationRelaySamlStatsStats(res)
		d.Set("stats", AamAuthenticationRelaySamlStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelaySamlStatsStats(ret edpt.DataAamAuthenticationRelaySamlStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request": ret.DtAamAuthenticationRelaySamlStats.Stats.Request,
			"success": ret.DtAamAuthenticationRelaySamlStats.Stats.Success,
			"failure": ret.DtAamAuthenticationRelaySamlStats.Stats.Failure,
			"error":   ret.DtAamAuthenticationRelaySamlStats.Stats.Error,
		},
	}
}

func getObjectAamAuthenticationRelaySamlStatsStats(d []interface{}) edpt.AamAuthenticationRelaySamlStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelaySamlStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request = in["request"].(int)
		ret.Success = in["success"].(int)
		ret.Failure = in["failure"].(int)
		ret.Error = in["error"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelaySamlStats(d *schema.ResourceData) edpt.AamAuthenticationRelaySamlStats {
	var ret edpt.AamAuthenticationRelaySamlStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelaySamlStatsStats(d.Get("stats").([]interface{}))
	return ret
}
