





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileAuthSamlIdp() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_auth_saml_idp`: SAML metadata file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileAuthSamlIdpCreate,
        UpdateContext: resourceFileAuthSamlIdpUpdate,
        ReadContext: resourceFileAuthSamlIdpRead,
        DeleteContext: resourceFileAuthSamlIdpDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "Destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "SAML metadata local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "Full path of the uploaded file",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        "verify_xml_signature": {
        Type: schema.TypeInt, Optional: true, Default: 0, Description: "Verify metadata's XML signature",
        },
        },
    }
}
func resourceFileAuthSamlIdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthSamlIdpCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthSamlIdp(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAuthSamlIdpRead(ctx, d, meta)
    }
    return diags
}

func resourceFileAuthSamlIdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthSamlIdpUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthSamlIdp(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAuthSamlIdpRead(ctx, d, meta)
    }
    return diags
}
func resourceFileAuthSamlIdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthSamlIdpDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthSamlIdp(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileAuthSamlIdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthSamlIdpRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthSamlIdp(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileAuthSamlIdp(d *schema.ResourceData) edpt.FileAuthSamlIdp {
    var ret edpt.FileAuthSamlIdp
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
	ret.Inst.VerifyXmlSignature = d.Get("verify_xml_signature").(int)
    return ret
}

