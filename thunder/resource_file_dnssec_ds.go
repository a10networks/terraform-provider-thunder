





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileDnssecDs() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_dnssec_ds`: dnssec ds file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileDnssecDsCreate,
        UpdateContext: resourceFileDnssecDsUpdate,
        ReadContext: resourceFileDnssecDsRead,
        DeleteContext: resourceFileDnssecDsDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'import': import; 'export': export;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "dnssec ds local file name",
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
func resourceFileDnssecDsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDsCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDs(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDnssecDsRead(ctx, d, meta)
    }
    return diags
}

func resourceFileDnssecDsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDsUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDs(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileDnssecDsRead(ctx, d, meta)
    }
    return diags
}
func resourceFileDnssecDsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDsDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDs(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileDnssecDsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileDnssecDsRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileDnssecDs(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileDnssecDs(d *schema.ResourceData) edpt.FileDnssecDs {
    var ret edpt.FileDnssecDs
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
    return ret
}

