package thunder

//Thunder resource NtpTrustedKey

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceNtpTrustedKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceNtpTrustedKeyCreate,
		Update: resourceNtpTrustedKeyUpdate,
		Read:   resourceNtpTrustedKeyRead,
		Delete: resourceNtpTrustedKeyDelete,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeInt,
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

func resourceNtpTrustedKeyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating NtpTrustedKey (Inside resourceNtpTrustedKeyCreate) ")
		name1 := d.Get("key").(int)
		data := dataToNtpTrustedKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpTrustedKey --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostNtpTrustedKey(client.Token, data, client.Host)

		return resourceNtpTrustedKeyRead(d, meta)

	}
	return nil
}

func resourceNtpTrustedKeyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading NtpTrustedKey (Inside resourceNtpTrustedKeyRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetNtpTrustedKey(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceNtpTrustedKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying NtpTrustedKey   (Inside resourceNtpTrustedKeyUpdate) ")
		data := dataToNtpTrustedKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpTrustedKey ")
		go_thunder.PutNtpTrustedKey(client.Token, name1, data, client.Host)

		return resourceNtpTrustedKeyRead(d, meta)

	}
	return nil
}

func resourceNtpTrustedKeyDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceNtpTrustedKeyDelete) " + name1)
		err := go_thunder.DeleteNtpTrustedKey(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToNtpTrustedKey(d *schema.ResourceData) go_thunder.NtpTrustedKey {
	var vc go_thunder.NtpTrustedKey
	var c go_thunder.NtpTrustedKeyInstance
	c.Key = d.Get("key").(int)

	vc.Key = c
	return vc
}
