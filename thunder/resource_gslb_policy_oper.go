package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbPolicyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_policy_oper`: Operational Status for the object policy\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbPolicyOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify policy name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"order": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceGslbPolicyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbPolicyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbPolicyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbPolicyOperOper := setObjectGslbPolicyOperOper(res)
		d.Set("oper", GslbPolicyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbPolicyOperOper(ret edpt.DataGslbPolicyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"metric_list": setSliceGslbPolicyOperOperMetricList(ret.DtGslbPolicyOper.Oper.MetricList),
		},
	}
}

func setSliceGslbPolicyOperOperMetricList(d []edpt.GslbPolicyOperOperMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["order"] = item.Order
		result = append(result, in)
	}
	return result
}

func getObjectGslbPolicyOperOper(d []interface{}) edpt.GslbPolicyOperOper {

	count1 := len(d)
	var ret edpt.GslbPolicyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MetricList = getSliceGslbPolicyOperOperMetricList(in["metric_list"].([]interface{}))
	}
	return ret
}

func getSliceGslbPolicyOperOperMetricList(d []interface{}) []edpt.GslbPolicyOperOperMetricList {

	count1 := len(d)
	ret := make([]edpt.GslbPolicyOperOperMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbPolicyOperOperMetricList
		oi.Type = in["type"].(string)
		oi.Order = in["order"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbPolicyOper(d *schema.ResourceData) edpt.GslbPolicyOper {
	var ret edpt.GslbPolicyOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectGslbPolicyOperOper(d.Get("oper").([]interface{}))
	return ret
}
