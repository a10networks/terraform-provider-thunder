package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMemory() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_memory`: Configure System Parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemMemoryCreate,
		UpdateContext: resourceSystemMemoryUpdate,
		ReadContext:   resourceSystemMemoryRead,
		DeleteContext: resourceSystemMemoryDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'usage-percentage': Memory Usage percentage;",
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
func resourceSystemMemoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemory(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMemoryRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemMemoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemory(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemMemoryRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemMemoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemory(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemMemoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemMemoryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemMemory(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemMemorySamplingEnable(d []interface{}) []edpt.SystemMemorySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemMemorySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemMemorySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemMemory(d *schema.ResourceData) edpt.SystemMemory {
	var ret edpt.SystemMemory
	ret.Inst.SamplingEnable = getSliceSystemMemorySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
