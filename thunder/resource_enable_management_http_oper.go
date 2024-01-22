package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementHttpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_enable_management_http_oper`: Operational Status for the object http\n\n__PLACEHOLDER__",
		ReadContext: resourceEnableManagementHttpOperRead,

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

func resourceEnableManagementHttpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementHttpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementHttpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		EnableManagementHttpOperOper := setObjectEnableManagementHttpOperOper(res)
		d.Set("oper", EnableManagementHttpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectEnableManagementHttpOperOper(ret edpt.DataEnableManagementHttpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"port_list": setSliceEnableManagementHttpOperOperPortList(ret.DtEnableManagementHttpOper.Oper.PortList),
		},
	}
}

func setSliceEnableManagementHttpOperOperPortList(d []edpt.EnableManagementHttpOperOperPortList) []map[string]interface{} {
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

func getObjectEnableManagementHttpOperOper(d []interface{}) edpt.EnableManagementHttpOperOper {

	count1 := len(d)
	var ret edpt.EnableManagementHttpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PortList = getSliceEnableManagementHttpOperOperPortList(in["port_list"].([]interface{}))
	}
	return ret
}

func getSliceEnableManagementHttpOperOperPortList(d []interface{}) []edpt.EnableManagementHttpOperOperPortList {

	count1 := len(d)
	ret := make([]edpt.EnableManagementHttpOperOperPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementHttpOperOperPortList
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

func dataToEndpointEnableManagementHttpOper(d *schema.ResourceData) edpt.EnableManagementHttpOper {
	var ret edpt.EnableManagementHttpOper

	ret.Oper = getObjectEnableManagementHttpOperOper(d.Get("oper").([]interface{}))
	return ret
}
