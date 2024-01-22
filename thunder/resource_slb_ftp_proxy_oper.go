package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFtpProxyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ftp_proxy_oper`: Operational Status for the object ftp-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFtpProxyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ftp_proxy_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_curr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"svrsel_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_route": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"feat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_tls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"prot": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pbsz": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"open": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"site": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"user": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pass": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"quit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_find_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"eprt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_find_eprt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_dont_care": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_auth_tls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pasv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_find_pasv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pasv_addr_ne_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smp_create_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_server_conn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_send_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"epsv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_find_epsv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_unsupported": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"adat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unsupported_pbsz_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unsupported_prot_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unsupported_command": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_to_clear": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_to_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_sequence": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rsv_persist_conn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smp_v6_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smp_v4_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insert_tuple_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cl_est_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ser_connecting_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_response_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cl_request_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_conn_start_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_serv_connecting_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_serv_connected_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_succ": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwd_to_internet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fwd_to_sg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ds_succ": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ds_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFtpProxyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpProxyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpProxyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFtpProxyOperOper := setObjectSlbFtpProxyOperOper(res)
		d.Set("oper", SlbFtpProxyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFtpProxyOperOper(ret edpt.DataSlbFtpProxyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ftp_proxy_cpu_list": setSliceSlbFtpProxyOperOperFtpProxyCpuList(ret.DtSlbFtpProxyOper.Oper.FtpProxyCpuList),
			"cpu_count":          ret.DtSlbFtpProxyOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbFtpProxyOperOperFtpProxyCpuList(d []edpt.SlbFtpProxyOperOperFtpProxyCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr"] = item.Curr
		in["total"] = item.Total
		in["data_curr"] = item.Data_curr
		in["data_total"] = item.Data_total
		in["request"] = item.Request
		in["svrsel_fail"] = item.Svrsel_fail
		in["no_route"] = item.No_route
		in["snat_fail"] = item.Snat_fail
		in["feat"] = item.Feat
		in["cc"] = item.Cc
		in["data_ssl"] = item.Data_ssl
		in["line_mem_freed"] = item.Line_mem_freed
		in["invalid_start_line"] = item.Invalid_start_line
		in["auth_tls"] = item.Auth_tls
		in["prot"] = item.Prot
		in["pbsz"] = item.Pbsz
		in["open"] = item.Open
		in["site"] = item.Site
		in["user"] = item.User
		in["pass"] = item.Pass
		in["quit"] = item.Quit
		in["port"] = item.Port
		in["cant_find_port"] = item.Cant_find_port
		in["eprt"] = item.Eprt
		in["cant_find_eprt"] = item.Cant_find_eprt
		in["request_dont_care"] = item.Request_dont_care
		in["line_too_long"] = item.Line_too_long
		in["client_auth_tls"] = item.Client_auth_tls
		in["pasv"] = item.Pasv
		in["cant_find_pasv"] = item.Cant_find_pasv
		in["pasv_addr_ne_server"] = item.Pasv_addr_ne_server
		in["smp_create_fail"] = item.Smp_create_fail
		in["data_server_conn_fail"] = item.Data_server_conn_fail
		in["data_send_fail"] = item.Data_send_fail
		in["epsv"] = item.Epsv
		in["cant_find_epsv"] = item.Cant_find_epsv
		in["auth_unsupported"] = item.Auth_unsupported
		in["adat"] = item.Adat
		in["unsupported_pbsz_value"] = item.Unsupported_pbsz_value
		in["unsupported_prot_value"] = item.Unsupported_prot_value
		in["unsupported_command"] = item.Unsupported_command
		in["control_to_clear"] = item.Control_to_clear
		in["control_to_ssl"] = item.Control_to_ssl
		in["bad_sequence"] = item.Bad_sequence
		in["rsv_persist_conn_fail"] = item.Rsv_persist_conn_fail
		in["smp_v6_fail"] = item.Smp_v6_fail
		in["smp_v4_fail"] = item.Smp_v4_fail
		in["insert_tuple_fail"] = item.Insert_tuple_fail
		in["cl_est_err"] = item.Cl_est_err
		in["ser_connecting_err"] = item.Ser_connecting_err
		in["server_response_err"] = item.Server_response_err
		in["cl_request_err"] = item.Cl_request_err
		in["data_conn_start_err"] = item.Data_conn_start_err
		in["data_serv_connecting_err"] = item.Data_serv_connecting_err
		in["data_serv_connected_err"] = item.Data_serv_connected_err
		in["auth_req"] = item.Auth_req
		in["auth_succ"] = item.Auth_succ
		in["auth_fail"] = item.Auth_fail
		in["fwd_to_internet"] = item.Fwd_to_internet
		in["fwd_to_sg"] = item.Fwd_to_sg
		in["drop"] = item.Drop
		in["ds_succ"] = item.Ds_succ
		in["ds_fail"] = item.Ds_fail
		result = append(result, in)
	}
	return result
}

func getObjectSlbFtpProxyOperOper(d []interface{}) edpt.SlbFtpProxyOperOper {

	count1 := len(d)
	var ret edpt.SlbFtpProxyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FtpProxyCpuList = getSliceSlbFtpProxyOperOperFtpProxyCpuList(in["ftp_proxy_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbFtpProxyOperOperFtpProxyCpuList(d []interface{}) []edpt.SlbFtpProxyOperOperFtpProxyCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbFtpProxyOperOperFtpProxyCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbFtpProxyOperOperFtpProxyCpuList
		oi.Curr = in["curr"].(int)
		oi.Total = in["total"].(int)
		oi.Data_curr = in["data_curr"].(int)
		oi.Data_total = in["data_total"].(int)
		oi.Request = in["request"].(int)
		oi.Svrsel_fail = in["svrsel_fail"].(int)
		oi.No_route = in["no_route"].(int)
		oi.Snat_fail = in["snat_fail"].(int)
		oi.Feat = in["feat"].(int)
		oi.Cc = in["cc"].(int)
		oi.Data_ssl = in["data_ssl"].(int)
		oi.Line_mem_freed = in["line_mem_freed"].(int)
		oi.Invalid_start_line = in["invalid_start_line"].(int)
		oi.Auth_tls = in["auth_tls"].(int)
		oi.Prot = in["prot"].(int)
		oi.Pbsz = in["pbsz"].(int)
		oi.Open = in["open"].(int)
		oi.Site = in["site"].(int)
		oi.User = in["user"].(int)
		oi.Pass = in["pass"].(int)
		oi.Quit = in["quit"].(int)
		oi.Port = in["port"].(int)
		oi.Cant_find_port = in["cant_find_port"].(int)
		oi.Eprt = in["eprt"].(int)
		oi.Cant_find_eprt = in["cant_find_eprt"].(int)
		oi.Request_dont_care = in["request_dont_care"].(int)
		oi.Line_too_long = in["line_too_long"].(int)
		oi.Client_auth_tls = in["client_auth_tls"].(int)
		oi.Pasv = in["pasv"].(int)
		oi.Cant_find_pasv = in["cant_find_pasv"].(int)
		oi.Pasv_addr_ne_server = in["pasv_addr_ne_server"].(int)
		oi.Smp_create_fail = in["smp_create_fail"].(int)
		oi.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		oi.Data_send_fail = in["data_send_fail"].(int)
		oi.Epsv = in["epsv"].(int)
		oi.Cant_find_epsv = in["cant_find_epsv"].(int)
		oi.Auth_unsupported = in["auth_unsupported"].(int)
		oi.Adat = in["adat"].(int)
		oi.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		oi.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		oi.Unsupported_command = in["unsupported_command"].(int)
		oi.Control_to_clear = in["control_to_clear"].(int)
		oi.Control_to_ssl = in["control_to_ssl"].(int)
		oi.Bad_sequence = in["bad_sequence"].(int)
		oi.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		oi.Smp_v6_fail = in["smp_v6_fail"].(int)
		oi.Smp_v4_fail = in["smp_v4_fail"].(int)
		oi.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		oi.Cl_est_err = in["cl_est_err"].(int)
		oi.Ser_connecting_err = in["ser_connecting_err"].(int)
		oi.Server_response_err = in["server_response_err"].(int)
		oi.Cl_request_err = in["cl_request_err"].(int)
		oi.Data_conn_start_err = in["data_conn_start_err"].(int)
		oi.Data_serv_connecting_err = in["data_serv_connecting_err"].(int)
		oi.Data_serv_connected_err = in["data_serv_connected_err"].(int)
		oi.Auth_req = in["auth_req"].(int)
		oi.Auth_succ = in["auth_succ"].(int)
		oi.Auth_fail = in["auth_fail"].(int)
		oi.Fwd_to_internet = in["fwd_to_internet"].(int)
		oi.Fwd_to_sg = in["fwd_to_sg"].(int)
		oi.Drop = in["drop"].(int)
		oi.Ds_succ = in["ds_succ"].(int)
		oi.Ds_fail = in["ds_fail"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbFtpProxyOper(d *schema.ResourceData) edpt.SlbFtpProxyOper {
	var ret edpt.SlbFtpProxyOper

	ret.Oper = getObjectSlbFtpProxyOperOper(d.Get("oper").([]interface{}))
	return ret
}
