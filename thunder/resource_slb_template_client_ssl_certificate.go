package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateClientSslCertificate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_client_ssl_certificate`: Server Certificate\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateClientSslCertificateCreate,
		UpdateContext: resourceSlbTemplateClientSslCertificateUpdate,
		ReadContext:   resourceSlbTemplateClientSslCertificateRead,
		DeleteContext: resourceSlbTemplateClientSslCertificateDelete,

		Schema: map[string]*schema.Schema{
			"cert": {
				Type: schema.TypeString, Required: true, Description: "Certificate Name",
			},
			"chain_cert": {
				Type: schema.TypeString, Optional: true, Description: "Chain Certificate (Chain Certificate Name)",
			},
			"key": {
				Type: schema.TypeString, Optional: true, Description: "Server Private Key (Key Name)",
			},
			"passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Certificate and Key Partition Shared",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateClientSslCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslCertificateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSslCertificate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateClientSslCertificateRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateClientSslCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslCertificateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSslCertificate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateClientSslCertificateRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateClientSslCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslCertificateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSslCertificate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateClientSslCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSslCertificateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSslCertificate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateClientSslCertificate(d *schema.ResourceData) edpt.SlbTemplateClientSslCertificate {
	var ret edpt.SlbTemplateClientSslCertificate
	ret.Inst.Cert = d.Get("cert").(string)
	ret.Inst.ChainCert = d.Get("chain_cert").(string)
	ret.Inst.Key = d.Get("key").(string)
	//omit key_encrypted
	ret.Inst.Passphrase = d.Get("passphrase").(string)
	ret.Inst.Shared = d.Get("shared").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
