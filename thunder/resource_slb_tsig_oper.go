package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTsigOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_tsig_oper`: Operational Status for the object tsig\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTsigOperRead,

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
								},
							},
						},
						"tsig_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTsigOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTsigOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTsigOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTsigOperOper := setObjectSlbTsigOperOper(res)
		d.Set("oper", SlbTsigOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTsigOperOper(ret edpt.DataSlbTsigOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter_entry": ret.DtSlbTsigOper.Oper.Filter_entry,
			"file_list":    setSliceSlbTsigOperOperFileList(ret.DtSlbTsigOper.Oper.FileList),
			"tsig_count":   ret.DtSlbTsigOper.Oper.TsigCount,
		},
	}
}

func setSliceSlbTsigOperOperFileList(d []edpt.SlbTsigOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		result = append(result, in)
	}
	return result
}

func getObjectSlbTsigOperOper(d []interface{}) edpt.SlbTsigOperOper {

	count1 := len(d)
	var ret edpt.SlbTsigOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Filter_entry = in["filter_entry"].(string)
		ret.FileList = getSliceSlbTsigOperOperFileList(in["file_list"].([]interface{}))
		ret.TsigCount = in["tsig_count"].(int)
	}
	return ret
}

func getSliceSlbTsigOperOperFileList(d []interface{}) []edpt.SlbTsigOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.SlbTsigOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTsigOperOperFileList
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTsigOper(d *schema.ResourceData) edpt.SlbTsigOper {
	var ret edpt.SlbTsigOper

	ret.Oper = getObjectSlbTsigOperOper(d.Get("oper").([]interface{}))
	return ret
}
