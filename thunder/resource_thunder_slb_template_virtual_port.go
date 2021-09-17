package thunder

//Thunder resource SlbTemplateVirtualPort

import (
	"context"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSlbTemplateVirtualPort() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateVirtualPortCreate,
		UpdateContext: resourceSlbTemplateVirtualPortUpdate,
		ReadContext:   resourceSlbTemplateVirtualPortRead,
		DeleteContext: resourceSlbTemplateVirtualPortDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"aflow": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allow_syn_otherflags": {
				Type:        schema.TypeInt,
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
			"pkt_rate_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pkt_rate_limit_reset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"log_options": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"when_rr_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allow_vip_to_rport_mapping": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dscp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"drop_unknown_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_unknown_conn": {
				Type:        schema.TypeInt,
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
			"snat_msl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snat_port_preserve": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"non_syn_initiation": {
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

func resourceSlbTemplateVirtualPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateVirtualPort (Inside resourceSlbTemplateVirtualPortCreate) ")
		name1 := d.Get("name").(string)
		data := dataToSlbTemplateVirtualPort(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateVirtualPort --")
		d.SetId(name1)
		err := go_thunder.PostSlbTemplateVirtualPort(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateVirtualPortRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateVirtualPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateVirtualPort (Inside resourceSlbTemplateVirtualPortRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetSlbTemplateVirtualPort(client.Token, name1, client.Host)
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

func resourceSlbTemplateVirtualPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying SlbTemplateVirtualPort   (Inside resourceSlbTemplateVirtualPortUpdate) ")
		data := dataToSlbTemplateVirtualPort(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateVirtualPort ")
		err := go_thunder.PutSlbTemplateVirtualPort(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateVirtualPortRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateVirtualPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateVirtualPortDelete) " + name1)
		err := go_thunder.DeleteSlbTemplateVirtualPort(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToSlbTemplateVirtualPort(d *schema.ResourceData) go_thunder.SlbTemplateVirtualPort {
	var vc go_thunder.SlbTemplateVirtualPort
	var c go_thunder.SlbTemplateVirtualPortInstance
	c.SlbTemplateVirtualPortInstanceName = d.Get("name").(string)
	c.SlbTemplateVirtualPortInstanceAflow = d.Get("aflow").(int)
	c.SlbTemplateVirtualPortInstanceAllowSynOtherflags = d.Get("allow_syn_otherflags").(int)
	c.SlbTemplateVirtualPortInstanceConnLimit = d.Get("conn_limit").(int)
	c.SlbTemplateVirtualPortInstanceConnLimitReset = d.Get("conn_limit_reset").(int)
	c.SlbTemplateVirtualPortInstanceConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	c.SlbTemplateVirtualPortInstanceConnRateLimit = d.Get("conn_rate_limit").(int)
	c.SlbTemplateVirtualPortInstanceRateInterval = d.Get("rate_interval").(string)
	c.SlbTemplateVirtualPortInstanceConnRateLimitReset = d.Get("conn_rate_limit_reset").(int)
	c.SlbTemplateVirtualPortInstanceConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	c.SlbTemplateVirtualPortInstancePktRateType = d.Get("pkt_rate_type").(string)
	c.SlbTemplateVirtualPortInstanceRate = d.Get("rate").(int)
	c.SlbTemplateVirtualPortInstancePktRateInterval = d.Get("pkt_rate_interval").(string)
	c.SlbTemplateVirtualPortInstancePktRateLimitReset = d.Get("pkt_rate_limit_reset").(int)
	c.SlbTemplateVirtualPortInstanceLogOptions = d.Get("log_options").(string)
	c.SlbTemplateVirtualPortInstanceWhenRrEnable = d.Get("when_rr_enable").(int)
	c.SlbTemplateVirtualPortInstanceAllowVipToRportMapping = d.Get("allow_vip_to_rport_mapping").(int)
	c.SlbTemplateVirtualPortInstanceDscp = d.Get("dscp").(int)
	c.SlbTemplateVirtualPortInstanceDropUnknownConn = d.Get("drop_unknown_conn").(int)
	c.SlbTemplateVirtualPortInstanceResetUnknownConn = d.Get("reset_unknown_conn").(int)
	c.SlbTemplateVirtualPortInstanceResetL7OnFailover = d.Get("reset_l7_on_failover").(int)
	c.SlbTemplateVirtualPortInstanceIgnoreTCPMsl = d.Get("ignore_tcp_msl").(int)
	c.SlbTemplateVirtualPortInstanceSnatMsl = d.Get("snat_msl").(int)
	c.SlbTemplateVirtualPortInstanceSnatPortPreserve = d.Get("snat_port_preserve").(int)
	c.SlbTemplateVirtualPortInstanceNonSynInitiation = d.Get("non_syn_initiation").(int)
	c.SlbTemplateVirtualPortInstanceUserTag = d.Get("user_tag").(string)

	vc.SlbTemplateVirtualPortInstanceName = c
	return vc
}
