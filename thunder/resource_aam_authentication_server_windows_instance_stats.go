package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerWindowsInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_windows_instance_stats`: Statistics for the object instance\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerWindowsInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Windows authentication server name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"krb_send_req_success": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos Request",
						},
						"krb_get_resp_success": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos Response",
						},
						"krb_timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos Timeout",
						},
						"krb_other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos Other Error",
						},
						"krb_pw_expiry": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos password expiry",
						},
						"krb_pw_change_success": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos password change success",
						},
						"krb_pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos password change failure",
						},
						"krb_validate_kdc_success": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos KDC Validation Success",
						},
						"krb_validate_kdc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Kerberos KDC Validation Failure",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerWindowsInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerWindowsInstanceStatsStats := setObjectAamAuthenticationServerWindowsInstanceStatsStats(res)
		d.Set("stats", AamAuthenticationServerWindowsInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerWindowsInstanceStatsStats(ret edpt.DataAamAuthenticationServerWindowsInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"krb_send_req_success":     ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_send_req_success,
			"krb_get_resp_success":     ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_get_resp_success,
			"krb_timeout_error":        ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_timeout_error,
			"krb_other_error":          ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_other_error,
			"krb_pw_expiry":            ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_pw_expiry,
			"krb_pw_change_success":    ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_pw_change_success,
			"krb_pw_change_failure":    ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_pw_change_failure,
			"krb_validate_kdc_success": ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_validate_kdc_success,
			"krb_validate_kdc_failure": ret.DtAamAuthenticationServerWindowsInstanceStats.Stats.Krb_validate_kdc_failure,
		},
	}
}

func getObjectAamAuthenticationServerWindowsInstanceStatsStats(d []interface{}) edpt.AamAuthenticationServerWindowsInstanceStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Krb_send_req_success = in["krb_send_req_success"].(int)
		ret.Krb_get_resp_success = in["krb_get_resp_success"].(int)
		ret.Krb_timeout_error = in["krb_timeout_error"].(int)
		ret.Krb_other_error = in["krb_other_error"].(int)
		ret.Krb_pw_expiry = in["krb_pw_expiry"].(int)
		ret.Krb_pw_change_success = in["krb_pw_change_success"].(int)
		ret.Krb_pw_change_failure = in["krb_pw_change_failure"].(int)
		ret.Krb_validate_kdc_success = in["krb_validate_kdc_success"].(int)
		ret.Krb_validate_kdc_failure = in["krb_validate_kdc_failure"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerWindowsInstanceStats(d *schema.ResourceData) edpt.AamAuthenticationServerWindowsInstanceStats {
	var ret edpt.AamAuthenticationServerWindowsInstanceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectAamAuthenticationServerWindowsInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
