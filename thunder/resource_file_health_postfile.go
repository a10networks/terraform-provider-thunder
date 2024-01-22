





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileHealthPostfile() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_health_postfile`: Address the HTTP Post data file\n\n__PLACEHOLDER__",
        CreateContext: resourceFileHealthPostfileCreate,
        UpdateContext: resourceFileHealthPostfileUpdate,
        ReadContext: resourceFileHealthPostfileRead,
        DeleteContext: resourceFileHealthPostfileDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "Specify the File Name",
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
func resourceFileHealthPostfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthPostfileCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthPostfile(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileHealthPostfileRead(ctx, d, meta)
    }
    return diags
}

func resourceFileHealthPostfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthPostfileUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthPostfile(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileHealthPostfileRead(ctx, d, meta)
    }
    return diags
}
func resourceFileHealthPostfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthPostfileDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthPostfile(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileHealthPostfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileHealthPostfileRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileHealthPostfile(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileHealthPostfile(d *schema.ResourceData) edpt.FileHealthPostfile {
    var ret edpt.FileHealthPostfile
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

