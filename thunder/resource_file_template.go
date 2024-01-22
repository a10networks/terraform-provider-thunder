





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileTemplate() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_template`: upgrade acos wizard template\n\n__PLACEHOLDER__",
        CreateContext: resourceFileTemplateCreate,
        UpdateContext: resourceFileTemplateUpdate,
        ReadContext: resourceFileTemplateRead,
        DeleteContext: resourceFileTemplateDelete,

        Schema: map[string]*schema.Schema{
        "act_type": {
        Type: schema.TypeString, Optional: true, Description: "Specify ACT package type (i.e. slug/destination name)",
        },
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "app": {
        Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
        Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "",
        },
        "act_name": {
        Type: schema.TypeString, Optional: true, Description: "",
        },
        "version": {
        Type: schema.TypeString, Optional: true, Description: "",
        },
        },
},
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "aflex local file name",
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
func resourceFileTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTemplateCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTemplate(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileTemplateRead(ctx, d, meta)
    }
    return diags
}

func resourceFileTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTemplateUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTemplate(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileTemplateRead(ctx, d, meta)
    }
    return diags
}
func resourceFileTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTemplateDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTemplate(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileTemplateRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileTemplate(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func getObjectFileTemplateApp333(d []interface{}) edpt.FileTemplateApp333 {
    
        count1 := len(d)
            var ret edpt.FileTemplateApp333
        if count1 > 0 {
            in := d[0].(map[string]interface{})
            ret.Action = in["action"].(string)
            ret.ActName = in["act_name"].(string)
            ret.Version = in["version"].(string)
            }
return ret
}



func dataToEndpointFileTemplate(d *schema.ResourceData) edpt.FileTemplate {
    var ret edpt.FileTemplate
	ret.Inst.ActType = d.Get("act_type").(string)
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.App = getObjectFileTemplateApp333(d.Get("app").([]interface{}))
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

