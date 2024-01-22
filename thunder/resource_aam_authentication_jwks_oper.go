package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationJwksOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_jwks_oper`: Operational Status for the object jwks\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationJwksOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"jwk_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"jwk_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"jwk_size": {
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

func resourceAamAuthenticationJwksOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationJwksOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationJwksOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationJwksOperOper := setObjectAamAuthenticationJwksOperOper(res)
		d.Set("oper", AamAuthenticationJwksOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationJwksOperOper(ret edpt.DataAamAuthenticationJwksOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"jwk_list": setSliceAamAuthenticationJwksOperOperJwkList(ret.DtAamAuthenticationJwksOper.Oper.JwkList),
		},
	}
}

func setSliceAamAuthenticationJwksOperOperJwkList(d []edpt.AamAuthenticationJwksOperOperJwkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["jwk_name"] = item.JwkName
		in["jwk_size"] = item.JwkSize
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationJwksOperOper(d []interface{}) edpt.AamAuthenticationJwksOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationJwksOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.JwkList = getSliceAamAuthenticationJwksOperOperJwkList(in["jwk_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationJwksOperOperJwkList(d []interface{}) []edpt.AamAuthenticationJwksOperOperJwkList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationJwksOperOperJwkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationJwksOperOperJwkList
		oi.JwkName = in["jwk_name"].(string)
		oi.JwkSize = in["jwk_size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationJwksOper(d *schema.ResourceData) edpt.AamAuthenticationJwksOper {
	var ret edpt.AamAuthenticationJwksOper

	ret.Oper = getObjectAamAuthenticationJwksOperOper(d.Get("oper").([]interface{}))
	return ret
}
