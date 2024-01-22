package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslScepCertLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_scep_cert_log_oper`: Operational Status for the object ssl-scep-cert-log\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslScepCertLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scep_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scep_log_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"scep_log_offset": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"scep_log_over": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSlbSslScepCertLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslScepCertLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslScepCertLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslScepCertLogOperOper := setObjectSlbSslScepCertLogOperOper(res)
		d.Set("oper", SlbSslScepCertLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslScepCertLogOperOper(ret edpt.DataSlbSslScepCertLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scep_log_list":   setSliceSlbSslScepCertLogOperOperScepLogList(ret.DtSlbSslScepCertLogOper.Oper.ScepLogList),
			"scep_log_offset": ret.DtSlbSslScepCertLogOper.Oper.ScepLogOffset,
			"scep_log_over":   ret.DtSlbSslScepCertLogOper.Oper.ScepLogOver,
			"name":            ret.DtSlbSslScepCertLogOper.Oper.Name,
			"follow":          ret.DtSlbSslScepCertLogOper.Oper.Follow,
			"from_start":      ret.DtSlbSslScepCertLogOper.Oper.FromStart,
			"num_lines":       ret.DtSlbSslScepCertLogOper.Oper.NumLines,
		},
	}
}

func setSliceSlbSslScepCertLogOperOperScepLogList(d []edpt.SlbSslScepCertLogOperOperScepLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["scep_log_data"] = item.ScepLogData
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslScepCertLogOperOper(d []interface{}) edpt.SlbSslScepCertLogOperOper {

	count1 := len(d)
	var ret edpt.SlbSslScepCertLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ScepLogList = getSliceSlbSslScepCertLogOperOperScepLogList(in["scep_log_list"].([]interface{}))
		ret.ScepLogOffset = in["scep_log_offset"].(int)
		ret.ScepLogOver = in["scep_log_over"].(int)
		ret.Name = in["name"].(string)
		ret.Follow = in["follow"].(int)
		ret.FromStart = in["from_start"].(int)
		ret.NumLines = in["num_lines"].(int)
	}
	return ret
}

func getSliceSlbSslScepCertLogOperOperScepLogList(d []interface{}) []edpt.SlbSslScepCertLogOperOperScepLogList {

	count1 := len(d)
	ret := make([]edpt.SlbSslScepCertLogOperOperScepLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslScepCertLogOperOperScepLogList
		oi.ScepLogData = in["scep_log_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslScepCertLogOper(d *schema.ResourceData) edpt.SlbSslScepCertLogOper {
	var ret edpt.SlbSslScepCertLogOper

	ret.Oper = getObjectSlbSslScepCertLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
