package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSystemBackupLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_system_backup_local`: Backup system files\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSystemBackupLocalCreate,
		UpdateContext: resourceFileSystemBackupLocalUpdate,
		ReadContext:   resourceFileSystemBackupLocalRead,
		DeleteContext: resourceFileSystemBackupLocalDelete,

		Schema: map[string]*schema.Schema{
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
		},
	}
}
func resourceFileSystemBackupLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemBackupLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemBackupLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemBackupLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSystemBackupLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemBackupLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemBackupLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSystemBackupLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSystemBackupLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemBackupLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemBackupLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSystemBackupLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSystemBackupLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSystemBackupLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSystemBackupLocal(d *schema.ResourceData) edpt.FileSystemBackupLocal {
	var ret edpt.FileSystemBackupLocal
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	return ret
}
