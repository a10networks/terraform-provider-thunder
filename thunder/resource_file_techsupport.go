





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileTechsupport() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_techsupport`: Generate the current showtech file\n\n__PLACEHOLDER__",
        CreateContext: resourceFileTechsupportCreate,
        UpdateContext: resourceFileTechsupportUpdate,
        ReadContext: resourceFileTechsupportRead,
        DeleteContext: resourceFileTechsupportDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'check': check; 'delete': delete; 'export': export;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "techsupport local file name",
        },
        "slot": {
        Type: schema.TypeInt, Optional: true, Description: "specify slot id in chassis",
        },
        "status": {
        Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
        Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
},
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileTechsupportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTechsupportCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTechsupport(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileTechsupportRead(ctx, d, meta)
    }
    return diags
}

func resourceFileTechsupportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTechsupportUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTechsupport(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileTechsupportRead(ctx, d, meta)
    }
    return diags
}
func resourceFileTechsupportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTechsupportDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTechsupport(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileTechsupportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTechsupportRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTechsupport(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func getObjectFileTechsupportStatus332(d []interface{}) edpt.FileTechsupportStatus332 {
    
     var ret edpt.FileTechsupportStatus332
return ret
}



func dataToEndpointFileTechsupport(d *schema.ResourceData) edpt.FileTechsupport {
    var ret edpt.FileTechsupport
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.Slot = d.Get("slot").(int)
	ret.Inst.Status = getObjectFileTechsupportStatus332(d.Get("status").([]interface{}))
	//omit uuid
    return ret
}

