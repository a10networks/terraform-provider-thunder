





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileRpz() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_rpz`: Response Policy Zone\n\n__PLACEHOLDER__",
        CreateContext: resourceFileRpzCreate,
        UpdateContext: resourceFileRpzUpdate,
        ReadContext: resourceFileRpzRead,
        DeleteContext: resourceFileRpzDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'import': import; 'export': export; 'check': check; 'delete': delete; 'replace': replace;",
        },
        "device": {
        Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "RPZ file name",
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
func resourceFileRpzCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileRpzCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileRpz(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileRpzRead(ctx, d, meta)
    }
    return diags
}

func resourceFileRpzUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileRpzUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileRpz(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileRpzRead(ctx, d, meta)
    }
    return diags
}
func resourceFileRpzDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileRpzDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileRpz(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileRpzRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileRpzRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileRpz(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileRpz(d *schema.ResourceData) edpt.FileRpz {
    var ret edpt.FileRpz
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

