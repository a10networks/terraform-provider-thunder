package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCrlSrcip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_crl_srcip`: Configure connection rate limit\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCrlSrcipCreate,
		UpdateContext: resourceSlbCrlSrcipUpdate,
		ReadContext:   resourceSlbCrlSrcipRead,
		DeleteContext: resourceSlbCrlSrcipDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'sessions_alloc': Sessions allocated; 'sessions_freed': Sessions freed; 'out_of_sessions': Out of sessions; 'too_many_sessions': Too many sessions consumed; 'called': Threshold check count; 'permitted': Honor threshold  count; 'threshold_exceed': Threshold exceeded count; 'lockout_drop': Lockout drops; 'log_msg_sent': Log messages sent;",
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
func resourceSlbCrlSrcipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCrlSrcipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCrlSrcip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCrlSrcipRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCrlSrcipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCrlSrcipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCrlSrcip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCrlSrcipRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCrlSrcipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCrlSrcipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCrlSrcip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCrlSrcipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCrlSrcipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCrlSrcip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbCrlSrcipSamplingEnable(d []interface{}) []edpt.SlbCrlSrcipSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbCrlSrcipSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbCrlSrcipSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbCrlSrcip(d *schema.ResourceData) edpt.SlbCrlSrcip {
	var ret edpt.SlbCrlSrcip
	ret.Inst.SamplingEnable = getSliceSlbCrlSrcipSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
