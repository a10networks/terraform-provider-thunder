package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_icap_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"app_serv_conn_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn Err Stats",
			},
			"app_serv_conn_no_pcb_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn no ES PCB Err Stats",
			},
			"chunk1_hdr_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Hdr Err1 Stats",
			},
			"chunk2_hdr_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Hdr Err2 Stats",
			},
			"chunk_bad_trail_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Chunk Bad Trail Err Stats",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"encap_hdr_incomplete_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encap HDR Incomplete Err Stats",
			},
			"http_resp_hdr_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Resp Hdr Err Stats",
			},
			"http_resp_line_parse_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Parse Err Stats",
			},
			"http_resp_line_read_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Read Err Stats",
			},
			"icap_line_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Line Err Stats",
			},
			"icap_ver_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Ver Err Stats",
			},
			"no_icap_resp_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No ICAP Resp Err Stats",
			},
			"no_payload_buff_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload Buff Err Stats",
			},
			"no_payload_next_buff_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload In Next Buff Err Stats",
			},
			"no_status_code_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Status Code Err Stats",
			},
			"prep_req_fail_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Prepare ICAP req fail Err Stats",
			},
			"req_hdr_incomplete_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Req Hdr Incomplete Err Stats",
			},
			"resp_hdr_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Err Stats",
			},
			"resp_hdr_incomplete_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Incomplete Err Stats",
			},
			"resp_line_parse_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Parse Err Stats",
			},
			"resp_line_read_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Read Err Stats",
			},
			"serv_sel_fail_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Select Fail Err Stats",
			},
			"start_icap_conn_fail_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Start ICAP conn fail Stats",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate
	ret.Inst.App_serv_conn_err = d.Get("app_serv_conn_err").(int)
	ret.Inst.App_serv_conn_no_pcb_err = d.Get("app_serv_conn_no_pcb_err").(int)
	ret.Inst.Chunk1_hdr_err = d.Get("chunk1_hdr_err").(int)
	ret.Inst.Chunk2_hdr_err = d.Get("chunk2_hdr_err").(int)
	ret.Inst.Chunk_bad_trail_err = d.Get("chunk_bad_trail_err").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Encap_hdr_incomplete_err = d.Get("encap_hdr_incomplete_err").(int)
	ret.Inst.Http_resp_hdr_err = d.Get("http_resp_hdr_err").(int)
	ret.Inst.Http_resp_line_parse_err = d.Get("http_resp_line_parse_err").(int)
	ret.Inst.Http_resp_line_read_err = d.Get("http_resp_line_read_err").(int)
	ret.Inst.Icap_line_err = d.Get("icap_line_err").(int)
	ret.Inst.Icap_ver_err = d.Get("icap_ver_err").(int)
	ret.Inst.No_icap_resp_err = d.Get("no_icap_resp_err").(int)
	ret.Inst.No_payload_buff_err = d.Get("no_payload_buff_err").(int)
	ret.Inst.No_payload_next_buff_err = d.Get("no_payload_next_buff_err").(int)
	ret.Inst.No_status_code_err = d.Get("no_status_code_err").(int)
	ret.Inst.Prep_req_fail_err = d.Get("prep_req_fail_err").(int)
	ret.Inst.Req_hdr_incomplete_err = d.Get("req_hdr_incomplete_err").(int)
	ret.Inst.Resp_hdr_err = d.Get("resp_hdr_err").(int)
	ret.Inst.Resp_hdr_incomplete_err = d.Get("resp_hdr_incomplete_err").(int)
	ret.Inst.Resp_line_parse_err = d.Get("resp_line_parse_err").(int)
	ret.Inst.Resp_line_read_err = d.Get("resp_line_read_err").(int)
	ret.Inst.Serv_sel_fail_err = d.Get("serv_sel_fail_err").(int)
	ret.Inst.Start_icap_conn_fail_err = d.Get("start_icap_conn_fail_err").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
