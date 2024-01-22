package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbImapProxyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_imap_proxy_oper`: Operational Status for the object imap-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbImapProxyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"imap_proxy_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"current_proxy_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy_conns": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_imap_request": {
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
									"start_tls_cmd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"login_packet": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"capability_packet": {
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
									"imap_line_too_long": {
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
									"realloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"boundary_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"negative_error": {
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

func resourceSlbImapProxyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbImapProxyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbImapProxyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbImapProxyOperOper := setObjectSlbImapProxyOperOper(res)
		d.Set("oper", SlbImapProxyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbImapProxyOperOper(ret edpt.DataSlbImapProxyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"imap_proxy_cpu_list": setSliceSlbImapProxyOperOperImapProxyCpuList(ret.DtSlbImapProxyOper.Oper.ImapProxyCpuList),
			"cpu_count":           ret.DtSlbImapProxyOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbImapProxyOperOperImapProxyCpuList(d []edpt.SlbImapProxyOperOperImapProxyCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["current_proxy_conns"] = item.Current_proxy_conns
		in["total_proxy_conns"] = item.Total_proxy_conns
		in["total_imap_request"] = item.Total_imap_request
		in["server_selection_failure"] = item.Server_selection_failure
		in["no_route_failure"] = item.No_route_failure
		in["source_nat_failure"] = item.Source_nat_failure
		in["start_tls_cmd"] = item.Start_tls_cmd
		in["login_packet"] = item.Login_packet
		in["capability_packet"] = item.Capability_packet
		in["request_line_freed"] = item.Request_line_freed
		in["inv_start_line"] = item.Inv_start_line
		in["other_cmd"] = item.Other_cmd
		in["imap_line_too_long"] = item.Imap_line_too_long
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
		in["realloc_error"] = item.Realloc_error
		in["alloc_error"] = item.Alloc_error
		in["boundary_error"] = item.Boundary_error
		in["negative_error"] = item.Negative_error
		result = append(result, in)
	}
	return result
}

func getObjectSlbImapProxyOperOper(d []interface{}) edpt.SlbImapProxyOperOper {

	count1 := len(d)
	var ret edpt.SlbImapProxyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ImapProxyCpuList = getSliceSlbImapProxyOperOperImapProxyCpuList(in["imap_proxy_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbImapProxyOperOperImapProxyCpuList(d []interface{}) []edpt.SlbImapProxyOperOperImapProxyCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbImapProxyOperOperImapProxyCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbImapProxyOperOperImapProxyCpuList
		oi.Current_proxy_conns = in["current_proxy_conns"].(int)
		oi.Total_proxy_conns = in["total_proxy_conns"].(int)
		oi.Total_imap_request = in["total_imap_request"].(int)
		oi.Server_selection_failure = in["server_selection_failure"].(int)
		oi.No_route_failure = in["no_route_failure"].(int)
		oi.Source_nat_failure = in["source_nat_failure"].(int)
		oi.Start_tls_cmd = in["start_tls_cmd"].(int)
		oi.Login_packet = in["login_packet"].(int)
		oi.Capability_packet = in["capability_packet"].(int)
		oi.Request_line_freed = in["request_line_freed"].(int)
		oi.Inv_start_line = in["inv_start_line"].(int)
		oi.Other_cmd = in["other_cmd"].(int)
		oi.Imap_line_too_long = in["imap_line_too_long"].(int)
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
		oi.Realloc_error = in["realloc_error"].(int)
		oi.Alloc_error = in["alloc_error"].(int)
		oi.Boundary_error = in["boundary_error"].(int)
		oi.Negative_error = in["negative_error"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbImapProxyOper(d *schema.ResourceData) edpt.SlbImapProxyOper {
	var ret edpt.SlbImapProxyOper

	ret.Oper = getObjectSlbImapProxyOperOper(d.Get("oper").([]interface{}))
	return ret
}
