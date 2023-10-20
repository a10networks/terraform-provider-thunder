package thunder

//Thunder resource - Rib Route

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRibRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRibRouteCreate,
		UpdateContext: resourceRibRouteUpdate,
		ReadContext:   resourceRibRouteRead,
		DeleteContext: resourceRibRouteDelete,

		Schema: map[string]*schema.Schema{
			"ip_nexthop_ipv4": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance_nexthop_ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"description_nexthop_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ip_next_hop": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip_nexthop_partition": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"partition_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"vrid_num_in_partition": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"description_partition_vrid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"description_nexthop_partition": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_dest_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_nexthop_lif": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lif": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"description_nexthop_lif": {
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
			"ip_nexthop_tunnel": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_next_hop_tunnel": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"description_nexthop_tunnel": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"distance_nexthop_tunnel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tunnel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceRibRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RibRoute (Inside resourceRibRouteCreate    ") // + name)
		ribRoute := dataToRibRoute(d)
		logger.Println("[INFO] received formatted data from method data to RibRoute --")
		d.SetId(ribRoute.UUID.Instance)
		err := go_thunder.PostRibRoute(client.Token, ribRoute, client.Host, ribRoute.UUID.Instance)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRibRouteRead(ctx, d, meta)
	}
	return diags
}

func resourceRibRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	var diags diag.Diagnostics
	logger.Println("[INFO] Reading Route (Inside resourceRibRouteRead)")

	if client.Host != "" {
		client := meta.(Thunder)

		name := d.Id()

		logger.Println("[INFO] Fetching route Read" + name)

		rib, err := go_thunder.GetRibRoute(client.Token, client.Host, name)
		if err != nil {
			return diag.FromErr(err)
		}
		if rib == nil {
			logger.Println("[INFO] No Route found " + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceRibRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		ribRoute := dataToRibRoute(d)
		logger.Println("[INFO] received formatted data from method data to RibRoute --")
		d.SetId(ribRoute.UUID.Instance)
		err := go_thunder.PutRibRoute(client.Token, ribRoute.UUID.Instance, ribRoute, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRibRouteRead(ctx, d, meta)
	}
	return diags
}

func resourceRibRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting route (Inside resourceRibRouteDelete) " + name)

		err := go_thunder.DeleteRibRoute(client.Token, name, client.Host)
		if err != nil {
			logger.Println("[ERROR] Unable to Delete service group", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

// Utility method to instantiate Rib Route structure
func dataToRibRoute(d *schema.ResourceData) go_thunder.Rib {

	logger := util.GetLoggerInstance()
	var r go_thunder.Rib
	var rInst go_thunder.RibInst

	ipNextHopCount := d.Get("ip_nexthop_ipv4.#").(int)

	rInst.IPNextHop = make([]go_thunder.IPNexthopIpv4, 0, ipNextHopCount)

	logger.Println("[INFO] IP Next hop count - " + strconv.Itoa(ipNextHopCount))

	for i := 0; i < ipNextHopCount; i++ {
		var hop go_thunder.IPNexthopIpv4
		prefix := fmt.Sprintf("ip_nexthop_ipv4.%d", i)
		logger.Println("[INFO] distance_nexthop_ip: " + strconv.Itoa(d.Get(prefix+".distance_nexthop_ip").(int)))
		hop.DistanceNexthopIP = d.Get(prefix + ".distance_nexthop_ip").(int)
		logger.Println("[INFO] ip_next_hop: " + d.Get(prefix+".ip_next_hop").(string))
		hop.IPNextHop = d.Get(prefix + ".ip_next_hop").(string)
		rInst.IPNextHop = append(rInst.IPNextHop, hop)
	}

	rInst.Instance = d.Get("ip_dest_addr").(string) + "+/" + d.Get("ip_mask").(string) //strings.Replace(mask, "/", "", -1)
	rInst.IPDestAddr = d.Get("ip_dest_addr").(string)
	rInst.IPMask = d.Get("ip_mask").(string)

	r.UUID = rInst

	return r
}
