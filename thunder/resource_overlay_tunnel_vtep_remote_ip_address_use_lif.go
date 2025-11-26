package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpAddressUseLif() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ip_address_use_lif`: Specify the interface lif\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpAddressUseLifCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpAddressUseLifUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpAddressUseLifRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpAddressUseLifDelete,

		Schema: map[string]*schema.Schema{
			"lif": {
				Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
			},
			"partition": {
				Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
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
func resourceOverlayTunnelVtepRemoteIpAddressUseLifCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseLifCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseLif(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressUseLifRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressUseLifUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseLifUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseLif(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpAddressUseLifRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpAddressUseLifDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseLifDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseLif(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpAddressUseLifRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpAddressUseLifRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpAddressUseLif(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpAddressUseLif(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpAddressUseLif {
	var ret edpt.OverlayTunnelVtepRemoteIpAddressUseLif
	ret.Inst.Lif = d.Get("lif").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	return ret
}
