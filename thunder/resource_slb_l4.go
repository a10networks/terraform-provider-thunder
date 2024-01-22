package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbL4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_l4`: Configure L4\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbL4Create,
		UpdateContext: resourceSlbL4Update,
		ReadContext:   resourceSlbL4Read,
		DeleteContext: resourceSlbL4Delete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'intcp': TCP received; 'synreceived': TCP SYN received; 'tcp_fwd_last_ack': L4 rcv fwd last ACK; 'tcp_rev_last_ack': L4 rcv rev last ACK; 'tcp_rev_fin': L4 rcv rev FIN; 'tcp_fwd_fin': L4 rcv fwd FIN; 'tcp_fwd_ackfin': L4 rcv fwd FIN|ACK; 'inudp': UDP received; 'syncookiessent': TCP SYN cookie snt; 'syncookiessent_ts': TCP SYN cookie snt ts; 'syncookiessentfailed': TCP SYN cookie snt fail; 'outrst': TCP out RST; 'outrst_nosyn': TCP out RST no SYN; 'outrst_broker': TCP out RST L4 proxy; 'outrst_ack_attack': TCP out RST ACK attack; 'outrst_aflex': TCP out RST aFleX; 'outrst_stale_sess': TCP out RST stale sess; 'syn_stale_sess': SYN stale sess drop; 'outrst_tcpproxy': TCP out RST TCP proxy; 'svrselfail': Server sel failure; 'noroute': IP out noroute; 'snat_fail': Source NAT failure; 'snat_no_fwd_route': Source NAT no fwd route; 'snat_no_rev_route': Source NAT no rev route; 'snat_icmp_error_process': Source NAT ICMP Process; 'snat_icmp_no_match': Source NAT ICMP No Match; 'smart_nat_id_mismatch': Auto NAT id mismatch; 'syncookiescheckfailed': TCP SYN cookie failed; 'novport_drop': NAT no session drops; 'no_vport_drop': vport not matching drops; 'nosyn_drop': No SYN pkt drops; 'nosyn_drop_fin': No SYN pkt drops - FIN; 'nosyn_drop_rst': No SYN pkt drops - RST; 'nosyn_drop_ack': No SYN pkt drops - ACK; 'connlimit_drop': Conn Limit drops; 'connlimit_reset': Conn Limit resets; 'conn_rate_limit_drop': Conn rate limit drops; 'conn_rate_limit_reset': Conn rate limit resets; 'proxy_nosock_drop': Proxy no sock drops; 'drop_aflex': aFleX drops; 'sess_aged_out': Session aged out; 'tcp_sess_aged_out': TCP Session aged out; 'udp_sess_aged_out': UDP Session aged out; 'other_sess_aged_out': Other Session aged out; 'tcp_no_slb': TCP no SLB; 'udp_no_slb': UDP no SLB; 'throttle_syn': SYN Throttle; 'drop_gslb': Drop GSLB; 'inband_hm_retry': Inband HM retry; 'inband_hm_reassign': Inband HM reassign; 'auto_reassign': Auto-reselect server; 'fast_aging_set': Fast aging set; 'fast_aging_reset': Fast aging reset; 'dns_policy_drop': DNS Policy Drop; 'tcp_invalid_drop': TCP invalid drop; 'anomaly_out_seq': Anomaly out of sequence; 'anomaly_zero_win': Anomaly zero window; 'anomaly_bad_content': Anomaly bad content; 'anomaly_pbslb_drop': Anomaly pbslb drop; 'no_resourse_drop': No resource drop; 'reset_unknown_conn': Reset unknown conn; 'reset_l7_on_failover': RST L7 on failover; 'ignore_msl': ignore msl; 'l2_dsr': L2 DSR received; 'l3_dsr': L3 DSR received; 'port_preserve_attempt': NAT Port Preserve Try; 'port_preserve_succ': NAT Port Preserve Succ; 'tcpsyndata_drop': TCP SYN With Data Drop; 'tcpotherflags_drop': TCP SYN Other Flags Drop; 'bw_rate_limit_exceed': BW-Limit Exceed drop; 'bw_watermark_drop': BW-Watermark drop; 'l4_cps_exceed': L4 CPS exceed drop; 'nat_cps_exceed': NAT CPS exceed drop; 'l7_cps_exceed': L7 CPS exceed drop; 'ssl_cps_exceed': SSL CPS exceed drop; 'ssl_tpt_exceed': SSL TPT exceed drop; 'ssl_watermark_drop': SSL TPT-Watermark drop; 'concurrent_conn_exceed': L3V Conn Limit Drop; 'svr_syn_handshake_fail': L4 server handshake fail; 'stateless_conn_timeout': L4 stateless Conn TO; 'tcp_ax_rexmit_syn': L4 AX re-xmit SYN; 'tcp_syn_rcv_ack': L4 rcv ACK on SYN; 'tcp_syn_rcv_rst': L4 rcv RST on SYN; 'tcp_sess_noest_aged_out': TCP no-Est Sess aged out; 'tcp_sess_noest_csyn_rcv_aged_out': no-Est CSYN rcv aged out; 'tcp_sess_noest_ssyn_xmit_aged_out': no-Est SSYN snt aged out; 'tcp_rexmit_syn': L4 rcv rexmit SYN; 'tcp_rexmit_syn_delq': L4 rcv rexmit SYN (delq); 'tcp_rexmit_synack': L4 rcv rexmit SYN|ACK; 'tcp_rexmit_synack_delq': L4 rcv rexmit SYN|ACK DQ; 'tcp_fwd_fin_dup': L4 rcv fwd FIN dup; 'tcp_rev_fin_dup': L4 rcv rev FIN dup; 'tcp_rev_ackfin': L4 rcv rev FIN|ACK; 'tcp_fwd_rst': L4 rcv fwd RST; 'tcp_rev_rst': L4 rcv rev RST; 'udp_req_oneplus_no_resp': L4 UDP reqs no rsp; 'udp_req_one_oneplus_resp': L4 UDP req rsps; 'udp_req_resp_notmatch': L4 UDP req/rsp not match; 'udp_req_more_resp': L4 UDP req greater than rsps; 'udp_resp_more_req': L4 UDP rsps greater than reqs; 'udp_req_oneplus': L4 UDP reqs; 'udp_resp_oneplus': L4 UDP rsps; 'out_seq_ack_drop': Out of sequence ACK drop; 'tcp_est': L4 TCP Established; 'synattack': L4 SYN attack; 'syn_rate': TCP SYN rate per sec; 'syncookie_buff_drop': TCP SYN cookie buff drop; 'syncookie_buff_queue': TCP SYN cookie buff queue; 'skip_insert_client_ip': Skip Insert-client-ip; 'synreceived_hw': TCP SYN (HW SYN cookie); 'dns_id_switch': DNS query id switch; 'server_down_del': Server Down Del switch; 'dnssec_switch': DNSSEC SG switch; 'rate_drop_reset_unkn': Rate Drop reset; 'tcp_connections_closed': TCP Connections Closed; 'gtp_c_invalid_port': Invalid Packet Received on GTP VIP; 'gtp_c_invalid_header': Invalid Header Received on GTP VIP; 'gtp_c_invalid_message': Non Create Session/PDP Context Request/Response Received on GTP VIP; 'reselect_svrselfail': Server reselect failure; 'snat_port_overload_fail': Snat port overload fail; 'snat_force_preserve_alloc': Snat port preserve allocated; 'snat_force_preserve_free': Snat port preserve freed; 'proxy_header_insert': PROXY protocol header inserted; 'proxy_header_rexmit': PROXY protocol header retransmitted; 'proxy_prot_error': PROXY protocol error; 'proxy_prot_drop': PROXY protocol drop; 'slb_gtp_proxy_pkt_rcv_rr': SLB GTP proxy packet received on RR; 'slb_gtp_proxy_smp_match': SLB GTP proxy helper session found; 'slb_gtp_proxy_smp_no_match': SLB GTP proxy helper session not found; 'slb_gtp_proxy_c_process_local_rr': SLB GTP proxy messageprocessed locally on RR; 'slb_gtp_proxy_smp_creation_failed': SLB GTP proxy helper session creation failed; 'slb_gtp_proxy_smp_created': SLB GTP proxy helper session created; 'slb_gtp_proxy_smp_free_not_found': SLB GTP proxy session helper not found during cleanup; 'slb_gtp_proxy_smp_freed': SLB GTP proxy session helper freed; 'slb_gtp_proxy_retx_requests': SLB GTP proxy retx requests; 'pbslb_entry_limit_exceed': pbslb entry limit Exceed; 'fast_path_reroute': Fast Path Reroute; 'fast_path_l2_reroute': Fast Path L2 Reroute;",
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
func resourceSlbL4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbL4Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbL4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbL4Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbL4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbL4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbL4SamplingEnable(d []interface{}) []edpt.SlbL4SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbL4SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbL4SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbL4(d *schema.ResourceData) edpt.SlbL4 {
	var ret edpt.SlbL4
	ret.Inst.SamplingEnable = getSliceSlbL4SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
