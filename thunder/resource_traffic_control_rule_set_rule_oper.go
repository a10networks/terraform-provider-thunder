package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetRuleOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_traffic_control_rule_set_rule_oper`: Operational Status for the object rule\n\n__PLACEHOLDER__",
		ReadContext: resourceTrafficControlRuleSetRuleOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hitcount": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceTrafficControlRuleSetRuleOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TrafficControlRuleSetRuleOperOper := setObjectTrafficControlRuleSetRuleOperOper(res)
		d.Set("oper", TrafficControlRuleSetRuleOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTrafficControlRuleSetRuleOperOper(ret edpt.DataTrafficControlRuleSetRuleOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status":   ret.DtTrafficControlRuleSetRuleOper.Oper.Status,
			"hitcount": ret.DtTrafficControlRuleSetRuleOper.Oper.Hitcount,
		},
	}
}

func getObjectTrafficControlRuleSetRuleOperOper(d []interface{}) edpt.TrafficControlRuleSetRuleOperOper {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetRuleOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.Hitcount = in["hitcount"].(int)
	}
	return ret
}

func dataToEndpointTrafficControlRuleSetRuleOper(d *schema.ResourceData) edpt.TrafficControlRuleSetRuleOper {
	var ret edpt.TrafficControlRuleSetRuleOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectTrafficControlRuleSetRuleOperOper(d.Get("oper").([]interface{}))
	return ret
}
