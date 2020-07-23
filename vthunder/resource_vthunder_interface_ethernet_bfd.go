package vthunder

//vThunder resource InterfaceEthernetBFD

import (
	"strconv"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceEthernetBFD() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceEthernetBFDCreate,
		Update: resourceInterfaceEthernetBFDUpdate,
		Read:   resourceInterfaceEthernetBFDRead,
		Delete: resourceInterfaceEthernetBFDDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"multiplier": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"min_rx": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"echo": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"demand": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"authentication": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"encrypted": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"key_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceEthernetBFDCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernetBFD (Inside resourceInterfaceEthernetBFDCreate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernetBFD(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernetBFD --")
		d.SetId(strconv.Itoa(name))
		go_vthunder.PostInterfaceEthernetBFD(client.Token, name, data, client.Host)

		return resourceInterfaceEthernetBFDRead(d, meta)

	}
	return nil
}

func resourceInterfaceEthernetBFDRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading InterfaceEthernetBFD (Inside resourceInterfaceEthernetBFDRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetInterfaceEthernetBFD(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceEthernetBFDUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceEthernetBFDRead(d, meta)
}

func resourceInterfaceEthernetBFDDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceEthernetBFDRead(d, meta)
}
func dataToInterfaceEthernetBFD(d *schema.ResourceData) go_vthunder.EthernetBFD {
	var vc go_vthunder.EthernetBFD
	var c go_vthunder.EthernetBFDInstance

	var obj1 go_vthunder.IntervalCfg
	prefix := "interval_cfg.0."
	obj1.Interval = d.Get(prefix + "interval").(int)
	obj1.MinRx = d.Get(prefix + "min_rx").(int)
	obj1.Multiplier = d.Get(prefix + "multiplier").(int)
	c.Interval = obj1

	var obj2 go_vthunder.Authentication
	prefix = "authentication.0."
	obj2.Encrypted = d.Get(prefix + "encrypted").(string)
	obj2.Password = d.Get(prefix + "password").(string)
	obj2.Method = d.Get(prefix + "method").(string)
	obj2.KeyID = d.Get(prefix + "key_id").(int)
	c.Encrypted = obj2

	c.Echo = d.Get("echo").(int)
	c.Demand = d.Get("demand").(int)

	vc.UUID = c
	return vc
}
