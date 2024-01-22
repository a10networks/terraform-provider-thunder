package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIsisAddressFamilyIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_isis_address_family_ipv6`: Address family\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIsisAddressFamilyIpv6Create,
		UpdateContext: resourceRouterIsisAddressFamilyIpv6Update,
		ReadContext:   resourceRouterIsisAddressFamilyIpv6Read,
		DeleteContext: resourceRouterIsisAddressFamilyIpv6Delete,

		Schema: map[string]*schema.Schema{
			"adjacency_check": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Check ISIS neighbor protocol support",
			},
			"default_information": {
				Type: schema.TypeString, Optional: true, Description: "'originate': Distribute a default route;",
			},
			"distance": {
				Type: schema.TypeInt, Optional: true, Default: 115, Description: "ISIS Administrative Distance (Distance value)",
			},
			"multi_topology_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multi_topology": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable multi-topology mode",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Level-1 only; 'level-1-2': Level-1-2; 'level-2': Level-2 only;",
						},
						"transition": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept and generate both IS-IS IPv6 and Multi-topology IPV6 TLVs",
						},
						"level_transition": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept and generate both IS-IS IPv6 and Multi-topology IPV6 TLVs",
						},
					},
				},
			},
			"redistribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"summary_prefix_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Default: "level-2", Description: "'level-1': Summarize into level-1 area; 'level-1-2': Summarize into both area and sub-domain; 'level-2': Summarize into level-2 sub-domain;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"tag": {
				Type: schema.TypeString, Required: true, Description: "Tag",
			},
		},
	}
}
func resourceRouterIsisAddressFamilyIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIsisAddressFamilyIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceRouterIsisAddressFamilyIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIsisAddressFamilyIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceRouterIsisAddressFamilyIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIsisAddressFamilyIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIsisAddressFamilyIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIsisAddressFamilyIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterIsisAddressFamilyIpv6MultiTopologyCfg(d []interface{}) edpt.RouterIsisAddressFamilyIpv6MultiTopologyCfg {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6MultiTopologyCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MultiTopology = in["multi_topology"].(int)
		ret.Level = in["level"].(string)
		ret.Transition = in["transition"].(int)
		ret.LevelTransition = in["level_transition"].(int)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6Redistribute1268(d []interface{}) edpt.RouterIsisAddressFamilyIpv6Redistribute1268 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6Redistribute1268
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterIsisAddressFamilyIpv6RedistributeRedistList1269(in["redist_list"].([]interface{}))
		ret.VipList = getSliceRouterIsisAddressFamilyIpv6RedistributeVipList1270(in["vip_list"].([]interface{}))
		ret.Isis = getObjectRouterIsisAddressFamilyIpv6RedistributeIsis1271(in["isis"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6RedistributeRedistList1269(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList1269 {

	count1 := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList1269, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6RedistributeRedistList1269
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6RedistributeVipList1270(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6RedistributeVipList1270 {

	count1 := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6RedistributeVipList1270, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6RedistributeVipList1270
		oi.VipType = in["vip_type"].(string)
		oi.VipMetric = in["vip_metric"].(int)
		oi.VipRouteMap = in["vip_route_map"].(string)
		oi.VipMetricType = in["vip_metric_type"].(string)
		oi.VipLevel = in["vip_level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsis1271(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsis1271 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsis1271
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1From = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1272(in["level_1_from"].([]interface{}))
		ret.Level2From = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1274(in["level_2_from"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1272(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1272 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1272
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Into1 = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11273(in["into_1"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11273(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11273 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11273
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level2 = in["level_2"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1274(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1274 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1274
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Into2 = getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21275(in["into_2"].([]interface{}))
	}
	return ret
}

func getObjectRouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21275(d []interface{}) edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21275 {

	count1 := len(d)
	var ret edpt.RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21275
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Level1 = in["level_1"].(int)
		ret.DistributeList = in["distribute_list"].(string)
	}
	return ret
}

func getSliceRouterIsisAddressFamilyIpv6SummaryPrefixList(d []interface{}) []edpt.RouterIsisAddressFamilyIpv6SummaryPrefixList {

	count1 := len(d)
	ret := make([]edpt.RouterIsisAddressFamilyIpv6SummaryPrefixList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIsisAddressFamilyIpv6SummaryPrefixList
		oi.Prefix = in["prefix"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIsisAddressFamilyIpv6(d *schema.ResourceData) edpt.RouterIsisAddressFamilyIpv6 {
	var ret edpt.RouterIsisAddressFamilyIpv6
	ret.Inst.AdjacencyCheck = d.Get("adjacency_check").(int)
	ret.Inst.DefaultInformation = d.Get("default_information").(string)
	ret.Inst.Distance = d.Get("distance").(int)
	ret.Inst.MultiTopologyCfg = getObjectRouterIsisAddressFamilyIpv6MultiTopologyCfg(d.Get("multi_topology_cfg").([]interface{}))
	ret.Inst.Redistribute = getObjectRouterIsisAddressFamilyIpv6Redistribute1268(d.Get("redistribute").([]interface{}))
	ret.Inst.SummaryPrefixList = getSliceRouterIsisAddressFamilyIpv6SummaryPrefixList(d.Get("summary_prefix_list").([]interface{}))
	//omit uuid
	ret.Inst.Tag = d.Get("tag").(string)
	return ret
}
