package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_smtp`: Configure SMTP\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSmtpCreate,
		UpdateContext: resourceSlbSmtpUpdate,
		ReadContext:   resourceSlbSmtpRead,
		DeleteContext: resourceSlbSmtpDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_proxy': Current proxy conns; 'total_proxy': Total proxy conns; 'request': SMTP requests; 'request_success': SMTP requests (success); 'no_proxy': No proxy error; 'client_reset': Client reset; 'server_reset': Server reset; 'no_tuple': No tuple error; 'parse_req_fail': Parse request failure; 'server_select_fail': Server selection failure; 'forward_req_fail': Forward request failure; 'forward_req_data_fail': Forward REQ data failure; 'req_retran': Request retransmit; 'req_ofo': Request pkt out-of-order; 'server_reselect': Server reselection; 'server_prem_close': Server premature close; 'new_server_conn': Server connection made; 'snat_fail': Source NAT failure; 'tcp_out_reset': TCP out reset; 'Aflex_switch': aFleX switching; 'Aflex_switch_ok': aFleX switching (succ); 'recv_client_command_EHLO': Recv client EHLO; 'recv_client_command_HELO': Recv client HELO; 'recv_client_command_MAIL': Recv client MAIL; 'recv_client_command_RCPT': Recv client RCPT; 'recv_client_command_DATA': Recv client DATA; 'recv_client_command_RSET': Recv client RSET; 'recv_client_command_VRFY': Recv client VRFY; 'recv_client_command_EXPN': Recv client EXPN; 'recv_client_command_HELP': Recv client HELP; 'recv_client_command_NOOP': Recv client NOOP; 'recv_client_command_QUIT': Recv client QUIT; 'recv_client_command_STARTTLS': Recv client STARTTLS; 'recv_client_command_others': Recv client other cmds; 'send_client_service_ready': Sent client serv-rdy; 'send_client_service_not_ready': Sent client serv-not-rdy; 'send_client_close_connection': Sent client close-conn; 'send_client_go_ahead': Sent client go-ahead; 'send_client_start_TLS_first': Sent client STARTTLS-1st; 'send_client_TLS_not_available': Sent client TLS-not-aval; 'send_client_no_command': Sent client no-such-cmd; 'send_server_cmd_reset': Sent server RSET; 'TLS_established': SSL session established; 'L4_switch': L4 switching; 'recv_server_service_not_ready': Recv server serv-not-rdy; 'recv_server_unknow_reply_code': Recv server unknown-code; 'client_domain_switch': Client domain switching; 'client_domain_switch_ok': Client domain sw (succ); 'LB_switch': LB switching; 'LB_switch_ok': LB switching (succ); 'read_request_line_fail': Read request line fail; 'get_all_headers_fail': Get all headers fail; 'too_many_headers': Too many headers; 'line_too_long': Line too long; 'line_across_packet': Line across packets; 'line_extend': Line extend; 'line_extend_fail': Line extend fail; 'line_table_extend': Table extend; 'line_table_extend_fail': Table extend fail; 'parse_request_line_fail': Parse request line fail; 'insert_resonse_line_fail': Ins response line fail; 'remove_resonse_line_fail': Del response line fail; 'parse_resonse_line_fail': Parse response line fail; 'Aflex_lb_reselect': aFleX lb reselect; 'Aflex_lb_reselect_ok': aFleX lb reselect (succ); 'server_STARTTLS_init': Init server side STARTTLS; 'server_STARTTLS_fail': Server side STARTTLS fail; 'rserver_STARTTLS_disable': real server not support STARTTLS; 'recv_client_command_TURN': Recv client TURN; 'recv_client_command_ETRN': Recv client ETRN; 'send_server_ehlo': Proxy sends server EHLO; 'fail_to_save_client_ehlo': Failed to save client EHLO; 'aflex_mail_fail': aFlex Mail event failed; 'drop_server_ehlo_ok': Server EHLO_OK dropped; 'client_ehlo_saved': Client EHLO saved;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbSmtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSmtpRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSmtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSmtpRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSmtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSmtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSmtpSamplingEnable(d []interface{}) []edpt.SlbSmtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSmtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSmtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSmtp(d *schema.ResourceData) edpt.SlbSmtp {
	var ret edpt.SlbSmtp
	ret.Inst.SamplingEnable = getSliceSlbSmtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
