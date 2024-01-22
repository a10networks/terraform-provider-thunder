package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepRemoteIpv6AddressUseLif() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_remote_ipv6_address_use_lif`: Specify the interface lif\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepRemoteIpv6AddressUseLifCreate,
		UpdateContext: resourceOverlayTunnelVtepRemoteIpv6AddressUseLifUpdate,
		ReadContext:   resourceOverlayTunnelVtepRemoteIpv6AddressUseLifRead,
		DeleteContext: resourceOverlayTunnelVtepRemoteIpv6AddressUseLifDelete,

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
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "Ipv6Address",
			},
		},
	}
}
func resourceOverlayTunnelVtepRemoteIpv6AddressUseLifCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseLifCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseLif(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressUseLifRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressUseLifUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseLifUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseLif(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRemoteIpv6AddressUseLifRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepRemoteIpv6AddressUseLifDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseLifDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseLif(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRemoteIpv6AddressUseLifRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRemoteIpv6AddressUseLifRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseLif(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepRemoteIpv6AddressUseLif(d *schema.ResourceData) edpt.OverlayTunnelVtepRemoteIpv6AddressUseLif {
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressUseLif
	ret.Inst.Lif = d.Get("lif").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	return ret
}
