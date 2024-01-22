package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbPop3ProxyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_pop3_proxy_oper`: Operational Status for the object pop3-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbPop3ProxyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"current_proxy_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_selection_failure": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_route_failure": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"source_nat_failure": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stls_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_line_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inv_start_line": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"other_cmd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pop3_line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_chn_ssl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_seq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_sel_persist_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_sel_smpv6_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_sel_smpv4_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_sel_ins_tpl_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_est_state_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_ctng_state_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_resp_state_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rq_state_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_pop3_request": {
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

func resourceSlbPop3ProxyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbPop3ProxyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbPop3ProxyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbPop3ProxyOperOper := setObjectSlbPop3ProxyOperOper(res)
		d.Set("oper", SlbPop3ProxyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbPop3ProxyOperOper(ret edpt.DataSlbPop3ProxyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l4_cpu_list": setSliceSlbPop3ProxyOperOperL4CpuList(ret.DtSlbPop3ProxyOper.Oper.L4CpuList),
			"cpu_count":   ret.DtSlbPop3ProxyOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbPop3ProxyOperOperL4CpuList(d []edpt.SlbPop3ProxyOperOperL4CpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["current_proxy_conns"] = item.Current_proxy_conns
		in["total_proxy_conns"] = item.Total_proxy_conns
		in["server_selection_failure"] = item.Server_selection_failure
		in["no_route_failure"] = item.No_route_failure
		in["source_nat_failure"] = item.Source_nat_failure
		in["stls_packet"] = item.Stls_packet
		in["request_line_freed"] = item.Request_line_freed
		in["inv_start_line"] = item.Inv_start_line
		in["other_cmd"] = item.Other_cmd
		in["pop3_line_too_long"] = item.Pop3_line_too_long
		in["control_chn_ssl"] = item.Control_chn_ssl
		in["bad_seq"] = item.Bad_seq
		in["serv_sel_persist_fail"] = item.Serv_sel_persist_fail
		in["serv_sel_smpv6_fail"] = item.Serv_sel_smpv6_fail
		in["serv_sel_smpv4_fail"] = item.Serv_sel_smpv4_fail
		in["serv_sel_ins_tpl_fail"] = item.Serv_sel_ins_tpl_fail
		in["client_est_state_err"] = item.Client_est_state_err
		in["serv_ctng_state_err"] = item.Serv_ctng_state_err
		in["serv_resp_state_err"] = item.Serv_resp_state_err
		in["client_rq_state_err"] = item.Client_rq_state_err
		in["total_pop3_request"] = item.Total_pop3_request
		result = append(result, in)
	}
	return result
}

func getObjectSlbPop3ProxyOperOper(d []interface{}) edpt.SlbPop3ProxyOperOper {

	count1 := len(d)
	var ret edpt.SlbPop3ProxyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4CpuList = getSliceSlbPop3ProxyOperOperL4CpuList(in["l4_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbPop3ProxyOperOperL4CpuList(d []interface{}) []edpt.SlbPop3ProxyOperOperL4CpuList {

	count1 := len(d)
	ret := make([]edpt.SlbPop3ProxyOperOperL4CpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbPop3ProxyOperOperL4CpuList
		oi.Current_proxy_conns = in["current_proxy_conns"].(int)
		oi.Total_proxy_conns = in["total_proxy_conns"].(int)
		oi.Server_selection_failure = in["server_selection_failure"].(int)
		oi.No_route_failure = in["no_route_failure"].(int)
		oi.Source_nat_failure = in["source_nat_failure"].(int)
		oi.Stls_packet = in["stls_packet"].(int)
		oi.Request_line_freed = in["request_line_freed"].(int)
		oi.Inv_start_line = in["inv_start_line"].(int)
		oi.Other_cmd = in["other_cmd"].(int)
		oi.Pop3_line_too_long = in["pop3_line_too_long"].(int)
		oi.Control_chn_ssl = in["control_chn_ssl"].(int)
		oi.Bad_seq = in["bad_seq"].(int)
		oi.Serv_sel_persist_fail = in["serv_sel_persist_fail"].(int)
		oi.Serv_sel_smpv6_fail = in["serv_sel_smpv6_fail"].(int)
		oi.Serv_sel_smpv4_fail = in["serv_sel_smpv4_fail"].(int)
		oi.Serv_sel_ins_tpl_fail = in["serv_sel_ins_tpl_fail"].(int)
		oi.Client_est_state_err = in["client_est_state_err"].(int)
		oi.Serv_ctng_state_err = in["serv_ctng_state_err"].(int)
		oi.Serv_resp_state_err = in["serv_resp_state_err"].(int)
		oi.Client_rq_state_err = in["client_rq_state_err"].(int)
		oi.Total_pop3_request = in["total_pop3_request"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbPop3ProxyOper(d *schema.ResourceData) edpt.SlbPop3ProxyOper {
	var ret edpt.SlbPop3ProxyOper

	ret.Oper = getObjectSlbPop3ProxyOperOper(d.Get("oper").([]interface{}))
	return ret
}
