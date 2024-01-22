





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceImportIpsProfile() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_import_ips_profile`: Modify/Enable/Disable IPS signatures\n\n__PLACEHOLDER__",
        CreateContext: resourceImportIpsProfileCreate,
        UpdateContext: resourceImportIpsProfileUpdate,
        ReadContext: resourceImportIpsProfileRead,
        DeleteContext: resourceImportIpsProfileDelete,

        Schema: map[string]*schema.Schema{
        "overwrite": {
        Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
        },
        "password": {
        Type: schema.TypeString, Optional: true, Description: "password for the remote site",
        },
        "profile_filename": {
        Type: schema.TypeString, Optional: true, Description: "Specify profile",
        },
        "remote_file": {
        Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
        },
        "use_mgmt_port": {
        Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
        },
        },
    }
}
func resourceImportIpsProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsProfileCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsProfile(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceImportIpsProfileRead(ctx, d, meta)
    }
    return diags
}

func resourceImportIpsProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsProfileUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsProfile(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceImportIpsProfileRead(ctx, d, meta)
    }
    return diags
}
func resourceImportIpsProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsProfileDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsProfile(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceImportIpsProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceImportIpsProfileRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointImportIpsProfile(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointImportIpsProfile(d *schema.ResourceData) edpt.ImportIpsProfile {
    var ret edpt.ImportIpsProfile
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.ProfileFilename = d.Get("profile_filename").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
    return ret
}

