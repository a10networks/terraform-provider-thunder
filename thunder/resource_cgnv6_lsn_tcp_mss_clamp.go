package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnTcpMssClamp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_tcp_mss_clamp`: LSN TCP MSS Clamping\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnTcpMssClampCreate,
		UpdateContext: resourceCgnv6LsnTcpMssClampUpdate,
		ReadContext:   resourceCgnv6LsnTcpMssClampRead,
		DeleteContext: resourceCgnv6LsnTcpMssClampDelete,

		Schema: map[string]*schema.Schema{
			"min": {
				Type: schema.TypeInt, Optional: true, Default: 456, Description: "Specify the min value allowed for the TCP MSS (Specify the min value allowed for the TCP MSS (default: ((576 - 60 - 60))))",
			},
			"mss_clamp_type": {
				Type: schema.TypeString, Optional: true, Default: "none", Description: "'fixed': Specify a fixed max value for the TCP MSS; 'subtract': Specify the value to subtract from the TCP MSS; 'none': No TCP MSS clamping (default);",
			},
			"mss_subtract": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the value to subtract from the TCP MSS (default: not configured)",
			},
			"mss_value": {
				Type: schema.TypeInt, Optional: true, Description: "The max value allowed for the TCP MSS (default: not configured)},",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LsnTcpMssClampCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpMssClampCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpMssClamp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnTcpMssClampRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnTcpMssClampUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpMssClampUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpMssClamp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnTcpMssClampRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnTcpMssClampDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpMssClampDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpMssClamp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnTcpMssClampRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnTcpMssClampRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnTcpMssClamp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LsnTcpMssClamp(d *schema.ResourceData) edpt.Cgnv6LsnTcpMssClamp {
	var ret edpt.Cgnv6LsnTcpMssClamp
	ret.Inst.Min = d.Get("min").(int)
	ret.Inst.MssClampType = d.Get("mss_clamp_type").(string)
	ret.Inst.MssSubtract = d.Get("mss_subtract").(int)
	ret.Inst.MssValue = d.Get("mss_value").(int)
	//omit uuid
	return ret
}
