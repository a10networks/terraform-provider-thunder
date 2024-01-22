package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileCaCertLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_ca_cert_local`: ssl certificate file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileCaCertLocalCreate,
		UpdateContext: resourceFileCaCertLocalUpdate,
		ReadContext:   resourceFileCaCertLocalRead,
		DeleteContext: resourceFileCaCertLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'replace': replace;",
			},
			"certificate_type": {
				Type: schema.TypeString, Optional: true, Description: "'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "ssl ca certificate local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"pfx_password": {
				Type: schema.TypeString, Optional: true, Description: "The password for certificate file (pfx type only)",
			},
			"pfx_password_export": {
				Type: schema.TypeString, Optional: true, Description: "The password for exported certificate file (pfx type only)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileCaCertLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCaCertLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCaCertLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCaCertLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileCaCertLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCaCertLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCaCertLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCaCertLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileCaCertLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCaCertLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCaCertLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileCaCertLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCaCertLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCaCertLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileCaCertLocal(d *schema.ResourceData) edpt.FileCaCertLocal {
	var ret edpt.FileCaCertLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.CertificateType = d.Get("certificate_type").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.PfxPassword = d.Get("pfx_password").(string)
	ret.Inst.PfxPasswordExport = d.Get("pfx_password_export").(string)
	//omit uuid
	return ret
}
