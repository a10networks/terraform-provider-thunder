package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDns64() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_dns64`: DNS64 Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbDns64Create,
		UpdateContext: resourceSlbDns64Update,
		ReadContext:   resourceSlbDns64Read,
		DeleteContext: resourceSlbDns64Delete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'query': Query; 'query-bad-pkt': Query Bad Packet; 'query-chg': Query Changed; 'query-parallel': Query Parallel; 'query-passive': Query Passive; 'resp': Response; 'resp-bad-pkt': Response Bad Packet; 'resp-bad-qr': Response Bad Query; 'resp-chg': Response Changed; 'resp-err': Response Error; 'resp-empty': Response Empty; 'resp-local': Response Local; 'adjust': Translated; 'cache': Cache; 'drop': Dropped;",
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
func resourceSlbDns64Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDns64Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDns64(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDns64Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbDns64Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDns64Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDns64(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDns64Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbDns64Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDns64Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDns64(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbDns64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDns64Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDns64(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbDns64SamplingEnable(d []interface{}) []edpt.SlbDns64SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbDns64SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDns64SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDns64(d *schema.ResourceData) edpt.SlbDns64 {
	var ret edpt.SlbDns64
	ret.Inst.SamplingEnable = getSliceSlbDns64SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
