package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_packet_capture_stats`: Statistics for the object packet-capture\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPacketCaptureStatsRead,

		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceVisibilityPacketCaptureStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPacketCaptureStatsAutomatedCaptures := setObjectVisibilityPacketCaptureStatsAutomatedCaptures(res)
		d.Set("automated_captures", VisibilityPacketCaptureStatsAutomatedCaptures)
		VisibilityPacketCaptureStatsCaptureConfigList := setSliceVisibilityPacketCaptureStatsCaptureConfigList(res)
		d.Set("capture_config_list", VisibilityPacketCaptureStatsCaptureConfigList)
		VisibilityPacketCaptureStatsStats := setObjectVisibilityPacketCaptureStatsStats(res)
		d.Set("stats", VisibilityPacketCaptureStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPacketCaptureStatsAutomatedCaptures(ret edpt.DataVisibilityPacketCaptureStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectVisibilityPacketCaptureStatsAutomatedCapturesStats(ret.DtVisibilityPacketCaptureStats.AutomatedCaptures.Stats),
		},
	}
}

func setObjectVisibilityPacketCaptureStatsAutomatedCapturesStats(d edpt.VisibilityPacketCaptureStatsAutomatedCapturesStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["total_failure"] = d.TotalFailure
	result = append(result, in)
	return result
}

func setSliceVisibilityPacketCaptureStatsCaptureConfigList(d edpt.DataVisibilityPacketCaptureStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtVisibilityPacketCaptureStats.CaptureConfigList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectVisibilityPacketCaptureStatsCaptureConfigListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityPacketCaptureStatsCaptureConfigListStats(d edpt.VisibilityPacketCaptureStatsCaptureConfigListStats) []map[string]interface{} {
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

func setObjectVisibilityPacketCaptureStatsStats(ret edpt.DataVisibilityPacketCaptureStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"concurrent_capture_created_by_ctr_increment":       ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreatedByCtrIncrement,
			"concurrent_capture_created_by_ctr_anomaly":         ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreatedByCtrAnomaly,
			"concurrent_capture_created_by_other_feature":       ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreatedByOtherFeature,
			"concurrent_capture_create_failed_by_ctr_increment": ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreateFailedByCtrIncrement,
			"concurrent_capture_create_failed_by_ctr_anomaly":   ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreateFailedByCtrAnomaly,
			"concurrent_capture_create_failed_by_other_feature": ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreateFailedByOtherFeature,
			"concurrent_capture_create_failed_oom":              ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureCreateFailedOom,
			"concurrent_capture_limit_reached":                  ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureLimitReached,
			"concurrent_capture_by_ctr_increment_freed":         ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureByCtrIncrementFreed,
			"concurrent_capture_by_ctr_anomaly_freed":           ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureByCtrAnomalyFreed,
			"concurrent_capture_by_ctr_other_feature_freed":     ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureByCtrOtherFeatureFreed,
			"global_capture_finished":                           ret.DtVisibilityPacketCaptureStats.Stats.GlobalCaptureFinished,
			"concurrent_capture_finished":                       ret.DtVisibilityPacketCaptureStats.Stats.ConcurrentCaptureFinished,
			"pktcapture_with_no_conn_success":                   ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithNoConnSuccess,
			"pktcapture_with_no_conn_failure":                   ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithNoConnFailure,
			"pktcapture_with_conn_but_not_tagged_success":       ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithConnButNotTaggedSuccess,
			"pktcapture_with_conn_but_not_tagged_failure":       ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithConnButNotTaggedFailure,
			"pktcapture_with_conn_success_global":               ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithConnSuccessGlobal,
			"pktcapture_with_conn_success":                      ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithConnSuccess,
			"pktcapture_with_conn_failure_global":               ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithConnFailureGlobal,
			"pktcapture_with_conn_failure":                      ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureWithConnFailure,
			"pktcapture_failure_wait_for_block":                 ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureFailureWaitForBlock,
			"pktcapture_failure_file_size_rchd":                 ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureFailureFileSizeRchd,
			"num_conns_tagged_global_increment":                 ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalIncrement,
			"num_conns_tagged_global_anomaly":                   ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalAnomaly,
			"num_conns_tagged_global_other_feature":             ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalOtherFeature,
			"num_conns_tagged_global_increment_fail":            ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalIncrementFail,
			"num_conns_tagged_global_anomaly_fail":              ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalAnomalyFail,
			"num_conns_tagged_global_other_feature_fail":        ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalOtherFeatureFail,
			"num_conns_tagged_global_increment_maxed":           ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalIncrementMaxed,
			"num_conns_tagged_global_anomaly_maxed":             ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalAnomalyMaxed,
			"num_conns_tagged_global_other_feature_maxed":       ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedGlobalOtherFeatureMaxed,
			"num_conns_tagged_increment":                        ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedIncrement,
			"num_conns_tagged_anomaly":                          ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedAnomaly,
			"num_conns_tagged_other_feature":                    ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedOtherFeature,
			"num_conns_tagged_increment_fail":                   ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedIncrementFail,
			"num_conns_tagged_anomaly_fail":                     ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedAnomalyFail,
			"num_conns_tagged_other_feature_fail":               ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedOtherFeatureFail,
			"num_conns_tagged_increment_maxed":                  ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedIncrementMaxed,
			"num_conns_tagged_anomaly_maxed":                    ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedAnomalyMaxed,
			"num_conns_tagged_other_feature_maxed":              ret.DtVisibilityPacketCaptureStats.Stats.NumConnsTaggedOtherFeatureMaxed,
			"num_conns_untagged":                                ret.DtVisibilityPacketCaptureStats.Stats.NumConnsUntagged,
			"pktcapture_triggered_by_increment":                 ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureTriggeredByIncrement,
			"pktcapture_triggered_by_anomaly":                   ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureTriggeredByAnomaly,
			"pktcapture_triggered_by_other_feature":             ret.DtVisibilityPacketCaptureStats.Stats.PktcaptureTriggeredByOtherFeature,
			"num_of_anomalies_detected":                         ret.DtVisibilityPacketCaptureStats.Stats.NumOfAnomaliesDetected,
			"num_of_anomalies_cleared":                          ret.DtVisibilityPacketCaptureStats.Stats.NumOfAnomaliesCleared,
			"num_pcaps_created":                                 ret.DtVisibilityPacketCaptureStats.Stats.NumPcapsCreated,
			"num_tmp_pcaps_created":                             ret.DtVisibilityPacketCaptureStats.Stats.NumTmpPcapsCreated,
			"num_pcaps_create_failed":                           ret.DtVisibilityPacketCaptureStats.Stats.NumPcapsCreateFailed,
			"pktcap_oom":                                        ret.DtVisibilityPacketCaptureStats.Stats.PktcapOom,
			"failed_disk_full":                                  ret.DtVisibilityPacketCaptureStats.Stats.FailedDiskFull,
			"conn_ext_failed":                                   ret.DtVisibilityPacketCaptureStats.Stats.ConnExtFailed,
			"skip_as_conn_already_recapture":                    ret.DtVisibilityPacketCaptureStats.Stats.SkipAsConnAlreadyRecapture,
			"skip_capture_as_conn_created_before_smp":           ret.DtVisibilityPacketCaptureStats.Stats.SkipCaptureAsConnCreatedBeforeSmp,
			"failed_as_return_completed_set":                    ret.DtVisibilityPacketCaptureStats.Stats.FailedAsReturnCompletedSet,
			"non_pkt_path":                                      ret.DtVisibilityPacketCaptureStats.Stats.NonPktPath,
			"pkt_already_captured":                              ret.DtVisibilityPacketCaptureStats.Stats.PktAlreadyCaptured,
			"wrong_ctr_incremented":                             ret.DtVisibilityPacketCaptureStats.Stats.WrongCtrIncremented,
			"auto_pcap_file_merged":                             ret.DtVisibilityPacketCaptureStats.Stats.AutoPcapFileMerged,
			"auto_pcap_file_merged_failed":                      ret.DtVisibilityPacketCaptureStats.Stats.AutoPcapFileMergedFailed,
			"num_global_tmpl_created":                           ret.DtVisibilityPacketCaptureStats.Stats.NumGlobalTmplCreated,
			"num_object_tmpl_created":                           ret.DtVisibilityPacketCaptureStats.Stats.NumObjectTmplCreated,
			"num_global_tmpl_deleted":                           ret.DtVisibilityPacketCaptureStats.Stats.NumGlobalTmplDeleted,
			"num_object_tmpl_deleted":                           ret.DtVisibilityPacketCaptureStats.Stats.NumObjectTmplDeleted,
			"num_capture_config_created":                        ret.DtVisibilityPacketCaptureStats.Stats.NumCaptureConfigCreated,
			"num_dynamic_capture_config_created":                ret.DtVisibilityPacketCaptureStats.Stats.NumDynamicCaptureConfigCreated,
			"num_capture_config_deleted":                        ret.DtVisibilityPacketCaptureStats.Stats.NumCaptureConfigDeleted,
			"num_dynamic_capture_config_deleted":                ret.DtVisibilityPacketCaptureStats.Stats.NumDynamicCaptureConfigDeleted,
			"num_capture_config_delete_q":                       ret.DtVisibilityPacketCaptureStats.Stats.NumCaptureConfigDeleteQ,
			"num_dynamic_capture_config_delete_q":               ret.DtVisibilityPacketCaptureStats.Stats.NumDynamicCaptureConfigDeleteQ,
			"num_capture_config_linked":                         ret.DtVisibilityPacketCaptureStats.Stats.NumCaptureConfigLinked,
			"num_dynamic_capture_config_linked":                 ret.DtVisibilityPacketCaptureStats.Stats.NumDynamicCaptureConfigLinked,
			"num_capture_config_unlinked":                       ret.DtVisibilityPacketCaptureStats.Stats.NumCaptureConfigUnlinked,
			"num_dynamic_capture_config_unlinked":               ret.DtVisibilityPacketCaptureStats.Stats.NumDynamicCaptureConfigUnlinked,
			"num_global_counters_registered":                    ret.DtVisibilityPacketCaptureStats.Stats.NumGlobalCountersRegistered,
			"num_global_counters_deregistered":                  ret.DtVisibilityPacketCaptureStats.Stats.NumGlobalCountersDeregistered,
			"num_per_object_counters_registered":                ret.DtVisibilityPacketCaptureStats.Stats.NumPerObjectCountersRegistered,
			"num_per_object_counters_deregistered":              ret.DtVisibilityPacketCaptureStats.Stats.NumPerObjectCountersDeregistered,
		},
	}
}

func getObjectVisibilityPacketCaptureStatsAutomatedCaptures(d []interface{}) edpt.VisibilityPacketCaptureStatsAutomatedCaptures {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureStatsAutomatedCaptures
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectVisibilityPacketCaptureStatsAutomatedCapturesStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityPacketCaptureStatsAutomatedCapturesStats(d []interface{}) edpt.VisibilityPacketCaptureStatsAutomatedCapturesStats {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureStatsAutomatedCapturesStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalFailure = in["total_failure"].(int)
	}
	return ret
}

func getSliceVisibilityPacketCaptureStatsCaptureConfigList(d []interface{}) []edpt.VisibilityPacketCaptureStatsCaptureConfigList {

	count1 := len(d)
	ret := make([]edpt.VisibilityPacketCaptureStatsCaptureConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityPacketCaptureStatsCaptureConfigList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectVisibilityPacketCaptureStatsCaptureConfigListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityPacketCaptureStatsCaptureConfigListStats(d []interface{}) edpt.VisibilityPacketCaptureStatsCaptureConfigListStats {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureStatsCaptureConfigListStats
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

func getObjectVisibilityPacketCaptureStatsStats(d []interface{}) edpt.VisibilityPacketCaptureStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureStatsStats
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

func dataToEndpointVisibilityPacketCaptureStats(d *schema.ResourceData) edpt.VisibilityPacketCaptureStats {
	var ret edpt.VisibilityPacketCaptureStats

	ret.AutomatedCaptures = getObjectVisibilityPacketCaptureStatsAutomatedCaptures(d.Get("automated_captures").([]interface{}))

	ret.CaptureConfigList = getSliceVisibilityPacketCaptureStatsCaptureConfigList(d.Get("capture_config_list").([]interface{}))

	ret.Stats = getObjectVisibilityPacketCaptureStatsStats(d.Get("stats").([]interface{}))
	return ret
}
