package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionDdosScriptOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_detection_ddos_script_oper`: Operational Status for the object ddos-script\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDetectionDdosScriptOperRead,

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
									"reference_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"file_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_records": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDetectionDdosScriptOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionDdosScriptOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionDdosScriptOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDetectionDdosScriptOperOper := setObjectDdosDetectionDdosScriptOperOper(res)
		d.Set("oper", DdosDetectionDdosScriptOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDetectionDdosScriptOperOper(ret edpt.DataDdosDetectionDdosScriptOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list":     setSliceDdosDetectionDdosScriptOperOperFileList(ret.DtDdosDetectionDdosScriptOper.Oper.FileList),
			"total_records": ret.DtDdosDetectionDdosScriptOper.Oper.TotalRecords,
		},
	}
}

func setSliceDdosDetectionDdosScriptOperOperFileList(d []edpt.DdosDetectionDdosScriptOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["reference_count"] = item.ReferenceCount
		in["file_size"] = item.FileSize
		result = append(result, in)
	}
	return result
}

func getObjectDdosDetectionDdosScriptOperOper(d []interface{}) edpt.DdosDetectionDdosScriptOperOper {

	count1 := len(d)
	var ret edpt.DdosDetectionDdosScriptOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceDdosDetectionDdosScriptOperOperFileList(in["file_list"].([]interface{}))
		ret.TotalRecords = in["total_records"].(int)
	}
	return ret
}

func getSliceDdosDetectionDdosScriptOperOperFileList(d []interface{}) []edpt.DdosDetectionDdosScriptOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionDdosScriptOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionDdosScriptOperOperFileList
		oi.File = in["file"].(string)
		oi.ReferenceCount = in["reference_count"].(int)
		oi.FileSize = in["file_size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDetectionDdosScriptOper(d *schema.ResourceData) edpt.DdosDetectionDdosScriptOper {
	var ret edpt.DdosDetectionDdosScriptOper

	ret.Oper = getObjectDdosDetectionDdosScriptOperOper(d.Get("oper").([]interface{}))
	return ret
}
