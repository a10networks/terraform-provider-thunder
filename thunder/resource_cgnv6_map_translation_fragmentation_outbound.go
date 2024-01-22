package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationFragmentationOutbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_fragmentation_outbound`: MAP-T fragmentation rules for outbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationFragmentationOutboundCreate,
		UpdateContext: resourceCgnv6MapTranslationFragmentationOutboundUpdate,
		ReadContext:   resourceCgnv6MapTranslationFragmentationOutboundRead,
		DeleteContext: resourceCgnv6MapTranslationFragmentationOutboundDelete,

		Schema: map[string]*schema.Schema{
			"frag_action": {
				Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'drop': Drop Silently; 'ipv4': Use IPv4 fragmentation for oversize packets (default); 'send-icmpv6': Send ICMP Type 2 Code 0 (Packet Too Big);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6MapTranslationFragmentationOutboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationOutboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationOutbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationFragmentationOutboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationOutboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationOutbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationFragmentationOutboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationFragmentationOutboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationOutboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationOutbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationFragmentationOutboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationOutboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationOutbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapTranslationFragmentationOutbound(d *schema.ResourceData) edpt.Cgnv6MapTranslationFragmentationOutbound {
	var ret edpt.Cgnv6MapTranslationFragmentationOutbound
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
