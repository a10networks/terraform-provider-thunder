package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSecureAttributes() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_secure_attributes`: GSLB Secure Settings per partition\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSecureAttributesCreate,
		UpdateContext: resourceGslbSecureAttributesUpdate,
		ReadContext:   resourceGslbSecureAttributesRead,
		DeleteContext: resourceGslbSecureAttributesDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "use-global-behavior", Description: "'enable': Enable Secure; 'disable': Disable Secure; 'enable-fallback': Fall back to non-secure if fail; 'use-global-behavior': Follow global configuration under gslb protocol (default);",
			},
			"gslb_cert": {
				Type: schema.TypeString, Optional: true, Description: "Certificate for Secure GSLB (Certificate name)",
			},
			"gslb_key": {
				Type: schema.TypeString, Optional: true, Description: "Private Key for secure gslb signing (Key name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbSecureAttributesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSecureAttributesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSecureAttributes(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSecureAttributesRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSecureAttributesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSecureAttributesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSecureAttributes(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSecureAttributesRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSecureAttributesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSecureAttributesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSecureAttributes(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSecureAttributesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSecureAttributesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSecureAttributes(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGslbSecureAttributes(d *schema.ResourceData) edpt.GslbSecureAttributes {
	var ret edpt.GslbSecureAttributes
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.GslbCert = d.Get("gslb_cert").(string)
	ret.Inst.GslbKey = d.Get("gslb_key").(string)
	//omit uuid
	return ret
}
