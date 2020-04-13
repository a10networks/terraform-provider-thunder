package vthunder

//vThunder resource SlbTemplateDynamicService

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateDynamicService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateDynamicServiceCreate,
		Update: resourceSlbTemplateDynamicServiceUpdate,
		Read:   resourceSlbTemplateDynamicServiceRead,
		Delete: resourceSlbTemplateDynamicServiceDelete,
		Schema: map[string]*schema.Schema{
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"dns_server": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_dns_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_dns_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateDynamicServiceCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateDynamicService (Inside resourceSlbTemplateDynamicServiceCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateDynamicService(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDynamicService --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateDynamicService(client.Token, data, client.Host)

		return resourceSlbTemplateDynamicServiceRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateDynamicServiceRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateDynamicService (Inside resourceSlbTemplateDynamicServiceRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateDynamicService(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateDynamicServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateDynamicService   (Inside resourceSlbTemplateDynamicServiceUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateDynamicService(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateDynamicService ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateDynamicService(client.Token, name, data, client.Host)

		return resourceSlbTemplateDynamicServiceRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateDynamicServiceDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateDynamicServiceDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateDynamicService(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateDynamicService(d *schema.ResourceData) go_vthunder.DynamicService {
	var vc go_vthunder.DynamicService
	var c go_vthunder.DynamicServiceInstance

	DNSServerCount := d.Get("dns_server.#").(int)
	c.Ipv6DNSServer = make([]go_vthunder.DNSServer, 0, DNSServerCount)

	for i := 0; i < DNSServerCount; i++ {
		var obj1 go_vthunder.DNSServer
		prefix := fmt.Sprintf("dns_server.%d.", i)
		obj1.Ipv6DNSServer = d.Get(prefix + "ipv6_dns_server").(string)
		obj1.Ipv4DNSServer = d.Get(prefix + "ipv4_dns_server").(string)
		c.Ipv6DNSServer = append(c.Ipv6DNSServer, obj1)
	}

	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
