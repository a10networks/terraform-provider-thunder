package vthunder

//vThunder resource Tcp

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"util"
)

func resourceSlbTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Create: resourceTcpCreate,
		Update: resourceTcpUpdate,
		Read:   resourceTcpRead,
		Delete: resourceTcpDelete,

		Schema: map[string]*schema.Schema{
			"insert_client_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"lan_fast_ack": {
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
			"reset_rev": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_follow_fin": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_fwd": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"del_session_on_server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTcpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating tcp (Inside resourceTcpCreate    " + name)
		v := dataToTcp(name, d)
		d.SetId(name)
		go_vthunder.PostTcp(client.Token, v, client.Host)

		return resourceTcpRead(d, meta)
	}
	return nil
}

func resourceTcpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading tcp (Inside resourceTcpRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching tcp Read" + name)

		tcp, err := go_vthunder.GetTcp(client.Token, name, client.Host)

		if tcp == nil {
			logger.Println("[INFO] No tcp found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceTcpUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying tcp (Inside resourceTcpUpdate    " + name)
		v := dataToTcp(name, d)
		d.SetId(name)
		go_vthunder.PutTcp(client.Token, name, v, client.Host)

		return resourceTcpRead(d, meta)
	}
	return nil
}

func resourceTcpDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting tcp (Inside resourceTcpDelete) " + name)

		err := go_vthunder.DeleteTcp(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete tcp  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate tcp structure
func dataToTcp(name string, d *schema.ResourceData) go_vthunder.TCP {
	var s go_vthunder.TCP

	var sInstance go_vthunder.TCPInstance

	sInstance.IdleTimeout = d.Get("idle_timeout").(int)
	sInstance.InsertClientIP = d.Get("insert_client_ip").(int)
	sInstance.LanFastAck = d.Get("lan_fast_ack").(int)
	sInstance.ResetFwd = d.Get("reset_fwd").(int)
	sInstance.ResetRev = d.Get("reset_rev").(int)
	sInstance.ResetFollowFin = d.Get("reset_follow_fin").(int)
	sInstance.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)

	s.Name = sInstance

	return s
}
