package thunder

//Thunder resource TemplateFix

import (
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateFix() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateFixCreate,
		Update: resourceTemplateFixUpdate,
		Read:   resourceTemplateFixRead,
		Delete: resourceTemplateFixDelete,
		Schema: map[string]*schema.Schema{
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tag_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"equals": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"switching_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"logging": {
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

func resourceTemplateFixCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateFix (Inside resourceTemplateFixCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateFix(d)
		logger.Println("[INFO] received formatted data from method data to TemplateFix --")
		d.SetId(name)
		go_thunder.PostTemplateFix(client.Token, data, client.Host)

		return resourceTemplateFixRead(d, meta)

	}
	return nil
}

func resourceTemplateFixRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading TemplateFix (Inside resourceTemplateFixRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateFix(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateFixUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateFix   (Inside resourceTemplateFixUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateFix(d)
		logger.Println("[INFO] received formatted data from method data to TemplateFix ")
		d.SetId(name)
		go_thunder.PutTemplateFix(client.Token, name, data, client.Host)

		return resourceTemplateFixRead(d, meta)

	}
	return nil
}

func resourceTemplateFixDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateFixDelete) " + name)
		err := go_thunder.DeleteTemplateFix(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateFix(d *schema.ResourceData) go_thunder.Fix {
	var vc go_thunder.Fix
	var c go_thunder.FixInstance

	tagSwitchCount := d.Get("tag_switching.#").(int)
	c.ServiceGroup = make([]go_thunder.TagSwitching, 0, tagSwitchCount)
	for i := 0; i < tagSwitchCount; i++ {
		var ts go_thunder.TagSwitching
		prefix := fmt.Sprintf("tag_switching.%d", i)
		ts.Equals = d.Get(prefix + ".equals").(string)
		ts.ServiceGroup = d.Get(prefix + ".service_group").(string)
		ts.SwitchingType = d.Get(prefix + ".switching_type").(string)
		c.ServiceGroup = append(c.ServiceGroup, ts)
	}

	c.Name = d.Get("name").(string)
	c.Logging = d.Get("logging").(string)
	c.InsertClientIP = d.Get("insert_client_ip").(int)
	c.UserTag = d.Get("user_tag").(string)

	vc.Name = c
	return vc
}
