package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackupLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_backup_log`: Backup log files\n\n__PLACEHOLDER__",
		CreateContext: resourceBackupLogCreate,
		UpdateContext: resourceBackupLogUpdate,
		ReadContext:   resourceBackupLogRead,
		DeleteContext: resourceBackupLogDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "all log",
			},
			"core": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup core files only",
			},
			"date": {
				Type: schema.TypeInt, Optional: true, Description: "specify number of days",
			},
			"day": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most recent day",
			},
			"expedite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Expedite the Backup",
			},
			"messages": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup messages file only",
			},
			"month": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most recent month",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify backup period",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"showtech": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup showtech files only",
			},
			"stats_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup web statistical data",
			},
			"store_name": {
				Type: schema.TypeString, Optional: true, Description: "Save backup store information",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"week": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most recent week",
			},
		},
	}
}
func resourceBackupLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupLogRead(ctx, d, meta)
	}
	return diags
}

func resourceBackupLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupLogRead(ctx, d, meta)
	}
	return diags
}
func resourceBackupLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBackupLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointBackupLog(d *schema.ResourceData) edpt.BackupLog {
	var ret edpt.BackupLog
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Core = d.Get("core").(int)
	ret.Inst.Date = d.Get("date").(int)
	ret.Inst.Day = d.Get("day").(int)
	ret.Inst.Expedite = d.Get("expedite").(int)
	ret.Inst.Messages = d.Get("messages").(int)
	ret.Inst.Month = d.Get("month").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.Showtech = d.Get("showtech").(int)
	ret.Inst.StatsData = d.Get("stats_data").(int)
	ret.Inst.StoreName = d.Get("store_name").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.Week = d.Get("week").(int)
	return ret
}
