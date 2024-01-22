package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_fix`: Configure FIX Proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbFixCreate,
		UpdateContext: resourceSlbFixUpdate,
		ReadContext:   resourceSlbFixRead,
		DeleteContext: resourceSlbFixDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_proxy': Current proxy conns; 'total_proxy': Total proxy conns; 'svrsel_fail': Server selection failure; 'noroute': No route failure; 'snat_fail': Source NAT failure; 'client_err': Client fail; 'server_err': Server fail; 'insert_clientip': Insert client IP; 'default_switching': Default switching; 'sender_switching': Sender ID switching; 'target_switching': Target ID switching; 'client_tls_conn': Client TLS conn; 'server_tls_conn': Server TLS conn;",
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
func resourceSlbFixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFixRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbFixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbFixRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbFixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbFixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbFixSamplingEnable(d []interface{}) []edpt.SlbFixSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbFixSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbFixSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbFix(d *schema.ResourceData) edpt.SlbFix {
	var ret edpt.SlbFix
	ret.Inst.SamplingEnable = getSliceSlbFixSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
