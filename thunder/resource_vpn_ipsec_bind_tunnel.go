package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsecBindTunnel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn_ipsec_bind_tunnel`: Binds tunnel interface to the IPsec connection\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnIpsecBindTunnelCreate,
		UpdateContext: resourceVpnIpsecBindTunnelUpdate,
		ReadContext:   resourceVpnIpsecBindTunnelRead,
		DeleteContext: resourceVpnIpsecBindTunnelDelete,

		Schema: map[string]*schema.Schema{
			"next_hop": {
				Type: schema.TypeString, Optional: true, Description: "IPsec Next Hop IP Address",
			},
			"next_hop_v6": {
				Type: schema.TypeString, Optional: true, Description: "IPsec Next Hop IPv6 Address",
			},
			"tunnel": {
				Type: schema.TypeInt, Optional: true, Description: "Tunnel interface index",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVpnIpsecBindTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecBindTunnelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecBindTunnel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecBindTunnelRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnIpsecBindTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecBindTunnelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecBindTunnel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecBindTunnelRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnIpsecBindTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecBindTunnelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecBindTunnel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnIpsecBindTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecBindTunnelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsecBindTunnel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVpnIpsecBindTunnel(d *schema.ResourceData) edpt.VpnIpsecBindTunnel {
	var ret edpt.VpnIpsecBindTunnel
	ret.Inst.NextHop = d.Get("next_hop").(string)
	ret.Inst.NextHopV6 = d.Get("next_hop_v6").(string)
	ret.Inst.Tunnel = d.Get("tunnel").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
