package thunder

//Thunder resource NtpAuthKey

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceNtpAuthKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceNtpAuthKeyCreate,
		Update: resourceNtpAuthKeyUpdate,
		Read:   resourceNtpAuthKeyRead,
		Delete: resourceNtpAuthKeyDelete,
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"alg_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"asc_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"hex_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"hex_encrypted": {
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

func resourceNtpAuthKeyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating NtpAuthKey (Inside resourceNtpAuthKeyCreate) ")
		name1 := d.Get("key").(int)
		data := dataToNtpAuthKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpAuthKey --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostNtpAuthKey(client.Token, data, client.Host)

		return resourceNtpAuthKeyRead(d, meta)

	}
	return nil
}

func resourceNtpAuthKeyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading NtpAuthKey (Inside resourceNtpAuthKeyRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetNtpAuthKey(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceNtpAuthKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying NtpAuthKey   (Inside resourceNtpAuthKeyUpdate) ")
		data := dataToNtpAuthKey(d)
		logger.Println("[INFO] received formatted data from method data to NtpAuthKey ")
		go_thunder.PutNtpAuthKey(client.Token, name1, data, client.Host)

		return resourceNtpAuthKeyRead(d, meta)

	}
	return nil
}

func resourceNtpAuthKeyDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceNtpAuthKeyDelete) " + name1)
		err := go_thunder.DeleteNtpAuthKey(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToNtpAuthKey(d *schema.ResourceData) go_thunder.NtpAuthKey {
	var vc go_thunder.NtpAuthKey
	var c go_thunder.NtpAuthKeyInstance
	c.Key = d.Get("key").(int)
	c.AlgType = d.Get("alg_type").(string)
	c.KeyType = d.Get("key_type").(string)
	c.AscKey = d.Get("asc_key").(string)
	c.Encrypted = d.Get("encrypted").(string)
	c.HexKey = d.Get("hex_key").(string)
	c.HexEncrypted = d.Get("hex_encrypted").(string)

	vc.Key = c
	return vc
}
