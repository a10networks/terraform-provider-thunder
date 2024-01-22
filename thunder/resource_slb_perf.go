package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPerf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_perf`: SLB performance stats\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbPerfCreate,
		UpdateContext: resourceSlbPerfUpdate,
		ReadContext:   resourceSlbPerfRead,
		DeleteContext: resourceSlbPerfDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-throughput-bits-per-sec': Total Throughput in bits/sec; 'l4-conns-per-sec': L4 Connections/sec; 'l7-conns-per-sec': L7 Connections/sec; 'l7-trans-per-sec': L7 Transactions/sec; 'ssl-conns-per-sec': SSL Connections/sec; 'ip-nat-conns-per-sec': IP NAT Connections/sec; 'total-new-conns-per-sec': Total New Connections Established/sec; 'total-curr-conns': Total Current Established Connections; 'l4-bandwidth': L4 Bandwidth in bits/sec; 'l7-bandwidth': L7 Bandwidth in bits/sec; 'serv-ssl-conns-per-sec': Server SSL Connections/sec; 'fw-conns-per-sec': FW Connections/sec; 'gifw-conns-per-sec': GiFW Connections/sec; 'l7-proxy-conns-per-sec': L7 Proxy Connections/sec; 'l7-proxy-trans-per-sec': L7 Proxy Transactions/sec;",
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
func resourceSlbPerfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPerfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPerf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPerfRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbPerfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPerfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPerf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPerfRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbPerfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPerfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPerf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbPerfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPerfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPerf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbPerfSamplingEnable(d []interface{}) []edpt.SlbPerfSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbPerfSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPerfSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPerf(d *schema.ResourceData) edpt.SlbPerf {
	var ret edpt.SlbPerf
	ret.Inst.SamplingEnable = getSliceSlbPerfSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
