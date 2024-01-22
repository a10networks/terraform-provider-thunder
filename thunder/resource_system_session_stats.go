package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSessionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_session_stats`: Statistics for the object session\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemSessionStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_l4_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total L4 Count",
						},
						"conn_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Count",
						},
						"conn_freed_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Freed",
						},
						"total_l4_packet_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total L4 Packet Count",
						},
						"total_l7_packet_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total L7 Packet Count",
						},
						"total_l4_conn_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total L4 Conn Proxy Count",
						},
						"total_l7_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total L7 Conn",
						},
						"total_tcp_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP Conn",
						},
						"curr_free_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Curr Free Conn",
						},
						"tcp_est_counter": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Established",
						},
						"tcp_half_open_counter": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Half Open",
						},
						"tcp_half_close_counter": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Half Closed",
						},
						"udp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Count",
						},
						"ip_counter": {
							Type: schema.TypeInt, Optional: true, Description: "IP Count",
						},
						"other_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Non TCP/UDP IP sessions",
						},
						"reverse_nat_tcp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse NAT TCP",
						},
						"reverse_nat_udp_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse NAT UDP",
						},
						"tcp_syn_half_open_counter": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Half Open",
						},
						"conn_smp_alloc_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Alloc",
						},
						"conn_smp_free_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Free",
						},
						"conn_smp_aged_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Aged",
						},
						"ssl_count_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Curr SSL Count",
						},
						"ssl_count_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL Count",
						},
						"server_ssl_count_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current SSL Server Count",
						},
						"server_ssl_count_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL Server Count",
						},
						"client_ssl_reuse_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL Client Reuse",
						},
						"server_ssl_reuse_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL Server Reuse",
						},
						"total_ip_nat_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total IP Nat Conn",
						},
						"total_l2l3_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Totl L2/L3 Connections",
						},
						"conn_type_0_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Type 0 Available",
						},
						"conn_type_1_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Type 1 Available",
						},
						"conn_type_2_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Type 2 Available",
						},
						"conn_type_3_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Type 3 Available",
						},
						"conn_type_4_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Type 4 Available",
						},
						"conn_smp_type_0_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Type 0 Available",
						},
						"conn_smp_type_1_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Type 1 Available",
						},
						"conn_smp_type_2_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Type 2 Available",
						},
						"conn_smp_type_3_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Type 3 Available",
						},
						"conn_smp_type_4_available": {
							Type: schema.TypeInt, Optional: true, Description: "Conn SMP Type 4 Available",
						},
						"sctp_half_open_counter": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Half Open",
						},
						"sctp_est_counter": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Established",
						},
						"conn_app_smp_alloc_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Conn APP SMP Alloc",
						},
						"diameter_conn_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter Conn Count",
						},
						"diameter_conn_freed_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter Conn Freed",
						},
						"total_fw_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total Firewall Conn",
						},
						"total_local_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total Local Conn",
						},
						"total_curr_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total Curr Conn",
						},
						"client_ssl_fatal_alert": {
							Type: schema.TypeInt, Optional: true, Description: "client ssl fatal alert",
						},
						"client_ssl_fin_rst": {
							Type: schema.TypeInt, Optional: true, Description: "client ssl fin rst",
						},
						"fp_session_fin_rst": {
							Type: schema.TypeInt, Optional: true, Description: "FP Session FIN/RST",
						},
						"server_ssl_fatal_alert": {
							Type: schema.TypeInt, Optional: true, Description: "server ssl fatal alert",
						},
						"server_ssl_fin_rst": {
							Type: schema.TypeInt, Optional: true, Description: "server ssl fin rst",
						},
						"client_template_int_err": {
							Type: schema.TypeInt, Optional: true, Description: "client template internal error",
						},
						"client_template_unknown_err": {
							Type: schema.TypeInt, Optional: true, Description: "client template unknown error",
						},
						"server_template_int_err": {
							Type: schema.TypeInt, Optional: true, Description: "server template int error",
						},
						"server_template_unknown_err": {
							Type: schema.TypeInt, Optional: true, Description: "server template unknown error",
						},
						"diameter_concurrent_user_sessions_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Diameter Concurrent User-Sessions",
						},
						"client_ssl_session_ticket_reuse_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL Client Session Ticket Reuse",
						},
						"server_ssl_session_ticket_reuse_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SSL Server Session Ticket Reuse",
						},
						"total_clientside_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total clientside early data connections",
						},
						"total_serverside_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total serverside early data connections",
						},
						"total_clientside_failed_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total clientside failed early data connections",
						},
						"total_serverside_failed_early_data_connections": {
							Type: schema.TypeInt, Optional: true, Description: "Total serverside failed early data connections",
						},
						"total_logging_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Total Logging Conn",
						},
						"gtp_c_est_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Established",
						},
						"gtp_c_half_open_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Half Open",
						},
						"gtp_u_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Count",
						},
						"gtp_c_echo_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-C Echo Count",
						},
						"gtp_u_echo_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP-U Echo Count",
						},
						"gtp_curr_free_conn": {
							Type: schema.TypeInt, Optional: true, Description: "GTP Current Available Conn",
						},
						"gtp_cum_conn_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP cumulative Conn Count",
						},
						"gtp_cum_conn_freed_counter": {
							Type: schema.TypeInt, Optional: true, Description: "GTP cumulative Conn Freed",
						},
						"fw_blacklist_sess": {
							Type: schema.TypeInt, Optional: true, Description: "Blacklist Sessions",
						},
						"fw_blacklist_sess_created": {
							Type: schema.TypeInt, Optional: true, Description: "Blacklist Session Created",
						},
						"fw_blacklist_sess_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Blacklist Session Freed",
						},
						"server_tcp_est_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Server TCP Established",
						},
						"server_tcp_half_open_counter": {
							Type: schema.TypeInt, Optional: true, Description: "Server TCP Half Open",
						},
					},
				},
			},
		},
	}
}

func resourceSystemSessionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSessionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSessionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemSessionStatsStats := setObjectSystemSessionStatsStats(res)
		d.Set("stats", SystemSessionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemSessionStatsStats(ret edpt.DataSystemSessionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_l4_conn":                                  ret.DtSystemSessionStats.Stats.Total_l4_conn,
			"conn_counter":                                   ret.DtSystemSessionStats.Stats.Conn_counter,
			"conn_freed_counter":                             ret.DtSystemSessionStats.Stats.Conn_freed_counter,
			"total_l4_packet_count":                          ret.DtSystemSessionStats.Stats.Total_l4_packet_count,
			"total_l7_packet_count":                          ret.DtSystemSessionStats.Stats.Total_l7_packet_count,
			"total_l4_conn_proxy":                            ret.DtSystemSessionStats.Stats.Total_l4_conn_proxy,
			"total_l7_conn":                                  ret.DtSystemSessionStats.Stats.Total_l7_conn,
			"total_tcp_conn":                                 ret.DtSystemSessionStats.Stats.Total_tcp_conn,
			"curr_free_conn":                                 ret.DtSystemSessionStats.Stats.Curr_free_conn,
			"tcp_est_counter":                                ret.DtSystemSessionStats.Stats.Tcp_est_counter,
			"tcp_half_open_counter":                          ret.DtSystemSessionStats.Stats.Tcp_half_open_counter,
			"tcp_half_close_counter":                         ret.DtSystemSessionStats.Stats.Tcp_half_close_counter,
			"udp_counter":                                    ret.DtSystemSessionStats.Stats.Udp_counter,
			"ip_counter":                                     ret.DtSystemSessionStats.Stats.Ip_counter,
			"other_counter":                                  ret.DtSystemSessionStats.Stats.Other_counter,
			"reverse_nat_tcp_counter":                        ret.DtSystemSessionStats.Stats.Reverse_nat_tcp_counter,
			"reverse_nat_udp_counter":                        ret.DtSystemSessionStats.Stats.Reverse_nat_udp_counter,
			"tcp_syn_half_open_counter":                      ret.DtSystemSessionStats.Stats.Tcp_syn_half_open_counter,
			"conn_smp_alloc_counter":                         ret.DtSystemSessionStats.Stats.Conn_smp_alloc_counter,
			"conn_smp_free_counter":                          ret.DtSystemSessionStats.Stats.Conn_smp_free_counter,
			"conn_smp_aged_counter":                          ret.DtSystemSessionStats.Stats.Conn_smp_aged_counter,
			"ssl_count_curr":                                 ret.DtSystemSessionStats.Stats.Ssl_count_curr,
			"ssl_count_total":                                ret.DtSystemSessionStats.Stats.Ssl_count_total,
			"server_ssl_count_curr":                          ret.DtSystemSessionStats.Stats.Server_ssl_count_curr,
			"server_ssl_count_total":                         ret.DtSystemSessionStats.Stats.Server_ssl_count_total,
			"client_ssl_reuse_total":                         ret.DtSystemSessionStats.Stats.Client_ssl_reuse_total,
			"server_ssl_reuse_total":                         ret.DtSystemSessionStats.Stats.Server_ssl_reuse_total,
			"total_ip_nat_conn":                              ret.DtSystemSessionStats.Stats.Total_ip_nat_conn,
			"total_l2l3_conn":                                ret.DtSystemSessionStats.Stats.Total_l2l3_conn,
			"conn_type_0_available":                          ret.DtSystemSessionStats.Stats.Conn_type_0_available,
			"conn_type_1_available":                          ret.DtSystemSessionStats.Stats.Conn_type_1_available,
			"conn_type_2_available":                          ret.DtSystemSessionStats.Stats.Conn_type_2_available,
			"conn_type_3_available":                          ret.DtSystemSessionStats.Stats.Conn_type_3_available,
			"conn_type_4_available":                          ret.DtSystemSessionStats.Stats.Conn_type_4_available,
			"conn_smp_type_0_available":                      ret.DtSystemSessionStats.Stats.Conn_smp_type_0_available,
			"conn_smp_type_1_available":                      ret.DtSystemSessionStats.Stats.Conn_smp_type_1_available,
			"conn_smp_type_2_available":                      ret.DtSystemSessionStats.Stats.Conn_smp_type_2_available,
			"conn_smp_type_3_available":                      ret.DtSystemSessionStats.Stats.Conn_smp_type_3_available,
			"conn_smp_type_4_available":                      ret.DtSystemSessionStats.Stats.Conn_smp_type_4_available,
			"sctp_half_open_counter":                         ret.DtSystemSessionStats.Stats.SctpHalfOpenCounter,
			"sctp_est_counter":                               ret.DtSystemSessionStats.Stats.SctpEstCounter,
			"conn_app_smp_alloc_counter":                     ret.DtSystemSessionStats.Stats.Conn_app_smp_alloc_counter,
			"diameter_conn_counter":                          ret.DtSystemSessionStats.Stats.Diameter_conn_counter,
			"diameter_conn_freed_counter":                    ret.DtSystemSessionStats.Stats.Diameter_conn_freed_counter,
			"total_fw_conn":                                  ret.DtSystemSessionStats.Stats.Total_fw_conn,
			"total_local_conn":                               ret.DtSystemSessionStats.Stats.Total_local_conn,
			"total_curr_conn":                                ret.DtSystemSessionStats.Stats.Total_curr_conn,
			"client_ssl_fatal_alert":                         ret.DtSystemSessionStats.Stats.Client_ssl_fatal_alert,
			"client_ssl_fin_rst":                             ret.DtSystemSessionStats.Stats.Client_ssl_fin_rst,
			"fp_session_fin_rst":                             ret.DtSystemSessionStats.Stats.Fp_session_fin_rst,
			"server_ssl_fatal_alert":                         ret.DtSystemSessionStats.Stats.Server_ssl_fatal_alert,
			"server_ssl_fin_rst":                             ret.DtSystemSessionStats.Stats.Server_ssl_fin_rst,
			"client_template_int_err":                        ret.DtSystemSessionStats.Stats.Client_template_int_err,
			"client_template_unknown_err":                    ret.DtSystemSessionStats.Stats.Client_template_unknown_err,
			"server_template_int_err":                        ret.DtSystemSessionStats.Stats.Server_template_int_err,
			"server_template_unknown_err":                    ret.DtSystemSessionStats.Stats.Server_template_unknown_err,
			"diameter_concurrent_user_sessions_counter":      ret.DtSystemSessionStats.Stats.Diameter_concurrent_user_sessions_counter,
			"client_ssl_session_ticket_reuse_total":          ret.DtSystemSessionStats.Stats.Client_ssl_session_ticket_reuse_total,
			"server_ssl_session_ticket_reuse_total":          ret.DtSystemSessionStats.Stats.Server_ssl_session_ticket_reuse_total,
			"total_clientside_early_data_connections":        ret.DtSystemSessionStats.Stats.Total_clientside_early_data_connections,
			"total_serverside_early_data_connections":        ret.DtSystemSessionStats.Stats.Total_serverside_early_data_connections,
			"total_clientside_failed_early_data_connections": ret.DtSystemSessionStats.Stats.Total_clientside_failed_early_dataConnections,
			"total_serverside_failed_early_data_connections": ret.DtSystemSessionStats.Stats.Total_serverside_failed_early_dataConnections,
			"total_logging_conn":                             ret.DtSystemSessionStats.Stats.Total_logging_conn,
			"gtp_c_est_counter":                              ret.DtSystemSessionStats.Stats.Gtp_c_est_counter,
			"gtp_c_half_open_counter":                        ret.DtSystemSessionStats.Stats.Gtp_c_half_open_counter,
			"gtp_u_counter":                                  ret.DtSystemSessionStats.Stats.Gtp_u_counter,
			"gtp_c_echo_counter":                             ret.DtSystemSessionStats.Stats.Gtp_c_echo_counter,
			"gtp_u_echo_counter":                             ret.DtSystemSessionStats.Stats.Gtp_u_echo_counter,
			"gtp_curr_free_conn":                             ret.DtSystemSessionStats.Stats.Gtp_curr_free_conn,
			"gtp_cum_conn_counter":                           ret.DtSystemSessionStats.Stats.Gtp_cum_conn_counter,
			"gtp_cum_conn_freed_counter":                     ret.DtSystemSessionStats.Stats.Gtp_cum_conn_freed_counter,
			"fw_blacklist_sess":                              ret.DtSystemSessionStats.Stats.Fw_blacklist_sess,
			"fw_blacklist_sess_created":                      ret.DtSystemSessionStats.Stats.Fw_blacklist_sess_created,
			"fw_blacklist_sess_freed":                        ret.DtSystemSessionStats.Stats.Fw_blacklist_sess_freed,
			"server_tcp_est_counter":                         ret.DtSystemSessionStats.Stats.Server_tcp_est_counter,
			"server_tcp_half_open_counter":                   ret.DtSystemSessionStats.Stats.Server_tcp_half_open_counter,
		},
	}
}

func getObjectSystemSessionStatsStats(d []interface{}) edpt.SystemSessionStatsStats {

	count1 := len(d)
	var ret edpt.SystemSessionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_l4_conn = in["total_l4_conn"].(int)
		ret.Conn_counter = in["conn_counter"].(int)
		ret.Conn_freed_counter = in["conn_freed_counter"].(int)
		ret.Total_l4_packet_count = in["total_l4_packet_count"].(int)
		ret.Total_l7_packet_count = in["total_l7_packet_count"].(int)
		ret.Total_l4_conn_proxy = in["total_l4_conn_proxy"].(int)
		ret.Total_l7_conn = in["total_l7_conn"].(int)
		ret.Total_tcp_conn = in["total_tcp_conn"].(int)
		ret.Curr_free_conn = in["curr_free_conn"].(int)
		ret.Tcp_est_counter = in["tcp_est_counter"].(int)
		ret.Tcp_half_open_counter = in["tcp_half_open_counter"].(int)
		ret.Tcp_half_close_counter = in["tcp_half_close_counter"].(int)
		ret.Udp_counter = in["udp_counter"].(int)
		ret.Ip_counter = in["ip_counter"].(int)
		ret.Other_counter = in["other_counter"].(int)
		ret.Reverse_nat_tcp_counter = in["reverse_nat_tcp_counter"].(int)
		ret.Reverse_nat_udp_counter = in["reverse_nat_udp_counter"].(int)
		ret.Tcp_syn_half_open_counter = in["tcp_syn_half_open_counter"].(int)
		ret.Conn_smp_alloc_counter = in["conn_smp_alloc_counter"].(int)
		ret.Conn_smp_free_counter = in["conn_smp_free_counter"].(int)
		ret.Conn_smp_aged_counter = in["conn_smp_aged_counter"].(int)
		ret.Ssl_count_curr = in["ssl_count_curr"].(int)
		ret.Ssl_count_total = in["ssl_count_total"].(int)
		ret.Server_ssl_count_curr = in["server_ssl_count_curr"].(int)
		ret.Server_ssl_count_total = in["server_ssl_count_total"].(int)
		ret.Client_ssl_reuse_total = in["client_ssl_reuse_total"].(int)
		ret.Server_ssl_reuse_total = in["server_ssl_reuse_total"].(int)
		ret.Total_ip_nat_conn = in["total_ip_nat_conn"].(int)
		ret.Total_l2l3_conn = in["total_l2l3_conn"].(int)
		ret.Conn_type_0_available = in["conn_type_0_available"].(int)
		ret.Conn_type_1_available = in["conn_type_1_available"].(int)
		ret.Conn_type_2_available = in["conn_type_2_available"].(int)
		ret.Conn_type_3_available = in["conn_type_3_available"].(int)
		ret.Conn_type_4_available = in["conn_type_4_available"].(int)
		ret.Conn_smp_type_0_available = in["conn_smp_type_0_available"].(int)
		ret.Conn_smp_type_1_available = in["conn_smp_type_1_available"].(int)
		ret.Conn_smp_type_2_available = in["conn_smp_type_2_available"].(int)
		ret.Conn_smp_type_3_available = in["conn_smp_type_3_available"].(int)
		ret.Conn_smp_type_4_available = in["conn_smp_type_4_available"].(int)
		ret.SctpHalfOpenCounter = in["sctp_half_open_counter"].(int)
		ret.SctpEstCounter = in["sctp_est_counter"].(int)
		ret.Conn_app_smp_alloc_counter = in["conn_app_smp_alloc_counter"].(int)
		ret.Diameter_conn_counter = in["diameter_conn_counter"].(int)
		ret.Diameter_conn_freed_counter = in["diameter_conn_freed_counter"].(int)
		ret.Total_fw_conn = in["total_fw_conn"].(int)
		ret.Total_local_conn = in["total_local_conn"].(int)
		ret.Total_curr_conn = in["total_curr_conn"].(int)
		ret.Client_ssl_fatal_alert = in["client_ssl_fatal_alert"].(int)
		ret.Client_ssl_fin_rst = in["client_ssl_fin_rst"].(int)
		ret.Fp_session_fin_rst = in["fp_session_fin_rst"].(int)
		ret.Server_ssl_fatal_alert = in["server_ssl_fatal_alert"].(int)
		ret.Server_ssl_fin_rst = in["server_ssl_fin_rst"].(int)
		ret.Client_template_int_err = in["client_template_int_err"].(int)
		ret.Client_template_unknown_err = in["client_template_unknown_err"].(int)
		ret.Server_template_int_err = in["server_template_int_err"].(int)
		ret.Server_template_unknown_err = in["server_template_unknown_err"].(int)
		ret.Diameter_concurrent_user_sessions_counter = in["diameter_concurrent_user_sessions_counter"].(int)
		ret.Client_ssl_session_ticket_reuse_total = in["client_ssl_session_ticket_reuse_total"].(int)
		ret.Server_ssl_session_ticket_reuse_total = in["server_ssl_session_ticket_reuse_total"].(int)
		ret.Total_clientside_early_data_connections = in["total_clientside_early_data_connections"].(int)
		ret.Total_serverside_early_data_connections = in["total_serverside_early_data_connections"].(int)
		ret.Total_clientside_failed_early_dataConnections = in["total_clientside_failed_early_data_connections"].(int)
		ret.Total_serverside_failed_early_dataConnections = in["total_serverside_failed_early_data_connections"].(int)
		ret.Total_logging_conn = in["total_logging_conn"].(int)
		ret.Gtp_c_est_counter = in["gtp_c_est_counter"].(int)
		ret.Gtp_c_half_open_counter = in["gtp_c_half_open_counter"].(int)
		ret.Gtp_u_counter = in["gtp_u_counter"].(int)
		ret.Gtp_c_echo_counter = in["gtp_c_echo_counter"].(int)
		ret.Gtp_u_echo_counter = in["gtp_u_echo_counter"].(int)
		ret.Gtp_curr_free_conn = in["gtp_curr_free_conn"].(int)
		ret.Gtp_cum_conn_counter = in["gtp_cum_conn_counter"].(int)
		ret.Gtp_cum_conn_freed_counter = in["gtp_cum_conn_freed_counter"].(int)
		ret.Fw_blacklist_sess = in["fw_blacklist_sess"].(int)
		ret.Fw_blacklist_sess_created = in["fw_blacklist_sess_created"].(int)
		ret.Fw_blacklist_sess_freed = in["fw_blacklist_sess_freed"].(int)
		ret.Server_tcp_est_counter = in["server_tcp_est_counter"].(int)
		ret.Server_tcp_half_open_counter = in["server_tcp_half_open_counter"].(int)
	}
	return ret
}

func dataToEndpointSystemSessionStats(d *schema.ResourceData) edpt.SystemSessionStats {
	var ret edpt.SystemSessionStats

	ret.Stats = getObjectSystemSessionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
