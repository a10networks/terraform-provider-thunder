package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_tcp`: tcp counters and global config\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTcpCreate,
		UpdateContext: resourceSystemTcpUpdate,
		ReadContext:   resourceSystemTcpRead,
		DeleteContext: resourceSystemTcpDelete,

		Schema: map[string]*schema.Schema{
			"rate_limit_reset_unknown_conn": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pkt_rate_for_reset_unknown_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"log_for_reset_unknown_conn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log when rate exceed",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'activeopens': Active open conns; 'passiveopens': Passive open conns; 'attemptfails': Connect attemp failures; 'estabresets': Resets rcvd on EST conn; 'insegs': Total in TCP packets; 'outsegs': Total out TCP packets; 'retranssegs': Retransmited packets; 'inerrs': Input errors; 'outrsts': Reset Sent; 'sock_alloc': Sockets allocated; 'orphan_count': Orphan sockets; 'mem_alloc': Memory alloc; 'recv_mem': Total rx buffer; 'send_mem': Total tx buffer; 'currestab': Currently EST conns; 'currsyssnt': TCP in SYN-SNT state; 'currsynrcv': TCP in SYN-RCV state; 'currfinw1': TCP in FIN-W1 state; 'currfinw2': TCP FIN-W2 state; 'currtimew': TCP TimeW state; 'currclose': TCP in Close state; 'currclsw': TCP in CloseW state; 'currlack': TCP in LastACK state; 'currlstn': TCP in Listen state; 'currclsg': TCP in Closing state; 'pawsactiverejected': TCP paw active rej; 'syn_rcv_rstack': Rcv RST|ACK on SYN; 'syn_rcv_rst': Rcv RST on SYN; 'syn_rcv_ack': Rcv ACK on SYN; 'ax_rexmit_syn': TCP rexmit SYN; 'tcpabortontimeout': TCP abort on timeout; 'noroute': TCPIP out noroute; 'exceedmss': MSS exceeded pkt dropped; 'tfo_conns': TFO Total Connections; 'tfo_actives': TFO Current Actives; 'tfo_denied': TFO Denied;",
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
func resourceSystemTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemTcpRateLimitResetUnknownConn1646(d []interface{}) edpt.SystemTcpRateLimitResetUnknownConn1646 {

	count1 := len(d)
	var ret edpt.SystemTcpRateLimitResetUnknownConn1646
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PktRateForResetUnknownConn = in["pkt_rate_for_reset_unknown_conn"].(int)
		ret.LogForResetUnknownConn = in["log_for_reset_unknown_conn"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSystemTcpSamplingEnable(d []interface{}) []edpt.SystemTcpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemTcpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemTcp(d *schema.ResourceData) edpt.SystemTcp {
	var ret edpt.SystemTcp
	ret.Inst.RateLimitResetUnknownConn = getObjectSystemTcpRateLimitResetUnknownConn1646(d.Get("rate_limit_reset_unknown_conn").([]interface{}))
	ret.Inst.SamplingEnable = getSliceSystemTcpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
