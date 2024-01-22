package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SixrdFragmentationOutbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sixrd_fragmentation_outbound`: sixrd fragmentation rules for outbound oversize packets (default: ipv6)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SixrdFragmentationOutboundCreate,
		UpdateContext: resourceCgnv6SixrdFragmentationOutboundUpdate,
		ReadContext:   resourceCgnv6SixrdFragmentationOutboundRead,
		DeleteContext: resourceCgnv6SixrdFragmentationOutboundDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "ipv6", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 Fragmentation for oversize packets (default); 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"count1": {
				Type: schema.TypeInt, Optional: true, Description: "Configure number of ICMP messages sent when DF set. Default is 1",
			},
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 Fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SixrdFragmentationOutboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationOutboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationOutbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SixrdFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SixrdFragmentationOutboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationOutboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationOutbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SixrdFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SixrdFragmentationOutboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationOutboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationOutbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SixrdFragmentationOutboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationOutboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationOutbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SixrdFragmentationOutbound(d *schema.ResourceData) edpt.Cgnv6SixrdFragmentationOutbound {
	var ret edpt.Cgnv6SixrdFragmentationOutbound
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.DfSet = d.Get("df_set").(string)
	//omit uuid
	return ret
}
