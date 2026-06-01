package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallVrid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_vrid`: Set VRRP-A vrid for stateful firewall\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallVridCreate,
		UpdateContext: resourceCgnv6StatefulFirewallVridUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallVridRead,
		DeleteContext: resourceCgnv6StatefulFirewallVridDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_value": {
				Type: schema.TypeInt, Optional: true, Description: "Set VRRP-A vrid for stateful firewall (IPv4 and IPv6)",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallVridCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallVrid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallVridRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallVridUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallVrid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallVridRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallVridDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallVrid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallVridRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallVrid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6StatefulFirewallVrid(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallVrid {
	var ret edpt.Cgnv6StatefulFirewallVrid
	//omit uuid
	ret.Inst.VridValue = d.Get("vrid_value").(int)
	return ret
}
