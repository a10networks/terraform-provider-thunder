package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicSslCert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_ssl_cert`: SSL Cert File(enter bulk when import an archive file)\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicSslCertCreate,
		UpdateContext: resourceImportPeriodicSslCertUpdate,
		ReadContext:   resourceImportPeriodicSslCertRead,
		DeleteContext: resourceImportPeriodicSslCertDelete,

		Schema: map[string]*schema.Schema{
			"certificate_type": {
				Type: schema.TypeString, Optional: true, Description: "'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"pfx_password": {
				Type: schema.TypeString, Optional: true, Description: "The password for certificate file (pfx type only)",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"ssl_cert": {
				Type: schema.TypeString, Required: true, Description: "SSL Cert File(enter bulk when import an archive file)",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceImportPeriodicSslCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicSslCertRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicSslCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicSslCertRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicSslCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicSslCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicSslCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicSslCert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicSslCert(d *schema.ResourceData) edpt.ImportPeriodicSslCert {
	var ret edpt.ImportPeriodicSslCert
	ret.Inst.CertificateType = d.Get("certificate_type").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.PfxPassword = d.Get("pfx_password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.SslCert = d.Get("ssl_cert").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
