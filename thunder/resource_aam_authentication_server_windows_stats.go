package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerWindowsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_windows_stats`: Statistics for the object windows\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerWindowsStatsRead,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kerberos_request_send": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Request",
						},
						"kerberos_response_get": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Response",
						},
						"kerberos_timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Timeout",
						},
						"kerberos_other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Other Error",
						},
						"kerberos_request_normal": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Normal Request",
						},
						"kerberos_request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Dropped Request",
						},
						"kerberos_response_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Success Response",
						},
						"kerberos_response_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Failure Response",
						},
						"kerberos_response_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Error Response",
						},
						"kerberos_response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Timeout Response",
						},
						"kerberos_response_other": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Other Response",
						},
						"kerberos_job_start_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Job Start Error",
						},
						"kerberos_polling_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos Polling Control Error",
						},
						"kerberos_pw_expiry": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos password expiry",
						},
						"kerberos_pw_change_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos password change success",
						},
						"kerberos_pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos password change failure",
						},
						"kerberos_validate_kdc_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos KDC Validation Success",
						},
						"kerberos_validate_kdc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos KDC Validation Failure",
						},
						"kerberos_generate_kdc_keytab_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos KDC Keytab Generation Success",
						},
						"kerberos_generate_kdc_keytab_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos KDC Keytab Generation Failure",
						},
						"kerberos_delete_kdc_keytab_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos KDC Keytab Deletion Success",
						},
						"kerberos_delete_kdc_keytab_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total Kerberos KDC Keytab Deletion Failure",
						},
						"kerberos_kdc_keytab_count": {
							Type: schema.TypeInt, Optional: true, Description: "Current Kerberos KDC Keytab Count",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationServerWindowsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerWindowsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerWindowsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerWindowsStatsInstanceList := setSliceAamAuthenticationServerWindowsStatsInstanceList(res)
		d.Set("instance_list", AamAuthenticationServerWindowsStatsInstanceList)
		AamAuthenticationServerWindowsStatsStats := setObjectAamAuthenticationServerWindowsStatsStats(res)
		d.Set("stats", AamAuthenticationServerWindowsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceAamAuthenticationServerWindowsStatsInstanceList(d edpt.DataAamAuthenticationServerWindowsStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtAamAuthenticationServerWindowsStats.InstanceList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectAamAuthenticationServerWindowsStatsInstanceListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectAamAuthenticationServerWindowsStatsInstanceListStats(d edpt.AamAuthenticationServerWindowsStatsInstanceListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["krb_send_req_success"] = d.Krb_send_req_success

	in["krb_get_resp_success"] = d.Krb_get_resp_success

	in["krb_timeout_error"] = d.Krb_timeout_error

	in["krb_other_error"] = d.Krb_other_error

	in["krb_pw_expiry"] = d.Krb_pw_expiry

	in["krb_pw_change_success"] = d.Krb_pw_change_success

	in["krb_pw_change_failure"] = d.Krb_pw_change_failure

	in["krb_validate_kdc_success"] = d.Krb_validate_kdc_success

	in["krb_validate_kdc_failure"] = d.Krb_validate_kdc_failure
	result = append(result, in)
	return result
}

func setObjectAamAuthenticationServerWindowsStatsStats(ret edpt.DataAamAuthenticationServerWindowsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"kerberos_request_send":                ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosRequestSend,
			"kerberos_response_get":                ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseGet,
			"kerberos_timeout_error":               ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosTimeoutError,
			"kerberos_other_error":                 ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosOtherError,
			"kerberos_request_normal":              ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosRequestNormal,
			"kerberos_request_dropped":             ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosRequestDropped,
			"kerberos_response_success":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseSuccess,
			"kerberos_response_failure":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseFailure,
			"kerberos_response_error":              ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseError,
			"kerberos_response_timeout":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseTimeout,
			"kerberos_response_other":              ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseOther,
			"kerberos_job_start_error":             ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosJobStartError,
			"kerberos_polling_control_error":       ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosPollingControlError,
			"kerberos_pw_expiry":                   ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosPwExpiry,
			"kerberos_pw_change_success":           ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosPwChangeSuccess,
			"kerberos_pw_change_failure":           ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosPwChangeFailure,
			"kerberos_validate_kdc_success":        ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosValidateKdcSuccess,
			"kerberos_validate_kdc_failure":        ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosValidateKdcFailure,
			"kerberos_generate_kdc_keytab_success": ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosGenerateKdcKeytabSuccess,
			"kerberos_generate_kdc_keytab_failure": ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosGenerateKdcKeytabFailure,
			"kerberos_delete_kdc_keytab_success":   ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosDeleteKdcKeytabSuccess,
			"kerberos_delete_kdc_keytab_failure":   ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosDeleteKdcKeytabFailure,
			"kerberos_kdc_keytab_count":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosKdcKeytabCount,
		},
	}
}

func getSliceAamAuthenticationServerWindowsStatsInstanceList(d []interface{}) []edpt.AamAuthenticationServerWindowsStatsInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerWindowsStatsInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerWindowsStatsInstanceList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectAamAuthenticationServerWindowsStatsInstanceListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerWindowsStatsInstanceListStats(d []interface{}) edpt.AamAuthenticationServerWindowsStatsInstanceListStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsStatsInstanceListStats
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

func getObjectAamAuthenticationServerWindowsStatsStats(d []interface{}) edpt.AamAuthenticationServerWindowsStatsStats {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerWindowsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KerberosRequestSend = in["kerberos_request_send"].(int)
		ret.KerberosResponseGet = in["kerberos_response_get"].(int)
		ret.KerberosTimeoutError = in["kerberos_timeout_error"].(int)
		ret.KerberosOtherError = in["kerberos_other_error"].(int)
		ret.KerberosRequestNormal = in["kerberos_request_normal"].(int)
		ret.KerberosRequestDropped = in["kerberos_request_dropped"].(int)
		ret.KerberosResponseSuccess = in["kerberos_response_success"].(int)
		ret.KerberosResponseFailure = in["kerberos_response_failure"].(int)
		ret.KerberosResponseError = in["kerberos_response_error"].(int)
		ret.KerberosResponseTimeout = in["kerberos_response_timeout"].(int)
		ret.KerberosResponseOther = in["kerberos_response_other"].(int)
		ret.KerberosJobStartError = in["kerberos_job_start_error"].(int)
		ret.KerberosPollingControlError = in["kerberos_polling_control_error"].(int)
		ret.KerberosPwExpiry = in["kerberos_pw_expiry"].(int)
		ret.KerberosPwChangeSuccess = in["kerberos_pw_change_success"].(int)
		ret.KerberosPwChangeFailure = in["kerberos_pw_change_failure"].(int)
		ret.KerberosValidateKdcSuccess = in["kerberos_validate_kdc_success"].(int)
		ret.KerberosValidateKdcFailure = in["kerberos_validate_kdc_failure"].(int)
		ret.KerberosGenerateKdcKeytabSuccess = in["kerberos_generate_kdc_keytab_success"].(int)
		ret.KerberosGenerateKdcKeytabFailure = in["kerberos_generate_kdc_keytab_failure"].(int)
		ret.KerberosDeleteKdcKeytabSuccess = in["kerberos_delete_kdc_keytab_success"].(int)
		ret.KerberosDeleteKdcKeytabFailure = in["kerberos_delete_kdc_keytab_failure"].(int)
		ret.KerberosKdcKeytabCount = in["kerberos_kdc_keytab_count"].(int)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerWindowsStats(d *schema.ResourceData) edpt.AamAuthenticationServerWindowsStats {
	var ret edpt.AamAuthenticationServerWindowsStats

	ret.InstanceList = getSliceAamAuthenticationServerWindowsStatsInstanceList(d.Get("instance_list").([]interface{}))

	ret.Stats = getObjectAamAuthenticationServerWindowsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
