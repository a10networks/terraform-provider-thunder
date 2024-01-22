





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileNgWafModule() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_ng_waf_module`: WAF module\n\n__PLACEHOLDER__",
        CreateContext: resourceFileNgWafModuleCreate,
        UpdateContext: resourceFileNgWafModuleUpdate,
        ReadContext: resourceFileNgWafModuleRead,
        DeleteContext: resourceFileNgWafModuleDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'import': import; 'replace': replace;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "NG WAF module local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        },
    }
}
func resourceFileNgWafModuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileNgWafModuleCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileNgWafModule(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileNgWafModuleRead(ctx, d, meta)
    }
    return diags
}

func resourceFileNgWafModuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileNgWafModuleUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileNgWafModule(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileNgWafModuleRead(ctx, d, meta)
    }
    return diags
}
func resourceFileNgWafModuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileNgWafModuleDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileNgWafModule(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileNgWafModuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileNgWafModuleRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileNgWafModule(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileNgWafModule(d *schema.ResourceData) edpt.FileNgWafModule {
    var ret edpt.FileNgWafModule
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

