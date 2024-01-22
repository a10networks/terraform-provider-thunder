package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPersistStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_persist_stats`: Statistics for the object persist\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPersistStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hash_tbl_trylock_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Hash tbl lock fail",
						},
						"hash_tbl_create_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Hash tbl create ok",
						},
						"hash_tbl_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Hash tbl create fail",
						},
						"hash_tbl_free": {
							Type: schema.TypeInt, Optional: true, Description: "Hash tbl free",
						},
						"hash_tbl_rst_updown": {
							Type: schema.TypeInt, Optional: true, Description: "Hash tbl reset (up/down)",
						},
						"hash_tbl_rst_adddel": {
							Type: schema.TypeInt, Optional: true, Description: "Hash tbl reset (add/del)",
						},
						"url_hash_pri": {
							Type: schema.TypeInt, Optional: true, Description: "URL hash persist (pri)",
						},
						"url_hash_enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "URL hash persist (enQ)",
						},
						"url_hash_sec": {
							Type: schema.TypeInt, Optional: true, Description: "URL hash persist (sec)",
						},
						"url_hash_fail": {
							Type: schema.TypeInt, Optional: true, Description: "URL hash persist fail",
						},
						"header_hash_pri": {
							Type: schema.TypeInt, Optional: true, Description: "Header hash persist(pri)",
						},
						"header_hash_enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "Header hash persist(enQ)",
						},
						"header_hash_sec": {
							Type: schema.TypeInt, Optional: true, Description: "Header hash persist(sec)",
						},
						"header_hash_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Header hash persist fail",
						},
						"src_ip": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP persist ok",
						},
						"src_ip_enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP persist enqueue",
						},
						"src_ip_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP persist fail",
						},
						"src_ip_new_sess_cache": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP new sess (cache)",
						},
						"src_ip_new_sess_cache_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP new sess fail (c)",
						},
						"src_ip_new_sess_sel": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP new sess (select)",
						},
						"src_ip_new_sess_sel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP new sess fail (s)",
						},
						"src_ip_hash_pri": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP hash persist(pri)",
						},
						"src_ip_hash_enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP hash persist(enQ)",
						},
						"src_ip_hash_sec": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP hash persist(sec)",
						},
						"src_ip_hash_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SRC IP hash persist fail",
						},
						"src_ip_enforce": {
							Type: schema.TypeInt, Optional: true, Description: "Enforce higher priority",
						},
						"dst_ip": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP persist ok",
						},
						"dst_ip_enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP persist enqueue",
						},
						"dst_ip_fail": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP persist fail",
						},
						"dst_ip_new_sess_cache": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP new sess (cache)",
						},
						"dst_ip_new_sess_cache_fail": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP new sess fail (c)",
						},
						"dst_ip_new_sess_sel": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP new sess (select)",
						},
						"dst_ip_new_sess_sel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP new sess fail (s)",
						},
						"dst_ip_hash_pri": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP hash persist(pri)",
						},
						"dst_ip_hash_enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP hash persist(enQ)",
						},
						"dst_ip_hash_sec": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP hash persist(sec)",
						},
						"dst_ip_hash_fail": {
							Type: schema.TypeInt, Optional: true, Description: "DST IP hash persist fail",
						},
						"cssl_sid_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Client SSL SID not found",
						},
						"cssl_sid_match": {
							Type: schema.TypeInt, Optional: true, Description: "Client SSL SID match",
						},
						"cssl_sid_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Client SSL SID not match",
						},
						"sssl_sid_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL SID not found",
						},
						"sssl_sid_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL SID reset",
						},
						"sssl_sid_match": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL SID match",
						},
						"sssl_sid_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL SID not match",
						},
						"ssl_sid_persist_ok": {
							Type: schema.TypeInt, Optional: true, Description: "SSL SID persist ok",
						},
						"ssl_sid_persist_fail": {
							Type: schema.TypeInt, Optional: true, Description: "SSL SID persist fail",
						},
						"ssl_sid_session_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Create SSL SID ok",
						},
						"ssl_sid_session_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Create SSL SID fail",
						},
						"cookie_persist_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Cookie persist ok",
						},
						"cookie_persist_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Cookie persist fail",
						},
						"cookie_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Persist cookie not found",
						},
						"cookie_pass_thru": {
							Type: schema.TypeInt, Optional: true, Description: "Persist cookie Pass-thru",
						},
						"cookie_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid persist cookie",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPersistStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPersistStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPersistStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPersistStatsStats := setObjectSlbPersistStatsStats(res)
		d.Set("stats", SlbPersistStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPersistStatsStats(ret edpt.DataSlbPersistStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hash_tbl_trylock_fail":      ret.DtSlbPersistStats.Stats.Hash_tbl_trylock_fail,
			"hash_tbl_create_ok":         ret.DtSlbPersistStats.Stats.Hash_tbl_create_ok,
			"hash_tbl_create_fail":       ret.DtSlbPersistStats.Stats.Hash_tbl_create_fail,
			"hash_tbl_free":              ret.DtSlbPersistStats.Stats.Hash_tbl_free,
			"hash_tbl_rst_updown":        ret.DtSlbPersistStats.Stats.Hash_tbl_rst_updown,
			"hash_tbl_rst_adddel":        ret.DtSlbPersistStats.Stats.Hash_tbl_rst_adddel,
			"url_hash_pri":               ret.DtSlbPersistStats.Stats.Url_hash_pri,
			"url_hash_enqueue":           ret.DtSlbPersistStats.Stats.Url_hash_enqueue,
			"url_hash_sec":               ret.DtSlbPersistStats.Stats.Url_hash_sec,
			"url_hash_fail":              ret.DtSlbPersistStats.Stats.Url_hash_fail,
			"header_hash_pri":            ret.DtSlbPersistStats.Stats.Header_hash_pri,
			"header_hash_enqueue":        ret.DtSlbPersistStats.Stats.Header_hash_enqueue,
			"header_hash_sec":            ret.DtSlbPersistStats.Stats.Header_hash_sec,
			"header_hash_fail":           ret.DtSlbPersistStats.Stats.Header_hash_fail,
			"src_ip":                     ret.DtSlbPersistStats.Stats.Src_ip,
			"src_ip_enqueue":             ret.DtSlbPersistStats.Stats.Src_ip_enqueue,
			"src_ip_fail":                ret.DtSlbPersistStats.Stats.Src_ip_fail,
			"src_ip_new_sess_cache":      ret.DtSlbPersistStats.Stats.Src_ip_new_sess_cache,
			"src_ip_new_sess_cache_fail": ret.DtSlbPersistStats.Stats.Src_ip_new_sess_cache_fail,
			"src_ip_new_sess_sel":        ret.DtSlbPersistStats.Stats.Src_ip_new_sess_sel,
			"src_ip_new_sess_sel_fail":   ret.DtSlbPersistStats.Stats.Src_ip_new_sess_sel_fail,
			"src_ip_hash_pri":            ret.DtSlbPersistStats.Stats.Src_ip_hash_pri,
			"src_ip_hash_enqueue":        ret.DtSlbPersistStats.Stats.Src_ip_hash_enqueue,
			"src_ip_hash_sec":            ret.DtSlbPersistStats.Stats.Src_ip_hash_sec,
			"src_ip_hash_fail":           ret.DtSlbPersistStats.Stats.Src_ip_hash_fail,
			"src_ip_enforce":             ret.DtSlbPersistStats.Stats.Src_ip_enforce,
			"dst_ip":                     ret.DtSlbPersistStats.Stats.Dst_ip,
			"dst_ip_enqueue":             ret.DtSlbPersistStats.Stats.Dst_ip_enqueue,
			"dst_ip_fail":                ret.DtSlbPersistStats.Stats.Dst_ip_fail,
			"dst_ip_new_sess_cache":      ret.DtSlbPersistStats.Stats.Dst_ip_new_sess_cache,
			"dst_ip_new_sess_cache_fail": ret.DtSlbPersistStats.Stats.Dst_ip_new_sess_cache_fail,
			"dst_ip_new_sess_sel":        ret.DtSlbPersistStats.Stats.Dst_ip_new_sess_sel,
			"dst_ip_new_sess_sel_fail":   ret.DtSlbPersistStats.Stats.Dst_ip_new_sess_sel_fail,
			"dst_ip_hash_pri":            ret.DtSlbPersistStats.Stats.Dst_ip_hash_pri,
			"dst_ip_hash_enqueue":        ret.DtSlbPersistStats.Stats.Dst_ip_hash_enqueue,
			"dst_ip_hash_sec":            ret.DtSlbPersistStats.Stats.Dst_ip_hash_sec,
			"dst_ip_hash_fail":           ret.DtSlbPersistStats.Stats.Dst_ip_hash_fail,
			"cssl_sid_not_found":         ret.DtSlbPersistStats.Stats.Cssl_sid_not_found,
			"cssl_sid_match":             ret.DtSlbPersistStats.Stats.Cssl_sid_match,
			"cssl_sid_not_match":         ret.DtSlbPersistStats.Stats.Cssl_sid_not_match,
			"sssl_sid_not_found":         ret.DtSlbPersistStats.Stats.Sssl_sid_not_found,
			"sssl_sid_reset":             ret.DtSlbPersistStats.Stats.Sssl_sid_reset,
			"sssl_sid_match":             ret.DtSlbPersistStats.Stats.Sssl_sid_match,
			"sssl_sid_not_match":         ret.DtSlbPersistStats.Stats.Sssl_sid_not_match,
			"ssl_sid_persist_ok":         ret.DtSlbPersistStats.Stats.Ssl_sid_persist_ok,
			"ssl_sid_persist_fail":       ret.DtSlbPersistStats.Stats.Ssl_sid_persist_fail,
			"ssl_sid_session_ok":         ret.DtSlbPersistStats.Stats.Ssl_sid_session_ok,
			"ssl_sid_session_fail":       ret.DtSlbPersistStats.Stats.Ssl_sid_session_fail,
			"cookie_persist_ok":          ret.DtSlbPersistStats.Stats.Cookie_persist_ok,
			"cookie_persist_fail":        ret.DtSlbPersistStats.Stats.Cookie_persist_fail,
			"cookie_not_found":           ret.DtSlbPersistStats.Stats.Cookie_not_found,
			"cookie_pass_thru":           ret.DtSlbPersistStats.Stats.Cookie_pass_thru,
			"cookie_invalid":             ret.DtSlbPersistStats.Stats.Cookie_invalid,
		},
	}
}

func getObjectSlbPersistStatsStats(d []interface{}) edpt.SlbPersistStatsStats {

	count1 := len(d)
	var ret edpt.SlbPersistStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hash_tbl_trylock_fail = in["hash_tbl_trylock_fail"].(int)
		ret.Hash_tbl_create_ok = in["hash_tbl_create_ok"].(int)
		ret.Hash_tbl_create_fail = in["hash_tbl_create_fail"].(int)
		ret.Hash_tbl_free = in["hash_tbl_free"].(int)
		ret.Hash_tbl_rst_updown = in["hash_tbl_rst_updown"].(int)
		ret.Hash_tbl_rst_adddel = in["hash_tbl_rst_adddel"].(int)
		ret.Url_hash_pri = in["url_hash_pri"].(int)
		ret.Url_hash_enqueue = in["url_hash_enqueue"].(int)
		ret.Url_hash_sec = in["url_hash_sec"].(int)
		ret.Url_hash_fail = in["url_hash_fail"].(int)
		ret.Header_hash_pri = in["header_hash_pri"].(int)
		ret.Header_hash_enqueue = in["header_hash_enqueue"].(int)
		ret.Header_hash_sec = in["header_hash_sec"].(int)
		ret.Header_hash_fail = in["header_hash_fail"].(int)
		ret.Src_ip = in["src_ip"].(int)
		ret.Src_ip_enqueue = in["src_ip_enqueue"].(int)
		ret.Src_ip_fail = in["src_ip_fail"].(int)
		ret.Src_ip_new_sess_cache = in["src_ip_new_sess_cache"].(int)
		ret.Src_ip_new_sess_cache_fail = in["src_ip_new_sess_cache_fail"].(int)
		ret.Src_ip_new_sess_sel = in["src_ip_new_sess_sel"].(int)
		ret.Src_ip_new_sess_sel_fail = in["src_ip_new_sess_sel_fail"].(int)
		ret.Src_ip_hash_pri = in["src_ip_hash_pri"].(int)
		ret.Src_ip_hash_enqueue = in["src_ip_hash_enqueue"].(int)
		ret.Src_ip_hash_sec = in["src_ip_hash_sec"].(int)
		ret.Src_ip_hash_fail = in["src_ip_hash_fail"].(int)
		ret.Src_ip_enforce = in["src_ip_enforce"].(int)
		ret.Dst_ip = in["dst_ip"].(int)
		ret.Dst_ip_enqueue = in["dst_ip_enqueue"].(int)
		ret.Dst_ip_fail = in["dst_ip_fail"].(int)
		ret.Dst_ip_new_sess_cache = in["dst_ip_new_sess_cache"].(int)
		ret.Dst_ip_new_sess_cache_fail = in["dst_ip_new_sess_cache_fail"].(int)
		ret.Dst_ip_new_sess_sel = in["dst_ip_new_sess_sel"].(int)
		ret.Dst_ip_new_sess_sel_fail = in["dst_ip_new_sess_sel_fail"].(int)
		ret.Dst_ip_hash_pri = in["dst_ip_hash_pri"].(int)
		ret.Dst_ip_hash_enqueue = in["dst_ip_hash_enqueue"].(int)
		ret.Dst_ip_hash_sec = in["dst_ip_hash_sec"].(int)
		ret.Dst_ip_hash_fail = in["dst_ip_hash_fail"].(int)
		ret.Cssl_sid_not_found = in["cssl_sid_not_found"].(int)
		ret.Cssl_sid_match = in["cssl_sid_match"].(int)
		ret.Cssl_sid_not_match = in["cssl_sid_not_match"].(int)
		ret.Sssl_sid_not_found = in["sssl_sid_not_found"].(int)
		ret.Sssl_sid_reset = in["sssl_sid_reset"].(int)
		ret.Sssl_sid_match = in["sssl_sid_match"].(int)
		ret.Sssl_sid_not_match = in["sssl_sid_not_match"].(int)
		ret.Ssl_sid_persist_ok = in["ssl_sid_persist_ok"].(int)
		ret.Ssl_sid_persist_fail = in["ssl_sid_persist_fail"].(int)
		ret.Ssl_sid_session_ok = in["ssl_sid_session_ok"].(int)
		ret.Ssl_sid_session_fail = in["ssl_sid_session_fail"].(int)
		ret.Cookie_persist_ok = in["cookie_persist_ok"].(int)
		ret.Cookie_persist_fail = in["cookie_persist_fail"].(int)
		ret.Cookie_not_found = in["cookie_not_found"].(int)
		ret.Cookie_pass_thru = in["cookie_pass_thru"].(int)
		ret.Cookie_invalid = in["cookie_invalid"].(int)
	}
	return ret
}

func dataToEndpointSlbPersistStats(d *schema.ResourceData) edpt.SlbPersistStats {
	var ret edpt.SlbPersistStats

	ret.Stats = getObjectSlbPersistStatsStats(d.Get("stats").([]interface{}))
	return ret
}
