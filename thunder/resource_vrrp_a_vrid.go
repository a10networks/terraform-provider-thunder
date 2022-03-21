package thunder

import (
	"context"
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVrrpAVrid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid`: Specify VRRP-A vrid\n\n",
		CreateContext: resourceVrrpAVridCreate,
		UpdateContext: resourceVrrpAVridUpdate,
		ReadContext:   resourceVrrpAVridRead,
		DeleteContext: resourceVrrpAVridDelete,
		Schema: map[string]*schema.Schema{
			"blade_parameters": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 150, Description: "VRRP-A priorty (Priority, default is 150)",
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"fail_over_policy_template": {
							Type: schema.TypeString, Optional: true, Description: "Apply a fail over policy template (VRRP-A fail over policy template name)",
							ValidateFunc: validation.StringLenBetween(1, 63),
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"tracking_options": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet Interface (Ethernet interface number)",
													ValidateFunc: validation.IntAtLeast(1),
												},
												"priority_cost": {
													Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
													ValidateFunc: validation.IntBetween(1, 255),
												},
											},
										},
									},
									"route": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_destination_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_destination": {
																Type: schema.TypeString, Optional: true, Description: "Destination prefix",
																ValidateFunc: validation.IsIPv4Address,
															},
															"mask": {
																Type: schema.TypeString, Optional: true, Description: "Destination prefix mask",
																ValidateFunc: axapi.IsIpv4NetmaskBrief,
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease if the route is missing (The amount the priority will decrease if the route is not present)",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"gateway": {
																Type: schema.TypeString, Optional: true, Description: "Match the route's gateway (next-hop) (default: match any)",
																ValidateFunc: validation.IsIPv4Address,
															},
															"distance": {
																Type: schema.TypeInt, Optional: true, Description: "Route's administrative distance (default: match any)",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"protocol": {
																Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Match any routing protocol (default); 'static': Match only static routes (added by user); 'dynamic': Match routes added by dynamic routing protocols (e.g. OSPF);",
																ValidateFunc: validation.StringInSlice([]string{"any", "static", "dynamic"}, false),
															},
														},
													},
												},
												"ipv6_destination_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_destination": {
																Type: schema.TypeString, Optional: true, Description: "IPv6 Destination Prefix",
																ValidateFunc: axapi.IsIpv6AddressPlen,
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease if the route is missing (The amount the priority will decrease if the route is not present)",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"gatewayv6": {
																Type: schema.TypeString, Optional: true, Description: "Match the route's gateway (next-hop) (default: match any)",
																ValidateFunc: validation.IsIPv6Address,
															},
															"distance": {
																Type: schema.TypeInt, Optional: true, Description: "Route's administrative distance (default: match any)",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"protocol": {
																Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Match any routing protocol (default); 'static': Match only static routes (added by user); 'dynamic': Match routes added by dynamic routing protocols (e.g. OSPF);",
																ValidateFunc: validation.StringInSlice([]string{"any", "static", "dynamic"}, false),
															},
														},
													},
												},
											},
										},
									},
									"trunk_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "trunk tracking (Trunk Number)",
													ValidateFunc: validation.IntBetween(1, 4096),
												},
												"priority_cost": {
													Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"per_port_pri": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "per port priority",
													ValidateFunc: validation.IntBetween(0, 255),
												},
											},
										},
									},
									"bgp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bgp_ipv4_address_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bgp_ipv4_address": {
																Type: schema.TypeString, Optional: true, Description: "bgp IP Address",
																ValidateFunc: validation.IsIPv4Address,
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
																ValidateFunc: validation.IntBetween(1, 255),
															},
														},
													},
												},
												"bgp_ipv6_address_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bgp_ipv6_address": {
																Type: schema.TypeString, Optional: true, Description: "IPV6 address",
																ValidateFunc: validation.IsIPv6Address,
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
																ValidateFunc: validation.IntBetween(1, 255),
															},
														},
													},
												},
											},
										},
									},
									"vlan_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vlan": {
													Type: schema.TypeInt, Optional: true, Description: "VLAN tracking (VLAN id)",
													ValidateFunc: validation.IntBetween(1, 4094),
												},
												"timeout": {
													Type: schema.TypeInt, Optional: true, Description: "",
													ValidateFunc: validation.IntBetween(2, 600),
												},
												"priority_cost": {
													Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
													ValidateFunc: validation.IntBetween(1, 255),
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"gateway": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_gateway_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_address": {
																Type: schema.TypeString, Required: true, Description: "IP Address",
																ValidateFunc: validation.IsIPv4Address,
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
															},
														},
													},
												},
												"ipv6_gateway_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_address": {
																Type: schema.TypeString, Required: true, Description: "IPV6 address",
																ValidateFunc: validation.IsIPv6Address,
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
																ValidateFunc: validation.IntBetween(1, 255),
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
					},
				},
			},
			"floating_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							ConflictsWith: []string{"floating_ip.0.ip_address_part_cfg"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "IP Address [shared partition only]",
										ValidateFunc: validation.IsIPv4Address,
									},
								},
							},
						},
						"ip_address_part_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							ConflictsWith: []string{"floating_ip.0.ip_address_cfg"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_partition": {
										Type: schema.TypeString, Optional: true, Description: "IP Address [private partition only]",
										ValidateFunc: validation.IsIPv4Address,
									},
								},
							},
						},
						"ipv6_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							ConflictsWith: []string{"floating_ip.0.ipv6_address_part_cfg"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address [shared partition only]",
										ValidateFunc: validation.IsIPv6Address,
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet (for link-local address only)",
										ValidateFunc: validation.IntAtLeast(1),
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (for link-local address only)",
										ValidateFunc: validation.IntBetween(1, 4096),
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "VE interface (for link-local address only)",
										ValidateFunc: validation.IntBetween(2, 4094),
									},
								},
							},
						},
						"ipv6_address_part_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							ConflictsWith: []string{"floating_ip.0.ipv6_address_cfg"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address_partition": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address [private partition only]",
										ValidateFunc: validation.IsIPv6Address,
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet (for link-local address only)",
										ValidateFunc: validation.IntAtLeast(1),
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (for link-local address only)",
										ValidateFunc: validation.IntBetween(1, 4096),
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "VE interface (for link-local address only)",
										ValidateFunc: validation.IntBetween(2, 4094),
									},
								},
							},
						},
					},
				},
			},
			"follow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid_lead": {
							Type: schema.TypeString, Optional: true, Description: "Define a VRRP-A VRID leader",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
					},
				},
			},
			"pair_follow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pair_follow": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow other VRRP-A vrid",
							ValidateFunc: validation.IntBetween(0, 1),
						},
						"vrid_lead": {
							Type: schema.TypeString, Optional: true, Description: "Define a VRRP-A VRID leader",
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
					},
				},
			},
			"preempt_mode": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "preemption threshold (preemption threshhold (0-255), default 0)",
							ValidateFunc: validation.IntBetween(0, 255),
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable preemption",
							ValidateFunc: validation.IntBetween(0, 1),
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'associated_vip_count': Number of vips associated to vrid; 'associated_vport_count': Number of vports associated to vrid; 'associated_natpool_count': Number of nat pools associated to vrid;",
							ValidateFunc: validation.StringInSlice([]string{"all", "associated_vip_count", "associated_vport_count", "associated_natpool_count"}, false),
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_val": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Specify ha VRRP-A vrid",
				ValidateFunc: validation.IntBetween(0, 31),
			},
		},
	}
}

func resourceVrrpAVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj, err := dataToEndpointVrrpAVrid(d, client.Partition == "shared")
		if err != nil {
			logger.Println(obj.Inst.FloatingIp)
			return diag.FromErr(err)
		}
		d.SetId(obj.GetId())
		err = obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj, err := dataToEndpointVrrpAVrid(d, client.Partition == "shared")
		if err != nil {
			return diag.FromErr(err)
		}
		err = obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj, err := dataToEndpointVrrpAVrid(d, client.Partition == "shared")
		if err != nil {
			return diag.FromErr(err)
		}
		err = obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj, err := dataToEndpointVrrpAVrid(d, client.Partition == "shared")
		if err != nil {
			return diag.FromErr(err)
		}
		err = obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAVridBladeParameters(d []interface{}) edpt.VrrpAVridBladeParameters {
	var ret edpt.VrrpAVridBladeParameters
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Priority = in["priority"].(int)
		ret.FailOverPolicyTemplate = in["fail_over_policy_template"].(string)
		//omit uuid
		ret.TrackingOptions = getObjectVrrpAVridBladeParametersTrackingOptions(in["tracking_options"].([]interface{}))
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptions(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptions {
	var ret edpt.VrrpAVridBladeParametersTrackingOptions
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Interface = getSliceVrrpAVridBladeParametersTrackingOptionsInterface(in["interface"].([]interface{}))
		ret.Route = getObjectVrrpAVridBladeParametersTrackingOptionsRoute(in["route"].([]interface{}))
		ret.TrunkCfg = getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg(in["trunk_cfg"].([]interface{}))
		ret.Bgp = getObjectVrrpAVridBladeParametersTrackingOptionsBgp(in["bgp"].([]interface{}))
		ret.VlanCfg = getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg(in["vlan_cfg"].([]interface{}))
		//omit uuid
		ret.Gateway = getObjectVrrpAVridBladeParametersTrackingOptionsGateway(in["gateway"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsInterface(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsInterface {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsInterface, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsInterface
		oi.Ethernet = in["ethernet"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsRoute(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsRoute {
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsRoute
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpDestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg(in["ip_destination_cfg"].([]interface{}))
		ret.Ipv6DestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg(in["ipv6_destination_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg
		oi.IpDestination = in["ip_destination"].(string)
		oi.Mask = in["mask"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.Gateway = in["gateway"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg
		oi.Ipv6Destination = in["ipv6_destination"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.Gatewayv6 = in["gatewayv6"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg
		oi.Trunk = in["trunk"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.PerPortPri = in["per_port_pri"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsBgp(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsBgp {
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsBgp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.BgpIpv4AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg(in["bgp_ipv4_address_cfg"].([]interface{}))
		ret.BgpIpv6AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg(in["bgp_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg
		oi.BgpIpv4Address = in["bgp_ipv4_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg
		oi.BgpIpv6Address = in["bgp_ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg
		oi.Vlan = in["vlan"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsGateway(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsGateway {
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsGateway
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Ipv4GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList(in["ipv4_gateway_list"].([]interface{}))
		ret.Ipv6GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList(in["ipv6_gateway_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList
		oi.IpAddress = in["ip_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList {
	count := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridFloatingIp(d []interface{}, isShared bool) (edpt.VrrpAVridFloatingIp, error) {
	var err error
	var ret edpt.VrrpAVridFloatingIp
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.IpAddressCfg = getSliceVrrpAVridFloatingIpIpAddressCfg(in["ip_address_cfg"].([]interface{}))
		ret.Ipv6AddressCfg = getSliceVrrpAVridFloatingIpIpv6AddressCfg(in["ipv6_address_cfg"].([]interface{}))
		ret.IpAddressPartCfg = getSliceVrrpAVridFloatingIpIpAddressPartCfg(in["ip_address_part_cfg"].([]interface{}))
		ret.Ipv6AddressPartCfg = getSliceVrrpAVridFloatingIpIpv6AddressPartCfg(in["ipv6_address_part_cfg"].([]interface{}))
		if isShared {
			if len(ret.IpAddressPartCfg) != 0 {
				return ret, errors.New("'ip_address_part_cfg' is unavaible in shared partition")
			}
			if len(ret.Ipv6AddressPartCfg) != 0 {
				return ret, errors.New("'ipv6_address_part_cfg' is unavaible in shared partition")
			}
		} else {
			if len(ret.IpAddressCfg) != 0 {
				return ret, errors.New("'ip_address_cfg' is unavaible in private partition")
			}
			if len(ret.Ipv6AddressCfg) != 0 {
				return ret, errors.New("'ipv6_address_cfg' is unavaible in private partition")
			}
		}
	}
	return ret, err
}

func getSliceVrrpAVridFloatingIpIpAddressCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpAddressCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpAddressCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridFloatingIpIpAddressCfg
		oi.IpAddress = in["ip_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridFloatingIpIpAddressPartCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpAddressPartCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpAddressPartCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridFloatingIpIpAddressPartCfg
		oi.IpAddressPartition = in["ip_address_partition"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridFloatingIpIpv6AddressCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpv6AddressCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpv6AddressCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridFloatingIpIpv6AddressCfg
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridFloatingIpIpv6AddressPartCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpv6AddressPartCfg {
	count := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpv6AddressPartCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridFloatingIpIpv6AddressPartCfg
		oi.Ipv6AddressPartition = in["ipv6_address_partition"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridFollow(d []interface{}) edpt.VrrpAVridFollow {
	var ret edpt.VrrpAVridFollow
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.VridLead = in["vrid_lead"].(string)
	}
	return ret
}

func getObjectVrrpAVridPairFollow(d []interface{}) edpt.VrrpAVridPairFollow {
	var ret edpt.VrrpAVridPairFollow
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.PairFollow = in["pair_follow"].(int)
		ret.VridLead = in["vrid_lead"].(string)
	}
	return ret
}

func getObjectVrrpAVridPreemptMode(d []interface{}) edpt.VrrpAVridPreemptMode {
	var ret edpt.VrrpAVridPreemptMode
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		ret.Threshold = in["threshold"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceVrrpAVridSamplingEnable(d []interface{}) []edpt.VrrpAVridSamplingEnable {
	count := len(d)
	ret := make([]edpt.VrrpAVridSamplingEnable, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAVrid(d *schema.ResourceData, isShared bool) (edpt.VrrpAVrid, error) {
	var ret edpt.VrrpAVrid
	var err error
	ret.Inst.BladeParameters = getObjectVrrpAVridBladeParameters(d.Get("blade_parameters").([]interface{}))
	ret.Inst.FloatingIp, err = getObjectVrrpAVridFloatingIp(d.Get("floating_ip").([]interface{}), isShared)
	if err != nil {
		return ret, err
	}
	_, ok := d.GetOk("follow")
	if isShared && ok {
		return ret, errors.New("'follow' is unavailable for shared partition")
	}
	ret.Inst.Follow = getObjectVrrpAVridFollow(d.Get("follow").([]interface{}))
	ret.Inst.PairFollow = getObjectVrrpAVridPairFollow(d.Get("pair_follow").([]interface{}))
	ret.Inst.PreemptMode = getObjectVrrpAVridPreemptMode(d.Get("preempt_mode").([]interface{}))
	ret.Inst.SamplingEnable = getSliceVrrpAVridSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VridVal = d.Get("vrid_val").(int)
	return ret, nil
}
