package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelWebrootLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_threat_intel_webroot_log_oper`: Operational Status for the object webroot-log\n\n__PLACEHOLDER__",
		ReadContext: resourceThreatIntelWebrootLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"webroot_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"webroot_log_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"webroot_log_offset": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"webroot_log_over": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"follow": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"from_start": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_lines": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceThreatIntelWebrootLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ThreatIntelWebrootLogOperOper := setObjectThreatIntelWebrootLogOperOper(res)
		d.Set("oper", ThreatIntelWebrootLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectThreatIntelWebrootLogOperOper(ret edpt.DataThreatIntelWebrootLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"webroot_log_list":   setSliceThreatIntelWebrootLogOperOperWebrootLogList(ret.DtThreatIntelWebrootLogOper.Oper.WebrootLogList),
			"webroot_log_offset": ret.DtThreatIntelWebrootLogOper.Oper.WebrootLogOffset,
			"webroot_log_over":   ret.DtThreatIntelWebrootLogOper.Oper.WebrootLogOver,
			"follow":             ret.DtThreatIntelWebrootLogOper.Oper.Follow,
			"from_start":         ret.DtThreatIntelWebrootLogOper.Oper.FromStart,
			"num_lines":          ret.DtThreatIntelWebrootLogOper.Oper.NumLines,
		},
	}
}

func setSliceThreatIntelWebrootLogOperOperWebrootLogList(d []edpt.ThreatIntelWebrootLogOperOperWebrootLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["webroot_log_data"] = item.WebrootLogData
		result = append(result, in)
	}
	return result
}

func getObjectThreatIntelWebrootLogOperOper(d []interface{}) edpt.ThreatIntelWebrootLogOperOper {

	count1 := len(d)
	var ret edpt.ThreatIntelWebrootLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.WebrootLogList = getSliceThreatIntelWebrootLogOperOperWebrootLogList(in["webroot_log_list"].([]interface{}))
		ret.WebrootLogOffset = in["webroot_log_offset"].(int)
		ret.WebrootLogOver = in["webroot_log_over"].(int)
		ret.Follow = in["follow"].(int)
		ret.FromStart = in["from_start"].(int)
		ret.NumLines = in["num_lines"].(int)
	}
	return ret
}

func getSliceThreatIntelWebrootLogOperOperWebrootLogList(d []interface{}) []edpt.ThreatIntelWebrootLogOperOperWebrootLogList {

	count1 := len(d)
	ret := make([]edpt.ThreatIntelWebrootLogOperOperWebrootLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ThreatIntelWebrootLogOperOperWebrootLogList
		oi.WebrootLogData = in["webroot_log_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointThreatIntelWebrootLogOper(d *schema.ResourceData) edpt.ThreatIntelWebrootLogOper {
	var ret edpt.ThreatIntelWebrootLogOper

	ret.Oper = getObjectThreatIntelWebrootLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
