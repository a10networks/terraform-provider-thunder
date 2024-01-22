package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Global() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_global`: CGN global paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6GlobalCreate,
		UpdateContext: resourceCgnv6GlobalUpdate,
		ReadContext:   resourceCgnv6GlobalRead,
		DeleteContext: resourceCgnv6GlobalDelete,

		Schema: map[string]*schema.Schema{
			"ping_sweep_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable ping sweep detection; 'disable': Disable ping sweep detection(default);",
			},
			"port_scan_detection": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable port scan detection; 'disable': Disable port scan detection(default);",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'tcp-total-ports-allocated': Total TCP ports allocated; 'udp-total-ports-allocated': Total UDP ports allocated; 'icmp-total-ports-allocated': Total ICMP ports allocated;",
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
func resourceCgnv6GlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Global(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6GlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6GlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Global(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6GlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6GlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Global(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6GlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6GlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Global(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6GlobalSamplingEnable(d []interface{}) []edpt.Cgnv6GlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6GlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6GlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Global(d *schema.ResourceData) edpt.Cgnv6Global {
	var ret edpt.Cgnv6Global
	ret.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	ret.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6GlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
