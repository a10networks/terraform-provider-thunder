package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_imap_vport_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"auth_unsupported": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported auth",
			},
			"bad_sequence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Sequence",
			},
			"cant_find_epsv": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find epsv",
			},
			"cant_find_pasv": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cant find pasv",
			},
			"cl_est_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client EST state erro",
			},
			"cl_request_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client RQ state error",
			},
			"data_send_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data send fail",
			},
			"data_server_conn_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for data svr conn fail",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"insert_tuple_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel insert tuple fail",
			},
			"invalid_start_line": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid start line",
			},
			"line_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for line too long",
			},
			"no_route": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for no route failure",
			},
			"rsv_persist_conn_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel Persist fail",
			},
			"ser_connecting_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv CTNG state error",
			},
			"server_response_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv RESP state error",
			},
			"smp_create_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for smp create fail",
			},
			"smp_v4_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv4 fail",
			},
			"smp_v6_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Serv Sel SMPv6 fail",
			},
			"snat_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for source nat failure",
			},
			"svrsel_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
			},
			"unsupported_pbsz_value": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PBSZ",
			},
			"unsupported_prot_value": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unsupported PROT",
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
func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesImapVportTmplTriggerStatsRate
	ret.Inst.Auth_unsupported = d.Get("auth_unsupported").(int)
	ret.Inst.Bad_sequence = d.Get("bad_sequence").(int)
	ret.Inst.Cant_find_epsv = d.Get("cant_find_epsv").(int)
	ret.Inst.Cant_find_pasv = d.Get("cant_find_pasv").(int)
	ret.Inst.Cl_est_err = d.Get("cl_est_err").(int)
	ret.Inst.Cl_request_err = d.Get("cl_request_err").(int)
	ret.Inst.Data_send_fail = d.Get("data_send_fail").(int)
	ret.Inst.Data_server_conn_fail = d.Get("data_server_conn_fail").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Insert_tuple_fail = d.Get("insert_tuple_fail").(int)
	ret.Inst.Invalid_start_line = d.Get("invalid_start_line").(int)
	ret.Inst.Line_too_long = d.Get("line_too_long").(int)
	ret.Inst.No_route = d.Get("no_route").(int)
	ret.Inst.Rsv_persist_conn_fail = d.Get("rsv_persist_conn_fail").(int)
	ret.Inst.Ser_connecting_err = d.Get("ser_connecting_err").(int)
	ret.Inst.Server_response_err = d.Get("server_response_err").(int)
	ret.Inst.Smp_create_fail = d.Get("smp_create_fail").(int)
	ret.Inst.Smp_v4_fail = d.Get("smp_v4_fail").(int)
	ret.Inst.Smp_v6_fail = d.Get("smp_v6_fail").(int)
	ret.Inst.Snat_fail = d.Get("snat_fail").(int)
	ret.Inst.Svrsel_fail = d.Get("svrsel_fail").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	ret.Inst.Unsupported_pbsz_value = d.Get("unsupported_pbsz_value").(int)
	ret.Inst.Unsupported_prot_value = d.Get("unsupported_prot_value").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
