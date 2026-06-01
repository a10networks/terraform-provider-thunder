package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpv6AddressVni() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ipv6_address_vni`: Virtual Segment Id configured on the remote VTEP\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpv6AddressVniCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpv6AddressVniUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpv6AddressVniRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpv6AddressVniDelete,

		Schema: map[string]*schema.Schema{
			"segment": {
				Type: schema.TypeInt, Required: true, Description: "VNI configured for the remote VTEP",
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
func resourceOverlayTunnelVtepRemoteIpv6AddressVniCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressVniCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressVni(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressVniRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressVniUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressVniUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressVni(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressVniRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpv6AddressVniDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressVniDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressVni(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressVniRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressVniRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressVni(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpv6AddressVni(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpv6AddressVni {
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressVni
	ret.Inst.Segment = d.Get("segment").(int)
	//omit uuid
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
