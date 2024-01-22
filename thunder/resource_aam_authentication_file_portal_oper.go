package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationFilePortalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_file_portal_oper`: Operational Status for the object portal\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationFilePortalOperRead,

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
						"portal_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationFilePortalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationFilePortalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationFilePortalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationFilePortalOperOper := setObjectAamAuthenticationFilePortalOperOper(res)
		d.Set("oper", AamAuthenticationFilePortalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationFilePortalOperOper(ret edpt.DataAamAuthenticationFilePortalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list":   setSliceAamAuthenticationFilePortalOperOperFileList(ret.DtAamAuthenticationFilePortalOper.Oper.FileList),
			"portal_name": ret.DtAamAuthenticationFilePortalOper.Oper.PortalName,
		},
	}
}

func setSliceAamAuthenticationFilePortalOperOperFileList(d []edpt.AamAuthenticationFilePortalOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["size"] = item.Size
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationFilePortalOperOper(d []interface{}) edpt.AamAuthenticationFilePortalOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationFilePortalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceAamAuthenticationFilePortalOperOperFileList(in["file_list"].([]interface{}))
		ret.PortalName = in["portal_name"].(string)
	}
	return ret
}

func getSliceAamAuthenticationFilePortalOperOperFileList(d []interface{}) []edpt.AamAuthenticationFilePortalOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationFilePortalOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationFilePortalOperOperFileList
		oi.File = in["file"].(string)
		oi.Size = in["size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationFilePortalOper(d *schema.ResourceData) edpt.AamAuthenticationFilePortalOper {
	var ret edpt.AamAuthenticationFilePortalOper

	ret.Oper = getObjectAamAuthenticationFilePortalOperOper(d.Get("oper").([]interface{}))
	return ret
}
