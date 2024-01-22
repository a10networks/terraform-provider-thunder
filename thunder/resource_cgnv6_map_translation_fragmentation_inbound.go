package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationFragmentationInbound() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_fragmentation_inbound`: MAP-T fragmentation rules for inbound oversize packets\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationFragmentationInboundCreate,
		UpdateContext: resourceCgnv6MapTranslationFragmentationInboundUpdate,
		ReadContext:   resourceCgnv6MapTranslationFragmentationInboundRead,
		DeleteContext: resourceCgnv6MapTranslationFragmentationInboundDelete,

		Schema: map[string]*schema.Schema{
			"df_set": {
				Type: schema.TypeString, Optional: true, Default: "send-icmp", Description: "'drop': Drop Silently; 'ipv6': Use IPv6 fragmentation for oversize packets; 'send-icmp': Send ICMP Type 3 Code 4 (Fragmentation Needed and DF Set) (default);",
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
func resourceCgnv6MapTranslationFragmentationInboundCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationInboundCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationInbound(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationFragmentationInboundRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationFragmentationInboundUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationInboundUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationInbound(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationFragmentationInboundRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationFragmentationInboundDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationInboundDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationInbound(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationFragmentationInboundRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationFragmentationInboundRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationFragmentationInbound(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapTranslationFragmentationInbound(d *schema.ResourceData) edpt.Cgnv6MapTranslationFragmentationInbound {
	var ret edpt.Cgnv6MapTranslationFragmentationInbound
	ret.Inst.DfSet = d.Get("df_set").(string)
	ret.Inst.FragAction = d.Get("frag_action").(string)
	//omit uuid
	return ret
}
