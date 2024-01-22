package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebServiceSecureRegenerate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_service_secure_regenerate`: Web-service secure regenerate operation\n\n__PLACEHOLDER__",
		CreateContext: resourceWebServiceSecureRegenerateCreate,
		UpdateContext: resourceWebServiceSecureRegenerateUpdate,
		ReadContext:   resourceWebServiceSecureRegenerateRead,
		DeleteContext: resourceWebServiceSecureRegenerateDelete,

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
func resourceWebServiceSecureRegenerateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureRegenerateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureRegenerate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureRegenerateRead(ctx, d, meta)
	}
	return diags
}

func resourceWebServiceSecureRegenerateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureRegenerateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureRegenerate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureRegenerateRead(ctx, d, meta)
	}
	return diags
}
func resourceWebServiceSecureRegenerateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureRegenerateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureRegenerate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebServiceSecureRegenerateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureRegenerateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureRegenerate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointWebServiceSecureRegenerate(d *schema.ResourceData) edpt.WebServiceSecureRegenerate {
	var ret edpt.WebServiceSecureRegenerate
	ret.Inst.Country = d.Get("country").(string)
	ret.Inst.DomainName = d.Get("domain_name").(string)
	ret.Inst.State = d.Get("state").(string)
	return ret
}
