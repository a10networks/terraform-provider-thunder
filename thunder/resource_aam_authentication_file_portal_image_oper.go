package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationFilePortalImageOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_file_portal_image_oper`: Operational Status for the object portal-image\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationFilePortalImageOperRead,

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
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationFilePortalImageOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationFilePortalImageOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationFilePortalImageOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationFilePortalImageOperOper := setObjectAamAuthenticationFilePortalImageOperOper(res)
		d.Set("oper", AamAuthenticationFilePortalImageOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationFilePortalImageOperOper(ret edpt.DataAamAuthenticationFilePortalImageOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list": setSliceAamAuthenticationFilePortalImageOperOperFileList(ret.DtAamAuthenticationFilePortalImageOper.Oper.FileList),
			"name":      ret.DtAamAuthenticationFilePortalImageOper.Oper.Name,
		},
	}
}

func setSliceAamAuthenticationFilePortalImageOperOperFileList(d []edpt.AamAuthenticationFilePortalImageOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["size"] = item.Size
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationFilePortalImageOperOper(d []interface{}) edpt.AamAuthenticationFilePortalImageOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationFilePortalImageOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceAamAuthenticationFilePortalImageOperOperFileList(in["file_list"].([]interface{}))
		ret.Name = in["name"].(string)
	}
	return ret
}

func getSliceAamAuthenticationFilePortalImageOperOperFileList(d []interface{}) []edpt.AamAuthenticationFilePortalImageOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationFilePortalImageOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationFilePortalImageOperOperFileList
		oi.File = in["file"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationFilePortalImageOper(d *schema.ResourceData) edpt.AamAuthenticationFilePortalImageOper {
	var ret edpt.AamAuthenticationFilePortalImageOper

	ret.Oper = getObjectAamAuthenticationFilePortalImageOperOper(d.Get("oper").([]interface{}))
	return ret
}
