package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCoreOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_core_oper`: Operational Status for the object core\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCoreOperRead,

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
									"update_time": {
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

func resourceSystemCoreOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCoreOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCoreOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCoreOperOper := setObjectSystemCoreOperOper(res)
		d.Set("oper", SystemCoreOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCoreOperOper(ret edpt.DataSystemCoreOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list": setSliceSystemCoreOperOperFileList(ret.DtSystemCoreOper.Oper.FileList),
		},
	}
}

func setSliceSystemCoreOperOperFileList(d []edpt.SystemCoreOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["update_time"] = item.UpdateTime
		in["size"] = item.Size
		result = append(result, in)
	}
	return result
}

func getObjectSystemCoreOperOper(d []interface{}) edpt.SystemCoreOperOper {

	count1 := len(d)
	var ret edpt.SystemCoreOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceSystemCoreOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemCoreOperOperFileList(d []interface{}) []edpt.SystemCoreOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.SystemCoreOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCoreOperOperFileList
		oi.File = in["file"].(string)
		oi.UpdateTime = in["update_time"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCoreOper(d *schema.ResourceData) edpt.SystemCoreOper {
	var ret edpt.SystemCoreOper

	ret.Oper = getObjectSystemCoreOperOper(d.Get("oper").([]interface{}))
	return ret
}
