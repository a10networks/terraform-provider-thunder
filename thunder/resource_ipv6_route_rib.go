package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteRib() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_rib`: Establish static routes\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteRibCreate,
		UpdateContext: resourceIpv6RouteRibUpdate,
		ReadContext:   resourceIpv6RouteRibRead,
		DeleteContext: resourceIpv6RouteRibDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_address": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address",
			},
			"ipv6_nexthop_ipv6": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_nexthop": {
							Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
						},
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Ethernet interface number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface (Virtual Ethernet interface number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"distance": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Distance value for this route",
						},
						"description": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
						},
					},
				},
			},
			"ipv6_nexthop_tunnel": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
						"ipv6_nexthop_tunnel_addr": {
							Type: schema.TypeString, Optional: true, Description: "Forwarding router's address",
						},
						"distance_nexthop_tunnel": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Distance value for this route",
						},
						"description": {
							Type: schema.TypeString, Optional: true, Description: "Description for static route",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6RouteRibCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteRibCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteRib(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteRibRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteRibUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteRibUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteRib(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteRibRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteRibDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteRibDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteRib(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteRibRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteRibRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteRib(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpv6RouteRibIpv6NexthopIpv6(d []interface{}) []edpt.Ipv6RouteRibIpv6NexthopIpv6 {

	count1 := len(d)
	ret := make([]edpt.Ipv6RouteRibIpv6NexthopIpv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6RouteRibIpv6NexthopIpv6
		oi.Ipv6Nexthop = in["ipv6_nexthop"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Distance = in["distance"].(int)
		oi.Description = in["description"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpv6RouteRibIpv6NexthopTunnel(d []interface{}) []edpt.Ipv6RouteRibIpv6NexthopTunnel {

	count1 := len(d)
	ret := make([]edpt.Ipv6RouteRibIpv6NexthopTunnel, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6RouteRibIpv6NexthopTunnel
		oi.Tunnel = in["tunnel"].(int)
		oi.Ipv6NexthopTunnelAddr = in["ipv6_nexthop_tunnel_addr"].(string)
		oi.DistanceNexthopTunnel = in["distance_nexthop_tunnel"].(int)
		oi.Description = in["description"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6RouteRib(d *schema.ResourceData) edpt.Ipv6RouteRib {
	var ret edpt.Ipv6RouteRib
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Ipv6NexthopIpv6 = getSliceIpv6RouteRibIpv6NexthopIpv6(d.Get("ipv6_nexthop_ipv6").([]interface{}))
	ret.Inst.Ipv6NexthopTunnel = getSliceIpv6RouteRibIpv6NexthopTunnel(d.Get("ipv6_nexthop_tunnel").([]interface{}))
	//omit uuid
	return ret
}
