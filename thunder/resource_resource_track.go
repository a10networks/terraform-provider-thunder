package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceResourceTrack() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_resource_track`: Define a resource track policy template\n\n__PLACEHOLDER__",
		CreateContext: resourceResourceTrackCreate,
		UpdateContext: resourceResourceTrackUpdate,
		ReadContext:   resourceResourceTrackRead,
		DeleteContext: resourceResourceTrackDelete,

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
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
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
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
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
						"gw_ipv4_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gw_ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "IP Address",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
									},
								},
							},
						},
						"gw_ipv6_address_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gw_ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "IPV6 address",
									},
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
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
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "resource track policy template name",
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
									"weight": {
										Type: schema.TypeInt, Optional: true, Description: "The amount the priority will decrease if the route is missing (The amount the priority will decrease if the route is not present)",
									},
									"gateway": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"distance": {
										Type: schema.TypeInt, Optional: true, Description: "Route's administrative distance(default: match any)",
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
									"weight": {
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
			"scaleout_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scaleout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "scaleout tracking",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
						},
					},
				},
			},
			"trunk_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "trunk tracking (trunk id)",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "failover event weight",
						},
						"per_port_weight": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Per port failover weight",
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
						"weight": {
							Type: schema.TypeInt, Optional: true, Description: "The failover event weight",
						},
					},
				},
			},
		},
	}
}
func resourceResourceTrackCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceResourceTrackCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointResourceTrack(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceResourceTrackRead(ctx, d, meta)
	}
	return diags
}

func resourceResourceTrackUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceResourceTrackUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointResourceTrack(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceResourceTrackRead(ctx, d, meta)
	}
	return diags
}
func resourceResourceTrackDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceResourceTrackDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointResourceTrack(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceResourceTrackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceResourceTrackRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointResourceTrack(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectResourceTrackBgp(d []interface{}) edpt.ResourceTrackBgp {

	count1 := len(d)
	var ret edpt.ResourceTrackBgp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BgpIpv4AddressCfg = getSliceResourceTrackBgpBgpIpv4AddressCfg(in["bgp_ipv4_address_cfg"].([]interface{}))
		ret.BgpIpv6AddressCfg = getSliceResourceTrackBgpBgpIpv6AddressCfg(in["bgp_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceResourceTrackBgpBgpIpv4AddressCfg(d []interface{}) []edpt.ResourceTrackBgpBgpIpv4AddressCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackBgpBgpIpv4AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackBgpBgpIpv4AddressCfg
		oi.BgpIpv4Address = in["bgp_ipv4_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceResourceTrackBgpBgpIpv6AddressCfg(d []interface{}) []edpt.ResourceTrackBgpBgpIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackBgpBgpIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackBgpBgpIpv6AddressCfg
		oi.BgpIpv6Address = in["bgp_ipv6_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectResourceTrackGateway(d []interface{}) edpt.ResourceTrackGateway {

	count1 := len(d)
	var ret edpt.ResourceTrackGateway
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GwIpv4AddressCfg = getSliceResourceTrackGatewayGwIpv4AddressCfg(in["gw_ipv4_address_cfg"].([]interface{}))
		ret.GwIpv6AddressCfg = getSliceResourceTrackGatewayGwIpv6AddressCfg(in["gw_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceResourceTrackGatewayGwIpv4AddressCfg(d []interface{}) []edpt.ResourceTrackGatewayGwIpv4AddressCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackGatewayGwIpv4AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackGatewayGwIpv4AddressCfg
		oi.GwIpv4Address = in["gw_ipv4_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceResourceTrackGatewayGwIpv6AddressCfg(d []interface{}) []edpt.ResourceTrackGatewayGwIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackGatewayGwIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackGatewayGwIpv6AddressCfg
		oi.GwIpv6Address = in["gw_ipv6_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceResourceTrackInterface(d []interface{}) []edpt.ResourceTrackInterface {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackInterface, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackInterface
		oi.Ethernet = in["ethernet"].(int)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectResourceTrackRoute(d []interface{}) edpt.ResourceTrackRoute {

	count1 := len(d)
	var ret edpt.ResourceTrackRoute
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDestinationCfg = getSliceResourceTrackRouteIpDestinationCfg(in["ip_destination_cfg"].([]interface{}))
		ret.Ipv6DestinationCfg = getSliceResourceTrackRouteIpv6DestinationCfg(in["ipv6_destination_cfg"].([]interface{}))
	}
	return ret
}

func getSliceResourceTrackRouteIpDestinationCfg(d []interface{}) []edpt.ResourceTrackRouteIpDestinationCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackRouteIpDestinationCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackRouteIpDestinationCfg
		oi.IpDestination = in["ip_destination"].(string)
		oi.Mask = in["mask"].(string)
		oi.Weight = in["weight"].(int)
		oi.Gateway = in["gateway"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceResourceTrackRouteIpv6DestinationCfg(d []interface{}) []edpt.ResourceTrackRouteIpv6DestinationCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackRouteIpv6DestinationCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackRouteIpv6DestinationCfg
		oi.Ipv6Destination = in["ipv6_destination"].(string)
		oi.Weight = in["weight"].(int)
		oi.Gatewayv6 = in["gatewayv6"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectResourceTrackScaleoutCfg(d []interface{}) edpt.ResourceTrackScaleoutCfg {

	count1 := len(d)
	var ret edpt.ResourceTrackScaleoutCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Scaleout = in["scaleout"].(int)
		ret.Weight = in["weight"].(int)
	}
	return ret
}

func getSliceResourceTrackTrunkCfg(d []interface{}) []edpt.ResourceTrackTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackTrunkCfg
		oi.Trunk = in["trunk"].(int)
		oi.Weight = in["weight"].(int)
		oi.PerPortWeight = in["per_port_weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceResourceTrackVlanCfg(d []interface{}) []edpt.ResourceTrackVlanCfg {

	count1 := len(d)
	ret := make([]edpt.ResourceTrackVlanCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ResourceTrackVlanCfg
		oi.Vlan = in["vlan"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointResourceTrack(d *schema.ResourceData) edpt.ResourceTrack {
	var ret edpt.ResourceTrack
	ret.Inst.Bgp = getObjectResourceTrackBgp(d.Get("bgp").([]interface{}))
	ret.Inst.Gateway = getObjectResourceTrackGateway(d.Get("gateway").([]interface{}))
	ret.Inst.Interface = getSliceResourceTrackInterface(d.Get("interface").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Route = getObjectResourceTrackRoute(d.Get("route").([]interface{}))
	ret.Inst.ScaleoutCfg = getObjectResourceTrackScaleoutCfg(d.Get("scaleout_cfg").([]interface{}))
	ret.Inst.TrunkCfg = getSliceResourceTrackTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VlanCfg = getSliceResourceTrackVlanCfg(d.Get("vlan_cfg").([]interface{}))
	return ret
}
