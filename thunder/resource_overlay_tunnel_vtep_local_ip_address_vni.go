package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepLocalIpAddressVni() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_local_ip_address_vni`: IP Address of the local tunnel end point\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepLocalIpAddressVniCreate,
		UpdateContext: resourceOverlayTunnelVtepLocalIpAddressVniUpdate,
		ReadContext:   resourceOverlayTunnelVtepLocalIpAddressVniRead,
		DeleteContext: resourceOverlayTunnelVtepLocalIpAddressVniDelete,

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
func resourceOverlayTunnelVtepLocalIpAddressVniCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressVniCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddressVni(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpAddressVniRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpAddressVniUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressVniUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddressVni(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpAddressVniRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepLocalIpAddressVniDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressVniDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddressVni(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpAddressVniRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressVniRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddressVni(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepLocalIpAddressVni(d *schema.ResourceData) edpt.OverlayTunnelVtepLocalIpAddressVni {
	var ret edpt.OverlayTunnelVtepLocalIpAddressVni
	ret.Inst.Gateway = d.Get("gateway").(int)
	ret.Inst.Lif = d.Get("lif").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.Segment = d.Get("segment").(int)
	//omit uuid
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
