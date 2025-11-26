package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6RipRouteMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_rip_route_map`: Route map set\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6RipRouteMapCreate,
		UpdateContext: resourceRouterIpv6RipRouteMapUpdate,
		ReadContext:   resourceRouterIpv6RipRouteMapRead,
		DeleteContext: resourceRouterIpv6RipRouteMapDelete,

		Schema: map[string]*schema.Schema{
			"map_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"map": {
							Type: schema.TypeString, Optional: true, Description: "Route map name",
						},
						"route_map_direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Route map set for input filtering; 'out': Route map set for output filtering;",
						},
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
						},
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
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
func resourceRouterIpv6RipRouteMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRouteMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRouteMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipRouteMapRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6RipRouteMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRouteMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRouteMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipRouteMapRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6RipRouteMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRouteMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRouteMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6RipRouteMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRouteMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRouteMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6RipRouteMapMapCfg(d []interface{}) []edpt.RouterIpv6RipRouteMapMapCfg {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRouteMapMapCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRouteMapMapCfg
		oi.Map = in["map"].(string)
		oi.RouteMapDirection = in["route_map_direction"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Loopback = in["loopback"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIpv6RipRouteMap(d *schema.ResourceData) edpt.RouterIpv6RipRouteMap {
	var ret edpt.RouterIpv6RipRouteMap
	ret.Inst.MapCfg = getSliceRouterIpv6RipRouteMapMapCfg(d.Get("map_cfg").([]interface{}))
	//omit uuid
	return ret
}
