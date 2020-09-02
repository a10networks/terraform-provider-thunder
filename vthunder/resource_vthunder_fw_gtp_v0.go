package vthunder

//vThunder resource FwGtpV0

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwGtpV0() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwGtpV0Create,
		Update: resourceFwGtpV0Update,
		Read:   resourceFwGtpV0Read,
		Delete: resourceFwGtpV0Delete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"gtpv0_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwGtpV0Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwGtpV0 (Inside resourceFwGtpV0Create) ")

		data := dataToFwGtpV0(d)
		logger.Println("[INFO] received formatted data from method data to FwGtpV0 --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwGtpV0(client.Token, data, client.Host)

		return resourceFwGtpV0Read(d, meta)

	}
	return nil
}

func resourceFwGtpV0Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwGtpV0 (Inside resourceFwGtpV0Read)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwGtpV0(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwGtpV0Update(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGtpV0Read(d, meta)
}

func resourceFwGtpV0Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwGtpV0Read(d, meta)
}
func dataToFwGtpV0(d *schema.ResourceData) go_vthunder.FwGtpV0 {
	var vc go_vthunder.FwGtpV0
	var c go_vthunder.FwGtpV0Instance
	c.Gtpv0Value = d.Get("gtpv0_value").(string)

	vc.UUID = c
	return vc
}
