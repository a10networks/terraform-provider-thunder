package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_icap`: Configure triggers for slb.icap object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_serv_conn_no_pcb_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn no ES PCB Err Stats",
						},
						"app_serv_conn_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn Err Stats",
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
						"no_payload_next_buff_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload In Next Buff Err Stats",
						},
						"no_payload_buff_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload Buff Err Stats",
						},
						"resp_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Incomplete Err Stats",
						},
						"serv_sel_fail_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Select Fail Err Stats",
						},
						"start_icap_conn_fail_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Start ICAP conn fail Stats",
						},
						"prep_req_fail_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Prepare ICAP req fail Err Stats",
						},
						"icap_ver_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Ver Err Stats",
						},
						"icap_line_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Line Err Stats",
						},
						"encap_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encap HDR Incomplete Err Stats",
						},
						"no_icap_resp_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No ICAP Resp Err Stats",
						},
						"resp_line_read_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Read Err Stats",
						},
						"resp_line_parse_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Parse Err Stats",
						},
						"resp_hdr_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Err Stats",
						},
						"req_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Req Hdr Incomplete Err Stats",
						},
						"no_status_code_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Status Code Err Stats",
						},
						"http_resp_line_read_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Read Err Stats",
						},
						"http_resp_line_parse_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Parse Err Stats",
						},
						"http_resp_hdr_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Resp Hdr Err Stats",
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
						"app_serv_conn_no_pcb_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn no ES PCB Err Stats",
						},
						"app_serv_conn_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for App Server Conn Err Stats",
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
						"no_payload_next_buff_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload In Next Buff Err Stats",
						},
						"no_payload_buff_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Payload Buff Err Stats",
						},
						"resp_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Incomplete Err Stats",
						},
						"serv_sel_fail_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server Select Fail Err Stats",
						},
						"start_icap_conn_fail_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Start ICAP conn fail Stats",
						},
						"prep_req_fail_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Prepare ICAP req fail Err Stats",
						},
						"icap_ver_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Ver Err Stats",
						},
						"icap_line_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for ICAP Line Err Stats",
						},
						"encap_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Encap HDR Incomplete Err Stats",
						},
						"no_icap_resp_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No ICAP Resp Err Stats",
						},
						"resp_line_read_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Read Err Stats",
						},
						"resp_line_parse_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Line Parse Err Stats",
						},
						"resp_hdr_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Resp Hdr Err Stats",
						},
						"req_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Req Hdr Incomplete Err Stats",
						},
						"no_status_code_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for No Status Code Err Stats",
						},
						"http_resp_line_read_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Read Err Stats",
						},
						"http_resp_line_parse_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Response Line Parse Err Stats",
						},
						"http_resp_hdr_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP Resp Hdr Err Stats",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2041(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2041 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2041
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.App_serv_conn_no_pcb_err = in["app_serv_conn_no_pcb_err"].(int)
		ret.App_serv_conn_err = in["app_serv_conn_err"].(int)
		ret.Chunk1_hdr_err = in["chunk1_hdr_err"].(int)
		ret.Chunk2_hdr_err = in["chunk2_hdr_err"].(int)
		ret.Chunk_bad_trail_err = in["chunk_bad_trail_err"].(int)
		ret.No_payload_next_buff_err = in["no_payload_next_buff_err"].(int)
		ret.No_payload_buff_err = in["no_payload_buff_err"].(int)
		ret.Resp_hdr_incomplete_err = in["resp_hdr_incomplete_err"].(int)
		ret.Serv_sel_fail_err = in["serv_sel_fail_err"].(int)
		ret.Start_icap_conn_fail_err = in["start_icap_conn_fail_err"].(int)
		ret.Prep_req_fail_err = in["prep_req_fail_err"].(int)
		ret.Icap_ver_err = in["icap_ver_err"].(int)
		ret.Icap_line_err = in["icap_line_err"].(int)
		ret.Encap_hdr_incomplete_err = in["encap_hdr_incomplete_err"].(int)
		ret.No_icap_resp_err = in["no_icap_resp_err"].(int)
		ret.Resp_line_read_err = in["resp_line_read_err"].(int)
		ret.Resp_line_parse_err = in["resp_line_parse_err"].(int)
		ret.Resp_hdr_err = in["resp_hdr_err"].(int)
		ret.Req_hdr_incomplete_err = in["req_hdr_incomplete_err"].(int)
		ret.No_status_code_err = in["no_status_code_err"].(int)
		ret.Http_resp_line_read_err = in["http_resp_line_read_err"].(int)
		ret.Http_resp_line_parse_err = in["http_resp_line_parse_err"].(int)
		ret.Http_resp_hdr_err = in["http_resp_hdr_err"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2042(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2042 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2042
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.App_serv_conn_no_pcb_err = in["app_serv_conn_no_pcb_err"].(int)
		ret.App_serv_conn_err = in["app_serv_conn_err"].(int)
		ret.Chunk1_hdr_err = in["chunk1_hdr_err"].(int)
		ret.Chunk2_hdr_err = in["chunk2_hdr_err"].(int)
		ret.Chunk_bad_trail_err = in["chunk_bad_trail_err"].(int)
		ret.No_payload_next_buff_err = in["no_payload_next_buff_err"].(int)
		ret.No_payload_buff_err = in["no_payload_buff_err"].(int)
		ret.Resp_hdr_incomplete_err = in["resp_hdr_incomplete_err"].(int)
		ret.Serv_sel_fail_err = in["serv_sel_fail_err"].(int)
		ret.Start_icap_conn_fail_err = in["start_icap_conn_fail_err"].(int)
		ret.Prep_req_fail_err = in["prep_req_fail_err"].(int)
		ret.Icap_ver_err = in["icap_ver_err"].(int)
		ret.Icap_line_err = in["icap_line_err"].(int)
		ret.Encap_hdr_incomplete_err = in["encap_hdr_incomplete_err"].(int)
		ret.No_icap_resp_err = in["no_icap_resp_err"].(int)
		ret.Resp_line_read_err = in["resp_line_read_err"].(int)
		ret.Resp_line_parse_err = in["resp_line_parse_err"].(int)
		ret.Resp_hdr_err = in["resp_hdr_err"].(int)
		ret.Req_hdr_incomplete_err = in["req_hdr_incomplete_err"].(int)
		ret.No_status_code_err = in["no_status_code_err"].(int)
		ret.Http_resp_line_read_err = in["http_resp_line_read_err"].(int)
		ret.Http_resp_line_parse_err = in["http_resp_line_parse_err"].(int)
		ret.Http_resp_hdr_err = in["http_resp_hdr_err"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2041(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2042(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
