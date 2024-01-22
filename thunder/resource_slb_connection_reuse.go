package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbConnectionReuse() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_connection_reuse`: Configure Connection Reuse\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbConnectionReuseCreate,
		UpdateContext: resourceSlbConnectionReuseUpdate,
		ReadContext:   resourceSlbConnectionReuseRead,
		DeleteContext: resourceSlbConnectionReuseDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'current_open': Open persist; 'current_active': Active persist; 'nbind': Total bind; 'nunbind': Total unbind; 'nestab': Total established; 'ntermi': Total terminated; 'ntermi_err': Total terminated by err; 'delay_unbind': Delayed unbind; 'long_resp': Long resp; 'miss_resp': Missed resp; 'unbound_data_rcv': Unbound data rcvd; 'pause_conn': Pause request; 'pause_conn_fail': Pause request fail; 'resume_conn': Resume request; 'not_remove_from_rport': Not remove from list;",
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
func resourceSlbConnectionReuseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbConnectionReuseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbConnectionReuse(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbConnectionReuseRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbConnectionReuseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbConnectionReuseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbConnectionReuse(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbConnectionReuseRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbConnectionReuseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbConnectionReuseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbConnectionReuse(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbConnectionReuseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbConnectionReuseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbConnectionReuse(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbConnectionReuseSamplingEnable(d []interface{}) []edpt.SlbConnectionReuseSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbConnectionReuseSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbConnectionReuseSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbConnectionReuse(d *schema.ResourceData) edpt.SlbConnectionReuse {
	var ret edpt.SlbConnectionReuse
	ret.Inst.SamplingEnable = getSliceSlbConnectionReuseSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
