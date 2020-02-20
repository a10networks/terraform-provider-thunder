package vthunder

//vThunder resource TemplateCipher

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateCipher() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateCipherCreate,
		Update: resourceTemplateCipherUpdate,
		Read:   resourceTemplateCipherRead,
		Delete: resourceTemplateCipherDelete,
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
			"cipher_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher_suite": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"priority": {
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

func resourceTemplateCipherCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateCipher (Inside resourceTemplateCipherCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateCipher(d)
		logger.Println("[INFO] received V from method data to TemplateCipher --")
		d.SetId(name)
		go_vthunder.PostTemplateCipher(client.Token, data, client.Host)

		return resourceTemplateCipherRead(d, meta)

	}
	return nil
}

func resourceTemplateCipherRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateCipher (Inside resourceTemplateCipherRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateCipher(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateCipherUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateCipher   (Inside resourceTemplateCipherUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateCipher(d)
		logger.Println("[INFO] received V from method data to TemplateCipher ")
		d.SetId(name)
		go_vthunder.PutTemplateCipher(client.Token, name, data, client.Host)

		return resourceTemplateCipherRead(d, meta)

	}
	return nil
}

func resourceTemplateCipherDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateCipherDelete) " + name)
		err := go_vthunder.DeleteTemplateCipher(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateCipher(d *schema.ResourceData) go_vthunder.Cipher {
	var vc go_vthunder.Cipher
	var c go_vthunder.CipherInstance

	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)

	ciphercfgCounter := d.Get("cipher_cfg.#").(int)
	c.Priority = make([]go_vthunder.CipherCfg, 0, ciphercfgCounter)
	for i := 0; i < ciphercfgCounter; i++ {
		var cf go_vthunder.CipherCfg
		prefix := fmt.Sprintf("cipher_cfg.%d", i)
		cf.Priority = d.Get(prefix + ".priority").(int)
		cf.CipherSuite = d.Get(prefix + ".cipher_suite").(string)
		c.Priority = append(c.Priority, cf)

	}

	vc.UUID = c
	return vc
}
