package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebServiceSecureGenerate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_service_secure_generate`: Web-service secure generate operation\n\n__PLACEHOLDER__",
		CreateContext: resourceWebServiceSecureGenerateCreate,
		UpdateContext: resourceWebServiceSecureGenerateUpdate,
		ReadContext:   resourceWebServiceSecureGenerateRead,
		DeleteContext: resourceWebServiceSecureGenerateDelete,

		Schema: map[string]*schema.Schema{
			"country": {
				Type: schema.TypeString, Optional: true, Description: "The country name",
			},
			"domain_name": {
				Type: schema.TypeString, Optional: true, Description: "The domain name",
			},
			"state": {
				Type: schema.TypeString, Optional: true, Description: "The location",
			},
		},
	}
}
func resourceWebServiceSecureGenerateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureGenerateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureGenerate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureGenerateRead(ctx, d, meta)
	}
	return diags
}

func resourceWebServiceSecureGenerateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureGenerateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureGenerate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureGenerateRead(ctx, d, meta)
	}
	return diags
}
func resourceWebServiceSecureGenerateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureGenerateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureGenerate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebServiceSecureGenerateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureGenerateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureGenerate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointWebServiceSecureGenerate(d *schema.ResourceData) edpt.WebServiceSecureGenerate {
	var ret edpt.WebServiceSecureGenerate
	ret.Inst.Country = d.Get("country").(string)
	ret.Inst.DomainName = d.Get("domain_name").(string)
	ret.Inst.State = d.Get("state").(string)
	return ret
}
