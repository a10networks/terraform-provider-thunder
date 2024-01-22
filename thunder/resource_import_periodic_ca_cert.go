package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicCaCert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_ca_cert`: SSL CA Cert File(enter bulk when import an archive file)\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicCaCertCreate,
		UpdateContext: resourceImportPeriodicCaCertUpdate,
		ReadContext:   resourceImportPeriodicCaCertRead,
		DeleteContext: resourceImportPeriodicCaCertDelete,

		Schema: map[string]*schema.Schema{
			"ca_cert": {
				Type: schema.TypeString, Required: true, Description: "SSL CA Cert File(enter bulk when import an archive file)",
			},
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
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceImportPeriodicCaCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicCaCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicCaCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicCaCertRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicCaCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicCaCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicCaCert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicCaCertRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicCaCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicCaCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicCaCert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicCaCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicCaCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicCaCert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicCaCert(d *schema.ResourceData) edpt.ImportPeriodicCaCert {
	var ret edpt.ImportPeriodicCaCert
	ret.Inst.CaCert = d.Get("ca_cert").(string)
	ret.Inst.CertificateType = d.Get("certificate_type").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.PfxPassword = d.Get("pfx_password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
