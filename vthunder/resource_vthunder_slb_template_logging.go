package vthunder

//vThunder resource logging

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"util"
)

func resourceSlbTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoggingCreate,
		Update: resourceLoggingUpdate,
		Read:   resourceLoggingRead,
		Delete: resourceLoggingDelete,

		Schema: map[string]*schema.Schema{
			"template_tcp_proxy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_tcp_proxy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pcre_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pool": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"keep_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"local_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pool_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tcp_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"keep_end": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceLoggingCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating Logging (Inside resourceLoggingCreate    " + name)
		v := dataToLogging(name, d)
		d.SetId(name)
		go_vthunder.PostLogging(client.Token, v, client.Host)

		return resourceLoggingRead(d, meta)
	}
	return nil
}

func resourceLoggingRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Logging (Inside resourcLoggingRead)")

	if client.Host != "" {
		client := meta.(vThunder)

		name := d.Id()

		logger.Println("[INFO] Fetching Logging Read" + name)

		logging, err := go_vthunder.GetLogging(client.Token, name, client.Host)

		if logging == nil {
			logger.Println("[INFO] No Logging found " + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceLoggingUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying Logging (Inside resourceLoggingUpdate    " + name)
		v := dataToLogging(name, d)
		d.SetId(name)
		go_vthunder.PutLogging(client.Token, name, v, client.Host)

		return resourceLoggingRead(d, meta)
	}
	return nil
}

func resourceLoggingDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting Logging (Inside resourceLoggingDelete) " + name)

		err := go_vthunder.DeleteLogging(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete Logging  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//utility method to instantiate Logging structure
func dataToLogging(name string, d *schema.ResourceData) go_vthunder.Logging {
	var s go_vthunder.Logging

	var sInstance go_vthunder.LoggingInstance

	sInstance.PoolShared = d.Get("pool_shared").(string)
	sInstance.Name = d.Get("name").(string)
	sInstance.Format = d.Get("format").(string)
	sInstance.Auto = d.Get("auto").(string)
	sInstance.KeepEnd = d.Get("keep_end").(int)
	sInstance.LocalLogging = d.Get("local_logging").(int)
	sInstance.Mask = d.Get("mask").(string)
	sInstance.TemplateTCPProxyShared = d.Get("template_tcp_proxy_shared").(string)
	sInstance.SharedPartitionTCPProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	sInstance.KeepStart = d.Get("keep_start").(int)
	sInstance.ServiceGroup = d.Get("service_group").(string)
	sInstance.PcreMask = d.Get("pcre_mask").(string)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.TCPProxy = d.Get("tcp_proxy").(string)
	sInstance.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	sInstance.Pool = d.Get("pool").(string)

	s.Name = sInstance

	return s
}
