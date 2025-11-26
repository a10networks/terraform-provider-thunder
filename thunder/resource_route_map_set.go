package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouteMapSet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_route_map_set`: Set values in destination routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouteMapSetCreate,
		UpdateContext: resourceRouteMapSetUpdate,
		ReadContext:   resourceRouteMapSetRead,
		DeleteContext: resourceRouteMapSetDelete,

		Schema: map[string]*schema.Schema{
			"aggregator": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregator_as": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"asn": {
										Type: schema.TypeInt, Optional: true, Description: "AS number",
									},
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address of aggregator",
									},
								},
							},
						},
					},
				},
			},
			"as_path": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prepend": {
							Type: schema.TypeString, Optional: true, Description: "Prepend to the as-path (AS number)",
						},
						"num": {
							Type: schema.TypeString, Optional: true, Description: "AS number",
						},
						"num2": {
							Type: schema.TypeString, Optional: true, Description: "AS number",
						},
					},
				},
			},
			"atomic_aggregate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "BGP atomic aggregate attribute",
			},
			"comm_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v_std": {
							Type: schema.TypeInt, Optional: true, Description: "Community-list number (standard)",
						},
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching communities",
						},
						"v_exp": {
							Type: schema.TypeInt, Optional: true, Description: "Community-list number (expanded)",
						},
						"v_exp_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching communities",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Community-list name",
						},
						"name_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching communities",
						},
					},
				},
			},
			"community": {
				Type: schema.TypeString, Optional: true, Description: "BGP community attribute",
			},
			"dampening_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dampening": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable route-flap dampening",
						},
						"dampening_half_time": {
							Type: schema.TypeInt, Optional: true, Description: "Reachability Half-life time for the penalty(minutes)",
						},
						"dampening_reuse": {
							Type: schema.TypeInt, Optional: true, Description: "Value to start reusing a route",
						},
						"dampening_supress": {
							Type: schema.TypeInt, Optional: true, Description: "Value to start suppressing a route",
						},
						"dampening_max_supress": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum duration to suppress a stable route(minutes)",
						},
						"dampening_penalty": {
							Type: schema.TypeInt, Optional: true, Description: "Un-reachability Half-life time for the penalty(minutes)",
						},
					},
				},
			},
			"ddos": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Class-List Name",
						},
						"class_list_cid": {
							Type: schema.TypeInt, Optional: true, Description: "Class-List Cid",
						},
						"zone": {
							Type: schema.TypeString, Optional: true, Description: "Zone Name",
						},
					},
				},
			},
			"extcommunity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rt": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "VPN extended community",
									},
								},
							},
						},
						"soo": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "VPN extended community",
									},
								},
							},
						},
					},
				},
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"next_hop": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "IP address of next hop",
									},
								},
							},
						},
					},
				},
			},
			"ipv6": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"next_hop_1": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "global address of next hop",
									},
									"local": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 address of next hop",
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
			"large_comm_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l_v_std": {
							Type: schema.TypeInt, Optional: true, Description: "Large Community-list number (standard)",
						},
						"l_v_std_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching large communities",
						},
						"l_v_exp": {
							Type: schema.TypeInt, Optional: true, Description: "Large Community-list number (expanded)",
						},
						"l_v_exp_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching large communities",
						},
						"l_name": {
							Type: schema.TypeString, Optional: true, Description: "Large Community-list name",
						},
						"large_name_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete matching large communities",
						},
					},
				},
			},
			"large_community": {
				Type: schema.TypeString, Optional: true, Description: "BGP large community attribute",
			},
			"level": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Export into a level-1 area; 'level-1-2': Export into level-1 and level-2; 'level-2': Export into level-2 sub-domain;",
						},
					},
				},
			},
			"local_preference": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"val": {
							Type: schema.TypeInt, Optional: true, Description: "Preference value",
						},
					},
				},
			},
			"metric": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "Metric Value (from -4294967295 to 4294967295)",
						},
					},
				},
			},
			"metric_type": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "'external': IS-IS external metric type; 'internal': IS-IS internal metric type; 'type-1': OSPF external type 1 metric; 'type-2': OSPF external type 2 metric;",
						},
					},
				},
			},
			"origin": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"egp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "remote EGP",
						},
						"igp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "local IGP",
						},
						"incomplete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "unknown heritage",
						},
					},
				},
			},
			"originator_id": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"originator_ip": {
							Type: schema.TypeString, Optional: true, Description: "IP address of originator",
						},
					},
				},
			},
			"tag": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Description: "Tag value",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"weight_val": {
							Type: schema.TypeInt, Optional: true, Description: "Weight value",
						},
					},
				},
			},
			"sequence": {
				Type: schema.TypeString, Required: true, Description: "Sequence",
			},
			"action": {
				Type: schema.TypeString, Required: true, Description: "Action",
			},
		},
	}
}
func resourceRouteMapSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapSetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapSet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouteMapSetRead(ctx, d, meta)
	}
	return diags
}

func resourceRouteMapSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapSetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapSet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouteMapSetRead(ctx, d, meta)
	}
	return diags
}
func resourceRouteMapSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapSetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapSet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouteMapSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapSetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapSet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouteMapSetAggregator(d []interface{}) edpt.RouteMapSetAggregator {

	count1 := len(d)
	var ret edpt.RouteMapSetAggregator
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AggregatorAs = getObjectRouteMapSetAggregatorAggregatorAs(in["aggregator_as"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetAggregatorAggregatorAs(d []interface{}) edpt.RouteMapSetAggregatorAggregatorAs {

	count1 := len(d)
	var ret edpt.RouteMapSetAggregatorAggregatorAs
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Asn = in["asn"].(int)
		ret.Ip = in["ip"].(string)
	}
	return ret
}

func getObjectRouteMapSetAsPath(d []interface{}) edpt.RouteMapSetAsPath {

	count1 := len(d)
	var ret edpt.RouteMapSetAsPath
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Prepend = in["prepend"].(string)
		ret.Num = in["num"].(string)
		ret.Num2 = in["num2"].(string)
	}
	return ret
}

func getObjectRouteMapSetCommList(d []interface{}) edpt.RouteMapSetCommList {

	count1 := len(d)
	var ret edpt.RouteMapSetCommList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VStd = in["v_std"].(int)
		ret.Delete = in["delete"].(int)
		ret.VExp = in["v_exp"].(int)
		ret.VExpDelete = in["v_exp_delete"].(int)
		ret.Name = in["name"].(string)
		ret.NameDelete = in["name_delete"].(int)
	}
	return ret
}

func getObjectRouteMapSetDampeningCfg(d []interface{}) edpt.RouteMapSetDampeningCfg {

	count1 := len(d)
	var ret edpt.RouteMapSetDampeningCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dampening = in["dampening"].(int)
		ret.DampeningHalfTime = in["dampening_half_time"].(int)
		ret.DampeningReuse = in["dampening_reuse"].(int)
		ret.DampeningSupress = in["dampening_supress"].(int)
		ret.DampeningMaxSupress = in["dampening_max_supress"].(int)
		ret.DampeningPenalty = in["dampening_penalty"].(int)
	}
	return ret
}

func getObjectRouteMapSetDdos(d []interface{}) edpt.RouteMapSetDdos {

	count1 := len(d)
	var ret edpt.RouteMapSetDdos
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClassListName = in["class_list_name"].(string)
		ret.ClassListCid = in["class_list_cid"].(int)
		ret.Zone = in["zone"].(string)
	}
	return ret
}

func getObjectRouteMapSetExtcommunity(d []interface{}) edpt.RouteMapSetExtcommunity {

	count1 := len(d)
	var ret edpt.RouteMapSetExtcommunity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rt = getObjectRouteMapSetExtcommunityRt(in["rt"].([]interface{}))
		ret.Soo = getObjectRouteMapSetExtcommunitySoo(in["soo"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetExtcommunityRt(d []interface{}) edpt.RouteMapSetExtcommunityRt {

	count1 := len(d)
	var ret edpt.RouteMapSetExtcommunityRt
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetExtcommunitySoo(d []interface{}) edpt.RouteMapSetExtcommunitySoo {

	count1 := len(d)
	var ret edpt.RouteMapSetExtcommunitySoo
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetIp(d []interface{}) edpt.RouteMapSetIp {

	count1 := len(d)
	var ret edpt.RouteMapSetIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NextHop = getObjectRouteMapSetIpNextHop(in["next_hop"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetIpNextHop(d []interface{}) edpt.RouteMapSetIpNextHop {

	count1 := len(d)
	var ret edpt.RouteMapSetIpNextHop
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = in["address"].(string)
	}
	return ret
}

func getObjectRouteMapSetIpv6(d []interface{}) edpt.RouteMapSetIpv6 {

	count1 := len(d)
	var ret edpt.RouteMapSetIpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NextHop1 = getObjectRouteMapSetIpv6NextHop1(in["next_hop_1"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetIpv6NextHop1(d []interface{}) edpt.RouteMapSetIpv6NextHop1 {

	count1 := len(d)
	var ret edpt.RouteMapSetIpv6NextHop1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = in["address"].(string)
		ret.Local = getObjectRouteMapSetIpv6NextHop1Local(in["local"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapSetIpv6NextHop1Local(d []interface{}) edpt.RouteMapSetIpv6NextHop1Local {

	count1 := len(d)
	var ret edpt.RouteMapSetIpv6NextHop1Local
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = in["address"].(string)
	}
	return ret
}

func getObjectRouteMapSetLargeCommList(d []interface{}) edpt.RouteMapSetLargeCommList {

	count1 := len(d)
	var ret edpt.RouteMapSetLargeCommList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LVStd = in["l_v_std"].(int)
		ret.LVStdDelete = in["l_v_std_delete"].(int)
		ret.LVExp = in["l_v_exp"].(int)
		ret.LVExpDelete = in["l_v_exp_delete"].(int)
		ret.LName = in["l_name"].(string)
		ret.LargeNameDelete = in["large_name_delete"].(int)
	}
	return ret
}

func getObjectRouteMapSetLevel(d []interface{}) edpt.RouteMapSetLevel {

	count1 := len(d)
	var ret edpt.RouteMapSetLevel
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetLocalPreference(d []interface{}) edpt.RouteMapSetLocalPreference {

	count1 := len(d)
	var ret edpt.RouteMapSetLocalPreference
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Val = in["val"].(int)
	}
	return ret
}

func getObjectRouteMapSetMetric(d []interface{}) edpt.RouteMapSetMetric {

	count1 := len(d)
	var ret edpt.RouteMapSetMetric
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetMetricType(d []interface{}) edpt.RouteMapSetMetricType {

	count1 := len(d)
	var ret edpt.RouteMapSetMetricType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapSetOrigin(d []interface{}) edpt.RouteMapSetOrigin {

	count1 := len(d)
	var ret edpt.RouteMapSetOrigin
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Egp = in["egp"].(int)
		ret.Igp = in["igp"].(int)
		ret.Incomplete = in["incomplete"].(int)
	}
	return ret
}

func getObjectRouteMapSetOriginatorId(d []interface{}) edpt.RouteMapSetOriginatorId {

	count1 := len(d)
	var ret edpt.RouteMapSetOriginatorId
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OriginatorIp = in["originator_ip"].(string)
	}
	return ret
}

func getObjectRouteMapSetTag(d []interface{}) edpt.RouteMapSetTag {

	count1 := len(d)
	var ret edpt.RouteMapSetTag
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
	}
	return ret
}

func getObjectRouteMapSetWeight(d []interface{}) edpt.RouteMapSetWeight {

	count1 := len(d)
	var ret edpt.RouteMapSetWeight
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WeightVal = in["weight_val"].(int)
	}
	return ret
}

func dataToEndpointRouteMapSet(d *schema.ResourceData) edpt.RouteMapSet {
	var ret edpt.RouteMapSet
	ret.Inst.Aggregator = getObjectRouteMapSetAggregator(d.Get("aggregator").([]interface{}))
	ret.Inst.AsPath = getObjectRouteMapSetAsPath(d.Get("as_path").([]interface{}))
	ret.Inst.AtomicAggregate = d.Get("atomic_aggregate").(int)
	ret.Inst.CommList = getObjectRouteMapSetCommList(d.Get("comm_list").([]interface{}))
	ret.Inst.Community = d.Get("community").(string)
	ret.Inst.DampeningCfg = getObjectRouteMapSetDampeningCfg(d.Get("dampening_cfg").([]interface{}))
	ret.Inst.Ddos = getObjectRouteMapSetDdos(d.Get("ddos").([]interface{}))
	ret.Inst.Extcommunity = getObjectRouteMapSetExtcommunity(d.Get("extcommunity").([]interface{}))
	ret.Inst.Ip = getObjectRouteMapSetIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectRouteMapSetIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.LargeCommList = getObjectRouteMapSetLargeCommList(d.Get("large_comm_list").([]interface{}))
	ret.Inst.LargeCommunity = d.Get("large_community").(string)
	ret.Inst.Level = getObjectRouteMapSetLevel(d.Get("level").([]interface{}))
	ret.Inst.LocalPreference = getObjectRouteMapSetLocalPreference(d.Get("local_preference").([]interface{}))
	ret.Inst.Metric = getObjectRouteMapSetMetric(d.Get("metric").([]interface{}))
	ret.Inst.MetricType = getObjectRouteMapSetMetricType(d.Get("metric_type").([]interface{}))
	ret.Inst.Origin = getObjectRouteMapSetOrigin(d.Get("origin").([]interface{}))
	ret.Inst.OriginatorId = getObjectRouteMapSetOriginatorId(d.Get("originator_id").([]interface{}))
	ret.Inst.Tag = getObjectRouteMapSetTag(d.Get("tag").([]interface{}))
	//omit uuid
	ret.Inst.Weight = getObjectRouteMapSetWeight(d.Get("weight").([]interface{}))
	ret.Inst.Sequence = d.Get("sequence").(string)
	ret.Inst.Action = d.Get("action").(string)
	return ret
}
