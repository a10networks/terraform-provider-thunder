





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileAuthJwks() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_auth_jwks`: JWK file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileAuthJwksCreate,
        UpdateContext: resourceFileAuthJwksUpdate,
        ReadContext: resourceFileAuthJwksRead,
        DeleteContext: resourceFileAuthJwksDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "Destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "JWK local file name",
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
func resourceFileAuthJwksCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthJwksCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthJwks(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAuthJwksRead(ctx, d, meta)
    }
    return diags
}

func resourceFileAuthJwksUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthJwksUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthJwks(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAuthJwksRead(ctx, d, meta)
    }
    return diags
}
func resourceFileAuthJwksDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthJwksDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthJwks(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileAuthJwksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAuthJwksRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAuthJwks(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileAuthJwks(d *schema.ResourceData) edpt.FileAuthJwks {
    var ret edpt.FileAuthJwks
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

