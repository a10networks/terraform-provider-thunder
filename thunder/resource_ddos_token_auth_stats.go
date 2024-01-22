package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTokenAuthStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_token_auth_stats`: Statistics for the object token-auth\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosTokenAuthStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token_authentication_matched": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Matched Packets",
						},
						"token_authentication_mismatched": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Mismatched Packets",
						},
						"token_authentication_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Invalid Packets",
						},
					},
				},
			},
		},
	}
}

func resourceDdosTokenAuthStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosTokenAuthStatsStats := setObjectDdosTokenAuthStatsStats(res)
		d.Set("stats", DdosTokenAuthStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosTokenAuthStatsStats(ret edpt.DataDdosTokenAuthStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"token_authentication_matched":    ret.DtDdosTokenAuthStats.Stats.Token_authentication_matched,
			"token_authentication_mismatched": ret.DtDdosTokenAuthStats.Stats.Token_authentication_mismatched,
			"token_authentication_invalid":    ret.DtDdosTokenAuthStats.Stats.Token_authentication_invalid,
		},
	}
}

func getObjectDdosTokenAuthStatsStats(d []interface{}) edpt.DdosTokenAuthStatsStats {

	count1 := len(d)
	var ret edpt.DdosTokenAuthStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Token_authentication_matched = in["token_authentication_matched"].(int)
		ret.Token_authentication_mismatched = in["token_authentication_mismatched"].(int)
		ret.Token_authentication_invalid = in["token_authentication_invalid"].(int)
	}
	return ret
}

func dataToEndpointDdosTokenAuthStats(d *schema.ResourceData) edpt.DdosTokenAuthStats {
	var ret edpt.DdosTokenAuthStats

	ret.Stats = getObjectDdosTokenAuthStatsStats(d.Get("stats").([]interface{}))
	return ret
}
