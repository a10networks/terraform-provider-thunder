package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64FragmentationInbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_fragmentation_inbound`: Fragmentation Behavior from NAT outside to NAT inside\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64FragmentationInboundCreate,
		UpdateContext: resourceCgnv6Nat64FragmentationInboundUpdate,
		ReadContext:   resourceCgnv6Nat64FragmentationInboundRead,
		DeleteContext: resourceCgnv6Nat64FragmentationInboundDelete,

		Schema: map[string]*schema.Schema{
			"count1": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure number of ICMP messages sent when DF set. Default is 1",
			},
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 fragmentation; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default);",
			},
			"frag_action": {
				Type: schema.TypeString, Optional: true, Default: "ipv6", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 fragmentation for oversize packets (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat64FragmentationInboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationInboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationInbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64FragmentationInboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64FragmentationInboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationInboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationInbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64FragmentationInboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64FragmentationInboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationInboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationInbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64FragmentationInboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationInboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationInbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64FragmentationInbound(d *schema.ResourceData) edpt.Cgnv6Nat64FragmentationInbound {
	var ret edpt.Cgnv6Nat64FragmentationInbound
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.DfSet = d.Get("df_set").(string)
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
