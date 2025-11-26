package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridBladeParameters() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid_blade_parameters`: blade parameters of VRRP-A vrid\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAVridBladeParametersCreate,
		UpdateContext: resourceVrrpAVridBladeParametersUpdate,
		ReadContext:   resourceVrrpAVridBladeParametersRead,
		DeleteContext: resourceVrrpAVridBladeParametersDelete,

		Schema: map[string]*schema.Schema{
			"fail_over_policy_template": {
				Type: schema.TypeString, Optional: true, Description: "Apply a fail over policy template (VRRP-A fail over policy template name)",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Default: 150, Description: "VRRP-A priorty (Priority, default is 150)",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid_val": {
				Type: schema.TypeString, Required: true, Description: "VridVal",
			},
		},
	}
}
func resourceVrrpAVridBladeParametersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParameters(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridBladeParametersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParameters(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAVridBladeParametersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParameters(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridBladeParametersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParameters(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAVridBladeParametersTrackingOptions3730(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptions3730 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptions3730
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interface = getSliceVrrpAVridBladeParametersTrackingOptionsInterface3731(in["interface"].([]interface{}))
		ret.Route = getObjectVrrpAVridBladeParametersTrackingOptionsRoute3732(in["route"].([]interface{}))
		ret.TrunkCfg = getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg3735(in["trunk_cfg"].([]interface{}))
		ret.Bgp = getObjectVrrpAVridBladeParametersTrackingOptionsBgp3736(in["bgp"].([]interface{}))
		ret.VlanCfg = getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg3739(in["vlan_cfg"].([]interface{}))
		//omit uuid
		ret.Gateway = getObjectVrrpAVridBladeParametersTrackingOptionsGateway3740(in["gateway"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsInterface3731(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsInterface3731 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsInterface3731, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsInterface3731
		oi.Ethernet = in["ethernet"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsRoute3732(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsRoute3732 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsRoute3732
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733(in["ip_destination_cfg"].([]interface{}))
		ret.Ipv6DestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734(in["ipv6_destination_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733
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

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734
		oi.Ipv6Destination = in["ipv6_destination"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.Gatewayv6 = in["gatewayv6"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg3735(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg3735 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg3735, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg3735
		oi.Trunk = in["trunk"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.PerPortPri = in["per_port_pri"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsBgp3736(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsBgp3736 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsBgp3736
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BgpIpv4AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737(in["bgp_ipv4_address_cfg"].([]interface{}))
		ret.BgpIpv6AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738(in["bgp_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737
		oi.BgpIpv4Address = in["bgp_ipv4_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738
		oi.BgpIpv6Address = in["bgp_ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg3739(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg3739 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg3739, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg3739
		oi.Vlan = in["vlan"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsGateway3740(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsGateway3740 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsGateway3740
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741(in["ipv4_gateway_list"].([]interface{}))
		ret.Ipv6GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742(in["ipv6_gateway_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741
		oi.IpAddress = in["ip_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742 {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAVridBladeParameters(d *schema.ResourceData) edpt.VrrpAVridBladeParameters {
	var ret edpt.VrrpAVridBladeParameters
	ret.Inst.FailOverPolicyTemplate = d.Get("fail_over_policy_template").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.TrackingOptions = getObjectVrrpAVridBladeParametersTrackingOptions3730(d.Get("tracking_options").([]interface{}))
	//omit uuid
	ret.Inst.VridVal = d.Get("vrid_val").(string)
	return ret
}
