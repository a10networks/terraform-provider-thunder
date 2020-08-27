package vthunder

//vThunder resource harmony controller profile

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceHarmonyControllerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceProfileCreate,
		Update: resourceProfileUpdate,
		Read:   resourceProfileRead,
		Delete: resourceProfileDelete,
		Schema: map[string]*schema.Schema{
			"thunder_mgmt_ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"use_mgmt_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"provider2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secret_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceProfileCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating Profile (Inside resourceProfileCreate)")

	if client.Host != "" {
		name := d.Get("user_name").(string)
		vc := dataToProfile(d)
		d.SetId(name)
		go_vthunder.PostProfile(client.Token, vc, client.Host)

		return resourceProfileRead(d, meta)
	}
	return nil
}

func resourceProfileRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Profile (Inside resourceProfileRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetProfile(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No Profile found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceProfileUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceProfileRead(d, meta)
}

func resourceProfileDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceProfileRead(d, meta)
}

//Utility method to instantiate Profile Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToProfile(d *schema.ResourceData) go_vthunder.Profile {
	var vc go_vthunder.ProfileInstance

	var c go_vthunder.Profile

	vc.Host = d.Get("host").(string)
	vc.Port = d.Get("port").(int)
	vc.UserName = d.Get("user_name").(string)
	vc.SecretValue = d.Get("secret_value").(string)
	vc.Provider2 = d.Get("provider2").(string)
	vc.Action = d.Get("action").(string)
	vc.UseMgmtPort = d.Get("use_mgmt_port").(int)
	vc.Region = d.Get("region").(string)
	vc.AvailabilityZone = d.Get("availability_zone").(string)

	var mgmtIp go_vthunder.ThunderMgmtIP

	mgmtIp.IPAddress = d.Get("thunder_mgmt_ip.0.ip_address").(string)

	vc.IPAddress = mgmtIp

	c.Host = vc

	return c
}
