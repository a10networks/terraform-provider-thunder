package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslAcmeCertLogOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ssl_acme_cert_log_oper`: Operational Status for the object ssl-acme-cert-log\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSslAcmeCertLogOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acme_log_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"acme_log_data": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"acme_log_offset": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"acme_log_over": {
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

func resourceSlbSslAcmeCertLogOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslAcmeCertLogOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslAcmeCertLogOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSslAcmeCertLogOperOper := setObjectSlbSslAcmeCertLogOperOper(res)
		d.Set("oper", SlbSslAcmeCertLogOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSslAcmeCertLogOperOper(ret edpt.DataSlbSslAcmeCertLogOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"acme_log_list":   setSliceSlbSslAcmeCertLogOperOperAcmeLogList(ret.DtSlbSslAcmeCertLogOper.Oper.AcmeLogList),
			"acme_log_offset": ret.DtSlbSslAcmeCertLogOper.Oper.AcmeLogOffset,
			"acme_log_over":   ret.DtSlbSslAcmeCertLogOper.Oper.AcmeLogOver,
			"name":            ret.DtSlbSslAcmeCertLogOper.Oper.Name,
			"follow":          ret.DtSlbSslAcmeCertLogOper.Oper.Follow,
			"from_start":      ret.DtSlbSslAcmeCertLogOper.Oper.FromStart,
			"num_lines":       ret.DtSlbSslAcmeCertLogOper.Oper.NumLines,
		},
	}
}

func setSliceSlbSslAcmeCertLogOperOperAcmeLogList(d []edpt.SlbSslAcmeCertLogOperOperAcmeLogList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["acme_log_data"] = item.AcmeLogData
		result = append(result, in)
	}
	return result
}

func getObjectSlbSslAcmeCertLogOperOper(d []interface{}) edpt.SlbSslAcmeCertLogOperOper {

	count1 := len(d)
	var ret edpt.SlbSslAcmeCertLogOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AcmeLogList = getSliceSlbSslAcmeCertLogOperOperAcmeLogList(in["acme_log_list"].([]interface{}))
		ret.AcmeLogOffset = in["acme_log_offset"].(int)
		ret.AcmeLogOver = in["acme_log_over"].(int)
		ret.Name = in["name"].(string)
		ret.Follow = in["follow"].(int)
		ret.FromStart = in["from_start"].(int)
		ret.NumLines = in["num_lines"].(int)
	}
	return ret
}

func getSliceSlbSslAcmeCertLogOperOperAcmeLogList(d []interface{}) []edpt.SlbSslAcmeCertLogOperOperAcmeLogList {

	count1 := len(d)
	ret := make([]edpt.SlbSslAcmeCertLogOperOperAcmeLogList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSslAcmeCertLogOperOperAcmeLogList
		oi.AcmeLogData = in["acme_log_data"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSslAcmeCertLogOper(d *schema.ResourceData) edpt.SlbSslAcmeCertLogOper {
	var ret edpt.SlbSslAcmeCertLogOper

	ret.Oper = getObjectSlbSslAcmeCertLogOperOper(d.Get("oper").([]interface{}))
	return ret
}
