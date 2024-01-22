package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6Redistribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_redistribute`: Redistribute information from another routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6RedistributeCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6RedistributeUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6RedistributeRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6RedistributeDelete,

		Schema: map[string]*schema.Schema{
			"connected_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connected": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Connected",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"floating_ip_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"floating_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Floating IP",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"ip_nat_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_nat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP NAT",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"ip_nat_list_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_nat_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP NAT list",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"isis_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isis": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ISO IS-IS",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"lw4o6_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lw4o6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "LW4O6 Prefix",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"nat_map_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat_map": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT MAP Prefix",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"nat64_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat64": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT64 Prefix",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"ospf_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"rip_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Routing Information Protocol (RIP)",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"static_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Static routes",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"static_nat_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_nat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Static NAT Prefix",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"only_flagged_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_flagged": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Selected Virtual IP (VIP)",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
						"only_not_flagged_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_not_flagged": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only not flagged",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
								},
							},
						},
					},
				},
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpAddressFamilyIpv6RedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6RedistributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Redistribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6RedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6RedistributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Redistribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6RedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6RedistributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Redistribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6RedistributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6Redistribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeIsisCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nat64 = in["nat64"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeOspfCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeRipCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVip(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVip {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVip
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg(d []interface{}) edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg {

	count1 := len(d)
	var ret edpt.RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func dataToEndpointRouterBgpAddressFamilyIpv6Redistribute(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6Redistribute {
	var ret edpt.RouterBgpAddressFamilyIpv6Redistribute
	ret.Inst.ConnectedCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeConnectedCfg(d.Get("connected_cfg").([]interface{}))
	ret.Inst.FloatingIpCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg(d.Get("floating_ip_cfg").([]interface{}))
	ret.Inst.IpNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatCfg(d.Get("ip_nat_cfg").([]interface{}))
	ret.Inst.IpNatListCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIpNatListCfg(d.Get("ip_nat_list_cfg").([]interface{}))
	ret.Inst.IsisCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeIsisCfg(d.Get("isis_cfg").([]interface{}))
	ret.Inst.Lw4o6Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg(d.Get("lw4o6_cfg").([]interface{}))
	ret.Inst.NatMapCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNatMapCfg(d.Get("nat_map_cfg").([]interface{}))
	ret.Inst.Nat64Cfg = getObjectRouterBgpAddressFamilyIpv6RedistributeNat64Cfg(d.Get("nat64_cfg").([]interface{}))
	ret.Inst.OspfCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeOspfCfg(d.Get("ospf_cfg").([]interface{}))
	ret.Inst.RipCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeRipCfg(d.Get("rip_cfg").([]interface{}))
	ret.Inst.StaticCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticCfg(d.Get("static_cfg").([]interface{}))
	ret.Inst.StaticNatCfg = getObjectRouterBgpAddressFamilyIpv6RedistributeStaticNatCfg(d.Get("static_nat_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Vip = getObjectRouterBgpAddressFamilyIpv6RedistributeVip(d.Get("vip").([]interface{}))
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
