package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSamlIdentityProviderOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_saml_identity_provider_oper`: Operational Status for the object identity-provider\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSamlIdentityProviderOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "SAML authentication identity provider name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"md": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cert": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"entity_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sso_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sso_location": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sso_binding": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"slo_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"slo_location": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"slo_binding": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"ars_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ars_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ars_location": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ars_binding": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"aqs_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aqs_location": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"aqs_binding": {
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

func resourceAamAuthenticationSamlIdentityProviderOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSamlIdentityProviderOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSamlIdentityProviderOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSamlIdentityProviderOperOper := setObjectAamAuthenticationSamlIdentityProviderOperOper(res)
		d.Set("oper", AamAuthenticationSamlIdentityProviderOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSamlIdentityProviderOperOper(ret edpt.DataAamAuthenticationSamlIdentityProviderOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"md":        ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.Md,
			"cert":      ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.Cert,
			"entity_id": ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.EntityId,
			"sso_list":  setSliceAamAuthenticationSamlIdentityProviderOperOperSsoList(ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.SsoList),
			"slo_list":  setSliceAamAuthenticationSamlIdentityProviderOperOperSloList(ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.SloList),
			"ars_list":  setSliceAamAuthenticationSamlIdentityProviderOperOperArsList(ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.ArsList),
			"aqs_list":  setSliceAamAuthenticationSamlIdentityProviderOperOperAqsList(ret.DtAamAuthenticationSamlIdentityProviderOper.Oper.AqsList),
		},
	}
}

func setSliceAamAuthenticationSamlIdentityProviderOperOperSsoList(d []edpt.AamAuthenticationSamlIdentityProviderOperOperSsoList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sso_location"] = item.SsoLocation
		in["sso_binding"] = item.SsoBinding
		result = append(result, in)
	}
	return result
}

func setSliceAamAuthenticationSamlIdentityProviderOperOperSloList(d []edpt.AamAuthenticationSamlIdentityProviderOperOperSloList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["slo_location"] = item.SloLocation
		in["slo_binding"] = item.SloBinding
		result = append(result, in)
	}
	return result
}

func setSliceAamAuthenticationSamlIdentityProviderOperOperArsList(d []edpt.AamAuthenticationSamlIdentityProviderOperOperArsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ars_index"] = item.ArsIndex
		in["ars_location"] = item.ArsLocation
		in["ars_binding"] = item.ArsBinding
		result = append(result, in)
	}
	return result
}

func setSliceAamAuthenticationSamlIdentityProviderOperOperAqsList(d []edpt.AamAuthenticationSamlIdentityProviderOperOperAqsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["aqs_location"] = item.AqsLocation
		in["aqs_binding"] = item.AqsBinding
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationSamlIdentityProviderOperOper(d []interface{}) edpt.AamAuthenticationSamlIdentityProviderOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationSamlIdentityProviderOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md = in["md"].(string)
		ret.Cert = in["cert"].(string)
		ret.EntityId = in["entity_id"].(string)
		ret.SsoList = getSliceAamAuthenticationSamlIdentityProviderOperOperSsoList(in["sso_list"].([]interface{}))
		ret.SloList = getSliceAamAuthenticationSamlIdentityProviderOperOperSloList(in["slo_list"].([]interface{}))
		ret.ArsList = getSliceAamAuthenticationSamlIdentityProviderOperOperArsList(in["ars_list"].([]interface{}))
		ret.AqsList = getSliceAamAuthenticationSamlIdentityProviderOperOperAqsList(in["aqs_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationSamlIdentityProviderOperOperSsoList(d []interface{}) []edpt.AamAuthenticationSamlIdentityProviderOperOperSsoList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlIdentityProviderOperOperSsoList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlIdentityProviderOperOperSsoList
		oi.SsoLocation = in["sso_location"].(string)
		oi.SsoBinding = in["sso_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationSamlIdentityProviderOperOperSloList(d []interface{}) []edpt.AamAuthenticationSamlIdentityProviderOperOperSloList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlIdentityProviderOperOperSloList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlIdentityProviderOperOperSloList
		oi.SloLocation = in["slo_location"].(string)
		oi.SloBinding = in["slo_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationSamlIdentityProviderOperOperArsList(d []interface{}) []edpt.AamAuthenticationSamlIdentityProviderOperOperArsList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlIdentityProviderOperOperArsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlIdentityProviderOperOperArsList
		oi.ArsIndex = in["ars_index"].(int)
		oi.ArsLocation = in["ars_location"].(string)
		oi.ArsBinding = in["ars_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationSamlIdentityProviderOperOperAqsList(d []interface{}) []edpt.AamAuthenticationSamlIdentityProviderOperOperAqsList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSamlIdentityProviderOperOperAqsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSamlIdentityProviderOperOperAqsList
		oi.AqsLocation = in["aqs_location"].(string)
		oi.AqsBinding = in["aqs_binding"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSamlIdentityProviderOper(d *schema.ResourceData) edpt.AamAuthenticationSamlIdentityProviderOper {
	var ret edpt.AamAuthenticationSamlIdentityProviderOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectAamAuthenticationSamlIdentityProviderOperOper(d.Get("oper").([]interface{}))
	return ret
}
