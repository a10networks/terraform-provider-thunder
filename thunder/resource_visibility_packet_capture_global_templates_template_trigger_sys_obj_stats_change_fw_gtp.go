package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_gtp`: Configure triggers for fw.gtp object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"out_of_session_memory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory for GTP-C",
						},
						"blade_out_of_session_memory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory on PU2",
						},
						"gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
						},
						"gtp_smp_c_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed",
						},
						"blade_gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed on PU2",
						},
						"blade_gtp_smp_c_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed on PU2",
						},
						"gtp_tunnel_rate_limit_entry_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
						},
						"gtp_u_tunnel_rate_limit_entry_create_fa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U Tunnel Level Rate Limit Entry Create Failure",
						},
						"gtp_rate_limit_smp_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure",
						},
						"gtp_rate_limit_t3_ctr_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure",
						},
						"gtp_rate_limit_entry_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure",
						},
						"blade_gtp_rate_limit_smp_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure on PU2",
						},
						"blade_gtp_rate_limit_t3_ctr_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure on PU2",
						},
						"blade_gtp_rate_limit_entry_create_failu": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure on PU2",
						},
						"gtp_smp_dec_sess_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
						},
						"gtp_u_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed",
						},
						"gtp_info_ext_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp",
						},
						"blade_gtp_smp_dec_sess_count_check_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP on PU2",
						},
						"blade_gtp_u_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed on PU2",
						},
						"blade_gtp_info_ext_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp on PU2",
						},
						"blade_gtp_smp_session_count_check_faile": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP on PU2",
						},
						"gtp_c_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed",
						},
						"blade_gtp_c_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed on PU2",
						},
						"gtp_u_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
						},
						"blade_gtp_u_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U SMP signature check Failed on PU2",
						},
						"gtp_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
						},
						"blade_gtp_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed on PU2",
						},
						"gtp_c_fail_conn_create_slow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C packet failed creating L4-session in slowpath",
						},
						"gtp_u_fail_conn_create_slow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U packet failed while creating L4-session in slowpath",
						},
						"gtp_pathm_fail_conn_create_slow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP path packet failed while creating L4-session in slowpath",
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
						"out_of_session_memory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory for GTP-C",
						},
						"blade_out_of_session_memory": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory on PU2",
						},
						"gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
						},
						"gtp_smp_c_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed",
						},
						"blade_gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed on PU2",
						},
						"blade_gtp_smp_c_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed on PU2",
						},
						"gtp_tunnel_rate_limit_entry_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
						},
						"gtp_u_tunnel_rate_limit_entry_create_fa": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U Tunnel Level Rate Limit Entry Create Failure",
						},
						"gtp_rate_limit_smp_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure",
						},
						"gtp_rate_limit_t3_ctr_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure",
						},
						"gtp_rate_limit_entry_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure",
						},
						"blade_gtp_rate_limit_smp_create_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure on PU2",
						},
						"blade_gtp_rate_limit_t3_ctr_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure on PU2",
						},
						"blade_gtp_rate_limit_entry_create_failu": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure on PU2",
						},
						"gtp_smp_dec_sess_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
						},
						"gtp_u_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed",
						},
						"gtp_info_ext_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp",
						},
						"blade_gtp_smp_dec_sess_count_check_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP on PU2",
						},
						"blade_gtp_u_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed on PU2",
						},
						"blade_gtp_info_ext_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp on PU2",
						},
						"blade_gtp_smp_session_count_check_faile": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP on PU2",
						},
						"gtp_c_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed",
						},
						"blade_gtp_c_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed on PU2",
						},
						"gtp_u_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
						},
						"blade_gtp_u_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U SMP signature check Failed on PU2",
						},
						"gtp_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
						},
						"blade_gtp_smp_sig_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed on PU2",
						},
						"gtp_c_fail_conn_create_slow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C packet failed creating L4-session in slowpath",
						},
						"gtp_u_fail_conn_create_slow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U packet failed while creating L4-session in slowpath",
						},
						"gtp_pathm_fail_conn_create_slow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP path packet failed while creating L4-session in slowpath",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2130(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2130 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2130
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.BladeOutOfSessionMemory = in["blade_out_of_session_memory"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCCheckFailed = in["gtp_smp_c_check_failed"].(int)
		ret.BladeGtpSmpPathCheckFailed = in["blade_gtp_smp_path_check_failed"].(int)
		ret.BladeGtpSmpCCheckFailed = in["blade_gtp_smp_c_check_failed"].(int)
		ret.GtpTunnelRateLimitEntryCreateFail = in["gtp_tunnel_rate_limit_entry_create_fail"].(int)
		ret.GtpUTunnelRateLimitEntryCreateFa = in["gtp_u_tunnel_rate_limit_entry_create_fa"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.BladeGtpRateLimitSmpCreateFailure = in["blade_gtp_rate_limit_smp_create_failure"].(int)
		ret.BladeGtpRateLimitT3CtrCreateFail = in["blade_gtp_rate_limit_t3_ctr_create_fail"].(int)
		ret.BladeGtpRateLimitEntryCreateFailu = in["blade_gtp_rate_limit_entry_create_failu"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		ret.GtpUSmpCheckFailed = in["gtp_u_smp_check_failed"].(int)
		ret.GtpInfoExtNotFound = in["gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpDecSessCountCheckFail = in["blade_gtp_smp_dec_sess_count_check_fail"].(int)
		ret.BladeGtpUSmpCheckFailed = in["blade_gtp_u_smp_check_failed"].(int)
		ret.BladeGtpInfoExtNotFound = in["blade_gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpSessionCountCheckFaile = in["blade_gtp_smp_session_count_check_faile"].(int)
		ret.GtpCSmpSigCheckFailed = in["gtp_c_smp_sig_check_failed"].(int)
		ret.BladeGtpCSmpSigCheckFailed = in["blade_gtp_c_smp_sig_check_failed"].(int)
		ret.GtpUSmpSigCheckFailed = in["gtp_u_smp_sig_check_failed"].(int)
		ret.BladeGtpUSmpSigCheckFailed = in["blade_gtp_u_smp_sig_check_failed"].(int)
		ret.GtpSmpSigCheckFailed = in["gtp_smp_sig_check_failed"].(int)
		ret.BladeGtpSmpSigCheckFailed = in["blade_gtp_smp_sig_check_failed"].(int)
		ret.GtpCFailConnCreateSlow = in["gtp_c_fail_conn_create_slow"].(int)
		ret.GtpUFailConnCreateSlow = in["gtp_u_fail_conn_create_slow"].(int)
		ret.GtpPathmFailConnCreateSlow = in["gtp_pathm_fail_conn_create_slow"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2131(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2131 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2131
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.BladeOutOfSessionMemory = in["blade_out_of_session_memory"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCCheckFailed = in["gtp_smp_c_check_failed"].(int)
		ret.BladeGtpSmpPathCheckFailed = in["blade_gtp_smp_path_check_failed"].(int)
		ret.BladeGtpSmpCCheckFailed = in["blade_gtp_smp_c_check_failed"].(int)
		ret.GtpTunnelRateLimitEntryCreateFail = in["gtp_tunnel_rate_limit_entry_create_fail"].(int)
		ret.GtpUTunnelRateLimitEntryCreateFa = in["gtp_u_tunnel_rate_limit_entry_create_fa"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.BladeGtpRateLimitSmpCreateFailure = in["blade_gtp_rate_limit_smp_create_failure"].(int)
		ret.BladeGtpRateLimitT3CtrCreateFail = in["blade_gtp_rate_limit_t3_ctr_create_fail"].(int)
		ret.BladeGtpRateLimitEntryCreateFailu = in["blade_gtp_rate_limit_entry_create_failu"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		ret.GtpUSmpCheckFailed = in["gtp_u_smp_check_failed"].(int)
		ret.GtpInfoExtNotFound = in["gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpDecSessCountCheckFail = in["blade_gtp_smp_dec_sess_count_check_fail"].(int)
		ret.BladeGtpUSmpCheckFailed = in["blade_gtp_u_smp_check_failed"].(int)
		ret.BladeGtpInfoExtNotFound = in["blade_gtp_info_ext_not_found"].(int)
		ret.BladeGtpSmpSessionCountCheckFaile = in["blade_gtp_smp_session_count_check_faile"].(int)
		ret.GtpCSmpSigCheckFailed = in["gtp_c_smp_sig_check_failed"].(int)
		ret.BladeGtpCSmpSigCheckFailed = in["blade_gtp_c_smp_sig_check_failed"].(int)
		ret.GtpUSmpSigCheckFailed = in["gtp_u_smp_sig_check_failed"].(int)
		ret.BladeGtpUSmpSigCheckFailed = in["blade_gtp_u_smp_sig_check_failed"].(int)
		ret.GtpSmpSigCheckFailed = in["gtp_smp_sig_check_failed"].(int)
		ret.BladeGtpSmpSigCheckFailed = in["blade_gtp_smp_sig_check_failed"].(int)
		ret.GtpCFailConnCreateSlow = in["gtp_c_fail_conn_create_slow"].(int)
		ret.GtpUFailConnCreateSlow = in["gtp_u_fail_conn_create_slow"].(int)
		ret.GtpPathmFailConnCreateSlow = in["gtp_pathm_fail_conn_create_slow"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2130(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2131(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Template_name = d.Get("template_name").(string)
	return ret
}
