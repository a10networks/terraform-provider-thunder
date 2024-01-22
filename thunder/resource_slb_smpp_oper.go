package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmppOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_smpp_oper`: Operational Status for the object smpp\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSmppOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smpp_cpu_fields": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"msg_proxy_current": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mem_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mem_cached": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mem_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_send_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_incomplete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_connection": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_parse": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_process": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_exceed_tmp_buff": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_send_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_fail_start_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_send_success": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_incomplete": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_parse": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_process": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_selec_connt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_exceed_tmp_buff": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_fail_send_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_create_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_start_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_start_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_conn_fail_snat": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_construct_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_reserve_pconn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_start_server_conn_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_conn_already_exists": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_insert_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_parse_msg_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_process_msg_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_no_vport": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_select_server": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_alloc_mem": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_unexpected_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l7_cpu_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l4_to_l7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l4_from_l7": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_to_l4_send_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l4_from_l4_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_l7_to_l4": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_mag_back": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_fail_dcmsg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_deprecated_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_hold_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_split_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_pipline_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_client_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"msg_proxy_server_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"payload_allocd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"payload_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_too_small": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_seq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ax_response_directly": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_client_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_client_by_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_client_from_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_client_by_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_client_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_by_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_from_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_by_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"select_server_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bind_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unbind_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enquire_link_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enquire_link_resp_recv": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enquire_link_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enquire_link_resp_send": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_conn_put_in_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_conn_get_from_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_put_in_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_get_from_list": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_conn_fail_bind": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"single_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fail_bind_msg": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSmppOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmppOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmppOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSmppOperOper := setObjectSlbSmppOperOper(res)
		d.Set("oper", SlbSmppOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSmppOperOper(ret edpt.DataSlbSmppOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"smpp_cpu_fields": setSliceSlbSmppOperOperSmppCpuFields(ret.DtSlbSmppOper.Oper.SmppCpuFields),
			"cpu_count":       ret.DtSlbSmppOper.Oper.CpuCount,
			"filter_type":     ret.DtSlbSmppOper.Oper.Filter_type,
		},
	}
}

func setSliceSlbSmppOperOperSmppCpuFields(d []edpt.SlbSmppOperOperSmppCpuFields) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["msg_proxy_current"] = item.Msg_proxy_current
		in["msg_proxy_total"] = item.Msg_proxy_total
		in["msg_proxy_mem_allocd"] = item.Msg_proxy_mem_allocd
		in["msg_proxy_mem_cached"] = item.Msg_proxy_mem_cached
		in["msg_proxy_mem_freed"] = item.Msg_proxy_mem_freed
		in["msg_proxy_client_recv"] = item.Msg_proxy_client_recv
		in["msg_proxy_client_send_success"] = item.Msg_proxy_client_send_success
		in["msg_proxy_client_incomplete"] = item.Msg_proxy_client_incomplete
		in["msg_proxy_client_drop"] = item.Msg_proxy_client_drop
		in["msg_proxy_client_connection"] = item.Msg_proxy_client_connection
		in["msg_proxy_client_fail"] = item.Msg_proxy_client_fail
		in["msg_proxy_client_fail_parse"] = item.Msg_proxy_client_fail_parse
		in["msg_proxy_client_fail_process"] = item.Msg_proxy_client_fail_process
		in["msg_proxy_client_fail_snat"] = item.Msg_proxy_client_fail_snat
		in["msg_proxy_client_exceed_tmp_buff"] = item.Msg_proxy_client_exceed_tmp_buff
		in["msg_proxy_client_fail_send_pkt"] = item.Msg_proxy_client_fail_send_pkt
		in["msg_proxy_client_fail_start_server_conn"] = item.Msg_proxy_client_fail_start_server_conn
		in["msg_proxy_server_recv"] = item.Msg_proxy_server_recv
		in["msg_proxy_server_send_success"] = item.Msg_proxy_server_send_success
		in["msg_proxy_server_incomplete"] = item.Msg_proxy_server_incomplete
		in["msg_proxy_server_drop"] = item.Msg_proxy_server_drop
		in["msg_proxy_server_fail"] = item.Msg_proxy_server_fail
		in["msg_proxy_server_fail_parse"] = item.Msg_proxy_server_fail_parse
		in["msg_proxy_server_fail_process"] = item.Msg_proxy_server_fail_process
		in["msg_proxy_server_fail_selec_connt"] = item.Msg_proxy_server_fail_selec_connt
		in["msg_proxy_server_fail_snat"] = item.Msg_proxy_server_fail_snat
		in["msg_proxy_server_exceed_tmp_buff"] = item.Msg_proxy_server_exceed_tmp_buff
		in["msg_proxy_server_fail_send_pkt"] = item.Msg_proxy_server_fail_send_pkt
		in["msg_proxy_create_server_conn"] = item.Msg_proxy_create_server_conn
		in["msg_proxy_start_server_conn"] = item.Msg_proxy_start_server_conn
		in["msg_proxy_fail_start_server_conn"] = item.Msg_proxy_fail_start_server_conn
		in["msg_proxy_server_conn_fail_snat"] = item.Msg_proxy_server_conn_fail_snat
		in["msg_proxy_fail_construct_server_conn"] = item.Msg_proxy_fail_construct_server_conn
		in["msg_proxy_fail_reserve_pconn"] = item.Msg_proxy_fail_reserve_pconn
		in["msg_proxy_start_server_conn_failed"] = item.Msg_proxy_start_server_conn_failed
		in["msg_proxy_server_conn_already_exists"] = item.Msg_proxy_server_conn_already_exists
		in["msg_proxy_fail_insert_server_conn"] = item.Msg_proxy_fail_insert_server_conn
		in["msg_proxy_parse_msg_fail"] = item.Msg_proxy_parse_msg_fail
		in["msg_proxy_process_msg_fail"] = item.Msg_proxy_process_msg_fail
		in["msg_proxy_no_vport"] = item.Msg_proxy_no_vport
		in["msg_proxy_fail_select_server"] = item.Msg_proxy_fail_select_server
		in["msg_proxy_fail_alloc_mem"] = item.Msg_proxy_fail_alloc_mem
		in["msg_proxy_unexpected_err"] = item.Msg_proxy_unexpected_err
		in["msg_proxy_l7_cpu_failed"] = item.Msg_proxy_l7_cpu_failed
		in["msg_proxy_l4_to_l7"] = item.Msg_proxy_l4_to_l7
		in["msg_proxy_l4_from_l7"] = item.Msg_proxy_l4_from_l7
		in["msg_proxy_to_l4_send_pkt"] = item.Msg_proxy_to_l4_send_pkt
		in["msg_proxy_l4_from_l4_send"] = item.Msg_proxy_l4_from_l4_send
		in["msg_proxy_l7_to_l4"] = item.Msg_proxy_l7_to_l4
		in["msg_proxy_mag_back"] = item.Msg_proxy_mag_back
		in["msg_proxy_fail_dcmsg"] = item.Msg_proxy_fail_dcmsg
		in["msg_proxy_deprecated_conn"] = item.Msg_proxy_deprecated_conn
		in["msg_proxy_hold_msg"] = item.Msg_proxy_hold_msg
		in["msg_proxy_split_pkt"] = item.Msg_proxy_split_pkt
		in["msg_proxy_pipline_msg"] = item.Msg_proxy_pipline_msg
		in["msg_proxy_client_reset"] = item.Msg_proxy_client_reset
		in["msg_proxy_server_reset"] = item.Msg_proxy_server_reset
		in["payload_allocd"] = item.Payload_allocd
		in["payload_freed"] = item.Payload_freed
		in["pkt_too_small"] = item.Pkt_too_small
		in["invalid_seq"] = item.Invalid_seq
		in["ax_response_directly"] = item.Ax_response_directly
		in["select_client_conn"] = item.Select_client_conn
		in["select_client_by_req"] = item.Select_client_by_req
		in["select_client_from_list"] = item.Select_client_from_list
		in["select_client_by_conn"] = item.Select_client_by_conn
		in["select_client_fail"] = item.Select_client_fail
		in["select_server_conn"] = item.Select_server_conn
		in["select_server_by_req"] = item.Select_server_by_req
		in["select_server_from_list"] = item.Select_server_from_list
		in["select_server_by_conn"] = item.Select_server_by_conn
		in["select_server_fail"] = item.Select_server_fail
		in["bind_conn"] = item.Bind_conn
		in["unbind_conn"] = item.Unbind_conn
		in["enquire_link_recv"] = item.Enquire_link_recv
		in["enquire_link_resp_recv"] = item.Enquire_link_resp_recv
		in["enquire_link_send"] = item.Enquire_link_send
		in["enquire_link_resp_send"] = item.Enquire_link_resp_send
		in["client_conn_put_in_list"] = item.Client_conn_put_in_list
		in["client_conn_get_from_list"] = item.Client_conn_get_from_list
		in["server_conn_put_in_list"] = item.Server_conn_put_in_list
		in["server_conn_get_from_list"] = item.Server_conn_get_from_list
		in["server_conn_fail_bind"] = item.Server_conn_fail_bind
		in["single_msg"] = item.Single_msg
		in["fail_bind_msg"] = item.Fail_bind_msg
		result = append(result, in)
	}
	return result
}

func getObjectSlbSmppOperOper(d []interface{}) edpt.SlbSmppOperOper {

	count1 := len(d)
	var ret edpt.SlbSmppOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SmppCpuFields = getSliceSlbSmppOperOperSmppCpuFields(in["smpp_cpu_fields"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
		ret.Filter_type = in["filter_type"].(string)
	}
	return ret
}

func getSliceSlbSmppOperOperSmppCpuFields(d []interface{}) []edpt.SlbSmppOperOperSmppCpuFields {

	count1 := len(d)
	ret := make([]edpt.SlbSmppOperOperSmppCpuFields, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSmppOperOperSmppCpuFields
		oi.Msg_proxy_current = in["msg_proxy_current"].(int)
		oi.Msg_proxy_total = in["msg_proxy_total"].(int)
		oi.Msg_proxy_mem_allocd = in["msg_proxy_mem_allocd"].(int)
		oi.Msg_proxy_mem_cached = in["msg_proxy_mem_cached"].(int)
		oi.Msg_proxy_mem_freed = in["msg_proxy_mem_freed"].(int)
		oi.Msg_proxy_client_recv = in["msg_proxy_client_recv"].(int)
		oi.Msg_proxy_client_send_success = in["msg_proxy_client_send_success"].(int)
		oi.Msg_proxy_client_incomplete = in["msg_proxy_client_incomplete"].(int)
		oi.Msg_proxy_client_drop = in["msg_proxy_client_drop"].(int)
		oi.Msg_proxy_client_connection = in["msg_proxy_client_connection"].(int)
		oi.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		oi.Msg_proxy_client_fail_parse = in["msg_proxy_client_fail_parse"].(int)
		oi.Msg_proxy_client_fail_process = in["msg_proxy_client_fail_process"].(int)
		oi.Msg_proxy_client_fail_snat = in["msg_proxy_client_fail_snat"].(int)
		oi.Msg_proxy_client_exceed_tmp_buff = in["msg_proxy_client_exceed_tmp_buff"].(int)
		oi.Msg_proxy_client_fail_send_pkt = in["msg_proxy_client_fail_send_pkt"].(int)
		oi.Msg_proxy_client_fail_start_server_conn = in["msg_proxy_client_fail_start_server_conn"].(int)
		oi.Msg_proxy_server_recv = in["msg_proxy_server_recv"].(int)
		oi.Msg_proxy_server_send_success = in["msg_proxy_server_send_success"].(int)
		oi.Msg_proxy_server_incomplete = in["msg_proxy_server_incomplete"].(int)
		oi.Msg_proxy_server_drop = in["msg_proxy_server_drop"].(int)
		oi.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		oi.Msg_proxy_server_fail_parse = in["msg_proxy_server_fail_parse"].(int)
		oi.Msg_proxy_server_fail_process = in["msg_proxy_server_fail_process"].(int)
		oi.Msg_proxy_server_fail_selec_connt = in["msg_proxy_server_fail_selec_connt"].(int)
		oi.Msg_proxy_server_fail_snat = in["msg_proxy_server_fail_snat"].(int)
		oi.Msg_proxy_server_exceed_tmp_buff = in["msg_proxy_server_exceed_tmp_buff"].(int)
		oi.Msg_proxy_server_fail_send_pkt = in["msg_proxy_server_fail_send_pkt"].(int)
		oi.Msg_proxy_create_server_conn = in["msg_proxy_create_server_conn"].(int)
		oi.Msg_proxy_start_server_conn = in["msg_proxy_start_server_conn"].(int)
		oi.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		oi.Msg_proxy_server_conn_fail_snat = in["msg_proxy_server_conn_fail_snat"].(int)
		oi.Msg_proxy_fail_construct_server_conn = in["msg_proxy_fail_construct_server_conn"].(int)
		oi.Msg_proxy_fail_reserve_pconn = in["msg_proxy_fail_reserve_pconn"].(int)
		oi.Msg_proxy_start_server_conn_failed = in["msg_proxy_start_server_conn_failed"].(int)
		oi.Msg_proxy_server_conn_already_exists = in["msg_proxy_server_conn_already_exists"].(int)
		oi.Msg_proxy_fail_insert_server_conn = in["msg_proxy_fail_insert_server_conn"].(int)
		oi.Msg_proxy_parse_msg_fail = in["msg_proxy_parse_msg_fail"].(int)
		oi.Msg_proxy_process_msg_fail = in["msg_proxy_process_msg_fail"].(int)
		oi.Msg_proxy_no_vport = in["msg_proxy_no_vport"].(int)
		oi.Msg_proxy_fail_select_server = in["msg_proxy_fail_select_server"].(int)
		oi.Msg_proxy_fail_alloc_mem = in["msg_proxy_fail_alloc_mem"].(int)
		oi.Msg_proxy_unexpected_err = in["msg_proxy_unexpected_err"].(int)
		oi.Msg_proxy_l7_cpu_failed = in["msg_proxy_l7_cpu_failed"].(int)
		oi.Msg_proxy_l4_to_l7 = in["msg_proxy_l4_to_l7"].(int)
		oi.Msg_proxy_l4_from_l7 = in["msg_proxy_l4_from_l7"].(int)
		oi.Msg_proxy_to_l4_send_pkt = in["msg_proxy_to_l4_send_pkt"].(int)
		oi.Msg_proxy_l4_from_l4_send = in["msg_proxy_l4_from_l4_send"].(int)
		oi.Msg_proxy_l7_to_l4 = in["msg_proxy_l7_to_l4"].(int)
		oi.Msg_proxy_mag_back = in["msg_proxy_mag_back"].(int)
		oi.Msg_proxy_fail_dcmsg = in["msg_proxy_fail_dcmsg"].(int)
		oi.Msg_proxy_deprecated_conn = in["msg_proxy_deprecated_conn"].(int)
		oi.Msg_proxy_hold_msg = in["msg_proxy_hold_msg"].(int)
		oi.Msg_proxy_split_pkt = in["msg_proxy_split_pkt"].(int)
		oi.Msg_proxy_pipline_msg = in["msg_proxy_pipline_msg"].(int)
		oi.Msg_proxy_client_reset = in["msg_proxy_client_reset"].(int)
		oi.Msg_proxy_server_reset = in["msg_proxy_server_reset"].(int)
		oi.Payload_allocd = in["payload_allocd"].(int)
		oi.Payload_freed = in["payload_freed"].(int)
		oi.Pkt_too_small = in["pkt_too_small"].(int)
		oi.Invalid_seq = in["invalid_seq"].(int)
		oi.Ax_response_directly = in["ax_response_directly"].(int)
		oi.Select_client_conn = in["select_client_conn"].(int)
		oi.Select_client_by_req = in["select_client_by_req"].(int)
		oi.Select_client_from_list = in["select_client_from_list"].(int)
		oi.Select_client_by_conn = in["select_client_by_conn"].(int)
		oi.Select_client_fail = in["select_client_fail"].(int)
		oi.Select_server_conn = in["select_server_conn"].(int)
		oi.Select_server_by_req = in["select_server_by_req"].(int)
		oi.Select_server_from_list = in["select_server_from_list"].(int)
		oi.Select_server_by_conn = in["select_server_by_conn"].(int)
		oi.Select_server_fail = in["select_server_fail"].(int)
		oi.Bind_conn = in["bind_conn"].(int)
		oi.Unbind_conn = in["unbind_conn"].(int)
		oi.Enquire_link_recv = in["enquire_link_recv"].(int)
		oi.Enquire_link_resp_recv = in["enquire_link_resp_recv"].(int)
		oi.Enquire_link_send = in["enquire_link_send"].(int)
		oi.Enquire_link_resp_send = in["enquire_link_resp_send"].(int)
		oi.Client_conn_put_in_list = in["client_conn_put_in_list"].(int)
		oi.Client_conn_get_from_list = in["client_conn_get_from_list"].(int)
		oi.Server_conn_put_in_list = in["server_conn_put_in_list"].(int)
		oi.Server_conn_get_from_list = in["server_conn_get_from_list"].(int)
		oi.Server_conn_fail_bind = in["server_conn_fail_bind"].(int)
		oi.Single_msg = in["single_msg"].(int)
		oi.Fail_bind_msg = in["fail_bind_msg"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSmppOper(d *schema.ResourceData) edpt.SlbSmppOper {
	var ret edpt.SlbSmppOper

	ret.Oper = getObjectSlbSmppOperOper(d.Get("oper").([]interface{}))
	return ret
}
