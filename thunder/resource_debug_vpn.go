package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugVpn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_vpn`: Debug IPsec vpn\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugVpnCreate,
		UpdateContext: resourceDebugVpnUpdate,
		ReadContext:   resourceDebugVpnRead,
		DeleteContext: resourceDebugVpnDelete,

		Schema: map[string]*schema.Schema{
			"ike_gateway": {
				Type: schema.TypeString, Optional: true, Description: "Specify IKE gateway name (gateway filter name)",
			},
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Debug level (Level 1-4)",
			},
			"strict": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only record the logs that can match Ike-gateway name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugVpnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVpnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVpn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVpnRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugVpnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVpnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVpn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugVpnRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugVpnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVpnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVpn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugVpnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugVpnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugVpn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugVpn(d *schema.ResourceData) edpt.DebugVpn {
	var ret edpt.DebugVpn
	ret.Inst.IkeGateway = d.Get("ike_gateway").(string)
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.Strict = d.Get("strict").(int)
	//omit uuid
	return ret
}
