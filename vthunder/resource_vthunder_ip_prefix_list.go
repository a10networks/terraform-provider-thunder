package vthunder

//vThunder resource IpPrefixList

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpPrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpPrefixListCreate,
		Update: resourceIpPrefixListUpdate,
		Read:   resourceIpPrefixListRead,
		Delete: resourceIpPrefixListDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"le": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipaddr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"any": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"seq": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ge": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpPrefixListCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpPrefixList (Inside resourceIpPrefixListCreate) ")
		name := d.Get("name").(string)
		data := dataToIpPrefixList(d)
		logger.Println("[INFO] received V from method data to IpPrefixList --")
		d.SetId(name)
		go_vthunder.PostIpPrefixList(client.Token, data, client.Host)

		return resourceIpPrefixListRead(d, meta)

	}
	return nil
}

func resourceIpPrefixListRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IpPrefixList (Inside resourceIpPrefixListRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetIpPrefixList(client.Token, name, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpPrefixListUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpPrefixListRead(d, meta)
}

func resourceIpPrefixListDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpPrefixListRead(d, meta)
}
func dataToIpPrefixList(d *schema.ResourceData) go_vthunder.PrefixList {
	var vc go_vthunder.PrefixList
	var c go_vthunder.PrefixListInstance
	c.Name = d.Get("name").(string)

	RulesCount := d.Get("rules.#").(int)
	c.Le = make([]go_vthunder.Rules, 0, RulesCount)

	for i := 0; i < RulesCount; i++ {
		var obj1 go_vthunder.Rules
		prefix := fmt.Sprintf("rules.%d.", i)
		obj1.Seq = d.Get(prefix + "seq").(int)
		obj1.Ipaddr = d.Get(prefix + "ipaddr").(string)
		obj1.Ge = d.Get(prefix + "ge").(int)
		obj1.Any = d.Get(prefix + "any").(int)
		obj1.Description = d.Get(prefix + "description").(string)
		obj1.Action = d.Get(prefix + "action").(string)
		obj1.Le = d.Get(prefix + "le").(int)
		c.Le = append(c.Le, obj1)
	}

	vc.UUID = c
	return vc
}
