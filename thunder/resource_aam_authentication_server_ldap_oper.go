package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerLdapOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_authentication_server_ldap_oper`: Operational Status for the object ldap\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAuthenticationServerLdapOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ldaps_server_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ldap_uri": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ldaps_idle_conn_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ldaps_idle_conn_fd_list": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ldaps_inuse_conn_num": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ldaps_inuse_conn_fd_list": {
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

func resourceAamAuthenticationServerLdapOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerLdapOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerLdapOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAuthenticationServerLdapOperOper := setObjectAamAuthenticationServerLdapOperOper(res)
		d.Set("oper", AamAuthenticationServerLdapOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAuthenticationServerLdapOperOper(ret edpt.DataAamAuthenticationServerLdapOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ldaps_server_list": setSliceAamAuthenticationServerLdapOperOperLdapsServerList(ret.DtAamAuthenticationServerLdapOper.Oper.LdapsServerList),
		},
	}
}

func setSliceAamAuthenticationServerLdapOperOperLdapsServerList(d []edpt.AamAuthenticationServerLdapOperOperLdapsServerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ldap_uri"] = item.LdapUri
		in["ldaps_idle_conn_num"] = item.LdapsIdleConnNum
		in["ldaps_idle_conn_fd_list"] = item.LdapsIdleConnFdList
		in["ldaps_inuse_conn_num"] = item.LdapsInuseConnNum
		in["ldaps_inuse_conn_fd_list"] = item.LdapsInuseConnFdList
		result = append(result, in)
	}
	return result
}

func getObjectAamAuthenticationServerLdapOperOper(d []interface{}) edpt.AamAuthenticationServerLdapOperOper {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerLdapOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LdapsServerList = getSliceAamAuthenticationServerLdapOperOperLdapsServerList(in["ldaps_server_list"].([]interface{}))
	}
	return ret
}

func getSliceAamAuthenticationServerLdapOperOperLdapsServerList(d []interface{}) []edpt.AamAuthenticationServerLdapOperOperLdapsServerList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerLdapOperOperLdapsServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerLdapOperOperLdapsServerList
		oi.LdapUri = in["ldap_uri"].(string)
		oi.LdapsIdleConnNum = in["ldaps_idle_conn_num"].(int)
		oi.LdapsIdleConnFdList = in["ldaps_idle_conn_fd_list"].(string)
		oi.LdapsInuseConnNum = in["ldaps_inuse_conn_num"].(int)
		oi.LdapsInuseConnFdList = in["ldaps_inuse_conn_fd_list"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerLdapOper(d *schema.ResourceData) edpt.AamAuthenticationServerLdapOper {
	var ret edpt.AamAuthenticationServerLdapOper

	ret.Oper = getObjectAamAuthenticationServerLdapOperOper(d.Get("oper").([]interface{}))
	return ret
}
