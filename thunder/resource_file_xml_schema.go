





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileXmlSchema() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_xml_schema`: XML-Schema file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileXmlSchemaCreate,
        UpdateContext: resourceFileXmlSchemaUpdate,
        ReadContext: resourceFileXmlSchemaRead,
        DeleteContext: resourceFileXmlSchemaDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;",
        },
        "dst_file": {
        Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "XML-schema local file name",
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
func resourceFileXmlSchemaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileXmlSchemaCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileXmlSchema(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileXmlSchemaRead(ctx, d, meta)
    }
    return diags
}

func resourceFileXmlSchemaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileXmlSchemaUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileXmlSchema(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileXmlSchemaRead(ctx, d, meta)
    }
    return diags
}
func resourceFileXmlSchemaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileXmlSchemaDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileXmlSchema(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileXmlSchemaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileXmlSchemaRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileXmlSchema(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileXmlSchema(d *schema.ResourceData) edpt.FileXmlSchema {
    var ret edpt.FileXmlSchema
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

