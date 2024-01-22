package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6RipRedistribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_rip_redistribute`: Redistribute information from another routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6RipRedistributeCreate,
		UpdateContext: resourceRouterIpv6RipRedistributeUpdate,
		ReadContext:   resourceRouterIpv6RipRedistributeRead,
		DeleteContext: resourceRouterIpv6RipRedistributeDelete,

		Schema: map[string]*schema.Schema{
			"redist_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'ip-nat': IP NAT; 'isis': ISO IS-IS; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'nat64': NAT64 Prefix; 'static-nat': Static NAT; 'ospf': Open Shortest Path First (OSPF); 'static': Static routes;",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Description: "Metric for redistributed routes (metric value)",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vip_type": {
							Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
						},
						"vip_metric": {
							Type: schema.TypeInt, Optional: true, Description: "Metric for redistributed routes (metric value)",
						},
						"vip_route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
		},
	}
}
func resourceRouterIpv6RipRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRedistributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRedistribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipRedistributeRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6RipRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRedistributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRedistribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6RipRedistributeRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6RipRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRedistributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRedistribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6RipRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6RipRedistributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6RipRedistribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6RipRedistributeRedistList(d []interface{}) []edpt.RouterIpv6RipRedistributeRedistList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRedistributeRedistList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRedistributeRedistList
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.RouteMap = in["route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6RipRedistributeVipList(d []interface{}) []edpt.RouterIpv6RipRedistributeVipList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6RipRedistributeVipList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6RipRedistributeVipList
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIpv6RipRedistribute(d *schema.ResourceData) edpt.RouterIpv6RipRedistribute {
	var ret edpt.RouterIpv6RipRedistribute
	ret.Inst.RedistList = getSliceRouterIpv6RipRedistributeRedistList(d.Get("redist_list").([]interface{}))
	//omit uuid
	ret.Inst.VipList = getSliceRouterIpv6RipRedistributeVipList(d.Get("vip_list").([]interface{}))
	return ret
}
