package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6StatefulFirewallStunTimeout() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_stateful_firewall_stun_timeout`: Set STUN timeout for endpoint-independent filtering\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6StatefulFirewallStunTimeoutCreate,
		UpdateContext: resourceCgnv6StatefulFirewallStunTimeoutUpdate,
		ReadContext:   resourceCgnv6StatefulFirewallStunTimeoutRead,
		DeleteContext: resourceCgnv6StatefulFirewallStunTimeoutDelete,

		Schema: map[string]*schema.Schema{
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Single Destination Port or Port Range Start",
			},
			"port_end": {
				Type: schema.TypeInt, Required: true, Description: "Port Range End",
			},
			"stun_timeout_val_port_range": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "STUN timeout (default: 2 minutes)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6StatefulFirewallStunTimeoutCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallStunTimeoutCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallStunTimeout(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallStunTimeoutRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6StatefulFirewallStunTimeoutUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallStunTimeoutUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallStunTimeout(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6StatefulFirewallStunTimeoutRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6StatefulFirewallStunTimeoutDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallStunTimeoutDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallStunTimeout(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6StatefulFirewallStunTimeoutRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6StatefulFirewallStunTimeoutRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6StatefulFirewallStunTimeout(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6StatefulFirewallStunTimeout(d *schema.ResourceData) edpt.Cgnv6StatefulFirewallStunTimeout {
	var ret edpt.Cgnv6StatefulFirewallStunTimeout
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortEnd = d.Get("port_end").(int)
	ret.Inst.StunTimeoutValPortRange = d.Get("stun_timeout_val_port_range").(int)
	//omit uuid
	return ret
}
