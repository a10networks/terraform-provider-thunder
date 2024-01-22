





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileSslCertPoc() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_ssl_cert_poc`: ssl certificate file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileSslCertPocCreate,
        UpdateContext: resourceFileSslCertPocUpdate,
        ReadContext: resourceFileSslCertPocRead,
        DeleteContext: resourceFileSslCertPocDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "certificate_type": {
        Type: schema.TypeString, Optional: true, Description: "'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "ssl certificate local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "public-key",
        },
        "pfx_password": {
        Type: schema.TypeString, Optional: true, Description: "The password for certificate file (pfx type only)",
        },
        },
    }
}
func resourceFileSslCertPocCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSslCertPocCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSslCertPoc(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSslCertPocRead(ctx, d, meta)
    }
    return diags
}

func resourceFileSslCertPocUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSslCertPocUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSslCertPoc(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSslCertPocRead(ctx, d, meta)
    }
    return diags
}
func resourceFileSslCertPocDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSslCertPocDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSslCertPoc(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileSslCertPocRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSslCertPocRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSslCertPoc(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileSslCertPoc(d *schema.ResourceData) edpt.FileSslCertPoc {
    var ret edpt.FileSslCertPoc
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.CertificateType = d.Get("certificate_type").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.PfxPassword = d.Get("pfx_password").(string)
    return ret
}

