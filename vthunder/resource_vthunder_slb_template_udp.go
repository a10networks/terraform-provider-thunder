package vthunder

//vThunder resource Udp

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"util"
)

func resourceSlbTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Create: resourceUdpCreate,
		Update: resourceUdpUpdate,
		Read:   resourceUdpRead,
		Delete: resourceUdpDelete,

		Schema: map[string]*schema.Schema{
			"qos": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"re_select_if_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"immediate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"idle_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_conn_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceUdpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating udp (Inside resourceUdpCreate    " + name)
		v := dataToUdp(name, d)
		d.SetId(name)
		go_vthunder.PostUdp(client.Token, v, client.Host)

		return resourceUdpRead(d, meta)
	}
	return nil
}

func resourceUdpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading udp (Inside resourceUdpRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching udp Read" + name)

		udp, err := go_vthunder.GetUdp(client.Token, name, client.Host)

		if udp == nil {
			logger.Println("[INFO] No udp found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceUdpUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying udp (Inside resourceUdpUpdate    " + name)
		v := dataToUdp(name, d)
		d.SetId(name)
		go_vthunder.PutUdp(client.Token, name, v, client.Host)

		return resourceUdpRead(d, meta)
	}
	return nil
}

func resourceUdpDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting udp (Inside resourceUdpDelete) " + name)

		err := go_vthunder.DeleteUdp(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete udp  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate udp structure
func dataToUdp(name string, d *schema.ResourceData) go_vthunder.UDP {
	var s go_vthunder.UDP

	var sInstance go_vthunder.UDPInstance

	sInstance.Qos = d.Get("qos").(int)
	sInstance.Name = d.Get("name").(string)
	sInstance.StatelessConnTimeout = d.Get("stateless_conn_timeout").(int)
	sInstance.IdleTimeout = d.Get("idle_timeout").(int)
	sInstance.ReSelectIfServerDown = d.Get("re_select_if_server_down").(int)
	sInstance.Immediate = d.Get("immediate").(int)
	sInstance.UserTag = d.Get("user_tag").(string)

	s.Name = sInstance

	return s
}
