package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementTelnetOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_enable_management_telnet_oper`: Operational Status for the object telnet\n\n__PLACEHOLDER__",
		ReadContext: resourceEnableManagementTelnetOperRead,

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
									"ipv4_acl": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_acl": {
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

func resourceEnableManagementTelnetOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementTelnetOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementTelnetOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		EnableManagementTelnetOperOper := setObjectEnableManagementTelnetOperOper(res)
		d.Set("oper", EnableManagementTelnetOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectEnableManagementTelnetOperOper(ret edpt.DataEnableManagementTelnetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_list": setSliceEnableManagementTelnetOperOperPortList(ret.DtEnableManagementTelnetOper.Oper.PortList),
		},
	}
}

func setSliceEnableManagementTelnetOperOperPortList(d []edpt.EnableManagementTelnetOperOperPortList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["management"] = item.Management
		in["ethernet"] = item.Ethernet
		in["ve"] = item.Ve
		in["tunnel"] = item.Tunnel
		in["action"] = item.Action
		in["ipv4_acl"] = item.Ipv4Acl
		in["ipv6_acl"] = item.Ipv6Acl
		result = append(result, in)
	}
	return result
}

func getObjectEnableManagementTelnetOperOper(d []interface{}) edpt.EnableManagementTelnetOperOper {

	count1 := len(d)
	var ret edpt.EnableManagementTelnetOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortList = getSliceEnableManagementTelnetOperOperPortList(in["port_list"].([]interface{}))
	}
	return ret
}

func getSliceEnableManagementTelnetOperOperPortList(d []interface{}) []edpt.EnableManagementTelnetOperOperPortList {

	count1 := len(d)
	ret := make([]edpt.EnableManagementTelnetOperOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementTelnetOperOperPortList
		oi.Management = in["management"].(int)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Tunnel = in["tunnel"].(int)
		oi.Action = in["action"].(string)
		oi.Ipv4Acl = in["ipv4_acl"].(string)
		oi.Ipv6Acl = in["ipv6_acl"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementTelnetOper(d *schema.ResourceData) edpt.EnableManagementTelnetOper {
	var ret edpt.EnableManagementTelnetOper

	ret.Oper = getObjectEnableManagementTelnetOperOper(d.Get("oper").([]interface{}))
	return ret
}
