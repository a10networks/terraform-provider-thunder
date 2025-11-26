package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpAddressUseGreKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ip_address_use_gre_key`: Specify the gre key\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyDelete,

		Schema: map[string]*schema.Schema{
			"gre_key": {
				Type: schema.TypeInt, Optional: true, Description: "key",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
			"ip_address": {
				Type: schema.TypeString, Required: true, Description: "IpAddress",
			},
		},
	}
}
func resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseGreKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseGreKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseGreKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseGreKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseGreKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpAddressUseGreKey(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpAddressUseGreKey {
	var ret edpt.OverlayTunnelVtepRemoteIpAddressUseGreKey
	ret.Inst.GreKey = d.Get("gre_key").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	return ret
}
