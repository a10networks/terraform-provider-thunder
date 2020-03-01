package vthunder

//vThunder resource SlbTemplateServer

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateServerCreate,
		Update: resourceSlbTemplateServerUpdate,
		Read:   resourceSlbTemplateServerRead,
		Delete: resourceSlbTemplateServerDelete,
		Schema: map[string]*schema.Schema{
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"till": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_dynamic_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_query_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"slow_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"initial_slow_start": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_ttl_ratio": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"times": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
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
			"log_selection_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"every": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"add": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic_server_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"weight": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"spoofing_cache": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_no_logging": {
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
			"bw_rate_limit_acct": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateServerCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateServer (Inside resourceSlbTemplateServerCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServer(d)
		logger.Println("[INFO] received V from method data to SlbTemplateServer --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateServer(client.Token, data, client.Host)

		return resourceSlbTemplateServerRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateServerRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateServer (Inside resourceSlbTemplateServerRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateServer(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateServerUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateServer   (Inside resourceSlbTemplateServerUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServer(d)
		logger.Println("[INFO] received V from method data to SlbTemplateServer ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateServer(client.Token, name, data, client.Host)

		return resourceSlbTemplateServerRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateServerDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateServerDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateServer(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateServer(d *schema.ResourceData) go_vthunder.TemplateServer {
	var vc go_vthunder.TemplateServer
	var c go_vthunder.TemplateServerInstance

	c.Add = d.Get("add").(int)
	c.BwRateLimit = d.Get("bw_rate_limit").(int)
	c.BwRateLimitAcct = d.Get("bw_rate_limit_acct").(string)
	c.BwRateLimitDuration = d.Get("bw_rate_limit_duration").(int)
	c.BwRateLimitNoLogging = d.Get("bw_rate_limit_no_logging").(int)
	c.BwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	c.ConnLimit = d.Get("conn_limit").(int)
	c.ConnLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.ConnRateLimit = d.Get("conn_rate_limit").(int)
	c.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.DNSQueryInterval = d.Get("dns_query_interval").(int)
	c.DynamicServerPrefix = d.Get("dynamic_server_prefix").(string)
	c.Every = d.Get("every").(int)
	c.ExtendedStats = d.Get("extended_stats").(int)
	c.HealthCheck = d.Get("health_check").(string)
	c.HealthCheckDisable = d.Get("health_check_disable").(int)
	c.InitialSlowStart = d.Get("initial_slow_start").(int)
	c.LogSelectionFailure = d.Get("log_selection_failure").(int)
	c.MaxDynamicServer = d.Get("max_dynamic_server").(int)
	c.MinTTLRatio = d.Get("min_ttl_ratio").(int)
	c.Name = d.Get("name").(string)
	c.RateInterval = d.Get("rate_interval").(string)
	c.Resume = d.Get("resume").(int)
	c.SlowStart = d.Get("slow_start").(int)
	c.SpoofingCache = d.Get("spoofing_cache").(int)
	c.StatsDataAction = d.Get("stats_data_action").(string)
	c.Till = d.Get("till").(int)
	c.Times = d.Get("times").(int)
	c.Weight = d.Get("weight").(int)

	vc.UUID = c
	return vc
}
