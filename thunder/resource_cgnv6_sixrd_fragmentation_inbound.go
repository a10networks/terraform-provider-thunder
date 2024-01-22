package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SixrdFragmentationInbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sixrd_fragmentation_inbound`: 6rd fragmentation rules for inbound oversize packets (default: send-icmpv6)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SixrdFragmentationInboundCreate,
		UpdateContext: resourceCgnv6SixrdFragmentationInboundUpdate,
		ReadContext:   resourceCgnv6SixrdFragmentationInboundRead,
		DeleteContext: resourceCgnv6SixrdFragmentationInboundDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "send-icmpv6", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets; 'ipv6': Use IPv6 Fragmentation for oversize packets; 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big) (default);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SixrdFragmentationInboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationInboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationInbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SixrdFragmentationInboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SixrdFragmentationInboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationInboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationInbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SixrdFragmentationInboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SixrdFragmentationInboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationInboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationInbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SixrdFragmentationInboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SixrdFragmentationInboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SixrdFragmentationInbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SixrdFragmentationInbound(d *schema.ResourceData) edpt.Cgnv6SixrdFragmentationInbound {
	var ret edpt.Cgnv6SixrdFragmentationInbound
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
