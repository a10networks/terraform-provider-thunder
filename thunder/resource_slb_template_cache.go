package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_cache`: RAM caching template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateCacheCreate,
		UpdateContext: resourceSlbTemplateCacheUpdate,
		ReadContext:   resourceSlbTemplateCacheRead,
		DeleteContext: resourceSlbTemplateCacheDelete,

		Schema: map[string]*schema.Schema{
			"accept_reload_req": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept reload requests via cache-control directives in HTTP headers",
			},
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 3600, Description: "Specify duration in seconds cached content valid, default is 3600 seconds (seconds that the cached content is valid (default 3600 seconds))",
			},
			"default_policy_nocache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify default policy to be to not cache",
			},
			"disable_insert_age": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable insertion of age header in response served from RAM cache",
			},
			"disable_insert_via": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable insertion of via header in response served from RAM cache",
			},
			"local_uri_policy": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_uri": {
							Type: schema.TypeString, Optional: true, Description: "Specify Local URI for caching (Specify URI pattern that the policy should be applied to, maximum 63 charaters)",
						},
					},
				},
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Description: "Specify logging template (Logging Config name)",
			},
			"max_cache_size": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Specify maximum cache size in megabytes, default is 80MB (RAM cache size in megabytes (default 80MB))",
			},
			"max_content_size": {
				Type: schema.TypeInt, Optional: true, Default: 81920, Description: "Maximum size (bytes) of response that can be cached - default 81920 (80KB)",
			},
			"min_content_size": {
				Type: schema.TypeInt, Optional: true, Default: 512, Description: "Minimum size (bytes) of response that can be cached - default 512",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify cache template name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"remove_cookies": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remove cookies in response and cache",
			},
			"replacement_policy": {
				Type: schema.TypeString, Optional: true, Default: "LFU", Description: "'LFU': LFU;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Cache hits; 'miss': Cache misses; 'bytes_served': Bytes served from cache; 'total_req': Total requests received; 'caching_req': Total requests to cache; 'nc_req_header': slbTemplateCacheNcReqHeader, help nc_req_header; 'nc_res_header': slbTemplateCacheNcResHeader, help nc_res_header; 'rv_success': rv_success; 'rv_failure': slbTemplateCacheRvFailure, help rv_failure; 'ims_request': ims_request; 'nm_response': nm_response; 'rsp_type_CL': rsp_type_CL; 'rsp_type_CE': rsp_type_CE; 'rsp_type_304': rsp_type_304; 'rsp_type_other': rsp_type_other; 'rsp_no_compress': rsp_no_compress; 'rsp_gzip': rsp_gzip; 'rsp_deflate': rsp_deflate; 'rsp_other': rsp_other; 'nocache_match': nocache_match; 'match': match; 'invalidate_match': invalidate_match; 'content_toobig': slbTemplateCacheContentToobig, help content_toobig; 'content_toosmall': slbTemplateCacheContentToosmall, help content_toosmall; 'entry_create_failures': slbTemplateCacheEntryCreateFailures, help entry_create_failures; 'mem_size': mem_size; 'entry_num': entry_num; 'replaced_entry': replaced_entry; 'aging_entry': aging_entry; 'cleaned_entry': cleaned_entry; 'rsp_type_stream': rsp_type_stream; 'header_save_error': header_save_error; 'rsp_br': rsp_br;",
						},
					},
				},
			},
			"uri_policy": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type: schema.TypeString, Optional: true, Description: "Specify URI for cache policy (Specify URI pattern that the policy should be applied to, maximum 63 charaters)",
						},
						"cache_action": {
							Type: schema.TypeString, Optional: true, Description: "'cache': Specify if certain URIs should be cached; 'nocache': Specify if certain URIs should not be cached;",
						},
						"cache_value": {
							Type: schema.TypeInt, Optional: true, Description: "Specify seconds that content should be cached, default is age specified in cache template",
						},
						"invalidate": {
							Type: schema.TypeString, Optional: true, Description: "Specify if URI should invalidate cache entries matching pattern (pattern that would match entries to be invalidated (64 chars max))",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"verify_host": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Verify request using host before sending response from RAM cache",
			},
		},
	}
}
func resourceSlbTemplateCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateCacheLocalUriPolicy(d []interface{}) []edpt.SlbTemplateCacheLocalUriPolicy {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateCacheLocalUriPolicy, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCacheLocalUriPolicy
		oi.LocalUri = in["local_uri"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateCacheSamplingEnable(d []interface{}) []edpt.SlbTemplateCacheSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateCacheSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCacheSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateCacheUriPolicy(d []interface{}) []edpt.SlbTemplateCacheUriPolicy {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateCacheUriPolicy, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCacheUriPolicy
		oi.Uri = in["uri"].(string)
		oi.CacheAction = in["cache_action"].(string)
		oi.CacheValue = in["cache_value"].(int)
		oi.Invalidate = in["invalidate"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateCache(d *schema.ResourceData) edpt.SlbTemplateCache {
	var ret edpt.SlbTemplateCache
	ret.Inst.AcceptReloadReq = d.Get("accept_reload_req").(int)
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.DefaultPolicyNocache = d.Get("default_policy_nocache").(int)
	ret.Inst.DisableInsertAge = d.Get("disable_insert_age").(int)
	ret.Inst.DisableInsertVia = d.Get("disable_insert_via").(int)
	ret.Inst.LocalUriPolicy = getSliceSlbTemplateCacheLocalUriPolicy(d.Get("local_uri_policy").([]interface{}))
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.MaxCacheSize = d.Get("max_cache_size").(int)
	ret.Inst.MaxContentSize = d.Get("max_content_size").(int)
	ret.Inst.MinContentSize = d.Get("min_content_size").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.RemoveCookies = d.Get("remove_cookies").(int)
	ret.Inst.ReplacementPolicy = d.Get("replacement_policy").(string)
	ret.Inst.SamplingEnable = getSliceSlbTemplateCacheSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UriPolicy = getSliceSlbTemplateCacheUriPolicy(d.Get("uri_policy").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VerifyHost = d.Get("verify_host").(int)
	return ret
}
