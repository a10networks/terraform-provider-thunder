





package thunder

import (
    "context"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"

)

func resourceFileLogBackup() *schema.Resource {
    return &schema.Resource{
    	Description: "`thunder_file_log_backup`: Backup system files\n\n__PLACEHOLDER__",
        CreateContext: resourceFileLogBackupCreate,
        UpdateContext: resourceFileLogBackupUpdate,
        ReadContext: resourceFileLogBackupRead,
        DeleteContext: resourceFileLogBackupDelete,

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
func resourceFileLogBackupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLogBackupCreate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLogBackup(d)
        d.SetId(obj.GetId())
        err := obj.Post(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileLogBackupRead(ctx, d, meta)
    }
    return diags
}

func resourceFileLogBackupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLogBackupUpdate()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLogBackup(d)
        err := obj.Put(client.Token, client.Host, logger)
        if err != nil {
            return diag.FromErr(err)
        }
        return resourceFileLogBackupRead(ctx, d, meta)
    }
    return diags
}
func resourceFileLogBackupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLogBackupDelete()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLogBackup(d)
        err := obj.Delete(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}



func resourceFileLogBackupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    client := meta.(Thunder)
    logger := client.log
    logger.Println("resourceFileLogBackupRead()")
    var diags diag.Diagnostics
    if client.Host != "" {
        obj := dataToEndpointFileLogBackup(d)
        err := obj.Get(client.Token, client.Host, d.Id(), logger)
        if err != nil {
            return diag.FromErr(err)
        }
    }
    return diags
}




func dataToEndpointFileLogBackup(d *schema.ResourceData) edpt.FileLogBackup {
    var ret edpt.FileLogBackup
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

