package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepLocalIpAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_local_ip_address`: IP Address of the local tunnel end point\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepLocalIpAddressCreate,
		UpdateContext: resourceOverlayTunnelVtepLocalIpAddressUpdate,
		ReadContext:   resourceOverlayTunnelVtepLocalIpAddressRead,
		DeleteContext: resourceOverlayTunnelVtepLocalIpAddressDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type: schema.TypeString, Required: true, Description: "Source Tunnel End Point IPv4 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vni_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"segment": {
							Type: schema.TypeInt, Required: true, Description: "Id of the segment that is being extended",
						},
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
						},
						"gateway": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "This is a Gateway segment id",
						},
						"lif": {
							Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"id1": {
				Type: schema.TypeString, Required: true, Description: "Id1",
			},
		},
	}
}
func resourceOverlayTunnelVtepLocalIpAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepLocalIpAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceOverlayTunnelVtepLocalIpAddressVniList(d []interface{}) []edpt.OverlayTunnelVtepLocalIpAddressVniList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepLocalIpAddressVniList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepLocalIpAddressVniList
		oi.Segment = in["segment"].(int)
		oi.Partition = in["partition"].(string)
		oi.Gateway = in["gateway"].(int)
		oi.Lif = in["lif"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointOverlayTunnelVtepLocalIpAddress(d *schema.ResourceData) edpt.OverlayTunnelVtepLocalIpAddress {
	var ret edpt.OverlayTunnelVtepLocalIpAddress
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	//omit uuid
	ret.Inst.VniList = getSliceOverlayTunnelVtepLocalIpAddressVniList(d.Get("vni_list").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
