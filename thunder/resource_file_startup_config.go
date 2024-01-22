





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileStartupConfig() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_startup_config`: Contents of Startup Configuration\n\n__PLACEHOLDER__",
        CreateContext: resourceFileStartupConfigCreate,
        UpdateContext: resourceFileStartupConfigUpdate,
        ReadContext: resourceFileStartupConfigRead,
        DeleteContext: resourceFileStartupConfigDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'import': import;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "startup-config local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileStartupConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileStartupConfigCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileStartupConfig(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileStartupConfigRead(ctx, d, meta)
    }
    return diags
}

func resourceFileStartupConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileStartupConfigUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileStartupConfig(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileStartupConfigRead(ctx, d, meta)
    }
    return diags
}
func resourceFileStartupConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileStartupConfigDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileStartupConfig(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileStartupConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileStartupConfigRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileStartupConfig(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileStartupConfig(d *schema.ResourceData) edpt.FileStartupConfig {
    var ret edpt.FileStartupConfig
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

