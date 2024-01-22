package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSslCertificate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ssl_certificate`: ssl certificate file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSslCertificateCreate,
		UpdateContext: resourceSslCertificateUpdate,
		ReadContext:   resourceSslCertificateRead,
		DeleteContext: resourceSslCertificateDelete,

		Schema: map[string]*schema.Schema{
			"certificate_type": {
				Type: schema.TypeString, Optional: true, Description: "'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "ssl certificate local file name",
			},
			"pfx_password": {
				Type: schema.TypeString, Optional: true, Description: "The password for certificate file (pfx type only)",
			},
			"public_key": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}
func resourceSslCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslCertificateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslCertificate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSslCertificateRead(ctx, d, meta)
	}
	return diags
}

func resourceSslCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslCertificateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslCertificate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSslCertificateRead(ctx, d, meta)
	}
	return diags
}
func resourceSslCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslCertificateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslCertificate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSslCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslCertificateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslCertificate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSslCertificate(d *schema.ResourceData) edpt.SslCertificate {
	var ret edpt.SslCertificate
	ret.Inst.CertificateType = d.Get("certificate_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PfxPassword = d.Get("pfx_password").(string)
	ret.Inst.PublicKey = d.Get("public_key").(string)
	return ret
}
