





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileHealthExternal() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_health_external`: Address the External Script Program\n\n__PLACEHOLDER__",
        CreateContext: resourceFileHealthExternalCreate,
        UpdateContext: resourceFileHealthExternalUpdate,
        ReadContext: resourceFileHealthExternalRead,
        DeleteContext: resourceFileHealthExternalDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "description": {
        Type: schema.TypeString, Optional: true, Description: "Describe the Program Function briefly",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "Specify the Program Name",
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
func resourceFileHealthExternalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthExternalCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthExternal(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileHealthExternalRead(ctx, d, meta)
    }
    return diags
}

func resourceFileHealthExternalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthExternalUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthExternal(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileHealthExternalRead(ctx, d, meta)
    }
    return diags
}
func resourceFileHealthExternalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthExternalDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthExternal(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileHealthExternalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthExternalRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthExternal(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileHealthExternal(d *schema.ResourceData) edpt.FileHealthExternal {
    var ret edpt.FileHealthExternal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

