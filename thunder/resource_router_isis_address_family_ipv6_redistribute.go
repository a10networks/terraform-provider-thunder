package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIsisAddressFamilyIpv6Redistribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_isis_address_family_ipv6_redistribute`: Redistribute information from another routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIsisAddressFamilyIpv6RedistributeCreate,
		UpdateContext: resourceRouterIsisAddressFamilyIpv6RedistributeUpdate,
		ReadContext:   resourceRouterIsisAddressFamilyIpv6RedistributeRead,
		DeleteContext: resourceRouterIsisAddressFamilyIpv6RedistributeDelete,

		Schema: map[string]*schema.Schema{
			"isis": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level_1_from": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"into_1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_2": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-area routes into level-2",
												},
												"distribute_list": {
													Type: schema.TypeString, Optional: true, Description: "Select routes (Access-list name)",
												},
											},
										},
									},
								},
							},
						},
						"level_2_from": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"into_2": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"level_1": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-area routes into level-2",
												},
												"distribute_list": {
													Type: schema.TypeString, Optional: true, Description: "Select routes (Access-list name)",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"redist_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'ip-nat': IP NAT; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'nat64': NAT64 Prefix; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Metric for redistributed routes (IS-IS default metric)",
						},
						"metric_type": {
							Type: schema.TypeString, Optional: true, Default: "internal", Description: "'external': Set IS-IS External metric type; 'internal': Set IS-IS Internal metric type;",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': IS-IS level-1 routes only; 'level-1-2': IS-IS level-1 and level-2 routes; 'level-2': IS-IS level-2 routes only;",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Metric for redistributed routes (IS-IS default metric)",
						},
						"vip_route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"vip_metric_type": {
							Type: schema.TypeString, Optional: true, Default: "internal", Description: "'external': Set IS-IS External metric type; 'internal': Set IS-IS Internal metric type;",
						},
						"vip_level": {
							Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': IS-IS level-1 routes only; 'level-1-2': IS-IS level-1 and level-2 routes; 'level-2': IS-IS level-2 routes only;",
						},
					},
				},
			},
			"tag": {
				Type: schema.TypeString, Required: true, Description: "Tag",
			},
		},
	}
}
func resourceRouterIsisAddressFamilyIpv6RedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6RedistributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6Redistribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIsisAddressFamilyIpv6RedistributeRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIsisAddressFamilyIpv6RedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6RedistributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6Redistribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIsisAddressFamilyIpv6RedistributeRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIsisAddressFamilyIpv6RedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6RedistributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6Redistribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIsisAddressFamilyIpv6RedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6RedistributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6Redistribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsis(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsis {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsis
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1From = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1From(in["level_1_from"].([]interface{}))
		ret.Level2From = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2From(in["level_2_from"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1From(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Into1 = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1(in["into_1"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level2 = in["level_2"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2From(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Into2 = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2(in["into_2"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1 = in["level_1"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6RedistributeRedistList(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList {

	count1 := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6RedistributeVipList(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6RedistributeVipList {

	count1 := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6RedistributeVipList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6RedistributeVipList
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		oi.VipMetricType = in["vip_metric_type"].(string)
		oi.VipLevel = in["vip_level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIsisAddressFamilyIpv6Redistribute(d *schema.ResourceData) edpt.RouterIsisAddressFamilyIpv6Redistribute {
	var ret edpt.RouterIsisAddressFamilyIpv6Redistribute
	ret.Inst.Isis = getObjectRouterIsisAddressFamilyIpv6RedistributeIsis(d.Get("isis").([]interface{}))
	ret.Inst.RedistList = getSliceRouterIsisAddressFamilyIpv6RedistributeRedistList(d.Get("redist_list").([]interface{}))
	//omit uuid
	ret.Inst.VipList = getSliceRouterIsisAddressFamilyIpv6RedistributeVipList(d.Get("vip_list").([]interface{}))
	ret.Inst.Tag = d.Get("tag").(string)
	return ret
}
