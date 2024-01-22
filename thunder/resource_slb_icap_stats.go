package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbIcapStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_icap_stats`: Statistics for the object icap\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbIcapStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reqmod_request": {
							Type: schema.TypeInt, Optional: true, Description: "Reqmod Request Stats",
						},
						"respmod_request": {
							Type: schema.TypeInt, Optional: true, Description: "Respmod Request Stats",
						},
						"reqmod_request_after_100": {
							Type: schema.TypeInt, Optional: true, Description: "Reqmod Request Sent After 100 Cont Stats",
						},
						"respmod_request_after_100": {
							Type: schema.TypeInt, Optional: true, Description: "Respmod Request Sent After 100 Cont Stats",
						},
						"reqmod_response": {
							Type: schema.TypeInt, Optional: true, Description: "Reqmod Response Stats",
						},
						"respmod_response": {
							Type: schema.TypeInt, Optional: true, Description: "Respmod Response Stats",
						},
						"reqmod_response_after_100": {
							Type: schema.TypeInt, Optional: true, Description: "Reqmod Response After 100 Cont Stats",
						},
						"respmod_response_after_100": {
							Type: schema.TypeInt, Optional: true, Description: "Respmod Response After 100 Cont Stats",
						},
						"chunk_no_allow_204": {
							Type: schema.TypeInt, Optional: true, Description: "Chunk so no Allow 204 Stats",
						},
						"len_exceed_no_allow_204": {
							Type: schema.TypeInt, Optional: true, Description: "Length Exceeded so no Allow 204 Stats",
						},
						"result_continue": {
							Type: schema.TypeInt, Optional: true, Description: "Result Continue Stats",
						},
						"result_icap_response": {
							Type: schema.TypeInt, Optional: true, Description: "Result ICAP Response Stats",
						},
						"result_100_continue": {
							Type: schema.TypeInt, Optional: true, Description: "Result 100 Continue Stats",
						},
						"result_other": {
							Type: schema.TypeInt, Optional: true, Description: "Result Other Stats",
						},
						"status_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status 2xx Stats",
						},
						"status_200": {
							Type: schema.TypeInt, Optional: true, Description: "Status 200 Stats",
						},
						"status_201": {
							Type: schema.TypeInt, Optional: true, Description: "Status 201 Stats",
						},
						"status_202": {
							Type: schema.TypeInt, Optional: true, Description: "Status 202 Stats",
						},
						"status_203": {
							Type: schema.TypeInt, Optional: true, Description: "Status 203 Stats",
						},
						"status_204": {
							Type: schema.TypeInt, Optional: true, Description: "Status 204 Stats",
						},
						"status_205": {
							Type: schema.TypeInt, Optional: true, Description: "Status 205 Stats",
						},
						"status_206": {
							Type: schema.TypeInt, Optional: true, Description: "Status 206 Stats",
						},
						"status_207": {
							Type: schema.TypeInt, Optional: true, Description: "Status 207 Stats",
						},
						"status_1xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status 1xx Stats",
						},
						"status_100": {
							Type: schema.TypeInt, Optional: true, Description: "Status 100 Stats",
						},
						"status_101": {
							Type: schema.TypeInt, Optional: true, Description: "Status 101 Stats",
						},
						"status_102": {
							Type: schema.TypeInt, Optional: true, Description: "Status 102 Stats",
						},
						"status_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status 3xx Stats",
						},
						"status_300": {
							Type: schema.TypeInt, Optional: true, Description: "Status 300 Stats",
						},
						"status_301": {
							Type: schema.TypeInt, Optional: true, Description: "Status 301 Stats",
						},
						"status_302": {
							Type: schema.TypeInt, Optional: true, Description: "Status 302 Stats",
						},
						"status_303": {
							Type: schema.TypeInt, Optional: true, Description: "Status 303 Stats",
						},
						"status_304": {
							Type: schema.TypeInt, Optional: true, Description: "Status 304 Stats",
						},
						"status_305": {
							Type: schema.TypeInt, Optional: true, Description: "Status 305 Stats",
						},
						"status_306": {
							Type: schema.TypeInt, Optional: true, Description: "Status 306 Stats",
						},
						"status_307": {
							Type: schema.TypeInt, Optional: true, Description: "Status 307 Stats",
						},
						"status_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status 4xx Stats",
						},
						"status_400": {
							Type: schema.TypeInt, Optional: true, Description: "Status 400 Stats",
						},
						"status_401": {
							Type: schema.TypeInt, Optional: true, Description: "Status 401 Stats",
						},
						"status_402": {
							Type: schema.TypeInt, Optional: true, Description: "Status 402 Stats",
						},
						"status_403": {
							Type: schema.TypeInt, Optional: true, Description: "Status 403 Stats",
						},
						"status_404": {
							Type: schema.TypeInt, Optional: true, Description: "Status 404 Stats",
						},
						"status_405": {
							Type: schema.TypeInt, Optional: true, Description: "Status 405 Stats",
						},
						"status_406": {
							Type: schema.TypeInt, Optional: true, Description: "Status 406 Stats",
						},
						"status_407": {
							Type: schema.TypeInt, Optional: true, Description: "Status 407 Stats",
						},
						"status_408": {
							Type: schema.TypeInt, Optional: true, Description: "Status 408 Stats",
						},
						"status_409": {
							Type: schema.TypeInt, Optional: true, Description: "Status 409 Stats",
						},
						"status_410": {
							Type: schema.TypeInt, Optional: true, Description: "Status 410 Stats",
						},
						"status_411": {
							Type: schema.TypeInt, Optional: true, Description: "Status 411 Stats",
						},
						"status_412": {
							Type: schema.TypeInt, Optional: true, Description: "Status 412 Stats",
						},
						"status_413": {
							Type: schema.TypeInt, Optional: true, Description: "Status 413 Stats",
						},
						"status_414": {
							Type: schema.TypeInt, Optional: true, Description: "Status 414 Stats",
						},
						"status_415": {
							Type: schema.TypeInt, Optional: true, Description: "Status 415 Stats",
						},
						"status_416": {
							Type: schema.TypeInt, Optional: true, Description: "Status 416 Stats",
						},
						"status_417": {
							Type: schema.TypeInt, Optional: true, Description: "Status 417 Stats",
						},
						"status_418": {
							Type: schema.TypeInt, Optional: true, Description: "Status 418 Stats",
						},
						"status_419": {
							Type: schema.TypeInt, Optional: true, Description: "Status 419 Stats",
						},
						"status_420": {
							Type: schema.TypeInt, Optional: true, Description: "Status 420 Stats",
						},
						"status_422": {
							Type: schema.TypeInt, Optional: true, Description: "Status 422 Stats",
						},
						"status_423": {
							Type: schema.TypeInt, Optional: true, Description: "Status 423 Stats",
						},
						"status_424": {
							Type: schema.TypeInt, Optional: true, Description: "Status 424 Stats",
						},
						"status_425": {
							Type: schema.TypeInt, Optional: true, Description: "Status 425 Stats",
						},
						"status_426": {
							Type: schema.TypeInt, Optional: true, Description: "Status 426 Stats",
						},
						"status_449": {
							Type: schema.TypeInt, Optional: true, Description: "Status 449 Stats",
						},
						"status_450": {
							Type: schema.TypeInt, Optional: true, Description: "Status 450 Stats",
						},
						"status_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status 5xx Stats",
						},
						"status_500": {
							Type: schema.TypeInt, Optional: true, Description: "Status 500 Stats",
						},
						"status_501": {
							Type: schema.TypeInt, Optional: true, Description: "Status 501 Stats",
						},
						"status_502": {
							Type: schema.TypeInt, Optional: true, Description: "Status 502 Stats",
						},
						"status_503": {
							Type: schema.TypeInt, Optional: true, Description: "Status 503 Stats",
						},
						"status_504": {
							Type: schema.TypeInt, Optional: true, Description: "Status 504 Stats",
						},
						"status_505": {
							Type: schema.TypeInt, Optional: true, Description: "Status 505 Stats",
						},
						"status_506": {
							Type: schema.TypeInt, Optional: true, Description: "Status 506 Stats",
						},
						"status_507": {
							Type: schema.TypeInt, Optional: true, Description: "Status 507 Stats",
						},
						"status_508": {
							Type: schema.TypeInt, Optional: true, Description: "Status 508 Stats",
						},
						"status_509": {
							Type: schema.TypeInt, Optional: true, Description: "Status 509 Stats",
						},
						"status_510": {
							Type: schema.TypeInt, Optional: true, Description: "Status 510 Stats",
						},
						"status_6xx": {
							Type: schema.TypeInt, Optional: true, Description: "Status 6xx Stats",
						},
						"status_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Status Unknown Stats",
						},
						"send_option_req": {
							Type: schema.TypeInt, Optional: true, Description: "Send Option Req Stats",
						},
						"app_serv_conn_no_pcb_err": {
							Type: schema.TypeInt, Optional: true, Description: "App Server Conn no ES PCB Err Stats",
						},
						"app_serv_conn_err": {
							Type: schema.TypeInt, Optional: true, Description: "App Server Conn Err Stats",
						},
						"chunk1_hdr_err": {
							Type: schema.TypeInt, Optional: true, Description: "Chunk Hdr Err1 Stats",
						},
						"chunk2_hdr_err": {
							Type: schema.TypeInt, Optional: true, Description: "Chunk Hdr Err2 Stats",
						},
						"chunk_bad_trail_err": {
							Type: schema.TypeInt, Optional: true, Description: "Chunk Bad Trail Err Stats",
						},
						"no_payload_next_buff_err": {
							Type: schema.TypeInt, Optional: true, Description: "No Payload In Next Buff Err Stats",
						},
						"no_payload_buff_err": {
							Type: schema.TypeInt, Optional: true, Description: "No Payload Buff Err Stats",
						},
						"resp_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Description: "Resp Hdr Incomplete Err Stats",
						},
						"serv_sel_fail_err": {
							Type: schema.TypeInt, Optional: true, Description: "Server Select Fail Err Stats",
						},
						"start_icap_conn_fail_err": {
							Type: schema.TypeInt, Optional: true, Description: "Start ICAP conn fail Stats",
						},
						"prep_req_fail_err": {
							Type: schema.TypeInt, Optional: true, Description: "Prepare ICAP req fail Err Stats",
						},
						"icap_ver_err": {
							Type: schema.TypeInt, Optional: true, Description: "ICAP Ver Err Stats",
						},
						"icap_line_err": {
							Type: schema.TypeInt, Optional: true, Description: "ICAP Line Err Stats",
						},
						"encap_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Description: "Encap HDR Incomplete Err Stats",
						},
						"no_icap_resp_err": {
							Type: schema.TypeInt, Optional: true, Description: "No ICAP Resp Err Stats",
						},
						"resp_line_read_err": {
							Type: schema.TypeInt, Optional: true, Description: "Resp Line Read Err Stats",
						},
						"resp_line_parse_err": {
							Type: schema.TypeInt, Optional: true, Description: "Resp Line Parse Err Stats",
						},
						"resp_hdr_err": {
							Type: schema.TypeInt, Optional: true, Description: "Resp Hdr Err Stats",
						},
						"req_hdr_incomplete_err": {
							Type: schema.TypeInt, Optional: true, Description: "Req Hdr Incomplete Err Stats",
						},
						"no_status_code_err": {
							Type: schema.TypeInt, Optional: true, Description: "No Status Code Err Stats",
						},
						"http_resp_line_read_err": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Response Line Read Err Stats",
						},
						"http_resp_line_parse_err": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Response Line Parse Err Stats",
						},
						"http_resp_hdr_err": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Resp Hdr Err Stats",
						},
						"recv_option_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Send Option Req Stats",
						},
						"http_connect_reqmod_request": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP CONNECT Reqmod Request Stats",
						},
					},
				},
			},
		},
	}
}

func resourceSlbIcapStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbIcapStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbIcapStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbIcapStatsStats := setObjectSlbIcapStatsStats(res)
		d.Set("stats", SlbIcapStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbIcapStatsStats(ret edpt.DataSlbIcapStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"reqmod_request":              ret.DtSlbIcapStats.Stats.Reqmod_request,
			"respmod_request":             ret.DtSlbIcapStats.Stats.Respmod_request,
			"reqmod_request_after_100":    ret.DtSlbIcapStats.Stats.Reqmod_request_after_100,
			"respmod_request_after_100":   ret.DtSlbIcapStats.Stats.Respmod_request_after_100,
			"reqmod_response":             ret.DtSlbIcapStats.Stats.Reqmod_response,
			"respmod_response":            ret.DtSlbIcapStats.Stats.Respmod_response,
			"reqmod_response_after_100":   ret.DtSlbIcapStats.Stats.Reqmod_response_after_100,
			"respmod_response_after_100":  ret.DtSlbIcapStats.Stats.Respmod_response_after_100,
			"chunk_no_allow_204":          ret.DtSlbIcapStats.Stats.Chunk_no_allow_204,
			"len_exceed_no_allow_204":     ret.DtSlbIcapStats.Stats.Len_exceed_no_allow_204,
			"result_continue":             ret.DtSlbIcapStats.Stats.Result_continue,
			"result_icap_response":        ret.DtSlbIcapStats.Stats.Result_icap_response,
			"result_100_continue":         ret.DtSlbIcapStats.Stats.Result_100_continue,
			"result_other":                ret.DtSlbIcapStats.Stats.Result_other,
			"status_2xx":                  ret.DtSlbIcapStats.Stats.Status_2xx,
			"status_200":                  ret.DtSlbIcapStats.Stats.Status_200,
			"status_201":                  ret.DtSlbIcapStats.Stats.Status_201,
			"status_202":                  ret.DtSlbIcapStats.Stats.Status_202,
			"status_203":                  ret.DtSlbIcapStats.Stats.Status_203,
			"status_204":                  ret.DtSlbIcapStats.Stats.Status_204,
			"status_205":                  ret.DtSlbIcapStats.Stats.Status_205,
			"status_206":                  ret.DtSlbIcapStats.Stats.Status_206,
			"status_207":                  ret.DtSlbIcapStats.Stats.Status_207,
			"status_1xx":                  ret.DtSlbIcapStats.Stats.Status_1xx,
			"status_100":                  ret.DtSlbIcapStats.Stats.Status_100,
			"status_101":                  ret.DtSlbIcapStats.Stats.Status_101,
			"status_102":                  ret.DtSlbIcapStats.Stats.Status_102,
			"status_3xx":                  ret.DtSlbIcapStats.Stats.Status_3xx,
			"status_300":                  ret.DtSlbIcapStats.Stats.Status_300,
			"status_301":                  ret.DtSlbIcapStats.Stats.Status_301,
			"status_302":                  ret.DtSlbIcapStats.Stats.Status_302,
			"status_303":                  ret.DtSlbIcapStats.Stats.Status_303,
			"status_304":                  ret.DtSlbIcapStats.Stats.Status_304,
			"status_305":                  ret.DtSlbIcapStats.Stats.Status_305,
			"status_306":                  ret.DtSlbIcapStats.Stats.Status_306,
			"status_307":                  ret.DtSlbIcapStats.Stats.Status_307,
			"status_4xx":                  ret.DtSlbIcapStats.Stats.Status_4xx,
			"status_400":                  ret.DtSlbIcapStats.Stats.Status_400,
			"status_401":                  ret.DtSlbIcapStats.Stats.Status_401,
			"status_402":                  ret.DtSlbIcapStats.Stats.Status_402,
			"status_403":                  ret.DtSlbIcapStats.Stats.Status_403,
			"status_404":                  ret.DtSlbIcapStats.Stats.Status_404,
			"status_405":                  ret.DtSlbIcapStats.Stats.Status_405,
			"status_406":                  ret.DtSlbIcapStats.Stats.Status_406,
			"status_407":                  ret.DtSlbIcapStats.Stats.Status_407,
			"status_408":                  ret.DtSlbIcapStats.Stats.Status_408,
			"status_409":                  ret.DtSlbIcapStats.Stats.Status_409,
			"status_410":                  ret.DtSlbIcapStats.Stats.Status_410,
			"status_411":                  ret.DtSlbIcapStats.Stats.Status_411,
			"status_412":                  ret.DtSlbIcapStats.Stats.Status_412,
			"status_413":                  ret.DtSlbIcapStats.Stats.Status_413,
			"status_414":                  ret.DtSlbIcapStats.Stats.Status_414,
			"status_415":                  ret.DtSlbIcapStats.Stats.Status_415,
			"status_416":                  ret.DtSlbIcapStats.Stats.Status_416,
			"status_417":                  ret.DtSlbIcapStats.Stats.Status_417,
			"status_418":                  ret.DtSlbIcapStats.Stats.Status_418,
			"status_419":                  ret.DtSlbIcapStats.Stats.Status_419,
			"status_420":                  ret.DtSlbIcapStats.Stats.Status_420,
			"status_422":                  ret.DtSlbIcapStats.Stats.Status_422,
			"status_423":                  ret.DtSlbIcapStats.Stats.Status_423,
			"status_424":                  ret.DtSlbIcapStats.Stats.Status_424,
			"status_425":                  ret.DtSlbIcapStats.Stats.Status_425,
			"status_426":                  ret.DtSlbIcapStats.Stats.Status_426,
			"status_449":                  ret.DtSlbIcapStats.Stats.Status_449,
			"status_450":                  ret.DtSlbIcapStats.Stats.Status_450,
			"status_5xx":                  ret.DtSlbIcapStats.Stats.Status_5xx,
			"status_500":                  ret.DtSlbIcapStats.Stats.Status_500,
			"status_501":                  ret.DtSlbIcapStats.Stats.Status_501,
			"status_502":                  ret.DtSlbIcapStats.Stats.Status_502,
			"status_503":                  ret.DtSlbIcapStats.Stats.Status_503,
			"status_504":                  ret.DtSlbIcapStats.Stats.Status_504,
			"status_505":                  ret.DtSlbIcapStats.Stats.Status_505,
			"status_506":                  ret.DtSlbIcapStats.Stats.Status_506,
			"status_507":                  ret.DtSlbIcapStats.Stats.Status_507,
			"status_508":                  ret.DtSlbIcapStats.Stats.Status_508,
			"status_509":                  ret.DtSlbIcapStats.Stats.Status_509,
			"status_510":                  ret.DtSlbIcapStats.Stats.Status_510,
			"status_6xx":                  ret.DtSlbIcapStats.Stats.Status_6xx,
			"status_unknown":              ret.DtSlbIcapStats.Stats.Status_unknown,
			"send_option_req":             ret.DtSlbIcapStats.Stats.Send_option_req,
			"app_serv_conn_no_pcb_err":    ret.DtSlbIcapStats.Stats.App_serv_conn_no_pcb_err,
			"app_serv_conn_err":           ret.DtSlbIcapStats.Stats.App_serv_conn_err,
			"chunk1_hdr_err":              ret.DtSlbIcapStats.Stats.Chunk1_hdr_err,
			"chunk2_hdr_err":              ret.DtSlbIcapStats.Stats.Chunk2_hdr_err,
			"chunk_bad_trail_err":         ret.DtSlbIcapStats.Stats.Chunk_bad_trail_err,
			"no_payload_next_buff_err":    ret.DtSlbIcapStats.Stats.No_payload_next_buff_err,
			"no_payload_buff_err":         ret.DtSlbIcapStats.Stats.No_payload_buff_err,
			"resp_hdr_incomplete_err":     ret.DtSlbIcapStats.Stats.Resp_hdr_incomplete_err,
			"serv_sel_fail_err":           ret.DtSlbIcapStats.Stats.Serv_sel_fail_err,
			"start_icap_conn_fail_err":    ret.DtSlbIcapStats.Stats.Start_icap_conn_fail_err,
			"prep_req_fail_err":           ret.DtSlbIcapStats.Stats.Prep_req_fail_err,
			"icap_ver_err":                ret.DtSlbIcapStats.Stats.Icap_ver_err,
			"icap_line_err":               ret.DtSlbIcapStats.Stats.Icap_line_err,
			"encap_hdr_incomplete_err":    ret.DtSlbIcapStats.Stats.Encap_hdr_incomplete_err,
			"no_icap_resp_err":            ret.DtSlbIcapStats.Stats.No_icap_resp_err,
			"resp_line_read_err":          ret.DtSlbIcapStats.Stats.Resp_line_read_err,
			"resp_line_parse_err":         ret.DtSlbIcapStats.Stats.Resp_line_parse_err,
			"resp_hdr_err":                ret.DtSlbIcapStats.Stats.Resp_hdr_err,
			"req_hdr_incomplete_err":      ret.DtSlbIcapStats.Stats.Req_hdr_incomplete_err,
			"no_status_code_err":          ret.DtSlbIcapStats.Stats.No_status_code_err,
			"http_resp_line_read_err":     ret.DtSlbIcapStats.Stats.Http_resp_line_read_err,
			"http_resp_line_parse_err":    ret.DtSlbIcapStats.Stats.Http_resp_line_parse_err,
			"http_resp_hdr_err":           ret.DtSlbIcapStats.Stats.Http_resp_hdr_err,
			"recv_option_resp":            ret.DtSlbIcapStats.Stats.Recv_option_resp,
			"http_connect_reqmod_request": ret.DtSlbIcapStats.Stats.Http_connect_reqmod_request,
		},
	}
}

func getObjectSlbIcapStatsStats(d []interface{}) edpt.SlbIcapStatsStats {

	count1 := len(d)
	var ret edpt.SlbIcapStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Reqmod_request = in["reqmod_request"].(int)
		ret.Respmod_request = in["respmod_request"].(int)
		ret.Reqmod_request_after_100 = in["reqmod_request_after_100"].(int)
		ret.Respmod_request_after_100 = in["respmod_request_after_100"].(int)
		ret.Reqmod_response = in["reqmod_response"].(int)
		ret.Respmod_response = in["respmod_response"].(int)
		ret.Reqmod_response_after_100 = in["reqmod_response_after_100"].(int)
		ret.Respmod_response_after_100 = in["respmod_response_after_100"].(int)
		ret.Chunk_no_allow_204 = in["chunk_no_allow_204"].(int)
		ret.Len_exceed_no_allow_204 = in["len_exceed_no_allow_204"].(int)
		ret.Result_continue = in["result_continue"].(int)
		ret.Result_icap_response = in["result_icap_response"].(int)
		ret.Result_100_continue = in["result_100_continue"].(int)
		ret.Result_other = in["result_other"].(int)
		ret.Status_2xx = in["status_2xx"].(int)
		ret.Status_200 = in["status_200"].(int)
		ret.Status_201 = in["status_201"].(int)
		ret.Status_202 = in["status_202"].(int)
		ret.Status_203 = in["status_203"].(int)
		ret.Status_204 = in["status_204"].(int)
		ret.Status_205 = in["status_205"].(int)
		ret.Status_206 = in["status_206"].(int)
		ret.Status_207 = in["status_207"].(int)
		ret.Status_1xx = in["status_1xx"].(int)
		ret.Status_100 = in["status_100"].(int)
		ret.Status_101 = in["status_101"].(int)
		ret.Status_102 = in["status_102"].(int)
		ret.Status_3xx = in["status_3xx"].(int)
		ret.Status_300 = in["status_300"].(int)
		ret.Status_301 = in["status_301"].(int)
		ret.Status_302 = in["status_302"].(int)
		ret.Status_303 = in["status_303"].(int)
		ret.Status_304 = in["status_304"].(int)
		ret.Status_305 = in["status_305"].(int)
		ret.Status_306 = in["status_306"].(int)
		ret.Status_307 = in["status_307"].(int)
		ret.Status_4xx = in["status_4xx"].(int)
		ret.Status_400 = in["status_400"].(int)
		ret.Status_401 = in["status_401"].(int)
		ret.Status_402 = in["status_402"].(int)
		ret.Status_403 = in["status_403"].(int)
		ret.Status_404 = in["status_404"].(int)
		ret.Status_405 = in["status_405"].(int)
		ret.Status_406 = in["status_406"].(int)
		ret.Status_407 = in["status_407"].(int)
		ret.Status_408 = in["status_408"].(int)
		ret.Status_409 = in["status_409"].(int)
		ret.Status_410 = in["status_410"].(int)
		ret.Status_411 = in["status_411"].(int)
		ret.Status_412 = in["status_412"].(int)
		ret.Status_413 = in["status_413"].(int)
		ret.Status_414 = in["status_414"].(int)
		ret.Status_415 = in["status_415"].(int)
		ret.Status_416 = in["status_416"].(int)
		ret.Status_417 = in["status_417"].(int)
		ret.Status_418 = in["status_418"].(int)
		ret.Status_419 = in["status_419"].(int)
		ret.Status_420 = in["status_420"].(int)
		ret.Status_422 = in["status_422"].(int)
		ret.Status_423 = in["status_423"].(int)
		ret.Status_424 = in["status_424"].(int)
		ret.Status_425 = in["status_425"].(int)
		ret.Status_426 = in["status_426"].(int)
		ret.Status_449 = in["status_449"].(int)
		ret.Status_450 = in["status_450"].(int)
		ret.Status_5xx = in["status_5xx"].(int)
		ret.Status_500 = in["status_500"].(int)
		ret.Status_501 = in["status_501"].(int)
		ret.Status_502 = in["status_502"].(int)
		ret.Status_503 = in["status_503"].(int)
		ret.Status_504 = in["status_504"].(int)
		ret.Status_505 = in["status_505"].(int)
		ret.Status_506 = in["status_506"].(int)
		ret.Status_507 = in["status_507"].(int)
		ret.Status_508 = in["status_508"].(int)
		ret.Status_509 = in["status_509"].(int)
		ret.Status_510 = in["status_510"].(int)
		ret.Status_6xx = in["status_6xx"].(int)
		ret.Status_unknown = in["status_unknown"].(int)
		ret.Send_option_req = in["send_option_req"].(int)
		ret.App_serv_conn_no_pcb_err = in["app_serv_conn_no_pcb_err"].(int)
		ret.App_serv_conn_err = in["app_serv_conn_err"].(int)
		ret.Chunk1_hdr_err = in["chunk1_hdr_err"].(int)
		ret.Chunk2_hdr_err = in["chunk2_hdr_err"].(int)
		ret.Chunk_bad_trail_err = in["chunk_bad_trail_err"].(int)
		ret.No_payload_next_buff_err = in["no_payload_next_buff_err"].(int)
		ret.No_payload_buff_err = in["no_payload_buff_err"].(int)
		ret.Resp_hdr_incomplete_err = in["resp_hdr_incomplete_err"].(int)
		ret.Serv_sel_fail_err = in["serv_sel_fail_err"].(int)
		ret.Start_icap_conn_fail_err = in["start_icap_conn_fail_err"].(int)
		ret.Prep_req_fail_err = in["prep_req_fail_err"].(int)
		ret.Icap_ver_err = in["icap_ver_err"].(int)
		ret.Icap_line_err = in["icap_line_err"].(int)
		ret.Encap_hdr_incomplete_err = in["encap_hdr_incomplete_err"].(int)
		ret.No_icap_resp_err = in["no_icap_resp_err"].(int)
		ret.Resp_line_read_err = in["resp_line_read_err"].(int)
		ret.Resp_line_parse_err = in["resp_line_parse_err"].(int)
		ret.Resp_hdr_err = in["resp_hdr_err"].(int)
		ret.Req_hdr_incomplete_err = in["req_hdr_incomplete_err"].(int)
		ret.No_status_code_err = in["no_status_code_err"].(int)
		ret.Http_resp_line_read_err = in["http_resp_line_read_err"].(int)
		ret.Http_resp_line_parse_err = in["http_resp_line_parse_err"].(int)
		ret.Http_resp_hdr_err = in["http_resp_hdr_err"].(int)
		ret.Recv_option_resp = in["recv_option_resp"].(int)
		ret.Http_connect_reqmod_request = in["http_connect_reqmod_request"].(int)
	}
	return ret
}

func dataToEndpointSlbIcapStats(d *schema.ResourceData) edpt.SlbIcapStats {
	var ret edpt.SlbIcapStats

	ret.Stats = getObjectSlbIcapStatsStats(d.Get("stats").([]interface{}))
	return ret
}
