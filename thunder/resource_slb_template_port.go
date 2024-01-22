package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_port`: Port template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePortCreate,
		UpdateContext: resourceSlbTemplatePortUpdate,
		ReadContext:   resourceSlbTemplatePortRead,
		DeleteContext: resourceSlbTemplatePortDelete,

		Schema: map[string]*schema.Schema{
			"add": {
				Type: schema.TypeInt, Optional: true, Description: "Slow start connection limit add by a number every interval (Add by this number every interval)",
			},
			"bw_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Configure bandwidth rate limit on real server port (Bandwidth rate limit in Kbps)",
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
			"dampening_flaps": {
				Type: schema.TypeInt, Optional: true, Description: "service dampening flaps count (max-flaps allowed in flap period)",
			},
			"decrement": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Decrease after every round of DNS query (default is 0)",
			},
			"del_session_on_server_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete session if the server/port goes down (either disabled/hm down)",
			},
			"dest_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Destination NAT",
			},
			"down_grace_period": {
				Type: schema.TypeInt, Optional: true, Description: "Port down grace period (Down grace period in seconds)",
			},
			"down_timer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The timer to bring the marked down server/port to up (default is 0, never bring up) (The timer to bring up server (in second, default is 0))",
			},
			"dscp": {
				Type: schema.TypeInt, Optional: true, Description: "Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)",
			},
			"dynamic_member_priority": {
				Type: schema.TypeInt, Optional: true, Default: 16, Description: "Set dynamic member's priority (Initial priority (default is 16))",
			},
			"every": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Slow start connection limit increment interval (default 10)",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on real server port",
			},
			"flap_period": {
				Type: schema.TypeInt, Optional: true, Description: "take service out of rotation if max-flaps exceeded within time in seconds",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Health monitor name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
			},
			"inband_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use inband traffic to detect port's health status",
			},
			"initial_slow_start": {
				Type: schema.TypeInt, Optional: true, Default: 128, Description: "Initial slow start connection limit (default 128)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Port template name",
			},
			"no_ssl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "No SSL",
			},
			"rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "second", Description: "'100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;",
			},
			"reassign": {
				Type: schema.TypeInt, Optional: true, Default: 25, Description: "Maximum reassign times before declear the server/port down (default is 25) (The maximum reassign number)",
			},
			"request_rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "second", Description: "'100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;",
			},
			"request_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Request rate limit",
			},
			"request_rate_no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
			},
			"resel_on_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "When receiving reset from server, do the server/port reselection (default is 0, don't do reselection)",
			},
			"reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when connection rate over limit",
			},
			"restore_svc_time": {
				Type: schema.TypeInt, Optional: true, Description: "put the service back to the rotation after time in seconds",
			},
			"resume": {
				Type: schema.TypeInt, Optional: true, Description: "Resume accepting new connection after connection number drops below threshold (Connection resume threshold)",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Maximum retry times before reassign this connection to another server/port (default is 2) (The maximum retry number)",
			},
			"shared_partition_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a NAT pool or pool-group from shared partition",
			},
			"slow_start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Slowly ramp up the connection number after port is up",
			},
			"source_nat": {
				Type: schema.TypeString, Optional: true, Description: "Source NAT (IP NAT Pool or pool group name)",
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for real server port; 'stats-data-disable': Disable statistical data collection for real server port;",
			},
			"sub_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Divide service group members into different sub groups (Sub group ID (default is 0))",
			},
			"template_port_pool_shared": {
				Type: schema.TypeString, Optional: true, Description: "Source NAT (IP NAT Pool or pool group name)",
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
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Weight (port weight)",
			},
		},
	}
}
func resourceSlbTemplatePortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePortRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePortRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePort(d *schema.ResourceData) edpt.SlbTemplatePort {
	var ret edpt.SlbTemplatePort
	ret.Inst.Add = d.Get("add").(int)
	ret.Inst.BwRateLimit = d.Get("bw_rate_limit").(int)
	ret.Inst.BwRateLimitDuration = d.Get("bw_rate_limit_duration").(int)
	ret.Inst.BwRateLimitNoLogging = d.Get("bw_rate_limit_no_logging").(int)
	ret.Inst.BwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	ret.Inst.DampeningFlaps = d.Get("dampening_flaps").(int)
	ret.Inst.Decrement = d.Get("decrement").(int)
	ret.Inst.DelSessionOnServerDown = d.Get("del_session_on_server_down").(int)
	ret.Inst.DestNat = d.Get("dest_nat").(int)
	ret.Inst.DownGracePeriod = d.Get("down_grace_period").(int)
	ret.Inst.DownTimer = d.Get("down_timer").(int)
	ret.Inst.Dscp = d.Get("dscp").(int)
	ret.Inst.DynamicMemberPriority = d.Get("dynamic_member_priority").(int)
	ret.Inst.Every = d.Get("every").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.FlapPeriod = d.Get("flap_period").(int)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.InbandHealthCheck = d.Get("inband_health_check").(int)
	ret.Inst.InitialSlowStart = d.Get("initial_slow_start").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NoSsl = d.Get("no_ssl").(int)
	ret.Inst.RateInterval = d.Get("rate_interval").(string)
	ret.Inst.Reassign = d.Get("reassign").(int)
	ret.Inst.RequestRateInterval = d.Get("request_rate_interval").(string)
	ret.Inst.RequestRateLimit = d.Get("request_rate_limit").(int)
	ret.Inst.RequestRateNoLogging = d.Get("request_rate_no_logging").(int)
	ret.Inst.ReselOnReset = d.Get("resel_on_reset").(int)
	ret.Inst.Reset = d.Get("reset").(int)
	ret.Inst.RestoreSvcTime = d.Get("restore_svc_time").(int)
	ret.Inst.Resume = d.Get("resume").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	ret.Inst.SlowStart = d.Get("slow_start").(int)
	ret.Inst.SourceNat = d.Get("source_nat").(string)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.SubGroup = d.Get("sub_group").(int)
	ret.Inst.TemplatePortPoolShared = d.Get("template_port_pool_shared").(string)
	ret.Inst.Till = d.Get("till").(int)
	ret.Inst.Times = d.Get("times").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	return ret
}
