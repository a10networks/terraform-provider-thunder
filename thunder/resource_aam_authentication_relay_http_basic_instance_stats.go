package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayHttpBasicInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_relay_http_basic_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationRelayHttpBasicInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify HTTP basic authentication relay name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"success": {
							Type: schema.TypeInt, Optional: true, Description: "Success",
						},
						"no_creds": {
							Type: schema.TypeInt, Optional: true, Description: "No Credential",
						},
						"bad_req": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Request",
						},
						"unauth": {
							Type: schema.TypeInt, Optional: true, Description: "Unauthorized",
						},
						"forbidden": {
							Type: schema.TypeInt, Optional: true, Description: "Forbidden",
						},
						"not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Not Found",
						},
						"server_error": {
							Type: schema.TypeInt, Optional: true, Description: "Internal Server Error",
						},
						"unavailable": {
							Type: schema.TypeInt, Optional: true, Description: "Service Unavailable",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationRelayHttpBasicInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayHttpBasicInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayHttpBasicInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationRelayHttpBasicInstanceStatsStats := setObjectAamAuthenticationRelayHttpBasicInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationRelayHttpBasicInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationRelayHttpBasicInstanceStatsStats(ret edpt.DataAamAuthenticationRelayHttpBasicInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"success":      ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.Success,
			"no_creds":     ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.NoCreds,
			"bad_req":      ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.BadReq,
			"unauth":       ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.Unauth,
			"forbidden":    ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.Forbidden,
			"not_found":    ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.NotFound,
			"server_error": ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.ServerError,
			"unavailable":  ret.DtAamAuthenticationRelayHttpBasicInstanceStats.Stats.Unavailable,
		},
	}
}

func getObjectAamAuthenticationRelayHttpBasicInstanceStatsStats(d []interface{}) edpt.AamAuthenticationRelayHttpBasicInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayHttpBasicInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Success = in["success"].(int)
		ret.NoCreds = in["no_creds"].(int)
		ret.BadReq = in["bad_req"].(int)
		ret.Unauth = in["unauth"].(int)
		ret.Forbidden = in["forbidden"].(int)
		ret.NotFound = in["not_found"].(int)
		ret.ServerError = in["server_error"].(int)
		ret.Unavailable = in["unavailable"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayHttpBasicInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationRelayHttpBasicInstanceStats {
	var ret edpt.AamAuthenticationRelayHttpBasicInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationRelayHttpBasicInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
