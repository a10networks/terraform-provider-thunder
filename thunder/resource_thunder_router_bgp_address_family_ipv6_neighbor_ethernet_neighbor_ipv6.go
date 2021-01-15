package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create,
		Update: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update,
		Read:   resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read,
		Delete: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete,
		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"peer_group_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		name2 := d.Get("ethernet").(int)
		data := dataToRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		go_thunder.PostRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(client.Token, name, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6   (Inside resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update) ")
		data := dataToRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 ")
		go_thunder.PutRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(client.Token, name1, name2, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(d, meta)
}

func dataToRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6
	var c go_thunder.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Instance
	c.Ethernet = d.Get("ethernet").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ethernet = c
	return vc
}
