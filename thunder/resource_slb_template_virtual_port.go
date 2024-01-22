package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateVirtualPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_virtual_port`: Virtual port template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateVirtualPortCreate,
		UpdateContext: resourceSlbTemplateVirtualPortUpdate,
		ReadContext:   resourceSlbTemplateVirtualPortRead,
		DeleteContext: resourceSlbTemplateVirtualPortDelete,

		Schema: map[string]*schema.Schema{
			"aflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use aFlow to eliminate the traffic surge",
			},
			"allow_syn_otherflags": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow initial SYN packet with other flags",
			},
			"allow_vip_to_rport_mapping": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow mapping of VIP to real port",
			},
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
			"drop_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop conection if receives TCP packet without SYN or RST flag and it does not belong to any existing connections",
			},
			"dscp": {
				Type: schema.TypeInt, Optional: true, Description: "Differentiated Services Code Point (DSCP to Real Server IP Mapping Value)",
			},
			"ignore_tcp_msl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "reclaim TCP resource immediately without MSL",
			},
			"log_options": {
				Type: schema.TypeString, Optional: true, Description: "'no-logging': Do not log over limit event; 'no-repeat-logging': log once for over limit event. Default is log once per minute;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Virtual port template name",
			},
			"non_syn_initiation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow initial TCP packet to be non-SYN",
			},
			"pkt_rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "second", Description: "'100ms': Source IP and port rate limit per 100ms; 'second': Source IP and port rate limit per second (default);",
			},
			"pkt_rate_limit_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "send client-side reset (reset after packet limit)",
			},
			"pkt_rate_type": {
				Type: schema.TypeString, Optional: true, Description: "'src-ip-port': Source IP and port rate limit; 'src-port': Source port rate limit;",
			},
			"rate": {
				Type: schema.TypeInt, Optional: true, Description: "Source IP and port rate limit (Packet rate limit)",
			},
			"rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "second", Description: "'100ms': Use 100 ms as sampling interval; 'second': Use 1 second as sampling interval;",
			},
			"reset_l7_on_failover": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset to L7 client and server connection upon a failover",
			},
			"reset_unknown_conn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset back if receives TCP packet without SYN or RST flag and it does not belong to any existing connections",
			},
			"snat_msl": {
				Type: schema.TypeInt, Optional: true, Description: "Source NAT MSL (Source NAT MSL value (seconds))",
			},
			"snat_port_preserve": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source NAT Port Preservation",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"when_rr_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only do rate limit if CPU RR triggered",
			},
		},
	}
}
func resourceSlbTemplateVirtualPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateVirtualPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateVirtualPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateVirtualPortRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateVirtualPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateVirtualPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateVirtualPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateVirtualPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateVirtualPort(d *schema.ResourceData) edpt.SlbTemplateVirtualPort {
	var ret edpt.SlbTemplateVirtualPort
	ret.Inst.Aflow = d.Get("aflow").(int)
	ret.Inst.AllowSynOtherflags = d.Get("allow_syn_otherflags").(int)
	ret.Inst.AllowVipToRportMapping = d.Get("allow_vip_to_rport_mapping").(int)
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnLimitNoLogging = d.Get("conn_limit_no_logging").(int)
	ret.Inst.ConnLimitReset = d.Get("conn_limit_reset").(int)
	ret.Inst.ConnRateLimit = d.Get("conn_rate_limit").(int)
	ret.Inst.ConnRateLimitNoLogging = d.Get("conn_rate_limit_no_logging").(int)
	ret.Inst.ConnRateLimitReset = d.Get("conn_rate_limit_reset").(int)
	ret.Inst.DropUnknownConn = d.Get("drop_unknown_conn").(int)
	ret.Inst.Dscp = d.Get("dscp").(int)
	ret.Inst.IgnoreTcpMsl = d.Get("ignore_tcp_msl").(int)
	ret.Inst.LogOptions = d.Get("log_options").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NonSynInitiation = d.Get("non_syn_initiation").(int)
	ret.Inst.PktRateInterval = d.Get("pkt_rate_interval").(string)
	ret.Inst.PktRateLimitReset = d.Get("pkt_rate_limit_reset").(int)
	ret.Inst.PktRateType = d.Get("pkt_rate_type").(string)
	ret.Inst.Rate = d.Get("rate").(int)
	ret.Inst.RateInterval = d.Get("rate_interval").(string)
	ret.Inst.ResetL7OnFailover = d.Get("reset_l7_on_failover").(int)
	ret.Inst.ResetUnknownConn = d.Get("reset_unknown_conn").(int)
	ret.Inst.SnatMsl = d.Get("snat_msl").(int)
	ret.Inst.SnatPortPreserve = d.Get("snat_port_preserve").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.WhenRrEnable = d.Get("when_rr_enable").(int)
	return ret
}
