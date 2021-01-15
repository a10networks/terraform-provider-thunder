package thunder

//Thunder resource RouterBgpNeighborVeNeighbor

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterBgpNeighborVeNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborVeNeighborCreate,
		Update: resourceRouterBgpNeighborVeNeighborUpdate,
		Read:   resourceRouterBgpNeighborVeNeighborRead,
		Delete: resourceRouterBgpNeighborVeNeighborDelete,
		Schema: map[string]*schema.Schema{
			"ve": {
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
		},
	}
}

func resourceRouterBgpNeighborVeNeighborCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborVeNeighbor (Inside resourceRouterBgpNeighborVeNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("ve").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborVeNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborVeNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		go_thunder.PostRouterBgpNeighborVeNeighbor(client.Token, name, data, client.Host)

		return resourceRouterBgpNeighborVeNeighborRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNeighborVeNeighborRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpNeighborVeNeighbor (Inside resourceRouterBgpNeighborVeNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborVeNeighbor(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpNeighborVeNeighborUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborVeNeighbor   (Inside resourceRouterBgpNeighborVeNeighborUpdate) ")
		data := dataToRouterBgpNeighborVeNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborVeNeighbor ")
		go_thunder.PutRouterBgpNeighborVeNeighbor(client.Token, name1, name2, data, client.Host)

		return resourceRouterBgpNeighborVeNeighborRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNeighborVeNeighborDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborVeNeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborVeNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToRouterBgpNeighborVeNeighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborVeNeighbor {
	var vc go_thunder.RouterBgpNeighborVeNeighbor
	var c go_thunder.RouterBgpNeighborVeNeighborInstance
	c.Ve = d.Get("ve").(int)
	c.Unnumbered = d.Get("unnumbered").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ve = c
	return vc
}
