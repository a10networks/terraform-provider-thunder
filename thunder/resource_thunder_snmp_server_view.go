package thunder

//Thunder resource SnmpServerView

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"util"
)

func resourceSnmpServerView() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerViewCreate,
		Update: resourceSnmpServerViewUpdate,
		Read:   resourceSnmpServerViewRead,
		Delete: resourceSnmpServerViewDelete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"oid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"viewname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerViewCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerView (Inside resourceSnmpServerViewCreate) ")
		name2 := d.Get("oid").(string)
		name1 := d.Get("viewname").(string)
		data := dataToSnmpServerView(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerView --")
		d.SetId(name1 + "," + name2)
		go_thunder.PostSnmpServerView(client.Token, data, client.Host)

		return resourceSnmpServerViewRead(d, meta)

	}
	return nil
}

func resourceSnmpServerViewRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerView (Inside resourceSnmpServerViewRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerView(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerViewUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerView   (Inside resourceSnmpServerViewUpdate) ")
		data := dataToSnmpServerView(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerView ")
		go_thunder.PutSnmpServerView(client.Token, name1, name2, data, client.Host)

		return resourceSnmpServerViewRead(d, meta)

	}
	return nil
}

func resourceSnmpServerViewDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerViewDelete) " + name1)
		err := go_thunder.DeleteSnmpServerView(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerView(d *schema.ResourceData) go_thunder.SnmpServerView {
	var vc go_thunder.SnmpServerView
	var c go_thunder.SnmpServerViewInstance
	c.Type = d.Get("type").(string)
	c.Oid = d.Get("oid").(string)
	c.Mask = d.Get("mask").(string)
	c.Viewname = d.Get("viewname").(string)

	vc.Type = c
	return vc
}
