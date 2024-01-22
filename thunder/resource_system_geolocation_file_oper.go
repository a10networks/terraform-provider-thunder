package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemGeolocationFileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_geolocation_file_oper`: Operational Status for the object geolocation-file\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemGeolocationFileOperRead,

		Schema: map[string]*schema.Schema{
			"error_info": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geofiles": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filename": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"template": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"percentage_loaded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lines": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error_warning": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"comment": {
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

func resourceSystemGeolocationFileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemGeolocationFileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemGeolocationFileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemGeolocationFileOperErrorInfo := setObjectSystemGeolocationFileOperErrorInfo(res)
		d.Set("error_info", SystemGeolocationFileOperErrorInfo)
		SystemGeolocationFileOperOper := setObjectSystemGeolocationFileOperOper(res)
		d.Set("oper", SystemGeolocationFileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemGeolocationFileOperErrorInfo(ret edpt.DataSystemGeolocationFileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectSystemGeolocationFileOperErrorInfoOper(ret.DtSystemGeolocationFileOper.ErrorInfo.Oper),
		},
	}
}

func setObjectSystemGeolocationFileOperErrorInfoOper(d edpt.SystemGeolocationFileOperErrorInfoOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["error_list"] = setSliceSystemGeolocationFileOperErrorInfoOperErrorList(d.ErrorList)

	in["filename"] = d.Filename
	result = append(result, in)
	return result
}

func setSliceSystemGeolocationFileOperErrorInfoOperErrorList(d []edpt.SystemGeolocationFileOperErrorInfoOperErrorList) []map[string]interface{} {
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

func setObjectSystemGeolocationFileOperOper(ret edpt.DataSystemGeolocationFileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geofiles": setSliceSystemGeolocationFileOperOperGeofiles(ret.DtSystemGeolocationFileOper.Oper.Geofiles),
		},
	}
}

func setSliceSystemGeolocationFileOperOperGeofiles(d []edpt.SystemGeolocationFileOperOperGeofiles) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["filename"] = item.Filename
		in["type"] = item.Type
		in["template"] = item.Template
		in["percentage_loaded"] = item.PercentageLoaded
		in["lines"] = item.Lines
		in["success"] = item.Success
		in["error_warning"] = item.ErrorWarning
		in["comment"] = item.Comment
		result = append(result, in)
	}
	return result
}

func getObjectSystemGeolocationFileOperErrorInfo(d []interface{}) edpt.SystemGeolocationFileOperErrorInfo {

	count1 := len(d)
	var ret edpt.SystemGeolocationFileOperErrorInfo
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectSystemGeolocationFileOperErrorInfoOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectSystemGeolocationFileOperErrorInfoOper(d []interface{}) edpt.SystemGeolocationFileOperErrorInfoOper {

	count1 := len(d)
	var ret edpt.SystemGeolocationFileOperErrorInfoOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ErrorList = getSliceSystemGeolocationFileOperErrorInfoOperErrorList(in["error_list"].([]interface{}))
		ret.Filename = in["filename"].(string)
	}
	return ret
}

func getSliceSystemGeolocationFileOperErrorInfoOperErrorList(d []interface{}) []edpt.SystemGeolocationFileOperErrorInfoOperErrorList {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocationFileOperErrorInfoOperErrorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocationFileOperErrorInfoOperErrorList
		oi.Line = in["line"].(int)
		oi.Offset = in["offset"].(int)
		oi.Error = in["error"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemGeolocationFileOperOper(d []interface{}) edpt.SystemGeolocationFileOperOper {

	count1 := len(d)
	var ret edpt.SystemGeolocationFileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Geofiles = getSliceSystemGeolocationFileOperOperGeofiles(in["geofiles"].([]interface{}))
	}
	return ret
}

func getSliceSystemGeolocationFileOperOperGeofiles(d []interface{}) []edpt.SystemGeolocationFileOperOperGeofiles {

	count1 := len(d)
	ret := make([]edpt.SystemGeolocationFileOperOperGeofiles, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemGeolocationFileOperOperGeofiles
		oi.Filename = in["filename"].(string)
		oi.Type = in["type"].(string)
		oi.Template = in["template"].(string)
		oi.PercentageLoaded = in["percentage_loaded"].(int)
		oi.Lines = in["lines"].(int)
		oi.Success = in["success"].(int)
		oi.ErrorWarning = in["error_warning"].(int)
		oi.Comment = in["comment"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemGeolocationFileOper(d *schema.ResourceData) edpt.SystemGeolocationFileOper {
	var ret edpt.SystemGeolocationFileOper

	ret.ErrorInfo = getObjectSystemGeolocationFileOperErrorInfo(d.Get("error_info").([]interface{}))

	ret.Oper = getObjectSystemGeolocationFileOperOper(d.Get("oper").([]interface{}))
	return ret
}
