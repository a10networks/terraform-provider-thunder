package thunder

//Thunder resource SnmpServerHostHostName

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
	"util"
)

func resourceSnmpServerHostHostName() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerHostHostNameCreate,
		Update: resourceSnmpServerHostHostNameUpdate,
		Read:   resourceSnmpServerHostHostNameRead,
		Delete: resourceSnmpServerHostHostNameDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"udp_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"v1_v2c_comm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerHostHostNameCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerHostHostName (Inside resourceSnmpServerHostHostNameCreate) ")
		name2 := d.Get("version").(string)
		name1 := d.Get("hostname").(string)
		data := dataToSnmpServerHostHostName(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerHostHostName --")
		d.SetId(name1 + "," + name2)
		go_thunder.PostSnmpServerHostHostName(client.Token, data, client.Host)

		return resourceSnmpServerHostHostNameRead(d, meta)

	}
	return nil
}

func resourceSnmpServerHostHostNameRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerHostHostName (Inside resourceSnmpServerHostHostNameRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSnmpServerHostHostName(client.Token, name1, name2, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerHostHostNameUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying SnmpServerHostHostName   (Inside resourceSnmpServerHostHostNameUpdate) ")
		data := dataToSnmpServerHostHostName(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerHostHostName ")
		go_thunder.PutSnmpServerHostHostName(client.Token, name1, name2, data, client.Host)

		return resourceSnmpServerHostHostNameRead(d, meta)

	}
	return nil
}

func resourceSnmpServerHostHostNameDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceSnmpServerHostHostNameDelete) " + name1)
		err := go_thunder.DeleteSnmpServerHostHostName(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToSnmpServerHostHostName(d *schema.ResourceData) go_thunder.SnmpServerHostHostName {
	var vc go_thunder.SnmpServerHostHostName
	var c go_thunder.SnmpServerHostHostNameInstance
	c.Hostname = d.Get("hostname").(string)
	c.UDPPort = d.Get("udp_port").(int)
	c.V1V2CComm = d.Get("v1_v2c_comm").(string)
	c.User = d.Get("user").(string)
	c.Version = d.Get("version").(string)

	vc.UUID = c
	return vc
}
