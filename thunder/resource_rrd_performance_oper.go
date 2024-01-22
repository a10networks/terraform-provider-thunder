package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdPerformanceOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_performance_oper`: Operational Status for the object performance\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdPerformanceOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"end_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"performance_data": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_puts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_cur_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_new_conn_l4": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_new_conn_l7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_new_conn_ipnat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_new_conn_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_new_conn_srv_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_new_conn_tot": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_l7_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rus_c_conns_tot": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rus_c_conns_cur": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rus_s_conns_tot": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rus_s_conns_cur": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rus_s_conns_act": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_cps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_used_sess": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_avail_sess": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_tcp_port_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_tcp_port_avail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_udp_port_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"lsn_udp_port_avail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_tcp_cur_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_sctp_cur_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_udp_cur_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_ip_cur_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ov_other_cur_conns": {
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

func resourceRrdPerformanceOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdPerformanceOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdPerformanceOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdPerformanceOperOper := setObjectRrdPerformanceOperOper(res)
		d.Set("oper", RrdPerformanceOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdPerformanceOperOper(ret edpt.DataRrdPerformanceOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time":       ret.DtRrdPerformanceOper.Oper.StartTime,
			"end_time":         ret.DtRrdPerformanceOper.Oper.EndTime,
			"performance_data": setSliceRrdPerformanceOperOperPerformanceData(ret.DtRrdPerformanceOper.Oper.PerformanceData),
		},
	}
}

func setSliceRrdPerformanceOperOperPerformanceData(d []edpt.RrdPerformanceOperOperPerformanceData) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["ov_puts"] = item.Ov_puts
		in["ov_cur_conns"] = item.Ov_cur_conns
		in["ov_new_conn_l4"] = item.Ov_new_conn_l4
		in["ov_new_conn_l7"] = item.Ov_new_conn_l7
		in["ov_new_conn_ipnat"] = item.Ov_new_conn_ipnat
		in["ov_new_conn_ssl"] = item.Ov_new_conn_ssl
		in["ov_new_conn_srv_ssl"] = item.Ov_new_conn_srv_ssl
		in["ov_new_conn_tot"] = item.Ov_new_conn_tot
		in["ov_l7_req"] = item.Ov_l7_req
		in["rus_c_conns_tot"] = item.Rus_c_conns_tot
		in["rus_c_conns_cur"] = item.Rus_c_conns_cur
		in["rus_s_conns_tot"] = item.Rus_s_conns_tot
		in["rus_s_conns_cur"] = item.Rus_s_conns_cur
		in["rus_s_conns_act"] = item.Rus_s_conns_act
		in["syn_recv"] = item.Syn_recv
		in["syn_fail"] = item.Syn_fail
		in["lsn_cps"] = item.Lsn_cps
		in["lsn_used_sess"] = item.Lsn_used_sess
		in["lsn_avail_sess"] = item.Lsn_avail_sess
		in["lsn_tcp_port_used"] = item.Lsn_tcp_port_used
		in["lsn_tcp_port_avail"] = item.Lsn_tcp_port_avail
		in["lsn_udp_port_used"] = item.Lsn_udp_port_used
		in["lsn_udp_port_avail"] = item.Lsn_udp_port_avail
		in["ov_tcp_cur_conns"] = item.Ov_tcp_cur_conns
		in["ov_sctp_cur_conns"] = item.Ov_sctp_cur_conns
		in["ov_udp_cur_conns"] = item.Ov_udp_cur_conns
		in["ov_ip_cur_conns"] = item.Ov_ip_cur_conns
		in["ov_other_cur_conns"] = item.Ov_other_cur_conns
		result = append(result, in)
	}
	return result
}

func getObjectRrdPerformanceOperOper(d []interface{}) edpt.RrdPerformanceOperOper {

	count1 := len(d)
	var ret edpt.RrdPerformanceOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.PerformanceData = getSliceRrdPerformanceOperOperPerformanceData(in["performance_data"].([]interface{}))
	}
	return ret
}

func getSliceRrdPerformanceOperOperPerformanceData(d []interface{}) []edpt.RrdPerformanceOperOperPerformanceData {

	count1 := len(d)
	ret := make([]edpt.RrdPerformanceOperOperPerformanceData, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdPerformanceOperOperPerformanceData
		oi.Time = in["time"].(int)
		oi.Ov_puts = in["ov_puts"].(int)
		oi.Ov_cur_conns = in["ov_cur_conns"].(int)
		oi.Ov_new_conn_l4 = in["ov_new_conn_l4"].(int)
		oi.Ov_new_conn_l7 = in["ov_new_conn_l7"].(int)
		oi.Ov_new_conn_ipnat = in["ov_new_conn_ipnat"].(int)
		oi.Ov_new_conn_ssl = in["ov_new_conn_ssl"].(int)
		oi.Ov_new_conn_srv_ssl = in["ov_new_conn_srv_ssl"].(int)
		oi.Ov_new_conn_tot = in["ov_new_conn_tot"].(int)
		oi.Ov_l7_req = in["ov_l7_req"].(int)
		oi.Rus_c_conns_tot = in["rus_c_conns_tot"].(int)
		oi.Rus_c_conns_cur = in["rus_c_conns_cur"].(int)
		oi.Rus_s_conns_tot = in["rus_s_conns_tot"].(int)
		oi.Rus_s_conns_cur = in["rus_s_conns_cur"].(int)
		oi.Rus_s_conns_act = in["rus_s_conns_act"].(int)
		oi.Syn_recv = in["syn_recv"].(int)
		oi.Syn_fail = in["syn_fail"].(int)
		oi.Lsn_cps = in["lsn_cps"].(int)
		oi.Lsn_used_sess = in["lsn_used_sess"].(int)
		oi.Lsn_avail_sess = in["lsn_avail_sess"].(int)
		oi.Lsn_tcp_port_used = in["lsn_tcp_port_used"].(int)
		oi.Lsn_tcp_port_avail = in["lsn_tcp_port_avail"].(int)
		oi.Lsn_udp_port_used = in["lsn_udp_port_used"].(int)
		oi.Lsn_udp_port_avail = in["lsn_udp_port_avail"].(int)
		oi.Ov_tcp_cur_conns = in["ov_tcp_cur_conns"].(int)
		oi.Ov_sctp_cur_conns = in["ov_sctp_cur_conns"].(int)
		oi.Ov_udp_cur_conns = in["ov_udp_cur_conns"].(int)
		oi.Ov_ip_cur_conns = in["ov_ip_cur_conns"].(int)
		oi.Ov_other_cur_conns = in["ov_other_cur_conns"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdPerformanceOper(d *schema.ResourceData) edpt.RrdPerformanceOper {
	var ret edpt.RrdPerformanceOper

	ret.Oper = getObjectRrdPerformanceOperOper(d.Get("oper").([]interface{}))
	return ret
}
