package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosGeoLocationFileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_geo_location_file_oper`: Operational Status for the object file\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosGeoLocationFileOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filename": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
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
									"error_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"error_line": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"error_information": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
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

func resourceDdosGeoLocationFileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosGeoLocationFileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosGeoLocationFileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosGeoLocationFileOperOper := setObjectDdosGeoLocationFileOperOper(res)
		d.Set("oper", DdosGeoLocationFileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosGeoLocationFileOperOper(ret edpt.DataDdosGeoLocationFileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list": setSliceDdosGeoLocationFileOperOperFileList(ret.DtDdosGeoLocationFileOper.Oper.FileList),
		},
	}
}

func setSliceDdosGeoLocationFileOperOperFileList(d []edpt.DdosGeoLocationFileOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["filename"] = item.Filename
		in["type"] = item.Type
		in["lines"] = item.Lines
		in["success"] = item.Success
		in["error_warning"] = item.ErrorWarning
		in["error_list"] = setSliceDdosGeoLocationFileOperOperFileListErrorList(item.ErrorList)
		result = append(result, in)
	}
	return result
}

func setSliceDdosGeoLocationFileOperOperFileListErrorList(d []edpt.DdosGeoLocationFileOperOperFileListErrorList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["error_line"] = item.ErrorLine
		in["error_information"] = item.ErrorInformation
		result = append(result, in)
	}
	return result
}

func getObjectDdosGeoLocationFileOperOper(d []interface{}) edpt.DdosGeoLocationFileOperOper {

	count1 := len(d)
	var ret edpt.DdosGeoLocationFileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceDdosGeoLocationFileOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosGeoLocationFileOperOperFileList(d []interface{}) []edpt.DdosGeoLocationFileOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.DdosGeoLocationFileOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosGeoLocationFileOperOperFileList
		oi.Filename = in["filename"].(string)
		oi.Type = in["type"].(string)
		oi.Lines = in["lines"].(int)
		oi.Success = in["success"].(int)
		oi.ErrorWarning = in["error_warning"].(int)
		oi.ErrorList = getSliceDdosGeoLocationFileOperOperFileListErrorList(in["error_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosGeoLocationFileOperOperFileListErrorList(d []interface{}) []edpt.DdosGeoLocationFileOperOperFileListErrorList {

	count1 := len(d)
	ret := make([]edpt.DdosGeoLocationFileOperOperFileListErrorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosGeoLocationFileOperOperFileListErrorList
		oi.ErrorLine = in["error_line"].(int)
		oi.ErrorInformation = in["error_information"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosGeoLocationFileOper(d *schema.ResourceData) edpt.DdosGeoLocationFileOper {
	var ret edpt.DdosGeoLocationFileOper

	ret.Oper = getObjectDdosGeoLocationFileOperOper(d.Get("oper").([]interface{}))
	return ret
}
