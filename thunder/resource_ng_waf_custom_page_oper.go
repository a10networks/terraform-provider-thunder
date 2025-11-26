package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNgWafCustomPageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ng_waf_custom_page_oper`: Operational Status for the object custom-page\n\n__PLACEHOLDER__",
		ReadContext: resourceNgWafCustomPageOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"file": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"size": {
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

func resourceNgWafCustomPageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafCustomPageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWafCustomPageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		NgWafCustomPageOperOper := setObjectNgWafCustomPageOperOper(res)
		d.Set("oper", NgWafCustomPageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectNgWafCustomPageOperOper(ret edpt.DataNgWafCustomPageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list": setSliceNgWafCustomPageOperOperFileList(ret.DtNgWafCustomPageOper.Oper.FileList),
		},
	}
}

func setSliceNgWafCustomPageOperOperFileList(d []edpt.NgWafCustomPageOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["size"] = item.Size
		result = append(result, in)
	}
	return result
}

func getObjectNgWafCustomPageOperOper(d []interface{}) edpt.NgWafCustomPageOperOper {

	count1 := len(d)
	var ret edpt.NgWafCustomPageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceNgWafCustomPageOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceNgWafCustomPageOperOperFileList(d []interface{}) []edpt.NgWafCustomPageOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.NgWafCustomPageOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafCustomPageOperOperFileList
		oi.File = in["file"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNgWafCustomPageOper(d *schema.ResourceData) edpt.NgWafCustomPageOper {
	var ret edpt.NgWafCustomPageOper

	ret.Oper = getObjectNgWafCustomPageOperOper(d.Get("oper").([]interface{}))
	return ret
}
