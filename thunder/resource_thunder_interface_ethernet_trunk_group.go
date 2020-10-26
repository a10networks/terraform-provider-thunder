package thunder

//Thunder resource InterfaceEthernetTrunkGroup

import (
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceEthernetTrunkGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceEthernetTrunkGroupCreate,
		Update: resourceInterfaceEthernetTrunkGroupUpdate,
		Read:   resourceInterfaceEthernetTrunkGroupRead,
		Delete: resourceInterfaceEthernetTrunkGroupDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"udld_timeout_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"slow": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"port_priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"admin_key": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"trunk_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceEthernetTrunkGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernetTrunkGroup (Inside resourceInterfaceEthernetTrunkGroupCreate) ")
		idNum := d.Get("ifnum").(int)
		name := d.Get("trunk_number").(int)
		data := dataToInterfaceEthernetTrunkGroup(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernetTrunkGroup --")
		d.SetId(strconv.Itoa(name) + "," + strconv.Itoa(idNum))
		go_thunder.PostInterfaceEthernetTrunkGroup(client.Token, idNum, data, client.Host)

		return resourceInterfaceEthernetTrunkGroupRead(d, meta)

	}
	return nil
}

func resourceInterfaceEthernetTrunkGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceEthernetTrunkGroup (Inside resourceInterfaceEthernetTrunkGroupRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name := id[0]
		idNum := id[1]
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceEthernetTrunkGroup(client.Token, idNum, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceEthernetTrunkGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying InterfaceEthernetTrunkGroup   (Inside resourceInterfaceEthernetTrunkGroupUpdate) ")
		idNum := d.Get("ifnum").(int)
		name := d.Get("trunk_number").(int)
		data := dataToInterfaceEthernetTrunkGroup(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernetTrunkGroup ")
		d.SetId("strconv.Itoa(name)")
		go_thunder.PutInterfaceEthernetTrunkGroup(client.Token, idNum, name, data, client.Host)

		return resourceInterfaceEthernetTrunkGroupRead(d, meta)

	}
	return nil
}

func resourceInterfaceEthernetTrunkGroupDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name := id[0]
		idNum := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceEthernetTrunkGroupDelete) " + name)
		err := go_thunder.DeleteInterfaceEthernetTrunkGroup(client.Token, idNum, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToInterfaceEthernetTrunkGroup(d *schema.ResourceData) go_thunder.EthernetTrunkGroup {
	var vc go_thunder.EthernetTrunkGroup
	var c go_thunder.EthernetTrunkGroupInstance
	c.TrunkNumber = d.Get("trunk_number").(int)
	c.UserTag = d.Get("user_tag").(string)

	var obj1 go_thunder.UdldTimeoutCfg
	prefix := "udld_timeout_cfg.0."
	obj1.Slow = d.Get(prefix + "slow").(int)
	obj1.Fast = d.Get(prefix + "fast").(int)
	c.Slow = obj1

	c.Mode = d.Get("mode").(string)
	c.Timeout = d.Get("timeout").(string)
	c.Type = d.Get("type").(string)
	c.AdminKey = d.Get("admin_key").(int)
	c.PortPriority = d.Get("port_priority").(int)

	vc.UUID = c
	return vc
}
