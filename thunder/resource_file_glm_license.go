





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileGlmLicense() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_glm_license`: glm license file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileGlmLicenseCreate,
        UpdateContext: resourceFileGlmLicenseUpdate,
        ReadContext: resourceFileGlmLicenseRead,
        DeleteContext: resourceFileGlmLicenseDelete,

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
func resourceFileGlmLicenseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmLicenseCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmLicense(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileGlmLicenseRead(ctx, d, meta)
    }
    return diags
}

func resourceFileGlmLicenseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmLicenseUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmLicense(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileGlmLicenseRead(ctx, d, meta)
    }
    return diags
}
func resourceFileGlmLicenseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmLicenseDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmLicense(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileGlmLicenseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileGlmLicenseRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileGlmLicense(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileGlmLicense(d *schema.ResourceData) edpt.FileGlmLicense {
    var ret edpt.FileGlmLicense
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

