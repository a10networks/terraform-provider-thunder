package thunder

//Thunder resource RouterOspfRedistribute

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfRedistribute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterOspfRedistributeCreate,
		UpdateContext: resourceRouterOspfRedistributeUpdate,
		ReadContext:   resourceRouterOspfRedistributeRead,
		DeleteContext: resourceRouterOspfRedistributeDelete,
		Schema: map[string]*schema.Schema{
			"redist_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"metric": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"route_map": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tag": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ospf_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"process_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_ospf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_type_ospf": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"route_map_ospf": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tag_ospf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"metric_ip_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"metric_type_ip_nat": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"route_map_ip_nat": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tag_ip_nat": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ip_nat_floating_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_nat_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ip_nat_floating_ip_forward": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"vip_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type_vip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"metric_vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"metric_type_vip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"route_map_vip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tag_vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"vip_floating_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"vip_floating_ip_forward": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
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
				Type:        schema.TypeInt,
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

func resourceRouterOspfRedistributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterOspfRedistribute (Inside resourceRouterOspfRedistributeCreate) ")
		name1 := d.Get("process_id").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterOspfRedistribute(d)
		logger.Println("[INFO] received formatted data from method data to RouterOspfRedistribute --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostRouterOspfRedistribute(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterOspfRedistributeRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterOspfRedistributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterOspfRedistribute (Inside resourceRouterOspfRedistributeRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterOspfRedistribute(client.Token, name1, client.Host)
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

func resourceRouterOspfRedistributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterOspfRedistributeRead(ctx, d, meta)
}

func resourceRouterOspfRedistributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterOspfRedistributeRead(ctx, d, meta)
}
func dataToRouterOspfRedistribute(d *schema.ResourceData) go_thunder.RouterOspfRedistribute {
	var vc go_thunder.RouterOspfRedistribute
	var c go_thunder.RouterOspfRedistributeInstance

	RedistListCount := d.Get("redist_list.#").(int)
	c.Type = make([]go_thunder.RouterOspfRedistributeRedistList, 0, RedistListCount)

	for i := 0; i < RedistListCount; i++ {
		var obj1 go_thunder.RouterOspfRedistributeRedistList
		prefix1 := fmt.Sprintf("redist_list.%d.", i)
		obj1.Type = d.Get(prefix1 + "type").(string)
		obj1.Metric = d.Get(prefix1 + "metric").(int)
		obj1.MetricType = d.Get(prefix1 + "metric_type").(string)
		obj1.RouteMap = d.Get(prefix1 + "route_map").(string)
		obj1.Tag = d.Get(prefix1 + "tag").(int)
		c.Type = append(c.Type, obj1)
	}

	OspfListCount := d.Get("ospf_list.#").(int)
	c.Ospf = make([]go_thunder.RouterOspfRedistributeOspfList, 0, OspfListCount)

	for i := 0; i < OspfListCount; i++ {
		var obj2 go_thunder.RouterOspfRedistributeOspfList
		prefix2 := fmt.Sprintf("ospf_list.%d.", i)
		obj2.Ospf = d.Get(prefix2 + "ospf").(int)
		obj2.ProcessID = d.Get(prefix2 + "process_id").(int)
		obj2.MetricOspf = d.Get(prefix2 + "metric_ospf").(int)
		obj2.MetricTypeOspf = d.Get(prefix2 + "metric_type_ospf").(string)
		obj2.RouteMapOspf = d.Get(prefix2 + "route_map_ospf").(string)
		obj2.TagOspf = d.Get(prefix2 + "tag_ospf").(int)
		c.Ospf = append(c.Ospf, obj2)
	}

	c.IPNat = d.Get("ip_nat").(int)
	c.MetricIPNat = d.Get("metric_ip_nat").(int)
	c.MetricTypeIPNat = d.Get("metric_type_ip_nat").(string)
	c.RouteMapIPNat = d.Get("route_map_ip_nat").(string)
	c.TagIPNat = d.Get("tag_ip_nat").(int)

	IPNatFloatingListCount := d.Get("ip_nat_floating_list.#").(int)
	c.IPNatPrefix = make([]go_thunder.RouterOspfRedistributeIPNatFloatingList, 0, IPNatFloatingListCount)

	for i := 0; i < IPNatFloatingListCount; i++ {
		var obj3 go_thunder.RouterOspfRedistributeIPNatFloatingList
		prefix3 := fmt.Sprintf("ip_nat_floating_list.%d.", i)
		obj3.IPNatPrefix = d.Get(prefix3 + "ip_nat_prefix").(string)
		obj3.IPNatFloatingIPForward = d.Get(prefix3 + "ip_nat_floating_ip_forward").(string)
		c.IPNatPrefix = append(c.IPNatPrefix, obj3)
	}

	VipListCount := d.Get("vip_list.#").(int)
	c.TypeVip = make([]go_thunder.RouterOspfRedistributeVipList, 0, VipListCount)

	for i := 0; i < VipListCount; i++ {
		var obj4 go_thunder.RouterOspfRedistributeVipList
		prefix4 := fmt.Sprintf("vip_list.%d.", i)
		obj4.TypeVip = d.Get(prefix4 + "type_vip").(string)
		obj4.MetricVip = d.Get(prefix4 + "metric_vip").(int)
		obj4.MetricTypeVip = d.Get(prefix4 + "metric_type_vip").(string)
		obj4.RouteMapVip = d.Get(prefix4 + "route_map_vip").(string)
		obj4.TagVip = d.Get(prefix4 + "tag_vip").(int)
		c.TypeVip = append(c.TypeVip, obj4)
	}

	VipFloatingListCount := d.Get("vip_floating_list.#").(int)
	c.VipAddress = make([]go_thunder.RouterOspfRedistributeVipFloatingList, 0, VipFloatingListCount)

	for i := 0; i < VipFloatingListCount; i++ {
		var obj5 go_thunder.RouterOspfRedistributeVipFloatingList
		prefix5 := fmt.Sprintf("vip_floating_list.%d.", i)
		obj5.VipAddress = d.Get(prefix5 + "vip_address").(string)
		obj5.VipFloatingIPForward = d.Get(prefix5 + "vip_floating_ip_forward").(string)
		c.VipAddress = append(c.VipAddress, obj5)
	}

	vc.RedistList = c
	return vc
}
