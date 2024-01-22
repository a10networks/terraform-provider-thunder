package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslCmpCertLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_cmp_cert_log_oper`: Operational Status for the object ssl-cmp-cert-log\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslCmpCertLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cmp_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmp_log_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"cmp_log_offset": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cmp_log_over": {
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

func resourceSlbSslCmpCertLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslCmpCertLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslCmpCertLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslCmpCertLogOperOper := setObjectSlbSslCmpCertLogOperOper(res)
		d.Set("oper", SlbSslCmpCertLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslCmpCertLogOperOper(ret edpt.DataSlbSslCmpCertLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"cmp_log_list":   setSliceSlbSslCmpCertLogOperOperCmpLogList(ret.DtSlbSslCmpCertLogOper.Oper.CmpLogList),
			"cmp_log_offset": ret.DtSlbSslCmpCertLogOper.Oper.CmpLogOffset,
			"cmp_log_over":   ret.DtSlbSslCmpCertLogOper.Oper.CmpLogOver,
			"name":           ret.DtSlbSslCmpCertLogOper.Oper.Name,
			"follow":         ret.DtSlbSslCmpCertLogOper.Oper.Follow,
			"from_start":     ret.DtSlbSslCmpCertLogOper.Oper.FromStart,
			"num_lines":      ret.DtSlbSslCmpCertLogOper.Oper.NumLines,
		},
	}
}

func setSliceSlbSslCmpCertLogOperOperCmpLogList(d []edpt.SlbSslCmpCertLogOperOperCmpLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cmp_log_data"] = item.CmpLogData
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslCmpCertLogOperOper(d []interface{}) edpt.SlbSslCmpCertLogOperOper {

	count1 := len(d)
	var ret edpt.SlbSslCmpCertLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CmpLogList = getSliceSlbSslCmpCertLogOperOperCmpLogList(in["cmp_log_list"].([]interface{}))
		ret.CmpLogOffset = in["cmp_log_offset"].(int)
		ret.CmpLogOver = in["cmp_log_over"].(int)
		ret.Name = in["name"].(string)
		ret.Follow = in["follow"].(int)
		ret.FromStart = in["from_start"].(int)
		ret.NumLines = in["num_lines"].(int)
	}
	return ret
}

func getSliceSlbSslCmpCertLogOperOperCmpLogList(d []interface{}) []edpt.SlbSslCmpCertLogOperOperCmpLogList {

	count1 := len(d)
	ret := make([]edpt.SlbSslCmpCertLogOperOperCmpLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslCmpCertLogOperOperCmpLogList
		oi.CmpLogData = in["cmp_log_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslCmpCertLogOper(d *schema.ResourceData) edpt.SlbSslCmpCertLogOper {
	var ret edpt.SlbSslCmpCertLogOper

	ret.Oper = getObjectSlbSslCmpCertLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
