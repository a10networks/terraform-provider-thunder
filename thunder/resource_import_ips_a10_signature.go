





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceImportIpsA10Signature() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_import_ips_a10_signature`: Update IPS A10 signatures\n\n__PLACEHOLDER__",
        CreateContext: resourceImportIpsA10SignatureCreate,
        UpdateContext: resourceImportIpsA10SignatureUpdate,
        ReadContext: resourceImportIpsA10SignatureRead,
        DeleteContext: resourceImportIpsA10SignatureDelete,

        Schema: map[string]*schema.Schema{
        "a10_signature_filename": {
        Type: schema.TypeString, Optional: true, Description: "Specify A10 signature file",
        },
        "password": {
        Type: schema.TypeString, Optional: true, Description: "password for the remote site",
        },
        "remote_file": {
        Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
        },
        "use_mgmt_port": {
        Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
        },
        },
    }
}
func resourceImportIpsA10SignatureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsA10SignatureCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsA10Signature(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceImportIpsA10SignatureRead(ctx, d, meta)
    }
    return diags
}

func resourceImportIpsA10SignatureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsA10SignatureUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsA10Signature(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceImportIpsA10SignatureRead(ctx, d, meta)
    }
    return diags
}
func resourceImportIpsA10SignatureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsA10SignatureDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsA10Signature(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceImportIpsA10SignatureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsA10SignatureRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsA10Signature(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointImportIpsA10Signature(d *schema.ResourceData) edpt.ImportIpsA10Signature {
    var ret edpt.ImportIpsA10Signature
	ret.Inst.A10SignatureFilename = d.Get("a10_signature_filename").(string)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
    return ret
}

