package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthUpReasonOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_up_reason_oper`: Operational Status for the object health-up-reason\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthUpReasonOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"up_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"up_reason": {
							Type: schema.TypeString, Optional: true, Description: "health up reason",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthUpReasonOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthUpReasonOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthUpReasonOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthUpReasonOperOper := setObjectSlbHealthUpReasonOperOper(res)
		d.Set("oper", SlbHealthUpReasonOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthUpReasonOperOper(ret edpt.DataSlbHealthUpReasonOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"up_id":     ret.DtSlbHealthUpReasonOper.Oper.Up_id,
			"up_reason": ret.DtSlbHealthUpReasonOper.Oper.Up_reason,
		},
	}
}

func getObjectSlbHealthUpReasonOperOper(d []interface{}) edpt.SlbHealthUpReasonOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthUpReasonOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Up_id = in["up_id"].(int)
		ret.Up_reason = in["up_reason"].(string)
	}
	return ret
}

func dataToEndpointSlbHealthUpReasonOper(d *schema.ResourceData) edpt.SlbHealthUpReasonOper {
	var ret edpt.SlbHealthUpReasonOper

	ret.Oper = getObjectSlbHealthUpReasonOperOper(d.Get("oper").([]interface{}))
	return ret
}
