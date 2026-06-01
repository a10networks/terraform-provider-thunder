package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallUdpIdleTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_udp_idle_timeout`: Configure UDP session idle timeout for IPv4 and IPv6\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallUdpIdleTimeoutCreate,
		UpdateContext: resourceCgnv6StatefulFirewallUdpIdleTimeoutUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallUdpIdleTimeoutRead,
		DeleteContext: resourceCgnv6StatefulFirewallUdpIdleTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"fast": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fast aging for idle sessions",
			},
			"idle_timeout_val_port_range": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Idle timeout for IPv4 and IPv6 TCP established sessions (Idle timeout for IPv4 and IPv6 TCP established sessions (default: 300 seconds))",
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Single Destination Port or Port Range Start",
			},
			"port_end": {
				Type: schema.TypeInt, Required: true, Description: "Port Range End",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallUdpIdleTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallUdpIdleTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallUdpIdleTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallUdpIdleTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallUdpIdleTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallUdpIdleTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallUdpIdleTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallUdpIdleTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallUdpIdleTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallUdpIdleTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallUdpIdleTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallUdpIdleTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallUdpIdleTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallUdpIdleTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6StatefulFirewallUdpIdleTimeout(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallUdpIdleTimeout {
	var ret edpt.Cgnv6StatefulFirewallUdpIdleTimeout
	ret.Inst.Fast = d.Get("fast").(int)
	ret.Inst.IdleTimeoutValPortRange = d.Get("idle_timeout_val_port_range").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortEnd = d.Get("port_end").(int)
	//omit uuid
	return ret
}
