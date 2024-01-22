package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6OspfRedistribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_ospf_redistribute`: Redistribute information from another routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6OspfRedistributeCreate,
		UpdateContext: resourceRouterIpv6OspfRedistributeUpdate,
		ReadContext:   resourceRouterIpv6OspfRedistributeRead,
		DeleteContext: resourceRouterIpv6OspfRedistributeDelete,

		Schema: map[string]*schema.Schema{
			"ip_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP-NAT",
			},
			"ip_nat_floating_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_nat_prefix": {
							Type: schema.TypeString, Optional: true, Description: "Address",
						},
						"ip_nat_floating_ip_forward": {
							Type: schema.TypeString, Optional: true, Description: "Floating-IP as forward address",
						},
					},
				},
			},
			"metric_ip_nat": {
				Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
			},
			"metric_type_ip_nat": {
				Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
			},
			"ospf_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
						},
						"process_id": {
							Type: schema.TypeString, Optional: true, Description: "OSPFV3 process tag",
						},
						"metric_ospf": {
							Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
						},
						"metric_type_ospf": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
						},
						"route_map_ospf": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"redist_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'nat64': NAT64 Prefix; 'lw4o6': LW4O6 Prefix; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
						},
						"metric_type": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"route_map_ip_nat": {
				Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_floating_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vip_address": {
							Type: schema.TypeString, Optional: true, Description: "Address",
						},
						"vip_floating_ip_forward": {
							Type: schema.TypeString, Optional: true, Description: "Floating-IP as forward address",
						},
					},
				},
			},
			"vip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type_vip": {
							Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
						},
						"metric_vip": {
							Type: schema.TypeInt, Optional: true, Description: "OSPFV3 default metric (OSPFV3 metric)",
						},
						"metric_type_vip": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPFV3 External Type 1 metrics; '2': Set OSPFV3 External Type 2 metrics;",
						},
						"route_map_vip": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"process_id": {
				Type: schema.TypeString, Required: true, Description: "ProcessId",
			},
		},
	}
}
func resourceRouterIpv6OspfRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfRedistributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfRedistribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfRedistributeRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6OspfRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfRedistributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfRedistribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfRedistributeRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6OspfRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfRedistributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfRedistribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6OspfRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfRedistributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfRedistribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6OspfRedistributeIpNatFloatingList(d []interface{}) []edpt.RouterIpv6OspfRedistributeIpNatFloatingList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeIpNatFloatingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeIpNatFloatingList
		oi.IpNatPrefix = in["ip_nat_prefix"].(string)
		oi.IpNatFloatingIpForward = in["ip_nat_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeOspfList(d []interface{}) []edpt.RouterIpv6OspfRedistributeOspfList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeOspfList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeOspfList
		oi.Ospf = in["ospf"].(int)
		oi.ProcessId = in["process_id"].(string)
		oi.MetricOspf = in["metric_ospf"].(int)
		oi.MetricTypeOspf = in["metric_type_ospf"].(string)
		oi.RouteMapOspf = in["route_map_ospf"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeRedistList(d []interface{}) []edpt.RouterIpv6OspfRedistributeRedistList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeRedistList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeRedistList
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeVipFloatingList(d []interface{}) []edpt.RouterIpv6OspfRedistributeVipFloatingList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeVipFloatingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeVipFloatingList
		oi.VipAddress = in["vip_address"].(string)
		oi.VipFloatingIpForward = in["vip_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfRedistributeVipList(d []interface{}) []edpt.RouterIpv6OspfRedistributeVipList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfRedistributeVipList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfRedistributeVipList
		oi.TypeVip = in["type_vip"].(string)
		oi.MetricVip = in["metric_vip"].(int)
		oi.MetricTypeVip = in["metric_type_vip"].(string)
		oi.RouteMapVip = in["route_map_vip"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIpv6OspfRedistribute(d *schema.ResourceData) edpt.RouterIpv6OspfRedistribute {
	var ret edpt.RouterIpv6OspfRedistribute
	ret.Inst.IpNat = d.Get("ip_nat").(int)
	ret.Inst.IpNatFloatingList = getSliceRouterIpv6OspfRedistributeIpNatFloatingList(d.Get("ip_nat_floating_list").([]interface{}))
	ret.Inst.MetricIpNat = d.Get("metric_ip_nat").(int)
	ret.Inst.MetricTypeIpNat = d.Get("metric_type_ip_nat").(string)
	ret.Inst.OspfList = getSliceRouterIpv6OspfRedistributeOspfList(d.Get("ospf_list").([]interface{}))
	ret.Inst.RedistList = getSliceRouterIpv6OspfRedistributeRedistList(d.Get("redist_list").([]interface{}))
	ret.Inst.RouteMapIpNat = d.Get("route_map_ip_nat").(string)
	//omit uuid
	ret.Inst.VipFloatingList = getSliceRouterIpv6OspfRedistributeVipFloatingList(d.Get("vip_floating_list").([]interface{}))
	ret.Inst.VipList = getSliceRouterIpv6OspfRedistributeVipList(d.Get("vip_list").([]interface{}))
	ret.Inst.ProcessId = d.Get("process_id").(string)
	return ret
}
