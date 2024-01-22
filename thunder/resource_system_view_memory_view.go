package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewMemoryView() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_view_memory_view`: Configure System Parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemViewMemoryViewCreate,
		UpdateContext: resourceSystemViewMemoryViewUpdate,
		ReadContext:   resourceSystemViewMemoryViewRead,
		DeleteContext: resourceSystemViewMemoryViewDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'usage-percentage': Usage percentage;",
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
func resourceSystemViewMemoryViewCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewMemoryViewCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewMemoryView(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemViewMemoryViewRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemViewMemoryViewUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewMemoryViewUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewMemoryView(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemViewMemoryViewRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemViewMemoryViewDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewMemoryViewDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewMemoryView(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemViewMemoryViewRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewMemoryViewRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewMemoryView(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemViewMemoryViewSamplingEnable(d []interface{}) []edpt.SystemViewMemoryViewSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemViewMemoryViewSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewMemoryViewSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemViewMemoryView(d *schema.ResourceData) edpt.SystemViewMemoryView {
	var ret edpt.SystemViewMemoryView
	ret.Inst.SamplingEnable = getSliceSystemViewMemoryViewSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
