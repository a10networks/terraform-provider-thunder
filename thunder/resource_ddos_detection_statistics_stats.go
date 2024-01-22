package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionStatisticsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_detection_statistics_stats`: Statistics for the object statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDetectionStatisticsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sflow_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Packets Received",
						},
						"not_supported_sflow_ver": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Packets Version Not Supported",
						},
						"netflow_pkts_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Packets Received",
						},
						"not_supproted_nflow_ver": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Packets Version Not Supported",
						},
						"agent_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Detection Agent Not Found",
						},
						"pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Flow Packets Dropped",
						},
						"report_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Report Allocate Failure",
						},
						"report_enqueue_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Report Enqueue Failure",
						},
						"sample_procssed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed",
						},
						"ip_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Received",
						},
						"ipv6_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Received",
						},
						"frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Received",
						},
						"dst_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Hit",
						},
						"dst_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Missed",
						},
						"dst_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Learned",
						},
						"dst_age": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Aged",
						},
						"dst_service_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Service Entry Hit",
						},
						"dst_service_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Service Entry Missed",
						},
						"dst_service_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Service Entry Learned",
						},
						"dst_service_age": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Service Entry Aged",
						},
						"src_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Hit",
						},
						"src_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Missed",
						},
						"src_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Learned",
						},
						"src_age": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Aged",
						},
						"entry_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Allocate Failure",
						},
						"geo_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Geolocation Entry Learned",
						},
						"geo_age": {
							Type: schema.TypeInt, Optional: true, Description: "Geolocation Entry Aged",
						},
						"unmatch_entry_port_zero": {
							Type: schema.TypeInt, Optional: true, Description: "Unmatched Entry Port-zero Packet",
						},
						"object_alloc_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Object Allocate Failure Out of Memory",
						},
						"invalid_event": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Event in Notification",
						},
						"rtbh_start_sent": {
							Type: schema.TypeInt, Optional: true, Description: "RTBH Start Notification Sent",
						},
						"rtbh_stop_sent": {
							Type: schema.TypeInt, Optional: true, Description: "RTBH Stop Notification Sent",
						},
						"rtbh_start_fail": {
							Type: schema.TypeInt, Optional: true, Description: "RTBH Start Notification Sent Failed",
						},
						"rtbh_stop_fail": {
							Type: schema.TypeInt, Optional: true, Description: "RTBH Stop Notification Sent Failed",
						},
						"invalid_proto": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Proto in Notification",
						},
						"dst_ip_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP Entry Learned",
						},
						"dst_ip_age": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP Entry Aged",
						},
						"n_subnet_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry of Network-object learned",
						},
						"n_subnet_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry of Network-object Aged",
						},
						"n_ip_learned": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry of Network-object Learned",
						},
						"n_ip_aged": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry of Network-object Aged",
						},
						"n_service_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry of Network-object Learned",
						},
						"n_service_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry of Network-object Aged",
						},
						"network_match_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Network-object Match Missed",
						},
						"session_match_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Session Match Missed",
						},
						"session_allocate_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Session Allocate Failed",
						},
						"session_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Session Learned",
						},
						"session_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Session Aged",
						},
						"src_port_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Src Port Entry Hit",
						},
						"src_port_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Src Port Entry Missed",
						},
						"src_port_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Src Port Entry Learned",
						},
						"src_port_age": {
							Type: schema.TypeInt, Optional: true, Description: "Src Port Entry Aged",
						},
						"n_service_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry of Network-object Not Found",
						},
						"n_subnet_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Subnet Entry of Network-object Create Failed",
						},
						"n_ip_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IP Entry of Network-object Create Failed",
						},
						"n_service_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Service Entry of Network-object Create Failed",
						},
						"db_unexpected_error": {
							Type: schema.TypeInt, Optional: true, Description: "Database Unexpected Error",
						},
						"db_oper_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Operation Failure",
						},
						"db_open_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Open Failure",
						},
						"db_n_subnet_table_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Subnet Entry Table Creation Failure",
						},
						"db_n_ip_table_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network IP Entry Table Creation Failure",
						},
						"db_n_svc_table_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Service Entry Table Creation Failure",
						},
						"db_n_subnet_save_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Subnet Entry Saving Attempt",
						},
						"db_n_subnet_save_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Subnet Entry Saving Failure",
						},
						"db_n_subnet_restore_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Subnet Entry Restoring Attempt",
						},
						"db_n_ip_save_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network IP Entry Saving Attempt",
						},
						"db_n_ip_save_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network IP Entry Saving Failure",
						},
						"db_n_ip_restore_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network IP Entry Restoring Attempt",
						},
						"db_n_svc_save_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Service Entry Saving Attempt",
						},
						"db_n_svc_save_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Service Entry Saving Failure",
						},
						"db_n_svc_restore_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Service Entry Restoring Attempt",
						},
						"db_n_static_subnet_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Static Subnet Entry Not Found",
						},
						"db_n_parent_entry_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Parent Entry Not Found",
						},
						"db_worker_enq_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Worker Enqueue Failure",
						},
						"db_n_subnet_table_purge_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Subnet Entry Table Purge Entries Failure",
						},
						"db_n_ip_table_purge_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network IP Entry Table Purge Entries Failure",
						},
						"db_n_svc_table_purge_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Database Network Service Entry Table Purge Entries Failure",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDetectionStatisticsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionStatisticsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionStatisticsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDetectionStatisticsStatsStats := setObjectDdosDetectionStatisticsStatsStats(res)
		d.Set("stats", DdosDetectionStatisticsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDetectionStatisticsStatsStats(ret edpt.DataDdosDetectionStatisticsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sflow_packets_received":           ret.DtDdosDetectionStatisticsStats.Stats.SflowPacketsReceived,
			"not_supported_sflow_ver":          ret.DtDdosDetectionStatisticsStats.Stats.NotSupportedSflowVer,
			"netflow_pkts_received":            ret.DtDdosDetectionStatisticsStats.Stats.NetflowPktsReceived,
			"not_supproted_nflow_ver":          ret.DtDdosDetectionStatisticsStats.Stats.NotSupprotedNflowVer,
			"agent_not_found":                  ret.DtDdosDetectionStatisticsStats.Stats.AgentNotFound,
			"pkt_drop":                         ret.DtDdosDetectionStatisticsStats.Stats.PktDrop,
			"report_alloc_fail":                ret.DtDdosDetectionStatisticsStats.Stats.ReportAllocFail,
			"report_enqueue_fail":              ret.DtDdosDetectionStatisticsStats.Stats.ReportEnqueueFail,
			"sample_procssed":                  ret.DtDdosDetectionStatisticsStats.Stats.SampleProcssed,
			"ip_rcvd":                          ret.DtDdosDetectionStatisticsStats.Stats.Ip_rcvd,
			"ipv6_rcvd":                        ret.DtDdosDetectionStatisticsStats.Stats.Ipv6_rcvd,
			"frag_rcvd":                        ret.DtDdosDetectionStatisticsStats.Stats.Frag_rcvd,
			"dst_hit":                          ret.DtDdosDetectionStatisticsStats.Stats.DstHit,
			"dst_miss":                         ret.DtDdosDetectionStatisticsStats.Stats.DstMiss,
			"dst_learn":                        ret.DtDdosDetectionStatisticsStats.Stats.DstLearn,
			"dst_age":                          ret.DtDdosDetectionStatisticsStats.Stats.DstAge,
			"dst_service_hit":                  ret.DtDdosDetectionStatisticsStats.Stats.DstServiceHit,
			"dst_service_miss":                 ret.DtDdosDetectionStatisticsStats.Stats.DstServiceMiss,
			"dst_service_learn":                ret.DtDdosDetectionStatisticsStats.Stats.DstServiceLearn,
			"dst_service_age":                  ret.DtDdosDetectionStatisticsStats.Stats.DstServiceAge,
			"src_hit":                          ret.DtDdosDetectionStatisticsStats.Stats.SrcHit,
			"src_miss":                         ret.DtDdosDetectionStatisticsStats.Stats.SrcMiss,
			"src_learn":                        ret.DtDdosDetectionStatisticsStats.Stats.SrcLearn,
			"src_age":                          ret.DtDdosDetectionStatisticsStats.Stats.SrcAge,
			"entry_alloc_fail":                 ret.DtDdosDetectionStatisticsStats.Stats.EntryAllocFail,
			"geo_learn":                        ret.DtDdosDetectionStatisticsStats.Stats.GeoLearn,
			"geo_age":                          ret.DtDdosDetectionStatisticsStats.Stats.GeoAge,
			"unmatch_entry_port_zero":          ret.DtDdosDetectionStatisticsStats.Stats.UnmatchEntryPortZero,
			"object_alloc_oom":                 ret.DtDdosDetectionStatisticsStats.Stats.ObjectAllocOom,
			"invalid_event":                    ret.DtDdosDetectionStatisticsStats.Stats.InvalidEvent,
			"rtbh_start_sent":                  ret.DtDdosDetectionStatisticsStats.Stats.RtbhStartSent,
			"rtbh_stop_sent":                   ret.DtDdosDetectionStatisticsStats.Stats.RtbhStopSent,
			"rtbh_start_fail":                  ret.DtDdosDetectionStatisticsStats.Stats.RtbhStartFail,
			"rtbh_stop_fail":                   ret.DtDdosDetectionStatisticsStats.Stats.RtbhStopFail,
			"invalid_proto":                    ret.DtDdosDetectionStatisticsStats.Stats.InvalidProto,
			"dst_ip_learn":                     ret.DtDdosDetectionStatisticsStats.Stats.DstIpLearn,
			"dst_ip_age":                       ret.DtDdosDetectionStatisticsStats.Stats.DstIpAge,
			"n_subnet_learned":                 ret.DtDdosDetectionStatisticsStats.Stats.NSubnetLearned,
			"n_subnet_aged":                    ret.DtDdosDetectionStatisticsStats.Stats.NSubnetAged,
			"n_ip_learned":                     ret.DtDdosDetectionStatisticsStats.Stats.NIpLearned,
			"n_ip_aged":                        ret.DtDdosDetectionStatisticsStats.Stats.NIpAged,
			"n_service_learned":                ret.DtDdosDetectionStatisticsStats.Stats.NServiceLearned,
			"n_service_aged":                   ret.DtDdosDetectionStatisticsStats.Stats.NServiceAged,
			"network_match_miss":               ret.DtDdosDetectionStatisticsStats.Stats.NetworkMatchMiss,
			"session_match_miss":               ret.DtDdosDetectionStatisticsStats.Stats.SessionMatchMiss,
			"session_allocate_fail":            ret.DtDdosDetectionStatisticsStats.Stats.SessionAllocateFail,
			"session_learned":                  ret.DtDdosDetectionStatisticsStats.Stats.SessionLearned,
			"session_aged":                     ret.DtDdosDetectionStatisticsStats.Stats.SessionAged,
			"src_port_hit":                     ret.DtDdosDetectionStatisticsStats.Stats.SrcPortHit,
			"src_port_miss":                    ret.DtDdosDetectionStatisticsStats.Stats.SrcPortMiss,
			"src_port_learn":                   ret.DtDdosDetectionStatisticsStats.Stats.SrcPortLearn,
			"src_port_age":                     ret.DtDdosDetectionStatisticsStats.Stats.SrcPortAge,
			"n_service_not_found":              ret.DtDdosDetectionStatisticsStats.Stats.NServiceNotFound,
			"n_subnet_create_fail":             ret.DtDdosDetectionStatisticsStats.Stats.NSubnetCreateFail,
			"n_ip_create_fail":                 ret.DtDdosDetectionStatisticsStats.Stats.NIpCreateFail,
			"n_service_create_fail":            ret.DtDdosDetectionStatisticsStats.Stats.NServiceCreateFail,
			"db_unexpected_error":              ret.DtDdosDetectionStatisticsStats.Stats.DbUnexpectedError,
			"db_oper_failure":                  ret.DtDdosDetectionStatisticsStats.Stats.DbOperFailure,
			"db_open_failure":                  ret.DtDdosDetectionStatisticsStats.Stats.DbOpenFailure,
			"db_n_subnet_table_create_failure": ret.DtDdosDetectionStatisticsStats.Stats.DbNSubnetTableCreateFailure,
			"db_n_ip_table_create_failure":     ret.DtDdosDetectionStatisticsStats.Stats.DbNIpTableCreateFailure,
			"db_n_svc_table_create_failure":    ret.DtDdosDetectionStatisticsStats.Stats.DbNSvcTableCreateFailure,
			"db_n_subnet_save_attempt":         ret.DtDdosDetectionStatisticsStats.Stats.DbNSubnetSaveAttempt,
			"db_n_subnet_save_failure":         ret.DtDdosDetectionStatisticsStats.Stats.DbNSubnetSaveFailure,
			"db_n_subnet_restore_attempt":      ret.DtDdosDetectionStatisticsStats.Stats.DbNSubnetRestoreAttempt,
			"db_n_ip_save_attempt":             ret.DtDdosDetectionStatisticsStats.Stats.DbNIpSaveAttempt,
			"db_n_ip_save_failure":             ret.DtDdosDetectionStatisticsStats.Stats.DbNIpSaveFailure,
			"db_n_ip_restore_attempt":          ret.DtDdosDetectionStatisticsStats.Stats.DbNIpRestoreAttempt,
			"db_n_svc_save_attempt":            ret.DtDdosDetectionStatisticsStats.Stats.DbNSvcSaveAttempt,
			"db_n_svc_save_failure":            ret.DtDdosDetectionStatisticsStats.Stats.DbNSvcSaveFailure,
			"db_n_svc_restore_attempt":         ret.DtDdosDetectionStatisticsStats.Stats.DbNSvcRestoreAttempt,
			"db_n_static_subnet_not_found":     ret.DtDdosDetectionStatisticsStats.Stats.DbNStaticSubnetNotFound,
			"db_n_parent_entry_not_found":      ret.DtDdosDetectionStatisticsStats.Stats.DbNParentEntryNotFound,
			"db_worker_enq_failure":            ret.DtDdosDetectionStatisticsStats.Stats.DbWorkerEnqFailure,
			"db_n_subnet_table_purge_failure":  ret.DtDdosDetectionStatisticsStats.Stats.DbNSubnetTablePurgeFailure,
			"db_n_ip_table_purge_failure":      ret.DtDdosDetectionStatisticsStats.Stats.DbNIpTablePurgeFailure,
			"db_n_svc_table_purge_failure":     ret.DtDdosDetectionStatisticsStats.Stats.DbNSvcTablePurgeFailure,
		},
	}
}

func getObjectDdosDetectionStatisticsStatsStats(d []interface{}) edpt.DdosDetectionStatisticsStatsStats {

	count1 := len(d)
	var ret edpt.DdosDetectionStatisticsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPacketsReceived = in["sflow_packets_received"].(int)
		ret.NotSupportedSflowVer = in["not_supported_sflow_ver"].(int)
		ret.NetflowPktsReceived = in["netflow_pkts_received"].(int)
		ret.NotSupprotedNflowVer = in["not_supproted_nflow_ver"].(int)
		ret.AgentNotFound = in["agent_not_found"].(int)
		ret.PktDrop = in["pkt_drop"].(int)
		ret.ReportAllocFail = in["report_alloc_fail"].(int)
		ret.ReportEnqueueFail = in["report_enqueue_fail"].(int)
		ret.SampleProcssed = in["sample_procssed"].(int)
		ret.Ip_rcvd = in["ip_rcvd"].(int)
		ret.Ipv6_rcvd = in["ipv6_rcvd"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.DstHit = in["dst_hit"].(int)
		ret.DstMiss = in["dst_miss"].(int)
		ret.DstLearn = in["dst_learn"].(int)
		ret.DstAge = in["dst_age"].(int)
		ret.DstServiceHit = in["dst_service_hit"].(int)
		ret.DstServiceMiss = in["dst_service_miss"].(int)
		ret.DstServiceLearn = in["dst_service_learn"].(int)
		ret.DstServiceAge = in["dst_service_age"].(int)
		ret.SrcHit = in["src_hit"].(int)
		ret.SrcMiss = in["src_miss"].(int)
		ret.SrcLearn = in["src_learn"].(int)
		ret.SrcAge = in["src_age"].(int)
		ret.EntryAllocFail = in["entry_alloc_fail"].(int)
		ret.GeoLearn = in["geo_learn"].(int)
		ret.GeoAge = in["geo_age"].(int)
		ret.UnmatchEntryPortZero = in["unmatch_entry_port_zero"].(int)
		ret.ObjectAllocOom = in["object_alloc_oom"].(int)
		ret.InvalidEvent = in["invalid_event"].(int)
		ret.RtbhStartSent = in["rtbh_start_sent"].(int)
		ret.RtbhStopSent = in["rtbh_stop_sent"].(int)
		ret.RtbhStartFail = in["rtbh_start_fail"].(int)
		ret.RtbhStopFail = in["rtbh_stop_fail"].(int)
		ret.InvalidProto = in["invalid_proto"].(int)
		ret.DstIpLearn = in["dst_ip_learn"].(int)
		ret.DstIpAge = in["dst_ip_age"].(int)
		ret.NSubnetLearned = in["n_subnet_learned"].(int)
		ret.NSubnetAged = in["n_subnet_aged"].(int)
		ret.NIpLearned = in["n_ip_learned"].(int)
		ret.NIpAged = in["n_ip_aged"].(int)
		ret.NServiceLearned = in["n_service_learned"].(int)
		ret.NServiceAged = in["n_service_aged"].(int)
		ret.NetworkMatchMiss = in["network_match_miss"].(int)
		ret.SessionMatchMiss = in["session_match_miss"].(int)
		ret.SessionAllocateFail = in["session_allocate_fail"].(int)
		ret.SessionLearned = in["session_learned"].(int)
		ret.SessionAged = in["session_aged"].(int)
		ret.SrcPortHit = in["src_port_hit"].(int)
		ret.SrcPortMiss = in["src_port_miss"].(int)
		ret.SrcPortLearn = in["src_port_learn"].(int)
		ret.SrcPortAge = in["src_port_age"].(int)
		ret.NServiceNotFound = in["n_service_not_found"].(int)
		ret.NSubnetCreateFail = in["n_subnet_create_fail"].(int)
		ret.NIpCreateFail = in["n_ip_create_fail"].(int)
		ret.NServiceCreateFail = in["n_service_create_fail"].(int)
		ret.DbUnexpectedError = in["db_unexpected_error"].(int)
		ret.DbOperFailure = in["db_oper_failure"].(int)
		ret.DbOpenFailure = in["db_open_failure"].(int)
		ret.DbNSubnetTableCreateFailure = in["db_n_subnet_table_create_failure"].(int)
		ret.DbNIpTableCreateFailure = in["db_n_ip_table_create_failure"].(int)
		ret.DbNSvcTableCreateFailure = in["db_n_svc_table_create_failure"].(int)
		ret.DbNSubnetSaveAttempt = in["db_n_subnet_save_attempt"].(int)
		ret.DbNSubnetSaveFailure = in["db_n_subnet_save_failure"].(int)
		ret.DbNSubnetRestoreAttempt = in["db_n_subnet_restore_attempt"].(int)
		ret.DbNIpSaveAttempt = in["db_n_ip_save_attempt"].(int)
		ret.DbNIpSaveFailure = in["db_n_ip_save_failure"].(int)
		ret.DbNIpRestoreAttempt = in["db_n_ip_restore_attempt"].(int)
		ret.DbNSvcSaveAttempt = in["db_n_svc_save_attempt"].(int)
		ret.DbNSvcSaveFailure = in["db_n_svc_save_failure"].(int)
		ret.DbNSvcRestoreAttempt = in["db_n_svc_restore_attempt"].(int)
		ret.DbNStaticSubnetNotFound = in["db_n_static_subnet_not_found"].(int)
		ret.DbNParentEntryNotFound = in["db_n_parent_entry_not_found"].(int)
		ret.DbWorkerEnqFailure = in["db_worker_enq_failure"].(int)
		ret.DbNSubnetTablePurgeFailure = in["db_n_subnet_table_purge_failure"].(int)
		ret.DbNIpTablePurgeFailure = in["db_n_ip_table_purge_failure"].(int)
		ret.DbNSvcTablePurgeFailure = in["db_n_svc_table_purge_failure"].(int)
	}
	return ret
}

func dataToEndpointDdosDetectionStatisticsStats(d *schema.ResourceData) edpt.DdosDetectionStatisticsStats {
	var ret edpt.DdosDetectionStatisticsStats

	ret.Stats = getObjectDdosDetectionStatisticsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
