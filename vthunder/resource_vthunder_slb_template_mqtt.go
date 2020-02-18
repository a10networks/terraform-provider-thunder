package vthunder

//vThunder resource mqtt

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"util"
)

func resourceSlbTemplateMqtt() *schema.Resource {
	return &schema.Resource{
		Create: resourceMqttCreate,
		Update: resourceMqttUpdate,
		Read:   resourceMqttRead,
		Delete: resourceMqttDelete,

		Schema: map[string]*schema.Schema{
			"clientid_hash_persist": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"clientid_hash_offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"clientid_hash_first": {
				Type:        schema.TypeInt,
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
			"clientid_hash_last": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceMqttCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating mqtt (Inside resourceMqttCreate    " + name)
		v := dataToMqtt(name, d)
		d.SetId(name)
		go_vthunder.PostMqtt(client.Token, v, client.Host)

		return resourceMqttRead(d, meta)
	}
	return nil
}

func resourceMqttRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading mqtt (Inside resourceMqttRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching mqtt Read" + name)

		mqtt, err := go_vthunder.GetMqtt(client.Token, name, client.Host)

		if mqtt == nil {
			logger.Println("[INFO] No mqtt found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceMqttUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying mqtt (Inside resourceMqttUpdate    " + name)
		v := dataToMqtt(name, d)
		d.SetId(name)
		go_vthunder.PutMqtt(client.Token, name, v, client.Host)

		return resourceMqttRead(d, meta)
	}
	return nil
}

func resourceMqttDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting mqtt (Inside resourceMqttDelete) " + name)

		err := go_vthunder.DeleteMqtt(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete mqtt  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate mqtt structure
func dataToMqtt(name string, d *schema.ResourceData) go_vthunder.Mqtt {
	var s go_vthunder.Mqtt

	var sInstance go_vthunder.MqttInstance

	sInstance.UUID = d.Get("uuid").(string)
	sInstance.ClientidHashLast = d.Get("clientid_hash_last").(int)
	sInstance.ClientidHashOffset = d.Get("clientid_hash_offset").(int)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.ClientidHashFirst = d.Get("clientid_hash_first").(int)
	sInstance.ClientidHashPersist = d.Get("clientid_hash_persist").(int)
	sInstance.Name = d.Get("name").(string)

	s.Name = sInstance

	return s
}
