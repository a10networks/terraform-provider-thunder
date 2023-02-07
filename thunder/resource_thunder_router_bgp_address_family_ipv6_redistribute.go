package thunder

//Thunder resource RouterBgpAddressFamilyIpv6Redistribute

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6Redistribute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpAddressFamilyIpv6RedistributeCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6RedistributeUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6RedistributeRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6RedistributeDelete,
		Schema: map[string]*schema.Schema{
			"connected_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connected": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"floating_ip_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"floating_ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"nat64_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat64": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"nat_map_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nat_map": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"lw4o6_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lw4o6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"static_nat_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip_nat_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_nat": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip_nat_list_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_nat_list": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"isis_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isis": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ospf_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"rip_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"static_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"vip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"only_flagged_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_flagged": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"only_not_flagged_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"only_not_flagged": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"route_map": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6RedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6Redistribute (Inside resourceRouterBgpAddressFamilyIpv6RedistributeCreate) ")
		name1 := d.Get("as_number").(int)
		data := dataToRouterBgpAddressFamilyIpv6Redistribute(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6Redistribute --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostRouterBgpAddressFamilyIpv6Redistribute(client.Token, strconv.Itoa(name1), data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6Redistribute (Inside resourceRouterBgpAddressFamilyIpv6RedistributeRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6Redistribute(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6RedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx, d, meta)
}

func resourceRouterBgpAddressFamilyIpv6RedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6RedistributeRead(ctx, d, meta)
}
func dataToRouterBgpAddressFamilyIpv6Redistribute(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6Redistribute {
	var vc go_thunder.RouterBgpAddressFamilyIpv6Redistribute
	var c go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstance

	var obj1 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfg
	prefix1 := "connected_cfg.0."
	obj1.RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfgConnected = d.Get(prefix1 + "connected").(int)
	obj1.RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfgRouteMap = d.Get(prefix1 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfgConnected = obj1

	var obj2 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfg
	prefix2 := "floating_ip_cfg.0."
	obj2.RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfgFloatingIP = d.Get(prefix2 + "floating_ip").(int)
	obj2.RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfgRouteMap = d.Get(prefix2 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceFloatingIPCfgFloatingIP = obj2

	var obj3 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceNat64Cfg
	prefix3 := "nat64_cfg.0."
	obj3.RouterBgpAddressFamilyIpv6RedistributeInstanceNat64CfgNat64 = d.Get(prefix3 + "nat64").(int)
	obj3.RouterBgpAddressFamilyIpv6RedistributeInstanceNat64CfgRouteMap = d.Get(prefix3 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceNat64CfgNat64 = obj3

	var obj4 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfg
	prefix4 := "nat_map_cfg.0."
	obj4.RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfgNatMap = d.Get(prefix4 + "nat_map").(int)
	obj4.RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfgRouteMap = d.Get(prefix4 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceNatMapCfgNatMap = obj4

	var obj5 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6Cfg
	prefix5 := "lw4o6_cfg.0."
	obj5.RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6CfgLw4O6 = d.Get(prefix5 + "lw4o6").(int)
	obj5.RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6CfgRouteMap = d.Get(prefix5 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceLw4O6CfgLw4O6 = obj5

	var obj6 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfg
	prefix6 := "static_nat_cfg.0."
	obj6.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfgStaticNat = d.Get(prefix6 + "static_nat").(int)
	obj6.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfgRouteMap = d.Get(prefix6 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticNatCfgStaticNat = obj6

	var obj7 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfg
	prefix7 := "ip_nat_cfg.0."
	obj7.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfgIPNat = d.Get(prefix7 + "ip_nat").(int)
	obj7.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfgRouteMap = d.Get(prefix7 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatCfgIPNat = obj7

	var obj8 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfg
	prefix8 := "ip_nat_list_cfg.0."
	obj8.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfgIPNatList = d.Get(prefix8 + "ip_nat_list").(int)
	obj8.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfgRouteMap = d.Get(prefix8 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceIPNatListCfgIPNatList = obj8

	var obj9 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfg
	prefix9 := "isis_cfg.0."
	obj9.RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfgIsis = d.Get(prefix9 + "isis").(int)
	obj9.RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfgRouteMap = d.Get(prefix9 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceIsisCfgIsis = obj9

	var obj10 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfg
	prefix10 := "ospf_cfg.0."
	obj10.RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfgOspf = d.Get(prefix10 + "ospf").(int)
	obj10.RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfgRouteMap = d.Get(prefix10 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceOspfCfgOspf = obj10

	var obj11 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfg
	prefix11 := "rip_cfg.0."
	obj11.RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfgRip = d.Get(prefix11 + "rip").(int)
	obj11.RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfgRouteMap = d.Get(prefix11 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceRipCfgRip = obj11

	var obj12 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfg
	prefix12 := "static_cfg.0."
	obj12.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfgStatic = d.Get(prefix12 + "static").(int)
	obj12.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfgRouteMap = d.Get(prefix12 + "route_map").(string)

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceStaticCfgStatic = obj12

	var obj13 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceVip
	prefix13 := "vip.0."

	var obj13_1 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfg
	prefix13_1 := prefix13 + "only_flagged_cfg.0."
	obj13_1.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfgOnlyFlagged = d.Get(prefix13_1 + "only_flagged").(int)
	obj13_1.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfgRouteMap = d.Get(prefix13_1 + "route_map").(string)

	obj13.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfgOnlyFlagged = obj13_1

	var obj13_2 go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfg
	prefix13_2 := prefix13 + "only_not_flagged_cfg.0."
	obj13_2.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfgOnlyNotFlagged = d.Get(prefix13_2 + "only_not_flagged").(int)
	obj13_2.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfgRouteMap = d.Get(prefix13_2 + "route_map").(string)

	obj13.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyNotFlaggedCfgOnlyNotFlagged = obj13_2

	c.RouterBgpAddressFamilyIpv6RedistributeInstanceVipOnlyFlaggedCfg = obj13

	vc.RouterBgpAddressFamilyIpv6RedistributeInstanceConnectedCfg = c
	return vc
}
