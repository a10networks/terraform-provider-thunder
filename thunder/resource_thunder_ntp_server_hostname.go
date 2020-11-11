package thunder

//Thunder resource NtpServerHostname

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceNtpServerHostname() *schema.Resource {
	return &schema.Resource{
		Create: resourceNtpServerHostnameCreate,
		Update: resourceNtpServerHostnameUpdate,
		Read:   resourceNtpServerHostnameRead,
		Delete: resourceNtpServerHostnameDelete,
		Schema: map[string]*schema.Schema{
			"host_servername": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"key": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"prefer": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
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

func resourceNtpServerHostnameCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating NtpServerHostname (Inside resourceNtpServerHostnameCreate) ")
		name1 := d.Get("host_servername").(string)
		data := dataToNtpServerHostname(d)
		logger.Println("[INFO] received formatted data from method data to NtpServerHostname --")
		d.SetId(name1)
		go_thunder.PostNtpServerHostname(client.Token, data, client.Host)

		return resourceNtpServerHostnameRead(d, meta)

	}
	return nil
}

func resourceNtpServerHostnameRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading NtpServerHostname (Inside resourceNtpServerHostnameRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetNtpServerHostname(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceNtpServerHostnameUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying NtpServerHostname   (Inside resourceNtpServerHostnameUpdate) ")
		data := dataToNtpServerHostname(d)
		logger.Println("[INFO] received formatted data from method data to NtpServerHostname ")
		go_thunder.PutNtpServerHostname(client.Token, name1, data, client.Host)

		return resourceNtpServerHostnameRead(d, meta)

	}
	return nil
}

func resourceNtpServerHostnameDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceNtpServerHostnameDelete) " + name1)
		err := go_thunder.DeleteNtpServerHostname(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return err
		}
		return nil
	}
	return nil
}

func dataToNtpServerHostname(d *schema.ResourceData) go_thunder.NtpServerHostname {
	var vc go_thunder.NtpServerHostname
	var c go_thunder.NtpServerHostnameInstance
	c.HostServername = d.Get("host_servername").(string)
	c.Key = d.Get("key").(int)
	c.Prefer = d.Get("prefer").(int)
	c.Action = d.Get("action").(string)

	vc.HostServername = c
	return vc
}
