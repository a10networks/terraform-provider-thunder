package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmtpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_smtp_stats`: Statistics for the object smtp\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSmtpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Current proxy conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total proxy conns",
						},
						"request": {
							Type: schema.TypeInt, Optional: true, Description: "SMTP requests",
						},
						"request_success": {
							Type: schema.TypeInt, Optional: true, Description: "SMTP requests (success)",
						},
						"no_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "No proxy error",
						},
						"client_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Client reset",
						},
						"server_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Server reset",
						},
						"no_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "No tuple error",
						},
						"parse_req_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse request failure",
						},
						"server_select_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection failure",
						},
						"forward_req_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Forward request failure",
						},
						"forward_req_data_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Forward REQ data failure",
						},
						"req_retran": {
							Type: schema.TypeInt, Optional: true, Description: "Request retransmit",
						},
						"req_ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Request pkt out-of-order",
						},
						"server_reselect": {
							Type: schema.TypeInt, Optional: true, Description: "Server reselection",
						},
						"server_prem_close": {
							Type: schema.TypeInt, Optional: true, Description: "Server premature close",
						},
						"new_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection made",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT failure",
						},
						"tcp_out_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out reset",
						},
						"aflex_switch": {
							Type: schema.TypeInt, Optional: true, Description: "aFleX switching",
						},
						"aflex_switch_ok": {
							Type: schema.TypeInt, Optional: true, Description: "aFleX switching (succ)",
						},
						"recv_client_command_ehlo": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client EHLO",
						},
						"recv_client_command_helo": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client HELO",
						},
						"recv_client_command_mail": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client MAIL",
						},
						"recv_client_command_rcpt": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client RCPT",
						},
						"recv_client_command_data": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client DATA",
						},
						"recv_client_command_rset": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client RSET",
						},
						"recv_client_command_vrfy": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client VRFY",
						},
						"recv_client_command_expn": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client EXPN",
						},
						"recv_client_command_help": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client HELP",
						},
						"recv_client_command_noop": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client NOOP",
						},
						"recv_client_command_quit": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client QUIT",
						},
						"recv_client_command_starttls": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client STARTTLS",
						},
						"recv_client_command_others": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client other cmds",
						},
						"send_client_service_ready": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client serv-rdy",
						},
						"send_client_service_not_ready": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client serv-not-rdy",
						},
						"send_client_close_connection": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client close-conn",
						},
						"send_client_go_ahead": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client go-ahead",
						},
						"send_client_start_tls_first": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client STARTTLS-1st",
						},
						"send_client_tls_not_available": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client TLS-not-aval",
						},
						"send_client_no_command": {
							Type: schema.TypeInt, Optional: true, Description: "Sent client no-such-cmd",
						},
						"send_server_cmd_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Sent server RSET",
						},
						"tls_established": {
							Type: schema.TypeInt, Optional: true, Description: "SSL session established",
						},
						"l4_switch": {
							Type: schema.TypeInt, Optional: true, Description: "L4 switching",
						},
						"recv_server_service_not_ready": {
							Type: schema.TypeInt, Optional: true, Description: "Recv server serv-not-rdy",
						},
						"recv_server_unknow_reply_code": {
							Type: schema.TypeInt, Optional: true, Description: "Recv server unknown-code",
						},
						"client_domain_switch": {
							Type: schema.TypeInt, Optional: true, Description: "Client domain switching",
						},
						"client_domain_switch_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Client domain sw (succ)",
						},
						"lb_switch": {
							Type: schema.TypeInt, Optional: true, Description: "LB switching",
						},
						"lb_switch_ok": {
							Type: schema.TypeInt, Optional: true, Description: "LB switching (succ)",
						},
						"read_request_line_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Read request line fail",
						},
						"get_all_headers_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Get all headers fail",
						},
						"too_many_headers": {
							Type: schema.TypeInt, Optional: true, Description: "Too many headers",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Line too long",
						},
						"line_across_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Line across packets",
						},
						"line_extend": {
							Type: schema.TypeInt, Optional: true, Description: "Line extend",
						},
						"line_extend_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Line extend fail",
						},
						"line_table_extend": {
							Type: schema.TypeInt, Optional: true, Description: "Table extend",
						},
						"line_table_extend_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Table extend fail",
						},
						"parse_request_line_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse request line fail",
						},
						"insert_resonse_line_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Ins response line fail",
						},
						"remove_resonse_line_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Del response line fail",
						},
						"parse_resonse_line_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parse response line fail",
						},
						"aflex_lb_reselect": {
							Type: schema.TypeInt, Optional: true, Description: "aFleX lb reselect",
						},
						"aflex_lb_reselect_ok": {
							Type: schema.TypeInt, Optional: true, Description: "aFleX lb reselect (succ)",
						},
						"server_starttls_init": {
							Type: schema.TypeInt, Optional: true, Description: "Init server side STARTTLS",
						},
						"server_starttls_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server side STARTTLS fail",
						},
						"rserver_starttls_disable": {
							Type: schema.TypeInt, Optional: true, Description: "real server not support STARTTLS",
						},
						"recv_client_command_turn": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client TURN",
						},
						"recv_client_command_etrn": {
							Type: schema.TypeInt, Optional: true, Description: "Recv client ETRN",
						},
						"send_server_ehlo": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy sends server EHLO",
						},
						"fail_to_save_client_ehlo": {
							Type: schema.TypeInt, Optional: true, Description: "Failed to save client EHLO",
						},
						"aflex_mail_fail": {
							Type: schema.TypeInt, Optional: true, Description: "aFlex Mail event failed",
						},
						"drop_server_ehlo_ok": {
							Type: schema.TypeInt, Optional: true, Description: "Server EHLO_OK dropped",
						},
						"client_ehlo_saved": {
							Type: schema.TypeInt, Optional: true, Description: "Client EHLO saved",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSmtpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmtpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmtpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSmtpStatsStats := setObjectSlbSmtpStatsStats(res)
		d.Set("stats", SlbSmtpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSmtpStatsStats(ret edpt.DataSlbSmtpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":                    ret.DtSlbSmtpStats.Stats.Curr_proxy,
			"total_proxy":                   ret.DtSlbSmtpStats.Stats.Total_proxy,
			"request":                       ret.DtSlbSmtpStats.Stats.Request,
			"request_success":               ret.DtSlbSmtpStats.Stats.Request_success,
			"no_proxy":                      ret.DtSlbSmtpStats.Stats.No_proxy,
			"client_reset":                  ret.DtSlbSmtpStats.Stats.Client_reset,
			"server_reset":                  ret.DtSlbSmtpStats.Stats.Server_reset,
			"no_tuple":                      ret.DtSlbSmtpStats.Stats.No_tuple,
			"parse_req_fail":                ret.DtSlbSmtpStats.Stats.Parse_req_fail,
			"server_select_fail":            ret.DtSlbSmtpStats.Stats.Server_select_fail,
			"forward_req_fail":              ret.DtSlbSmtpStats.Stats.Forward_req_fail,
			"forward_req_data_fail":         ret.DtSlbSmtpStats.Stats.Forward_req_data_fail,
			"req_retran":                    ret.DtSlbSmtpStats.Stats.Req_retran,
			"req_ofo":                       ret.DtSlbSmtpStats.Stats.Req_ofo,
			"server_reselect":               ret.DtSlbSmtpStats.Stats.Server_reselect,
			"server_prem_close":             ret.DtSlbSmtpStats.Stats.Server_prem_close,
			"new_server_conn":               ret.DtSlbSmtpStats.Stats.New_server_conn,
			"snat_fail":                     ret.DtSlbSmtpStats.Stats.Snat_fail,
			"tcp_out_reset":                 ret.DtSlbSmtpStats.Stats.Tcp_out_reset,
			"aflex_switch":                  ret.DtSlbSmtpStats.Stats.Aflex_switch,
			"aflex_switch_ok":               ret.DtSlbSmtpStats.Stats.Aflex_switch_ok,
			"recv_client_command_ehlo":      ret.DtSlbSmtpStats.Stats.Recv_client_command_ehlo,
			"recv_client_command_helo":      ret.DtSlbSmtpStats.Stats.Recv_client_command_helo,
			"recv_client_command_mail":      ret.DtSlbSmtpStats.Stats.Recv_client_command_mail,
			"recv_client_command_rcpt":      ret.DtSlbSmtpStats.Stats.Recv_client_command_rcpt,
			"recv_client_command_data":      ret.DtSlbSmtpStats.Stats.Recv_client_command_data,
			"recv_client_command_rset":      ret.DtSlbSmtpStats.Stats.Recv_client_command_rset,
			"recv_client_command_vrfy":      ret.DtSlbSmtpStats.Stats.Recv_client_command_vrfy,
			"recv_client_command_expn":      ret.DtSlbSmtpStats.Stats.Recv_client_command_expn,
			"recv_client_command_help":      ret.DtSlbSmtpStats.Stats.Recv_client_command_help,
			"recv_client_command_noop":      ret.DtSlbSmtpStats.Stats.Recv_client_command_noop,
			"recv_client_command_quit":      ret.DtSlbSmtpStats.Stats.Recv_client_command_quit,
			"recv_client_command_starttls":  ret.DtSlbSmtpStats.Stats.Recv_client_command_starttls,
			"recv_client_command_others":    ret.DtSlbSmtpStats.Stats.Recv_client_command_others,
			"send_client_service_ready":     ret.DtSlbSmtpStats.Stats.Send_client_service_ready,
			"send_client_service_not_ready": ret.DtSlbSmtpStats.Stats.Send_client_service_not_ready,
			"send_client_close_connection":  ret.DtSlbSmtpStats.Stats.Send_client_close_connection,
			"send_client_go_ahead":          ret.DtSlbSmtpStats.Stats.Send_client_go_ahead,
			"send_client_start_tls_first":   ret.DtSlbSmtpStats.Stats.Send_client_start_tls_first,
			"send_client_tls_not_available": ret.DtSlbSmtpStats.Stats.Send_client_tls_not_available,
			"send_client_no_command":        ret.DtSlbSmtpStats.Stats.Send_client_no_command,
			"send_server_cmd_reset":         ret.DtSlbSmtpStats.Stats.Send_server_cmd_reset,
			"tls_established":               ret.DtSlbSmtpStats.Stats.Tls_established,
			"l4_switch":                     ret.DtSlbSmtpStats.Stats.L4_switch,
			"recv_server_service_not_ready": ret.DtSlbSmtpStats.Stats.Recv_server_service_not_ready,
			"recv_server_unknow_reply_code": ret.DtSlbSmtpStats.Stats.Recv_server_unknow_reply_code,
			"client_domain_switch":          ret.DtSlbSmtpStats.Stats.Client_domain_switch,
			"client_domain_switch_ok":       ret.DtSlbSmtpStats.Stats.Client_domain_switch_ok,
			"lb_switch":                     ret.DtSlbSmtpStats.Stats.Lb_switch,
			"lb_switch_ok":                  ret.DtSlbSmtpStats.Stats.Lb_switch_ok,
			"read_request_line_fail":        ret.DtSlbSmtpStats.Stats.Read_request_line_fail,
			"get_all_headers_fail":          ret.DtSlbSmtpStats.Stats.Get_all_headers_fail,
			"too_many_headers":              ret.DtSlbSmtpStats.Stats.Too_many_headers,
			"line_too_long":                 ret.DtSlbSmtpStats.Stats.Line_too_long,
			"line_across_packet":            ret.DtSlbSmtpStats.Stats.Line_across_packet,
			"line_extend":                   ret.DtSlbSmtpStats.Stats.Line_extend,
			"line_extend_fail":              ret.DtSlbSmtpStats.Stats.Line_extend_fail,
			"line_table_extend":             ret.DtSlbSmtpStats.Stats.Line_table_extend,
			"line_table_extend_fail":        ret.DtSlbSmtpStats.Stats.Line_table_extend_fail,
			"parse_request_line_fail":       ret.DtSlbSmtpStats.Stats.Parse_request_line_fail,
			"insert_resonse_line_fail":      ret.DtSlbSmtpStats.Stats.Insert_resonse_line_fail,
			"remove_resonse_line_fail":      ret.DtSlbSmtpStats.Stats.Remove_resonse_line_fail,
			"parse_resonse_line_fail":       ret.DtSlbSmtpStats.Stats.Parse_resonse_line_fail,
			"aflex_lb_reselect":             ret.DtSlbSmtpStats.Stats.Aflex_lb_reselect,
			"aflex_lb_reselect_ok":          ret.DtSlbSmtpStats.Stats.Aflex_lb_reselect_ok,
			"server_starttls_init":          ret.DtSlbSmtpStats.Stats.Server_starttls_init,
			"server_starttls_fail":          ret.DtSlbSmtpStats.Stats.Server_starttls_fail,
			"rserver_starttls_disable":      ret.DtSlbSmtpStats.Stats.Rserver_starttls_disable,
			"recv_client_command_turn":      ret.DtSlbSmtpStats.Stats.Recv_client_command_turn,
			"recv_client_command_etrn":      ret.DtSlbSmtpStats.Stats.Recv_client_command_etrn,
			"send_server_ehlo":              ret.DtSlbSmtpStats.Stats.Send_server_ehlo,
			"fail_to_save_client_ehlo":      ret.DtSlbSmtpStats.Stats.Fail_to_save_client_ehlo,
			"aflex_mail_fail":               ret.DtSlbSmtpStats.Stats.Aflex_mail_fail,
			"drop_server_ehlo_ok":           ret.DtSlbSmtpStats.Stats.Drop_server_ehlo_ok,
			"client_ehlo_saved":             ret.DtSlbSmtpStats.Stats.Client_ehlo_saved,
		},
	}
}

func getObjectSlbSmtpStatsStats(d []interface{}) edpt.SlbSmtpStatsStats {

	count1 := len(d)
	var ret edpt.SlbSmtpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Request = in["request"].(int)
		ret.Request_success = in["request_success"].(int)
		ret.No_proxy = in["no_proxy"].(int)
		ret.Client_reset = in["client_reset"].(int)
		ret.Server_reset = in["server_reset"].(int)
		ret.No_tuple = in["no_tuple"].(int)
		ret.Parse_req_fail = in["parse_req_fail"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Forward_req_fail = in["forward_req_fail"].(int)
		ret.Forward_req_data_fail = in["forward_req_data_fail"].(int)
		ret.Req_retran = in["req_retran"].(int)
		ret.Req_ofo = in["req_ofo"].(int)
		ret.Server_reselect = in["server_reselect"].(int)
		ret.Server_prem_close = in["server_prem_close"].(int)
		ret.New_server_conn = in["new_server_conn"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Tcp_out_reset = in["tcp_out_reset"].(int)
		ret.Aflex_switch = in["aflex_switch"].(int)
		ret.Aflex_switch_ok = in["aflex_switch_ok"].(int)
		ret.Recv_client_command_ehlo = in["recv_client_command_ehlo"].(int)
		ret.Recv_client_command_helo = in["recv_client_command_helo"].(int)
		ret.Recv_client_command_mail = in["recv_client_command_mail"].(int)
		ret.Recv_client_command_rcpt = in["recv_client_command_rcpt"].(int)
		ret.Recv_client_command_data = in["recv_client_command_data"].(int)
		ret.Recv_client_command_rset = in["recv_client_command_rset"].(int)
		ret.Recv_client_command_vrfy = in["recv_client_command_vrfy"].(int)
		ret.Recv_client_command_expn = in["recv_client_command_expn"].(int)
		ret.Recv_client_command_help = in["recv_client_command_help"].(int)
		ret.Recv_client_command_noop = in["recv_client_command_noop"].(int)
		ret.Recv_client_command_quit = in["recv_client_command_quit"].(int)
		ret.Recv_client_command_starttls = in["recv_client_command_starttls"].(int)
		ret.Recv_client_command_others = in["recv_client_command_others"].(int)
		ret.Send_client_service_ready = in["send_client_service_ready"].(int)
		ret.Send_client_service_not_ready = in["send_client_service_not_ready"].(int)
		ret.Send_client_close_connection = in["send_client_close_connection"].(int)
		ret.Send_client_go_ahead = in["send_client_go_ahead"].(int)
		ret.Send_client_start_tls_first = in["send_client_start_tls_first"].(int)
		ret.Send_client_tls_not_available = in["send_client_tls_not_available"].(int)
		ret.Send_client_no_command = in["send_client_no_command"].(int)
		ret.Send_server_cmd_reset = in["send_server_cmd_reset"].(int)
		ret.Tls_established = in["tls_established"].(int)
		ret.L4_switch = in["l4_switch"].(int)
		ret.Recv_server_service_not_ready = in["recv_server_service_not_ready"].(int)
		ret.Recv_server_unknow_reply_code = in["recv_server_unknow_reply_code"].(int)
		ret.Client_domain_switch = in["client_domain_switch"].(int)
		ret.Client_domain_switch_ok = in["client_domain_switch_ok"].(int)
		ret.Lb_switch = in["lb_switch"].(int)
		ret.Lb_switch_ok = in["lb_switch_ok"].(int)
		ret.Read_request_line_fail = in["read_request_line_fail"].(int)
		ret.Get_all_headers_fail = in["get_all_headers_fail"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_across_packet = in["line_across_packet"].(int)
		ret.Line_extend = in["line_extend"].(int)
		ret.Line_extend_fail = in["line_extend_fail"].(int)
		ret.Line_table_extend = in["line_table_extend"].(int)
		ret.Line_table_extend_fail = in["line_table_extend_fail"].(int)
		ret.Parse_request_line_fail = in["parse_request_line_fail"].(int)
		ret.Insert_resonse_line_fail = in["insert_resonse_line_fail"].(int)
		ret.Remove_resonse_line_fail = in["remove_resonse_line_fail"].(int)
		ret.Parse_resonse_line_fail = in["parse_resonse_line_fail"].(int)
		ret.Aflex_lb_reselect = in["aflex_lb_reselect"].(int)
		ret.Aflex_lb_reselect_ok = in["aflex_lb_reselect_ok"].(int)
		ret.Server_starttls_init = in["server_starttls_init"].(int)
		ret.Server_starttls_fail = in["server_starttls_fail"].(int)
		ret.Rserver_starttls_disable = in["rserver_starttls_disable"].(int)
		ret.Recv_client_command_turn = in["recv_client_command_turn"].(int)
		ret.Recv_client_command_etrn = in["recv_client_command_etrn"].(int)
		ret.Send_server_ehlo = in["send_server_ehlo"].(int)
		ret.Fail_to_save_client_ehlo = in["fail_to_save_client_ehlo"].(int)
		ret.Aflex_mail_fail = in["aflex_mail_fail"].(int)
		ret.Drop_server_ehlo_ok = in["drop_server_ehlo_ok"].(int)
		ret.Client_ehlo_saved = in["client_ehlo_saved"].(int)
	}
	return ret
}

func dataToEndpointSlbSmtpStats(d *schema.ResourceData) edpt.SlbSmtpStats {
	var ret edpt.SlbSmtpStats

	ret.Stats = getObjectSlbSmtpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
