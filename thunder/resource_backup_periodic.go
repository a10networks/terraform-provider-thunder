package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackupPeriodic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_backup_periodic`: Configure backup periodically\n\n__PLACEHOLDER__",
		CreateContext: resourceBackupPeriodicCreate,
		UpdateContext: resourceBackupPeriodicUpdate,
		ReadContext:   resourceBackupPeriodicRead,
		DeleteContext: resourceBackupPeriodicDelete,

		Schema: map[string]*schema.Schema{
			"day": {
				Type: schema.TypeInt, Optional: true, Description: "Specify interval days",
			},
			"encrypt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Encrypt the backup file",
			},
			"fixed_nat": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup fixed-nat port mapping files",
			},
			"hour": {
				Type: schema.TypeInt, Optional: true, Description: "Specify interval hours",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup log files",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"store_name": {
				Type: schema.TypeString, Optional: true, Description: "profile name to store remote url",
			},
			"system": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup system files",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"week": {
				Type: schema.TypeInt, Optional: true, Description: "Specify interval weeks",
			},
		},
	}
}
func resourceBackupPeriodicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupPeriodicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupPeriodic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupPeriodicRead(ctx, d, meta)
	}
	return diags
}

func resourceBackupPeriodicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupPeriodicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupPeriodic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupPeriodicRead(ctx, d, meta)
	}
	return diags
}
func resourceBackupPeriodicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupPeriodicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupPeriodic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBackupPeriodicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupPeriodicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupPeriodic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointBackupPeriodic(d *schema.ResourceData) edpt.BackupPeriodic {
	var ret edpt.BackupPeriodic
	ret.Inst.Day = d.Get("day").(int)
	ret.Inst.Encrypt = d.Get("encrypt").(int)
	ret.Inst.FixedNat = d.Get("fixed_nat").(int)
	ret.Inst.Hour = d.Get("hour").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.StoreName = d.Get("store_name").(string)
	ret.Inst.System = d.Get("system").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	ret.Inst.Week = d.Get("week").(int)
	return ret
}
