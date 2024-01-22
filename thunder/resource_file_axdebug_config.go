





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileAxdebugConfig() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_axdebug_config`: axdebug file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileAxdebugConfigCreate,
        UpdateContext: resourceFileAxdebugConfigUpdate,
        ReadContext: resourceFileAxdebugConfigRead,
        DeleteContext: resourceFileAxdebugConfigDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'export': export; 'delete': delete;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "axdebug local config file name",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileAxdebugConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugConfigCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugConfig(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAxdebugConfigRead(ctx, d, meta)
    }
    return diags
}

func resourceFileAxdebugConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugConfigUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugConfig(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAxdebugConfigRead(ctx, d, meta)
    }
    return diags
}
func resourceFileAxdebugConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugConfigDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugConfig(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileAxdebugConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugConfigRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugConfig(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileAxdebugConfig(d *schema.ResourceData) edpt.FileAxdebugConfig {
    var ret edpt.FileAxdebugConfig
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	//omit uuid
    return ret
}

