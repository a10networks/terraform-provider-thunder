package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemThroughput() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_throughput`: System throughput\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemThroughputCreate,
		UpdateContext: resourceSystemThroughputUpdate,
		ReadContext:   resourceSystemThroughputRead,
		DeleteContext: resourceSystemThroughputDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'global-system-throughput-bits-per-sec': Global System throughput in bits/sec; 'per-part-throughput-bits-per-sec': Partition throughput in bits/sec;",
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
func resourceSystemThroughputCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemThroughputCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemThroughput(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemThroughputRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemThroughputUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemThroughputUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemThroughput(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemThroughputRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemThroughputDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemThroughputDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemThroughput(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemThroughputRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemThroughputRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemThroughput(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemThroughputSamplingEnable(d []interface{}) []edpt.SystemThroughputSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemThroughputSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemThroughputSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemThroughput(d *schema.ResourceData) edpt.SystemThroughput {
	var ret edpt.SystemThroughput
	ret.Inst.SamplingEnable = getSliceSystemThroughputSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
