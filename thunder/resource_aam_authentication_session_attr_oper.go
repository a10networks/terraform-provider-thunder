package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationSessionAttrOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_session_attr_oper`: Operational Status for the object session-attr\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationSessionAttrOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attr_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"attr_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"attr_value": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"sid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"partition": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceAamAuthenticationSessionAttrOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationSessionAttrOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationSessionAttrOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationSessionAttrOperOper := setObjectAamAuthenticationSessionAttrOperOper(res)
		d.Set("oper", AamAuthenticationSessionAttrOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationSessionAttrOperOper(ret edpt.DataAamAuthenticationSessionAttrOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"attr_list": setSliceAamAuthenticationSessionAttrOperOperAttrList(ret.DtAamAuthenticationSessionAttrOper.Oper.AttrList),
			"sid":       ret.DtAamAuthenticationSessionAttrOper.Oper.Sid,
			"partition": ret.DtAamAuthenticationSessionAttrOper.Oper.Partition,
		},
	}
}

func setSliceAamAuthenticationSessionAttrOperOperAttrList(d []edpt.AamAuthenticationSessionAttrOperOperAttrList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["attr_name"] = item.AttrName
		in["attr_type"] = item.AttrType
		in["attr_value"] = item.AttrValue
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationSessionAttrOperOper(d []interface{}) edpt.AamAuthenticationSessionAttrOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationSessionAttrOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AttrList = getSliceAamAuthenticationSessionAttrOperOperAttrList(in["attr_list"].([]interface{}))
		ret.Sid = in["sid"].(int)
		ret.Partition = in["partition"].(string)
	}
	return ret
}

func getSliceAamAuthenticationSessionAttrOperOperAttrList(d []interface{}) []edpt.AamAuthenticationSessionAttrOperOperAttrList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationSessionAttrOperOperAttrList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationSessionAttrOperOperAttrList
		oi.AttrName = in["attr_name"].(string)
		oi.AttrType = in["attr_type"].(string)
		oi.AttrValue = in["attr_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationSessionAttrOper(d *schema.ResourceData) edpt.AamAuthenticationSessionAttrOper {
	var ret edpt.AamAuthenticationSessionAttrOper

	ret.Oper = getObjectAamAuthenticationSessionAttrOperOper(d.Get("oper").([]interface{}))
	return ret
}
