





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileAxdebugCapture() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_axdebug_capture`: axdebug file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileAxdebugCaptureCreate,
        UpdateContext: resourceFileAxdebugCaptureUpdate,
        ReadContext: resourceFileAxdebugCaptureRead,
        DeleteContext: resourceFileAxdebugCaptureDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'delete': delete;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "axdebug local capture file name",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileAxdebugCaptureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugCaptureCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugCapture(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAxdebugCaptureRead(ctx, d, meta)
    }
    return diags
}

func resourceFileAxdebugCaptureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugCaptureUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugCapture(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAxdebugCaptureRead(ctx, d, meta)
    }
    return diags
}
func resourceFileAxdebugCaptureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugCaptureDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugCapture(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileAxdebugCaptureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAxdebugCaptureRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAxdebugCapture(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileAxdebugCapture(d *schema.ResourceData) edpt.FileAxdebugCapture {
    var ret edpt.FileAxdebugCapture
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	//omit uuid
    return ret
}

