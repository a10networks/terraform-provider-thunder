package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_persist`: Configure triggers for slb.persist object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistDelete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hash_tbl_trylock_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl lock fail",
						},
						"hash_tbl_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl create fail",
						},
						"hash_tbl_rst_updown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (up/down)",
						},
						"hash_tbl_rst_adddel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (add/del)",
						},
						"url_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for URL hash persist fail",
						},
						"header_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header hash persist fail",
						},
						"src_ip_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP persist fail",
						},
						"src_ip_new_sess_cache_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (c)",
						},
						"src_ip_new_sess_sel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (s)",
						},
						"src_ip_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP hash persist fail",
						},
						"dst_ip_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP persist fail",
						},
						"dst_ip_new_sess_cache_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (c)",
						},
						"dst_ip_new_sess_sel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (s)",
						},
						"dst_ip_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP hash persist fail",
						},
						"cssl_sid_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not found",
						},
						"cssl_sid_not_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not match",
						},
						"sssl_sid_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not found",
						},
						"sssl_sid_not_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not match",
						},
						"ssl_sid_persist_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL SID persist fail",
						},
						"ssl_sid_session_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Create SSL SID fail",
						},
						"cookie_persist_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cookie persist fail",
						},
						"cookie_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Persist cookie not found",
						},
						"cookie_invalid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid persist cookie",
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
						"hash_tbl_trylock_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl lock fail",
						},
						"hash_tbl_create_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl create fail",
						},
						"hash_tbl_rst_updown": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (up/down)",
						},
						"hash_tbl_rst_adddel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Hash tbl reset (add/del)",
						},
						"url_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for URL hash persist fail",
						},
						"header_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header hash persist fail",
						},
						"src_ip_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP persist fail",
						},
						"src_ip_new_sess_cache_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (c)",
						},
						"src_ip_new_sess_sel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP new sess fail (s)",
						},
						"src_ip_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SRC IP hash persist fail",
						},
						"dst_ip_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP persist fail",
						},
						"dst_ip_new_sess_cache_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (c)",
						},
						"dst_ip_new_sess_sel_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP new sess fail (s)",
						},
						"dst_ip_hash_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DST IP hash persist fail",
						},
						"cssl_sid_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not found",
						},
						"cssl_sid_not_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Client SSL SID not match",
						},
						"sssl_sid_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not found",
						},
						"sssl_sid_not_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Server SSL SID not match",
						},
						"ssl_sid_persist_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SSL SID persist fail",
						},
						"ssl_sid_session_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Create SSL SID fail",
						},
						"cookie_persist_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cookie persist fail",
						},
						"cookie_not_found": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Persist cookie not found",
						},
						"cookie_invalid": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid persist cookie",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2059(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2059 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2059
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hash_tbl_trylock_fail = in["hash_tbl_trylock_fail"].(int)
		ret.Hash_tbl_create_fail = in["hash_tbl_create_fail"].(int)
		ret.Hash_tbl_rst_updown = in["hash_tbl_rst_updown"].(int)
		ret.Hash_tbl_rst_adddel = in["hash_tbl_rst_adddel"].(int)
		ret.Url_hash_fail = in["url_hash_fail"].(int)
		ret.Header_hash_fail = in["header_hash_fail"].(int)
		ret.Src_ip_fail = in["src_ip_fail"].(int)
		ret.Src_ip_new_sess_cache_fail = in["src_ip_new_sess_cache_fail"].(int)
		ret.Src_ip_new_sess_sel_fail = in["src_ip_new_sess_sel_fail"].(int)
		ret.Src_ip_hash_fail = in["src_ip_hash_fail"].(int)
		ret.Dst_ip_fail = in["dst_ip_fail"].(int)
		ret.Dst_ip_new_sess_cache_fail = in["dst_ip_new_sess_cache_fail"].(int)
		ret.Dst_ip_new_sess_sel_fail = in["dst_ip_new_sess_sel_fail"].(int)
		ret.Dst_ip_hash_fail = in["dst_ip_hash_fail"].(int)
		ret.Cssl_sid_not_found = in["cssl_sid_not_found"].(int)
		ret.Cssl_sid_not_match = in["cssl_sid_not_match"].(int)
		ret.Sssl_sid_not_found = in["sssl_sid_not_found"].(int)
		ret.Sssl_sid_not_match = in["sssl_sid_not_match"].(int)
		ret.Ssl_sid_persist_fail = in["ssl_sid_persist_fail"].(int)
		ret.Ssl_sid_session_fail = in["ssl_sid_session_fail"].(int)
		ret.Cookie_persist_fail = in["cookie_persist_fail"].(int)
		ret.Cookie_not_found = in["cookie_not_found"].(int)
		ret.Cookie_invalid = in["cookie_invalid"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2060(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2060 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2060
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Hash_tbl_trylock_fail = in["hash_tbl_trylock_fail"].(int)
		ret.Hash_tbl_create_fail = in["hash_tbl_create_fail"].(int)
		ret.Hash_tbl_rst_updown = in["hash_tbl_rst_updown"].(int)
		ret.Hash_tbl_rst_adddel = in["hash_tbl_rst_adddel"].(int)
		ret.Url_hash_fail = in["url_hash_fail"].(int)
		ret.Header_hash_fail = in["header_hash_fail"].(int)
		ret.Src_ip_fail = in["src_ip_fail"].(int)
		ret.Src_ip_new_sess_cache_fail = in["src_ip_new_sess_cache_fail"].(int)
		ret.Src_ip_new_sess_sel_fail = in["src_ip_new_sess_sel_fail"].(int)
		ret.Src_ip_hash_fail = in["src_ip_hash_fail"].(int)
		ret.Dst_ip_fail = in["dst_ip_fail"].(int)
		ret.Dst_ip_new_sess_cache_fail = in["dst_ip_new_sess_cache_fail"].(int)
		ret.Dst_ip_new_sess_sel_fail = in["dst_ip_new_sess_sel_fail"].(int)
		ret.Dst_ip_hash_fail = in["dst_ip_hash_fail"].(int)
		ret.Cssl_sid_not_found = in["cssl_sid_not_found"].(int)
		ret.Cssl_sid_not_match = in["cssl_sid_not_match"].(int)
		ret.Sssl_sid_not_found = in["sssl_sid_not_found"].(int)
		ret.Sssl_sid_not_match = in["sssl_sid_not_match"].(int)
		ret.Ssl_sid_persist_fail = in["ssl_sid_persist_fail"].(int)
		ret.Ssl_sid_session_fail = in["ssl_sid_session_fail"].(int)
		ret.Cookie_persist_fail = in["cookie_persist_fail"].(int)
		ret.Cookie_not_found = in["cookie_not_found"].(int)
		ret.Cookie_invalid = in["cookie_invalid"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2059(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2060(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
