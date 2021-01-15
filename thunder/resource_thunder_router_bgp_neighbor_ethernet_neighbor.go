package thunder

//Thunder resource RouterBgpNeighborEthernetNeighbor

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterBgpNeighborEthernetNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborEthernetNeighborCreate,
		Update: resourceRouterBgpNeighborEthernetNeighborUpdate,
		Read:   resourceRouterBgpNeighborEthernetNeighborRead,
		Delete: resourceRouterBgpNeighborEthernetNeighborDelete,
		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"unnumbered": {
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

func resourceRouterBgpNeighborEthernetNeighborCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborEthernetNeighbor (Inside resourceRouterBgpNeighborEthernetNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("ethernet").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborEthernetNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborEthernetNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		go_thunder.PostRouterBgpNeighborEthernetNeighbor(client.Token, name, data, client.Host)

		return resourceRouterBgpNeighborEthernetNeighborRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNeighborEthernetNeighborRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpNeighborEthernetNeighbor (Inside resourceRouterBgpNeighborEthernetNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborEthernetNeighbor(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpNeighborEthernetNeighborUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborEthernetNeighbor   (Inside resourceRouterBgpNeighborEthernetNeighborUpdate) ")
		data := dataToRouterBgpNeighborEthernetNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborEthernetNeighbor ")
		go_thunder.PutRouterBgpNeighborEthernetNeighbor(client.Token, name1, name2, data, client.Host)

		return resourceRouterBgpNeighborEthernetNeighborRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNeighborEthernetNeighborDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborEthernetNeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborEthernetNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToRouterBgpNeighborEthernetNeighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborEthernetNeighbor {
	var vc go_thunder.RouterBgpNeighborEthernetNeighbor
	var c go_thunder.RouterBgpNeighborEthernetNeighborInstance
	c.Ethernet = d.Get("ethernet").(int)
	c.Unnumbered = d.Get("unnumbered").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ethernet = c
	return vc
}
