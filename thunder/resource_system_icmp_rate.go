package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIcmpRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_icmp_rate`: Icmp rate limit statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIcmpRateCreate,
		UpdateContext: resourceSystemIcmpRateUpdate,
		ReadContext:   resourceSystemIcmpRateRead,
		DeleteContext: resourceSystemIcmpRateDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'over_limit_drop': Over limit drops; 'limit_intf_drop': Interfaces rate limit drops; 'limit_vserver_drop': Virtual Server rate limit drops; 'limit_total_drop': Total rate limit drops; 'lockup_time_left': Lockup time left; 'curr_rate': Current rate; 'v6_over_limit_drop': Over limit drops (v6); 'v6_limit_intf_drop': Interfaces rate limit drops (v6); 'v6_limit_vserver_drop': Virtual Server rate limit drops (v6); 'v6_limit_total_drop': Total rate limit drops (v6); 'v6_lockup_time_left': Lockup time left (v6); 'v6_curr_rate': Current rate (v6);",
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
func resourceSystemIcmpRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmpRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIcmpRateRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIcmpRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmpRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIcmpRateRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIcmpRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmpRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIcmpRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIcmpRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIcmpRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemIcmpRateSamplingEnable(d []interface{}) []edpt.SystemIcmpRateSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemIcmpRateSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemIcmpRateSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemIcmpRate(d *schema.ResourceData) edpt.SystemIcmpRate {
	var ret edpt.SystemIcmpRate
	ret.Inst.SamplingEnable = getSliceSystemIcmpRateSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
