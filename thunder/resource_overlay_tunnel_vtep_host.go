package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepHost() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_host`: IP Address of the local tunnel end point\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepHostCreate,
		UpdateContext: resourceOverlayTunnelVtepHostUpdate,
		ReadContext:   resourceOverlayTunnelVtepHostRead,
		DeleteContext: resourceOverlayTunnelVtepHostDelete,

		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type: schema.TypeString, Required: true, Description: "IPv4 address of the overlay host",
			},
			"overlay_mac_addr": {
				Type: schema.TypeString, Required: true, Description: "MAC Address of the overlay host",
			},
			"remote_vtep": {
				Type: schema.TypeString, Required: true, Description: "Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vni": {
				Type: schema.TypeInt, Required: true, Description: "Configure the segment id ( VNI of the remote host)",
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
		},
	}
}
func resourceOverlayTunnelVtepHostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepHostCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepHost(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepHostRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepHostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepHostUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepHost(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepHostRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepHostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepHostDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepHost(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepHostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepHostRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepHost(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointOverlayTunnelVtepHost(d *schema.ResourceData) edpt.OverlayTunnelVtepHost {
	var ret edpt.OverlayTunnelVtepHost
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.OverlayMacAddr = d.Get("overlay_mac_addr").(string)
	ret.Inst.RemoteVtep = d.Get("remote_vtep").(string)
	//omit uuid
	ret.Inst.Vni = d.Get("vni").(int)
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
