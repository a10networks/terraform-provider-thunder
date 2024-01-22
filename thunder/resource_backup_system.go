package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackupSystem() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_backup_system`: Backup system files\n\n__PLACEHOLDER__",
		CreateContext: resourceBackupSystemCreate,
		UpdateContext: resourceBackupSystemUpdate,
		ReadContext:   resourceBackupSystemRead,
		DeleteContext: resourceBackupSystemDelete,

		Schema: map[string]*schema.Schema{
			"encrypt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Encrypt the backup file",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"store_name": {
				Type: schema.TypeString, Optional: true, Description: "Save backup store information",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceBackupSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupSystemCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupSystem(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupSystemRead(ctx, d, meta)
	}
	return diags
}

func resourceBackupSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupSystemUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupSystem(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupSystemRead(ctx, d, meta)
	}
	return diags
}
func resourceBackupSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupSystemDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupSystem(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBackupSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupSystemRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupSystem(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointBackupSystem(d *schema.ResourceData) edpt.BackupSystem {
	var ret edpt.BackupSystem
	ret.Inst.Encrypt = d.Get("encrypt").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.StoreName = d.Get("store_name").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
