package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_auth_server_win`: Configure triggers for aam.authentication.server.windows object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kerberos_timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout",
						},
						"kerberos_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Other Error",
						},
						"kerberos_request_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Dropped Request",
						},
						"kerberos_response_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Failure Response",
						},
						"kerberos_response_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Error Response",
						},
						"kerberos_response_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout Response",
						},
						"kerberos_job_start_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Job Start Error",
						},
						"kerberos_polling_control_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Polling Control Error",
						},
						"kerberos_pw_expiry": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password expiry",
						},
						"kerberos_pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password change failure",
						},
						"kerberos_validate_kdc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Validation Failure",
						},
						"kerberos_generate_kdc_keytab_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Generation Failure",
						},
						"kerberos_delete_kdc_keytab_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Deletion Failure",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_rate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_exceeded_by": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
						},
						"duration": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
						},
						"kerberos_timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout",
						},
						"kerberos_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Other Error",
						},
						"kerberos_request_dropped": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Dropped Request",
						},
						"kerberos_response_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Failure Response",
						},
						"kerberos_response_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Error Response",
						},
						"kerberos_response_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Timeout Response",
						},
						"kerberos_job_start_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Job Start Error",
						},
						"kerberos_polling_control_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos Polling Control Error",
						},
						"kerberos_pw_expiry": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password expiry",
						},
						"kerberos_pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos password change failure",
						},
						"kerberos_validate_kdc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Validation Failure",
						},
						"kerberos_generate_kdc_keytab_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Generation Failure",
						},
						"kerberos_delete_kdc_keytab_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Kerberos KDC Keytab Deletion Failure",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"template_name": {
				Type: schema.TypeString, Required: true, Description: "Template_name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2073(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2073 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2073
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KerberosTimeoutError = in["kerberos_timeout_error"].(int)
		ret.KerberosOtherError = in["kerberos_other_error"].(int)
		ret.KerberosRequestDropped = in["kerberos_request_dropped"].(int)
		ret.KerberosResponseFailure = in["kerberos_response_failure"].(int)
		ret.KerberosResponseError = in["kerberos_response_error"].(int)
		ret.KerberosResponseTimeout = in["kerberos_response_timeout"].(int)
		ret.KerberosJobStartError = in["kerberos_job_start_error"].(int)
		ret.KerberosPollingControlError = in["kerberos_polling_control_error"].(int)
		ret.KerberosPwExpiry = in["kerberos_pw_expiry"].(int)
		ret.KerberosPwChangeFailure = in["kerberos_pw_change_failure"].(int)
		ret.KerberosValidateKdcFailure = in["kerberos_validate_kdc_failure"].(int)
		ret.KerberosGenerateKdcKeytabFailure = in["kerberos_generate_kdc_keytab_failure"].(int)
		ret.KerberosDeleteKdcKeytabFailure = in["kerberos_delete_kdc_keytab_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2074(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2074 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2074
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.KerberosTimeoutError = in["kerberos_timeout_error"].(int)
		ret.KerberosOtherError = in["kerberos_other_error"].(int)
		ret.KerberosRequestDropped = in["kerberos_request_dropped"].(int)
		ret.KerberosResponseFailure = in["kerberos_response_failure"].(int)
		ret.KerberosResponseError = in["kerberos_response_error"].(int)
		ret.KerberosResponseTimeout = in["kerberos_response_timeout"].(int)
		ret.KerberosJobStartError = in["kerberos_job_start_error"].(int)
		ret.KerberosPollingControlError = in["kerberos_polling_control_error"].(int)
		ret.KerberosPwExpiry = in["kerberos_pw_expiry"].(int)
		ret.KerberosPwChangeFailure = in["kerberos_pw_change_failure"].(int)
		ret.KerberosValidateKdcFailure = in["kerberos_validate_kdc_failure"].(int)
		ret.KerberosGenerateKdcKeytabFailure = in["kerberos_generate_kdc_keytab_failure"].(int)
		ret.KerberosDeleteKdcKeytabFailure = in["kerberos_delete_kdc_keytab_failure"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2073(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2074(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Template_name = d.Get("template_name").(string)
	return ret
}
