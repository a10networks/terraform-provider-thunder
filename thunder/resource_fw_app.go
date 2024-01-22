package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwApp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_app`: Show application firewall supported protocols\n\n__PLACEHOLDER__",
		CreateContext: resourceFwAppCreate,
		UpdateContext: resourceFwAppUpdate,
		ReadContext:   resourceFwAppRead,
		DeleteContext: resourceFwAppDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'dummy': Entry for a10countergen;",
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
func resourceFwAppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAppCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAppRead(ctx, d, meta)
	}
	return diags
}

func resourceFwAppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAppUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwAppRead(ctx, d, meta)
	}
	return diags
}
func resourceFwAppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAppDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwAppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwAppRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwApp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwAppSamplingEnable(d []interface{}) []edpt.FwAppSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwAppSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwAppSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwApp(d *schema.ResourceData) edpt.FwApp {
	var ret edpt.FwApp
	ret.Inst.SamplingEnable = getSliceFwAppSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
