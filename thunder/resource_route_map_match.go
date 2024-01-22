package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouteMapMatch() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_route_map_match`: Match values from routing table\n\n__PLACEHOLDER__",
		CreateContext: resourceRouteMapMatchCreate,
		UpdateContext: resourceRouteMapMatchUpdate,
		ReadContext:   resourceRouteMapMatchRead,
		DeleteContext: resourceRouteMapMatchDelete,

		Schema: map[string]*schema.Schema{
			"as_path": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "AS path access-list name",
						},
					},
				},
			},
			"community": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "One or more Community Lists (numbered or named)",
									},
									"exact_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do exact matching of communities",
									},
								},
							},
						},
					},
				},
			},
			"extcommunity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extcommunity_l_name": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "One or more Extended Community Lists (numbered or named)",
									},
									"exact_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do exact matching of ecommunities",
									},
								},
							},
						},
					},
				},
			},
			"group": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type: schema.TypeInt, Optional: true, Description: "HA or VRRP-A group id",
						},
						"ha_state": {
							Type: schema.TypeString, Optional: true, Description: "'active': HA or VRRP-A in Active State; 'standby': HA or VRRP-A in Standby State;",
						},
					},
				},
			},
			"interface": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"loopback": {
							Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk interface number)",
						},
						"ve": {
							Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
						},
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
						},
					},
				},
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl1": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number",
									},
									"acl2": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "IP access-list name",
									},
									"prefix_list": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IP prefix-list name",
												},
											},
										},
									},
								},
							},
						},
						"next_hop": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl1": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number",
									},
									"acl2": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "IP access-list name",
									},
									"prefix_list_1": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IP prefix-list name",
												},
											},
										},
									},
								},
							},
						},
						"peer": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl1": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number",
									},
									"acl2": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "IP access-list name",
									},
								},
							},
						},
						"rib": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exact": {
										Type: schema.TypeString, Optional: true, Description: "Exact match a prefix (Prefix IPv4 address)",
									},
									"reachable": {
										Type: schema.TypeString, Optional: true, Description: "IP address is reachable",
									},
									"unreachable": {
										Type: schema.TypeString, Optional: true, Description: "IP address is unreachable",
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
						"address_1": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 access-list name",
									},
									"prefix_list_2": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "IPv6 prefix-list name",
												},
											},
										},
									},
								},
							},
						},
						"next_hop_1": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"next_hop_acl_name": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 access-list name",
									},
									"v6_addr": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address of next hop",
									},
									"prefix_list_name": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 prefix-list name",
									},
								},
							},
						},
						"peer_1": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acl1": {
										Type: schema.TypeInt, Optional: true, Description: "IPv6 access-list number",
									},
									"acl2": {
										Type: schema.TypeInt, Optional: true, Description: "IP access-list number (expanded range)",
									},
									"name": {
										Type: schema.TypeString, Optional: true, Description: "IP access-list name",
									},
								},
							},
						},
						"rib": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exact": {
										Type: schema.TypeString, Optional: true, Description: "Exact match a prefix (Prefix IPv6 address)",
									},
									"reachable": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address is reachable",
									},
									"unreachable": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address is unreachable",
									},
								},
							},
						},
					},
				},
			},
			"large_community": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l_name_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "One or more Large Community Lists (numbered or named)",
									},
									"exact_match": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do exact matching of large-communities",
									},
								},
							},
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
							Type: schema.TypeInt, Optional: true, Description: "Metric value",
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
			"route_type": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeString, Optional: true, Description: "'type-1': Match OSPF External Type 1 metrics; 'type-2': Match OSPF External Type 2 metrics;",
									},
								},
							},
						},
					},
				},
			},
			"scaleout": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type: schema.TypeInt, Optional: true, Description: "Scaleout Cluster-id",
						},
						"operational_state": {
							Type: schema.TypeString, Optional: true, Description: "'up': Scaleout is up and running; 'down': Scaleout is down or disabled;",
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
			"sequence": {
				Type: schema.TypeString, Required: true, Description: "Sequence",
			},
			"action": {
				Type: schema.TypeString, Required: true, Description: "Action",
			},
		},
	}
}
func resourceRouteMapMatchCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapMatchCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapMatch(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouteMapMatchRead(ctx, d, meta)
	}
	return diags
}

func resourceRouteMapMatchUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapMatchUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapMatch(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouteMapMatchRead(ctx, d, meta)
	}
	return diags
}
func resourceRouteMapMatchDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapMatchDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapMatch(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouteMapMatchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouteMapMatchRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouteMapMatch(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouteMapMatchAsPath(d []interface{}) edpt.RouteMapMatchAsPath {

	count1 := len(d)
	var ret edpt.RouteMapMatchAsPath
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchCommunity(d []interface{}) edpt.RouteMapMatchCommunity {

	count1 := len(d)
	var ret edpt.RouteMapMatchCommunity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NameCfg = getObjectRouteMapMatchCommunityNameCfg(in["name_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchCommunityNameCfg(d []interface{}) edpt.RouteMapMatchCommunityNameCfg {

	count1 := len(d)
	var ret edpt.RouteMapMatchCommunityNameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getObjectRouteMapMatchExtcommunity(d []interface{}) edpt.RouteMapMatchExtcommunity {

	count1 := len(d)
	var ret edpt.RouteMapMatchExtcommunity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExtcommunityLName = getObjectRouteMapMatchExtcommunityExtcommunityLName(in["extcommunity_l_name"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchExtcommunityExtcommunityLName(d []interface{}) edpt.RouteMapMatchExtcommunityExtcommunityLName {

	count1 := len(d)
	var ret edpt.RouteMapMatchExtcommunityExtcommunityLName
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getObjectRouteMapMatchGroup(d []interface{}) edpt.RouteMapMatchGroup {

	count1 := len(d)
	var ret edpt.RouteMapMatchGroup
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupId = in["group_id"].(int)
		ret.HaState = in["ha_state"].(string)
	}
	return ret
}

func getObjectRouteMapMatchInterface(d []interface{}) edpt.RouteMapMatchInterface {

	count1 := len(d)
	var ret edpt.RouteMapMatchInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethernet = in["ethernet"].(int)
		ret.Loopback = in["loopback"].(int)
		ret.Trunk = in["trunk"].(int)
		ret.Ve = in["ve"].(int)
		ret.Tunnel = in["tunnel"].(int)
	}
	return ret
}

func getObjectRouteMapMatchIp(d []interface{}) edpt.RouteMapMatchIp {

	count1 := len(d)
	var ret edpt.RouteMapMatchIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address = getObjectRouteMapMatchIpAddress(in["address"].([]interface{}))
		ret.NextHop = getObjectRouteMapMatchIpNextHop(in["next_hop"].([]interface{}))
		ret.Peer = getObjectRouteMapMatchIpPeer(in["peer"].([]interface{}))
		ret.Rib = getObjectRouteMapMatchIpRib(in["rib"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpAddress(d []interface{}) edpt.RouteMapMatchIpAddress {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpAddress
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
		ret.PrefixList = getObjectRouteMapMatchIpAddressPrefixList(in["prefix_list"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpAddressPrefixList(d []interface{}) edpt.RouteMapMatchIpAddressPrefixList {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpAddressPrefixList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpNextHop(d []interface{}) edpt.RouteMapMatchIpNextHop {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpNextHop
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
		ret.PrefixList1 = getObjectRouteMapMatchIpNextHopPrefixList1(in["prefix_list_1"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpNextHopPrefixList1(d []interface{}) edpt.RouteMapMatchIpNextHopPrefixList1 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpNextHopPrefixList1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpPeer(d []interface{}) edpt.RouteMapMatchIpPeer {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpPeer
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpRib(d []interface{}) edpt.RouteMapMatchIpRib {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpRib
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exact = in["exact"].(string)
		ret.Reachable = in["reachable"].(string)
		ret.Unreachable = in["unreachable"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6(d []interface{}) edpt.RouteMapMatchIpv6 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Address1 = getObjectRouteMapMatchIpv6Address1(in["address_1"].([]interface{}))
		ret.NextHop1 = getObjectRouteMapMatchIpv6NextHop1(in["next_hop_1"].([]interface{}))
		ret.Peer1 = getObjectRouteMapMatchIpv6Peer1(in["peer_1"].([]interface{}))
		ret.Rib = getObjectRouteMapMatchIpv6Rib(in["rib"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpv6Address1(d []interface{}) edpt.RouteMapMatchIpv6Address1 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Address1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.PrefixList2 = getObjectRouteMapMatchIpv6Address1PrefixList2(in["prefix_list_2"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchIpv6Address1PrefixList2(d []interface{}) edpt.RouteMapMatchIpv6Address1PrefixList2 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Address1PrefixList2
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6NextHop1(d []interface{}) edpt.RouteMapMatchIpv6NextHop1 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6NextHop1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NextHopAclName = in["next_hop_acl_name"].(string)
		ret.V6Addr = in["v6_addr"].(string)
		ret.PrefixListName = in["prefix_list_name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6Peer1(d []interface{}) edpt.RouteMapMatchIpv6Peer1 {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Peer1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Acl1 = in["acl1"].(int)
		ret.Acl2 = in["acl2"].(int)
		ret.Name = in["name"].(string)
	}
	return ret
}

func getObjectRouteMapMatchIpv6Rib(d []interface{}) edpt.RouteMapMatchIpv6Rib {

	count1 := len(d)
	var ret edpt.RouteMapMatchIpv6Rib
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exact = in["exact"].(string)
		ret.Reachable = in["reachable"].(string)
		ret.Unreachable = in["unreachable"].(string)
	}
	return ret
}

func getObjectRouteMapMatchLargeCommunity(d []interface{}) edpt.RouteMapMatchLargeCommunity {

	count1 := len(d)
	var ret edpt.RouteMapMatchLargeCommunity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LNameCfg = getObjectRouteMapMatchLargeCommunityLNameCfg(in["l_name_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchLargeCommunityLNameCfg(d []interface{}) edpt.RouteMapMatchLargeCommunityLNameCfg {

	count1 := len(d)
	var ret edpt.RouteMapMatchLargeCommunityLNameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.ExactMatch = in["exact_match"].(int)
	}
	return ret
}

func getObjectRouteMapMatchLocalPreference(d []interface{}) edpt.RouteMapMatchLocalPreference {

	count1 := len(d)
	var ret edpt.RouteMapMatchLocalPreference
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Val = in["val"].(int)
	}
	return ret
}

func getObjectRouteMapMatchMetric(d []interface{}) edpt.RouteMapMatchMetric {

	count1 := len(d)
	var ret edpt.RouteMapMatchMetric
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
	}
	return ret
}

func getObjectRouteMapMatchOrigin(d []interface{}) edpt.RouteMapMatchOrigin {

	count1 := len(d)
	var ret edpt.RouteMapMatchOrigin
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Egp = in["egp"].(int)
		ret.Igp = in["igp"].(int)
		ret.Incomplete = in["incomplete"].(int)
	}
	return ret
}

func getObjectRouteMapMatchRouteType(d []interface{}) edpt.RouteMapMatchRouteType {

	count1 := len(d)
	var ret edpt.RouteMapMatchRouteType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.External = getObjectRouteMapMatchRouteTypeExternal(in["external"].([]interface{}))
	}
	return ret
}

func getObjectRouteMapMatchRouteTypeExternal(d []interface{}) edpt.RouteMapMatchRouteTypeExternal {

	count1 := len(d)
	var ret edpt.RouteMapMatchRouteTypeExternal
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectRouteMapMatchScaleout(d []interface{}) edpt.RouteMapMatchScaleout {

	count1 := len(d)
	var ret edpt.RouteMapMatchScaleout
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClusterId = in["cluster_id"].(int)
		ret.OperationalState = in["operational_state"].(string)
	}
	return ret
}

func getObjectRouteMapMatchTag(d []interface{}) edpt.RouteMapMatchTag {

	count1 := len(d)
	var ret edpt.RouteMapMatchTag
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
	}
	return ret
}

func dataToEndpointRouteMapMatch(d *schema.ResourceData) edpt.RouteMapMatch {
	var ret edpt.RouteMapMatch
	ret.Inst.AsPath = getObjectRouteMapMatchAsPath(d.Get("as_path").([]interface{}))
	ret.Inst.Community = getObjectRouteMapMatchCommunity(d.Get("community").([]interface{}))
	ret.Inst.Extcommunity = getObjectRouteMapMatchExtcommunity(d.Get("extcommunity").([]interface{}))
	ret.Inst.Group = getObjectRouteMapMatchGroup(d.Get("group").([]interface{}))
	ret.Inst.Interface = getObjectRouteMapMatchInterface(d.Get("interface").([]interface{}))
	ret.Inst.Ip = getObjectRouteMapMatchIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getObjectRouteMapMatchIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.LargeCommunity = getObjectRouteMapMatchLargeCommunity(d.Get("large_community").([]interface{}))
	ret.Inst.LocalPreference = getObjectRouteMapMatchLocalPreference(d.Get("local_preference").([]interface{}))
	ret.Inst.Metric = getObjectRouteMapMatchMetric(d.Get("metric").([]interface{}))
	ret.Inst.Origin = getObjectRouteMapMatchOrigin(d.Get("origin").([]interface{}))
	ret.Inst.RouteType = getObjectRouteMapMatchRouteType(d.Get("route_type").([]interface{}))
	ret.Inst.Scaleout = getObjectRouteMapMatchScaleout(d.Get("scaleout").([]interface{}))
	ret.Inst.Tag = getObjectRouteMapMatchTag(d.Get("tag").([]interface{}))
	//omit uuid
	ret.Inst.Sequence = d.Get("sequence").(string)
	ret.Inst.Action = d.Get("action").(string)
	return ret
}
