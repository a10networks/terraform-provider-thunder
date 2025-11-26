package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNgWafCustomSignalsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ng_waf_custom_signals_oper`: Operational Status for the object custom-signals\n\n__PLACEHOLDER__",
		ReadContext: resourceNgWafCustomSignalsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signal_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"signal": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceNgWafCustomSignalsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafCustomSignalsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWafCustomSignalsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NgWafCustomSignalsOperOper := setObjectNgWafCustomSignalsOperOper(res)
		d.Set("oper", NgWafCustomSignalsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNgWafCustomSignalsOperOper(ret edpt.DataNgWafCustomSignalsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"signal_list": setSliceNgWafCustomSignalsOperOperSignalList(ret.DtNgWafCustomSignalsOper.Oper.SignalList),
		},
	}
}

func setSliceNgWafCustomSignalsOperOperSignalList(d []edpt.NgWafCustomSignalsOperOperSignalList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["signal"] = item.Signal
		result = append(result, in)
	}
	return result
}

func getObjectNgWafCustomSignalsOperOper(d []interface{}) edpt.NgWafCustomSignalsOperOper {

	count1 := len(d)
	var ret edpt.NgWafCustomSignalsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SignalList = getSliceNgWafCustomSignalsOperOperSignalList(in["signal_list"].([]interface{}))
	}
	return ret
}

func getSliceNgWafCustomSignalsOperOperSignalList(d []interface{}) []edpt.NgWafCustomSignalsOperOperSignalList {

	count1 := len(d)
	ret := make([]edpt.NgWafCustomSignalsOperOperSignalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafCustomSignalsOperOperSignalList
		oi.Signal = in["signal"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNgWafCustomSignalsOper(d *schema.ResourceData) edpt.NgWafCustomSignalsOper {
	var ret edpt.NgWafCustomSignalsOper

	ret.Oper = getObjectNgWafCustomSignalsOperOper(d.Get("oper").([]interface{}))
	return ret
}
