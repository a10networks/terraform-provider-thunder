





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileSystemBackup() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_system_backup`: Backup system files\n\n__PLACEHOLDER__",
        CreateContext: resourceFileSystemBackupCreate,
        UpdateContext: resourceFileSystemBackupUpdate,
        ReadContext: resourceFileSystemBackupRead,
        DeleteContext: resourceFileSystemBackupDelete,

        Schema: map[string]*schema.Schema{
        "file_handle": {
        Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
        },
        },
    }
}
func resourceFileSystemBackupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSystemBackupCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSystemBackup(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSystemBackupRead(ctx, d, meta)
    }
    return diags
}

func resourceFileSystemBackupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSystemBackupUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSystemBackup(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileSystemBackupRead(ctx, d, meta)
    }
    return diags
}
func resourceFileSystemBackupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSystemBackupDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSystemBackup(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileSystemBackupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileSystemBackupRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileSystemBackup(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileSystemBackup(d *schema.ResourceData) edpt.FileSystemBackup {
    var ret edpt.FileSystemBackup
	ret.Inst.FileHandle = d.Get("file_handle").(string)
    return ret
}

