package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthDownReasonOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_down_reason_oper`: Operational Status for the object health-down-reason\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthDownReasonOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"down_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"down_reason": {
							Type: schema.TypeString, Optional: true, Description: "health down reason",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthDownReasonOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthDownReasonOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthDownReasonOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthDownReasonOperOper := setObjectSlbHealthDownReasonOperOper(res)
		d.Set("oper", SlbHealthDownReasonOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthDownReasonOperOper(ret edpt.DataSlbHealthDownReasonOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"down_id":     ret.DtSlbHealthDownReasonOper.Oper.Down_id,
			"down_reason": ret.DtSlbHealthDownReasonOper.Oper.Down_reason,
		},
	}
}

func getObjectSlbHealthDownReasonOperOper(d []interface{}) edpt.SlbHealthDownReasonOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthDownReasonOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Down_id = in["down_id"].(int)
		ret.Down_reason = in["down_reason"].(string)
	}
	return ret
}

func dataToEndpointSlbHealthDownReasonOper(d *schema.ResourceData) edpt.SlbHealthDownReasonOper {
	var ret edpt.SlbHealthDownReasonOper

	ret.Oper = getObjectSlbHealthDownReasonOperOper(d.Get("oper").([]interface{}))
	return ret
}
