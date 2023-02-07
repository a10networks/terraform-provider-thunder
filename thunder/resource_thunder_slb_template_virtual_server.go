package thunder

//Thunder resource SlbTemplateVirtualServer

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateVirtualServer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateVirtualServerCreate,
		UpdateContext: resourceSlbTemplateVirtualServerUpdate,
		ReadContext:   resourceSlbTemplateVirtualServerRead,
		DeleteContext: resourceSlbTemplateVirtualServerDelete,
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
			"conn_rate_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_limit_no_logging": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmp_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmp_lockup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmp_lockup_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmpv6_rate_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmpv6_lockup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"icmpv6_lockup_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tcp_stack_tfo_active_conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tcp_stack_tfo_cookie_time_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"tcp_stack_tfo_backoff_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"subnet_gratuitous_arp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_when_all_ports_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"disable_when_any_port_down": {
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

func resourceSlbTemplateVirtualServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateVirtualServer (Inside resourceSlbTemplateVirtualServerCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateVirtualServer(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateVirtualServer --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateVirtualServer(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateVirtualServerRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateVirtualServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateVirtualServer (Inside resourceSlbTemplateVirtualServerRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateVirtualServer(client.Token, name1, client.Host)
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

func resourceSlbTemplateVirtualServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateVirtualServer   (Inside resourceSlbTemplateVirtualServerUpdate) ")
		data := dataToSlbTemplateVirtualServer(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateVirtualServer ")
		err := go_thunder.PutSlbTemplateVirtualServer(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateVirtualServerRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateVirtualServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateVirtualServerDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateVirtualServer(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateVirtualServer(d *schema.ResourceData) go_thunder.SlbTemplateVirtualServer {
	var vc go_thunder.SlbTemplateVirtualServer
	var c go_thunder.SlbTemplateVirtualServerInstance
	c.SlbTemplateVirtualServerInstanceName = d.Get("name").(string)
	c.SlbTemplateVirtualServerInstanceConnLimit = d.Get("conn_limit").(int)
	c.SlbTemplateVirtualServerInstanceConnLimitReset = d.Get("conn_limit_reset").(int)
	c.SlbTemplateVirtualServerInstanceConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	c.SlbTemplateVirtualServerInstanceConnRateLimit = d.Get("conn_rate_limit").(int)
	c.SlbTemplateVirtualServerInstanceRateInterval = d.Get("rate_interval").(string)
	c.SlbTemplateVirtualServerInstanceConnRateLimitReset = d.Get("conn_rate_limit_reset").(int)
	c.SlbTemplateVirtualServerInstanceConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.SlbTemplateVirtualServerInstanceIcmpRateLimit = d.Get("icmp_rate_limit").(int)
	c.SlbTemplateVirtualServerInstanceIcmpLockup = d.Get("icmp_lockup").(int)
	c.SlbTemplateVirtualServerInstanceIcmpLockupPeriod = d.Get("icmp_lockup_period").(int)
	c.SlbTemplateVirtualServerInstanceIcmpv6RateLimit = d.Get("icmpv6_rate_limit").(int)
	c.SlbTemplateVirtualServerInstanceIcmpv6Lockup = d.Get("icmpv6_lockup").(int)
	c.SlbTemplateVirtualServerInstanceIcmpv6LockupPeriod = d.Get("icmpv6_lockup_period").(int)
	c.SlbTemplateVirtualServerInstanceTCPStackTfoActiveConnLimit = d.Get("tcp_stack_tfo_active_conn_limit").(int)
	c.SlbTemplateVirtualServerInstanceTCPStackTfoCookieTimeLimit = d.Get("tcp_stack_tfo_cookie_time_limit").(int)
	c.SlbTemplateVirtualServerInstanceTCPStackTfoBackoffTime = d.Get("tcp_stack_tfo_backoff_time").(int)
	c.SlbTemplateVirtualServerInstanceSubnetGratuitousArp = d.Get("subnet_gratuitous_arp").(int)
	c.SlbTemplateVirtualServerInstanceDisableWhenAllPortsDown = d.Get("disable_when_all_ports_down").(int)
	c.SlbTemplateVirtualServerInstanceDisableWhenAnyPortDown = d.Get("disable_when_any_port_down").(int)
	c.SlbTemplateVirtualServerInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateVirtualServerInstanceName = c
	return vc
}
