package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVrid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid`: Specify VRRP-A vrid\n\n__PLACEHOLDER__",
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
						},
						"fail_over_policy_template": {
							Type: schema.TypeString, Optional: true, Description: "Apply a fail over policy template (VRRP-A fail over policy template name)",
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
												},
												"priority_cost": {
													Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
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
															},
															"mask": {
																Type: schema.TypeString, Optional: true, Description: "Destination prefix mask",
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease if the route is missing (The amount the priority will decrease if the route is not present)",
															},
															"gateway": {
																Type: schema.TypeString, Optional: true, Description: "Match the route's gateway (next-hop) (default: match any)",
															},
															"distance": {
																Type: schema.TypeInt, Optional: true, Description: "Route's administrative distance (default: match any)",
															},
															"protocol": {
																Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Match any routing protocol (default); 'static': Match only static routes (added by user); 'dynamic': Match routes added by dynamic routing protocols (e.g. OSPF);",
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
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease if the route is missing (The amount the priority will decrease if the route is not present)",
															},
															"gatewayv6": {
																Type: schema.TypeString, Optional: true, Description: "Match the route's gateway (next-hop) (default: match any)",
															},
															"distance": {
																Type: schema.TypeInt, Optional: true, Description: "Route's administrative distance (default: match any)",
															},
															"protocol": {
																Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Match any routing protocol (default); 'static': Match only static routes (added by user); 'dynamic': Match routes added by dynamic routing protocols (e.g. OSPF);",
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
												},
												"priority_cost": {
													Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
												},
												"per_port_pri": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "per port priority",
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
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
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
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
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
												},
												"timeout": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"priority_cost": {
													Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
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
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
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
															},
															"priority_cost": {
																Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease",
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type: schema.TypeString, Optional: true, Description: "IP Address",
									},
								},
							},
						},
						"ip_address_part_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_partition": {
										Type: schema.TypeString, Optional: true, Description: "IP Address",
									},
								},
							},
						},
						"ipv6_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet (for link-local address only)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (for link-local address only)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "VE interface (for link-local address only)",
									},
								},
							},
						},
						"ipv6_address_part_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address_partition": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet (for link-local address only)",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (for link-local address only)",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "VE interface (for link-local address only)",
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
						},
						"vrid_lead": {
							Type: schema.TypeString, Optional: true, Description: "Define a VRRP-A VRID leader",
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
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable preemption",
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
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_val": {
				Type: schema.TypeInt, Required: true, Description: "Specify ha VRRP-A vrid",
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
		obj := dataToEndpointVrrpAVrid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVrid(d)
		err := obj.Put(client.Token, client.Host, logger)
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
		obj := dataToEndpointVrrpAVrid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVrid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAVridBladeParameters3644(d []interface{}) edpt.VrrpAVridBladeParameters3644 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParameters3644
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Priority = in["priority"].(int)
		ret.FailOverPolicyTemplate = in["fail_over_policy_template"].(string)
		//omit uuid
		ret.TrackingOptions = getObjectVrrpAVridBladeParametersTrackingOptions3645(in["tracking_options"].([]interface{}))
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptions3645(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptions3645 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptions3645
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interface = getSliceVrrpAVridBladeParametersTrackingOptionsInterface3646(in["interface"].([]interface{}))
		ret.Route = getObjectVrrpAVridBladeParametersTrackingOptionsRoute3647(in["route"].([]interface{}))
		ret.TrunkCfg = getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg3650(in["trunk_cfg"].([]interface{}))
		ret.Bgp = getObjectVrrpAVridBladeParametersTrackingOptionsBgp3651(in["bgp"].([]interface{}))
		ret.VlanCfg = getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg3654(in["vlan_cfg"].([]interface{}))
		//omit uuid
		ret.Gateway = getObjectVrrpAVridBladeParametersTrackingOptionsGateway3655(in["gateway"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsInterface3646(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsInterface3646 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsInterface3646, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsInterface3646
		oi.Ethernet = in["ethernet"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsRoute3647(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsRoute3647 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsRoute3647
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648(in["ip_destination_cfg"].([]interface{}))
		ret.Ipv6DestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649(in["ipv6_destination_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648
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

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649
		oi.Ipv6Destination = in["ipv6_destination"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.Gatewayv6 = in["gatewayv6"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg3650(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg3650 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg3650, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg3650
		oi.Trunk = in["trunk"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.PerPortPri = in["per_port_pri"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsBgp3651(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsBgp3651 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsBgp3651
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BgpIpv4AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652(in["bgp_ipv4_address_cfg"].([]interface{}))
		ret.BgpIpv6AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653(in["bgp_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652
		oi.BgpIpv4Address = in["bgp_ipv4_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653
		oi.BgpIpv6Address = in["bgp_ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg3654(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg3654 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg3654, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg3654
		oi.Vlan = in["vlan"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsGateway3655(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsGateway3655 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsGateway3655
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656(in["ipv4_gateway_list"].([]interface{}))
		ret.Ipv6GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657(in["ipv6_gateway_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656
		oi.IpAddress = in["ip_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridFloatingIp(d []interface{}) edpt.VrrpAVridFloatingIp {

	count1 := len(d)
	var ret edpt.VrrpAVridFloatingIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpAddressCfg = getSliceVrrpAVridFloatingIpIpAddressCfg(in["ip_address_cfg"].([]interface{}))
		ret.IpAddressPartCfg = getSliceVrrpAVridFloatingIpIpAddressPartCfg(in["ip_address_part_cfg"].([]interface{}))
		ret.Ipv6AddressCfg = getSliceVrrpAVridFloatingIpIpv6AddressCfg(in["ipv6_address_cfg"].([]interface{}))
		ret.Ipv6AddressPartCfg = getSliceVrrpAVridFloatingIpIpv6AddressPartCfg(in["ipv6_address_part_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridFloatingIpIpAddressCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpAddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpAddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridFloatingIpIpAddressCfg
		oi.IpAddress = in["ip_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridFloatingIpIpAddressPartCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpAddressPartCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpAddressPartCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridFloatingIpIpAddressPartCfg
		oi.IpAddressPartition = in["ip_address_partition"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridFloatingIpIpv6AddressCfg(d []interface{}) []edpt.VrrpAVridFloatingIpIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpv6AddressCfg, 0, count1)
	for _, item := range d {
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

	count1 := len(d)
	ret := make([]edpt.VrrpAVridFloatingIpIpv6AddressPartCfg, 0, count1)
	for _, item := range d {
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

	count1 := len(d)
	var ret edpt.VrrpAVridFollow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VridLead = in["vrid_lead"].(string)
	}
	return ret
}

func getObjectVrrpAVridPairFollow(d []interface{}) edpt.VrrpAVridPairFollow {

	count1 := len(d)
	var ret edpt.VrrpAVridPairFollow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PairFollow = in["pair_follow"].(int)
		ret.VridLead = in["vrid_lead"].(string)
	}
	return ret
}

func getObjectVrrpAVridPreemptMode(d []interface{}) edpt.VrrpAVridPreemptMode {

	count1 := len(d)
	var ret edpt.VrrpAVridPreemptMode
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Threshold = in["threshold"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceVrrpAVridSamplingEnable(d []interface{}) []edpt.VrrpAVridSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAVrid(d *schema.ResourceData) edpt.VrrpAVrid {
	var ret edpt.VrrpAVrid
	ret.Inst.BladeParameters = getObjectVrrpAVridBladeParameters3644(d.Get("blade_parameters").([]interface{}))
	ret.Inst.FloatingIp = getObjectVrrpAVridFloatingIp(d.Get("floating_ip").([]interface{}))
	ret.Inst.Follow = getObjectVrrpAVridFollow(d.Get("follow").([]interface{}))
	ret.Inst.PairFollow = getObjectVrrpAVridPairFollow(d.Get("pair_follow").([]interface{}))
	ret.Inst.PreemptMode = getObjectVrrpAVridPreemptMode(d.Get("preempt_mode").([]interface{}))
	ret.Inst.SamplingEnable = getSliceVrrpAVridSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VridVal = d.Get("vrid_val").(int)
	return ret
}
