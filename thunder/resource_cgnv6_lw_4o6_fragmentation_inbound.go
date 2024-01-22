package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6FragmentationInbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_fragmentation_inbound`: LW-4O6 fragmentation rules for inbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6FragmentationInboundCreate,
		UpdateContext: resourceCgnv6Lw4o6FragmentationInboundUpdate,
		ReadContext:   resourceCgnv6Lw4o6FragmentationInboundRead,
		DeleteContext: resourceCgnv6Lw4o6FragmentationInboundDelete,

		Schema: map[string]*schema.Schema{
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets; 'ipv6': Use IPv6 fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default);",
			},
			"frag_action": {
				Type: schema.TypeString, Optional: true, Default: "ipv6", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets; 'ipv6': Use IPv6 fragmentation for oversize packets (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Lw4o6FragmentationInboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationInboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationInbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6FragmentationInboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6FragmentationInboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationInboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationInbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6FragmentationInboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6FragmentationInboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationInboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationInbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6FragmentationInboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6FragmentationInboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6FragmentationInbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Lw4o6FragmentationInbound(d *schema.ResourceData) edpt.Cgnv6Lw4o6FragmentationInbound {
	var ret edpt.Cgnv6Lw4o6FragmentationInbound
	ret.Inst.DfSet = d.Get("df_set").(string)
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
