package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsNegativeDnsCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_negative_dns_cache`: setting the nagative dns cache\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsNegativeDnsCacheCreate,
		UpdateContext: resourceSlbTemplateDnsNegativeDnsCacheUpdate,
		ReadContext:   resourceSlbTemplateDnsNegativeDnsCacheRead,
		DeleteContext: resourceSlbTemplateDnsNegativeDnsCacheDelete,

		Schema: map[string]*schema.Schema{
			"bypass_query_threshold": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "the threshold bypass the query, default is 100",
			},
			"enable_negative_dns_cache": {
				Type: schema.TypeInt, Required: true, Description: "Enable DNS negative cache (Need to turn-on the dns-cache for this feature)",
			},
			"max_negative_cache_ttl": {
				Type: schema.TypeInt, Optional: true, Default: 7200, Description: "Max negative cache ttl, default is 2 hours",
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
func resourceSlbTemplateDnsNegativeDnsCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsNegativeDnsCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsNegativeDnsCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsNegativeDnsCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsNegativeDnsCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsNegativeDnsCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsNegativeDnsCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsNegativeDnsCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsNegativeDnsCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsNegativeDnsCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsNegativeDnsCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsNegativeDnsCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsNegativeDnsCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsNegativeDnsCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsNegativeDnsCache(d *schema.ResourceData) edpt.SlbTemplateDnsNegativeDnsCache {
	var ret edpt.SlbTemplateDnsNegativeDnsCache
	ret.Inst.BypassQueryThreshold = d.Get("bypass_query_threshold").(int)
	ret.Inst.EnableNegativeDnsCache = d.Get("enable_negative_dns_cache").(int)
	ret.Inst.MaxNegativeCacheTtl = d.Get("max_negative_cache_ttl").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
