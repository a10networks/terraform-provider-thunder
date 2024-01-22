package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlSessionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_saml_session_oper`: Operational Status for the object session\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSamlSessionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sp_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"session_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nameid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"client_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"id_provider": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"auth_instant": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"expire_time": {
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
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationSamlSessionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlSessionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlSessionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSamlSessionOperOper := setObjectAamAuthenticationSamlSessionOperOper(res)
		d.Set("oper", AamAuthenticationSamlSessionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSamlSessionOperOper(ret edpt.DataAamAuthenticationSamlSessionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sp_list": setSliceAamAuthenticationSamlSessionOperOperSpList(ret.DtAamAuthenticationSamlSessionOper.Oper.SpList),
			"name":    ret.DtAamAuthenticationSamlSessionOper.Oper.Name,
		},
	}
}

func setSliceAamAuthenticationSamlSessionOperOperSpList(d []edpt.AamAuthenticationSamlSessionOperOperSpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sp_id"] = item.SpId
		in["session_count"] = item.SessionCount
		in["session_list"] = setSliceAamAuthenticationSamlSessionOperOperSpListSessionList(item.SessionList)
		result = append(result, in)
	}
	return result
}

func setSliceAamAuthenticationSamlSessionOperOperSpListSessionList(d []edpt.AamAuthenticationSamlSessionOperOperSpListSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["nameid"] = item.Nameid
		in["client_addr"] = item.ClientAddr
		in["id_provider"] = item.IdProvider
		in["auth_instant"] = item.AuthInstant
		in["expire_time"] = item.ExpireTime
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationSamlSessionOperOper(d []interface{}) edpt.AamAuthenticationSamlSessionOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlSessionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SpList = getSliceAamAuthenticationSamlSessionOperOperSpList(in["sp_list"].([]interface{}))
		ret.Name = in["name"].(string)
	}
	return ret
}

func getSliceAamAuthenticationSamlSessionOperOperSpList(d []interface{}) []edpt.AamAuthenticationSamlSessionOperOperSpList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlSessionOperOperSpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlSessionOperOperSpList
		oi.SpId = in["sp_id"].(string)
		oi.SessionCount = in["session_count"].(int)
		oi.SessionList = getSliceAamAuthenticationSamlSessionOperOperSpListSessionList(in["session_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationSamlSessionOperOperSpListSessionList(d []interface{}) []edpt.AamAuthenticationSamlSessionOperOperSpListSessionList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlSessionOperOperSpListSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlSessionOperOperSpListSessionList
		oi.Nameid = in["nameid"].(string)
		oi.ClientAddr = in["client_addr"].(string)
		oi.IdProvider = in["id_provider"].(string)
		oi.AuthInstant = in["auth_instant"].(string)
		oi.ExpireTime = in["expire_time"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlSessionOper(d *schema.ResourceData) edpt.AamAuthenticationSamlSessionOper {
	var ret edpt.AamAuthenticationSamlSessionOper

	ret.Oper = getObjectAamAuthenticationSamlSessionOperOper(d.Get("oper").([]interface{}))
	return ret
}
