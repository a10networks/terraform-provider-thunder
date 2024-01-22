





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileAwsAccesskey() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_aws_accesskey`: The aws accesskey for admin user\n\n__PLACEHOLDER__",
        CreateContext: resourceFileAwsAccesskeyCreate,
        UpdateContext: resourceFileAwsAccesskeyUpdate,
        ReadContext: resourceFileAwsAccesskeyRead,
        DeleteContext: resourceFileAwsAccesskeyDelete,

        Schema: map[string]*schema.Schema{
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        "user": {
        Type: schema.TypeString, Optional: true, Description: "user name of the pub key",
        },
        "uuid": {
        Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
        },
        },
    }
}
func resourceFileAwsAccesskeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAwsAccesskeyCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAwsAccesskey(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAwsAccesskeyRead(ctx, d, meta)
    }
    return diags
}

func resourceFileAwsAccesskeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAwsAccesskeyUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAwsAccesskey(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileAwsAccesskeyRead(ctx, d, meta)
    }
    return diags
}
func resourceFileAwsAccesskeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAwsAccesskeyDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAwsAccesskey(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileAwsAccesskeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileAwsAccesskeyRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileAwsAccesskey(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileAwsAccesskey(d *schema.ResourceData) edpt.FileAwsAccesskey {
    var ret edpt.FileAwsAccesskey
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
    return ret
}

