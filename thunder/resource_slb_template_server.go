package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_server`: Server template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateServerCreate,
		UpdateContext: resourceSlbTemplateServerUpdate,
		ReadContext:   resourceSlbTemplateServerRead,
		DeleteContext: resourceSlbTemplateServerDelete,

		Schema: map[string]*schema.Schema{
			"add": {
				Type: schema.TypeInt, Optional: true, Description: "Slow start connection limit add by a number every interval (Add by this number every interval)",
			},
			"bw_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure bandwidth rate limit on real server (Bandwidth rate limit in Kbps)",
			},
			"bw_rate_limit_acct": {
				Type: schema.TypeString, Optional: true, Default: "all", Description: "'to-server-only': Only account for traffic sent to server; 'from-server-only': Only account for traffic received from server; 'all': Account for all traffic sent to and received from server;",
			},
			"bw_rate_limit_duration": {
				Type: schema.TypeInt, Optional: true, Description: "Duration in seconds the observed rate needs to honor",
			},
			"bw_rate_limit_no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log bandwidth rate limit related state transitions",
			},
			"bw_rate_limit_resume": {
				Type: schema.TypeInt, Optional: true, Description: "Resume server selection after bandwidth drops below this threshold (in Kbps) (Bandwidth rate limit resume threshold (in Kbps))",
			},
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection limit",
			},
			"conn_limit_no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
			},
			"conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
			},
			"conn_rate_limit_no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
			},
			"dns_fail_interval": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "The interval to retry when DNS failed to query (DNS failure interval (in second, default is 30))",
			},
			"dns_query_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "The interval to query DNS server for the hostname (DNS query interval (in minute, default is 10))",
			},
			"dynamic_server_prefix": {
				Type: schema.TypeString, Optional: true, Default: "DRS", Description: "Prefix of dynamic server (Prefix of dynamic server (default is \"DRS\"))",
			},
			"every": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Slow start connection limit increment interval (default 10)",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on real server",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Health monitor name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
			},
			"initial_slow_start": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Initial slow start connection limit (default 128)",
			},
			"log_selection_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable real-time logging for server selection failure event",
			},
			"max_dynamic_server": {
				Type: schema.TypeInt, Optional: true, Default: 255, Description: "Maximum dynamic server number (Maximum dynamic server number (default is 255))",
			},
			"min_ttl_ratio": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Minimum TTL to DNS query interval ratio (Minimum TTL ratio (default is 2))",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server template name",
			},
			"rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "second", Description: "'100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;",
			},
			"resume": {
				Type: schema.TypeInt, Optional: true, Description: "Resume accepting new connection after connection number drops below threshold (Connection resume threshold)",
			},
			"slow_start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Slowly ramp up the connection number after server is up",
			},
			"spoofing_cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Servers under the template are spoofing cache",
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for real server; 'stats-data-disable': Disable statistical data collection for real server;",
			},
			"till": {
				Type: schema.TypeInt, Optional: true, Default: 4096, Description: "Slow start ends when slow start connection limit reaches a number (default 4096) (Slow start ends when connection limit reaches this number)",
			},
			"times": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Slow start connection limit multiply by a number every interval (default 2) (Multiply by this number every interval)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Weight for the Real Servers (Connection Weight (default is 1))",
			},
		},
	}
}
func resourceSlbTemplateServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateServer(d *schema.ResourceData) edpt.SlbTemplateServer {
	var ret edpt.SlbTemplateServer
	ret.Inst.Add = d.Get("add").(int)
	ret.Inst.BwRateLimit = d.Get("bw_rate_limit").(int)
	ret.Inst.BwRateLimitAcct = d.Get("bw_rate_limit_acct").(string)
	ret.Inst.BwRateLimitDuration = d.Get("bw_rate_limit_duration").(int)
	ret.Inst.BwRateLimitNoLogging = d.Get("bw_rate_limit_no_logging").(int)
	ret.Inst.BwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	ret.Inst.DnsFailInterval = d.Get("dns_fail_interval").(int)
	ret.Inst.DnsQueryInterval = d.Get("dns_query_interval").(int)
	ret.Inst.DynamicServerPrefix = d.Get("dynamic_server_prefix").(string)
	ret.Inst.Every = d.Get("every").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.InitialSlowStart = d.Get("initial_slow_start").(int)
	ret.Inst.LogSelectionFailure = d.Get("log_selection_failure").(int)
	ret.Inst.MaxDynamicServer = d.Get("max_dynamic_server").(int)
	ret.Inst.MinTtlRatio = d.Get("min_ttl_ratio").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RateInterval = d.Get("rate_interval").(string)
	ret.Inst.Resume = d.Get("resume").(int)
	ret.Inst.SlowStart = d.Get("slow_start").(int)
	ret.Inst.SpoofingCache = d.Get("spoofing_cache").(int)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.Till = d.Get("till").(int)
	ret.Inst.Times = d.Get("times").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	return ret
}
