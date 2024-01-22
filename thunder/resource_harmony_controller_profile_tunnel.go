package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfileTunnel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_profile_tunnel`: tunnel status\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerProfileTunnelCreate,
		UpdateContext: resourceHarmonyControllerProfileTunnelUpdate,
		ReadContext:   resourceHarmonyControllerProfileTunnelRead,
		DeleteContext: resourceHarmonyControllerProfileTunnelDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Tunnel Enable; 'disable': Tunnel Disable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceHarmonyControllerProfileTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileTunnelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileTunnel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileTunnelRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerProfileTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileTunnelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileTunnel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileTunnelRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerProfileTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileTunnelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileTunnel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerProfileTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileTunnelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileTunnel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHarmonyControllerProfileTunnel(d *schema.ResourceData) edpt.HarmonyControllerProfileTunnel {
	var ret edpt.HarmonyControllerProfileTunnel
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
