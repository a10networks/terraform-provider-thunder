package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepalive() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ipv6_address_gre_keepalive`: Specify the gre keepalive\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveDelete,

		Schema: map[string]*schema.Schema{
			"retry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Keepalive multiplier",
			},
			"retry_time": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Keepalive retry interval in seconds",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "Ipv6Address",
			},
		},
	}
}
func resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressGreKeepalive(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressGreKeepalive(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressGreKeepalive(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressGreKeepaliveRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressGreKeepalive(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpv6AddressGreKeepalive(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpv6AddressGreKeepalive {
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressGreKeepalive
	ret.Inst.RetryCount = d.Get("retry_count").(int)
	ret.Inst.RetryTime = d.Get("retry_time").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	return ret
}
