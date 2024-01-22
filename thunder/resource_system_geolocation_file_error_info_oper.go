package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocationFileErrorInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geolocation_file_error_info_oper`: Operational Status for the object error-info\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocationFileErrorInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"offset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"filename": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemGeolocationFileErrorInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocationFileErrorInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocationFileErrorInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocationFileErrorInfoOperOper := setObjectSystemGeolocationFileErrorInfoOperOper(res)
		d.Set("oper", SystemGeolocationFileErrorInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocationFileErrorInfoOperOper(ret edpt.DataSystemGeolocationFileErrorInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"error_list": setSliceSystemGeolocationFileErrorInfoOperOperErrorList(ret.DtSystemGeolocationFileErrorInfoOper.Oper.ErrorList),
			"filename":   ret.DtSystemGeolocationFileErrorInfoOper.Oper.Filename,
		},
	}
}

func setSliceSystemGeolocationFileErrorInfoOperOperErrorList(d []edpt.SystemGeolocationFileErrorInfoOperOperErrorList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["line"] = item.Line
		in["offset"] = item.Offset
		in["error"] = item.Error
		result = append(result, in)
	}
	return result
}

func getObjectSystemGeolocationFileErrorInfoOperOper(d []interface{}) edpt.SystemGeolocationFileErrorInfoOperOper {

	count1 := len(d)
	var ret edpt.SystemGeolocationFileErrorInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ErrorList = getSliceSystemGeolocationFileErrorInfoOperOperErrorList(in["error_list"].([]interface{}))
		ret.Filename = in["filename"].(string)
	}
	return ret
}

func getSliceSystemGeolocationFileErrorInfoOperOperErrorList(d []interface{}) []edpt.SystemGeolocationFileErrorInfoOperOperErrorList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocationFileErrorInfoOperOperErrorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocationFileErrorInfoOperOperErrorList
		oi.Line = in["line"].(int)
		oi.Offset = in["offset"].(int)
		oi.Error = in["error"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocationFileErrorInfoOper(d *schema.ResourceData) edpt.SystemGeolocationFileErrorInfoOper {
	var ret edpt.SystemGeolocationFileErrorInfoOper

	ret.Oper = getObjectSystemGeolocationFileErrorInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
