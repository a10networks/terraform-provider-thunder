





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileGlmCert() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_glm_cert`: GLM certificate\n\n__PLACEHOLDER__",
        CreateContext: resourceFileGlmCertCreate,
        UpdateContext: resourceFileGlmCertUpdate,
        ReadContext: resourceFileGlmCertRead,
        DeleteContext: resourceFileGlmCertDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'import': import;",
        },
        "device": {
        Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "glm license local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        },
    }
}
func resourceFileGlmCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmCertCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmCert(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileGlmCertRead(ctx, d, meta)
    }
    return diags
}

func resourceFileGlmCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmCertUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmCert(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileGlmCertRead(ctx, d, meta)
    }
    return diags
}
func resourceFileGlmCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmCertDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmCert(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileGlmCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmCertRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmCert(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileGlmCert(d *schema.ResourceData) edpt.FileGlmCert {
    var ret edpt.FileGlmCert
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

