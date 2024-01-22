package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_imap_proxy`: Configure triggers for slb.imap-proxy object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyDelete,

		Schema: map[string]*schema.Schema{
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2043(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2043 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2043
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

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2044(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2044 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2044
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

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2043(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2044(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
