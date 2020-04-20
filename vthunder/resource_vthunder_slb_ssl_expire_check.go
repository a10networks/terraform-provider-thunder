package vthunder

//vThunder resource SlbSSLExpireCheck

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSSLExpireCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSSLExpireCheckCreate,
		Update: resourceSlbSSLExpireCheckUpdate,
		Read:   resourceSlbSSLExpireCheckRead,
		Delete: resourceSlbSSLExpireCheckDelete,
		Schema: map[string]*schema.Schema{
			"exception": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"certificate_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"expire_address1": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"before": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"interval_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_expire_email_address": {
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

func resourceSlbSSLExpireCheckCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSSLExpireCheck (Inside resourceSlbSSLExpireCheckCreate) ")

		data := dataToSlbSSLExpireCheck(d)
		logger.Println("[INFO] received SSL Expire Check from method data to SlbSSLExpireCheck --")
		d.SetId("1")
		go_vthunder.PostSlbSSLExpireCheck(client.Token, data, client.Host)

		return resourceSlbSSLExpireCheckRead(d, meta)

	}
	return nil
}

func resourceSlbSSLExpireCheckRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbSSLExpireCheck (Inside resourceSlbSSLExpireCheckRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetSlbSSLExpireCheck(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSSLExpireCheckUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSSLExpireCheckRead(d, meta)
}

func resourceSlbSSLExpireCheckDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSSLExpireCheckRead(d, meta)
}
func dataToSlbSSLExpireCheck(d *schema.ResourceData) go_vthunder.SSLExpireCheck {
	var vc go_vthunder.SSLExpireCheck
	var c go_vthunder.SSLExpireCheckInstance
	c.SslExpireEmailAddress = d.Get("ssl_expire_email_address").(string)

	var obj1 go_vthunder.Exception
	prefix := "exception.0."
	obj1.Action = d.Get(prefix + "action").(string)
	obj1.CertificateName = d.Get(prefix + "certificate_name").(string)
	c.Action = obj1

	c.ExpireAddress1 = d.Get("expire_address1").(string)
	c.IntervalDays = d.Get("interval_days").(int)
	c.Before = d.Get("before").(int)
	vc.UUID = c
	return vc
}
