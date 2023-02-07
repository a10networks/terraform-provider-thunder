package thunder

//Thunder resource RouterBgpRedistribute

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpRedistribute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpRedistributeCreate,
		UpdateContext: resourceRouterBgpRedistributeUpdate,
		ReadContext:   resourceRouterBgpRedistributeRead,
		DeleteContext: resourceRouterBgpRedistributeDelete,
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

func resourceRouterBgpRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpRedistribute (Inside resourceRouterBgpRedistributeCreate) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpRedistribute(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpRedistribute --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostRouterBgpRedistribute(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpRedistributeRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpRedistribute (Inside resourceRouterBgpRedistributeRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpRedistribute(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceRouterBgpRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpRedistributeRead(ctx, d, meta)
}

func resourceRouterBgpRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpRedistributeRead(ctx, d, meta)
}
func dataToRouterBgpRedistribute(d *schema.ResourceData) go_thunder.RouterBgpRedistribute {
	var vc go_thunder.RouterBgpRedistribute
	var c go_thunder.RouterBgpRedistributeInstance

	var obj1 go_thunder.RouterBgpRedistributeConnectedCfg
	prefix1 := "connected_cfg.0."
	obj1.Connected = d.Get(prefix1 + "connected").(int)
	obj1.RouteMap = d.Get(prefix1 + "route_map").(string)

	c.Connected = obj1

	var obj2 go_thunder.RouterBgpRedistributeFloatingIPCfg
	prefix2 := "floating_ip_cfg.0."
	obj2.FloatingIP = d.Get(prefix2 + "floating_ip").(int)
	obj2.RouteMap = d.Get(prefix2 + "route_map").(string)

	c.FloatingIP = obj2

	var obj3 go_thunder.RouterBgpRedistributeLw4O6Cfg
	prefix3 := "lw4o6_cfg.0."
	obj3.Lw4O6 = d.Get(prefix3 + "lw4o6").(int)
	obj3.RouteMap = d.Get(prefix3 + "route_map").(string)

	c.Lw4O6 = obj3

	var obj4 go_thunder.RouterBgpRedistributeStaticNatCfg
	prefix4 := "static_nat_cfg.0."
	obj4.StaticNat = d.Get(prefix4 + "static_nat").(int)
	obj4.RouteMap = d.Get(prefix4 + "route_map").(string)

	c.StaticNat = obj4

	var obj5 go_thunder.RouterBgpRedistributeIPNatCfg
	prefix5 := "ip_nat_cfg.0."
	obj5.IPNat = d.Get(prefix5 + "ip_nat").(int)
	obj5.RouteMap = d.Get(prefix5 + "route_map").(string)

	c.IPNat = obj5

	var obj6 go_thunder.RouterBgpRedistributeIPNatListCfg
	prefix6 := "ip_nat_list_cfg.0."
	obj6.IPNatList = d.Get(prefix6 + "ip_nat_list").(int)
	obj6.RouteMap = d.Get(prefix6 + "route_map").(string)

	c.IPNatList = obj6

	var obj7 go_thunder.RouterBgpRedistributeIsisCfg
	prefix7 := "isis_cfg.0."
	obj7.Isis = d.Get(prefix7 + "isis").(int)
	obj7.RouteMap = d.Get(prefix7 + "route_map").(string)

	c.Isis = obj7

	var obj8 go_thunder.RouterBgpRedistributeOspfCfg
	prefix8 := "ospf_cfg.0."
	obj8.Ospf = d.Get(prefix8 + "ospf").(int)
	obj8.RouteMap = d.Get(prefix8 + "route_map").(string)

	c.Ospf = obj8

	var obj9 go_thunder.RouterBgpRedistributeRipCfg
	prefix9 := "rip_cfg.0."
	obj9.Rip = d.Get(prefix9 + "rip").(int)
	obj9.RouteMap = d.Get(prefix9 + "route_map").(string)

	c.Rip = obj9

	var obj10 go_thunder.RouterBgpRedistributeStaticCfg
	prefix10 := "static_cfg.0."
	obj10.Static = d.Get(prefix10 + "static").(int)
	obj10.RouteMap = d.Get(prefix10 + "route_map").(string)

	c.Static = obj10

	var obj11 go_thunder.RouterBgpRedistributeNatMapCfg
	prefix11 := "nat_map_cfg.0."
	obj11.NatMap = d.Get(prefix11 + "nat_map").(int)
	obj11.RouteMap = d.Get(prefix11 + "route_map").(string)

	c.NatMap = obj11

	var obj12 go_thunder.RouterBgpRedistributeVip
	prefix12 := "vip.0."

	var obj12_1 go_thunder.RouterBgpRedistributeOnlyFlaggedCfg
	prefix12_1 := prefix12 + "only_flagged_cfg.0."
	obj12_1.OnlyFlagged = d.Get(prefix12_1 + "only_flagged").(int)
	obj12_1.RouteMap = d.Get(prefix12_1 + "route_map").(string)

	obj12.OnlyFlagged = obj12_1

	var obj12_2 go_thunder.RouterBgpRedistributeOnlyNotFlaggedCfg
	prefix12_2 := prefix12 + "only_not_flagged_cfg.0."
	obj12_2.OnlyNotFlagged = d.Get(prefix12_2 + "only_not_flagged").(int)
	obj12_2.RouteMap = d.Get(prefix12_2 + "route_map").(string)

	obj12.OnlyNotFlagged = obj12_2

	c.OnlyFlaggedCfg = obj12

	vc.ConnectedCfg = c
	return vc
}
