package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDnsCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_dns_cache`: DNS Cache Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemDnsCacheCreate,
		UpdateContext: resourceSystemDnsCacheUpdate,
		ReadContext:   resourceSystemDnsCacheRead,
		DeleteContext: resourceSystemDnsCacheDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_q': Total query; 'total_r': Total server response; 'hit': Total cache hit; 'bad_q': Query not passed; 'encode_q': Query encoded; 'multiple_q': Query with multiple questions; 'oversize_q': Query exceed cache size; 'bad_r': Response not passed; 'oversize_r': Response exceed cache size; 'encode_r': Response encoded; 'multiple_r': Response with multiple questions; 'answer_r': Response with multiple answers; 'ttl_r': Response with short TTL; 'ageout': Total aged out; 'bad_answer': Bad Answer; 'ageout_weight': Total aged for lower weight; 'total_log': Total stats log sent; 'total_alloc': Total allocated; 'total_freed': Total freed; 'current_allocate': Current allocate; 'current_data_allocate': Current data allocate; 'resolver_queue_full': Resolver task queue full; 'truncated_r': Response with Truncation bit set;",
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
func resourceSystemDnsCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDnsCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemDnsCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDnsCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemDnsCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemDnsCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemDnsCacheSamplingEnable(d []interface{}) []edpt.SystemDnsCacheSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemDnsCacheSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsCacheSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemDnsCache(d *schema.ResourceData) edpt.SystemDnsCache {
	var ret edpt.SystemDnsCache
	ret.Inst.SamplingEnable = getSliceSystemDnsCacheSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
