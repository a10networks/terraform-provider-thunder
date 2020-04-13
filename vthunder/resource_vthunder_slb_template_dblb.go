package vthunder

//vThunder resource TemplateDBLB

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateDBLB() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateDBLBCreate,
		Update: resourceTemplateDBLBUpdate,
		Read:   resourceTemplateDBLBRead,
		Delete: resourceTemplateDBLBDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"calc_sha1": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sha1_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"class_list": {
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

func resourceTemplateDBLBCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateDBLB (Inside resourceTemplateDBLBCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateDBLB(d)
		logger.Println("[INFO] received formatted data from method data to TemplateDBLB --")
		d.SetId(name)
		go_vthunder.PostTemplateDBLB(client.Token, data, client.Host)

		return resourceTemplateDBLBRead(d, meta)

	}
	return nil
}

func resourceTemplateDBLBRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateDBLB (Inside resourceTemplateDBLBRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateDBLB(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateDBLBUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateDBLB   (Inside resourceTemplateDBLBUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateDBLB(d)
		logger.Println("[INFO] received formatted data from method data to TemplateDBLB ")
		d.SetId(name)
		go_vthunder.PutTemplateDBLB(client.Token, name, data, client.Host)

		return resourceTemplateDBLBRead(d, meta)

	}
	return nil
}

func resourceTemplateDBLBDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateDBLBDelete) " + name)
		err := go_vthunder.DeleteTemplateDBLB(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateDBLB(d *schema.ResourceData) go_vthunder.DBLB {
	var vc go_vthunder.DBLB
	var c go_vthunder.DblbInstance
	var CalcSha1Obj go_vthunder.CalcSha1
	c.Name = d.Get("name").(string)
	c.ServerVersion = d.Get("server_version").(string)
	CalcSha1Obj.Sha1Value = d.Get("calc_sha1.0.sha1_value").(string)
	c.Sha1Value = CalcSha1Obj
	c.UserTag = d.Get("user_tag").(string)
	c.ClassList = d.Get("class_list").(string)

	vc.UUID = c
	return vc
}
