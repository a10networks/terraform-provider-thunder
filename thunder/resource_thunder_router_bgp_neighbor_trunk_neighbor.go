package thunder

//Thunder resource RouterBgpNeighborTrunkNeighbor

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouterBgpNeighborTrunkNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNeighborTrunkNeighborCreate,
		Update: resourceRouterBgpNeighborTrunkNeighborUpdate,
		Read:   resourceRouterBgpNeighborTrunkNeighborRead,
		Delete: resourceRouterBgpNeighborTrunkNeighborDelete,
		Schema: map[string]*schema.Schema{
			"trunk": {
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

func resourceRouterBgpNeighborTrunkNeighborCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborTrunkNeighbor (Inside resourceRouterBgpNeighborTrunkNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("trunk").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborTrunkNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborTrunkNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		go_thunder.PostRouterBgpNeighborTrunkNeighbor(client.Token, name, data, client.Host)

		return resourceRouterBgpNeighborTrunkNeighborRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNeighborTrunkNeighborRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpNeighborTrunkNeighbor (Inside resourceRouterBgpNeighborTrunkNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborTrunkNeighbor(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpNeighborTrunkNeighborUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborTrunkNeighbor   (Inside resourceRouterBgpNeighborTrunkNeighborUpdate) ")
		data := dataToRouterBgpNeighborTrunkNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborTrunkNeighbor ")
		go_thunder.PutRouterBgpNeighborTrunkNeighbor(client.Token, name1, name2, data, client.Host)

		return resourceRouterBgpNeighborTrunkNeighborRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNeighborTrunkNeighborDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborTrunkNeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborTrunkNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToRouterBgpNeighborTrunkNeighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborTrunkNeighbor {
	var vc go_thunder.RouterBgpNeighborTrunkNeighbor
	var c go_thunder.RouterBgpNeighborTrunkNeighborInstance
	c.Trunk = d.Get("trunk").(int)
	c.Unnumbered = d.Get("unnumbered").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Trunk = c
	return vc
}
