package vthunder

//vThunder resource FwTcpRstCloseImmediate

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwTcpRstCloseImmediate() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwTcpRstCloseImmediateCreate,
		Update: resourceFwTcpRstCloseImmediateUpdate,
		Read:   resourceFwTcpRstCloseImmediateRead,
		Delete: resourceFwTcpRstCloseImmediateDelete,
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

func resourceFwTcpRstCloseImmediateCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTcpRstCloseImmediate (Inside resourceFwTcpRstCloseImmediateCreate) ")

		data := dataToFwTcpRstCloseImmediate(d)
		logger.Println("[INFO] received formatted data from method data to FwTcpRstCloseImmediate --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwTcpRstCloseImmediate(client.Token, data, client.Host)

		return resourceFwTcpRstCloseImmediateRead(d, meta)

	}
	return nil
}

func resourceFwTcpRstCloseImmediateRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwTcpRstCloseImmediate (Inside resourceFwTcpRstCloseImmediateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwTcpRstCloseImmediate(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwTcpRstCloseImmediateUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpRstCloseImmediateRead(d, meta)
}

func resourceFwTcpRstCloseImmediateDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTcpRstCloseImmediateRead(d, meta)
}
func dataToFwTcpRstCloseImmediate(d *schema.ResourceData) go_vthunder.FwTcpRstCloseImmediate {
	var vc go_vthunder.FwTcpRstCloseImmediate
	var c go_vthunder.FwTCPRstCloseImmediateInstance
	c.Status = d.Get("status").(string)

	vc.Status = c
	return vc
}
