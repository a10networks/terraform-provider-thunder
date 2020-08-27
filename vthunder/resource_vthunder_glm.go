package vthunder

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceGlm() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlmCreate,
		Update: resourceGlmUpdate,
		Read:   resourceGlmRead,
		Delete: resourceGlmDelete,

		Schema: map[string]*schema.Schema{
			"use_mgmt_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"appliance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"enterprise": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"proxy_server": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"secret_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"encrypted": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"burst": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"send": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license_request": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"enable_requests": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"allocate_bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceGlmCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating Glm (Inside resourceGlmCreate)")

	if client.Host != "" {
		vc := dataToGlm(d)
		go_vthunder.PostGlm(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceGlmRead(d, meta)
	}
	return nil
}

func resourceGlmRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Glm (Inside resourceGlmRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		vc, err := go_vthunder.GetGlm(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No Glm found")
			//d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceGlmUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceGlmRead(d, meta)
}

func resourceGlmDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceGlmRead(d, meta)
}

//Utility method to instantiate Glm Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToGlm(d *schema.ResourceData) go_vthunder.Glm {
	var vc go_vthunder.Glm

	var c go_vthunder.GlmInstance

	c.Token = d.Get("token").(string)
	c.UseMgmtPort = d.Get("use_mgmt_port").(int)
	c.EnableRequests = d.Get("enable_requests").(int)
	c.AllocateBandwidth = d.Get("allocate_bandwidth").(int)

	if d.Get("send.#").(int) > 0 {
		//var send go_vthunder.Send
		send2 := &go_vthunder.Send{LicenseRequest: d.Get("send.0.license_request").(int)}
		//send.LicenseRequest = d.Get("send.0.license_request").(int)
		c.LicenseRequest = send2
	}

	vc.UUID = c

	return vc
}
