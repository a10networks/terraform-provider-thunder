package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPersist() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_persist`: Configure persist\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbPersistCreate,
		UpdateContext: resourceSlbPersistUpdate,
		ReadContext:   resourceSlbPersistRead,
		DeleteContext: resourceSlbPersistDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hash_tbl_trylock_fail': Hash tbl lock fail; 'hash_tbl_create_ok': Hash tbl create ok; 'hash_tbl_create_fail': Hash tbl create fail; 'hash_tbl_free': Hash tbl free; 'hash_tbl_rst_updown': Hash tbl reset (up/down); 'hash_tbl_rst_adddel': Hash tbl reset (add/del); 'url_hash_pri': URL hash persist (pri); 'url_hash_enqueue': URL hash persist (enQ); 'url_hash_sec': URL hash persist (sec); 'url_hash_fail': URL hash persist fail; 'header_hash_pri': Header hash persist(pri); 'header_hash_enqueue': Header hash persist(enQ); 'header_hash_sec': Header hash persist(sec); 'header_hash_fail': Header hash persist fail; 'src_ip': SRC IP persist ok; 'src_ip_enqueue': SRC IP persist enqueue; 'src_ip_fail': SRC IP persist fail; 'src_ip_new_sess_cache': SRC IP new sess (cache); 'src_ip_new_sess_cache_fail': SRC IP new sess fail (c); 'src_ip_new_sess_sel': SRC IP new sess (select); 'src_ip_new_sess_sel_fail': SRC IP new sess fail (s); 'src_ip_hash_pri': SRC IP hash persist(pri); 'src_ip_hash_enqueue': SRC IP hash persist(enQ); 'src_ip_hash_sec': SRC IP hash persist(sec); 'src_ip_hash_fail': SRC IP hash persist fail; 'src_ip_enforce': Enforce higher priority; 'dst_ip': DST IP persist ok; 'dst_ip_enqueue': DST IP persist enqueue; 'dst_ip_fail': DST IP persist fail; 'dst_ip_new_sess_cache': DST IP new sess (cache); 'dst_ip_new_sess_cache_fail': DST IP new sess fail (c); 'dst_ip_new_sess_sel': DST IP new sess (select); 'dst_ip_new_sess_sel_fail': DST IP new sess fail (s); 'dst_ip_hash_pri': DST IP hash persist(pri); 'dst_ip_hash_enqueue': DST IP hash persist(enQ); 'dst_ip_hash_sec': DST IP hash persist(sec); 'dst_ip_hash_fail': DST IP hash persist fail; 'cssl_sid_not_found': Client SSL SID not found; 'cssl_sid_match': Client SSL SID match; 'cssl_sid_not_match': Client SSL SID not match; 'sssl_sid_not_found': Server SSL SID not found; 'sssl_sid_reset': Server SSL SID reset; 'sssl_sid_match': Server SSL SID match; 'sssl_sid_not_match': Server SSL SID not match; 'ssl_sid_persist_ok': SSL SID persist ok; 'ssl_sid_persist_fail': SSL SID persist fail; 'ssl_sid_session_ok': Create SSL SID ok; 'ssl_sid_session_fail': Create SSL SID fail; 'cookie_persist_ok': Cookie persist ok; 'cookie_persist_fail': Cookie persist fail; 'cookie_not_found': Persist cookie not found; 'cookie_pass_thru': Persist cookie Pass-thru; 'cookie_invalid': Invalid persist cookie;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbPersistCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPersistCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPersist(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPersistRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbPersistUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPersistUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPersist(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPersistRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbPersistDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPersistDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPersist(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbPersistRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPersistRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPersist(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbPersistSamplingEnable(d []interface{}) []edpt.SlbPersistSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbPersistSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPersistSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPersist(d *schema.ResourceData) edpt.SlbPersist {
	var ret edpt.SlbPersist
	ret.Inst.SamplingEnable = getSliceSlbPersistSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
