package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"cookie_invalid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid persist cookie",
			},
			"cookie_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Persist cookie not found",
			},
			"cookie_persist_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cookie persist fail",
			},
			"cssl_sid_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not found",
			},
			"cssl_sid_not_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not match",
			},
			"dst_ip_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP persist fail",
			},
			"dst_ip_hash_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP hash persist fail",
			},
			"dst_ip_new_sess_cache_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (c)",
			},
			"dst_ip_new_sess_sel_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (s)",
			},
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"hash_tbl_create_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl create fail",
			},
			"hash_tbl_rst_adddel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (add/del)",
			},
			"hash_tbl_rst_updown": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (up/down)",
			},
			"hash_tbl_trylock_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl lock fail",
			},
			"header_hash_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header hash persist fail",
			},
			"src_ip_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP persist fail",
			},
			"src_ip_hash_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP hash persist fail",
			},
			"src_ip_new_sess_cache_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (c)",
			},
			"src_ip_new_sess_sel_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (s)",
			},
			"ssl_sid_persist_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL SID persist fail",
			},
			"ssl_sid_session_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Create SSL SID fail",
			},
			"sssl_sid_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not found",
			},
			"sssl_sid_not_match": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not match",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
			},
			"url_hash_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for URL hash persist fail",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate
	ret.Inst.Cookie_invalid = d.Get("cookie_invalid").(int)
	ret.Inst.Cookie_not_found = d.Get("cookie_not_found").(int)
	ret.Inst.Cookie_persist_fail = d.Get("cookie_persist_fail").(int)
	ret.Inst.Cssl_sid_not_found = d.Get("cssl_sid_not_found").(int)
	ret.Inst.Cssl_sid_not_match = d.Get("cssl_sid_not_match").(int)
	ret.Inst.Dst_ip_fail = d.Get("dst_ip_fail").(int)
	ret.Inst.Dst_ip_hash_fail = d.Get("dst_ip_hash_fail").(int)
	ret.Inst.Dst_ip_new_sess_cache_fail = d.Get("dst_ip_new_sess_cache_fail").(int)
	ret.Inst.Dst_ip_new_sess_sel_fail = d.Get("dst_ip_new_sess_sel_fail").(int)
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Hash_tbl_create_fail = d.Get("hash_tbl_create_fail").(int)
	ret.Inst.Hash_tbl_rst_adddel = d.Get("hash_tbl_rst_adddel").(int)
	ret.Inst.Hash_tbl_rst_updown = d.Get("hash_tbl_rst_updown").(int)
	ret.Inst.Hash_tbl_trylock_fail = d.Get("hash_tbl_trylock_fail").(int)
	ret.Inst.Header_hash_fail = d.Get("header_hash_fail").(int)
	ret.Inst.Src_ip_fail = d.Get("src_ip_fail").(int)
	ret.Inst.Src_ip_hash_fail = d.Get("src_ip_hash_fail").(int)
	ret.Inst.Src_ip_new_sess_cache_fail = d.Get("src_ip_new_sess_cache_fail").(int)
	ret.Inst.Src_ip_new_sess_sel_fail = d.Get("src_ip_new_sess_sel_fail").(int)
	ret.Inst.Ssl_sid_persist_fail = d.Get("ssl_sid_persist_fail").(int)
	ret.Inst.Ssl_sid_session_fail = d.Get("ssl_sid_session_fail").(int)
	ret.Inst.Sssl_sid_not_found = d.Get("sssl_sid_not_found").(int)
	ret.Inst.Sssl_sid_not_match = d.Get("sssl_sid_not_match").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	ret.Inst.Url_hash_fail = d.Get("url_hash_fail").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
