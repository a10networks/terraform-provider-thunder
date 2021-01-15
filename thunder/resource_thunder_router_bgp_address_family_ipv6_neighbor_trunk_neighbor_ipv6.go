package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Create,
		Update: resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Update,
		Read:   resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read,
		Delete: resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Delete,
		Schema: map[string]*schema.Schema{
			"trunk": {
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

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Create) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("trunk").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6 --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		go_thunder.PostRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(client.Token, name, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Update(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6   (Inside resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Update) ")
		data := dataToRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6 ")
		go_thunder.PutRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(client.Token, name1, name2, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Read(d, meta)

}

func dataToRouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6 {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6
	var c go_thunder.RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6Instance
	c.Trunk = d.Get("trunk").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Trunk = c
	return vc
}
