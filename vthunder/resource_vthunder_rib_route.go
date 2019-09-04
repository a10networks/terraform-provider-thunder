package vthunder

//vThunder resource - Rib Route

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
	"util"
)

func resourceRibRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRibRouteCreate,
		Update: resourceRibRouteUpdate,
		Read:   resourceRibRouteRead,
		Delete: resourceRibRouteDelete,

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

func resourceRibRouteCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RibRoute (Inside resourceRibRouteCreate    ") // + name)
		ribRoute := dataToRibRoute(d)
		logger.Println("[INFO] received V from method data to RibRoute --")
		d.SetId(ribRoute.UUID.Instance)
		go_vthunder.PostRibRoute(client.Token, ribRoute, client.Host, ribRoute.UUID.Instance)

		return resourceRibRouteRead(d, meta)
	}
	return nil
}

func resourceRibRouteRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Route (Inside resourceRibRouteRead)")

	if client.Host != "" {
		client := meta.(vThunder)
		name := d.Id()

		logger.Println("[INFO] Fetching route Read" + name)

		rib, err := go_vthunder.GetRibRoute(client.Token, client.Host, name)

		if rib == nil {
			logger.Println("[INFO] No Route found " + name)
		}

		return err
	}
	return nil
}

func resourceRibRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		ribRoute := dataToRibRoute(d)
		logger.Println("[INFO] received V from method data to RibRoute --")
		d.SetId(ribRoute.UUID.Instance)
		go_vthunder.PutRibRoute(client.Token, ribRoute.UUID.Instance, ribRoute, client.Host)

		return resourceRibRouteRead(d, meta)
	}
	return nil
}

func resourceRibRouteDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting route (Inside resourceRibRouteDelete) " + name)

		err := go_vthunder.DeleteRibRoute(client.Token, name, client.Host)
		if err != nil {
			logger.Println("[ERROR] Unable to Delete service group", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//Utility method to instantiate Rib Route structure
func dataToRibRoute(d *schema.ResourceData) go_vthunder.Rib {

	logger := util.GetLoggerInstance()
	var r go_vthunder.Rib
	var rInst go_vthunder.RibInst

	ipNextHopCount := d.Get("ip_nexthop_ipv4.#").(int)

	rInst.IPNextHop = make([]go_vthunder.IPNexthopIpv4, 0, ipNextHopCount)

	logger.Println("[INFO] IP Next hop count - " + strconv.Itoa(ipNextHopCount))

	for i := 0; i < ipNextHopCount; i++ {
		var hop go_vthunder.IPNexthopIpv4
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
