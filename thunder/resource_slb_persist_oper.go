package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPersistOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_persist_oper`: Operational Status for the object persist\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPersistOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"persist_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hash_tbl_trylock_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash_tbl_create_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash_tbl_create_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash_tbl_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash_tbl_rst_updown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash_tbl_rst_adddel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_hash_pri": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_hash_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_hash_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"url_hash_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_hash_pri": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_hash_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_hash_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_hash_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_new_sess_cache": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_new_sess_cache_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_new_sess_sel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_new_sess_sel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_hash_pri": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_hash_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_hash_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_hash_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"src_ip_enforce": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_new_sess_cache": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_new_sess_cache_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_new_sess_sel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_new_sess_sel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_hash_pri": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_hash_enqueue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_hash_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_ip_hash_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cssl_sid_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cssl_sid_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cssl_sid_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sssl_sid_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sssl_sid_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sssl_sid_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sssl_sid_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_sid_persist_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_sid_persist_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_sid_session_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_sid_session_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_persist_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_persist_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_pass_thru": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cookie_invalid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbPersistOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPersistOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPersistOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPersistOperOper := setObjectSlbPersistOperOper(res)
		d.Set("oper", SlbPersistOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPersistOperOper(ret edpt.DataSlbPersistOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"persist_cpu_list": setSliceSlbPersistOperOperPersistCpuList(ret.DtSlbPersistOper.Oper.PersistCpuList),
			"cpu_count":        ret.DtSlbPersistOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbPersistOperOperPersistCpuList(d []edpt.SlbPersistOperOperPersistCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["hash_tbl_trylock_fail"] = item.Hash_tbl_trylock_fail
		in["hash_tbl_create_ok"] = item.Hash_tbl_create_ok
		in["hash_tbl_create_fail"] = item.Hash_tbl_create_fail
		in["hash_tbl_free"] = item.Hash_tbl_free
		in["hash_tbl_rst_updown"] = item.Hash_tbl_rst_updown
		in["hash_tbl_rst_adddel"] = item.Hash_tbl_rst_adddel
		in["url_hash_pri"] = item.Url_hash_pri
		in["url_hash_enqueue"] = item.Url_hash_enqueue
		in["url_hash_sec"] = item.Url_hash_sec
		in["url_hash_fail"] = item.Url_hash_fail
		in["header_hash_pri"] = item.Header_hash_pri
		in["header_hash_enqueue"] = item.Header_hash_enqueue
		in["header_hash_sec"] = item.Header_hash_sec
		in["header_hash_fail"] = item.Header_hash_fail
		in["src_ip"] = item.Src_ip
		in["src_ip_enqueue"] = item.Src_ip_enqueue
		in["src_ip_fail"] = item.Src_ip_fail
		in["src_ip_new_sess_cache"] = item.Src_ip_new_sess_cache
		in["src_ip_new_sess_cache_fail"] = item.Src_ip_new_sess_cache_fail
		in["src_ip_new_sess_sel"] = item.Src_ip_new_sess_sel
		in["src_ip_new_sess_sel_fail"] = item.Src_ip_new_sess_sel_fail
		in["src_ip_hash_pri"] = item.Src_ip_hash_pri
		in["src_ip_hash_enqueue"] = item.Src_ip_hash_enqueue
		in["src_ip_hash_sec"] = item.Src_ip_hash_sec
		in["src_ip_hash_fail"] = item.Src_ip_hash_fail
		in["src_ip_enforce"] = item.Src_ip_enforce
		in["dst_ip"] = item.Dst_ip
		in["dst_ip_enqueue"] = item.Dst_ip_enqueue
		in["dst_ip_fail"] = item.Dst_ip_fail
		in["dst_ip_new_sess_cache"] = item.Dst_ip_new_sess_cache
		in["dst_ip_new_sess_cache_fail"] = item.Dst_ip_new_sess_cache_fail
		in["dst_ip_new_sess_sel"] = item.Dst_ip_new_sess_sel
		in["dst_ip_new_sess_sel_fail"] = item.Dst_ip_new_sess_sel_fail
		in["dst_ip_hash_pri"] = item.Dst_ip_hash_pri
		in["dst_ip_hash_enqueue"] = item.Dst_ip_hash_enqueue
		in["dst_ip_hash_sec"] = item.Dst_ip_hash_sec
		in["dst_ip_hash_fail"] = item.Dst_ip_hash_fail
		in["cssl_sid_not_found"] = item.Cssl_sid_not_found
		in["cssl_sid_match"] = item.Cssl_sid_match
		in["cssl_sid_not_match"] = item.Cssl_sid_not_match
		in["sssl_sid_not_found"] = item.Sssl_sid_not_found
		in["sssl_sid_reset"] = item.Sssl_sid_reset
		in["sssl_sid_match"] = item.Sssl_sid_match
		in["sssl_sid_not_match"] = item.Sssl_sid_not_match
		in["ssl_sid_persist_ok"] = item.Ssl_sid_persist_ok
		in["ssl_sid_persist_fail"] = item.Ssl_sid_persist_fail
		in["ssl_sid_session_ok"] = item.Ssl_sid_session_ok
		in["ssl_sid_session_fail"] = item.Ssl_sid_session_fail
		in["cookie_persist_ok"] = item.Cookie_persist_ok
		in["cookie_persist_fail"] = item.Cookie_persist_fail
		in["cookie_not_found"] = item.Cookie_not_found
		in["cookie_pass_thru"] = item.Cookie_pass_thru
		in["cookie_invalid"] = item.Cookie_invalid
		result = append(result, in)
	}
	return result
}

func getObjectSlbPersistOperOper(d []interface{}) edpt.SlbPersistOperOper {

	count1 := len(d)
	var ret edpt.SlbPersistOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PersistCpuList = getSliceSlbPersistOperOperPersistCpuList(in["persist_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbPersistOperOperPersistCpuList(d []interface{}) []edpt.SlbPersistOperOperPersistCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbPersistOperOperPersistCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPersistOperOperPersistCpuList
		oi.Hash_tbl_trylock_fail = in["hash_tbl_trylock_fail"].(int)
		oi.Hash_tbl_create_ok = in["hash_tbl_create_ok"].(int)
		oi.Hash_tbl_create_fail = in["hash_tbl_create_fail"].(int)
		oi.Hash_tbl_free = in["hash_tbl_free"].(int)
		oi.Hash_tbl_rst_updown = in["hash_tbl_rst_updown"].(int)
		oi.Hash_tbl_rst_adddel = in["hash_tbl_rst_adddel"].(int)
		oi.Url_hash_pri = in["url_hash_pri"].(int)
		oi.Url_hash_enqueue = in["url_hash_enqueue"].(int)
		oi.Url_hash_sec = in["url_hash_sec"].(int)
		oi.Url_hash_fail = in["url_hash_fail"].(int)
		oi.Header_hash_pri = in["header_hash_pri"].(int)
		oi.Header_hash_enqueue = in["header_hash_enqueue"].(int)
		oi.Header_hash_sec = in["header_hash_sec"].(int)
		oi.Header_hash_fail = in["header_hash_fail"].(int)
		oi.Src_ip = in["src_ip"].(int)
		oi.Src_ip_enqueue = in["src_ip_enqueue"].(int)
		oi.Src_ip_fail = in["src_ip_fail"].(int)
		oi.Src_ip_new_sess_cache = in["src_ip_new_sess_cache"].(int)
		oi.Src_ip_new_sess_cache_fail = in["src_ip_new_sess_cache_fail"].(int)
		oi.Src_ip_new_sess_sel = in["src_ip_new_sess_sel"].(int)
		oi.Src_ip_new_sess_sel_fail = in["src_ip_new_sess_sel_fail"].(int)
		oi.Src_ip_hash_pri = in["src_ip_hash_pri"].(int)
		oi.Src_ip_hash_enqueue = in["src_ip_hash_enqueue"].(int)
		oi.Src_ip_hash_sec = in["src_ip_hash_sec"].(int)
		oi.Src_ip_hash_fail = in["src_ip_hash_fail"].(int)
		oi.Src_ip_enforce = in["src_ip_enforce"].(int)
		oi.Dst_ip = in["dst_ip"].(int)
		oi.Dst_ip_enqueue = in["dst_ip_enqueue"].(int)
		oi.Dst_ip_fail = in["dst_ip_fail"].(int)
		oi.Dst_ip_new_sess_cache = in["dst_ip_new_sess_cache"].(int)
		oi.Dst_ip_new_sess_cache_fail = in["dst_ip_new_sess_cache_fail"].(int)
		oi.Dst_ip_new_sess_sel = in["dst_ip_new_sess_sel"].(int)
		oi.Dst_ip_new_sess_sel_fail = in["dst_ip_new_sess_sel_fail"].(int)
		oi.Dst_ip_hash_pri = in["dst_ip_hash_pri"].(int)
		oi.Dst_ip_hash_enqueue = in["dst_ip_hash_enqueue"].(int)
		oi.Dst_ip_hash_sec = in["dst_ip_hash_sec"].(int)
		oi.Dst_ip_hash_fail = in["dst_ip_hash_fail"].(int)
		oi.Cssl_sid_not_found = in["cssl_sid_not_found"].(int)
		oi.Cssl_sid_match = in["cssl_sid_match"].(int)
		oi.Cssl_sid_not_match = in["cssl_sid_not_match"].(int)
		oi.Sssl_sid_not_found = in["sssl_sid_not_found"].(int)
		oi.Sssl_sid_reset = in["sssl_sid_reset"].(int)
		oi.Sssl_sid_match = in["sssl_sid_match"].(int)
		oi.Sssl_sid_not_match = in["sssl_sid_not_match"].(int)
		oi.Ssl_sid_persist_ok = in["ssl_sid_persist_ok"].(int)
		oi.Ssl_sid_persist_fail = in["ssl_sid_persist_fail"].(int)
		oi.Ssl_sid_session_ok = in["ssl_sid_session_ok"].(int)
		oi.Ssl_sid_session_fail = in["ssl_sid_session_fail"].(int)
		oi.Cookie_persist_ok = in["cookie_persist_ok"].(int)
		oi.Cookie_persist_fail = in["cookie_persist_fail"].(int)
		oi.Cookie_not_found = in["cookie_not_found"].(int)
		oi.Cookie_pass_thru = in["cookie_pass_thru"].(int)
		oi.Cookie_invalid = in["cookie_invalid"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPersistOper(d *schema.ResourceData) edpt.SlbPersistOper {
	var ret edpt.SlbPersistOper

	ret.Oper = getObjectSlbPersistOperOper(d.Get("oper").([]interface{}))
	return ret
}
