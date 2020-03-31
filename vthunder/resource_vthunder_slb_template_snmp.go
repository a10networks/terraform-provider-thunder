package vthunder

//vThunder resource SlbTemplateSNMP

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateSNMP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateSNMPCreate,
		Update: resourceSlbTemplateSNMPUpdate,
		Read:   resourceSlbTemplateSNMPRead,
		Delete: resourceSlbTemplateSNMPDelete,
		Schema: map[string]*schema.Schema{
			"security_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"context_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"snmp_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"oid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"context_engine_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"community": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interface": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"security_engine_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priv_proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auth_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auth_proto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateSNMPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateSNMP (Inside resourceSlbTemplateSNMPCreate) ")
		snmp_name := d.Get("snmp_name").(string)
		data := dataToSlbTemplateSNMP(d)
		logger.Println("[INFO] received V from method data to SlbTemplateSNMP --")
		d.SetId(snmp_name)
		go_vthunder.PostSlbTemplateSNMP(client.Token, data, client.Host)

		return resourceSlbTemplateSNMPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateSNMPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateSNMP (Inside resourceSlbTemplateSNMPRead)")

	if client.Host != "" {
		snmp_name := d.Id()
		logger.Println("[INFO] Fetching service Read" + snmp_name)
		data, err := go_vthunder.GetSlbTemplateSNMP(client.Token, snmp_name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + snmp_name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateSNMPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateSNMP   (Inside resourceSlbTemplateSNMPUpdate) ")
		snmp_name := d.Get("snmp_name").(string)
		data := dataToSlbTemplateSNMP(d)
		logger.Println("[INFO] received V from method data to SlbTemplateSNMP ")
		d.SetId(snmp_name)
		go_vthunder.PutSlbTemplateSNMP(client.Token, snmp_name, data, client.Host)

		return resourceSlbTemplateSNMPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateSNMPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		snmp_name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateSNMPDelete) " + snmp_name)
		err := go_vthunder.DeleteSlbTemplateSNMP(client.Token, snmp_name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", snmp_name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateSNMP(d *schema.ResourceData) go_vthunder.SNMP {
	var vc go_vthunder.SNMP
	var c go_vthunder.SNMPInstance

	c.Username = d.Get("username").(string)
	c.Oid = d.Get("oid").(string)
	c.PrivProto = d.Get("priv_proto").(string)
	c.ContextName = d.Get("context_name").(string)
	c.AuthKey = d.Get("auth_key").(string)
	c.Interval = d.Get("interval").(int)
	c.ContextEngineID = d.Get("context_engine_id").(string)
	c.SecurityLevel = d.Get("security_level").(string)
	c.Community = d.Get("community").(string)
	c.AuthProto = d.Get("auth_proto").(string)
	c.Host = d.Get("host").(string)
	c.Version = d.Get("version").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.Interface = d.Get("interface").(int)
	c.PrivKey = d.Get("priv_key").(string)
	c.SecurityEngineID = d.Get("security_engine_id").(string)
	c.Port = d.Get("port").(int)
	c.SnmpName = d.Get("snmp_name").(string)
	vc.UUID = c
	return vc
}
