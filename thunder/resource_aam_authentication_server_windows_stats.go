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
									"ntlm_proto_nego_success": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Protocol Negotiation Success",
									},
									"ntlm_proto_nego_failure": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Protocol Negotiation Failure",
									},
									"ntlm_session_setup_success": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Session Setup Success",
									},
									"ntlm_session_setup_failure": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Session Setup Failure",
									},
									"ntlm_prepare_req_success": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Prepare Request Success",
									},
									"ntlm_prepare_req_error": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Prepare Request Error",
									},
									"ntlm_auth_success": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Authentication Success",
									},
									"ntlm_auth_failure": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Authentication Failure",
									},
									"ntlm_timeout_error": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Timeout",
									},
									"ntlm_other_error": {
										Type: schema.TypeInt, Optional: true, Description: "NTLM Other Error",
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
						"ntlm_authentication_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Authentication Success",
						},
						"ntlm_authentication_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Authentication Failure",
						},
						"ntlm_proto_negotiation_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Protocol Negotiation Success",
						},
						"ntlm_proto_negotiation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Protocol Negotiation Failure",
						},
						"ntlm_session_setup_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Session Setup Success",
						},
						"ntlm_session_setup_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Session Setup Failure",
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
						"ntlm_prepare_req_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Prepare Request Success",
						},
						"ntlm_prepare_req_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Prepare Request Failed",
						},
						"ntlm_timeout_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Timeout",
						},
						"ntlm_other_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Other Error",
						},
						"ntlm_request_normal": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Normal Request",
						},
						"ntlm_request_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Dropped Request",
						},
						"ntlm_response_success": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Success Response",
						},
						"ntlm_response_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Failure Response",
						},
						"ntlm_response_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Error Response",
						},
						"ntlm_response_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Timeout Response",
						},
						"ntlm_response_other": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Other Response",
						},
						"ntlm_job_start_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Job Start Error",
						},
						"ntlm_polling_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total NTLM Polling Control Error",
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

	in["ntlm_proto_nego_success"] = d.Ntlm_proto_nego_success

	in["ntlm_proto_nego_failure"] = d.Ntlm_proto_nego_failure

	in["ntlm_session_setup_success"] = d.Ntlm_session_setup_success

	in["ntlm_session_setup_failure"] = d.Ntlm_session_setup_failure

	in["ntlm_prepare_req_success"] = d.Ntlm_prepare_req_success

	in["ntlm_prepare_req_error"] = d.Ntlm_prepare_req_error

	in["ntlm_auth_success"] = d.Ntlm_auth_success

	in["ntlm_auth_failure"] = d.Ntlm_auth_failure

	in["ntlm_timeout_error"] = d.Ntlm_timeout_error

	in["ntlm_other_error"] = d.Ntlm_other_error

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
			"ntlm_authentication_success":          ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmAuthenticationSuccess,
			"ntlm_authentication_failure":          ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmAuthenticationFailure,
			"ntlm_proto_negotiation_success":       ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmProtoNegotiationSuccess,
			"ntlm_proto_negotiation_failure":       ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmProtoNegotiationFailure,
			"ntlm_session_setup_success":           ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmSessionSetupSuccess,
			"ntlm_session_setup_failed":            ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmSessionSetupFailed,
			"kerberos_request_normal":              ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosRequestNormal,
			"kerberos_request_dropped":             ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosRequestDropped,
			"kerberos_response_success":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseSuccess,
			"kerberos_response_failure":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseFailure,
			"kerberos_response_error":              ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseError,
			"kerberos_response_timeout":            ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseTimeout,
			"kerberos_response_other":              ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosResponseOther,
			"kerberos_job_start_error":             ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosJobStartError,
			"kerberos_polling_control_error":       ret.DtAamAuthenticationServerWindowsStats.Stats.KerberosPollingControlError,
			"ntlm_prepare_req_success":             ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmPrepareReqSuccess,
			"ntlm_prepare_req_failed":              ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmPrepareReqFailed,
			"ntlm_timeout_error":                   ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmTimeoutError,
			"ntlm_other_error":                     ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmOtherError,
			"ntlm_request_normal":                  ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmRequestNormal,
			"ntlm_request_dropped":                 ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmRequestDropped,
			"ntlm_response_success":                ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmResponseSuccess,
			"ntlm_response_failure":                ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmResponseFailure,
			"ntlm_response_error":                  ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmResponseError,
			"ntlm_response_timeout":                ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmResponseTimeout,
			"ntlm_response_other":                  ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmResponseOther,
			"ntlm_job_start_error":                 ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmJobStartError,
			"ntlm_polling_control_error":           ret.DtAamAuthenticationServerWindowsStats.Stats.NtlmPollingControlError,
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
		ret.Ntlm_proto_nego_success = in["ntlm_proto_nego_success"].(int)
		ret.Ntlm_proto_nego_failure = in["ntlm_proto_nego_failure"].(int)
		ret.Ntlm_session_setup_success = in["ntlm_session_setup_success"].(int)
		ret.Ntlm_session_setup_failure = in["ntlm_session_setup_failure"].(int)
		ret.Ntlm_prepare_req_success = in["ntlm_prepare_req_success"].(int)
		ret.Ntlm_prepare_req_error = in["ntlm_prepare_req_error"].(int)
		ret.Ntlm_auth_success = in["ntlm_auth_success"].(int)
		ret.Ntlm_auth_failure = in["ntlm_auth_failure"].(int)
		ret.Ntlm_timeout_error = in["ntlm_timeout_error"].(int)
		ret.Ntlm_other_error = in["ntlm_other_error"].(int)
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
		ret.NtlmAuthenticationSuccess = in["ntlm_authentication_success"].(int)
		ret.NtlmAuthenticationFailure = in["ntlm_authentication_failure"].(int)
		ret.NtlmProtoNegotiationSuccess = in["ntlm_proto_negotiation_success"].(int)
		ret.NtlmProtoNegotiationFailure = in["ntlm_proto_negotiation_failure"].(int)
		ret.NtlmSessionSetupSuccess = in["ntlm_session_setup_success"].(int)
		ret.NtlmSessionSetupFailed = in["ntlm_session_setup_failed"].(int)
		ret.KerberosRequestNormal = in["kerberos_request_normal"].(int)
		ret.KerberosRequestDropped = in["kerberos_request_dropped"].(int)
		ret.KerberosResponseSuccess = in["kerberos_response_success"].(int)
		ret.KerberosResponseFailure = in["kerberos_response_failure"].(int)
		ret.KerberosResponseError = in["kerberos_response_error"].(int)
		ret.KerberosResponseTimeout = in["kerberos_response_timeout"].(int)
		ret.KerberosResponseOther = in["kerberos_response_other"].(int)
		ret.KerberosJobStartError = in["kerberos_job_start_error"].(int)
		ret.KerberosPollingControlError = in["kerberos_polling_control_error"].(int)
		ret.NtlmPrepareReqSuccess = in["ntlm_prepare_req_success"].(int)
		ret.NtlmPrepareReqFailed = in["ntlm_prepare_req_failed"].(int)
		ret.NtlmTimeoutError = in["ntlm_timeout_error"].(int)
		ret.NtlmOtherError = in["ntlm_other_error"].(int)
		ret.NtlmRequestNormal = in["ntlm_request_normal"].(int)
		ret.NtlmRequestDropped = in["ntlm_request_dropped"].(int)
		ret.NtlmResponseSuccess = in["ntlm_response_success"].(int)
		ret.NtlmResponseFailure = in["ntlm_response_failure"].(int)
		ret.NtlmResponseError = in["ntlm_response_error"].(int)
		ret.NtlmResponseTimeout = in["ntlm_response_timeout"].(int)
		ret.NtlmResponseOther = in["ntlm_response_other"].(int)
		ret.NtlmJobStartError = in["ntlm_job_start_error"].(int)
		ret.NtlmPollingControlError = in["ntlm_polling_control_error"].(int)
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
