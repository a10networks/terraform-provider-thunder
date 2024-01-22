package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallAlgRtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_alg_rtp`: Configure FTP ALG for NAT stateful firewall (default: enabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallAlgRtpCreate,
		UpdateContext: resourceCgnv6StatefulFirewallAlgRtpUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallAlgRtpRead,
		DeleteContext: resourceCgnv6StatefulFirewallAlgRtpDelete,

		Schema: map[string]*schema.Schema{
			"rtp_stun_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "RTP/RTCP STUN timeout (default: 5 minutes)}",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallAlgRtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgRtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgRtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgRtpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgRtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgRtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgRtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallAlgRtpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallAlgRtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgRtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgRtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallAlgRtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallAlgRtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallAlgRtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6StatefulFirewallAlgRtp(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallAlgRtp {
	var ret edpt.Cgnv6StatefulFirewallAlgRtp
	ret.Inst.RtpStunTimeout = d.Get("rtp_stun_timeout").(int)
	//omit uuid
	return ret
}
