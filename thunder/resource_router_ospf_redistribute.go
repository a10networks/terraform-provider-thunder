package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfRedistribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ospf_redistribute`: Redistribute information from another routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterOspfRedistributeCreate,
		UpdateContext: resourceRouterOspfRedistributeUpdate,
		ReadContext:   resourceRouterOspfRedistributeRead,
		DeleteContext: resourceRouterOspfRedistributeDelete,

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
				Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
			},
			"metric_type_ip_nat": {
				Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
			},
			"ospf_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
						},
						"process_id": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF process ID",
						},
						"metric_ospf": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type_ospf": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
						},
						"route_map_ospf": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"tag_ospf": {
							Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
						},
					},
				},
			},
			"redist_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"tag": {
							Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
						},
					},
				},
			},
			"route_map_ip_nat": {
				Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
			},
			"tag_ip_nat": {
				Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
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
							Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type_vip": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
						},
						"route_map_vip": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"tag_vip": {
							Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
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
func resourceRouterOspfRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfRedistributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfRedistribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfRedistributeRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterOspfRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfRedistributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfRedistribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfRedistributeRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterOspfRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfRedistributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfRedistribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterOspfRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfRedistributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspfRedistribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterOspfRedistributeIpNatFloatingList(d []interface{}) []edpt.RouterOspfRedistributeIpNatFloatingList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeIpNatFloatingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeIpNatFloatingList
		oi.IpNatPrefix = in["ip_nat_prefix"].(string)
		oi.IpNatFloatingIpForward = in["ip_nat_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeOspfList(d []interface{}) []edpt.RouterOspfRedistributeOspfList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeOspfList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeOspfList
		oi.Ospf = in["ospf"].(int)
		oi.ProcessId = in["process_id"].(int)
		oi.MetricOspf = in["metric_ospf"].(int)
		oi.MetricTypeOspf = in["metric_type_ospf"].(string)
		oi.RouteMapOspf = in["route_map_ospf"].(string)
		oi.TagOspf = in["tag_ospf"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeRedistList(d []interface{}) []edpt.RouterOspfRedistributeRedistList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeRedistList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeRedistList
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Tag = in["tag"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeVipFloatingList(d []interface{}) []edpt.RouterOspfRedistributeVipFloatingList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeVipFloatingList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeVipFloatingList
		oi.VipAddress = in["vip_address"].(string)
		oi.VipFloatingIpForward = in["vip_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeVipList(d []interface{}) []edpt.RouterOspfRedistributeVipList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeVipList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeVipList
		oi.TypeVip = in["type_vip"].(string)
		oi.MetricVip = in["metric_vip"].(int)
		oi.MetricTypeVip = in["metric_type_vip"].(string)
		oi.RouteMapVip = in["route_map_vip"].(string)
		oi.TagVip = in["tag_vip"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterOspfRedistribute(d *schema.ResourceData) edpt.RouterOspfRedistribute {
	var ret edpt.RouterOspfRedistribute
	ret.Inst.IpNat = d.Get("ip_nat").(int)
	ret.Inst.IpNatFloatingList = getSliceRouterOspfRedistributeIpNatFloatingList(d.Get("ip_nat_floating_list").([]interface{}))
	ret.Inst.MetricIpNat = d.Get("metric_ip_nat").(int)
	ret.Inst.MetricTypeIpNat = d.Get("metric_type_ip_nat").(string)
	ret.Inst.OspfList = getSliceRouterOspfRedistributeOspfList(d.Get("ospf_list").([]interface{}))
	ret.Inst.RedistList = getSliceRouterOspfRedistributeRedistList(d.Get("redist_list").([]interface{}))
	ret.Inst.RouteMapIpNat = d.Get("route_map_ip_nat").(string)
	ret.Inst.TagIpNat = d.Get("tag_ip_nat").(int)
	//omit uuid
	ret.Inst.VipFloatingList = getSliceRouterOspfRedistributeVipFloatingList(d.Get("vip_floating_list").([]interface{}))
	ret.Inst.VipList = getSliceRouterOspfRedistributeVipList(d.Get("vip_list").([]interface{}))
	ret.Inst.ProcessId = d.Get("process_id").(string)
	return ret
}
