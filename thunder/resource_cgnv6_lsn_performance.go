package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPerformance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_performance`: Large-Scale NAT performance statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnPerformanceCreate,
		UpdateContext: resourceCgnv6LsnPerformanceUpdate,
		ReadContext:   resourceCgnv6LsnPerformanceRead,
		DeleteContext: resourceCgnv6LsnPerformanceDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'data-sessions-current-epoch': data-sessions-current-epoch; 'fullcone-created-current-epoch': fullcone-created-current-epoch; 'user-quote-created-current-epoch': user-quote-created-current-epoch; 'data-sessions-previous-epoch-first': data-sessions-previous-epoch-first; 'fullcone-created-previous-epoch-first': fullcone-created-previous-epoch-first; 'user-quote-created-previous-epoch-first': user-quote-created-previous-epoch-first; 'data-sessions-previous-epoch-last': data-sessions-previous-epoch-last; 'fullcone-created-previous-epoch-last': fullcone-created-previous-epoch-last; 'user-quote-created-previous-epoch-last': user-quote-created-previous-epoch-last;",
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
func resourceCgnv6LsnPerformanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerformanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerformance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPerformanceRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnPerformanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerformanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerformance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnPerformanceRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnPerformanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerformanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerformance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnPerformanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerformanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerformance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnPerformanceSamplingEnable(d []interface{}) []edpt.Cgnv6LsnPerformanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnPerformanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnPerformanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnPerformance(d *schema.ResourceData) edpt.Cgnv6LsnPerformance {
	var ret edpt.Cgnv6LsnPerformance
	ret.Inst.SamplingEnable = getSliceCgnv6LsnPerformanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
