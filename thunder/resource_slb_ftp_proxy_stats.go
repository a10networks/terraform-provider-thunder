package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbFtpProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_ftp_proxy_stats`: Statistics for the object ftp-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbFtpProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current proxy conns",
						},
						"total": {
							Type: schema.TypeInt, Optional: true, Description: "Total proxy conns",
						},
						"svrsel_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection failure",
						},
						"no_route": {
							Type: schema.TypeInt, Optional: true, Description: "no route failure",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "source nat failure",
						},
						"feat": {
							Type: schema.TypeInt, Optional: true, Description: "feat packet",
						},
						"cc": {
							Type: schema.TypeInt, Optional: true, Description: "clear ctrl port packet",
						},
						"data_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "data ssl force",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "line too long",
						},
						"line_mem_freed": {
							Type: schema.TypeInt, Optional: true, Description: "request line freed",
						},
						"invalid_start_line": {
							Type: schema.TypeInt, Optional: true, Description: "invalid start line",
						},
						"auth_tls": {
							Type: schema.TypeInt, Optional: true, Description: "auth tls cmd",
						},
						"prot": {
							Type: schema.TypeInt, Optional: true, Description: "prot cmd",
						},
						"pbsz": {
							Type: schema.TypeInt, Optional: true, Description: "pbsz cmd",
						},
						"pasv": {
							Type: schema.TypeInt, Optional: true, Description: "pasv cmd",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "port cmd",
						},
						"request_dont_care": {
							Type: schema.TypeInt, Optional: true, Description: "other cmd",
						},
						"client_auth_tls": {
							Type: schema.TypeInt, Optional: true, Description: "client auth tls",
						},
						"cant_find_pasv": {
							Type: schema.TypeInt, Optional: true, Description: "cant find pasv",
						},
						"pasv_addr_ne_server": {
							Type: schema.TypeInt, Optional: true, Description: "psv addr not equal to svr",
						},
						"smp_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "smp create fail",
						},
						"data_server_conn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "data svr conn fail",
						},
						"data_send_fail": {
							Type: schema.TypeInt, Optional: true, Description: "data send fail",
						},
						"epsv": {
							Type: schema.TypeInt, Optional: true, Description: "epsv command",
						},
						"cant_find_epsv": {
							Type: schema.TypeInt, Optional: true, Description: "cant find epsv",
						},
						"data_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current Data Proxy",
						},
						"data_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Data Proxy",
						},
						"auth_unsupported": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported auth",
						},
						"adat": {
							Type: schema.TypeInt, Optional: true, Description: "adat cmd",
						},
						"unsupported_pbsz_value": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported PBSZ",
						},
						"unsupported_prot_value": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported PROT",
						},
						"unsupported_command": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported cmd",
						},
						"control_to_clear": {
							Type: schema.TypeInt, Optional: true, Description: "Control chn clear txt",
						},
						"control_to_ssl": {
							Type: schema.TypeInt, Optional: true, Description: "Control chn ssl",
						},
						"bad_sequence": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Sequence",
						},
						"rsv_persist_conn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel Persist fail",
						},
						"smp_v6_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel SMPv6 fail",
						},
						"smp_v4_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel SMPv4 fail",
						},
						"insert_tuple_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Serv Sel insert tuple fail",
						},
						"cl_est_err": {
							Type: schema.TypeInt, Optional: true, Description: "Client EST state erro",
						},
						"ser_connecting_err": {
							Type: schema.TypeInt, Optional: true, Description: "Serv CTNG state error",
						},
						"server_response_err": {
							Type: schema.TypeInt, Optional: true, Description: "Serv RESP state error",
						},
						"cl_request_err": {
							Type: schema.TypeInt, Optional: true, Description: "Client RQ state error",
						},
						"data_conn_start_err": {
							Type: schema.TypeInt, Optional: true, Description: "Data Start state error",
						},
						"data_serv_connecting_err": {
							Type: schema.TypeInt, Optional: true, Description: "Data Serv CTNG error",
						},
						"data_serv_connected_err": {
							Type: schema.TypeInt, Optional: true, Description: "Data Serv CTED error",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "Total FTP Request",
						},
						"auth_req": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Request",
						},
						"auth_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Success",
						},
						"auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Failure",
						},
						"fwd_to_internet": {
							Type: schema.TypeInt, Optional: true, Description: "Forward to Internet",
						},
						"fwd_to_sg": {
							Type: schema.TypeInt, Optional: true, Description: "Total Forward to Service-group",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Description: "Total FTP Drop",
						},
						"ds_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Host Domain Name is resolved",
						},
						"ds_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Host Domain Name isn't resolved",
						},
						"open": {
							Type: schema.TypeInt, Optional: true, Description: "open cmd",
						},
						"site": {
							Type: schema.TypeInt, Optional: true, Description: "site cmd",
						},
						"user": {
							Type: schema.TypeInt, Optional: true, Description: "user cmd",
						},
						"pass": {
							Type: schema.TypeInt, Optional: true, Description: "pass cmd",
						},
						"quit": {
							Type: schema.TypeInt, Optional: true, Description: "quit cmd",
						},
						"eprt": {
							Type: schema.TypeInt, Optional: true, Description: "eprt cmd",
						},
						"cant_find_port": {
							Type: schema.TypeInt, Optional: true, Description: "cant find port",
						},
						"cant_find_eprt": {
							Type: schema.TypeInt, Optional: true, Description: "cant find eprt",
						},
					},
				},
			},
		},
	}
}

func resourceSlbFtpProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbFtpProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbFtpProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbFtpProxyStatsStats := setObjectSlbFtpProxyStatsStats(res)
		d.Set("stats", SlbFtpProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbFtpProxyStatsStats(ret edpt.DataSlbFtpProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr":                     ret.DtSlbFtpProxyStats.Stats.Curr,
			"total":                    ret.DtSlbFtpProxyStats.Stats.Total,
			"svrsel_fail":              ret.DtSlbFtpProxyStats.Stats.Svrsel_fail,
			"no_route":                 ret.DtSlbFtpProxyStats.Stats.No_route,
			"snat_fail":                ret.DtSlbFtpProxyStats.Stats.Snat_fail,
			"feat":                     ret.DtSlbFtpProxyStats.Stats.Feat,
			"cc":                       ret.DtSlbFtpProxyStats.Stats.Cc,
			"data_ssl":                 ret.DtSlbFtpProxyStats.Stats.Data_ssl,
			"line_too_long":            ret.DtSlbFtpProxyStats.Stats.Line_too_long,
			"line_mem_freed":           ret.DtSlbFtpProxyStats.Stats.Line_mem_freed,
			"invalid_start_line":       ret.DtSlbFtpProxyStats.Stats.Invalid_start_line,
			"auth_tls":                 ret.DtSlbFtpProxyStats.Stats.Auth_tls,
			"prot":                     ret.DtSlbFtpProxyStats.Stats.Prot,
			"pbsz":                     ret.DtSlbFtpProxyStats.Stats.Pbsz,
			"pasv":                     ret.DtSlbFtpProxyStats.Stats.Pasv,
			"port":                     ret.DtSlbFtpProxyStats.Stats.Port,
			"request_dont_care":        ret.DtSlbFtpProxyStats.Stats.Request_dont_care,
			"client_auth_tls":          ret.DtSlbFtpProxyStats.Stats.Client_auth_tls,
			"cant_find_pasv":           ret.DtSlbFtpProxyStats.Stats.Cant_find_pasv,
			"pasv_addr_ne_server":      ret.DtSlbFtpProxyStats.Stats.Pasv_addr_ne_server,
			"smp_create_fail":          ret.DtSlbFtpProxyStats.Stats.Smp_create_fail,
			"data_server_conn_fail":    ret.DtSlbFtpProxyStats.Stats.Data_server_conn_fail,
			"data_send_fail":           ret.DtSlbFtpProxyStats.Stats.Data_send_fail,
			"epsv":                     ret.DtSlbFtpProxyStats.Stats.Epsv,
			"cant_find_epsv":           ret.DtSlbFtpProxyStats.Stats.Cant_find_epsv,
			"data_curr":                ret.DtSlbFtpProxyStats.Stats.Data_curr,
			"data_total":               ret.DtSlbFtpProxyStats.Stats.Data_total,
			"auth_unsupported":         ret.DtSlbFtpProxyStats.Stats.Auth_unsupported,
			"adat":                     ret.DtSlbFtpProxyStats.Stats.Adat,
			"unsupported_pbsz_value":   ret.DtSlbFtpProxyStats.Stats.Unsupported_pbsz_value,
			"unsupported_prot_value":   ret.DtSlbFtpProxyStats.Stats.Unsupported_prot_value,
			"unsupported_command":      ret.DtSlbFtpProxyStats.Stats.Unsupported_command,
			"control_to_clear":         ret.DtSlbFtpProxyStats.Stats.Control_to_clear,
			"control_to_ssl":           ret.DtSlbFtpProxyStats.Stats.Control_to_ssl,
			"bad_sequence":             ret.DtSlbFtpProxyStats.Stats.Bad_sequence,
			"rsv_persist_conn_fail":    ret.DtSlbFtpProxyStats.Stats.Rsv_persist_conn_fail,
			"smp_v6_fail":              ret.DtSlbFtpProxyStats.Stats.Smp_v6_fail,
			"smp_v4_fail":              ret.DtSlbFtpProxyStats.Stats.Smp_v4_fail,
			"insert_tuple_fail":        ret.DtSlbFtpProxyStats.Stats.Insert_tuple_fail,
			"cl_est_err":               ret.DtSlbFtpProxyStats.Stats.Cl_est_err,
			"ser_connecting_err":       ret.DtSlbFtpProxyStats.Stats.Ser_connecting_err,
			"server_response_err":      ret.DtSlbFtpProxyStats.Stats.Server_response_err,
			"cl_request_err":           ret.DtSlbFtpProxyStats.Stats.Cl_request_err,
			"data_conn_start_err":      ret.DtSlbFtpProxyStats.Stats.Data_conn_start_err,
			"data_serv_connecting_err": ret.DtSlbFtpProxyStats.Stats.Data_serv_connecting_err,
			"data_serv_connected_err":  ret.DtSlbFtpProxyStats.Stats.Data_serv_connected_err,
			"request":                  ret.DtSlbFtpProxyStats.Stats.Request,
			"auth_req":                 ret.DtSlbFtpProxyStats.Stats.Auth_req,
			"auth_succ":                ret.DtSlbFtpProxyStats.Stats.Auth_succ,
			"auth_fail":                ret.DtSlbFtpProxyStats.Stats.Auth_fail,
			"fwd_to_internet":          ret.DtSlbFtpProxyStats.Stats.Fwd_to_internet,
			"fwd_to_sg":                ret.DtSlbFtpProxyStats.Stats.Fwd_to_sg,
			"drop":                     ret.DtSlbFtpProxyStats.Stats.Drop,
			"ds_succ":                  ret.DtSlbFtpProxyStats.Stats.Ds_succ,
			"ds_fail":                  ret.DtSlbFtpProxyStats.Stats.Ds_fail,
			"open":                     ret.DtSlbFtpProxyStats.Stats.Open,
			"site":                     ret.DtSlbFtpProxyStats.Stats.Site,
			"user":                     ret.DtSlbFtpProxyStats.Stats.User,
			"pass":                     ret.DtSlbFtpProxyStats.Stats.Pass,
			"quit":                     ret.DtSlbFtpProxyStats.Stats.Quit,
			"eprt":                     ret.DtSlbFtpProxyStats.Stats.Eprt,
			"cant_find_port":           ret.DtSlbFtpProxyStats.Stats.Cant_find_port,
			"cant_find_eprt":           ret.DtSlbFtpProxyStats.Stats.Cant_find_eprt,
		},
	}
}

func getObjectSlbFtpProxyStatsStats(d []interface{}) edpt.SlbFtpProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbFtpProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr = in["curr"].(int)
		ret.Total = in["total"].(int)
		ret.Svrsel_fail = in["svrsel_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Feat = in["feat"].(int)
		ret.Cc = in["cc"].(int)
		ret.Data_ssl = in["data_ssl"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_mem_freed = in["line_mem_freed"].(int)
		ret.Invalid_start_line = in["invalid_start_line"].(int)
		ret.Auth_tls = in["auth_tls"].(int)
		ret.Prot = in["prot"].(int)
		ret.Pbsz = in["pbsz"].(int)
		ret.Pasv = in["pasv"].(int)
		ret.Port = in["port"].(int)
		ret.Request_dont_care = in["request_dont_care"].(int)
		ret.Client_auth_tls = in["client_auth_tls"].(int)
		ret.Cant_find_pasv = in["cant_find_pasv"].(int)
		ret.Pasv_addr_ne_server = in["pasv_addr_ne_server"].(int)
		ret.Smp_create_fail = in["smp_create_fail"].(int)
		ret.Data_server_conn_fail = in["data_server_conn_fail"].(int)
		ret.Data_send_fail = in["data_send_fail"].(int)
		ret.Epsv = in["epsv"].(int)
		ret.Cant_find_epsv = in["cant_find_epsv"].(int)
		ret.Data_curr = in["data_curr"].(int)
		ret.Data_total = in["data_total"].(int)
		ret.Auth_unsupported = in["auth_unsupported"].(int)
		ret.Adat = in["adat"].(int)
		ret.Unsupported_pbsz_value = in["unsupported_pbsz_value"].(int)
		ret.Unsupported_prot_value = in["unsupported_prot_value"].(int)
		ret.Unsupported_command = in["unsupported_command"].(int)
		ret.Control_to_clear = in["control_to_clear"].(int)
		ret.Control_to_ssl = in["control_to_ssl"].(int)
		ret.Bad_sequence = in["bad_sequence"].(int)
		ret.Rsv_persist_conn_fail = in["rsv_persist_conn_fail"].(int)
		ret.Smp_v6_fail = in["smp_v6_fail"].(int)
		ret.Smp_v4_fail = in["smp_v4_fail"].(int)
		ret.Insert_tuple_fail = in["insert_tuple_fail"].(int)
		ret.Cl_est_err = in["cl_est_err"].(int)
		ret.Ser_connecting_err = in["ser_connecting_err"].(int)
		ret.Server_response_err = in["server_response_err"].(int)
		ret.Cl_request_err = in["cl_request_err"].(int)
		ret.Data_conn_start_err = in["data_conn_start_err"].(int)
		ret.Data_serv_connecting_err = in["data_serv_connecting_err"].(int)
		ret.Data_serv_connected_err = in["data_serv_connected_err"].(int)
		ret.Request = in["request"].(int)
		ret.Auth_req = in["auth_req"].(int)
		ret.Auth_succ = in["auth_succ"].(int)
		ret.Auth_fail = in["auth_fail"].(int)
		ret.Fwd_to_internet = in["fwd_to_internet"].(int)
		ret.Fwd_to_sg = in["fwd_to_sg"].(int)
		ret.Drop = in["drop"].(int)
		ret.Ds_succ = in["ds_succ"].(int)
		ret.Ds_fail = in["ds_fail"].(int)
		ret.Open = in["open"].(int)
		ret.Site = in["site"].(int)
		ret.User = in["user"].(int)
		ret.Pass = in["pass"].(int)
		ret.Quit = in["quit"].(int)
		ret.Eprt = in["eprt"].(int)
		ret.Cant_find_port = in["cant_find_port"].(int)
		ret.Cant_find_eprt = in["cant_find_eprt"].(int)
	}
	return ret
}

func dataToEndpointSlbFtpProxyStats(d *schema.ResourceData) edpt.SlbFtpProxyStats {
	var ret edpt.SlbFtpProxyStats

	ret.Stats = getObjectSlbFtpProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
