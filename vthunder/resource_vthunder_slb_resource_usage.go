package vthunder

//vThunder resource SlbResourceUsage

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbResourceUsage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbResourceUsageCreate,
		Update: resourceSlbResourceUsageUpdate,
		Read:   resourceSlbResourceUsageRead,
		Delete: resourceSlbResourceUsageDelete,
		Schema: map[string]*schema.Schema{
			"persist_srcip_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"virtual_server_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pbslb_subnet_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"http_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"real_port_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"nat_pool_addr_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fast_udp_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cache_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_threshold_res_usage_percent": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fast_tcp_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"virtual_port_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_ssl_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_reuse_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"proxy_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stream_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"real_server_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_monitor_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persist_cookie_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client_ssl_template_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbResourceUsageCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbResourceUsage (Inside resourceSlbResourceUsageCreate) ")

		data := dataToSlbResourceUsage(d)
		logger.Println("[INFO] received formatted data from method data to SlbResourceUsage --")
		d.SetId("1")
		go_vthunder.PostSlbResourceUsage(client.Token, data, client.Host)

		return resourceSlbResourceUsageRead(d, meta)

	}
	return nil
}

func resourceSlbResourceUsageRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbResourceUsage (Inside resourceSlbResourceUsageRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbResourceUsage(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbResourceUsageUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbResourceUsageRead(d, meta)
}

func resourceSlbResourceUsageDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbResourceUsageRead(d, meta)
}

func dataToSlbResourceUsage(d *schema.ResourceData) go_vthunder.ResourceUsage {
	var vc go_vthunder.ResourceUsage
	var c go_vthunder.ResourceUsageInstance
	c.PbslbSubnetCount = d.Get("pbslb_subnet_count").(int)
	c.ClientSslTemplateCount = d.Get("client_ssl_template_count").(int)
	c.StreamTemplateCount = d.Get("stream_template_count").(int)
	c.NatPoolAddrCount = d.Get("nat_pool_addr_count").(int)
	c.VirtualPortCount = d.Get("virtual_port_count").(int)
	c.ServerSslTemplateCount = d.Get("server_ssl_template_count").(int)
	c.RealServerCount = d.Get("real_server_count").(int)
	c.SlbThresholdResUsagePercent = d.Get("slb_threshold_res_usage_percent").(int)
	c.ServiceGroupCount = d.Get("service_group_count").(int)
	c.VirtualServerCount = d.Get("virtual_server_count").(int)
	c.CacheTemplateCount = d.Get("cache_template_count").(int)
	c.FastTCPTemplateCount = d.Get("fast_tcp_template_count").(int)
	c.ConnReuseTemplateCount = d.Get("conn_reuse_template_count").(int)
	c.HealthMonitorCount = d.Get("health_monitor_count").(int)
	c.PersistCookieTemplateCount = d.Get("persist_cookie_template_count").(int)
	c.PersistSrcipTemplateCount = d.Get("persist_srcip_template_count").(int)
	c.ProxyTemplateCount = d.Get("proxy_template_count").(int)
	c.HTTPTemplateCount = d.Get("http_template_count").(int)
	c.RealPortCount = d.Get("real_port_count").(int)
	c.FastUDPTemplateCount = d.Get("fast_udp_template_count").(int)

	vc.SlbThresholdResUsagePercent = c
	return vc
}
