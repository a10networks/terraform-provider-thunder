package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_imap_vport_tmpl`: Configure template for counter.imap_vport\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplDelete,

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
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
						},
						"no_route": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
						},
						"invalid_start_line": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
						},
						"cant_find_pasv": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
						},
						"smp_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
						},
						"data_server_conn_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
						},
						"data_send_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
						},
						"cant_find_epsv": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
						},
						"auth_unsupported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
						},
						"unsupported_pbsz_value": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
						},
						"unsupported_prot_value": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
						},
						"bad_sequence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
						},
						"rsv_persist_conn_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
						},
						"smp_v6_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
						},
						"smp_v4_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
						},
						"insert_tuple_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
						},
						"cl_est_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
						},
						"ser_connecting_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
						},
						"server_response_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
						},
						"cl_request_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
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
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
						},
						"no_route": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
						},
						"invalid_start_line": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
						},
						"cant_find_pasv": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
						},
						"smp_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
						},
						"data_server_conn_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
						},
						"data_send_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
						},
						"cant_find_epsv": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
						},
						"auth_unsupported": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
						},
						"unsupported_pbsz_value": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
						},
						"unsupported_prot_value": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
						},
						"bad_sequence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
						},
						"rsv_persist_conn_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
						},
						"smp_v6_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
						},
						"smp_v4_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
						},
						"insert_tuple_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
						},
						"cl_est_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
						},
						"ser_connecting_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
						},
						"server_response_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
						},
						"cl_request_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
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
func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc2690(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc2690 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc2690
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate2691(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate2691 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate2691
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsSeverity2692(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsSeverity2692 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsSeverity2692
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsInc2690(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate2691(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsSeverity2692(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
