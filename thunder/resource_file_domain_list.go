





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileDomainList() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_domain_list`: domain list file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileDomainListCreate,
        UpdateContext: resourceFileDomainListUpdate,
        ReadContext: resourceFileDomainListRead,
        DeleteContext: resourceFileDomainListDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "domain list local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        "size": {
        Type: schema.TypeInt, Optional: true, Description: "domain list file size in byte",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileDomainListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDomainListCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDomainList(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDomainListRead(ctx, d, meta)
    }
    return diags
}

func resourceFileDomainListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDomainListUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDomainList(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDomainListRead(ctx, d, meta)
    }
    return diags
}
func resourceFileDomainListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDomainListDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDomainList(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileDomainListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDomainListRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDomainList(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileDomainList(d *schema.ResourceData) edpt.FileDomainList {
    var ret edpt.FileDomainList
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.Size = d.Get("size").(int)
	//omit uuid
    return ret
}

