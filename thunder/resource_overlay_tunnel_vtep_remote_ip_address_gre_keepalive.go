package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpAddressGreKeepalive() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ip_address_gre_keepalive`: Specify the gre keepalive\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveDelete,

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
			"ip_address": {
				Type: schema.TypeString, Required: true, Description: "IpAddress",
			},
		},
	}
}
func resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressGreKeepalive(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressGreKeepalive(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressGreKeepalive(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressGreKeepaliveRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressGreKeepalive(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpAddressGreKeepalive(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpAddressGreKeepalive {
	var ret edpt.OverlayTunnelVtepRemoteIpAddressGreKeepalive
	ret.Inst.RetryCount = d.Get("retry_count").(int)
	ret.Inst.RetryTime = d.Get("retry_time").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	return ret
}
