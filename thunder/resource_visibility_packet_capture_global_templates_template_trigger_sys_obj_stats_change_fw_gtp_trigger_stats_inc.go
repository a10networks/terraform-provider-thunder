package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_fw_gtp_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"blade_gtp_c_smp_sig_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed on PU2",
			},
			"blade_gtp_info_ext_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp on PU2",
			},
			"blade_gtp_rate_limit_entry_create_failu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure on PU2",
			},
			"blade_gtp_rate_limit_smp_create_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure on PU2",
			},
			"blade_gtp_rate_limit_t3_ctr_create_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure on PU2",
			},
			"blade_gtp_smp_c_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed on PU2",
			},
			"blade_gtp_smp_dec_sess_count_check_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP on PU2",
			},
			"blade_gtp_smp_path_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed on PU2",
			},
			"blade_gtp_smp_session_count_check_faile": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP on PU2",
			},
			"blade_gtp_smp_sig_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed on PU2",
			},
			"blade_gtp_u_smp_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed on PU2",
			},
			"blade_gtp_u_smp_sig_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U SMP signature check Failed on PU2",
			},
			"blade_out_of_session_memory": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory on PU2",
			},
			"gtp_c_fail_conn_create_slow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C packet failed creating L4-session in slowpath",
			},
			"gtp_c_smp_sig_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP signature check Failed",
			},
			"gtp_info_ext_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-Info ext not found while freeing C-smp",
			},
			"gtp_pathm_fail_conn_create_slow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP path packet failed while creating L4-session in slowpath",
			},
			"gtp_rate_limit_entry_create_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Entry Create Failure",
			},
			"gtp_rate_limit_smp_create_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit SMP Create Failure",
			},
			"gtp_rate_limit_t3_ctr_create_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Rate Limit Dynamic Counters Create Failure",
			},
			"gtp_smp_c_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C SMP check Failed",
			},
			"gtp_smp_dec_sess_count_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
			},
			"gtp_smp_path_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
			},
			"gtp_smp_sig_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
			},
			"gtp_tunnel_rate_limit_entry_create_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
			},
			"gtp_u_fail_conn_create_slow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U packet failed while creating L4-session in slowpath",
			},
			"gtp_u_smp_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP U-SMP check Failed",
			},
			"gtp_u_smp_sig_check_failed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP signature check Failed",
			},
			"gtp_u_tunnel_rate_limit_entry_create_fa": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U Tunnel Level Rate Limit Entry Create Failure",
			},
			"out_of_session_memory": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory for GTP-C",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc
	ret.Inst.BladeGtpCSmpSigCheckFailed = d.Get("blade_gtp_c_smp_sig_check_failed").(int)
	ret.Inst.BladeGtpInfoExtNotFound = d.Get("blade_gtp_info_ext_not_found").(int)
	ret.Inst.BladeGtpRateLimitEntryCreateFailu = d.Get("blade_gtp_rate_limit_entry_create_failu").(int)
	ret.Inst.BladeGtpRateLimitSmpCreateFailure = d.Get("blade_gtp_rate_limit_smp_create_failure").(int)
	ret.Inst.BladeGtpRateLimitT3CtrCreateFail = d.Get("blade_gtp_rate_limit_t3_ctr_create_fail").(int)
	ret.Inst.BladeGtpSmpCCheckFailed = d.Get("blade_gtp_smp_c_check_failed").(int)
	ret.Inst.BladeGtpSmpDecSessCountCheckFail = d.Get("blade_gtp_smp_dec_sess_count_check_fail").(int)
	ret.Inst.BladeGtpSmpPathCheckFailed = d.Get("blade_gtp_smp_path_check_failed").(int)
	ret.Inst.BladeGtpSmpSessionCountCheckFaile = d.Get("blade_gtp_smp_session_count_check_faile").(int)
	ret.Inst.BladeGtpSmpSigCheckFailed = d.Get("blade_gtp_smp_sig_check_failed").(int)
	ret.Inst.BladeGtpUSmpCheckFailed = d.Get("blade_gtp_u_smp_check_failed").(int)
	ret.Inst.BladeGtpUSmpSigCheckFailed = d.Get("blade_gtp_u_smp_sig_check_failed").(int)
	ret.Inst.BladeOutOfSessionMemory = d.Get("blade_out_of_session_memory").(int)
	ret.Inst.GtpCFailConnCreateSlow = d.Get("gtp_c_fail_conn_create_slow").(int)
	ret.Inst.GtpCSmpSigCheckFailed = d.Get("gtp_c_smp_sig_check_failed").(int)
	ret.Inst.GtpInfoExtNotFound = d.Get("gtp_info_ext_not_found").(int)
	ret.Inst.GtpPathmFailConnCreateSlow = d.Get("gtp_pathm_fail_conn_create_slow").(int)
	ret.Inst.GtpRateLimitEntryCreateFailure = d.Get("gtp_rate_limit_entry_create_failure").(int)
	ret.Inst.GtpRateLimitSmpCreateFailure = d.Get("gtp_rate_limit_smp_create_failure").(int)
	ret.Inst.GtpRateLimitT3CtrCreateFailure = d.Get("gtp_rate_limit_t3_ctr_create_failure").(int)
	ret.Inst.GtpSmpCCheckFailed = d.Get("gtp_smp_c_check_failed").(int)
	ret.Inst.GtpSmpDecSessCountCheckFailed = d.Get("gtp_smp_dec_sess_count_check_failed").(int)
	ret.Inst.GtpSmpPathCheckFailed = d.Get("gtp_smp_path_check_failed").(int)
	ret.Inst.GtpSmpSigCheckFailed = d.Get("gtp_smp_sig_check_failed").(int)
	ret.Inst.GtpTunnelRateLimitEntryCreateFail = d.Get("gtp_tunnel_rate_limit_entry_create_fail").(int)
	ret.Inst.GtpUFailConnCreateSlow = d.Get("gtp_u_fail_conn_create_slow").(int)
	ret.Inst.GtpUSmpCheckFailed = d.Get("gtp_u_smp_check_failed").(int)
	ret.Inst.GtpUSmpSigCheckFailed = d.Get("gtp_u_smp_sig_check_failed").(int)
	ret.Inst.GtpUTunnelRateLimitEntryCreateFa = d.Get("gtp_u_tunnel_rate_limit_entry_create_fa").(int)
	ret.Inst.OutOfSessionMemory = d.Get("out_of_session_memory").(int)
	//omit uuid
	ret.Inst.Template_name = d.Get("template_name").(string)
	return ret
}
