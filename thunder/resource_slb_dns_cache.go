package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsCache() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_dns_cache`: DNS Cache Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbDnsCacheCreate,
		UpdateContext: resourceSlbDnsCacheUpdate,
		ReadContext:   resourceSlbDnsCacheRead,
		DeleteContext: resourceSlbDnsCacheDelete,

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
func resourceSlbDnsCacheCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsCacheCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsCache(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDnsCacheRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDnsCacheUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsCacheUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsCache(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDnsCacheRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbDnsCacheDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsCacheDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsCache(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbDnsCacheRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsCacheRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsCache(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbDnsCacheSamplingEnable(d []interface{}) []edpt.SlbDnsCacheSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbDnsCacheSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsCacheSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDnsCache(d *schema.ResourceData) edpt.SlbDnsCache {
	var ret edpt.SlbDnsCache
	ret.Inst.SamplingEnable = getSliceSlbDnsCacheSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
