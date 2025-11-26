package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsServerOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_tacacs_server_oper`: Operational Status for the object tacacs-server\n\n__PLACEHOLDER__",
		ReadContext: resourceTacacsServerOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tacacs_server_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_open": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_aborts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_errors": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_failconn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_rev": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"socket_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"monitor_oper": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"con_fail_attempts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_fail_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_fail_auth": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_available": {
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

func resourceTacacsServerOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTacacsServerOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTacacsServerOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TacacsServerOperOper := setObjectTacacsServerOperOper(res)
		d.Set("oper", TacacsServerOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTacacsServerOperOper(ret edpt.DataTacacsServerOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tacacs_server_list": setSliceTacacsServerOperOperTacacsServerList(ret.DtTacacsServerOper.Oper.TacacsServerList),
		},
	}
}

func setSliceTacacsServerOperOperTacacsServerList(d []edpt.TacacsServerOperOperTacacsServerList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["port"] = item.Port
		in["socket_open"] = item.Socket_open
		in["socket_close"] = item.Socket_close
		in["socket_aborts"] = item.Socket_aborts
		in["socket_errors"] = item.Socket_errors
		in["socket_timeout"] = item.Socket_timeout
		in["socket_failconn"] = item.Socket_failconn
		in["socket_rev"] = item.Socket_rev
		in["socket_send"] = item.Socket_send
		in["monitor_oper"] = item.Monitor_oper
		in["con_fail_attempts"] = item.Con_fail_attempts
		in["total_fail_conn"] = item.Total_fail_conn
		in["total_fail_auth"] = item.Total_fail_auth
		in["last_available"] = item.Last_available
		result = append(result, in)
	}
	return result
}

func getObjectTacacsServerOperOper(d []interface{}) edpt.TacacsServerOperOper {

	count1 := len(d)
	var ret edpt.TacacsServerOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TacacsServerList = getSliceTacacsServerOperOperTacacsServerList(in["tacacs_server_list"].([]interface{}))
	}
	return ret
}

func getSliceTacacsServerOperOperTacacsServerList(d []interface{}) []edpt.TacacsServerOperOperTacacsServerList {

	count1 := len(d)
	ret := make([]edpt.TacacsServerOperOperTacacsServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TacacsServerOperOperTacacsServerList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.Socket_open = in["socket_open"].(int)
		oi.Socket_close = in["socket_close"].(int)
		oi.Socket_aborts = in["socket_aborts"].(int)
		oi.Socket_errors = in["socket_errors"].(int)
		oi.Socket_timeout = in["socket_timeout"].(int)
		oi.Socket_failconn = in["socket_failconn"].(int)
		oi.Socket_rev = in["socket_rev"].(int)
		oi.Socket_send = in["socket_send"].(int)
		oi.Monitor_oper = in["monitor_oper"].(int)
		oi.Con_fail_attempts = in["con_fail_attempts"].(int)
		oi.Total_fail_conn = in["total_fail_conn"].(int)
		oi.Total_fail_auth = in["total_fail_auth"].(int)
		oi.Last_available = in["last_available"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTacacsServerOper(d *schema.ResourceData) edpt.TacacsServerOper {
	var ret edpt.TacacsServerOper

	ret.Oper = getObjectTacacsServerOperOper(d.Get("oper").([]interface{}))
	return ret
}
