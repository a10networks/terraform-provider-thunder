package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRpz() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_rpz`: Manage RPZ Scripts\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbRpzCreate,
		UpdateContext: resourceSlbRpzUpdate,
		ReadContext:   resourceSlbRpzRead,
		DeleteContext: resourceSlbRpzDelete,

		Schema: map[string]*schema.Schema{
			"check": {
				Type: schema.TypeString, Optional: true, Description: "Check Response Policy Zone file",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'set_bw_error': Total RPZ Set Class-list Error; 'parse_error': Total RPZ Parse Error;",
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
func resourceSlbRpzCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpz(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbRpzRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbRpzUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpz(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbRpzRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbRpzDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpz(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbRpzRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpz(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbRpzSamplingEnable(d []interface{}) []edpt.SlbRpzSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbRpzSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRpzSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbRpz(d *schema.ResourceData) edpt.SlbRpz {
	var ret edpt.SlbRpz
	ret.Inst.Check = d.Get("check").(string)
	ret.Inst.SamplingEnable = getSliceSlbRpzSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
