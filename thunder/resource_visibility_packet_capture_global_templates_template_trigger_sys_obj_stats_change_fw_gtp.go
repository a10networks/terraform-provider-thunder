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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory",
						},
						"gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
						},
						"gtp_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP check Failed",
						},
						"gtp_smp_session_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP",
						},
						"gtp_c_ref_count_smp_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C session count on C-smp exceeded 2",
						},
						"gtp_u_smp_in_rml_with_sess": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U smp is marked RML with U-session",
						},
						"gtp_tunnel_rate_limit_entry_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
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
						"gtp_smp_dec_sess_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Out of Tunnel Memory",
						},
						"gtp_smp_path_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP PATH check Failed",
						},
						"gtp_smp_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP SMP check Failed",
						},
						"gtp_smp_session_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is not in range of 0-11 in GTP-C SMP",
						},
						"gtp_c_ref_count_smp_exceeded": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-C session count on C-smp exceeded 2",
						},
						"gtp_u_smp_in_rml_with_sess": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U smp is marked RML with U-session",
						},
						"gtp_tunnel_rate_limit_entry_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP Tunnel Level Rate Limit Entry Create Failure",
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
						"gtp_smp_dec_sess_count_check_failed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for GTP-U session count is 0 in GTP-C SMP",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
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

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2009(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2009 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2009
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCheckFailed = in["gtp_smp_check_failed"].(int)
		ret.GtpSmpSessionCountCheckFailed = in["gtp_smp_session_count_check_failed"].(int)
		ret.GtpCRefCountSmpExceeded = in["gtp_c_ref_count_smp_exceeded"].(int)
		ret.GtpUSmpInRmlWithSess = in["gtp_u_smp_in_rml_with_sess"].(int)
		ret.GtpTunnelRateLimitEntryCreateFail = in["gtp_tunnel_rate_limit_entry_create_fail"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2010(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2010 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2010
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.OutOfSessionMemory = in["out_of_session_memory"].(int)
		ret.GtpSmpPathCheckFailed = in["gtp_smp_path_check_failed"].(int)
		ret.GtpSmpCheckFailed = in["gtp_smp_check_failed"].(int)
		ret.GtpSmpSessionCountCheckFailed = in["gtp_smp_session_count_check_failed"].(int)
		ret.GtpCRefCountSmpExceeded = in["gtp_c_ref_count_smp_exceeded"].(int)
		ret.GtpUSmpInRmlWithSess = in["gtp_u_smp_in_rml_with_sess"].(int)
		ret.GtpTunnelRateLimitEntryCreateFail = in["gtp_tunnel_rate_limit_entry_create_fail"].(int)
		ret.GtpRateLimitSmpCreateFailure = in["gtp_rate_limit_smp_create_failure"].(int)
		ret.GtpRateLimitT3CtrCreateFailure = in["gtp_rate_limit_t3_ctr_create_failure"].(int)
		ret.GtpRateLimitEntryCreateFailure = in["gtp_rate_limit_entry_create_failure"].(int)
		ret.GtpSmpDecSessCountCheckFailed = in["gtp_smp_dec_sess_count_check_failed"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2009(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2010(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
