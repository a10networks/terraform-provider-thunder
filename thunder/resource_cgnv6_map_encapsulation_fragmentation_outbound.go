package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapEncapsulationFragmentationOutbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_encapsulation_fragmentation_outbound`: MAP-E fragmentation rules for outbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapEncapsulationFragmentationOutboundCreate,
		UpdateContext: resourceCgnv6MapEncapsulationFragmentationOutboundUpdate,
		ReadContext:   resourceCgnv6MapEncapsulationFragmentationOutboundRead,
		DeleteContext: resourceCgnv6MapEncapsulationFragmentationOutboundDelete,

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
func resourceCgnv6MapEncapsulationFragmentationOutboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationFragmentationOutboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationFragmentationOutbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapEncapsulationFragmentationOutboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationFragmentationOutboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationFragmentationOutbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapEncapsulationFragmentationOutboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationFragmentationOutboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationFragmentationOutbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapEncapsulationFragmentationOutboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationFragmentationOutboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationFragmentationOutbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapEncapsulationFragmentationOutbound(d *schema.ResourceData) edpt.Cgnv6MapEncapsulationFragmentationOutbound {
	var ret edpt.Cgnv6MapEncapsulationFragmentationOutbound
	ret.Inst.DfSet = d.Get("df_set").(string)
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
