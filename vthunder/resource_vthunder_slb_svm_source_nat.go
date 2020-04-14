package vthunder

//vThunder resource SlbSvmSourceNat

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSvmSourceNat() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSvmSourceNatCreate,
		Update: resourceSlbSvmSourceNatUpdate,
		Read:   resourceSlbSvmSourceNatRead,
		Delete: resourceSlbSvmSourceNatDelete,
		Schema: map[string]*schema.Schema{
			"pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbSvmSourceNatCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSvmSourceNat (Inside resourceSlbSvmSourceNatCreate) ")

		data := dataToSlbSvmSourceNat(d)
		logger.Println("[INFO] received formatted data from method data to SlbSvmSourceNat --")
		d.SetId("1")
		go_vthunder.PostSlbSvmSourceNat(client.Token, data, client.Host)

		return resourceSlbSvmSourceNatRead(d, meta)

	}
	return nil
}

func resourceSlbSvmSourceNatRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbSvmSourceNat (Inside resourceSlbSvmSourceNatRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbSvmSourceNat(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSvmSourceNatUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSvmSourceNatRead(d, meta)
}

func resourceSlbSvmSourceNatDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSvmSourceNatRead(d, meta)
}

func dataToSlbSvmSourceNat(d *schema.ResourceData) go_vthunder.SvmSourceNat {
	var vc go_vthunder.SvmSourceNat
	var c go_vthunder.SvmSourceNatInstance
	c.Pool = d.Get("pool").(string)

	vc.Pool = c
	return vc
}
