





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileDebugMonitor() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_debug_monitor`: debug monitor file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileDebugMonitorCreate,
        UpdateContext: resourceFileDebugMonitorUpdate,
        ReadContext: resourceFileDebugMonitorRead,
        DeleteContext: resourceFileDebugMonitorDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'export': export;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "debug monitor local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        },
    }
}
func resourceFileDebugMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDebugMonitorCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDebugMonitor(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDebugMonitorRead(ctx, d, meta)
    }
    return diags
}

func resourceFileDebugMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDebugMonitorUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDebugMonitor(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDebugMonitorRead(ctx, d, meta)
    }
    return diags
}
func resourceFileDebugMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDebugMonitorDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDebugMonitor(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileDebugMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDebugMonitorRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDebugMonitor(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileDebugMonitor(d *schema.ResourceData) edpt.FileDebugMonitor {
    var ret edpt.FileDebugMonitor
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

