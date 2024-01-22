package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTcpOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_tcp_oper`: Operational Status for the object tcp\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemTcpOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currestab": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activeopens": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"passiveopens": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"attemptfails": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"insegs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"outsegs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retranssegs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"estabresets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"outrsts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"noroute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tfo_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tfo_actives": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tfo_denied": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inerrs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sock_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"orphan_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mem_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_mem": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_mem": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currsyssnt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currsynrcv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currfinw1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currfinw2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currtimew": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currclose": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currclsw": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currlack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currlstn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"currclsg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pawsactiverejected": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_rcv_rstack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_rcv_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_rcv_ack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcpabortontimeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ax_rexmit_syn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"exceedmss": {
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
			"rate_limit_reset_unknown_conn": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"unknown_conn_rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unknown_conn_current_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unknown_conn_rate_limit_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceSystemTcpOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTcpOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTcpOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemTcpOperOper := setObjectSystemTcpOperOper(res)
		d.Set("oper", SystemTcpOperOper)
		SystemTcpOperRateLimitResetUnknownConn := setObjectSystemTcpOperRateLimitResetUnknownConn(res)
		d.Set("rate_limit_reset_unknown_conn", SystemTcpOperRateLimitResetUnknownConn)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemTcpOperOper(ret edpt.DataSystemTcpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_cpu_list": setSliceSystemTcpOperOperTcpCpuList(ret.DtSystemTcpOper.Oper.TcpCpuList),
			"cpu_count":    ret.DtSystemTcpOper.Oper.CpuCount,
		},
	}
}

func setSliceSystemTcpOperOperTcpCpuList(d []edpt.SystemTcpOperOperTcpCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["currestab"] = item.Currestab
		in["activeopens"] = item.Activeopens
		in["passiveopens"] = item.Passiveopens
		in["attemptfails"] = item.Attemptfails
		in["insegs"] = item.Insegs
		in["outsegs"] = item.Outsegs
		in["retranssegs"] = item.Retranssegs
		in["estabresets"] = item.Estabresets
		in["outrsts"] = item.Outrsts
		in["noroute"] = item.Noroute
		in["tfo_conns"] = item.Tfo_conns
		in["tfo_actives"] = item.Tfo_actives
		in["tfo_denied"] = item.Tfo_denied
		in["inerrs"] = item.Inerrs
		in["sock_alloc"] = item.Sock_alloc
		in["orphan_count"] = item.Orphan_count
		in["mem_alloc"] = item.Mem_alloc
		in["recv_mem"] = item.Recv_mem
		in["send_mem"] = item.Send_mem
		in["currsyssnt"] = item.Currsyssnt
		in["currsynrcv"] = item.Currsynrcv
		in["currfinw1"] = item.Currfinw1
		in["currfinw2"] = item.Currfinw2
		in["currtimew"] = item.Currtimew
		in["currclose"] = item.Currclose
		in["currclsw"] = item.Currclsw
		in["currlack"] = item.Currlack
		in["currlstn"] = item.Currlstn
		in["currclsg"] = item.Currclsg
		in["pawsactiverejected"] = item.Pawsactiverejected
		in["syn_rcv_rstack"] = item.Syn_rcv_rstack
		in["syn_rcv_rst"] = item.Syn_rcv_rst
		in["syn_rcv_ack"] = item.Syn_rcv_ack
		in["tcpabortontimeout"] = item.Tcpabortontimeout
		in["ax_rexmit_syn"] = item.Ax_rexmit_syn
		in["exceedmss"] = item.Exceedmss
		result = append(result, in)
	}
	return result
}

func setObjectSystemTcpOperRateLimitResetUnknownConn(ret edpt.DataSystemTcpOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectSystemTcpOperRateLimitResetUnknownConnOper(ret.DtSystemTcpOper.RateLimitResetUnknownConn.Oper),
		},
	}
}

func setObjectSystemTcpOperRateLimitResetUnknownConnOper(d edpt.SystemTcpOperRateLimitResetUnknownConnOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["unknown_conn_rate_limit"] = d.UnknownConnRateLimit

	in["unknown_conn_current_rate"] = d.UnknownConnCurrentRate

	in["unknown_conn_rate_limit_drop"] = d.UnknownConnRateLimitDrop
	result = append(result, in)
	return result
}

func getObjectSystemTcpOperOper(d []interface{}) edpt.SystemTcpOperOper {

	count1 := len(d)
	var ret edpt.SystemTcpOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpCpuList = getSliceSystemTcpOperOperTcpCpuList(in["tcp_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSystemTcpOperOperTcpCpuList(d []interface{}) []edpt.SystemTcpOperOperTcpCpuList {

	count1 := len(d)
	ret := make([]edpt.SystemTcpOperOperTcpCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemTcpOperOperTcpCpuList
		oi.Currestab = in["currestab"].(int)
		oi.Activeopens = in["activeopens"].(int)
		oi.Passiveopens = in["passiveopens"].(int)
		oi.Attemptfails = in["attemptfails"].(int)
		oi.Insegs = in["insegs"].(int)
		oi.Outsegs = in["outsegs"].(int)
		oi.Retranssegs = in["retranssegs"].(int)
		oi.Estabresets = in["estabresets"].(int)
		oi.Outrsts = in["outrsts"].(int)
		oi.Noroute = in["noroute"].(int)
		oi.Tfo_conns = in["tfo_conns"].(int)
		oi.Tfo_actives = in["tfo_actives"].(int)
		oi.Tfo_denied = in["tfo_denied"].(int)
		oi.Inerrs = in["inerrs"].(int)
		oi.Sock_alloc = in["sock_alloc"].(int)
		oi.Orphan_count = in["orphan_count"].(int)
		oi.Mem_alloc = in["mem_alloc"].(int)
		oi.Recv_mem = in["recv_mem"].(int)
		oi.Send_mem = in["send_mem"].(int)
		oi.Currsyssnt = in["currsyssnt"].(int)
		oi.Currsynrcv = in["currsynrcv"].(int)
		oi.Currfinw1 = in["currfinw1"].(int)
		oi.Currfinw2 = in["currfinw2"].(int)
		oi.Currtimew = in["currtimew"].(int)
		oi.Currclose = in["currclose"].(int)
		oi.Currclsw = in["currclsw"].(int)
		oi.Currlack = in["currlack"].(int)
		oi.Currlstn = in["currlstn"].(int)
		oi.Currclsg = in["currclsg"].(int)
		oi.Pawsactiverejected = in["pawsactiverejected"].(int)
		oi.Syn_rcv_rstack = in["syn_rcv_rstack"].(int)
		oi.Syn_rcv_rst = in["syn_rcv_rst"].(int)
		oi.Syn_rcv_ack = in["syn_rcv_ack"].(int)
		oi.Tcpabortontimeout = in["tcpabortontimeout"].(int)
		oi.Ax_rexmit_syn = in["ax_rexmit_syn"].(int)
		oi.Exceedmss = in["exceedmss"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSystemTcpOperRateLimitResetUnknownConn(d []interface{}) edpt.SystemTcpOperRateLimitResetUnknownConn {

	count1 := len(d)
	var ret edpt.SystemTcpOperRateLimitResetUnknownConn
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectSystemTcpOperRateLimitResetUnknownConnOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectSystemTcpOperRateLimitResetUnknownConnOper(d []interface{}) edpt.SystemTcpOperRateLimitResetUnknownConnOper {

	count1 := len(d)
	var ret edpt.SystemTcpOperRateLimitResetUnknownConnOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UnknownConnRateLimit = in["unknown_conn_rate_limit"].(int)
		ret.UnknownConnCurrentRate = in["unknown_conn_current_rate"].(int)
		ret.UnknownConnRateLimitDrop = in["unknown_conn_rate_limit_drop"].(int)
	}
	return ret
}

func dataToEndpointSystemTcpOper(d *schema.ResourceData) edpt.SystemTcpOper {
	var ret edpt.SystemTcpOper

	ret.Oper = getObjectSystemTcpOperOper(d.Get("oper").([]interface{}))

	ret.RateLimitResetUnknownConn = getObjectSystemTcpOperRateLimitResetUnknownConn(d.Get("rate_limit_reset_unknown_conn").([]interface{}))
	return ret
}
