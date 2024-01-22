package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_smtp_vport_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"forward_req_data_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward REQ data failure",
			},
			"forward_req_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Forward request failure",
			},
			"get_all_headers_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Get all headers fail",
			},
			"insert_resonse_line_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Ins response line fail",
			},
			"line_extend_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line extend fail",
			},
			"line_table_extend_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Table extend fail",
			},
			"line_too_long": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Line too long",
			},
			"no_proxy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No proxy error",
			},
			"parse_req_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request failure",
			},
			"parse_request_line_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse request line fail",
			},
			"parse_resonse_line_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Parse response line fail",
			},
			"read_request_line_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Read request line fail",
			},
			"recv_server_unknow_reply_code": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Recv server unknown-code",
			},
			"remove_resonse_line_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Del response line fail",
			},
			"send_client_service_not_ready": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Sent client serv-not-rdy",
			},
			"server_select_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server selection failure",
			},
			"server_starttls_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server side STARTTLS fail",
			},
			"snat_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Source NAT failure",
			},
			"too_many_headers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Too many headers",
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
func resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSmtpVportTmplTriggerStatsInc
	ret.Inst.Forward_req_data_fail = d.Get("forward_req_data_fail").(int)
	ret.Inst.Forward_req_fail = d.Get("forward_req_fail").(int)
	ret.Inst.Get_all_headers_fail = d.Get("get_all_headers_fail").(int)
	ret.Inst.Insert_resonse_line_fail = d.Get("insert_resonse_line_fail").(int)
	ret.Inst.Line_extend_fail = d.Get("line_extend_fail").(int)
	ret.Inst.Line_table_extend_fail = d.Get("line_table_extend_fail").(int)
	ret.Inst.Line_too_long = d.Get("line_too_long").(int)
	ret.Inst.No_proxy = d.Get("no_proxy").(int)
	ret.Inst.Parse_req_fail = d.Get("parse_req_fail").(int)
	ret.Inst.Parse_request_line_fail = d.Get("parse_request_line_fail").(int)
	ret.Inst.Parse_resonse_line_fail = d.Get("parse_resonse_line_fail").(int)
	ret.Inst.Read_request_line_fail = d.Get("read_request_line_fail").(int)
	ret.Inst.Recv_server_unknow_reply_code = d.Get("recv_server_unknow_reply_code").(int)
	ret.Inst.Remove_resonse_line_fail = d.Get("remove_resonse_line_fail").(int)
	ret.Inst.Send_client_service_not_ready = d.Get("send_client_service_not_ready").(int)
	ret.Inst.Server_select_fail = d.Get("server_select_fail").(int)
	ret.Inst.Server_starttls_fail = d.Get("server_starttls_fail").(int)
	ret.Inst.Snat_fail = d.Get("snat_fail").(int)
	ret.Inst.Too_many_headers = d.Get("too_many_headers").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
