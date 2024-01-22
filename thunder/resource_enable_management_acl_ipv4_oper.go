package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementAclIpv4Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_enable_management_acl_ipv4_oper`: Operational Status for the object acl-ipv4\n\n__PLACEHOLDER__",
		ReadContext: resourceEnableManagementAclIpv4OperRead,

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

func resourceEnableManagementAclIpv4OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementAclIpv4OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementAclIpv4Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		EnableManagementAclIpv4OperOper := setObjectEnableManagementAclIpv4OperOper(res)
		d.Set("oper", EnableManagementAclIpv4OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectEnableManagementAclIpv4OperOper(ret edpt.DataEnableManagementAclIpv4Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_list": setSliceEnableManagementAclIpv4OperOperPortList(ret.DtEnableManagementAclIpv4Oper.Oper.PortList),
		},
	}
}

func setSliceEnableManagementAclIpv4OperOperPortList(d []edpt.EnableManagementAclIpv4OperOperPortList) []map[string]interface{} {
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

func getObjectEnableManagementAclIpv4OperOper(d []interface{}) edpt.EnableManagementAclIpv4OperOper {

	count1 := len(d)
	var ret edpt.EnableManagementAclIpv4OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortList = getSliceEnableManagementAclIpv4OperOperPortList(in["port_list"].([]interface{}))
	}
	return ret
}

func getSliceEnableManagementAclIpv4OperOperPortList(d []interface{}) []edpt.EnableManagementAclIpv4OperOperPortList {

	count1 := len(d)
	ret := make([]edpt.EnableManagementAclIpv4OperOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementAclIpv4OperOperPortList
		oi.Management = in["management"].(int)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementAclIpv4Oper(d *schema.ResourceData) edpt.EnableManagementAclIpv4Oper {
	var ret edpt.EnableManagementAclIpv4Oper

	ret.Oper = getObjectEnableManagementAclIpv4OperOper(d.Get("oper").([]interface{}))
	return ret
}
