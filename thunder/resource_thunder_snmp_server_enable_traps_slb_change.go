package thunder

//Thunder resource SnmpServerEnableTrapsSlbChange

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSlbChange() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsSlbChangeCreate,
		Update: resourceSnmpServerEnableTrapsSlbChangeUpdate,
		Read:   resourceSnmpServerEnableTrapsSlbChangeRead,
		Delete: resourceSnmpServerEnableTrapsSlbChangeDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resource_usage_warning": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ssl_cert_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ssl_cert_expire": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"system_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"connection_resource_event": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsSlbChangeCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSlbChange (Inside resourceSnmpServerEnableTrapsSlbChangeCreate) ")

		data := dataToSnmpServerEnableTrapsSlbChange(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSlbChange --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsSlbChange(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsSlbChangeRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsSlbChangeRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSlbChange (Inside resourceSnmpServerEnableTrapsSlbChangeRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSlbChange(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsSlbChangeUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSlbChangeRead(d, meta)
}

func resourceSnmpServerEnableTrapsSlbChangeDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSlbChangeRead(d, meta)
}
func dataToSnmpServerEnableTrapsSlbChange(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSlbChange {
	var vc go_thunder.SnmpServerEnableTrapsSlbChange
	var c go_thunder.SnmpServerEnableTrapsSlbChangeInstance
	c.All = d.Get("all").(int)
	c.ResourceUsageWarning = d.Get("resource_usage_warning").(int)
	c.SslCertChange = d.Get("ssl_cert_change").(int)
	c.SslCertExpire = d.Get("ssl_cert_expire").(int)
	c.SystemThreshold = d.Get("system_threshold").(int)
	c.Server = d.Get("server").(int)
	c.Vip = d.Get("vip").(int)
	c.ConnectionResourceEvent = d.Get("connection_resource_event").(int)
	c.ServerPort = d.Get("server_port").(int)
	c.VipPort = d.Get("vip_port").(int)

	vc.All = c
	return vc
}
