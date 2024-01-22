package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureCaptureConfigStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_packet_capture_capture_config_stats`: Statistics for the object capture-config\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityPacketCaptureCaptureConfigStatsRead,

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
	}
}

func resourceVisibilityPacketCaptureCaptureConfigStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureCaptureConfigStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureCaptureConfigStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityPacketCaptureCaptureConfigStatsStats := setObjectVisibilityPacketCaptureCaptureConfigStatsStats(res)
		d.Set("stats", VisibilityPacketCaptureCaptureConfigStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityPacketCaptureCaptureConfigStatsStats(ret edpt.DataVisibilityPacketCaptureCaptureConfigStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"concurrent_capture_created_by_ctr_increment":       ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureCreatedByCtrIncrement,
			"concurrent_capture_created_by_ctr_anomaly":         ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureCreatedByCtrAnomaly,
			"concurrent_capture_create_failed_by_ctr_increment": ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureCreateFailedByCtrIncrement,
			"concurrent_capture_create_failed_by_ctr_anomaly":   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureCreateFailedByCtrAnomaly,
			"concurrent_capture_create_failed_by_other_feature": ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureCreateFailedByOtherFeature,
			"concurrent_capture_create_failed_oom":              ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureCreateFailedOom,
			"concurrent_capture_limit_reached":                  ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureLimitReached,
			"concurrent_capture_by_ctr_increment_freed":         ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureByCtrIncrementFreed,
			"concurrent_capture_by_ctr_anomaly_freed":           ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureByCtrAnomalyFreed,
			"concurrent_capture_by_ctr_other_feature_freed":     ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureByCtrOtherFeatureFreed,
			"global_capture_finished":                           ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.GlobalCaptureFinished,
			"concurrent_capture_finished":                       ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConcurrentCaptureFinished,
			"pktcapture_with_no_conn_success":                   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithNoConnSuccess,
			"pktcapture_with_no_conn_failure":                   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithNoConnFailure,
			"pktcapture_with_conn_but_not_tagged_success":       ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithConnButNotTaggedSuccess,
			"pktcapture_with_conn_but_not_tagged_failure":       ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithConnButNotTaggedFailure,
			"pktcapture_with_conn_success_global":               ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithConnSuccessGlobal,
			"pktcapture_with_conn_success":                      ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithConnSuccess,
			"pktcapture_with_conn_failure_global":               ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithConnFailureGlobal,
			"pktcapture_with_conn_failure":                      ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureWithConnFailure,
			"pktcapture_failure_wait_for_block":                 ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureFailureWaitForBlock,
			"pktcapture_failure_file_size_rchd":                 ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureFailureFileSizeRchd,
			"num_conns_tagged_global_increment":                 ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalIncrement,
			"num_conns_tagged_global_anomaly":                   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalAnomaly,
			"num_conns_tagged_global_other_feature":             ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalOtherFeature,
			"num_conns_tagged_global_increment_fail":            ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalIncrementFail,
			"num_conns_tagged_global_anomaly_fail":              ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalAnomalyFail,
			"num_conns_tagged_global_other_feature_fail":        ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalOtherFeatureFail,
			"num_conns_tagged_global_increment_maxed":           ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalIncrementMaxed,
			"num_conns_tagged_global_anomaly_maxed":             ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalAnomalyMaxed,
			"num_conns_tagged_global_other_feature_maxed":       ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedGlobalOtherFeatureMaxed,
			"num_conns_tagged_increment":                        ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedIncrement,
			"num_conns_tagged_anomaly":                          ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedAnomaly,
			"num_conns_tagged_other_feature":                    ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedOtherFeature,
			"num_conns_tagged_increment_fail":                   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedIncrementFail,
			"num_conns_tagged_anomaly_fail":                     ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedAnomalyFail,
			"num_conns_tagged_other_feature_fail":               ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedOtherFeatureFail,
			"num_conns_tagged_increment_maxed":                  ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedIncrementMaxed,
			"num_conns_tagged_anomaly_maxed":                    ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedAnomalyMaxed,
			"num_conns_tagged_other_feature_maxed":              ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsTaggedOtherFeatureMaxed,
			"num_conns_untagged":                                ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumConnsUntagged,
			"pktcapture_triggered_by_increment":                 ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureTriggeredByIncrement,
			"pktcapture_triggered_by_anomaly":                   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureTriggeredByAnomaly,
			"pktcapture_triggered_by_other_feature":             ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcaptureTriggeredByOtherFeature,
			"num_of_anomalies_detected":                         ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumOfAnomaliesDetected,
			"num_of_anomalies_cleared":                          ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumOfAnomaliesCleared,
			"num_pcaps_created":                                 ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumPcapsCreated,
			"num_tmp_pcaps_created":                             ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumTmpPcapsCreated,
			"num_pcaps_create_failed":                           ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumPcapsCreateFailed,
			"pktcap_oom":                                        ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktcapOom,
			"failed_disk_full":                                  ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.FailedDiskFull,
			"conn_ext_failed":                                   ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.ConnExtFailed,
			"skip_as_conn_already_recapture":                    ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.SkipAsConnAlreadyRecapture,
			"skip_capture_as_conn_created_before_smp":           ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.SkipCaptureAsConnCreatedBeforeSmp,
			"failed_as_return_completed_set":                    ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.FailedAsReturnCompletedSet,
			"non_pkt_path":                                      ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NonPktPath,
			"pkt_already_captured":                              ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.PktAlreadyCaptured,
			"wrong_ctr_incremented":                             ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.WrongCtrIncremented,
			"auto_pcap_file_merged":                             ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.AutoPcapFileMerged,
			"auto_pcap_file_merged_failed":                      ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.AutoPcapFileMergedFailed,
			"num_dynamic_capture_config_created":                ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumDynamicCaptureConfigCreated,
			"num_dynamic_capture_config_delete_q":               ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumDynamicCaptureConfigDeleteQ,
			"num_dynamic_capture_config_deleted":                ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumDynamicCaptureConfigDeleted,
			"num_global_counters_registered":                    ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumGlobalCountersRegistered,
			"num_global_counters_deregistered":                  ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumGlobalCountersDeregistered,
			"num_per_object_counters_registered":                ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumPerObjectCountersRegistered,
			"num_per_object_counters_deregistered":              ret.DtVisibilityPacketCaptureCaptureConfigStats.Stats.NumPerObjectCountersDeregistered,
		},
	}
}

func getObjectVisibilityPacketCaptureCaptureConfigStatsStats(d []interface{}) edpt.VisibilityPacketCaptureCaptureConfigStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureCaptureConfigStatsStats
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

func dataToEndpointVisibilityPacketCaptureCaptureConfigStats(d *schema.ResourceData) edpt.VisibilityPacketCaptureCaptureConfigStats {
	var ret edpt.VisibilityPacketCaptureCaptureConfigStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectVisibilityPacketCaptureCaptureConfigStatsStats(d.Get("stats").([]interface{}))
	return ret
}
