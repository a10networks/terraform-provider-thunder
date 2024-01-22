package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_auth_server_win_inst_tmpl`: Configure template for aam.authentication.server.windows.instance\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
			},
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"krb_timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Timeout",
						},
						"krb_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Other Error",
						},
						"krb_pw_expiry": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password expiry",
						},
						"krb_pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password change failure",
						},
						"ntlm_proto_nego_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Protocol Negotiation Failure",
						},
						"ntlm_session_setup_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Session Setup Failure",
						},
						"ntlm_prepare_req_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Prepare Request Error",
						},
						"ntlm_auth_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Authentication Failure",
						},
						"ntlm_timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Timeout",
						},
						"ntlm_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Other Error",
						},
						"krb_validate_kdc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos KDC Validation Failure",
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
						"krb_timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Timeout",
						},
						"krb_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos Other Error",
						},
						"krb_pw_expiry": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password expiry",
						},
						"krb_pw_change_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos password change failure",
						},
						"ntlm_proto_nego_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Protocol Negotiation Failure",
						},
						"ntlm_session_setup_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Session Setup Failure",
						},
						"ntlm_prepare_req_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Prepare Request Error",
						},
						"ntlm_auth_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Authentication Failure",
						},
						"ntlm_timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Timeout",
						},
						"ntlm_other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for NTLM Other Error",
						},
						"krb_validate_kdc_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Kerberos KDC Validation Failure",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
						},
						"error_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
						},
						"error_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
						},
						"error_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
						},
						"drop_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
						},
						"drop_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
						},
						"drop_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc2654(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc2654 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc2654
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Krb_timeout_error = in["krb_timeout_error"].(int)
		ret.Krb_other_error = in["krb_other_error"].(int)
		ret.Krb_pw_expiry = in["krb_pw_expiry"].(int)
		ret.Krb_pw_change_failure = in["krb_pw_change_failure"].(int)
		ret.Ntlm_proto_nego_failure = in["ntlm_proto_nego_failure"].(int)
		ret.Ntlm_session_setup_failure = in["ntlm_session_setup_failure"].(int)
		ret.Ntlm_prepare_req_error = in["ntlm_prepare_req_error"].(int)
		ret.Ntlm_auth_failure = in["ntlm_auth_failure"].(int)
		ret.Ntlm_timeout_error = in["ntlm_timeout_error"].(int)
		ret.Ntlm_other_error = in["ntlm_other_error"].(int)
		ret.Krb_validate_kdc_failure = in["krb_validate_kdc_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsRate2655(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsRate2655 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsRate2655
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Krb_timeout_error = in["krb_timeout_error"].(int)
		ret.Krb_other_error = in["krb_other_error"].(int)
		ret.Krb_pw_expiry = in["krb_pw_expiry"].(int)
		ret.Krb_pw_change_failure = in["krb_pw_change_failure"].(int)
		ret.Ntlm_proto_nego_failure = in["ntlm_proto_nego_failure"].(int)
		ret.Ntlm_session_setup_failure = in["ntlm_session_setup_failure"].(int)
		ret.Ntlm_prepare_req_error = in["ntlm_prepare_req_error"].(int)
		ret.Ntlm_auth_failure = in["ntlm_auth_failure"].(int)
		ret.Ntlm_timeout_error = in["ntlm_timeout_error"].(int)
		ret.Ntlm_other_error = in["ntlm_other_error"].(int)
		ret.Krb_validate_kdc_failure = in["krb_validate_kdc_failure"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsSeverity2656(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsSeverity2656 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsSeverity2656
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsInc2654(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsRate2655(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplTriggerStatsSeverity2656(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
