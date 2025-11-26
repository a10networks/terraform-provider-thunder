package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepLocalIpv6AddressVni() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_local_ipv6_address_vni`: IP Address of the local tunnel end point\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepLocalIpv6AddressVniCreate,
		UpdateContext: resourceOverlayTunnelVtepLocalIpv6AddressVniUpdate,
		ReadContext:   resourceOverlayTunnelVtepLocalIpv6AddressVniRead,
		DeleteContext: resourceOverlayTunnelVtepLocalIpv6AddressVniDelete,

		Schema: map[string]*schema.Schema{
			"gateway": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "This is a Gateway segment id",
			},
			"lif": {
				Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
			},
			"partition": {
				Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
			},
			"segment": {
				Type: schema.TypeInt, Required: true, Description: "Id of the segment that is being extended",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
		},
	}
}
func resourceOverlayTunnelVtepLocalIpv6AddressVniCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressVniCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6AddressVni(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpv6AddressVniRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpv6AddressVniUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressVniUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6AddressVni(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpv6AddressVniRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepLocalIpv6AddressVniDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressVniDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6AddressVni(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpv6AddressVniRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressVniRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6AddressVni(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepLocalIpv6AddressVni(d *schema.ResourceData) edpt.OverlayTunnelVtepLocalIpv6AddressVni {
	var ret edpt.OverlayTunnelVtepLocalIpv6AddressVni
	ret.Inst.Gateway = d.Get("gateway").(int)
	ret.Inst.Lif = d.Get("lif").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.Segment = d.Get("segment").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
