package vthunder

//vThunder resource FwHelperSessions

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwHelperSessions() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwHelperSessionsCreate,
		Update: resourceFwHelperSessionsUpdate,
		Read:   resourceFwHelperSessionsRead,
		Delete: resourceFwHelperSessionsDelete,
		Schema: map[string]*schema.Schema{
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mode": {
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

func resourceFwHelperSessionsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwHelperSessions (Inside resourceFwHelperSessionsCreate) ")

		data := dataToFwHelperSessions(d)
		logger.Println("[INFO] received formatted data from method data to FwHelperSessions --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwHelperSessions(client.Token, data, client.Host)

		return resourceFwHelperSessionsRead(d, meta)

	}
	return nil
}

func resourceFwHelperSessionsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwHelperSessions (Inside resourceFwHelperSessionsRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwHelperSessions(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwHelperSessionsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwHelperSessionsRead(d, meta)
}

func resourceFwHelperSessionsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwHelperSessionsRead(d, meta)
}
func dataToFwHelperSessions(d *schema.ResourceData) go_vthunder.FwHelperSessions {
	var vc go_vthunder.FwHelperSessions
	var c go_vthunder.FwHelperSessionsInstance
	c.IdleTimeout = d.Get("idle_timeout").(int)
	c.Limit = d.Get("limit").(int)
	c.Mode = d.Get("mode").(string)

	vc.IdleTimeout = c
	return vc
}
