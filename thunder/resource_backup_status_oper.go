package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackupStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_backup_status_oper`: Operational Status for the object status\n\n__PLACEHOLDER__",
		ReadContext: resourceBackupStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"message": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"progress": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceBackupStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		BackupStatusOperOper := setObjectBackupStatusOperOper(res)
		d.Set("oper", BackupStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectBackupStatusOperOper(ret edpt.DataBackupStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status":   ret.DtBackupStatusOper.Oper.Status,
			"message":  ret.DtBackupStatusOper.Oper.Message,
			"size":     ret.DtBackupStatusOper.Oper.Size,
			"progress": ret.DtBackupStatusOper.Oper.Progress,
		},
	}
}

func getObjectBackupStatusOperOper(d []interface{}) edpt.BackupStatusOperOper {

	count1 := len(d)
	var ret edpt.BackupStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(int)
		ret.Message = in["message"].(string)
		ret.Size = in["size"].(int)
		ret.Progress = in["progress"].(int)
	}
	return ret
}

func dataToEndpointBackupStatusOper(d *schema.ResourceData) edpt.BackupStatusOper {
	var ret edpt.BackupStatusOper

	ret.Oper = getObjectBackupStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
