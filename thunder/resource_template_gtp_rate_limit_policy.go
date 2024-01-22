package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateGtpRateLimitPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_gtp_rate_limit_policy`: Configure GTP Rate Limit policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateGtpRateLimitPolicyCreate,
		UpdateContext: resourceTemplateGtpRateLimitPolicyUpdate,
		ReadContext:   resourceTemplateGtpRateLimitPolicyRead,
		DeleteContext: resourceTemplateGtpRateLimitPolicyDelete,

		Schema: map[string]*schema.Schema{
			"gtp_u_downlink_byte_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U bytes in downlink direction",
			},
			"gtp_u_downlink_packet_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U packets in downlink direction",
			},
			"gtp_u_max_concurrent_tunnels": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed GTP-U tunnels per Peer/APN Filter",
			},
			"gtp_u_total_byte_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U bytes total in both directions",
			},
			"gtp_u_total_packet_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U packets total in both directions",
			},
			"gtp_u_tunnel_create_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U tunnels",
			},
			"gtp_u_uplink_byte_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U bytes in uplink direction",
			},
			"gtp_u_uplink_packet_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed of rate of GTP-U packets in uplink direction",
			},
			"lockout": {
				Type: schema.TypeInt, Optional: true, Description: "Lockout traffic from the source for a certain time period after rate exceeded (Lockout duration in minutes)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the GTP Rate Limit Policy",
			},
			"rate_limit_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'monitor': Forward packet exceeding limit; 'drop': drop packet exceeding limit(default);",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v0_agg_message_type_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed GTPv0-C message rate of all types",
			},
			"v1_agg_message_type_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed GTPv1-C message rate of all types",
			},
			"v1_create_pdp_request_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed rate of GTPv1-C Create PDP Request message type",
			},
			"v1_update_pdp_request_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed rate of GTPv1-C Update PDP Request message type",
			},
			"v2_agg_message_type_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed GTPv2-C message rate of all types",
			},
			"v2_create_session_request_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed rate of GTPv2-C Create Session Request message type",
			},
			"v2_modify_bearer_request_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum allowed rate of GTPv2-C Modify Bearer Request message type",
			},
		},
	}
}
func resourceTemplateGtpRateLimitPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpRateLimitPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpRateLimitPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpRateLimitPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateGtpRateLimitPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpRateLimitPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpRateLimitPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateGtpRateLimitPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateGtpRateLimitPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpRateLimitPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpRateLimitPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateGtpRateLimitPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateGtpRateLimitPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateGtpRateLimitPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateGtpRateLimitPolicy(d *schema.ResourceData) edpt.TemplateGtpRateLimitPolicy {
	var ret edpt.TemplateGtpRateLimitPolicy
	ret.Inst.GtpUDownlinkByteRate = d.Get("gtp_u_downlink_byte_rate").(int)
	ret.Inst.GtpUDownlinkPacketRate = d.Get("gtp_u_downlink_packet_rate").(int)
	ret.Inst.GtpUMaxConcurrentTunnels = d.Get("gtp_u_max_concurrent_tunnels").(int)
	ret.Inst.GtpUTotalByteRate = d.Get("gtp_u_total_byte_rate").(int)
	ret.Inst.GtpUTotalPacketRate = d.Get("gtp_u_total_packet_rate").(int)
	ret.Inst.GtpUTunnelCreateRate = d.Get("gtp_u_tunnel_create_rate").(int)
	ret.Inst.GtpUUplinkByteRate = d.Get("gtp_u_uplink_byte_rate").(int)
	ret.Inst.GtpUUplinkPacketRate = d.Get("gtp_u_uplink_packet_rate").(int)
	ret.Inst.Lockout = d.Get("lockout").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RateLimitAction = d.Get("rate_limit_action").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.V0AggMessageTypeRate = d.Get("v0_agg_message_type_rate").(int)
	ret.Inst.V1AggMessageTypeRate = d.Get("v1_agg_message_type_rate").(int)
	ret.Inst.V1CreatePdpRequestRate = d.Get("v1_create_pdp_request_rate").(int)
	ret.Inst.V1UpdatePdpRequestRate = d.Get("v1_update_pdp_request_rate").(int)
	ret.Inst.V2AggMessageTypeRate = d.Get("v2_agg_message_type_rate").(int)
	ret.Inst.V2CreateSessionRequestRate = d.Get("v2_create_session_request_rate").(int)
	ret.Inst.V2ModifyBearerRequestRate = d.Get("v2_modify_bearer_request_rate").(int)
	return ret
}
