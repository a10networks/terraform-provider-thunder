package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbGeolocationFileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_gslb_geolocation_file_oper`: Operational Status for the object geolocation-file\n\n__PLACEHOLDER__",
		ReadContext: resourceGslbGeolocationFileOperRead,

		Schema: map[string]*schema.Schema{
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

func resourceGslbGeolocationFileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGeolocationFileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGeolocationFileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		GslbGeolocationFileOperOper := setObjectGslbGeolocationFileOperOper(res)
		d.Set("oper", GslbGeolocationFileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectGslbGeolocationFileOperOper(ret edpt.DataGslbGeolocationFileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"geofiles": setSliceGslbGeolocationFileOperOperGeofiles(ret.DtGslbGeolocationFileOper.Oper.Geofiles),
		},
	}
}

func setSliceGslbGeolocationFileOperOperGeofiles(d []edpt.GslbGeolocationFileOperOperGeofiles) []map[string]interface{} {
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

func getObjectGslbGeolocationFileOperOper(d []interface{}) edpt.GslbGeolocationFileOperOper {

	count1 := len(d)
	var ret edpt.GslbGeolocationFileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Geofiles = getSliceGslbGeolocationFileOperOperGeofiles(in["geofiles"].([]interface{}))
	}
	return ret
}

func getSliceGslbGeolocationFileOperOperGeofiles(d []interface{}) []edpt.GslbGeolocationFileOperOperGeofiles {

	count1 := len(d)
	ret := make([]edpt.GslbGeolocationFileOperOperGeofiles, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbGeolocationFileOperOperGeofiles
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

func dataToEndpointGslbGeolocationFileOper(d *schema.ResourceData) edpt.GslbGeolocationFileOper {
	var ret edpt.GslbGeolocationFileOper

	ret.Oper = getObjectGslbGeolocationFileOperOper(d.Get("oper").([]interface{}))
	return ret
}
