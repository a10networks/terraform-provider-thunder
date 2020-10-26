package thunder

//Thunder resource TemplateCipher

import (
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
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
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateCipher (Inside resourceTemplateCipherCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateCipher(d)
		logger.Println("[INFO] received formatted data from method data to TemplateCipher --")
		d.SetId(name)
		go_thunder.PostTemplateCipher(client.Token, data, client.Host)

		return resourceTemplateCipherRead(d, meta)

	}
	return nil
}

func resourceTemplateCipherRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading TemplateCipher (Inside resourceTemplateCipherRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetTemplateCipher(client.Token, name, client.Host)
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
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateCipher   (Inside resourceTemplateCipherUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateCipher(d)
		logger.Println("[INFO] received V from method data to TemplateCipher ")
		d.SetId(name)
		go_thunder.PutTemplateCipher(client.Token, name, data, client.Host)

		return resourceTemplateCipherRead(d, meta)

	}
	return nil
}

func resourceTemplateCipherDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateCipherDelete) " + name)
		err := go_thunder.DeleteTemplateCipher(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateCipher(d *schema.ResourceData) go_thunder.Cipher {
	var vc go_thunder.Cipher
	var c go_thunder.CipherInstance

	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)

	ciphercfgCounter := d.Get("cipher_cfg.#").(int)
	c.Priority = make([]go_thunder.CipherCfg, 0, ciphercfgCounter)
	for i := 0; i < ciphercfgCounter; i++ {
		var cf go_thunder.CipherCfg
		prefix := fmt.Sprintf("cipher_cfg.%d", i)
		cf.Priority = d.Get(prefix + ".priority").(int)
		cf.CipherSuite = d.Get(prefix + ".cipher_suite").(string)
		c.Priority = append(c.Priority, cf)

	}

	vc.UUID = c
	return vc
}
