package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAVridBladeParametersTrackingOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_vrid_blade_parameters_tracking_options`: VRRP-A tracking\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAVridBladeParametersTrackingOptionsCreate,
		UpdateContext: resourceVrrpAVridBladeParametersTrackingOptionsUpdate,
		ReadContext:   resourceVrrpAVridBladeParametersTrackingOptionsRead,
		DeleteContext: resourceVrrpAVridBladeParametersTrackingOptionsDelete,

		Schema: map[string]*schema.Schema{
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"vrid_val": {
				Type: schema.TypeString, Required: true, Description: "VridVal",
			},
		},
	}
}
func resourceVrrpAVridBladeParametersTrackingOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersTrackingOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAVridBladeParametersTrackingOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAVridBladeParametersTrackingOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAVridBladeParametersTrackingOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAVridBladeParametersTrackingOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAVridBladeParametersTrackingOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAVridBladeParametersTrackingOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAVridBladeParametersTrackingOptionsBgp(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsBgp {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsBgp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BgpIpv4AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg(in["bgp_ipv4_address_cfg"].([]interface{}))
		ret.BgpIpv6AddressCfg = getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg(in["bgp_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg
		oi.BgpIpv4Address = in["bgp_ipv4_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg
		oi.BgpIpv6Address = in["bgp_ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsGateway3630(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsGateway3630 {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsGateway3630
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList(in["ipv4_gateway_list"].([]interface{}))
		ret.Ipv6GatewayList = getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList(in["ipv6_gateway_list"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList, 0, count1)
	for _, item := range d {
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

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.PriorityCost = in["priority_cost"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsInterface(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsInterface {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsInterface, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsInterface
		oi.Ethernet = in["ethernet"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAVridBladeParametersTrackingOptionsRoute(d []interface{}) edpt.VrrpAVridBladeParametersTrackingOptionsRoute {

	count1 := len(d)
	var ret edpt.VrrpAVridBladeParametersTrackingOptionsRoute
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg(in["ip_destination_cfg"].([]interface{}))
		ret.Ipv6DestinationCfg = getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg(in["ipv6_destination_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg, 0, count1)
	for _, item := range d {
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

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg, 0, count1)
	for _, item := range d {
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

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsTrunkCfg
		oi.Trunk = in["trunk"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		oi.PerPortPri = in["per_port_pri"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg(d []interface{}) []edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAVridBladeParametersTrackingOptionsVlanCfg
		oi.Vlan = in["vlan"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.PriorityCost = in["priority_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAVridBladeParametersTrackingOptions(d *schema.ResourceData) edpt.VrrpAVridBladeParametersTrackingOptions {
	var ret edpt.VrrpAVridBladeParametersTrackingOptions
	ret.Inst.Bgp = getObjectVrrpAVridBladeParametersTrackingOptionsBgp(d.Get("bgp").([]interface{}))
	ret.Inst.Gateway = getObjectVrrpAVridBladeParametersTrackingOptionsGateway3630(d.Get("gateway").([]interface{}))
	ret.Inst.Interface = getSliceVrrpAVridBladeParametersTrackingOptionsInterface(d.Get("interface").([]interface{}))
	ret.Inst.Route = getObjectVrrpAVridBladeParametersTrackingOptionsRoute(d.Get("route").([]interface{}))
	ret.Inst.TrunkCfg = getSliceVrrpAVridBladeParametersTrackingOptionsTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VlanCfg = getSliceVrrpAVridBladeParametersTrackingOptionsVlanCfg(d.Get("vlan_cfg").([]interface{}))
	ret.Inst.VridVal = d.Get("vrid_val").(string)
	return ret
}
