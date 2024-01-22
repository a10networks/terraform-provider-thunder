





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileSyslog() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_syslog`: syslog file information and management commands\n\n__PLACEHOLDER__",
        CreateContext: resourceFileSyslogCreate,
        UpdateContext: resourceFileSyslogUpdate,
        ReadContext: resourceFileSyslogRead,
        DeleteContext: resourceFileSyslogDelete,

        Schema: map[string]*schema.Schema{
        "action": {
        Type: schema.TypeString, Optional: true, Description: "'export': export;",
        },
        "file": {
        Type: schema.TypeString, Optional: true, Description: "syslog local file name",
        },
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        },
    }
}
func resourceFileSyslogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSyslogCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSyslog(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSyslogRead(ctx, d, meta)
    }
    return diags
}

func resourceFileSyslogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSyslogUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSyslog(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSyslogRead(ctx, d, meta)
    }
    return diags
}
func resourceFileSyslogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSyslogDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSyslog(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileSyslogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSyslogRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSyslog(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileSyslog(d *schema.ResourceData) edpt.FileSyslog {
    var ret edpt.FileSyslog
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

