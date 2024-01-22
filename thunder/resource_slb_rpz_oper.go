package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRpzOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rpz_oper`: Operational Status for the object rpz\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRpzOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_entry": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dns_template_bound": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"rpz_file_size_max": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rpz_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbRpzOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRpzOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRpzOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRpzOperOper := setObjectSlbRpzOperOper(res)
		d.Set("oper", SlbRpzOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRpzOperOper(ret edpt.DataSlbRpzOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter_entry":      ret.DtSlbRpzOper.Oper.Filter_entry,
			"file_list":         setSliceSlbRpzOperOperFileList(ret.DtSlbRpzOper.Oper.FileList),
			"rpz_file_size_max": ret.DtSlbRpzOper.Oper.RpzFileSizeMax,
			"rpz_count":         ret.DtSlbRpzOper.Oper.RpzCount,
		},
	}
}

func setSliceSlbRpzOperOperFileList(d []edpt.SlbRpzOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["dns_template_bound"] = item.Dns_template_bound
		result = append(result, in)
	}
	return result
}

func getObjectSlbRpzOperOper(d []interface{}) edpt.SlbRpzOperOper {

	count1 := len(d)
	var ret edpt.SlbRpzOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Filter_entry = in["filter_entry"].(string)
		ret.FileList = getSliceSlbRpzOperOperFileList(in["file_list"].([]interface{}))
		ret.RpzFileSizeMax = in["rpz_file_size_max"].(int)
		ret.RpzCount = in["rpz_count"].(int)
	}
	return ret
}

func getSliceSlbRpzOperOperFileList(d []interface{}) []edpt.SlbRpzOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.SlbRpzOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRpzOperOperFileList
		oi.File = in["file"].(string)
		oi.Dns_template_bound = in["dns_template_bound"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbRpzOper(d *schema.ResourceData) edpt.SlbRpzOper {
	var ret edpt.SlbRpzOper

	ret.Oper = getObjectSlbRpzOperOper(d.Get("oper").([]interface{}))
	return ret
}
