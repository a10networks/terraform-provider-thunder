package thunder

//Thunder resource SlbResourceUsage

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbResourceUsage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbResourceUsageCreate,
		UpdateContext: resourceSlbResourceUsageUpdate,
		ReadContext:   resourceSlbResourceUsageRead,
		DeleteContext: resourceSlbResourceUsageDelete,
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

func resourceSlbResourceUsageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbResourceUsage (Inside resourceSlbResourceUsageCreate) ")

		data := dataToSlbResourceUsage(d)
		logger.Println("[INFO] received formatted data from method data to SlbResourceUsage --")
		d.SetId("1")
		err := go_thunder.PostSlbResourceUsage(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbResourceUsageRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbResourceUsageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbResourceUsage (Inside resourceSlbResourceUsageRead)")

	if client.Host != "" {

		data, err := go_thunder.GetSlbResourceUsage(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {

			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbResourceUsageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbResourceUsageRead(ctx, d, meta)
}

func resourceSlbResourceUsageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSlbResourceUsageRead(ctx, d, meta)
}

func dataToSlbResourceUsage(d *schema.ResourceData) go_thunder.ResourceUsage {
	var vc go_thunder.ResourceUsage
	var c go_thunder.ResourceUsageInstance
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
