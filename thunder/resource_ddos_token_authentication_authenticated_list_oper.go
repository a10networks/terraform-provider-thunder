package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTokenAuthenticationAuthenticatedListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_token_authentication_authenticated_list_oper`: Operational Status for the object authenticated-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosTokenAuthenticationAuthenticatedListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"player_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_ip_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"dst_ip_str": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"src_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dst_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"token": {
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

func resourceDdosTokenAuthenticationAuthenticatedListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTokenAuthenticationAuthenticatedListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTokenAuthenticationAuthenticatedListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosTokenAuthenticationAuthenticatedListOperOper := setObjectDdosTokenAuthenticationAuthenticatedListOperOper(res)
		d.Set("oper", DdosTokenAuthenticationAuthenticatedListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosTokenAuthenticationAuthenticatedListOperOper(ret edpt.DataDdosTokenAuthenticationAuthenticatedListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"player_list": setSliceDdosTokenAuthenticationAuthenticatedListOperOperPlayerList(ret.DtDdosTokenAuthenticationAuthenticatedListOper.Oper.PlayerList),
		},
	}
}

func setSliceDdosTokenAuthenticationAuthenticatedListOperOperPlayerList(d []edpt.DdosTokenAuthenticationAuthenticatedListOperOperPlayerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_ip_str"] = item.SrcIpStr
		in["dst_ip_str"] = item.DstIpStr
		in["src_port"] = item.SrcPort
		in["dst_port"] = item.DstPort
		in["token"] = item.Token
		result = append(result, in)
	}
	return result
}

func getObjectDdosTokenAuthenticationAuthenticatedListOperOper(d []interface{}) edpt.DdosTokenAuthenticationAuthenticatedListOperOper {

	count1 := len(d)
	var ret edpt.DdosTokenAuthenticationAuthenticatedListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PlayerList = getSliceDdosTokenAuthenticationAuthenticatedListOperOperPlayerList(in["player_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosTokenAuthenticationAuthenticatedListOperOperPlayerList(d []interface{}) []edpt.DdosTokenAuthenticationAuthenticatedListOperOperPlayerList {

	count1 := len(d)
	ret := make([]edpt.DdosTokenAuthenticationAuthenticatedListOperOperPlayerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTokenAuthenticationAuthenticatedListOperOperPlayerList
		oi.SrcIpStr = in["src_ip_str"].(string)
		oi.DstIpStr = in["dst_ip_str"].(string)
		oi.SrcPort = in["src_port"].(int)
		oi.DstPort = in["dst_port"].(int)
		oi.Token = in["token"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosTokenAuthenticationAuthenticatedListOper(d *schema.ResourceData) edpt.DdosTokenAuthenticationAuthenticatedListOper {
	var ret edpt.DdosTokenAuthenticationAuthenticatedListOper

	ret.Oper = getObjectDdosTokenAuthenticationAuthenticatedListOperOper(d.Get("oper").([]interface{}))
	return ret
}
