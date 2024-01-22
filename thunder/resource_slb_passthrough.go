package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPassthrough() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_passthrough`: passthrough session stats\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbPassthroughCreate,
		UpdateContext: resourceSlbPassthroughUpdate,
		ReadContext:   resourceSlbPassthroughRead,
		DeleteContext: resourceSlbPassthroughDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_conn': Current connections; 'total_conn': Total connections; 'total_fwd_bytes': Forward bytes; 'total_fwd_packets': Forward packets; 'total_rev_bytes': Reverse bytes; 'total_rev_packets': Reverse packets; 'curr_pconn': Persistent connections;",
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
func resourceSlbPassthroughCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPassthroughCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPassthrough(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPassthroughRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbPassthroughUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPassthroughUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPassthrough(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbPassthroughRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbPassthroughDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPassthroughDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPassthrough(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbPassthroughRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPassthroughRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPassthrough(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbPassthroughSamplingEnable(d []interface{}) []edpt.SlbPassthroughSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbPassthroughSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPassthroughSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPassthrough(d *schema.ResourceData) edpt.SlbPassthrough {
	var ret edpt.SlbPassthrough
	ret.Inst.SamplingEnable = getSliceSlbPassthroughSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
