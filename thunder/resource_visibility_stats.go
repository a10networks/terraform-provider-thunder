package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_stats`: Statistics for the object visibility\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityStatsRead,

		Schema: map[string]*schema.Schema{
			"mon_entity_telemetry_data": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"in_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN pkts",
									},
									"out_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT pkts",
									},
									"in_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN bytes",
									},
									"out_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT bytes",
									},
									"errors": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric ERRORS",
									},
									"in_small_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN SMALL pkt",
									},
									"in_frag": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN frag",
									},
									"out_small_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT SMALL pkt",
									},
									"out_frag": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT frag",
									},
									"new_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric New Sessions",
									},
									"avg_data_cpu_util": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric Average data CPU utilization",
									},
									"outside_intf_util": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric Outside interface utilization",
									},
									"concurrent_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"in_bytes_per_out_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN bytes per OUT bytes",
									},
									"drop_pkts_per_pkts": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric Drop pkts per pkts",
									},
									"tcp_in_syn": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN syn",
									},
									"tcp_out_syn": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT syn",
									},
									"tcp_in_fin": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN fin",
									},
									"tcp_out_fin": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT fin",
									},
									"tcp_in_payload": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN payload",
									},
									"tcp_out_payload": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT payload",
									},
									"tcp_in_rexmit": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN rexmit",
									},
									"tcp_out_rexmit": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT rexmit",
									},
									"tcp_in_rst": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN rst",
									},
									"tcp_out_rst": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT rst",
									},
									"tcp_in_empty_ack": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP_IN EMPTY ack",
									},
									"tcp_out_empty_ack": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT EMPTY ack",
									},
									"tcp_in_zero_wnd": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN ZERO wnd",
									},
									"tcp_out_zero_wnd": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT ZERO wnd",
									},
									"tcp_conn_miss": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP connection miss",
									},
									"tcp_fwd_syn_per_fin": {
										Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP FWD SYN per FIN",
									},
								},
							},
						},
					},
				},
			},
			"packet_capture": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"concurrent_capture_created_by_ctr_increment": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic 3 tuple based capture created (ctr increment based)",
									},
									"concurrent_capture_created_by_ctr_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic 3 tuple based capture created (ctr anomaly based)",
									},
									"concurrent_capture_created_by_other_feature": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic 3 tuple based capture created (Other feature based)",
									},
									"concurrent_capture_create_failed_by_ctr_increment": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture(ctr increment based) create failed",
									},
									"concurrent_capture_create_failed_by_ctr_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture(ctr anomaly based) create failed",
									},
									"concurrent_capture_create_failed_by_other_feature": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture(Other feature based) create failed",
									},
									"concurrent_capture_create_failed_oom": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture create failed, OOM",
									},
									"concurrent_capture_limit_reached": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture configured concurrent limit reached",
									},
									"concurrent_capture_by_ctr_increment_freed": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture(ctr increment based) freed",
									},
									"concurrent_capture_by_ctr_anomaly_freed": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture(ctr anomaly based) freed",
									},
									"concurrent_capture_by_ctr_other_feature_freed": {
										Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture(Other feature based) freed",
									},
									"global_capture_finished": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times global capture finished capturing",
									},
									"concurrent_capture_finished": {
										Type: schema.TypeInt, Optional: true, Description: "Number of Dynamic captures(3 tuple based) finished capturing",
									},
									"pktcapture_with_no_conn_success": {
										Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets without conn",
									},
									"pktcapture_with_no_conn_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets without conn",
									},
									"pktcapture_with_conn_but_not_tagged_success": {
										Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets with untagged conn",
									},
									"pktcapture_with_conn_but_not_tagged_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets with untagged conn",
									},
									"pktcapture_with_conn_success_global": {
										Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets with tagged conn (global capture)",
									},
									"pktcapture_with_conn_success": {
										Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets with tagged conn (dynamic capture)",
									},
									"pktcapture_with_conn_failure_global": {
										Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets with tagged conn (global capture)",
									},
									"pktcapture_with_conn_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets with tagged conn (dynamic capture)",
									},
									"pktcapture_failure_wait_for_block": {
										Type: schema.TypeInt, Optional: true, Description: "Capture fail, waiting to get free buffer",
									},
									"pktcapture_failure_file_size_rchd": {
										Type: schema.TypeInt, Optional: true, Description: "Capture fail, file size reached",
									},
									"num_conns_tagged_global_increment": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr increment, Global)",
									},
									"num_conns_tagged_global_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr anomaly, Global)",
									},
									"num_conns_tagged_global_other_feature": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on Other feature, Global)",
									},
									"num_conns_tagged_global_increment_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr increment, Global)",
									},
									"num_conns_tagged_global_anomaly_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr anomaly, Global)",
									},
									"num_conns_tagged_global_other_feature_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on Other feature, Global)",
									},
									"num_conns_tagged_global_increment_maxed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr increment, Global)",
									},
									"num_conns_tagged_global_anomaly_maxed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr anomaly, Global)",
									},
									"num_conns_tagged_global_other_feature_maxed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on Other feature, Global)",
									},
									"num_conns_tagged_increment": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr increment, dynamic)",
									},
									"num_conns_tagged_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr anomaly, dynamic)",
									},
									"num_conns_tagged_other_feature": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on Other feature, dynamic)",
									},
									"num_conns_tagged_increment_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr increment, dynamic)",
									},
									"num_conns_tagged_anomaly_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr anomaly, dynamic)",
									},
									"num_conns_tagged_other_feature_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on Other feature, dynamic)",
									},
									"num_conns_tagged_increment_maxed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr increment, dynamic)",
									},
									"num_conns_tagged_anomaly_maxed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr anomaly, dynamic)",
									},
									"num_conns_tagged_other_feature_maxed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on Other feature, dynamic)",
									},
									"num_conns_untagged": {
										Type: schema.TypeInt, Optional: true, Description: "Number of conns untagged (done with conn limit or capture)",
									},
									"pktcapture_triggered_by_increment": {
										Type: schema.TypeInt, Optional: true, Description: "Capture triggered by counter increment",
									},
									"pktcapture_triggered_by_anomaly": {
										Type: schema.TypeInt, Optional: true, Description: "Capture triggered by counter anomaly",
									},
									"pktcapture_triggered_by_other_feature": {
										Type: schema.TypeInt, Optional: true, Description: "Capture triggered by Other feature",
									},
									"num_of_anomalies_detected": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times ctr Anomaly detected",
									},
									"num_of_anomalies_cleared": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times ctr Anomaly cleared",
									},
									"num_pcaps_created": {
										Type: schema.TypeInt, Optional: true, Description: "Number of pcapng files created",
									},
									"num_tmp_pcaps_created": {
										Type: schema.TypeInt, Optional: true, Description: "Number of temporary pcapng files created",
									},
									"num_pcaps_create_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Number of pcapng files creation failed",
									},
									"pktcap_oom": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Automated Packet capture infra OOM",
									},
									"failed_disk_full": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Capture fail, Disk limit reached",
									},
									"conn_ext_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Error, Conn extension creation fail",
									},
									"skip_as_conn_already_recapture": {
										Type: schema.TypeInt, Optional: true, Description: "Skip creating capture, conn was already captured",
									},
									"skip_capture_as_conn_created_before_smp": {
										Type: schema.TypeInt, Optional: true, Description: "Skip capturing, conn was created before the capture started",
									},
									"failed_as_return_completed_set": {
										Type: schema.TypeInt, Optional: true, Description: "Skip capturing, capture-config marked completed",
									},
									"non_pkt_path": {
										Type: schema.TypeInt, Optional: true, Description: "Skip capturing, not packet processing path",
									},
									"pkt_already_captured": {
										Type: schema.TypeInt, Optional: true, Description: "Skip capturing, packet already captured",
									},
									"wrong_ctr_incremented": {
										Type: schema.TypeInt, Optional: true, Description: "Counter increment issue",
									},
									"auto_pcap_file_merged": {
										Type: schema.TypeInt, Optional: true, Description: "Auto pcapng files merged",
									},
									"auto_pcap_file_merged_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Auto pcapng files merged failed",
									},
									"num_global_tmpl_created": {
										Type: schema.TypeInt, Optional: true, Description: "Number of global templates created",
									},
									"num_object_tmpl_created": {
										Type: schema.TypeInt, Optional: true, Description: "Number of object templates created",
									},
									"num_global_tmpl_deleted": {
										Type: schema.TypeInt, Optional: true, Description: "Number of global templates deleted",
									},
									"num_object_tmpl_deleted": {
										Type: schema.TypeInt, Optional: true, Description: "Number of object templates deleted",
									},
									"num_capture_config_created": {
										Type: schema.TypeInt, Optional: true, Description: "Number of capture-config created",
									},
									"num_dynamic_capture_config_created": {
										Type: schema.TypeInt, Optional: true, Description: "Number of dynamic capture-config created",
									},
									"num_capture_config_deleted": {
										Type: schema.TypeInt, Optional: true, Description: "Number of capture-config deleted",
									},
									"num_dynamic_capture_config_deleted": {
										Type: schema.TypeInt, Optional: true, Description: "Number of dynamic capture-config deleted",
									},
									"num_capture_config_delete_q": {
										Type: schema.TypeInt, Optional: true, Description: "Number of capture-config set for deletion",
									},
									"num_dynamic_capture_config_delete_q": {
										Type: schema.TypeInt, Optional: true, Description: "Number of dynamic capture-config set for deletion",
									},
									"num_capture_config_linked": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times capture-config linked to template",
									},
									"num_dynamic_capture_config_linked": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times dynamic capture-config linked to template",
									},
									"num_capture_config_unlinked": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times capture-config unlinked from template",
									},
									"num_dynamic_capture_config_unlinked": {
										Type: schema.TypeInt, Optional: true, Description: "Number of times dynamic capture-config unlinked from template",
									},
									"num_global_counters_registered": {
										Type: schema.TypeInt, Optional: true, Description: "Number of global objects registered",
									},
									"num_global_counters_deregistered": {
										Type: schema.TypeInt, Optional: true, Description: "Number of global objects deregistered",
									},
									"num_per_object_counters_registered": {
										Type: schema.TypeInt, Optional: true, Description: "Number of per instance objects registered",
									},
									"num_per_object_counters_deregistered": {
										Type: schema.TypeInt, Optional: true, Description: "Number of per instance objects deregistered",
									},
								},
							},
						},
						"capture_config_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Specify the name of the capture-config",
									},
									"stats": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"concurrent_capture_created_by_ctr_increment": {
													Type: schema.TypeInt, Optional: true, Description: "Dynamic 3 tuple based capture created (ctr increment based)",
												},
												"concurrent_capture_created_by_ctr_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "Dynamic 3 tuple based capture created (ctr anomaly based)",
												},
												"concurrent_capture_create_failed_by_ctr_increment": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture(ctr increment based) create failed",
												},
												"concurrent_capture_create_failed_by_ctr_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture(ctr anomaly based) create failed",
												},
												"concurrent_capture_create_failed_by_other_feature": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture(Other feature based) create failed",
												},
												"concurrent_capture_create_failed_oom": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Dynamic Capture create failed, OOM",
												},
												"concurrent_capture_limit_reached": {
													Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture configured concurrent limit reached",
												},
												"concurrent_capture_by_ctr_increment_freed": {
													Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture(ctr increment based) freed",
												},
												"concurrent_capture_by_ctr_anomaly_freed": {
													Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture(ctr anomaly based) freed",
												},
												"concurrent_capture_by_ctr_other_feature_freed": {
													Type: schema.TypeInt, Optional: true, Description: "Dynamic Capture(Other feature based) freed",
												},
												"global_capture_finished": {
													Type: schema.TypeInt, Optional: true, Description: "Number of times global capture finished capturing",
												},
												"concurrent_capture_finished": {
													Type: schema.TypeInt, Optional: true, Description: "Number of Dynamic captures(3 tuple based) finished capturing",
												},
												"pktcapture_with_no_conn_success": {
													Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets without conn",
												},
												"pktcapture_with_no_conn_failure": {
													Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets without conn",
												},
												"pktcapture_with_conn_but_not_tagged_success": {
													Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets with untagged conn",
												},
												"pktcapture_with_conn_but_not_tagged_failure": {
													Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets with untagged conn",
												},
												"pktcapture_with_conn_success_global": {
													Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets with tagged conn (global capture)",
												},
												"pktcapture_with_conn_success": {
													Type: schema.TypeInt, Optional: true, Description: "Capture success, Packets with tagged conn (dynamic capture)",
												},
												"pktcapture_with_conn_failure_global": {
													Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets with tagged conn (global capture)",
												},
												"pktcapture_with_conn_failure": {
													Type: schema.TypeInt, Optional: true, Description: "Capture fail, Packets with tagged conn (dynamic capture)",
												},
												"pktcapture_failure_wait_for_block": {
													Type: schema.TypeInt, Optional: true, Description: "Capture fail, waiting to get free buffer",
												},
												"pktcapture_failure_file_size_rchd": {
													Type: schema.TypeInt, Optional: true, Description: "Capture fail, file size reached",
												},
												"num_conns_tagged_global_increment": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr increment, Global)",
												},
												"num_conns_tagged_global_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr anomaly, Global)",
												},
												"num_conns_tagged_global_other_feature": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on Other feature, Global)",
												},
												"num_conns_tagged_global_increment_fail": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr increment, Global)",
												},
												"num_conns_tagged_global_anomaly_fail": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr anomaly, Global)",
												},
												"num_conns_tagged_global_other_feature_fail": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on Other feature, Global)",
												},
												"num_conns_tagged_global_increment_maxed": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr increment, Global)",
												},
												"num_conns_tagged_global_anomaly_maxed": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr anomaly, Global)",
												},
												"num_conns_tagged_global_other_feature_maxed": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on Other feature, Global)",
												},
												"num_conns_tagged_increment": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr increment, dynamic)",
												},
												"num_conns_tagged_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on ctr anomaly, dynamic)",
												},
												"num_conns_tagged_other_feature": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag success (based on Other feature, dynamic)",
												},
												"num_conns_tagged_increment_fail": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr increment, dynamic)",
												},
												"num_conns_tagged_anomaly_fail": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on ctr anomaly, dynamic)",
												},
												"num_conns_tagged_other_feature_fail": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail (based on Other feature, dynamic)",
												},
												"num_conns_tagged_increment_maxed": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr increment, dynamic)",
												},
												"num_conns_tagged_anomaly_maxed": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on ctr anomaly, dynamic)",
												},
												"num_conns_tagged_other_feature_maxed": {
													Type: schema.TypeInt, Optional: true, Description: "Conn tag fail, reached limit (based on Other feature, dynamic)",
												},
												"num_conns_untagged": {
													Type: schema.TypeInt, Optional: true, Description: "Number of conns untagged (done with conn limit or capture)",
												},
												"pktcapture_triggered_by_increment": {
													Type: schema.TypeInt, Optional: true, Description: "Capture triggered by counter increment",
												},
												"pktcapture_triggered_by_anomaly": {
													Type: schema.TypeInt, Optional: true, Description: "Capture triggered by counter anomaly",
												},
												"pktcapture_triggered_by_other_feature": {
													Type: schema.TypeInt, Optional: true, Description: "Capture triggered by Other feature",
												},
												"num_of_anomalies_detected": {
													Type: schema.TypeInt, Optional: true, Description: "Number of times ctr Anomaly detected",
												},
												"num_of_anomalies_cleared": {
													Type: schema.TypeInt, Optional: true, Description: "Number of times ctr Anomaly cleared",
												},
												"num_pcaps_created": {
													Type: schema.TypeInt, Optional: true, Description: "Number of pcapng files created",
												},
												"num_tmp_pcaps_created": {
													Type: schema.TypeInt, Optional: true, Description: "Number of temporary pcapng files created",
												},
												"num_pcaps_create_failed": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Number of pcapng files creation failed",
												},
												"pktcap_oom": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Automated Packet capture infra OOM",
												},
												"failed_disk_full": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Capture fail, Disk limit reached",
												},
												"conn_ext_failed": {
													Type: schema.TypeInt, Optional: true, Description: "Error, Conn extension creation fail",
												},
												"skip_as_conn_already_recapture": {
													Type: schema.TypeInt, Optional: true, Description: "Skip creating capture, conn was already captured",
												},
												"skip_capture_as_conn_created_before_smp": {
													Type: schema.TypeInt, Optional: true, Description: "Skip capturing, conn was created before the capture started",
												},
												"failed_as_return_completed_set": {
													Type: schema.TypeInt, Optional: true, Description: "Skip capturing, capture-config marked completed",
												},
												"non_pkt_path": {
													Type: schema.TypeInt, Optional: true, Description: "Skip capturing, not packet processing path",
												},
												"pkt_already_captured": {
													Type: schema.TypeInt, Optional: true, Description: "Skip capturing, packet already captured",
												},
												"wrong_ctr_incremented": {
													Type: schema.TypeInt, Optional: true, Description: "Counter increment issue",
												},
												"auto_pcap_file_merged": {
													Type: schema.TypeInt, Optional: true, Description: "Auto pcapng files merged",
												},
												"auto_pcap_file_merged_failed": {
													Type: schema.TypeInt, Optional: true, Description: "Auto pcapng files merged failed",
												},
												"num_dynamic_capture_config_created": {
													Type: schema.TypeInt, Optional: true, Description: "Number of dynamic capture-config created",
												},
												"num_dynamic_capture_config_delete_q": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"num_dynamic_capture_config_deleted": {
													Type: schema.TypeInt, Optional: true, Description: "Number of dynamic capture-config deleted",
												},
												"num_global_counters_registered": {
													Type: schema.TypeInt, Optional: true, Description: "Number of global objects registered",
												},
												"num_global_counters_deregistered": {
													Type: schema.TypeInt, Optional: true, Description: "Number of global objects deregistered",
												},
												"num_per_object_counters_registered": {
													Type: schema.TypeInt, Optional: true, Description: "Number of per instance objects registered",
												},
												"num_per_object_counters_deregistered": {
													Type: schema.TypeInt, Optional: true, Description: "Number of per instance objects deregistered",
												},
											},
										},
									},
								},
							},
						},
						"automated_captures": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stats": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"total_failure": {
													Type: schema.TypeInt, Optional: true, Description: "Total failures",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"reporting": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_transmit_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Total log transmit failures",
									},
									"buffer_alloc_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Total reporting buffer allocation failures",
									},
									"notif_jobs_in_queue": {
										Type: schema.TypeInt, Optional: true, Description: "Total notification jobs in queue",
									},
									"enqueue_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Total enqueue jobs failed",
									},
									"enqueue_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Total enqueue jobs passed",
									},
									"dequeued": {
										Type: schema.TypeInt, Optional: true, Description: "Total jobs dequeued",
									},
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mon_entity_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity limit exceed failures",
						},
						"ha_entity_create_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total montior entity HA create messages sent",
						},
						"ha_entity_delete_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total montior entity HA delete messages sent",
						},
						"ha_entity_anomaly_on_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total anomaly on HA messages sent",
						},
						"ha_entity_anomaly_off_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total anomaly off HA messages sent",
						},
						"ha_entity_periodic_sync_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity periodic sync messages sent",
						},
						"out_of_memory_alloc_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Out of memory allocation failures",
						},
						"lw_mon_entity_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total Light-weight entities created",
						},
						"lw_mon_entity_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Total Light-weight entities deleted",
						},
						"lw_mon_entity_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Light weight limit exceeded errors",
						},
						"lw_out_of_memory_alloc_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Light Weight Out-of-memory allocation failures",
						},
						"mon_entity_rrd_file_timestamp_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity rrd file timestamp errors",
						},
						"mon_entity_rrd_update_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity rrd update error",
						},
						"mon_entity_rrd_last_update_fetch_failed_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity rrd last update fetch failed error",
						},
						"mon_entity_rrd_tune_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity rrd tune error",
						},
						"mon_entity_rrd_out_of_memory_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity rrd load failed, out of memory error",
						},
						"mon_entity_rrd_file_create_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total monitor entity rrd file create error",
						},
					},
				},
			},
			"topn": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"heap_alloc_success": {
										Type: schema.TypeInt, Optional: true, Description: "Total heap node allocated",
									},
									"heap_alloc_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Total heap node alloc failed",
									},
									"heap_alloc_oom": {
										Type: schema.TypeInt, Optional: true, Description: "Total heap node alloc failed Out of Memory",
									},
									"obj_reg_success": {
										Type: schema.TypeInt, Optional: true, Description: "Total object node allocated",
									},
									"obj_reg_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Total object node alloc failed",
									},
									"obj_reg_oom": {
										Type: schema.TypeInt, Optional: true, Description: "Total object node alloc failed Out of Memory",
									},
									"heap_deleted": {
										Type: schema.TypeInt, Optional: true, Description: "Total Heap node deleted",
									},
									"obj_deleted": {
										Type: schema.TypeInt, Optional: true, Description: "Total Object node deleted",
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

func resourceVisibilityStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityStatsMonEntityTelemetryData := setObjectVisibilityStatsMonEntityTelemetryData(res)
		d.Set("mon_entity_telemetry_data", VisibilityStatsMonEntityTelemetryData)
		VisibilityStatsPacketCapture := setObjectVisibilityStatsPacketCapture(res)
		d.Set("packet_capture", VisibilityStatsPacketCapture)
		VisibilityStatsReporting := setObjectVisibilityStatsReporting(res)
		d.Set("reporting", VisibilityStatsReporting)
		VisibilityStatsStats := setObjectVisibilityStatsStats(res)
		d.Set("stats", VisibilityStatsStats)
		VisibilityStatsTopn := setObjectVisibilityStatsTopn(res)
		d.Set("topn", VisibilityStatsTopn)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityStatsMonEntityTelemetryData(ret edpt.DataVisibilityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVisibilityStatsMonEntityTelemetryDataStats(ret.DtVisibilityStats.MonEntityTelemetryData.Stats),
		},
	}
}

func setObjectVisibilityStatsMonEntityTelemetryDataStats(d edpt.VisibilityStatsMonEntityTelemetryDataStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["in_pkts"] = d.In_pkts

	in["out_pkts"] = d.Out_pkts

	in["in_bytes"] = d.In_bytes

	in["out_bytes"] = d.Out_bytes

	in["errors"] = d.Errors

	in["in_small_pkt"] = d.In_small_pkt

	in["in_frag"] = d.In_frag

	in["out_small_pkt"] = d.Out_small_pkt

	in["out_frag"] = d.Out_frag

	in["new_conn"] = d.NewConn

	in["avg_data_cpu_util"] = d.Avg_data_cpu_util

	in["outside_intf_util"] = d.Outside_intf_util

	in["concurrent_conn"] = d.ConcurrentConn

	in["in_bytes_per_out_bytes"] = d.In_bytes_per_out_bytes

	in["drop_pkts_per_pkts"] = d.Drop_pkts_per_pkts

	in["tcp_in_syn"] = d.Tcp_in_syn

	in["tcp_out_syn"] = d.Tcp_out_syn

	in["tcp_in_fin"] = d.Tcp_in_fin

	in["tcp_out_fin"] = d.Tcp_out_fin

	in["tcp_in_payload"] = d.Tcp_in_payload

	in["tcp_out_payload"] = d.Tcp_out_payload

	in["tcp_in_rexmit"] = d.Tcp_in_rexmit

	in["tcp_out_rexmit"] = d.Tcp_out_rexmit

	in["tcp_in_rst"] = d.Tcp_in_rst

	in["tcp_out_rst"] = d.Tcp_out_rst

	in["tcp_in_empty_ack"] = d.Tcp_in_empty_ack

	in["tcp_out_empty_ack"] = d.Tcp_out_empty_ack

	in["tcp_in_zero_wnd"] = d.Tcp_in_zero_wnd

	in["tcp_out_zero_wnd"] = d.Tcp_out_zero_wnd

	in["tcp_conn_miss"] = d.Tcp_conn_miss

	in["tcp_fwd_syn_per_fin"] = d.Tcp_fwd_syn_per_fin
	result = append(result, in)
	return result
}

func setObjectVisibilityStatsPacketCapture(ret edpt.DataVisibilityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats":               setObjectVisibilityStatsPacketCaptureStats(ret.DtVisibilityStats.PacketCapture.Stats),
			"capture_config_list": setSliceVisibilityStatsPacketCaptureCaptureConfigList(ret.DtVisibilityStats.PacketCapture.CaptureConfigList),
			"automated_captures":  setObjectVisibilityStatsPacketCaptureAutomatedCaptures(ret.DtVisibilityStats.PacketCapture.AutomatedCaptures),
		},
	}
}

func setObjectVisibilityStatsPacketCaptureStats(d edpt.VisibilityStatsPacketCaptureStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["concurrent_capture_created_by_ctr_increment"] = d.ConcurrentCaptureCreatedByCtrIncrement

	in["concurrent_capture_created_by_ctr_anomaly"] = d.ConcurrentCaptureCreatedByCtrAnomaly

	in["concurrent_capture_created_by_other_feature"] = d.ConcurrentCaptureCreatedByOtherFeature

	in["concurrent_capture_create_failed_by_ctr_increment"] = d.ConcurrentCaptureCreateFailedByCtrIncrement

	in["concurrent_capture_create_failed_by_ctr_anomaly"] = d.ConcurrentCaptureCreateFailedByCtrAnomaly

	in["concurrent_capture_create_failed_by_other_feature"] = d.ConcurrentCaptureCreateFailedByOtherFeature

	in["concurrent_capture_create_failed_oom"] = d.ConcurrentCaptureCreateFailedOom

	in["concurrent_capture_limit_reached"] = d.ConcurrentCaptureLimitReached

	in["concurrent_capture_by_ctr_increment_freed"] = d.ConcurrentCaptureByCtrIncrementFreed

	in["concurrent_capture_by_ctr_anomaly_freed"] = d.ConcurrentCaptureByCtrAnomalyFreed

	in["concurrent_capture_by_ctr_other_feature_freed"] = d.ConcurrentCaptureByCtrOtherFeatureFreed

	in["global_capture_finished"] = d.GlobalCaptureFinished

	in["concurrent_capture_finished"] = d.ConcurrentCaptureFinished

	in["pktcapture_with_no_conn_success"] = d.PktcaptureWithNoConnSuccess

	in["pktcapture_with_no_conn_failure"] = d.PktcaptureWithNoConnFailure

	in["pktcapture_with_conn_but_not_tagged_success"] = d.PktcaptureWithConnButNotTaggedSuccess

	in["pktcapture_with_conn_but_not_tagged_failure"] = d.PktcaptureWithConnButNotTaggedFailure

	in["pktcapture_with_conn_success_global"] = d.PktcaptureWithConnSuccessGlobal

	in["pktcapture_with_conn_success"] = d.PktcaptureWithConnSuccess

	in["pktcapture_with_conn_failure_global"] = d.PktcaptureWithConnFailureGlobal

	in["pktcapture_with_conn_failure"] = d.PktcaptureWithConnFailure

	in["pktcapture_failure_wait_for_block"] = d.PktcaptureFailureWaitForBlock

	in["pktcapture_failure_file_size_rchd"] = d.PktcaptureFailureFileSizeRchd

	in["num_conns_tagged_global_increment"] = d.NumConnsTaggedGlobalIncrement

	in["num_conns_tagged_global_anomaly"] = d.NumConnsTaggedGlobalAnomaly

	in["num_conns_tagged_global_other_feature"] = d.NumConnsTaggedGlobalOtherFeature

	in["num_conns_tagged_global_increment_fail"] = d.NumConnsTaggedGlobalIncrementFail

	in["num_conns_tagged_global_anomaly_fail"] = d.NumConnsTaggedGlobalAnomalyFail

	in["num_conns_tagged_global_other_feature_fail"] = d.NumConnsTaggedGlobalOtherFeatureFail

	in["num_conns_tagged_global_increment_maxed"] = d.NumConnsTaggedGlobalIncrementMaxed

	in["num_conns_tagged_global_anomaly_maxed"] = d.NumConnsTaggedGlobalAnomalyMaxed

	in["num_conns_tagged_global_other_feature_maxed"] = d.NumConnsTaggedGlobalOtherFeatureMaxed

	in["num_conns_tagged_increment"] = d.NumConnsTaggedIncrement

	in["num_conns_tagged_anomaly"] = d.NumConnsTaggedAnomaly

	in["num_conns_tagged_other_feature"] = d.NumConnsTaggedOtherFeature

	in["num_conns_tagged_increment_fail"] = d.NumConnsTaggedIncrementFail

	in["num_conns_tagged_anomaly_fail"] = d.NumConnsTaggedAnomalyFail

	in["num_conns_tagged_other_feature_fail"] = d.NumConnsTaggedOtherFeatureFail

	in["num_conns_tagged_increment_maxed"] = d.NumConnsTaggedIncrementMaxed

	in["num_conns_tagged_anomaly_maxed"] = d.NumConnsTaggedAnomalyMaxed

	in["num_conns_tagged_other_feature_maxed"] = d.NumConnsTaggedOtherFeatureMaxed

	in["num_conns_untagged"] = d.NumConnsUntagged

	in["pktcapture_triggered_by_increment"] = d.PktcaptureTriggeredByIncrement

	in["pktcapture_triggered_by_anomaly"] = d.PktcaptureTriggeredByAnomaly

	in["pktcapture_triggered_by_other_feature"] = d.PktcaptureTriggeredByOtherFeature

	in["num_of_anomalies_detected"] = d.NumOfAnomaliesDetected

	in["num_of_anomalies_cleared"] = d.NumOfAnomaliesCleared

	in["num_pcaps_created"] = d.NumPcapsCreated

	in["num_tmp_pcaps_created"] = d.NumTmpPcapsCreated

	in["num_pcaps_create_failed"] = d.NumPcapsCreateFailed

	in["pktcap_oom"] = d.PktcapOom

	in["failed_disk_full"] = d.FailedDiskFull

	in["conn_ext_failed"] = d.ConnExtFailed

	in["skip_as_conn_already_recapture"] = d.SkipAsConnAlreadyRecapture

	in["skip_capture_as_conn_created_before_smp"] = d.SkipCaptureAsConnCreatedBeforeSmp

	in["failed_as_return_completed_set"] = d.FailedAsReturnCompletedSet

	in["non_pkt_path"] = d.NonPktPath

	in["pkt_already_captured"] = d.PktAlreadyCaptured

	in["wrong_ctr_incremented"] = d.WrongCtrIncremented

	in["auto_pcap_file_merged"] = d.AutoPcapFileMerged

	in["auto_pcap_file_merged_failed"] = d.AutoPcapFileMergedFailed

	in["num_global_tmpl_created"] = d.NumGlobalTmplCreated

	in["num_object_tmpl_created"] = d.NumObjectTmplCreated

	in["num_global_tmpl_deleted"] = d.NumGlobalTmplDeleted

	in["num_object_tmpl_deleted"] = d.NumObjectTmplDeleted

	in["num_capture_config_created"] = d.NumCaptureConfigCreated

	in["num_dynamic_capture_config_created"] = d.NumDynamicCaptureConfigCreated

	in["num_capture_config_deleted"] = d.NumCaptureConfigDeleted

	in["num_dynamic_capture_config_deleted"] = d.NumDynamicCaptureConfigDeleted

	in["num_capture_config_delete_q"] = d.NumCaptureConfigDeleteQ

	in["num_dynamic_capture_config_delete_q"] = d.NumDynamicCaptureConfigDeleteQ

	in["num_capture_config_linked"] = d.NumCaptureConfigLinked

	in["num_dynamic_capture_config_linked"] = d.NumDynamicCaptureConfigLinked

	in["num_capture_config_unlinked"] = d.NumCaptureConfigUnlinked

	in["num_dynamic_capture_config_unlinked"] = d.NumDynamicCaptureConfigUnlinked

	in["num_global_counters_registered"] = d.NumGlobalCountersRegistered

	in["num_global_counters_deregistered"] = d.NumGlobalCountersDeregistered

	in["num_per_object_counters_registered"] = d.NumPerObjectCountersRegistered

	in["num_per_object_counters_deregistered"] = d.NumPerObjectCountersDeregistered
	result = append(result, in)
	return result
}

func setSliceVisibilityStatsPacketCaptureCaptureConfigList(d []edpt.VisibilityStatsPacketCaptureCaptureConfigList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectVisibilityStatsPacketCaptureCaptureConfigListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityStatsPacketCaptureCaptureConfigListStats(d edpt.VisibilityStatsPacketCaptureCaptureConfigListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["concurrent_capture_created_by_ctr_increment"] = d.ConcurrentCaptureCreatedByCtrIncrement

	in["concurrent_capture_created_by_ctr_anomaly"] = d.ConcurrentCaptureCreatedByCtrAnomaly

	in["concurrent_capture_create_failed_by_ctr_increment"] = d.ConcurrentCaptureCreateFailedByCtrIncrement

	in["concurrent_capture_create_failed_by_ctr_anomaly"] = d.ConcurrentCaptureCreateFailedByCtrAnomaly

	in["concurrent_capture_create_failed_by_other_feature"] = d.ConcurrentCaptureCreateFailedByOtherFeature

	in["concurrent_capture_create_failed_oom"] = d.ConcurrentCaptureCreateFailedOom

	in["concurrent_capture_limit_reached"] = d.ConcurrentCaptureLimitReached

	in["concurrent_capture_by_ctr_increment_freed"] = d.ConcurrentCaptureByCtrIncrementFreed

	in["concurrent_capture_by_ctr_anomaly_freed"] = d.ConcurrentCaptureByCtrAnomalyFreed

	in["concurrent_capture_by_ctr_other_feature_freed"] = d.ConcurrentCaptureByCtrOtherFeatureFreed

	in["global_capture_finished"] = d.GlobalCaptureFinished

	in["concurrent_capture_finished"] = d.ConcurrentCaptureFinished

	in["pktcapture_with_no_conn_success"] = d.PktcaptureWithNoConnSuccess

	in["pktcapture_with_no_conn_failure"] = d.PktcaptureWithNoConnFailure

	in["pktcapture_with_conn_but_not_tagged_success"] = d.PktcaptureWithConnButNotTaggedSuccess

	in["pktcapture_with_conn_but_not_tagged_failure"] = d.PktcaptureWithConnButNotTaggedFailure

	in["pktcapture_with_conn_success_global"] = d.PktcaptureWithConnSuccessGlobal

	in["pktcapture_with_conn_success"] = d.PktcaptureWithConnSuccess

	in["pktcapture_with_conn_failure_global"] = d.PktcaptureWithConnFailureGlobal

	in["pktcapture_with_conn_failure"] = d.PktcaptureWithConnFailure

	in["pktcapture_failure_wait_for_block"] = d.PktcaptureFailureWaitForBlock

	in["pktcapture_failure_file_size_rchd"] = d.PktcaptureFailureFileSizeRchd

	in["num_conns_tagged_global_increment"] = d.NumConnsTaggedGlobalIncrement

	in["num_conns_tagged_global_anomaly"] = d.NumConnsTaggedGlobalAnomaly

	in["num_conns_tagged_global_other_feature"] = d.NumConnsTaggedGlobalOtherFeature

	in["num_conns_tagged_global_increment_fail"] = d.NumConnsTaggedGlobalIncrementFail

	in["num_conns_tagged_global_anomaly_fail"] = d.NumConnsTaggedGlobalAnomalyFail

	in["num_conns_tagged_global_other_feature_fail"] = d.NumConnsTaggedGlobalOtherFeatureFail

	in["num_conns_tagged_global_increment_maxed"] = d.NumConnsTaggedGlobalIncrementMaxed

	in["num_conns_tagged_global_anomaly_maxed"] = d.NumConnsTaggedGlobalAnomalyMaxed

	in["num_conns_tagged_global_other_feature_maxed"] = d.NumConnsTaggedGlobalOtherFeatureMaxed

	in["num_conns_tagged_increment"] = d.NumConnsTaggedIncrement

	in["num_conns_tagged_anomaly"] = d.NumConnsTaggedAnomaly

	in["num_conns_tagged_other_feature"] = d.NumConnsTaggedOtherFeature

	in["num_conns_tagged_increment_fail"] = d.NumConnsTaggedIncrementFail

	in["num_conns_tagged_anomaly_fail"] = d.NumConnsTaggedAnomalyFail

	in["num_conns_tagged_other_feature_fail"] = d.NumConnsTaggedOtherFeatureFail

	in["num_conns_tagged_increment_maxed"] = d.NumConnsTaggedIncrementMaxed

	in["num_conns_tagged_anomaly_maxed"] = d.NumConnsTaggedAnomalyMaxed

	in["num_conns_tagged_other_feature_maxed"] = d.NumConnsTaggedOtherFeatureMaxed

	in["num_conns_untagged"] = d.NumConnsUntagged

	in["pktcapture_triggered_by_increment"] = d.PktcaptureTriggeredByIncrement

	in["pktcapture_triggered_by_anomaly"] = d.PktcaptureTriggeredByAnomaly

	in["pktcapture_triggered_by_other_feature"] = d.PktcaptureTriggeredByOtherFeature

	in["num_of_anomalies_detected"] = d.NumOfAnomaliesDetected

	in["num_of_anomalies_cleared"] = d.NumOfAnomaliesCleared

	in["num_pcaps_created"] = d.NumPcapsCreated

	in["num_tmp_pcaps_created"] = d.NumTmpPcapsCreated

	in["num_pcaps_create_failed"] = d.NumPcapsCreateFailed

	in["pktcap_oom"] = d.PktcapOom

	in["failed_disk_full"] = d.FailedDiskFull

	in["conn_ext_failed"] = d.ConnExtFailed

	in["skip_as_conn_already_recapture"] = d.SkipAsConnAlreadyRecapture

	in["skip_capture_as_conn_created_before_smp"] = d.SkipCaptureAsConnCreatedBeforeSmp

	in["failed_as_return_completed_set"] = d.FailedAsReturnCompletedSet

	in["non_pkt_path"] = d.NonPktPath

	in["pkt_already_captured"] = d.PktAlreadyCaptured

	in["wrong_ctr_incremented"] = d.WrongCtrIncremented

	in["auto_pcap_file_merged"] = d.AutoPcapFileMerged

	in["auto_pcap_file_merged_failed"] = d.AutoPcapFileMergedFailed

	in["num_dynamic_capture_config_created"] = d.NumDynamicCaptureConfigCreated

	in["num_dynamic_capture_config_delete_q"] = d.NumDynamicCaptureConfigDeleteQ

	in["num_dynamic_capture_config_deleted"] = d.NumDynamicCaptureConfigDeleted

	in["num_global_counters_registered"] = d.NumGlobalCountersRegistered

	in["num_global_counters_deregistered"] = d.NumGlobalCountersDeregistered

	in["num_per_object_counters_registered"] = d.NumPerObjectCountersRegistered

	in["num_per_object_counters_deregistered"] = d.NumPerObjectCountersDeregistered
	result = append(result, in)
	return result
}

func setObjectVisibilityStatsPacketCaptureAutomatedCaptures(d edpt.VisibilityStatsPacketCaptureAutomatedCaptures) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["stats"] = setObjectVisibilityStatsPacketCaptureAutomatedCapturesStats(d.Stats)
	result = append(result, in)
	return result
}

func setObjectVisibilityStatsPacketCaptureAutomatedCapturesStats(d edpt.VisibilityStatsPacketCaptureAutomatedCapturesStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["total_failure"] = d.TotalFailure
	result = append(result, in)
	return result
}

func setObjectVisibilityStatsReporting(ret edpt.DataVisibilityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVisibilityStatsReportingStats(ret.DtVisibilityStats.Reporting.Stats),
		},
	}
}

func setObjectVisibilityStatsReportingStats(d edpt.VisibilityStatsReportingStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["log_transmit_failure"] = d.LogTransmitFailure

	in["buffer_alloc_failure"] = d.BufferAllocFailure

	in["notif_jobs_in_queue"] = d.NotifJobsInQueue

	in["enqueue_fail"] = d.EnqueueFail

	in["enqueue_pass"] = d.EnqueuePass

	in["dequeued"] = d.Dequeued
	result = append(result, in)
	return result
}

func setObjectVisibilityStatsStats(ret edpt.DataVisibilityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mon_entity_limit_exceed":                     ret.DtVisibilityStats.Stats.MonEntityLimitExceed,
			"ha_entity_create_sent":                       ret.DtVisibilityStats.Stats.HaEntityCreateSent,
			"ha_entity_delete_sent":                       ret.DtVisibilityStats.Stats.HaEntityDeleteSent,
			"ha_entity_anomaly_on_sent":                   ret.DtVisibilityStats.Stats.HaEntityAnomalyOnSent,
			"ha_entity_anomaly_off_sent":                  ret.DtVisibilityStats.Stats.HaEntityAnomalyOffSent,
			"ha_entity_periodic_sync_sent":                ret.DtVisibilityStats.Stats.HaEntityPeriodicSyncSent,
			"out_of_memory_alloc_failures":                ret.DtVisibilityStats.Stats.OutOfMemoryAllocFailures,
			"lw_mon_entity_created":                       ret.DtVisibilityStats.Stats.LwMonEntityCreated,
			"lw_mon_entity_deleted":                       ret.DtVisibilityStats.Stats.LwMonEntityDeleted,
			"lw_mon_entity_limit_exceed":                  ret.DtVisibilityStats.Stats.LwMonEntityLimitExceed,
			"lw_out_of_memory_alloc_failures":             ret.DtVisibilityStats.Stats.LwOutOfMemoryAllocFailures,
			"mon_entity_rrd_file_timestamp_err":           ret.DtVisibilityStats.Stats.MonEntityRrdFileTimestampErr,
			"mon_entity_rrd_update_err":                   ret.DtVisibilityStats.Stats.MonEntityRrdUpdateErr,
			"mon_entity_rrd_last_update_fetch_failed_err": ret.DtVisibilityStats.Stats.MonEntityRrdLastUpdateFetchFailedErr,
			"mon_entity_rrd_tune_err":                     ret.DtVisibilityStats.Stats.MonEntityRrdTuneErr,
			"mon_entity_rrd_out_of_memory_err":            ret.DtVisibilityStats.Stats.MonEntityRrdOutOfMemoryErr,
			"mon_entity_rrd_file_create_err":              ret.DtVisibilityStats.Stats.MonEntityRrdFileCreateErr,
		},
	}
}

func setObjectVisibilityStatsTopn(ret edpt.DataVisibilityStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVisibilityStatsTopnStats(ret.DtVisibilityStats.Topn.Stats),
		},
	}
}

func setObjectVisibilityStatsTopnStats(d edpt.VisibilityStatsTopnStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["heap_alloc_success"] = d.HeapAllocSuccess

	in["heap_alloc_failed"] = d.HeapAllocFailed

	in["heap_alloc_oom"] = d.HeapAllocOom

	in["obj_reg_success"] = d.ObjRegSuccess

	in["obj_reg_failed"] = d.ObjRegFailed

	in["obj_reg_oom"] = d.ObjRegOom

	in["heap_deleted"] = d.HeapDeleted

	in["obj_deleted"] = d.ObjDeleted
	result = append(result, in)
	return result
}

func getObjectVisibilityStatsMonEntityTelemetryData(d []interface{}) edpt.VisibilityStatsMonEntityTelemetryData {

	count1 := len(d)
	var ret edpt.VisibilityStatsMonEntityTelemetryData
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityStatsMonEntityTelemetryDataStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityStatsMonEntityTelemetryDataStats(d []interface{}) edpt.VisibilityStatsMonEntityTelemetryDataStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsMonEntityTelemetryDataStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.In_pkts = in["in_pkts"].(int)
		ret.Out_pkts = in["out_pkts"].(int)
		ret.In_bytes = in["in_bytes"].(int)
		ret.Out_bytes = in["out_bytes"].(int)
		ret.Errors = in["errors"].(int)
		ret.In_small_pkt = in["in_small_pkt"].(int)
		ret.In_frag = in["in_frag"].(int)
		ret.Out_small_pkt = in["out_small_pkt"].(int)
		ret.Out_frag = in["out_frag"].(int)
		ret.NewConn = in["new_conn"].(int)
		ret.Avg_data_cpu_util = in["avg_data_cpu_util"].(int)
		ret.Outside_intf_util = in["outside_intf_util"].(int)
		ret.ConcurrentConn = in["concurrent_conn"].(int)
		ret.In_bytes_per_out_bytes = in["in_bytes_per_out_bytes"].(int)
		ret.Drop_pkts_per_pkts = in["drop_pkts_per_pkts"].(int)
		ret.Tcp_in_syn = in["tcp_in_syn"].(int)
		ret.Tcp_out_syn = in["tcp_out_syn"].(int)
		ret.Tcp_in_fin = in["tcp_in_fin"].(int)
		ret.Tcp_out_fin = in["tcp_out_fin"].(int)
		ret.Tcp_in_payload = in["tcp_in_payload"].(int)
		ret.Tcp_out_payload = in["tcp_out_payload"].(int)
		ret.Tcp_in_rexmit = in["tcp_in_rexmit"].(int)
		ret.Tcp_out_rexmit = in["tcp_out_rexmit"].(int)
		ret.Tcp_in_rst = in["tcp_in_rst"].(int)
		ret.Tcp_out_rst = in["tcp_out_rst"].(int)
		ret.Tcp_in_empty_ack = in["tcp_in_empty_ack"].(int)
		ret.Tcp_out_empty_ack = in["tcp_out_empty_ack"].(int)
		ret.Tcp_in_zero_wnd = in["tcp_in_zero_wnd"].(int)
		ret.Tcp_out_zero_wnd = in["tcp_out_zero_wnd"].(int)
		ret.Tcp_conn_miss = in["tcp_conn_miss"].(int)
		ret.Tcp_fwd_syn_per_fin = in["tcp_fwd_syn_per_fin"].(int)
	}
	return ret
}

func getObjectVisibilityStatsPacketCapture(d []interface{}) edpt.VisibilityStatsPacketCapture {

	count1 := len(d)
	var ret edpt.VisibilityStatsPacketCapture
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityStatsPacketCaptureStats(in["stats"].([]interface{}))
		ret.CaptureConfigList = getSliceVisibilityStatsPacketCaptureCaptureConfigList(in["capture_config_list"].([]interface{}))
		ret.AutomatedCaptures = getObjectVisibilityStatsPacketCaptureAutomatedCaptures(in["automated_captures"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityStatsPacketCaptureStats(d []interface{}) edpt.VisibilityStatsPacketCaptureStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsPacketCaptureStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentCaptureCreatedByCtrIncrement = in["concurrent_capture_created_by_ctr_increment"].(int)
		ret.ConcurrentCaptureCreatedByCtrAnomaly = in["concurrent_capture_created_by_ctr_anomaly"].(int)
		ret.ConcurrentCaptureCreatedByOtherFeature = in["concurrent_capture_created_by_other_feature"].(int)
		ret.ConcurrentCaptureCreateFailedByCtrIncrement = in["concurrent_capture_create_failed_by_ctr_increment"].(int)
		ret.ConcurrentCaptureCreateFailedByCtrAnomaly = in["concurrent_capture_create_failed_by_ctr_anomaly"].(int)
		ret.ConcurrentCaptureCreateFailedByOtherFeature = in["concurrent_capture_create_failed_by_other_feature"].(int)
		ret.ConcurrentCaptureCreateFailedOom = in["concurrent_capture_create_failed_oom"].(int)
		ret.ConcurrentCaptureLimitReached = in["concurrent_capture_limit_reached"].(int)
		ret.ConcurrentCaptureByCtrIncrementFreed = in["concurrent_capture_by_ctr_increment_freed"].(int)
		ret.ConcurrentCaptureByCtrAnomalyFreed = in["concurrent_capture_by_ctr_anomaly_freed"].(int)
		ret.ConcurrentCaptureByCtrOtherFeatureFreed = in["concurrent_capture_by_ctr_other_feature_freed"].(int)
		ret.GlobalCaptureFinished = in["global_capture_finished"].(int)
		ret.ConcurrentCaptureFinished = in["concurrent_capture_finished"].(int)
		ret.PktcaptureWithNoConnSuccess = in["pktcapture_with_no_conn_success"].(int)
		ret.PktcaptureWithNoConnFailure = in["pktcapture_with_no_conn_failure"].(int)
		ret.PktcaptureWithConnButNotTaggedSuccess = in["pktcapture_with_conn_but_not_tagged_success"].(int)
		ret.PktcaptureWithConnButNotTaggedFailure = in["pktcapture_with_conn_but_not_tagged_failure"].(int)
		ret.PktcaptureWithConnSuccessGlobal = in["pktcapture_with_conn_success_global"].(int)
		ret.PktcaptureWithConnSuccess = in["pktcapture_with_conn_success"].(int)
		ret.PktcaptureWithConnFailureGlobal = in["pktcapture_with_conn_failure_global"].(int)
		ret.PktcaptureWithConnFailure = in["pktcapture_with_conn_failure"].(int)
		ret.PktcaptureFailureWaitForBlock = in["pktcapture_failure_wait_for_block"].(int)
		ret.PktcaptureFailureFileSizeRchd = in["pktcapture_failure_file_size_rchd"].(int)
		ret.NumConnsTaggedGlobalIncrement = in["num_conns_tagged_global_increment"].(int)
		ret.NumConnsTaggedGlobalAnomaly = in["num_conns_tagged_global_anomaly"].(int)
		ret.NumConnsTaggedGlobalOtherFeature = in["num_conns_tagged_global_other_feature"].(int)
		ret.NumConnsTaggedGlobalIncrementFail = in["num_conns_tagged_global_increment_fail"].(int)
		ret.NumConnsTaggedGlobalAnomalyFail = in["num_conns_tagged_global_anomaly_fail"].(int)
		ret.NumConnsTaggedGlobalOtherFeatureFail = in["num_conns_tagged_global_other_feature_fail"].(int)
		ret.NumConnsTaggedGlobalIncrementMaxed = in["num_conns_tagged_global_increment_maxed"].(int)
		ret.NumConnsTaggedGlobalAnomalyMaxed = in["num_conns_tagged_global_anomaly_maxed"].(int)
		ret.NumConnsTaggedGlobalOtherFeatureMaxed = in["num_conns_tagged_global_other_feature_maxed"].(int)
		ret.NumConnsTaggedIncrement = in["num_conns_tagged_increment"].(int)
		ret.NumConnsTaggedAnomaly = in["num_conns_tagged_anomaly"].(int)
		ret.NumConnsTaggedOtherFeature = in["num_conns_tagged_other_feature"].(int)
		ret.NumConnsTaggedIncrementFail = in["num_conns_tagged_increment_fail"].(int)
		ret.NumConnsTaggedAnomalyFail = in["num_conns_tagged_anomaly_fail"].(int)
		ret.NumConnsTaggedOtherFeatureFail = in["num_conns_tagged_other_feature_fail"].(int)
		ret.NumConnsTaggedIncrementMaxed = in["num_conns_tagged_increment_maxed"].(int)
		ret.NumConnsTaggedAnomalyMaxed = in["num_conns_tagged_anomaly_maxed"].(int)
		ret.NumConnsTaggedOtherFeatureMaxed = in["num_conns_tagged_other_feature_maxed"].(int)
		ret.NumConnsUntagged = in["num_conns_untagged"].(int)
		ret.PktcaptureTriggeredByIncrement = in["pktcapture_triggered_by_increment"].(int)
		ret.PktcaptureTriggeredByAnomaly = in["pktcapture_triggered_by_anomaly"].(int)
		ret.PktcaptureTriggeredByOtherFeature = in["pktcapture_triggered_by_other_feature"].(int)
		ret.NumOfAnomaliesDetected = in["num_of_anomalies_detected"].(int)
		ret.NumOfAnomaliesCleared = in["num_of_anomalies_cleared"].(int)
		ret.NumPcapsCreated = in["num_pcaps_created"].(int)
		ret.NumTmpPcapsCreated = in["num_tmp_pcaps_created"].(int)
		ret.NumPcapsCreateFailed = in["num_pcaps_create_failed"].(int)
		ret.PktcapOom = in["pktcap_oom"].(int)
		ret.FailedDiskFull = in["failed_disk_full"].(int)
		ret.ConnExtFailed = in["conn_ext_failed"].(int)
		ret.SkipAsConnAlreadyRecapture = in["skip_as_conn_already_recapture"].(int)
		ret.SkipCaptureAsConnCreatedBeforeSmp = in["skip_capture_as_conn_created_before_smp"].(int)
		ret.FailedAsReturnCompletedSet = in["failed_as_return_completed_set"].(int)
		ret.NonPktPath = in["non_pkt_path"].(int)
		ret.PktAlreadyCaptured = in["pkt_already_captured"].(int)
		ret.WrongCtrIncremented = in["wrong_ctr_incremented"].(int)
		ret.AutoPcapFileMerged = in["auto_pcap_file_merged"].(int)
		ret.AutoPcapFileMergedFailed = in["auto_pcap_file_merged_failed"].(int)
		ret.NumGlobalTmplCreated = in["num_global_tmpl_created"].(int)
		ret.NumObjectTmplCreated = in["num_object_tmpl_created"].(int)
		ret.NumGlobalTmplDeleted = in["num_global_tmpl_deleted"].(int)
		ret.NumObjectTmplDeleted = in["num_object_tmpl_deleted"].(int)
		ret.NumCaptureConfigCreated = in["num_capture_config_created"].(int)
		ret.NumDynamicCaptureConfigCreated = in["num_dynamic_capture_config_created"].(int)
		ret.NumCaptureConfigDeleted = in["num_capture_config_deleted"].(int)
		ret.NumDynamicCaptureConfigDeleted = in["num_dynamic_capture_config_deleted"].(int)
		ret.NumCaptureConfigDeleteQ = in["num_capture_config_delete_q"].(int)
		ret.NumDynamicCaptureConfigDeleteQ = in["num_dynamic_capture_config_delete_q"].(int)
		ret.NumCaptureConfigLinked = in["num_capture_config_linked"].(int)
		ret.NumDynamicCaptureConfigLinked = in["num_dynamic_capture_config_linked"].(int)
		ret.NumCaptureConfigUnlinked = in["num_capture_config_unlinked"].(int)
		ret.NumDynamicCaptureConfigUnlinked = in["num_dynamic_capture_config_unlinked"].(int)
		ret.NumGlobalCountersRegistered = in["num_global_counters_registered"].(int)
		ret.NumGlobalCountersDeregistered = in["num_global_counters_deregistered"].(int)
		ret.NumPerObjectCountersRegistered = in["num_per_object_counters_registered"].(int)
		ret.NumPerObjectCountersDeregistered = in["num_per_object_counters_deregistered"].(int)
	}
	return ret
}

func getSliceVisibilityStatsPacketCaptureCaptureConfigList(d []interface{}) []edpt.VisibilityStatsPacketCaptureCaptureConfigList {

	count1 := len(d)
	ret := make([]edpt.VisibilityStatsPacketCaptureCaptureConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityStatsPacketCaptureCaptureConfigList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectVisibilityStatsPacketCaptureCaptureConfigListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityStatsPacketCaptureCaptureConfigListStats(d []interface{}) edpt.VisibilityStatsPacketCaptureCaptureConfigListStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsPacketCaptureCaptureConfigListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConcurrentCaptureCreatedByCtrIncrement = in["concurrent_capture_created_by_ctr_increment"].(int)
		ret.ConcurrentCaptureCreatedByCtrAnomaly = in["concurrent_capture_created_by_ctr_anomaly"].(int)
		ret.ConcurrentCaptureCreateFailedByCtrIncrement = in["concurrent_capture_create_failed_by_ctr_increment"].(int)
		ret.ConcurrentCaptureCreateFailedByCtrAnomaly = in["concurrent_capture_create_failed_by_ctr_anomaly"].(int)
		ret.ConcurrentCaptureCreateFailedByOtherFeature = in["concurrent_capture_create_failed_by_other_feature"].(int)
		ret.ConcurrentCaptureCreateFailedOom = in["concurrent_capture_create_failed_oom"].(int)
		ret.ConcurrentCaptureLimitReached = in["concurrent_capture_limit_reached"].(int)
		ret.ConcurrentCaptureByCtrIncrementFreed = in["concurrent_capture_by_ctr_increment_freed"].(int)
		ret.ConcurrentCaptureByCtrAnomalyFreed = in["concurrent_capture_by_ctr_anomaly_freed"].(int)
		ret.ConcurrentCaptureByCtrOtherFeatureFreed = in["concurrent_capture_by_ctr_other_feature_freed"].(int)
		ret.GlobalCaptureFinished = in["global_capture_finished"].(int)
		ret.ConcurrentCaptureFinished = in["concurrent_capture_finished"].(int)
		ret.PktcaptureWithNoConnSuccess = in["pktcapture_with_no_conn_success"].(int)
		ret.PktcaptureWithNoConnFailure = in["pktcapture_with_no_conn_failure"].(int)
		ret.PktcaptureWithConnButNotTaggedSuccess = in["pktcapture_with_conn_but_not_tagged_success"].(int)
		ret.PktcaptureWithConnButNotTaggedFailure = in["pktcapture_with_conn_but_not_tagged_failure"].(int)
		ret.PktcaptureWithConnSuccessGlobal = in["pktcapture_with_conn_success_global"].(int)
		ret.PktcaptureWithConnSuccess = in["pktcapture_with_conn_success"].(int)
		ret.PktcaptureWithConnFailureGlobal = in["pktcapture_with_conn_failure_global"].(int)
		ret.PktcaptureWithConnFailure = in["pktcapture_with_conn_failure"].(int)
		ret.PktcaptureFailureWaitForBlock = in["pktcapture_failure_wait_for_block"].(int)
		ret.PktcaptureFailureFileSizeRchd = in["pktcapture_failure_file_size_rchd"].(int)
		ret.NumConnsTaggedGlobalIncrement = in["num_conns_tagged_global_increment"].(int)
		ret.NumConnsTaggedGlobalAnomaly = in["num_conns_tagged_global_anomaly"].(int)
		ret.NumConnsTaggedGlobalOtherFeature = in["num_conns_tagged_global_other_feature"].(int)
		ret.NumConnsTaggedGlobalIncrementFail = in["num_conns_tagged_global_increment_fail"].(int)
		ret.NumConnsTaggedGlobalAnomalyFail = in["num_conns_tagged_global_anomaly_fail"].(int)
		ret.NumConnsTaggedGlobalOtherFeatureFail = in["num_conns_tagged_global_other_feature_fail"].(int)
		ret.NumConnsTaggedGlobalIncrementMaxed = in["num_conns_tagged_global_increment_maxed"].(int)
		ret.NumConnsTaggedGlobalAnomalyMaxed = in["num_conns_tagged_global_anomaly_maxed"].(int)
		ret.NumConnsTaggedGlobalOtherFeatureMaxed = in["num_conns_tagged_global_other_feature_maxed"].(int)
		ret.NumConnsTaggedIncrement = in["num_conns_tagged_increment"].(int)
		ret.NumConnsTaggedAnomaly = in["num_conns_tagged_anomaly"].(int)
		ret.NumConnsTaggedOtherFeature = in["num_conns_tagged_other_feature"].(int)
		ret.NumConnsTaggedIncrementFail = in["num_conns_tagged_increment_fail"].(int)
		ret.NumConnsTaggedAnomalyFail = in["num_conns_tagged_anomaly_fail"].(int)
		ret.NumConnsTaggedOtherFeatureFail = in["num_conns_tagged_other_feature_fail"].(int)
		ret.NumConnsTaggedIncrementMaxed = in["num_conns_tagged_increment_maxed"].(int)
		ret.NumConnsTaggedAnomalyMaxed = in["num_conns_tagged_anomaly_maxed"].(int)
		ret.NumConnsTaggedOtherFeatureMaxed = in["num_conns_tagged_other_feature_maxed"].(int)
		ret.NumConnsUntagged = in["num_conns_untagged"].(int)
		ret.PktcaptureTriggeredByIncrement = in["pktcapture_triggered_by_increment"].(int)
		ret.PktcaptureTriggeredByAnomaly = in["pktcapture_triggered_by_anomaly"].(int)
		ret.PktcaptureTriggeredByOtherFeature = in["pktcapture_triggered_by_other_feature"].(int)
		ret.NumOfAnomaliesDetected = in["num_of_anomalies_detected"].(int)
		ret.NumOfAnomaliesCleared = in["num_of_anomalies_cleared"].(int)
		ret.NumPcapsCreated = in["num_pcaps_created"].(int)
		ret.NumTmpPcapsCreated = in["num_tmp_pcaps_created"].(int)
		ret.NumPcapsCreateFailed = in["num_pcaps_create_failed"].(int)
		ret.PktcapOom = in["pktcap_oom"].(int)
		ret.FailedDiskFull = in["failed_disk_full"].(int)
		ret.ConnExtFailed = in["conn_ext_failed"].(int)
		ret.SkipAsConnAlreadyRecapture = in["skip_as_conn_already_recapture"].(int)
		ret.SkipCaptureAsConnCreatedBeforeSmp = in["skip_capture_as_conn_created_before_smp"].(int)
		ret.FailedAsReturnCompletedSet = in["failed_as_return_completed_set"].(int)
		ret.NonPktPath = in["non_pkt_path"].(int)
		ret.PktAlreadyCaptured = in["pkt_already_captured"].(int)
		ret.WrongCtrIncremented = in["wrong_ctr_incremented"].(int)
		ret.AutoPcapFileMerged = in["auto_pcap_file_merged"].(int)
		ret.AutoPcapFileMergedFailed = in["auto_pcap_file_merged_failed"].(int)
		ret.NumDynamicCaptureConfigCreated = in["num_dynamic_capture_config_created"].(int)
		ret.NumDynamicCaptureConfigDeleteQ = in["num_dynamic_capture_config_delete_q"].(int)
		ret.NumDynamicCaptureConfigDeleted = in["num_dynamic_capture_config_deleted"].(int)
		ret.NumGlobalCountersRegistered = in["num_global_counters_registered"].(int)
		ret.NumGlobalCountersDeregistered = in["num_global_counters_deregistered"].(int)
		ret.NumPerObjectCountersRegistered = in["num_per_object_counters_registered"].(int)
		ret.NumPerObjectCountersDeregistered = in["num_per_object_counters_deregistered"].(int)
	}
	return ret
}

func getObjectVisibilityStatsPacketCaptureAutomatedCaptures(d []interface{}) edpt.VisibilityStatsPacketCaptureAutomatedCaptures {

	count1 := len(d)
	var ret edpt.VisibilityStatsPacketCaptureAutomatedCaptures
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityStatsPacketCaptureAutomatedCapturesStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityStatsPacketCaptureAutomatedCapturesStats(d []interface{}) edpt.VisibilityStatsPacketCaptureAutomatedCapturesStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsPacketCaptureAutomatedCapturesStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalFailure = in["total_failure"].(int)
	}
	return ret
}

func getObjectVisibilityStatsReporting(d []interface{}) edpt.VisibilityStatsReporting {

	count1 := len(d)
	var ret edpt.VisibilityStatsReporting
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityStatsReportingStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityStatsReportingStats(d []interface{}) edpt.VisibilityStatsReportingStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsReportingStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogTransmitFailure = in["log_transmit_failure"].(int)
		ret.BufferAllocFailure = in["buffer_alloc_failure"].(int)
		ret.NotifJobsInQueue = in["notif_jobs_in_queue"].(int)
		ret.EnqueueFail = in["enqueue_fail"].(int)
		ret.EnqueuePass = in["enqueue_pass"].(int)
		ret.Dequeued = in["dequeued"].(int)
	}
	return ret
}

func getObjectVisibilityStatsStats(d []interface{}) edpt.VisibilityStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonEntityLimitExceed = in["mon_entity_limit_exceed"].(int)
		ret.HaEntityCreateSent = in["ha_entity_create_sent"].(int)
		ret.HaEntityDeleteSent = in["ha_entity_delete_sent"].(int)
		ret.HaEntityAnomalyOnSent = in["ha_entity_anomaly_on_sent"].(int)
		ret.HaEntityAnomalyOffSent = in["ha_entity_anomaly_off_sent"].(int)
		ret.HaEntityPeriodicSyncSent = in["ha_entity_periodic_sync_sent"].(int)
		ret.OutOfMemoryAllocFailures = in["out_of_memory_alloc_failures"].(int)
		ret.LwMonEntityCreated = in["lw_mon_entity_created"].(int)
		ret.LwMonEntityDeleted = in["lw_mon_entity_deleted"].(int)
		ret.LwMonEntityLimitExceed = in["lw_mon_entity_limit_exceed"].(int)
		ret.LwOutOfMemoryAllocFailures = in["lw_out_of_memory_alloc_failures"].(int)
		ret.MonEntityRrdFileTimestampErr = in["mon_entity_rrd_file_timestamp_err"].(int)
		ret.MonEntityRrdUpdateErr = in["mon_entity_rrd_update_err"].(int)
		ret.MonEntityRrdLastUpdateFetchFailedErr = in["mon_entity_rrd_last_update_fetch_failed_err"].(int)
		ret.MonEntityRrdTuneErr = in["mon_entity_rrd_tune_err"].(int)
		ret.MonEntityRrdOutOfMemoryErr = in["mon_entity_rrd_out_of_memory_err"].(int)
		ret.MonEntityRrdFileCreateErr = in["mon_entity_rrd_file_create_err"].(int)
	}
	return ret
}

func getObjectVisibilityStatsTopn(d []interface{}) edpt.VisibilityStatsTopn {

	count1 := len(d)
	var ret edpt.VisibilityStatsTopn
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityStatsTopnStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityStatsTopnStats(d []interface{}) edpt.VisibilityStatsTopnStats {

	count1 := len(d)
	var ret edpt.VisibilityStatsTopnStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HeapAllocSuccess = in["heap_alloc_success"].(int)
		ret.HeapAllocFailed = in["heap_alloc_failed"].(int)
		ret.HeapAllocOom = in["heap_alloc_oom"].(int)
		ret.ObjRegSuccess = in["obj_reg_success"].(int)
		ret.ObjRegFailed = in["obj_reg_failed"].(int)
		ret.ObjRegOom = in["obj_reg_oom"].(int)
		ret.HeapDeleted = in["heap_deleted"].(int)
		ret.ObjDeleted = in["obj_deleted"].(int)
	}
	return ret
}

func dataToEndpointVisibilityStats(d *schema.ResourceData) edpt.VisibilityStats {
	var ret edpt.VisibilityStats

	ret.MonEntityTelemetryData = getObjectVisibilityStatsMonEntityTelemetryData(d.Get("mon_entity_telemetry_data").([]interface{}))

	ret.PacketCapture = getObjectVisibilityStatsPacketCapture(d.Get("packet_capture").([]interface{}))

	ret.Reporting = getObjectVisibilityStatsReporting(d.Get("reporting").([]interface{}))

	ret.Stats = getObjectVisibilityStatsStats(d.Get("stats").([]interface{}))

	ret.Topn = getObjectVisibilityStatsTopn(d.Get("topn").([]interface{}))
	return ret
}
