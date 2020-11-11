package thunder

//Thunder resource SnmpServerGroup

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerGroupCreate,
		Update: resourceSnmpServerGroupUpdate,
		Read:   resourceSnmpServerGroupRead,
		Delete: resourceSnmpServerGroupDelete,
		Schema: map[string]*schema.Schema{
			"read": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"groupname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"v3": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerGroup (Inside resourceSnmpServerGroupCreate) ")
		name1 := d.Get("groupname").(string)
		data := dataToSnmpServerGroup(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerGroup --")
		d.SetId(name1)
		go_thunder.PostSnmpServerGroup(client.Token, data, client.Host)

		return resourceSnmpServerGroupRead(d, meta)

	}
	return nil
}

func resourceSnmpServerGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerGroup (Inside resourceSnmpServerGroupRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerGroup(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SnmpServerGroup   (Inside resourceSnmpServerGroupUpdate) ")
		data := dataToSnmpServerGroup(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerGroup ")
		go_thunder.PutSnmpServerGroup(client.Token, name1, data, client.Host)

		return resourceSnmpServerGroupRead(d, meta)

	}
	return nil
}

func resourceSnmpServerGroupDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerGroupDelete) " + name1)
		err := go_thunder.DeleteSnmpServerGroup(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerGroup(d *schema.ResourceData) go_thunder.SnmpServerGroup {
	var vc go_thunder.SnmpServerGroup
	var c go_thunder.SnmpServerGroupInstance
	c.Read = d.Get("read").(string)
	c.Groupname = d.Get("groupname").(string)
	c.V3 = d.Get("v3").(string)

	vc.Read = c
	return vc
}
