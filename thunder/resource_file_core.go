





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileCore() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_core`: file operation for core files\n\n__PLACEHOLDER__",
        CreateContext: resourceFileCoreCreate,
        UpdateContext: resourceFileCoreUpdate,
        ReadContext: resourceFileCoreRead,
        DeleteContext: resourceFileCoreDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'export': export; 'delete': delete;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "local core file name",
        },
        "slot": {
        Type: schema.TypeInt, Optional: true, Description: "specify slot id in chassis",
        },
        },
    }
}
func resourceFileCoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileCoreCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileCore(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileCoreRead(ctx, d, meta)
    }
    return diags
}

func resourceFileCoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileCoreUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileCore(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileCoreRead(ctx, d, meta)
    }
    return diags
}
func resourceFileCoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileCoreDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileCore(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileCoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileCoreRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileCore(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileCore(d *schema.ResourceData) edpt.FileCore {
    var ret edpt.FileCore
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.Slot = d.Get("slot").(int)
    return ret
}

