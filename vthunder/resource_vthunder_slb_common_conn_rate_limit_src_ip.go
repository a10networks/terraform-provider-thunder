package vthunder

//vThunder resource SlbCommonConnRateLimitSrcIP

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbCommonConnRateLimitSrcIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbCommonConnRateLimitSrcIPCreate,
		Update: resourceSlbCommonConnRateLimitSrcIPUpdate,
		Read:   resourceSlbCommonConnRateLimitSrcIPRead,
		Delete: resourceSlbCommonConnRateLimitSrcIPDelete,
		Schema: map[string]*schema.Schema{
			"lock_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"exceed_action": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"limit_period": {
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

func resourceSlbCommonConnRateLimitSrcIPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbCommonConnRateLimitSrcIP (Inside resourceSlbCommonConnRateLimitSrcIPCreate) ")
		name := d.Get("protocol").(string)
		data := dataToSlbCommonConnRateLimitSrcIP(d)
		logger.Println("[INFO] received formatted data from method data to SlbCommonConnRateLimitSrcIP --")
		d.SetId(name)
		go_vthunder.PostSlbCommonConnRateLimitSrcIP(client.Token, data, client.Host)

		return resourceSlbCommonConnRateLimitSrcIPRead(d, meta)

	}
	return nil
}

func resourceSlbCommonConnRateLimitSrcIPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbCommonConnRateLimitSrcIP (Inside resourceSlbCommonConnRateLimitSrcIPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbCommonConnRateLimitSrcIP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbCommonConnRateLimitSrcIPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbCommonConnRateLimitSrcIP   (Inside resourceSlbCommonConnRateLimitSrcIPUpdate) ")
		name := d.Get("protocol").(string)
		data := dataToSlbCommonConnRateLimitSrcIP(d)
		logger.Println("[INFO] received formatted data from method data to SlbCommonConnRateLimitSrcIP ")
		d.SetId(name)
		go_vthunder.PutSlbCommonConnRateLimitSrcIP(client.Token, name, data, client.Host)

		return resourceSlbCommonConnRateLimitSrcIPRead(d, meta)

	}
	return nil
}

func resourceSlbCommonConnRateLimitSrcIPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbCommonConnRateLimitSrcIPDelete) " + name)
		err := go_vthunder.DeleteSlbCommonConnRateLimitSrcIP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbCommonConnRateLimitSrcIP(d *schema.ResourceData) go_vthunder.CommonConnRateLimitSrcIP {
	var vc go_vthunder.CommonConnRateLimitSrcIP
	var c go_vthunder.CommonConnRateLimitSrcIPInstance

	c.Protocol = d.Get("protocol").(string)
	c.Log = d.Get("log").(int)
	c.LockOut = d.Get("lock_out").(int)
	c.LimitPeriod = d.Get("limit_period").(string)
	c.Limit = d.Get("limit").(int)
	c.ExceedAction = d.Get("exceed_action").(int)
	c.Shared = d.Get("shared").(int)

	vc.UUID = c
	return vc
}
