package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementAclIpv6Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_enable_management_acl_ipv6_oper`: Operational Status for the object acl-ipv6\n\n__PLACEHOLDER__",
		ReadContext: resourceEnableManagementAclIpv6OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"management": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"action": {
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

func resourceEnableManagementAclIpv6OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementAclIpv6OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementAclIpv6Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		EnableManagementAclIpv6OperOper := setObjectEnableManagementAclIpv6OperOper(res)
		d.Set("oper", EnableManagementAclIpv6OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectEnableManagementAclIpv6OperOper(ret edpt.DataEnableManagementAclIpv6Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_list": setSliceEnableManagementAclIpv6OperOperPortList(ret.DtEnableManagementAclIpv6Oper.Oper.PortList),
		},
	}
}

func setSliceEnableManagementAclIpv6OperOperPortList(d []edpt.EnableManagementAclIpv6OperOperPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["management"] = item.Management
		in["ethernet"] = item.Ethernet
		in["ve"] = item.Ve
		in["tunnel"] = item.Tunnel
		in["action"] = item.Action
		result = append(result, in)
	}
	return result
}

func getObjectEnableManagementAclIpv6OperOper(d []interface{}) edpt.EnableManagementAclIpv6OperOper {

	count1 := len(d)
	var ret edpt.EnableManagementAclIpv6OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortList = getSliceEnableManagementAclIpv6OperOperPortList(in["port_list"].([]interface{}))
	}
	return ret
}

func getSliceEnableManagementAclIpv6OperOperPortList(d []interface{}) []edpt.EnableManagementAclIpv6OperOperPortList {

	count1 := len(d)
	ret := make([]edpt.EnableManagementAclIpv6OperOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementAclIpv6OperOperPortList
		oi.Management = in["management"].(int)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementAclIpv6Oper(d *schema.ResourceData) edpt.EnableManagementAclIpv6Oper {
	var ret edpt.EnableManagementAclIpv6Oper

	ret.Oper = getObjectEnableManagementAclIpv6OperOper(d.Get("oper").([]interface{}))
	return ret
}
