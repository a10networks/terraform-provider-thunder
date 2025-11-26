package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtepLocalIpv6Address() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep_local_ipv6_address`: IP Address of the local tunnel end point\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepLocalIpv6AddressCreate,
		UpdateContext: resourceOverlayTunnelVtepLocalIpv6AddressUpdate,
		ReadContext:   resourceOverlayTunnelVtepLocalIpv6AddressRead,
		DeleteContext: resourceOverlayTunnelVtepLocalIpv6AddressDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "Source Tunnel End Point IPv6 address",
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
func resourceOverlayTunnelVtepLocalIpv6AddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6Address(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpv6AddressRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpv6AddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6Address(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepLocalIpv6AddressRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepLocalIpv6AddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6Address(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepLocalIpv6AddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepLocalIpv6AddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtepLocalIpv6Address(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceOverlayTunnelVtepLocalIpv6AddressVniList(d []interface{}) []edpt.OverlayTunnelVtepLocalIpv6AddressVniList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepLocalIpv6AddressVniList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepLocalIpv6AddressVniList
		oi.Segment = in["segment"].(int)
		oi.Partition = in["partition"].(string)
		oi.Gateway = in["gateway"].(int)
		oi.Lif = in["lif"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointOverlayTunnelVtepLocalIpv6Address(d *schema.ResourceData) edpt.OverlayTunnelVtepLocalIpv6Address {
	var ret edpt.OverlayTunnelVtepLocalIpv6Address
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	//omit uuid
	ret.Inst.VniList = getSliceOverlayTunnelVtepLocalIpv6AddressVniList(d.Get("vni_list").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(string)
	return ret
}
