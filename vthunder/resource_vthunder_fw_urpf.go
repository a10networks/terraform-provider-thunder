package vthunder

//vThunder resource FwUrpf

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwUrpf() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwUrpfCreate,
		Update: resourceFwUrpfUpdate,
		Read:   resourceFwUrpfRead,
		Delete: resourceFwUrpfDelete,
		Schema: map[string]*schema.Schema{
			"status": {
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

func resourceFwUrpfCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwUrpf (Inside resourceFwUrpfCreate) ")

		data := dataToFwUrpf(d)
		logger.Println("[INFO] received formatted data from method data to FwUrpf --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwUrpf(client.Token, data, client.Host)

		return resourceFwUrpfRead(d, meta)

	}
	return nil
}

func resourceFwUrpfRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwUrpf (Inside resourceFwUrpfRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwUrpf(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwUrpfUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwUrpfRead(d, meta)
}

func resourceFwUrpfDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwUrpfRead(d, meta)
}
func dataToFwUrpf(d *schema.ResourceData) go_vthunder.FwUrpf {
	var vc go_vthunder.FwUrpf
	var c go_vthunder.FwUrpfInstance
	c.Status = d.Get("status").(string)

	vc.Status = c
	return vc
}
