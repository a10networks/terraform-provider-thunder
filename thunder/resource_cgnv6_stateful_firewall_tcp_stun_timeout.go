package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallTcpStunTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_tcp_stun_timeout`: Set STUN timeout for endpoint-independent filtering\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallTcpStunTimeoutCreate,
		UpdateContext: resourceCgnv6StatefulFirewallTcpStunTimeoutUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallTcpStunTimeoutRead,
		DeleteContext: resourceCgnv6StatefulFirewallTcpStunTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Single Destination Port or Port Range Start",
			},
			"port_end": {
				Type: schema.TypeInt, Required: true, Description: "Port Range End",
			},
			"stun_timeout_val_port_range": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "STUN timeout (default: 2minutes)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallTcpStunTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpStunTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpStunTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallTcpStunTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallTcpStunTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpStunTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpStunTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallTcpStunTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallTcpStunTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpStunTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpStunTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallTcpStunTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallTcpStunTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallTcpStunTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6StatefulFirewallTcpStunTimeout(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallTcpStunTimeout {
	var ret edpt.Cgnv6StatefulFirewallTcpStunTimeout
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortEnd = d.Get("port_end").(int)
	ret.Inst.StunTimeoutValPortRange = d.Get("stun_timeout_val_port_range").(int)
	//omit uuid
	return ret
}
