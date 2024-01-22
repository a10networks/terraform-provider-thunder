package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlMetadataOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_saml_metadata_oper`: Operational Status for the object metadata\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSamlMetadataOperRead,

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
									"binding": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceAamAuthenticationSamlMetadataOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlMetadataOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlMetadataOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSamlMetadataOperOper := setObjectAamAuthenticationSamlMetadataOperOper(res)
		d.Set("oper", AamAuthenticationSamlMetadataOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSamlMetadataOperOper(ret edpt.DataAamAuthenticationSamlMetadataOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"file_list": setSliceAamAuthenticationSamlMetadataOperOperFileList(ret.DtAamAuthenticationSamlMetadataOper.Oper.FileList),
		},
	}
}

func setSliceAamAuthenticationSamlMetadataOperOperFileList(d []edpt.AamAuthenticationSamlMetadataOperOperFileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["file"] = item.File
		in["size"] = item.Size
		in["binding"] = item.Binding
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationSamlMetadataOperOper(d []interface{}) edpt.AamAuthenticationSamlMetadataOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlMetadataOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FileList = getSliceAamAuthenticationSamlMetadataOperOperFileList(in["file_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationSamlMetadataOperOperFileList(d []interface{}) []edpt.AamAuthenticationSamlMetadataOperOperFileList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlMetadataOperOperFileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlMetadataOperOperFileList
		oi.File = in["file"].(string)
		oi.Size = in["size"].(int)
		oi.Binding = in["binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlMetadataOper(d *schema.ResourceData) edpt.AamAuthenticationSamlMetadataOper {
	var ret edpt.AamAuthenticationSamlMetadataOper

	ret.Oper = getObjectAamAuthenticationSamlMetadataOperOper(d.Get("oper").([]interface{}))
	return ret
}
