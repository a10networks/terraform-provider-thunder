package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6FragmentationOutbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_fragmentation_outbound`: LW-4O6 fragmentation rules for outbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6FragmentationOutboundCreate,
		UpdateContext: resourceCgnv6Lw4o6FragmentationOutboundUpdate,
		ReadContext:   resourceCgnv6Lw4o6FragmentationOutboundRead,
		DeleteContext: resourceCgnv6Lw4o6FragmentationOutboundDelete,

		Schema: map[string]*schema.Schema{
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set)(default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"frag_action": {
				Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets (default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Lw4o6FragmentationOutboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationOutboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationOutbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6FragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6FragmentationOutboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationOutboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationOutbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6FragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6FragmentationOutboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationOutboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationOutbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6FragmentationOutboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationOutboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationOutbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Lw4o6FragmentationOutbound(d *schema.ResourceData) edpt.Cgnv6Lw4o6FragmentationOutbound {
	var ret edpt.Cgnv6Lw4o6FragmentationOutbound
	ret.Inst.DfSet = d.Get("df_set").(string)
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
