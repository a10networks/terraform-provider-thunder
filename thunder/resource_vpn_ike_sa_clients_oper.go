package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIkeSaClientsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vpn_ike_sa_clients_oper`: Operational Status for the object ike-sa-clients\n\n__PLACEHOLDER__",
		ReadContext: resourceVpnIkeSaClientsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_sa_clients_local_ip": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ike_sa_clients_remote_gw": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ike_sa_clients_remote_gw_ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_remote_gw_remote_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_remote_gw_user_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_remote_gw_idle_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_remote_gw_session_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ike_sa_clients_remote_gw_bytes": {
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

func resourceVpnIkeSaClientsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIkeSaClientsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIkeSaClientsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VpnIkeSaClientsOperOper := setObjectVpnIkeSaClientsOperOper(res)
		d.Set("oper", VpnIkeSaClientsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVpnIkeSaClientsOperOper(ret edpt.DataVpnIkeSaClientsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"name":                     ret.DtVpnIkeSaClientsOper.Oper.Name,
			"ike_sa_clients_local_ip":  ret.DtVpnIkeSaClientsOper.Oper.IkeSaClientsLocalIp,
			"ike_sa_clients_remote_gw": setSliceVpnIkeSaClientsOperOperIkeSaClientsRemoteGw(ret.DtVpnIkeSaClientsOper.Oper.IkeSaClientsRemoteGw),
		},
	}
}

func setSliceVpnIkeSaClientsOperOperIkeSaClientsRemoteGw(d []edpt.VpnIkeSaClientsOperOperIkeSaClientsRemoteGw) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ike_sa_clients_remote_gw_ip"] = item.IkeSaClientsRemoteGwIp
		in["ike_sa_clients_remote_gw_remote_id"] = item.IkeSaClientsRemoteGwRemoteId
		in["ike_sa_clients_remote_gw_user_id"] = item.IkeSaClientsRemoteGwUserId
		in["ike_sa_clients_remote_gw_idle_time"] = item.IkeSaClientsRemoteGwIdleTime
		in["ike_sa_clients_remote_gw_session_time"] = item.IkeSaClientsRemoteGwSessionTime
		in["ike_sa_clients_remote_gw_bytes"] = item.IkeSaClientsRemoteGwBytes
		result = append(result, in)
	}
	return result
}

func getObjectVpnIkeSaClientsOperOper(d []interface{}) edpt.VpnIkeSaClientsOperOper {

	count1 := len(d)
	var ret edpt.VpnIkeSaClientsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Name = in["name"].(string)
		ret.IkeSaClientsLocalIp = in["ike_sa_clients_local_ip"].(string)
		ret.IkeSaClientsRemoteGw = getSliceVpnIkeSaClientsOperOperIkeSaClientsRemoteGw(in["ike_sa_clients_remote_gw"].([]interface{}))
	}
	return ret
}

func getSliceVpnIkeSaClientsOperOperIkeSaClientsRemoteGw(d []interface{}) []edpt.VpnIkeSaClientsOperOperIkeSaClientsRemoteGw {

	count1 := len(d)
	ret := make([]edpt.VpnIkeSaClientsOperOperIkeSaClientsRemoteGw, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIkeSaClientsOperOperIkeSaClientsRemoteGw
		oi.IkeSaClientsRemoteGwIp = in["ike_sa_clients_remote_gw_ip"].(string)
		oi.IkeSaClientsRemoteGwRemoteId = in["ike_sa_clients_remote_gw_remote_id"].(string)
		oi.IkeSaClientsRemoteGwUserId = in["ike_sa_clients_remote_gw_user_id"].(string)
		oi.IkeSaClientsRemoteGwIdleTime = in["ike_sa_clients_remote_gw_idle_time"].(string)
		oi.IkeSaClientsRemoteGwSessionTime = in["ike_sa_clients_remote_gw_session_time"].(string)
		oi.IkeSaClientsRemoteGwBytes = in["ike_sa_clients_remote_gw_bytes"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVpnIkeSaClientsOper(d *schema.ResourceData) edpt.VpnIkeSaClientsOper {
	var ret edpt.VpnIkeSaClientsOper

	ret.Oper = getObjectVpnIkeSaClientsOperOper(d.Get("oper").([]interface{}))
	return ret
}
