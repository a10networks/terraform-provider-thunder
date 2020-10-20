package thunder

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
	client := meta.(Thunder)

	logger.Println("[INFO] Creating Glm (Inside resourceGlmCreate)")

	if client.Host != "" {
		vc := dataToGlm(d)
		go_thunder.PostGlm(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceGlmRead(d, meta)
	}
	return nil
}

func resourceGlmRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Glm (Inside resourceGlmRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		vc, err := go_thunder.GetGlm(client.Token, client.Host)

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

func dataToGlm(d *schema.ResourceData) go_thunder.Glm {
	var vc go_thunder.Glm
	var c go_thunder.GlmInstance
	c.UseMgmtPort = d.Get("use_mgmt_port").(int)
	c.Burst = d.Get("burst").(int)
	c.Interval = d.Get("interval").(int)

	c.Token = d.Get("token").(string)
	c.Enterprise = d.Get("enterprise").(string)

	var obj2 go_thunder.ProxyServer
	prefix := "proxy_server.0."
	obj2.Username = d.Get(prefix + "username").(string)
	obj2.Encrypted = d.Get(prefix + "encrypted").(string)
	obj2.Host = d.Get(prefix + "host").(string)
	obj2.Password = d.Get(prefix + "password").(int)
	obj2.Port = d.Get(prefix + "port").(int)
	obj2.SecretString = d.Get(prefix + "secret_string").(string)
	c.Username = obj2

	c.ApplianceName = d.Get("appliance_name").(string)
	c.EnableRequests = d.Get("enable_requests").(int)
	c.AllocateBandwidth = d.Get("allocate_bandwidth").(int)
	c.Port = d.Get("port").(int)
	vc.UUID = c
	return vc
}
