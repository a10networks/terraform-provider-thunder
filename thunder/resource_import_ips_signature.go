





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceImportIpsSignature() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_import_ips_signature`: Add IPS custom signatures\n\n__PLACEHOLDER__",
        CreateContext: resourceImportIpsSignatureCreate,
        UpdateContext: resourceImportIpsSignatureUpdate,
        ReadContext: resourceImportIpsSignatureRead,
        DeleteContext: resourceImportIpsSignatureDelete,

        Schema: map[string]*schema.Schema{
        "password": {
        Type: schema.TypeString, Optional: true, Description: "password for the remote site",
        },
        "remote_file": {
        Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
        },
        "signature_filename": {
        Type: schema.TypeString, Optional: true, Description: "Specify signature file",
        },
        "use_mgmt_port": {
        Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
        },
        },
    }
}
func resourceImportIpsSignatureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsSignatureCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsSignature(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceImportIpsSignatureRead(ctx, d, meta)
    }
    return diags
}

func resourceImportIpsSignatureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsSignatureUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsSignature(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceImportIpsSignatureRead(ctx, d, meta)
    }
    return diags
}
func resourceImportIpsSignatureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsSignatureDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsSignature(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceImportIpsSignatureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsSignatureRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsSignature(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointImportIpsSignature(d *schema.ResourceData) edpt.ImportIpsSignature {
    var ret edpt.ImportIpsSignature
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.SignatureFilename = d.Get("signature_filename").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
    return ret
}

