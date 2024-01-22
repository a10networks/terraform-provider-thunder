package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHwCompress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_hw_compress`: Configure HW compression\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHwCompressCreate,
		UpdateContext: resourceSlbHwCompressUpdate,
		ReadContext:   resourceSlbHwCompressRead,
		DeleteContext: resourceSlbHwCompressDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request_count': Total request count; 'submit_count': Total submit count; 'response_count': Total response count; 'failure_count': Total failure count; 'failure_code': Last failure code; 'ring_full_count': Compression queue full; 'max_outstanding_request_count': Max queued request count; 'max_outstanding_submit_count': Max queued submit count;",
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
func resourceSlbHwCompressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHwCompressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHwCompress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHwCompressRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbHwCompressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHwCompressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHwCompress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHwCompressRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbHwCompressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHwCompressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHwCompress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHwCompressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHwCompressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHwCompress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHwCompressSamplingEnable(d []interface{}) []edpt.SlbHwCompressSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHwCompressSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHwCompressSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHwCompress(d *schema.ResourceData) edpt.SlbHwCompress {
	var ret edpt.SlbHwCompress
	ret.Inst.SamplingEnable = getSliceSlbHwCompressSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
