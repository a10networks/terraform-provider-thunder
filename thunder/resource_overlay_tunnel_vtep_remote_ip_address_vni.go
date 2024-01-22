package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpAddressVni() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ip_address_vni`: Virtual Segment Id configured on the remote VTEP\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpAddressVniCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpAddressVniUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpAddressVniRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpAddressVniDelete,

		Schema: map[string]*schema.Schema{
			"segment": {
				Type: schema.TypeInt, Required: true, Description: "VNI configured for the remote VTEP",
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
func resourceOverlayTunnelVtepRemoteIpAddressVniCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressVniCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressVni(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressVniRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressVniUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressVniUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressVni(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressVniRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpAddressVniDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressVniDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressVni(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressVniRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressVniRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressVni(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpAddressVni(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpAddressVni {
	var ret edpt.OverlayTunnelVtepRemoteIpAddressVni
	ret.Inst.Segment = d.Get("segment").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	return ret
}
