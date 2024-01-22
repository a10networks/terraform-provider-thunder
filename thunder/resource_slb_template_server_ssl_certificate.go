package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServerSslCertificate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_server_ssl_certificate`: Client Certificate\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateServerSslCertificateCreate,
		UpdateContext: resourceSlbTemplateServerSslCertificateUpdate,
		ReadContext:   resourceSlbTemplateServerSslCertificateRead,
		DeleteContext: resourceSlbTemplateServerSslCertificateDelete,

		Schema: map[string]*schema.Schema{
			"cert": {
				Type: schema.TypeString, Optional: true, Description: "Certificate Name",
			},
			"key": {
				Type: schema.TypeString, Optional: true, Description: "Client private-key (Key Name)",
			},
			"passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client Certificate and Key Partition Shared",
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
func resourceSlbTemplateServerSslCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslCertificateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSslCertificate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerSslCertificateRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateServerSslCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslCertificateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSslCertificate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerSslCertificateRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateServerSslCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslCertificateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSslCertificate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateServerSslCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSslCertificateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSslCertificate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateServerSslCertificate(d *schema.ResourceData) edpt.SlbTemplateServerSslCertificate {
	var ret edpt.SlbTemplateServerSslCertificate
	ret.Inst.Cert = d.Get("cert").(string)
	//omit encrypted
	ret.Inst.Key = d.Get("key").(string)
	ret.Inst.Passphrase = d.Get("passphrase").(string)
	ret.Inst.Shared = d.Get("shared").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
