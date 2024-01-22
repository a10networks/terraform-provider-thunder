package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsPersistentCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_dns_persistent_cache`: DNS Persistent Cache Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbDnsPersistentCacheCreate,
		UpdateContext: resourceSlbDnsPersistentCacheUpdate,
		ReadContext:   resourceSlbDnsPersistentCacheRead,
		DeleteContext: resourceSlbDnsPersistentCacheDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_entry': Total persistent cache entry; 'entry_saved': Total saved cache entry; 'entry_deleted': Total deleted cache entry; 'database_busy': Database busy; 'database_error': Database error;",
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
func resourceSlbDnsPersistentCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsPersistentCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsPersistentCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDnsPersistentCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDnsPersistentCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsPersistentCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsPersistentCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDnsPersistentCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbDnsPersistentCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsPersistentCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsPersistentCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbDnsPersistentCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsPersistentCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsPersistentCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbDnsPersistentCacheSamplingEnable(d []interface{}) []edpt.SlbDnsPersistentCacheSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbDnsPersistentCacheSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsPersistentCacheSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDnsPersistentCache(d *schema.ResourceData) edpt.SlbDnsPersistentCache {
	var ret edpt.SlbDnsPersistentCache
	ret.Inst.SamplingEnable = getSliceSlbDnsPersistentCacheSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
