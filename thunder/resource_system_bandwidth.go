package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemBandwidth() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_bandwidth`: System bandwidth information\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemBandwidthCreate,
		UpdateContext: resourceSystemBandwidthUpdate,
		ReadContext:   resourceSystemBandwidthRead,
		DeleteContext: resourceSystemBandwidthDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'input-bytes-per-sec': In Bytes per second; 'output-bytes-per-sec': Out Bytes per second;",
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
func resourceSystemBandwidthCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBandwidthCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBandwidth(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemBandwidthRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemBandwidthUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBandwidthUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBandwidth(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemBandwidthRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemBandwidthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBandwidthDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBandwidth(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemBandwidthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemBandwidthRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemBandwidth(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemBandwidthSamplingEnable(d []interface{}) []edpt.SystemBandwidthSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemBandwidthSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemBandwidthSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemBandwidth(d *schema.ResourceData) edpt.SystemBandwidth {
	var ret edpt.SystemBandwidth
	ret.Inst.SamplingEnable = getSliceSystemBandwidthSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
