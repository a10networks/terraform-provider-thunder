package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSportRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_sport_rate_limit`: Configure source port rate-limit\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSportRateLimitCreate,
		UpdateContext: resourceSlbSportRateLimitUpdate,
		ReadContext:   resourceSlbSportRateLimitRead,
		DeleteContext: resourceSlbSportRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'alloc_sport': Alloc'd src port entry; 'alloc_sportip': Alloc'd src port-ip entry; 'freed_sport': Freed src port entry; 'freed_sportip': Freed src port-ip entry; 'total_drop': Total rate exceed drop; 'total_reset': Total rate exceed reset; 'total_log': Total log sent;",
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
func resourceSlbSportRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSportRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSportRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSportRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSportRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSportRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSportRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSportRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSportRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSportRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSportRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSportRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSportRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSportRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSportRateLimitSamplingEnable(d []interface{}) []edpt.SlbSportRateLimitSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSportRateLimitSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSportRateLimitSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSportRateLimit(d *schema.ResourceData) edpt.SlbSportRateLimit {
	var ret edpt.SlbSportRateLimit
	ret.Inst.SamplingEnable = getSliceSlbSportRateLimitSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
