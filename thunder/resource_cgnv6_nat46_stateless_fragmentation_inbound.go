package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessFragmentationInbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat46_stateless_fragmentation_inbound`: Stateless NAT46 fragmentation rules for inbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat46StatelessFragmentationInboundCreate,
		UpdateContext: resourceCgnv6Nat46StatelessFragmentationInboundUpdate,
		ReadContext:   resourceCgnv6Nat46StatelessFragmentationInboundRead,
		DeleteContext: resourceCgnv6Nat46StatelessFragmentationInboundDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets (default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat46StatelessFragmentationInboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationInboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationInbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessFragmentationInboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat46StatelessFragmentationInboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationInboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationInbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessFragmentationInboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat46StatelessFragmentationInboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationInboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationInbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat46StatelessFragmentationInboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationInboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationInbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat46StatelessFragmentationInbound(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessFragmentationInbound {
	var ret edpt.Cgnv6Nat46StatelessFragmentationInbound
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
