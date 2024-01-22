package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTcpMssClamp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_tcp_mss_clamp`: FW TCP MSS Clamping\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTcpMssClampCreate,
		UpdateContext: resourceFwTcpMssClampUpdate,
		ReadContext:   resourceFwTcpMssClampRead,
		DeleteContext: resourceFwTcpMssClampDelete,

		Schema: map[string]*schema.Schema{
			"min": {
				Type: schema.TypeInt, Optional: true, Default: 456, Description: "Specify the min value allowed for the TCP MSS (Specify the min value allowed for the TCP MSS (default: ((576 - 60 - 60))))",
			},
			"mss_clamp_type": {
				Type: schema.TypeString, Optional: true, Description: "'fixed': Specify a fixed max value for the TCP MSS; 'subtract': Specify the value to subtract from the TCP MSS;",
			},
			"mss_subtract": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the value to subtract from the TCP MSS (default: not configured)",
			},
			"mss_value": {
				Type: schema.TypeInt, Optional: true, Description: "The max value allowed for the TCP MSS (default: not configured)}",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwTcpMssClampCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpMssClampCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpMssClamp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpMssClampRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTcpMssClampUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpMssClampUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpMssClamp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTcpMssClampRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTcpMssClampDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpMssClampDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpMssClamp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTcpMssClampRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTcpMssClampRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTcpMssClamp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwTcpMssClamp(d *schema.ResourceData) edpt.FwTcpMssClamp {
	var ret edpt.FwTcpMssClamp
	ret.Inst.Min = d.Get("min").(int)
	ret.Inst.MssClampType = d.Get("mss_clamp_type").(string)
	ret.Inst.MssSubtract = d.Get("mss_subtract").(int)
	ret.Inst.MssValue = d.Get("mss_value").(int)
	//omit uuid
	return ret
}
