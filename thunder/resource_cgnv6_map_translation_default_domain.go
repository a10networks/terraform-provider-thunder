package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationDefaultDomain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_default_domain`: Set Default MAP-T domain\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationDefaultDomainCreate,
		UpdateContext: resourceCgnv6MapTranslationDefaultDomainUpdate,
		ReadContext:   resourceCgnv6MapTranslationDefaultDomainRead,
		DeleteContext: resourceCgnv6MapTranslationDefaultDomainDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "MAP-T domain name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6MapTranslationDefaultDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDefaultDomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDefaultDomain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDefaultDomainRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationDefaultDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDefaultDomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDefaultDomain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDefaultDomainRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationDefaultDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDefaultDomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDefaultDomain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationDefaultDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDefaultDomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDefaultDomain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapTranslationDefaultDomain(d *schema.ResourceData) edpt.Cgnv6MapTranslationDefaultDomain {
	var ret edpt.Cgnv6MapTranslationDefaultDomain
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
