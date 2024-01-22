





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileAuthPortalImage() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_auth_portal_image`: Authentication portal image file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileAuthPortalImageCreate,
        UpdateContext: resourceFileAuthPortalImageUpdate,
        ReadContext: resourceFileAuthPortalImageRead,
        DeleteContext: resourceFileAuthPortalImageDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "Destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "Authentication portal image local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "Full path of the uploaded file",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileAuthPortalImageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthPortalImageCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthPortalImage(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAuthPortalImageRead(ctx, d, meta)
    }
    return diags
}

func resourceFileAuthPortalImageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthPortalImageUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthPortalImage(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAuthPortalImageRead(ctx, d, meta)
    }
    return diags
}
func resourceFileAuthPortalImageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthPortalImageDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthPortalImage(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileAuthPortalImageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthPortalImageRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthPortalImage(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileAuthPortalImage(d *schema.ResourceData) edpt.FileAuthPortalImage {
    var ret edpt.FileAuthPortalImage
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

