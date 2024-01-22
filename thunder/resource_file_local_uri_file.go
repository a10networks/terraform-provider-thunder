





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileLocalUriFile() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_local_uri_file`: local uri file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileLocalUriFileCreate,
        UpdateContext: resourceFileLocalUriFileUpdate,
        ReadContext: resourceFileLocalUriFileRead,
        DeleteContext: resourceFileLocalUriFileDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "local uri file name",
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
func resourceFileLocalUriFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLocalUriFileCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLocalUriFile(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileLocalUriFileRead(ctx, d, meta)
    }
    return diags
}

func resourceFileLocalUriFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLocalUriFileUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLocalUriFile(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileLocalUriFileRead(ctx, d, meta)
    }
    return diags
}
func resourceFileLocalUriFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLocalUriFileDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLocalUriFile(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileLocalUriFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLocalUriFileRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLocalUriFile(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileLocalUriFile(d *schema.ResourceData) edpt.FileLocalUriFile {
    var ret edpt.FileLocalUriFile
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

