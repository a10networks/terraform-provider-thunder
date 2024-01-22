package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHmDplaneStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_hm_dplane_stats`: Statistics for the object hm-dplane\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHmDplaneStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Current HM Entries",
						},
						"total_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total HM Entries Created",
						},
						"total_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Total HM entries inserted",
						},
						"total_ready_to_free": {
							Type: schema.TypeInt, Optional: true, Description: "Total HM entries ready to free",
						},
						"total_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total HM entries freed",
						},
						"err_entry_create_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Creation Failure",
						},
						"err_entry_create_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Entry creation out of memory",
						},
						"err_entry_insert_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry insert failed",
						},
						"total_tcp_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP errors in health-checks sent",
						},
						"err_smart_nat_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Error creating smart-nat instance",
						},
						"err_smart_nat_port_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Error obtaining smart-nat source port",
						},
						"err_l4_sess_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Error allocating L4 session for HM",
						},
						"err_hm_tcp_conn_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Error in initiating TCP connection for HM",
						},
						"hm_tcp_conn_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP connections sent for HM",
						},
						"entry_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Entry deleted",
						},
						"err_entry_create_vip_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Error in creating HM internal VIP",
						},
						"total_match_resp_code": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP received response with match response code",
						},
						"total_match_default_resp_code": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP received response with match 200 response code",
						},
						"total_maintenance_received": {
							Type: schema.TypeInt, Optional: true, Description: "Total maintenace response received",
						},
						"total_wrong_status_received": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP received response with wrong response code",
						},
						"err_no_hm_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Error no HM entry found",
						},
						"err_ssl_cert_name_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Error SSL cert name mismatch",
						},
						"err_server_syn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Error SSL server SYN timeout",
						},
						"err_http2_callback": {
							Type: schema.TypeInt, Optional: true, Description: "Error HTTP2 callback",
						},
						"err_l7_sess_process_tcp_estab_failed": {
							Type: schema.TypeInt, Optional: true, Description: "L7 session process TCP established failed",
						},
						"err_l7_sess_process_tcp_data_failed": {
							Type: schema.TypeInt, Optional: true, Description: "L7 session process TCP data failed",
						},
						"err_http2_ver_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Error HTTP2 version mismatch",
						},
						"smart_nat_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Total smart-nat allocation successful",
						},
						"smart_nat_release": {
							Type: schema.TypeInt, Optional: true, Description: "Total smart-nat release successful",
						},
						"smart_nat_alloc_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total smart-nat allocation failed",
						},
						"smart_nat_release_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Total smart-nat release failed",
						},
						"total_server_quic_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total start server QUIC connections",
						},
						"total_server_quic_conn_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total start server QUIC connections error",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHmDplaneStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHmDplaneStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHmDplaneStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHmDplaneStatsStats := setObjectSlbHmDplaneStatsStats(res)
		d.Set("stats", SlbHmDplaneStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHmDplaneStatsStats(ret edpt.DataSlbHmDplaneStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_entries":                         ret.DtSlbHmDplaneStats.Stats.Curr_entries,
			"total_created":                        ret.DtSlbHmDplaneStats.Stats.Total_created,
			"total_inserted":                       ret.DtSlbHmDplaneStats.Stats.Total_inserted,
			"total_ready_to_free":                  ret.DtSlbHmDplaneStats.Stats.Total_ready_to_free,
			"total_freed":                          ret.DtSlbHmDplaneStats.Stats.Total_freed,
			"err_entry_create_failed":              ret.DtSlbHmDplaneStats.Stats.Err_entry_create_failed,
			"err_entry_create_oom":                 ret.DtSlbHmDplaneStats.Stats.Err_entry_create_oom,
			"err_entry_insert_failed":              ret.DtSlbHmDplaneStats.Stats.Err_entry_insert_failed,
			"total_tcp_err":                        ret.DtSlbHmDplaneStats.Stats.Total_tcp_err,
			"err_smart_nat_alloc":                  ret.DtSlbHmDplaneStats.Stats.Err_smart_nat_alloc,
			"err_smart_nat_port_alloc":             ret.DtSlbHmDplaneStats.Stats.Err_smart_nat_port_alloc,
			"err_l4_sess_alloc":                    ret.DtSlbHmDplaneStats.Stats.Err_l4_sess_alloc,
			"err_hm_tcp_conn_sent":                 ret.DtSlbHmDplaneStats.Stats.Err_hm_tcp_conn_sent,
			"hm_tcp_conn_sent":                     ret.DtSlbHmDplaneStats.Stats.Hm_tcp_conn_sent,
			"entry_deleted":                        ret.DtSlbHmDplaneStats.Stats.Entry_deleted,
			"err_entry_create_vip_failed":          ret.DtSlbHmDplaneStats.Stats.Err_entry_create_vip_failed,
			"total_match_resp_code":                ret.DtSlbHmDplaneStats.Stats.Total_match_resp_code,
			"total_match_default_resp_code":        ret.DtSlbHmDplaneStats.Stats.Total_match_default_resp_code,
			"total_maintenance_received":           ret.DtSlbHmDplaneStats.Stats.Total_maintenance_received,
			"total_wrong_status_received":          ret.DtSlbHmDplaneStats.Stats.Total_wrong_status_received,
			"err_no_hm_entry":                      ret.DtSlbHmDplaneStats.Stats.Err_no_hm_entry,
			"err_ssl_cert_name_mismatch":           ret.DtSlbHmDplaneStats.Stats.Err_ssl_cert_name_mismatch,
			"err_server_syn_timeout":               ret.DtSlbHmDplaneStats.Stats.Err_server_syn_timeout,
			"err_http2_callback":                   ret.DtSlbHmDplaneStats.Stats.Err_http2_callback,
			"err_l7_sess_process_tcp_estab_failed": ret.DtSlbHmDplaneStats.Stats.Err_l7_sess_process_tcp_estab_failed,
			"err_l7_sess_process_tcp_data_failed":  ret.DtSlbHmDplaneStats.Stats.Err_l7_sess_process_tcp_data_failed,
			"err_http2_ver_mismatch":               ret.DtSlbHmDplaneStats.Stats.Err_http2_ver_mismatch,
			"smart_nat_alloc":                      ret.DtSlbHmDplaneStats.Stats.Smart_nat_alloc,
			"smart_nat_release":                    ret.DtSlbHmDplaneStats.Stats.Smart_nat_release,
			"smart_nat_alloc_failed":               ret.DtSlbHmDplaneStats.Stats.Smart_nat_alloc_failed,
			"smart_nat_release_failed":             ret.DtSlbHmDplaneStats.Stats.Smart_nat_release_failed,
			"total_server_quic_conn":               ret.DtSlbHmDplaneStats.Stats.Total_server_quic_conn,
			"total_server_quic_conn_err":           ret.DtSlbHmDplaneStats.Stats.Total_server_quic_conn_err,
		},
	}
}

func getObjectSlbHmDplaneStatsStats(d []interface{}) edpt.SlbHmDplaneStatsStats {

	count1 := len(d)
	var ret edpt.SlbHmDplaneStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_entries = in["curr_entries"].(int)
		ret.Total_created = in["total_created"].(int)
		ret.Total_inserted = in["total_inserted"].(int)
		ret.Total_ready_to_free = in["total_ready_to_free"].(int)
		ret.Total_freed = in["total_freed"].(int)
		ret.Err_entry_create_failed = in["err_entry_create_failed"].(int)
		ret.Err_entry_create_oom = in["err_entry_create_oom"].(int)
		ret.Err_entry_insert_failed = in["err_entry_insert_failed"].(int)
		ret.Total_tcp_err = in["total_tcp_err"].(int)
		ret.Err_smart_nat_alloc = in["err_smart_nat_alloc"].(int)
		ret.Err_smart_nat_port_alloc = in["err_smart_nat_port_alloc"].(int)
		ret.Err_l4_sess_alloc = in["err_l4_sess_alloc"].(int)
		ret.Err_hm_tcp_conn_sent = in["err_hm_tcp_conn_sent"].(int)
		ret.Hm_tcp_conn_sent = in["hm_tcp_conn_sent"].(int)
		ret.Entry_deleted = in["entry_deleted"].(int)
		ret.Err_entry_create_vip_failed = in["err_entry_create_vip_failed"].(int)
		ret.Total_match_resp_code = in["total_match_resp_code"].(int)
		ret.Total_match_default_resp_code = in["total_match_default_resp_code"].(int)
		ret.Total_maintenance_received = in["total_maintenance_received"].(int)
		ret.Total_wrong_status_received = in["total_wrong_status_received"].(int)
		ret.Err_no_hm_entry = in["err_no_hm_entry"].(int)
		ret.Err_ssl_cert_name_mismatch = in["err_ssl_cert_name_mismatch"].(int)
		ret.Err_server_syn_timeout = in["err_server_syn_timeout"].(int)
		ret.Err_http2_callback = in["err_http2_callback"].(int)
		ret.Err_l7_sess_process_tcp_estab_failed = in["err_l7_sess_process_tcp_estab_failed"].(int)
		ret.Err_l7_sess_process_tcp_data_failed = in["err_l7_sess_process_tcp_data_failed"].(int)
		ret.Err_http2_ver_mismatch = in["err_http2_ver_mismatch"].(int)
		ret.Smart_nat_alloc = in["smart_nat_alloc"].(int)
		ret.Smart_nat_release = in["smart_nat_release"].(int)
		ret.Smart_nat_alloc_failed = in["smart_nat_alloc_failed"].(int)
		ret.Smart_nat_release_failed = in["smart_nat_release_failed"].(int)
		ret.Total_server_quic_conn = in["total_server_quic_conn"].(int)
		ret.Total_server_quic_conn_err = in["total_server_quic_conn_err"].(int)
	}
	return ret
}

func dataToEndpointSlbHmDplaneStats(d *schema.ResourceData) edpt.SlbHmDplaneStats {
	var ret edpt.SlbHmDplaneStats

	ret.Stats = getObjectSlbHmDplaneStatsStats(d.Get("stats").([]interface{}))
	return ret
}
