package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmtpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_smtp_oper`: Operational Status for the object smtp\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSmtpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smtp_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_req_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_select_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"forward_req_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"forward_req_data_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_retran": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_reselect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_prem_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"new_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_out_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_ehlo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_helo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_mail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_rcpt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_data": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_rset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_vrfy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_expn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_help": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_noop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_quit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_starttls": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_turn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_etrn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_client_command_others": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_server_service_not_ready": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_server_unknow_reply_code": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_service_ready": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_service_not_ready": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_close_connection": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_go_ahead": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_start_tls_first": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_tls_not_available": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_client_no_command": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_server_cmd_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tls_established": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l4_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_switch_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_domain_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_domain_switch_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lb_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lb_switch_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"read_request_line_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"get_all_headers_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"too_many_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_across_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_extend": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_extend_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_table_extend": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"line_table_extend_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_request_line_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insert_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"remove_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"parse_resonse_line_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_lb_reselect": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_lb_reselect_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_starttls_init": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_starttls_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rserver_starttls_disable": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_server_ehlo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_to_save_client_ehlo": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_mail_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drop_server_ehlo_ok": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_ehlo_saved": {
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

func resourceSlbSmtpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmtpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmtpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSmtpOperOper := setObjectSlbSmtpOperOper(res)
		d.Set("oper", SlbSmtpOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSmtpOperOper(ret edpt.DataSlbSmtpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"smtp_cpu_list": setSliceSlbSmtpOperOperSmtpCpuList(ret.DtSlbSmtpOper.Oper.SmtpCpuList),
			"cpu_count":     ret.DtSlbSmtpOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbSmtpOperOperSmtpCpuList(d []edpt.SlbSmtpOperOperSmtpCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr_proxy"] = item.Curr_proxy
		in["total_proxy"] = item.Total_proxy
		in["request"] = item.Request
		in["request_success"] = item.Request_success
		in["no_proxy"] = item.No_proxy
		in["client_reset"] = item.Client_reset
		in["server_reset"] = item.Server_reset
		in["no_tuple"] = item.No_tuple
		in["parse_req_fail"] = item.Parse_req_fail
		in["server_select_fail"] = item.Server_select_fail
		in["forward_req_fail"] = item.Forward_req_fail
		in["forward_req_data_fail"] = item.Forward_req_data_fail
		in["req_retran"] = item.Req_retran
		in["req_ofo"] = item.Req_ofo
		in["server_reselect"] = item.Server_reselect
		in["server_prem_close"] = item.Server_prem_close
		in["new_server_conn"] = item.New_server_conn
		in["snat_fail"] = item.Snat_fail
		in["tcp_out_reset"] = item.Tcp_out_reset
		in["recv_client_command_ehlo"] = item.Recv_client_command_ehlo
		in["recv_client_command_helo"] = item.Recv_client_command_helo
		in["recv_client_command_mail"] = item.Recv_client_command_mail
		in["recv_client_command_rcpt"] = item.Recv_client_command_rcpt
		in["recv_client_command_data"] = item.Recv_client_command_data
		in["recv_client_command_rset"] = item.Recv_client_command_rset
		in["recv_client_command_vrfy"] = item.Recv_client_command_vrfy
		in["recv_client_command_expn"] = item.Recv_client_command_expn
		in["recv_client_command_help"] = item.Recv_client_command_help
		in["recv_client_command_noop"] = item.Recv_client_command_noop
		in["recv_client_command_quit"] = item.Recv_client_command_quit
		in["recv_client_command_starttls"] = item.Recv_client_command_starttls
		in["recv_client_command_turn"] = item.Recv_client_command_turn
		in["recv_client_command_etrn"] = item.Recv_client_command_etrn
		in["recv_client_command_others"] = item.Recv_client_command_others
		in["recv_server_service_not_ready"] = item.Recv_server_service_not_ready
		in["recv_server_unknow_reply_code"] = item.Recv_server_unknow_reply_code
		in["send_client_service_ready"] = item.Send_client_service_ready
		in["send_client_service_not_ready"] = item.Send_client_service_not_ready
		in["send_client_close_connection"] = item.Send_client_close_connection
		in["send_client_go_ahead"] = item.Send_client_go_ahead
		in["send_client_start_tls_first"] = item.Send_client_start_tls_first
		in["send_client_tls_not_available"] = item.Send_client_tls_not_available
		in["send_client_no_command"] = item.Send_client_no_command
		in["send_server_cmd_reset"] = item.Send_server_cmd_reset
		in["tls_established"] = item.Tls_established
		in["l4_switch"] = item.L4_switch
		in["aflex_switch"] = item.Aflex_switch
		in["aflex_switch_ok"] = item.Aflex_switch_ok
		in["client_domain_switch"] = item.Client_domain_switch
		in["client_domain_switch_ok"] = item.Client_domain_switch_ok
		in["lb_switch"] = item.Lb_switch
		in["lb_switch_ok"] = item.Lb_switch_ok
		in["read_request_line_fail"] = item.Read_request_line_fail
		in["get_all_headers_fail"] = item.Get_all_headers_fail
		in["too_many_headers"] = item.Too_many_headers
		in["line_too_long"] = item.Line_too_long
		in["line_across_packet"] = item.Line_across_packet
		in["line_extend"] = item.Line_extend
		in["line_extend_fail"] = item.Line_extend_fail
		in["line_table_extend"] = item.Line_table_extend
		in["line_table_extend_fail"] = item.Line_table_extend_fail
		in["parse_request_line_fail"] = item.Parse_request_line_fail
		in["insert_resonse_line_fail"] = item.Insert_resonse_line_fail
		in["remove_resonse_line_fail"] = item.Remove_resonse_line_fail
		in["parse_resonse_line_fail"] = item.Parse_resonse_line_fail
		in["aflex_lb_reselect"] = item.Aflex_lb_reselect
		in["aflex_lb_reselect_ok"] = item.Aflex_lb_reselect_ok
		in["server_starttls_init"] = item.Server_starttls_init
		in["server_starttls_fail"] = item.Server_starttls_fail
		in["rserver_starttls_disable"] = item.Rserver_starttls_disable
		in["send_server_ehlo"] = item.Send_server_ehlo
		in["fail_to_save_client_ehlo"] = item.Fail_to_save_client_ehlo
		in["aflex_mail_fail"] = item.Aflex_mail_fail
		in["drop_server_ehlo_ok"] = item.Drop_server_ehlo_ok
		in["client_ehlo_saved"] = item.Client_ehlo_saved
		result = append(result, in)
	}
	return result
}

func getObjectSlbSmtpOperOper(d []interface{}) edpt.SlbSmtpOperOper {

	count1 := len(d)
	var ret edpt.SlbSmtpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SmtpCpuList = getSliceSlbSmtpOperOperSmtpCpuList(in["smtp_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbSmtpOperOperSmtpCpuList(d []interface{}) []edpt.SlbSmtpOperOperSmtpCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbSmtpOperOperSmtpCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSmtpOperOperSmtpCpuList
		oi.Curr_proxy = in["curr_proxy"].(int)
		oi.Total_proxy = in["total_proxy"].(int)
		oi.Request = in["request"].(int)
		oi.Request_success = in["request_success"].(int)
		oi.No_proxy = in["no_proxy"].(int)
		oi.Client_reset = in["client_reset"].(int)
		oi.Server_reset = in["server_reset"].(int)
		oi.No_tuple = in["no_tuple"].(int)
		oi.Parse_req_fail = in["parse_req_fail"].(int)
		oi.Server_select_fail = in["server_select_fail"].(int)
		oi.Forward_req_fail = in["forward_req_fail"].(int)
		oi.Forward_req_data_fail = in["forward_req_data_fail"].(int)
		oi.Req_retran = in["req_retran"].(int)
		oi.Req_ofo = in["req_ofo"].(int)
		oi.Server_reselect = in["server_reselect"].(int)
		oi.Server_prem_close = in["server_prem_close"].(int)
		oi.New_server_conn = in["new_server_conn"].(int)
		oi.Snat_fail = in["snat_fail"].(int)
		oi.Tcp_out_reset = in["tcp_out_reset"].(int)
		oi.Recv_client_command_ehlo = in["recv_client_command_ehlo"].(int)
		oi.Recv_client_command_helo = in["recv_client_command_helo"].(int)
		oi.Recv_client_command_mail = in["recv_client_command_mail"].(int)
		oi.Recv_client_command_rcpt = in["recv_client_command_rcpt"].(int)
		oi.Recv_client_command_data = in["recv_client_command_data"].(int)
		oi.Recv_client_command_rset = in["recv_client_command_rset"].(int)
		oi.Recv_client_command_vrfy = in["recv_client_command_vrfy"].(int)
		oi.Recv_client_command_expn = in["recv_client_command_expn"].(int)
		oi.Recv_client_command_help = in["recv_client_command_help"].(int)
		oi.Recv_client_command_noop = in["recv_client_command_noop"].(int)
		oi.Recv_client_command_quit = in["recv_client_command_quit"].(int)
		oi.Recv_client_command_starttls = in["recv_client_command_starttls"].(int)
		oi.Recv_client_command_turn = in["recv_client_command_turn"].(int)
		oi.Recv_client_command_etrn = in["recv_client_command_etrn"].(int)
		oi.Recv_client_command_others = in["recv_client_command_others"].(int)
		oi.Recv_server_service_not_ready = in["recv_server_service_not_ready"].(int)
		oi.Recv_server_unknow_reply_code = in["recv_server_unknow_reply_code"].(int)
		oi.Send_client_service_ready = in["send_client_service_ready"].(int)
		oi.Send_client_service_not_ready = in["send_client_service_not_ready"].(int)
		oi.Send_client_close_connection = in["send_client_close_connection"].(int)
		oi.Send_client_go_ahead = in["send_client_go_ahead"].(int)
		oi.Send_client_start_tls_first = in["send_client_start_tls_first"].(int)
		oi.Send_client_tls_not_available = in["send_client_tls_not_available"].(int)
		oi.Send_client_no_command = in["send_client_no_command"].(int)
		oi.Send_server_cmd_reset = in["send_server_cmd_reset"].(int)
		oi.Tls_established = in["tls_established"].(int)
		oi.L4_switch = in["l4_switch"].(int)
		oi.Aflex_switch = in["aflex_switch"].(int)
		oi.Aflex_switch_ok = in["aflex_switch_ok"].(int)
		oi.Client_domain_switch = in["client_domain_switch"].(int)
		oi.Client_domain_switch_ok = in["client_domain_switch_ok"].(int)
		oi.Lb_switch = in["lb_switch"].(int)
		oi.Lb_switch_ok = in["lb_switch_ok"].(int)
		oi.Read_request_line_fail = in["read_request_line_fail"].(int)
		oi.Get_all_headers_fail = in["get_all_headers_fail"].(int)
		oi.Too_many_headers = in["too_many_headers"].(int)
		oi.Line_too_long = in["line_too_long"].(int)
		oi.Line_across_packet = in["line_across_packet"].(int)
		oi.Line_extend = in["line_extend"].(int)
		oi.Line_extend_fail = in["line_extend_fail"].(int)
		oi.Line_table_extend = in["line_table_extend"].(int)
		oi.Line_table_extend_fail = in["line_table_extend_fail"].(int)
		oi.Parse_request_line_fail = in["parse_request_line_fail"].(int)
		oi.Insert_resonse_line_fail = in["insert_resonse_line_fail"].(int)
		oi.Remove_resonse_line_fail = in["remove_resonse_line_fail"].(int)
		oi.Parse_resonse_line_fail = in["parse_resonse_line_fail"].(int)
		oi.Aflex_lb_reselect = in["aflex_lb_reselect"].(int)
		oi.Aflex_lb_reselect_ok = in["aflex_lb_reselect_ok"].(int)
		oi.Server_starttls_init = in["server_starttls_init"].(int)
		oi.Server_starttls_fail = in["server_starttls_fail"].(int)
		oi.Rserver_starttls_disable = in["rserver_starttls_disable"].(int)
		oi.Send_server_ehlo = in["send_server_ehlo"].(int)
		oi.Fail_to_save_client_ehlo = in["fail_to_save_client_ehlo"].(int)
		oi.Aflex_mail_fail = in["aflex_mail_fail"].(int)
		oi.Drop_server_ehlo_ok = in["drop_server_ehlo_ok"].(int)
		oi.Client_ehlo_saved = in["client_ehlo_saved"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSmtpOper(d *schema.ResourceData) edpt.SlbSmtpOper {
	var ret edpt.SlbSmtpOper

	ret.Oper = getObjectSlbSmtpOperOper(d.Get("oper").([]interface{}))
	return ret
}
