package thunder

//Thunder resource SlbTemplateRespmodIcap

import (
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateRespmodIcap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateRespmodIcapCreate,
		Update: resourceSlbTemplateRespmodIcapUpdate,
		Read:   resourceSlbTemplateRespmodIcapRead,
		Delete: resourceSlbTemplateRespmodIcapDelete,
		Schema: map[string]*schema.Schema{
			"preview": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"include_protocol_in_uri": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_tcp_proxy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_persist_source_ip_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_tcp_proxy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_http_server_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cylance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"min_payload_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bypass_ip_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"fail_close": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tcp_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"logging": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"source_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_ssl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_source_ip_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"log_only_allowed_method": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateRespmodIcapCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateRespmodIcap (Inside resourceSlbTemplateRespmodIcapCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateRespmodIcap(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateRespmodIcap --")
		d.SetId(name)
		go_thunder.PostSlbTemplateRespmodIcap(client.Token, data, client.Host)

		return resourceSlbTemplateRespmodIcapRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateRespmodIcapRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateRespmodIcap (Inside resourceSlbTemplateRespmodIcapRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateRespmodIcap(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateRespmodIcapUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateRespmodIcap   (Inside resourceSlbTemplateRespmodIcapUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateRespmodIcap(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateRespmodIcap ")
		d.SetId(name)
		go_thunder.PutSlbTemplateRespmodIcap(client.Token, name, data, client.Host)

		return resourceSlbTemplateRespmodIcapRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateRespmodIcapDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateRespmodIcapDelete) " + name)
		err := go_thunder.DeleteSlbTemplateRespmodIcap(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateRespmodIcap(d *schema.ResourceData) go_thunder.RespmodIcap {
	var vc go_thunder.RespmodIcap
	var c go_thunder.RespmodIcapInstance

	c.MinPayloadSize = d.Get("min_payload_size").(int)
	c.SharedPartitionPersistSourceIPTemplate = d.Get("shared_partition_persist_source_ip_template").(int)
	c.TemplateTCPProxyShared = d.Get("template_tcp_proxy_shared").(string)
	c.ServiceURL = d.Get("service_url").(string)
	c.SharedPartitionTCPProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	c.ServiceGroup = d.Get("service_group").(string)
	c.TCPProxy = d.Get("tcp_proxy").(string)
	c.Preview = d.Get("preview").(int)
	c.DisableHTTPServerReset = d.Get("disable_http_server_reset").(int)
	c.ServerSsl = d.Get("server_ssl").(string)
	c.FailClose = d.Get("fail_close").(int)

	BypassIpCfgCount := d.Get("bypass_ip_cfg.#").(int)
	c.Mask = make([]go_thunder.BypassIPCfg2, 0, BypassIpCfgCount)

	for i := 0; i < BypassIpCfgCount; i++ {
		var obj1 go_thunder.BypassIPCfg2
		prefix := fmt.Sprintf("bypass_ip_cfg.%d.", i)
		obj1.Mask = d.Get(prefix + "mask").(string)
		obj1.BypassIP = d.Get(prefix + "bypass_ip").(string)
		c.Mask = append(c.Mask, obj1)
	}

	c.TemplatePersistSourceIPShared = d.Get("template_persist_source_ip_shared").(string)
	c.IncludeProtocolInURI = d.Get("include_protocol_in_uri").(int)
	c.Logging = d.Get("logging").(string)
	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.LogOnlyAllowedMethod = d.Get("log_only_allowed_method").(int)
	c.Action = d.Get("action").(string)
	c.Cylance = d.Get("cylance").(int)
	c.SourceIP = d.Get("source_ip").(string)

	vc.UUID = c
	return vc
}
