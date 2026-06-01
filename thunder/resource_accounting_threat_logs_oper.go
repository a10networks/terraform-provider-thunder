package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAccountingThreatLogsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_accounting_threat_logs_oper`: Operational Status for the object threat-logs\n\n__PLACEHOLDER__",
		ReadContext: resourceAccountingThreatLogsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"result": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"msg": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAccountingThreatLogsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAccountingThreatLogsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAccountingThreatLogsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AccountingThreatLogsOperOper := setObjectAccountingThreatLogsOperOper(res)
		d.Set("oper", AccountingThreatLogsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAccountingThreatLogsOperOper(ret edpt.DataAccountingThreatLogsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status": ret.DtAccountingThreatLogsOper.Oper.Status,
			"result": ret.DtAccountingThreatLogsOper.Oper.Result,
			"msg":    ret.DtAccountingThreatLogsOper.Oper.Msg,
		},
	}
}

func getObjectAccountingThreatLogsOperOper(d []interface{}) edpt.AccountingThreatLogsOperOper {

	count1 := len(d)
	var ret edpt.AccountingThreatLogsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.Result = in["result"].(int)
		ret.Msg = in["msg"].(string)
	}
	return ret
}

func dataToEndpointAccountingThreatLogsOper(d *schema.ResourceData) edpt.AccountingThreatLogsOper {
	var ret edpt.AccountingThreatLogsOper

	ret.Oper = getObjectAccountingThreatLogsOperOper(d.Get("oper").([]interface{}))
	return ret
}
