package thunder

//Thunder resource RouterBgpAddressFamilyIpv6Redistribute

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6Redistribute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAddressFamilyIpv6RedistributeCreate,
		Update: resourceRouterBgpAddressFamilyIpv6RedistributeUpdate,
		Read:   resourceRouterBgpAddressFamilyIpv6RedistributeRead,
		Delete: resourceRouterBgpAddressFamilyIpv6RedistributeDelete,
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
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6RedistributeCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6Redistribute (Inside resourceRouterBgpAddressFamilyIpv6RedistributeCreate) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpAddressFamilyIpv6Redistribute(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6Redistribute --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostRouterBgpAddressFamilyIpv6Redistribute(client.Token, name, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6RedistributeRead(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6RedistributeRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6Redistribute (Inside resourceRouterBgpAddressFamilyIpv6RedistributeRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6Redistribute(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6RedistributeUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6RedistributeRead(d, meta)
}

func resourceRouterBgpAddressFamilyIpv6RedistributeDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6RedistributeRead(d, meta)
}
func dataToRouterBgpAddressFamilyIpv6Redistribute(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6Redistribute {
	var vc go_thunder.RouterBgpAddressFamilyIpv6Redistribute
	var c go_thunder.RouterBgpAddressFamilyIpv6RedistributeInstance

	var obj1 go_thunder.RouterBgpAddressFamilyIpv6RedistributeConnectedCfg
	prefix1 := "connected_cfg.0."
	obj1.Connected = d.Get(prefix1 + "connected").(int)
	obj1.RouteMap = d.Get(prefix1 + "route_map").(string)

	c.Connected = obj1

	var obj2 go_thunder.RouterBgpAddressFamilyIpv6RedistributeFloatingIPCfg
	prefix2 := "floating_ip_cfg.0."
	obj2.FloatingIP = d.Get(prefix2 + "floating_ip").(int)
	obj2.RouteMap = d.Get(prefix2 + "route_map").(string)

	c.FloatingIP = obj2

	var obj3 go_thunder.RouterBgpAddressFamilyIpv6RedistributeNat64Cfg
	prefix3 := "nat64_cfg.0."
	obj3.Nat64 = d.Get(prefix3 + "nat64").(int)
	obj3.RouteMap = d.Get(prefix3 + "route_map").(string)

	c.Nat64 = obj3

	var obj4 go_thunder.RouterBgpAddressFamilyIpv6RedistributeNatMapCfg
	prefix4 := "nat_map_cfg.0."
	obj4.NatMap = d.Get(prefix4 + "nat_map").(int)
	obj4.RouteMap = d.Get(prefix4 + "route_map").(string)

	c.NatMap = obj4

	var obj5 go_thunder.RouterBgpAddressFamilyIpv6RedistributeLw4O6Cfg
	prefix5 := "lw4o6_cfg.0."
	obj5.Lw4O6 = d.Get(prefix5 + "lw4o6").(int)
	obj5.RouteMap = d.Get(prefix5 + "route_map").(string)

	c.Lw4O6 = obj5

	var obj6 go_thunder.RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg
	prefix6 := "static_nat_cfg.0."
	obj6.StaticNat = d.Get(prefix6 + "static_nat").(int)
	obj6.RouteMap = d.Get(prefix6 + "route_map").(string)

	c.StaticNat = obj6

	var obj7 go_thunder.RouterBgpAddressFamilyIpv6RedistributeIPNatCfg
	prefix7 := "ip_nat_cfg.0."
	obj7.IPNat = d.Get(prefix7 + "ip_nat").(int)
	obj7.RouteMap = d.Get(prefix7 + "route_map").(string)

	c.IPNat = obj7

	var obj8 go_thunder.RouterBgpAddressFamilyIpv6RedistributeIPNatListCfg
	prefix8 := "ip_nat_list_cfg.0."
	obj8.IPNatList = d.Get(prefix8 + "ip_nat_list").(int)
	obj8.RouteMap = d.Get(prefix8 + "route_map").(string)

	c.IPNatList = obj8

	var obj9 go_thunder.RouterBgpAddressFamilyIpv6RedistributeIsisCfg
	prefix9 := "isis_cfg.0."
	obj9.Isis = d.Get(prefix9 + "isis").(int)
	obj9.RouteMap = d.Get(prefix9 + "route_map").(string)

	c.Isis = obj9

	var obj10 go_thunder.RouterBgpAddressFamilyIpv6RedistributeOspfCfg
	prefix10 := "ospf_cfg.0."
	obj10.Ospf = d.Get(prefix10 + "ospf").(int)
	obj10.RouteMap = d.Get(prefix10 + "route_map").(string)

	c.Ospf = obj10

	var obj11 go_thunder.RouterBgpAddressFamilyIpv6RedistributeRipCfg
	prefix11 := "rip_cfg.0."
	obj11.Rip = d.Get(prefix11 + "rip").(int)
	obj11.RouteMap = d.Get(prefix11 + "route_map").(string)

	c.Rip = obj11

	var obj12 go_thunder.RouterBgpAddressFamilyIpv6RedistributeStaticCfg
	prefix12 := "static_cfg.0."
	obj12.Static = d.Get(prefix12 + "static").(int)
	obj12.RouteMap = d.Get(prefix12 + "route_map").(string)

	c.Static = obj12

	var obj13 go_thunder.RouterBgpAddressFamilyIpv6RedistributeVip
	prefix13 := "vip.0."

	var obj13_1 go_thunder.RouterBgpAddressFamilyIpv6RedistributeOnlyFlaggedCfg
	prefix13_1 := prefix13 + "only_flagged_cfg.0."
	obj13_1.OnlyFlagged = d.Get(prefix13_1 + "only_flagged").(int)
	obj13_1.RouteMap = d.Get(prefix13_1 + "route_map").(string)

	obj13.OnlyFlagged = obj13_1

	var obj13_2 go_thunder.RouterBgpAddressFamilyIpv6RedistributeOnlyNotFlaggedCfg
	prefix13_2 := prefix13 + "only_not_flagged_cfg.0."
	obj13_2.OnlyNotFlagged = d.Get(prefix13_2 + "only_not_flagged").(int)
	obj13_2.RouteMap = d.Get(prefix13_2 + "route_map").(string)

	obj13.OnlyNotFlagged = obj13_2

	c.OnlyFlaggedCfg = obj13

	vc.ConnectedCfg = c
	return vc
}
