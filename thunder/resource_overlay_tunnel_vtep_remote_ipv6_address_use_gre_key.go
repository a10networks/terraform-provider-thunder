package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ipv6_address_use_gre_key`: Specify the gre key\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyDelete,

		Schema: map[string]*schema.Schema{
			"gre_key": {
				Type: schema.TypeInt, Optional: true, Description: "key",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "Ipv6Address",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
		},
	}
}
func resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseGreKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseGreKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseGreKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseGreKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseGreKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseGreKey(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpv6AddressUseGreKey {
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressUseGreKey
	ret.Inst.GreKey = d.Get("gre_key").(int)
	//omit uuid
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
