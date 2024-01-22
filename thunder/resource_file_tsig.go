





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileTsig() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_tsig`: Transaction SIGnatures\n\n__PLACEHOLDER__",
        CreateContext: resourceFileTsigCreate,
        UpdateContext: resourceFileTsigUpdate,
        ReadContext: resourceFileTsigRead,
        DeleteContext: resourceFileTsigDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'import': import; 'export': export; 'check': check; 'delete': delete; 'replace': replace;",
        },
        "device": {
        Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "TSIG file name",
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
func resourceFileTsigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTsigCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTsig(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileTsigRead(ctx, d, meta)
    }
    return diags
}

func resourceFileTsigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTsigUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTsig(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileTsigRead(ctx, d, meta)
    }
    return diags
}
func resourceFileTsigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTsigDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTsig(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileTsigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTsigRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTsig(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileTsig(d *schema.ResourceData) edpt.FileTsig {
    var ret edpt.FileTsig
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

