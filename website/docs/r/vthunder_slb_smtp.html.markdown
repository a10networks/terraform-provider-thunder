---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_smtp"
sidebar_current: "docs-vthunder-slb-smtp"
description: |-
    Provides details about vthunder SLB smtp resource for A10
---

# vthunder\_slb\_smtp

`vthunder_slb_smtp` Provides details about vthunder SLB smtp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_smtp" "smtp" {
	sampling_enable {
		counters1 = "all"
	} 
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'curr_proxy’: Current proxy conns; 'total_proxy’: Total proxy conns; 'request’: SMTP requests; 'request_success’: SMTP requests (success); 'no_proxy’: No proxy error; 'client_reset’: Client reset; 'server_reset’: Server reset; 'no_tuple’: No tuple error; 'parse_req_fail’: Parse request failure; 'server_select_fail’: Server selection failure; 'forward_req_fail’: Forward request failure; 'forward_req_data_fail’: Forward REQ data failure; 'req_retran’: Request retransmit; 'req_ofo’: Request pkt out-of-order; 'server_reselect’: Server reselection; 'server_prem_close’: Server premature close; 'new_server_conn’: Server connection made; 'snat_fail’: Source NAT failure; 'tcp_out_reset’: TCP out reset; 'recv_client_command_EHLO’: Recv client EHLO; 'recv_client_command_HELO’: Recv client HELO; 'recv_client_command_MAIL’: Recv client MAIL; 'recv_client_command_RCPT’: Recv client RCPT; 'recv_client_command_DATA’: Recv client DATA; 'recv_client_command_RSET’: Recv client RSET; 'recv_client_command_VRFY’: Recv client VRFY; 'recv_client_command_EXPN’: Recv client EXPN; 'recv_client_command_HELP’: Recv client HELP; 'recv_client_command_NOOP’: Recv client NOOP; 'recv_client_command_QUIT’: Recv client QUIT; 'recv_client_command_STARTTLS’: Recv client STARTTLS; 'recv_client_command_others’: Recv client other cmds; 'recv_server_service_not_ready’: Recv server serv-not-rdy; 'recv_server_unknow_reply_code’: Recv server unknown-code; 'send_client_service_ready’: Sent client serv-rdy; 'send_client_service_not_ready’: Sent client serv-not-rdy; 'send_client_close_connection’: Sent client close-conn; 'send_client_go_ahead’: Sent client go-ahead; 'send_client_start_TLS_first’: Sent client STARTTLS-1st; 'send_client_TLS_not_available’: Sent client TLS-not-aval; 'send_client_no_command’: Sent client no-such-cmd; 'send_server_cmd_reset’: Sent server RSET; 'TLS_established’: SSL session established; 'L4_switch’: L4 switching; 'Aflex_switch’: aFleX switching; 'Aflex_switch_ok’: aFleX switching (succ); 'client_domain_switch’: Client domain switching; 'client_domain_switch_ok’: Client domain sw (succ); 'LB_switch’: LB switching; 'LB_switch_ok’: LB switching (succ); 'read_request_line_fail’: Read request line fail; 'get_all_headers_fail’: Get all headers fail; 'too_many_headers’: Too many headers; 'line_too_long’: Line too long; 'line_across_packet’: Line across packets; 'line_extend’: Line extend; 'line_extend_fail’: Line extend fail; 'line_table_extend’: Table extend; 'line_table_extend_fail’: Table extend fail; 'parse_request_line_fail’: Parse request line fail; 'insert_resonse_line_fail’: Ins response line fail; 'remove_resonse_line_fail’: Del response line fail; 'parse_resonse_line_fail’: Parse response line fail; 'Aflex_lb_reselect’: aFleX lb reselect; 'Aflex_lb_reselect_ok’: aFleX lb reselect (succ); 'server_STARTTLS_init’: Init server side STARTTLS; 'server_STARTTLS_fail’: Server side STARTTLS fail; 'rserver_STARTTLS_disable’: real server not support STARTTLS; 'recv_client_command_TURN’: Recv client TURN; 'recv_client_command_ETRN’: Recv client ETRN;