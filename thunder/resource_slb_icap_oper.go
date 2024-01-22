package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIcapOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_icap_oper`: Operational Status for the object icap\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbIcapOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icap_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"reqmod_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"respmod_request": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reqmod_request_after_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"respmod_request_after_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reqmod_response": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"respmod_response": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reqmod_response_after_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"respmod_response_after_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_no_allow_204": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"len_exceed_no_allow_204": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"result_continue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"result_icap_response": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"result_100_continue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"result_other": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_200": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_201": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_202": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_203": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_204": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_205": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_206": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_207": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_100": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_101": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_102": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_300": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_301": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_302": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_303": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_304": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_305": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_306": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_307": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_400": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_401": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_402": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_403": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_404": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_405": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_406": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_407": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_408": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_409": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_410": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_411": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_412": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_413": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_414": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_415": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_416": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_417": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_418": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_419": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_420": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_422": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_423": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_424": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_425": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_426": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_449": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_450": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_500": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_501": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_502": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_503": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_504": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_505": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_506": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_507": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_508": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_509": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_510": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_6xx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"status_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"send_option_req": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_serv_conn_no_pcb_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"app_serv_conn_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk1_hdr_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk2_hdr_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"chunk_bad_trail_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_payload_next_buff_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_payload_buff_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resp_hdr_incomplete_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"serv_sel_fail_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"start_icap_conn_fail_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"prep_req_fail_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icap_ver_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"icap_line_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"encap_hdr_incomplete_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_icap_resp_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resp_line_read_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resp_line_parse_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resp_hdr_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_hdr_incomplete_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_status_code_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_resp_line_read_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_resp_line_parse_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_resp_hdr_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"recv_option_resp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_connect_reqmod_request": {
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

func resourceSlbIcapOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIcapOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIcapOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbIcapOperOper := setObjectSlbIcapOperOper(res)
		d.Set("oper", SlbIcapOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbIcapOperOper(ret edpt.DataSlbIcapOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"icap_cpu_list": setSliceSlbIcapOperOperIcapCpuList(ret.DtSlbIcapOper.Oper.IcapCpuList),
			"cpu_count":     ret.DtSlbIcapOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbIcapOperOperIcapCpuList(d []edpt.SlbIcapOperOperIcapCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["reqmod_request"] = item.Reqmod_request
		in["respmod_request"] = item.Respmod_request
		in["reqmod_request_after_100"] = item.Reqmod_request_after_100
		in["respmod_request_after_100"] = item.Respmod_request_after_100
		in["reqmod_response"] = item.Reqmod_response
		in["respmod_response"] = item.Respmod_response
		in["reqmod_response_after_100"] = item.Reqmod_response_after_100
		in["respmod_response_after_100"] = item.Respmod_response_after_100
		in["chunk_no_allow_204"] = item.Chunk_no_allow_204
		in["len_exceed_no_allow_204"] = item.Len_exceed_no_allow_204
		in["result_continue"] = item.Result_continue
		in["result_icap_response"] = item.Result_icap_response
		in["result_100_continue"] = item.Result_100_continue
		in["result_other"] = item.Result_other
		in["status_2xx"] = item.Status_2xx
		in["status_200"] = item.Status_200
		in["status_201"] = item.Status_201
		in["status_202"] = item.Status_202
		in["status_203"] = item.Status_203
		in["status_204"] = item.Status_204
		in["status_205"] = item.Status_205
		in["status_206"] = item.Status_206
		in["status_207"] = item.Status_207
		in["status_1xx"] = item.Status_1xx
		in["status_100"] = item.Status_100
		in["status_101"] = item.Status_101
		in["status_102"] = item.Status_102
		in["status_3xx"] = item.Status_3xx
		in["status_300"] = item.Status_300
		in["status_301"] = item.Status_301
		in["status_302"] = item.Status_302
		in["status_303"] = item.Status_303
		in["status_304"] = item.Status_304
		in["status_305"] = item.Status_305
		in["status_306"] = item.Status_306
		in["status_307"] = item.Status_307
		in["status_4xx"] = item.Status_4xx
		in["status_400"] = item.Status_400
		in["status_401"] = item.Status_401
		in["status_402"] = item.Status_402
		in["status_403"] = item.Status_403
		in["status_404"] = item.Status_404
		in["status_405"] = item.Status_405
		in["status_406"] = item.Status_406
		in["status_407"] = item.Status_407
		in["status_408"] = item.Status_408
		in["status_409"] = item.Status_409
		in["status_410"] = item.Status_410
		in["status_411"] = item.Status_411
		in["status_412"] = item.Status_412
		in["status_413"] = item.Status_413
		in["status_414"] = item.Status_414
		in["status_415"] = item.Status_415
		in["status_416"] = item.Status_416
		in["status_417"] = item.Status_417
		in["status_418"] = item.Status_418
		in["status_419"] = item.Status_419
		in["status_420"] = item.Status_420
		in["status_422"] = item.Status_422
		in["status_423"] = item.Status_423
		in["status_424"] = item.Status_424
		in["status_425"] = item.Status_425
		in["status_426"] = item.Status_426
		in["status_449"] = item.Status_449
		in["status_450"] = item.Status_450
		in["status_5xx"] = item.Status_5xx
		in["status_500"] = item.Status_500
		in["status_501"] = item.Status_501
		in["status_502"] = item.Status_502
		in["status_503"] = item.Status_503
		in["status_504"] = item.Status_504
		in["status_505"] = item.Status_505
		in["status_506"] = item.Status_506
		in["status_507"] = item.Status_507
		in["status_508"] = item.Status_508
		in["status_509"] = item.Status_509
		in["status_510"] = item.Status_510
		in["status_6xx"] = item.Status_6xx
		in["status_unknown"] = item.Status_unknown
		in["send_option_req"] = item.Send_option_req
		in["app_serv_conn_no_pcb_err"] = item.App_serv_conn_no_pcb_err
		in["app_serv_conn_err"] = item.App_serv_conn_err
		in["chunk1_hdr_err"] = item.Chunk1_hdr_err
		in["chunk2_hdr_err"] = item.Chunk2_hdr_err
		in["chunk_bad_trail_err"] = item.Chunk_bad_trail_err
		in["no_payload_next_buff_err"] = item.No_payload_next_buff_err
		in["no_payload_buff_err"] = item.No_payload_buff_err
		in["resp_hdr_incomplete_err"] = item.Resp_hdr_incomplete_err
		in["serv_sel_fail_err"] = item.Serv_sel_fail_err
		in["start_icap_conn_fail_err"] = item.Start_icap_conn_fail_err
		in["prep_req_fail_err"] = item.Prep_req_fail_err
		in["icap_ver_err"] = item.Icap_ver_err
		in["icap_line_err"] = item.Icap_line_err
		in["encap_hdr_incomplete_err"] = item.Encap_hdr_incomplete_err
		in["no_icap_resp_err"] = item.No_icap_resp_err
		in["resp_line_read_err"] = item.Resp_line_read_err
		in["resp_line_parse_err"] = item.Resp_line_parse_err
		in["resp_hdr_err"] = item.Resp_hdr_err
		in["req_hdr_incomplete_err"] = item.Req_hdr_incomplete_err
		in["no_status_code_err"] = item.No_status_code_err
		in["http_resp_line_read_err"] = item.Http_resp_line_read_err
		in["http_resp_line_parse_err"] = item.Http_resp_line_parse_err
		in["http_resp_hdr_err"] = item.Http_resp_hdr_err
		in["recv_option_resp"] = item.Recv_option_resp
		in["http_connect_reqmod_request"] = item.Http_connect_reqmod_request
		result = append(result, in)
	}
	return result
}

func getObjectSlbIcapOperOper(d []interface{}) edpt.SlbIcapOperOper {

	count1 := len(d)
	var ret edpt.SlbIcapOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcapCpuList = getSliceSlbIcapOperOperIcapCpuList(in["icap_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbIcapOperOperIcapCpuList(d []interface{}) []edpt.SlbIcapOperOperIcapCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbIcapOperOperIcapCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbIcapOperOperIcapCpuList
		oi.Reqmod_request = in["reqmod_request"].(int)
		oi.Respmod_request = in["respmod_request"].(int)
		oi.Reqmod_request_after_100 = in["reqmod_request_after_100"].(int)
		oi.Respmod_request_after_100 = in["respmod_request_after_100"].(int)
		oi.Reqmod_response = in["reqmod_response"].(int)
		oi.Respmod_response = in["respmod_response"].(int)
		oi.Reqmod_response_after_100 = in["reqmod_response_after_100"].(int)
		oi.Respmod_response_after_100 = in["respmod_response_after_100"].(int)
		oi.Chunk_no_allow_204 = in["chunk_no_allow_204"].(int)
		oi.Len_exceed_no_allow_204 = in["len_exceed_no_allow_204"].(int)
		oi.Result_continue = in["result_continue"].(int)
		oi.Result_icap_response = in["result_icap_response"].(int)
		oi.Result_100_continue = in["result_100_continue"].(int)
		oi.Result_other = in["result_other"].(int)
		oi.Status_2xx = in["status_2xx"].(int)
		oi.Status_200 = in["status_200"].(int)
		oi.Status_201 = in["status_201"].(int)
		oi.Status_202 = in["status_202"].(int)
		oi.Status_203 = in["status_203"].(int)
		oi.Status_204 = in["status_204"].(int)
		oi.Status_205 = in["status_205"].(int)
		oi.Status_206 = in["status_206"].(int)
		oi.Status_207 = in["status_207"].(int)
		oi.Status_1xx = in["status_1xx"].(int)
		oi.Status_100 = in["status_100"].(int)
		oi.Status_101 = in["status_101"].(int)
		oi.Status_102 = in["status_102"].(int)
		oi.Status_3xx = in["status_3xx"].(int)
		oi.Status_300 = in["status_300"].(int)
		oi.Status_301 = in["status_301"].(int)
		oi.Status_302 = in["status_302"].(int)
		oi.Status_303 = in["status_303"].(int)
		oi.Status_304 = in["status_304"].(int)
		oi.Status_305 = in["status_305"].(int)
		oi.Status_306 = in["status_306"].(int)
		oi.Status_307 = in["status_307"].(int)
		oi.Status_4xx = in["status_4xx"].(int)
		oi.Status_400 = in["status_400"].(int)
		oi.Status_401 = in["status_401"].(int)
		oi.Status_402 = in["status_402"].(int)
		oi.Status_403 = in["status_403"].(int)
		oi.Status_404 = in["status_404"].(int)
		oi.Status_405 = in["status_405"].(int)
		oi.Status_406 = in["status_406"].(int)
		oi.Status_407 = in["status_407"].(int)
		oi.Status_408 = in["status_408"].(int)
		oi.Status_409 = in["status_409"].(int)
		oi.Status_410 = in["status_410"].(int)
		oi.Status_411 = in["status_411"].(int)
		oi.Status_412 = in["status_412"].(int)
		oi.Status_413 = in["status_413"].(int)
		oi.Status_414 = in["status_414"].(int)
		oi.Status_415 = in["status_415"].(int)
		oi.Status_416 = in["status_416"].(int)
		oi.Status_417 = in["status_417"].(int)
		oi.Status_418 = in["status_418"].(int)
		oi.Status_419 = in["status_419"].(int)
		oi.Status_420 = in["status_420"].(int)
		oi.Status_422 = in["status_422"].(int)
		oi.Status_423 = in["status_423"].(int)
		oi.Status_424 = in["status_424"].(int)
		oi.Status_425 = in["status_425"].(int)
		oi.Status_426 = in["status_426"].(int)
		oi.Status_449 = in["status_449"].(int)
		oi.Status_450 = in["status_450"].(int)
		oi.Status_5xx = in["status_5xx"].(int)
		oi.Status_500 = in["status_500"].(int)
		oi.Status_501 = in["status_501"].(int)
		oi.Status_502 = in["status_502"].(int)
		oi.Status_503 = in["status_503"].(int)
		oi.Status_504 = in["status_504"].(int)
		oi.Status_505 = in["status_505"].(int)
		oi.Status_506 = in["status_506"].(int)
		oi.Status_507 = in["status_507"].(int)
		oi.Status_508 = in["status_508"].(int)
		oi.Status_509 = in["status_509"].(int)
		oi.Status_510 = in["status_510"].(int)
		oi.Status_6xx = in["status_6xx"].(int)
		oi.Status_unknown = in["status_unknown"].(int)
		oi.Send_option_req = in["send_option_req"].(int)
		oi.App_serv_conn_no_pcb_err = in["app_serv_conn_no_pcb_err"].(int)
		oi.App_serv_conn_err = in["app_serv_conn_err"].(int)
		oi.Chunk1_hdr_err = in["chunk1_hdr_err"].(int)
		oi.Chunk2_hdr_err = in["chunk2_hdr_err"].(int)
		oi.Chunk_bad_trail_err = in["chunk_bad_trail_err"].(int)
		oi.No_payload_next_buff_err = in["no_payload_next_buff_err"].(int)
		oi.No_payload_buff_err = in["no_payload_buff_err"].(int)
		oi.Resp_hdr_incomplete_err = in["resp_hdr_incomplete_err"].(int)
		oi.Serv_sel_fail_err = in["serv_sel_fail_err"].(int)
		oi.Start_icap_conn_fail_err = in["start_icap_conn_fail_err"].(int)
		oi.Prep_req_fail_err = in["prep_req_fail_err"].(int)
		oi.Icap_ver_err = in["icap_ver_err"].(int)
		oi.Icap_line_err = in["icap_line_err"].(int)
		oi.Encap_hdr_incomplete_err = in["encap_hdr_incomplete_err"].(int)
		oi.No_icap_resp_err = in["no_icap_resp_err"].(int)
		oi.Resp_line_read_err = in["resp_line_read_err"].(int)
		oi.Resp_line_parse_err = in["resp_line_parse_err"].(int)
		oi.Resp_hdr_err = in["resp_hdr_err"].(int)
		oi.Req_hdr_incomplete_err = in["req_hdr_incomplete_err"].(int)
		oi.No_status_code_err = in["no_status_code_err"].(int)
		oi.Http_resp_line_read_err = in["http_resp_line_read_err"].(int)
		oi.Http_resp_line_parse_err = in["http_resp_line_parse_err"].(int)
		oi.Http_resp_hdr_err = in["http_resp_hdr_err"].(int)
		oi.Recv_option_resp = in["recv_option_resp"].(int)
		oi.Http_connect_reqmod_request = in["http_connect_reqmod_request"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbIcapOper(d *schema.ResourceData) edpt.SlbIcapOper {
	var ret edpt.SlbIcapOper

	ret.Oper = getObjectSlbIcapOperOper(d.Get("oper").([]interface{}))
	return ret
}
