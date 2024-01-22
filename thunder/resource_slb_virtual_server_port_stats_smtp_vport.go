package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbVirtualServerPortStats54() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_virtual_server_port_stats_smtp_vport`: Statistics for the object port\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbVirtualServerPortStats54Create,
		UpdateContext: resourceSlbVirtualServerPortStats54Update,
		ReadContext:   resourceSlbVirtualServerPortStats54Read,
		DeleteContext: resourceSlbVirtualServerPortStats54Delete,

		Schema: map[string]*schema.Schema{
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': TCP LB service; 'udp': UDP Port; 'others': for no tcp/udp protocol, do IP load balancing; 'diameter': diameter port; 'dns-tcp': DNS service over TCP; 'dns-udp': DNS service over UDP; 'fast-http': Fast HTTP Port; 'fix': FIX Port; 'ftp': File Transfer Protocol Port; 'ftp-proxy': ftp proxy port; 'http': HTTP Port; 'https': HTTPS port; 'imap': imap proxy port; 'mlb': Message based load balancing; 'mms': Microsoft Multimedia Service Port; 'mysql': mssql port; 'mssql': mssql; 'pop3': pop3 proxy port; 'radius': RADIUS Port; 'rtsp': Real Time Streaming Protocol Port; 'sip': Session initiation protocol over UDP; 'sip-tcp': Session initiation protocol over TCP; 'sips': Session initiation protocol over TLS; 'smpp-tcp': SMPP service over TCP; 'spdy': spdy port; 'spdys': spdys port; 'smtp': SMTP Port; 'mqtt': MQTT Port; 'mqtts': MQTTS Port; 'ssl-proxy': Generic SSL proxy; 'ssli': SSL insight; 'ssh': SSH Port; 'tcp-proxy': Generic TCP proxy; 'tftp': TFTP Port; 'fast-fix': Fast FIX port; 'http-over-quic': HTTP3-over-quic port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smtp_vport": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmd_ehlo": {
										Type: schema.TypeInt, Optional: true, Description: "EHLO commands",
									},
									"cmd_helo": {
										Type: schema.TypeInt, Optional: true, Description: "HELO commands",
									},
									"cmd_mail": {
										Type: schema.TypeInt, Optional: true, Description: "MAIL commands",
									},
									"cmd_rcpt": {
										Type: schema.TypeInt, Optional: true, Description: "RCPT commands",
									},
									"cmd_data": {
										Type: schema.TypeInt, Optional: true, Description: "DATA commands",
									},
									"cmd_rset": {
										Type: schema.TypeInt, Optional: true, Description: "RSET commands",
									},
									"cmd_vrfy": {
										Type: schema.TypeInt, Optional: true, Description: "VRFY commands",
									},
									"cmd_expn": {
										Type: schema.TypeInt, Optional: true, Description: "EXPN commands",
									},
									"cmd_help": {
										Type: schema.TypeInt, Optional: true, Description: "HELP commands",
									},
									"cmd_noop": {
										Type: schema.TypeInt, Optional: true, Description: "NOOP commands",
									},
									"cmd_starttls": {
										Type: schema.TypeInt, Optional: true, Description: "STARTTLS commands",
									},
									"cmd_turn": {
										Type: schema.TypeInt, Optional: true, Description: "TURN commands",
									},
									"cmd_etrn": {
										Type: schema.TypeInt, Optional: true, Description: "ETRN commands",
									},
									"cmd_quit": {
										Type: schema.TypeInt, Optional: true, Description: "QUIT commands",
									},
									"cmd_local": {
										Type: schema.TypeInt, Optional: true, Description: "Local extension commands",
									},
									"cmd_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "Unknown commands",
									},
									"code_200": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 200",
									},
									"code_211": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 211",
									},
									"code_214": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 214",
									},
									"code_220": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 220",
									},
									"code_221": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 221",
									},
									"code_250": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 250",
									},
									"code_251": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 251",
									},
									"code_252": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 252",
									},
									"code_354": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 354",
									},
									"code_421": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 421",
									},
									"code_450": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 450",
									},
									"code_451": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 451",
									},
									"code_452": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 452",
									},
									"code_455": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 455",
									},
									"code_500": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 500",
									},
									"code_501": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 501",
									},
									"code_502": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 502",
									},
									"code_503": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 503",
									},
									"code_504": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 504",
									},
									"code_521": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 521",
									},
									"code_530": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 530",
									},
									"code_550": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 550",
									},
									"code_551": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 551",
									},
									"code_552": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 552",
									},
									"code_553": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 553",
									},
									"code_554": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 554",
									},
									"code_555": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 555",
									},
									"code_556": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 556",
									},
									"code_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 2XX",
									},
									"code_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 3XX",
									},
									"code_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 4XX",
									},
									"code_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code 5XX",
									},
									"code_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "Reply code unknown",
									},
									"reply_10u": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 10u",
									},
									"reply_20u": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 20u",
									},
									"reply_50u": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 50u",
									},
									"reply_100u": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 100u",
									},
									"reply_200u": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 200u",
									},
									"reply_500u": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 500u",
									},
									"reply_1m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 1m",
									},
									"reply_2m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 2m",
									},
									"reply_5m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 5m",
									},
									"reply_10m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 10m",
									},
									"reply_20m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 20m",
									},
									"reply_50m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 5m",
									},
									"reply_100m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 100m",
									},
									"reply_200m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 200m",
									},
									"reply_500m": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 500m",
									},
									"reply_1s": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 1s",
									},
									"reply_2s": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 2s",
									},
									"reply_5s": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time less than 5s",
									},
									"reply_over_5s": {
										Type: schema.TypeInt, Optional: true, Description: "Reply time greater than equal to 5s",
									},
									"data_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 1KB",
									},
									"data_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 2KB",
									},
									"data_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 4KB",
									},
									"data_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 8KB",
									},
									"data_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 16KB",
									},
									"data_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 32KB",
									},
									"data_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 64KB",
									},
									"data_sz_128k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 128KB",
									},
									"data_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 256KB",
									},
									"data_sz_512k": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 512KB",
									},
									"data_sz_1m": {
										Type: schema.TypeInt, Optional: true, Description: "Data size less than 1MB",
									},
									"data_sz_gt_1m": {
										Type: schema.TypeInt, Optional: true, Description: "Data size greater than 1MB",
									},
									"total_commands": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of commands received",
									},
									"total_replies": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of replies received",
									},
									"curr_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Current number of SMTP connections",
									},
									"total_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Total number of SMTP connections",
									},
									"peak_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Largest number of concurrent SMTP connections",
									},
									"client_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Total bytes in SMTP commands",
									},
									"server_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Total bytes in SMTP replies",
									},
									"aflex_switch": {
										Type: schema.TypeInt, Optional: true, Description: "aFlex SMTP switching",
									},
									"aflex_switch_ok": {
										Type: schema.TypeInt, Optional: true, Description: "aFlex SMTP switching success",
									},
									"aflex_ehlo_sent": {
										Type: schema.TypeInt, Optional: true, Description: "aFlex SMTP EHLO sent",
									},
									"send_server_ehlo": {
										Type: schema.TypeInt, Optional: true, Description: "Proxy sends server EHLO",
									},
									"fail_to_save_client_ehlo": {
										Type: schema.TypeInt, Optional: true, Description: "Failed to save client EHLO",
									},
									"aflex_mail_fail": {
										Type: schema.TypeInt, Optional: true, Description: "aFlex Mail event abort",
									},
									"drop_server_ehlo_ok": {
										Type: schema.TypeInt, Optional: true, Description: "Server EHLO-OK dropped",
									},
									"client_ehlo_saved": {
										Type: schema.TypeInt, Optional: true, Description: "Client EHLO saved",
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
									"recv_client_command_others": {
										Type: schema.TypeInt, Optional: true, Description: "Recv client other cmds",
									},
									"recv_client_command_starttls": {
										Type: schema.TypeInt, Optional: true, Description: "Recv client STARTTLS",
									},
									"recv_client_command_rcpt": {
										Type: schema.TypeInt, Optional: true, Description: "Recv client RCPT",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbVirtualServerPortStats54Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats54Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats54(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats54Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbVirtualServerPortStats54Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats54Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats54(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbVirtualServerPortStats54Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbVirtualServerPortStats54Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats54Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats54(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbVirtualServerPortStats54Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbVirtualServerPortStats54Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbVirtualServerPortStats54(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbVirtualServerPortStats54Stats(d []interface{}) edpt.SlbVirtualServerPortStats54Stats {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats54Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Smtp_vport = getObjectSlbVirtualServerPortStats54StatsSmtp_vport(in["smtp_vport"].([]interface{}))
	}
	return ret
}

func getObjectSlbVirtualServerPortStats54StatsSmtp_vport(d []interface{}) edpt.SlbVirtualServerPortStats54StatsSmtp_vport {

	count1 := len(d)
	var ret edpt.SlbVirtualServerPortStats54StatsSmtp_vport
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Cmd_ehlo = in["cmd_ehlo"].(int)
		ret.Cmd_helo = in["cmd_helo"].(int)
		ret.Cmd_mail = in["cmd_mail"].(int)
		ret.Cmd_rcpt = in["cmd_rcpt"].(int)
		ret.Cmd_data = in["cmd_data"].(int)
		ret.Cmd_rset = in["cmd_rset"].(int)
		ret.Cmd_vrfy = in["cmd_vrfy"].(int)
		ret.Cmd_expn = in["cmd_expn"].(int)
		ret.Cmd_help = in["cmd_help"].(int)
		ret.Cmd_noop = in["cmd_noop"].(int)
		ret.Cmd_starttls = in["cmd_starttls"].(int)
		ret.Cmd_turn = in["cmd_turn"].(int)
		ret.Cmd_etrn = in["cmd_etrn"].(int)
		ret.Cmd_quit = in["cmd_quit"].(int)
		ret.Cmd_local = in["cmd_local"].(int)
		ret.Cmd_unknown = in["cmd_unknown"].(int)
		ret.Code_200 = in["code_200"].(int)
		ret.Code_211 = in["code_211"].(int)
		ret.Code_214 = in["code_214"].(int)
		ret.Code_220 = in["code_220"].(int)
		ret.Code_221 = in["code_221"].(int)
		ret.Code_250 = in["code_250"].(int)
		ret.Code_251 = in["code_251"].(int)
		ret.Code_252 = in["code_252"].(int)
		ret.Code_354 = in["code_354"].(int)
		ret.Code_421 = in["code_421"].(int)
		ret.Code_450 = in["code_450"].(int)
		ret.Code_451 = in["code_451"].(int)
		ret.Code_452 = in["code_452"].(int)
		ret.Code_455 = in["code_455"].(int)
		ret.Code_500 = in["code_500"].(int)
		ret.Code_501 = in["code_501"].(int)
		ret.Code_502 = in["code_502"].(int)
		ret.Code_503 = in["code_503"].(int)
		ret.Code_504 = in["code_504"].(int)
		ret.Code_521 = in["code_521"].(int)
		ret.Code_530 = in["code_530"].(int)
		ret.Code_550 = in["code_550"].(int)
		ret.Code_551 = in["code_551"].(int)
		ret.Code_552 = in["code_552"].(int)
		ret.Code_553 = in["code_553"].(int)
		ret.Code_554 = in["code_554"].(int)
		ret.Code_555 = in["code_555"].(int)
		ret.Code_556 = in["code_556"].(int)
		ret.Code_2xx = in["code_2xx"].(int)
		ret.Code_3xx = in["code_3xx"].(int)
		ret.Code_4xx = in["code_4xx"].(int)
		ret.Code_5xx = in["code_5xx"].(int)
		ret.Code_unknown = in["code_unknown"].(int)
		ret.Reply_10u = in["reply_10u"].(int)
		ret.Reply_20u = in["reply_20u"].(int)
		ret.Reply_50u = in["reply_50u"].(int)
		ret.Reply_100u = in["reply_100u"].(int)
		ret.Reply_200u = in["reply_200u"].(int)
		ret.Reply_500u = in["reply_500u"].(int)
		ret.Reply_1m = in["reply_1m"].(int)
		ret.Reply_2m = in["reply_2m"].(int)
		ret.Reply_5m = in["reply_5m"].(int)
		ret.Reply_10m = in["reply_10m"].(int)
		ret.Reply_20m = in["reply_20m"].(int)
		ret.Reply_50m = in["reply_50m"].(int)
		ret.Reply_100m = in["reply_100m"].(int)
		ret.Reply_200m = in["reply_200m"].(int)
		ret.Reply_500m = in["reply_500m"].(int)
		ret.Reply_1s = in["reply_1s"].(int)
		ret.Reply_2s = in["reply_2s"].(int)
		ret.Reply_5s = in["reply_5s"].(int)
		ret.Reply_over_5s = in["reply_over_5s"].(int)
		ret.Data_sz_1k = in["data_sz_1k"].(int)
		ret.Data_sz_2k = in["data_sz_2k"].(int)
		ret.Data_sz_4k = in["data_sz_4k"].(int)
		ret.Data_sz_8k = in["data_sz_8k"].(int)
		ret.Data_sz_16k = in["data_sz_16k"].(int)
		ret.Data_sz_32k = in["data_sz_32k"].(int)
		ret.Data_sz_64k = in["data_sz_64k"].(int)
		ret.Data_sz_128k = in["data_sz_128k"].(int)
		ret.Data_sz_256k = in["data_sz_256k"].(int)
		ret.Data_sz_512k = in["data_sz_512k"].(int)
		ret.Data_sz_1m = in["data_sz_1m"].(int)
		ret.Data_sz_gt_1m = in["data_sz_gt_1m"].(int)
		ret.Total_commands = in["total_commands"].(int)
		ret.Total_replies = in["total_replies"].(int)
		ret.Curr_conn = in["curr_conn"].(int)
		ret.Total_conn = in["total_conn"].(int)
		ret.Peak_conn = in["peak_conn"].(int)
		ret.Client_bytes = in["client_bytes"].(int)
		ret.Server_bytes = in["server_bytes"].(int)
		ret.Aflex_switch = in["aflex_switch"].(int)
		ret.Aflex_switch_ok = in["aflex_switch_ok"].(int)
		ret.Aflex_ehlo_sent = in["aflex_ehlo_sent"].(int)
		ret.Send_server_ehlo = in["send_server_ehlo"].(int)
		ret.Fail_to_save_client_ehlo = in["fail_to_save_client_ehlo"].(int)
		ret.Aflex_mail_fail = in["aflex_mail_fail"].(int)
		ret.Drop_server_ehlo_ok = in["drop_server_ehlo_ok"].(int)
		ret.Client_ehlo_saved = in["client_ehlo_saved"].(int)
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
		ret.Recv_client_command_others = in["recv_client_command_others"].(int)
		ret.Recv_client_command_starttls = in["recv_client_command_starttls"].(int)
		ret.Recv_client_command_rcpt = in["recv_client_command_rcpt"].(int)
	}
	return ret
}

func dataToEndpointSlbVirtualServerPortStats54(d *schema.ResourceData) edpt.SlbVirtualServerPortStats54 {
	var ret edpt.SlbVirtualServerPortStats54
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectSlbVirtualServerPortStats54Stats(d.Get("stats").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
