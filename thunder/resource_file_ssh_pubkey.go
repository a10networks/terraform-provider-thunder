





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileSshPubkey() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_ssh_pubkey`: The ssh pubkey for admin user\n\n__PLACEHOLDER__",
        CreateContext: resourceFileSshPubkeyCreate,
        UpdateContext: resourceFileSshPubkeyUpdate,
        ReadContext: resourceFileSshPubkeyRead,
        DeleteContext: resourceFileSshPubkeyDelete,

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
func resourceFileSshPubkeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSshPubkeyCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSshPubkey(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSshPubkeyRead(ctx, d, meta)
    }
    return diags
}

func resourceFileSshPubkeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSshPubkeyUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSshPubkey(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSshPubkeyRead(ctx, d, meta)
    }
    return diags
}
func resourceFileSshPubkeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSshPubkeyDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSshPubkey(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileSshPubkeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSshPubkeyRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSshPubkey(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileSshPubkey(d *schema.ResourceData) edpt.FileSshPubkey {
    var ret edpt.FileSshPubkey
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.User = d.Get("user").(string)
	//omit uuid
    return ret
}

