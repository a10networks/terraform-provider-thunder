package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateVirtualServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_virtual_server`: Virtual server template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateVirtualServerCreate,
		UpdateContext: resourceSlbTemplateVirtualServerUpdate,
		ReadContext:   resourceSlbTemplateVirtualServerRead,
		DeleteContext: resourceSlbTemplateVirtualServerDelete,

		Schema: map[string]*schema.Schema{
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection limit",
			},
			"conn_limit_no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
			},
			"conn_limit_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when connection over limit",
			},
			"conn_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Connection rate limit",
			},
			"conn_rate_limit_no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
			},
			"conn_rate_limit_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send client reset when connection rate over limit",
			},
			"disable_when_all_ports_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Virtual Server when all member ports are down",
			},
			"disable_when_any_port_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Virtual Server when any member port is down",
			},
			"icmp_lockup": {
				Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
			},
			"icmp_lockup_period": {
				Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
			},
			"icmp_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "ICMP rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)",
			},
			"icmpv6_lockup": {
				Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
			},
			"icmpv6_lockup_period": {
				Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
			},
			"icmpv6_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "ICMPv6 rate limit (Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Virtual server template name",
			},
			"rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "second", Description: "'100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;",
			},
			"subnet_gratuitous_arp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send gratuitous ARP for every IP in the subnet virtual server",
			},
			"tcp_stack_tfo_active_conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The allowed active layer 7 tcp fast-open connection limit, default is zero (number)",
			},
			"tcp_stack_tfo_backoff_time": {
				Type: schema.TypeInt, Optional: true, Default: 600, Description: "The time tcp stack will wait before allowing new fast-open requests after security condition, default 600 seconds (number)",
			},
			"tcp_stack_tfo_cookie_time_limit": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "The time limit (in seconds) that a layer 7 tcp fast-open cookie is valid, default is 60 seconds (number)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateVirtualServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateVirtualServerRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateVirtualServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateVirtualServerRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateVirtualServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateVirtualServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateVirtualServer(d *schema.ResourceData) edpt.SlbTemplateVirtualServer {
	var ret edpt.SlbTemplateVirtualServer
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	ret.Inst.ConnLimitReset = d.Get("conn_limit_reset").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	ret.Inst.ConnRateLimitReset = d.Get("conn_rate_limit_reset").(int)
	ret.Inst.DisableWhenAllPortsDown = d.Get("disable_when_all_ports_down").(int)
	ret.Inst.DisableWhenAnyPortDown = d.Get("disable_when_any_port_down").(int)
	ret.Inst.IcmpLockup = d.Get("icmp_lockup").(int)
	ret.Inst.IcmpLockupPeriod = d.Get("icmp_lockup_period").(int)
	ret.Inst.IcmpRateLimit = d.Get("icmp_rate_limit").(int)
	ret.Inst.Icmpv6Lockup = d.Get("icmpv6_lockup").(int)
	ret.Inst.Icmpv6LockupPeriod = d.Get("icmpv6_lockup_period").(int)
	ret.Inst.Icmpv6RateLimit = d.Get("icmpv6_rate_limit").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RateInterval = d.Get("rate_interval").(string)
	ret.Inst.SubnetGratuitousArp = d.Get("subnet_gratuitous_arp").(int)
	ret.Inst.TcpStackTfoActiveConnLimit = d.Get("tcp_stack_tfo_active_conn_limit").(int)
	ret.Inst.TcpStackTfoBackoffTime = d.Get("tcp_stack_tfo_backoff_time").(int)
	ret.Inst.TcpStackTfoCookieTimeLimit = d.Get("tcp_stack_tfo_cookie_time_limit").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
