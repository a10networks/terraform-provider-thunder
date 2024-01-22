package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAppPerformance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_app_performance`: Application perfomance stats\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAppPerformanceCreate,
		UpdateContext: resourceSystemAppPerformanceUpdate,
		ReadContext:   resourceSystemAppPerformanceRead,
		DeleteContext: resourceSystemAppPerformanceDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-throughput-bits-per-sec': Total Throughput in bits/sec; 'l4-conns-per-sec': L4 Connections/sec; 'l7-conns-per-sec': L7 Connections/sec; 'l7-trans-per-sec': L7 Transactions/sec; 'ssl-conns-per-sec': SSL Connections/sec; 'ip-nat-conns-per-sec': IP NAT Connections/sec; 'total-new-conns-per-sec': Total New Connections Established/sec; 'total-curr-conns': Total Current Established Connections; 'l4-bandwidth': L4 Bandwidth in bits/sec; 'l7-bandwidth': L7 Bandwidth in bits/sec; 'serv-ssl-conns-per-sec': Server SSL Connections/sec; 'fw-conns-per-sec': FW Connections/sec; 'gifw-conns-per-sec': GiFW Connections/sec;",
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
func resourceSystemAppPerformanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppPerformanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppPerformance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAppPerformanceRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAppPerformanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppPerformanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppPerformance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAppPerformanceRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAppPerformanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppPerformanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppPerformance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAppPerformanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAppPerformanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAppPerformance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemAppPerformanceSamplingEnable(d []interface{}) []edpt.SystemAppPerformanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemAppPerformanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemAppPerformanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemAppPerformance(d *schema.ResourceData) edpt.SystemAppPerformance {
	var ret edpt.SystemAppPerformance
	ret.Inst.SamplingEnable = getSliceSystemAppPerformanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
