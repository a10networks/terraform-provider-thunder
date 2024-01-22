package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewShowBackupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_show_backup_oper`: Operational Status for the object show-backup\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewShowBackupOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_show_1": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"backup_show_2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"backup_show_3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemViewShowBackupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewShowBackupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewShowBackupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewShowBackupOperOper := setObjectSystemViewShowBackupOperOper(res)
		d.Set("oper", SystemViewShowBackupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewShowBackupOperOper(ret edpt.DataSystemViewShowBackupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"backup_show_1": ret.DtSystemViewShowBackupOper.Oper.BackupShow1,
			"backup_show_2": ret.DtSystemViewShowBackupOper.Oper.BackupShow2,
			"backup_show_3": ret.DtSystemViewShowBackupOper.Oper.BackupShow3,
		},
	}
}

func getObjectSystemViewShowBackupOperOper(d []interface{}) edpt.SystemViewShowBackupOperOper {

	count1 := len(d)
	var ret edpt.SystemViewShowBackupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BackupShow1 = in["backup_show_1"].(string)
		ret.BackupShow2 = in["backup_show_2"].(string)
		ret.BackupShow3 = in["backup_show_3"].(string)
	}
	return ret
}

func dataToEndpointSystemViewShowBackupOper(d *schema.ResourceData) edpt.SystemViewShowBackupOper {
	var ret edpt.SystemViewShowBackupOper

	ret.Oper = getObjectSystemViewShowBackupOperOper(d.Get("oper").([]interface{}))
	return ret
}
