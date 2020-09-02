package vthunder

//vThunder resource FwVrid

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwVrid() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwVridCreate,
		Update: resourceFwVridUpdate,
		Read:   resourceFwVridRead,
		Delete: resourceFwVridDelete,
		Schema: map[string]*schema.Schema{
			"vrid": {
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

func resourceFwVridCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwVrid (Inside resourceFwVridCreate) ")

		data := dataToFwVrid(d)
		logger.Println("[INFO] received formatted data from method data to FwVrid --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwVrid(client.Token, data, client.Host)

		return resourceFwVridRead(d, meta)

	}
	return nil
}

func resourceFwVridRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwVrid (Inside resourceFwVridRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwVrid(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwVridUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwVridRead(d, meta)
}

func resourceFwVridDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwVridRead(d, meta)
}
func dataToFwVrid(d *schema.ResourceData) go_vthunder.FwVrid {
	var vc go_vthunder.FwVrid
	var c go_vthunder.FwVridInstance
	c.Vrid = d.Get("vrid").(int)

	vc.Vrid = c
	return vc
}
