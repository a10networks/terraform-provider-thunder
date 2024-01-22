package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpWindowCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_tcp_window_check`: Configure TCP window check\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTcpWindowCheckCreate,
		UpdateContext: resourceFwTcpWindowCheckUpdate,
		ReadContext:   resourceFwTcpWindowCheckRead,
		DeleteContext: resourceFwTcpWindowCheckDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'outside-window': packet dropped counter for outside of tcp window;",
						},
					},
				},
			},
			"status": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable TCP window check (default); 'disable': Disable TCP window check;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwTcpWindowCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpWindowCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpWindowCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpWindowCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTcpWindowCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpWindowCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpWindowCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpWindowCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTcpWindowCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpWindowCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpWindowCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTcpWindowCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpWindowCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpWindowCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTcpWindowCheckSamplingEnable(d []interface{}) []edpt.FwTcpWindowCheckSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwTcpWindowCheckSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTcpWindowCheckSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTcpWindowCheck(d *schema.ResourceData) edpt.FwTcpWindowCheck {
	var ret edpt.FwTcpWindowCheck
	ret.Inst.SamplingEnable = getSliceFwTcpWindowCheckSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
