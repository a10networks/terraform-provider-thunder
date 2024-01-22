package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRcCacheGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_rc_cache_global`: global ram cache stats\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbRcCacheGlobalCreate,
		UpdateContext: resourceSlbRcCacheGlobalUpdate,
		ReadContext:   resourceSlbRcCacheGlobalRead,
		DeleteContext: resourceSlbRcCacheGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Cache Hits; 'miss': Cache Misses; 'bytes_served': Bytes Served; 'total_req': Total Requests; 'caching_req': Cacheable Requests; 'nc_req_header': No-cache Request; 'nc_res_header': Not cacheable; 'rv_success': Revalidation Successes; 'rv_failure': Revalidation Failures; 'ims_request': IMS Requests; 'nm_response': Responses from cache 304 Not Modified; 'rsp_type_CL': Responses from server 200 OK - Cont Len; 'rsp_type_CE': Responses from server 200 OK - Chnk Enc; 'rsp_type_304': Responses from server 304 Not Modified; 'rsp_type_other': Responses from server 200 OK - Other; 'rsp_no_compress': Responses from cache 200 OK - No Comp; 'rsp_gzip': Responses from cache 200 OK - Gzip; 'rsp_deflate': Responses from cache 200 OK - Deflate; 'rsp_other': Responses from cache Other; 'nocache_match': Policy URI nocache; 'match': Policy URI cache; 'invalidate_match': Policy URI invalidate; 'content_toobig': Policy Content Too Big; 'content_toosmall': Policy Content Too Small; 'entry_create_failures': Entry Create failures; 'mem_size': Memory Used; 'entry_num': Entry Cached; 'replaced_entry': Entry Replaced; 'aging_entry': Entry Aged Out; 'cleaned_entry': Entry Cleaned; 'rsp_type_stream': Responses from http2 server 200 OK - Stream; 'rsp_br': Responses from cache 200 OK - Brotli;",
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
func resourceSlbRcCacheGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRcCacheGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRcCacheGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbRcCacheGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbRcCacheGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRcCacheGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRcCacheGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbRcCacheGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbRcCacheGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRcCacheGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRcCacheGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbRcCacheGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRcCacheGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRcCacheGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbRcCacheGlobalSamplingEnable(d []interface{}) []edpt.SlbRcCacheGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbRcCacheGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRcCacheGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbRcCacheGlobal(d *schema.ResourceData) edpt.SlbRcCacheGlobal {
	var ret edpt.SlbRcCacheGlobal
	ret.Inst.SamplingEnable = getSliceSlbRcCacheGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
