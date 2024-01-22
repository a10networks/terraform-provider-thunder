package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSecureGaming() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_secure_gaming`: SLB Secure Gaming info\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSecureGamingCreate,
		UpdateContext: resourceSlbSecureGamingUpdate,
		ReadContext:   resourceSlbSecureGamingRead,
		DeleteContext: resourceSlbSecureGamingDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'secure_gaming_pass': Secure Gaming passed; 'secure_gaming_drop': Secure Gaming dropped;",
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
func resourceSlbSecureGamingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSecureGamingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSecureGaming(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSecureGamingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSecureGamingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSecureGamingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSecureGaming(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSecureGamingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSecureGamingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSecureGamingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSecureGaming(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSecureGamingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSecureGamingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSecureGaming(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSecureGamingSamplingEnable(d []interface{}) []edpt.SlbSecureGamingSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSecureGamingSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSecureGamingSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSecureGaming(d *schema.ResourceData) edpt.SlbSecureGaming {
	var ret edpt.SlbSecureGaming
	ret.Inst.SamplingEnable = getSliceSlbSecureGamingSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
