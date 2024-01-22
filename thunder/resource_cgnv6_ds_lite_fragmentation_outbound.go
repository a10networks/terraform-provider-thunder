package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6DsLiteFragmentationOutbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_ds_lite_fragmentation_outbound`: DS-Lite Outbound Fragmentation Rules (default: ipv4)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6DsLiteFragmentationOutboundCreate,
		UpdateContext: resourceCgnv6DsLiteFragmentationOutboundUpdate,
		ReadContext:   resourceCgnv6DsLiteFragmentationOutboundRead,
		DeleteContext: resourceCgnv6DsLiteFragmentationOutboundDelete,

		Schema: map[string]*schema.Schema{
			"count1": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure number of ICMP messages sent when DF set. Default is 1",
			},
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 Fragmentation; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"frag_action": {
				Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets (default); 'send-icmpv6': Send ICMPv6 Type 2 Code 0 (Packet Too Big);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6DsLiteFragmentationOutboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteFragmentationOutboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteFragmentationOutbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6DsLiteFragmentationOutboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteFragmentationOutboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteFragmentationOutbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6DsLiteFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6DsLiteFragmentationOutboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteFragmentationOutboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteFragmentationOutbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6DsLiteFragmentationOutboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6DsLiteFragmentationOutboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6DsLiteFragmentationOutbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6DsLiteFragmentationOutbound(d *schema.ResourceData) edpt.Cgnv6DsLiteFragmentationOutbound {
	var ret edpt.Cgnv6DsLiteFragmentationOutbound
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.DfSet = d.Get("df_set").(string)
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
