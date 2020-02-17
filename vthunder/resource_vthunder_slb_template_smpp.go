package vthunder

//vThunder resource smpp

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"util"
)

func resourceSlbTemplateSmpp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSmppCreate,
		Update: resourceSmppUpdate,
		Read:   resourceSmppRead,
		Delete: resourceSmppDelete,

		Schema: map[string]*schema.Schema{
			"client_enquire_link": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_enquire_link": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_selection_per_request": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_enquire_link_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSmppCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating smpp (Inside resourceSmppCreate    " + name)
		v := dataToSmpp(name, d)
		d.SetId(name)
		go_vthunder.PostSmpp(client.Token, v, client.Host)

		return resourceSmppRead(d, meta)
	}
	return nil
}

func resourceSmppRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading smpp (Inside resourceSmppRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching smppp Read" + name)

		smpp, err := go_vthunder.GetSmpp(client.Token, name, client.Host)

		if smpp == nil {
			logger.Println("[INFO] No smpp found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSmppUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying smpp (Inside resourceSmppUpdate    " + name)
		v := dataToSmpp(name, d)
		d.SetId(name)
		go_vthunder.PutSmpp(client.Token, name, v, client.Host)

		return resourceSmppRead(d, meta)
	}
	return nil
}

func resourceSmppDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting smpp (Inside resourceSmppDelete) " + name)

		err := go_vthunder.DeleteSmpp(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete smpp  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate smpp structure
func dataToSmpp(name string, d *schema.ResourceData) go_vthunder.Smpp {
	var s go_vthunder.Smpp

	var sInstance go_vthunder.SmppInstance
	sInstance.Name = d.Get("name").(string)
	sInstance.ServerEnquireLink = d.Get("server_enquire_link").(int)
	sInstance.ServerSelectionPerRequest = d.Get("server_selection_per_request").(int)
	sInstance.ClientEnquireLink = d.Get("client_enquire_link").(int)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.ServerEnquireLinkVal = d.Get("server_enquire_link_val").(int)
	sInstance.User = d.Get("user").(string)
	sInstance.Password = d.Get("password").(string)
	sInstance.UUID = d.Get("uuid").(string)

	s.Name = sInstance

	return s
}
