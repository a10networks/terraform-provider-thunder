package thunder

//Thunder resource SlbTemplateServer

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateServerCreate,
		UpdateContext: resourceSlbTemplateServerUpdate,
		ReadContext:   resourceSlbTemplateServerRead,
		DeleteContext: resourceSlbTemplateServerDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_query_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dns_fail_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic_server_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_selection_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_dynamic_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_ttl_ratio": {
				Type:        schema.TypeInt,
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
			"add": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"times": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"every": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"till": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_acct": {
				Type:        schema.TypeString,
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
			"bw_rate_limit_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateServer (Inside resourceSlbTemplateServerCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateServer(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServer --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateServer(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateServer (Inside resourceSlbTemplateServerRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateServer(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceSlbTemplateServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateServer   (Inside resourceSlbTemplateServerUpdate) ")
		data := dataToSlbTemplateServer(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServer ")
		err := go_thunder.PutSlbTemplateServer(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateServerRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateServerDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateServer(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateServer(d *schema.ResourceData) go_thunder.SlbTemplateServer {
	var vc go_thunder.SlbTemplateServer
	var c go_thunder.SlbTemplateServerInstance
	c.SlbTemplateServerInstanceName = d.Get("name").(string)
	c.SlbTemplateServerInstanceConnLimit = d.Get("conn_limit").(int)
	c.SlbTemplateServerInstanceResume = d.Get("resume").(int)
	c.SlbTemplateServerInstanceConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	c.SlbTemplateServerInstanceConnRateLimit = d.Get("conn_rate_limit").(int)
	c.SlbTemplateServerInstanceRateInterval = d.Get("rate_interval").(string)
	c.SlbTemplateServerInstanceConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.SlbTemplateServerInstanceDNSQueryInterval = d.Get("dns_query_interval").(int)
	c.SlbTemplateServerInstanceDNSFailInterval = d.Get("dns_fail_interval").(int)
	c.SlbTemplateServerInstanceDynamicServerPrefix = d.Get("dynamic_server_prefix").(string)
	c.SlbTemplateServerInstanceExtendedStats = d.Get("extended_stats").(int)
	c.SlbTemplateServerInstanceLogSelectionFailure = d.Get("log_selection_failure").(int)
	c.SlbTemplateServerInstanceHealthCheck = d.Get("health_check").(string)
	c.SlbTemplateServerInstanceHealthCheckDisable = d.Get("health_check_disable").(int)
	c.SlbTemplateServerInstanceMaxDynamicServer = d.Get("max_dynamic_server").(int)
	c.SlbTemplateServerInstanceMinTTLRatio = d.Get("min_ttl_ratio").(int)
	c.SlbTemplateServerInstanceWeight = d.Get("weight").(int)
	c.SlbTemplateServerInstanceSpoofingCache = d.Get("spoofing_cache").(int)
	c.SlbTemplateServerInstanceStatsDataAction = d.Get("stats_data_action").(string)
	c.SlbTemplateServerInstanceSlowStart = d.Get("slow_start").(int)
	c.SlbTemplateServerInstanceInitialSlowStart = d.Get("initial_slow_start").(int)
	c.SlbTemplateServerInstanceAdd = d.Get("add").(int)
	c.SlbTemplateServerInstanceTimes = d.Get("times").(int)
	c.SlbTemplateServerInstanceEvery = d.Get("every").(int)
	c.SlbTemplateServerInstanceTill = d.Get("till").(int)
	c.SlbTemplateServerInstanceBwRateLimitAcct = d.Get("bw_rate_limit_acct").(string)
	c.SlbTemplateServerInstanceBwRateLimit = d.Get("bw_rate_limit").(int)
	c.SlbTemplateServerInstanceBwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	c.SlbTemplateServerInstanceBwRateLimitDuration = d.Get("bw_rate_limit_duration").(int)
	c.SlbTemplateServerInstanceBwRateLimitNoLogging = d.Get("bw_rate_limit_no_logging").(int)
	c.SlbTemplateServerInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateServerInstanceName = c
	return vc
}
