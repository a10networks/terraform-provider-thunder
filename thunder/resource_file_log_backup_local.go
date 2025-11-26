package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileLogBackupLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_log_backup_local`: Backup system files\n\n__PLACEHOLDER__",
		CreateContext: resourceFileLogBackupLocalCreate,
		UpdateContext: resourceFileLogBackupLocalUpdate,
		ReadContext:   resourceFileLogBackupLocalRead,
		DeleteContext: resourceFileLogBackupLocalDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "all log",
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
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"month": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most recent month",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify backup period",
			},
			"stats_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Backup web statistical data",
			},
			"week": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Most recent week",
			},
		},
	}
}
func resourceFileLogBackupLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLogBackupLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLogBackupLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileLogBackupLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileLogBackupLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLogBackupLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLogBackupLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileLogBackupLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileLogBackupLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLogBackupLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLogBackupLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileLogBackupLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileLogBackupLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileLogBackupLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileLogBackupLocal(d *schema.ResourceData) edpt.FileLogBackupLocal {
	var ret edpt.FileLogBackupLocal
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.Date = d.Get("date").(int)
	ret.Inst.Day = d.Get("day").(int)
	ret.Inst.Expedite = d.Get("expedite").(int)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.Month = d.Get("month").(int)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.StatsData = d.Get("stats_data").(int)
	ret.Inst.Week = d.Get("week").(int)
	return ret
}
