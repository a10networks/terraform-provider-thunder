package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpRedistribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_redistribute`: Redistribute information from another routing protocol\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpRedistributeCreate,
		UpdateContext: resourceRouterBgpRedistributeUpdate,
		ReadContext:   resourceRouterBgpRedistributeRead,
		DeleteContext: resourceRouterBgpRedistributeDelete,

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
func resourceRouterBgpRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpRedistributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpRedistribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpRedistributeRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpRedistributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpRedistribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpRedistributeRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpRedistributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpRedistribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpRedistributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpRedistribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRouterBgpRedistributeConnectedCfg(d []interface{}) edpt.RouterBgpRedistributeConnectedCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeConnectedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Connected = in["connected"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeFloatingIpCfg(d []interface{}) edpt.RouterBgpRedistributeFloatingIpCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeFloatingIpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FloatingIp = in["floating_ip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatCfg(d []interface{}) edpt.RouterBgpRedistributeIpNatCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIpNatCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNat = in["ip_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIpNatListCfg(d []interface{}) edpt.RouterBgpRedistributeIpNatListCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIpNatListCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpNatList = in["ip_nat_list"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeIsisCfg(d []interface{}) edpt.RouterBgpRedistributeIsisCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeIsisCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = in["isis"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeLw4o6Cfg(d []interface{}) edpt.RouterBgpRedistributeLw4o6Cfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeLw4o6Cfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lw4o6 = in["lw4o6"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeNatMapCfg(d []interface{}) edpt.RouterBgpRedistributeNatMapCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeNatMapCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NatMap = in["nat_map"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeOspfCfg(d []interface{}) edpt.RouterBgpRedistributeOspfCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeOspfCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeRipCfg(d []interface{}) edpt.RouterBgpRedistributeRipCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeRipCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticCfg(d []interface{}) edpt.RouterBgpRedistributeStaticCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeStaticCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Static = in["static"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeStaticNatCfg(d []interface{}) edpt.RouterBgpRedistributeStaticNatCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeStaticNatCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNat = in["static_nat"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVip(d []interface{}) edpt.RouterBgpRedistributeVip {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVip
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyFlaggedCfg(in["only_flagged_cfg"].([]interface{}))
		ret.OnlyNotFlaggedCfg = getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg(in["only_not_flagged_cfg"].([]interface{}))
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyFlaggedCfg(d []interface{}) edpt.RouterBgpRedistributeVipOnlyFlaggedCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVipOnlyFlaggedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyFlagged = in["only_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func getObjectRouterBgpRedistributeVipOnlyNotFlaggedCfg(d []interface{}) edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg {

	count1 := len(d)
	var ret edpt.RouterBgpRedistributeVipOnlyNotFlaggedCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OnlyNotFlagged = in["only_not_flagged"].(int)
		ret.RouteMap = in["route_map"].(string)
	}
	return ret
}

func dataToEndpointRouterBgpRedistribute(d *schema.ResourceData) edpt.RouterBgpRedistribute {
	var ret edpt.RouterBgpRedistribute
	ret.Inst.ConnectedCfg = getObjectRouterBgpRedistributeConnectedCfg(d.Get("connected_cfg").([]interface{}))
	ret.Inst.FloatingIpCfg = getObjectRouterBgpRedistributeFloatingIpCfg(d.Get("floating_ip_cfg").([]interface{}))
	ret.Inst.IpNatCfg = getObjectRouterBgpRedistributeIpNatCfg(d.Get("ip_nat_cfg").([]interface{}))
	ret.Inst.IpNatListCfg = getObjectRouterBgpRedistributeIpNatListCfg(d.Get("ip_nat_list_cfg").([]interface{}))
	ret.Inst.IsisCfg = getObjectRouterBgpRedistributeIsisCfg(d.Get("isis_cfg").([]interface{}))
	ret.Inst.Lw4o6Cfg = getObjectRouterBgpRedistributeLw4o6Cfg(d.Get("lw4o6_cfg").([]interface{}))
	ret.Inst.NatMapCfg = getObjectRouterBgpRedistributeNatMapCfg(d.Get("nat_map_cfg").([]interface{}))
	ret.Inst.OspfCfg = getObjectRouterBgpRedistributeOspfCfg(d.Get("ospf_cfg").([]interface{}))
	ret.Inst.RipCfg = getObjectRouterBgpRedistributeRipCfg(d.Get("rip_cfg").([]interface{}))
	ret.Inst.StaticCfg = getObjectRouterBgpRedistributeStaticCfg(d.Get("static_cfg").([]interface{}))
	ret.Inst.StaticNatCfg = getObjectRouterBgpRedistributeStaticNatCfg(d.Get("static_nat_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Vip = getObjectRouterBgpRedistributeVip(d.Get("vip").([]interface{}))
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
