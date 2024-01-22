package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpsProfileOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ips_profile_oper`: Operational Status for the object profile\n\n__PLACEHOLDER__",
		ReadContext: resourceIpsProfileOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"profile_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"is_predefined": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reference_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"signature_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"signature_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"enable": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"missing_signature": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"message": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"profile_detail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceIpsProfileOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsProfileOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsProfileOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpsProfileOperOper := setObjectIpsProfileOperOper(res)
		d.Set("oper", IpsProfileOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpsProfileOperOper(ret edpt.DataIpsProfileOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"profile_list":   setSliceIpsProfileOperOperProfileList(ret.DtIpsProfileOper.Oper.ProfileList),
			"name":           ret.DtIpsProfileOper.Oper.Name,
			"sid":            ret.DtIpsProfileOper.Oper.Sid,
			"profile_detail": ret.DtIpsProfileOper.Oper.ProfileDetail,
		},
	}
}

func setSliceIpsProfileOperOperProfileList(d []edpt.IpsProfileOperOperProfileList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["profile_name"] = item.ProfileName
		in["is_predefined"] = item.IsPredefined
		in["reference_count"] = item.ReferenceCount
		in["signature_count"] = item.SignatureCount
		in["signature_list"] = setSliceIpsProfileOperOperProfileListSignatureList(item.SignatureList)
		result = append(result, in)
	}
	return result
}

func setSliceIpsProfileOperOperProfileListSignatureList(d []edpt.IpsProfileOperOperProfileListSignatureList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sid"] = item.Sid
		in["enable"] = item.Enable
		in["action"] = item.Action
		in["missing_signature"] = item.MissingSignature
		in["message"] = item.Message
		result = append(result, in)
	}
	return result
}

func getObjectIpsProfileOperOper(d []interface{}) edpt.IpsProfileOperOper {

	count1 := len(d)
	var ret edpt.IpsProfileOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ProfileList = getSliceIpsProfileOperOperProfileList(in["profile_list"].([]interface{}))
		ret.Name = in["name"].(string)
		ret.Sid = in["sid"].(int)
		ret.ProfileDetail = in["profile_detail"].(int)
	}
	return ret
}

func getSliceIpsProfileOperOperProfileList(d []interface{}) []edpt.IpsProfileOperOperProfileList {

	count1 := len(d)
	ret := make([]edpt.IpsProfileOperOperProfileList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpsProfileOperOperProfileList
		oi.ProfileName = in["profile_name"].(string)
		oi.IsPredefined = in["is_predefined"].(int)
		oi.ReferenceCount = in["reference_count"].(int)
		oi.SignatureCount = in["signature_count"].(int)
		oi.SignatureList = getSliceIpsProfileOperOperProfileListSignatureList(in["signature_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpsProfileOperOperProfileListSignatureList(d []interface{}) []edpt.IpsProfileOperOperProfileListSignatureList {

	count1 := len(d)
	ret := make([]edpt.IpsProfileOperOperProfileListSignatureList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpsProfileOperOperProfileListSignatureList
		oi.Sid = in["sid"].(int)
		oi.Enable = in["enable"].(int)
		oi.Action = in["action"].(string)
		oi.MissingSignature = in["missing_signature"].(int)
		oi.Message = in["message"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpsProfileOper(d *schema.ResourceData) edpt.IpsProfileOper {
	var ret edpt.IpsProfileOper

	ret.Oper = getObjectIpsProfileOperOper(d.Get("oper").([]interface{}))
	return ret
}
