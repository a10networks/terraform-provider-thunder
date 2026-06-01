package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallTcpIdleTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_tcp_idle_timeout`: Configure TCP established-session idle timeout for IPv4 and IPv6\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallTcpIdleTimeoutCreate,
		UpdateContext: resourceCgnv6StatefulFirewallTcpIdleTimeoutUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallTcpIdleTimeoutRead,
		DeleteContext: resourceCgnv6StatefulFirewallTcpIdleTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"idle_timeout_val_port_range": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "Set Idle timeout for IPv4 and IPv6 TCP established sessions (Idle timeout for IPv4 and IPv6 TCP established sessions (default: 300 seconds))",
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
func resourceCgnv6StatefulFirewallTcpIdleTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpIdleTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpIdleTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallTcpIdleTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallTcpIdleTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpIdleTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpIdleTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallTcpIdleTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallTcpIdleTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpIdleTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpIdleTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallTcpIdleTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpIdleTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpIdleTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6StatefulFirewallTcpIdleTimeout(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallTcpIdleTimeout {
	var ret edpt.Cgnv6StatefulFirewallTcpIdleTimeout
	ret.Inst.IdleTimeoutValPortRange = d.Get("idle_timeout_val_port_range").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortEnd = d.Get("port_end").(int)
	//omit uuid
	return ret
}
