package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemUpgradeStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_upgrade_status_oper`: Operational Status for the object upgrade-status\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemUpgradeStatusOperRead,

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
						"rollback_pri": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rollback_sec": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemUpgradeStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemUpgradeStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemUpgradeStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemUpgradeStatusOperOper := setObjectSystemUpgradeStatusOperOper(res)
		d.Set("oper", SystemUpgradeStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemUpgradeStatusOperOper(ret edpt.DataSystemUpgradeStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status":       ret.DtSystemUpgradeStatusOper.Oper.Status,
			"message":      ret.DtSystemUpgradeStatusOper.Oper.Message,
			"rollback_pri": ret.DtSystemUpgradeStatusOper.Oper.RollbackPri,
			"rollback_sec": ret.DtSystemUpgradeStatusOper.Oper.RollbackSec,
		},
	}
}

func getObjectSystemUpgradeStatusOperOper(d []interface{}) edpt.SystemUpgradeStatusOperOper {

	count1 := len(d)
	var ret edpt.SystemUpgradeStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(int)
		ret.Message = in["message"].(string)
		ret.RollbackPri = in["rollback_pri"].(string)
		ret.RollbackSec = in["rollback_sec"].(string)
	}
	return ret
}

func dataToEndpointSystemUpgradeStatusOper(d *schema.ResourceData) edpt.SystemUpgradeStatusOper {
	var ret edpt.SystemUpgradeStatusOper

	ret.Oper = getObjectSystemUpgradeStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
