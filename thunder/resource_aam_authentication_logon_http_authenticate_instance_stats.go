package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationLogonHttpAuthenticateInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_logon_http_authenticate_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationLogonHttpAuthenticateInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify HTTP-Authenticate logon name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spn_krb_request": {
							Type: schema.TypeInt, Optional: true, Description: "SPN Kerberos Request",
						},
						"spn_krb_success": {
							Type: schema.TypeInt, Optional: true, Description: "SPN Kerberos Success",
						},
						"spn_krb_faiure": {
							Type: schema.TypeInt, Optional: true, Description: "SPN Kerberos Failure",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationLogonHttpAuthenticateInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationLogonHttpAuthenticateInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationLogonHttpAuthenticateInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationLogonHttpAuthenticateInstanceStatsStats := setObjectAamAuthenticationLogonHttpAuthenticateInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationLogonHttpAuthenticateInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationLogonHttpAuthenticateInstanceStatsStats(ret edpt.DataAamAuthenticationLogonHttpAuthenticateInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"spn_krb_request": ret.DtAamAuthenticationLogonHttpAuthenticateInstanceStats.Stats.Spn_krb_request,
			"spn_krb_success": ret.DtAamAuthenticationLogonHttpAuthenticateInstanceStats.Stats.Spn_krb_success,
			"spn_krb_faiure":  ret.DtAamAuthenticationLogonHttpAuthenticateInstanceStats.Stats.Spn_krb_faiure,
		},
	}
}

func getObjectAamAuthenticationLogonHttpAuthenticateInstanceStatsStats(d []interface{}) edpt.AamAuthenticationLogonHttpAuthenticateInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Spn_krb_request = in["spn_krb_request"].(int)
		ret.Spn_krb_success = in["spn_krb_success"].(int)
		ret.Spn_krb_faiure = in["spn_krb_faiure"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationLogonHttpAuthenticateInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationLogonHttpAuthenticateInstanceStats {
	var ret edpt.AamAuthenticationLogonHttpAuthenticateInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationLogonHttpAuthenticateInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
