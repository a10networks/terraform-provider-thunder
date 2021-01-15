package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Create,
		Update: resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Update,
		Read:   resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Read,
		Delete: resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Delete,
		Schema: map[string]*schema.Schema{
			"ve": {
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

func resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Create) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("ve").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6 --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		go_thunder.PostRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6(client.Token, name, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6 (Inside resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Read)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Update(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6   (Inside resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Update) ")
		data := dataToRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6 ")
		go_thunder.PutRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6(client.Token, name1, name2, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Read(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Read(d, meta)

}

func dataToRouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6 {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6
	var c go_thunder.RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6Instance
	c.Ve = d.Get("ve").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ve = c
	return vc
}
