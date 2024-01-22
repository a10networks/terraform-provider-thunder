package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAFailOverPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_fail_over_policy_template`: Define a VRRP-A failover policy template\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAFailOverPolicyTemplateCreate,
		UpdateContext: resourceVrrpAFailOverPolicyTemplateUpdate,
		ReadContext:   resourceVrrpAFailOverPolicyTemplateRead,
		DeleteContext: resourceVrrpAFailOverPolicyTemplateDelete,

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
				Type: schema.TypeString, Required: true, Description: "VRRP-A fail over policy template name",
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
func resourceVrrpAFailOverPolicyTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAFailOverPolicyTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAFailOverPolicyTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAFailOverPolicyTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAFailOverPolicyTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAFailOverPolicyTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAFailOverPolicyTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAFailOverPolicyTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAFailOverPolicyTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAFailOverPolicyTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAFailOverPolicyTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAFailOverPolicyTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAFailOverPolicyTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAFailOverPolicyTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVrrpAFailOverPolicyTemplateBgp(d []interface{}) edpt.VrrpAFailOverPolicyTemplateBgp {

	count1 := len(d)
	var ret edpt.VrrpAFailOverPolicyTemplateBgp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BgpIpv4AddressCfg = getSliceVrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg(in["bgp_ipv4_address_cfg"].([]interface{}))
		ret.BgpIpv6AddressCfg = getSliceVrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg(in["bgp_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg
		oi.BgpIpv4Address = in["bgp_ipv4_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg
		oi.BgpIpv6Address = in["bgp_ipv6_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAFailOverPolicyTemplateGateway(d []interface{}) edpt.VrrpAFailOverPolicyTemplateGateway {

	count1 := len(d)
	var ret edpt.VrrpAFailOverPolicyTemplateGateway
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GwIpv4AddressCfg = getSliceVrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg(in["gw_ipv4_address_cfg"].([]interface{}))
		ret.GwIpv6AddressCfg = getSliceVrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg(in["gw_ipv6_address_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg
		oi.GwIpv4Address = in["gw_ipv4_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg
		oi.GwIpv6Address = in["gw_ipv6_address"].(string)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateInterface(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateInterface {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateInterface, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateInterface
		oi.Ethernet = in["ethernet"].(int)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVrrpAFailOverPolicyTemplateRoute(d []interface{}) edpt.VrrpAFailOverPolicyTemplateRoute {

	count1 := len(d)
	var ret edpt.VrrpAFailOverPolicyTemplateRoute
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpDestinationCfg = getSliceVrrpAFailOverPolicyTemplateRouteIpDestinationCfg(in["ip_destination_cfg"].([]interface{}))
		ret.Ipv6DestinationCfg = getSliceVrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg(in["ipv6_destination_cfg"].([]interface{}))
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateRouteIpDestinationCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateRouteIpDestinationCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateRouteIpDestinationCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateRouteIpDestinationCfg
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

func getSliceVrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg
		oi.Ipv6Destination = in["ipv6_destination"].(string)
		oi.Weight = in["weight"].(int)
		oi.Gatewayv6 = in["gatewayv6"].(string)
		oi.Distance = in["distance"].(int)
		oi.Protocol = in["protocol"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateTrunkCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateTrunkCfg
		oi.Trunk = in["trunk"].(int)
		oi.Weight = in["weight"].(int)
		oi.PerPortWeight = in["per_port_weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVrrpAFailOverPolicyTemplateVlanCfg(d []interface{}) []edpt.VrrpAFailOverPolicyTemplateVlanCfg {

	count1 := len(d)
	ret := make([]edpt.VrrpAFailOverPolicyTemplateVlanCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VrrpAFailOverPolicyTemplateVlanCfg
		oi.Vlan = in["vlan"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.Weight = in["weight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVrrpAFailOverPolicyTemplate(d *schema.ResourceData) edpt.VrrpAFailOverPolicyTemplate {
	var ret edpt.VrrpAFailOverPolicyTemplate
	ret.Inst.Bgp = getObjectVrrpAFailOverPolicyTemplateBgp(d.Get("bgp").([]interface{}))
	ret.Inst.Gateway = getObjectVrrpAFailOverPolicyTemplateGateway(d.Get("gateway").([]interface{}))
	ret.Inst.Interface = getSliceVrrpAFailOverPolicyTemplateInterface(d.Get("interface").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Route = getObjectVrrpAFailOverPolicyTemplateRoute(d.Get("route").([]interface{}))
	ret.Inst.TrunkCfg = getSliceVrrpAFailOverPolicyTemplateTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VlanCfg = getSliceVrrpAFailOverPolicyTemplateVlanCfg(d.Get("vlan_cfg").([]interface{}))
	return ret
}
