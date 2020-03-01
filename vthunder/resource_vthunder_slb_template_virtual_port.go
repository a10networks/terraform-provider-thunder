package vthunder

//vThunder resource SlbTemplateVirtualPort

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateVirtualPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateVirtualPortCreate,
		Update: resourceSlbTemplateVirtualPortUpdate,
		Read:   resourceSlbTemplateVirtualPortRead,
		Delete: resourceSlbTemplateVirtualPortDelete,
		Schema: map[string]*schema.Schema{
			"allow_syn_otherflags": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_port_preserve": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"aflow": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dscp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_msl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allow_vip_to_rport_mapping": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reset_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_options": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reset_l7_on_failover": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ignore_tcp_msl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"drop_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"when_rr_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"non_syn_initiation": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateVirtualPortCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateVirtualPort (Inside resourceSlbTemplateVirtualPortCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateVirtualPort(d)
		logger.Println("[INFO] received V from method data to SlbTemplateVirtualPort --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateVirtualPort(client.Token, data, client.Host)

		return resourceSlbTemplateVirtualPortRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateVirtualPortRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateVirtualPort (Inside resourceSlbTemplateVirtualPortRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateVirtualPort(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateVirtualPortUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateVirtualPort   (Inside resourceSlbTemplateVirtualPortUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateVirtualPort(d)
		logger.Println("[INFO] received V from method data to SlbTemplateVirtualPort ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateVirtualPort(client.Token, name, data, client.Host)

		return resourceSlbTemplateVirtualPortRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateVirtualPortDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateVirtualPortDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateVirtualPort(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateVirtualPort(d *schema.ResourceData) go_vthunder.VirtualPort {
	var vc go_vthunder.VirtualPort
	var c go_vthunder.VirtualPortInstance

	c.ResetUnknownConn = d.Get("reset_unknown_conn").(int)
	c.IgnoreTCPMsl = d.Get("ignore_tcp_msl").(int)
	c.Rate = d.Get("rate").(int)
	c.SnatMsl = d.Get("snat_msl").(int)
	c.AllowSynOtherflags = d.Get("allow_syn_otherflags").(int)
	c.Aflow = d.Get("aflow").(int)
	c.ConnLimit = d.Get("conn_limit").(int)
	c.DropUnknownConn = d.Get("drop_unknown_conn").(int)
	c.ResetL7OnFailover = d.Get("reset_l7_on_failover").(int)
	c.PktRateType = d.Get("pkt_rate_type").(string)
	c.RateInterval = d.Get("rate_interval").(string)
	c.SnatPortPreserve = d.Get("snat_port_preserve").(int)
	c.ConnRateLimitReset = d.Get("conn_rate_limit_reset").(int)
	c.WhenRrEnable = d.Get("when_rr_enable").(int)
	c.NonSynInitiation = d.Get("non_syn_initiation").(int)
	c.ConnLimitReset = d.Get("conn_limit_reset").(int)
	c.Dscp = d.Get("dscp").(int)
	c.PktRateLimitReset = d.Get("pkt_rate_limit_reset").(int)
	c.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	c.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.LogOptions = d.Get("log_options").(string)
	c.Name = d.Get("name").(string)
	c.AllowVipToRportMapping = d.Get("allow_vip_to_rport_mapping").(int)
	c.PktRateInterval = d.Get("pkt_rate_interval").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.ConnRateLimit = d.Get("conn_rate_limit").(int)
	vc.UUID = c
	return vc
}
