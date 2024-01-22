package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessFragmentationOutbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat46_stateless_fragmentation_outbound`: Stateless NAT46 fragmentation rules for outbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat46StatelessFragmentationOutboundCreate,
		UpdateContext: resourceCgnv6Nat46StatelessFragmentationOutboundUpdate,
		ReadContext:   resourceCgnv6Nat46StatelessFragmentationOutboundRead,
		DeleteContext: resourceCgnv6Nat46StatelessFragmentationOutboundDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "ipv6", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 Fragmentation for oversize packets (default);",
			},
			"count1": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure number of ICMP messages sent when DF set. Default is 1",
			},
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 Fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat46StatelessFragmentationOutboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationOutboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationOutbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat46StatelessFragmentationOutboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationOutboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationOutbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat46StatelessFragmentationOutboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationOutboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationOutbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat46StatelessFragmentationOutboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessFragmentationOutboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessFragmentationOutbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat46StatelessFragmentationOutbound(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessFragmentationOutbound {
	var ret edpt.Cgnv6Nat46StatelessFragmentationOutbound
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.DfSet = d.Get("df_set").(string)
	//omit uuid
	return ret
}
